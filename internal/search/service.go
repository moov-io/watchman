package search

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/concurrencychamp"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/indices"
	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/internal/minmaxmed"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	LatestStats() download.Stats
	UpdateEntities(stats download.Stats)

	Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error)
}

func NewService(logger log.Logger, config Config) (Service, error) {
	cm, err := concurrencychamp.NewConcurrencyManager(config.Goroutines.Default, config.Goroutines.Min, config.Goroutines.Max)
	if err != nil {
		return nil, fmt.Errorf("creating search service: %w", err)
	}
	return &service{
		logger: logger,
		config: config,
		cm:     cm,
	}, nil
}

type service struct {
	logger log.Logger
	config Config

	latestStats  download.Stats
	sync.RWMutex // protects latestStats (which has entities and list hashes)

	cm *concurrencychamp.ConcurrencyManager
}

func (s *service) LatestStats() download.Stats {
	// Grab a read-lock over our data
	s.RLock()
	defer s.RUnlock()

	// Only bring over what fields we need
	out := download.Stats{
		Lists:      s.latestStats.Lists,
		ListHashes: s.latestStats.ListHashes,
		StartedAt:  s.latestStats.StartedAt,
		EndedAt:    s.latestStats.EndedAt,
		Version:    watchman.Version,
	}
	return out
}

func (s *service) UpdateEntities(stats download.Stats) {
	s.Lock()
	defer s.Unlock()

	s.latestStats = stats
}

func (s *service) Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
	ctx, span := telemetry.StartSpan(ctx, "search", trace.WithAttributes(
		attribute.String("entity.type", string(query.Type)),
	))
	defer span.End()

	// Grab a read-lock over our data
	s.RLock()
	defer s.RUnlock()

	out, err := s.performSearch(ctx, query, opts)
	if err != nil {
		return nil, fmt.Errorf("v2 search: %w", err)
	}
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
		return nil, fmt.Errorf("getGoroutineCount: %w", err)
	}
	start := time.Now()

	indices.ProcessSliceFn(s.latestStats.Entities, goroutineCount, func(index search.Entity[search.Value]) {
		start := time.Now()

		var score float64
		if !opts.Debug {
			score = search.Similarity(query, index)
		} else {
			var buf bytes.Buffer
			buf.Grow(1700) // approximate size of debug logs

			scores := search.DetailedSimilarity(&buf, query, index)
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
			return 0, fmt.Errorf("parsing SEARCH_GOROUTINE_COUNT=%q failed: %v", fromEnv, err)
		}
		return int(n), nil
	}
	return cm.PickConcurrency(), nil
}
