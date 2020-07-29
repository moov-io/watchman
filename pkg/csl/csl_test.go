// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

import (
	"io/ioutil"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestCSL(t *testing.T) {
	logger := log.NewNopLogger()
	dir, err := ioutil.TempDir("", "csl")
	if err != nil {
		t.Fatal(err)
	}

	file, err := Download(logger, dir)
	if err != nil {
		t.Fatal(err)
	}

	cslRecords, err := Read(file)
	if err != nil {
		t.Fatal(err)
	}
	if len(cslRecords.SSIs) == 0 {
		t.Error("parsed zero CSL SSI records")
	}
	if len(cslRecords.ELs) == 0 {
		t.Error("parsed zero CSL EL records")
	}
}
