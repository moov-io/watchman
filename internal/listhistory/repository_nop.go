// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"

	"github.com/moov-io/watchman/pkg/search"
)

type nopRepository struct{}

var _ Repository = nopRepository{}

func (nopRepository) GetRef(context.Context, string, string) (*Ref, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) PreviousRef(context.Context, string, int64) (*Ref, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) ListRefs(context.Context, string) ([]Ref, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) ListLatestRefs(context.Context) ([]Ref, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) LiveHashes(context.Context, string) (map[string]string, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) ApplyDelta(context.Context, ApplyParams) (*Ref, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) ListEntities(context.Context, string, int64, EntityFilter) ([]EntityRecord, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) GetEntity(context.Context, string, int64, string) (*EntityRecord, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) ListSummaries(context.Context, string, int64) ([]EntitySummary, error) {
	return nil, ErrNotConfigured
}

func (nopRepository) EntitiesBySourceIDs(context.Context, string, int64, []string) (map[string]search.Entity[search.Value], error) {
	return nil, ErrNotConfigured
}
