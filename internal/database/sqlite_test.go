// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	// "database/sql"
	// "io/ioutil"
	// "os"
	// "path/filepath"
	"testing"
)

// type testSqliteDB struct {
// 	db *sql.DB

// 	dir string // temp dir created for sqlite files
// }

// func (r *testSqliteDB) close() error {
// 	if err := r.db.Close(); err != nil {
// 		return err
// 	}
// 	return os.RemoveAll(r.dir)
// }

func TestSqlite__getSqlitePath(t *testing.T) {
	if v := getSqlitePath(); v != "ofac.db" {
		t.Errorf("got %s", v)
	}
}
