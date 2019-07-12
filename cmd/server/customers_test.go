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
	"time"

	"github.com/cardonator/ofac"
	"github.com/moov-io/base"

	"github.com/gorilla/mux"
)

var (
	// customerSearcher is a searcher instance with one individual entity contained. It's designed to be used
	// in tests that expect an individual SDN.
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

func createTestCustomerRepository(t *testing.T) *sqliteCustomerRepository {
	t.Helper()

	db, err := createTestSqliteDB()
	if err != nil {
		t.Fatal(err)
	}
	return &sqliteCustomerRepository{db.db}
}

func TestCustomers__id(t *testing.T) {
	router := mux.NewRouter()

	// Happy path
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/random-cust-id", nil)
	router.Methods("GET").Path("/customers/{customerID}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCustomerID(w, r); v != "random-cust-id" {
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
	router.Methods("GET").Path("/customers/{customerID}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCustomerID(w, req); v != "" {
			t.Errorf("didn't expect customerID, got %s", v)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Don't pass req through mux so mux.Vars finds nothing
	if v := getCustomerID(w, req); v != "" {
		t.Errorf("expected empty, but got %q", v)
	}
}

func TestCustomer_getById(t *testing.T) {
	repo := createTestCustomerRepository(t)
	defer repo.close()

	// make sure we only return SDNType == "individual"
	// We do this by proviing a searcher with non-individual results
	cust, err := getCustomerByID("21206", companySearcher, repo)
	if cust != nil {
		t.Fatalf("expected no Customer, but got %#v", cust)
	}
	if err != nil {
		t.Fatalf("expected no error, but got %#v", err)
	}
}

func TestCustomer_get(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/306", nil)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

func TestCustomer_EmptyHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/customers/foo", nil)

	customerRepo := createTestCustomerRepository(t)
	defer customerRepo.close()

	getCustomer(nil, customerSearcher, customerRepo)(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_addWatch(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": "foo"}`)
	req := httptest.NewRequest("POST", "/customers/foo/watch", body)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, nil, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_addWatchMissingAuthToken(t *testing.T) {
	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": ""}`)

	req := httptest.NewRequest("POST", "/customers/foo/watch", body)
	req.Header.Set("x-user-id", "test")

	w := httptest.NewRecorder()

	// Setup test HTTP server
	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_addNameWatch(t *testing.T) {
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": "foo"}`)
	req := httptest.NewRequest("POST", "/customers/watch?name=foo", body)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

func TestCustomer_updateUnsafe(t *testing.T) {
	w := httptest.NewRecorder()

	body := strings.NewReader(`{"status": "unsafe"}`)
	req := httptest.NewRequest("PUT", "/customers/foo", body)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_updateException(t *testing.T) {
	w := httptest.NewRecorder()

	body := strings.NewReader(`{"status": "exception"}`)
	req := httptest.NewRequest("PUT", "/customers/foo", body)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_updateUnknown(t *testing.T) {
	w := httptest.NewRecorder()

	body := strings.NewReader(`{"status": "unknown"}`) // has status, but not blocked or unblocked
	req := httptest.NewRequest("PUT", "/customers/foo", body)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomer_updateNoUserId(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PUT", "/customers/foo", nil)

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %d but got: %d", http.StatusBadRequest, w.Code)
	}
}

func TestCustomer_updateNoBody(t *testing.T) {
	w := httptest.NewRecorder()

	req := httptest.NewRequest("PUT", "/customers/foo", nil)
	req.Header.Set("x-user-id", "test")

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
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

	custRepo := createTestCustomerRepository(t)
	defer custRepo.close()
	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCustomerRoutes(nil, router, customerSearcher, custRepo, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCustomerRepository(t *testing.T) {
	repo := createTestCustomerRepository(t)
	defer repo.close()

	customerID, userID := base.ID(), base.ID()

	status, err := repo.getCustomerStatus(customerID)
	if err != nil {
		t.Fatal(err)
	}
	if status != nil {
		t.Fatal("should give nil CustomerStatus")
	}

	// block customer
	status = &CustomerStatus{UserID: userID, Status: CustomerUnsafe, CreatedAt: time.Now()}
	if err := repo.upsertCustomerStatus(customerID, status); err != nil {
		t.Errorf("addCustomerBlock: shouldn't error, but got %v", err)
	}
	status = nil

	// verify
	status, err = repo.getCustomerStatus(customerID)
	if err != nil {
		t.Error(err)
	}
	if status == nil {
		t.Errorf("empty CustomerStatus")
	}
	if status.UserID == "" || string(status.Status) == "" {
		t.Errorf("invalid CustomerStatus: %#v", status)
	}
	if status.Status != CustomerUnsafe {
		t.Errorf("status.Status=%v", status.Status)
	}

	// unblock
	status = &CustomerStatus{UserID: userID, Status: CustomerException, CreatedAt: time.Now()}
	if err := repo.upsertCustomerStatus(customerID, status); err != nil {
		t.Errorf("addCustomerBlock: shouldn't error, but got %v", err)
	}
	status = nil

	status, err = repo.getCustomerStatus(customerID)
	if err != nil {
		t.Error(err)
	}
	if status == nil {
		t.Errorf("empty CustomerStatus")
	}
	if status.UserID == "" || string(status.Status) == "" {
		t.Errorf("invalid CustomerStatus: %#v", status)
	}
	if status.Status != CustomerException {
		t.Errorf("status.Status=%v", status.Status)
	}

	// edgae case
	status, err = repo.getCustomerStatus("")
	if status != nil {
		t.Error("empty customerID shouldn return nil status")
	}
	if err == nil {
		t.Error("but an error should be returned")
	}
}
