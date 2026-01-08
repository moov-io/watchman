//go:build embeddings

package embeddings

import "time"

// Config holds configuration for the embeddings service.
type Config struct {
	// Enabled determines if embedding-based search is active.
	// When false, the service returns nil and falls back to Jaro-Winkler.
	Enabled bool `json:"enabled"`

	// ModelPath is the directory containing the ONNX model files.
	// Expected files: model.onnx, tokenizer.json, tokenizer_config.json
	ModelPath string `json:"modelPath"`

	// CacheSize is the maximum number of embeddings to cache in memory.
	// Helps reduce latency for repeated queries.
	CacheSize int `json:"cacheSize"`

	// CrossScriptOnly when true, embeddings are only used for non-Latin queries.
	// Latin-only queries fall back to Jaro-Winkler for better performance.
	CrossScriptOnly bool `json:"crossScriptOnly"`

	// SimilarityThreshold is the minimum cosine similarity to consider a match.
	// Range: 0.0 to 1.0
	SimilarityThreshold float64 `json:"similarityThreshold"`

	// BatchSize is the number of texts to encode in a single inference call.
	// Larger batches are more efficient but use more memory.
	BatchSize int `json:"batchSize"`

	// IndexBuildTimeout is the maximum time allowed for building the index.
	IndexBuildTimeout time.Duration `json:"indexBuildTimeout"`
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Enabled:             false, // Opt-in feature
		ModelPath:           "./models/multilingual-minilm",
		CacheSize:           10000,
		CrossScriptOnly:     true, // Hybrid approach: embeddings for cross-script only
		SimilarityThreshold: 0.70,
		BatchSize:           32,
		IndexBuildTimeout:   10 * time.Minute,
	}
}

// Validate checks the configuration for errors.
func (c Config) Validate() error {
	if !c.Enabled {
		return nil // No validation needed when disabled
	}

	if c.ModelPath == "" {
		return ErrModelPathRequired
	}

	if c.SimilarityThreshold < 0 || c.SimilarityThreshold > 1 {
		return ErrInvalidThreshold
	}

	if c.BatchSize < 1 {
		return ErrInvalidBatchSize
	}

	if c.CacheSize < 0 {
		return ErrInvalidCacheSize
	}

	return nil
}
