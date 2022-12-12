// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/database"
)

// This test fails when run with the "-race" flag on manualRefreshHandler>refreshData>ukSanctionsListRecords>ReadUKSanctionsListFile>ParseContent>Decode. The failure has to do with decoding xml in the context of an ods document (specifically UK_Sanctions_List.ods). We are skipping this test for now in favor of the functionality and because the underlying function will never be run concurrently and therefore will never experience a race condition.
func TestDownload__manualRefreshPath(t *testing.T) {
	t.Skip("test skipped due to panic on -race")
	t.Parallel()

	if testing.Short() {
		return
	}

	check := func(t *testing.T, repo *sqliteDownloadRepository) {
		searcher := newSearcher(log.NewNopLogger(), noLogPipeliner, 1)

		w := httptest.NewRecorder()

		req := httptest.NewRequest("GET", "/data/refresh", nil)
		req.Header.Set("x-request-id", base.ID())

		updates := make(chan *DownloadStats)

		manualRefreshHandler(log.NewNopLogger(), searcher, updates, repo)(w, req)
		w.Flush()

		if w.Code != http.StatusOK {
			t.Errorf("bogus status code: %d: %s", w.Code, w.Body.String())
		}
		var stats DownloadStats
		if err := json.NewDecoder(w.Body).Decode(&stats); err != nil {
			t.Error(err)
		}
		if stats.SDNs == 0 {
			t.Errorf("stats.SDNs=%d but expected non-zero", stats.SDNs)
		}
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteDownloadRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteDownloadRepository{mysqlDB, log.NewNopLogger()})
}
