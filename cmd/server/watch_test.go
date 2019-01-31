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
