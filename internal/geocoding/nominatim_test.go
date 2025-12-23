package geocoding

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestNominatimGeocoder_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.RawQuery, "format=json")
		require.Contains(t, r.URL.RawQuery, "q=")
		require.Contains(t, r.URL.RawQuery, "limit=1")

		// Check User-Agent header
		userAgent := r.Header.Get("User-Agent")
		require.Contains(t, userAgent, "moov-io/watchman")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{
			"lat": "40.7128",
			"lon": "-74.0060",
			"display_name": "123 Main St, New York, NY, USA",
			"class": "building",
			"type": "house",
			"importance": 0.9
		}]`))
	}))
	defer server.Close()

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
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
	require.Equal(t, 40.7128, coords.Latitude)
	require.Equal(t, -74.0060, coords.Longitude)
	require.Equal(t, "rooftop", coords.Accuracy)
}

func TestNominatimGeocoder_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "unknown"})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestNominatimGeocoder_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
		BaseURL: server.URL,
	})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{Line1: "test"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "status 429")
	require.Nil(t, coords)
}

func TestNominatimGeocoder_EmptyAddress(t *testing.T) {
	geocoder, err := NewNominatimGeocoder(ProviderConfig{})
	require.NoError(t, err)

	coords, err := geocoder.Geocode(context.Background(), search.Address{})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestNominatimGeocoder_Name(t *testing.T) {
	geocoder, err := NewNominatimGeocoder(ProviderConfig{})
	require.NoError(t, err)
	require.Equal(t, "nominatim", geocoder.Name())
}

func TestNominatimClassToAccuracy(t *testing.T) {
	tests := []struct {
		class    string
		typ      string
		expected string
	}{
		{"building", "", "rooftop"},
		{"place", "house", "rooftop"},
		{"place", "building", "rooftop"},
		{"place", "street", "street"},
		{"place", "road", "street"},
		{"place", "city", "city"},
		{"place", "town", "city"},
		{"place", "village", "city"},
		{"place", "state", "state"},
		{"place", "country", "country"},
		{"highway", "", "street"},
		{"boundary", "administrative", "city"},
		{"boundary", "other", "approximate"},
		{"unknown", "", "approximate"},
		{"", "", "approximate"},
	}

	for _, tt := range tests {
		t.Run(tt.class+"/"+tt.typ, func(t *testing.T) {
			require.Equal(t, tt.expected, nominatimClassToAccuracy(tt.class, tt.typ))
		})
	}
}

func TestNominatimGeocoder_NoAPIKeyRequired(t *testing.T) {
	// Nominatim doesn't require an API key
	geocoder, err := NewNominatimGeocoder(ProviderConfig{})
	require.NoError(t, err)
	require.NotNil(t, geocoder)
}
