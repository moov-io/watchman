package minmaxmed

import (
	"math"
	"sync"
	"time"

	"github.com/JaderDias/movingmedian"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Observor is a structure which records the minimum and maximum for a streaming set
// of observations and timings.
//
// Observations can be included by calling Add(n).
type Observor struct {
	mu       sync.RWMutex
	min, max int64
	sum      int64
	count    int64

	median movingmedian.MovingMedian
}

// New returns an Observor
func New(medianWindowSize int) *Observor {
	return &Observor{
		min:    math.MaxInt64,
		median: movingmedian.NewMovingMedian(medianWindowSize),
	}
}

func (o *Observor) AddDuration(v time.Duration) {
	o.Add(v.Milliseconds())
}

func (o *Observor) Add(n int64) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// track min / max
	if n < o.min {
		o.min = n
	}
	if n > o.max {
		o.max = n
	}

	o.sum += n
	o.count += 1

	o.median.Push(float64(n))
}

func (o *Observor) AddEvent(s trace.Span) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.count == 0 {
		return
	}

	s.AddEvent("stats",
		trace.WithAttributes(
			attribute.Int64("min_ms", o.min),
			attribute.Int64("max_ms", o.max),
			attribute.Float64("median_ms", o.median.Median()),
			attribute.Float64("average_ms", float64(o.sum)/float64(o.count)),
			attribute.Int64("observations", o.count),
		),
	)
}
