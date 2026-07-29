// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/fshelp"
	"github.com/moov-io/watchman/internal/index"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestGetRefreshInterval(t *testing.T) {
	conf := download.Config{
		RefreshInterval: 2 * time.Minute,
	}
	got := getRefreshInterval(conf)
	require.Equal(t, 2*time.Minute, got)

	t.Setenv("DATA_REFRESH_INTERVAL", "1h")

	got = getRefreshInterval(conf)
	require.Equal(t, 1*time.Hour, got)
}

func TestDownloader_setupPeriodicRefreshing(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	logger := log.NewTestLogger()

	pkg, err := fshelp.FindPkgDir()
	require.NoError(t, err)

	conf := download.Config{
		InitialDataDirectory: filepath.Join(pkg, "ofac", "testdata"),
	}

	dl, err := download.NewDownloader(logger, conf, nil)
	require.NoError(t, err)

	indexedLists := index.NewLists(nil) // only in-memory

	r := download.NewRefresher(ctx, logger, dl, indexedLists, nil)

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancelFunc()
	}()

	errs := make(chan error, 1)
	err = setupPeriodicRefreshing(ctx, logger, errs, conf, r)
	require.NoError(t, err)

	cancelFunc()
	require.NoError(t, <-errs)
}

// scriptedDownloader returns successive results from RefreshAll. After the
// script is exhausted it repeats the last entry.
type scriptedDownloader struct {
	mu    sync.Mutex
	calls int
	steps []struct {
		stats download.Stats
		err   error
	}
}

func (s *scriptedDownloader) RefreshAll(ctx context.Context) (download.Stats, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.calls++
	idx := s.calls - 1
	if idx >= len(s.steps) {
		idx = len(s.steps) - 1
	}
	step := s.steps[idx]
	return step.stats, step.err
}

func (s *scriptedDownloader) callCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.calls
}

func TestDownloader_setupPeriodicRefreshing_scheduledFailureKeepsRunning(t *testing.T) {
	// A failed scheduled refresh must not push onto errs (which would shut down
	// the process). The server should keep the last successful data and retry later.
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	ok := download.Stats{
		Lists:      map[string]int{"us_ofac": 1},
		ListHashes: map[string]string{"us_ofac": "abc"},
	}
	dl := &scriptedDownloader{
		steps: []struct {
			stats download.Stats
			err   error
		}{
			{stats: ok},                      // startup succeeds
			{err: errors.New("origin down")}, // first scheduled tick fails
			{stats: ok},                      // later tick recovers
		},
	}

	conf := download.Config{
		RefreshInterval: 20 * time.Millisecond,
	}
	r := download.NewRefresher(ctx, log.NewTestLogger(), dl, index.NewLists(nil), nil)

	errs := make(chan error, 4)
	err := setupPeriodicRefreshing(ctx, log.NewTestLogger(), errs, conf, r)
	require.NoError(t, err)
	require.Equal(t, download.StateSucceeded, r.Status().State)

	// Wait until the failing scheduled refresh has run.
	require.Eventually(t, func() bool {
		return dl.callCount() >= 2 && r.Status().State == download.StateFailed
	}, 2*time.Second, 5*time.Millisecond)

	require.Contains(t, r.Status().LastError, "origin down")

	// Wait for a successful recovery tick so we know the loop kept going.
	require.Eventually(t, func() bool {
		return dl.callCount() >= 3 && r.Status().State == download.StateSucceeded
	}, 2*time.Second, 5*time.Millisecond)

	// No fatal error should have been reported while we were running.
	select {
	case err := <-errs:
		t.Fatalf("expected no fatal error from scheduled refresh failure, got %v", err)
	default:
	}

	cancelFunc()
	require.NoError(t, <-errs)
}
