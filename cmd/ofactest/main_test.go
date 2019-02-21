// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestOFAC_getBasePath(t *testing.T) {
	cases := []struct {
		addr     string
		local    bool
		expected string
	}{
		{"http://localhost:8084", false, "http://localhost:8084"},
		{"http://localhost:8084/", false, "http://localhost:8084"},
		{defaultApiAddress, true, "http://localhost:8084"},
		{defaultApiAddress, false, defaultApiAddress + "/v1/ofac"},
	}
	for i := range cases {
		got := getBasePath(cases[i].addr, cases[i].local)
		if got != cases[i].expected {
			t.Errorf("idx=%d got=%q expected=%q", i, got, cases[i].expected)
		}
	}
}
