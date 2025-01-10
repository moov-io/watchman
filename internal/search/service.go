package search

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/moov-io/watchman/internal/indices"
	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"go4.org/syncutil"
)

type Service interface {
	UpdateEntities(entities []search.Entity[search.Value])

	Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]search.SearchedEntity[search.Value], error)
}

func NewService(logger log.Logger) Service {
	gate := syncutil.NewGate(100) // TODO(adam):

	return &service{
		logger: logger,
		Gate:   gate,
	}
}

type service struct {
	logger   log.Logger
	entities []search.Entity[search.Value]

	sync.RWMutex   // protects entities
	*syncutil.Gate // limits concurrent processing
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

	// Grab a worker slot
	s.Gate.Start()
	defer s.Gate.Done()

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

	indices := indices.New(len(s.entities), opts.Limit/3) // limit goroutines

	var wg sync.WaitGroup
	wg.Add(len(indices))

	fmt.Printf("indices: %#v ", indices)

	start := time.Now()
	for idx := range indices {
		start := idx
		var end int

		// We don't have another group
		if start+1 >= len(indices) {
			end = len(s.entities)
		} else {
			end = indices[start+1]
		}
		fmt.Printf("start=%d  end=%v\n", start, end)

		go func() {
			defer wg.Done()

			performSubSearch(items, query, s.entities[indices[start]:end], opts)
		}()
	}
	wg.Wait()

	fmt.Printf("concurrent search took: %v\n", time.Since(start))
	start = time.Now()

	results := items.Items()
	var out []search.SearchedEntity[search.Value]

	for _, entity := range results {
		if entity == nil || entity.Value == nil {
			continue
		}

		if entity.Weight <= 0.001 {
			continue
		}

		out = append(out, search.SearchedEntity[search.Value]{
			Entity: entity.Value.(search.Entity[search.Value]), // TODO(adam):
			Match:  entity.Weight,
		})
	}

	fmt.Printf("result mapping took: %v\n", time.Since(start))

	return out, nil
}

func performSubSearch(items *largest.Items, query search.Entity[search.Value], entities []search.Entity[search.Value], opts SearchOpts) {
	for _, entity := range entities {
		score := search.DebugSimilarity(nil, query, entity)

		if slices.Contains(opts.DebugSourceIDs, entity.SourceID) {
			fmt.Printf("%#v\n", entity)
			fmt.Println("")
			fmt.Printf("%#v\n", entity.SourceData)
			fmt.Println("")
			fmt.Printf("%#v\n", entity.Business)
		}

		items.Add(&largest.Item{
			Value:  entity,
			Weight: score,
		})
	}
}
