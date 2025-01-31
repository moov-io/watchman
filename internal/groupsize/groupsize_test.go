package groupsize

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
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
