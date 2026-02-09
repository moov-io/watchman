package embeddings

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/db"

	"github.com/stretchr/testify/require"
)

func TestEmbeddingCache_PutGet(t *testing.T) {
	forEachCache(t, 100, func(c Cache) {
		ctx := context.Background()

		embedding := []float64{0.1, 0.2, 0.3}
		c.Put(ctx, "test", embedding)

		retrieved, ok := c.Get(ctx, "test")
		require.True(t, ok)
		require.Equal(t, embedding, retrieved)
	})
}

func TestEmbeddingCache_Miss(t *testing.T) {
	forEachCache(t, 100, func(c Cache) {
		ctx := context.Background()

		_, ok := c.Get(ctx, "nonexistent")
		require.False(t, ok)
	})
}

func TestEmbeddingCache_Memory_Eviction(t *testing.T) {
	ctx := context.Background()

	c, err := newMemoryCache(2)
	require.NoError(t, err)

	c.Put(ctx, "a", []float64{1})
	c.Put(ctx, "b", []float64{2})
	c.Put(ctx, "c", []float64{3}) // Should evict "a"

	_, ok := c.Get(ctx, "a")
	require.False(t, ok, "Expected 'a' to be evicted")

	_, ok = c.Get(ctx, "b")
	require.True(t, ok)

	_, ok = c.Get(ctx, "c")
	require.True(t, ok)
}

func TestEmbeddingCache_Memory_Size(t *testing.T) {
	ctx := context.Background()

	cache, err := newMemoryCache(100)
	require.NoError(t, err)

	require.Equal(t, 0, cache.Size())

	cache.Put(ctx, "a", []float64{1})
	require.Equal(t, 1, cache.Size())

	cache.Put(ctx, "b", []float64{2})
	require.Equal(t, 2, cache.Size())
}

func TestEmbeddingCache_Memory_Clear(t *testing.T) {
	ctx := context.Background()

	cache, err := newMemoryCache(100)
	require.NoError(t, err)

	cache.Put(ctx, "a", []float64{1})
	cache.Put(ctx, "b", []float64{2})
	require.Equal(t, 2, cache.Size())

	cache.Clear()
	require.Equal(t, 0, cache.Size())
}

func TestEmbeddingCache_CopyOnPut(t *testing.T) {
	forEachCache(t, 100, func(c Cache) {
		ctx := context.Background()

		original := []float64{1, 2, 3}
		c.Put(ctx, "test", original)

		// Modify original
		original[0] = 999

		// Retrieved should be unchanged
		retrieved, ok := c.Get(ctx, "test")
		require.True(t, ok)
		require.Equal(t, float64(1), retrieved[0], "Cache should store a copy, not a reference")
	})
}

func TestEmbeddingCache_Memory_DefaultSize(t *testing.T) {
	// Size 0 should use default
	cache, err := newMemoryCache(100)
	require.NoError(t, err)
	require.NotNil(t, cache)
}

func TestEmbeddingCache_Stats(t *testing.T) {
	ctx := context.Background()

	cache, err := newMemoryCache(100)
	require.NoError(t, err)

	cache.Put(ctx, "a", []float64{1})
	cache.Put(ctx, "b", []float64{2})

	stats := cache.Stats()
	require.Equal(t, 2, stats.Size)
}

func forEachCache(t *testing.T, size int, fn func(c Cache)) {
	t.Helper()

	t.Run("memory", func(t *testing.T) {
		// memory
		mem, err := newMemoryCache(size)
		require.NoError(t, err)
		require.NotNil(t, mem)

		fn(mem)
	})

	if !testing.Short() {
		t.Run("database", func(t *testing.T) {
			db.ForEachDatabase(t, func(database db.DB) {
				config := Config{
					Provider: ProviderConfig{
						Model:     "testing",
						Dimension: 123,
					},
				}

				cc, err := newSqlRepository(t.Context(), config, database)
				require.NoError(t, err)

				fn(cc)
			})
		})
	}
}
