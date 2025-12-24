package search

import (
	"cmp"
	"context"
	"io"
	"slices"
	"sync"
)

type MockClient struct {
	Err error

	ListInfoResponse   ListInfoResponse
	IngestFileResponse IngestFileResponse
	ExportFileResponse []Entity[Value]

	ListInfoErr   error
	SearchErr     error
	IngestFileErr error
	ExportFileErr error

	mu       sync.RWMutex
	Index    []Entity[Value]
	Searches []Entity[Value]
}

var _ Client = (&MockClient{})

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) ListInfo(ctx context.Context) (ListInfoResponse, error) {
	err := cmp.Or(c.ListInfoErr, c.Err)
	if err != nil {
		var out ListInfoResponse
		return out, err
	}
	return c.ListInfoResponse, nil
}

func (c *MockClient) SearchByEntity(ctx context.Context, query Entity[Value], opts SearchOpts) (SearchResponse, error) {
	err := cmp.Or(c.SearchErr, c.Err)
	if err != nil {
		var out SearchResponse
		return out, err
	}

	// Make sure to prepare the Query
	query = query.Normalize()

	// Record the search
	c.mu.Lock()
	c.Searches = append(c.Searches, query)
	c.mu.Unlock()

	// Grab read lock
	c.mu.RLock()
	defer c.mu.RUnlock()

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

func (c *MockClient) IngestFile(ctx context.Context, fileType string, file io.Reader) (IngestFileResponse, error) {
	err := cmp.Or(c.IngestFileErr, c.Err)
	if err != nil {
		var out IngestFileResponse
		return out, err
	}
	return c.IngestFileResponse, nil
}

func (c *MockClient) ExportFile(ctx context.Context, fileType string) ([]Entity[Value], error) {
	err := cmp.Or(c.ExportFileErr, c.Err)
	if err != nil {
		return nil, err
	}
	return c.ExportFileResponse, nil
}

func (c *MockClient) Normalize() {
	for idx := range c.Index {
		c.Index[idx] = c.Index[idx].Normalize()
	}
}
