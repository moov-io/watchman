package minmaxmed

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// Original implementation wrapping
type OriginalObservor struct {
	*Observor
}

func NewOriginal(size int) *OriginalObservor {
	return &OriginalObservor{New(size)}
}

// Optimized implementation wrapping
type OptimizedObservor struct {
	*Observor
}

func NewOptimized(size int) *OptimizedObservor {
	return &OptimizedObservor{New(size)}
}

var (
	windowSizes = []int{10, 100, 1000, 10000}
	benchSizes  = []int{1000, 10000, 100000, 1000000}
)

func generateTestData(n int) []int64 {
	data := make([]int64, n)
	r := rand.New(rand.NewSource(42)) // Fixed seed for reproducibility
	for i := range data {
		data[i] = r.Int63n(1000000) // Random values between 0 and 1M
	}
	return data
}

// BenchmarkAdd tests basic add operations
func BenchmarkAdd(b *testing.B) {
	for _, windowSize := range windowSizes {
		b.Run("Original/WindowSize="+itoa(windowSize), func(b *testing.B) {
			obs := NewOriginal(windowSize)
			data := generateTestData(b.N)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				obs.Add(data[i%len(data)])
			}
		})

		b.Run("Optimized/WindowSize="+itoa(windowSize), func(b *testing.B) {
			obs := NewOptimized(windowSize)
			data := generateTestData(b.N)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				obs.Add(data[i%len(data)])
			}
		})
	}
}

// BenchmarkAddDuration tests duration add operations
func BenchmarkAddDuration(b *testing.B) {
	durations := make([]time.Duration, b.N)
	r := rand.New(rand.NewSource(42))
	for i := range durations {
		durations[i] = time.Duration(r.Int63n(1000000)) * time.Millisecond
	}

	for _, windowSize := range windowSizes {
		b.Run("Original/WindowSize="+itoa(windowSize), func(b *testing.B) {
			obs := NewOriginal(windowSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				obs.AddDuration(durations[i%len(durations)])
			}
		})

		b.Run("Optimized/WindowSize="+itoa(windowSize), func(b *testing.B) {
			obs := NewOptimized(windowSize)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				obs.AddDuration(durations[i%len(durations)])
			}
		})
	}
}

// BenchmarkConcurrentAdd tests pure concurrent add operations
func BenchmarkConcurrentAdd(b *testing.B) {
	for _, windowSize := range windowSizes {
		for _, goroutines := range []int{2, 4, 8, 16} {
			name := fmt.Sprintf("Original/WindowSize=%d/Goroutines=%d", windowSize, goroutines)
			b.Run(name, func(b *testing.B) {
				obs := NewOriginal(windowSize)
				data := generateTestData(b.N)
				b.ResetTimer()
				b.SetParallelism(goroutines)
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						obs.Add(data[i%len(data)])
						i++
					}
				})
			})

			name = fmt.Sprintf("Optimized/WindowSize=%d/Goroutines=%d", windowSize, goroutines)
			b.Run(name, func(b *testing.B) {
				obs := NewOptimized(windowSize)
				data := generateTestData(b.N)
				b.ResetTimer()
				b.SetParallelism(goroutines)
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						obs.Add(data[i%len(data)])
						i++
					}
				})
			})
		}
	}
}

// BenchmarkConcurrentReadWrite tests mixed read/write patterns with AddEvent
func BenchmarkConcurrentReadWrite(b *testing.B) {
	mockSpan := trace.SpanFromContext(context.Background())

	for _, windowSize := range windowSizes {
		for _, goroutines := range []int{2, 4, 8, 16} {
			name := fmt.Sprintf("Original/WindowSize=%d/Goroutines=%d", windowSize, goroutines)
			b.Run(name, func(b *testing.B) {
				obs := NewOriginal(windowSize)
				data := generateTestData(b.N)

				// Pre-fill with some data
				for i := 0; i < windowSize; i++ {
					if i < len(data) {
						obs.Add(data[i])
					}
				}

				b.ResetTimer()
				b.SetParallelism(goroutines)
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						if i%2 == 0 {
							// Write operation
							obs.Add(data[i%len(data)])
						} else {
							// Read operation via AddEvent
							obs.AddEvent(mockSpan)
						}
						i++
					}
				})
			})

			name = fmt.Sprintf("Optimized/WindowSize=%d/Goroutines=%d", windowSize, goroutines)
			b.Run(name, func(b *testing.B) {
				obs := NewOptimized(windowSize)
				data := generateTestData(b.N)

				// Pre-fill with some data
				for i := 0; i < windowSize; i++ {
					if i < len(data) {
						obs.Add(data[i])
					}
				}

				b.ResetTimer()
				b.SetParallelism(goroutines)
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						if i%2 == 0 {
							// Write operation
							obs.Add(data[i%len(data)])
						} else {
							// Read operation via AddEvent
							obs.AddEvent(mockSpan)
						}
						i++
					}
				})
			})
		}
	}
}

// Utility function since we can't use strconv.Itoa directly in benchmarks
func itoa(n int) string {
	switch {
	case n == 10:
		return "10"
	case n == 100:
		return "100"
	case n == 1000:
		return "1000"
	case n == 10000:
		return "10000"
	case n == 100000:
		return "100000"
	case n == 1000000:
		return "1000000"
	default:
		return "unknown"
	}
}
