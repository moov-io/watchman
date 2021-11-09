// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"path/filepath"
	"testing"
)

func TestDPL__read(t *testing.T) {
	dpls, err := Read(filepath.Join("..", "..", "test", "testdata", "dpl.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if len(dpls) != 546 {
		t.Errorf("found %d DPL records", len(dpls))
	}

	// this file is formatted incorrectly for DPL, so we expect all rows to be skipped
	got, err := Read(filepath.Join("..", "..", "test", "testdata", "sdn.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("found %d DPL records, wanted 0", len(got))
	}
}
