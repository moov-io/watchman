package search

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/decisions"
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

	t.Run("website", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Websites: []string{"https://www.suex.io/"},
			},
		}

		results, err := svc.Search(ctx, query.Normalize(), opts)
		require.NoError(t, err)
		require.NotEmpty(t, results)
		require.Equal(t, "33151", results[0].SourceID)
		require.Greater(t, results[0].Match, 0.0)
		require.Less(t, results[0].Match, 1.0)
	})

	t.Run("domain", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Domains: []string{"pay.suex.io"},
			},
		}

		results, err := svc.Search(ctx, query.Normalize(), opts)
		require.NoError(t, err)
		require.NotEmpty(t, results)
		require.Equal(t, "33151", results[0].SourceID)
		require.Greater(t, results[0].Match, 0.0)
		require.Less(t, results[0].Match, 1.0)
	})

	t.Run("domain multi-label public suffix", func(t *testing.T) {
		query := search.Entity[search.Value]{
			Type: search.EntityBusiness,
			Contact: search.ContactInfo{
				Domains: []string{"mail.nitc.co.ir"},
			},
		}

		results, err := svc.Search(ctx, query.Normalize(), opts)
		require.NoError(t, err)
		require.NotEmpty(t, results)
		require.Equal(t, "15117", results[0].SourceID)
		require.Greater(t, results[0].Match, 0.0)
		require.Less(t, results[0].Match, 1.0)
	})
}

func TestService_Search_AttachesDecisions(t *testing.T) {
	t.Setenv("DECISIONS_ENABLED", "")
	t.Setenv("DECISIONS_BASE_URL", "")
	t.Setenv("DECISIONS_API_KEY", "")
	t.Setenv("DECISIONS_MODEL", "")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v1/systemone", r.URL.Path)

		var req decisions.Request
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.NotEmpty(t, req.Questions)

		answers := make(map[string]decisions.Answer, len(req.Questions))
		for name, question := range req.Questions {
			switch question.Type {
			case "noul":
				answers[name] = decisions.Answer{Type: "noul", Noul: 0.93}
			case "choice":
				answers[name] = decisions.Answer{Type: "choice", Choice: search.DecisionActionNeedsReview}
			case "score":
				answers[name] = decisions.Answer{Type: "score", Score: 2.0}
			}
		}
		require.NoError(t, json.NewEncoder(w).Encode(decisions.Response{
			Model:   "nimble",
			Answers: answers,
		}))
	}))
	t.Cleanup(server.Close)

	svc := testServiceWithDecisions(t, decisions.Config{
		Enabled: true,
		BaseURL: server.URL,
		Model:   "nimble",
		Timeout: time.Second,
	})

	results, err := svc.Search(t.Context(), search.Entity[search.Value]{
		Name: "Dmitry Yuryevich Khoroshev",
		Type: search.EntityPerson,
	}.Normalize(), SearchOpts{Limit: 3, MinMatch: 0.5})
	require.NoError(t, err)
	require.NotEmpty(t, results)
	require.NotNil(t, results[0].Decision)
	require.Equal(t, 0.93, results[0].Decision.SameEntity)
	require.Equal(t, search.DecisionActionNeedsReview, results[0].Decision.Action)
	require.Equal(t, "nimble", results[0].Decision.Model)
	require.Greater(t, results[0].Match, 0.5)
}

func TestService_Search_DecisionsFailOpen(t *testing.T) {
	t.Setenv("DECISIONS_ENABLED", "")
	t.Setenv("DECISIONS_BASE_URL", "")
	t.Setenv("DECISIONS_API_KEY", "")
	t.Setenv("DECISIONS_MODEL", "")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	t.Cleanup(server.Close)

	svc := testServiceWithDecisions(t, decisions.Config{
		Enabled: true,
		BaseURL: server.URL,
		Model:   "nimble",
		Timeout: time.Second,
	})

	results, err := svc.Search(t.Context(), search.Entity[search.Value]{
		Name: "Dmitry Yuryevich Khoroshev",
		Type: search.EntityPerson,
	}.Normalize(), SearchOpts{Limit: 3, MinMatch: 0.5})
	require.NoError(t, err)
	require.NotEmpty(t, results)
	require.Nil(t, results[0].Decision)
}

func testService(tb testing.TB) Service {
	tb.Helper()
	return testServiceWithDecisions(tb, DefaultConfig().Decisions)
}

func testServiceWithDecisions(tb testing.TB, conf decisions.Config) Service {
	tb.Helper()
	tb.Setenv("DECISIONS_ENABLED", "")
	tb.Setenv("DECISIONS_API_KEY", "")
	tb.Setenv("DECISIONS_BASE_URL", "")
	tb.Setenv("DECISIONS_MODEL", "")
	tb.Setenv("DECISIONS_TIMEOUT", "")
	tb.Setenv("DECISIONS_MAX_ENTITIES", "")
	tb.Setenv("DECISIONS_MIN_MATCH", "")

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
	indexedLists := index.NewLists(nil)
	searchConfig := DefaultConfig()
	searchConfig.Decisions = conf
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
