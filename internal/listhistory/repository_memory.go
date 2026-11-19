// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"slices"
	"strings"
	"sync"

	"github.com/moov-io/watchman/pkg/search"
)

type memVersion struct {
	listName   string
	sourceID   string
	entityType string
	name       string
	hash       string
	entity     search.Entity[search.Value]
	fromRefID  int64
	toRefID    *int64
}

type memoryRepository struct {
	mu        sync.Mutex
	nextRefID int64
	refs      []Ref
	versions  []memVersion
}

func newMemoryRepository() *memoryRepository {
	return &memoryRepository{}
}

var _ Repository = (*memoryRepository)(nil)

func (r *memoryRepository) GetRef(_ context.Context, listName, refToken string) (*Ref, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.getRefLocked(listName, refToken)
}

func (r *memoryRepository) getRefLocked(listName, refToken string) (*Ref, error) {
	token := strings.TrimSpace(refToken)
	if token == "" || strings.EqualFold(token, RefLatest) {
		var latest *Ref
		for i := range r.refs {
			if r.refs[i].ListName != listName {
				continue
			}
			if latest == nil || r.refs[i].ID > latest.ID {
				cp := r.refs[i]
				latest = &cp
			}
		}
		if latest == nil {
			return nil, ErrNotFound
		}
		return latest, nil
	}
	for i := range r.refs {
		if r.refs[i].ListName == listName && r.refs[i].Ref == token {
			cp := r.refs[i]
			return &cp, nil
		}
	}
	return nil, ErrNotFound
}

func (r *memoryRepository) PreviousRef(_ context.Context, listName string, refID int64) (*Ref, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var prev *Ref
	for i := range r.refs {
		if r.refs[i].ListName != listName || r.refs[i].ID >= refID {
			continue
		}
		if prev == nil || r.refs[i].ID > prev.ID {
			cp := r.refs[i]
			prev = &cp
		}
	}
	if prev == nil {
		return nil, ErrNotFound
	}
	return prev, nil
}

func (r *memoryRepository) ListRefs(_ context.Context, listName string) ([]Ref, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]Ref, 0)
	for i := range r.refs {
		if r.refs[i].ListName == listName {
			out = append(out, r.refs[i])
		}
	}
	slices.SortFunc(out, func(a, b Ref) int {
		if a.ID == b.ID {
			return 0
		}
		if a.ID > b.ID {
			return -1
		}
		return 1
	})
	return out, nil
}

func (r *memoryRepository) ListLatestRefs(_ context.Context) ([]Ref, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	latest := make(map[string]Ref)
	for i := range r.refs {
		cur, ok := latest[r.refs[i].ListName]
		if !ok || r.refs[i].ID > cur.ID {
			latest[r.refs[i].ListName] = r.refs[i]
		}
	}
	names := make([]string, 0, len(latest))
	for name := range latest {
		names = append(names, name)
	}
	slices.Sort(names)
	out := make([]Ref, 0, len(names))
	for _, name := range names {
		out = append(out, latest[name])
	}
	return out, nil
}

func (r *memoryRepository) LiveHashes(_ context.Context, listName string) (map[string]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make(map[string]string)
	for i := range r.versions {
		v := r.versions[i]
		if v.listName == listName && v.toRefID == nil {
			out[v.sourceID] = v.hash
		}
	}
	return out, nil
}

func (r *memoryRepository) ApplyDelta(_ context.Context, params ApplyParams) (*Ref, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if existing, err := r.getRefLocked(params.ListName, params.ListHash); err == nil {
		return existing, nil
	}

	r.nextRefID++
	ref := Ref{
		ID:           r.nextRefID,
		ListName:     params.ListName,
		Ref:          params.ListHash,
		DownloadedAt: params.DownloadedAt.UTC(),
		EntityCount:  params.EntityCount,
		AddedCount:   len(params.Delta.Added),
		RemovedCount: len(params.Delta.Removed),
		ChangedCount: len(params.Delta.Changed),
	}
	r.refs = append(r.refs, ref)

	closeIDs := make(map[string]struct{}, len(params.Delta.Removed)+len(params.Delta.Changed))
	for _, id := range params.Delta.Removed {
		closeIDs[id] = struct{}{}
	}
	for i := range params.Delta.Changed {
		closeIDs[params.Delta.Changed[i].SourceID] = struct{}{}
	}
	for i := range r.versions {
		v := &r.versions[i]
		if v.listName != params.ListName || v.toRefID != nil {
			continue
		}
		if _, ok := closeIDs[v.sourceID]; ok {
			id := ref.ID
			v.toRefID = &id
		}
	}

	add := append(slices.Clone(params.Delta.Added), params.Delta.Changed...)
	for i := range add {
		hash, err := hashEntity(add[i])
		if err != nil {
			return nil, err
		}
		r.versions = append(r.versions, memVersion{
			listName:   params.ListName,
			sourceID:   add[i].SourceID,
			entityType: string(add[i].Type),
			name:       truncateName(add[i].Name),
			hash:       hash,
			entity:     add[i].Normalize(),
			fromRefID:  ref.ID,
		})
	}
	cp := ref
	return &cp, nil
}

