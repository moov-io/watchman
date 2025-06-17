package ingest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/moov-io/base/database"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
)

type Repository interface {
	Upsert(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error
	Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error)
	ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error)
}

func NewRepository(db db.DB) Repository {
	if db == nil {
		return &MockRepository{}
	}
	return &sqlRepository{db: db}
}

type sqlRepository struct {
	db db.DB
}

func (r *sqlRepository) Upsert(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error {
	// Delete the existing rows
	err := r.deleteEntities(ctx, fileType)
	if err != nil {
		return err
	}

	for idx := range entities {
		err = r.upsertEntity(ctx, entities[idx])
		if err != nil {
			return fmt.Errorf("upserting (%s) %s/%s entity: %w", fileType, entities[idx].Source, entities[idx].SourceID, err)
		}
	}

	return nil
}

func (r *sqlRepository) deleteEntities(ctx context.Context, fileType string) error {
	qry := "DELETE FROM ingested_entities WHERE source = ?;"

	_, err := r.db.ExecContext(ctx, qry, fileType)
	if err != nil {
		return fmt.Errorf("deleting %s entities: %w", fileType, err)
	}
	return nil
}

func (r *sqlRepository) upsertEntity(ctx context.Context, entity search.Entity[search.Value]) error {
	qry := `INSERT INTO ingested_entities (type, source, source_id, entity) VALUES (?, ?, ?, ?);`

	bs, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	_, err = r.db.ExecContext(ctx, qry,
		string(entity.Type),
		string(entity.Source),
		string(entity.SourceID),
		bs,
	)
	if err != nil {
		// Update if we collide on INSERT
		if database.UniqueViolation(err) {
			qry := `UPDATE ingested_entities SET entity = ? WHERE source = ? AND source_id = ?;`

			_, err := r.db.ExecContext(ctx, qry,
				// SET
				bs,
				// WHERE
				string(entity.Source),
				string(entity.SourceID),
			)
			if err != nil {
				return fmt.Errorf("updating ingested entity: %w", err)
			}
			return nil
		}
		return fmt.Errorf("inserting ingested entity: %w", err)
	}

	return nil
}

func (r *sqlRepository) Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error) {
	qry := `SELECT entity from ingested_entities where source_id = ? AND source = ? LIMIT 1;`

	rows, err := r.queryScanEntities(ctx, qry, sourceID, string(source))
	if err != nil {
		return nil, fmt.Errorf("getting ingested entity: %w", err)
	}
	if len(rows) > 0 {
		return &rows[0], nil
	}
	return nil, errors.New("no entity found")
}

func (r *sqlRepository) ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error) {
	qry := `SELECT entity from ingested_entities where source_id > ? AND source = ? LIMIT ?;`

	rows, err := r.queryScanEntities(ctx, qry, lastSourceID, string(source), limit)
	if err != nil {
		return nil, fmt.Errorf("listing ingested entities by source: %w", err)
	}
	return rows, nil
}

func (r *sqlRepository) queryScanEntities(ctx context.Context, qry string, args ...interface{}) ([]search.Entity[search.Value], error) {
	rows, err := r.db.QueryContext(ctx, qry, args...)
	if err != nil {
		return nil, fmt.Errorf("query for ingested entities: %w", err)
	}
	defer rows.Close()

	var out []search.Entity[search.Value]
	for rows.Next() {
		var data string
		err = rows.Scan(&data)
		if err != nil {
			return nil, fmt.Errorf("scanning entity json: %w", err)
		}

		var row search.Entity[search.Value]
		err = json.NewDecoder(strings.NewReader(data)).Decode(&row)
		if err != nil {
			return nil, fmt.Errorf("json decode: %w", err)
		}

		out = append(out, row.Normalize())
	}
	return out, rows.Err()
}
