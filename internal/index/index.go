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

	ingestRepository ingest.Repository
}

func (l *lists) GetEntities(ctx context.Context, source search.SourceList) ([]search.Entity[search.Value], error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	_, exists := l.latestStats.Lists[string(source)]

	// Let api-request use our inmem entities
	exists = exists || source == search.SourceAPIRequest
	if string(source) == "" {
		exists = true
	}

	if exists {
		return l.latestStats.Entities, nil
	}

	// Check the repository
	if l.ingestRepository != nil {
		// TODO(adam): need to support pagination
		return l.ingestRepository.ListBySource(ctx, "", source, 1000)
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
	l.mu.Lock()
	defer l.mu.Unlock()

	l.latestStats = latest
}

func (l *lists) GetTFIDFIndex() *tfidf.Index {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.latestStats.TFIDFIndex
}
