package geocoding

import (
	"cmp"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestGoogleGeocoder_Success(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
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

		geocoder, err := NewGoogleGeocoder(ProviderConfig{
			APIKey:  "test-key",
			BaseURL: server.URL,
		})
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
		require.InDelta(t, 40.762363, coords.Latitude, 0.0001)
		require.InDelta(t, -73.8313912, coords.Longitude, 0.0001)
		require.Equal(t, "rooftop", coords.Accuracy)
	})

	t.Run("live", func(t *testing.T) {
		if testing.Short() {
			t.Skip("-short flag provided")
		}
		apiKey := cmp.Or(os.Getenv("GOOGLE_MAPS_API_KEY"), os.Getenv("GOOGLE_API_KEY"))
		if apiKey == "" {
			t.Skip("GOOGLE_MAPS_API_KEY not set")
		}

		geocoder, err := NewGoogleGeocoder(ProviderConfig{
			APIKey:  apiKey,
			Timeout: 30 * time.Second,
		})
		require.NoError(t, err)

		addr := search.Address{
			Line1:   "Statue of Liberty",
			City:    "New York",
			State:   "NY",
			Country: "US",
		}

		coords, err := geocoder.Geocode(context.Background(), addr)
		require.NoError(t, err)
		require.NotNil(t, coords)
		require.InDelta(t, 40.69, coords.Latitude, 0.01)
		require.InDelta(t, -74.04, coords.Longitude, 0.01)
		t.Logf("Google live: lat=%.6f lng=%.6f accuracy=%s", coords.Latitude, coords.Longitude, coords.Accuracy)
	})
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
