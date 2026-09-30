package search

import (
	"context"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/fshelp"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/ofac"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestAttachDebugDetails_ThresholdInsideReturnedSet(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "Ada Lovelace",
		Type: search.EntityPerson,
		Person: &search.Person{
			Name: "Ada Lovelace",
		},
	}.Normalize()

	out := []search.SearchedEntity[search.Value]{
		{Entity: query, Match: 0.75},
		{Entity: query, Match: 0.79},
		{Entity: query, Match: 0.80},
		{Entity: query, Match: 0.91},
	}

	attachDebugDetails(query, nil, SearchOpts{Debug: true, DebugMinMatch: 0.80}, out)

	require.Empty(t, out[0].Debug)
	require.Empty(t, out[0].Details.Pieces)
	require.Empty(t, out[1].Debug)
	require.Empty(t, out[1].Details.Pieces)

	require.NotEmpty(t, out[2].Debug)
	require.NotEmpty(t, out[2].Details.Pieces)
	require.NotEmpty(t, out[3].Debug)
	require.NotEmpty(t, out[3].Details.Pieces)
}

func TestService_Search(t *testing.T) {
	ctx := context.Background()
	opts := SearchOpts{Limit: 10, MinMatch: 0.01, Debug: testing.Verbose()}

	svc := testService(t)

	t.Run("minMatch 0.75 with debug 0.80", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Name: "Mohammad",
			Type: search.EntityPerson,
		}.Normalize()

		results, err := svc.Search(ctx, query, SearchOpts{
			Limit:         20,
			MinMatch:      0.75,
			Debug:         true,
			DebugMinMatch: 0.80,
		})
		require.NoError(t, err)
		require.NotEmpty(t, results)

		var belowDebug int
		for _, ent := range results {
			require.GreaterOrEqual(t, ent.Match, 0.75)
			if ent.Match >= 0.80 {
				require.NotEmpty(t, ent.Debug, "match=%.4f should include debug", ent.Match)
				require.NotEmpty(t, ent.Details.Pieces)
			} else {
				require.Empty(t, ent.Debug, "match=%.4f should omit debug", ent.Match)
				require.Empty(t, ent.Details.Pieces)
				belowDebug++
			}
		}
		require.Greater(t, belowDebug, 0, "Mohammad should return hits in [0.75, 0.80) without debug")

		exact, err := svc.Search(ctx, search.Entity[search.Value]{
			Name: "Dmitry Yuryevich Khoroshev",
			Type: search.EntityPerson,
		}.Normalize(), SearchOpts{
			Limit:         5,
			MinMatch:      0.75,
			Debug:         true,
			DebugMinMatch: 0.80,
		})
		require.NoError(t, err)
		require.NotEmpty(t, exact)

		var withPieces int
		for _, ent := range exact {
			require.GreaterOrEqual(t, ent.Match, 0.75)
			if ent.Match >= 0.80 {
				require.NotEmpty(t, ent.Debug, "match=%.4f should include debug", ent.Match)
				require.NotEmpty(t, ent.Details.Pieces)
				withPieces++
			} else {
				require.Empty(t, ent.Debug, "match=%.4f should omit debug", ent.Match)
			}
		}
		require.Greater(t, withPieces, 0, "full-name Khoroshev should attach debug on hits >= 0.80")
	})

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

			_, err := base64.StdEncoding.DecodeString(results[0].Debug)
			require.NoError(t, err)
		}
		require.Greater(t, len(results), 0)

		res := results[0]
		require.InDelta(t, 1.00, res.Match, 0.001) // 36216
	})
}

func testService(tb testing.TB) Service {
	tb.Helper()

	pkg, err := fshelp.FindPkgDir()
	require.NoError(tb, err)

	files := testInputs(tb,
		filepath.Join(pkg, "sources", "ofac", "testdata", "sdn.csv"),
		filepath.Join(pkg, "sources", "ofac", "testdata", "alt.csv"),
		filepath.Join(pkg, "sources", "ofac", "testdata", "add.csv"),
		filepath.Join(pkg, "sources", "ofac", "testdata", "sdn_comments.csv"),
	)
	ofacRecords, err := ofac.Read(files)
	require.NoError(tb, err)

	entities := ofac.GroupIntoEntities(ofacRecords.SDNs, ofacRecords.Addresses, ofacRecords.SDNComments, ofacRecords.AlternateIdentities)

	logger := log.NewTestLogger()

	indexedLists := index.NewLists(nil) // only in-mem

	searchConfig := DefaultConfig()
	svc, err := NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(tb, err)

	indexedLists.Update(download.Stats{
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
