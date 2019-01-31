// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
)

type downloadStats struct {
	SDNs      int `json:"SDNs"`
	Alts      int `json:"altNames"`
	Addresses int `json:"addresses"`
}

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

type downloadRepository interface {
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
