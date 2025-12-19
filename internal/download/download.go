package download

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"golang.org/x/sync/errgroup"
)

type Downloader interface {
	RefreshAll(ctx context.Context) (Stats, error)
}

func NewDownloader(logger log.Logger, conf Config) (Downloader, error) {
	return &downloader{
		logger: logger,
		conf:   conf,
	}, nil
}

type downloader struct {
	logger log.Logger
	conf   Config
}

func (dl *downloader) RefreshAll(ctx context.Context) (Stats, error) {
	ctx, span := telemetry.StartSpan(ctx, "refresh-all")
	defer span.End()

	stats := Stats{
		Lists:      make(map[string]int),
		ListHashes: make(map[string]string),
		StartedAt:  time.Now().In(time.UTC),
	}
	logger := dl.logger.Info().With(log.Fields{
		"initial_data_directory": log.String(expandInitialDir(initialDataDirectory(dl.conf))),
	})
	start := time.Now()
	logger.Info().Log("starting list refresh")

	g, ctx := errgroup.WithContext(ctx)
	preparedLists := make(chan preparedList, 10)

	// Start a goroutine to accumulate results
	resultsDone := make(chan struct{})
	go func() {
		defer close(resultsDone)

		for list := range preparedLists {
			logger.Info().Logf("adding %d entities from %v", len(list.Entities), list.ListName)

			stats.Lists[string(list.ListName)] = len(list.Entities)
			stats.ListHashes[string(list.ListName)] = list.Hash

			stats.Entities = append(stats.Entities, list.Entities...)
		}
	}()

	// Create a WaitGroup to track all producers
	var producerWg sync.WaitGroup

	// Track what lists have been requested
	var requestedLists []search.SourceList

	// OFAC Records
	if slices.Contains(dl.conf.IncludedLists, search.SourceUSOFAC) {
		requestedLists = append(requestedLists, search.SourceUSOFAC)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()
			err := loadOFACRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				return fmt.Errorf("loading OFAC records: %w", err)
			}
			return nil
		})
	}

	// OFAC Non-SDN Records
	if slices.Contains(dl.conf.IncludedLists, search.SourceUSNonSDN) {
		requestedLists = append(requestedLists, search.SourceUSNonSDN)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()
			err := loadUSNonSDNRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				return fmt.Errorf("loading OFAC Non-SDN records: %w", err)
			}
			return nil
		})
	}

	// CSL Records
	if slices.Contains(dl.conf.IncludedLists, search.SourceUSCSL) {
		requestedLists = append(requestedLists, search.SourceUSCSL)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()
			err := loadCSLUSRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				return fmt.Errorf("loading US CSL records: %w", err)
			}
			return nil
		})
	}

	// Compare the configured lists against those we actually loaded.
	// Any extra lists are an error as we don't want to silently ignore them.
	if len(requestedLists) > len(dl.conf.IncludedLists) {
		close(preparedLists)

		return stats, fmt.Errorf("loaded extra lists: %#v loaded compared to %#v configured", requestedLists, dl.conf.IncludedLists)
	}
	if extra := findExtraLists(dl.conf.IncludedLists, requestedLists); extra != "" {
		close(preparedLists)

		return stats, fmt.Errorf("unknown lists: %v", extra)
	}

	// Add a goroutine to close the channel when all producers are done
	g.Go(func() error {
		producerWg.Wait()
		close(preparedLists)
		return nil
	})

	// Wait for both producers and consumer to finish
	err := g.Wait()
	<-resultsDone

	if err != nil {
		return stats, fmt.Errorf("problem loading lists: %v", err)
	}

	logger.Info().Logf("finished all lists: %v", time.Since(start))

	stats.EndedAt = time.Now().In(time.UTC)

	return stats, nil
}

func findExtraLists(config, loaded []search.SourceList) string {
	var extra []search.SourceList

	for _, c := range config {
		var found bool
		for _, l := range loaded {
			if c == l {
				found = true
				break
			}
		}
		if !found {
			extra = append(extra, c)
		}
	}

	var buf bytes.Buffer
	for idx, e := range extra {
		if idx > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(string(e))
	}
	return buf.String()
}

type preparedList struct {
	ListName search.SourceList
	Entities []search.Entity[search.Value]

	Hash string
}

func expandInitialDir(initialDir string) string {
	dir, err := filepath.Abs(initialDir)
	if err != nil {
		dir = initialDir
	}
	return dir
}

func initialDataDirectory(conf Config) string {
	return cmp.Or(os.Getenv("INITIAL_DATA_DIRECTORY"), conf.InitialDataDirectory)
}
