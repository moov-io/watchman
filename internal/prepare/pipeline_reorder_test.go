// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipeline__reorderSDNStep(t *testing.T) {
	got := ReorderSDNName("Last, First Middle", "individual")
	require.Equal(t, "First Middle Last", got)
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
		{"AL-ZAYDI, Shibl Muhsin 'Ubayd", "Shibl Muhsin 'Ubayd AL-ZAYDI"},
		// Issue 115
		{"Bush, George W", "George W Bush"},
		{"RIZO MORENO, Jorge Luis", "Jorge Luis RIZO MORENO"},
	}
	for i := range cases {
		got := ReorderSDNName(cases[i].input, "individual")
		require.Equal(t, cases[i].expected, got)
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
		got := ReorderSDNName(cases[i].input, "") // blank refers to a company
		require.Equal(t, cases[i].expected, got)
	}
}
