package embeddings

import (
	"context"
	"fmt"
	"sync"

	"github.com/moov-io/watchman/internal/db"

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
	Encode(ctx context.Context, text string) ([]float64, error)

	// EncodeBatch encodes multiple texts efficiently in a single API call.
	EncodeBatch(ctx context.Context, texts []string) ([][]float64, error)

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

	provider Provider
	index    *vectorIndex
	cache    Cache

	mu sync.RWMutex
}

// NewService creates a new embeddings service.
// Returns nil if config.Enabled is false.
func NewService(logger log.Logger, config Config, database db.DB) (Service, error) {
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
		attribute.String("cache_type", config.Cache.Type),
		attribute.Int("cache_size", config.Cache.Size),
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
	cache, err := NewCache(ctx, config, database)
	if err != nil {
		return nil, fmt.Errorf("embeddings: failed to create cache: %w", err)
	}

	logger.Info().Logf("embeddings: service initialized (provider=%s, dimension=%d, %s_cache_size=%d, cross_script_only=%v)",
		provider.Name(), provider.Dimension(), config.Cache.Type, config.Cache.Size, config.CrossScriptOnly)

	return &service{
		logger:   logger,
		config:   config,
		provider: provider,
		index:    newVectorIndex(provider.Dimension()),
		cache:    cache,
	}, nil
}

// Encode converts text to a normalized embedding vector.
func (s *service) Encode(ctx context.Context, text string) ([]float64, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-encode")
	defer span.End()

	// Check cache first
	if emb, ok := s.cache.Get(ctx, text); ok {
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
	s.cache.Put(ctx, text, embeddings[0])

	return embeddings[0], nil
}

// EncodeBatch encodes multiple texts efficiently.
func (s *service) EncodeBatch(ctx context.Context, texts []string) ([][]float64, error) {
	_, span := telemetry.StartSpan(ctx, "embeddings-encode-batch", trace.WithAttributes(
		attribute.Int("batch_size", len(texts)),
	))
	defer span.End()

	// Check cache for all texts
	result := make([][]float64, len(texts))
	uncachedIndices := make([]int, 0, len(texts))
	uncachedTexts := make([]string, 0, len(texts))

	for i, text := range texts {
		if emb, ok := s.cache.Get(ctx, text); ok {
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
			s.cache.Put(ctx, uncachedTexts[i], embeddings[i])
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

	// Check cache for all names first
	allEmbeddings := make([][]float64, len(names))
	uncachedIndices := make([]int, 0)
	uncachedNames := make([]string, 0)

	for i, name := range names {
		if emb, ok := s.cache.Get(ctx, name); ok {
			allEmbeddings[i] = emb
		} else {
			uncachedIndices = append(uncachedIndices, i)
			uncachedNames = append(uncachedNames, name)
		}
	}

	cacheHits := len(names) - len(uncachedNames)
	span.SetAttributes(
		attribute.Int("cache_hits", cacheHits),
		attribute.Int("cache_misses", len(uncachedNames)),
	)

	if cacheHits > 0 {
		s.logger.Info().Logf("embeddings: cache hits: %d/%d (%.1f%%)",
			cacheHits, len(names), float64(cacheHits)/float64(len(names))*100)
	}

	// Encode uncached names in batches
	if len(uncachedNames) > 0 {
		batchSize := s.config.BatchSize

		for i := 0; i < len(uncachedNames); i += batchSize {
			// Check for context cancellation
			select {
			case <-ctx.Done():
				return fmt.Errorf("embeddings: build index cancelled: %w", ctx.Err())
			default:
			}

			end := i + batchSize
			if end > len(uncachedNames) {
				end = len(uncachedNames)
			}

			batch := uncachedNames[i:end]
			embeddings, err := s.provider.Embed(ctx, batch)
			if err != nil {
				return fmt.Errorf("embeddings: failed to encode batch %d: %w", i/batchSize, err)
			}

			// Fill in results and cache
			for j, emb := range embeddings {
				originalIdx := uncachedIndices[i+j]
				allEmbeddings[originalIdx] = emb
				s.cache.Put(ctx, uncachedNames[i+j], emb)
			}

			// Log progress for large indexes
			if len(uncachedNames) > 1000 && (i+batchSize)%(batchSize*10) == 0 {
				s.logger.Info().Logf("embeddings: encoded %d/%d uncached entities", i+batchSize, len(uncachedNames))
			}
		}
	}

	// Build the index
	s.mu.Lock()
	s.index = newVectorIndex(s.provider.Dimension())
	s.index.Add(allEmbeddings, ids, names)
	s.mu.Unlock()

	s.logger.Info().Logf("embeddings: index built with %d entities (%d from cache, %d newly encoded)",
		len(names), cacheHits, len(uncachedNames))

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
	results, err := s.index.Search(queryEmb, k)
	if err != nil {
		return nil, err
	}

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
