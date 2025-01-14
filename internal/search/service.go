package search

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/moov-io/watchman/internal/indices"
	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
)

type Service interface {
	UpdateEntities(entities []search.Entity[search.Value])

	Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error)
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

type service struct {
	logger   log.Logger
	entities []search.Entity[search.Value]

	sync.RWMutex // protects entities
}

func (s *service) UpdateEntities(entities []search.Entity[search.Value]) {
	s.Lock()
	defer s.Unlock()

	s.entities = entities
}

func (s *service) Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error) {
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
	items := largest.NewItems(opts.Limit, opts.MinMatch)

	indices.ProcessSliceFn(s.entities, getGroupCount(opts), func(index search.Entity[search.Value]) {
		score := search.DebugSimilarity(nil, query, index) // TODO(adam): add proper debug functionality?

		if slices.Contains(opts.DebugSourceIDs, index.SourceID) {
			// fmt.Printf("%#v\n", index)
			// fmt.Println("")
			// fmt.Printf("%#v\n", index.SourceData)
			// fmt.Println("")
			// fmt.Printf("%#v\n", index.Business)
		}

		items.Add(largest.Item{
			Value:  index,
			Weight: score,
		})
	})

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
	defaultGroupCount = 20 // rough estimate from local testing // TODO(adam): more benchmarks
)

func getGroupCount(opts SearchOpts) int {
	fromEnv := strings.TrimSpace(os.Getenv("SEARCH_GROUP_COUNT")) // not exported, used for initial testing
	if fromEnv != "" {
		n, err := strconv.ParseUint(fromEnv, 10, 8)
		if err != nil {
			panic(fmt.Sprintf("ERROR: parsing SEARCH_GROUP_COUNT=%q failed: %v", fromEnv, err)) //nolint:forbidigo
		}
		return int(n)
	}
	return defaultGroupCount
}
