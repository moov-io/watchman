package concurrencychamp

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
)

// rollingStats keeps a fixed-size ring buffer of durations, plus incremental
// mean & variance using Welford's algorithm.
type rollingStats struct {
	mu       sync.RWMutex // Use RWMutex instead of Mutex for better read concurrency
	buffer   []time.Duration
	capacity int32
	index    int32 // Use atomic for index operations
	count    int32 // Use atomic for count operations
	mean     float64
	m2       float64 // sum of squares of differences from the mean
}

type ConcurrencyManager struct {
	mu sync.Mutex

	champion int // current best concurrency
	minC     int
	maxC     int

	lastSwitchTime time.Time
	switchCooldown time.Duration // e.g. 1 minute

	// ring-buffers for each concurrency under test
	stats map[int]*rollingStats

	// Traffic distribution:
	// champion gets ~80% of requests, challengers ~10% each
	trafficWeights map[int]float64

	// Confidence / decision parameters
	confidenceLevel float64 // e.g., 1.96 for ~95%
	windowSize      int32   // rolling window for each concurrency
	minSamples      int32   // min samples before we do a compare
	minImprovement  float64

	// cleanup
	lastCleanup     time.Time
	lastCleanupUnix int64         // atomic timestamp for cleanup
	cleanupInterval time.Duration // e.g., 1 hour
	totalSamples    int32         // atomic counter
	evaluateChan    chan struct{} // channel for async evaluation
}

type interval struct {
	low, high float64
}

type evaluationResult struct {
	concurrency int
	mean        float64
	ci          interval
}

// NewConcurrencyManager sets up with a chosen champion concurrency and typical config
func NewConcurrencyManager(initialChampion, minC, maxC int) (*ConcurrencyManager, error) {
	if initialChampion < minC || initialChampion > maxC {
		return nil, fmt.Errorf("initial champion %d must be between min %d and max %d", initialChampion, minC, maxC)
	}
	if minC <= 0 {
		return nil, fmt.Errorf("minimum concurrency must be positive, got %d", minC)
	}
	if maxC <= minC {
		return nil, fmt.Errorf("maximum concurrency %d must be greater than minimum %d", maxC, minC)
	}

	cm := &ConcurrencyManager{
		champion:        initialChampion,
		minC:            minC,
		maxC:            maxC,
		stats:           make(map[int]*rollingStats),
		trafficWeights:  make(map[int]float64),
		confidenceLevel: 1.645, // approximate z-score for 90% CI
		minSamples:      10,
		minImprovement:  0.02,
		windowSize:      100,         // Add explicit window size
		switchCooldown:  time.Second, // Add explicit cooldown
		cleanupInterval: 10 * time.Minute,
		lastCleanup:     time.Now(),
		lastCleanupUnix: time.Now().Unix(),
		evaluateChan:    make(chan struct{}, 1),
	}

	cm.setChampion(initialChampion)

	// Start evaluation goroutine
	go cm.evaluationLoop()

	// Run cleanup in an async goroutine
	go cm.startBackgroundCleanup()

	return cm, nil
}

// newRollingStats creates a ring buffer of a given capacity.
func newRollingStats(capacity int32) *rollingStats {
	if capacity <= 0 {
		capacity = 100 // Default to reasonable size
	}
	return &rollingStats{
		buffer:   make([]time.Duration, capacity),
		capacity: capacity,
	}
}

// add inserts a new duration, updating mean and variance. Welford's method in ring form.
func (rs *rollingStats) add(d time.Duration) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	newVal := float64(d.Microseconds())
	if math.IsNaN(newVal) || math.IsInf(newVal, 0) {
		return // Skip invalid values
	}

	// If we're overwriting an old value, remove its contribution
	old := rs.buffer[atomic.LoadInt32(&rs.index)]
	if atomic.LoadInt32(&rs.count) >= rs.capacity {
		// we have a full buffer, so subtract old from the Welford's calc
		oldVal := float64(old.Microseconds())
		n := float64(atomic.LoadInt32(&rs.count))

		// Remove oldVal's contribution
		delta := oldVal - rs.mean
		rs.mean = rs.mean - delta/n
		rs.m2 = rs.m2 - delta*(oldVal-rs.mean)
	} else {
		atomic.AddInt32(&rs.count, 1)
	}

	// Insert the new duration into the ring buffer
	rs.buffer[rs.index] = d
	atomic.StoreInt32(&rs.index, (atomic.LoadInt32(&rs.index)+1)%rs.capacity)

	// Add new value
	n := float64(atomic.LoadInt32(&rs.count))
	delta := newVal - rs.mean
	rs.mean = rs.mean + delta/n
	delta2 := newVal - rs.mean
	rs.m2 = rs.m2 + delta*delta2
}

// getStats returns (mean, stddev, sampleSize) from the rolling stats
func (rs *rollingStats) getStats() (mean float64, stddev float64, n int32) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	count := atomic.LoadInt32(&rs.count)
	if count < 2 {
		return rs.mean, 0, count
	}
	variance := rs.m2 / float64(count-1)
	return rs.mean, math.Sqrt(variance), count
}

