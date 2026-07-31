package corpusfile

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/ofac"

	"github.com/stretchr/testify/require"
)

func TestSplitEntity_ClearsSourceData(t *testing.T) {
	e := ofactest.FindEntity(t, "29702")
	require.NotNil(t, e.SourceData)

	split := SplitEntity(e)
	require.Nil(t, split.Hot.SourceData)
	require.NotNil(t, split.Cold.SourceData)
	// Raw display fields move cold; prepared scoring fields stay hot.
	require.Nil(t, split.Hot.Addresses)
	require.Equal(t, e.Addresses, split.Cold.Addresses)
	require.Equal(t, e.PreparedFields.Addresses, split.Hot.PreparedFields.Addresses)
	require.Empty(t, split.Hot.Contact.PhoneNumbers)
	require.Equal(t, e.Contact.PhoneNumbers, split.Cold.PhoneNumbers)
	require.Equal(t, e.PreparedFields.Contact.PhoneNumbers, split.Hot.PreparedFields.Contact.PhoneNumbers)

	// Scoring fields remain
	require.Equal(t, e.Name, split.Hot.Name)
	require.Equal(t, e.Type, split.Hot.Type)
	require.Equal(t, e.PreparedFields.Name, split.Hot.PreparedFields.Name)

	// Similarity must not depend on cold fields
	other := ofactest.FindEntity(t, "44525")
	require.InDelta(t,
		search.Similarity(e, other),
		search.Similarity(split.Hot, SplitEntity(other).Hot),
		1e-9,
	)

	full := Reattach(split.Hot, split.Cold)
	require.Equal(t, e.SourceData, full.SourceData)
	require.Equal(t, e.Addresses, full.Addresses)
	require.Equal(t, e.Contact.PhoneNumbers, full.Contact.PhoneNumbers)
}

func TestSnapshot_RoundTrip(t *testing.T) {
	entities := loadOFACEntities(t)
	require.NotEmpty(t, entities)

	path := filepath.Join(t.TempDir(), "corpus.snap")
	snap, err := BuildSnapshot(path, entities)
	require.NoError(t, err)
	t.Cleanup(func() { snap.Close() })

	require.Equal(t, len(entities), snap.Len())
	require.Equal(t, uint32(Version), snap.Header().Version)

	// Spot-check a few entities
	for _, id := range []string{"29702", "44525", "50972"} {
		idx := findBySourceID(entities, id)
		require.GreaterOrEqual(t, idx, 0, "missing %s", id)

		hot := snap.Hot(idx)
		require.Nil(t, hot.SourceData)
		require.Equal(t, entities[idx].Name, hot.Name)

		// Similarity must match baseline (SourceData unused)
		base := search.Similarity(entities[idx], entities[(idx+1)%len(entities)])
		got := search.Similarity(hot, snap.Hot((idx+1)%len(entities)))
		require.Equal(t, base, got)

		hydrated, err := snap.Hydrate(idx)
		require.NoError(t, err)
		sdn, ok := hydrated.SourceData.(ofac.SDN)
		require.True(t, ok)
		require.Equal(t, id, sdn.EntityID)
	}

	size, err := snap.FileSize()
	require.NoError(t, err)
	require.Greater(t, size, int64(HeaderSize))
	t.Logf("snapshot: %d entities, file=%d bytes (%.2f MB)", snap.Len(), size, float64(size)/(1<<20))
}

func TestHeader_RoundTrip(t *testing.T) {
	h := Header{
		Version:         Version,
		Flags:           FlagHasStringTable,
		EntityCount:     42,
		StringTableOff:  128,
		StringTableLen:  1000,
		ColdDirOff:      1128,
		ColdDirLen:      42 * coldDirEntrySize,
		ColdBlobOff:     2000,
		ColdBlobLen:     99999,
		CreatedUnixNano: 1234567890,
	}
	copy(h.Magic[:], Magic)

	var buf [HeaderSize]byte
	PutHeader(buf[:], h)
	got, err := ReadHeader(buf[:])
	require.NoError(t, err)
	require.Equal(t, h, got)
}

func loadOFACEntities(tb testing.TB) []search.Entity[search.Value] {
	tb.Helper()
	// Warm ofactest (small fixture for unit tests is fine).
	_ = ofactest.FindEntity(tb, "29702")
	stats, err := ofactest.GetDownloader(tb).RefreshAll(tb.Context())
	require.NoError(tb, err)
	require.NotEmpty(tb, stats.Entities)
	return stats.Entities
}

func findBySourceID(entities []search.Entity[search.Value], id string) int {
	for i := range entities {
		if entities[i].SourceID == id {
			return i
		}
	}
	return -1
}
