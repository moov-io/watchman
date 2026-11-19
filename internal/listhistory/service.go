// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	defaultListLimit = 50
	maxListLimit     = 100
	defaultDiffLimit = 100
	maxDiffLimit     = 500
)

type snapshot struct {
	ctx    context.Context
	byList map[string][]search.Entity[search.Value]
	hashes map[string]string
	at     time.Time
}

type Service struct {
	logger  log.Logger
	repo    Repository
	enabled bool

	mu     sync.Mutex
	listMu map[string]*sync.Mutex
	jobs   chan snapshot
	wg     sync.WaitGroup
}

var _ download.ListHistory = (*Service)(nil)

func NewService(logger log.Logger, repo Repository) *Service {
	if repo == nil {
		repo = nopRepository{}
	}
	_, nop := repo.(nopRepository)
	s := &Service{
		logger:  logger,
		repo:    repo,
		enabled: !nop,
		listMu:  make(map[string]*sync.Mutex),
	}
	if s.enabled {
		// Buffer two snapshots so a refresh can finish while the previous
		// store is still writing. A third send blocks the refresher.
		s.jobs = make(chan snapshot, 2)
		go s.worker()
	}
	return s
}

func (s *Service) Enabled() bool {
	return s != nil && s.enabled
}

// StoreRefreshAsync copies list entities from stats and queues the delta for a
// single worker so overlapping refreshes apply in call order. The caller may
// drop stats.Entities after this returns. A full buffer blocks the caller.
func (s *Service) StoreRefreshAsync(ctx context.Context, stats download.Stats) {
	if !s.Enabled() {
		return
	}
	snap := snapshotStats(stats)
	snap.ctx = context.WithoutCancel(ctx)
	s.wg.Add(1)
	s.jobs <- snap
}

func (s *Service) worker() {
	for snap := range s.jobs {
		func(snap snapshot) {
			defer s.wg.Done()
			defer func() {
				if rec := recover(); rec != nil {
					s.logger.Error().Logf("storing list history panicked: %v", rec)
				}
			}()
			ctx := snap.ctx
			if ctx == nil {
				ctx = context.Background()
			}
			if err := s.storeSnapshot(ctx, snap); err != nil {
				s.logger.Error().Logf("storing list history: %v", err)
			}
		}(snap)
	}
}

