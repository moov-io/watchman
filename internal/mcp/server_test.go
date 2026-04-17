package mcp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestMCPHandler(t *testing.T) {
	logger := log.NewTestLogger()

	// Set up real search service with test data
	indexedLists := index.NewLists(nil) // only in-mem
	searchConfig := search.DefaultConfig()
	service, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	server, err := NewServer(logger, service, config.MCPSigning{}, config.MCPAgentPass{})
	require.NoError(t, err)

	handler := server.Handler()

	t.Run("handler returns http.Handler", func(t *testing.T) {
		require.NotNil(t, handler)
		require.Implements(t, (*http.Handler)(nil), handler)
	})
}

func TestSearchEntitiesTool(t *testing.T) {
	logger := log.NewTestLogger()

	// Set up real search service with test data
	indexedLists := index.NewLists(nil)
	searchConfig := search.DefaultConfig()
	service, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	server, err := NewServer(logger, service, config.MCPSigning{}, config.MCPAgentPass{})
	require.NoError(t, err)

	// Test the handleSearchEntities function directly
	req := &mcp.CallToolRequest{}
	args := SearchEntitiesRequest{
		Request: SearchEntityRequest{
			Name: "Logan",
			Type: pubsearch.EntityPerson,
		},
		Limit:    &[]int{10}[0],
		MinMatch: &[]float64{0.8}[0],
	}

	result, extra, err := server.HandleSearchEntities(context.Background(), req, args)
	require.NoError(t, err)
	require.Nil(t, extra)
	require.NotNil(t, result)
	require.Len(t, result.Content, 1)

	// Parse the JSON response
	content := result.Content[0].(*mcp.TextContent)
	var response pubsearch.SearchResponse
	err = json.Unmarshal([]byte(content.Text), &response)
	require.NoError(t, err)
	// Note: Test data may not contain matching entities, so we just verify the response structure
	require.NotNil(t, response.Query)
}

// TestSearchEntitiesRequest tests the JSON marshaling/unmarshaling
func TestSearchEntitiesRequest(t *testing.T) {
	req := SearchEntitiesRequest{
		Request: SearchEntityRequest{
			Name: "Test Person",
			Type: pubsearch.EntityPerson,
			Person: &pubsearch.Person{
				Name: "Test Person",
			},
		},
		Limit:    &[]int{5}[0],
		MinMatch: &[]float64{0.9}[0],
	}

	// Test marshaling
	data, err := json.Marshal(req)
	require.NoError(t, err)
	require.Contains(t, string(data), "Test Person")
	require.Contains(t, string(data), "person")

	// Test unmarshaling
	var unmarshaled SearchEntitiesRequest
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)
	require.Equal(t, "Test Person", unmarshaled.Request.Name)
	require.Equal(t, pubsearch.EntityPerson, unmarshaled.Request.Type)
	require.Equal(t, 5, *unmarshaled.Limit)
	require.Equal(t, 0.9, *unmarshaled.MinMatch)
}

// TestMCPSigningEnabled verifies that the MCP server initialises correctly
// when MCPS message signing is enabled via config (key paths in a temp dir),
// and that search_entities responses are emitted in the signed-envelope shape.
func TestMCPSigningEnabled(t *testing.T) {
	logger := log.NewTestLogger()

	indexedLists := index.NewLists(nil)
	searchConfig := search.DefaultConfig()
	service, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	tmpDir := t.TempDir()
	signingConf := config.MCPSigning{
		Enabled: true,
		KeyPath: filepath.Join(tmpDir, "test-mcps.key"),
		PubPath: filepath.Join(tmpDir, "test-mcps.pub"),
	}

	server, err := NewServer(logger, service, signingConf, config.MCPAgentPass{})
	require.NoError(t, err)
	require.True(t, server.signing, "signing should be enabled after successful key init")
	require.NotNil(t, server.keyPair, "keyPair should be populated when signing is enabled")

	// Run a search and confirm the response is emitted as a signed envelope
	// rather than a raw pubsearch.SearchResponse.
	req := &mcp.CallToolRequest{}
	args := SearchEntitiesRequest{
		Request: SearchEntityRequest{
			Name: "Logan",
			Type: pubsearch.EntityPerson,
		},
		Limit:    &[]int{5}[0],
		MinMatch: &[]float64{0.8}[0],
	}

	result, extra, err := server.HandleSearchEntities(context.Background(), req, args)
	require.NoError(t, err)
	require.Nil(t, extra)
	require.NotNil(t, result)
	require.Len(t, result.Content, 1)

	content := result.Content[0].(*mcp.TextContent)
	var envelope map[string]interface{}
	err = json.Unmarshal([]byte(content.Text), &envelope)
	require.NoError(t, err)

	// A signed MCPS envelope carries signature + passport fields alongside the payload.
	// We assert presence without binding to a specific signature implementation so the
	// test stays stable across mcps-go versions.
	_, hasSignature := envelope["signature"]
	_, hasPassport := envelope["passport"]
	require.True(t, hasSignature || hasPassport,
		"signed response should include signature or passport fields; got keys: %v", keysOf(envelope))
}

func keysOf(m map[string]interface{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func TestMCPHTTPIntegration(t *testing.T) {
	logger := log.NewTestLogger()

	// Set up real search service with test data
	indexedLists := index.NewLists(nil)
	searchConfig := search.DefaultConfig()
	service, err := search.NewService(logger, searchConfig, nil, indexedLists)
	require.NoError(t, err)

	dl := ofactest.GetDownloader(t)
	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	indexedLists.Update(stats)

	server, err := NewServer(logger, service, config.MCPSigning{}, config.MCPAgentPass{})
	require.NoError(t, err)

	handler := server.Handler()

	// Create a test HTTP server
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	t.Run("MCP tools/call search_entities", func(t *testing.T) {
		// Call the search_entities tool directly (stateless mode)
		searchReq := `{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "tools/call",
			"params": {
				"name": "search_entities",
				"arguments": {
					"request": {
						"name": "Logan",
						"entityType": "person",
						"sourceList": "api-request",
						"sourceID": "test",
						"person": null,
						"business": null,
						"organization": null,
						"aircraft": null,
						"vessel": null,
						"contact": {
							"emailAddresses": [],
							"phoneNumbers": [],
							"faxNumbers": [],
							"websites": []
						},
						"addresses": [],
						"cryptoAddresses": [],
						"affiliations": [],
						"sanctionsInfo": null,
						"historicalInfo": [],
						"sourceData": null
					},
					"limit": 5,
					"minMatch": 0.8
				}
			}
		}`

		req, err := http.NewRequest("POST", testServer.URL, strings.NewReader(searchReq))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json,text/event-stream")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), `"result"`)

		// Parse the response
		var jsonResp map[string]interface{}
		err = json.Unmarshal(body, &jsonResp)
		require.NoError(t, err)

		result, ok := jsonResp["result"].(map[string]interface{})
		require.True(t, ok, "Response should contain result")

		content, ok := result["content"].([]interface{})
		require.True(t, ok, "Result should contain content")
		require.Len(t, content, 1)

		contentItem, ok := content[0].(map[string]interface{})
		require.True(t, ok)

		text, ok := contentItem["text"].(string)
		require.True(t, ok)

		// Parse the search response
		var searchResp pubsearch.SearchResponse
		err = json.Unmarshal([]byte(text), &searchResp)
		require.NoError(t, err)
		// Note: Test data may not contain matching entities, so we just verify the response structure
		require.NotNil(t, searchResp.Query)
		require.NotNil(t, searchResp.Entities) // Should be an empty slice, not nil
	})
}
