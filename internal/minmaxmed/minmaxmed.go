package minmaxmed

import (
	"math"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Observor is a structure which records the minimum and maximum for a streaming set
// of observations and timings.
type Observor struct {
	mu sync.RWMutex

	// Combined stats in a single cache line
	stats struct {
		min   int64
		max   int64
		sum   int64
		count int64
	}

	// Optimized median calculation using a circular buffer and selection algorithm
	window      []float64 // circular buffer for window
	pos         int       // current position in circular buffer
	windowSize  int       // size of the window
	windowCount int       // current count in window
}

// New returns an Observor
func New(medianWindowSize int) *Observor {
	return &Observor{
		stats: struct {
			min   int64
			max   int64
			sum   int64
			count int64
		}{
			min: math.MaxInt64,
		},
		window:     make([]float64, medianWindowSize),
		windowSize: medianWindowSize,
	}
}

func (o *Observor) AddDuration(v time.Duration) {
	o.Add(v.Milliseconds())
}

// Add adds a new observation
func (o *Observor) Add(n int64) {
	o.mu.Lock()

	// Update min/max/sum/count
	if n < o.stats.min {
		o.stats.min = n
	}
	if n > o.stats.max {
		o.stats.max = n
	}
	o.stats.sum += n
	o.stats.count++

	// Update window for median calculation
	o.window[o.pos] = float64(n)
	o.pos = (o.pos + 1) % o.windowSize
	if o.windowCount < o.windowSize {
		o.windowCount++
	}

	o.mu.Unlock()
}

// quickSelect implements Floyd-Rivest selection algorithm
// Returns the k-th smallest element in arr[left:right+1]
func quickSelect(arr []float64, left, right, k int) float64 {
	for right > left {
		// Use median-of-three partitioning
		pivot := left + (right-left)/2
		if right-left > 40 {
			// For larger arrays, use ninther
			s := (right - left) / 8
			medianOfThree(arr, left, left+s, left+2*s)
			medianOfThree(arr, pivot-s, pivot, pivot+s)
			medianOfThree(arr, right-2*s, right-s, right)
			medianOfThree(arr, left+s, pivot, right-s)
			pivot = left + s
		}

		// Partition around pivot
		pivotVal := arr[pivot]
		arr[pivot], arr[right] = arr[right], arr[pivot]
		store := left

		for i := left; i < right; i++ {
			if arr[i] < pivotVal {
				arr[store], arr[i] = arr[i], arr[store]
				store++
			}
		}
		arr[right], arr[store] = arr[store], arr[right]

		// Adjust bounds for next iteration
		if store == k {
			return arr[store]
		} else if store < k {
			left = store + 1
		} else {
			right = store - 1
		}
	}
	return arr[left]
}

// medianOfThree sorts three elements
func medianOfThree(arr []float64, a, b, c int) {
	if arr[a] > arr[b] {
		arr[a], arr[b] = arr[b], arr[a]
	}
	if arr[b] > arr[c] {
		arr[b], arr[c] = arr[c], arr[b]
		if arr[a] > arr[b] {
			arr[a], arr[b] = arr[b], arr[a]
		}
	}
}

// getMedian calculates the current median using QuickSelect
func (o *Observor) getMedian() float64 {
	if o.windowCount == 0 {
		return 0.0
	}

	// Create a copy of the current window
	temp := make([]float64, o.windowCount)
	if o.windowCount < o.windowSize {
		copy(temp, o.window[:o.windowCount])
	} else {
		// Handle wrap-around in circular buffer
		n := copy(temp, o.window[o.pos:])
		copy(temp[n:], o.window[:o.pos])
	}

	if o.windowCount%2 == 1 {
		return quickSelect(temp, 0, o.windowCount-1, o.windowCount/2)
	}

	// For even count, average the two middle values
	lower := quickSelect(temp, 0, o.windowCount-1, (o.windowCount-1)/2)
	upper := quickSelect(temp, 0, o.windowCount-1, o.windowCount/2)
	return (lower + upper) / 2.0
}

func (o *Observor) AddEvent(s trace.Span) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.stats.count == 0 {
		return
	}

	s.AddEvent("stats",
		trace.WithAttributes(
			attribute.Int64("min_ms", o.stats.min),
			attribute.Int64("max_ms", o.stats.max),
			attribute.Float64("median_ms", o.getMedian()),
			attribute.Float64("average_ms", float64(o.stats.sum)/float64(o.stats.count)),
			attribute.Int64("observations", o.stats.count),
		),
	)
}
