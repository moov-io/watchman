// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/gorilla/mux"
)

func TestSDN__id(t *testing.T) {
	router := mux.NewRouter()

	// Happy path
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/sdn/random-sdn-id", nil)
	router.Methods("GET").Path("/sdn/{sdnId}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getSDNId(w, r); v != "random-sdn-id" {
			t.Errorf("got %s", v)
		}
		if w.Code != http.StatusOK {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Unhappy case
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/sdn", nil)
	router.Methods("GET").Path("/sdn/{sdnId}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getSDNId(w, req); v != "" {
			t.Errorf("didn't expect SDN, got %s", v)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Don't pass req through mux so mux.Vars finds nothing
	if v := getSDNId(w, req); v != "" {
		t.Errorf("expected empty, but got %q", v)
	}
}

func TestSDN__Address(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/sdn/173/addresses", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(log.NewNopLogger(), router, addressSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var addresses []*ofac.Address
	if err := json.NewDecoder(w.Body).Decode(&addresses); err != nil {
		t.Fatal(err)
	}
	if len(addresses) != 1 {
		t.Fatalf("got %#v", addresses)
	}
	if addresses[0].EntityID != "173" {
		t.Errorf("got %s", addresses[0].EntityID)
	}
}

func TestSDN__AltNames(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/sdn/4691/alts", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(log.NewNopLogger(), router, altSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var alts []*ofac.AlternateIdentity
	if err := json.NewDecoder(w.Body).Decode(&alts); err != nil {
		t.Fatal(err)
	}
	if len(alts) != 1 {
		t.Fatalf("got %#v", alts)
	}
	if alts[0].EntityID != "4691" {
		t.Errorf("got %s", alts[0].EntityID)
	}
}

func TestSDN__Get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/sdn/2681", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(log.NewNopLogger(), router, sdnSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var sdn *ofac.SDN
	if err := json.NewDecoder(w.Body).Decode(&sdn); err != nil {
		t.Fatal(err)
	}
	if sdn == nil || sdn.EntityID != "2681" {
		t.Errorf("got %#v", sdn)
	}
}
