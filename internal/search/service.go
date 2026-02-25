package search

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/concurrencychamp"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/embeddings"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/indices"
	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/internal/minmaxmed"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// embeddingSearchLimitMultiplier is used to fetch extra results from embedding search
// to account for filtering (by type, threshold, etc.) before returning final results.
const embeddingSearchLimitMultiplier = 2

type Service interface {
	LatestStats() download.Stats

	Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error)

	// RebuildEmbeddingIndex rebuilds the embedding index from current entities.
	// This should be called after the entity list has been updated.
	RebuildEmbeddingIndex(ctx context.Context) error
}

func NewService(logger log.Logger, config Config, database db.DB, indexedLists index.Lists) (Service, error) {
	cm, err := concurrencychamp.NewConcurrencyManager(config.Goroutines.Default, config.Goroutines.Min, config.Goroutines.Max)
	if err != nil {
		return nil, fmt.Errorf("creating search service: %w", err)
	}

	// Initialize embeddings service (optional, for cross-script matching)
	var embeddingsSvc embeddings.Service
	if config.Embeddings.Enabled {
		embeddingsSvc, err = embeddings.NewService(logger, config.Embeddings, database)
		if err != nil {
			return nil, fmt.Errorf("creating embeddings service: %w", err)
		}
	}

	return &service{
		logger:       logger,
		config:       config,
		indexedLists: indexedLists,
		cm:           cm,
		embeddings:   embeddingsSvc,
	}, nil
}

type service struct {
	logger log.Logger
	config Config

	indexedLists index.Lists
	embeddings   embeddings.Service

	cm *concurrencychamp.ConcurrencyManager
}

func (s *service) LatestStats() download.Stats {
	return s.indexedLists.LatestStats()
}

func (s *service) Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
	ctx, span := telemetry.StartSpan(ctx, "search", trace.WithAttributes(
		attribute.String("query.type", string(query.Type)),
		attribute.String("query.source", string(query.Source)),
		attribute.String("query.source_id", string(query.SourceID)),
		attribute.String("request_id", opts.RequestID),
		attribute.Bool("query.debug", opts.Debug),
		attribute.StringSlice("query.debug_source_ids", opts.DebugSourceIDs),
	))
	defer span.End()

	// Check if we should use embedding-based search for cross-script queries
	if s.shouldUseEmbeddings(query.Name) {
		span.SetAttributes(attribute.Bool("search.use_embeddings", true))
		out, err := s.performEmbeddingSearch(ctx, query, opts)
		if err != nil {
			// Fall back to Jaro-Winkler on embedding search failure
			s.logger.Error().Logf("embedding search failed, falling back to Jaro-Winkler: %v", err)
		} else {
			return out, nil
		}
	}

	span.SetAttributes(attribute.Bool("search.use_embeddings", false))
	out, err := s.performSearch(ctx, query, opts)
	if err != nil {
		s.logger.Error().Logf("v2 search failed: %v", err)
		return nil, fmt.Errorf("v2 search: %w", err)
	}
	return out, nil
}

// shouldUseEmbeddings determines if the query should use embedding-based search.
// Returns true for non-Latin scripts (Arabic, Cyrillic, Chinese, etc.) when embeddings are enabled.
func (s *service) shouldUseEmbeddings(queryName string) bool {
	if s.embeddings == nil {
		return false
	}
	return s.embeddings.ShouldUseEmbeddings(queryName)
}

// performEmbeddingSearch executes a search using neural embeddings.
// This is used for cross-script matching (e.g., Arabic query -> Latin results).
func (s *service) performEmbeddingSearch(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
	_, span := telemetry.StartSpan(ctx, "perform-embedding-search", trace.WithAttributes(
		attribute.Int("opts.limit", opts.Limit),
		attribute.Float64("opts.min_match", opts.MinMatch),
	))
	defer span.End()

	// Search using embeddings (fetch extra to account for filtering)
	results, err := s.embeddings.Search(ctx, query.Name, opts.Limit*embeddingSearchLimitMultiplier)
	if err != nil {
		s.logger.Error().Logf("embedding search failed: %v", err)
		return nil, fmt.Errorf("embedding search: %w", err)
	}

	// Get all entities for lookup
	searchEntities, err := s.indexedLists.GetEntities(ctx, query.Source)
	if err != nil {
		s.logger.Error().Logf("getting indexed entities failed: %v", err)
		return nil, fmt.Errorf("getting indexed entities: %w", err)
	}

	// Build map for fast entity lookup by SourceID
	entityMap := make(map[string]search.Entity[search.Value], len(searchEntities))
	for _, e := range searchEntities {
		entityMap[e.SourceID] = e
	}

	// Convert results to SearchedEntity
	var out []search.SearchedEntity[search.Value]
	for _, result := range results {
		if result.Score < opts.MinMatch {
			continue
		}

		entity, ok := entityMap[result.ID]
		if !ok {
			continue
		}

		// Apply type filter if specified
		if query.Type != "" && query.Type != entity.Type {
			continue
		}

		out = append(out, search.SearchedEntity[search.Value]{
			Entity: entity,
			Match:  result.Score,
		})

		if len(out) >= opts.Limit {
			break
		}
	}

	span.SetAttributes(attribute.Int("results_count", len(out)))

	return out, nil
}

type SearchOpts struct {
	Limit    int
	MinMatch float64

	RequestID      string
	Debug          bool
	DebugSourceIDs []string
}

type debugRespone struct {
	scores search.SimilarityScore
	buffer *bytes.Buffer
}

