package embeddings

import (
	"context"

	lru "github.com/hashicorp/golang-lru/v2"
)

// memoryCache provides an LRU cache for embedding vectors.
// This reduces inference latency for repeated queries.
type memoryCache struct {
	cache    *lru.Cache[string, []float64]
	capacity int
}

var _ Cache = (&memoryCache{})

// newMemoryCache creates a new embedding cache with the specified size.
func newMemoryCache(size int) (*memoryCache, error) {
	if size <= 0 {
		size = 1000 // Default size
	}

	c, err := lru.New[string, []float64](size)
	if err != nil {
		return nil, err
	}

	return &memoryCache{
		cache:    c,
		capacity: size,
	}, nil
}

// Get retrieves an embedding from the cache.
// Returns the embedding and true if found, nil and false otherwise.
func (c *memoryCache) Get(_ context.Context, text string) ([]float64, bool) {
	return c.cache.Get(text)
}

// Put adds an embedding to the cache.
func (c *memoryCache) Put(_ context.Context, text string, embedding []float64) {
	// Make a copy to avoid external modifications
	emb := make([]float64, len(embedding))
	copy(emb, embedding)
	c.cache.Add(text, emb)
}

// Size returns the current number of items in the cache.
func (c *memoryCache) Size() int {
	return c.cache.Len()
}

// Clear removes all items from the cache.
func (c *memoryCache) Clear() {
	c.cache.Purge()
}

// Stats returns cache statistics.
type CacheStats struct {
	Size     int
	Capacity int
}

// Stats returns current cache statistics.
func (c *memoryCache) Stats() CacheStats {
	return CacheStats{
		Size:     c.cache.Len(),
		Capacity: c.capacity,
	}
}
