// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package opensanctions

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping download test in short mode (file is ~670MB)")
	}

	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	require.Len(t, files, 1)

	file, found := files["peps_senzing.json"]
	require.True(t, found)
	require.NotNil(t, file)

	// Read first 1KB to verify it's valid JSON
	// Note: Download() returns a streaming response, so only the bytes we read
	// are actually downloaded. Full file (~670MB) is only downloaded when fully read.
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	require.NoError(t, err)
	require.Greater(t, n, 0)

	// Should start with '[' (JSON array) or '{' (JSON object)
	content := string(buf[:n])
	require.True(t, strings.HasPrefix(strings.TrimSpace(content), "[") ||
		strings.HasPrefix(strings.TrimSpace(content), "{"),
		"expected JSON content, got: %s", content[:min(100, len(content))])

	require.NoError(t, file.Close())
}

// TestDownloadAndParse downloads and parses the full OpenSanctions PEP file.
// This test is slow (~670MB download) and should only be run manually.
func TestDownloadAndParse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping full download test in short mode (file is ~670MB)")
	}
	if os.Getenv("TEST_OPENSANCTIONS_FULL") == "" {
		t.Skip("set TEST_OPENSANCTIONS_FULL=1 to run this test")
	}

	files, err := Download(context.Background(), log.NewNopLogger(), "")
	require.NoError(t, err)
	defer files.Close()

	results, err := Read(files)
	require.NoError(t, err)
	require.NotNil(t, results)

	// OpenSanctions PEP dataset should have ~2M entities
	t.Logf("Downloaded and parsed %d entities", len(results.Entities))
	require.Greater(t, len(results.Entities), 100000, "expected at least 100k entities")

	// Verify hash was computed
	require.NotEmpty(t, results.ListHash)
	t.Logf("List hash: %s", results.ListHash)
}

func TestDownload_initialDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "initial-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	mk := func(t *testing.T, name string, body string) {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(body), 0600); err != nil {
			t.Fatalf("writing %s: %v", path, err)
		}
	}

	// create the file locally
	mk(t, "peps_senzing.json", `[{"DATA_SOURCE":"TEST","RECORD_ID":"1"}]`)

	files, err := Download(context.Background(), log.NewNopLogger(), dir)
	require.NoError(t, err)
	require.Len(t, files, 1)

	for filename, fd := range files {
		if strings.EqualFold("peps_senzing.json", filepath.Base(filename)) {
			bs, err := io.ReadAll(fd)
			require.NoError(t, err)

			require.Equal(t, `[{"DATA_SOURCE":"TEST","RECORD_ID":"1"}]`, string(bs))
		} else {
			t.Fatalf("unknown file: %v", filename)
		}
	}
}
