// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripTrailingVesselPlace(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"MV SIAM ORCHID 7, Bangkok", "MV SIAM ORCHID 7"},
		{"M.T. THONG NAWA 3 (Laem Chabang)", "M.T. THONG NAWA 3"},
		{"T/B CHAO LE 8, Sihanoukville", "T/B CHAO LE 8"},
		{"MT GOLDEN KINNAREE, Laem Chabang", "MT GOLDEN KINNAREE"},
		{"MV Fraternité des Mers, port d'Abidjan", "MV Fraternité des Mers"},
		{"MV Soja Paranaense, Santos", "MV Soja Paranaense"},
		{"N/M Bahia Trader, Paranaguá", "N/M Bahia Trader"},
		{"M/T Ondara Leigh (ex-Pelican Morrow)", "M/T Ondara Leigh (ex-Pelican Morrow)"},
		{"Solenne Harbour", "Solenne Harbour"},
		{"", ""},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, StripTrailingVesselPlace(tc.in))
		})
	}
}
