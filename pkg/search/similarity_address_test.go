package search

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/moov-io/watchman/internal/norm"

	"github.com/stretchr/testify/require"
)

func TestCompareAddress(t *testing.T) {
	tests := []struct {
		name     string
		query    Address
		index    Address
		expected float64
	}{
		{
			name: "only_line1_exact",
			query: Address{
				Line1: "123 Main St",
			},
			index: Address{
				Line1: "123 Main St",
			},
			expected: 1.0,
		},
		{
			name: "only_line1_close",
			query: Address{
				Line1: "123 Main Street",
			},
			index: Address{
				Line1: "123 Main St",
			},
			expected: 0.876, // High but not exact due to Street vs St
		},
		{
			name: "only_line1_different_number",
			query: Address{
				Line1: "124 Main St", // Different building number
			},
			index: Address{
				Line1: "123 Main St",
			},
			expected: 0.941,
		},
		{
			name: "only_city",
			query: Address{
				City: "New York",
			},
			index: Address{
				City: "New York",
			},
			expected: 1.0,
		},
		{
			name: "similar_cities",
			query: Address{
				City: "Los Angeles",
			},
			index: Address{
				City: "Los Angles", // Common misspelling
			},
			expected: 0.958,
		},
		{
			name: "only_postal",
			query: Address{
				PostalCode: "90210",
			},
			index: Address{
				PostalCode: "90210",
			},
			expected: 1.0,
		},
		{
			name: "similar_addresses_different_units",
			query: Address{
				Line1: "123 Main St Apt 4B",
				City:  "New York",
				State: "NY",
			},
			index: Address{
				Line1: "123 Main St Apt 4C", // Different unit
				City:  "New York",
				State: "NY",
			},
			expected: 0.987,
		},
		{
			name: "similar_addresses_different_line2",
			query: Address{
				Line1: "123 Main St",
				Line2: "Apt 4B",
				City:  "New York",
				State: "NY",
			},
			index: Address{
				Line1: "123 Main St",
				Line2: "Apt 4C", // Different unit
				City:  "New York",
				State: "NY",
			},
			expected: 0.9877,
		},
		{
			name: "country_code_vs_name",
			query: Address{
				Country: "United States",
			},
			index: Address{
				Country: "US",
			},
			expected: 1.0,
		},
		{
			name: "complex_partial_match",
			query: Address{
				Line1:      "1234 Broadway Suite 500",
				City:       "New York",
				State:      "NY",
				PostalCode: "10013",
				Country:    "US",
			},
			index: Address{
				Line1:      "1234 Broadway",
				City:       "New York",
				PostalCode: "10013",
			},
			expected: 0.8958,
		},
		{
			name: "tricky_similar_but_different",
			query: Address{
				Line1: "45 Park Avenue South",
				City:  "New York",
				State: "NY",
			},
			index: Address{
				Line1: "45 Park Avenue North", // Different direction
				City:  "New York",
				State: "NY",
			},
			expected: 0.905,
		},
		{
			name: "ambiguous_addresses",
			query: Address{
				Line1: "100 Washington St",
				City:  "Boston",
			},
			index: Address{
				Line1: "100 Washington Ave", // Different street type
				City:  "Boston",
			},
			expected: 0.9644,
		},
		{
			name: "missing_fields_comparison",
			query: Address{
				Line1: "555 Market St",
				City:  "San Francisco",
			},
			index: Address{
				Line1:   "555 Market St",
				Line2:   "Floor 2",
				City:    "San Francisco",
				State:   "CA",
				Country: "US",
			},
			expected: 1.0, // Should still be high as all provided fields match
		},
		{
			// See https://github.com/moov-io/watchman/issues/625
			name: "tricky half address",
			query: Address{
				Line1: "1/a block 2 gulshan-e-iqbal",
			},
			index: Address{
				Line1: "ST 1/A, Block 2, Gulshan-e-Iqbal",
			},
			expected: 0.9885,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			score := compareAddress(&buf, normalizeAddress(tt.query), normalizeAddress(tt.index))

			if testing.Verbose() {
				fmt.Println(buf.String())
			}

			require.InDelta(t, tt.expected, score, 0.001, "addresses should have expected similarity score (got %.2f, want %.2f)", score, tt.expected)
		})
	}
}

func TestCompareAddress_Normalized(t *testing.T) {
	tests := []struct {
		name     string
		query    Address
		index    Address
		expected float64
	}{
		{
			name: "country_code_vs_name",
			query: Address{
				Country: "United States",
			},
			index: Address{
				Country: "US",
			},
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Normalize country (like callers of compareAddress do)
			tt.query.Country = norm.Country(tt.query.Country)
			tt.index.Country = norm.Country(tt.index.Country)

			var buf bytes.Buffer
			score := compareAddress(&buf, normalizeAddress(tt.query), normalizeAddress(tt.index))

			if testing.Verbose() {
				fmt.Println(buf.String())
			}

			require.InDelta(t, tt.expected, score, 0.001,
				"addresses should have expected similarity score (got %.2f, want %.2f)", score, tt.expected)
		})
	}
}

func TestCompareAddressesNoMatch(t *testing.T) {
	var buf bytes.Buffer

	tests := []struct {
		name     string
		query    Address
		index    Address
		expected float64
	}{
		{
			name: "completely_different_addresses",
			query: Address{
				Line1: "123 Main St",
				City:  "Boston",
				State: "MA",
			},
			index: Address{
				Line1: "456 Oak Ave",
				City:  "Chicago",
				State: "IL",
			},
			expected: 0.0,
		},
		{
			name: "similar_looking_but_different",
			query: Address{
				Line1: "1 World Trade Center",
				City:  "New York",
				State: "NY",
			},
			index: Address{
				Line1: "2 World Trade Center", // Different building
				City:  "New York",
				State: "NY",
			},
			expected: 0.982,
		},
		{
			name: "transposed_numbers",
			query: Address{
				Line1: "123 Main St",
				City:  "Anytown",
			},
			index: Address{
				Line1: "321 Main St",
				City:  "Anytown",
			},
			expected: 0.867,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := compareAddress(&buf, normalizeAddress(tt.query), normalizeAddress(tt.index))
			require.InDelta(t, tt.expected, score, 0.001, "different addresses should have low similarity score: %.2f", score)
		})
	}
}
