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
	ctx := context.Background()
	opts := SearchOpts{Limit: 10, MinMatch: 0.01}

	svc := testService(t)

	t.Run("basic", func(t *testing.T) {
		results, err := svc.Search(ctx, search.Entity[search.Value]{
			Name: "SHIPPING LIMITED",
			Type: search.EntityBusiness,
		}, opts)
		require.NoError(t, err)
		require.Greater(t, len(results), 0)

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

		res := results[0]
		require.InDelta(t, 1.00, res.Match, 0.001)

		// 36216
	})
}

func testService(tb testing.TB) Service {
	files := testInputs(tb,
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "sdn.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "alt.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "add.csv"),
		filepath.Join("..", "..", "pkg", "ofac", "testdata", "sdn_comments.csv"),
	)
	ofacRecords, err := ofac.Read(files)
	require.NoError(tb, err)

	entities := ofac.GroupIntoEntities(ofacRecords.SDNs, ofacRecords.Addresses, ofacRecords.SDNComments, ofacRecords.AlternateIdentities)

	logger := log.NewTestLogger()

	svc := NewService(logger)
	svc.UpdateEntities(entities)

	return svc
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
