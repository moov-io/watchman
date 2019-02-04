// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type Download struct {
	Timestamp time.Time `json:"timestamp"`
	SDNs      int       `json:"SDNs"`
	Alts      int       `json:"altNames"`
	Addresses int       `json:"addresses"`
}

type downloadStats struct {
	SDNs      int `json:"SDNs"`
	Alts      int `json:"altNames"`
	Addresses int `json:"addresses"`
}

// periodicDataRefresh will forever block for interval's duration and then download and reparse the OFAC data.
// Download stats are recorded as part of a successful re-download and parse.
func (s *searcher) periodicDataRefresh(interval time.Duration, downloadRepo downloadRepository, updates chan *downloadStats) {
	for {
		time.Sleep(interval)
		stats, err := s.refreshData()
		if err != nil {
			if s.logger != nil {
				s.logger.Log("main", fmt.Sprintf("ERROR: refreshing OFAC data: %v", err))
			}
		} else {
			downloadRepo.recordStats(stats)
			if s.logger != nil {
				s.logger.Log("main", fmt.Sprintf("OFAC data refreshed - Addresses=%d AltNames=%d SDNs=%d", stats.Addresses, stats.Alts, stats.SDNs))
			}
			updates <- stats // send stats for re-search and watch notifications
		}
	}
}

// refreshData reaches out to the OFAC website to download the latest files and then runs ofac.Reader to
// parse and index data for searches.
func (s *searcher) refreshData() (*downloadStats, error) {
	if s.logger != nil {
		s.logger.Log("download", "Starting refresh of OFAC data")
	}

	// Download files
	dir, err := (&ofac.Downloader{}).GetFiles()
	if err != nil {
		return nil, fmt.Errorf("ERROR: downloading OFAC data: %v", err)
	}

	// Parse each OFAC file
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

	// Precompute new data
	sdns := precomputeSDNs(r.SDNs)
	adds := precomputeAddresses(r.Addresses)
	alts := precomputeAlts(r.AlternateIdentities)

	stats := &downloadStats{
		SDNs:      len(sdns),
		Alts:      len(alts),
		Addresses: len(adds),
	}

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	s.Unlock()

	if s.logger != nil {
		s.logger.Log("download", "Finished refresh of OFAC data")
	}

	return stats, nil
}

func addDownloadRoutes(logger log.Logger, r *mux.Router, repo downloadRepository) {
	r.Methods("GET").Path("/downloads").HandlerFunc(getLatestDownloads(logger, repo))
}

func getLatestDownloads(logger log.Logger, repo downloadRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		limit := extractSearchLimit(r)
		downloads, err := repo.latestDownloads(limit)
		if err != nil {
			moovhttp.Problem(w, err)
			return
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

	query := `insert into ofac_download_stats (downloaded_at, sdns, alt_names, addresses) values (?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), stats.SDNs, stats.Alts, stats.Addresses)
	return err
}

func (r *sqliteDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := `select downloaded_at, sdns, alt_names, addresses from ofac_download_stats order by downloaded_at desc limit ?;`
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
		if err := rows.Scan(&dl.Timestamp, &dl.SDNs, &dl.Alts, &dl.Addresses); err == nil {
			downloads = append(downloads, dl)
		}
	}
	return downloads, nil
}
