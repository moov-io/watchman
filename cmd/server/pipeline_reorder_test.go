// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"
)

func TestPipeline__reorderSDNStep(t *testing.T) {
	nn := &Name{
		Processed: "Last, First Middle",
		sdn: &ofac.SDN{
			SDNType: "individual",
		},
	}

	step := &reorderSDNStep{}
	if err := step.apply(nn); err != nil {
		t.Fatal(err)
	}

	if nn.Processed != "First Middle Last" {
		t.Errorf("nn.Processed=%v", nn.Processed)
	}
}

func TestReorderSDNName(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"Jane Doe", "Jane Doe"}, // no change, control (without commas)
		{"Doe Other, Jane", "Jane Doe Other"},
		{"Last, First Middle", "First Middle Last"},
		{"FELIX B. MADURO S.A.", "FELIX B. MADURO S.A."}, // keep .'s in a name
		{"MADURO MOROS, Nicolas", "Nicolas MADURO MOROS"},
		{"IBRAHIM, Sadr", "Sadr IBRAHIM"},
		{"AL ZAWAHIRI, Dr. Ayman", "Dr. Ayman AL ZAWAHIRI"},
		// Issue 115
		{"Bush, George W", "George W Bush"},
		{"RIZO MORENO, Jorge Luis", "Jorge Luis RIZO MORENO"},
	}
	for i := range cases {
		guess := reorderSDNName(cases[i].input, "individual")
		if guess != cases[i].expected {
			t.Errorf("reorderSDNName(%q)=%q expected %q", cases[i].input, guess, cases[i].expected)
		}
	}

	// Entities
	cases = []struct {
		input, expected string
	}{
		// Issue 483
		{"11420 CORP.", "11420 CORP."},
		{"11,420.2-1 CORP.", "11,420.2-1 CORP."},
	}
	for i := range cases {
		guess := reorderSDNName(cases[i].input, "") // blank refers to a company
		if guess != cases[i].expected {
			t.Errorf("reorderSDNName(%q)=%q expected %q", cases[i].input, guess, cases[i].expected)
		}
	}
}
