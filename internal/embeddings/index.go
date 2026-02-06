package embeddings

import (
	"sort"
	"sync"
)

// vectorIndex provides approximate nearest neighbor search for embeddings.
// It uses a simple brute-force approach that is sufficient for sanctions lists
// (typically 10,000-50,000 entities).
//
// For larger datasets, consider using a more sophisticated index like HNSW.
type vectorIndex struct {
	mu      sync.RWMutex
	vectors [][]float64
	ids     []string
	names   []string
	dim     int
}

// newVectorIndex creates a new empty vector index.
func newVectorIndex(dim int) *vectorIndex {
	return &vectorIndex{
		dim:     dim,
		vectors: make([][]float64, 0),
		ids:     make([]string, 0),
		names:   make([]string, 0),
	}
}

// Add adds vectors to the index with their corresponding IDs and names.
func (idx *vectorIndex) Add(vectors [][]float64, ids, names []string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.vectors = append(idx.vectors, vectors...)
	idx.ids = append(idx.ids, ids...)
	idx.names = append(idx.names, names...)
}

// Search finds the k most similar vectors to the query.
// Returns results sorted by similarity score (highest first).
func (idx *vectorIndex) Search(query []float64, k int) []SearchResult {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if len(idx.vectors) == 0 {
		return nil
	}

	// Compute similarity scores for all vectors
	type scored struct {
		idx   int
		score float64
	}

	scores := make([]scored, len(idx.vectors))
	for i, vec := range idx.vectors {
		scores[i] = scored{
			idx:   i,
			score: dotProduct(query, vec),
		}
	}

	// Sort by score descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	// Return top k results
	if k > len(scores) {
		k = len(scores)
	}

	results := make([]SearchResult, k)
	for i := 0; i < k; i++ {
		s := scores[i]
		results[i] = SearchResult{
			ID:    idx.ids[s.idx],
			Name:  idx.names[s.idx],
			Score: s.score,
		}
	}

	return results
}

// Size returns the number of vectors in the index.
func (idx *vectorIndex) Size() int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return len(idx.vectors)
}

// Clear removes all vectors from the index.
func (idx *vectorIndex) Clear() {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.vectors = make([][]float64, 0)
	idx.ids = make([]string, 0)
	idx.names = make([]string, 0)
}

// GetVector returns the embedding vector for a given ID.
// Returns nil if the ID is not found.
func (idx *vectorIndex) GetVector(id string) []float64 {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	for i, vid := range idx.ids {
		if vid == id {
			return idx.vectors[i]
		}
	}
	return nil
}
