// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mcp

import (
	"context"
	"encoding/json"
	"testing"

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T) *Server {
	t.Helper()
	logger := log.NewTestLogger()
	indexedLists := index.NewLists(nil)
	searchConfig := search.DefaultConfig()
	svc, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	srv, err := NewServer(logger, svc, config.MCPConfig{})
	require.NoError(t, err)
	return srv
}

func TestSearchEntities_ReturnsSearchResponse(t *testing.T) {
	srv := newTestServer(t)

	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request: SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp pubsearch.SearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.NotNil(t, resp.Query)
}

func TestSearchEntities_IncludeDetails_PopulatesScorePieces(t *testing.T) {
	srv := newTestServer(t)

	includeDetails := true
	minMatch := 0.50

	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:        SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
		MinMatch:       &minMatch,
		IncludeDetails: &includeDetails,
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp pubsearch.SearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.NotEmpty(t, resp.Entities, "expected at least one match for 'Logan'")

	// When includeDetails=true, every returned entity must carry SimilarityScore pieces.
	for _, entity := range resp.Entities {
		require.NotEmpty(t, entity.Details.Pieces,
			"entity %s/%s should have Details.Pieces when includeDetails=true",
			entity.Source, entity.SourceID)
	}
}

func TestSearchEntities_IncludeDetails_FalseOmitsPieces(t *testing.T) {
	srv := newTestServer(t)

	includeDetails := false
	minMatch := 0.50

	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:        SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
		MinMatch:       &minMatch,
		IncludeDetails: &includeDetails,
	})
	require.NoError(t, err)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp pubsearch.SearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.NotEmpty(t, resp.Entities)
	for _, entity := range resp.Entities {
		require.Empty(t, entity.Details.Pieces,
			"Details.Pieces should be empty when includeDetails=false")
	}
}

func TestSearchEntities_SwiftBIC(t *testing.T) {
	srv := newTestServer(t)

	minMatch := 0.80
	limit := 5
	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:  SearchEntityRequest{SwiftBICs: []string{"HNRBBY2X"}},
		Limit:    &limit,
		MinMatch: &minMatch,
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp pubsearch.SearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))
	require.NotEmpty(t, resp.Entities)
	require.Equal(t, "12028", resp.Entities[0].SourceID)
	require.InDelta(t, 1.0, resp.Entities[0].Match, 0.001)
	require.Contains(t, resp.Entities[0].SwiftBICs, "HNRBBY2X")
}

func TestSearchEntities_AlgorithmSoundex(t *testing.T) {
	srv := newTestServer(t)

	alg := "soundex"
	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:   SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
		Algorithm: &alg,
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)
}

func TestSearchEntities_AlgorithmSoftBidist(t *testing.T) {
	srv := newTestServer(t)

	alg := "soft-bidist"
	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:   SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
		Algorithm: &alg,
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)
}

func TestSearchEntities_AlgorithmInvalid(t *testing.T) {
	srv := newTestServer(t)

	alg := "levenshtein"
	_, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request:   SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
		Algorithm: &alg,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown algorithm")
}
