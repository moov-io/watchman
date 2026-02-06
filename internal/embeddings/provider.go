package embeddings

import (
	"context"
)

// EmbeddingProvider defines the interface for embedding generation backends.
// Implementations can use remote APIs (OpenAI, Ollama, OpenRouter, etc.)
type EmbeddingProvider interface {
	// Embed generates embeddings for a batch of texts.
	// Returns L2-normalized vectors suitable for cosine similarity.
	// The returned vectors MUST be L2-normalized.
	Embed(ctx context.Context, texts []string) ([][]float64, error)

	// Dimension returns the embedding dimension for this provider.
	Dimension() int

	// Name returns the provider name for logging/telemetry.
	Name() string

	// Close releases any resources held by the provider.
	Close() error
}
