package download_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/download"

	"github.com/stretchr/testify/require"
)

func TestClient_GetFiles_InitialDir(t *testing.T) {
	logger := log.NewTestLogger()

	dl := download.New(logger, nil)
	require.NotNil(t, dl)

	ctx := context.Background()
	dir := filepath.Join("..", "sources", "ofac", "testdata")

	namesAndSources := make(map[string]string)
	namesAndSources["sdn.csv"] = "https://example.com"
	namesAndSources["alt.csv"] = "https://example.com"
	namesAndSources["sdn_comments.csv"] = "https://example.com"

	files, err := dl.GetFiles(ctx, dir, namesAndSources)
	require.NoError(t, err)
	require.Len(t, files, 3)

	require.NotNil(t, files["sdn.csv"])
	require.NotNil(t, files["alt.csv"])
	require.NotNil(t, files["sdn_comments.csv"])
}
