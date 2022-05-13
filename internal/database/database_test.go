// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"errors"
	"testing"

	"github.com/moov-io/base/log"
)

func TestDatabase(t *testing.T) {
	logger := log.NewNopLogger()

	db, err := New(logger, "other")
	if db != nil || err == nil {
		t.Errorf("db=%#v expected error", db)
	}

	db, err = New(logger, "sqlite")
	if db == nil || err != nil {
		t.Errorf("sqlite never errors on initial Connect")
	}
	db.Close()

	db, err = New(logger, "mysql")
	if db != nil || err == nil {
		t.Errorf("expected mysql to fail, but got no error")
	}
}

func TestUniqueViolation(t *testing.T) {
	err := errors.New(`problem upserting depository="282f6ffcd9ba5b029afbf2b739ee826e22d9df3b", userId="f25f48968da47ef1adb5b6531a1c2197295678ce": Error 1062: Duplicate entry '282f6ffcd9ba5b029afbf2b739ee826e22d9df3b' for key 'PRIMARY'`)
	if !UniqueViolation(err) {
		t.Error("should have matched unique violation")
	}

	err = errors.New(`problem upserting depository="7d676c65eccd48090ff238a0d5e35eb6126c23f2", userId="80cfe1311d9eb7659d02cba9ee6cb04ed3739a85": UNIQUE constraint failed: depositories.depository_id`)
	if !UniqueViolation(err) {
		t.Error("should have matched unique violation")
	}
}
