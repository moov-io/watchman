package download_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestClient_DownloadFile(t *testing.T) {
	logger := log.NewTestLogger()

	dl := download.New(logger, nil)
	require.NotNil(t, dl)

	namesAndSources := map[string]string{
		"client": "file://client.go",
		"test":   "file://client_test.go",
	}

	files, err := dl.GetFiles(context.Background(), "", namesAndSources)
	require.NoError(t, err)
	require.Len(t, files, 2)

	require.NotNil(t, files["client"])
	require.NotNil(t, files["test"])

	require.NoError(t, files["client"].Close())
	require.NoError(t, files["test"].Close())
}

func TestClient_WithAdditionalHeaders(t *testing.T) {
	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	dl := download.New(log.NewTestLogger(), server.Client(),
		download.WithAdditionalHeaders(map[string]string{"Authorization": "Token abc123"}))

	files, err := dl.GetFiles(context.Background(), "", map[string]string{
		"list.json": server.URL + "/list.json",
	})
	require.NoError(t, err)
	require.Len(t, files, 1)
	t.Cleanup(func() { files.Close() })

	body, err := io.ReadAll(files["list.json"])
	require.NoError(t, err)
	require.Equal(t, "ok", string(body))
	require.Equal(t, "Token abc123", gotAuth)
}
