package ingest

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
)

const listPageSize = 1000

type Repository interface {
	Upsert(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error
	Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error)
	ListBySource(ctx context.Context, lastSourceID string, source search.SourceList, limit int) ([]search.Entity[search.Value], error)
	ListAll(ctx context.Context) ([]search.Entity[search.Value], error)
	Checksums(ctx context.Context) ([]SourceChecksum, error)
}

func NewRepository(database db.DB) Repository {
	if database == nil {
		return &MockRepository{}
	}
	return &sqlRepository{db: database}
}

type sqlRepository struct {
	db db.DB
}

func (r *sqlRepository) Upsert(ctx context.Context, fileType string, entities []search.Entity[search.Value]) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin ingest upsert: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	if err := r.deleteEntities(ctx, tx, fileType); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM ingest_sources WHERE source = ?;", fileType); err != nil {
		return fmt.Errorf("deleting %s ingest checksum: %w", fileType, err)
	}

	if len(entities) == 0 {
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit empty ingest upsert: %w", err)
		}
		return nil
	}

	indexed := append([]search.Entity[search.Value](nil), entities...)
	for i := range indexed {
		indexed[i].Source = search.SourceList(fileType)
	}
	sort.Slice(indexed, func(i, j int) bool {
		return indexed[i].SourceID < indexed[j].SourceID
	})

	h := sha256.New()
	for i := range indexed {
		raw, err := json.Marshal(indexed[i])
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}
		writeChecksum(h, indexed[i].SourceID, raw)
		if err := r.insertEntity(ctx, tx, indexed[i], raw); err != nil {
			return fmt.Errorf("upserting (%s) %s/%s entity: %w", fileType, indexed[i].Source, indexed[i].SourceID, err)
		}
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO ingest_sources (source, entity_count, checksum, updated_at) VALUES (?, ?, ?, ?);`,
		fileType,
		len(indexed),
		hex.EncodeToString(h.Sum(nil)),
		time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("inserting %s ingest checksum: %w", fileType, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit ingest upsert: %w", err)
	}
	return nil
}

func (r *sqlRepository) deleteEntities(ctx context.Context, tx db.Tx, fileType string) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM ingested_entities WHERE source = ?;", fileType)
	if err != nil {
		return fmt.Errorf("deleting %s entities: %w", fileType, err)
	}
	return nil
}

func (r *sqlRepository) insertEntity(ctx context.Context, tx db.Tx, entity search.Entity[search.Value], raw []byte) error {
	_, err := tx.ExecContext(ctx,
		`INSERT INTO ingested_entities (type, source, source_id, entity) VALUES (?, ?, ?, ?);`,
		string(entity.Type),
		string(entity.Source),
		string(entity.SourceID),
		raw,
	)
	if err != nil {
		return fmt.Errorf("inserting ingested entity: %w", err)
	}
	return nil
}

func (r *sqlRepository) Get(ctx context.Context, sourceID string, source search.SourceList) (*search.Entity[search.Value], error) {
	qry := `SELECT entity FROM ingested_entities WHERE source_id = ? AND source = ? LIMIT 1;`

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
	if limit <= 0 {
		limit = listPageSize
	}
	qry := `SELECT entity FROM ingested_entities WHERE source = ? AND source_id > ? ORDER BY source_id ASC LIMIT ?;`

	rows, err := r.queryScanEntities(ctx, qry, string(source), lastSourceID, limit)
	if err != nil {
		return nil, fmt.Errorf("listing ingested entities by source: %w", err)
	}
	return rows, nil
}

func (r *sqlRepository) ListAll(ctx context.Context) ([]search.Entity[search.Value], error) {
	var all []search.Entity[search.Value]
	lastSource, lastID := "", ""
	for {
		batch, err := r.listAfter(ctx, lastSource, lastID, listPageSize)
		if err != nil {
			return nil, err
		}
		if len(batch) == 0 {
			break
		}
		all = append(all, batch...)
		last := batch[len(batch)-1]
		lastSource, lastID = string(last.Source), last.SourceID
		if len(batch) < listPageSize {
			break
		}
	}
	return all, nil
}

func (r *sqlRepository) listAfter(ctx context.Context, lastSource, lastSourceID string, limit int) ([]search.Entity[search.Value], error) {
	qry := `SELECT entity FROM ingested_entities
WHERE source > ? OR (source = ? AND source_id > ?)
ORDER BY source ASC, source_id ASC
LIMIT ?;`

	rows, err := r.queryScanEntities(ctx, qry, lastSource, lastSource, lastSourceID, limit)
	if err != nil {
		return nil, fmt.Errorf("listing all ingested entities: %w", err)
	}
	return rows, nil
}

func (r *sqlRepository) Checksums(ctx context.Context) ([]SourceChecksum, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT source, entity_count, checksum FROM ingest_sources ORDER BY source ASC;`)
	if err != nil {
		return nil, fmt.Errorf("listing ingest checksums: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var out []SourceChecksum
	for rows.Next() {
		var row SourceChecksum
		if err := rows.Scan(&row.Source, &row.EntityCount, &row.Checksum); err != nil {
			return nil, fmt.Errorf("scanning ingest checksum: %w", err)
		}
		out = append(out, row)
	}
	return out, rows.Err()
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
