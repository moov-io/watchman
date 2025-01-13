// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipeline__normalizeStep(t *testing.T) {
	got := LowerAndRemovePunctuation("Nicolás Maduro")
	require.Equal(t, "nicolas maduro", got)
}

// TestLowerAndRemovePunctuation ensures we are trimming and UTF-8 normalizing strings
// as expected. This is needed since our datafiles are normalized for us.
func TestLowerAndRemovePunctuation(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"remove accents", "nicolás maduro", "nicolas maduro"},
		{"convert IAcute", "Delcy Rodríguez", "delcy rodriguez"},
		{"issue 58", "Raúl Castro", "raul castro"},
		{"remove hyphen", "ANGLO-CARIBBEAN ", "anglo caribbean"},
		// Issue 483
		{"issue 483 #1", "11420 CORP.", "11420 corp"},
		{"issue 483 #2", "11,420.2-1 CORP.", "114202 1 corp"},
	}
	for _, tc := range tests {
		guess := LowerAndRemovePunctuation(tc.input)
		require.Equal(t, tc.expected, guess)
	}
}
