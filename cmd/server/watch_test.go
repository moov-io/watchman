// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base"

	"github.com/go-kit/kit/log"
)

func createTestCustomerWatchRepository(t *testing.T) *sqliteWatchRepository {
	t.Helper()

	db, err := createTestSqliteDB()
	if err != nil {
		t.Fatal(err)
	}

	return &sqliteWatchRepository{
		db.db,
		log.NewNopLogger(),
	}
}

func TestCustomerWatch(t *testing.T) {
	repo := createTestCustomerWatchRepository(t)

	customerId := base.ID()
	if err := repo.removeCustomerWatch(customerId, base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// add watch, then remove
	watchId, err := repo.addCustomerWatch(customerId, watchRequest{Webhook: "https://moov.io"})
	if err != nil {
		t.Errorf("customerId=%q got error: %v", customerId, err)
	}
	if watchId == "" {
		t.Error("empty watchId")
	}

	// remove
	if err := repo.removeCustomerWatch(customerId, watchId); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestCustomerNameWatch(t *testing.T) {
	repo := createTestCustomerWatchRepository(t)

	if err := repo.removeCustomerNameWatch(base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// Add
	name := base.ID()
	watchId, err := repo.addCustomerNameWatch(name, "https://moov.io")
	if err != nil {
		t.Errorf("name=%q got error: %v", name, err)
	}
	if watchId == "" {
		t.Error("empty watchId")
	}

	// Remove
	if err := repo.removeCustomerNameWatch(watchId); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestWatchCursor(t *testing.T) {
	repo := createTestCustomerWatchRepository(t)
	defer repo.close()

	cur := repo.getWatchesCursor(2) // batchSize

	// insert some watches
	watchId1, _ := repo.addCustomerWatch(base.ID(), watchRequest{Webhook: "https://moov.io/1"})
	watchId2, _ := repo.addCustomerWatch(base.ID(), watchRequest{Webhook: "https://moov.io/2"})
	watchId3, _ := repo.addCustomerWatch(base.ID(), watchRequest{Webhook: "https://moov.io/3"})

	// get first batch (should have 2 watches)
	firstBatch, err := cur.Next()
	if len(firstBatch) != 2 || err != nil {
		t.Fatalf("len(firstBatch)=%d expected 2, err=%v", len(firstBatch), err)
	}
	for i := range firstBatch {
		switch firstBatch[i].id {
		case watchId1:
			if firstBatch[i].webhook != "https://moov.io/1" {
				t.Errorf("watch %#v didn't match", firstBatch[i])
			}
		case watchId2:
			if firstBatch[i].webhook != "https://moov.io/2" {
				t.Errorf("watch %#v didn't match", firstBatch[i])
			}
		default:
			t.Errorf("unknown watch: %v", firstBatch[i])
		}
	}

	// second batch (should only have watchId3)
	secondBatch, err := cur.Next()
	if len(secondBatch) != 1 || err != nil {
		t.Fatalf("len(secondBatch)=%d expected 1, err=%v", len(secondBatch), err)
	}
	if secondBatch[0].id != watchId3 || secondBatch[0].webhook != "https://moov.io/3" {
		t.Errorf("unknown watch: %v", secondBatch[0])
	}
}
