package geocoding

import (
	"cmp"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestGoogleGeocoder_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.RawQuery, "key=test-key")
		require.Contains(t, r.URL.RawQuery, "address=")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [{
				"geometry": {
					"location": {"lat": 40.762363, "lng": -73.8313912},
					"location_type": "ROOFTOP"
				}
			}],
			"status": "OK"
		}`))
	}))
	defer server.Close()

	var apiKey string
	if !testing.Short() {
		apiKey = os.Getenv("GOOGLE_MAPS_API_KEY")
	}
	conf := ProviderConfig{
		APIKey: cmp.Or(apiKey, "test-key"),
	}
	if apiKey == "" {
		conf.BaseURL = server.URL
	}

	geocoder, err := NewGoogleGeocoder(conf)
	require.NoError(t, err)

	addr := search.Address{
		Line1:   "123 Main St",
		City:    "New York",
		State:   "NY",
		Country: "US",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords)

	require.InDelta(t, 40.762363, coords.Latitude, 0.001)
	require.InDelta(t, -73.8313912, coords.Longitude, 0.001)
	require.Equal(t, "rooftop", coords.Accuracy)
}

func TestGoogleGeocoder_ZeroResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [],
			"status": "ZERO_RESULTS"
		}`))
	}))
	defer server.Close()

	geocoder, err := NewGoogleGeocoder(ProviderConfig{
		APIKey:  "test-key",
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "unknown"})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestGoogleGeocoder_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [],
			"status": "REQUEST_DENIED",
			"error_message": "API key is invalid"
		}`))
	}))
	defer server.Close()

	geocoder, err := NewGoogleGeocoder(ProviderConfig{
		APIKey:  "invalid-key",
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "test"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "REQUEST_DENIED")
	require.Nil(t, coords)
}

func TestGoogleGeocoder_EmptyAddress(t *testing.T) {
	geocoder, err := NewGoogleGeocoder(ProviderConfig{
		APIKey: "test-key",
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestGoogleGeocoder_MissingAPIKey(t *testing.T) {
	_, err := NewGoogleGeocoder(ProviderConfig{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "API key is required")
}

func TestGoogleGeocoder_Name(t *testing.T) {
	geocoder, err := NewGoogleGeocoder(ProviderConfig{APIKey: "test"})
	require.NoError(t, err)
	require.Equal(t, "google", geocoder.Name())
}

func TestGoogleLocationTypeToAccuracy(t *testing.T) {
	tests := []struct {
		locationType string
		expected     string
	}{
		{"ROOFTOP", "rooftop"},
		{"RANGE_INTERPOLATED", "street"},
		{"GEOMETRIC_CENTER", "city"},
		{"APPROXIMATE", "approximate"},
		{"UNKNOWN", "approximate"},
		{"", "approximate"},
	}

	for _, tt := range tests {
		t.Run(tt.locationType, func(t *testing.T) {
			require.Equal(t, tt.expected, googleLocationTypeToAccuracy(tt.locationType))
		})
	}
}
