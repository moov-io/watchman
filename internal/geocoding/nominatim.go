package geocoding

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

const nominatimBaseURL = "https://nominatim.openstreetmap.org/search"

// NominatimGeocoder implements the Geocoder interface using OpenStreetMap's Nominatim API.
// Note: The public Nominatim server has strict usage policies (max 1 request/second).
// For production use, consider hosting your own Nominatim instance.
type NominatimGeocoder struct {
	baseURL string
	client  *http.Client
}

type nominatimResult struct {
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	DisplayName string  `json:"display_name"`
	Class       string  `json:"class"`
	Type        string  `json:"type"`
	Importance  float64 `json:"importance"`
}

// NewNominatimGeocoder creates a new Nominatim geocoder.
// APIKey is not required for Nominatim (it's free and open).
func NewNominatimGeocoder(conf ProviderConfig) (*NominatimGeocoder, error) {
	baseURL := conf.BaseURL
	if baseURL == "" {
		baseURL = nominatimBaseURL
	}

	timeout := conf.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &NominatimGeocoder{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// Geocode converts an address to coordinates using the Nominatim API.
func (g *NominatimGeocoder) Geocode(ctx context.Context, addr search.Address) (*Coordinates, error) {
	query := addr.Format()
	if query == "" {
		return nil, nil
	}

	reqURL := fmt.Sprintf("%s?format=json&q=%s&limit=1",
		g.baseURL,
		url.QueryEscape(query),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Nominatim requires a descriptive User-Agent
	req.Header.Set("User-Agent", "moov-io/watchman (https://github.com/moov-io/watchman)")

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nominatim returned status %d", resp.StatusCode)
	}

	var results []nominatimResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	if len(results) == 0 {
		return nil, nil // No results found
	}

	r := results[0]
	lat, err := strconv.ParseFloat(r.Lat, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing latitude: %w", err)
	}

	lon, err := strconv.ParseFloat(r.Lon, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing longitude: %w", err)
	}

	return &Coordinates{
		Latitude:  lat,
		Longitude: lon,
		Accuracy:  nominatimClassToAccuracy(r.Class, r.Type),
	}, nil
}

// Name returns the provider name.
func (g *NominatimGeocoder) Name() string {
	return "nominatim"
}

// nominatimClassToAccuracy converts Nominatim class/type to accuracy string.
func nominatimClassToAccuracy(class, typ string) string {
	switch class {
	case "building":
		return "rooftop"
	case "place":
		switch typ {
		case "house", "building":
			return "rooftop"
		case "street", "road":
			return "street"
		case "city", "town", "village", "hamlet":
			return "city"
		case "state", "province", "region":
			return "state"
		case "country":
			return "country"
		}
	case "highway":
		return "street"
	case "boundary":
		switch typ {
		case "administrative":
			return "city"
		}
	}
	return "approximate"
}
