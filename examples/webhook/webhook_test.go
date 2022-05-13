// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/gorilla/mux"
)

var (
	exampleCustomer = Customer{
		ID: "320",
		SDN: &ofac.SDN{
			EntityID: "306",
			SDNName:  "BANCO NACIONAL DE CUBA",
			SDNType:  "individual",
			Programs: []string{"CUBA"},
			Title:    "",
			Remarks:  "a.k.a. 'BNC'.",
		},
		Addresses: []*ofac.Address{
			{
				EntityID:                    "306",
				AddressID:                   "201",
				Address:                     "Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku",
				CityStateProvincePostalCode: "Tokyo 103",
				Country:                     "Japan",
			},
		},
		Alts: []*ofac.AlternateIdentity{
			{
				EntityID:      "306",
				AlternateID:   "220",
				AlternateType: "aka",
				AlternateName: "NATIONAL BANK OF CUBA",
			},
		},
	}
)

func TestWebhookRoute(t *testing.T) {
	w := httptest.NewRecorder()
	logger := log.NewNopLogger()

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exampleCustomer); err != nil {
		t.Fatal(err)
	}

	router := mux.NewRouter()
	addWebhookRoute(logger, router)

	req := httptest.NewRequest("POST", "/ofac", &body)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestWebhookRoute__bad(t *testing.T) {
	logger := log.NewNopLogger()

	router := mux.NewRouter()
	addWebhookRoute(logger, router)

	// no body
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/ofac", nil)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}

	// malformed body (wrong JSON)
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/ofac", strings.NewReader(`{"thing": "other-object"}`))
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}
