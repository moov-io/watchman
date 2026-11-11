// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/moov-io/base/database"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
)

type sqlRepository struct {
	db db.DB
}

var _ Repository = (*sqlRepository)(nil)

const refColumns = `id, list_name, list_hash, downloaded_at, entity_count, added_count, removed_count, changed_count`

func (r *sqlRepository) GetRef(ctx context.Context, listName, refToken string) (*Ref, error) {
	token := strings.TrimSpace(refToken)
	var row *sql.Row
	if token == "" || strings.EqualFold(token, RefLatest) {
		row = r.db.QueryRowContext(ctx,
			`SELECT `+refColumns+` FROM list_refs WHERE list_name = ? ORDER BY id DESC LIMIT 1;`,
			listName)
	} else {
		row = r.db.QueryRowContext(ctx,
			`SELECT `+refColumns+` FROM list_refs WHERE list_name = ? AND list_hash = ? LIMIT 1;`,
			listName, token)
	}
	ref, err := scanRef(row)
	if err != nil {
		return nil, fmt.Errorf("getting list ref: %w", err)
	}
	return ref, nil
}

func (r *sqlRepository) PreviousRef(ctx context.Context, listName string, refID int64) (*Ref, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT `+refColumns+` FROM list_refs WHERE list_name = ? AND id < ? ORDER BY id DESC LIMIT 1;`,
		listName, refID)
	ref, err := scanRef(row)
	if err != nil {
		return nil, fmt.Errorf("getting previous list ref: %w", err)
	}
	return ref, nil
}

func (r *sqlRepository) ListRefs(ctx context.Context, listName string) ([]Ref, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT `+refColumns+` FROM list_refs WHERE list_name = ? ORDER BY id DESC;`,
		listName)
	if err != nil {
		return nil, fmt.Errorf("listing list refs: %w", err)
	}
	defer rows.Close()
	return scanRefs(rows)
}

func (r *sqlRepository) ListLatestRefs(ctx context.Context) ([]Ref, error) {
	rows, err := r.db.QueryContext(ctx, `
SELECT r.id, r.list_name, r.list_hash, r.downloaded_at, r.entity_count, r.added_count, r.removed_count, r.changed_count
FROM list_refs r
INNER JOIN (
  SELECT list_name, MAX(id) AS max_id FROM list_refs GROUP BY list_name
) t ON r.id = t.max_id
ORDER BY r.list_name ASC;`)
	if err != nil {
		return nil, fmt.Errorf("listing latest list refs: %w", err)
	}
	defer rows.Close()
	return scanRefs(rows)
}

