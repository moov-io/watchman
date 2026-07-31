package corpusfile

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/fshelp"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

// Corpus benchmarks compare today's full in-memory working set (Baseline)
// against the tiered hot-RAM + cold-disk prototype (Tiered).
//
//	go test ./internal/corpusfile/ -bench='BenchmarkCorpus_' -benchmem -count=5
//
// Key metrics:
//   - HeapAlive / heap bytes held by the corpus after GC
//   - Snapshot file size (tiered only)
//   - Similarity ns/op (score path must not regress)
//   - Hydrate ns/op (cold read cost for top-N responses)
//   - Candidate selection + score over a name query

func BenchmarkCorpus_Baseline_Size(b *testing.B) {
	// GC-stable size proxy for today's full in-memory []Entity (with SourceData).
	entities := loadBenchEntities(b)
	var report SizeReport
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := MeasureSizes(entities)
		require.NoError(b, err)
		report = r
		b.ReportMetric(float64(r.Entities), "entities")
		b.ReportMetric(float64(r.FullEntityBytes)/(1<<20), "full_MB")
		b.ReportMetric(float64(r.FullEntityBytes)/float64(r.Entities), "full_B/entity")
		b.ReportMetric(float64(r.HotBytes)/(1<<20), "hot_MB")
		b.ReportMetric(float64(r.ColdBytes)/(1<<20), "cold_MB")
	}
	runtime.KeepAlive(report)
}

func BenchmarkCorpus_Tiered_Size(b *testing.B) {
	// Size proxy for tiered layout: hot gob + on-disk snapshot file.
	entities := loadBenchEntities(b)
	dir := b.TempDir()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := MeasureSizes(entities)
		require.NoError(b, err)
		path := filepath.Join(dir, fmt.Sprintf("size-%d.snap", i))
		snap, err := BuildSnapshot(path, entities)
		require.NoError(b, err)
		size, err := snap.FileSize()
		require.NoError(b, err)
		require.NoError(b, snap.Close())
		r.SnapshotFileBytes = size

		b.ReportMetric(float64(r.Entities), "entities")
		b.ReportMetric(float64(r.HotBytes)/(1<<20), "hot_MB")
		b.ReportMetric(float64(r.HotBytes)/float64(r.Entities), "hot_B/entity")
		b.ReportMetric(float64(r.SnapshotFileBytes)/(1<<20), "cold_file_MB")
		b.ReportMetric(float64(r.SnapshotFileBytes)/float64(r.Entities), "cold_B/entity")
		// Retained scoring set is hot only; cold lives on disk.
		b.ReportMetric(float64(r.HotBytes)/(1<<20), "retained_MB")
	}
}

func BenchmarkCorpus_Baseline_Similarity(b *testing.B) {
	entities := loadBenchEntities(b)
	queries := sampleQueries(b, entities, 64)
	b.ReportMetric(float64(len(entities)), "entities")
	b.ResetTimer()

	var sink float64
	for i := 0; i < b.N; i++ {
		q := queries[i%len(queries)]
		// Score against a rotating window to approximate candidate-set work.
		start := (i * 17) % len(entities)
		for j := 0; j < 256; j++ {
			sink += search.Similarity(q, entities[(start+j)%len(entities)])
		}
	}
	runtime.KeepAlive(sink)
}

func BenchmarkCorpus_Tiered_Similarity(b *testing.B) {
	entities := loadBenchEntities(b)
	snap, err := BuildSnapshot(filepath.Join(b.TempDir(), "corpus.snap"), entities)
	require.NoError(b, err)
	b.Cleanup(func() { snap.Close() })

	hot := snap.HotSlice()
	queries := sampleQueries(b, hot, 64)
	b.ReportMetric(float64(len(hot)), "entities")
	b.ResetTimer()

	var sink float64
	for i := 0; i < b.N; i++ {
		q := queries[i%len(queries)]
		start := (i * 17) % len(hot)
		for j := 0; j < 256; j++ {
			sink += search.Similarity(q, hot[(start+j)%len(hot)])
		}
	}
	runtime.KeepAlive(sink)
}

func BenchmarkCorpus_Tiered_HydrateTopN(b *testing.B) {
	entities := loadBenchEntities(b)
	snap, err := BuildSnapshot(filepath.Join(b.TempDir(), "corpus.snap"), entities)
	require.NoError(b, err)
	b.Cleanup(func() { snap.Close() })

	// Simulate hydrating top-20 hits for an API response.
	const topN = 20
	idxs := make([]int, topN)
	for i := range idxs {
		idxs[i] = (i * 97) % snap.Len()
	}

	b.ReportMetric(float64(topN), "hydrate_n")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, idx := range idxs {
			e, err := snap.Hydrate(idx)
			if err != nil {
				b.Fatal(err)
			}
			if e.SourceData == nil {
				b.Fatal("expected SourceData after hydrate")
			}
		}
	}
}

