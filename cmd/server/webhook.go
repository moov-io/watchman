// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/moov-io/base/strx"
	"go4.org/syncutil"
)

var (
	// webhookGate is a goroutine-safe throttler designed to only allow N
	// goroutines to run at any given time.
	webhookGate *syncutil.Gate

	webhookHTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			IdleConnTimeout:     1 * time.Minute,
		},
		// never follow a redirect as it could lead to a DoS or us being redirected
		// to an unexpected location.
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)

func init() {
	maxWorkers, err := strconv.ParseInt(strx.Or(os.Getenv("WEBHOOK_MAX_WORKERS"), "10"), 10, 32)
	if err == nil {
		webhookGate = syncutil.NewGate(int(maxWorkers))
	}
}

// callWebhook will take `body` as JSON and make a POST request to the provided webhook url.
// Returned is the HTTP status code.
func callWebhook(body *bytes.Buffer, webhook string, authToken string) (int, error) {
	webhook, err := validateWebhook(webhook)
	if err != nil {
		return 0, err
	}

	// Setup HTTP request
	req, err := http.NewRequest("POST", webhook, body)
	if err != nil {
		return 0, fmt.Errorf("unknown error webhook: %v", err)
	}
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	// Guard HTTP calls in-flight
	if webhookGate != nil {
		webhookGate.Start()
		defer webhookGate.Done()
	}

	resp, err := webhookHTTPClient.Do(req)
	if resp == nil || err != nil {
		if resp == nil {
			return 0, fmt.Errorf("unable to call webhook: %v", err)
		}
		return resp.StatusCode, fmt.Errorf("HTTP problem with webhook %v", err)
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return resp.StatusCode, fmt.Errorf("callWebhook: bogus status code: %d", resp.StatusCode)
	}
	return resp.StatusCode, nil
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
