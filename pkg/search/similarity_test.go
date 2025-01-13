package search_test

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_EdgeCases(t *testing.T) {
	var query search.Entity[search.Value]
	index := findOFACEntity(t, "47371")

	got := search.Similarity(query, index)
	require.InDelta(t, 0.0, got, 0.001)

}
