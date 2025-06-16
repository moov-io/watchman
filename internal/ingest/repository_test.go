package ingest

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/ofactest"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := NewRepository(db)

		ctx := context.Background()
		entity := ofactest.FindEntity(t, "44525")

		err := repo.Upsert(ctx, entity)
		require.NoError(t, err)

		found, err := repo.Get(ctx, entity.SourceID, entity.Source)
		require.NoError(t, err)
		require.NotNil(t, found)

		// Empty SourceData // TODO(adam): can we unmarshal
		entity.SourceData = make(map[string]interface{})
		found.SourceData = make(map[string]interface{})

		// Compare objects
		require.Equal(t, entity.Normalize(), found.Normalize())
	})
}
