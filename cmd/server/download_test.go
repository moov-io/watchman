// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestDownloadAndParse(t *testing.T) {
	if testing.Short() {
		return
	}

	reader, err := getAndParseOFACData()
	if err != nil {
		t.Fatal(err)
	}
	if reader == nil {
		t.Fatal("nil ofac.Reader")
	}

	if len(reader.Addresses) == 0 {
		t.Errorf("empty Addresses")
	}
	if len(reader.AlternateIdentities) == 0 {
		t.Errorf("empty AlternateIdentities")
	}
	if len(reader.SDNs) == 0 {
		t.Errorf("empty SDNs")
	}
}
