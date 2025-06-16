package ingest

import (
	"context"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
)

type Repository interface {
	Add(ctx context.Context, entity search.Entity[search.Value]) error
	Get(ctx context.Context, sourceID string, source search.SourceList) (search.Entity[search.Value], error)
	ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error)
}

func NewRepsoitory(db db.DB) Repository {
	return &sqlRepository{db: db}
}

type sqlRepository struct {
	db db.DB
}

func (r *sqlRepository) Add(ctx context.Context, entity search.Entity[search.Value]) error {
	qry := `INSERT INTO ingested_entities (type, source, source_id, entity) VALUES (?, ?, ?, ?);`

}

func (r *sqlRepository) Get(ctx context.Context, sourceID string, source search.SourceList) (search.Entity[search.Value], error) {

}

func (r *sqlRepository) ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error) {

}
