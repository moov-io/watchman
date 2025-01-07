package search

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/pkg/ofac"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestService_Search(t *testing.T) {
	files := testInputs(t,
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "sdn.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "alt.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "add.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "sdn_comments.csv"),
	)
	ofacRecords, err := ofac.Read(files)
	require.NoError(t, err)

	sdns := depointer(ofacRecords.SDNs)
	addrs := depointer(ofacRecords.Addresses)
	comments := depointer(ofacRecords.SDNComments)
	alts := depointer(ofacRecords.AlternateIdentities)
	entities := ofac.ToEntities(sdns, addrs, comments, alts)

	ctx := context.Background()
	logger := log.NewTestLogger()

	opts := SearchOpts{Limit: 10, MinMatch: 0.01}

	svc := NewService(logger, entities)

	t.Run("basic", func(t *testing.T) {
		results, err := svc.Search(ctx, search.Entity[search.Value]{
			Name: "SHIPPING LIMITED",
		}, opts)
		require.NoError(t, err)
		require.Greater(t, len(results), 1)

		t.Logf("got %d results", len(results))
		t.Logf("")
		t.Logf("%#v", results[0])
		t.Logf("")
		t.Logf("%#v", results[1])
	})

	t.Run("crypto address", func(t *testing.T) {
		results, err := svc.Search(ctx, search.Entity[search.Value]{
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "12VrYZgS1nmf9KHHped24xBb1aLLRpV2cT"},
			},
		}, opts)
		require.NoError(t, err)
		require.Greater(t, len(results), 0)

		t.Logf("got %d results", len(results))
		t.Logf("")
		t.Logf("%#v", results[0])
	})
}

func TestService_makeIndicies(t *testing.T) {
	indices := makeIndices(122, 5)
	require.Len(t, indices, 7)

	expected := []int{0, 24, 48, 72, 96, 120, 122}
	require.Equal(t, expected, indices)
}

func testInputs(tb testing.TB, paths ...string) map[string]io.ReadCloser {
	tb.Helper()

	input := make(map[string]io.ReadCloser)
	for _, path := range paths {
		_, filename := filepath.Split(path)

		fd, err := os.Open(path)
		require.NoError(tb, err)

		input[filename] = fd
	}
	return input
}

func depointer[T any](input []*T) []T {
	out := make([]T, len(input))
	for i := range input {
		if input[i] != nil {
			out[i] = *input[i]
		}
	}
	return out
}