func BenchmarkCorpus_Baseline_BuildCorpus(b *testing.B) {
	entities := loadBenchEntities(b)
	b.ReportMetric(float64(len(entities)), "entities")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lists := index.NewLists(nil)
		// download.Stats shape — Update builds partitions + inverted indexes.
		lists.Update(benchStats(entities))
		runtime.KeepAlive(lists)
	}
}

func BenchmarkCorpus_Tiered_BuildCorpus(b *testing.B) {
	entities := loadBenchEntities(b)
	dir := b.TempDir()
	b.ReportMetric(float64(len(entities)), "entities")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		path := filepath.Join(dir, fmt.Sprintf("build-%d.snap", i))
		snap, err := BuildSnapshot(path, entities)
		require.NoError(b, err)

		lists := index.NewLists(nil)
		lists.Update(benchStats(snap.HotSlice()))
		require.NoError(b, snap.Close())
		runtime.KeepAlive(lists)
	}
}

func BenchmarkCorpus_Baseline_SelectAndScore(b *testing.B) {
	entities := loadBenchEntities(b)
	lists := index.NewLists(nil)
	lists.Update(benchStats(entities))

	query := search.Entity[search.Value]{
		Name:   "Vladimir Putin",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
	}.Normalize()

	b.ReportMetric(float64(len(entities)), "entities")
	b.ResetTimer()

	var hits int
	for i := 0; i < b.N; i++ {
		cands, err := lists.SelectCandidates(b.Context(), query)
		require.NoError(b, err)
		hits = len(cands)
		for j := range cands {
			_ = search.Similarity(query, cands[j])
		}
	}
	b.ReportMetric(float64(hits), "candidates")
}

func BenchmarkCorpus_Tiered_SelectAndScore(b *testing.B) {
	entities := loadBenchEntities(b)
	snap, err := BuildSnapshot(filepath.Join(b.TempDir(), "corpus.snap"), entities)
	require.NoError(b, err)
	b.Cleanup(func() { snap.Close() })

	lists := index.NewLists(nil)
	lists.Update(benchStats(snap.HotSlice()))

	query := search.Entity[search.Value]{
		Name:   "Vladimir Putin",
		Type:   search.EntityPerson,
		Source: search.SourceUSOFAC,
	}.Normalize()

	b.ReportMetric(float64(len(entities)), "entities")
	b.ResetTimer()

	var hits int
	for i := 0; i < b.N; i++ {
		cands, err := lists.SelectCandidates(b.Context(), query)
		require.NoError(b, err)
		hits = len(cands)
		for j := range cands {
			_ = search.Similarity(query, cands[j])
		}
	}
	b.ReportMetric(float64(hits), "candidates")
}

// --- helpers ---

var (
	benchEntitiesOnce sync.Once
	benchEntities     []search.Entity[search.Value]
	benchEntitiesErr  error
)

// loadBenchEntities returns a cached full OFAC corpus for latency benches.
// Heap benches must call mustLoadOFACTestdata instead so SourceData can be dropped.
func loadBenchEntities(b *testing.B) []search.Entity[search.Value] {
	b.Helper()
	benchEntitiesOnce.Do(func() {
		benchEntities, benchEntitiesErr = loadOFACTestdata()
	})
	require.NoError(b, benchEntitiesErr)
	require.NotEmpty(b, benchEntities)
	return benchEntities
}

func mustLoadOFACTestdata(tb testing.TB) []search.Entity[search.Value] {
	tb.Helper()
	entities, err := loadOFACTestdata()
	require.NoError(tb, err)
	require.NotEmpty(tb, entities)
	return entities
}

func loadOFACTestdata() ([]search.Entity[search.Value], error) {
	pkg, err := fshelp.FindPkgDir()
	if err != nil {
		return nil, err
	}
	conf := download.Config{
		InitialDataDirectory: filepath.Join(pkg, "sources", "ofac", "testdata"),
		IncludedLists:        []search.SourceList{search.SourceUSOFAC},
	}
	dl, err := download.NewDownloader(log.NewNopLogger(), conf, nil)
	if err != nil {
		return nil, err
	}
	stats, err := dl.RefreshAll(context.Background())
	if err != nil {
		return nil, err
	}
	return stats.Entities, nil
}

func sampleQueries(tb testing.TB, entities []search.Entity[search.Value], n int) []search.Entity[search.Value] {
	tb.Helper()
	if n > len(entities) {
		n = len(entities)
	}
	out := make([]search.Entity[search.Value], n)
	step := len(entities) / n
	if step < 1 {
		step = 1
	}
	for i := 0; i < n; i++ {
		e := entities[i*step]
		e.SourceID = "" // avoid exact SourceID short-circuit in Similarity
		e.SourceData = nil
		out[i] = e
	}
	return out
}

func benchStats(entities []search.Entity[search.Value]) download.Stats {
	lists := make(map[string]int)
	for i := range entities {
		lists[string(entities[i].Source)]++
	}
	return download.Stats{
		Entities:  entities,
		Lists:     lists,
		StartedAt: time.Now().Add(-time.Second),
		EndedAt:   time.Now(),
	}
}
