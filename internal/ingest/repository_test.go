package ingest

import (
	"context"
	"fmt"
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

		found, err := repo.Get(ctx, entity.SourceID, "fincen-person")
		require.NoError(t, err)
		require.NotNil(t, found)

		// Empty SourceData // TODO(adam): can we unmarshal?
		entity.SourceData = make(map[string]interface{})
		found.SourceData = make(map[string]interface{})
		entity.Source = "fincen-person"

		// Compare objects
		require.Equal(t, entity.Normalize(), found.Normalize())

		// List
		entities, err = repo.ListBySource(ctx, "", "fincen-person", 10)
		require.NoError(t, err)
		require.Len(t, entities, 1)
		require.Equal(t, entity.SourceID, entities[0].SourceID)

		// List again
		entities, err = repo.ListBySource(ctx, "12345", "fincen-person", 10)
		require.NoError(t, err)
		require.Len(t, entities, 1)
		require.Equal(t, entity.SourceID, entities[0].SourceID)

		// Find nothing
		entities, err = repo.ListBySource(ctx, entity.SourceID, "fincen-person", 10)
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

		found, err := repo.Get(ctx, entity.SourceID, "fincen-person")
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

		found, err := repo.Get(ctx, entity.SourceID, "fincen-person")
		require.NoError(t, err)
		require.NotNil(t, found)

		require.Equal(t, "john doe", found.Name)
	})
}

func TestRepository_ChecksumsAndListAll(t *testing.T) {
	db.ForEachDatabase(t, func(db db.DB) {
		repo := NewRepository(db)
		ctx := context.Background()

		var entities []search.Entity[search.Value]
		for i := 0; i < 1205; i++ {
			entities = append(entities, search.Entity[search.Value]{
				Name:     fmt.Sprintf("person %05d", i),
				Type:     search.EntityPerson,
				SourceID: fmt.Sprintf("%05d", i),
			})
		}

		err := repo.Upsert(ctx, "fincen-person", entities)
		require.NoError(t, err)

		sums, err := repo.Checksums(ctx)
		require.NoError(t, err)
		require.Len(t, sums, 1)
		require.Equal(t, "fincen-person", sums[0].Source)
		require.Equal(t, 1205, sums[0].EntityCount)
		require.Len(t, sums[0].Checksum, 64)

		all, err := repo.ListAll(ctx)
		require.NoError(t, err)
		require.Len(t, all, 1205)
		require.Equal(t, "00000", all[0].SourceID)
		require.Equal(t, "01204", all[len(all)-1].SourceID)

		page, err := repo.ListBySource(ctx, "", "fincen-person", 1000)
		require.NoError(t, err)
		require.Len(t, page, 1000)

		page2, err := repo.ListBySource(ctx, page[len(page)-1].SourceID, "fincen-person", 1000)
		require.NoError(t, err)
		require.Len(t, page2, 205)

		err = repo.Upsert(ctx, "fincen-person", nil)
		require.NoError(t, err)
		sums, err = repo.Checksums(ctx)
		require.NoError(t, err)
		require.Empty(t, sums)
	})
}
