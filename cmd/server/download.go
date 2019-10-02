// Copyright 2018 The Moov Authors
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
	"path/filepath"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	lastOFACDataRefreshSuccess = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "last_ofac_data_refresh_success",
		Help: "Unix timestamp of when OFAC data was last refreshed successfully",
	}, nil)
)

func init() {
	prometheus.MustRegister(lastOFACDataRefreshSuccess)
}

// Download holds counts for each type of OFAC and BIS Denied Persons List data parsed from files and a
// timestamp of when the download happened.
type Download struct {
	Timestamp     time.Time `json:"timestamp"`
	SDNs          int       `json:"SDNs"`
	Alts          int       `json:"altNames"`
	Addresses     int       `json:"addresses"`
	DeniedPersons int       `json:"deniedPersons"`
}

type downloadStats struct {
	SDNs          int       `json:"SDNs"`
	Alts          int       `json:"altNames"`
	Addresses     int       `json:"addresses"`
	DeniedPersons int       `json:"deniedPersons"`
	RefreshedAt   time.Time `json:"timestamp"`
}

// periodicDataRefresh will forever block for interval's duration and then download and reparse the OFAC data.
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
				s.logger.Log("main", fmt.Sprintf("ERROR: refreshing OFAC and/or BIS DPL data: %v", err))
			}
		} else {
			downloadRepo.recordStats(stats)
			if s.logger != nil {
				s.logger.Log("main", fmt.Sprintf("OFAC and BIS DPL data refreshed - Addresses=%d AltNames=%d SDNs=%d DPL=%d", stats.Addresses, stats.Alts, stats.SDNs, stats.DeniedPersons))
			}
			updates <- stats // send stats for re-search and watch notifications
		}
	}
}

// refreshData reaches out to the OFAC and BIS Denied Persons List websites to download the latest
// files and then runs ofac.Reader to parse and index data for searches.
func (s *searcher) refreshData(initialDir string) (*downloadStats, error) {
	if s.logger != nil {
		s.logger.Log("download", "Starting refresh of OFAC and DPL data")
	}

	// Download files
	dir, err := (&ofac.Downloader{
		Logger: s.logger,
	}).GetFiles(initialDir)
	if err != nil {
		return nil, fmt.Errorf("ERROR: downloading OFAC and DPL data: %v", err)
	}

	// Parse each file
	r := &ofac.Reader{}
	r.FileName = filepath.Join(dir, "add.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading add.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "alt.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading alt.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "sdn.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading sdn.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "dpl.txt")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading dpl.txt: %v", err)
	}

	// Precompute new data once for slight performance win
	sdns := precomputeSDNs(r.SDNs)
	adds := precomputeAddresses(r.Addresses)
	alts := precomputeAlts(r.AlternateIdentities)
	dps := precomputeDPs(r.DeniedPersons)

	stats := &downloadStats{
		SDNs:          len(sdns),
		Alts:          len(alts),
		Addresses:     len(adds),
		DeniedPersons: len(dps),
	}
	stats.RefreshedAt = lastRefresh(initialDir)

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	s.DPs = dps
	s.lastRefreshedAt = stats.RefreshedAt
	s.Unlock()

	if s.logger != nil {
		s.logger.Log("download", "Finished refresh of OFAC and BIS DPL data")
	}

	// record successful data refresh
	lastOFACDataRefreshSuccess.WithLabelValues().Set(float64(time.Now().Unix()))

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

		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			userID := moovhttp.GetUserID(r)
			logger.Log("download", "get latest downloads", "requestID", requestID, "userID", userID)
		}

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

	query := `insert into ofac_download_stats (downloaded_at, sdns, alt_names, addresses, denied_persons) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stats.RefreshedAt, stats.SDNs, stats.Alts, stats.Addresses, stats.DeniedPersons)
	return err
}

func (r *sqliteDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := `select downloaded_at, sdns, alt_names, addresses, denied_persons from ofac_download_stats order by downloaded_at desc limit ?;`
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
		if err := rows.Scan(&dl.Timestamp, &dl.SDNs, &dl.Alts, &dl.Addresses, &dl.DeniedPersons); err == nil {
			downloads = append(downloads, dl)
		}
	}
	return downloads, rows.Err()
}
