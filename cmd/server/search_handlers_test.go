// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/moov-io/ofac"

	"github.com/gorilla/mux"
)

func TestSearch__Address(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?address=ibex+house&limit=1", nil)

	router := mux.NewRouter()
	addSearchRoutes(nil, router, addressSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	if v := w.Body.String(); !strings.Contains(v, `"match":0.89`) {
		t.Errorf("%#v", v)
	}

	var wrapper struct {
		Addresses []*ofac.Address `json:"addresses"`
	}
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if wrapper.Addresses[0].EntityID != "173" {
		t.Errorf("%#v", wrapper.Addresses[0])
	}
}

func TestSearch__Name(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?name=AL+ZAWAHIRI&limit=1", nil)

	router := mux.NewRouter()
	addSearchRoutes(nil, router, sdnSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	if v := w.Body.String(); !strings.Contains(v, `"match":0.91`) {
		t.Error(v)
	}

	var wrapper struct {
		SDNs []*ofac.SDN `json:"SDNs"`
	}
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if wrapper.SDNs[0].EntityID != "2676" {
		t.Errorf("%#v", wrapper.SDNs[0])
	}
}

func TestSearch__AltName(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?altName=sogo+KENKYUSHO&limit=1", nil)

	router := mux.NewRouter()
	addSearchRoutes(nil, router, altSearcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	if v := w.Body.String(); !strings.Contains(v, `"match":0.783`) {
		t.Error(v)
	}

	var wrapper struct {
		Alts []*ofac.AlternateIdentity `json:"altNames"`
	}
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if wrapper.Alts[0].EntityID != "4691" {
		t.Errorf("%#v", wrapper.Alts[0])
	}
}
