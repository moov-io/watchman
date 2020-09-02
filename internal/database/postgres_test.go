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
func TestPostgres__basic_11_0(t *testing.T) {
	db := CreateTestPostgresDB(t, "11.0")
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

// Test against postgres 10.0
func TestPostgres__basic_10_0(t *testing.T) {
	db := CreateTestPostgresDB(t, "10.0")
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

// Test against postgres 9.1
func TestPostgres__basic_9_1(t *testing.T) {
	db := CreateTestPostgresDB(t, "9.1")
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
	err := errors.New(`problem upserting depository="534ff5dba099ba334eeedd54773d11733ababdf5", userId="d341ee56453abb34ac899871ee266cd88321aa23": Error: duplicate key value violates unique constraint`)
	if !UniqueViolation(err) {
		t.Error("should have matched unique violation")
	}
}
