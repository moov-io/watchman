package stringscore

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJaroWinklerAlgorithms(t *testing.T) {
	cases := []struct {
		query, index string

		bestPairsCombination  float64
		bestPairsJaroWinkler  float64
		jaroWinkler           float64
		jaroWinklerFavoritism float64
	}{
		{
			query:                 "st 1/a, block 2, gulshan-e-iqbal",
			index:                 "1/a, block 2, gulshan-e-iqbal",
			bestPairsCombination:  0.963,
			bestPairsJaroWinkler:  0.963,
			jaroWinkler:           0.8,
			jaroWinklerFavoritism: 0.8,
		},
	}
	for _, tc := range cases {
		t.Run(tc.query, func(t *testing.T) {
			// Run each algorithm
			queryTokens := strings.Fields(tc.query)
			indexTokens := strings.Fields(tc.index)

			t.Run("best pairs combination", func(t *testing.T) {
				got := BestPairCombinationJaroWinkler(queryTokens, indexTokens)
				require.InDelta(t, tc.bestPairsCombination, got, 0.001)
			})

			t.Run("best pairs jaro winkler", func(t *testing.T) {
				got := BestPairsJaroWinkler(queryTokens, indexTokens)
				require.InDelta(t, tc.bestPairsJaroWinkler, got, 0.001)
			})

			t.Run("jaro winkler", func(t *testing.T) {
				got := JaroWinkler(tc.query, tc.index)
				require.InDelta(t, tc.jaroWinkler, got, 0.001)
			})

			t.Run("jaro winkler with favoritism", func(t *testing.T) {
				got := JaroWinklerWithFavoritism(tc.query, tc.index, exactMatchFavoritism)
				require.InDelta(t, tc.jaroWinklerFavoritism, got, 0.001)
			})
		})
	}
}
