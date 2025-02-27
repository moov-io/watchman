// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func testInputs(tb testing.TB, paths ...string) map[string]io.ReadCloser {
	tb.Helper()

	input := make(map[string]io.ReadCloser)
	for _, path := range paths {
		_, filename := filepath.Split(path)

		fd, err := os.Open(path)
		require.NoError(tb, err)

		input[filename] = fd
	}
	return input
}

func TestParseTime(t *testing.T) {
	t.Run("DOB", func(t *testing.T) {
		tt, _ := parseTime(dobPatterns, "01 Apr 1950")
		require.Equal(t, "1950-04-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "01 Feb 1958 to 28 Feb 1958")
		require.Equal(t, "1958-02-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "1928")
		require.Equal(t, "1928-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "1928 to 1930")
		require.Equal(t, "1928-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "Sep 1958")
		require.Equal(t, "1958-09-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 01 Jan 1961")
		require.Equal(t, "1961-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 1934")
		require.Equal(t, "1934-01-01", tt.Format(time.DateOnly))

		tt, _ = parseTime(dobPatterns, "circa 1979-1982")
		require.Equal(t, "1979-01-01", tt.Format(time.DateOnly))
	})

	t.Run("invalidDate", func(t *testing.T) {
		when := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
		require.True(t, invalidDate(when))

		when = time.Now()
		require.False(t, invalidDate(when))
	})
}

func TestParseGovernmentIDs(t *testing.T) {
	tests := []struct {
		name    string
		remarks []string
		want    []search.GovernmentID
	}{
		{
			name: "passport only",
			remarks: []string{
				"Passport A123456 (Iran) expires 2024",
			},
			want: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "Iran",
					Identifier: "A123456",
				},
			},
		},
		{
			name: "drivers license only",
			remarks: []string{
				"Driver's License No. 04900377 (Moldova) issued 02 Jul 2004",
			},
			want: []search.GovernmentID{
				{
					Type:       search.GovernmentIDDriversLicense,
					Country:    "Moldova",
					Identifier: "04900377",
				},
			},
		},
		{
			name: "multiple IDs",
			remarks: []string{
				"Passport A123456 (Iran) expires 2024",
				"Driver's License No. 04900377 (Moldova) issued 02 Jul 2004",
			},
			want: []search.GovernmentID{
				{
					Type:       search.GovernmentIDPassport,
					Country:    "Iran",
					Identifier: "A123456",
				},
				{
					Type:       search.GovernmentIDDriversLicense,
					Country:    "Moldova",
					Identifier: "04900377",
				},
			},
		},
		{
			name: "various driver license formats",
			remarks: []string{
				"Driver License M600161650080 (United States)",
				"Drivers License No. B-12345 (Canada)",
				"Driver's License Number 987654321 (Mexico)",
			},
			want: []search.GovernmentID{
				{
					Type:       search.GovernmentIDDriversLicense,
					Country:    "United States",
					Identifier: "M600161650080",
				},
				{
					Type:       search.GovernmentIDDriversLicense,
					Country:    "Canada",
					Identifier: "12345",
				},
				{
					Type:       search.GovernmentIDDriversLicense,
					Country:    "Mexico",
					Identifier: "987654321",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseGovernmentIDs(tt.remarks)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParseRemarks(t *testing.T) {
	tests := []struct {
		name             string
		remarks          []string
		wantAffiliations []search.Affiliation
		wantSanctions    *search.SanctionsInfo
		wantHistorical   []search.HistoricalInfo
		wantTitles       []string
	}{
		{
			name: "complete remarks",
			remarks: []string{
				"Linked To: ISLAMIC REVOLUTIONARY GUARD CORPS (IRGC)-QODS FORCE; " +
					"Subsidiary Of: BANK OF IRAN; " +
					"Additional Sanctions Information - Subject to Secondary Sanctions; " +
					"Former Name: TEHRAN BANK; Former Name: GLORY; " +
					"Title: Director; Title: Board Member",
			},
			wantAffiliations: []search.Affiliation{
				{
					EntityName: "ISLAMIC REVOLUTIONARY GUARD CORPS (IRGC)-QODS FORCE",
					Type:       "Linked To",
				},
				{
					EntityName: "BANK OF IRAN",
					Type:       "Subsidiary Of",
				},
			},
			wantSanctions: &search.SanctionsInfo{
				Description: "Subject to Secondary Sanctions",
				Secondary:   true,
			},
			wantHistorical: []search.HistoricalInfo{
				{
					Type:  "Former Name",
					Value: "TEHRAN BANK",
				},
				{
					Type:  "Former Name",
					Value: "GLORY",
				},
			},
			wantTitles: []string{
				"Director",
				"Board Member",
			},
		},
		{
			name: "deduplication test",
			remarks: []string{
				"Linked To: CORP A; Linked To: CORP A",         // Duplicate affiliation
				"Former Name: OLD NAME; Former Name: OLD NAME", // Duplicate historical
				"Title: CEO; Title: CEO",                       // Duplicate title
			},
			wantAffiliations: []search.Affiliation{
				{
					EntityName: "CORP A",
					Type:       "Linked To",
				},
			},
			wantSanctions: nil,
			wantHistorical: []search.HistoricalInfo{
				{
					Type:  "Former Name",
					Value: "OLD NAME",
				},
			},
			wantTitles: []string{"CEO"},
		},
		{
			name: "multiple relationships",
			remarks: []string{
				"Linked To: CORP A; Controlled By: CORP B; Owned By: CORP C",
			},
			wantAffiliations: []search.Affiliation{
				{
					EntityName: "CORP A",
					Type:       "Linked To",
				},
				{
					EntityName: "CORP B",
					Type:       "Subsidiary Of",
				},
				{
					EntityName: "CORP C",
					Type:       "Subsidiary Of",
				},
			},
			wantSanctions:  nil,
			wantHistorical: nil,
			wantTitles:     nil,
		},
		{
			name:             "empty remarks",
			remarks:          []string{},
			wantAffiliations: nil,
			wantSanctions:    nil,
			wantHistorical:   nil,
			wantTitles:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAff, gotSanc, gotHist, gotTitles := parseRemarks(tt.remarks)

			// Sort affiliations for consistent comparison
			sort.Slice(gotAff, func(i, j int) bool {
				if gotAff[i].Type != gotAff[j].Type {
					return gotAff[i].Type < gotAff[j].Type
				}
				return gotAff[i].EntityName < gotAff[j].EntityName
			})
			sort.Slice(tt.wantAffiliations, func(i, j int) bool {
				if tt.wantAffiliations[i].Type != tt.wantAffiliations[j].Type {
					return tt.wantAffiliations[i].Type < tt.wantAffiliations[j].Type
				}
				return tt.wantAffiliations[i].EntityName < tt.wantAffiliations[j].EntityName
			})

			// Test affiliations
			require.Equal(t, tt.wantAffiliations, gotAff)

			// Test sanctions info
			if tt.wantSanctions == nil {
				require.Nil(t, gotSanc)
			} else {
				require.Equal(t, tt.wantSanctions.Description, gotSanc.Description)
				require.Equal(t, tt.wantSanctions.Secondary, gotSanc.Secondary)
			}

			// Sort historical info
			sort.Slice(gotHist, func(i, j int) bool {
				return gotHist[i].Value < gotHist[j].Value
			})
			sort.Slice(tt.wantHistorical, func(i, j int) bool {
				return tt.wantHistorical[i].Value < tt.wantHistorical[j].Value
			})

			// Test historical info
			require.Equal(t, tt.wantHistorical, gotHist)

			// Sort titles
			sort.Strings(gotTitles)
			sort.Strings(tt.wantTitles)

			// Test titles
			require.Equal(t, tt.wantTitles, gotTitles)
		})
	}
}

func TestParseCryptoAddresses(t *testing.T) {
	tests := []struct {
		name    string
		remarks []string
		want    []search.CryptoAddress
	}{
		{
			name: "single address",
			remarks: []string{
				"Digital Currency Address - XBT bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			},
			want: []search.CryptoAddress{
				{
					Currency: "XBT",
					Address:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
				},
			},
		},
		{
			name: "multiple addresses",
			remarks: []string{
				"Digital Currency Address - XBT bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
				"Digital Currency Address - ETH 0xb794f5ea0ba39494ce839613fffba74279579268",
			},
			want: []search.CryptoAddress{
				{
					Currency: "XBT",
					Address:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
				},
				{
					Currency: "ETH",
					Address:  "0xb794f5ea0ba39494ce839613fffba74279579268",
				},
			},
		},
		{
			name: "duplicate addresses",
			remarks: []string{
				"Digital Currency Address - XBT bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
				"Digital Currency Address - XBT bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			},
			want: []search.CryptoAddress{
				{
					Currency: "XBT",
					Address:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
				},
			},
		},
		{
			name: "no addresses",
			remarks: []string{
				"Some other remark",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCryptoAddresses(tt.remarks)
			require.Equal(t, tt.want, got)
		})
	}
}
