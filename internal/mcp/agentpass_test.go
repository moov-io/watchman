// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mcp

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/razashariff/agentpass-go/testca"
	"github.com/stretchr/testify/require"
)

// echoHandler returns 200 if the middleware allows the request
// through. Captures agent context from the request for assertions.
var echoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	agent := agentFromContext(r.Context())
	if agent != nil {
		w.Header().Set("X-Agent-ID", agent.AgentID)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
})

// buildGate creates an agentPassGate backed by the test bundle's
// CA. Writes the CA PEM to a temp file so initAgentPass can load
// it via the normal file path.
func buildGate(t *testing.T, bundle testca.Bundle, minTrust int, scopes []string) *agentPassGate {
	t.Helper()
	caPath := t.TempDir() + "/ca.pem"
	require.NoError(t, os.WriteFile(caPath, bundle.CAPEM, 0o644))

	gate, err := initAgentPass(log.NewTestLogger(), config.MCPAgentPass{
		Enabled:         true,
		TrustAnchorPath: caPath,
		MinTrustLevel:   minTrust,
		RequiredScopes:  scopes,
	})
	require.NoError(t, err)
	require.NotNil(t, gate)
	return gate
}

// certHeader returns the base64-encoded agent PEM suitable for the
// X-AgentPass-Certificate HTTP header.
func certHeader(pem []byte) string {
	return base64.StdEncoding.EncodeToString(pem)
}

func TestAgentPassMiddleware_Happy(t *testing.T) {
	bundle := testca.Build(testca.AgentOptions{TrustLevel: 2, Scopes: []string{"sanctions:search"}})
	gate := buildGate(t, bundle, 0, nil)

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, certHeader(bundle.AgentPEM))
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "agent-001", w.Header().Get("X-Agent-ID"))
}

func TestAgentPassMiddleware_NoCert_Rejects(t *testing.T) {
	bundle := testca.Build(testca.AgentOptions{TrustLevel: 2})
	gate := buildGate(t, bundle, 0, nil)

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_BadBase64_Rejects(t *testing.T) {
	bundle := testca.Build(testca.AgentOptions{TrustLevel: 2})
	gate := buildGate(t, bundle, 0, nil)

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, "not-valid-base64!!!")
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_ExpiredCert_Rejects(t *testing.T) {
	now := time.Now()
	bundle := testca.Build(testca.AgentOptions{
		TrustLevel: 2,
		NotBefore:  now.Add(-48 * time.Hour),
		NotAfter:   now.Add(-1 * time.Hour),
	})
	gate := buildGate(t, bundle, 0, nil)

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, certHeader(bundle.AgentPEM))
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_UntrustedCA_Rejects(t *testing.T) {
	bundleA := testca.Build(testca.AgentOptions{TrustLevel: 2})
	bundleB := testca.Build(testca.AgentOptions{TrustLevel: 2})
	gate := buildGate(t, bundleA, 0, nil) // trust only CA-A

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, certHeader(bundleB.AgentPEM)) // cert signed by CA-B
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_TrustLevelTooLow_Rejects(t *testing.T) {
	bundle := testca.Build(testca.AgentOptions{TrustLevel: 1})
	gate := buildGate(t, bundle, 2, nil) // require minimum L2

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, certHeader(bundle.AgentPEM))
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_MissingScope_Rejects(t *testing.T) {
	bundle := testca.Build(testca.AgentOptions{TrustLevel: 2, Scopes: []string{"payments"}})
	gate := buildGate(t, bundle, 0, []string{"sanctions:search"})

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader("{}"))
	req.Header.Set(headerName, certHeader(bundle.AgentPEM))
	w := httptest.NewRecorder()

	gate.middleware(echoHandler).ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAgentPassMiddleware_Disabled_ReturnsNil(t *testing.T) {
	gate, err := initAgentPass(log.NewTestLogger(), config.MCPAgentPass{
		Enabled: false,
	})
	require.NoError(t, err)
	require.Nil(t, gate)
}

func TestAgentPassMiddleware_NoTrustAnchorPath_Errors(t *testing.T) {
	_, err := initAgentPass(log.NewTestLogger(), config.MCPAgentPass{
		Enabled:         true,
		TrustAnchorPath: "", // missing
	})
	require.Error(t, err)
}

// TestAcmeCorpWatchmanAgent simulates a real-world deployment where
// "Acme Corp" operates an AI agent that screens entities against
// sanctions lists via Watchman's MCP endpoint.
//
// Run with: go test -v -run TestAcmeCorpWatchmanAgent ./internal/mcp/
//
// The test demonstrates:
//   1. Acme Corp's trusted agent (L3, sanctions:search scope) is accepted
//   2. A rogue agent signed by an unknown CA is rejected
//   3. An under-privileged agent (L1, no scopes) is rejected
func TestAcmeCorpWatchmanAgent(t *testing.T) {
	// --- Acme Corp sets up their internal CA ---
	acmeCA := testca.Build(testca.AgentOptions{
		TrustLevel: 3,
		Scopes:     []string{"sanctions:search", "payments:screen"},
	})

	// --- Watchman operator configures AgentPass ---
	gate := buildGate(t, acmeCA, 2, []string{"sanctions:search"})

	// --- Scenario 1: Acme Corp's trusted agent screens an entity ---
	t.Run("acme_agent_screens_entity", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{
			"jsonrpc": "2.0",
			"method": "tools/call",
			"params": {"name": "search_entities", "arguments": {"name": "John Doe"}}
		}`))
		req.Header.Set(headerName, certHeader(acmeCA.AgentPEM))
		w := httptest.NewRecorder()

		gate.middleware(echoHandler).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code, "trusted acme-corp agent should pass")
		require.Equal(t, "agent-001", w.Header().Get("X-Agent-ID"))
	})

	// --- Scenario 2: Rogue agent from unknown CA tries to screen ---
	t.Run("rogue_agent_blocked", func(t *testing.T) {
		rogueCA := testca.Build(testca.AgentOptions{
			TrustLevel: 4,
			Scopes:     []string{"sanctions:search"},
		})

		req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{}`))
		req.Header.Set(headerName, certHeader(rogueCA.AgentPEM))
		w := httptest.NewRecorder()

		gate.middleware(echoHandler).ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code, "rogue agent must be rejected")
	})

	// --- Scenario 3: Under-privileged agent (low trust, no scopes) ---
	t.Run("underprivileged_agent_blocked", func(t *testing.T) {
		weakAgent := testca.Build(testca.AgentOptions{
			TrustLevel: 1,
			Scopes:     []string{"read-only"},
		})
		weakGate := buildGate(t, weakAgent, 2, []string{"sanctions:search"})

		req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{}`))
		req.Header.Set(headerName, certHeader(weakAgent.AgentPEM))
		w := httptest.NewRecorder()

		weakGate.middleware(echoHandler).ServeHTTP(w, req)

		require.Equal(t, http.StatusUnauthorized, w.Code, "L1 agent without sanctions:search scope must be rejected")
	})
}
