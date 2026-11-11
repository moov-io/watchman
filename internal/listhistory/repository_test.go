// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package listhistory

import (
	"context"
	"testing"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestMemoryRepository(t *testing.T) {
	testRepository(t, NewMemoryRepository())
}

func TestSQLRepository(t *testing.T) {
	db.ForEachDatabase(t, func(database db.DB) {
		testRepository(t, NewRepository(database))
	})
}

func testRepository(t *testing.T, repo Repository) {
	t.Helper()
	ctx := context.Background()
	listName := "us_ofac"
	at1 := testTime(1)
	at2 := testTime(2)

	alice := testPerson("1", "Alice")
	bob := testPerson("2", "Bob")
	carol := testPerson("3", "Carol")
	holdings := testBusiness("pct", "100% Holdings")

	_, err := repo.GetRef(ctx, listName, RefLatest)
	require.ErrorIs(t, err, ErrNotFound)

	first, err := repo.ApplyDelta(ctx, ApplyParams{
		ListName:     listName,
		ListHash:     "aaa",
		DownloadedAt: at1,
		EntityCount:  3,
		Delta: Delta{
			Added: []search.Entity[search.Value]{alice, bob, holdings},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "aaa", first.Ref)
	require.Equal(t, 3, first.EntityCount)
	require.Equal(t, 3, first.AddedCount)

	again, err := repo.ApplyDelta(ctx, ApplyParams{
		ListName:     listName,
		ListHash:     "aaa",
		DownloadedAt: at2,
		EntityCount:  99,
		Delta:        Delta{Added: []search.Entity[search.Value]{carol}},
	})
	require.NoError(t, err)
	require.Equal(t, first.ID, again.ID)
	require.Equal(t, 3, again.EntityCount)

	latest, err := repo.GetRef(ctx, listName, "")
	require.NoError(t, err)
	require.Equal(t, "aaa", latest.Ref)

	byHash, err := repo.GetRef(ctx, listName, "aaa")
	require.NoError(t, err)
	require.Equal(t, first.ID, byHash.ID)

	live, err := repo.LiveHashes(ctx, listName)
	require.NoError(t, err)
	require.Len(t, live, 3)
	require.Contains(t, live, "1")
	require.Contains(t, live, "2")
	require.Contains(t, live, "pct")

	alice2 := testPerson("1", "Alice Ann")
	second, err := repo.ApplyDelta(ctx, ApplyParams{
		ListName:     listName,
		ListHash:     "bbb",
		DownloadedAt: at2,
		EntityCount:  3,
		Delta: Delta{
			Added:   []search.Entity[search.Value]{carol},
			Changed: []search.Entity[search.Value]{alice2},
			Removed: []string{"2"},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "bbb", second.Ref)
	require.Equal(t, 1, second.AddedCount)
	require.Equal(t, 1, second.RemovedCount)
	require.Equal(t, 1, second.ChangedCount)

	t.Run("windowed entity versions", func(t *testing.T) {
		atFirst, err := repo.GetEntity(ctx, listName, first.ID, "1")
		require.NoError(t, err)
		require.Equal(t, "Alice", atFirst.Name)

		atSecond, err := repo.GetEntity(ctx, listName, second.ID, "1")
		require.NoError(t, err)
		require.Equal(t, "Alice Ann", atSecond.Name)
		require.NotEmpty(t, atSecond.Entity.PreparedFields.Name)

		_, err = repo.GetEntity(ctx, listName, second.ID, "2")
		require.ErrorIs(t, err, ErrNotFound)

		bobAtFirst, err := repo.GetEntity(ctx, listName, first.ID, "2")
		require.NoError(t, err)
		require.Equal(t, "Bob", bobAtFirst.Name)

		carolAtSecond, err := repo.GetEntity(ctx, listName, second.ID, "3")
		require.NoError(t, err)
		require.Equal(t, "Carol", carolAtSecond.Name)
	})

	t.Run("list entities filter and page", func(t *testing.T) {
		rows, err := repo.ListEntities(ctx, listName, second.ID, EntityFilter{Limit: 1})
		require.NoError(t, err)
		require.Len(t, rows, 2) // limit+1
		require.Equal(t, "1", rows[0].SourceID)

		page2, err := repo.ListEntities(ctx, listName, second.ID, EntityFilter{After: rows[0].SourceID, Limit: 10})
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(page2), 1)
		require.Greater(t, page2[0].SourceID, rows[0].SourceID)

		named, err := repo.ListEntities(ctx, listName, second.ID, EntityFilter{Query: "carol", Limit: 10})
		require.NoError(t, err)
		require.Len(t, named, 1)
		require.Equal(t, "3", named[0].SourceID)

		businesses, err := repo.ListEntities(ctx, listName, second.ID, EntityFilter{
			EntityType: string(search.EntityBusiness),
			Limit:      10,
		})
		require.NoError(t, err)
		require.Len(t, businesses, 1)
		require.Equal(t, "pct", businesses[0].SourceID)

		escaped, err := repo.ListEntities(ctx, listName, second.ID, EntityFilter{Query: "100%", Limit: 10})
		require.NoError(t, err)
		require.Len(t, escaped, 1)
		require.Equal(t, "pct", escaped[0].SourceID)
	})

	t.Run("summaries and by source id", func(t *testing.T) {
		sums, err := repo.ListSummaries(ctx, listName, second.ID)
		require.NoError(t, err)
		require.Len(t, sums, 3)

		found, err := repo.EntitiesBySourceIDs(ctx, listName, second.ID, []string{"1", "3", "missing"})
		require.NoError(t, err)
		require.Len(t, found, 2)
		require.Equal(t, "Alice Ann", found["1"].Name)
		require.Equal(t, "Carol", found["3"].Name)
	})

	t.Run("refs", func(t *testing.T) {
		refs, err := repo.ListRefs(ctx, listName)
		require.NoError(t, err)
		require.Len(t, refs, 2)
		require.Equal(t, "bbb", refs[0].Ref)
		require.Equal(t, "aaa", refs[1].Ref)

		prev, err := repo.PreviousRef(ctx, listName, second.ID)
		require.NoError(t, err)
		require.Equal(t, first.ID, prev.ID)

		_, err = repo.PreviousRef(ctx, listName, first.ID)
		require.ErrorIs(t, err, ErrNotFound)

		latestRefs, err := repo.ListLatestRefs(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, latestRefs)
		var ofac *Ref
		for i := range latestRefs {
			if latestRefs[i].ListName == listName {
				ofac = &latestRefs[i]
			}
		}
		require.NotNil(t, ofac)
		require.Equal(t, "bbb", ofac.Ref)
	})
}
