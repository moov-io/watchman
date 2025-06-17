// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"cmp"
	"context"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/index"
)

func setupPeriodicRefreshing(ctx context.Context, logger log.Logger, errs chan error, conf download.Config, downloader download.Downloader, indexedLists index.Lists) error {
	err := refreshAllSources(logger, downloader, indexedLists)
	if err != nil {
		return err
	}

	// Setup periodic refreshing
	ticker := time.NewTicker(getRefreshInterval(conf))
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				errs <- nil
				return

			case <-ticker.C:
				err := refreshAllSources(logger, downloader, indexedLists)
				if err != nil {
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

func refreshAllSources(logger log.Logger, downloader download.Downloader, indexedLists index.Lists) error {
	ctx, span := telemetry.StartSpan(context.Background(), "refresh-all-sources")
	defer span.End()

	stats, err := downloader.RefreshAll(ctx)
	if err != nil {
		return err
	}

	logger.Info().Logf("data refreshed - %v entities from %v lists took %v (using %.2fGB)",
		len(stats.Entities), len(stats.Lists), stats.EndedAt.Sub(stats.StartedAt), getCurrentMemoryUsed())

	// Replace in-mem entities
	indexedLists.Update(stats)

	return nil
}

func getCurrentMemoryUsed() float64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return float64(mem.Alloc) / 1024.0 / 1024.0 / 1024.0 // divide by 1GB
}
