package index

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
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
	// It applies source/type partitioning, name/crypto inverted indexes, and hashed
	// government-ID / address blocking keys, with safe fallbacks that never reduce
	// recall below a full partition scan.
	SelectCandidates(ctx context.Context, query search.Entity[search.Value]) (Candidates, error)
	Update(latest download.Stats)
	LatestStats() download.Stats
	GetTFIDFIndex() *tfidf.Index
	// RefreshIngest rebuilds the ingested-file corpus from the repository.
	RefreshIngest(ctx context.Context) error
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

	ingestRepository  ingest.Repository
	ingestMu          sync.Mutex // serializes ingest corpus rebuilds
	ingestCorpus      *corpus
	ingestFingerprint string
	ingestCounts      map[string]int
	ingestHashes      map[string]string
}

func (l *lists) downloadedSource(source search.SourceList) bool {
	if source.IsRequestType() {
		return true
	}
	src := string(source)
	if src == "" {
		return true
	}
	_, ok := l.latestStats.Lists[src]
	return ok
}

func (l *lists) GetEntities(ctx context.Context, source search.SourceList) ([]search.Entity[search.Value], error) {
	if err := l.prepareIngest(ctx, source); err != nil {
		return nil, err
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	src := string(source)
	downloaded := l.downloadedSource(source)

	if src == "" {
		var out []search.Entity[search.Value]
		if l.corpus != nil {
			out = append(out, l.corpus.entities...)
		}
		if l.ingestCorpus != nil {
			out = append(out, l.ingestCorpus.entities...)
		}
		if out != nil {
			return out, nil
		}
		return l.latestStats.Entities, nil
	}

	if downloaded && !source.IsRequestType() && l.corpus != nil {
		idxs, _ := l.corpus.partitionIndices(source, "")
		return l.corpus.materialize(idxs), nil
	}
	if source.IsRequestType() && l.corpus != nil {
		return l.corpus.entities, nil
	}
	if l.ingestLoaded(source) {
		idxs, _ := l.ingestCorpus.partitionIndices(source, "")
		return l.ingestCorpus.materialize(idxs), nil
	}
	if downloaded && l.corpus != nil {
		return l.corpus.entities, nil
	}
	if l.ingestRepository != nil {
		return nil, nil
	}

	return nil, fmt.Errorf("source %s not found", source)
}

func (l *lists) SelectCandidates(ctx context.Context, query search.Entity[search.Value]) (Candidates, error) {
	source := query.Source
	if err := l.prepareIngest(ctx, source); err != nil {
		return Candidates{}, err
	}

	l.mu.RLock()
	defer l.mu.RUnlock()

	downloaded := l.downloadedSource(source)
	ingested := l.ingestLoaded(source)

	if string(source) == "" {
		var dl, ing Candidates
		if l.corpus != nil {
			dl = l.corpus.selectCandidates(query, CandidateOpts{})
		}
		if l.ingestCorpus != nil && l.ingestFingerprint != "" {
			ing = l.ingestCorpus.selectCandidates(query, CandidateOpts{})
		}
		merged := mergeCandidates(dl, ing)
		if merged.Len() > 0 || l.corpus != nil || l.ingestCorpus != nil {
			return merged, nil
		}
		return candidatesFromEntities(l.latestStats.Entities, l.latestStats.TFIDFIndex), nil
	}

	if downloaded && l.corpus != nil {
		return l.corpus.selectCandidates(query, CandidateOpts{}), nil
	}
	if ingested {
		return l.ingestCorpus.selectCandidates(query, CandidateOpts{}), nil
	}
	if downloaded {
		return candidatesFromEntities(l.latestStats.Entities, l.latestStats.TFIDFIndex), nil
	}
	if l.ingestRepository != nil {
		return Candidates{}, nil
	}

	return Candidates{}, fmt.Errorf("source %s not found", source)
}

func (l *lists) prepareIngest(ctx context.Context, source search.SourceList) error {
	if l.ingestRepository == nil {
		return nil
	}
	src := string(source)
	if src != "" && !source.IsRequestType() {
		l.mu.RLock()
		_, downloaded := l.latestStats.Lists[src]
		l.mu.RUnlock()
		if downloaded {
			return nil
		}
	}
	return l.ensureIngest(ctx)
}

func (l *lists) LatestStats() download.Stats {
	l.mu.RLock()
	defer l.mu.RUnlock()

	lists := l.latestStats.Lists
	hashes := l.latestStats.ListHashes
	if len(l.ingestCounts) > 0 {
		lists = copyStringInt(l.latestStats.Lists)
		hashes = copyStringString(l.latestStats.ListHashes)
		if lists == nil {
			lists = make(map[string]int, len(l.ingestCounts))
		}
		if hashes == nil {
			hashes = make(map[string]string, len(l.ingestHashes))
		}
		for src, n := range l.ingestCounts {
			lists[src] = n
			hashes[src] = l.ingestHashes[src]
		}
	}

	return download.Stats{
		Lists:      lists,
		ListHashes: hashes,
		StartedAt:  l.latestStats.StartedAt,
		EndedAt:    l.latestStats.EndedAt,
		Version:    watchman.Version,
	}
}

// Update replaces the searchable corpus with a newly downloaded generation.
//
// Memory ownership after Update:
//   - The entity slice and TF-IDF index live only on the in-memory corpus.
//   - latestStats keeps list metadata (counts, hashes, timestamps) only.
//
// Callers should drop their own references to stats.Entities after Update so the
// previous generation (and download temps) can be collected. Update runs a GC
// and FreeOSMemory after the swap to reduce refresh peak RSS / OOM risk when the
// old and new corpora would otherwise coexist until the next natural GC cycle.
func (l *lists) Update(latest download.Stats) {
	// Build search corpus outside the write lock (CPU-heavy).
	// corpus.entities aliases latest.Entities (no extra copy of the slice).
	c := buildCorpus(latest.Entities, latest.TFIDFIndex)

	// Metadata only — do not retain a second root to the entity slice / TF-IDF
	// via latestStats (LatestStats() already omitted Entities; this makes the
	// in-process graph match that intent).
	meta := latest
	meta.Entities = nil
	meta.TFIDFIndex = nil

	l.mu.Lock()
	// Overwriting l.corpus drops the only lists-owned root to the previous
	// generation; nothing else in this frame retains it.
	l.latestStats = meta
	l.corpus = c
	l.mu.Unlock()

	// Refresh is infrequent; reclaim the prior generation promptly so cgroup
	// peaks are closer to one corpus than two. FreeOSMemory returns idle heap
	// pages to the OS (important after large refreshes).
	runtime.GC()
	debug.FreeOSMemory()
}

func (l *lists) GetTFIDFIndex() *tfidf.Index {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.corpus != nil {
		return l.corpus.tfidf
	}
	return l.latestStats.TFIDFIndex
}
