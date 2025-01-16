package search_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_OFAC_SDN_Person(t *testing.T) {
	indexEntity := ofactest.FindEntity(t, "48603")

	// 48603,"KHOROSHEV, Dmitry Yuryevich","individual","CYBER2",-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,
	// "DOB 17 Apr 1993; POB Russian Federation; nationality Russia; citizen Russia; Email Address khoroshev1@icloud.com;
	// alt. Email Address sitedev5@yandex.ru; Gender Male; Digital Currency Address - XBT bc1qvhnfknw852ephxyc5hm4q520zmvf9maphetc9z;
	// Secondary sanctions risk: Ukraine-/Russia-Related Sanctions Regulations, 31 CFR 589.201; Passport 2018278055 (Russia);
	// alt. Passport 2006801524 (Russia); Tax ID No. 366110340670 (Russia); a.k.a. 'LOCKBITSUPP'."

	// Let's define a sample time to represent the person's birthday
	birthDate := time.Date(1993, time.April, 17, 0, 0, 0, 0, time.UTC)
	laterBirthDate := birthDate.Add(30 * 24 * time.Hour)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Exact match - Person with all fields",
			query: search.Entity[any]{
				Name: "Dmitry Yuryevich KHOROSHEV",
				Type: search.EntityPerson,
				Person: &search.Person{
					Name:      "Dmitry Yuryevich KHOROSHEV",
					BirthDate: &birthDate,
					Gender:    search.GenderMale,
				},
				Contact: search.ContactInfo{
					EmailAddresses: []string{"khoroshev1@icloud.com"},
				},
			},
			expected: 1.00,
		},
		{
			name: "Partial match - Missing birthdate",
			query: search.Entity[any]{
				Name: "Dmitry Yuryevich KHOROSHEV",
				Type: search.EntityPerson,
				Person: &search.Person{
					Name: "Dmitry Yuryevich KHOROSHEV",
				},
				Contact: search.ContactInfo{
					EmailAddresses: []string{"khoroshev1@icloud.com"},
				},
			},
			expected: 0.98,
		},
		{
			name: "Name match only",
			query: search.Entity[any]{
				Name: "Dmitry Yuryevich KHOROSHEV",
				Type: search.EntityPerson,
			},
			expected: 0.98, // TODO(adam): should be lower?
		},
		{
			name: "Fuzzy name match - Alternate identity",
			query: search.Entity[any]{
				Name: "Dmitri Yuryevich",
				Type: search.EntityPerson,
			},
			expected: 0.95,
		},
		{
			name: "Close name but different person details",
			query: search.Entity[any]{
				Name: "Dmitri Yuryvich",
				Type: search.EntityPerson,
				Person: &search.Person{
					BirthDate: &laterBirthDate,
					Gender:    search.GenderMale,
				},
			},
			expected: 0.862,
		},
		{
			name: "Mismatch - Wrong name and no matching details",
			query: search.Entity[any]{
				Name: "SARAH JONES",
				Type: search.EntityPerson,
				Person: &search.Person{
					Gender: "F",
				},
			},
			expected: 0.3110,
		},
		{
			name: "Wrong entity type",
			query: search.Entity[any]{
				Name: "JOHN SMITH",
				Type: search.EntityVessel, // intentionally vessel to mismatch
			},
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := search.DebugSimilarity(debug(t), tc.query, indexEntity)
			require.InDelta(t, tc.expected, score, 0.02)
		})
	}
}

func TestSimilarity_OFAC_SDN_Business(t *testing.T) {
	indexEntity := ofactest.FindEntity(t, "50544")

	// 50544,"AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS",-0- ,"RUSSIA-EO14024",-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,
	// "Website www.dialog.info; alt. Website www.dialog-regions.ru; Secondary sanctions risk: See Section 11 of Executive Order 14024.;
	// Organization Established Date 21 Jul 2020; Tax ID No. 9709063550 (Russia); Business Registration Number 1207700248030 (Russia);
	// a.k.a. 'DIALOGUE REGIONS'; a.k.a. 'DIALOG REGIONY'; a.k.a. 'DIALOGUE'; Linked To: AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG."

	businessCreatedAt := time.Date(2020, time.July, 21, 12, 0, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Exact match - All fields",
			query: search.Entity[any]{
				Name: "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name:    "AUTONOMOUS NON-PROFIT ORGANIZATION DIALOG REGIONS",
					Created: &businessCreatedAt,
					GovernmentIDs: []search.GovernmentID{
						{
							Type:       search.GovernmentIDBusinessRegisration,
							Country:    "Russia",
							Identifier: "1207700248030",
						},
					},
				},
			},
			expected: 1.0,
		},
		{
			name: "Partial match - Fuzzy name, same identifier",
			query: search.Entity[any]{
				Name: "AUTO NON-PROFIT ORGANIZATION",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "AUTO NON-PROFIT ORGANIZATION",
					GovernmentIDs: []search.GovernmentID{
						{
							Type:       search.GovernmentIDTax,
							Country:    "Russia",
							Identifier: "9709063550",
						},
					},
				},
			},
			expected: 0.9647,
		},
		{
			name: "Alt name only + missing ID",
			query: search.Entity[any]{
				Name: "DIALOGUE REGIONS",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "DIALOGUE REGIONS",
				},
			},
			expected: 0.9555,
		},
		{
			name: "Different name, different ID",
			query: search.Entity[any]{
				Name: "BETA Solutions",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "BETA Solutions",
					GovernmentIDs: []search.GovernmentID{
						{
							Type:       search.GovernmentIDBusinessRegisration,
							Country:    "US",
							Identifier: "99999",
						},
					},
				},
			},
			expected: 0.0954,
		},
		{
			name: "Wrong entity type",
			query: search.Entity[any]{
				Name: "ACME Corporation",
				Type: search.EntityVessel,
			},
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := search.DebugSimilarity(debug(t), tc.query, indexEntity)
			require.InDelta(t, tc.expected, score, 0.05)
		})
	}
}

