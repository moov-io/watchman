package webui

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
)

func TestStaticHandlerGzipWASM(t *testing.T) {
	payload := bytesRepeat("watchman-ui-wasm", 40)
	handler := newStaticHandler(fstest.MapFS{
		"ui.wasm":    &fstest.MapFile{Data: payload},
		"index.html": &fstest.MapFile{Data: []byte("<html>Watchman</html>")},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/ui.wasm", nil)
	req.Header.Set("Accept-Encoding", "deflate, gzip;q=1.0")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "gzip", rec.Header().Get("Content-Encoding"))
	require.Equal(t, "application/wasm", rec.Header().Get("Content-Type"))
	require.Contains(t, rec.Header().Get("Vary"), "Accept-Encoding")
	require.Contains(t, rec.Header().Get("Cache-Control"), "max-age=3600")
	encoded := append([]byte(nil), rec.Body.Bytes()...)
	require.Less(t, len(encoded), len(payload))

	reader, err := gzip.NewReader(bytes.NewReader(encoded))
	require.NoError(t, err)
	got, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.Equal(t, payload, got)

	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req)
	require.Equal(t, encoded, rec2.Body.Bytes())
}

func TestStaticHandlerRawWASMAndHTML(t *testing.T) {
	payload := []byte("raw-wasm-bytes")
	handler := newStaticHandler(fstest.MapFS{
		"app.wasm":   &fstest.MapFile{Data: payload},
		"index.html": &fstest.MapFile{Data: []byte("hello ui")},
	}, nil)

	raw := httptest.NewRecorder()
	handler.ServeHTTP(raw, httptest.NewRequest(http.MethodGet, "/app.wasm", nil))
	require.Equal(t, http.StatusOK, raw.Code)
	require.Equal(t, payload, raw.Body.Bytes())
	require.Empty(t, raw.Header().Get("Content-Encoding"))

	page := httptest.NewRecorder()
	handler.ServeHTTP(page, httptest.NewRequest(http.MethodGet, "/", nil))
	require.Equal(t, http.StatusOK, page.Code)
	require.Contains(t, page.Body.String(), "hello ui")
}

func TestGzipMissingWASMFallsThrough(t *testing.T) {
	handler := newStaticHandler(fstest.MapFS{
		"index.html": &fstest.MapFile{Data: []byte("only html")},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/missing.wasm", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusNotFound, rec.Code)
	require.Empty(t, rec.Header().Get("Content-Encoding"))
}

func bytesRepeat(part string, n int) []byte {
	out := make([]byte, 0, len(part)*n)
	for i := 0; i < n; i++ {
		out = append(out, part...)
	}
	return out
}
