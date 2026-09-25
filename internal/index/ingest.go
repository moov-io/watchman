package index

import (
	"context"
	"fmt"

	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/pkg/search"

	"go.opentelemetry.io/otel/attribute"
)

// RefreshIngest rebuilds the in-memory ingest corpus from the repository.
// POST /v2/ingest calls this so the next search sees the new rows without waiting
// for a checksum miss. Startup also calls it to warm the cache.
func (l *lists) RefreshIngest(ctx context.Context) error {
	if l.ingestRepository == nil {
		return nil
	}

	l.ingestMu.Lock()
	defer l.ingestMu.Unlock()
	return l.rebuildIngestLocked(ctx)
}

func (l *lists) ensureIngest(ctx context.Context) error {
	if l.ingestRepository == nil {
		return nil
	}

	sums, err := l.ingestRepository.Checksums(ctx)
	if err != nil {
		return fmt.Errorf("ingest checksums: %w", err)
	}
	fp := ingest.Fingerprint(sums)

	l.mu.RLock()
	hit := l.ingestCorpus != nil && l.ingestFingerprint == fp
	l.mu.RUnlock()
	if hit {
		return nil
	}

	l.ingestMu.Lock()
	defer l.ingestMu.Unlock()

	l.mu.RLock()
	hit = l.ingestCorpus != nil && l.ingestFingerprint == fp
	l.mu.RUnlock()
	if hit {
		return nil
	}

	return l.rebuildIngestLocked(ctx)
}

func (l *lists) rebuildIngestLocked(ctx context.Context) error {
	ctx, span := telemetry.StartSpan(ctx, "ingest-rebuild-corpus")
	defer span.End()

	sums, err := l.ingestRepository.Checksums(ctx)
	if err != nil {
		return fmt.Errorf("ingest checksums: %w", err)
	}
	fp := ingest.Fingerprint(sums)

	var entities []search.Entity[search.Value]
	if fp != "" {
		entities, err = l.ingestRepository.ListAll(ctx)
		if err != nil {
			return fmt.Errorf("listing ingested entities: %w", err)
		}
	}

	c := buildCorpus(entities, nil)

	counts := make(map[string]int, len(sums))
	hashes := make(map[string]string, len(sums))
	total := 0
	for _, s := range sums {
		counts[s.Source] = s.EntityCount
		hashes[s.Source] = s.Checksum
		total += s.EntityCount
	}

	span.SetAttributes(
		attribute.Int("ingest.source_count", len(sums)),
		attribute.Int("ingest.entity_count", total),
		attribute.String("ingest.fingerprint", fp),
	)

	l.mu.Lock()
	l.ingestCorpus = c
	l.ingestFingerprint = fp
	l.ingestCounts = counts
	l.ingestHashes = hashes
	l.mu.Unlock()
	return nil
}

func (l *lists) ingestLoaded(source search.SourceList) bool {
	if l.ingestCorpus == nil {
		return false
	}
	src := string(source)
	if src == "" {
		return l.ingestFingerprint != ""
	}
	_, ok := l.ingestCounts[src]
	return ok
}

func mergeCandidates(a, b Candidates) Candidates {
	if a.Len() == 0 {
		return b
	}
	if b.Len() == 0 {
		return a
	}

	n := a.Len() + b.Len()
	ents := make([]search.Entity[search.Value], 0, n)
	idxs := make([]int, n)
	for i := 0; i < a.Len(); i++ {
		ents = append(ents, a.At(i))
		idxs[i] = i
	}
	base := len(ents)
	for i := 0; i < b.Len(); i++ {
		ents = append(ents, b.At(i))
		idxs[base+i] = base + i
	}

	tfidf := a.TFIDF
	if tfidf == nil {
		tfidf = b.TFIDF
	}
	return Candidates{Entities: ents, Indices: idxs, TFIDF: tfidf}
}

func copyStringInt(in map[string]int) map[string]int {
	if in == nil {
		return nil
	}
	out := make(map[string]int, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func copyStringString(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
