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

	"github.com/gorilla/mux"
)

func TestCustomer_get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/foo", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
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
	req := httptest.NewRequest("POST", "/customers/foo/watch", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
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

func TestCustomer_addNameWatch(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/customers/watch?name=foo", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
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

func TestCustomer_update(t *testing.T) {
	w := httptest.NewRecorder()

	body := strings.NewReader(`{"status": "Blocked"}`)
	req := httptest.NewRequest("PUT", "/customers/foo", body)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
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

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %d but got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestCustomer_removeWatch(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/customers/foo/watch", nil)
	req.Header.Set("x-user-id", "test")

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
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

	router := mux.NewRouter()
	addCustomerRoutes(nil, router)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}
