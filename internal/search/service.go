package search

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/groupsize"
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

func NewService(logger log.Logger) Service {
	cm, err := groupsize.NewConcurrencyManager(defaultGroupSize, 1, 100)
	if err != nil {
		panic(err)
	}
	return &service{
		logger: logger,
		cm:     cm,
	}
}

type service struct {
	logger log.Logger

	latestStats  download.Stats
	sync.RWMutex // protects latestStats (which has entities and list hashes)

	cm *groupsize.ConcurrencyManager
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
	DebugSourceIDs []string
}

func (s *service) performSearch(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
	ctx, span := telemetry.StartSpan(ctx, "perform-search", trace.WithAttributes(
		attribute.Int("opts.limit", opts.Limit),
		attribute.Float64("opts.min_match", opts.MinMatch),
	))
	defer span.End()

	stats := minmaxmed.New(10) // window size
	items := largest.NewItems(opts.Limit, opts.MinMatch)

	groupSize := getGroupSize(s.cm)
	start := time.Now()

	indices.ProcessSliceFn(s.latestStats.Entities, groupSize, func(index search.Entity[search.Value]) {
		start := time.Now()
		score := search.DebugSimilarity(nil, query, index) // TODO(adam): add proper debug functionality?
		stats.AddDuration(time.Since(start))

		items.Add(largest.Item{
			Value:  index,
			Weight: score,
		})
	})

	diff := time.Since(start)
	s.cm.RecordDuration(groupSize, diff)

	span.SetAttributes(
		attribute.Int("search.group_size", groupSize),
		attribute.Int64("search.duration", diff.Milliseconds()),
	)

	// After processing the list add stats to the span
	stats.AddEvent(span)

	results := items.Items()
	var out []search.SearchedEntity[search.Value]

	for _, res := range results {
		if res.Value.SourceID == "" || res.Weight <= 0.001 {
			continue
		}

		out = append(out, search.SearchedEntity[search.Value]{
			Entity: res.Value,
			Match:  res.Weight,
		})
	}

	return out, nil
}

const (
	defaultGroupSize = 20 // rough estimate from local testing
)

func getGroupSize(cm *groupsize.ConcurrencyManager) int {
	// After local benchmarking this is a tradeoff between the fastest / most efficient group size picking
	// and offering configurability to users.
	//
	// Using an atomic cache to store ParseUint's result is ~75% slower than just calling strconv.ParseUint every time.
	// This may be an inaccurate result on other hardware/platforms.
	//
	// Using groupsize.ConcurrencyManager provides the quickest searches while using an insignificant amount of memory
	// compared to what similarity scoring uses.
	fromEnv := strings.TrimSpace(os.Getenv("SEARCH_GROUP_COUNT"))
	if fromEnv != "" {
		n, err := strconv.ParseUint(fromEnv, 10, 8)
		if err != nil {
			panic(fmt.Sprintf("ERROR: parsing SEARCH_GROUP_COUNT=%q failed: %v", fromEnv, err)) //nolint:forbigido
		}
		return int(n)
	}
	return cm.PickConcurrency()
}
