package concurrencychamp

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConcurrencyManager(t *testing.T) {
	t.Run("Champion Switching", func(t *testing.T) {
		cm, err := NewConcurrencyManager(25, 1, 100)
		require.NoError(t, err)

		// Set key parameters for testing
		cm.switchCooldown = time.Millisecond
		cm.windowSize = 50
		cm.minSamples = 10
		cm.minImprovement = 0.02

		initial := cm.champion
		better := initial + 1

		// Record dramatically different durations many times
		for i := 0; i < 1000; i++ { // More samples
			cm.RecordDuration(initial, 1000*time.Millisecond) // 1 second
			cm.RecordDuration(better, 1*time.Millisecond)     // 1 millisecond
		}

		// Force multiple evaluations with time for the evaluationLoop to process
		for i := 0; i < 10; i++ {
			cm.evaluate()
			time.Sleep(10 * time.Millisecond)
		}

		if cm.champion == initial {
			initialStats := cm.stats[initial]
			betterStats := cm.stats[better]

			initialMean, initialStd, initialCount := initialStats.getStats()
			betterMean, betterStd, betterCount := betterStats.getStats()

			t.Logf("Initial champion %d stats: mean=%v std=%v count=%d",
				initial, initialMean, initialStd, initialCount)
			t.Logf("Challenger %d stats: mean=%v std=%v count=%d",
				better, betterMean, betterStd, betterCount)
		}

		require.NotEqual(t, initial, cm.champion,
			"champion should switch to better performer")
	})
}

func BenchmarkConcurrencyManager(b *testing.B) {
	tests := []struct {
		name          string
		goroutines    int
		cleanupFreq   time.Duration
		recordingFreq time.Duration
		workload      func() time.Duration // Add varying workloads
	}{
		{
			name:          "LowContention",
			goroutines:    2,
			cleanupFreq:   time.Second,
			recordingFreq: time.Millisecond,
			workload:      func() time.Duration { return time.Millisecond },
		},
		{
			name:          "HighContention",
			goroutines:    32,
			cleanupFreq:   100 * time.Millisecond,
			recordingFreq: time.Microsecond,
			workload: func() time.Duration {
				// Simulate varying latencies
				return time.Duration(rand.Int63n(int64(10 * time.Millisecond)))
			},
		},
		{
			name:          "BurstyWorkload",
			goroutines:    16,
			cleanupFreq:   time.Second,
			recordingFreq: time.Millisecond,
			workload: func() time.Duration {
				if rand.Float64() < 0.1 {
					return 100 * time.Millisecond // Occasional spike
				}
				return time.Millisecond
			},
		},
	}

	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			cm, err := NewConcurrencyManager(25, 1, 100)
			require.NoError(b, err)

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					champ := cm.PickConcurrency()
					cm.RecordDuration(champ, tc.workload())
				}
			})
		})
	}
}

func TestRaceConditions(t *testing.T) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	require.NoError(t, err)

	var wg sync.WaitGroup
	// Deliberately create race conditions
	for i := 0; i < 100; i++ {
		wg.Add(3)

		// Concurrent stats recording
		go func() {
			defer wg.Done()
			champ := cm.PickConcurrency()
			cm.RecordDuration(champ, time.Millisecond)
		}()

		// Concurrent evaluation
		go func() {
			defer wg.Done()
			cm.evaluate()
		}()

		// Concurrent cleanup
		go func() {
			defer wg.Done()
			cm.cleanupOldStats()
		}()
	}

	wg.Wait()
}

func TestStressConcurrencyManager(t *testing.T) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	require.NoError(t, err)

	duration := 5 * time.Second
	if testing.Short() {
		duration = 1 * time.Second
	}

	var wg sync.WaitGroup
	done := make(chan struct{})

	// Start multiple goroutines doing different operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					champ := cm.PickConcurrency()
					cm.RecordDuration(champ, time.Duration(rand.Int63n(int64(10*time.Millisecond))))
				}
			}
		}()
	}

	time.Sleep(duration)
	close(done)
	wg.Wait()
}

func dowork(data []byte, zeros int) ([]byte, int) {
	if zeros > 32 {
		return nil, 0
	}

	prefix := make([]byte, zeros)
	nonce := 0
	result := make([]byte, 32)

	for {
		h := sha256.New()
		binary.LittleEndian.PutUint64(result, uint64(nonce))
		h.Write(result[:8])
		h.Write(data)
		hash := h.Sum(nil)

		if bytes.HasPrefix(hash, prefix) {
			return hash, nonce
		}
		nonce++
	}
}
