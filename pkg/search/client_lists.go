package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// ListRef is one stored snapshot of a downloaded list. Ref is the list file hash.
type ListRef struct {
	ListName     string    `json:"listName"`
	Ref          string    `json:"ref"`
	DownloadedAt time.Time `json:"downloadedAt"`
	EntityCount  int       `json:"entityCount"`
	AddedCount   int       `json:"addedCount"`
	RemovedCount int       `json:"removedCount"`
	ChangedCount int       `json:"changedCount"`
}

// ListsResponse is GET /v2/lists.
type ListsResponse struct {
	Lists []ListRef `json:"lists"`
}

// ListRefsResponse is GET /v2/lists/{listName}/refs.
type ListRefsResponse struct {
	ListName string    `json:"listName"`
	Refs     []ListRef `json:"refs"`
}

// ListEntityRecord is a list row at a ref, including the prepared entity.
type ListEntityRecord struct {
	SourceID   string        `json:"sourceID"`
	Name       string        `json:"name"`
	EntityType EntityType    `json:"entityType"`
	EntityHash string        `json:"entityHash,omitempty"`
	Entity     Entity[Value] `json:"entity"`
}

// ListEntitiesResponse is GET /v2/lists/{listName}/entities.
type ListEntitiesResponse struct {
	ListName string             `json:"listName"`
	Ref      ListRef            `json:"ref"`
	Entities []ListEntityRecord `json:"entities"`
	After    string             `json:"after,omitempty"`
}

// ListChange is one added, removed, or changed entity in a list diff.
type ListChange struct {
	ChangeType string         `json:"changeType"`
	SourceID   string         `json:"sourceID"`
	Name       string         `json:"name"`
	EntityType EntityType     `json:"entityType"`
	Before     *Entity[Value] `json:"before,omitempty"`
	After      *Entity[Value] `json:"after,omitempty"`
}

// ListDiffCounts is the full (unpaginated) size of a list diff.
type ListDiffCounts struct {
	Added   int `json:"added"`
	Removed int `json:"removed"`
	Changed int `json:"changed"`
}

// ListDiffResponse is GET /v2/lists/{listName}/diff.
type ListDiffResponse struct {
	ListName string         `json:"listName"`
	From     ListRef        `json:"from"`
	To       ListRef        `json:"to"`
	Counts   ListDiffCounts `json:"counts"`
	Added    []ListChange   `json:"added"`
	Removed  []ListChange   `json:"removed"`
	Changed  []ListChange   `json:"changed"`
	After    string         `json:"after,omitempty"`
}

// ListEntitiesOpts customizes GET /v2/lists/{listName}/entities.
type ListEntitiesOpts struct {
	Ref        string
	Query      string
	EntityType string
	Limit      int
	After      string
}

// ListDiffOpts customizes GET /v2/lists/{listName}/diff.
type ListDiffOpts struct {
	From   string
	To     string
	Detail string
	Limit  int
	After  string
}

func (c *client) Lists(ctx context.Context) (ListsResponse, error) {
	var out ListsResponse
	if err := c.getJSON(ctx, "/v2/lists", nil, &out); err != nil {
		return out, fmt.Errorf("lists GET: %w", err)
	}
	return out, nil
}

func (c *client) ListRefs(ctx context.Context, listName string) (ListRefsResponse, error) {
	var out ListRefsResponse
	path := "/v2/lists/" + url.PathEscape(listName) + "/refs"
	if err := c.getJSON(ctx, path, nil, &out); err != nil {
		return out, fmt.Errorf("list refs GET: %w", err)
	}
	return out, nil
}

func (c *client) ListEntities(ctx context.Context, listName string, opts ListEntitiesOpts) (ListEntitiesResponse, error) {
	var out ListEntitiesResponse
	q := make(url.Values)
	if opts.Ref != "" {
		q.Set("ref", opts.Ref)
	}
	if opts.Query != "" {
		q.Set("q", opts.Query)
	}
	if opts.EntityType != "" {
		q.Set("type", opts.EntityType)
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.After != "" {
		q.Set("after", opts.After)
	}
	path := "/v2/lists/" + url.PathEscape(listName) + "/entities"
	if err := c.getJSON(ctx, path, q, &out); err != nil {
		return out, fmt.Errorf("list entities GET: %w", err)
	}
	return out, nil
}

func (c *client) ListEntity(ctx context.Context, listName, sourceID, ref string) (ListEntityRecord, error) {
	var out ListEntityRecord
	q := make(url.Values)
	if ref != "" {
		q.Set("ref", ref)
	}
	path := "/v2/lists/" + url.PathEscape(listName) + "/entities/" + url.PathEscape(sourceID)
	if err := c.getJSON(ctx, path, q, &out); err != nil {
		return out, fmt.Errorf("list entity GET: %w", err)
	}
	return out, nil
}

func (c *client) ListDiff(ctx context.Context, listName string, opts ListDiffOpts) (ListDiffResponse, error) {
	var out ListDiffResponse
	q := make(url.Values)
	if opts.From != "" {
		q.Set("from", opts.From)
	}
	if opts.To != "" {
		q.Set("to", opts.To)
	}
	if opts.Detail != "" {
		q.Set("detail", opts.Detail)
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.After != "" {
		q.Set("after", opts.After)
	}
	path := "/v2/lists/" + url.PathEscape(listName) + "/diff"
	if err := c.getJSON(ctx, path, q, &out); err != nil {
		return out, fmt.Errorf("list diff GET: %w", err)
	}
	return out, nil
}

func (c *client) getJSON(ctx context.Context, path string, query url.Values, dest any) error {
	addr := c.baseAddress + path
	if len(query) > 0 {
		addr += "?" + query.Encode()
	}

	req, err := retryablehttp.NewRequest("GET", addr, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return withoutRequestURL(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var er struct {
			Error string `json:"error"`
		}
		_ = json.Unmarshal(body, &er)
		if er.Error != "" {
			return fmt.Errorf("status %d: %s", resp.StatusCode, er.Error)
		}
		return fmt.Errorf("status %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	return nil
}
