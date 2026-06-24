package refresh

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"

	"github.com/stretchr/testify/require"
)

// fakeDownloader is a controllable download.Downloader for tests. When started is
// non-nil it sends once when RefreshAll begins; when release is non-nil RefreshAll
// blocks until release is closed. This lets tests observe the "running" state.
type fakeDownloader struct {
	stats   download.Stats
	err     error
	started chan struct{}
	release chan struct{}

	mu    sync.Mutex
	calls int
}

func (f *fakeDownloader) RefreshAll(ctx context.Context) (download.Stats, error) {
	f.mu.Lock()
	f.calls++
	f.mu.Unlock()

	if f.started != nil {
		f.started <- struct{}{}
	}
	if f.release != nil {
		<-f.release
	}
	if f.err != nil {
		return download.Stats{}, f.err
	}
	return f.stats, nil
}

func (f *fakeDownloader) callCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}

func okStats() download.Stats {
	return download.Stats{
		Lists:      map[string]int{"us_ofac": 3},
		ListHashes: map[string]string{"us_ofac": "abc"},
	}
}

func newTestManager(t *testing.T, dl download.Downloader) (*manager, index.Lists) {
	t.Helper()

	indexedLists := index.NewLists(nil) // only in-mem
	mgr := NewManager(context.Background(), log.NewTestLogger(), dl, indexedLists, nil).(*manager)
	return mgr, indexedLists
}

func TestManager_RefreshNowSuccess(t *testing.T) {
	dl := &fakeDownloader{stats: okStats()}
	mgr, indexedLists := newTestManager(t, dl)

	require.Equal(t, StateIdle, mgr.Status().State)

	err := mgr.RefreshNow(context.Background(), TriggerStartup)
	require.NoError(t, err)

	st := mgr.Status()
	require.Equal(t, StateSucceeded, st.State)
	require.Equal(t, TriggerStartup, st.Trigger)
	require.NotNil(t, st.StartedAt)
	require.NotNil(t, st.FinishedAt)
	require.NotEmpty(t, st.Duration)
	require.Empty(t, st.LastError)

	// The in-memory index was updated.
	require.Equal(t, map[string]int{"us_ofac": 3}, indexedLists.LatestStats().Lists)
}

func TestManager_RefreshNowFailure(t *testing.T) {
	dl := &fakeDownloader{err: errors.New("boom")}
	mgr, indexedLists := newTestManager(t, dl)

	err := mgr.RefreshNow(context.Background(), TriggerManual)
	require.Error(t, err)

	st := mgr.Status()
	require.Equal(t, StateFailed, st.State)
	require.Equal(t, "boom", st.LastError)
	require.NotNil(t, st.FinishedAt)

	// The in-memory index was not updated on failure.
	require.Empty(t, indexedLists.LatestStats().Lists)

	// A failed refresh must release the slot so a later refresh can recover.
	dl.err = nil
	require.NoError(t, mgr.RefreshNow(context.Background(), TriggerScheduled))
	require.Equal(t, StateSucceeded, mgr.Status().State)
}

func TestManager_ConcurrencyGuard(t *testing.T) {
	dl := &fakeDownloader{
		stats:   okStats(),
		started: make(chan struct{}),
		release: make(chan struct{}),
	}
	mgr, _ := newTestManager(t, dl)

	done := make(chan error, 1)
	go func() {
		done <- mgr.RefreshNow(context.Background(), TriggerScheduled)
	}()

	<-dl.started // a refresh is now in progress

	require.Equal(t, StateRunning, mgr.Status().State)

	// A second synchronous refresh is rejected while one is running.
	require.ErrorIs(t, mgr.RefreshNow(context.Background(), TriggerManual), ErrAlreadyRunning)
	// An async trigger is rejected too.
	require.False(t, mgr.TriggerAsync(TriggerManual))

	close(dl.release)
	require.NoError(t, <-done)

	require.Equal(t, StateSucceeded, mgr.Status().State)
	require.Equal(t, 1, dl.callCount()) // only one refresh actually ran
}

func TestManager_TriggerAsync(t *testing.T) {
	dl := &fakeDownloader{stats: okStats()}
	mgr, _ := newTestManager(t, dl)

	require.True(t, mgr.TriggerAsync(TriggerManual))

	require.Eventually(t, func() bool {
		return mgr.Status().State == StateSucceeded
	}, 2*time.Second, 5*time.Millisecond)

	require.Equal(t, TriggerManual, mgr.Status().Trigger)
}
