//go:build embeddings

package embeddings

import (
	"context"
	"fmt"
	"sync"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// SearchResult represents a single search result from embedding similarity search.
type SearchResult struct {
	ID    string  // Entity ID from the sanctions list
	Name  string  // Original name that was indexed
	Score float64 // Cosine similarity score (0.0 to 1.0)
}

// Service provides semantic embedding-based name matching.
// It uses neural network embeddings to find similar names across different scripts
// (Arabic, Cyrillic, Chinese, etc.) and Latin text.
type Service interface {
	// Encode converts text to a normalized embedding vector.
	// The returned vector has dimensions matching the configured provider.
	Encode(ctx context.Context, text string) ([]float32, error)

	// EncodeBatch encodes multiple texts efficiently in a single API call.
	EncodeBatch(ctx context.Context, texts []string) ([][]float32, error)

	// BuildIndex creates a searchable index from entity names.
	// This must be called before Search() can be used.
	// ids and names must have the same length.
	BuildIndex(ctx context.Context, names []string, ids []string) error

	// Search finds similar names using vector similarity.
	// Returns up to k results sorted by similarity score (highest first).
	Search(ctx context.Context, query string, k int) ([]SearchResult, error)

	// Similarity computes cosine similarity between two texts.
	// Returns a value between 0.0 (dissimilar) and 1.0 (identical).
	Similarity(ctx context.Context, text1, text2 string) (float64, error)

	// ShouldUseEmbeddings determines if a query should use embeddings vs Jaro-Winkler.
	// When CrossScriptOnly is enabled, returns true only for non-Latin queries.
	ShouldUseEmbeddings(query string) bool

	// IndexSize returns the number of items in the search index.
	IndexSize() int

	// Shutdown releases resources held by the service.
	Shutdown()
}

// service implements the Service interface.
type service struct {
	logger log.Logger
	config Config

	provider EmbeddingProvider
	index    *vectorIndex
	cache    *embeddingCache

	mu sync.RWMutex
}

// NewService creates a new embeddings service.
// Returns nil if config.Enabled is false.
func NewService(logger log.Logger, config Config) (Service, error) {
	if !config.Enabled {
		return nil, nil
	}

	// Apply environment variable overrides
	config.LoadFromEnv()

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("embeddings: invalid config: %w", err)
	}

	ctx, span := telemetry.StartSpan(context.Background(), "embeddings-setup", trace.WithAttributes(
		attribute.String("provider", config.Provider.Name),
		attribute.String("model", config.Provider.Model),
		attribute.Int("dimension", config.Provider.Dimension),
		attribute.Int("cache_size", config.CacheSize),
		attribute.Bool("cross_script_only", config.CrossScriptOnly),
	))
	defer span.End()

	// Create provider based on configuration
	provider, err := createProvider(config.Provider)
	if err != nil {
		return nil, fmt.Errorf("embeddings: failed to create provider: %w", err)
	}

	logger.Info().Logf("embeddings: using %s provider", provider.Name())

	// Create cache
	cache, err := newCache(config.CacheSize)
	if err != nil {
		return nil, fmt.Errorf("embeddings: failed to create cache: %w", err)
	}

	logger.Info().Logf("embeddings: service initialized (provider=%s, dimension=%d, cache_size=%d, cross_script_only=%v)",
		provider.Name(), provider.Dimension(), config.CacheSize, config.CrossScriptOnly)

	_ = ctx // silence unused variable warning

	return &service{
		logger:   logger,
		config:   config,
		provider: provider,
		index:    newVectorIndex(provider.Dimension()),
		cache:    cache,
	}, nil
}

// createProvider creates an embedding provider based on configuration.
func createProvider(config ProviderConfig) (EmbeddingProvider, error) {
	switch config.Name {
	case "openai", "ollama", "openrouter", "azure", "":
		// All use OpenAI-compatible API format
		return NewOpenAIProvider(config)
	case "mock":
		return NewMockProvider(config.Dimension), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", config.Name)
	}
}

