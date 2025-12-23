package geocoding

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
)

// Service provides geocoding with caching and rate limiting.
// It manages L1 (in-memory) and L2 (database) caches and applies
// rate limiting to external API calls.
type Service struct {
	logger   log.Logger
	conf     Config
	geocoder Geocoder
	limiter  *rate.Limiter

	// L1 cache (in-memory LRU with TTL)
	l1Cache *lru.Cache[string, cacheEntry]
	l1TTL   time.Duration
	mu      sync.RWMutex

	// L2 cache (database)
	l2Repo Repository
}

type cacheEntry struct {
	coords    *Coordinates
	timestamp time.Time
}

// NewService creates a new geocoding service with the provided configuration.
// Returns nil if geocoding is disabled.
func NewService(logger log.Logger, conf Config, database db.DB) (*Service, error) {
	if !conf.Enabled {
		logger.Info().Log("geocoding service is disabled")
		return nil, nil
	}

	// Create geocoder based on provider config
	geocoder, err := createGeocoder(conf.Provider)
	if err != nil {
		return nil, fmt.Errorf("creating geocoder: %w", err)
	}

	// Create rate limiter
	limiter := rate.NewLimiter(
		rate.Limit(conf.RateLimit.RequestsPerSecond),
		conf.RateLimit.Burst,
	)

	// Set defaults if not configured
	l1MaxSize := conf.Cache.L1MaxSize
	if l1MaxSize <= 0 {
		l1MaxSize = 10000
	}

	l1TTL := conf.Cache.L1TTL
	if l1TTL <= 0 {
		l1TTL = 24 * time.Hour
	}

	// Create L1 cache
	l1Cache, err := lru.New[string, cacheEntry](l1MaxSize)
	if err != nil {
		return nil, fmt.Errorf("creating L1 cache: %w", err)
	}

	// Create L2 repository if enabled and database available
	var l2Repo Repository
	if conf.Cache.L2Enabled && database != nil {
		l2Repo = NewRepository(database)
	}

	logger.Info().Logf("geocoding service enabled with provider=%s, rateLimit=%.1f/sec, l1CacheSize=%d",
		geocoder.Name(), conf.RateLimit.RequestsPerSecond, l1MaxSize)

	return &Service{
		logger:   logger,
		conf:     conf,
		geocoder: geocoder,
		limiter:  limiter,
		l1Cache:  l1Cache,
		l1TTL:    l1TTL,
		l2Repo:   l2Repo,
	}, nil
}

// GeocodeAddress geocodes a single address with caching and rate limiting.
// Returns nil if geocoding fails or is disabled.
func (s *Service) GeocodeAddress(ctx context.Context, addr search.Address) (*Coordinates, error) {
	if s == nil {
		return nil, nil
	}

	ctx, span := telemetry.StartSpan(ctx, "geocode-address", trace.WithAttributes(
		attribute.String("geocoder.provider", s.geocoder.Name()),
	))
	defer span.End()

	cacheKey := s.addressCacheKey(addr)

	// Check L1 cache
	if coords := s.checkL1Cache(cacheKey); coords != nil {
		span.SetAttributes(attribute.String("geocoder.cache", "l1-hit"))
		return coords, nil
	}

	// Check L2 cache
	if s.l2Repo != nil {
		coords, err := s.l2Repo.Get(ctx, cacheKey)
		if err == nil && coords != nil {
			s.setL1Cache(cacheKey, coords)
			span.SetAttributes(attribute.String("geocoder.cache", "l2-hit"))
			return coords, nil
		}
	}

	span.SetAttributes(attribute.String("geocoder.cache", "miss"))

	// Rate limit before calling external service
	if err := s.limiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limit wait: %w", err)
	}

	// Call geocoder
	coords, err := s.geocoder.Geocode(ctx, addr)
	if err != nil {
		s.logger.Warn().LogErrorf("geocoding failed for address %q: %v", addr.Format(), err)
		return nil, nil // Graceful degradation
	}

	if coords != nil {
		// Store in L1 cache
		s.setL1Cache(cacheKey, coords)

		// Store in L2 cache
		if s.l2Repo != nil {
			if err := s.l2Repo.Set(ctx, cacheKey, coords); err != nil {
				s.logger.Warn().LogErrorf("failed to store in L2 cache: %v", err)
			}
		}
	}

	return coords, nil
}

// GeocodeAddresses geocodes multiple addresses and returns them with coordinates populated.
// This method is the primary integration point for entity mapping.
func (s *Service) GeocodeAddresses(ctx context.Context, addresses []search.Address) []search.Address {
	if s == nil || len(addresses) == 0 {
		return addresses
	}

	result := make([]search.Address, len(addresses))
	copy(result, addresses)

	for i := range result {
		coords, err := s.GeocodeAddress(ctx, result[i])
		if err != nil {
			continue
		}
		if coords != nil {
			result[i].Latitude = coords.Latitude
			result[i].Longitude = coords.Longitude
		}
	}

	return result
}

// addressCacheKey generates a normalized cache key for an address.
func (s *Service) addressCacheKey(addr search.Address) string {
	return fmt.Sprintf("%s|%s|%s|%s|%s|%s",
		addr.Line1, addr.Line2, addr.City,
		addr.PostalCode, addr.State, addr.Country)
}

// checkL1Cache returns cached coordinates if present and not expired.
func (s *Service) checkL1Cache(key string) *Coordinates {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if entry, ok := s.l1Cache.Get(key); ok {
		if time.Since(entry.timestamp) < s.l1TTL {
			return entry.coords
		}
		// TTL expired, remove from cache
		s.l1Cache.Remove(key)
	}
	return nil
}

// setL1Cache stores coordinates in the L1 cache.
func (s *Service) setL1Cache(key string, coords *Coordinates) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.l1Cache.Add(key, cacheEntry{
		coords:    coords,
		timestamp: time.Now(),
	})
}

// createGeocoder creates a geocoder based on provider configuration.
func createGeocoder(conf ProviderConfig) (Geocoder, error) {
	switch conf.Name {
	case "opencage":
		return NewOpenCageGeocoder(conf)
	case "google":
		return NewGoogleGeocoder(conf)
	case "nominatim":
		return NewNominatimGeocoder(conf)
	case "":
		return nil, fmt.Errorf("geocoding provider name is required")
	default:
		return nil, fmt.Errorf("unknown geocoding provider: %s", conf.Name)
	}
}
