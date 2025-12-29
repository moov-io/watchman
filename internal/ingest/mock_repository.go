package ingest

import (
	"context"
	"slices"
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
	r.entities[fileType] = slices.Clone(entities)

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
