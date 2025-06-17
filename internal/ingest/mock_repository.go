package ingest

import (
	"context"

	"github.com/moov-io/watchman/pkg/search"
)

type MockRepository struct {
	Err error
}

func (r *MockRepository) Upsert(ctx context.Context, entity search.Entity[search.Value]) error {
	return r.Err
}

func (r *MockRepository) Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error) {
	return nil, r.Err
}

func (r *MockRepository) ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error) {
	return nil, r.Err
}
