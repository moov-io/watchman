// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

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

// DownloadStats holds counts for each type of list data parsed from files and a
// timestamp of when the download happened.
type DownloadStats struct {
	// US Office of Foreign Assets Control (OFAC)
	SDNs      int `json:"SDNs"`
	Alts      int `json:"altNames"`
	Addresses int `json:"addresses"`

	// US Bureau of Industry and Security (BIS)
	DeniedPersons int `json:"deniedPersons"`

	// Consolidated Screening List (CSL)
	BISEntities       int `json:"bisEntities"`
	MilitaryEndUsers  int `json:"militaryEndUsers"`
	SectoralSanctions int `json:"sectoralSanctions"`

	Errors      []error   `json:"-"`
	RefreshedAt time.Time `json:"timestamp"`
}

func (ss *DownloadStats) Error() string {
	var buf bytes.Buffer
	for i := range ss.Errors {
		buf.WriteString(ss.Errors[i].Error() + "\n")
	}
	return buf.String()
}

func (ss *DownloadStats) MarshalJSON() ([]byte, error) {
	type Aux struct {
		DownloadStats
		Errors []string `json:"errors"`
	}
	var errors []string
	for i := range ss.Errors {
		errors = append(errors, ss.Errors[i].Error())
	}
	return json.Marshal(Aux{
		DownloadStats: *ss,
		Errors:        errors,
	})
}

