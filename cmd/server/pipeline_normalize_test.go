// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestPipeline__normalizeStep(t *testing.T) {
	nn := &Name{Processed: "Nicolás Maduro"}

	step := &normalizeStep{}
	if err := step.apply(nn); err != nil {
		t.Fatal(err)
	}

	if nn.Processed != "nicolas maduro" {
		t.Errorf("nn.Processed=%v", nn.Processed)
	}
}

// TestPrecompute ensures we are trimming and UTF-8 normalizing strings
// as expected. This is needed since our datafiles are normalized for us.
func TestPrecompute(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"nicolás maduro", "nicolas maduro"},
		{"Delcy Rodríguez", "delcy rodriguez"},
		{"Raúl Castro", "raul castro"},
		{"ANGLO-CARIBBEAN ", "anglo caribbean"},
	}
	for i := range cases {
		guess := precompute(cases[i].input)
		if guess != cases[i].expected {
			t.Errorf("precompute(%q)=%q expected %q", cases[i].input, guess, cases[i].expected)
		}
	}
}
