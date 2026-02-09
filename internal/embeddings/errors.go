package embeddings

import "errors"

var (
	// ErrInvalidThreshold indicates the similarity threshold is out of range.
	ErrInvalidThreshold = errors.New("embeddings: similarity threshold must be between 0 and 1")

	// ErrInvalidBatchSize indicates the batch size is invalid.
	ErrInvalidBatchSize = errors.New("embeddings: batch size must be at least 1")

	// ErrInvalidCacheSize indicates the cache size is invalid.
	ErrInvalidCacheSize = errors.New("embeddings: cache size must be non-negative")

	// ErrIndexNotBuilt indicates the search index has not been built.
	ErrIndexNotBuilt = errors.New("embeddings: search index not built")

	// ErrServiceDisabled indicates the embeddings service is disabled.
	ErrServiceDisabled = errors.New("embeddings: service is disabled")

	// ErrBaseURLRequired indicates the provider base URL was not specified.
	ErrBaseURLRequired = errors.New("embeddings: provider base URL is required")

	// ErrModelRequired indicates the provider model name was not specified.
	ErrModelRequired = errors.New("embeddings: provider model name is required")

	// ErrInvalidDimension indicates the embedding dimension is invalid.
	ErrInvalidDimension = errors.New("embeddings: embedding dimension must be positive")

	// ErrProviderFailure indicates a provider request failed.
	ErrProviderFailure = errors.New("embeddings: provider request failed")

	// ErrDimensionMismatch indicates the embedding dimension does not match configuration.
	ErrDimensionMismatch = errors.New("embeddings: embedding dimension mismatch")

	// ErrInvalidResponse indicates an invalid response from the provider.
	ErrInvalidResponse = errors.New("embeddings: invalid response from provider")
)
