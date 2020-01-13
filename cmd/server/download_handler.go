// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	moovhttp "github.com/moov-io/base/http"
)

const (
	manualRefreshPath = "/data/refresh"
)

// manualRefreshHandler will register an endpoint on the admin server data refresh endpoint
func manualRefreshHandler(logger log.Logger, searcher *searcher, downloadRepo downloadRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("main", "admin: refreshing data")
		if stats, err := searcher.refreshData(""); err != nil {
			logger.Log("main", fmt.Sprintf("ERROR: admin: problem refreshing data: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if err := downloadRepo.recordStats(stats); err != nil {
				moovhttp.Problem(w, err)
				return
			}
			logger.Log(
				"main", fmt.Sprintf("admin: finished data refreshed %v ago", time.Since(stats.RefreshedAt)),
				"SDNs", stats.SDNs, "AltNames", stats.Alts, "Addresses", stats.Addresses, "SSI", stats.SectoralSanctions,
				"DPL", stats.DeniedPersons, "BISEntities", stats.BISEntities,
			)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(stats)
		}
	}
}
