package embeddings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmbeddingCache_PutGet(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	embedding := []float64{0.1, 0.2, 0.3}
	cache.Put("test", embedding)

	retrieved, ok := cache.Get("test")
	require.True(t, ok)
	require.Equal(t, embedding, retrieved)
}

func TestEmbeddingCache_Miss(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	_, ok := cache.Get("nonexistent")
	require.False(t, ok)
}

func TestEmbeddingCache_Eviction(t *testing.T) {
	// Small cache size
	cache, err := newCache(2)
	require.NoError(t, err)

	cache.Put("a", []float64{1})
	cache.Put("b", []float64{2})
	cache.Put("c", []float64{3}) // Should evict "a"

	_, ok := cache.Get("a")
	require.False(t, ok, "Expected 'a' to be evicted")

	_, ok = cache.Get("b")
	require.True(t, ok)

	_, ok = cache.Get("c")
	require.True(t, ok)
}

func TestEmbeddingCache_Size(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	require.Equal(t, 0, cache.Size())

	cache.Put("a", []float64{1})
	require.Equal(t, 1, cache.Size())

	cache.Put("b", []float64{2})
	require.Equal(t, 2, cache.Size())
}

func TestEmbeddingCache_Clear(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	cache.Put("a", []float64{1})
	cache.Put("b", []float64{2})
	require.Equal(t, 2, cache.Size())

	cache.Clear()
	require.Equal(t, 0, cache.Size())
}

func TestEmbeddingCache_CopyOnPut(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	original := []float64{1, 2, 3}
	cache.Put("test", original)

	// Modify original
	original[0] = 999

	// Retrieved should be unchanged
	retrieved, ok := cache.Get("test")
	require.True(t, ok)
	require.Equal(t, float64(1), retrieved[0], "Cache should store a copy, not a reference")
}

func TestEmbeddingCache_DefaultSize(t *testing.T) {
	// Size 0 should use default
	cache, err := newCache(0)
	require.NoError(t, err)
	require.NotNil(t, cache)
}

func TestEmbeddingCache_Stats(t *testing.T) {
	cache, err := newCache(100)
	require.NoError(t, err)

	cache.Put("a", []float64{1})
	cache.Put("b", []float64{2})

	stats := cache.Stats()
	require.Equal(t, 2, stats.Size)
}
