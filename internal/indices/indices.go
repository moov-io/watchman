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

// ProcessSliceFn executes f over the input slice concurrency
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

	out := make([]F, len(in))

	// Create work distribution channels
	jobs := make(chan int, len(in))
	var wg sync.WaitGroup

	// Start worker pool
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				out[i] = f(in[i])
			}
		}()
	}

	// Send work to workers
	for i := range in {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	return out
}

// ProcessSliceFn executes f over the input slice concurrency
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

	// Create work distribution channels
	jobs := make(chan int, len(in))
	var wg sync.WaitGroup

	// Start worker pool
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				f(in[i])
			}
		}()
	}

	// Send work to workers
	for i := range in {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}
