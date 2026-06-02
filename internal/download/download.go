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

// knownDownloadableLists is the set of standard built-in lists that watchman
// can download.  It is used when validating IgnoredDownloadErrors entries.
var knownDownloadableLists = []search.SourceList{
	search.SourceEUCSL,
	search.SourceUKCSL,
	search.SourceUSCSL,
	search.SourceUSOFAC,
	search.SourceUSNonSDN,
	search.SourceUSFinCEN311,
	search.SourceUNCSL,
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

	// Validate IgnoredDownloadErrors before spawning any goroutines so we can
	// return a clean error without having to tear down the errgroup machinery.
	ignoredLists := getIgnoredDownloadErrors(dl.conf)
	if err := validateIgnoredDownloadErrors(dl.conf, ignoredLists); err != nil {
		return stats, err
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

			// Catch any non-normalized entities before they're added into the index.
			var normalized int
			for idx := range entities {
				if entities[idx].PreparedFields.Name == "" || len(entities[idx].PreparedFields.NameFields) == 0 {
					normalized++
					entities[idx] = entities[idx].Normalize()
				}
			}
			if normalized > 0 {
				logger.Warn().Logf("normalized %d entities before index inclusion - normalize in the mapper", normalized)
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
				if slices.Contains(ignoredLists, search.SourceUSOFAC) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUSOFAC, err)
					return nil
				}
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
				if slices.Contains(ignoredLists, search.SourceUSNonSDN) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUSNonSDN, err)
					return nil
				}
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
				if slices.Contains(ignoredLists, search.SourceUSCSL) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUSCSL, err)
					return nil
				}
				return fmt.Errorf("loading US CSL records: %w", err)
			}
			return nil
		})
	}

	// FinCEN 311 Records
	if slices.Contains(requestedLists, search.SourceUSFinCEN311) {
		listsLoaded = append(listsLoaded, search.SourceUSFinCEN311)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()

			err := loadFinCEN311Records(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				if slices.Contains(ignoredLists, search.SourceUSFinCEN311) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUSFinCEN311, err)
					return nil
				}
				return fmt.Errorf("loading FinCEN 311 records: %w", err)
			}
			return nil
		})
	}

	// UK CSL Records
	if slices.Contains(requestedLists, search.SourceUKCSL) {
		listsLoaded = append(listsLoaded, search.SourceUKCSL)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()

			err := loadUKCSLRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				if slices.Contains(ignoredLists, search.SourceUKCSL) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUKCSL, err)
					return nil
				}
				return fmt.Errorf("loading UK CSL records: %w", err)
			}
			return nil
		})
	}

	// EU CSL Records
	if slices.Contains(requestedLists, search.SourceEUCSL) {
		listsLoaded = append(listsLoaded, search.SourceEUCSL)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()

			err := loadEUCSLRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				if slices.Contains(ignoredLists, search.SourceEUCSL) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceEUCSL, err)
					return nil
				}
				return fmt.Errorf("loading EU CSL records: %w", err)
			}
			return nil
		})
	}

	// OpenSanctions lists
	for _, list := range dl.conf.OpenSanctions.Lists {
		listsLoaded = append(listsLoaded, normalizeListName(list.SourceList))
	}
	producerWg.Add(1)
	g.Go(func() error {
		defer producerWg.Done()

		err := loadOpensanctionsRecords(ctx, logger, dl.conf, ignoredLists, preparedLists)
		if err != nil {
			return fmt.Errorf("loading opensanctions lists: %w", err)
		}
		return nil
	})

	// Senzing lists
	for _, list := range dl.conf.Senzing {
		listsLoaded = append(listsLoaded, normalizeListName(list.SourceList))
	}
	producerWg.Add(1)
	g.Go(func() error {
		defer producerWg.Done()

		err := loadSenzingRecords(ctx, logger, dl.conf, ignoredLists, preparedLists)
		if err != nil {
			return fmt.Errorf("loading senzing lists: %w", err)
		}
		return nil
	})

	// UN CSL Records
	if slices.Contains(requestedLists, search.SourceUNCSL) {
		listsLoaded = append(listsLoaded, search.SourceUNCSL)

		producerWg.Add(1)
		g.Go(func() error {
			defer producerWg.Done()

			err := loadUNCSLRecords(ctx, logger, dl.conf, preparedLists)
			if err != nil {
				if slices.Contains(ignoredLists, search.SourceUNCSL) {
					logger.Warn().Logf("ignoring error loading %s: %v", search.SourceUNCSL, err)
					return nil
				}
				return fmt.Errorf("loading UN CSL records: %w", err)
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
	for _, v := range conf.IncludedLists {
		if list := normalizeListName(v); list != "" {
			out = append(out, list)
		}
	}

	fromEnvStr := strings.TrimSpace(os.Getenv("INCLUDED_LISTS"))
	if fromEnvStr != "" {
		for _, v := range strings.Split(fromEnvStr, ",") {
			if list := normalizeListName(search.SourceList(v)); list != "" {
				out = append(out, list)
			}
		}
	}

	// Now add senzing lists (normalized so that custom names and IncludedLists
	// entries are treated case-insensitively and whitespace is ignored, matching
	// the behavior of IgnoredDownloadErrors and INCLUDED_LISTS env var).
	for _, list := range conf.OpenSanctions.Lists {
		if list := normalizeListName(list.SourceList); list != "" {
			out = append(out, list)
		}
	}
	for _, list := range conf.Senzing {
		if list := normalizeListName(list.SourceList); list != "" {
			out = append(out, list)
		}
	}

	// Sort and remove duplicates
	slices.Sort(out)
	return slices.Compact(out)
}

// normalizeListName lowercases and trims a SourceList value. User-supplied list
// names (from YAML or env) may have accidental case or whitespace; we normalize
// so that matching against the lowercase constants and deduping work reliably.
func normalizeListName(s search.SourceList) search.SourceList {
	trimmed := strings.TrimSpace(string(s))
	if trimmed == "" {
		return ""
	}
	return search.SourceList(strings.ToLower(trimmed))
}

// getIgnoredDownloadErrors returns the deduplicated, sorted set of source-list
// names for which download/parse errors should be suppressed.  It merges the
// YAML-configured IgnoredDownloadErrors field with the IGNORED_DOWNLOAD_ERRORS
// environment variable (comma-separated, same format as IncludedLists).
func getIgnoredDownloadErrors(conf Config) []search.SourceList {
	out := make([]search.SourceList, 0, len(conf.IgnoredDownloadErrors))
	for _, v := range conf.IgnoredDownloadErrors {
		if list := normalizeListName(v); list != "" {
			out = append(out, list)
		}
	}

	if fromEnvStr := strings.TrimSpace(os.Getenv("IGNORED_DOWNLOAD_ERRORS")); fromEnvStr != "" {
		for _, v := range strings.Split(fromEnvStr, ",") {
			if list := normalizeListName(search.SourceList(v)); list != "" {
				out = append(out, list)
			}
		}
	}

	slices.Sort(out)
	return slices.Compact(out)
}

// validateIgnoredDownloadErrors returns an error if any entry in ignoredLists is
// not a known standard downloadable list and is not a configured
// OpenSanctions or Senzing source list.
func validateIgnoredDownloadErrors(conf Config, ignoredLists []search.SourceList) error {
	if len(ignoredLists) == 0 {
		return nil
	}

	// Build the full set of valid names from standard + custom lists.
	allValid := make([]search.SourceList, len(knownDownloadableLists))
	copy(allValid, knownDownloadableLists)
	for _, l := range conf.OpenSanctions.Lists {
		allValid = append(allValid, normalizeListName(l.SourceList))
	}
	for _, l := range conf.Senzing {
		allValid = append(allValid, normalizeListName(l.SourceList))
	}

	for _, ignored := range ignoredLists {
		// Normalize the ignored entry for the lookup so that validation itself is
		// robust to casing/whitespace in the (already mostly-normalized) input.
		if !slices.Contains(allValid, normalizeListName(ignored)) {
			return fmt.Errorf("unknown list %q in IgnoredDownloadErrors: not a known downloadable SourceList or configured OpenSanctions/Senzing list", ignored)
		}
	}
	return nil
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
