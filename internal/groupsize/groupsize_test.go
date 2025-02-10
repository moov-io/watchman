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

func BenchmarkConcurrentConcurrencyManager(b *testing.B) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	if err != nil {
		b.Fatal(err)
	}

	data := make([]byte, 32)
	_, err = rand.Read(data)
	if err != nil {
		b.Fatal(err)
	}

	// Test with different concurrency levels
	for _, numGoroutines := range []int{1, 4, 8, 16, 32} {
		b.Run(b.Name()+"-goroutines-"+string(rune('0'+numGoroutines)), func(b *testing.B) {
			var wg sync.WaitGroup
			b.ResetTimer()

			// Each goroutine will do b.N/numGoroutines operations
			opsPerGoroutine := b.N / numGoroutines

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					ticker := time.NewTicker(100 * time.Microsecond)
					defer ticker.Stop()

					cleanupCount := opsPerGoroutine / 100
					if cleanupCount < 1 {
						cleanupCount = 1
					}

					for i := 0; i < cleanupCount; i++ {
						<-ticker.C
						cm.cleanupOldStats()
					}
				}()
			}

			// Add a goroutine that triggers cleanup periodically
			wg.Add(1)
			go func() {
				defer wg.Done()
				ticker := time.NewTicker(time.Microsecond)
				defer ticker.Stop()

				for i := 0; i < opsPerGoroutine; i++ {
					<-ticker.C
					cm.cleanupOldStats()
				}
			}()

			wg.Wait()
		})
	}
}

// BenchmarkPickAndRecord measures just the locking overhead without the actual work being done
func BenchmarkPickAndRecord(b *testing.B) {
	cm, err := NewConcurrencyManager(25, 1, 100)
	if err != nil {
		b.Fatal(err)
	}

	for _, numGoroutines := range []int{1, 4, 8, 16, 32} {
		b.Run(b.Name()+"-goroutines-"+string(rune('0'+numGoroutines)), func(b *testing.B) {
			var wg sync.WaitGroup
			b.ResetTimer()

			opsPerGoroutine := b.N / numGoroutines

			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					for j := 0; j < opsPerGoroutine; j++ {
						size := cm.PickConcurrency()
						cm.RecordDuration(size, time.Millisecond)
					}
				}()
			}

			wg.Wait()
		})
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
