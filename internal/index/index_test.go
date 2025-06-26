package index_test

import (
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/db"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/internal/ingest"

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
