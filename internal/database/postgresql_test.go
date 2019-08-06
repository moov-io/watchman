// Copyright 2019 DigitalMint [Carlos Saavedra]
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"errors"
	"testing"

	"github.com/go-kit/kit/log"
)

// Tests connection to the PostgreSQL database
func TestPostgreSQL__basic(t *testing.T) {
	// create a phony postgresql
	m := postgreSQLConnection(log.NewNopLogger(), "", "", "", "", false, 5432)

	conn, err := m.Connect()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	// check connection
	if err != nil {
		t.Fatalf("conn=%#v expected error", conn)
	}
}

// Tests Postgres' unique violation error
func TestPostgreSQLUniqueViolation(t *testing.T) {
	err := errors.New(`problem upserting depository="282f6ffcd9ba5b029afbf2b739ee826e22d9df3b", userId="f25f48968da47ef1adb5b6531a1c2197295678ce": ERROR: duplicate key '282f6ffcd9ba5b029afbf2b739ee826e22d9df3b' for key 'PRIMARY'`)
	if !UniqueViolation(err) {
		t.Error("should have matched unique violation")
	}
}
