package geocoding

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/moov-io/watchman/internal/db"
)

// Repository defines the interface for the L2 (database) cache.
type Repository interface {
	// Get retrieves coordinates from the cache by key.
	// Returns nil if not found.
	Get(ctx context.Context, cacheKey string) (*Coordinates, error)

	// Set stores coordinates in the cache.
	Set(ctx context.Context, cacheKey string, coords *Coordinates) error
}

type sqlRepository struct {
	db db.DB
}

// NewRepository creates a new database repository for geocoding cache.
func NewRepository(database db.DB) Repository {
	if database == nil {
		return nil
	}
	return &sqlRepository{db: database}
}

// Get retrieves coordinates from the database cache.
func (r *sqlRepository) Get(ctx context.Context, cacheKey string) (*Coordinates, error) {
	query := `SELECT latitude, longitude, accuracy FROM geocoding_cache WHERE cache_key = ? LIMIT 1`

	var coords Coordinates
	row := r.db.QueryRowContext(ctx, query, cacheKey)
	err := row.Scan(&coords.Latitude, &coords.Longitude, &coords.Accuracy)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying geocoding cache: %w", err)
	}

	return &coords, nil
}

// Set stores coordinates in the database cache.
// Uses upsert to handle duplicate keys.
func (r *sqlRepository) Set(ctx context.Context, cacheKey string, coords *Coordinates) error {
	// Use INSERT ... ON CONFLICT for PostgreSQL compatibility
	// The db wrapper will handle rebinding for MySQL
	query := `
		INSERT INTO geocoding_cache (cache_key, latitude, longitude, accuracy, created_at)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT (cache_key) DO UPDATE SET
			latitude = EXCLUDED.latitude,
			longitude = EXCLUDED.longitude,
			accuracy = EXCLUDED.accuracy,
			created_at = EXCLUDED.created_at
	`

	_, err := r.db.ExecContext(ctx, query,
		cacheKey,
		coords.Latitude,
		coords.Longitude,
		coords.Accuracy,
		time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("inserting into geocoding cache: %w", err)
	}

	return nil
}
