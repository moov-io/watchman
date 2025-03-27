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

func TestDownloader_RefreshAll_InitialDir(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
	}

	dl, err := download.NewDownloader(logger, conf)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	require.Greater(t, len(stats.Entities), 100)

	name := string(search.SourceUSOFAC)
	require.Greater(t, stats.Lists[name], 100)
	require.NotEmpty(t, stats.ListHashes[name])
}
