package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/geocoding/geocodetest"
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

	dl, err := download.NewDownloader(logger, conf, nil)
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

func TestDownloader_Geocode(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
	}

	g := &geocodetest.RandomGeocoder{}

	dl, err := download.NewDownloader(logger, conf, g)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)
	require.Greater(t, len(stats.Entities), 100)

	for _, entity := range stats.Entities {
		for _, addr := range entity.Addresses {
			require.NotEmpty(t, addr.Latitude)
			require.NotEmpty(t, addr.Longitude)
		}
	}
}
