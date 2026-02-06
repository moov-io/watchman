package embeddings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVectorIndex_AddAndSearch(t *testing.T) {
	idx := newVectorIndex(3)

	// Add some test vectors (normalized)
	vectors := [][]float64{
		{1, 0, 0},     // vec0
		{0, 1, 0},     // vec1
		{0, 0, 1},     // vec2
		{0.7, 0.7, 0}, // vec3 - similar to vec0 and vec1
	}
	ids := []string{"id0", "id1", "id2", "id3"}
	names := []string{"name0", "name1", "name2", "name3"}

	idx.Add(vectors, ids, names)

	require.Equal(t, 4, idx.Size())

	// Search for similar vectors
	query := []float64{1, 0, 0} // Should be most similar to vec0
	results := idx.Search(query, 2)

	require.Len(t, results, 2)
	require.Equal(t, "id0", results[0].ID) // Best match
	require.Equal(t, "name0", results[0].Name)
	require.InDelta(t, 1.0, results[0].Score, 0.001) // Perfect match
}

func TestVectorIndex_SearchTopK(t *testing.T) {
	idx := newVectorIndex(2)

	// Add vectors
	vectors := [][]float64{
		{1, 0},
		{0.9, 0.1},
		{0.8, 0.2},
		{0.7, 0.3},
		{0.6, 0.4},
	}
	ids := []string{"a", "b", "c", "d", "e"}
	names := []string{"A", "B", "C", "D", "E"}

	idx.Add(vectors, ids, names)

	// Request k=3
	query := []float64{1, 0}
	results := idx.Search(query, 3)

	require.Len(t, results, 3)
	// Results should be sorted by score descending
	require.GreaterOrEqual(t, results[0].Score, results[1].Score)
	require.GreaterOrEqual(t, results[1].Score, results[2].Score)
}

func TestVectorIndex_EmptyIndex(t *testing.T) {
	idx := newVectorIndex(3)

	query := []float64{1, 0, 0}
	results := idx.Search(query, 5)

	require.Empty(t, results)
}

func TestVectorIndex_SearchKLargerThanIndex(t *testing.T) {
	idx := newVectorIndex(2)

	vectors := [][]float64{{1, 0}, {0, 1}}
	ids := []string{"a", "b"}
	names := []string{"A", "B"}

	idx.Add(vectors, ids, names)

	// Request more than available
	results := idx.Search([]float64{1, 0}, 10)

	require.Len(t, results, 2) // Should return all available
}

func TestVectorIndex_Clear(t *testing.T) {
	idx := newVectorIndex(2)

	vectors := [][]float64{{1, 0}, {0, 1}}
	idx.Add(vectors, []string{"a", "b"}, []string{"A", "B"})

	require.Equal(t, 2, idx.Size())

	idx.Clear()

	require.Equal(t, 0, idx.Size())
}

func TestVectorIndex_GetVector(t *testing.T) {
	idx := newVectorIndex(2)

	vectors := [][]float64{{1, 0}, {0, 1}}
	idx.Add(vectors, []string{"a", "b"}, []string{"A", "B"})

	vec := idx.GetVector("a")
	require.NotNil(t, vec)
	require.Equal(t, []float64{1, 0}, vec)

	vec = idx.GetVector("nonexistent")
	require.Nil(t, vec)
}

func TestDotProduct(t *testing.T) {
	tests := []struct {
		name     string
		a        []float64
		b        []float64
		expected float64
	}{
		{"orthogonal vectors", []float64{1, 0}, []float64{0, 1}, 0.0},
		{"identical vectors", []float64{1, 0}, []float64{1, 0}, 1.0},
		{"opposite vectors", []float64{1, 0}, []float64{-1, 0}, -1.0},
		{"partial overlap", []float64{0.7, 0.7}, []float64{1, 0}, 0.7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := dotProduct(tc.a, tc.b)
			require.InDelta(t, tc.expected, result, 0.01)
		})
	}
}

func TestNormalizeL2(t *testing.T) {
	vec := []float64{3, 4} // 3-4-5 triangle
	normalized := normalizeL2(vec)

	// Check magnitude is 1
	var mag float64
	for _, v := range normalized {
		mag += float64(v * v)
	}
	require.InDelta(t, 1.0, mag, 0.0001)

	// Check proportions preserved
	require.InDelta(t, 0.6, normalized[0], 0.001)
	require.InDelta(t, 0.8, normalized[1], 0.001)
}

func TestNormalizeL2_ZeroVector(t *testing.T) {
	vec := []float64{0, 0, 0}
	normalized := normalizeL2(vec)

	// Should return unchanged (avoid division by zero)
	require.Equal(t, vec, normalized)
}
