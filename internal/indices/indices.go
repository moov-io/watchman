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

// ProcessSlice processes items concurrently using a worker pool
func ProcessSlice[T any, F any](in []T, workers int, f func(T) F) []F {
	if len(in) == 0 {
		return nil
	}

	// For very small slices, process sequentially
	if len(in) < workers {
		out := make([]F, len(in))
		for i, item := range in {
			out[i] = f(item)
		}
		return out
	}

	// Calculate chunk size
	chunkSize := len(in) / workers
	if chunkSize < 1 {
		chunkSize = 1
	}

	// Pre-allocate output slice
	out := make([]F, len(in))
	var wg sync.WaitGroup

	// Process chunks directly, writing to pre-allocated slice
	for i := 0; i < len(in); i += chunkSize {
		wg.Add(1)
		start := i
		end := i + chunkSize
		if end > len(in) {
			end = len(in)
		}

		go func(start, end int) {
			defer wg.Done()
			// Process chunk and write directly to output slice
			for i, item := range in[start:end] {
				out[start+i] = f(item)
			}
		}(start, end)
	}

	wg.Wait()
	return out
}

// ProcessSliceFn processes items in chunks concurrently
func ProcessSliceFn[T any](in []T, workers int, f func(T)) {
	if len(in) == 0 {
		return
	}

	// For very small slices, process sequentially
	if len(in) < workers {
		for _, item := range in {
			f(item)
		}
		return
	}

	// Calculate chunk size
	chunkSize := len(in) / workers
	if chunkSize < 1 {
		chunkSize = 1
	}

	var wg sync.WaitGroup

	// Process chunks directly
	for i := 0; i < len(in); i += chunkSize {
		wg.Add(1)
		start := i
		end := i + chunkSize
		if end > len(in) {
			end = len(in)
		}

		go func(start, end int) {
			defer wg.Done()
			for _, item := range in[start:end] {
				f(item)
			}
		}(start, end)
	}

	wg.Wait()
}
