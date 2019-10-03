// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/ofac/internal/database"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func TestSearcher__refreshInterval(t *testing.T) {
	if v := getOFACRefreshInterval(log.NewNopLogger(), ""); v.String() != "12h0m0s" {
		t.Errorf("Got %v", v)
	}
	if v := getOFACRefreshInterval(log.NewNopLogger(), "60s"); v.String() != "1m0s" {
		t.Errorf("Got %v", v)
	}
	if v := getOFACRefreshInterval(log.NewNopLogger(), "off"); v != 0*time.Second {
		t.Errorf("got %v", v)
	}

	// cover another branch
	s := &searcher{
		logger: log.NewNopLogger(),
	}
	s.periodicDataRefresh(0*time.Second, nil, nil)
}

func TestSearcher__refreshData(t *testing.T) {
	if testing.Short() {
		return
	}

	s := &searcher{}
	stats, err := s.refreshData("")
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Addresses) == 0 || stats.Addresses == 0 {
		t.Errorf("empty Addresses=%d stats.Addresses=%d", len(s.Addresses), stats.Addresses)
	}
	if len(s.Alts) == 0 || stats.Alts == 0 {
		t.Errorf("empty Alts=%d or stats.Alts=%d", len(s.Alts), stats.Alts)
	}
	if len(s.SDNs) == 0 || stats.SDNs == 0 {
		t.Errorf("empty SDNs=%d or stats.SDNs=%d", len(s.SDNs), stats.SDNs)
	}
}

func TestDownload_record(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteDownloadRepository) {
		stats := &downloadStats{SDNs: 1, Alts: 12, Addresses: 42, DeniedPersons: 13}
		if err := repo.recordStats(stats); err != nil {
			t.Fatal(err)
		}

		downloads, err := repo.latestDownloads(5)
		if err != nil {
			t.Fatal(err)
		}
		if len(downloads) != 1 {
			t.Errorf("got %d downloads", len(downloads))
		}
		dl := downloads[0]
		if dl.SDNs != stats.SDNs {
			t.Errorf("dl.SDNs=%d stats.SDNs=%d", dl.SDNs, stats.SDNs)
		}
		if dl.Alts != stats.Alts {
			t.Errorf("dl.Alts=%d stats.Alts=%d", dl.Alts, stats.Alts)
		}
		if dl.Addresses != stats.Addresses {
			t.Errorf("dl.Addresses=%d stats.Addresses=%d", dl.Addresses, stats.Addresses)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteDownloadRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.CreateTestMySQLDB(t)
	defer mysqlDB.Close()
	check(t, &sqliteDownloadRepository{mysqlDB.DB, log.NewNopLogger()})
}

func TestDownload_route(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteDownloadRepository) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/downloads", nil)
		req.Header.Set("x-user-id", "test")

		repo.recordStats(&downloadStats{SDNs: 1, Alts: 421, Addresses: 1511, DeniedPersons: 731})

		router := mux.NewRouter()
		addDownloadRoutes(nil, router, repo)
		router.ServeHTTP(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d", w.Code)
		}

		var downloads []Download
		if err := json.NewDecoder(w.Body).Decode(&downloads); err != nil {
			t.Error(err)
		}
		if len(downloads) != 1 {
			t.Errorf("got %d downloads: %v", len(downloads), downloads)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteDownloadRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.CreateTestMySQLDB(t)
	defer mysqlDB.Close()
	check(t, &sqliteDownloadRepository{mysqlDB.DB, log.NewNopLogger()})
}

func TestDownload__lastRefresh(t *testing.T) {
	start := time.Now()
	time.Sleep(5 * time.Millisecond) // force start to be before our calls

	if when := lastRefresh(""); when.Before(start) {
		t.Errorf("expected time.Now()")
	}

	// make a temp dir (initially with nothing in it)
	dir, err := ioutil.TempDir("", "lastRefresh")
	if err != nil {
		t.Fatal(err)
	}

	if when := lastRefresh(dir); !when.IsZero() {
		t.Errorf("expected zero time: %v", t)
	}

	// add a file and get it's mtime
	path := filepath.Join(dir, "out.txt")
	if err := ioutil.WriteFile(path, []byte("hello, world"), 0600); err != nil {
		t.Fatal(err)
	}
	if info, err := os.Stat(path); err != nil {
		t.Fatal(err)
	} else {
		if when := lastRefresh(dir); !when.Equal(info.ModTime()) {
			t.Errorf("t=%v", when)
		}
	}
}
