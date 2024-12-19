//go:build libpostal

package address

import (
	"fmt"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	postal "github.com/openvenues/gopostal/parser"
	"github.com/stretchr/testify/require"
)

func TestParseAddress(t *testing.T) {
	cases := []struct {
		input    string
		expected search.Address
	}{
		{
			input: "101 Maple Street Apt 202 Bigcity, New York 11222",
			expected: search.Address{
				Line1:      "101 maple street",
				Line2:      "apt 202",
				City:       "bigcity",
				PostalCode: "11222",
				State:      "new york",
			},
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%#v", tc.expected)

		t.Run(name, func(t *testing.T) {
			got := ParseAddress(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}

func TestOrganizeLibpostalComponents(t *testing.T) {
	cases := []struct {
		parts    []postal.ParsedComponent
		expected search.Address
	}{
		{
			parts: []postal.ParsedComponent{
				{Label: "house_number", Value: "101"},
				{Label: "road", Value: "Main Street"},
				{Label: "city", Value: "Springfield"},
				{Label: "state", Value: "Illinois"},
				{Label: "postcode", Value: "62704"},
				{Label: "country", Value: "United States"},
			},
			expected: search.Address{
				Line1:      "101 Main Street",
				City:       "Springfield",
				PostalCode: "62704",
				State:      "Illinois",
				Country:    "United States",
			},
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%#v", tc.expected)

		t.Run(name, func(t *testing.T) {
			got := organizeLibpostalComponents(tc.parts)
			require.Equal(t, tc.expected, got)
		})
	}
}
