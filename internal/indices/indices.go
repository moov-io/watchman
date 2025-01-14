package indices

import (
	"sync"
)

// New creates slice indices for parallel processing
func New(total, groups int) []int {
	if groups <= 0 {
		return []int{0, total}
	}
	if groups == 1 || groups >= total {
		return []int{0, total}
	}

	chunkSize := total / groups
	remaining := total % groups

	// Pre-allocate slice with exact capacity
	xs := make([]int, 0, groups+1)
	xs = append(xs, 0)

	pos := 0
	for i := 0; i < groups-1; i++ {
		pos += chunkSize
		if remaining > 0 {
			pos++
			remaining--
		}
		xs = append(xs, pos)
	}
	return append(xs, total)
}

// ProcessSlice processes input slice concurrently using the provided function
func ProcessSlice[T any, F any](in []T, groups int, f func(T) F) []F {
	results := make(chan F, len(in))

	var wg sync.WaitGroup
	wg.Add(len(in))

	for _, elm := range in {
		go func(v T) {
			defer wg.Done()
			results <- f(v)
		}(elm)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	out := make([]F, 0, len(in))
	for result := range results {
		out = append(out, result)
	}

	return out
}

func ProcessSliceFn[T any](in []T, groups int, f func(T)) {
	if len(in) == 0 {
		return
	}

	indices := New(len(in), groups)
	numGroups := len(indices) - 1 // Number of actual chunks

	//	fmt.Printf("in=%d  groups=%v  indices=%v  numGroups=%v\n", len(in), groups, indices, numGroups)

	// Use WaitGroup for synchronization
	var wg sync.WaitGroup
	wg.Add(numGroups)

	// Process each chunk concurrently
	for i := 0; i < numGroups; i++ {
		start := indices[i]
		end := indices[i+1]

		go func(start, end int) {
			defer wg.Done()

			// Process chunk and write directly to pre-allocated output slice
			for _, v := range in[start:end] {
				f(v)
			}
		}(start, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
