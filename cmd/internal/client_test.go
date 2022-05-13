// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package internal

import (
	"testing"
)

func TestWatchman_addr(t *testing.T) {
	cases := []struct {
		addr     string
		local    bool
		expected string
	}{
		{"http://localhost:8084", false, "http://localhost:8084"},
		{"http://localhost:8084/", false, "http://localhost:8084"},
		{DefaultApiAddress, true, "http://localhost:8084"},
		{DefaultApiAddress, false, DefaultApiAddress + "/v1/watchman"},
	}
	for i := range cases {
		got := addr(cases[i].addr, cases[i].local)
		if got != cases[i].expected {
			t.Errorf("idx=%d got=%q expected=%q", i, got, cases[i].expected)
		}
	}
}
