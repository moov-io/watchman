package search

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client interface {
	ListInfo(ctx context.Context) (ListInfoResponse, error)
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

type ListInfoResponse struct {
	Lists      map[string]int    `json:"lists"`
	ListHashes map[string]string `json:"listHashes"`

	StartedAt time.Time `json:"startedAt"`
	EndedAt   time.Time `json:"endedAt"`
}

func (c *client) ListInfo(ctx context.Context) (ListInfoResponse, error) {
	addr := c.baseAddress + "/v2/listinfo"

	var out ListInfoResponse
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return out, fmt.Errorf("creating listinfo request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return out, fmt.Errorf("listinfo GET: %w", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return out, fmt.Errorf("decoding listinfo response: %w", err)
	}
	return out, nil
}

type SearchResponse struct {
	Entities []SearchedEntity[Value] `json:"entities"`
}

type SearchOpts struct {
	Limit    int
	MinMatch float64
}

func (c *client) SearchByEntity(ctx context.Context, entity Entity[Value], opts SearchOpts) (SearchResponse, error) {
	var out SearchResponse

	// Build the URL
	addr, err := url.Parse(c.baseAddress + "/v2/search")
	if err != nil {
		return out, fmt.Errorf("problem creating baseAddress: %w", err)
	}

	// Set query parameters
	q := addr.Query()
	q.Set("type", string(entity.Type))

	if entity.Name != "" {
		q.Set("name", entity.Name)
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	addr.RawQuery = q.Encode()

	// Make the request
	req, err := http.NewRequest("GET", addr.String(), nil)
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
