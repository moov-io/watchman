package download

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

// fakeDownloader implements Downloader for tests.
type fakeDownloader struct {
	stats   Stats
	err     error
	started chan struct{}
	release chan struct{}

	mu    sync.Mutex
	calls int
}

func (f *fakeDownloader) RefreshAll(ctx context.Context) (Stats, error) {
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
		return Stats{}, f.err
	}
	return f.stats, nil
}

func (f *fakeDownloader) callCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}

func okStats() Stats {
	return Stats{
		Lists:      map[string]int{"us_ofac": 3},
		ListHashes: map[string]string{"us_ofac": "abc"},
	}
}

// testStatsUpdater is a minimal statsUpdater for tests that records the last update.
type testStatsUpdater struct {
	last Stats
}

func (t *testStatsUpdater) Update(s Stats) { t.last = s }

func newTestRefresher(t *testing.T, dl Downloader) (*Refresher, *testStatsUpdater) {
	t.Helper()
	up := &testStatsUpdater{}
	r := NewRefresher(context.Background(), log.NewTestLogger(), dl, up, nil)
	return r, up
}

func TestRefresher_RefreshNowSuccess(t *testing.T) {
	dl := &fakeDownloader{stats: okStats()}
	r, up := newTestRefresher(t, dl)

	require.Equal(t, StateIdle, r.Status().State)

	err := r.RefreshNow(context.Background(), TriggerStartup)
	require.NoError(t, err)

	st := r.Status()
	require.Equal(t, StateSucceeded, st.State)
	require.Equal(t, TriggerStartup, st.Trigger)
	require.NotNil(t, st.StartedAt)
	require.NotNil(t, st.FinishedAt)
	require.NotEmpty(t, st.Duration)
	require.Empty(t, st.LastError)

	require.Equal(t, map[string]int{"us_ofac": 3}, up.last.Lists)
}

func TestRefresher_RefreshNowFailure(t *testing.T) {
	dl := &fakeDownloader{err: errors.New("boom")}
	r, up := newTestRefresher(t, dl)

	err := r.RefreshNow(context.Background(), TriggerManual)
	require.Error(t, err)

	st := r.Status()
	require.Equal(t, StateFailed, st.State)
	require.Equal(t, "boom", st.LastError)
	require.NotNil(t, st.FinishedAt)

	require.Empty(t, up.last.Lists)

	// Slot must be released
	dl.err = nil
	require.NoError(t, r.RefreshNow(context.Background(), TriggerScheduled))
	require.Equal(t, StateSucceeded, r.Status().State)
}

func TestRefresher_ConcurrencyGuard(t *testing.T) {
	dl := &fakeDownloader{
		stats:   okStats(),
		started: make(chan struct{}),
		release: make(chan struct{}),
	}
	r, _ := newTestRefresher(t, dl)

	done := make(chan error, 1)
	go func() {
		done <- r.RefreshNow(context.Background(), TriggerScheduled)
	}()

	<-dl.started

	require.Equal(t, StateRunning, r.Status().State)

	require.ErrorIs(t, r.RefreshNow(context.Background(), TriggerManual), ErrAlreadyRunning)
	require.False(t, r.TriggerAsync(TriggerManual))

	close(dl.release)
	require.NoError(t, <-done)

	require.Equal(t, StateSucceeded, r.Status().State)
	require.Equal(t, 1, dl.callCount())
}

func TestRefresher_TriggerAsync(t *testing.T) {
	dl := &fakeDownloader{stats: okStats()}
	r, _ := newTestRefresher(t, dl)

	require.True(t, r.TriggerAsync(TriggerManual))

	require.Eventually(t, func() bool {
		return r.Status().State == StateSucceeded
	}, 2*time.Second, 5*time.Millisecond)

	require.Equal(t, TriggerManual, r.Status().Trigger)
}
