// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")
)

func addSearchRoutes(logger log.Logger, r *mux.Router, reader *ofac.Reader) {
	r.Methods("GET").Path("/search").HandlerFunc(search(logger, reader))
}

type addressSearchRequest struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
}

func (req addressSearchRequest) empty() bool {
	return req.Address == "" && req.City == "" && req.State == "" &&
		req.Providence == "" && req.Zip == "" && req.Country == ""
}

func readAddressSearchRequest(u *url.URL) addressSearchRequest {
	return addressSearchRequest{
		Address:    u.Query().Get("address"),
		City:       u.Query().Get("city"),
		State:      u.Query().Get("state"),
		Providence: u.Query().Get("providence"),
		Zip:        u.Query().Get("zip"),
		Country:    u.Query().Get("country"),
	}
}

func search(logger log.Logger, reader *ofac.Reader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		// Search by Name // TODO(adam): handle multiple?
		if name := r.URL.Query().Get("name"); name != "" {
			searchByName(logger, reader, name)(w, r)
			return
		}

		// Search by Alt Name
		if alt := r.URL.Query().Get("altName"); alt != "" {
			searchByAltName(logger, reader, alt)(w, r)
			return
		}

		// Search Addresses
		if req := readAddressSearchRequest(r.URL); !req.empty() {
			searchByAddress(logger, reader, req)(w, r)
			return
		}

		// Fallback if no search params were found
		moovhttp.Problem(w, errNoSearchParams)
	}
}

func searchByAddress(logger log.Logger, reader *ofac.Reader, req addressSearchRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answers []ofac.Address
		for i := range reader.Addresses {
			add := reader.Addresses[i]
			if strings.Contains(add.Address, req.Address) {
				answers = append(answers, add)
			}
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answers); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func searchByName(logger log.Logger, reader *ofac.Reader, nameSlug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answers []ofac.SDN // TODO(adam): pointers on ofac.Reader.SDNArray also
		for i := range reader.SDNs {
			sdn := reader.SDNs[i]
			if strings.Contains(sdn.SDNName, nameSlug) {
				answers = append(answers, sdn)
			}
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answers); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func searchByAltName(logger log.Logger, reader *ofac.Reader, altSlug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answers []*ofac.AlternateIdentity
		for i := range reader.AlternateIdentities {
			alt := reader.AlternateIdentities[i]
			if strings.Contains(alt.AlternateName, altSlug) {
				answers = append(answers, &alt)
			}
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answers); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