func (r *sqlRepository) LiveHashes(ctx context.Context, listName string) (map[string]string, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT source_id, entity_hash FROM list_entity_versions WHERE list_name = ? AND to_ref_id IS NULL;`,
		listName)
	if err != nil {
		return nil, fmt.Errorf("listing live entity hashes: %w", err)
	}
	defer rows.Close()

	out := make(map[string]string)
	for rows.Next() {
		var sourceID, hash string
		if err := rows.Scan(&sourceID, &hash); err != nil {
			return nil, fmt.Errorf("scanning live entity hash: %w", err)
		}
		out[sourceID] = hash
	}
	return out, rows.Err()
}

func (r *sqlRepository) ApplyDelta(ctx context.Context, params ApplyParams) (*Ref, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin list history delta: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	_, err = tx.ExecContext(ctx, `
INSERT INTO list_refs (list_name, list_hash, downloaded_at, entity_count, added_count, removed_count, changed_count)
VALUES (?, ?, ?, ?, ?, ?, ?);`,
		params.ListName,
		params.ListHash,
		params.DownloadedAt.UTC(),
		params.EntityCount,
		len(params.Delta.Added),
		len(params.Delta.Removed),
		len(params.Delta.Changed),
	)
	if err != nil {
		if database.UniqueViolation(err) {
			// Postgres aborts the transaction after 23505. Roll back before
			// reading the row the other writer inserted.
			_ = tx.Rollback()
			return r.GetRef(ctx, params.ListName, params.ListHash)
		}
		return nil, fmt.Errorf("inserting list ref: %w", err)
	}

	var refID int64
	err = tx.QueryRowContext(ctx,
		`SELECT id FROM list_refs WHERE list_name = ? AND list_hash = ?;`,
		params.ListName, params.ListHash).Scan(&refID)
	if err != nil {
		return nil, fmt.Errorf("reading inserted list ref id: %w", err)
	}

	closeIDs := make([]string, 0, len(params.Delta.Removed)+len(params.Delta.Changed))
	closeIDs = append(closeIDs, params.Delta.Removed...)
	for i := range params.Delta.Changed {
		closeIDs = append(closeIDs, params.Delta.Changed[i].SourceID)
	}
	if err := r.closeVersions(ctx, tx, params.ListName, refID, closeIDs); err != nil {
		return nil, err
	}

	incoming := append(slices.Clone(params.Delta.Added), params.Delta.Changed...)
	if err := r.insertVersions(ctx, tx, params.ListName, refID, incoming); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit list history delta: %w", err)
	}

	return &Ref{
		ID:           refID,
		ListName:     params.ListName,
		Ref:          params.ListHash,
		DownloadedAt: params.DownloadedAt.UTC(),
		EntityCount:  params.EntityCount,
		AddedCount:   len(params.Delta.Added),
		RemovedCount: len(params.Delta.Removed),
		ChangedCount: len(params.Delta.Changed),
	}, nil
}

func (r *sqlRepository) closeVersions(ctx context.Context, tx db.Tx, listName string, refID int64, sourceIDs []string) error {
	for i := 0; i < len(sourceIDs); i += closeChunk {
		end := i + closeChunk
		if end > len(sourceIDs) {
			end = len(sourceIDs)
		}
		chunk := sourceIDs[i:end]
		placeholders := strings.Repeat("?,", len(chunk))
		placeholders = strings.TrimSuffix(placeholders, ",")
		args := make([]any, 0, 2+len(chunk))
		args = append(args, refID, listName)
		for _, id := range chunk {
			args = append(args, id)
		}
		qry := `UPDATE list_entity_versions SET to_ref_id = ? WHERE list_name = ? AND to_ref_id IS NULL AND source_id IN (` + placeholders + `);`
		if _, err := tx.ExecContext(ctx, qry, args...); err != nil {
			return fmt.Errorf("closing list entity versions: %w", err)
		}
	}
	return nil
}

func (r *sqlRepository) insertVersions(ctx context.Context, tx db.Tx, listName string, refID int64, entities []search.Entity[search.Value]) error {
	for i := 0; i < len(entities); i += insertChunk {
		end := i + insertChunk
		if end > len(entities) {
			end = len(entities)
		}
		chunk := entities[i:end]
		var b strings.Builder
		b.WriteString(`INSERT INTO list_entity_versions (list_name, source_id, entity_type, name, entity_hash, entity, from_ref_id) VALUES `)
		args := make([]any, 0, len(chunk)*7)
		for j := range chunk {
			raw, hash, err := marshalEntity(chunk[j])
			if err != nil {
				return err
			}
			if j > 0 {
				b.WriteByte(',')
			}
			b.WriteString("(?,?,?,?,?,?,?)")
			args = append(args,
				listName,
				chunk[j].SourceID,
				string(chunk[j].Type),
				truncateName(chunk[j].Name),
				hash,
				raw,
				refID,
			)
		}
		if _, err := tx.ExecContext(ctx, b.String(), args...); err != nil {
			return fmt.Errorf("inserting list entity versions: %w", err)
		}
	}
	return nil
}

func (r *sqlRepository) ListEntities(ctx context.Context, listName string, refID int64, filter EntityFilter) ([]EntityRecord, error) {
	limit := filter.Limit
	if limit <= 0 {
		limit = 50
	}
	qry, args := entityWindowQuery(listName, refID, filter, true, limit+1)
	rows, err := r.db.QueryContext(ctx, qry, args...)
	if err != nil {
		return nil, fmt.Errorf("listing list entities: %w", err)
	}
	defer rows.Close()

	var out []EntityRecord
	for rows.Next() {
		var rec EntityRecord
		var raw, entityType string
		if err := rows.Scan(&rec.SourceID, &entityType, &rec.Name, &rec.EntityHash, &raw); err != nil {
			return nil, fmt.Errorf("scanning list entity: %w", err)
		}
		entity, err := decodeEntity(raw)
		if err != nil {
			return nil, err
		}
		rec.EntityType = search.EntityType(entityType)
		rec.Entity = entity
		out = append(out, rec)
	}
	if out == nil {
		out = []EntityRecord{}
	}
	return out, rows.Err()
}

func (r *sqlRepository) GetEntity(ctx context.Context, listName string, refID int64, sourceID string) (*EntityRecord, error) {
	row := r.db.QueryRowContext(ctx, `
SELECT source_id, entity_type, name, entity_hash, entity
FROM list_entity_versions
WHERE list_name = ? AND from_ref_id <= ? AND (to_ref_id IS NULL OR to_ref_id > ?) AND source_id = ?
ORDER BY from_ref_id DESC
LIMIT 1;`, listName, refID, refID, sourceID)
	var rec EntityRecord
	var raw, entityType string
	err := row.Scan(&rec.SourceID, &entityType, &rec.Name, &rec.EntityHash, &raw)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("getting list entity: %w", err)
	}
	entity, err := decodeEntity(raw)
	if err != nil {
		return nil, err
	}
	rec.EntityType = search.EntityType(entityType)
	rec.Entity = entity
	return &rec, nil
}

func (r *sqlRepository) ListSummaries(ctx context.Context, listName string, refID int64) ([]EntitySummary, error) {
	qry, args := entityWindowQuery(listName, refID, EntityFilter{}, false, 0)
	rows, err := r.db.QueryContext(ctx, qry, args...)
	if err != nil {
		return nil, fmt.Errorf("listing list entity summaries: %w", err)
	}
	defer rows.Close()

	out := make([]EntitySummary, 0)
	for rows.Next() {
		var rec EntitySummary
		var entityType string
		if err := rows.Scan(&rec.SourceID, &entityType, &rec.Name, &rec.EntityHash); err != nil {
			return nil, fmt.Errorf("scanning list entity summary: %w", err)
		}
		rec.EntityType = search.EntityType(entityType)
		out = append(out, rec)
	}
	return out, rows.Err()
}

func (r *sqlRepository) EntitiesBySourceIDs(ctx context.Context, listName string, refID int64, ids []string) (map[string]search.Entity[search.Value], error) {
	out := make(map[string]search.Entity[search.Value], len(ids))
	if len(ids) == 0 {
		return out, nil
	}
	for i := 0; i < len(ids); i += closeChunk {
		end := i + closeChunk
		if end > len(ids) {
			end = len(ids)
		}
		chunk := ids[i:end]
		placeholders := strings.Repeat("?,", len(chunk))
		placeholders = strings.TrimSuffix(placeholders, ",")
		args := make([]any, 0, 3+len(chunk))
		args = append(args, listName, refID, refID)
		for _, id := range chunk {
			args = append(args, id)
		}
		qry := `SELECT source_id, entity FROM list_entity_versions
WHERE list_name = ? AND from_ref_id <= ? AND (to_ref_id IS NULL OR to_ref_id > ?)
AND source_id IN (` + placeholders + `)
ORDER BY from_ref_id DESC;`
		rows, err := r.db.QueryContext(ctx, qry, args...)
		if err != nil {
			return nil, fmt.Errorf("getting list entities by source id: %w", err)
		}
		if err := scanEntitiesBySourceIDs(rows, out); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func scanEntitiesBySourceIDs(rows *sql.Rows, out map[string]search.Entity[search.Value]) error {
	defer rows.Close()
	for rows.Next() {
		var sourceID, raw string
		if err := rows.Scan(&sourceID, &raw); err != nil {
			return fmt.Errorf("scanning list entity json: %w", err)
		}
		if _, exists := out[sourceID]; exists {
			continue
		}
		entity, err := decodeEntity(raw)
		if err != nil {
			return err
		}
		out[sourceID] = entity
	}
	return rows.Err()
}

func entityWindowQuery(listName string, refID int64, filter EntityFilter, includeEntity bool, limit int) (string, []any) {
	cols := `source_id, entity_type, name, entity_hash`
	if includeEntity {
		cols += `, entity`
	}
	var b strings.Builder
	b.WriteString(`SELECT `)
	b.WriteString(cols)
	b.WriteString(` FROM list_entity_versions WHERE list_name = ? AND from_ref_id <= ? AND (to_ref_id IS NULL OR to_ref_id > ?)`)
	args := []any{listName, refID, refID}

	if filter.EntityType != "" {
		b.WriteString(` AND entity_type = ?`)
		args = append(args, filter.EntityType)
	}
	if filter.After != "" {
		b.WriteString(` AND source_id > ?`)
		args = append(args, filter.After)
	}
	if q := strings.TrimSpace(filter.Query); q != "" {
		b.WriteString(` AND (LOWER(name) LIKE ? ESCAPE '!' OR LOWER(source_id) LIKE ? ESCAPE '!')`)
		like := likeContains(q)
		args = append(args, like, like)
	}
	b.WriteString(` ORDER BY source_id ASC`)
	if limit > 0 {
		b.WriteString(` LIMIT ?`)
		args = append(args, limit)
	}
	b.WriteByte(';')
	return b.String(), args
}

func scanRef(row *sql.Row) (*Ref, error) {
	var ref Ref
	var downloadedAt time.Time
	err := row.Scan(&ref.ID, &ref.ListName, &ref.Ref, &downloadedAt, &ref.EntityCount, &ref.AddedCount, &ref.RemovedCount, &ref.ChangedCount)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	ref.DownloadedAt = downloadedAt.UTC()
	return &ref, nil
}

func scanRefs(rows *sql.Rows) ([]Ref, error) {
	out := make([]Ref, 0)
	for rows.Next() {
		var ref Ref
		var downloadedAt time.Time
		if err := rows.Scan(&ref.ID, &ref.ListName, &ref.Ref, &downloadedAt, &ref.EntityCount, &ref.AddedCount, &ref.RemovedCount, &ref.ChangedCount); err != nil {
			return nil, fmt.Errorf("scanning list ref: %w", err)
		}
		ref.DownloadedAt = downloadedAt.UTC()
		out = append(out, ref)
	}
	return out, rows.Err()
}
