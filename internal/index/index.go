package index

import (
	"context"
	"fmt"
	"sync"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

type Lists interface {
	GetEntities(ctx context.Context, source search.SourceList) ([]search.Entity[search.Value], error)
	// SelectCandidates returns the subset of entities that should be scored for the query.
	// It applies source/type partitioning and name/crypto inverted indexes, with safe
	// fallbacks that never reduce recall below a full partition scan.
	SelectCandidates(ctx context.Context, query search.Entity[search.Value]) ([]search.Entity[search.Value], error)
	Update(latest download.Stats)
	LatestStats() download.Stats
	GetTFIDFIndex() *tfidf.Index
}

func NewLists(ingestRepository ingest.Repository) Lists {
	return &lists{
		ingestRepository: ingestRepository,
	}
}

type lists struct {
	mu          sync.RWMutex
	latestStats download.Stats
	corpus      *corpus

	ingestRepository ingest.Repository
}

func (l *lists) GetEntities(ctx context.Context, source search.SourceList) ([]search.Entity[search.Value], error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	_, exists := l.latestStats.Lists[string(source)]

	// Let api-request use our inmem entities
	exists = exists || source.IsRequestType()
	if string(source) == "" {
		exists = true
	}

	if exists {
		if l.corpus != nil {
			// Prefer partitioned view when a specific source is requested.
			// Always return the partition (even if empty) so we never leak entities
			// from other sources when the requested partition has no rows.
			if src := string(source); src != "" && !source.IsRequestType() {
				idxs, _ := l.corpus.partitionIndices(source, "")
				return l.corpus.materialize(idxs), nil
			}
			return l.corpus.entities, nil
		}
		return l.latestStats.Entities, nil
	}

	// Check the repository
	if l.ingestRepository != nil {
		// TODO(adam): need to support pagination
		return l.ingestRepository.ListBySource(ctx, "", source, 1000)
	}

	return nil, fmt.Errorf("source %s not found", source)
}

func (l *lists) SelectCandidates(ctx context.Context, query search.Entity[search.Value]) ([]search.Entity[search.Value], error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	source := query.Source
	_, exists := l.latestStats.Lists[string(source)]
	exists = exists || source.IsRequestType() || string(source) == ""

	if exists && l.corpus != nil {
		return l.corpus.selectCandidates(query, CandidateOpts{}), nil
	}

	// Ingested-only source: pull from repository and return as-is (no inverted index)
	if !exists && l.ingestRepository != nil {
		return l.ingestRepository.ListBySource(ctx, "", source, 1000)
	}

	if exists {
		// Corpus not built yet — return full list
		return l.latestStats.Entities, nil
	}

	return nil, fmt.Errorf("source %s not found", source)
}

func (l *lists) LatestStats() download.Stats {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// Only bring over what fields we need
	out := download.Stats{
		Lists:      l.latestStats.Lists,
		ListHashes: l.latestStats.ListHashes,
		StartedAt:  l.latestStats.StartedAt,
		EndedAt:    l.latestStats.EndedAt,
		Version:    watchman.Version,
	}
	return out
}

func (l *lists) Update(latest download.Stats) {
	// Build search corpus outside the write lock (CPU-heavy)
	c := buildCorpus(latest.Entities, latest.TFIDFIndex)

	l.mu.Lock()
	defer l.mu.Unlock()

	l.latestStats = latest
	l.corpus = c
}

func (l *lists) GetTFIDFIndex() *tfidf.Index {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.corpus != nil {
		return l.corpus.tfidf
	}
	return l.latestStats.TFIDFIndex
}
