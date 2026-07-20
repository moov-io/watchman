package search

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/watchman/internal/concurrencychamp"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/embeddings"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// embeddingSearchLimitMultiplier is used to fetch extra results from embedding search
// to account for filtering (by type, threshold, etc.) before returning final results.
const embeddingSearchLimitMultiplier = 2

// smallCandidateBypass is the max candidate-set size that skips the search admission
// semaphore. Exact crypto/ID hits and tightly pruned name queries stay off the queue
// so they are not blocked behind full-partition scans.
const smallCandidateBypass = 100

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

	inFlight := config.MaxInFlight
	if inFlight <= 0 {
		inFlight = runtime.GOMAXPROCS(0)
		if inFlight < 1 {
			inFlight = 1
		}
	}

	return &service{
		logger:       logger,
		config:       config,
		indexedLists: indexedLists,
		cm:           cm,
		embeddings:   embeddingsSvc,
		searchSem:    make(chan struct{}, inFlight),
	}, nil
}

type service struct {
	logger log.Logger
	config Config

	indexedLists index.Lists
	embeddings   embeddings.Service

	cm *concurrencychamp.ConcurrencyManager

	// searchSem limits concurrent full-corpus searches to avoid goroutine oversubscription
	searchSem chan struct{}
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
		// Should we debug any specific entities
		if slices.Contains(opts.DebugSourceIDs, e.SourceID) {
			s.logger.Debug().With(log.Fields{
				"debug_source_id": log.String(e.SourceID),
			}).Logf("indexed entity: %#v", e)
		}

		entityMap[e.SourceID] = e
	}

	// Convert results to SearchedEntity
	var out []search.SearchedEntity[search.Value]
	for _, result := range results {
		// Should we debug any specific entities
		if slices.Contains(opts.DebugSourceIDs, result.ID) {
			s.logger.Debug().With(log.Fields{
				"debug_source_id": log.String(result.ID),
			}).Logf("embeddings results: %#v", result)
		}

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
	// Candidate selection first (cheap, RLock only) so we can skip admission control
	// for tiny result sets and avoid queuing them behind full-partition scans.
	searchEntities, err := s.indexedLists.SelectCandidates(ctx, query)
	if err != nil {
		s.logger.Error().Logf("selecting candidate entities failed: %v", err)
		return nil, fmt.Errorf("selecting candidate entities: %w", err)
	}

	// Admission control: bound concurrent large scans so worker pools do not oversubscribe CPUs.
	if len(searchEntities) > smallCandidateBypass {
		select {
		case s.searchSem <- struct{}{}:
			defer func() { <-s.searchSem }()
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	_, span := telemetry.StartSpan(ctx, "perform-search", trace.WithAttributes(
		attribute.Int("opts.limit", opts.Limit),
		attribute.Float64("opts.min_match", opts.MinMatch),
		attribute.Int("index.candidate_count", len(searchEntities)),
	))
	defer span.End()

	goroutineCount, err := getGoroutineCount(s.cm)
	if err != nil {
		s.logger.Error().Logf("getGoroutineCount failed: %v", err)
		return nil, fmt.Errorf("getGoroutineCount: %w", err)
	}
	start := time.Now()

	// Precompute query term weights once per search when TF-IDF is enabled
	tfidfIndex := s.indexedLists.GetTFIDFIndex()
	if tfidfIndex != nil && tfidfIndex.Enabled() && len(query.PreparedFields.NameFields) > 0 {
		query.PreparedFields.NameWeights = tfidfIndex.GetWeights(query.PreparedFields.NameFields)
	}

	hasDebugIDs := false
	for _, id := range opts.DebugSourceIDs {
		if id != "" {
			hasDebugIDs = true
			break
		}
	}

	// Size the worker pool to the candidate set — pruning often leaves few entities.
	if goroutineCount < 1 {
		goroutineCount = 1
	}
	numWorkers := goroutineCount
	if numWorkers <= 1 || len(searchEntities) < numWorkers {
		numWorkers = 1
	}

	items := largest.NewItems[search.Entity[search.Value]](opts.Limit, opts.MinMatch)
	var debugs *largest.Items[debugRespone]
	if opts.Debug {
		debugs = largest.NewItems[debugRespone](opts.Limit, opts.MinMatch)
	}

	if numWorkers <= 1 {
		// Common path after candidate pruning: score directly, no local heaps/merge.
		scoreEntities(searchEntities, query, tfidfIndex, opts, hasDebugIDs, s.logger, items, debugs)
	} else {
		// Per-worker local top-K (no shared mutex on the hot path), then merge
		localItems := make([]*largest.Items[search.Entity[search.Value]], numWorkers)
		var localDebugs []*largest.Items[debugRespone]
		if opts.Debug {
			localDebugs = make([]*largest.Items[debugRespone], numWorkers)
		}
		for i := 0; i < numWorkers; i++ {
			localItems[i] = largest.NewItems[search.Entity[search.Value]](opts.Limit, opts.MinMatch)
			if opts.Debug {
				localDebugs[i] = largest.NewItems[debugRespone](opts.Limit, opts.MinMatch)
			}
		}

		chunkSize := (len(searchEntities) + numWorkers - 1) / numWorkers
		var wg sync.WaitGroup
		for worker := 0; worker < numWorkers; worker++ {
			startIdx := worker * chunkSize
			if startIdx >= len(searchEntities) {
				break
			}
			endIdx := startIdx + chunkSize
			if endIdx > len(searchEntities) {
				endIdx = len(searchEntities)
			}
			wg.Add(1)
			go func(w, start, end int) {
				defer wg.Done()
				var debugLocal *largest.Items[debugRespone]
				if opts.Debug {
					debugLocal = localDebugs[w]
				}
				scoreEntities(searchEntities[start:end], query, tfidfIndex, opts, hasDebugIDs, s.logger, localItems[w], debugLocal)
			}(worker, startIdx, endIdx)
		}
		wg.Wait()

		for i := range localItems {
			items.Merge(localItems[i])
			if opts.Debug {
				debugs.Merge(localDebugs[i])
			}
		}
	}

	diff := time.Since(start)
	s.cm.RecordDuration(numWorkers, diff)

	span.SetAttributes(
		attribute.Int("index.searched_entities", len(searchEntities)),
		attribute.Int("search.goroutine_count", numWorkers),
		attribute.Int64("search.duration", diff.Milliseconds()),
	)

	results := items.Items()
	var debugLogs []largest.Item[debugRespone]
	if debugs != nil {
		debugLogs = debugs.Items()
	}
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

func scoreEntities(
	entities []search.Entity[search.Value],
	query search.Entity[search.Value],
	tfidfIndex *tfidf.Index,
	opts SearchOpts,
	hasDebugIDs bool,
	logger log.Logger,
	items *largest.Items[search.Entity[search.Value]],
	debugs *largest.Items[debugRespone],
) {
	for i := range entities {
		indexEntity := entities[i]
		isDebugEntity := hasDebugIDs && slices.Contains(opts.DebugSourceIDs, indexEntity.SourceID)

		if isDebugEntity {
			logger.Debug().With(log.Fields{
				"debug_source_id": log.String(indexEntity.SourceID),
			}).Logf("indexed entity: %#v", indexEntity)
		}

		var score float64
		if !opts.Debug {
			score = search.SimilarityWithTFIDF(query, indexEntity, tfidfIndex)
		} else {
			var buf bytes.Buffer
			buf.Grow(1700) // approximate size of debug logs

			scores := search.DebugSimilarityWithTFIDF(&buf, query, indexEntity, tfidfIndex)
			score = scores.FinalScore

			if isDebugEntity {
				logger.Debug().With(log.Fields{
					"debug_source_id": log.String(indexEntity.SourceID),
				}).Logf("similarity score: %#v", scores)

				logger.Debug().With(log.Fields{
					"debug_source_id": log.String(indexEntity.SourceID),
				}).Logf("scoring debug: %#v", buf.String())
			}

			if debugs != nil {
				debugs.AddLocal(largest.Item[debugRespone]{
					Value: debugRespone{
						scores: scores,
						buffer: &buf,
					},
					Weight: score,
				})
			}
		}

		if isDebugEntity {
			logger.Debug().With(log.Fields{
				"debug_source_id": log.String(indexEntity.SourceID),
			}).Logf("final score: %.5f", score)
		}

		items.AddLocal(largest.Item[search.Entity[search.Value]]{
			Value:  indexEntity,
			Weight: score,
		})
	}
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
