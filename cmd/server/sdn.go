// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// GET /sdn/:id/addresses
// - get addresses for a given SDN

// GET /sdn/:id/alternateNames
// - get alternate names for a given SDN

// GET /sdn/:id
// - get SDN information

func addSDNRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("GET").Path("/sdn/{id}/addresses").HandlerFunc(getSDNAddresses(logger))
	r.Methods("GET").Path("/sdn/{id}/alts").HandlerFunc(getSDNAltNames(logger))
	r.Methods("GET").Path("/sdn/{id}").HandlerFunc(getSDN(logger))
}

func getSDNAddresses(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func getSDNAltNames(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func getSDN(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
