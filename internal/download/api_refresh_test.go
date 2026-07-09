package download

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/base/log"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func newTestRefreshController(t *testing.T, dl Downloader) (*mux.Router, *Refresher) {
	t.Helper()

	r := NewRefresher(context.Background(), log.NewTestLogger(), dl, nil, nil)
	ctrl := NewRefreshController(log.NewTestLogger(), r)

	router := mux.NewRouter()
	ctrl.AppendRoutes(router)
	return router, r
}

func decodeStatus(t *testing.T, w *httptest.ResponseRecorder) Status {
	t.Helper()
	var st Status
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &st))
	return st
}

func TestRefreshAPI_StatusIdle(t *testing.T) {
	router, _ := newTestRefreshController(t, &fakeDownloader{stats: okStats()})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("GET", "/v2/data/refresh", nil))

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, StateIdle, decodeStatus(t, w).State)
}

func TestRefreshAPI_RefreshAsync(t *testing.T) {
	dl := &fakeDownloader{stats: okStats(), started: make(chan struct{}), release: make(chan struct{})}
	router, r := newTestRefreshController(t, dl)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusAccepted, w.Code)
	require.Equal(t, StateRunning, decodeStatus(t, w).State)

	<-dl.started

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusConflict, w2.Code)
	require.Equal(t, StateRunning, decodeStatus(t, w2).State)

	close(dl.release)
	require.Eventually(t, func() bool {
		return r.Status().State == StateSucceeded
	}, 2*time.Second, 5*time.Millisecond)
}

func TestRefreshAPI_AlwaysEnabled(t *testing.T) {
	router, _ := newTestRefreshController(t, &fakeDownloader{stats: okStats()})

	// POST is always registered now
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest("POST", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusAccepted, w.Code)

	// GET status always available
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, httptest.NewRequest("GET", "/v2/data/refresh", nil))
	require.Equal(t, http.StatusOK, w2.Code)
}
