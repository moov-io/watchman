// Copyright 2022 The Moov Authors
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

	"github.com/moov-io/base"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/database"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/gorilla/mux"
)

var (
	// companySearcher is a searcher instance with one company entity contained. It's designed to be used
	// in tests that expect a non-individual SDN.
	companySearcher *searcher
)

func init() {
	companySearcher = newSearcher(log.NewNopLogger(), noLogPipeliner, 1)
	companySearcher.SDNs = precomputeSDNs([]*ofac.SDN{
		{
			EntityID: "21206",
			SDNName:  "AL-HISN",
			Programs: []string{"SYRIA"},
			Remarks:  "Linked To: MAKHLUF, Rami.",
		},
	}, nil, noLogPipeliner)
	companySearcher.Addresses = precomputeAddresses([]*ofac.Address{
		{
			EntityID:                    "21206",
			AddressID:                   "32272",
			Address:                     "Jurmana",
			CityStateProvincePostalCode: "Damascus",
			Country:                     "Syria",
		},
	})
	companySearcher.Alts = precomputeAlts([]*ofac.AlternateIdentity{
		{
			EntityID:      "21206",
			AlternateID:   "33627",
			AlternateType: "aka",
			AlternateName: "AL-HISN FIRM",
		},
	}, noLogPipeliner)
}

func createTestCompanyRepository(t *testing.T) *sqliteCompanyRepository {
	t.Helper()

	db := database.CreateTestSqliteDB(t)
	return &sqliteCompanyRepository{db.DB, log.NewNopLogger()}
}

func TestCompanies__id(t *testing.T) {
	router := mux.NewRouter()

	// Happy path
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/companies/random-company-id", nil)
	router.Methods("GET").Path("/companies/{companyID}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCompanyID(w, r); v != "random-company-id" {
			t.Errorf("got %s", v)
		}
		if w.Code != http.StatusOK {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Unhappy case
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/ofac/companies", nil)
	router.Methods("GET").Path("/companies/{companyID}").HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if v := getCompanyID(w, req); v != "" {
			t.Errorf("didn't expect companyID, got %s", v)
		}
		if w.Code != http.StatusBadRequest {
			t.Errorf("got status code %d", w.Code)
		}
	})
	router.ServeHTTP(w, req)

	// Don't pass req through mux so mux.Vars finds nothing
	if v := getCompanyID(w, req); v != "" {
		t.Errorf("expected empty, but got %q", v)
	}
}

