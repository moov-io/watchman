// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	lastDataRefreshSuccess = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "last_data_refresh_success",
		Help: "Unix timestamp of when data was last refreshed successfully",
	}, nil)

	lastDataRefreshFailure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "last_data_refresh_failure",
		Help: "Unix timestamp of the most recent failure to refresh data",
	}, []string{"source"})

	lastDataRefreshCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "last_data_refresh_count",
		Help: "Count of records for a given sanction or entity list",
	}, []string{"source"})
)

func init() {
	prometheus.MustRegister(lastDataRefreshSuccess)
	prometheus.MustRegister(lastDataRefreshCount)
	prometheus.MustRegister(lastDataRefreshFailure)
}

// Download holds counts for each type of list data parsed from files and a
// timestamp of when the download happened.
type Download struct {
	Timestamp time.Time `json:"timestamp"`

	// US Office of Foreign Assets Control (OFAC)
	SDNs              int `json:"SDNs"`
	Alts              int `json:"altNames"`
	Addresses         int `json:"addresses"`
	SectoralSanctions int `json:"sectoralSanctions"`

	// US Bureau of Industry and Security (BIS)
	DeniedPersons int `json:"deniedPersons"`
	BISEntities   int `json:"bisEntities"`
}

type downloadStats struct {
	// US Office of Foreign Assets Control (OFAC)
	SDNs              int `json:"SDNs"`
	Alts              int `json:"altNames"`
	Addresses         int `json:"addresses"`
	SectoralSanctions int `json:"sectoralSanctions"`

	// US Bureau of Industry and Security (BIS)
	DeniedPersons int `json:"deniedPersons"`
	BISEntities   int `json:"bisEntities"`

	RefreshedAt time.Time `json:"timestamp"`
}

// periodicDataRefresh will forever block for interval's duration and then download and reparse the data.
// Download stats are recorded as part of a successful re-download and parse.
func (s *searcher) periodicDataRefresh(interval time.Duration, downloadRepo downloadRepository, updates chan *downloadStats) {
	if interval == 0*time.Second {
		s.logger.Log("download", fmt.Sprintf("not scheduling periodic refreshing duration=%v", interval))
		return
	}
	for {
		time.Sleep(interval)
		stats, err := s.refreshData("")
		if err != nil {
			if s.logger != nil {
				s.logger.Log("main", fmt.Sprintf("ERROR: refreshing data: %v", err))
			}
		} else {
			downloadRepo.recordStats(stats)
			if s.logger != nil {
				s.logger.Log(
					"main", fmt.Sprintf("data refreshed %v ago", time.Since(stats.RefreshedAt)),
					"SDNs", stats.SDNs, "AltNames", stats.Alts, "Addresses", stats.Addresses, "SSI", stats.SectoralSanctions,
					"DPL", stats.DeniedPersons, "BISEntities", stats.BISEntities,
				)
			}
			updates <- stats // send stats for re-search and watch notifications
		}
	}
}

func ofacRecords(logger log.Logger, initialDir string) (*ofac.Results, error) {
	files, err := ofac.Download(logger, initialDir)
	if err != nil {
		return nil, fmt.Errorf("download: %v", err)
	}
	if len(files) == 0 {
		return nil, errors.New("no OFAC Results")
	}

	var res *ofac.Results

	for i := range files {
		if i == 0 {
			res, err = ofac.Read(files[i])
			if err != nil {
				return nil, fmt.Errorf("read: %v", err)
			}
		} else {
			rr, err := ofac.Read(files[i])
			if err != nil {
				return nil, fmt.Errorf("read and replace: %v", err)
			}

			res.Addresses = append(res.Addresses, rr.Addresses...)
			res.AlternateIdentities = append(res.AlternateIdentities, rr.AlternateIdentities...)
			res.SDNs = append(res.SDNs, rr.SDNs...)
			res.SDNComments = append(res.SDNComments, rr.SDNComments...)
		}
	}
	return res, err
}

func dplRecords(logger log.Logger, initialDir string) ([]*dpl.DPL, error) {
	file, err := dpl.Download(logger, initialDir)
	if err != nil {
		return nil, err
	}
	return dpl.Read(file)
}

func cslRecords(logger log.Logger, initialDir string) (*csl.CSL, error) {
	file, err := csl.Download(logger, initialDir)
	if err != nil {
		logger.Log("download", "WARN: skipping CSL download", "description", err)
		return &csl.CSL{}, nil
	}
	cslRecords, err := csl.Read(file)
	if err != nil {
		return nil, err
	}
	return cslRecords, err
}

