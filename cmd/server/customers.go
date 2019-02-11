// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoAuthToken  = errors.New("no authToken provided for webhook")
	errNoCustomerId = errors.New("no customerId found")
	errNoNameParam  = errors.New("no name parameter found")
	errNoUserId     = errors.New("no userId (X-User-Id header) found")
)

type Customer struct {
	ID        string                    `json:"id"`
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Comments  []*ofac.SDNComments       `json:"comments"`

	// Blocked represents if the userId making this request has manually blocked this SDN
	Blocked bool `json:"blocked"`
}

type customerWatchResponse struct {
	WatchID string `json:"watchId"`
}

func addCustomerRoutes(logger log.Logger, r *mux.Router, searcher *searcher, custRepo *sqliteCustomerRepository, watchRepo *sqliteWatchRepository) {
	r.Methods("GET").Path("/customers/{customerId}").HandlerFunc(getCustomer(logger, searcher, custRepo))
	r.Methods("PUT").Path("/customers/{customerId}").HandlerFunc(updateCustomerStatus(logger, searcher, custRepo))

	r.Methods("POST").Path("/customers/{customerId}/watch").HandlerFunc(addCustomerWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/customers/{customerId}/watch/{watchId}").HandlerFunc(removeCustomerWatch(logger, searcher, watchRepo))

	r.Methods("POST").Path("/customers/watch").HandlerFunc(addCustomerNameWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/customers/watch/{watchId}").HandlerFunc(removeCustomerNameWatch(logger, searcher, watchRepo))
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
func getCustomerById(id string, userId string, searcher *searcher, custRepo customerRepository) (*Customer, error) {
	sdn := searcher.FindSDN(id)
	if sdn == nil {
		return nil, fmt.Errorf("Customer %s not found", id)
	}
	if userId != "" && custRepo == nil {
		return nil, errors.New("non-empty but nil customerRepository provided - logic bug")
	}
	var blocked bool
	if custRepo != nil {
		block, err := custRepo.isCustomerBlocked(sdn.EntityID, userId)
		if err != nil {
			return nil, fmt.Errorf("problem reading Customer=%s block status: %v", id, err)
		}
		blocked = block
	}
	return &Customer{
		ID:        id,
		SDN:       sdn,
		Addresses: searcher.FindAddresses(100, id),
		Alts:      searcher.FindAlts(100, id),
		Blocked:   blocked,
	}, nil
}

type customerRepository interface {
	addCustomerBlock(customerId string, userId string) error
	isCustomerBlocked(customerId string, userId string) (bool, error)
	removeCustomerBlock(customerId string, userId string) error
}

type sqliteCustomerRepository struct {
	db *sql.DB
}

func (r *sqliteCustomerRepository) close() error {
	return r.db.Close()
}

func (r *sqliteCustomerRepository) addCustomerBlock(customerId string, userId string) error {
	query := `insert or ignore into customer_blocks (customer_id, user_id, created_at) values (?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customerId, userId, time.Now())
	return err
}

func (r *sqliteCustomerRepository) isCustomerBlocked(customerId string, userId string) (bool, error) {
	if customerId == "" {
		return false, errors.New("isCustomerBlocked: no Customer.ID")
	}
	if userId == "" {
		return false, nil // no userId, so no overrides possible
	}

	query := `select customer_id from customer_blocks where customer_id = ? and user_id = ? and deleted_at is null limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return true, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(customerId, userId)

	var custId string // read and check Customer.ID matches
	if err := row.Scan(&custId); err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return true, err
	}
	return strings.EqualFold(custId, customerId), nil
}

func (r *sqliteCustomerRepository) removeCustomerBlock(customerId string, userId string) error {
	query := `update customer_blocks set deleted_at = ? where customer_id = ? and user_id = ? and deleted_at is null;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(time.Now(), customerId, userId)
	return err
}

func getCustomer(logger log.Logger, searcher *searcher, custRepo customerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		id := getCustomerId(w, r)
		if id == "" {
			return
		}
		userId := moovhttp.GetUserId(r)
		customer, err := getCustomerById(id, userId, searcher, custRepo)
		if err != nil {
			moovhttp.Problem(w, err)
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

func updateCustomerStatus(logger log.Logger, searcher *searcher, custRepo customerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		custId := getCustomerId(w, r)
		userId := moovhttp.GetUserId(r)
		if userId == "" {
			moovhttp.Problem(w, errNoUserId)
			return
		}

		var req customerStatusRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		req.Status = strings.ToLower(strings.TrimSpace(req.Status))
		if req.Status == "blocked" {
			if err := custRepo.addCustomerBlock(custId, userId); err != nil {
				moovhttp.Problem(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		if req.Status == "unblock" || req.Status == "unblocked" {
			if err := custRepo.removeCustomerBlock(custId, userId); err != nil {
				moovhttp.Problem(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		// no match, 400
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{}"))
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
