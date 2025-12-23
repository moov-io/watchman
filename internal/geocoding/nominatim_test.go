package geocoding

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

const useMockNominatim = true

func pickNominatimBaseURL(server *httptest.Server) string {
	if useMockNominatim {
		return server.URL
	}
	return nominatimBaseURL
}

func TestNominatimGeocoder_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.RawQuery, "format=json")
		require.Contains(t, r.URL.RawQuery, "q=")
		require.Contains(t, r.URL.RawQuery, "limit=1")

		// Check User-Agent header
		userAgent := r.Header.Get("User-Agent")
		require.Contains(t, userAgent, "moov-io/watchman")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"place_id":421093542,"licence":"Data © OpenStreetMap contributors, ODbL 1.0. http://osm.org/copyright","osm_type":"relation","osm_id":19761182,"lat":"38.8976387","lon":"-77.0365528","class":"office","type":"government","place_rank":30,"importance":0.6863355973183977,"addresstype":"office","name":"White House","display_name":"White House, 1600, Pennsylvania Avenue Northwest, Downtown, Ward 2, Washington, District of Columbia, 20500, United States","boundingbox":["38.8974904","38.8977959","-77.0368541","-77.0362517"]}]`))
	}))
	defer server.Close()

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
		BaseURL: pickNominatimBaseURL(server),
	})
	require.NoError(t, err)

	addr := search.Address{
		Line1:   "1600 Pennsylvania Avenue NW",
		City:    "Washington",
		State:   "DC",
		Country: "US",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords)
	require.InDelta(t, 38.8976387, coords.Latitude, 0.001)
	require.InDelta(t, -77.0365528, coords.Longitude, 0.001)
	require.Equal(t, "approximate", coords.Accuracy)
}

func TestNominatimGeocoder_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
		BaseURL: pickNominatimBaseURL(server),
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
		BaseURL: pickNominatimBaseURL(server),
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
