package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestDownloader_RefreshAll_USTel(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSTEL,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	// Verify entities were loaded
	require.Greater(t, len(stats.Entities), 0, "expected at least some US TEL entities")

	// Verify list tracking
	name := string(search.SourceUSTEL)
	require.Greater(t, stats.Lists[name], 0)
	require.NotEmpty(t, stats.ListHashes[name], "expected list hash to be computed")

	// Check a few entities
	for _, entity := range stats.Entities {
		require.Equal(t, search.SourceUSTEL, entity.Source, "entity should have US TEL source")
		require.NotEmpty(t, entity.SourceID, "entity should have source ID")
		require.NotEmpty(t, entity.Name, "entity should have name")
		require.NotNil(t, entity.SourceData, "entity should have SourceData")
	}
}
