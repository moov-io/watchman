package ingest

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestFingerprint(t *testing.T) {
	a := []SourceChecksum{
		{Source: "b", EntityCount: 2, Checksum: "bb"},
		{Source: "a", EntityCount: 1, Checksum: "aa"},
	}
	b := []SourceChecksum{
		{Source: "a", EntityCount: 1, Checksum: "aa"},
		{Source: "b", EntityCount: 2, Checksum: "bb"},
	}
	require.Equal(t, Fingerprint(a), Fingerprint(b))
	require.NotEqual(t, Fingerprint(a), Fingerprint(a[:1]))
	require.Empty(t, Fingerprint(nil))
}

func TestChecksumEntities_Stable(t *testing.T) {
	e1 := search.Entity[search.Value]{Name: "Ann", SourceID: "2", Type: search.EntityPerson}
	e2 := search.Entity[search.Value]{Name: "Bob", SourceID: "1", Type: search.EntityPerson}

	a, err := checksumEntities([]search.Entity[search.Value]{e1, e2})
	require.NoError(t, err)
	b, err := checksumEntities([]search.Entity[search.Value]{e2, e1})
	require.NoError(t, err)
	require.Equal(t, a, b)
	require.Len(t, a, 64)
}
