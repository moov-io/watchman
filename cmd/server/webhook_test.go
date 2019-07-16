// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base"
)

var (
	// customerWebhook reads a Customer in JSON from the incoming request and replies
	// with the Customer.ID
	customerWebhook = func(w http.ResponseWriter, r *http.Request) {
		var cust Customer
		if err := json.NewDecoder(r.Body).Decode(&cust); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			if cust.ID == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(cust.ID))
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

func TestWebhook_call(t *testing.T) {
	if testing.Short() {
		return
	}

	server := httptest.NewTLSServer(http.HandlerFunc(customerWebhook))
	defer server.Close()

	// override to add test TLS certificate
	if tr, ok := webhookHTTPClient.Transport.(*http.Transport); ok {
		if ctr, ok := server.Client().Transport.(*http.Transport); ok {
			tr.TLSClientConfig.RootCAs = ctr.TLSClientConfig.RootCAs
		} else {
			t.Errorf("unknown server.Client().Transport type: %T", server.Client().Transport)
		}
	} else {
		t.Fatalf("%T %#v", webhookHTTPClient.Transport, webhookHTTPClient.Transport)
	}

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()

	// execute webhook with arbitrary Customer
	body, err := getCustomerBody(customerSearcher, "watchID", "306", 1.0, custRepo)
	if body == nil {
		t.Fatalf("nil body: %v", err)
	}
	if _, err := callWebhook(base.ID(), body, server.URL, "authToken"); err != nil {
		t.Fatal(err)
	}
}

func TestWebhook_record(t *testing.T) {
	db, err := createTestSqliteDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.close()

	repo := sqliteWebhookRepository{db.db}
	defer repo.close()

	if err := repo.recordWebhook(base.ID(), time.Now(), 200); err != nil {
		t.Fatal(err)
	}
}