func TestSimilarity_OFAC_SDN_Vessel(t *testing.T) {
	t.Run("47371", func(t *testing.T) {
		indexEntity := ofactest.FindEntity(t, "47371")

		// 47371,"NS LEADER","vessel","RUSSIA-EO14024",-0- ,"A8LU7","Crude Oil Tanker",-0- ,-0- ,"Gabon",-0- ,
		// "Secondary sanctions risk: See Section 11 of Executive Order 14024.; Identification Number IMO 9339301;
		// MMSI 636013272; Linked To: NS LEADER SHIPPING INCORPORATED."

		testCases := []struct {
			name     string
			query    search.Entity[any]
			expected float64
		}{
			{
				name: "Exact match - All fields",
				query: search.Entity[any]{
					Name: "NS LEADER",
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						CallSign:  "A8LU7",
						MMSI:      "636013272",
						IMONumber: "9339301",
					},
				},
				expected: 1.0,
			},
			{
				name: "High confidence - Key fields match",
				query: search.Entity[any]{
					Name: "NS LEADER",
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						CallSign: "A8LU7",
					},
				},
				expected: 1.0,
			},
			{
				name: "Similar name with matching identifiers",
				query: search.Entity[any]{
					Name: "NS LEADER II", // Similar but not exact
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						MMSI: "636013272",
					},
				},
				expected: 1.0,
			},
			{
				name: "Matching identifiers with different name",
				query: search.Entity[any]{
					Name: "Sea Transporter",
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						IMONumber: "9339301",
					},
				},
				expected: 1.0,
			},
			{
				name: "Similar vessel details but key mismatch",
				query: search.Entity[any]{
					Name: "NS LEADER",
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						Flag:     "Iran",     // wrong
						CallSign: "BTANK124", // other callsign
					},
				},
				expected: 0.431,
			},
			{
				name: "Complete mismatch",
				query: search.Entity[any]{
					Name: "GOLDEN FREIGHTER",
					Type: search.EntityVessel,
					Vessel: &search.Vessel{
						Name:     "GOLDEN FREIGHTER",
						Type:     search.VesselTypeCargo,
						Flag:     "LR",
						CallSign: "GOLD999",
					},
				},
				expected: 0.2104,
			},
			{
				name: "Wrong entity type",
				query: search.Entity[any]{
					Name: "BLUE TANKER",
					Type: search.EntityBusiness,
				},
				expected: 0.0,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				score := search.DebugSimilarity(debug(t), tc.query, indexEntity)
				require.InDelta(t, tc.expected, score, 0.02)

				// Additional assertions for specific score thresholds
				if tc.expected >= 0.95 {
					require.GreaterOrEqual(t, score, 0.95, "High confidence matches should score >= 0.95")
				}
				if tc.expected <= 0.40 {
					require.LessOrEqual(t, score, 0.40, "Clear mismatches should score <= 0.40")
				}
			})
		}
	})
}

func TestSimilarity_Edge_Cases(t *testing.T) {
	baseSDN := ofac.SDN{
		EntityID: "123",
		SDNName:  "TEST ENTITY",
		SDNType:  "vessel",
	}
	indexEntity := ofac.ToEntity(baseSDN, nil, nil, nil)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name:     "Empty query",
			query:    search.Entity[any]{},
			expected: 0,
		},
		{
			name: "Name only",
			query: search.Entity[any]{
				Name: "TEST ENTITY",
				Type: search.EntityVessel,
			},
			expected: 1.0,
		},
		{
			name: "Mismatched types",
			query: search.Entity[any]{
				Name: "TEST ENTITY",
				Type: search.EntityBusiness,
			},
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := search.DebugSimilarity(debug(t), tc.query, indexEntity)
			require.InDelta(t, tc.expected, score, 0.02)
		})
	}
}

func debug(t *testing.T) io.Writer {
	t.Helper()

	if testing.Verbose() {
		buf := new(bytes.Buffer)
		t.Cleanup(func() {
			fmt.Printf("\n%s\n", buf.String())
		})
		return buf
	}

	return nil
}
