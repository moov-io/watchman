package geocoding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

const googleBaseURL = "https://maps.googleapis.com/maps/api/geocode/json"

// GoogleGeocoder implements the Geocoder interface using the Google Maps API.
type GoogleGeocoder struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

type googleResponse struct {
	Results      []googleResult `json:"results"`
	Status       string         `json:"status"`
	ErrorMessage string         `json:"error_message"`
}

type googleResult struct {
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		LocationType string `json:"location_type"`
	} `json:"geometry"`
}

// NewGoogleGeocoder creates a new Google Maps geocoder.
func NewGoogleGeocoder(conf ProviderConfig) (*GoogleGeocoder, error) {
	if conf.APIKey == "" {
		return nil, fmt.Errorf("Google Maps API key is required")
	}

	baseURL := conf.BaseURL
	if baseURL == "" {
		baseURL = googleBaseURL
	}

	timeout := conf.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &GoogleGeocoder{
		apiKey:  conf.APIKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// Geocode converts an address to coordinates using the Google Maps API.
func (g *GoogleGeocoder) Geocode(ctx context.Context, addr search.Address) (*Coordinates, error) {
	query := addr.Format()
	if query == "" {
		return nil, nil
	}

	reqURL := fmt.Sprintf("%s?key=%s&address=%s",
		g.baseURL,
		url.QueryEscape(g.apiKey),
		url.QueryEscape(query),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	var result googleResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	switch result.Status {
	case "OK":
		// Success, continue processing
	case "ZERO_RESULTS":
		return nil, nil // No results found
	case "OVER_QUERY_LIMIT", "REQUEST_DENIED", "INVALID_REQUEST", "UNKNOWN_ERROR":
		return nil, fmt.Errorf("google geocoding error: %s - %s", result.Status, result.ErrorMessage)
	default:
		return nil, fmt.Errorf("google geocoding unexpected status: %s", result.Status)
	}

	if len(result.Results) == 0 {
		return nil, nil
	}

	r := result.Results[0]
	return &Coordinates{
		Latitude:  r.Geometry.Location.Lat,
		Longitude: r.Geometry.Location.Lng,
		Accuracy:  googleLocationTypeToAccuracy(r.Geometry.LocationType),
	}, nil
}

// Name returns the provider name.
func (g *GoogleGeocoder) Name() string {
	return "google"
}

// googleLocationTypeToAccuracy converts Google location type to accuracy string.
func googleLocationTypeToAccuracy(locationType string) string {
	switch locationType {
	case "ROOFTOP":
		return "rooftop"
	case "RANGE_INTERPOLATED":
		return "street"
	case "GEOMETRIC_CENTER":
		return "city"
	case "APPROXIMATE":
		return "approximate"
	default:
		return "approximate"
	}
}
