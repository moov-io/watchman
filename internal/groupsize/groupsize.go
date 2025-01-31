package groupsize

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// rollingStats keeps a fixed-size ring buffer of durations, plus incremental
// mean & variance using Welford's algorithm.
type rollingStats struct {
	mu       sync.Mutex
	buffer   []time.Duration // ring buffer
	capacity int
	index    int
	count    int
	// For incremental mean/variance
	mean float64
	m2   float64 // sum of squares of differences from the mean
}

// newRollingStats creates a ring buffer of a given capacity.
func newRollingStats(capacity int) *rollingStats {
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
	old := rs.buffer[rs.index]
	if rs.count >= rs.capacity {
		// we have a full buffer, so subtract old from the Welford's calc
		oldVal := float64(old.Microseconds())
		n := float64(rs.count)

		// Remove oldVal's contribution
		delta := oldVal - rs.mean
		rs.mean = rs.mean - delta/(n)
		rs.m2 = rs.m2 - delta*(oldVal-rs.mean)
	} else {
		rs.count++
	}

	// Insert the new duration into the ring buffer
	rs.buffer[rs.index] = d
	rs.index = (rs.index + 1) % rs.capacity

	// Add new value
	n := float64(rs.count)
	delta := newVal - rs.mean
	rs.mean = rs.mean + delta/n
	delta2 := newVal - rs.mean
	rs.m2 = rs.m2 + delta*delta2
}

// getStats returns (mean, stddev, sampleSize) from the rolling stats
func (rs *rollingStats) getStats() (mean float64, stddev float64, n int) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	if rs.count < 2 {
		// Not enough data for stddev
		return rs.mean, 0, rs.count
	}
	variance := rs.m2 / float64(rs.count-1)
	return rs.mean, math.Sqrt(variance), rs.count
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

	// For picking concurrency randomly
	randSource *rand.Rand

	// Confidence / decision parameters
	confidenceLevel float64 // e.g., 1.96 for ~95%
	windowSize      int     // rolling window for each concurrency
	minSamples      int     // min samples before we do a compare
	minImprovement  float64

	// cleanup
	lastCleanup     time.Time
	cleanupInterval time.Duration // e.g., 1 hour
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
		randSource:      rand.New(rand.NewSource(time.Now().UnixNano())),
		confidenceLevel: 1.96, // approximate z-score for 95% CI
		windowSize:      100,  // how many samples in rolling window
		minSamples:      20,   // minimum number of samples for each concurrency to compare
		minImprovement:  0.05,
		cleanupInterval: time.Hour,
		lastCleanup:     time.Now(),
	}
	cm.setChampion(initialChampion)
	return cm, nil
}

// setChampion updates the champion and sets up challengers (champ+1 and champ-1) if in range.
func (cm *ConcurrencyManager) setChampion(c int) {
	cm.champion = c

	// Clean up old stats before setting new weights
	if time.Since(cm.lastCleanup) > cm.cleanupInterval {
		cm.cleanupOldStats()
		cm.lastCleanup = time.Now()
	}

	// Clear traffic weights and stats
	cm.trafficWeights = make(map[int]float64)

	// Champion gets 80% traffic
	cm.trafficWeights[c] = 0.8
	cm.ensureStats(c)

	// Potential challengers
	up := c + 1
	down := c - 1

	var challengers []int
	if up <= cm.maxC {
		challengers = append(challengers, up)
	}
	if down >= cm.minC {
		challengers = append(challengers, down)
	}

	if len(challengers) == 1 {
		// If only one valid challenger, give it 20% traffic
		cm.trafficWeights[challengers[0]] = 0.2
		cm.ensureStats(challengers[0])
	} else if len(challengers) == 2 {
		// If two challengers, give each 10% traffic
		cm.trafficWeights[challengers[0]] = 0.1
		cm.trafficWeights[challengers[1]] = 0.1
		cm.ensureStats(challengers[0])
		cm.ensureStats(challengers[1])
	}
}

// ensureStats ensures a rollingStats object exists for concurrency c
func (cm *ConcurrencyManager) ensureStats(c int) {
	if _, ok := cm.stats[c]; !ok {
		cm.stats[c] = newRollingStats(cm.windowSize)
	}
}

// cleanupOldStats will remove outdated stats from the trafficWeights
func (cm *ConcurrencyManager) cleanupOldStats() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for c := range cm.stats {
		if _, isActive := cm.trafficWeights[c]; !isActive {
			delete(cm.stats, c)
		}
	}
}

// PickConcurrency randomly chooses a concurrency among champion/challengers based on traffic weights
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

	r := cm.randSource.Float64() * total
	for c, w := range cm.trafficWeights {
		r -= w
		if r <= 0 {
			return c
		}
	}

	return cm.champion
}

// RecordDuration logs the latency for the given concurrency
func (cm *ConcurrencyManager) RecordDuration(concurrency int, d time.Duration) {
	if d < 0 {
		return // Skip negative durations (clock issues)
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Check if cleanup is needed
	if time.Since(cm.lastCleanup) > cm.cleanupInterval {
		cm.cleanupOldStats()
		cm.lastCleanup = time.Now()
	}

	st := cm.stats[concurrency]
	if st == nil {
		st = newRollingStats(cm.windowSize)
		cm.stats[concurrency] = st
	}
	st.add(d)

	// Check if all tested concurrencies have at least minSamples
	if cm.allHaveMinSamples() {
		cm.evaluate()
	}
}

// allHaveMinSamples checks if champion and any challengers each have enough samples
func (cm *ConcurrencyManager) allHaveMinSamples() bool {
	for c := range cm.trafficWeights {
		_, _, n := cm.stats[c].getStats()
		if n < cm.minSamples {
			return false
		}
	}
	return true
}

type evaluationResult struct {
	concurrency int
	mean        float64
	ci          interval
}

func (cm *ConcurrencyManager) evaluate() {
	now := time.Now()
	if now.Sub(cm.lastSwitchTime) < cm.switchCooldown {
		return
	}

	results := make([]evaluationResult, 0, len(cm.trafficWeights))

	// Gather all stats first
	for c := range cm.trafficWeights {
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
		if r.concurrency == cm.champion {
			champResult = r
			break
		}
	}

	// Find best challenger
	var bestChallenger evaluationResult
	bestImprovement := 0.0

	for _, r := range results {
		if r.concurrency == cm.champion {
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
		cm.lastSwitchTime = now
	}
}

type interval struct {
	low, high float64
}

// confidenceInterval returns (low, high) for the mean at a given confidence level (z).
// mean in microseconds, std in microseconds, count is sample size.
func confidenceInterval(mean, std float64, count int, z float64) interval {
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
