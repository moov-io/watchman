package download_test

import (
	"context"
	"os"
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

// badSenzingFile creates a temp file containing invalid JSON and returns its
// path.  The caller is responsible for cleanup (deferred via t.Cleanup).
func badSenzingFile(t *testing.T) string {
	t.Helper()
	f, err := os.CreateTemp("", "bad-senzing-*.jsonl")
	require.NoError(t, err)
	_, err = f.WriteString("{invalid json content")
	require.NoError(t, err)
	require.NoError(t, f.Close())
	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name()
}

// emptySenzingFile creates an empty temp .jsonl file and returns its path.
// The caller is responsible for cleanup (deferred via t.Cleanup).
func emptySenzingFile(t *testing.T) string {
	t.Helper()
	f, err := os.CreateTemp("", "empty-senzing-*.jsonl")
	require.NoError(t, err)
	require.NoError(t, f.Close())
	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name()
}

func TestDownloader_RefreshAll_IgnoresConfiguredDownloadError(t *testing.T) {
	logger := log.NewTestLogger()
	badFile := badSenzingFile(t)

	const listName search.SourceList = "senzing-bad-ignored"

	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		ErrorOnEmptyList:     true,
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
		Senzing: []download.SenzingList{
			{SourceList: listName, Location: "file://" + badFile},
		},
		IgnoredDownloadErrors: []search.SourceList{listName},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err, "error for configured IgnoredDownloadErrors list should be suppressed")
	require.NotNil(t, stats)

	// The failing list must not appear in stats.
	require.Empty(t, stats.Lists["senzing-bad-ignored"])

	// The successful OFAC list must still be loaded.
	ofacName := string(search.SourceUSOFAC)
	require.Greater(t, stats.Lists[ofacName], 100, "OFAC entities should be present")
	require.NotEmpty(t, stats.ListHashes[ofacName], "OFAC hash should be set")
	require.Greater(t, len(stats.Entities), 100, "total entities should include OFAC records")
}

