package download

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_us"
	"github.com/moov-io/watchman/pkg/sources/ofac"

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
		"initial_data_directory": log.String(dl.conf.InitialDataDirectory),
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

	// OFAC Records
	if slices.Contains(dl.conf.IncludedLists, search.SourceUSOFAC) {
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

	// CSL Records
	if slices.Contains(dl.conf.IncludedLists, search.SourceUSCSL) {
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

type preparedList struct {
	ListName search.SourceList
	Entities []search.Entity[search.Value]

	Hash string
}

func loadOFACRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-ofac-records")
	defer span.End()

	start := time.Now()
	files, err := ofac.Download(ctx, logger, conf.InitialDataDirectory)
	if err != nil {
		return fmt.Errorf("OFAC download: %v", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d OFAC files found", len(files))
	}

	logger.Debug().Logf("finished OFAC download: %v", time.Since(start))
	start = time.Now()

	res, err := ofac.Read(files)
	if err != nil {
		return fmt.Errorf("parsing OFAC: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := ofac.GroupIntoEntities(res.SDNs, res.Addresses, res.SDNComments, res.AlternateIdentities)

	logger.Debug().Logf("finished OFAC preparation: %v", time.Since(start))
	span.AddEvent("finished OFAC preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US OFAC")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSOFAC,
		Entities: entities,
		Hash:     res.ListHash,
	}
	return nil
}

func loadCSLUSRecords(ctx context.Context, logger log.Logger, conf Config, responseCh chan preparedList) error {
	ctx, span := telemetry.StartSpan(ctx, "load-us-csl-records")
	defer span.End()

	start := time.Now()
	files, err := csl_us.Download(ctx, logger, conf.InitialDataDirectory)
	if err != nil {
		return fmt.Errorf("US CSL download: %w", err)
	}
	defer files.Close()

	span.AddEvent("finished downloading")

	if len(files) == 0 {
		return fmt.Errorf("unexpected %d US CSL files found", len(files))
	}

	logger.Debug().Logf("finished US CSL download: %v", time.Since(start))
	start = time.Now()

	res, err := csl_us.Read(files)
	if err != nil {
		return fmt.Errorf("parsing US CSL: %w", err)
	}
	span.AddEvent("finished parsing")

	entities := csl_us.ConvertSanctionsData(res.SanctionsData)

	logger.Debug().Logf("finished US CSL preparation: %v", time.Since(start))
	span.AddEvent("finished US CSL preparation")

	if len(entities) == 0 && conf.ErrorOnEmptyList {
		return errors.New("no entities parsed from US CSL")
	}

	responseCh <- preparedList{
		ListName: search.SourceUSCSL,
		Entities: entities,
		Hash:     res.ListHash,
	}

	return nil
}
