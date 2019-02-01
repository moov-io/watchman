// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"net/http"

	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoSDNId = errors.New("no SDN Id provided")
)

func addSDNRoutes(logger log.Logger, r *mux.Router, searcher *searcher) {
	r.Methods("GET").Path("/sdn/{sdnId}/addresses").HandlerFunc(getSDNAddresses(logger, searcher))
	r.Methods("GET").Path("/sdn/{sdnId}/alts").HandlerFunc(getSDNAltNames(logger, searcher))
	r.Methods("GET").Path("/sdn/{sdnId}").HandlerFunc(getSDN(logger, searcher))
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
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		id, limit := getSDNId(w, r), extractSearchLimit(r)
		if id == "" {
			return
		}
		addresses := searcher.FindAddresses(limit, func(add *Address) bool {
			return add.Address.EntityID == id
		})
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(addresses); err != nil {
			moovhttp.Problem(w, err) // TODO(adam): JSON errors should moovhttp.InternalError (wrapped, see auth's http.go)
			return
		}
	}
}

func getSDNAltNames(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		id, limit := getSDNId(w, r), extractSearchLimit(r)
		if id == "" {
			return
		}
		alts := searcher.FindAlts(limit, func(alt *Alt) bool {
			return alt.AlternateIdentity.EntityID == id
		})
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
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		id := getSDNId(w, r)
		if id == "" {
			return
		}
		sdns := searcher.FindSDNs(1, func(s *SDN) bool {
			return s.SDN.EntityID == id
		})
		if len(sdns) != 1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sdns[0]); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