// Encode converts text to a normalized embedding vector.
func (s *service) Encode(ctx context.Context, text string) ([]float32, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-encode")
	defer span.End()

	// Check cache first
	if emb, ok := s.cache.Get(text); ok {
		span.SetAttributes(attribute.Bool("cache_hit", true))
		return emb, nil
	}

	span.SetAttributes(attribute.Bool("cache_hit", false))

	// Encode with provider
	embeddings, err := s.provider.Embed(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	if len(embeddings) == 0 {
		return nil, ErrInvalidResponse
	}

	// Validate dimension
	if len(embeddings[0]) != s.provider.Dimension() {
		return nil, fmt.Errorf("%w: got %d, expected %d",
			ErrDimensionMismatch, len(embeddings[0]), s.provider.Dimension())
	}

	// Cache the result
	s.cache.Put(text, embeddings[0])

	return embeddings[0], nil
}

// EncodeBatch encodes multiple texts efficiently.
func (s *service) EncodeBatch(ctx context.Context, texts []string) ([][]float32, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-encode-batch", trace.WithAttributes(
		attribute.Int("batch_size", len(texts)),
	))
	defer span.End()

	// Check cache for all texts
	result := make([][]float32, len(texts))
	uncachedIndices := make([]int, 0, len(texts))
	uncachedTexts := make([]string, 0, len(texts))

	for i, text := range texts {
		if emb, ok := s.cache.Get(text); ok {
			result[i] = emb
		} else {
			uncachedIndices = append(uncachedIndices, i)
			uncachedTexts = append(uncachedTexts, text)
		}
	}

	span.SetAttributes(
		attribute.Int("cache_hits", len(texts)-len(uncachedTexts)),
		attribute.Int("cache_misses", len(uncachedTexts)),
	)

	// Encode uncached texts
	if len(uncachedTexts) > 0 {
		embeddings, err := s.provider.Embed(ctx, uncachedTexts)
		if err != nil {
			return nil, err
		}

		// Fill in results and cache
		for i, idx := range uncachedIndices {
			result[idx] = embeddings[i]
			s.cache.Put(uncachedTexts[i], embeddings[i])
		}
	}

	return result, nil
}

// BuildIndex creates a searchable index from entity names.
func (s *service) BuildIndex(ctx context.Context, names []string, ids []string) error {
	// Apply timeout from config
	var cancel context.CancelFunc
	if s.config.IndexBuildTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, s.config.IndexBuildTimeout)
		defer cancel()
	}

	ctx, span := telemetry.StartSpan(ctx, "embeddings-build-index", trace.WithAttributes(
		attribute.Int("entity_count", len(names)),
	))
	defer span.End()

	if len(names) != len(ids) {
		return fmt.Errorf("embeddings: names and ids must have same length")
	}

	s.logger.Info().Logf("embeddings: building index for %d entities", len(names))

	// Encode all names in batches
	batchSize := s.config.BatchSize
	allEmbeddings := make([][]float32, 0, len(names))

	for i := 0; i < len(names); i += batchSize {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return fmt.Errorf("embeddings: build index cancelled: %w", ctx.Err())
		default:
		}

		end := i + batchSize
		if end > len(names) {
			end = len(names)
		}

		batch := names[i:end]
		embeddings, err := s.provider.Embed(ctx, batch)
		if err != nil {
			return fmt.Errorf("embeddings: failed to encode batch %d: %w", i/batchSize, err)
		}

		allEmbeddings = append(allEmbeddings, embeddings...)

		// Log progress for large indexes
		if len(names) > 1000 && (i+batchSize)%(batchSize*10) == 0 {
			s.logger.Info().Logf("embeddings: indexed %d/%d entities", i+batchSize, len(names))
		}
	}

	// Build the index
	s.mu.Lock()
	s.index = newVectorIndex(s.provider.Dimension())
	s.index.Add(allEmbeddings, ids, names)
	s.mu.Unlock()

	s.logger.Info().Logf("embeddings: index built with %d entities", len(names))

	return nil
}

// Search finds similar names using vector similarity.
func (s *service) Search(ctx context.Context, query string, k int) ([]SearchResult, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-search", trace.WithAttributes(
		attribute.String("query", query),
		attribute.Int("k", k),
	))
	defer span.End()

	// Encode query BEFORE acquiring lock to avoid blocking index during inference
	queryEmb, err := s.Encode(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("embeddings: failed to encode query: %w", err)
	}

	// Now acquire lock for index search
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.index.Size() == 0 {
		return nil, ErrIndexNotBuilt
	}

	// Search index
	results := s.index.Search(queryEmb, k)

	span.SetAttributes(attribute.Int("results_count", len(results)))

	return results, nil
}

// Similarity computes cosine similarity between two texts.
func (s *service) Similarity(ctx context.Context, text1, text2 string) (float64, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-similarity")
	defer span.End()

	// Encode both texts
	emb1, err := s.Encode(ctx, text1)
	if err != nil {
		return 0, err
	}

	emb2, err := s.Encode(ctx, text2)
	if err != nil {
		return 0, err
	}

	// Compute cosine similarity (vectors are L2-normalized, so dot product = cosine)
	return dotProduct(emb1, emb2), nil
}

// ShouldUseEmbeddings determines if a query should use embeddings.
func (s *service) ShouldUseEmbeddings(query string) bool {
	if !s.config.CrossScriptOnly {
		return true // Always use embeddings
	}
	return IsNonLatin(query)
}

// IndexSize returns the number of items in the search index.
func (s *service) IndexSize() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.index.Size()
}

// Shutdown releases resources held by the service.
func (s *service) Shutdown() {
	s.logger.Info().Log("embeddings: shutting down")
	if s.provider != nil {
		s.provider.Close()
	}
}
