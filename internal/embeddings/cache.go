package embeddings

import (
	"context"
	"strings"

	"github.com/moov-io/watchman/internal/db"
)

type Cache interface {
	Get(ctx context.Context, text string) ([]float64, bool)
	Put(ctx context.Context, text string, embedding []float64)
}

func NewCache(ctx context.Context, config Config, database db.DB) (Cache, error) {
	switch strings.ToLower(config.Cache.Type) {
	case "", "none":
		return &noopCache{}, nil

	case "memory":
		return newMemoryCache(config.Cache.Size)

	case "sql", "mysql", "postgres":
		return newSqlRepository(ctx, config, database)
	}

	return nil, nil
}

type noopCache struct{}

var _ Cache = (&noopCache{})

func (c *noopCache) Get(ctx context.Context, text string) ([]float64, bool) {
	return nil, false
}

func (c *noopCache) Put(ctx context.Context, text string, embedding []float64) {
	return
}
