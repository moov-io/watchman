// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestSplitNameIntoAlts(t *testing.T) {
	cases := []struct {
		input string

		name     string
		altNames []string
	}{
		{
			input: "Pakistan Atomic Energy Commission (PAEC), and subordinate entity Nuclear reactors (including power plants), fuel reprocessing and enrichment facilities, all uranium processing, conversion and enrichment facilities, heavy water production facilities and any collocated ammonia plants.",
			name:  "Pakistan Atomic Energy Commission",
			altNames: []string{
				"PAEC",
				"Nuclear reactors (including power plants), fuel reprocessing and enrichment facilities, all uranium processing, conversion and enrichment facilities, heavy water production facilities and any collocated ammonia plants.",
			},
		},
		{
			input: "Pakistan Atomic Energy Commission (PAEC), and subordinate entity Pakistan Institute for Nuclear Science and Technology (PINSTECH)",
			name:  "Pakistan Atomic Energy Commission",
			altNames: []string{
				"PAEC",
				"Pakistan Institute for Nuclear Science and Technology (PINSTECH)",
			},
		},
		{
			input: "Universal Enterprise Limited [aka General Technology Limited, aka Beijing Luo Luo Tech Development Limited, aka Tiger Force Electronics, aka Foshan Nanhai Winhope Trade Company] (Hong Kong Entity), and its sub-units and successors",
			name:  "Universal Enterprise Limited",
			altNames: []string{
				"General Technology Limited",
				"Beijing Luo Luo Tech Development Limited",
				"Tiger Force Electronics",
				"Foshan Nanhai Winhope Trade Company",
				"Hong Kong Entity",
			},
		},
		{
			input: "Innovative Equipment, and its sub-units and successors",
			name:  "Innovative Equipment",
		},
		{
			input: "China Electronics Technology Group Corporation 13th Research Institute (CETC 13).  Subordinate Institution - MT Microsystems",
			name:  "China Electronics Technology Group Corporation 13th Research Institute",
			altNames: []string{
				"MT Microsystems",
				"CETC 13",
			},
		},
		{
			input: "Huawei Technologies Co., Ltd. (Huawei).  Affiliated Entity: Huawei Terminal (Shenzhen) Co., Ltd.",
			name:  "Huawei Technologies Co., Ltd.",
			altNames: []string{
				"Huawei",
				"Huawei Terminal (Shenzhen) Co., Ltd.",
			},
		},
		{
			input: "The Ministry of Defence of the Republic of Belarus, including the Armed Forces of Belarus and all operating units wherever located.  This includes the national armed services (army and air force), as well as the national guard and national police, government intelligence or reconnaissance organizations of the Republic of Belarus.  All addresses located in Belarus.",
			name:  "The Ministry of Defence of the Republic of Belarus",
			altNames: []string{
				"the Armed Forces of Belarus and all operating units wherever located.  This includes the national armed services (army and air force), as well as the national guard and national police, government intelligence or reconnaissance organizations of the Republic of Belarus.  All addresses located in Belarus.",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			name, alts := splitNameIntoAlts(tc.input)

			require.Equal(t, tc.name, name)
			require.ElementsMatch(t, tc.altNames, alts)
		})
	}
}

func TestMapAddresses(t *testing.T) {
	cases := []struct {
		input    string
		expected []search.Address
	}{
		{
			input: "P.O. Box 1114, Islamabad, PK",
			expected: []search.Address{
				{
					Line1:   "P.O. Box 1114",
					City:    "Islamabad",
					Country: "Pakistan",
				},
			},
		},
		{
			input: "29-M, Civic Centre, Model Town Ext., Lahore, 43700, PK; Office No. 610, 6th Floor, Progressive Centre, 30-A, Block No. 6, P.E.C.H.S., Karachi, PK; P.O. Box 245221, Riyadh 11312, Kingdom of Saudi Arabia, Riyadh, SA; P.O. Box 97, Abu Dhabi, AE",
			expected: []search.Address{
				{
					Line1:      "29-M",
					Line2:      "Civic Centre Model Town Ext.",
					City:       "Lahore",
					PostalCode: "43700",
					Country:    "Pakistan",
				},
				{
					Line1:   "Office No. 610",
					Line2:   "6th Floor Progressive Centre 30-A Block No. 6",
					City:    "Karachi",
					Country: "Pakistan",
				},
				{
					Line1:   "P.O. Box 245221",
					Line2:   "Riyadh 11312",
					City:    "Riyadh",
					Country: "Saudi Arabia",
				},
				{
					Line1:   "P.O. Box 97",
					City:    "Abu Dhabi",
					Country: "United Arab Emirates",
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := mapAddresses(SanctionsEntry{
				Addresses: tc.input,
			})

			require.ElementsMatch(t, tc.expected, got)
		})
	}
}
