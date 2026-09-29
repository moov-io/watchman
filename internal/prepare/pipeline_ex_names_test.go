// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitExNames(t *testing.T) {
	cases := []struct {
		in      string
		cleaned string
		former  []string
	}{
		{"Ocean Pioneer (ex-Cape Diamond)", "Ocean Pioneer", []string{"Cape Diamond"}},
		{"Star Voyager (ex-Southern Glory)", "Star Voyager", []string{"Southern Glory"}},
		{"M/T Ondara Leigh (ex-Pelican Morrow)", "M/T Ondara Leigh", []string{"Pelican Morrow"}},
		{"Wheatland Glory (ex-Pacific Sower)", "Wheatland Glory", []string{"Pacific Sower"}},
		{"M.T. THONG NAWA 3 (Laem Chabang)", "M.T. THONG NAWA 3 (Laem Chabang)", nil},
		{"Solenne Harbour", "Solenne Harbour", nil},
		{"", "", nil},
		{"Foo (ex-Bar) (ex-Baz)", "Foo", []string{"Bar", "Baz"}},
		{"Foo (EX Cape Diamond)", "Foo", []string{"Cape Diamond"}},
		{"Ocean Pioneer (example)", "Ocean Pioneer (example)", nil},
		{"Star Voyager (ex)", "Star Voyager (ex)", nil},
		{"M/T Rosalind Meadow (f.k.a. Meadow Spirit)", "M/T Rosalind Meadow", []string{"Meadow Spirit"}},
		{"Nordbryggen Marine Coatings AS (f.k.a. Bryggen Paint AS)", "Nordbryggen Marine Coatings AS", []string{"Bryggen Paint AS"}},
		{"Ostrowski Grain Partners LLC (f.k.a. Ostrowski Feed & Grain LLC)", "Ostrowski Grain Partners LLC", []string{"Ostrowski Feed & Grain LLC"}},
		{"BEIRUT CEDAR TRADING SAL, f.k.a. LEBANON CEDAR IMPORT EXPORT SAL", "BEIRUT CEDAR TRADING SAL", []string{"LEBANON CEDAR IMPORT EXPORT SAL"}},
		{"Dalgleish Rope Works Ltd, formerly Dalgleish & Muir Ropes Ltd", "Dalgleish Rope Works Ltd", []string{"Dalgleish & Muir Ropes Ltd"}},
		{"Olmsbury Grain Corporation f/k/a Olmsbury Milling Corporation", "Olmsbury Grain Corporation", []string{"Olmsbury Milling Corporation"}},
		{"Kenari Food Distribution Sdn. Bhd. (formerly known as Kenari Frozen Foods Sdn. Bhd.)", "Kenari Food Distribution Sdn. Bhd.", []string{"Kenari Frozen Foods Sdn. Bhd."}},
		{"Tamarind Crest Commodities Pte. Ltd. (formerly Tamarind Crest Resources Pte. Ltd.)", "Tamarind Crest Commodities Pte. Ltd.", []string{"Tamarind Crest Resources Pte. Ltd."}},
		{"MV Karachi Crescent (ex Mandvi Bay, ex Wanshan Pearl)", "MV Karachi Crescent", []string{"Mandvi Bay", "Wanshan Pearl"}},
		{"MV Stamatia Horizon (ex-Elenitsa Z, ex-Aegean Lynx)", "MV Stamatia Horizon", []string{"Elenitsa Z", "Aegean Lynx"}},
		{"Superior Harvester (ex Duluth Harvester, ex Iron Range Harvester)", "Superior Harvester", []string{"Duluth Harvester", "Iron Range Harvester"}},
		{"Theodosia K (ex-Lindos Harrier until 2019)", "Theodosia K", []string{"Lindos Harrier"}},
		{"MV Apapa Falcon (ex Warri Osprey, 2020)", "MV Apapa Falcon", []string{"Warri Osprey"}},
		{"Kalliopi Dawn ex Orion Tanager", "Kalliopi Dawn", []string{"Orion Tanager"}},
		{"COASTAL DEFENDER EX ATLA", "COASTAL DEFENDER", []string{"ATLA"}},
		{"MT QUINTARA SPIRIT EX PALOMERA", "MT QUINTARA SPIRIT", []string{"PALOMERA"}},
		{"Qamar Al Khaleej ex-Rimal Star", "Qamar Al Khaleej", []string{"Rimal Star"}},
		{"Export Trading Ltd", "Export Trading Ltd", nil},
		{"Example Holdings", "Example Holdings", nil},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			cleaned, former := SplitExNames(tc.in)
			require.Equal(t, tc.cleaned, cleaned)
			require.Equal(t, tc.former, former)
		})
	}
}
