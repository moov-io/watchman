// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanonicalizeEnglishLegalForms(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"quillmont hydraulics inc", "quillmont hydraulics inc"},
		{"quillmont hydraulics incorporated", "quillmont hydraulics inc"},
		{"harrowfield bearings limited", "harrowfield bearings ltd"},
		{"harrowfield bearings ltd", "harrowfield bearings ltd"},
		{"pellbrook cold storage limited liability company", "pellbrook cold storage llc"},
		{"pellbrook cold storage llc", "pellbrook cold storage llc"},
		{"carrowmere holdings public limited company", "carrowmere holdings plc"},
		{"carrowmere holdings plc", "carrowmere holdings plc"},
		{"dongguan fengtai plastic products co ltd", "dongguan fengtai plastic products co ltd"},
		{"dongguan fengtai plastic products company limited", "dongguan fengtai plastic products co ltd"},
		{"northgate valve supply corporation", "northgate valve supply corp"},
		{"tamarind bay seafood company", "tamarind bay seafood co"},
		{"korvessa industrial gmbh", "korvessa industrial gmbh"},
		{"", ""},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, CanonicalizeEnglishLegalForms(tc.in))
		})
	}

	t.Run("spaced letter llc from punctuation strip", func(t *testing.T) {
		in := "pellbrook cold storage " + strings.Join([]string{"l", "l", "c"}, " ")
		require.Equal(t, "pellbrook cold storage llc", CanonicalizeEnglishLegalForms(in))
	})
	t.Run("spaced letter plc from punctuation strip", func(t *testing.T) {
		in := "carrowmere holdings " + strings.Join([]string{"p", "l", "c"}, " ")
		require.Equal(t, "carrowmere holdings plc", CanonicalizeEnglishLegalForms(in))
	})
}
