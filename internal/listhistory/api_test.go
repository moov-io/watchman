// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestAPI_NotConfigured(t *testing.T) {
	router := mux.NewRouter()
	NewController(log.NewTestLogger(), NewService(log.NewTestLogger(), nil)).AppendRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/v2/lists", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestAPI_ListHistory(t *testing.T) {
	repo := NewMemoryRepository()
	svc := NewService(log.NewTestLogger(), repo)
	router := mux.NewRouter()
	NewController(log.NewTestLogger(), svc).AppendRoutes(router)

	server := httptest.NewServer(router)
	t.Cleanup(server.Close)
	client := pubsearch.NewClient(nil, server.URL)

	ctx := context.Background()
	alice := testPerson("1", "Alice")
	bob := testPerson("2", "Bob")
	carol := testPerson("3", "Carol")
	svc.StoreRefreshAsync(ctx, testStats("aaa", testTime(1), alice, bob))
	require.NoError(t, svc.Wait(ctx))
	alice2 := testPerson("1", "Alice Ann")
	svc.StoreRefreshAsync(ctx, testStats("bbb", testTime(2), alice2, carol))
	require.NoError(t, svc.Wait(ctx))

	t.Run("lists", func(t *testing.T) {
		out, err := client.Lists(ctx)
		require.NoError(t, err)
		require.Len(t, out.Lists, 1)
		require.Equal(t, "us_ofac", out.Lists[0].ListName)
		require.Equal(t, "bbb", out.Lists[0].Ref)
	})

	t.Run("refs", func(t *testing.T) {
		out, err := client.ListRefs(ctx, "us_ofac")
		require.NoError(t, err)
		require.Equal(t, "us_ofac", out.ListName)
		require.Len(t, out.Refs, 2)
		require.Equal(t, "bbb", out.Refs[0].Ref)
	})

	t.Run("entities", func(t *testing.T) {
		out, err := client.ListEntities(ctx, "us_ofac", pubsearch.ListEntitiesOpts{Limit: 1})
		require.NoError(t, err)
		require.Equal(t, "bbb", out.Ref.Ref)
		require.Len(t, out.Entities, 1)
		require.Equal(t, "1", out.Entities[0].SourceID)
		require.NotEmpty(t, out.After)

		page2, err := client.ListEntities(ctx, "us_ofac", pubsearch.ListEntitiesOpts{After: out.After})
		require.NoError(t, err)
		require.NotEmpty(t, page2.Entities)
		require.Greater(t, page2.Entities[0].SourceID, out.Entities[0].SourceID)

		named, err := client.ListEntities(ctx, "us_ofac", pubsearch.ListEntitiesOpts{Query: "carol"})
		require.NoError(t, err)
		require.Len(t, named.Entities, 1)
		require.Equal(t, "3", named.Entities[0].SourceID)
	})

	t.Run("entity", func(t *testing.T) {
		rec, err := client.ListEntity(ctx, "us_ofac", "2", "aaa")
		require.NoError(t, err)
		require.Equal(t, "Bob", rec.Name)
		require.Equal(t, pubsearch.EntityPerson, rec.EntityType)
	})

	t.Run("diff", func(t *testing.T) {
		out, err := client.ListDiff(ctx, "us_ofac", pubsearch.ListDiffOpts{Detail: DetailFull})
		require.NoError(t, err)
		require.Equal(t, "aaa", out.From.Ref)
		require.Equal(t, "bbb", out.To.Ref)
		require.Equal(t, 1, out.Counts.Added)
		require.Equal(t, 1, out.Counts.Removed)
		require.Equal(t, 1, out.Counts.Changed)
		require.NotNil(t, out.Added[0].After)
		require.Equal(t, "Carol", out.Added[0].After.Name)
	})

	t.Run("errors", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v2/lists?foo=1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/nope/refs", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/US_OFAC/refs", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/us_ofac/entities?ref=not-hex", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/us_ofac/entities?type=planet", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/us_ofac/diff?detail=tiny", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/us_ofac/entities/missing", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)

		req = httptest.NewRequest(http.MethodGet, "/v2/lists/us_ofac/diff?after=nope", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestAPI_EmptyLists(t *testing.T) {
	svc := NewService(log.NewTestLogger(), NewMemoryRepository())
	router := mux.NewRouter()
	NewController(log.NewTestLogger(), svc).AppendRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/v2/lists", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	var body ListsResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&body))
	require.NotNil(t, body.Lists)
	require.Empty(t, body.Lists)
}
