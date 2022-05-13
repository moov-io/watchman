// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"

	"github.com/gorilla/mux"
)

var (
	errNoSDNId = errors.New("no SDN Id provided")
)

func addSDNRoutes(logger log.Logger, r *mux.Router, searcher *searcher) {
	r.Methods("GET").Path("/ofac/sdn/{sdnId}/addresses").HandlerFunc(getSDNAddresses(logger, searcher))
	r.Methods("GET").Path("/ofac/sdn/{sdnId}/alts").HandlerFunc(getSDNAltNames(logger, searcher))
	r.Methods("GET").Path("/ofac/sdn/{sdnId}").HandlerFunc(getSDN(logger, searcher))
}

func getSDNId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["sdnId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoSDNId)
		return ""
	}
	return v
}

func getSDNAddresses(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		id, limit := getSDNId(w, r), extractSearchLimit(r)
		if id == "" {
			return
		}

		addresses := searcher.FindAddresses(limit, id)

		logger.Info().With(log.Fields{
			"requestID": log.String(moovhttp.GetRequestID(r)),
		}).Logf("get sdn=%s addresses", id)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(addresses); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func getSDNAltNames(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		id, limit := getSDNId(w, r), extractSearchLimit(r)
		if id == "" {
			return
		}

		alts := searcher.FindAlts(limit, id)

		logger.Info().With(log.Fields{
			"requestID": log.String(moovhttp.GetRequestID(r)),
		}).Logf("get sdn=%s alt names", id)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(alts); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func getSDN(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		id := getSDNId(w, r)
		if id == "" {
			return
		}
		sdn := searcher.FindSDN(id)
		if sdn == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Info().With(log.Fields{
			"requestID": log.String(moovhttp.GetRequestID(r)),
		}).Logf("get sdn=%s", id)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sdn); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