func (s *Service) Wait(ctx context.Context) error {
	if s == nil {
		return nil
	}
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func snapshotStats(stats download.Stats) snapshot {
	byList := make(map[string][]search.Entity[search.Value])
	for i := range stats.Entities {
		src := string(stats.Entities[i].Source)
		if src == "" {
			continue
		}
		byList[src] = append(byList[src], stats.Entities[i])
	}
	at := stats.EndedAt
	if at.IsZero() {
		at = time.Now().UTC()
	}
	return snapshot{
		byList: byList,
		hashes: maps.Clone(stats.ListHashes),
		at:     at.UTC(),
	}
}

func (s *Service) lockList(listName string) func() {
	s.mu.Lock()
	m, ok := s.listMu[listName]
	if !ok {
		m = &sync.Mutex{}
		s.listMu[listName] = m
	}
	s.mu.Unlock()
	m.Lock()
	return m.Unlock
}

func (s *Service) storeSnapshot(ctx context.Context, snap snapshot) error {
	ctx, span := telemetry.StartSpan(ctx, "store-list-history")
	defer span.End()

	span.SetAttributes(attribute.Int("list_count", len(snap.hashes)))

	var firstErr error
	for listName, hash := range snap.hashes {
		if hash == "" {
			continue
		}
		if err := s.storeList(ctx, listName, hash, snap.byList[listName], snap.at); err != nil {
			if firstErr == nil {
				firstErr = err
			}
			s.logger.Error().With(log.Fields{
				"list_name": log.String(listName),
			}).Logf("storing list history delta failed: %v", err)
		}
	}
	return firstErr
}

func (s *Service) storeList(ctx context.Context, listName, hash string, entities []search.Entity[search.Value], at time.Time) error {
	unlock := s.lockList(listName)
	defer unlock()

	ctx, cancel := context.WithTimeout(ctx, 15*time.Minute)
	defer cancel()

	ctx, span := telemetry.StartSpan(ctx, "store-list-history-list", trace.WithAttributes(
		attribute.String("list_name", listName),
	))
	defer span.End()

	latest, err := s.repo.GetRef(ctx, listName, RefLatest)
	if err != nil && !errors.Is(err, ErrNotFound) {
		span.RecordError(err)
		return fmt.Errorf("loading latest %s ref: %w", listName, err)
	}
	if latest != nil && latest.Ref == hash {
		span.SetAttributes(attribute.Bool("unchanged_hash", true))
		return nil
	}
	if latest != nil {
		existing, err := s.repo.GetRef(ctx, listName, hash)
		if err != nil && !errors.Is(err, ErrNotFound) {
			span.RecordError(err)
			return fmt.Errorf("loading %s ref: %w", listName, err)
		}
		if existing != nil {
			// Same file was stored earlier (cache replay or rollback). latest
			// stays on the last new hash; callers should use listinfo.listHashes.
			span.SetAttributes(attribute.Bool("known_hash", true))
			return nil
		}
	}
	if latest != nil && latest.DownloadedAt.After(at) {
		span.SetAttributes(attribute.Bool("stale_snapshot", true))
		return nil
	}

	oldHashes := map[string]string{}
	if latest != nil {
		oldHashes, err = s.repo.LiveHashes(ctx, listName)
		if err != nil {
			span.RecordError(err)
			return fmt.Errorf("loading live %s hashes: %w", listName, err)
		}
	}

	delta, err := computeDelta(oldHashes, entities)
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("computing %s delta: %w", listName, err)
	}

	entityCount := len(oldHashes) + len(delta.Added) - len(delta.Removed)
	ref, err := s.repo.ApplyDelta(ctx, ApplyParams{
		ListName:     listName,
		ListHash:     hash,
		DownloadedAt: at,
		EntityCount:  entityCount,
		Delta:        delta,
	})
	if err != nil {
		span.RecordError(err)
		return fmt.Errorf("applying %s delta: %w", listName, err)
	}

	span.SetAttributes(
		attribute.Int("added_count", len(delta.Added)),
		attribute.Int("removed_count", len(delta.Removed)),
		attribute.Int("changed_count", len(delta.Changed)),
		attribute.Int("entity_count", entityCount),
	)
	s.logger.Info().With(log.Fields{
		"list_name": log.String(listName),
		"ref":       log.String(ref.Ref),
	}).Log("stored list history")
	return nil
}

func (s *Service) ListLatest(ctx context.Context) ([]Ref, error) {
	if !s.Enabled() {
		return nil, ErrNotConfigured
	}
	refs, err := s.repo.ListLatestRefs(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing latest list refs: %w", err)
	}
	if refs == nil {
		refs = []Ref{}
	}
	return refs, nil
}

func (s *Service) ListRefs(ctx context.Context, listName string) ([]Ref, error) {
	if !s.Enabled() {
		return nil, ErrNotConfigured
	}
	refs, err := s.repo.ListRefs(ctx, listName)
	if err != nil {
		return nil, fmt.Errorf("listing refs: %w", err)
	}
	if refs == nil {
		refs = []Ref{}
	}
	return refs, nil
}

func (s *Service) ListEntities(ctx context.Context, listName, refToken string, filter EntityFilter) (EntitiesResponse, error) {
	var out EntitiesResponse
	if !s.Enabled() {
		return out, ErrNotConfigured
	}
	ref, err := s.repo.GetRef(ctx, listName, refToken)
	if err != nil {
		return out, fmt.Errorf("resolving ref: %w", err)
	}
	if filter.Limit <= 0 {
		filter.Limit = defaultListLimit
	}
	if filter.Limit > maxListLimit {
		filter.Limit = maxListLimit
	}
	rows, err := s.repo.ListEntities(ctx, listName, ref.ID, filter)
	if err != nil {
		return out, fmt.Errorf("listing entities: %w", err)
	}
	after := ""
	if len(rows) > filter.Limit {
		rows = rows[:filter.Limit]
		after = rows[len(rows)-1].SourceID
	}
	if rows == nil {
		rows = []EntityRecord{}
	}
	out = EntitiesResponse{
		ListName: listName,
		Ref:      *ref,
		Entities: rows,
		After:    after,
	}
	return out, nil
}

