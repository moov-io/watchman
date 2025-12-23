package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"

	"github.com/stretchr/testify/require"
)

func TestDownloader_Senzing(t *testing.T) {
	logger := log.NewTestLogger()

	location, err := filepath.Abs(filepath.Join("..", "..", "pkg", "sources", "senzing", "testdata", "persons.jsonl"))
	require.NoError(t, err)

	t.Log(location)

	conf := download.Config{
		Senzing: []download.SenzingList{
			{
				SourceList: "senzing-persons",
				Location:   "file://" + location,
			},
		},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)
	require.Len(t, stats.Entities, 3)

	require.Equal(t, 3, stats.Lists["senzing-persons"])
	require.NotEmpty(t, stats.ListHashes["senzing-persons"])
}
