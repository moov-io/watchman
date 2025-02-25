package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareBusinessExactIDs(t *testing.T) {
	cases := []struct {
		query, index Business
		expected     ScorePiece
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
			expected: ScorePiece{
				Score:          1.0,
				Weight:         50,
				Matched:        true,
				Required:       true,
				Exact:          true,
				FieldsCompared: 1,
				PieceType:      "identifiers",
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
