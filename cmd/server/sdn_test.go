// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/ofac"

	"github.com/gorilla/mux"
)

func TestSDN__Address(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/foo/addresses", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var addresses []*ofac.Address
	if err := json.NewDecoder(w.Body).Decode(&addresses); err != nil {
		t.Fatal(err)
	}
	if len(addresses) != 2 {
		t.Fatalf("got %#v", addresses)
	}
	if addresses[0].EntityID != "587" || addresses[1].EntityID != "651" {
		t.Errorf("got %s and %s", addresses[0].EntityID, addresses[1].EntityID)
	}
}

func TestSDN__AltNames(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/foo/alts", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var alts []*ofac.AlternateIdentity
	if err := json.NewDecoder(w.Body).Decode(&alts); err != nil {
		t.Fatal(err)
	}
	if len(alts) != 2 {
		t.Fatalf("got %#v", alts)
	}
	if alts[0].EntityID != "815" || alts[1].EntityID != "4359" {
		t.Errorf("got %s and %s", alts[0].EntityID, alts[1].EntityID)
	}
}

func TestSDN__Get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/foo", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var sdn *ofac.SDN
	if err := json.NewDecoder(w.Body).Decode(&sdn); err != nil {
		t.Fatal(err)
	}
	if sdn == nil || sdn.EntityID != "3754" {
		t.Errorf("got %#v", sdn)
	}
}
