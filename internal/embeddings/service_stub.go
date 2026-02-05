//go:build !embeddings

package embeddings

import (
	"context"
	"time"

	"github.com/moov-io/base/log"
)

// Config holds configuration for the embeddings service.
// This is a stub when embeddings build tag is not set.
type Config struct {
	Enabled             bool           `json:"enabled"`
	Provider            ProviderConfig `json:"provider"`
	CacheSize           int            `json:"cacheSize"`
	CrossScriptOnly     bool           `json:"crossScriptOnly"`
	SimilarityThreshold float64        `json:"similarityThreshold"`
	BatchSize           int            `json:"batchSize"`
	IndexBuildTimeout   time.Duration  `json:"indexBuildTimeout"`
}

// ProviderConfig holds settings for an embedding provider.
type ProviderConfig struct {
	Name             string            `json:"name"`
	BaseURL          string            `json:"baseURL"`
	APIKey           string            `json:"apiKey,omitempty"`
	Model            string            `json:"model"`
	Dimension        int               `json:"dimension"`
	NormalizeVectors bool              `json:"normalizeVectors"`
	Timeout          time.Duration     `json:"timeout"`
	RateLimit        RateLimitConfig   `json:"rateLimit"`
	Retry            RetryConfig       `json:"retry"`
	Headers          map[string]string `json:"headers,omitempty"`
}

// RateLimitConfig controls the rate of API requests.
type RateLimitConfig struct {
	RequestsPerSecond float64 `json:"requestsPerSecond"`
	Burst             int     `json:"burst"`
}

// RetryConfig controls retry behavior for failed requests.
type RetryConfig struct {
	MaxRetries     int           `json:"maxRetries"`
	InitialBackoff time.Duration `json:"initialBackoff"`
	MaxBackoff     time.Duration `json:"maxBackoff"`
}

// DefaultConfig returns a Config with embeddings disabled.
func DefaultConfig() Config {
	return Config{Enabled: false}
}

// SearchResult represents a single search result from embedding similarity search.
type SearchResult struct {
	ID    string
	Name  string
	Score float64
}

// Service provides semantic embedding-based name matching.
// This is a stub interface when embeddings build tag is not set.
type Service interface {
	Encode(ctx context.Context, text string) ([]float32, error)
	EncodeBatch(ctx context.Context, texts []string) ([][]float32, error)
	BuildIndex(ctx context.Context, names []string, ids []string) error
	Search(ctx context.Context, query string, k int) ([]SearchResult, error)
	Similarity(ctx context.Context, text1, text2 string) (float64, error)
	ShouldUseEmbeddings(query string) bool
	IndexSize() int
	Shutdown()
}

// NewService returns nil when embeddings build tag is not set.
// This allows the search service to gracefully handle the absence of embeddings.
func NewService(logger log.Logger, config Config) (Service, error) {
	return nil, nil
}
