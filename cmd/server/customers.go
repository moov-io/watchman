// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// GET /customers/:id
// - get customer information and matches

// POST /customers/watch?name=...
// - monitor customer by name, re-parse on each search

// POST /customers/:id/watch
// - monitor customer

// PUT /customers/:id
// - mark customer as blocked or unblocked

// DELETE /customers/:id/watch
// - stop watching customer

// DELETE /customers/watch?name=...
// - stop watching customer name

type Customer struct {
	ID        string                    `json:"id"`
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Comments  []*ofac.SDNComments       `json:"comments"`
}

type customerWatchResponse struct {
	WatchID string `json:"watchId"`
}

func addCustomerRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("GET").Path("/customers/{id}").HandlerFunc(getCustomer(logger))
	r.Methods("PUT").Path("/customers/{id}").HandlerFunc(updateCustomer(logger))

	r.Methods("POST").Path("/customers/{id}/watch").HandlerFunc(addCustomerWatch(logger))
	r.Methods("DELETE").Path("/customers/{id}/watch").HandlerFunc(removeCustomerWatch(logger))

	r.Methods("POST").Path("/customers/watch").HandlerFunc(addCustomerNameWatch(logger))
	r.Methods("DELETE").Path("/customers/watch/{watchId}").HandlerFunc(removeCustomerNameWatch(logger))
}

func getCustomer(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		customer := Customer{
			ID: "13ou1fohfkajfah", // "random"
			SDN: &ofac.SDN{
				EntityID: "2676",
				SDNName:  "AL ZAWAHIRI, Dr. Ayman",
				SDNType:  "individual",
				Program:  "SDGT] [SDT",
				Title:    "Operational and Military Leader of JIHAD GROUP",
				Remarks:  "DOB 19 Jun 1951; POB Giza, Egypt; Passport 1084010 (Egypt); alt. Passport 19820215; Operational and Military Leader of JIHAD GROUP.",
			},
			Addresses: []*ofac.Address{
				{
					EntityID:                    "173",
					AddressID:                   "129",
					Address:                     "Ibex House, The Minories",
					CityStateProvincePostalCode: "London EC3N 1DY",
					Country:                     "United Kingdom",
				},
			},
			Alts: []*ofac.AlternateIdentity{
				{
					EntityID:      "559",
					AlternateID:   "481",
					AlternateType: "aka",
					AlternateName: "CIMEX",
				},
			},
		}
		if err := json.NewEncoder(w).Encode(customer); err != nil {
			moovhttp.Problem(w, err) // TODO(adam): replace with wrapped moovhttp.InternalError
			return
		}
	}
}

func addCustomerNameWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		// TODO: read ?name=... param

		if err := json.NewEncoder(w).Encode(customerWatchResponse{"cust-name-watch"}); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func addCustomerWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(customerWatchResponse{"cust-watch"}); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func updateCustomer(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK) // TODO
	}
}

func removeCustomerWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK) // TODO
	}
}

func removeCustomerNameWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK) // TODO
	}
}
