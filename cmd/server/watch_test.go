// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/database"

	"github.com/stretchr/testify/require"
)

func createTestWatchRepository(t *testing.T) *sqliteWatchRepository {
	t.Helper()

	db := database.CreateTestSqliteDB(t)
	return &sqliteWatchRepository{db.DB, log.NewNopLogger()}
}

func TestCompanyWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteWatchRepository) {
		companyID := base.ID()
		err := repo.removeCompanyWatch(companyID, base.ID())
		require.NoError(t, err)

		// add watch, then remove
		watchID, err := repo.addCompanyWatch(companyID, watchRequest{Webhook: "https://moov.io"})
		require.NoError(t, err)
		require.NotEmpty(t, watchID)

		// remove
		err = repo.removeCompanyWatch(companyID, watchID)
		require.NoError(t, err)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}

func TestCompanyNameWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteWatchRepository) {
		err := repo.removeCompanyNameWatch(base.ID())
		require.NoError(t, err)

		// Add
		name := base.ID()
		watchID, err := repo.addCompanyNameWatch(name, "https://moov.io", "authToken")
		require.NoError(t, err)
		require.NotEmpty(t, watchID)

		// Remove
		err = repo.removeCompanyNameWatch(watchID)
		require.NoError(t, err)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}

func TestCustomerWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteWatchRepository) {
		customerID := base.ID()
		err := repo.removeCustomerWatch(customerID, base.ID())
		require.NoError(t, err)

		// add watch, then remove
		watchID, err := repo.addCustomerWatch(customerID, watchRequest{Webhook: "https://moov.io"})
		require.NoError(t, err)
		require.NotEmpty(t, watchID)

		// remove
		err = repo.removeCustomerWatch(customerID, watchID)
		require.NoError(t, err)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}

func TestCustomerNameWatch(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteWatchRepository) {
		err := repo.removeCustomerNameWatch(base.ID())
		require.NoError(t, err)

		// Add
		name := base.ID()
		watchID, err := repo.addCustomerNameWatch(name, "https://moov.io", "authToken")
		require.NoError(t, err)
		require.NotEmpty(t, watchID)

		// Remove
		err = repo.removeCustomerNameWatch(watchID)
		require.NoError(t, err)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}

func TestWatchCursor_ID(t *testing.T) {
	t.Parallel()

	check := func(t *testing.T, repo *sqliteWatchRepository) {
		cur := repo.getWatchesCursor(log.NewNopLogger(), 4) // batchSize is divided in 4 parts to equally grab customer, customer name, company, and company name watches

		// insert some watches
		watchID1, _ := repo.addCustomerWatch(base.ID(), watchRequest{Webhook: "https://moov.io/1"})
		watchID2, _ := repo.addCustomerWatch(base.ID(), watchRequest{Webhook: "https://moov.io/2"})
		watchID3, _ := repo.addCompanyWatch(base.ID(), watchRequest{Webhook: "https://moov.io/3"})

		// get first batch (should have 2 watches)
		firstBatch, err := cur.Next()
		if len(firstBatch) != 2 || err != nil {
			t.Fatalf("len(firstBatch)=%d expected 2, err=%v", len(firstBatch), err)
		}
		for i := range firstBatch {
			switch firstBatch[i].id {
			case watchID1:
				if firstBatch[i].webhook != "https://moov.io/1" || firstBatch[i].customerID == "" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
			case watchID3:
				if firstBatch[i].webhook != "https://moov.io/3" || firstBatch[i].companyID == "" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
			default:
				t.Errorf("unknown watch: %v", firstBatch[i])
			}
		}

		// second batch (should only have watchID3)
		secondBatch, err := cur.Next()
		if len(secondBatch) != 1 || err != nil {
			t.Fatalf("len(secondBatch)=%d expected 1, err=%v", len(secondBatch), err)
		}
		require.Equal(t, watchID2, secondBatch[0].id)
		require.Equal(t, "https://moov.io/2", secondBatch[0].webhook)
		require.NotEmpty(t, secondBatch[0].customerID)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}

func TestWatchCursor_Names(t *testing.T) {
	check := func(t *testing.T, repo *sqliteWatchRepository) {
		cur := repo.getWatchesCursor(log.NewNopLogger(), 4)

		// insert some watches
		watchID1, _ := repo.addCustomerNameWatch("foo corp", "https://moov.io/1", base.ID())
		watchID2, _ := repo.addCustomerNameWatch("jane doe", "https://moov.io/2", base.ID())
		watchID3, _ := repo.addCompanyNameWatch("bar corp", "https://moov.io/3", base.ID())

		// get first batch (should have 2 watches)
		firstBatch, err := cur.Next()
		if len(firstBatch) != 2 || err != nil {
			t.Fatalf("len(firstBatch)=%d expected 2, err=%v", len(firstBatch), err)
		}
		for i := range firstBatch {
			switch firstBatch[i].id {
			case watchID1:
				if firstBatch[i].webhook != "https://moov.io/1" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
				if firstBatch[i].customerName != "foo corp" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
			case watchID3:
				if firstBatch[i].webhook != "https://moov.io/3" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
				if firstBatch[i].companyName != "bar corp" {
					t.Errorf("watch %#v didn't match", firstBatch[i])
				}
			default:
				t.Errorf("unknown watch: %v", firstBatch[i])
			}
		}

		// second batch (should only have watchID3)
		secondBatch, err := cur.Next()
		if len(secondBatch) != 1 || err != nil {
			t.Fatalf("len(secondBatch)=%d expected 1, err=%v", len(secondBatch), err)
		}
		require.Equal(t, watchID2, secondBatch[0].id)
		require.Equal(t, "https://moov.io/2", secondBatch[0].webhook)
		require.Equal(t, "jane doe", secondBatch[0].customerName)
	}

	// SQLite tests
	sqliteDB := database.CreateTestSqliteDB(t)
	defer sqliteDB.Close()
	check(t, &sqliteWatchRepository{sqliteDB.DB, log.NewNopLogger()})

	// MySQL tests
	mysqlDB := database.TestMySQLConnection(t)
	check(t, &sqliteWatchRepository{mysqlDB, log.NewNopLogger()})
}
