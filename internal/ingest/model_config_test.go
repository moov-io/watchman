package ingest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_maxBodyBytes(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		require.Equal(t, defaultMaxBodyBytes, Config{}.maxBodyBytes())
	})

	t.Run("from config", func(t *testing.T) {
		require.Equal(t, int64(1024), Config{MaxBodyBytes: 1024}.maxBodyBytes())
	})

	t.Run("env overrides config", func(t *testing.T) {
		t.Setenv("INGEST_MAX_BODY_BYTES", "4096")
		require.Equal(t, int64(4096), Config{MaxBodyBytes: 1024}.maxBodyBytes())
	})

	t.Run("invalid env uses config", func(t *testing.T) {
		t.Setenv("INGEST_MAX_BODY_BYTES", "nope")
		require.Equal(t, int64(2048), Config{MaxBodyBytes: 2048}.maxBodyBytes())
	})
}
