package embeddings

import (
	"context"
	"fmt"
)

// Provider defines the interface for embedding generation backends.
// Implementations can use remote APIs (OpenAI, Ollama, OpenRouter, etc.)
type Provider interface {
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

// createProvider creates an embedding provider based on configuration.
func createProvider(config ProviderConfig) (Provider, error) {
	switch config.Name {
	case "openai", "ollama", "openrouter", "azure", "":
		// All use OpenAI-compatible API format
		return NewOpenRouterProvider(config)
	case "mock":
		return NewMockProvider(config.Dimension), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", config.Name)
	}
}
