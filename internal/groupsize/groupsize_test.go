package groupsize

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConcurrencyManager(t *testing.T) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	require.NoError(t, err)

	size := cm.PickConcurrency()
	require.Greater(t, size, 0)

	for i := 0; i < 10_000; i++ {
		data := make([]byte, 32)
		rand.Read(data)

		size := cm.PickConcurrency()

		start := time.Now()
		dowork(data, 1)
		cm.RecordDuration(size, time.Since(start))
	}

	size = cm.PickConcurrency()
	t.Logf("after doing work: %d", size)
	require.Greater(t, size, 0)

	t.Run("cleanup", func(t *testing.T) {
		start := time.Now()
		cm.cleanupOldStats()
		diff := time.Since(start)

		t.Logf("cleanup took %v", diff)

		require.Less(t, diff, 5*time.Millisecond)
	})
}

func BenchmarkConcurrencyManager(b *testing.B) {
	// Test cases covering key scenarios
	tests := []struct {
		name          string
		goroutines    int
		cleanupFreq   time.Duration
		recordingFreq time.Duration
	}{
		{"SingleThreaded", 1, time.Second, time.Millisecond},
		{"ModerateContention", 4, time.Second, time.Millisecond},
		{"HighContention", 16, time.Second, time.Millisecond},
	}

	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			cm, err := NewConcurrencyManager(25, 1, 100)
			if err != nil {
				b.Fatal(err)
			}

			var wg sync.WaitGroup
			opsPerGoroutine := b.N / tc.goroutines

			// Start cleanup goroutine
			wg.Add(1)
			go func() {
				defer wg.Done()
				ticker := time.NewTicker(tc.cleanupFreq)
				defer ticker.Stop()

				for i := 0; i < opsPerGoroutine; i++ {
					<-ticker.C
					cm.cleanupOldStats()
				}
			}()

			// Start worker goroutines
			for i := 0; i < tc.goroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					ticker := time.NewTicker(tc.recordingFreq)
					defer ticker.Stop()

					for i := 0; i < opsPerGoroutine; i++ {
						<-ticker.C
						size := cm.PickConcurrency()
						cm.RecordDuration(size, time.Millisecond)
					}
				}()
			}

			wg.Wait()
		})
	}
}

func BenchmarkEvaluationLatency(b *testing.B) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	if err != nil {
		b.Fatal(err)
	}

	// Pre-populate with some data
	for i := 0; i < 1000; i++ {
		size := cm.PickConcurrency()
		cm.RecordDuration(size, time.Millisecond)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cm.evaluate()
	}
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
