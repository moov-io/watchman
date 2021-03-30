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
	tests := []struct {
		name, input, expected string
	}{
		{"remove accents", "nicolás maduro", "nicolas maduro"},
		{"convert IAcute", "Delcy Rodríguez", "delcy rodriguez"},
		{"issue 58", "Raúl Castro", "raul castro"},
		{"remove hyphen", "ANGLO-CARIBBEAN ", "anglo caribbean"},
	}
	for i, tc := range tests {
		guess := precompute(tc.input)
		if guess != tc.expected {
			t.Errorf("case: %d name: %s precompute(%q)=%q expected %q", i, tc.name, tc.input, guess, tc.expected)
		}
	}
}
