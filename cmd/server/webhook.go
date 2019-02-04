// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go4.org/syncutil"
)

var (
	// webhookGate is a goroutine-safe throttler designed to only allow N
	// goroutines to run at any given time.
	webhookGate = syncutil.NewGate(10)

	webhookHTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			IdleConnTimeout:     1 * time.Minute,
		},
		// never follow a redirect
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)


// callWebhook will encode Customer as JSON and make a POST request to the provided webhook url.
func callWebhook(watchId string, customer *Customer, webhook string) error {
	webhook, err := validateWebhook(webhook)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(customer); err != nil {
		return fmt.Errorf("problem creating JSON for watch %s: %v", watchId, err)
	}
	req, err := http.NewRequest("POST", webhook, &body)
	if err != nil {
		return fmt.Errorf("unknown error with watch %s: %v", watchId, err)
	}

	// Guard HTTP calls in-flight
	webhookGate.Start()
	defer webhookGate.Done()

	resp, err := webhookHTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP problem with watch %s: %v", watchId, err)
	}
	resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return fmt.Errorf("callWebhook: bogus status code: %d", resp.StatusCode)
	}
	return nil
}

// validateWebhook performs some basic checks against the incoming webhook and
// returns a normalized value.
//
// - Must be an HTTPS url
// - Must be a valid URL
func validateWebhook(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("%s is not a valid URL: %v", raw, err)
	}
	if u.Scheme != "https" {
		return "", fmt.Errorf("%s is not an HTTPS url", u.String())
	}
	return u.String(), nil
}