// periodicDataRefresh will forever block for interval's duration and then download and reparse the data.
// Download stats are recorded as part of a successful re-download and parse.
func (s *searcher) periodicDataRefresh(interval time.Duration, downloadRepo downloadRepository, updates chan *DownloadStats) {
	if interval == 0*time.Second {
		s.logger.Logf("not scheduling periodic refreshing duration=%v", interval)
		return
	}
	for {
		time.Sleep(interval)
		stats, err := s.refreshData("")
		if err != nil {
			if s.logger != nil {
				s.logger.Info().Logf("ERROR: refreshing data: %v", err)
			}
		} else {
			downloadRepo.recordStats(stats)
			if s.logger != nil {
				s.logger.Info().With(log.Fields{
					// OFAC
					"SDNs":      log.Int(stats.SDNs),
					"AltNames":  log.Int(stats.Alts),
					"Addresses": log.Int(stats.Addresses),

					// BIS
					"DPL": log.Int(stats.DeniedPersons),

					// CSL
					"BISEntities":      log.Int(stats.BISEntities),
					"MilitaryEndUsers": log.Int(stats.MilitaryEndUsers),
					"SSI":              log.Int(stats.SectoralSanctions),
				}).Logf("data refreshed %v ago", time.Since(stats.RefreshedAt))
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
			rr, err := ofac.Read(files[i])
			if err != nil {
				return nil, fmt.Errorf("read: %v", err)
			}
			if rr != nil {
				res = rr
			}
		} else {
			rr, err := ofac.Read(files[i])
			if err != nil {
				return nil, fmt.Errorf("read and replace: %v", err)
			}
			if rr != nil {
				res.Addresses = append(res.Addresses, rr.Addresses...)
				res.AlternateIdentities = append(res.AlternateIdentities, rr.AlternateIdentities...)
				res.SDNs = append(res.SDNs, rr.SDNs...)
				res.SDNComments = append(res.SDNComments, rr.SDNComments...)
			}
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
		logger.Warn().LogErrorf("skipping CSL download: %v", err)
		return &csl.CSL{}, nil
	}
	cslRecords, err := csl.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return cslRecords, err
}

// refreshData reaches out to the various websites to download the latest
// files, runs each list's parser, and index data for searches.
func (s *searcher) refreshData(initialDir string) (*DownloadStats, error) {
	if s.logger != nil {
		s.logger.Log("Starting refresh of data")

		if initialDir != "" {
			s.logger.Logf("reading files from %s", initialDir)
		}
	}

	stats := &DownloadStats{
		RefreshedAt: lastRefresh(initialDir),
	}

	lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))

	results, err := ofacRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))
		stats.Errors = append(stats.Errors, fmt.Errorf("OFAC: %v", err))
	}

	sdns := precomputeSDNs(results.SDNs, results.Addresses, s.pipe)
	adds := precomputeAddresses(results.Addresses)
	alts := precomputeAlts(results.AlternateIdentities, s.pipe)

	deniedPersons, err := dplRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("DPs").Set(float64(time.Now().Unix()))
		stats.Errors = append(stats.Errors, fmt.Errorf("DPL: %v", err))
	}
	dps := precomputeDPs(deniedPersons, s.pipe)

	consolidatedLists, err := cslRecords(s.logger, initialDir)
	if err != nil {
		lastDataRefreshFailure.WithLabelValues("CSL").Set(float64(time.Now().Unix()))
		stats.Errors = append(stats.Errors, fmt.Errorf("CSL: %v", err))
	}
	els := precomputeCSLEntities[csl.EL](consolidatedLists.ELs, s.pipe)
	meus := precomputeCSLEntities[csl.MEU](consolidatedLists.MEUs, s.pipe)
	ssis := precomputeCSLEntities[csl.SSI](consolidatedLists.SSIs, s.pipe)

	// OFAC
	stats.SDNs = len(sdns)
	stats.Alts = len(alts)
	stats.Addresses = len(adds)
	// BIS
	stats.DeniedPersons = len(dps)
	// CSL
	stats.BISEntities = len(els)
	stats.MilitaryEndUsers = len(meus)
	stats.SectoralSanctions = len(ssis)

	// record prometheus metrics
	lastDataRefreshCount.WithLabelValues("SDNs").Set(float64(len(sdns)))
	lastDataRefreshCount.WithLabelValues("SSIs").Set(float64(len(ssis)))
	lastDataRefreshCount.WithLabelValues("BISEntities").Set(float64(len(els)))
	lastDataRefreshCount.WithLabelValues("MilitaryEndUsers").Set(float64(len(meus)))
	lastDataRefreshCount.WithLabelValues("DPs").Set(float64(len(dps)))

	if len(stats.Errors) > 0 {
		return stats, stats
	}

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	// OFAC
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	// BIS
	s.DPs = dps
	// CSL
	s.BISEntities = els
	s.MilitaryEndUsers = meus
	s.SSIs = ssis
	// metadata
	s.lastRefreshedAt = stats.RefreshedAt
	s.Unlock()

	if s.logger != nil {
		s.logger.Log("Finished refresh of data")
	}

	// record successful data refresh
	lastDataRefreshSuccess.WithLabelValues().Set(float64(time.Now().Unix()))

	return stats, nil
}

// lastRefresh returns a time.Time for the oldest file in dir or the current time if empty.
func lastRefresh(dir string) time.Time {
	if dir == "" {
		return time.Now().In(time.UTC)
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
	return oldest.In(time.UTC)
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

		logger.Info().With(log.Fields{
			"requestID": log.String(moovhttp.GetRequestID(r)),
		}).Log("get latest downloads")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(downloads); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

type downloadRepository interface {
	latestDownloads(limit int) ([]DownloadStats, error)
	recordStats(stats *DownloadStats) error
}

type sqliteDownloadRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqliteDownloadRepository) close() error {
	return r.db.Close()
}

func (r *sqliteDownloadRepository) recordStats(stats *DownloadStats) error {
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

func (r *sqliteDownloadRepository) latestDownloads(limit int) ([]DownloadStats, error) {
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

	var downloads []DownloadStats
	for rows.Next() {
		var dl DownloadStats
		if err := rows.Scan(&dl.RefreshedAt, &dl.SDNs, &dl.Alts, &dl.Addresses, &dl.SectoralSanctions, &dl.DeniedPersons, &dl.BISEntities); err == nil {
			downloads = append(downloads, dl)
		}
	}
	return downloads, rows.Err()
}
