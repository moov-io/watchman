// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/base"

	"github.com/go-kit/kit/log"
)

func createTestWatchRepository(t *testing.T) *sqliteWatchRepository {
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

func TestCompanyWatch(t *testing.T) {
	repo := createTestWatchRepository(t)

	companyID := base.ID()
	if err := repo.removeCompanyWatch(companyID, base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// add watch, then remove
	watchID, err := repo.addCompanyWatch(companyID, watchRequest{Webhook: "https://moov.io"})
	if err != nil {
		t.Errorf("companyID=%q got error: %v", companyID, err)
	}
	if watchID == "" {
		t.Error("empty watchID")
	}

	// remove
	if err := repo.removeCompanyWatch(companyID, watchID); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestCompanyNameWatch(t *testing.T) {
	repo := createTestWatchRepository(t)

	if err := repo.removeCompanyNameWatch(base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// Add
	name := base.ID()
	watchID, err := repo.addCompanyNameWatch(name, "https://moov.io", "authToken")
	if err != nil {
		t.Errorf("name=%q got error: %v", name, err)
	}
	if watchID == "" {
		t.Error("empty watchID")
	}

	// Remove
	if err := repo.removeCompanyNameWatch(watchID); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestCustomerWatch(t *testing.T) {
	repo := createTestWatchRepository(t)

	customerID := base.ID()
	if err := repo.removeCustomerWatch(customerID, base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// add watch, then remove
	watchID, err := repo.addCustomerWatch(customerID, watchRequest{Webhook: "https://moov.io"})
	if err != nil {
		t.Errorf("customerID=%q got error: %v", customerID, err)
	}
	if watchID == "" {
		t.Error("empty watchID")
	}

	// remove
	if err := repo.removeCustomerWatch(customerID, watchID); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestCustomerNameWatch(t *testing.T) {
	repo := createTestWatchRepository(t)

	if err := repo.removeCustomerNameWatch(base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	// Add
	name := base.ID()
	watchID, err := repo.addCustomerNameWatch(name, "https://moov.io", "authToken")
	if err != nil {
		t.Errorf("name=%q got error: %v", name, err)
	}
	if watchID == "" {
		t.Error("empty watchID")
	}

	// Remove
	if err := repo.removeCustomerNameWatch(watchID); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestWatchCursor_ID(t *testing.T) {
	repo := createTestWatchRepository(t)
	defer repo.close()

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
	if secondBatch[0].id != watchID2 || secondBatch[0].webhook != "https://moov.io/2" || secondBatch[0].customerID == "" {
		t.Errorf("unknown watch: %v", secondBatch[0])
	}
}

func TestWatchCursor_Names(t *testing.T) {
	repo := createTestWatchRepository(t)
	defer repo.close()

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
	if secondBatch[0].id != watchID2 || secondBatch[0].webhook != "https://moov.io/2" || secondBatch[0].customerName != "jane doe" {
		t.Errorf("unknown watch: %v", secondBatch[0])
	}
}
