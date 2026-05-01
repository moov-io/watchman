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
	"github.com/razashariff/agentpass-go"
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

func contextWithAgent(ctx context.Context, agentID string, trustLevel int, issuerID string) context.Context {
	v := &agentpass.Verified{
		Agent: agentpass.Agent{
			AgentID:    agentID,
			TrustLevel: trustLevel,
			IssuerID:   issuerID,
		},
	}
	return context.WithValue(ctx, agentContextKey, v)
}

func TestSearchEntities_SearchedAtAlwaysPresent(t *testing.T) {
	srv := newTestServer(t)

	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request: SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	var resp mcpSearchResponse
	text := result.Content[0].(*mcpsdk.TextContent).Text
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.Greater(t, resp.SearchedAt, int64(0), "searchedAt should be a non-zero unix timestamp")
}

func TestSearchEntities_NoAgentContext(t *testing.T) {
	srv := newTestServer(t)

	result, _, err := srv.HandleSearchEntities(context.Background(), &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request: SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
	})
	require.NoError(t, err)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp mcpSearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.Nil(t, resp.AgentContext, "agentContext should be absent when no agent is authenticated")
}

func TestSearchEntities_AgentContextInResponse(t *testing.T) {
	srv := newTestServer(t)

	ctx := contextWithAgent(context.Background(), "acme-agent-001", 2, "acme-corp")

	result, _, err := srv.HandleSearchEntities(ctx, &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request: SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
	})
	require.NoError(t, err)

	text := result.Content[0].(*mcpsdk.TextContent).Text
	var resp mcpSearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.NotNil(t, resp.AgentContext, "agentContext must be present when agent is authenticated")
	require.Equal(t, "acme-agent-001", resp.AgentContext.AgentID)
	require.Equal(t, 2, resp.AgentContext.TrustLevel)
	require.Equal(t, "acme-corp", resp.AgentContext.IssuerID)
}

func TestSearchEntities_AgentContextInSignedResponse(t *testing.T) {
	logger := log.NewTestLogger()
	indexedLists := index.NewLists(nil)
	searchConfig := search.DefaultConfig()
	svc, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	tmpDir := t.TempDir()
	srv, err := NewServer(logger, svc, config.MCPConfig{
		Signing: config.MCPSigning{
			Enabled: true,
			KeyPath: tmpDir + "/test.key",
			PubPath: tmpDir + "/test.pub",
		},
	})
	require.NoError(t, err)

	ctx := contextWithAgent(context.Background(), "signing-agent-007", 3, "partner-ca")

	result, _, err := srv.HandleSearchEntities(ctx, &mcpsdk.CallToolRequest{}, SearchEntitiesRequest{
		Request: SearchEntityRequest{Name: "Logan", Type: pubsearch.EntityPerson},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	// MCPS-signed envelope wraps the payload — unwrap and verify agent context is inside.
	// mcps-go SignedMessage has: mcps_version, passport_id, nonce, timestamp, signature, message.
	text := result.Content[0].(*mcpsdk.TextContent).Text
	var envelope map[string]json.RawMessage
	require.NoError(t, json.Unmarshal([]byte(text), &envelope))

	_, hasSignature := envelope["signature"]
	require.True(t, hasSignature, "response should be a signed envelope")

	// The signed payload is in the "message" field.
	message, ok := envelope["message"]
	require.True(t, ok, "signed envelope must contain message field")

	var inner mcpSearchResponse
	require.NoError(t, json.Unmarshal(message, &inner))
	require.NotNil(t, inner.AgentContext)
	require.Equal(t, "signing-agent-007", inner.AgentContext.AgentID)
	require.Equal(t, 3, inner.AgentContext.TrustLevel)
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
	var resp mcpSearchResponse
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
	var resp mcpSearchResponse
	require.NoError(t, json.Unmarshal([]byte(text), &resp))

	require.NotEmpty(t, resp.Entities)
	for _, entity := range resp.Entities {
		require.Empty(t, entity.Details.Pieces,
			"Details.Pieces should be empty when includeDetails=false")
	}
}
