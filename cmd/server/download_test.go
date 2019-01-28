// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestSearcher__refreshInterval(t *testing.T) {
	if v := getOFACRefreshInterval(nil, ""); v.String() != "12h0m0s" {
		t.Errorf("Got %v", v)
	}
	// override
	if v := getOFACRefreshInterval(nil, "60s"); v.String() != "1m0s" {
		t.Errorf("Got %v", v)
	}
}

func TestSearcher__refreshData(t *testing.T) {
	if testing.Short() {
		return
	}

	s := &searcher{}
	if err := s.refreshData(); err != nil {
		t.Fatal(err)
	}
	if len(s.Addresses) == 0 {
		t.Errorf("empty Addresses")
	}
	if len(s.Alts) == 0 {
		t.Errorf("empty Alts")
	}
	if len(s.SDNs) == 0 {
		t.Errorf("empty SDNs")
	}
}
