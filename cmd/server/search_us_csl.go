// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"reflect"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"
)

func searchUSCSL(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		requestID := moovhttp.GetRequestID(r)

		limit := extractSearchLimit(r)
		filters := buildFilterRequest(r.URL)
		minMatch := extractSearchMinMatch(r)

		name := r.URL.Query().Get("name")
		resp := buildFullSearchResponseWith(searcher, cslGatherings, filters, limit, minMatch, name)

		logger.Info().With(log.Fields{
			"name":      log.String(name),
			"requestID": log.String(requestID),
		}).Log("performing US CSL search")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func precomputeCSLEntities[T any](items []*T, pipe *pipeliner) []*Result[T] {
	out := make([]*Result[T], len(items))

	for i, item := range items {
		name := cslName(item)
		if err := pipe.Do(name); err != nil {
			pipe.logger.LogErrorf("problem pipelining %T: %v", item, err)
			continue
		}

		var altNames []string

		elm := reflect.ValueOf(item).Elem()
		for i := 0; i < elm.NumField(); i++ {
			name := elm.Type().Field(i).Name
			_type := elm.Type().Field(i).Type.String()

			if name == "AlternateNames" && _type == "[]string" {
				alts, ok := elm.Field(i).Interface().([]string)
				if !ok {
					continue
				}
				for j := range alts {
					alt := &Name{Processed: alts[j]}
					pipe.Do(alt)
					altNames = append(altNames, alt.Processed)
				}
			}
		}

		out[i] = &Result[T]{
			Data:            *item,
			precomputedName: name.Processed,
			precomputedAlts: altNames,
		}
	}

	return out
}

// TopBISEntities searches BIS Entity List records by name and alias
func (s *searcher) TopBISEntities(limit int, minMatch float64, name string) []*Result[csl.EL] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start() // TODO(adam): This used to be on a pre-record gate, so this may have different perf metrics
	defer s.Gate.Done()

	return topResults[csl.EL](limit, minMatch, name, s.BISEntities)
}

// TopMEUs searches Military End User records by name and alias
func (s *searcher) TopMEUs(limit int, minMatch float64, name string) []*Result[csl.MEU] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.MEU](limit, minMatch, name, s.MilitaryEndUsers)
}

// TopSSIs searches Sectoral Sanctions records by Name and Alias
func (s *searcher) TopSSIs(limit int, minMatch float64, name string) []*Result[csl.SSI] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.SSI](limit, minMatch, name, s.SSIs)
}

// TopUVLs search Unverified Lists records by Name and Alias
func (s *searcher) TopUVLs(limit int, minMatch float64, name string) []*Result[csl.UVL] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.UVL](limit, minMatch, name, s.UVLs)
}

// TopISNs searches Nonproliferation Sanctions records by Name and Alias
func (s *searcher) TopISNs(limit int, minMatch float64, name string) []*Result[csl.ISN] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.ISN](limit, minMatch, name, s.ISNs)
}

// TopFSEs searches Foreign Sanctions Evaders records by Name and Alias
func (s *searcher) TopFSEs(limit int, minMatch float64, name string) []*Result[csl.FSE] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.FSE](limit, minMatch, name, s.FSEs)
}

// TopPLCs searches Palestinian Legislative Council records by Name and Alias
func (s *searcher) TopPLCs(limit int, minMatch float64, name string) []*Result[csl.PLC] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.PLC](limit, minMatch, name, s.PLCs)
}

// TopCAPs searches the CAPTA list by Name and Alias
func (s *searcher) TopCAPs(limit int, minMatch float64, name string) []*Result[csl.CAP] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.CAP](limit, minMatch, name, s.CAPs)
}

// TopDTCs searches the ITAR Debarred list by Name and Alias
func (s *searcher) TopDTCs(limit int, minMatch float64, name string) []*Result[csl.DTC] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.DTC](limit, minMatch, name, s.DTCs)
}

// TopCMICs searches the Non-SDN Chinese Military Industrial Complex list by Name and Alias
func (s *searcher) TopCMICs(limit int, minMatch float64, name string) []*Result[csl.CMIC] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.CMIC](limit, minMatch, name, s.CMICs)
}

// TopNS_MBS searches the Non-SDN Menu Based Sanctions list by Name and Alias
func (s *searcher) TopNS_MBS(limit int, minMatch float64, name string) []*Result[csl.NS_MBS] {
	s.RLock()
	defer s.RUnlock()

	s.Gate.Start()
	defer s.Gate.Done()

	return topResults[csl.NS_MBS](limit, minMatch, name, s.NS_MBSs)
}
