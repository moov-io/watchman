// +build integration

package geocoding

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

// Integration tests for geocoding providers.
// These tests require real API keys and make actual network requests.
//
// Run with: go test -tags=integration -v ./internal/geocoding/...
//
// Required environment variables:
//   OPENCAGE_API_KEY - OpenCage API key (get free key at https://opencagedata.com)
//   GOOGLE_API_KEY   - Google Maps API key (optional)
//
// Note: These tests are rate-limited and should not be run frequently.

func TestIntegration_OpenCage_RealAPI(t *testing.T) {
	apiKey := os.Getenv("OPENCAGE_API_KEY")
	if apiKey == "" {
		t.Skip("OPENCAGE_API_KEY not set, skipping integration test")
	}

	geocoder, err := NewOpenCageGeocoder(ProviderConfig{
		APIKey:  apiKey,
		Timeout: 30 * time.Second,
	})
	require.NoError(t, err)

	// Test with a well-known address
	addr := search.Address{
		Line1:   "1600 Amphitheatre Parkway",
		City:    "Mountain View",
		State:   "CA",
		Country: "US",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords, "expected coordinates for Google HQ address")

	// Google HQ is approximately at 37.4220, -122.0841
	require.InDelta(t, 37.42, coords.Latitude, 0.01, "latitude should be close to 37.42")
	require.InDelta(t, -122.08, coords.Longitude, 0.01, "longitude should be close to -122.08")

	t.Logf("OpenCage returned: lat=%.6f, lng=%.6f, accuracy=%s",
		coords.Latitude, coords.Longitude, coords.Accuracy)
}

func TestIntegration_OpenCage_UnknownAddress(t *testing.T) {
	apiKey := os.Getenv("OPENCAGE_API_KEY")
	if apiKey == "" {
		t.Skip("OPENCAGE_API_KEY not set, skipping integration test")
	}

	geocoder, err := NewOpenCageGeocoder(ProviderConfig{
		APIKey:  apiKey,
		Timeout: 30 * time.Second,
	})
	require.NoError(t, err)

	// Test with a nonsense address
	addr := search.Address{
		Line1:   "xyzzy123456 nonexistent street",
		City:    "Nowhere City",
		Country: "XX",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	// Should return nil or very low confidence result
	t.Logf("OpenCage returned for unknown address: %+v", coords)
}

func TestIntegration_Nominatim_RealAPI(t *testing.T) {
	// Nominatim doesn't require API key but has strict rate limits
	// Be respectful: max 1 request per second for public API

	geocoder, err := NewNominatimGeocoder(ProviderConfig{
		Timeout: 30 * time.Second,
	})
	require.NoError(t, err)

	// Test with a well-known address
	addr := search.Address{
		Line1:   "Empire State Building",
		City:    "New York",
		State:   "NY",
		Country: "US",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords, "expected coordinates for Empire State Building")

	// Empire State Building is approximately at 40.7484, -73.9857
	require.InDelta(t, 40.75, coords.Latitude, 0.01, "latitude should be close to 40.75")
	require.InDelta(t, -73.99, coords.Longitude, 0.01, "longitude should be close to -73.99")

	t.Logf("Nominatim returned: lat=%.6f, lng=%.6f, accuracy=%s",
		coords.Latitude, coords.Longitude, coords.Accuracy)
}

func TestIntegration_Google_RealAPI(t *testing.T) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		t.Skip("GOOGLE_API_KEY not set, skipping integration test")
	}

	geocoder, err := NewGoogleGeocoder(ProviderConfig{
		APIKey:  apiKey,
		Timeout: 30 * time.Second,
	})
	require.NoError(t, err)

	// Test with a well-known address
	addr := search.Address{
		Line1:   "Statue of Liberty",
		City:    "New York",
		State:   "NY",
		Country: "US",
	}

	coords, err := geocoder.Geocode(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords, "expected coordinates for Statue of Liberty")

	// Statue of Liberty is approximately at 40.6892, -74.0445
	require.InDelta(t, 40.69, coords.Latitude, 0.01, "latitude should be close to 40.69")
	require.InDelta(t, -74.04, coords.Longitude, 0.01, "longitude should be close to -74.04")

	t.Logf("Google returned: lat=%.6f, lng=%.6f, accuracy=%s",
		coords.Latitude, coords.Longitude, coords.Accuracy)
}

func TestIntegration_FullService_WithCaching(t *testing.T) {
	apiKey := os.Getenv("OPENCAGE_API_KEY")
	if apiKey == "" {
		t.Skip("OPENCAGE_API_KEY not set, skipping integration test")
	}

	conf := Config{
		Enabled: true,
		Provider: ProviderConfig{
			Name:    "opencage",
			APIKey:  apiKey,
			Timeout: 30 * time.Second,
		},
		RateLimit: RateLimitConfig{
			RequestsPerSecond: 1, // Free tier limit
			Burst:             1,
		},
		Cache: CacheConfig{
			L1MaxSize: 100,
			L1TTL:     time.Hour,
			L2Enabled: false, // No database in this test
		},
	}

	svc, err := NewService(log.NewTestLogger(), conf, nil)
	require.NoError(t, err)
	require.NotNil(t, svc)

	addr := search.Address{
		Line1:   "10 Downing Street",
		City:    "London",
		Country: "UK",
	}

	// First call - should hit the API
	start := time.Now()
	coords1, err := svc.GeocodeAddress(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords1)
	apiDuration := time.Since(start)

	t.Logf("First call (API): lat=%.6f, lng=%.6f, took=%v",
		coords1.Latitude, coords1.Longitude, apiDuration)

	// Second call - should hit L1 cache (much faster)
	start = time.Now()
	coords2, err := svc.GeocodeAddress(context.Background(), addr)
	require.NoError(t, err)
	require.NotNil(t, coords2)
	cacheDuration := time.Since(start)

	t.Logf("Second call (cache): lat=%.6f, lng=%.6f, took=%v",
		coords2.Latitude, coords2.Longitude, cacheDuration)

	// Cache should be significantly faster
	require.Less(t, cacheDuration, apiDuration/10,
		"cached call should be at least 10x faster than API call")

	// Results should be identical
	require.Equal(t, coords1.Latitude, coords2.Latitude)
	require.Equal(t, coords1.Longitude, coords2.Longitude)
}