func (s *Service) GetEntity(ctx context.Context, listName, refToken, sourceID string) (*EntityRecord, error) {
	if !s.Enabled() {
		return nil, ErrNotConfigured
	}
	ref, err := s.repo.GetRef(ctx, listName, refToken)
	if err != nil {
		return nil, fmt.Errorf("resolving ref: %w", err)
	}
	rec, err := s.repo.GetEntity(ctx, listName, ref.ID, sourceID)
	if err != nil {
		return nil, fmt.Errorf("getting entity: %w", err)
	}
	return rec, nil
}

func (s *Service) Diff(ctx context.Context, listName, fromToken, toToken, after, detail string, limit int) (DiffResponse, error) {
	var out DiffResponse
	if !s.Enabled() {
		return out, ErrNotConfigured
	}
	if limit <= 0 {
		limit = defaultDiffLimit
	}
	if limit > maxDiffLimit {
		limit = maxDiffLimit
	}

	toRef, err := s.repo.GetRef(ctx, listName, toToken)
	if err != nil {
		return out, fmt.Errorf("resolving to ref: %w", err)
	}

	var fromRef *Ref
	if strings.TrimSpace(fromToken) == "" {
		fromRef, err = s.repo.PreviousRef(ctx, listName, toRef.ID)
		if errors.Is(err, ErrNotFound) {
			fromRef = toRef
			err = nil
		}
	} else {
		fromRef, err = s.repo.GetRef(ctx, listName, fromToken)
	}
	if err != nil {
		return out, fmt.Errorf("resolving from ref: %w", err)
	}

	if fromRef.ID > toRef.ID {
		fromRef, toRef = toRef, fromRef
	}

	fromRows, err := s.repo.ListSummaries(ctx, listName, fromRef.ID)
	if err != nil {
		return out, fmt.Errorf("listing from snapshot: %w", err)
	}
	toRows, err := s.repo.ListSummaries(ctx, listName, toRef.ID)
	if err != nil {
		return out, fmt.Errorf("listing to snapshot: %w", err)
	}

	fromMap := make(map[string]EntitySummary, len(fromRows))
	for _, row := range fromRows {
		fromMap[row.SourceID] = row
	}
	toMap := make(map[string]EntitySummary, len(toRows))
	for _, row := range toRows {
		toMap[row.SourceID] = row
	}

	var added, removed, changed []Change
	for _, row := range toRows {
		prev, ok := fromMap[row.SourceID]
		switch {
		case !ok:
			added = append(added, Change{
				ChangeType: ChangeAdded,
				SourceID:   row.SourceID,
				Name:       row.Name,
				EntityType: row.EntityType,
			})
		case prev.EntityHash != row.EntityHash:
			changed = append(changed, Change{
				ChangeType: ChangeChanged,
				SourceID:   row.SourceID,
				Name:       row.Name,
				EntityType: row.EntityType,
			})
		}
	}
	for _, row := range fromRows {
		if _, ok := toMap[row.SourceID]; !ok {
			removed = append(removed, Change{
				ChangeType: ChangeRemoved,
				SourceID:   row.SourceID,
				Name:       row.Name,
				EntityType: row.EntityType,
			})
		}
	}

	counts := DiffCounts{Added: len(added), Removed: len(removed), Changed: len(changed)}
	added, removed, changed, next := pageChanges(added, removed, changed, after, limit)

	if strings.EqualFold(detail, DetailFull) {
		if err := s.attachDiffEntities(ctx, listName, fromRef.ID, toRef.ID, added, removed, changed); err != nil {
			return out, err
		}
	}

	out = DiffResponse{
		ListName: listName,
		From:     *fromRef,
		To:       *toRef,
		Counts:   counts,
		Added:    added,
		Removed:  removed,
		Changed:  changed,
		After:    next,
	}
	return out, nil
}

