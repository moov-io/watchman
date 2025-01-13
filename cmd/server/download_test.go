// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/search"

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

	conf := download.Config{
		InitialDataDirectory: filepath.Join("..", "..", "pkg", "ofac", "testdata"),
	}

	dl, err := download.NewDownloader(logger, conf)
	require.NoError(t, err)

	searchService := search.NewService(logger)

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancelFunc()
	}()

	errs := make(chan error, 1)
	err = setupPeriodicRefreshing(ctx, logger, errs, conf, dl, searchService)
	require.NoError(t, err)

	cancelFunc()
	require.NoError(t, <-errs)
}
