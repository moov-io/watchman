// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func TestSearcher__refreshInterval(t *testing.T) {
	if v := getOFACRefreshInterval(nil, ""); v.String() != "12h0m0s" {
		t.Errorf("Got %v", v)
	}
	// override
	if v := getOFACRefreshInterval(nil, "60s"); v.String() != "1m0s" {
		t.Errorf("Got %v", v)
	}
}

func TestSearcher__refreshData(t *testing.T) {
	if testing.Short() {
		return
	}

	s := &searcher{}
	stats, err := s.refreshData()
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

func createTestDownloadRepository(t *testing.T) *sqliteDownloadRepository {
	t.Helper()

	db, err := createTestSqliteDB()
	if err != nil {
		t.Fatal(err)
	}

	return &sqliteDownloadRepository{db.db, log.NewNopLogger()}
}

func TestDownload_record(t *testing.T) {
	repo := createTestDownloadRepository(t)
	defer repo.close()

	stats := &downloadStats{1, 12, 42}
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

func TestDownload_route(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/downloads", nil)
	req.Header.Set("x-user-id", "test")

	repo := createTestDownloadRepository(t)
	defer repo.close()

	// save a record
	repo.recordStats(&downloadStats{1, 421, 1511})

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
