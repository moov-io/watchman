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

func TestSearch__extractSearchLimit(t *testing.T) {
	// Too high, fallback to hard max
	req := httptest.NewRequest("GET", "/?limit=1000", nil)
	if limit := extractSearchLimit(req); limit != hardResultsLimit {
		t.Errorf("got limit of %d", limit)
	}

	// No limit, use default
	req = httptest.NewRequest("GET", "/", nil)
	if limit := extractSearchLimit(req); limit != softResultsLimit {
		t.Errorf("got limit of %d", limit)
	}

	// Between soft and hard max
	req = httptest.NewRequest("GET", "/?limit=25", nil)
	if limit := extractSearchLimit(req); limit != 25 {
		t.Errorf("got limit of %d", limit)
	}

	// Lower than soft max
	req = httptest.NewRequest("GET", "/?limit=1", nil)
	if limit := extractSearchLimit(req); limit != 1 {
		t.Errorf("got limit of %d", limit)
	}
}

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
	req := httptest.NewRequest("GET", "/search?address=Ibex", nil)
	req.Header.Set("x-user-id", "test")

	searcher := &searcher{
		Addresses: precomputeAddresses([]*ofac.Address{
			{ // Real OFAC entry -- 173,129,"Ibex House, The Minories","London EC3N 1DY","United Kingdom",-0-
				EntityID:                    "173",
				AddressID:                   "129",
				Address:                     "Ibex House, The Minories",
				CityStateProvincePostalCode: "London EC3N 1DY",
				Country:                     "United Kingdom",
			},
			{ // 735,447,"Piarco Airport","Port au Prince","Haiti",-0-
				EntityID:                    "735",
				AddressID:                   "447",
				Address:                     "Piarco Airport",
				CityStateProvincePostalCode: "Port au Prince",
				Country:                     "Haiti",
			},
		}),
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var wrapper searchResponse
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.Addresses) != 1 {
		t.Fatalf("got %#v", wrapper.Addresses)
	}
	if wrapper.Addresses[0].EntityID != "173" {
		t.Errorf("got %#v", wrapper.Addresses[0])
	}

	// Search with more data in ?address=...
	w = httptest.NewRecorder()
	req.URL.Query().Set("address", "ibex+the")
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
	wrapper.Addresses = nil
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.Addresses) != 1 {
		t.Fatalf("got %#v", wrapper.Addresses)
	}
}

func TestSearch__Name(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?name=ZAWAHIRI", nil)
	req.Header.Set("x-user-id", "test")

	searcher := &searcher{
		SDNs: precomputeSDNs([]*ofac.SDN{
			{ // Real OFAC entry
				EntityID: "2676",
				SDNName:  "AL ZAWAHIRI, Dr. Ayman",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Operational and Military Leader of JIHAD GROUP",
				Remarks:  "DOB 19 Jun 1951; POB Giza, Egypt; Passport 1084010 (Egypt); alt. Passport 19820215; Operational and Military Leader of JIHAD GROUP.",
			},
			{
				EntityID: "2681",
				SDNName:  "HAWATMA, Nayif",
				SDNType:  "individual",
				Program:  "SDT",
				Title:    "Secretary General of DEMOCRATIC FRONT FOR THE LIBERATION OF PALESTINE - HAWATMEH FACTION",
				Remarks:  "DOB 1933; Secretary General of DEMOCRATIC FRONT FOR THE LIBERATION OF PALESTINE - HAWATMEH FACTION.",
			},
		}),
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var wrapper searchResponse
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.SDNs) != 1 {
		t.Fatalf("got %#v", wrapper.SDNs)
	}
	if wrapper.SDNs[0].EntityID != "2676" {
		t.Errorf("got %#v", wrapper.SDNs[0])
	}
}

func TestSearch__NameMultiple(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?name=al+ayman", nil)
	req.Header.Set("x-user-id", "test")

	searcher := &searcher{
		SDNs: precomputeSDNs([]*ofac.SDN{
			{ // Real OFAC entry
				EntityID: "2676",
				SDNName:  "AL ZAWAHIRI, Dr. Ayman",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Operational and Military Leader of JIHAD GROUP",
				Remarks:  "DOB 19 Jun 1951; POB Giza, Egypt; Passport 1084010 (Egypt); alt. Passport 19820215; Operational and Military Leader of JIHAD GROUP.",
			},
			{
				EntityID: "2681",
				SDNName:  "HAWATMA, Nayif",
				SDNType:  "individual",
				Program:  "SDT",
				Title:    "Secretary General of DEMOCRATIC FRONT FOR THE LIBERATION OF PALESTINE - HAWATMEH FACTION",
				Remarks:  "DOB 1933; Secretary General of DEMOCRATIC FRONT FOR THE LIBERATION OF PALESTINE - HAWATMEH FACTION.",
			},
		}),
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var wrapper searchResponse
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.SDNs) != 1 {
		t.Fatalf("got %#v", wrapper.SDNs)
	}
	if wrapper.SDNs[0].EntityID != "2676" {
		t.Errorf("got %#v", wrapper.SDNs[0])
	}
}

func TestSearch__AltName(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?altName=CIMEX", nil)
	req.Header.Set("x-user-id", "test")

	searcher := &searcher{
		Alts: precomputeAlts([]*ofac.AlternateIdentity{
			{ // Real OFAC entry
				EntityID:      "559",
				AlternateID:   "481",
				AlternateType: "aka",
				AlternateName: "CIMEX",
			},
			{
				EntityID:      "4691",
				AlternateID:   "3887",
				AlternateType: "aka",
				AlternateName: "A.I.C. SOGO KENKYUSHO",
			},
		}),
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var wrapper searchResponse
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.AltNames) != 1 {
		t.Fatalf("got %#v", wrapper.AltNames)
	}
	if wrapper.AltNames[0].EntityID != "559" {
		t.Errorf("got %#v", wrapper.AltNames[0])
	}
}

func TestSearch__AltNameMultiple(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/search?altName=SOGO+kenkyusho", nil)
	req.Header.Set("x-user-id", "test")

	searcher := &searcher{
		Alts: precomputeAlts([]*ofac.AlternateIdentity{
			{ // Real OFAC entry
				EntityID:      "559",
				AlternateID:   "481",
				AlternateType: "aka",
				AlternateName: "CIMEX",
			},
			{
				EntityID:      "4691",
				AlternateID:   "3887",
				AlternateType: "aka",
				AlternateName: "A.I.C. SOGO KENKYUSHO",
			},
		}),
	}

	router := mux.NewRouter()
	addSearchRoutes(nil, router, searcher)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var wrapper searchResponse
	if err := json.NewDecoder(w.Body).Decode(&wrapper); err != nil {
		t.Fatal(err)
	}
	if len(wrapper.AltNames) != 1 {
		t.Fatalf("got %#v", wrapper.AltNames)
	}
	if wrapper.AltNames[0].EntityID != "4691" {
		t.Errorf("got %#v", wrapper.AltNames[0])
	}
}