func (cm *ConcurrencyManager) setChampion(c int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.champion = c

	// Clear traffic weights and stats
	cm.trafficWeights = make(map[int]float64)

	// Champion gets 70% traffic
	cm.trafficWeights[c] = 0.7
	cm.ensureStats(c)

	// Define step sizes - both fine and coarse adjustments
	steps := []int{1, 5}

	// Collect all valid challengers
	challengers := make([]int, 0, 4) // Pre-allocate with reasonable capacity
	for _, step := range steps {
		up := c + step
		down := c - step

		if up <= cm.maxC {
			challengers = append(challengers, up)
		}
		if down >= cm.minC {
			challengers = append(challengers, down)
		}
	}

	// Distribute remaining 30% traffic among challengers
	if len(challengers) > 0 {
		weight := 0.3 / float64(len(challengers))
		for _, challenger := range challengers {
			cm.trafficWeights[challenger] = weight
			cm.ensureStats(challenger)
		}
	}
}

func (cm *ConcurrencyManager) ensureStats(c int) {
	if _, ok := cm.stats[c]; !ok {
		cm.stats[c] = newRollingStats(cm.windowSize)
	}
}

func (cm *ConcurrencyManager) startBackgroundCleanup() {
	go func() {
		ticker := time.NewTicker(cm.cleanupInterval)
		defer ticker.Stop()

		for range ticker.C {
			cm.cleanupOldStats()
		}
	}()
}

func (cm *ConcurrencyManager) cleanupOldStats() {
	if !time.Unix(atomic.LoadInt64(&cm.lastCleanupUnix), 0).Add(cm.cleanupInterval).Before(time.Now()) {
		return // Another goroutine already cleaned up
	}

	active := make(map[int]bool)
	for c := range cm.trafficWeights {
		active[c] = true
	}

	for c := range cm.stats {
		if !active[c] {
			delete(cm.stats, c)
		}
	}

	cm.lastCleanup = time.Now()
	atomic.StoreInt64(&cm.lastCleanupUnix, time.Now().Unix())
}

func (cm *ConcurrencyManager) PickConcurrency() int {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	total := 0.0
	for _, w := range cm.trafficWeights {
		total += w
	}

	if total == 0 {
		return cm.champion
	}

	r := rand.Float64() * total
	for c, w := range cm.trafficWeights {
		r -= w
		if r <= 0 {
			return c
		}
	}

	return cm.champion
}

func (cm *ConcurrencyManager) RecordDuration(concurrency int, d time.Duration) {
	var st *rollingStats
	cm.mu.Lock()
	st = cm.stats[concurrency]
	if st == nil {
		st = newRollingStats(cm.windowSize)
		cm.stats[concurrency] = st
	}
	cm.mu.Unlock()

	st.add(d)

	// Increment sample count atomically
	if atomic.AddInt32(&cm.totalSamples, 1) >= cm.minSamples {
		select {
		case cm.evaluateChan <- struct{}{}:
		default:
		}
	}
}

func (cm *ConcurrencyManager) evaluationLoop() {
	for range cm.evaluateChan {
		if cm.allHaveMinSamples() {
			cm.evaluate()
		}
	}
}

func (cm *ConcurrencyManager) allHaveMinSamples() bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for c := range cm.trafficWeights {
		_, _, n := cm.stats[c].getStats()
		if n < cm.minSamples {
			return false
		}
	}
	return true
}

func (cm *ConcurrencyManager) evaluate() {
	now := time.Now()
	if now.Sub(cm.lastSwitchTime) < cm.switchCooldown {
		return
	}

	cm.mu.Lock()
	// Gather data under lock
	trafficWeightsCopy := make(map[int]float64, len(cm.trafficWeights))
	for k, v := range cm.trafficWeights {
		trafficWeightsCopy[k] = v
	}
	champion := cm.champion
	cm.mu.Unlock()

	results := make([]evaluationResult, 0, len(trafficWeightsCopy))

	// Gather all stats first
	for c := range trafficWeightsCopy {
		mean, std, count := cm.stats[c].getStats()
		if count < cm.minSamples {
			continue
		}
		ci := confidenceInterval(mean, std, count, cm.confidenceLevel)
		results = append(results, evaluationResult{
			concurrency: c,
			mean:        mean,
			ci:          ci,
		})
	}

	if len(results) < 2 {
		return // Need at least champion and one challenger
	}

	// Find champion stats
	var champResult evaluationResult
	for _, r := range results {
		if r.concurrency == champion {
			champResult = r
			break
		}
	}

	// Find best challenger
	var bestChallenger evaluationResult
	bestImprovement := 0.0

	for _, r := range results {
		if r.concurrency == champion {
			continue
		}

		improvement := (champResult.mean - r.mean) / champResult.mean
		if r.ci.high < champResult.ci.low && improvement >= cm.minImprovement {
			if improvement > bestImprovement {
				bestChallenger = r
				bestImprovement = improvement
			}
		}
	}

	if bestImprovement > 0 {
		cm.setChampion(bestChallenger.concurrency)
		cm.mu.Lock()
		cm.lastSwitchTime = now
		cm.mu.Unlock()
	}
}

// confidenceInterval returns (low, high) for the mean at a given confidence level (z).
// mean in microseconds, std in microseconds, count is sample size.
func confidenceInterval(mean, std float64, count int32, z float64) interval {
	var ci interval
	if count < 2 || std == 0 {
		// degenerate
		ci.low = mean
		ci.high = mean
		return ci
	}
	sem := std / math.Sqrt(float64(count))
	delta := z * sem
	ci.low = mean - delta
	ci.high = mean + delta
	return ci
}
