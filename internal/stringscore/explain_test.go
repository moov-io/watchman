package stringscore_test

import (
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"

	"github.com/stretchr/testify/require"
)

func TestExplainBestPairCombinationMatchesScore(t *testing.T) {
	t.Parallel()

	cfg := stringscore.DefaultScoringConfig()
	cases := [][2]string{
		{"john smith", "john smith"},
		{"JSC ARGUMENT", "JSCARGUMENT"},
		{"mohamed salem", "abdarrahmane ould mohamed el houcein ould mohamed salem"},
		{prepare.LowerAndRemovePunctuation("A.I.C. SOGO KENKYUSHO"), "sogo kenkyusho"},
		{"dmitry khoroshev", "dmitrii khoroshev"},
	}

	for _, tc := range cases {
		query := strings.Fields(strings.ToLower(tc[0]))
		index := strings.Fields(strings.ToLower(tc[1]))
		want := stringscore.BestPairCombinationJaroWinklerWithConfig(query, index, cfg)
		got, pairs := stringscore.ExplainBestPairCombinationJaroWinkler(query, index, cfg)
		require.InDelta(t, want, got, 1e-12, "%q vs %q", tc[0], tc[1])
		require.NotEmpty(t, pairs, "%q vs %q", tc[0], tc[1])
	}
}

func TestExplainBestPairCombinationIncludesUnmatchedTokens(t *testing.T) {
	t.Parallel()

	query := []string{"john", "smith"}
	index := []string{"john", "adams"}
	_, pairs := stringscore.ExplainBestPairCombinationJaroWinkler(query, index, stringscore.DefaultScoringConfig())

	var unmatchedIndex string
	for _, pair := range pairs {
		if pair.Search == "" {
			unmatchedIndex = pair.Index
		}
	}
	require.Equal(t, "adams", unmatchedIndex)
}

func TestExplainBestPairCombinationZeroScoreKeepsUnmatchedPairs(t *testing.T) {
	t.Parallel()

	cfg := stringscore.DefaultScoringConfig()
	query := []string{}
	index := []string{"smith"}
	want := stringscore.BestPairCombinationJaroWinklerWithConfig(query, index, cfg)
	got, pairs := stringscore.ExplainBestPairCombinationJaroWinkler(query, index, cfg)
	require.InDelta(t, want, got, 1e-12)
	require.InDelta(t, 0, got, 1e-12)
	require.NotEmpty(t, pairs)

	var unmatchedIndex string
	for _, pair := range pairs {
		if pair.Search == "" {
			unmatchedIndex = pair.Index
		}
	}
	require.Equal(t, "smith", unmatchedIndex)

	wantW := stringscore.BestPairCombinationJaroWinklerWeightedWithConfig(query, index, nil, []float64{1}, cfg)
	gotW, pairsW := stringscore.ExplainBestPairCombinationJaroWinklerWeighted(query, index, nil, []float64{1}, cfg)
	require.InDelta(t, wantW, gotW, 1e-12)
	require.NotEmpty(t, pairsW)
}

func TestExplainBestPairCombinationWeightedMatchesScore(t *testing.T) {
	t.Parallel()

	cfg := stringscore.DefaultScoringConfig()
	query := []string{"dmitry", "khoroshev"}
	index := []string{"dmitrii", "khoroshev"}
	qWeights := []float64{0.4, 0.6}
	iWeights := []float64{0.5, 0.5}

	want := stringscore.BestPairCombinationJaroWinklerWeightedWithConfig(query, index, qWeights, iWeights, cfg)
	got, pairs := stringscore.ExplainBestPairCombinationJaroWinklerWeighted(query, index, qWeights, iWeights, cfg)
	require.InDelta(t, want, got, 1e-12)
	require.NotEmpty(t, pairs)
}
