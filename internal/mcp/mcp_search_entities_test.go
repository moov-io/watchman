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
