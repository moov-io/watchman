package geocoding

import (
	"context"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestService_Disabled(t *testing.T) {
	conf := Config{Enabled: false}
	svc, err := NewService(log.NewTestLogger(), conf, nil)
	require.NoError(t, err)
	require.Nil(t, svc)
}

func TestService_GeocodeAddress_NilService(t *testing.T) {
	var svc *Service
	coords, err := svc.GeocodeAddress(context.Background(), search.Address{})
	require.NoError(t, err)
	require.Nil(t, coords)
}

func TestService_GeocodeAddresses_NilService(t *testing.T) {
	var svc *Service
	addresses := []search.Address{{Line1: "123 Main St"}}
	result := svc.GeocodeAddresses(context.Background(), addresses)
	require.Equal(t, addresses, result)
}

func TestService_AddressCacheKey(t *testing.T) {
	svc := &Service{}

	addr1 := search.Address{Line1: "123 Main", City: "NYC", Country: "US"}
	addr2 := search.Address{Line1: "123 Main", City: "NYC", Country: "US"}
	addr3 := search.Address{Line1: "456 Oak", City: "LA", Country: "US"}

	key1 := svc.addressCacheKey(addr1)
	key2 := svc.addressCacheKey(addr2)
	key3 := svc.addressCacheKey(addr3)

	require.Equal(t, key1, key2, "same addresses should have same cache key")
	require.NotEqual(t, key1, key3, "different addresses should have different cache keys")
}

func TestService_L1Cache(t *testing.T) {
	conf := Config{
		Enabled: true,
		Provider: ProviderConfig{
			Name:   "nominatim", // Nominatim doesn't require API key
			APIKey: "",
		},
		RateLimit: RateLimitConfig{
			RequestsPerSecond: 10,
			Burst:             10,
		},
		Cache: CacheConfig{
			L1MaxSize: 100,
			L1TTL:     time.Hour,
		},
	}

	svc, err := NewService(log.NewTestLogger(), conf, nil)
	require.NoError(t, err)
	require.NotNil(t, svc)

	// Test cache key generation
	addr := search.Address{
		Line1:   "123 Main St",
		City:    "New York",
		State:   "NY",
		Country: "US",
	}

	key := svc.addressCacheKey(addr)
	require.NotEmpty(t, key)

	// Test L1 cache operations
	coords := &Coordinates{Latitude: 40.7128, Longitude: -74.0060, Accuracy: "rooftop"}
	svc.setL1Cache(key, coords)

	cached := svc.checkL1Cache(key)
	require.NotNil(t, cached)
	require.Equal(t, 40.7128, cached.Latitude)
	require.Equal(t, -74.0060, cached.Longitude)
	require.Equal(t, "rooftop", cached.Accuracy)
}

func TestService_L1Cache_TTLExpiration(t *testing.T) {
	conf := Config{
		Enabled: true,
		Provider: ProviderConfig{
			Name: "nominatim",
		},
		RateLimit: RateLimitConfig{
			RequestsPerSecond: 10,
			Burst:             10,
		},
		Cache: CacheConfig{
			L1MaxSize: 100,
			L1TTL:     1 * time.Millisecond, // Very short TTL for testing
		},
	}

	svc, err := NewService(log.NewTestLogger(), conf, nil)
	require.NoError(t, err)
	require.NotNil(t, svc)

	key := "test-key"
	coords := &Coordinates{Latitude: 40.7128, Longitude: -74.0060}
	svc.setL1Cache(key, coords)

	// Wait for TTL to expire
	time.Sleep(10 * time.Millisecond)

	cached := svc.checkL1Cache(key)
	require.Nil(t, cached, "cache entry should be expired")
}

func TestDefaultConfig(t *testing.T) {
	conf := DefaultConfig()

	require.False(t, conf.Enabled)
	require.Equal(t, "opencage", conf.Provider.Name)
	require.Equal(t, 10*time.Second, conf.Provider.Timeout)
	require.Equal(t, float64(1), conf.RateLimit.RequestsPerSecond)
	require.Equal(t, 5, conf.RateLimit.Burst)
	require.Equal(t, 10000, conf.Cache.L1MaxSize)
	require.Equal(t, 24*time.Hour, conf.Cache.L1TTL)
	require.True(t, conf.Cache.L2Enabled)
}
