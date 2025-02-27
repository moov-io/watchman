package search

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/ofac"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestService_Search(t *testing.T) {
	ctx := context.Background()
	opts := SearchOpts{Limit: 10, MinMatch: 0.01, Debug: testing.Verbose()}

	svc := testService(t)

	t.Run("basic", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "SHIPPING LIMITED",
			Type: search.EntityBusiness,
		}
		results, err := svc.Search(ctx, query.Normalize(), opts)
		require.NoError(t, err)
		require.Greater(t, len(results), 0)

		t.Logf("got %d results", len(results))
		t.Logf("")
		t.Logf("%#v", results[0])
		t.Logf("")
		t.Logf("%#v", results[1])
	})

	t.Run("crypto address", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			CryptoAddresses: []search.CryptoAddress{
				{Currency: "XBT", Address: "12VrYZgS1nmf9KHHped24xBb1aLLRpV2cT"},
			},
		}
		results, err := svc.Search(ctx, query.Normalize(), opts)
		require.NoError(t, err)

		t.Logf("got %d results", len(results))
		if len(results) > 0 {
			t.Logf("match: %.2f", results[0].Match)
			t.Logf("%#v", results[0].Entity)

			bs, err := base64.StdEncoding.DecodeString(results[0].Debug)
			require.NoError(t, err)
			fmt.Println(string(bs))
		}
		require.Greater(t, len(results), 0)

		res := results[0]
		require.InDelta(t, 1.00, res.Match, 0.001) // 36216
	})
}

func testService(tb testing.TB) Service {
	tb.Helper()

	files := testInputs(tb,
		filepath.Join("..", "..", "pkg", "sources", "ofac", "testdata", "sdn.csv"),
		filepath.Join("..", "..", "pkg", "sources", "ofac", "testdata", "alt.csv"),
		filepath.Join("..", "..", "pkg", "sources", "ofac", "testdata", "add.csv"),
		filepath.Join("..", "..", "pkg", "sources", "ofac", "testdata", "sdn_comments.csv"),
	)
	ofacRecords, err := ofac.Read(files)
	require.NoError(tb, err)

	entities := ofac.GroupIntoEntities(ofacRecords.SDNs, ofacRecords.Addresses, ofacRecords.SDNComments, ofacRecords.AlternateIdentities)

	logger := log.NewTestLogger()

	searchConfig := DefaultConfig()
	svc, err := NewService(logger, searchConfig)
	require.NoError(tb, err)

	svc.UpdateEntities(download.Stats{
		Entities: entities,
	})

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
