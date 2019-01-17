// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// POST /search/address
// - Search for address records matching the given search criteria.

// POST /search/name?k=v
// - fuzzy name search
// - See: https://github.com/moov-io/ofac/issues/6

// POST /search/alt?k=v
// - fuzzy alternate name search

// POST /search/company?k=v
// - fuzzy company name search

func addSearchRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("POST").Path("/search/address").HandlerFunc(searchByAddress(logger))
	r.Methods("POST").Path("/search/name").HandlerFunc(searchByName(logger))
	r.Methods("POST").Path("/search/alt").HandlerFunc(searchByAltName(logger))
	// r.Methods("POST").Path("/search/company").HandlerFunc(searchByAddress()) // TODO
}

type addressSearchRequest struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
}

func searchByAddress(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func searchByName(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)

		// ?name=foo
	}
}

func searchByAltName(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)

		// ?name=foo
	}
}
