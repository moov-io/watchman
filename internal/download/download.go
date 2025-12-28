package download

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"

	"golang.org/x/sync/errgroup"
)

// GeocodingService defines the interface for geocoding addresses.
// This is used to decouple the download package from the geocoding implementation.
type GeocodingService interface {
	GeocodeAddresses(ctx context.Context, addresses []search.Address) []search.Address
}

type Downloader interface {
	RefreshAll(ctx context.Context) (Stats, error)
}

func NewDownloader(logger log.Logger, conf Config, geocoder GeocodingService) (Downloader, error) {
	return &downloader{
		logger:   logger,
		conf:     conf,
		geocoder: geocoder,
	}, nil
}

type downloader struct {
	logger   log.Logger
	conf     Config
	geocoder GeocodingService
}

func (dl *downloader) RefreshAll(ctx context.Context) (Stats, error) {
	ctx, span := telemetry.StartSpan(ctx, "refresh-all")
	defer span.End()

	start := time.Now()

	stats := Stats{
		Lists:      make(map[string]int),
		ListHashes: make(map[string]string),
		StartedAt:  time.Now().In(time.UTC),
	}
	logger := dl.logger.Info().With(log.Fields{
		"initial_data_directory": log.String(expandInitialDir(initialDataDirectory(dl.conf))),
	})

	g, ctx := errgroup.WithContext(ctx)
	preparedLists := make(chan preparedList, 10)

	// Start a goroutine to accumulate results
	resultsDone := make(chan struct{})
	go func() {
		defer close(resultsDone)

		for list := range preparedLists {
			entities := list.Entities

			// Apply geocoding to entities if service is available
			if dl.geocoder != nil {
				entities = dl.geocodeEntities(ctx, entities)
			}

			logger.Info().Logf("adding %d entities from %v", len(entities), list.ListName)

			stats.Lists[string(list.ListName)] = len(entities)
			stats.ListHashes[string(list.ListName)] = list.Hash

			stats.Entities = append(stats.Entities, entities...)
		}
	}()

	// Create a WaitGroup to track all producers
	var producerWg sync.WaitGroup

	// Track what lists have been requested and loaded
	requestedLists := getIncludedLists(dl.conf)
	var listsLoaded []search.SourceList

	if len(requestedLists) == 0 {
		logger.Warn().Log("no lists have been configured!")
	}
	logger.Info().Logf("starting list refresh of %v", requestedLists)

	// OFAC Records
	if slices.Contains(requestedLists, search.SourceUSOFAC) {
		listsLoaded = append(listsLoaded, search.SourceUSOFAC)

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
	if slices.Contains(requestedLists, search.SourceUSNonSDN) {
		listsLoaded = append(listsLoaded, search.SourceUSNonSDN)

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
	if slices.Contains(requestedLists, search.SourceUSCSL) {
		listsLoaded = append(listsLoaded, search.SourceUSCSL)

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

	// Senzing lists
	for idx := range dl.conf.Senzing {
		listsLoaded = append(listsLoaded, dl.conf.Senzing[idx].SourceList)
	}

	producerWg.Add(1)
	g.Go(func() error {
		defer producerWg.Done()

		err := loadSenzingRecords(ctx, logger, dl.conf, preparedLists)
		if err != nil {
			return fmt.Errorf("loading senzing lists: %w", err)
		}
		return nil
	})

	// OpenSanctions PEP Records
	if slices.Contains(requestedLists, search.SourceOpenSanctionsPEP) {
		listsLoaded = append(listsLoaded, search.SourceOpenSanctionsPEP)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()
			err := loadOpenSanctionsPEPRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				return fmt.Errorf("loading OpenSanctions PEP records: %w", err)
			}
			return nil
		})
	}

	// Compare the configured lists against those we actually loaded.
	// Any extra lists are an error as we don't want to silently ignore them.
	if len(listsLoaded) > len(requestedLists) {
		close(preparedLists)

		return stats, fmt.Errorf("loaded extra lists: %#v loaded compared to %#v configured", listsLoaded, requestedLists)
	}
	if extra := findExtraLists(requestedLists, listsLoaded); extra != "" {
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

	// Build TF-IDF index from all entity names
	stats.TFIDFIndex = buildTFIDFIndex(logger, stats.Entities)

	stats.EndedAt = time.Now().In(time.UTC)

	return stats, nil
}

// buildTFIDFIndex creates a TF-IDF index from all entity names.
// It extracts NameFields and AltNameFields from each entity's PreparedFields.
func buildTFIDFIndex(logger log.Logger, entities []search.Entity[search.Value]) *tfidf.Index {
	cfg := tfidf.ConfigFromEnvironment()
	idx := tfidf.NewIndex(cfg)

	if !cfg.Enabled {
		logger.Info().Log("TF-IDF indexing disabled")
		return idx
	}

	start := time.Now()

	// Collect all name terms as documents
	// Each entity's name (and alt names) is treated as a separate document
	var documents [][]string

	for i := range entities {
		// Add primary name fields
		if len(entities[i].PreparedFields.NameFields) > 0 {
			documents = append(documents, entities[i].PreparedFields.NameFields)
		}

		// Add alternate name fields
		for _, altFields := range entities[i].PreparedFields.AltNameFields {
			if len(altFields) > 0 {
				documents = append(documents, altFields)
			}
		}
	}

	idx.Build(documents)

	stats := idx.Stats()
	logger.Info().Logf("built TF-IDF index: %d documents, %d unique terms in %v",
		stats.TotalDocuments, stats.UniqueTerms, time.Since(start))

	return idx
}

func getIncludedLists(conf Config) []search.SourceList {
	out := make([]search.SourceList, 0, len(conf.IncludedLists))
	out = append(out, conf.IncludedLists...)

	fromEnvStr := strings.TrimSpace(os.Getenv("INCLUDED_LISTS"))
	if fromEnvStr != "" {
		for _, v := range strings.Split(fromEnvStr, ",") {
			list := strings.ToLower(strings.TrimSpace(v))
			if list != "" {
				out = append(out, search.SourceList(list))
			}
		}
	}

	// Now add Senzing
	for i := range conf.Senzing {
		out = append(out, conf.Senzing[i].SourceList)
	}

	// Sort and remove duplicates
	slices.Sort(out)
	return slices.Compact(out)
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

// geocodeEntities applies geocoding to all addresses in the given entities.
func (dl *downloader) geocodeEntities(ctx context.Context, entities []search.Entity[search.Value]) []search.Entity[search.Value] {
	if dl.geocoder == nil {
		return entities
	}

	var wg sync.WaitGroup
	wg.Add(len(entities))

	for i := range entities {
		go func(i int) {
			defer wg.Done()

			if len(entities[i].Addresses) > 0 {
				entities[i].Addresses = dl.geocoder.GeocodeAddresses(ctx, entities[i].Addresses)
			}
		}(i)
	}

	wg.Wait()

	return entities
}
