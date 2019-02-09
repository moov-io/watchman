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
	errNoAuthToken  = errors.New("no authToken provided for webhook")
	errNoCustomerId = errors.New("no customerId found")
	errNoNameParam  = errors.New("no name parameter found")
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

func addCustomerRoutes(logger log.Logger, r *mux.Router, searcher *searcher, repo *sqliteWatchRepository) {
	r.Methods("GET").Path("/customers/{customerId}").HandlerFunc(getCustomer(logger, searcher))
	r.Methods("PUT").Path("/customers/{customerId}").HandlerFunc(updateCustomerStatus(logger, searcher))

	r.Methods("POST").Path("/customers/{customerId}/watch").HandlerFunc(addCustomerWatch(logger, searcher, repo))
	r.Methods("DELETE").Path("/customers/{customerId}/watch/{watchId}").HandlerFunc(removeCustomerWatch(logger, searcher, repo))

	r.Methods("POST").Path("/customers/watch").HandlerFunc(addCustomerNameWatch(logger, searcher, repo))
	r.Methods("DELETE").Path("/customers/watch/{watchId}").HandlerFunc(removeCustomerNameWatch(logger, searcher, repo))
}

func getCustomerId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["customerId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoCustomerId)
		return ""
	}
	return v
}

// getCustomerId returns an SDN and any addresses or alt names by the entity id provided.
func getCustomerById(id string, searcher *searcher) *Customer {
	sdn := searcher.FindSDN(id)
	if sdn == nil {
		return nil
	}
	return &Customer{
		ID:        id,
		SDN:       sdn,
		Addresses: searcher.FindAddresses(100, id),
		Alts:      searcher.FindAlts(100, id),
	}
}

func getCustomer(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		id := getCustomerId(w, r)
		if id == "" {
			return
		}
		customer := getCustomerById(id, searcher)
		if customer == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customer); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func addCustomerNameWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		name := r.URL.Query().Get("name")
		if name == "" {
			moovhttp.Problem(w, errNoNameParam)
			return
		}

		var req watchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if req.AuthToken == "" {
			moovhttp.Problem(w, errNoAuthToken)
			return
		}
		webhook, err := validateWebhook(req.Webhook)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		watchId, err := repo.addCustomerNameWatch(name, webhook, req.AuthToken)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customerWatchResponse{watchId}); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func addCustomerWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req watchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if req.AuthToken == "" {
			moovhttp.Problem(w, errNoAuthToken)
			return
		}
		webhook, err := validateWebhook(req.Webhook)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		req.Webhook = webhook

		customerId := getCustomerId(w, r)
		watchId, err := repo.addCustomerWatch(customerId, req)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customerWatchResponse{watchId}); err != nil {
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
		w = wrapResponseWriter(logger, w, r)

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

func removeCustomerWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		customerId, watchId := getCustomerId(w, r), getWatchId(w, r)
		if customerId == "" || watchId == "" {
			return
		}
		if err := repo.removeCustomerWatch(customerId, watchId); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func removeCustomerNameWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		watchId := getWatchId(w, r)
		if watchId == "" {
			return
		}
		if err := repo.removeCustomerNameWatch(watchId); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
