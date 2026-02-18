// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloader_RefreshAll_EUCSL(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceEUCSL,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	// Should have entities
	require.Greater(t, len(stats.Entities), 100, "expected at least 100 EU CSL entities")

	// Check list stats
	name := string(search.SourceEUCSL)
	require.Greater(t, stats.Lists[name], 100, "expected EU CSL list to have over 100 entries")
	require.NotEmpty(t, stats.ListHashes[name], "expected EU CSL list hash")

	// Verify entity types
	var persons, businesses int
	for _, entity := range stats.Entities {
		switch entity.Type {
		case search.EntityPerson:
			persons++
		case search.EntityBusiness:
			businesses++
		}
	}

	assert.Greater(t, persons, 0, "expected at least some person entities")
	assert.Greater(t, businesses, 0, "expected at least some business entities")

	t.Logf("EU CSL: %d total entities (%d persons, %d businesses)", len(stats.Entities), persons, businesses)
}

func TestDownloader_RefreshAll_EUCSL_WithOFAC(t *testing.T) {
	logger := log.NewTestLogger()
	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
			search.SourceEUCSL,
		},
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)
	require.NotNil(t, dl)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err)
	require.NotNil(t, stats)

	// Should have entities from both lists
	require.Greater(t, len(stats.Entities), 200, "expected entities from both lists")

	// Check both lists are present
	require.Greater(t, stats.Lists[string(search.SourceUSOFAC)], 0, "expected OFAC entries")
	require.Greater(t, stats.Lists[string(search.SourceEUCSL)], 0, "expected EU CSL entries")

	t.Logf("Combined: OFAC=%d, EU CSL=%d, Total=%d",
		stats.Lists[string(search.SourceUSOFAC)],
		stats.Lists[string(search.SourceEUCSL)],
		len(stats.Entities))
}
