// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

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
	BISEntities                      int `json:"bisEntities"`
	MilitaryEndUsers                 int `json:"militaryEndUsers"`
	SectoralSanctions                int `json:"sectoralSanctions"`
	Unverified                       int `json:"unverifiedCSL"`
	NonProliferationSanctions        int `json:"nonProliferationSanctions"`
	ForeignSanctionsEvaders          int `json:"foreignSanctionsEvaders"`
	PalestinianLegislativeCouncil    int `json:"palestinianLegislativeCouncil"`
	CAPTA                            int `json:"CAPTA"`
	ITARDebarred                     int `json:"ITARDebarred"`
	ChineseMilitaryIndustrialComplex int `json:"chineseMilitaryIndustrialComplex"`
	NonSDNMenuBasedSanctions         int `json:"nonSDNMenuBasedSanctions"`

	// EU Consolidated Sanctions List
	EUCSL int `json:"europeanSanctionsList"`

	// UK Consolidated Sanctions List
	UKCSL int `json:"ukConsolidatedSanctionsList"`

	// UK Sanctions List
	UKSanctionsList int `json:"ukSanctionsList"`

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
	errors := make([]string, 0, len(ss.Errors))
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
func (s *searcher) periodicDataRefresh(interval time.Duration, updates chan *DownloadStats) {
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
					"UVL":              log.Int(stats.Unverified),
					"ISN":              log.Int(stats.NonProliferationSanctions),
					"FSE":              log.Int(stats.ForeignSanctionsEvaders),
					"PLC":              log.Int(stats.PalestinianLegislativeCouncil),
					"CAP":              log.Int(stats.CAPTA),
					"DTC":              log.Int(stats.ITARDebarred),
					"CMIC":             log.Int(stats.ChineseMilitaryIndustrialComplex),
					"NS_MBS":           log.Int(stats.NonSDNMenuBasedSanctions),
					"EU_CSL":           log.Int(stats.EUCSL),
					"UK_CSL":           log.Int(stats.UKCSL),
				}).Logf("data refreshed %v ago", time.Since(stats.RefreshedAt))
			}
			updates <- stats // send stats back
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
	res, err := ofac.Read(files)
	if err != nil {
		return nil, err
	}

	// Merge comments into SDNs
	res.SDNs = mergeSpilloverRecords(res.SDNs, res.SDNComments)
	return res, nil
}

func mergeSpilloverRecords(sdns []*ofac.SDN, comments []*ofac.SDNComments) []*ofac.SDN {
	for i := range sdns {
		for j := range comments {
			if sdns[i].EntityID == comments[j].EntityID {
				sdns[i].Remarks += comments[j].RemarksExtended
			}
		}
	}
	return sdns
}

func dplRecords(logger log.Logger, initialDir string) ([]*dpl.DPL, error) {
	file, err := dpl.Download(logger, initialDir)
	if err != nil {
		return nil, err
	}

	return dpl.Read(file["dpl.txt"])
}

func cslRecords(logger log.Logger, initialDir string) (*csl.CSL, error) {
	file, err := csl.Download(logger, initialDir)
	if err != nil {
		logger.Warn().Logf("skipping CSL download: %v", err)
		return &csl.CSL{}, nil
	}
	cslRecords, err := csl.ReadFile(file["csl.csv"])
	if err != nil {
		return nil, err
	}
	return cslRecords, nil
}

func euCSLRecords(logger log.Logger, initialDir string) ([]*csl.EUCSLRecord, error) {
	file, err := csl.DownloadEU(logger, initialDir)
	if err != nil {
		logger.Warn().Logf("skipping EU CSL download: %v", err)
		// no error to return because we skip the download
		return nil, nil
	}

	cslRecords, _, err := csl.ParseEU(file["eu_csl.csv"])
	if err != nil {
		return nil, err
	}
	return cslRecords, err

}

func ukCSLRecords(logger log.Logger, initialDir string) ([]*csl.UKCSLRecord, error) {
	file, err := csl.DownloadUKCSL(logger, initialDir)
	if err != nil {
		logger.Warn().Logf("skipping UK CSL download: %v", err)
		// no error to return because we skip the download
		return nil, nil
	}
	cslRecords, _, err := csl.ReadUKCSLFile(file["ConList.csv"])
	if err != nil {
		return nil, err
	}
	return cslRecords, err
}

