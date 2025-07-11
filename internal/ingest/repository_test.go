package ingest

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := NewRepository(db)

		ctx := context.Background()
		entity := ofactest.FindEntity(t, "44525")
		entities := []search.Entity[search.Value]{entity}

		err := repo.Upsert(ctx, "fincen-person", entities)
		require.NoError(t, err)

		found, err := repo.Get(ctx, entity.SourceID, entity.Source)
		require.NoError(t, err)
		require.NotNil(t, found)

		// Empty SourceData // TODO(adam): can we unmarshal?
		entity.SourceData = make(map[string]interface{})
		found.SourceData = make(map[string]interface{})

		// Compare objects
		require.Equal(t, entity.Normalize(), found.Normalize())

		// List
		entities, err = repo.ListBySource(ctx, "", entity.Source, 10)
		require.NoError(t, err)
		require.Len(t, entities, 1)
		require.Equal(t, entity.SourceID, entities[0].SourceID)

		// List again
		entities, err = repo.ListBySource(ctx, "12345", entity.Source, 10)
		require.NoError(t, err)
		require.Len(t, entities, 1)
		require.Equal(t, entity.SourceID, entities[0].SourceID)

		// Find nothing
		entities, err = repo.ListBySource(ctx, entity.SourceID, entity.Source, 10)
		require.NoError(t, err)
		require.Empty(t, entities)
	})
}

func TestRepository_Normalize(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := NewRepository(db)

		ctx := context.Background()
		entity := ofactest.FindEntity(t, "44525")
		entities := []search.Entity[search.Value]{entity}

		err := repo.Upsert(ctx, "fincen-person", entities)
		require.NoError(t, err)

		found, err := repo.Get(ctx, entity.SourceID, entity.Source)
		require.NoError(t, err)
		require.NotNil(t, found)

		require.NotEmpty(t, found.PreparedFields.NameFields)
	})
}

func TestRepository_Upsert(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := NewRepository(db)

		ctx := context.Background()
		entity := ofactest.FindEntity(t, "44525")
		entity2 := ofactest.FindEntity(t, "48727")
		entities := []search.Entity[search.Value]{entity, entity2}

		err := repo.Upsert(ctx, "fincen-person", entities)
		require.NoError(t, err)

		// Remove entity2 from the upsert list
		entities = []search.Entity[search.Value]{entity}
		entities[0].Name = "john doe"

		err = repo.Upsert(ctx, "fincen-person", entities)
		require.NoError(t, err)

		found, err := repo.Get(ctx, entity.SourceID, entity.Source)
		require.NoError(t, err)
		require.NotNil(t, found)

		require.Equal(t, "john doe", found.Name)
	})
}
