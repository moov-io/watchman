package cslustest

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func FindEntity(tb testing.TB, entityID string) search.Entity[search.Value] {
	tb.Helper()

	logger := log.NewTestLogger()

	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "pkg", "csl_us", "testdata"),
	}
	conf.IncludedLists = append(conf.IncludedLists, search.SourceUSCSL)

	dl, err := download.NewDownloader(logger, conf)
	require.NoError(tb, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(tb, err)

	for _, entity := range stats.Entities {
		if entity.SourceID == entityID && entity.Source == search.SourceUSCSL {
			return entity
		}
	}

	tb.Fatalf("US CSL entity %s not found", entityID)

	return search.Entity[search.Value]{}
}