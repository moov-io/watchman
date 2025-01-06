package search_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_OFAC_SDN_Vessel(t *testing.T) {
	baseSDN := ofac.SDN{
		EntityID:               "123",
		SDNName:                "BLUE TANKER",
		SDNType:                "vessel",
		Programs:               []string{"SDGT", "IRGC"},
		CallSign:               "BTANK123",
		VesselType:             "Cargo",
		Tonnage:                "15000",
		GrossRegisteredTonnage: "18000",
		VesselFlag:             "PA",
		VesselOwner:            "GLOBAL SHIPPING CORP",
		Remarks:                "Known aliases: Sea Transporter, Ocean Carrier",
	}

	// Create the base entity to match against
	indexEntity := ofac.ToEntity(baseSDN,
		[]ofac.Address{
			{
				AddressID:                   "addr1",
				Address:                     "123 Harbor Drive",
				CityStateProvincePostalCode: "Panama City 12345",
				Country:                     "PA",
			},
		},
		[]ofac.SDNComments{},
		[]ofac.AlternateIdentity{
			{
				AlternateID:   "alt1",
				AlternateType: "Vessel Registration",
				AlternateName: "REG123456",
			},
		},
	)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Exact match - All fields",
			query: search.Entity[any]{
				Name: "BLUE TANKER",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:                   "BLUE TANKER",
					Type:                   search.VesselTypeCargo,
					Flag:                   "PA",
					Tonnage:                15000,
					CallSign:               "BTANK123",
					GrossRegisteredTonnage: 18000,
					Owner:                  "GLOBAL SHIPPING CORP",
					IMONumber:              "IMO123456",
				},
				Addresses: []search.Address{
					{
						Line1:      "123 Harbor Drive",
						City:       "Panama City",
						Country:    "PA",
						PostalCode: "12345",
					},
				},
			},
			expected: 1.0,
		},
		{
			name: "High confidence - Key fields match",
			query: search.Entity[any]{
				Name: "BLUE TANKER",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:      "BLUE TANKER",
					Type:      search.VesselTypeCargo,
					IMONumber: "IMO123456",
					CallSign:  "BTANK123",
				},
			},
			expected: 1.0,
		},
		{
			name: "Similar name with matching identifiers",
			query: search.Entity[any]{
				Name: "Blue Tanker II", // Similar but not exact
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:      "Blue Tanker II",
					CallSign:  "BTANK123", // Exact match
					IMONumber: "IMO123456",
					Type:      search.VesselTypeCargo,
					Flag:      "PA",
				},
			},
			expected: 1.0,
		},
		{
			name: "Matching identifiers with different name",
			query: search.Entity[any]{
				Name: "Sea Transporter", // Known alias
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:     "Sea Transporter",
					CallSign: "BTANK123",
					Type:     search.VesselTypeCargo,
				},
			},
			expected: 0.814,
		},
		{
			name: "Similar vessel details but key mismatch",
			query: search.Entity[any]{
				Name: "BLUE TANKER",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:     "BLUE TANKER",
					Type:     search.VesselTypeCargo,
					Flag:     "PA",
					CallSign: "BTANK124", // One digit off
					Tonnage:  14800,      // Close but not exact
				},
			},
			expected: 0.431,
		},
		{
			name: "Partial info with some matches",
			query: search.Entity[any]{
				Name: "BLUE TANKER",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:  "BLUE TANKER",
					Flag:  "PA",
					Owner: "GLOBAL SHIPPING CORP",
				},
			},
			expected: 1.0,
		},
		{
			name: "Different vessel with similar name",
			query: search.Entity[any]{
				Name: "BLUE TANKER STAR",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:     "BLUE TANKER STAR",
					Type:     search.VesselTypeCargo,
					Flag:     "SG", // Different flag
					CallSign: "BSTAR789",
					Owner:    "STAR SHIPPING LTD",
				},
			},
			expected: 0.36,
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
			expected: 0.154,
		},
		{
			name: "Wrong entity type",
			query: search.Entity[any]{
				Name: "BLUE TANKER",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "BLUE TANKER",
				},
			},
			expected: 0.667,
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
}

