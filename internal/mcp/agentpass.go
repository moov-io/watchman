// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mcp

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/config"
	"github.com/razashariff/agentpass-go"
)

// agentContextKey is the context key for storing verified agent
// info. It is an unexported type to avoid collisions with other
// packages that may use context values.
type agentContextKeyType struct{}

var agentContextKey = agentContextKeyType{}

// agentPassGate holds the runtime state for the AgentPass
// verification middleware. It is initialised once at server
// startup and reused for every request.
type agentPassGate struct {
	logger         log.Logger
	pool           *agentpass.CertPool
	minTrust       int
	requiredScopes []string
}

// initAgentPass loads the trust anchors and prepares the AgentPass
// gate. It returns nil if AgentPass is not enabled so that the
// caller can skip middleware wrapping without additional checks.
func initAgentPass(logger log.Logger, conf config.MCPAgentPass) (*agentPassGate, error) {
	if !conf.Enabled {
		return nil, nil
	}

	if conf.TrustAnchorPath == "" {
		return nil, logger.Error().LogErrorf("agentpass: trust_anchor_path is required when agentpass is enabled").Err()
	}

	pool := agentpass.NewCertPool()
	if err := pool.AddFile(conf.TrustAnchorPath); err != nil {
		return nil, logger.Error().LogErrorf("agentpass: failed to load trust anchors from %s: %v", conf.TrustAnchorPath, err).Err()
	}

	logger.Info().Logf("agentpass: loaded %d trust anchor(s) from %s", pool.Size(), conf.TrustAnchorPath)
	logger.Info().Logf("agentpass: min_trust_level=L%d required_scopes=%v", conf.MinTrustLevel, conf.RequiredScopes)

	return &agentPassGate{
		logger:         logger,
		pool:           pool,
		minTrust:       conf.MinTrustLevel,
		requiredScopes: conf.RequiredScopes,
	}, nil
}

// headerName is the HTTP header that agents use to present their
// certificate. The value must be base64-encoded PEM.
const headerName = "X-AgentPass-Certificate"

// middleware returns an http.Handler that verifies the agent
// certificate before passing the request to the next handler.
// Requests without a valid certificate receive a 401 response
// and the entity screen never runs.
func (g *agentPassGate) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get(headerName)
		if raw == "" {
			g.logger.Warn().Logf("agentpass: rejected request with no %s header from %s", headerName, r.RemoteAddr)
			http.Error(w, "agent certificate required", http.StatusUnauthorized)
			return
		}

		pemBytes, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			g.logger.Warn().Logf("agentpass: rejected request with invalid base64 from %s", r.RemoteAddr)
			http.Error(w, "invalid certificate encoding", http.StatusUnauthorized)
			return
		}

		opts := []agentpass.VerifyOption{
			agentpass.WithMinTrust(g.minTrust),
		}
		if len(g.requiredScopes) > 0 {
			opts = append(opts, agentpass.WithRequiredScopes(g.requiredScopes...))
		}

		verified, err := agentpass.Verify(pemBytes, g.pool, opts...)
		if err != nil {
			g.logger.Warn().Logf("agentpass: rejected agent from %s: %v", r.RemoteAddr, err)
			http.Error(w, "agent verification failed", http.StatusUnauthorized)
			return
		}

		g.logger.Info().Logf("agentpass: verified agent=%s trust=L%d issuer=%s serial=%s",
			verified.AgentID, verified.TrustLevel, verified.IssuerID, verified.Serial)

		ctx := context.WithValue(r.Context(), agentContextKey, verified)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// agentFromContext retrieves the verified agent from the request
// context. Returns nil if AgentPass is not enabled or the request
// was not agent-authenticated (e.g. feature flag is off).
func agentFromContext(ctx context.Context) *agentpass.Verified {
	v, _ := ctx.Value(agentContextKey).(*agentpass.Verified)
	return v
}
