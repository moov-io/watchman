package mcp

import (
	"context"
	"encoding/json"
	"fmt"

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
		Limit:    25,
		MinMatch: 0.0,
	}

	if args.Limit != nil {
		opts.Limit = *args.Limit
	}
	if args.MinMatch != nil {
		opts.MinMatch = *args.MinMatch
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

	response := pubsearch.SearchResponse{
		Query:    searchReq,
		Entities: entities,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	// Sign the response if MCPS signing is enabled
	if s.signing && s.keyPair != nil {
		passport := &mcps.Passport{
			ID:       "watchman-mcp",
			Subject:  "watchman",
			Version:  mcps.Version,
			IssuedAt: 0, // stateless -- no expiry tracking
			Issuer:   "moov-io/watchman",
		}

		signed, signErr := mcps.SignMessage(responseJSON, s.keyPair, passport)
		if signErr != nil {
			// Log but don't fail the request -- signing is non-blocking
			s.logger.Error().LogErrorf("MCPS: failed to sign response: %v", signErr)
		} else {
			signedJSON, _ := json.Marshal(signed)
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
