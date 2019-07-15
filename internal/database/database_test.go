// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"testing"

	"github.com/go-kit/kit/log"
)

func TestDatabase(t *testing.T) {
	db, err := New(log.NewNopLogger(), "other")
	if db != nil || err == nil {
		t.Errorf("db=%#v expected error", db)
	}
}
