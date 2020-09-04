// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"errors"
	"testing"

	"github.com/go-kit/kit/log"
)

// Test against postgres 11.0
func TestPostgres__basic(t *testing.T) {
	db := CreateTestPostgresDB(t)
	defer db.Close()

	if err := db.DB.Ping(); err != nil {
		t.Fatal(err)
	}

	// create a phony Postgres
	m := postgresConnection(log.NewNopLogger(), "user", "pass", "127.0.0.1:4432", "watchman")

	conn, err := m.Connect()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if conn != nil || err == nil {
		t.Fatalf("conn=%#v expected error", conn)
	}
}

func TestPostgresUniqueViolation(t *testing.T) {
	err := errors.New(`problem upserting depository="534ff5dba099ba334eeedd54773d11733ababdf5", userId="d341ee56453abb34ac899871ee266cd88321aa23": [Err] ERROR:  duplicate key value violates unique constraint`)
	if !UniqueViolation(err) {
		t.Error("should have matched unique violation")
	}
}
