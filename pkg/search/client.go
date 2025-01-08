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
	SearchByEntity(ctx context.Context, entity Entity[Value]) ([]SearchedEntity[Value], error)
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

func (c *client) SearchByEntity(ctx context.Context, entity Entity[Value]) ([]SearchedEntity[Value], error) {
	addr := c.baseAddress + "/v2/search"
	// addr +=

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, fmt.Errorf("creating search request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search by entity: %w", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var out []SearchedEntity[Value]
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil, fmt.Errorf("decoding search by entity response: %w", err)
	}
	return out, nil
}
