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

var (
	customerSearcher = &searcher{
		SDNs: precomputeSDNs([]*ofac.SDN{
			{
				EntityID: "306",
				SDNName:  "BANCO NACIONAL DE CUBA",
				SDNType:  "individual",
				Program:  "CUBA",
				Title:    "",
				Remarks:  "a.k.a. 'BNC'.",
			},
		}),
		Addresses: precomputeAddresses([]*ofac.Address{
			{
				EntityID:                    "306",
				AddressID:                   "201",
				Address:                     "Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku",
				CityStateProvincePostalCode: "Tokyo 103",
				Country:                     "Japan",
			},
		}),
		Alts: precomputeAlts([]*ofac.AlternateIdentity{
			{
				EntityID:      "306",
				AlternateID:   "220",
				AlternateType: "aka",
				AlternateName: "NATIONAL BANK OF CUBA",
			},
		}),
	}
)

func TestCustomers__id(t *testing.T) {
	router := mux.NewRouter()

	// Happy path
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/random-cust-id", nil)
	router.Methods("GET").Path("/customers/{customerId}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCustomerId(w, r); v != "random-cust-id" {
			t.Errorf("got %s", v)
		}
		if w.Code != http.StatusOK {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Unhappy case
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/customers", nil)
	router.Methods("GET").Path("/customers/{customerId}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCustomerId(w, req); v != "" {
			t.Errorf("didn't expect customerId, got %s", v)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Don't pass req through mux so mux.Vars finds nothing
	if v := getCustomerId(w, req); v != "" {
		t.Errorf("expected empty, but got %q", v)
	}
}

func TestCustomer_get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/306", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var cust Customer
	if err := json.NewDecoder(w.Body).Decode(&cust); err != nil {
		t.Fatal(err)
	}
	if cust.ID == "" {
		t.Fatalf("empty ofac.Customer: %#v", cust)
	}
	if cust.SDN == nil {
		t.Fatal("missing cust.SDN")
	}
	if len(cust.Addresses) != 1 {
		t.Errorf("cust.Addresses: %#v", cust.Addresses)
	}
	if len(cust.Alts) != 1 {
		t.Errorf("cust.Alts: %#v", cust.Alts)
	}
}

func TestCustomer_addWatch(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"webhook": "https://moov.io"}`)
	req := httptest.NewRequest("POST", "/customers/foo/watch", body)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var watch customerWatchResponse
	if err := json.NewDecoder(w.Body).Decode(&watch); err != nil {
		t.Fatal(err)
	}
	if watch.WatchID == "" {
		t.Error("empty watch.WatchID")
	}
}

func TestCustomer_addWatchNoBody(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/customers/foo/watch", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_addNameWatch(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"webhook": "https://moov.io"}`)
	req := httptest.NewRequest("POST", "/customers/watch?name=foo", body)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}

	var watch customerWatchResponse
	if err := json.NewDecoder(w.Body).Decode(&watch); err != nil {
		t.Fatal(err)
	}
	if watch.WatchID == "" {
		t.Error("empty watch.WatchID")
	}
}

func TestCustomer_addCustomerNameWatchNoBody(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/customers/watch?name=foo", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}

	// reset
	w = httptest.NewRecorder()
	if w.Code != http.StatusOK {
		t.Errorf("bad state reset: %d", w.Code)
	}

	req.URL.Query().Set("name", "")
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_update(t *testing.T) {
	w := httptest.NewRecorder()

	body := strings.NewReader(`{"status": "Blocked"}`)
	req := httptest.NewRequest("PUT", "/customers/foo", body)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_updateNoBody(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PUT", "/customers/foo", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %d but got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestCustomer_removeWatch(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/customers/foo/watch/watch-id", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_removeNameWatch(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/customers/watch/foo", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}
