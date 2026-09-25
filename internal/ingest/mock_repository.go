package ingest

import (
	"context"
	"slices"
	"sort"
	"sync"

	"github.com/moov-io/watchman/pkg/search"
)

type MockRepository struct {
	Err error

	mu       sync.RWMutex
	entities map[string][]search.Entity[search.Value] // keyed by fileType/source
}

var _ Repository = (&MockRepository{})

func (r *MockRepository) Upsert(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error {
	if r.Err != nil {
		return r.Err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.entities == nil {
		r.entities = make(map[string][]search.Entity[search.Value])
	}
	cloned := slices.Clone(entities)
	for i := range cloned {
		cloned[i].Source = search.SourceList(fileType)
	}
	r.entities[fileType] = cloned

	return nil
}

func (r *MockRepository) Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error) {
	if r.Err != nil {
		return nil, r.Err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.entities == nil {
		return nil, nil
	}

	entities := r.entities[string(source)]
	for i := range entities {
		if entities[i].SourceID == sourceID {
			return &entities[i], nil
		}
	}

	return nil, nil
}

func (r *MockRepository) ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error) {
	if r.Err != nil {
		return nil, r.Err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.entities == nil {
		return nil, nil
	}

	entities := r.entities[string(source)]

	// Handle pagination
	startIdx := 0
	if lastSourceID != "" {
		for i, e := range entities {
			if e.SourceID == lastSourceID {
				startIdx = i + 1
				break
			}
		}
	}

	if startIdx >= len(entities) {
		return nil, nil
	}

	endIdx := startIdx + limit
	if endIdx > len(entities) {
		endIdx = len(entities)
	}

	return slices.Clone(entities[startIdx:endIdx]), nil
}

func (r *MockRepository) ListAll(ctx context.Context) ([]search.Entity[search.Value], error) {
	if r.Err != nil {
		return nil, r.Err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	var all []search.Entity[search.Value]
	for _, entities := range r.entities {
		all = append(all, entities...)
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].Source != all[j].Source {
			return all[i].Source < all[j].Source
		}
		return all[i].SourceID < all[j].SourceID
	})
	return slices.Clone(all), nil
}

func (r *MockRepository) Checksums(ctx context.Context) ([]SourceChecksum, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	sources := make([]string, 0, len(r.entities))
	for src := range r.entities {
		sources = append(sources, src)
	}
	sort.Strings(sources)

	out := make([]SourceChecksum, 0, len(sources))
	for _, src := range sources {
		ents := r.entities[src]
		sum, err := checksumEntities(ents)
		if err != nil {
			return nil, err
		}
		out = append(out, SourceChecksum{
			Source:      src,
			EntityCount: len(ents),
			Checksum:    sum,
		})
	}
	return out, nil
}