func TestCompany_getById(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		// make sure we only return SDNType != "individual"
		// We do this by proviing a searcher with individual results
		company, err := getCompanyByID("306", customerSearcher, repo)
		if company != nil {
			t.Fatalf("expected no Company, but got %#v", company)
		}
		if err != nil {
			t.Fatalf("expected no error, but got %#v", err)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_get(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/ofac/companies/21206", nil)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}

		var company Company
		if err := json.NewDecoder(w.Body).Decode(&company); err != nil {
			t.Fatal(err)
		}
		if company.ID == "" {
			t.Fatalf("empty ofac.Company: %#v", company)
		}
		if company.SDN == nil {
			t.Fatal("missing company.SDN")
		}
		if len(company.Addresses) != 1 {
			t.Errorf("company.Addresses: %#v", company.Addresses)
		}
		if len(company.Alts) != 1 {
			t.Errorf("company.Alts: %#v", company.Alts)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_EmptyHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ofac/companies/foo", nil)

	companyRepo := createTestCompanyRepository(t)
	defer companyRepo.close()

	getCompany(nil, companySearcher, companyRepo)(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCompany_addWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": "foo"}`)
		req := httptest.NewRequest("POST", "/ofac/companies/foo/watch", body)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}

		var watch companyWatchResponse
		if err := json.NewDecoder(w.Body).Decode(&watch); err != nil {
			t.Fatal(err)
		}
		if watch.WatchID == "" {
			t.Error("empty watch.WatchID")
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_addWatchNoBody(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/ofac/companies/foo/watch", nil)
	req.Header.Set("x-user-id", "test")

	watchRepo := createTestWatchRepository(t)
	defer watchRepo.close()

	router := mux.NewRouter()
	addCompanyRoutes(log.NewNopLogger(), router, companySearcher, nil, watchRepo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus status code: %d", w.Code)
	}
}

func TestCompany_addWatchMissingAuthToken(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": ""}`)

		req := httptest.NewRequest("POST", "/ofac/companies/foo/watch", body)
		req.Header.Set("x-user-id", "test")

		w := httptest.NewRecorder()

		// Setup test HTTP server
		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusBadRequest {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_addNameWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		body := strings.NewReader(`{"webhook": "https://moov.io", "authToken": "foo"}`)
		req := httptest.NewRequest("POST", "/ofac/companies/watch?name=foo", body)
		req.Header.Set("x-user-id", "test")
		req.Header.Set("x-request-id", base.ID())

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}

		var watch companyWatchResponse
		if err := json.NewDecoder(w.Body).Decode(&watch); err != nil {
			t.Fatal(err)
		}
		if watch.WatchID == "" {
			t.Error("empty watch.WatchID")
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_addCompanyNameWatchNoBody(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/ofac/companies/watch?name=foo", nil)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
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

		req = httptest.NewRequest("POST", "/ofac/companies/watch", nil)
		req.Header.Set("x-user-id", "test")

		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusBadRequest {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_updateUnsafe(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()

		body := strings.NewReader(`{"status": "unsafe"}`)
		req := httptest.NewRequest("PUT", "/ofac/companies/foo", body)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_updateException(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()

		body := strings.NewReader(`{"status": "exception"}`)
		req := httptest.NewRequest("PUT", "/ofac/companies/foo", body)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_updateUnknown(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()

		body := strings.NewReader(`{"status": "unknown"}`) // has status, but not blocked or unblocked
		req := httptest.NewRequest("PUT", "/ofac/companies/foo", body)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusBadRequest {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_updateNoUserId(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", "/ofac/companies/foo", nil)

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected %d but got: %d", http.StatusBadRequest, w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_updateNoBody(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", "/ofac/companies/foo", nil)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected %d but got: %d", http.StatusBadRequest, w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_removeWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/ofac/companies/foo/watch/watch-id", nil)
		req.Header.Set("x-user-id", "test")

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompany_removeNameWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/ofac/companies/watch/foo", nil)
		req.Header.Set("x-user-id", "test")
		req.Header.Set("x-request-id", base.ID())

		watchRepo := createTestWatchRepository(t)
		defer watchRepo.close()

		router := mux.NewRouter()
		addCompanyRoutes(log.NewNopLogger(), router, companySearcher, repo, watchRepo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompanyRepository(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteCompanyRepository) {
		companyID, userID := base.ID(), base.ID()

		status, err := repo.getCompanyStatus(companyID)
		if err != nil {
			t.Fatal(err)
		}
		if status != nil {
			t.Fatal("should give nil CompanyStatus")
		}

		// block company
		status = &CompanyStatus{UserID: userID, Status: CompanyUnsafe, CreatedAt: time.Now()}
		if err := repo.upsertCompanyStatus(companyID, status); err != nil {
			t.Errorf("addCompanyBlock: shouldn't error, but got %v", err)
		}

		// verify
		status, err = repo.getCompanyStatus(companyID)
		if err != nil {
			t.Error(err)
		}
		if status == nil {
			t.Fatal("empty CompanyStatus")
		}
		if status.UserID == "" || string(status.Status) == "" {
			t.Errorf("invalid CompanyStatus: %#v", status)
		}
		if status.Status != CompanyUnsafe {
			t.Errorf("status.Status=%v", status.Status)
		}

		// unblock
		status = &CompanyStatus{UserID: userID, Status: CompanyException, CreatedAt: time.Now()}
		if err := repo.upsertCompanyStatus(companyID, status); err != nil {
			t.Errorf("addCompanyBlock: shouldn't error, but got %v", err)
		}

		status, err = repo.getCompanyStatus(companyID)
		if err != nil {
			t.Error(err)
		}
		if status == nil {
			t.Fatal("empty CompanyStatus")
		}
		if status.UserID == "" || string(status.Status) == "" {
			t.Errorf("invalid CompanyStatus: %#v", status)
		}
		if status.Status != CompanyException {
			t.Errorf("status.Status=%v", status.Status)
		}

		// edgae case
		status, err = repo.getCompanyStatus("")
		if status != nil {
			t.Error("empty companyID shouldn return nil status")
		}
		if err == nil {
			t.Error("but an error should be returned")
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteCompanyRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteCompanyRepository{mysqlDB, log.NewNopLogger()})
}
