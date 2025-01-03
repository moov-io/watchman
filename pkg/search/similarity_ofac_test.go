package search_test

import (
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_OFAC_SDN_Vessel(t *testing.T) {
	fullSDN := ofac.SDN{
		EntityID:               "123",
		SDNName:                "TANKER VESSEL",
		SDNType:                "vessel",
		Programs:               []string{"SDGT", "IRGC"},
		CallSign:               "ABCD1234",
		VesselType:             "Cargo",
		Tonnage:                "10000",
		GrossRegisteredTonnage: "12000",
		VesselFlag:             "US",
		VesselOwner:            "BIG SHIPPING INC.",
		Remarks:                "Test remarks",
	}

	indexEntity := ofac.ToEntity(fullSDN, nil, nil)

	testCases := []struct {
		name     string
		query    search.Entity[any]
		expected float64
	}{
		{
			name: "Perfect match",
			query: search.Entity[any]{
				Name: "TANKER VESSEL",
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:                   "TANKER VESSEL",
					Type:                   search.VesselType("Cargo"),
					Flag:                   "US",
					Tonnage:                10000,
					CallSign:               "ABCD1234",
					GrossRegisteredTonnage: 12000,
					Owner:                  "BIG SHIPPING INC.",
				},
			},
			expected: 0.80,
		},
		{
			name: "Partial match (some fields differ)",
			query: search.Entity[any]{
				Name: "Tanker Vessel", // close match (capitalization differs)
				Type: search.EntityVessel,
				Vessel: &search.Vessel{
					Name:                   "Tanker Vessel",
					Type:                   search.VesselType("Cargo"),
					Flag:                   "GB", // mismatch
					Tonnage:                9500, // partial mismatch
					CallSign:               "ABCD1234",
					GrossRegisteredTonnage: 12000,
					Owner:                  "BIG SHIPPING Inc", // minor difference
				},
			},
			expected: 0.75,
		},
		{
			name: "Mismatch (completely different)",
			query: search.Entity[any]{
				Name: "Random Business",
				Type: search.EntityBusiness,
				Business: &search.Business{
					Name: "Random Business",
				},
			},
			expected: 0.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := search.Similarity(tc.query, indexEntity)
			require.InDelta(t, tc.expected, score, 0.001)
		})
	}
}
