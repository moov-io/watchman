package search

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	SearchByEntity(ctx context.Context, entity Entity[Value], opts SearchOpts) (SearchResponse, error)
}

func NewClient(httpClient *http.Client, baseAddress string) Client {
	httpClient = cmp.Or(httpClient, &http.Client{
		Timeout: 5 * time.Second,
	})

	return &client{
		httpClient:  httpClient,
		baseAddress: baseAddress,
	}
}

type client struct {
	httpClient  *http.Client
	baseAddress string
}

type SearchResponse struct {
	Entities []SearchedEntity[Value] `json:"entities"`
}

type SearchOpts struct {
	Limit    int
	MinMatch float64
}

func (c *client) SearchByEntity(ctx context.Context, entity Entity[Value], opts SearchOpts) (SearchResponse, error) {
	addr := c.baseAddress + "/v2/search"
	addr += "?name=" + entity.Name // TODO(adam): escape, use proper setters

	if opts.Limit > 0 {
		addr += fmt.Sprintf("&limit=%d", opts.Limit)
	}

	var out SearchResponse

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return out, fmt.Errorf("creating search request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return out, fmt.Errorf("search by entity: %w", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return out, fmt.Errorf("decoding search by entity response: %w", err)
	}
	return out, nil
}
