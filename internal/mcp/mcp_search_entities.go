// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/moov-io/watchman/internal/search"
	pubsearch "github.com/moov-io/watchman/pkg/search"
	mcps "github.com/razashariff/mcps-go"
)

type SearchEntityRequest struct {
	Name   string               `json:"name,omitempty" jsonschema:"Entity name"`
	Type   pubsearch.EntityType `json:"entityType,omitempty" jsonschema:"Entity type (person, business, organization, aircraft, vessel)"`
	Source pubsearch.SourceList `json:"sourceList,omitempty" jsonschema:"Source list"`

	SourceID string `json:"sourceID,omitempty" jsonschema:"Source data identifier"`

	Person       *pubsearch.Person       `json:"person,omitempty" jsonschema:"Person details"`
	Business     *pubsearch.Business     `json:"business,omitempty" jsonschema:"Business details"`
	Organization *pubsearch.Organization `json:"organization,omitempty" jsonschema:"Organization details"`
	Aircraft     *pubsearch.Aircraft     `json:"aircraft,omitempty" jsonschema:"Aircraft details"`
	Vessel       *pubsearch.Vessel       `json:"vessel,omitempty" jsonschema:"Vessel details"`

	Contact         pubsearch.ContactInfo     `json:"contact,omitempty" jsonschema:"Contact information"`
	Addresses       []pubsearch.Address       `json:"addresses,omitempty" jsonschema:"Addresses"`
	CryptoAddresses []pubsearch.CryptoAddress `json:"cryptoAddresses,omitempty" jsonschema:"Crypto addresses"`

	Affiliations   []pubsearch.Affiliation    `json:"affiliations,omitempty" jsonschema:"Affiliations"`
	SanctionsInfo  *pubsearch.SanctionsInfo   `json:"sanctionsInfo,omitempty" jsonschema:"Sanctions information"`
	HistoricalInfo []pubsearch.HistoricalInfo `json:"historicalInfo,omitempty" jsonschema:"Historical information"`

	SourceData pubsearch.Value `json:"sourceData,omitempty" jsonschema:"Original source data"`
}

type SearchEntitiesRequest struct {
	Request SearchEntityRequest `json:"request" jsonschema:"Entity search request object"`

	Limit    *int     `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default 10)"`
	MinMatch *float64 `json:"minMatch,omitempty" jsonschema:"Minimum match score threshold (default 0.0)"`

	IncludeDetails *bool `json:"includeDetails,omitempty" jsonschema:"Include field-level match score breakdown per result (names, addresses, IDs, etc.)"`
}

// AgentContext carries the verified identity of the AI agent that made this search request.
// Only present when the MCP server has AgentPass authentication enabled and the request
// carried a valid X-AgentPass-Certificate header.
type AgentContext struct {
	AgentID    string `json:"agentId"`
	TrustLevel int    `json:"trustLevel"`
	IssuerID   string `json:"issuerId"`
}

// mcpSearchResponse is the envelope returned by the search_entities MCP tool.
// Uses explicit fields instead of embedding pubsearch.SearchResponse to avoid inheriting
// a custom UnmarshalJSON that would swallow SearchedAt and AgentContext during tests.
type mcpSearchResponse struct {
	Query        pubsearch.Entity[pubsearch.Value]           `json:"query"`
	Entities     []pubsearch.SearchedEntity[pubsearch.Value] `json:"entities"`
	AgentContext *AgentContext                               `json:"agentContext,omitempty"`
	SearchedAt   int64                                       `json:"searchedAt"`
}

func (s *Server) HandleSearchEntities(ctx context.Context, req *mcp.CallToolRequest, args SearchEntitiesRequest) (*mcp.CallToolResult, any, error) {
	searchReq := pubsearch.Entity[pubsearch.Value]{
		Name:            args.Request.Name,
		Type:            args.Request.Type,
		Source:          args.Request.Source,
		SourceID:        args.Request.SourceID,
		Person:          args.Request.Person,
		Business:        args.Request.Business,
		Organization:    args.Request.Organization,
		Aircraft:        args.Request.Aircraft,
		Vessel:          args.Request.Vessel,
		Contact:         args.Request.Contact,
		Addresses:       args.Request.Addresses,
		CryptoAddresses: args.Request.CryptoAddresses,
		Affiliations:    args.Request.Affiliations,
		SanctionsInfo:   args.Request.SanctionsInfo,
		HistoricalInfo:  args.Request.HistoricalInfo,
		SourceData:      args.Request.SourceData,
	}

	// Set default entity details if not provided
	if args.Request.Type == pubsearch.EntityPerson && searchReq.Person == nil {
		searchReq.Person = &pubsearch.Person{Name: args.Request.Name}
	}

	opts := search.SearchOpts{
		Limit:    10,
		MinMatch: 0.0,
	}

	if args.Limit != nil {
		opts.Limit = *args.Limit
	}
	if args.MinMatch != nil {
		opts.MinMatch = *args.MinMatch
	}
	if args.IncludeDetails != nil && *args.IncludeDetails {
		opts.Debug = true
	}

	// Log agent identity so compliance audit trails can tie agent to search activity.
	agent := agentFromContext(ctx)
	if agent != nil {
		s.logger.Info().Logf("mcp: search by agent=%s trust=L%d issuer=%s name=%q type=%s",
			agent.AgentID, agent.TrustLevel, agent.IssuerID, args.Request.Name, args.Request.Type)
	}

	// Normalize the request
	searchReq = searchReq.Normalize()

	// Perform the search
	entities, err := s.service.Search(ctx, searchReq, opts)
	if err != nil {
		return nil, nil, fmt.Errorf("search failed: %w", err)
	}

	// Ensure entities is not nil for JSON marshaling
	if entities == nil {
		entities = []pubsearch.SearchedEntity[pubsearch.Value]{}
	}

	resp := mcpSearchResponse{
		Query:      searchReq,
		Entities:   entities,
		SearchedAt: time.Now().Unix(),
	}

	// Stamp verified agent identity into the response so that MCPS-signed
	// envelopes carry cryptographic proof of who requested this search.
	if agent != nil {
		resp.AgentContext = &AgentContext{
			AgentID:    agent.AgentID,
			TrustLevel: agent.TrustLevel,
			IssuerID:   agent.IssuerID,
		}
	}

	responseJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	// Sign the response if MCPS signing is enabled
	if s.signing && s.keyPair != nil {
		passport := &mcps.Passport{
			ID:       "watchman-mcp",
			Subject:  "watchman",
			Version:  mcps.Version,
			IssuedAt: time.Now().Unix(),
			Issuer:   "moov-io/watchman",
		}

		signed, signErr := mcps.SignMessage(responseJSON, s.keyPair, passport)
		if signErr != nil {
			// Log but don't fail the request -- signing is non-blocking
			s.logger.Error().LogErrorf("MCPS: failed to sign response: %v", signErr)
		} else {
			signedJSON, err := json.Marshal(signed)
			if err != nil {
				err = s.logger.Error().LogErrorf("MCPS: failed to marshal signed response: %v", err).Err()

				return &mcp.CallToolResult{
					Content: []mcp.Content{
						&mcp.TextContent{Text: err.Error()},
					},
					IsError: true,
				}, nil, nil
			}
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: string(signedJSON)},
				},
			}, nil, nil
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(responseJSON)},
		},
	}, nil, nil
}
