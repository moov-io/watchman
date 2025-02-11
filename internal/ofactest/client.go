package ofactest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is an HTTP client which makes calls to the OFAC Sanctions List Search application ("Sanctions List Search")
// hosted at https://sanctionslist.ofac.treas.gov/Home/index.html
//
// This client is mostly used for comparison of Watchman against the official OFAC tool.
type Client interface {
	Search(ctx context.Context, params SearchParams) ([]SearchResult, error)
}

func NewClient() Client {
	return &client{
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

type client struct {
	httpClient *http.Client
}

const (
	searchBaseAddress = "https://sanctionslistservice.ofac.treas.gov/api/Search/Search"
)

type SearchParams struct {
	Name          string   `json:"name"`
	City          string   `json:"city"`
	IDNumber      string   `json:"idNumber"`
	StateProvince string   `json:"stateProvince"`
	NameScore     int      `json:"nameScore"`
	Country       string   `json:"country"`
	Programs      []string `json:"programs"`
	Type          string   `json:"type"`
	Address       string   `json:"address"`
	List          string   `json:"list"`
}

type SearchResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	Programs  string `json:"programs"`
	Lists     string `json:"lists"`
	NameScore int    `json:"nameScore"`
}

type ErrorResponse struct {
	Name          string        `json:"name"`
	City          string        `json:"city"`
	IDNumber      string        `json:"idNumber"`
	StateProvince string        `json:"stateProvince"`
	NameScore     int           `json:"nameScore"`
	Success       bool          `json:"success"`
	ErrorMessage  string        `json:"errorMessage"`
	Country       string        `json:"country"`
	Programs      []interface{} `json:"programs"`
	Type          string        `json:"type"`
	Address       string        `json:"address"`
	List          string        `json:"list"`
}

func (er ErrorResponse) Error() string {
	return er.ErrorMessage
}

func (c *client) Search(ctx context.Context, params SearchParams) ([]SearchResult, error) {
	// params.Programs must be a defined array
	if params.Programs == nil {
		params.Programs = []string{}
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(params)
	if err != nil {
		return nil, fmt.Errorf("encoding search params: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", searchBaseAddress, &body)
	if err != nil {
		return nil, fmt.Errorf("creating search request: %w", err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.6")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://sanctionslist.ofac.treas.gov")
	req.Header.Set("Referer", "https://sanctionslist.ofac.treas.gov/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making search request: %w", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	bs, _ := io.ReadAll(resp.Body)

	var results []SearchResult
	err = json.NewDecoder(bytes.NewReader(bs)).Decode(&results)
	if err != nil {
		var error ErrorResponse
		if err2 := json.Unmarshal(bs, &error); err2 == nil && error.ErrorMessage != "" {
			err = error
		}
		return nil, fmt.Errorf("reading search response: %s: %w", string(bs), err)
	}

	return results, nil
}
