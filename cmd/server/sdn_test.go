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

var (
	sdnSearcher = &searcher{
		SDNs: precomputeSDNs([]*ofac.SDN{
			{
				EntityID: "3754",
				SDNName:  "ABU MARZOOK, Mousa Mohammed",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Political Leader in Amman, Jordan and Damascus, Syria for HAMAS",
				Remarks:  "DOB 09 Feb 1951; POB Gaza, Egypt; Passport 92/664 (Egypt); SSN 523-33-8386 (United States); Political Leader in Amman, Jordan and Damascus, Syria for HAMAS; a.k.a. 'ABU-'UMAR'.",
			},
		}),
		Addresses: precomputeAddresses([]*ofac.Address{
			{
				// 587,356,"Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku","Tokyo 103","Japan",-0-
				EntityID:                    "587",
				AddressID:                   "356",
				Address:                     "Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku",
				CityStateProvincePostalCode: "Tokyo 103",
				Country:                     "Japan",
			},
			{
				// 651,376,"Case Postale 236, 10 Bis Rue Du Vieux College 12-11","Geneva","Switzerland",-0-
				EntityID:                    "651",
				AddressID:                   "376",
				Address:                     "Case Postale 236, 10 Bis Rue Du Vieux College 12-11",
				CityStateProvincePostalCode: "Geneva",
				Country:                     "Switzerland",
			},
		}),
		Alts: precomputeAlts([]*ofac.AlternateIdentity{
			{
				// 815,641,"aka","GALAX INC.",-0-
				EntityID:      "815",
				AlternateID:   "641",
				AlternateType: "aka",
				AlternateName: "GALAX INC",
			},
			{
				// 4359,3565,"fka","INDUSTRIA AVICOLA PALMASECA S.A.",-0-
				EntityID:      "4359",
				AlternateID:   "3565",
				AlternateType: "fkaa",
				AlternateName: "INDUSTRIA AVICOLA PALMASECA S.A.",
			},
		}),
	}
)

func TestSDN__id(t *testing.T) {
	router := mux.NewRouter()

	// Happy path
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/random-sdn-id", nil)
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
}

func TestSDN__Address(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/587/addresses", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router, sdnSearcher)
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
	if addresses[0].EntityID != "587" {
		t.Errorf("got %s", addresses[0].EntityID)
	}
}

func TestSDN__AltNames(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/815/alts", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router, sdnSearcher)
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
	if alts[0].EntityID != "815" {
		t.Errorf("got %s", alts[0].EntityID)
	}
}

func TestSDN__Get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sdn/3754", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addSDNRoutes(nil, router, sdnSearcher)
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
