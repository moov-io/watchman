package search

import (
	"context"
)

type MockClient struct {
	Err            error
	SearchResponse SearchResponse
}

var _ Client = (&MockClient{})

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) SearchByEntity(ctx context.Context, entity Entity[Value], opts SearchOpts) (SearchResponse, error) {
	if c.Err != nil {
		return SearchResponse{}, c.Err
	}
	return c.SearchResponse, c.Err
}
