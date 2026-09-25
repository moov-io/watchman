package geocoding

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestOpenCageGeocoder_Success(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Contains(t, r.URL.RawQuery, "key=test-key")
			require.Contains(t, r.URL.RawQuery, "q=123+Main+St+New+York+NY+US")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"results": [{
					"geometry": {"lat": 40.754057, "lng": -73.956462},
					"confidence": 9
				}],
				"status": {"code": 200, "message": "OK"}
			}`))
		}))
		defer server.Close()

		geocoder, err := NewOpenCageGeocoder(ProviderConfig{
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
		require.InDelta(t, 40.754057, coords.Latitude, 0.0001)
		require.InDelta(t, -73.956462, coords.Longitude, 0.0001)
		require.Equal(t, "rooftop", coords.Accuracy)
	})

	t.Run("live", func(t *testing.T) {
		if testing.Short() {
			t.Skip("-short flag provided")
		}
		apiKey := os.Getenv("OPENCAGE_API_KEY")
		if apiKey == "" {
			t.Skip("OPENCAGE_API_KEY not set")
		}

		geocoder, err := NewOpenCageGeocoder(ProviderConfig{
			APIKey:  apiKey,
			Timeout: 30 * time.Second,
		})
		require.NoError(t, err)

		addr := search.Address{
			Line1:   "1600 Amphitheatre Parkway",
			City:    "Mountain View",
			State:   "CA",
			Country: "US",
		}

		coords, err := geocoder.Geocode(context.Background(), addr)
		require.NoError(t, err)
		require.NotNil(t, coords)
		require.InDelta(t, 37.42, coords.Latitude, 0.01)
		require.InDelta(t, -122.08, coords.Longitude, 0.01)
		t.Logf("OpenCage live: lat=%.6f lng=%.6f accuracy=%s", coords.Latitude, coords.Longitude, coords.Accuracy)
	})
}

func TestOpenCageGeocoder_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [],
			"status": {"code": 200, "message": "OK"}
		}`))
	}))
	defer server.Close()

	geocoder, err := NewOpenCageGeocoder(ProviderConfig{
		APIKey:  "test-key",
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "unknown"})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestOpenCageGeocoder_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [],
			"status": {"code": 401, "message": "Invalid API key"}
		}`))
	}))
	defer server.Close()

	geocoder, err := NewOpenCageGeocoder(ProviderConfig{
		APIKey:  "invalid-key",
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "test"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid API key")
	require.Nil(t, coords)
}

func TestOpenCageGeocoder_EmptyAddress(t *testing.T) {
	geocoder, err := NewOpenCageGeocoder(ProviderConfig{
		APIKey: "test-key",
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestOpenCageGeocoder_MissingAPIKey(t *testing.T) {
	_, err := NewOpenCageGeocoder(ProviderConfig{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "API key is required")
}

func TestOpenCageGeocoder_Name(t *testing.T) {
	geocoder, err := NewOpenCageGeocoder(ProviderConfig{APIKey: "test"})
	require.NoError(t, err)
	require.Equal(t, "opencage", geocoder.Name())
}

func TestConfidenceToAccuracy(t *testing.T) {
	tests := []struct {
		confidence int
		expected   string
	}{
		{10, "rooftop"},
		{9, "rooftop"},
		{8, "street"},
		{7, "street"},
		{6, "city"},
		{5, "city"},
		{4, "state"},
		{3, "state"},
		{2, "country"},
		{1, "country"},
		{0, "country"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tt.expected, confidenceToAccuracy(tt.confidence))
		})
	}
}
