package tfidf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/moov-io/base/strx"
)

// Config holds TF-IDF configuration options
type Config struct {
	// Enabled controls whether TF-IDF weighting is applied to name matching.
	// Default: false (opt-in feature)
	Enabled bool

	// SmoothingFactor (k) in the IDF formula: log((N+k)/(df+k))
	// Prevents division by zero and smooths IDF values.
	// Default: 1.0
	SmoothingFactor float64

	// MinIDF is the floor for IDF values. Prevents very common terms
	// from having zero or negative weight.
	// Default: 0.1
	MinIDF float64

	// MaxIDF is the ceiling for IDF values. Prevents single-occurrence
	// terms from dominating the score.
	// Default: 5.0
	MaxIDF float64
}

// DefaultConfig returns sensible defaults for TF-IDF configuration.
func DefaultConfig() Config {
	return Config{
		Enabled:         false,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
}

// ConfigFromEnvironment reads TF-IDF configuration from environment variables.
//
// Environment variables:
//   - TFIDF_ENABLED: Enable/disable TF-IDF (default: false)
//   - TFIDF_SMOOTHING: Smoothing factor (default: 1.0)
//   - TFIDF_MIN_IDF: Minimum IDF value (default: 0.1)
//   - TFIDF_MAX_IDF: Maximum IDF value (default: 5.0)
func ConfigFromEnvironment() Config {
	cfg := DefaultConfig()

	cfg.Enabled = strx.Yes(os.Getenv("TFIDF_ENABLED"))
	cfg.SmoothingFactor = readFloat("TFIDF_SMOOTHING", cfg.SmoothingFactor)
	cfg.MinIDF = readFloat("TFIDF_MIN_IDF", cfg.MinIDF)
	cfg.MaxIDF = readFloat("TFIDF_MAX_IDF", cfg.MaxIDF)

	return cfg
}

func readFloat(envVar string, defaultValue float64) float64 {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse %s=%q as float64: %w", envVar, value, err)) //nolint:forbidigo
	}
	return n
}