func pageChanges(added, removed, changed []Change, after string, limit int) ([]Change, []Change, []Change, string) {
	if added == nil {
		added = []Change{}
	}
	if removed == nil {
		removed = []Change{}
	}
	if changed == nil {
		changed = []Change{}
	}

	type item struct {
		kind string
		idx  int
	}
	seq := make([]item, 0, len(added)+len(changed)+len(removed))
	for i := range added {
		seq = append(seq, item{kind: ChangeAdded, idx: i})
	}
	for i := range changed {
		seq = append(seq, item{kind: ChangeChanged, idx: i})
	}
	for i := range removed {
		seq = append(seq, item{kind: ChangeRemoved, idx: i})
	}

	start := 0
	if after != "" {
		kind, id, ok := splitCursor(after)
		if ok {
			for i, it := range seq {
				cur := changeAt(added, removed, changed, it)
				if it.kind == kind && cur.SourceID == id {
					start = i + 1
					break
				}
			}
		}
	}
	if start > len(seq) {
		start = len(seq)
	}
	end := start + limit
	if end > len(seq) {
		end = len(seq)
	}
	page := seq[start:end]

	var outAdded, outRemoved, outChanged []Change
	for _, it := range page {
		cur := changeAt(added, removed, changed, it)
		switch it.kind {
		case ChangeAdded:
			outAdded = append(outAdded, cur)
		case ChangeRemoved:
			outRemoved = append(outRemoved, cur)
		case ChangeChanged:
			outChanged = append(outChanged, cur)
		}
	}
	if outAdded == nil {
		outAdded = []Change{}
	}
	if outRemoved == nil {
		outRemoved = []Change{}
	}
	if outChanged == nil {
		outChanged = []Change{}
	}

	next := ""
	if end < len(seq) && len(page) > 0 {
		last := changeAt(added, removed, changed, page[len(page)-1])
		next = page[len(page)-1].kind + ":" + last.SourceID
	}
	return outAdded, outRemoved, outChanged, next
}

func changeAt(added, removed, changed []Change, it struct {
	kind string
	idx  int
}) Change {
	switch it.kind {
	case ChangeAdded:
		return added[it.idx]
	case ChangeRemoved:
		return removed[it.idx]
	default:
		return changed[it.idx]
	}
}

func splitCursor(after string) (string, string, bool) {
	kind, id, ok := strings.Cut(after, ":")
	if !ok || kind == "" || id == "" {
		return "", "", false
	}
	switch kind {
	case ChangeAdded, ChangeRemoved, ChangeChanged:
		return kind, id, true
	default:
		return "", "", false
	}
}

func (s *Service) attachDiffEntities(ctx context.Context, listName string, fromID, toID int64, added, removed, changed []Change) error {
	addedIDs := sourceIDs(added)
	removedIDs := sourceIDs(removed)
	changedIDs := sourceIDs(changed)

	toEntities, err := s.repo.EntitiesBySourceIDs(ctx, listName, toID, append(addedIDs, changedIDs...))
	if err != nil {
		return fmt.Errorf("loading to entities: %w", err)
	}
	fromEntities, err := s.repo.EntitiesBySourceIDs(ctx, listName, fromID, append(removedIDs, changedIDs...))
	if err != nil {
		return fmt.Errorf("loading from entities: %w", err)
	}

	for i := range added {
		if ent, ok := toEntities[added[i].SourceID]; ok {
			cp := ent
			added[i].After = &cp
		}
	}
	for i := range removed {
		if ent, ok := fromEntities[removed[i].SourceID]; ok {
			cp := ent
			removed[i].Before = &cp
		}
	}
	for i := range changed {
		if ent, ok := fromEntities[changed[i].SourceID]; ok {
			cp := ent
			changed[i].Before = &cp
		}
		if ent, ok := toEntities[changed[i].SourceID]; ok {
			cp := ent
			changed[i].After = &cp
		}
	}
	return nil
}

func sourceIDs(changes []Change) []string {
	out := make([]string, 0, len(changes))
	for i := range changes {
		out = append(out, changes[i].SourceID)
	}
	return out
}

func validListName(s string) bool {
	if s == "" || len(s) > 64 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' || c == '-' {
			continue
		}
		return false
	}
	return true
}

func validRefToken(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" || strings.EqualFold(s, RefLatest) {
		return true
	}
	if len(s) > 64 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		}
		return false
	}
	return true
}