func (r *memoryRepository) visibleAt(v memVersion, listName string, refID int64) bool {
	if v.listName != listName {
		return false
	}
	if v.fromRefID > refID {
		return false
	}
	if v.toRefID != nil && *v.toRefID <= refID {
		return false
	}
	return true
}

func (r *memoryRepository) ListEntities(_ context.Context, listName string, refID int64, filter EntityFilter) ([]EntityRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	limit := filter.Limit
	if limit <= 0 {
		limit = 50
	}
	want := limit + 1
	out := make([]EntityRecord, 0)
	for i := range r.versions {
		v := r.versions[i]
		if !r.visibleAt(v, listName, refID) {
			continue
		}
		if filter.EntityType != "" && v.entityType != filter.EntityType {
			continue
		}
		if filter.After != "" && v.sourceID <= filter.After {
			continue
		}
		if !matchesQuery(v.name, v.sourceID, filter.Query) {
			continue
		}
		out = append(out, EntityRecord{
			EntitySummary: EntitySummary{
				SourceID:   v.sourceID,
				Name:       v.name,
				EntityType: search.EntityType(v.entityType),
				EntityHash: v.hash,
			},
			Entity: v.entity,
		})
	}
	slices.SortFunc(out, func(a, b EntityRecord) int {
		return strings.Compare(a.SourceID, b.SourceID)
	})
	if len(out) > want {
		out = out[:want]
	}
	return out, nil
}

func (r *memoryRepository) GetEntity(_ context.Context, listName string, refID int64, sourceID string) (*EntityRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var best *memVersion
	for i := range r.versions {
		v := r.versions[i]
		if !r.visibleAt(v, listName, refID) || v.sourceID != sourceID {
			continue
		}
		if best == nil || v.fromRefID > best.fromRefID {
			cp := v
			best = &cp
		}
	}
	if best == nil {
		return nil, ErrNotFound
	}
	return &EntityRecord{
		EntitySummary: EntitySummary{
			SourceID:   best.sourceID,
			Name:       best.name,
			EntityType: search.EntityType(best.entityType),
			EntityHash: best.hash,
		},
		Entity: best.entity,
	}, nil
}

func (r *memoryRepository) ListSummaries(_ context.Context, listName string, refID int64) ([]EntitySummary, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]EntitySummary, 0)
	for i := range r.versions {
		v := r.versions[i]
		if !r.visibleAt(v, listName, refID) {
			continue
		}
		out = append(out, EntitySummary{
			SourceID:   v.sourceID,
			Name:       v.name,
			EntityType: search.EntityType(v.entityType),
			EntityHash: v.hash,
		})
	}
	slices.SortFunc(out, func(a, b EntitySummary) int {
		return strings.Compare(a.SourceID, b.SourceID)
	})
	return out, nil
}

func (r *memoryRepository) EntitiesBySourceIDs(_ context.Context, listName string, refID int64, ids []string) (map[string]search.Entity[search.Value], error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	want := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		want[id] = struct{}{}
	}
	out := make(map[string]search.Entity[search.Value], len(ids))
	fromRef := make(map[string]int64, len(ids))
	for i := range r.versions {
		v := r.versions[i]
		if !r.visibleAt(v, listName, refID) {
			continue
		}
		if _, ok := want[v.sourceID]; !ok {
			continue
		}
		if prev, ok := fromRef[v.sourceID]; ok && prev >= v.fromRefID {
			continue
		}
		fromRef[v.sourceID] = v.fromRefID
		out[v.sourceID] = v.entity
	}
	return out, nil
}
