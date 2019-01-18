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

	if len(reader.AddressArray) == 0 {
		t.Errorf("empty AddressArray")
	}
	if len(reader.AlternateIdentityArray) == 0 {
		t.Errorf("empty AlternateIdentityArray")
	}
	if len(reader.SDNArray) == 0 {
		t.Errorf("empty SDNArray")
	}
}
