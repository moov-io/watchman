package index_test

import (
	"context"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestIndex_Stats(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := ingest.NewRepository(db)
		lists := index.NewLists(repo)

		// find empty stats
		found := lists.LatestStats()
		require.True(t, found.StartedAt.IsZero())

		// update
		stats := download.Stats{
			StartedAt: time.Now().UTC(),
		}
		lists.Update(stats)

		// find populated stats
		found = lists.LatestStats()
		require.False(t, found.StartedAt.IsZero())
	})
}

func TestIndex_GetEntities(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := ingest.NewRepository(db)
		lists := index.NewLists(repo)

		entity := ofactest.FindEntity(t, "11195")
		entity.Source = "custom"
		entity.PreparedFields = search.PreparedFields{} // empty
		require.Empty(t, entity.PreparedFields.Addresses)

		entities := []search.Entity[search.Value]{entity}

		ctx := context.Background()
		err := repo.Upsert(ctx, "custom", entities)
		require.NoError(t, err)

		// Verify the entity is found
		found, err := repo.ListBySource(ctx, "", "custom", 10)
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, "11195", found[0].SourceID)
		require.Len(t, found[0].PreparedFields.Addresses, 5) // repo does call .Normalize()

		// Read Entities and ensure PrecomputedFields is populated
		found, err = lists.GetEntities(ctx, "custom")
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, "11195", found[0].SourceID)
		require.Len(t, found[0].PreparedFields.Addresses, 5) // index.Lists does call .Normalize()
	})
}
