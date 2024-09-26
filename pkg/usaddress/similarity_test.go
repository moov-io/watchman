package usaddress

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressSimilarity(t *testing.T) {
	baseAddress := Address{
		PrimaryNumber: "123",
		StreetPredir:  "N",
		StreetName:    "MAIN",
		StreetSuffix:  "ST",
		StreetPostdir: "",
		SecondaryUnit: "APT 4B",
		City:          "ANYTOWN",
		State:         "CA",
		ZIPCode:       "90210",
		Plus4:         "1234",
	}

	tests := []struct {
		name     string
		addr1    Address
		addr2    Address
		expected float64
	}{
		{
			name:     "Identical Addresses",
			addr1:    baseAddress,
			addr2:    baseAddress,
			expected: 1.0, // Expect full similarity
		},
		{
			name:  "Different Primary Number",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.PrimaryNumber = "124"
				return a
			}(),
			expected: 0.956,
		},
		{
			name:  "Different Street Name",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.StreetName = "FLOWER"
				return a
			}(),
			expected: 0.75,
		},
		{
			name:  "Different ZIPCode",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.ZIPCode = "90211"
				return a
			}(),
			expected: 0.984,
		},
		{
			name:  "Different City",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.City = "CEDAR BEVERTON"
				return a
			}(),
			expected: 0.94,
		},
		{
			name:  "Different State",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.State = "NY"
				return a
			}(),
			expected: 0.95,
		},
		{
			name:  "Different Street Suffix",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.StreetSuffix = "AVE"
				return a
			}(),
			expected: 0.95,
		},
		{
			name:  "Different StreetPredir",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.StreetPredir = "S"
				return a
			}(),
			expected: 0.975,
		},
		{
			name:  "Different StreetPostdir",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.StreetPostdir = "NW"
				return a
			}(),
			expected: 0.975,
		},
		{
			name:  "Different Secondary Unit",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.SecondaryUnit = "APT 5C"
				return a
			}(),
			expected: 0.997,
		},
		{
			name:  "Different Plus4",
			addr1: baseAddress,
			addr2: func() Address {
				a := baseAddress
				a.Plus4 = "5678"
				return a
			}(),
			expected: 0.975,
		},
		{
			name: "POBox Instead of Street Address",
			addr1: func() Address {
				a := baseAddress
				a.POBox = "PO BOX 789"
				// Clear street address components
				a.PrimaryNumber = ""
				a.StreetName = ""
				a.StreetSuffix = ""
				a.StreetPredir = ""
				a.StreetPostdir = ""
				a.SecondaryUnit = ""
				return a
			}(),
			addr2: func() Address {
				a := baseAddress
				a.POBox = "PO BOX 789"
				// Clear street address components
				a.PrimaryNumber = ""
				a.StreetName = ""
				a.StreetSuffix = ""
				a.StreetPredir = ""
				a.StreetPostdir = ""
				a.SecondaryUnit = ""
				return a
			}(),
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			similarity := tt.addr1.Similarity(tt.addr2)
			require.InDelta(t, tt.expected, similarity, 0.01)
		})
	}
}
