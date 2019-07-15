// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"runtime"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestSQLite__basic(t *testing.T) {
	db := CreateTestSqliteDB(t)
	defer db.Close()

	if err := db.DB.Ping(); err != nil {
		t.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		t.Skip("/dev/null doesn't exist on Windows")
	}

	// error case
	s := sqliteConnection(log.NewNopLogger(), "/tmp/path/doesnt/exist")

	conn, err := s.Connect()
	defer conn.Close()

	if err := conn.Ping(); err == nil {
		t.Fatal("expected error")
	}
	if err == nil {
		t.Fatalf("conn=%#v expected error", conn)
	}
}

func TestSqlite__getSqlitePath(t *testing.T) {
	if v := getSqlitePath(); v != "ofac.db" {
		t.Errorf("got %s", v)
	}
}
