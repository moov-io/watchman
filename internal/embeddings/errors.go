//go:build embeddings

package embeddings

import "errors"

var (
	// ErrModelPathRequired indicates the model path was not specified.
	ErrModelPathRequired = errors.New("embeddings: model path is required")

	// ErrInvalidThreshold indicates the similarity threshold is out of range.
	ErrInvalidThreshold = errors.New("embeddings: similarity threshold must be between 0 and 1")

	// ErrInvalidBatchSize indicates the batch size is invalid.
	ErrInvalidBatchSize = errors.New("embeddings: batch size must be at least 1")

	// ErrInvalidCacheSize indicates the cache size is invalid.
	ErrInvalidCacheSize = errors.New("embeddings: cache size must be non-negative")

	// ErrModelNotLoaded indicates the model has not been loaded.
	ErrModelNotLoaded = errors.New("embeddings: model not loaded")

	// ErrIndexNotBuilt indicates the search index has not been built.
	ErrIndexNotBuilt = errors.New("embeddings: search index not built")

	// ErrServiceDisabled indicates the embeddings service is disabled.
	ErrServiceDisabled = errors.New("embeddings: service is disabled")
)
