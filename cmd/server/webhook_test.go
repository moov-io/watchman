// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	redirect = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://example.com")
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write([]byte("didn't redirect"))
	}
)

// TestWebhook_retry ensures the webhookHTTPClient never follows a redirect.
// This is done to prevent infinite (or costly) redirect cycles which can degrade performance.
func TestWebhook_retry(t *testing.T) {
	if testing.Short() {
		return
	}

	server := httptest.NewServer(http.HandlerFunc(redirect))
	defer server.Close()

	// normal client, ensure redirect is followed
	resp, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure we landed on example.com
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if !bytes.Contains(bs, []byte("iana.org")) {
		t.Errorf("resp.Body=%s", string(bs))
	}

	// Now ensure our webhookHTTPClient doesn't follow the redirect
	resp, err = webhookHTTPClient.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	bs, _ = ioutil.ReadAll(resp.Body)
	if !bytes.Contains(bs, []byte("didn't redirect")) {
		t.Errorf("resp.Body=%s", string(bs))
	}
}

func TestWebhook_validate(t *testing.T) {
	out, err := validateWebhook("")
	if err == nil {
		t.Error("expected error")
	}
	if out != "" {
		t.Errorf("got out=%q", out)
	}

	// happy path
	out, err = validateWebhook("https://ofac.example.com/callback")
	if err != nil {
		t.Error(err)
	}
	if out != "https://ofac.example.com/callback" {
		t.Errorf("got out=%q", out)
	}

	// HTTP endpoint
	out, err = validateWebhook("http://bad.example.com/callback")
	if err == nil {
		t.Error("expected error, but got none")
	}
	if out != "" {
		t.Errorf("out=%q", out)
	}
}
