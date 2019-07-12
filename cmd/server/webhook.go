// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"database/sql"
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
		// never follow a redirect as it could lead to a DoS or us being redirected
		// to an unexpected location.
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
)

// callWebhook will take `body` as JSON and make a POST request to the provided webhook url.
// Returned is the HTTP status code.
func callWebhook(watchID string, body *bytes.Buffer, webhook string, authToken string) (int, error) {
	webhook, err := validateWebhook(webhook)
	if err != nil {
		return 0, err
	}

	// Setup HTTP request
	req, err := http.NewRequest("POST", webhook, body)
	if err != nil {
		return 0, fmt.Errorf("unknown error with watch %s: %v", watchID, err)
	}
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}

	// Guard HTTP calls in-flight
	webhookGate.Start()
	defer webhookGate.Done()

	resp, err := webhookHTTPClient.Do(req)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("HTTP problem with watch %s: %v", watchID, err)
	}
	resp.Body.Close()
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

type webhookRepository interface {
	recordWebhook(watchID string, attemptedAt time.Time, status int) error
}

type sqliteWebhookRepository struct {
	db *sql.DB
}

func (r *sqliteWebhookRepository) close() error {
	return r.db.Close()
}

func (r *sqliteWebhookRepository) recordWebhook(watchID string, attemptedAt time.Time, status int) error {
	query := `insert into webhook_stats (watch_id, attempted_at, status) values (?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(watchID, attemptedAt, status)
	return err
}