// refreshData reaches out to the various websites to download the latest
// files, runs each list's parser, and index data for searches.
func (s *searcher) refreshData(initialDir string) (*downloadStats, error) {
	if s.logger != nil {
		s.logger.Log("download", "Starting refresh of data")

		if initialDir != "" {
			s.logger.Log("download", fmt.Sprintf("reading files from %s", initialDir))
		}
	}

	lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))

	results, err := ofacRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))

		return nil, fmt.Errorf("OFAC records: %v", err)
	}

	sdns := precomputeSDNs(results.SDNs, results.Addresses, s.pipe)
	adds := precomputeAddresses(results.Addresses)
	alts := precomputeAlts(results.AlternateIdentities)

	deniedPersons, err := dplRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("DPs").Set(float64(time.Now().Unix()))

		return nil, fmt.Errorf("DPL records: %v", err)
	}
	dps := precomputeDPs(deniedPersons, s.pipe)

	consolidatedLists, err := cslRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("CSL").Set(float64(time.Now().Unix()))

		return nil, fmt.Errorf("CSL records: %v", err)
	}
	ssis := precomputeSSIs(consolidatedLists.SSIs, s.pipe)
	els := precomputeBISEntities(consolidatedLists.ELs, s.pipe)

	stats := &downloadStats{
		// OFAC
		SDNs:              len(sdns),
		Alts:              len(alts),
		Addresses:         len(adds),
		SectoralSanctions: len(ssis),
		// BIS
		BISEntities:   len(els),
		DeniedPersons: len(dps),
	}
	stats.RefreshedAt = lastRefresh(initialDir)

	// record prometheus metrics
	lastDataRefreshCount.WithLabelValues("SDNs").Set(float64(len(sdns)))
	lastDataRefreshCount.WithLabelValues("SSIs").Set(float64(len(ssis)))
	lastDataRefreshCount.WithLabelValues("BISEntities").Set(float64(len(els)))
	lastDataRefreshCount.WithLabelValues("DPs").Set(float64(len(dps)))

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	// OFAC
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	s.SSIs = ssis
	// BIS
	s.DPs = dps
	s.BISEntities = els
	// metadata
	s.lastRefreshedAt = stats.RefreshedAt
	s.Unlock()

	if s.logger != nil {
		s.logger.Log("download", "Finished refresh of data")
	}

	// record successful data refresh
	lastDataRefreshSuccess.WithLabelValues().Set(float64(time.Now().Unix()))

	return stats, nil
}

// lastRefresh returns a time.Time for the oldest file in dir or the current time if empty.
func lastRefresh(dir string) time.Time {
	if dir == "" {
		return time.Now()
	}

	infos, err := ioutil.ReadDir(dir)
	if len(infos) == 0 || err != nil {
		return time.Time{} // zero time because there's no initial data
	}

	oldest := infos[0].ModTime()
	for i := range infos[1:] {
		if t := infos[i].ModTime(); t.Before(oldest) {
			oldest = t
		}
	}
	return oldest
}

func addDownloadRoutes(logger log.Logger, r *mux.Router, repo downloadRepository) {
	r.Methods("GET").Path("/downloads").HandlerFunc(getLatestDownloads(logger, repo))
}

func getLatestDownloads(logger log.Logger, repo downloadRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		limit := extractSearchLimit(r)
		downloads, err := repo.latestDownloads(limit)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("download", "get latest downloads", "requestID", moovhttp.GetRequestID(r))

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(downloads); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

type downloadRepository interface {
	latestDownloads(limit int) ([]Download, error)
	recordStats(stats *downloadStats) error
}

type sqliteDownloadRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqliteDownloadRepository) close() error {
	return r.db.Close()
}

func (r *sqliteDownloadRepository) recordStats(stats *downloadStats) error {
	if stats == nil {
		return errors.New("recordStats: nil downloadStats")
	}

	query := `insert into download_stats (downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities) values (?, ?, ?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stats.RefreshedAt, stats.SDNs, stats.Alts, stats.Addresses, stats.SectoralSanctions, stats.DeniedPersons, stats.BISEntities)
	return err
}

func (r *sqliteDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := `select downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities from download_stats order by downloaded_at desc limit ?;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var downloads []Download
	for rows.Next() {
		var dl Download
		if err := rows.Scan(&dl.Timestamp, &dl.SDNs, &dl.Alts, &dl.Addresses, &dl.SectoralSanctions, &dl.DeniedPersons, &dl.BISEntities); err == nil {
			downloads = append(downloads, dl)
		}
	}
	return downloads, rows.Err()
}
