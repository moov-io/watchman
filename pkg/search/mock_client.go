package search

import (
	"cmp"
	"context"
	"slices"
)

type MockClient struct {
	Err error

	Index    []Entity[Value]
	Searches []Entity[Value]
}

var _ Client = (&MockClient{})

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) SearchByEntity(ctx context.Context, query Entity[Value], opts SearchOpts) (SearchResponse, error) {
	if c.Err != nil {
		return SearchResponse{}, c.Err
	}

	// Record the search
	c.Searches = append(c.Searches, query)

	var resp SearchResponse
	for _, index := range c.Index {
		resp.Entities = append(resp.Entities, SearchedEntity[Value]{
			Entity: index,
			Match:  Similarity(query, index),
		})
	}

	// Sort the results, highest match first
	slices.SortFunc(resp.Entities, func(e1, e2 SearchedEntity[Value]) int {
		return -1 * cmp.Compare(e1.Match, e2.Match) // invert, make it DESC
	})

	// Truncate
	if len(resp.Entities) > opts.Limit {
		resp.Entities = resp.Entities[:opts.Limit]
	}

	return resp, nil
}
