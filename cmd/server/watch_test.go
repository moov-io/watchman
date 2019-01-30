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

	if err := repo.removeCustomerWatch(base.ID()); err != nil {
		t.Errorf("expected no error: %v", err)
	}

	customerId := base.ID()

	// add watch, then remove
	watchId, err := repo.addCustomerWatch(customerId, watchRequest{Webhook: "https://moov.io"})
	if err != nil {
		t.Errorf("customerId=%q got error: %v", customerId, err)
	}
	if watchId == "" {
		t.Error("empty watchId")
	}

	// remove
	if err := repo.removeCustomerWatch(watchId); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}
