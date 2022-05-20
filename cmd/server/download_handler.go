// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
)

const (
	manualRefreshPath = "/data/refresh"
)

// manualRefreshHandler will register an endpoint on the admin server data refresh endpoint
func manualRefreshHandler(logger log.Logger, searcher *searcher, updates chan *DownloadStats, downloadRepo downloadRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("admin: refreshing data")

		if stats, err := searcher.refreshData(""); err != nil {
			logger.LogErrorf("ERROR: admin: problem refreshing data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if err := downloadRepo.recordStats(stats); err != nil {
				moovhttp.Problem(w, err)
				return
			}

			go func() {
				updates <- stats
			}()

			logger.Info().With(log.Fields{
				"SDNs":        log.Int(stats.SDNs),
				"AltNames":    log.Int(stats.Alts),
				"Addresses":   log.Int(stats.Addresses),
				"SSI":         log.Int(stats.SectoralSanctions),
				"DPL":         log.Int(stats.DeniedPersons),
				"BISEntities": log.Int(stats.BISEntities),
			}).Logf("admin: finished data refresh %v ago", time.Since(stats.RefreshedAt))

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(stats)
		}
	}
}
