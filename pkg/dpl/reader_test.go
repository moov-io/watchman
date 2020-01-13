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

	if _, err := Read(filepath.Join("..", "..", "test", "testdata", "sdn.csv")); err == nil {
		t.Error("expected error")
	}
}
