package refresh

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func newTestController(t *testing.T, dl download.Downloader, allowManual bool) (*mux.Router, Manager) {
	t.Helper()

	indexedLists := index.NewLists(nil)
	mgr := NewManager(context.Background(), log.NewTestLogger(), dl, indexedLists, nil)
	controller := NewController(log.NewTestLogger(), mgr, allowManual)

	router := mux.NewRouter()
	controller.AppendRoutes(router)
	return router, mgr
}

func decodeStatus(t *testing.T, w *httptest.ResponseRecorder) Status {
	t.Helper()

	var st Status
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &st))
	return st
}

func TestAPI_StatusIdle(t *testing.T) {
	router, _ := newTestController(t, &fakeDownloader{stats: okStats()}, true)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", "/v2/data/refresh", nil))

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, StateIdle, decodeStatus(t, w).State)
}

func TestAPI_RefreshAsync(t *testing.T) {
	dl := &fakeDownloader{stats: okStats(), started: make(chan struct{}), release: make(chan struct{})}
	router, mgr := newTestController(t, dl, true)

	// Kick off an async refresh.
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusAccepted, w.Code)
	require.Equal(t, StateRunning, decodeStatus(t, w).State)

	<-dl.started // the refresh has begun and is blocked on release

	// A second POST while one is running returns 409 Conflict.
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusConflict, w2.Code)
	require.Equal(t, StateRunning, decodeStatus(t, w2).State)

	close(dl.release)
	require.Eventually(t, func() bool {
		return mgr.Status().State == StateSucceeded
	}, 2*time.Second, 5*time.Millisecond)
}

func TestAPI_RefreshDisabled(t *testing.T) {
	router, _ := newTestController(t, &fakeDownloader{stats: okStats()}, false)

	// POST is not registered when disabled -> 405 (the path exists for GET).
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusMethodNotAllowed, w.Code)

	// GET status remains available.
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, httptest.NewRequest("GET", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusOK, w2.Code)
}
