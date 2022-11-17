// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"
)

// search EUCLS
func searchEUCSL(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		requestID := moovhttp.GetRequestID(r)

		limit := extractSearchLimit(r)
		filters := buildFilterRequest(r.URL)
		minMatch := extractSearchMinMatch(r)

		name := r.URL.Query().Get("name")
		resp := buildFullSearchResponseWith(searcher, euGatherings, filters, limit, minMatch, name)

		logger.Info().With(log.Fields{
			"name":      log.String(name),
			"requestID": log.String(requestID),
		}).Log("performing EU-CSL search")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

// TopEUCSL searches the EU Sanctions list by Name and Alias
func (s *searcher) TopEUCSL(limit int, minMatch float64, name string) []*Result[csl.EUCSLRecord] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.EUCSLRecord](limit, minMatch, name, s.EUCSL)
}