func TestDownloader_RefreshAll_DoesNotIgnoreUnconfiguredDownloadError(t *testing.T) {
	logger := log.NewTestLogger()
	badFile := badSenzingFile(t)

	const listName search.SourceList = "senzing-bad-not-ignored"

	conf := download.Config{
		Senzing: []download.SenzingList{
			{SourceList: listName, Location: "file://" + badFile},
		},
		// IgnoredDownloadErrors intentionally empty
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	_, err = dl.RefreshAll(context.Background())
	require.Error(t, err, "error for list not in IgnoredDownloadErrors must propagate")
}

// TestDownloader_RefreshAll_IgnoredDownloadError_MixedBatch verifies that when
// a Senzing batch contains both an ignored and a non-ignored list, a download
// failure for the ignored list does not prevent the non-ignored list from loading.
func TestDownloader_RefreshAll_IgnoredDownloadError_MixedBatch(t *testing.T) {
	logger := log.NewTestLogger()

	goodFile, err := filepath.Abs(filepath.Join("..", "..", "pkg", "sources", "senzing", "testdata", "persons.jsonl"))
	require.NoError(t, err)

	const ignoredList search.SourceList = "senzing-missing-ignored"
	const goodList search.SourceList = "senzing-good"

	conf := download.Config{
		Senzing: []download.SenzingList{
			{SourceList: ignoredList, Location: "file:///nonexistent/path.jsonl"},
			{SourceList: goodList, Location: "file://" + goodFile},
		},
		IgnoredDownloadErrors: []search.SourceList{ignoredList},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err, "ignored list download failure must not fail the refresh")
	require.NotNil(t, stats)

	// The ignored list must not appear.
	require.Empty(t, stats.Lists[string(ignoredList)])

	// The non-ignored list must have loaded successfully.
	require.Equal(t, 3, stats.Lists[string(goodList)], "non-ignored list should load")
	require.NotEmpty(t, stats.ListHashes[string(goodList)])
}

func TestDownloader_RefreshAll_UnknownIgnoredDownloadErrors(t *testing.T) {
	logger := log.NewTestLogger()

	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
		// "completely-unknown" is neither a standard list nor a configured
		// OpenSanctions/Senzing list, so RefreshAll must reject it.
		IgnoredDownloadErrors: []search.SourceList{"completely-unknown"},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	_, err = dl.RefreshAll(context.Background())
	require.Error(t, err, "unknown list in IgnoredDownloadErrors must be rejected")
	require.Contains(t, err.Error(), "completely-unknown")
}

// TestDownloader_RefreshAll_IgnoredDownloadError_EmptySenzingList verifies that
// an empty Senzing-format list with ErrorOnEmptyList=true is ignored when the
// list is listed in IgnoredDownloadErrors.
func TestDownloader_RefreshAll_IgnoredDownloadError_EmptySenzingList(t *testing.T) {
	logger := log.NewTestLogger()
	emptyFile := emptySenzingFile(t)

	const listName search.SourceList = "senzing-empty-ignored"

	conf := download.Config{
		ErrorOnEmptyList: true,
		Senzing: []download.SenzingList{
			{SourceList: listName, Location: "file://" + emptyFile},
		},
		IgnoredDownloadErrors: []search.SourceList{listName},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err, "empty list error for ignored list should be suppressed")
	require.NotNil(t, stats)

	// The ignored empty list must not appear in stats.
	require.Empty(t, stats.Lists[string(listName)])
}

// TestDownloader_RefreshAll_DoesNotIgnoreUnconfiguredEmptySenzingList verifies
// that an empty Senzing list with ErrorOnEmptyList=true is fatal when the list
// is not listed in IgnoredDownloadErrors.
func TestDownloader_RefreshAll_DoesNotIgnoreUnconfiguredEmptySenzingList(t *testing.T) {
	logger := log.NewTestLogger()
	emptyFile := emptySenzingFile(t)

	const listName search.SourceList = "senzing-empty-not-ignored"

	conf := download.Config{
		ErrorOnEmptyList: true,
		Senzing: []download.SenzingList{
			{SourceList: listName, Location: "file://" + emptyFile},
		},
		// IgnoredDownloadErrors intentionally empty for this list
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	_, err = dl.RefreshAll(context.Background())
	require.Error(t, err, "empty list error for non-ignored list must propagate")
}

// TestDownloader_RefreshAll_IgnoredDownloadError_CustomListCasing validates that
// a custom (Senzing) list name declared with unusual casing/whitespace can still
// be successfully referenced in IgnoredDownloadErrors using different casing.
// This exercises the normalization in validateIgnoredDownloadErrors (and the
// related get/processing paths) so that validation does not spuriously reject
// the entry and the ignore behavior works at runtime.
func TestDownloader_RefreshAll_IgnoredDownloadError_CustomListCasing(t *testing.T) {
	logger := log.NewTestLogger()
	badFile := badSenzingFile(t)

	const declared search.SourceList = "Senzing-Casing-Test"
	const inIgnored search.SourceList = " senzing-casing-test " // different case + surrounding spaces

	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "test", "testdata"),
		IncludedLists: []search.SourceList{
			search.SourceUSOFAC,
		},
		Senzing: []download.SenzingList{
			{SourceList: declared, Location: "file://" + badFile},
		},
		IgnoredDownloadErrors: []search.SourceList{inIgnored},
	}
	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	stats, err := dl.RefreshAll(context.Background())
	require.NoError(t, err, "casing mismatch between custom list declaration and IgnoredDownloadErrors must not cause validation failure or load error")
	require.NotNil(t, stats)

	// The ignored custom list (under its normalized name) must not appear.
	require.Empty(t, stats.Lists["senzing-casing-test"])

	// A non-ignored standard list must still have loaded successfully.
	ofacName := string(search.SourceUSOFAC)
	require.Greater(t, stats.Lists[ofacName], 100, "OFAC entities should be present")
}
