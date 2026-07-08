// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"cmp"
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
)

func setupPeriodicRefreshing(ctx context.Context, logger log.Logger, errs chan error, conf download.Config, r *download.Refresher) error {
	// Initial, blocking load. A failure here is fatal to startup.
	if err := r.RefreshNow(ctx, download.TriggerStartup); err != nil {
		return err
	}

	// Setup periodic refreshing
	refreshEvery := getRefreshInterval(conf)
	ticker := time.NewTicker(refreshEvery)
	logger.Info().Logf("refreshing data every %v", refreshEvery)

	go func() {
		defer ticker.Stop() // stop ticker only when we're shutdown

		for {
			select {
			case <-ctx.Done():
				errs <- nil
				return

			case <-ticker.C:
				// A manual refresh may be in progress; skip this tick rather than fail.
				err := r.RefreshNow(ctx, download.TriggerScheduled)
				if err != nil && !errors.Is(err, download.ErrAlreadyRunning) {
					errs <- err
				}
			}
		}
	}()

	return nil
}

const (
	defaultRefreshInterval = 12 * time.Hour
)

func getRefreshInterval(conf download.Config) time.Duration {
	override := strings.TrimSpace(os.Getenv("DATA_REFRESH_INTERVAL"))
	if override != "" {
		dur, err := time.ParseDuration(override)
		if err == nil {
			return dur
		}
	}
	return cmp.Or(conf.RefreshInterval, defaultRefreshInterval)
}
