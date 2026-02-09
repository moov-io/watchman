package embeddings

import (
	"container/heap"
	"fmt"
	"runtime"
	"sort"
	"sync"
)

// vectorIndex provides approximate nearest neighbor search for embeddings.
// Optimized for ~50k entries with multi-core parallel search.
//
// TODO(adam): look at Hierarchical Navigable Small World Graphs (HNSW)
// https://github.com/fogfish/hnsw
type vectorIndex struct {
	mu         sync.RWMutex
	vectorData []float64 // Flat array: all vectors concatenated for cache efficiency
	ids        []string
	names      []string
	idMap      map[string]int // O(1) lookup by ID
	dim        int
	count      int // Number of vectors stored
}

// newVectorIndex creates a new empty vector index.
func newVectorIndex(dim int) *vectorIndex {
	const estimatedCapacity = 50000
	return &vectorIndex{
		dim:        dim,
		vectorData: make([]float64, 0, estimatedCapacity*dim),
		ids:        make([]string, 0, estimatedCapacity),
		names:      make([]string, 0, estimatedCapacity),
		idMap:      make(map[string]int, estimatedCapacity),
	}
}

// Add adds vectors to the index with their corresponding IDs and names.
func (idx *vectorIndex) Add(vectors [][]float64, ids, names []string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	baseIdx := idx.count

	// Flatten vectors into contiguous memory
	for _, vec := range vectors {
		idx.vectorData = append(idx.vectorData, vec...)
	}

	idx.ids = append(idx.ids, ids...)
	idx.names = append(idx.names, names...)

	// Build ID index for O(1) lookups
	for i, id := range ids {
		idx.idMap[id] = baseIdx + i
	}

	idx.count += len(vectors)
}

// Search finds the k most similar vectors to the query.
// Returns results sorted by similarity score (highest first).
// Uses parallel computation across CPU cores for optimal performance.
func (idx *vectorIndex) Search(query []float64, k int) []SearchResult {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if idx.count == 0 {
		return nil
	}

	scores := make([]scored, idx.count)

	// Parallel similarity computation across all cores
	numWorkers := runtime.GOMAXPROCS(0)
	chunkSize := (idx.count + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		start := w * chunkSize
		end := start + chunkSize
		if end > idx.count {
			end = idx.count
		}

		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				vecStart := i * idx.dim
				scores[i] = scored{
					idx:   i,
					score: dotProduct(query, idx.vectorData[vecStart:vecStart+idx.dim]),
				}
			}
		}(start, end)
	}

	wg.Wait()

	// Use heap-based partial sort for top-k: O(n log k) instead of O(n log n)
	if k < idx.count {
		topK := selectTopK(scores, k)
		results := make([]SearchResult, k)
		for i := 0; i < k; i++ {
			s := topK[i]
			results[i] = SearchResult{
				ID:    idx.ids[s.idx],
				Name:  idx.names[s.idx],
				Score: s.score,
			}
		}
		return results
	}

	// If k >= count, sort everything
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	results := make([]SearchResult, idx.count)
	for i := 0; i < idx.count; i++ {
		s := scores[i]
		results[i] = SearchResult{
			ID:    idx.ids[s.idx],
			Name:  idx.names[s.idx],
			Score: s.score,
		}
	}

	return results
}

// selectTopK uses a min-heap to efficiently find the k highest scores.
// Time complexity: O(n log k) vs O(n log n) for full sort.
func selectTopK(scores []scored, k int) []scored {
	h := &minHeap{data: make([]scored, 0, k)}
	heap.Init(h)

	for i := range scores {
		if h.Len() < k {
			heap.Push(h, scores[i])
		} else if scores[i].score > h.data[0].score {
			heap.Pop(h)
			heap.Push(h, scores[i])
		}
	}

	// Extract in descending order
	result := make([]scored, k)
	for i := k - 1; i >= 0; i-- {
		v := heap.Pop(h)
		s, ok := v.(scored)
		if !ok {
			panic(fmt.Errorf("unexpected %T", v)) //nolint:forbidigo
		}
		result[i] = s
	}

	return result
}

// Size returns the number of vectors in the index.
func (idx *vectorIndex) Size() int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return idx.count
}

// Clear removes all vectors from the index.
func (idx *vectorIndex) Clear() {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	const estimatedCapacity = 50000
	idx.vectorData = make([]float64, 0, estimatedCapacity*idx.dim)
	idx.ids = make([]string, 0, estimatedCapacity)
	idx.names = make([]string, 0, estimatedCapacity)
	idx.idMap = make(map[string]int, estimatedCapacity)
	idx.count = 0
}

// GetVector returns the embedding vector for a given ID.
// Returns nil if the ID is not found.
func (idx *vectorIndex) GetVector(id string) []float64 {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if i, ok := idx.idMap[id]; ok {
		start := i * idx.dim
		end := start + idx.dim
		// Return a copy to prevent external modification
		result := make([]float64, idx.dim)
		copy(result, idx.vectorData[start:end])
		return result
	}
	return nil
}

// minHeap implements heap.Interface for selecting top-k scores.
type minHeap struct {
	data []scored
}

type scored struct {
	idx   int
	score float64
}

func (h minHeap) Len() int           { return len(h.data) }
func (h minHeap) Less(i, j int) bool { return h.data[i].score < h.data[j].score }
func (h minHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *minHeap) Push(x interface{}) {
	if s, ok := x.(scored); ok {
		h.data = append(h.data, s)
	}
}

func (h *minHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