func (s *service) performSearch(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
	_, span := telemetry.StartSpan(ctx, "perform-search", trace.WithAttributes(
		attribute.Int("opts.limit", opts.Limit),
		attribute.Float64("opts.min_match", opts.MinMatch),
	))
	defer span.End()

	stats := minmaxmed.New(10) // window size
	items := largest.NewItems[search.Entity[search.Value]](opts.Limit, opts.MinMatch)

	var debugs *largest.Items[debugRespone]
	if opts.Debug {
		debugs = largest.NewItems[debugRespone](opts.Limit, opts.MinMatch)
	}

	goroutineCount, err := getGoroutineCount(s.cm)
	if err != nil {
		s.logger.Error().Logf("getGoroutineCount failed: %v", err)
		return nil, fmt.Errorf("getGoroutineCount: %w", err)
	}
	start := time.Now()

	// Check if the query is targeting ingested files
	searchEntities, err := s.indexedLists.GetEntities(ctx, query.Source)
	if err != nil {
		s.logger.Error().Logf("getting indexed entities failed: %v", err)
		return nil, fmt.Errorf("getting indexed entities: %w", err)
	}

	// Get TF-IDF index for weighted name matching
	tfidfIndex := s.indexedLists.GetTFIDFIndex()

	indices.ProcessSliceFn(searchEntities, goroutineCount, func(index search.Entity[search.Value]) {
		start := time.Now()

		var score float64
		if !opts.Debug {
			score = search.SimilarityWithTFIDF(query, index, tfidfIndex)
		} else {
			var buf bytes.Buffer
			buf.Grow(1700) // approximate size of debug logs

			scores := search.DetailedSimilarityWithTFIDF(&buf, query, index, tfidfIndex)
			score = scores.FinalScore

			// Add debug buffer to be stored
			debugs.Add(largest.Item[debugRespone]{
				Value: debugRespone{
					scores: scores,
					buffer: &buf,
				},
				Weight: score,
			})
		}
		stats.AddDuration(time.Since(start))

		items.Add(largest.Item[search.Entity[search.Value]]{
			Value:  index,
			Weight: score,
		})
	})

	diff := time.Since(start)
	s.cm.RecordDuration(goroutineCount, diff)

	span.SetAttributes(
		attribute.Int("index.searched_entities", len(searchEntities)),
		attribute.Int("search.goroutine_count", goroutineCount),
		attribute.Int64("search.duration", diff.Milliseconds()),
	)

	// After processing the list add stats to the span
	stats.AddEvent(span)

	results := items.Items()
	debugLogs := debugs.Items()
	var out []search.SearchedEntity[search.Value]

	for idx, res := range results {
		if res.Value.SourceID == "" || res.Weight <= 0.001 {
			continue
		}

		searched := search.SearchedEntity[search.Value]{
			Entity: res.Value,
			Match:  res.Weight,
		}

		if len(debugLogs) > idx {
			scores := debugLogs[idx].Value
			searched.Details = scores.scores

			if scores.buffer != nil {
				searched.Debug = base64.StdEncoding.EncodeToString(scores.buffer.Bytes())
			}
		}

		out = append(out, searched)
	}

	return out, nil
}

func getGoroutineCount(cm *concurrencychamp.ConcurrencyManager) (int, error) {
	// After local benchmarking this is a tradeoff between the fastest / most efficient group size picking
	// and offering configurability to users.
	//
	// Using an atomic cache to store ParseUint's result is ~75% slower than just calling strconv.ParseUint every time.
	// This may be an inaccurate result on other hardware/platforms.
	//
	// Using concurrencychamp.ConcurrencyManager provides the quickest searches while using an insignificant amount of memory
	// compared to what similarity scoring uses.
	fromEnv := strings.TrimSpace(os.Getenv("SEARCH_GOROUTINE_COUNT"))
	if fromEnv != "" {
		n, err := strconv.ParseUint(fromEnv, 10, 8)
		if err != nil {
			return 0, fmt.Errorf("parsing SEARCH_GOROUTINE_COUNT=%q failed: %w", fromEnv, err)
		}
		return int(n), nil
	}
	return cm.PickConcurrency(), nil
}

// RebuildEmbeddingIndex rebuilds the embedding index from current entities.
// This extracts entity names and IDs, then calls BuildIndex on the embeddings service.
func (s *service) RebuildEmbeddingIndex(ctx context.Context) error {
	if s.embeddings == nil {
		return nil // Embeddings not enabled
	}

	ctx, span := telemetry.StartSpan(ctx, "rebuild-embedding-index")
	defer span.End()

	// Get all entities
	entities, err := s.indexedLists.GetEntities(ctx, "")
	if err != nil {
		s.logger.Error().Logf("getting entities for embedding index failed: %v", err)
		return fmt.Errorf("getting entities for embedding index: %w", err)
	}

	if len(entities) == 0 {
		s.logger.Info().Log("embeddings: no entities to index")
		return nil
	}

	// Extract names and IDs
	names := make([]string, len(entities))
	ids := make([]string, len(entities))
	for i, e := range entities {
		names[i] = e.Name
		ids[i] = e.SourceID
	}

	s.logger.Info().Logf("embeddings: rebuilding index with %d entities", len(entities))

	// Build the embedding index
	if err := s.embeddings.BuildIndex(ctx, names, ids); err != nil {
		s.logger.Error().Logf("building embedding index failed: %v", err)
		return fmt.Errorf("building embedding index: %w", err)
	}

	s.logger.Info().Logf("embeddings: index rebuilt successfully")
	return nil
}
