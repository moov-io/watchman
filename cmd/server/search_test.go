// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/moov-io/ofac"

	"github.com/gorilla/mux"
)

func TestSearch__addressSearchRequest(t *testing.T) {
	u, _ := url.Parse("https://moov.io/search?address=add&city=new+york&state=ny&providence=prov&zip=44433&country=usa")
	req := readAddressSearchRequest(u)
	if req.Address != "add" {
		t.Errorf("req.Address=%s", req.Address)
	}
	if req.City != "new york" {
		t.Errorf("req.City=%s", req.City)
	}
	if req.State != "ny" {
		t.Errorf("req.State=%s", req.State)
	}
	if req.Providence != "prov" {
		t.Errorf("req.Providence=%s", req.Providence)
	}
	if req.Zip != "44433" {
		t.Errorf("req.Zip=%s", req.Zip)
	}
	if req.Country != "usa" {
		t.Errorf("req.Country=%s", req.Country)
	}
	if req.empty() {
		t.Error("req is not empty")
	}

	req = addressSearchRequest{}
	if !req.empty() {
		t.Error("req is empty now")
	}
	req.Address = "1600 1st St"
	if req.empty() {
		t.Error("req is not empty now")
	}
}

func TestSearch__Address(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?address=Ibex+House", nil)
	req.Header.Set("x-user-id", "test")

	reader := &ofac.Reader{
		AddressArray: []ofac.Address{
			{ // Real OFAC entry -- 173,129,"Ibex House, The Minories","London EC3N 1DY","United Kingdom",-0-
				EntityID:                    "173",
				AddressID:                   "129",
				Address:                     "Ibex House, The Minories",
				CityStateProvincePostalCode: "London EC3N 1DY",
				Country:                     "United Kingdom",
			},
		},
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, reader)
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
		t.Errorf("got %#v", addresses[0])
	}
}

func TestSearch__Name(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?name=ZAWAHIRI", nil)
	req.Header.Set("x-user-id", "test")

	reader := &ofac.Reader{
		SDNArray: []ofac.SDN{
			{ // Real OFAC entry
				EntityID: "2676",
				SDNName:  "AL ZAWAHIRI, Dr. Ayman",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Operational and Military Leader of JIHAD GROUP",
				Remarks:  "DOB 19 Jun 1951; POB Giza, Egypt; Passport 1084010 (Egypt); alt. Passport 19820215; Operational and Military Leader of JIHAD GROUP.",
			},
		},
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, reader)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var names []*ofac.SDN
	if err := json.NewDecoder(w.Body).Decode(&names); err != nil {
		t.Fatal(err)
	}
	if len(names) != 1 {
		t.Fatalf("got %#v", names)
	}
	if names[0].EntityID != "2676" {
		t.Errorf("got %#v", names[0])
	}
}

func TestSearch__AltName(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?altName=CIMEX", nil)
	req.Header.Set("x-user-id", "test")

	reader := &ofac.Reader{
		AlternateIdentityArray: []ofac.AlternateIdentity{
			{ // Real OFAC entry
				EntityID:      "559",
				AlternateID:   "481",
				AlternateType: "aka",
				AlternateName: "CIMEX",
			},
		},
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, reader)
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
	if alts[0].EntityID != "559" {
		t.Errorf("got %#v", alts[0])
	}
}
