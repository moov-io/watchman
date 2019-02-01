// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var (
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

// validateWebhook performs some basic checks against the incoming webhook and
// returns a normalized value.
//
// - Must be an HTTPS url
// - Must be a valid URL
func validateWebhook(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme != "https" {
		return "", fmt.Errorf("%s is not an HTTPS url", u.String())
	}
	return u.String(), nil
}
