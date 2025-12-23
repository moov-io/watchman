package geocoding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/pkg/search"
)

const openCageBaseURL = "https://api.opencagedata.com/geocode/v1/json"

// OpenCageGeocoder implements the Geocoder interface using the OpenCage API.
type OpenCageGeocoder struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

type openCageResponse struct {
	Results []openCageResult `json:"results"`
	Status  openCageStatus   `json:"status"`
}

type openCageResult struct {
	Geometry struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"geometry"`
	Confidence int `json:"confidence"`
}

type openCageStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewOpenCageGeocoder creates a new OpenCage geocoder.
func NewOpenCageGeocoder(conf ProviderConfig) (*OpenCageGeocoder, error) {
	if conf.APIKey == "" {
		return nil, fmt.Errorf("OpenCage API key is required")
	}

	baseURL := conf.BaseURL
	if baseURL == "" {
		baseURL = openCageBaseURL
	}

	timeout := conf.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &OpenCageGeocoder{
		apiKey:  conf.APIKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// Geocode converts an address to coordinates using the OpenCage API.
func (g *OpenCageGeocoder) Geocode(ctx context.Context, addr search.Address) (*Coordinates, error) {
	query := addr.Format()
	if query == "" {
		return nil, nil
	}

	u, err := url.Parse(g.baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing base url: %w", err)
	}

	q := u.Query()
	q.Set("key", g.apiKey)
	q.Set("q", query)
	q.Set("no_annotations", "1")
	q.Set("limit", "1")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", fmt.Sprintf("moov-io/watchman:%s", watchman.Version))

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("opencage returned status %d", resp.StatusCode)
	}

	var result openCageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if result.Status.Code != 200 {
		return nil, fmt.Errorf("opencage error: %s (code %d)", result.Status.Message, result.Status.Code)
	}

	if len(result.Results) == 0 {
		return nil, nil // No results found
	}

	r := result.Results[0]
	return &Coordinates{
		Latitude:  r.Geometry.Lat,
		Longitude: r.Geometry.Lng,
		Accuracy:  confidenceToAccuracy(r.Confidence),
	}, nil
}

// Name returns the provider name.
func (g *OpenCageGeocoder) Name() string {
	return "opencage"
}

// confidenceToAccuracy converts OpenCage confidence level to accuracy string.
// Confidence ranges from 1 (low) to 10 (high).
func confidenceToAccuracy(confidence int) string {
	switch {
	case confidence >= 9:
		return "rooftop"
	case confidence >= 7:
		return "street"
	case confidence >= 5:
		return "city"
	case confidence >= 3:
		return "state"
	default:
		return "country"
	}
}
