package geocoding

import "time"

// Config holds the configuration for the geocoding service.
type Config struct {
	// Enabled controls whether geocoding is active.
	Enabled bool

	// Provider configuration for the geocoding API.
	Provider ProviderConfig

	// RateLimit configuration to control API request rates.
	RateLimit RateLimitConfig

	// Cache configuration for L1 (in-memory) and L2 (database) caching.
	Cache CacheConfig
}

// ProviderConfig holds settings for a geocoding provider.
type ProviderConfig struct {
	// Name of the provider: "opencage", "google", or "nominatim"
	Name string

	// APIKey for authentication with the provider.
	// Can be set via GEOCODING_API_KEY environment variable.
	APIKey string

	// BaseURL allows overriding the default API endpoint.
	// Useful for self-hosted Nominatim instances.
	BaseURL string

	// Timeout for API requests. Default: 10s
	Timeout time.Duration
}

// RateLimitConfig controls the rate of geocoding API requests.
type RateLimitConfig struct {
	// RequestsPerSecond defines the sustained request rate.
	// Free tier typically allows 1/sec, paid tiers allow up to 40/sec.
	RequestsPerSecond float64

	// Burst allows temporary exceeding of the rate limit.
	Burst int
}

// CacheConfig controls caching behavior for geocoding results.
type CacheConfig struct {
	// L1MaxSize is the maximum number of entries in the in-memory LRU cache.
	// Default: 10000
	L1MaxSize int

	// L1TTL is the time-to-live for L1 cache entries.
	// Default: 24h
	L1TTL time.Duration

	// L2Enabled enables the database-backed persistent cache.
	// Requires a database connection.
	L2Enabled bool
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Enabled: false,
		Provider: ProviderConfig{
			Name:    "opencage",
			Timeout: 10 * time.Second,
		},
		RateLimit: RateLimitConfig{
			RequestsPerSecond: 1,
			Burst:             5,
		},
		Cache: CacheConfig{
			L1MaxSize: 10000,
			L1TTL:     24 * time.Hour,
			L2Enabled: true,
		},
	}
}
