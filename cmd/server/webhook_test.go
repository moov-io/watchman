// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base/log"
)

var (
	downloadWebhook = func(w http.ResponseWriter, r *http.Request) {
		var stats DownloadStats
		if err := json.NewDecoder(r.Body).Decode(&stats); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			if stats.SDNs != 101 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	}

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
	bs, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !bytes.Contains(bs, []byte("iana.org")) {
		t.Errorf("resp.Body=%s", string(bs))
	}

	// Now ensure our webhookHTTPClient doesn't follow the redirect
	resp, err = webhookHTTPClient.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bs, _ = io.ReadAll(resp.Body)
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

func TestWebhook_call(t *testing.T) {
	if testing.Short() {
		return
	}

	server := httptest.NewTLSServer(http.HandlerFunc(downloadWebhook))
	defer server.Close()

	// override to add test TLS certificate
	if tr, ok := webhookHTTPClient.Transport.(*http.Transport); ok {
		if ctr, ok := server.Client().Transport.(*http.Transport); ok {
			tr.TLSClientConfig = new(tls.Config)
			tr.TLSClientConfig.RootCAs = ctr.TLSClientConfig.RootCAs
		} else {
			t.Errorf("unknown server.Client().Transport type: %T", server.Client().Transport)
		}
	} else {
		t.Fatalf("%T %#v", webhookHTTPClient.Transport, webhookHTTPClient.Transport)
	}

	stats := &DownloadStats{
		SDNs:        101,
		RefreshedAt: time.Now().In(time.UTC),
	}

	t.Setenv("DOWNLOAD_WEBHOOK_URL", server.URL)
	t.Setenv("DOWNLOAD_WEBHOOK_AUTH_TOKEN", "authToken")

	logger := log.NewTestLogger()
	err := callDownloadWebook(logger, stats)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWebhook__CallErr(t *testing.T) {
	var body bytes.Buffer
	body.WriteString(`{"foo": "bar"}`)

	status, err := callWebhook(&body, "https://localhost/12345", "12345")
	if err == nil {
		t.Fatal(err)
	}
	if status != 0 {
		t.Errorf("bogus HTTP status: %d", status)
	}
}