func ukSanctionsListRecords(logger log.Logger, initialDir string) ([]*csl.UKSanctionsListRecord, error) {
	file, err := csl.DownloadUKSanctionsList(logger, initialDir)
	if err != nil {
		logger.Warn().Logf("skipping UK Sanctions List download: %v", err)
		// no error to return because we skip the download
		return nil, nil
	}

	records, _, err := csl.ReadUKSanctionsListFile(file["UK_Sanctions_List.ods"])
	if err != nil {
		return nil, err
	}
	return records, err
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

	var err error
	lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))

	var ofacResults *ofac.Results
	withOFACList := cmp.Or(os.Getenv("WITH_OFAC_LIST"), "true")
	if strx.Yes(withOFACList) {
		ofacResults, err = ofacRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("SDNs").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("OFAC: %v", err))
		}
	}
	if ofacResults == nil {
		ofacResults = &ofac.Results{}
	}

	sdns := precomputeSDNs(ofacResults.SDNs, ofacResults.Addresses, s.pipe)
	adds := precomputeAddresses(ofacResults.Addresses)
	alts := precomputeAlts(ofacResults.AlternateIdentities, s.pipe)
	sdnComments := ofacResults.SDNComments

	var deniedPersons []*dpl.DPL
	withDPLList := cmp.Or(os.Getenv("WITH_US_DPL_LIST"), "true")
	if strx.Yes(withDPLList) {
		deniedPersons, err = dplRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("DPs").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("DPL: %v", err))
		}
	}
	dps := precomputeDPs(deniedPersons, s.pipe)

	var euCSLs []*Result[csl.EUCSLRecord]
	withEUScreeningList := cmp.Or(os.Getenv("WITH_EU_SCREENING_LIST"), "true")
	if strx.Yes(withEUScreeningList) {
		euConsolidatedList, err := euCSLRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("EUCSL").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("EUCSL: %v", err))
		}
		euCSLs = precomputeCSLEntities[csl.EUCSLRecord](euConsolidatedList, s.pipe)
	}

	var ukCSLs []*Result[csl.UKCSLRecord]
	withUKCSLSanctionsList := cmp.Or(os.Getenv("WITH_UK_CSL_SANCTIONS_LIST"), "true")
	if strx.Yes(withUKCSLSanctionsList) {
		ukConsolidatedList, err := ukCSLRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("UKCSL").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("UKCSL: %v", err))
		}
		ukCSLs = precomputeCSLEntities[csl.UKCSLRecord](ukConsolidatedList, s.pipe)
	}

	var ukSLs []*Result[csl.UKSanctionsListRecord]
	withUKSanctionsList := os.Getenv("WITH_UK_SANCTIONS_LIST")
	if strings.ToLower(withUKSanctionsList) == "true" {
		ukSanctionsList, err := ukSanctionsListRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("UKSanctionsList").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("UKSanctionsList: %v", err))
		}
		ukSLs = precomputeCSLEntities[csl.UKSanctionsListRecord](ukSanctionsList, s.pipe)

		stats.UKSanctionsList = len(ukSLs)
		lastDataRefreshCount.WithLabelValues("UKSL").Set(float64(len(ukSLs)))
	}

	// csl records from US downloaded here
	var usConsolidatedLists *csl.CSL
	withUSConsolidatedLists := cmp.Or(os.Getenv("WITH_US_CSL_SANCTIONS_LIST"), "true")
	if strx.Yes(withUSConsolidatedLists) {
		usConsolidatedLists, err = cslRecords(s.logger, initialDir)
		if err != nil {
			lastDataRefreshFailure.WithLabelValues("CSL").Set(float64(time.Now().Unix()))
			stats.Errors = append(stats.Errors, fmt.Errorf("US CSL: %v", err))
		}
	}
	if usConsolidatedLists == nil {
		usConsolidatedLists = new(csl.CSL)
	}

	els := precomputeCSLEntities[csl.EL](usConsolidatedLists.ELs, s.pipe)
	meus := precomputeCSLEntities[csl.MEU](usConsolidatedLists.MEUs, s.pipe)
	ssis := precomputeCSLEntities[csl.SSI](usConsolidatedLists.SSIs, s.pipe)
	uvls := precomputeCSLEntities[csl.UVL](usConsolidatedLists.UVLs, s.pipe)
	isns := precomputeCSLEntities[csl.ISN](usConsolidatedLists.ISNs, s.pipe)
	fses := precomputeCSLEntities[csl.FSE](usConsolidatedLists.FSEs, s.pipe)
	plcs := precomputeCSLEntities[csl.PLC](usConsolidatedLists.PLCs, s.pipe)
	caps := precomputeCSLEntities[csl.CAP](usConsolidatedLists.CAPs, s.pipe)
	dtcs := precomputeCSLEntities[csl.DTC](usConsolidatedLists.DTCs, s.pipe)
	cmics := precomputeCSLEntities[csl.CMIC](usConsolidatedLists.CMICs, s.pipe)
	ns_mbss := precomputeCSLEntities[csl.NS_MBS](usConsolidatedLists.NS_MBSs, s.pipe)

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
	stats.Unverified = len(uvls)
	stats.NonProliferationSanctions = len(isns)
	stats.ForeignSanctionsEvaders = len(fses)
	stats.PalestinianLegislativeCouncil = len(plcs)
	stats.CAPTA = len(caps)
	stats.ITARDebarred = len(dtcs)
	stats.ChineseMilitaryIndustrialComplex = len(cmics)
	stats.NonSDNMenuBasedSanctions = len(ns_mbss)
	// EU - CSL
	stats.EUCSL = len(euCSLs)

	// UK - CSL
	stats.UKCSL = len(ukCSLs)

	// record prometheus metrics
	lastDataRefreshCount.WithLabelValues("SDNs").Set(float64(len(sdns)))
	lastDataRefreshCount.WithLabelValues("SSIs").Set(float64(len(ssis)))
	lastDataRefreshCount.WithLabelValues("BISEntities").Set(float64(len(els)))
	lastDataRefreshCount.WithLabelValues("MilitaryEndUsers").Set(float64(len(meus)))
	lastDataRefreshCount.WithLabelValues("DPs").Set(float64(len(dps)))
	lastDataRefreshCount.WithLabelValues("UVLs").Set(float64(len(uvls)))
	lastDataRefreshCount.WithLabelValues("ISNs").Set(float64(len(isns)))
	lastDataRefreshCount.WithLabelValues("FSEs").Set(float64(len(fses)))
	lastDataRefreshCount.WithLabelValues("PLCs").Set(float64(len(plcs)))
	lastDataRefreshCount.WithLabelValues("CAPs").Set(float64(len(caps)))
	lastDataRefreshCount.WithLabelValues("DTCs").Set(float64(len(dtcs)))
	lastDataRefreshCount.WithLabelValues("CMICs").Set(float64(len(cmics)))
	lastDataRefreshCount.WithLabelValues("NS_MBSs").Set(float64(len(ns_mbss)))
	// EU CSL
	lastDataRefreshCount.WithLabelValues("EUCSL").Set(float64(len(euCSLs)))
	// UK CSL
	lastDataRefreshCount.WithLabelValues("UKCSL").Set(float64(len(ukCSLs)))

	if len(stats.Errors) > 0 {
		return stats, stats
	}

	// Set new records after precomputation (to minimize lock contention)
	s.Lock()
	// OFAC
	s.SDNs = sdns
	s.Addresses = adds
	s.Alts = alts
	s.SDNComments = sdnComments
	// BIS
	s.DPs = dps
	// CSL
	s.BISEntities = els
	s.MilitaryEndUsers = meus
	s.SSIs = ssis
	s.UVLs = uvls
	s.ISNs = isns
	s.FSEs = fses
	s.PLCs = plcs
	s.CAPs = caps
	s.DTCs = dtcs
	s.CMICs = cmics
	s.NS_MBSs = ns_mbss
	//EUCSL
	s.EUCSL = euCSLs
	//UKCSL
	s.UKCSL = ukCSLs
	s.UKSanctionsList = ukSLs
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

	fds, err := os.ReadDir(dir)
	if len(fds) == 0 || err != nil {
		return time.Time{} // zero time because there's no initial data
	}

	oldest := time.Now().In(time.UTC)
	for i := range fds {
		info, err := fds[i].Info()
		if err != nil {
			continue
		}
		if t := info.ModTime(); t.Before(oldest) {
			oldest = t
		}
	}
	return oldest.In(time.UTC)
}
