package search

import (
	"context"
	"fmt"
	"sync"

	"github.com/moov-io/watchman/internal/largest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"go4.org/syncutil"
)

type Service interface {
	Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]SearchedEntity[search.Value], error)
}

func NewService(logger log.Logger, entities []search.Entity[search.Value]) Service {
	fmt.Printf("v2search NewService(%d entity types)\n", len(entities)) //nolint:forbidigo

	gate := syncutil.NewGate(100) // TODO(adam):

	return &service{
		logger:   logger,
		entities: entities,
		Gate:     gate,
	}
}

type service struct {
	logger   log.Logger
	entities []search.Entity[search.Value]

	sync.RWMutex   // protects entities
	*syncutil.Gate // limits concurrent processing
}

func (s *service) Search(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]SearchedEntity[search.Value], error) {
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
}

func (s *service) performSearch(ctx context.Context, query search.Entity[search.Value], opts SearchOpts) ([]SearchedEntity[search.Value], error) {
	items := largest.NewItems(opts.Limit, opts.MinMatch)

	indices := makeIndices(len(s.entities), opts.Limit/3) // limit goroutines

	var wg sync.WaitGroup
	wg.Add(len(indices))

	for idx := range indices {
		start := idx
		var end int

		// We don't have another group
		if start+1 >= len(indices) {
			end = len(s.entities)
		} else {
			end = indices[start+1]
		}

		go func() {
			defer wg.Done()

			performSubSearch(items, query, s.entities[indices[start]:end])
		}()
	}
	wg.Wait()

	results := items.Items()
	var out []SearchedEntity[search.Value]

	for _, entity := range results {
		if entity == nil || entity.Value == nil {
			continue
		}

		if entity.Weight <= 0.001 {
			continue
		}

		out = append(out, SearchedEntity[search.Value]{
			Entity: entity.Value.(search.Entity[search.Value]),
			Match:  entity.Weight,
		})
	}

	return out, nil
}

func performSubSearch(items *largest.Items, query search.Entity[search.Value], entities []search.Entity[search.Value]) {
	for _, entity := range entities {
		score := search.DebugSimilarity(nil, query, entity)

		if entity.Name == "HYDRA MARKET" {
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

func makeIndices(total, groups int) []int {
	if groups == 1 || groups >= total {
		return []int{total}
	}
	xs := []int{0}
	i := 0
	for {
		if i > total {
			break
		}
		i += total / groups
		if i < total {
			xs = append(xs, i)
		}
	}
	return append(xs, total)
}
