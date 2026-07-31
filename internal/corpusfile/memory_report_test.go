package corpusfile

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCorpus_MemoryReport prints a GC-stable before/after size comparison using
// gob-encoded byte lengths (see MeasureSizes) plus on-disk snapshot size.
//
//	go test ./internal/corpusfile/ -run TestCorpus_MemoryReport -v
func TestCorpus_MemoryReport(t *testing.T) {
	entities := mustLoadOFACTestdata(t)

	breakdown, err := MeasureFieldBreakdown(entities)
	require.NoError(t, err)
	t.Logf("field breakdown (OFAC package testdata)\n%s", breakdown.String())

	report, err := MeasureSizes(entities)
	require.NoError(t, err)

	path := filepath.Join(t.TempDir(), "corpus.snap")
	snap, err := BuildSnapshot(path, entities)
	require.NoError(t, err)
	t.Cleanup(func() { _ = snap.Close() })

	size, err := snap.FileSize()
	require.NoError(t, err)
	report.SnapshotFileBytes = size

	t.Logf("corpus size report (OFAC package testdata)")
	t.Logf("  %s", report.String())
	t.Logf("BASELINE retained working set proxy: full gob = %.2f MB", float64(report.FullEntityBytes)/(1<<20))
	t.Logf("TIERED retained working set proxy:   hot gob  = %.2f MB + cold file %.2f MB",
		float64(report.HotBytes)/(1<<20), float64(report.SnapshotFileBytes)/(1<<20))
	if report.FullEntityBytes > 0 {
		saved := float64(report.FullEntityBytes-report.HotBytes) / float64(report.FullEntityBytes) * 100
		t.Logf("cold-stripped share of full gob:    %.1f%% (%.2f MB)",
			saved, float64(report.FullEntityBytes-report.HotBytes)/(1<<20))
	}
}

func TestEncodeCold_ZstdSmallerThanGob(t *testing.T) {
	entities := mustLoadOFACTestdata(t)
	require.Greater(t, len(entities), 100)

	var gobTotal, zstdTotal int64
	for i := 0; i < 500 && i < len(entities); i++ {
		cold := SplitEntity(entities[i]).Cold
		g, err := EncodeColdCodec(cold, CodecGob)
		require.NoError(t, err)
		z, err := EncodeColdCodec(cold, CodecGobZstd)
		require.NoError(t, err)
		gobTotal += int64(len(g))
		zstdTotal += int64(len(z))

		// Round-trip zstd
		decoded, err := DecodeCold(z)
		require.NoError(t, err)
		require.Equal(t, cold.SourceData, decoded.SourceData)
	}
	t.Logf("sample cold payloads: gob=%.2fMB zstd=%.2fMB ratio=%.2f",
		float64(gobTotal)/(1<<20), float64(zstdTotal)/(1<<20), float64(zstdTotal)/float64(gobTotal))
	require.Less(t, zstdTotal, gobTotal)
}