func TestSimilarity_OFAC_SDN_Person(t *testing.T) {
	// Create a base SDN that represents a Person
	baseSDN := ofac.SDN{
		EntityID: "999",
		SDNName:  "JOHN SMITH",
		SDNType:  "individual", // or "person", depending on your data
		Title:    "MR",
		Remarks:  "Some remarks about JOHN SMITH; Gender Male; DOB 05 Jan 1959",
	}

	// For demonstration, we can embed a "DOB" note in Remarks or store it in an address,
	// but your actual logic might parse or store it differently. We'll keep it simple.
	sdnAddress := ofac.Address{
		EntityID:                    "999",
		AddressID:                   "addr-person-1",
		Address:                     "1234 MAIN ST",
		CityStateProvincePostalCode: "LOS ANGELES CA 90001",
		Country:                     "UNITED STATES",
		AddressRemarks:              "", // Example way to store
	}

	// Alternate identity or name
	altIdentity := ofac.AlternateIdentity{
		EntityID:         "999",
		AlternateID:      "alt-person-1",
		AlternateType:    "fka",
		AlternateName:    "JONATHAN SMITH",
		AlternateRemarks: "Formerly known as Jonathan",
	}

	// Convert this SDN into your internal Entity
	indexEntity := ofac.ToEntity(baseSDN, []ofac.Address{sdnAddress}, []ofac.SDNComments{}, []ofac.AlternateIdentity{altIdentity})

	// For the Person-specific comparison logic to work (compare birth date, gender, titles, etc.),
	// we’ll fill out the Person struct in the query. The watchman code’s Person comparison
	// checks for exact birth date matches, gender, titles, etc.

	// Let's define a sample time to represent the person's birthday
	dob := time.Date(1959, time.January, 5, 10, 32, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Exact match - Person with all fields",
			query: search.Entity[any]{
				Name: "JOHN SMITH",
				Type: search.EntityPerson,
				Person: &search.Person{
					BirthDate: &dob,
					Gender:    "male",
					Titles:    []string{"MR"},
				},
			},
			expected: 1.0,
		},
		{
			name: "Partial match - Missing birthdate",
			query: search.Entity[any]{
				Name: "JOHN SMITH",
				Type: search.EntityPerson,
				Person: &search.Person{
					// No BirthDate provided
					Gender: "male",
					Titles: []string{"MR"},
				},
			},
			expected: 1.0,
		},
		{
			name: "Fuzzy name match - Alternate identity",
			query: search.Entity[any]{
				Name: "Jonathan Smith",
				Type: search.EntityPerson,
				Person: &search.Person{
					Titles: []string{"MR"},
				},
			},
			// Expect a high score, but maybe slightly less than perfect
			// if "Jonathan" vs. "John" is considered a close fuzzy match or if alt name is recognized.
			expected: 0.667,
		},
		{
			name: "Close name but different person details",
			query: search.Entity[any]{
				Name: "JOHN SMYTH", // Slightly different spelling
				Type: search.EntityPerson,
				Person: &search.Person{
					// Different birth date
					BirthDate: func() *time.Time {
						d := time.Date(1975, 1, 1, 0, 0, 0, 0, time.UTC)
						return &d
					}(),
					Gender: "male",
					Titles: []string{"MR"},
				},
			},
			expected: 0.7938,
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
			expected: 0.4559,
		},
		{
			name: "Wrong entity type",
			query: search.Entity[any]{
				Name: "JOHN SMITH",
				Type: search.EntityVessel, // intentionally vessel to mismatch
			},
			expected: 0.667,
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
	// Create an OFAC SDN that represents a Business
	baseSDN := ofac.SDN{
		EntityID: "B-101",
		SDNName:  "ACME Corporation",
		SDNType:  "",
		Remarks:  "Organization Established Date 15 Feb 2010; Registration Number 12345;",
		// (Optionally add more fields, e.g. Programs, etc.)
	}

	// We'll define an Address or AlternateIdentity if needed
	sdnAddr := ofac.Address{
		EntityID:  "B-101",
		AddressID: "addr-bus-1",
		Address:   "1000 Industrial Way",
		Country:   "US",
		// ...
	}

	// For demonstration, a single AlternateIdentity for the business (like a DBA name)
	altID := ofac.AlternateIdentity{
		EntityID:      "B-101",
		AlternateID:   "alt-bus-1",
		AlternateType: "DBA",
		AlternateName: "ACME Co",
	}

	// Convert this SDN into your internal Entity (which should set Type=Business,
	// and fill out Business{Name, AltNames, etc.})
	// The details here depend on your ofac.ToEntity(...) implementation.
	indexEntity := ofac.ToEntity(baseSDN, []ofac.Address{sdnAddr}, nil, []ofac.AlternateIdentity{altID})

	// For this example, let's assume your ofac.ToEntity() sets:
	//   indexEntity.Type = search.EntityBusiness
	//   indexEntity.Business = &search.Business{
	//       Name:       "ACME Corporation",
	//       AltNames:   []string{"ACME Co"},
	//       Created:    &time.Date(2010, time.February, 15, ...),
	//       Dissolved:  nil,
	//       Identifier: []search.Identifier{
	//           { Name: "Registration", Identifier: "12345", Country: "US" },
	//       },
	//   }
	//
	// Adjust your actual mapping logic accordingly.

	// We'll define a Created date to match the "Organization Established Date" data
	businessCreatedAt := time.Date(2010, time.February, 15, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Exact match - All fields",
			query: search.Entity[any]{
				Name: "ACME Corporation",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name:     "ACME Corporation",
					AltNames: []string{"ACME Co"},
					Created:  &businessCreatedAt,
					Identifier: []search.Identifier{
						{
							Name:       "Registration",
							Country:    "US",
							Identifier: "12345",
						},
					},
				},
			},
			// If everything lines up exactly, we expect near 1.0
			expected: 1.0,
		},
		{
			name: "Partial match - Fuzzy name, same identifier",
			query: search.Entity[any]{
				Name: "ACME Corp",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "ACME Corp", // Fuzzy match to "ACME Corporation"
					Identifier: []search.Identifier{
						{
							Name:       "Registration Number",
							Identifier: "12345", // exact ID match
						},
					},
				},
			},
			// If your fuzzy logic sees "ACME Inc." ~ "ACME Corporation" ~0.8 or so,
			// plus an exact ID match, you might get something ~0.9 or 0.95
			expected: 0.8438,
		},
		{
			name: "Alt name only + missing ID",
			query: search.Entity[any]{
				Name: "ACME Co", // matches alt name
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "ACME Co",
					// No identifier
				},
			},
			// Expect a moderate score (fuzzy match or alt name success),
			// but not 1.0 since the ID is missing, or other fields not matched.
			expected: 0.7,
		},
		{
			name: "Different name, different ID",
			query: search.Entity[any]{
				Name: "BETA Solutions",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "BETA Solutions",
					Identifier: []search.Identifier{
						{Name: "Registration", Country: "US", Identifier: "99999"},
					},
				},
			},
			// Expect a low score for mismatch
			expected: 0.3347,
		},
		{
			name: "Wrong entity type",
			query: search.Entity[any]{
				Name: "ACME Corporation",
				Type: search.EntityVessel,
			},
			// Mismatched types typically yield a near-zero or a partial coverage score
			expected: 0.6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := search.DebugSimilarity(debug(t), tc.query, indexEntity)
			require.InDelta(t, tc.expected, score, 0.05)
		})
	}
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
			},
			expected: 0.667,
		},
		{
			name: "Mismatched types",
			query: search.Entity[any]{
				Name: "TEST ENTITY",
				Type: search.EntityBusiness,
			},
			expected: 0.667,
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
