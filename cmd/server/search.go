// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func addSearchRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("GET").Path("/search/address").HandlerFunc(searchByAddress(logger)) // TODO(adam): GET's ?
	r.Methods("GET").Path("/search/name").HandlerFunc(searchByName(logger))
	r.Methods("GET").Path("/search/alt").HandlerFunc(searchByAltName(logger))
	// r.Methods("GET").Path("/search/company").HandlerFunc(searchByAddress()) // TODO
}

type addressSearchRequest struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
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

func searchByAddress(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		_ = readAddressSearchRequest(r.URL) // TODO(adam): do something with req

		w.WriteHeader(http.StatusOK)

		addresses := []*ofac.Address{
			{ // Real OFAC entry -- 173,129,"Ibex House, The Minories","London EC3N 1DY","United Kingdom",-0-
				EntityID:                    "173",
				AddressID:                   "129",
				Address:                     "Ibex House, The Minories",
				CityStateProvincePostalCode: "London EC3N 1DY",
				Country:                     "United Kingdom",
			},
		}
		if err := json.NewEncoder(w).Encode(addresses); err != nil {
			moovhttp.Problem(w, err) // TODO(adam): JSON errors should moovhttp.InternalError (wrapped, see auth's http.go)
			return
		}
	}
}

func searchByName(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		// ?name=foo // TODO(adam): grab from *http.Request

		w.WriteHeader(http.StatusOK)

		sdns := []*ofac.SDN{
			{ // Real OFAC entry
				EntityID: "2676",
				SDNName:  "AL ZAWAHIRI, Dr. Ayman",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Operational and Military Leader of JIHAD GROUP",
				Remarks:  "DOB 19 Jun 1951; POB Giza, Egypt; Passport 1084010 (Egypt); alt. Passport 19820215; Operational and Military Leader of JIHAD GROUP.",
			},
		}
		if err := json.NewEncoder(w).Encode(sdns); err != nil {
			moovhttp.Problem(w, err)
			return
		}
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

		alts := []*ofac.AlternateIdentity{
			{ // Real OFAC entry
				EntityID:      "559",
				AlternateID:   "481",
				AlternateType: "aka",
				AlternateName: "CIMEX",
			},
		}
		if err := json.NewEncoder(w).Encode(alts); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
