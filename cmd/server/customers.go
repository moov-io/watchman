// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoCustomerId = errors.New("no customerId found in path")
)

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

func addCustomerRoutes(logger log.Logger, r *mux.Router, searcher *searcher) {
	r.Methods("GET").Path("/customers/{customerId}").HandlerFunc(getCustomer(logger, searcher))
	r.Methods("PUT").Path("/customers/{customerId}").HandlerFunc(updateCustomerStatus(logger, searcher))

	r.Methods("POST").Path("/customers/{customerId}/watch").HandlerFunc(addCustomerWatch(logger, searcher))
	r.Methods("DELETE").Path("/customers/{customerId}/watch").HandlerFunc(removeCustomerWatch(logger, searcher))

	r.Methods("POST").Path("/customers/watch").HandlerFunc(addCustomerNameWatch(logger, searcher))
	r.Methods("DELETE").Path("/customers/watch/{watchId}").HandlerFunc(removeCustomerNameWatch(logger, searcher))
}

func getCustomerId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["customerId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoCustomerId)
		return ""
	}
	return v
}

func getCustomer(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		id := getCustomerId(w, r)
		limit := extractSearchLimit(r)

		sdns := searcher.FindSDNs(1, func(sdn *SDN) bool {
			return sdn.SDN.EntityID == id
		})
		if len(sdns) != 1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Find customer and linked data
		customer := Customer{
			ID:  id,
			SDN: sdns[0],
			Addresses: searcher.FindAddresses(limit, func(add *Address) bool {
				return add.Address.EntityID == id
			}),
			Alts: searcher.FindAlts(limit, func(alt *Alt) bool {
				return alt.AlternateIdentity.EntityID == id
			}),
		}
		if err := json.NewEncoder(w).Encode(customer); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func addCustomerNameWatch(logger log.Logger, searcher *searcher) http.HandlerFunc {
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

func addCustomerWatch(logger log.Logger, searcher *searcher) http.HandlerFunc {
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

type customerStatusRequest struct {
	// Status represents a manual 'Blocked' value for a customer.
	Status string `json:"status"` // TODO(adam): better name for Default ?
}

func updateCustomerStatus(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		var req customerStatusRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		req.Status = strings.ToLower(strings.TrimSpace(req.Status))
		switch req.Status {
		case "blocked":
			// update customer status to blocked
		default:
			if req.Status == "" {
				moovhttp.Problem(w, errors.New("no status provided"))
				return
			}
			// remove blocked from customer status
		}
		w.WriteHeader(http.StatusOK)
	}
}

func removeCustomerWatch(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK) // TODO
	}
}

func removeCustomerNameWatch(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK) // TODO
	}
}
