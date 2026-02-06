package embeddings

import (
	"os"
	"strconv"
	"time"
)

// Config holds configuration for the embeddings service.
type Config struct {
	// Enabled determines if embedding-based search is active.
	// When false, the service returns nil and falls back to Jaro-Winkler.
	Enabled bool `json:"enabled"`

	// Provider configuration for the embedding API.
	Provider ProviderConfig `json:"provider"`

	// CacheSize is the maximum number of embeddings to cache in memory.
	// Helps reduce latency for repeated queries.
	CacheSize int `json:"cacheSize"`

	// CrossScriptOnly when true, embeddings are only used for non-Latin queries.
	// Latin-only queries fall back to Jaro-Winkler for better performance.
	CrossScriptOnly bool `json:"crossScriptOnly"`

	// SimilarityThreshold is the minimum cosine similarity to consider a match.
	// Range: 0.0 to 1.0
	SimilarityThreshold float64 `json:"similarityThreshold"`

	// BatchSize is the number of texts to encode in a single API call.
	// Larger batches are more efficient but use more memory.
	BatchSize int `json:"batchSize"`

	// IndexBuildTimeout is the maximum time allowed for building the index.
	IndexBuildTimeout time.Duration `json:"indexBuildTimeout"`
}

// ProviderConfig holds settings for an embedding provider.
type ProviderConfig struct {
	// Name of the provider: "openai", "ollama", "openrouter", "azure"
	// All providers use OpenAI-compatible API format.
	// Default: "ollama"
	Name string `json:"name"`

	// BaseURL is the API endpoint.
	// Examples:
	//   - Ollama: "http://localhost:11434/v1"
	//   - OpenAI: "https://api.openai.com/v1"
	//   - OpenRouter: "https://openrouter.ai/api/v1"
	//   - Azure: "https://{resource}.openai.azure.com/openai/deployments/{deployment}"
	BaseURL string `json:"baseURL"`

	// APIKey for authentication. Optional for local providers (Ollama).
	// Can also be set via EMBEDDINGS_API_KEY environment variable.
	APIKey string `json:"apiKey,omitempty"`

	// Model name to use for embeddings.
	// Examples:
	//   - Ollama: "nomic-embed-text", "mxbai-embed-large"
	//   - OpenAI: "text-embedding-3-small", "text-embedding-3-large"
	//   - OpenRouter: "openai/text-embedding-3-small"
	Model string `json:"model"`

	// Dimension of the embedding vectors.
	// Must match the model's output dimension.
	// Common values: 384 (MiniLM), 768 (nomic-embed-text), 1536 (OpenAI small), 3072 (OpenAI large)
	Dimension int `json:"dimension"`

	// NormalizeVectors determines if vectors should be L2-normalized after API response.
	// Set to false if the API already returns normalized vectors (e.g., OpenAI).
	// Set to true for providers that return unnormalized vectors (e.g., some Ollama models).
	NormalizeVectors bool `json:"normalizeVectors"`

	// Timeout for API requests. Default: 30s
	Timeout time.Duration `json:"timeout"`

	// RateLimit configuration for the embedding API.
	RateLimit RateLimitConfig `json:"rateLimit"`

	// Retry configuration for failed requests.
	Retry RetryConfig `json:"retry"`

	// Headers allows adding custom HTTP headers to requests.
	// Useful for authentication or routing (e.g., OpenRouter HTTP-Referer).
	Headers map[string]string `json:"headers,omitempty"`
}

// RateLimitConfig controls the rate of API requests.
type RateLimitConfig struct {
	// RequestsPerSecond defines the sustained request rate.
	// Default: 10
	RequestsPerSecond float64 `json:"requestsPerSecond"`

	// Burst allows temporary exceeding of the rate limit.
	// Default: 20
	Burst int `json:"burst"`
}

// RetryConfig controls retry behavior for failed requests.
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts.
	// Default: 3
	MaxRetries int `json:"maxRetries"`

	// InitialBackoff is the initial backoff duration.
	// Default: 1s
	InitialBackoff time.Duration `json:"initialBackoff"`

	// MaxBackoff is the maximum backoff duration.
	// Default: 30s
	MaxBackoff time.Duration `json:"maxBackoff"`
}

// DefaultConfig returns a Config with sensible defaults.
// Note: Model and Dimension must be explicitly configured when Enabled=true.
// No default model is provided because cross-script matching quality varies
// significantly between models. Users should choose based on their requirements.
func DefaultConfig() Config {
	return Config{
		Enabled: false, // Opt-in feature, requires explicit model configuration
		Provider: ProviderConfig{
			Name: "mock",
			// Model and Dimension intentionally left empty - must be configured
			NormalizeVectors: true,
			Timeout:          30 * time.Second,
			RateLimit: RateLimitConfig{
				RequestsPerSecond: 10,
				Burst:             20,
			},
			Retry: RetryConfig{
				MaxRetries:     3,
				InitialBackoff: 1 * time.Second,
				MaxBackoff:     30 * time.Second,
			},
		},
		CacheSize:           10000,
		CrossScriptOnly:     true, // Hybrid approach: embeddings for cross-script only
		SimilarityThreshold: 0.70,
		BatchSize:           32,
		IndexBuildTimeout:   10 * time.Minute,
	}
}

// LoadFromEnv applies environment variable overrides to the configuration.
func (c *Config) LoadFromEnv() {
	if apiKey := os.Getenv("EMBEDDINGS_API_KEY"); apiKey != "" {
		c.Provider.APIKey = apiKey
	}
	if baseURL := os.Getenv("EMBEDDINGS_BASE_URL"); baseURL != "" {
		c.Provider.BaseURL = baseURL
	}
	if model := os.Getenv("EMBEDDINGS_MODEL"); model != "" {
		c.Provider.Model = model
	}
	if dim := os.Getenv("EMBEDDINGS_DIMENSION"); dim != "" {
		if d, err := strconv.Atoi(dim); err == nil && d > 0 {
			c.Provider.Dimension = d
		}
	}
}

// Validate checks the configuration for errors.
func (c Config) Validate() error {
	if !c.Enabled {
		return nil // No validation needed when disabled
	}

	if c.Provider.BaseURL == "" {
		return ErrBaseURLRequired
	}

	if c.Provider.Model == "" {
		return ErrModelRequired
	}

	if c.Provider.Dimension <= 0 {
		return ErrInvalidDimension
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
