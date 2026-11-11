// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"errors"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
)

var (
	ErrNotConfigured = errors.New("list history requires a configured database")
	ErrNotFound      = errors.New("not found")
)

type Repository interface {
	GetRef(ctx context.Context, listName, refToken string) (*Ref, error)
	PreviousRef(ctx context.Context, listName string, refID int64) (*Ref, error)
	ListRefs(ctx context.Context, listName string) ([]Ref, error)
	ListLatestRefs(ctx context.Context) ([]Ref, error)

	LiveHashes(ctx context.Context, listName string) (map[string]string, error)
	ApplyDelta(ctx context.Context, params ApplyParams) (*Ref, error)

	ListEntities(ctx context.Context, listName string, refID int64, filter EntityFilter) ([]EntityRecord, error)
	GetEntity(ctx context.Context, listName string, refID int64, sourceID string) (*EntityRecord, error)
	ListSummaries(ctx context.Context, listName string, refID int64) ([]EntitySummary, error)
	EntitiesBySourceIDs(ctx context.Context, listName string, refID int64, ids []string) (map[string]search.Entity[search.Value], error)
}

func NewRepository(database db.DB) Repository {
	if database == nil {
		return nopRepository{}
	}
	return &sqlRepository{db: database}
}

func NewMemoryRepository() Repository {
	return newMemoryRepository()
}
