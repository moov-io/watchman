package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareBusinessExactIDs(t *testing.T) {
	cases := []struct {
		query, index Business
		expected     scorePiece
	}{
		{
			query: Business{
				GovernmentIDs: []GovernmentID{
					{Type: GovernmentIDTax, Country: "United States", Identifier: "522083095"},
				},
			},
			index: Business{
				GovernmentIDs: []GovernmentID{
					{Type: GovernmentIDTax, Country: "United States", Identifier: "52-2083095"},
				},
			},
			expected: scorePiece{
				score:          1.0,
				weight:         50,
				matched:        true,
				required:       true,
				exact:          true,
				fieldsCompared: 1,
				pieceType:      "identifiers",
			},
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%#v", tc.query.GovernmentIDs), func(t *testing.T) {
			got := compareBusinessExactIDs(nil, &tc.query, &tc.index, criticalIdWeight)
			require.Equal(t, tc.expected, got)
		})
	}
}
