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

	"github.com/cardonator/ofac"
	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoAuthToken  = errors.New("no authToken provided for webhook")
	errNoCustomerID = errors.New("no Customer ID found")
	errNoNameParam  = errors.New("no name parameter found")
	errNoUserID     = errors.New("no userID (X-User-Id header) found")
)

// Customer is an individual on one or more SDN list(s)
type Customer struct {
	ID string `json:"id"`
	// Federal Data
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Comments  []*ofac.SDNComments       `json:"comments"`
	// Metadata
	Status *CustomerStatus `json:"status"`
	Match  float64         `json:"match,omitempty"`
}

// CustomerBlockStatus can be either CustomerUnsafe or CustomerException
type CustomerBlockStatus string

const (
	// CustomerUnsafe customers have been manually marked to block all transactions with
	CustomerUnsafe CustomerBlockStatus = "unsafe"
	// CustomerException customers have been manually marked to allow transactions with
	CustomerException CustomerBlockStatus = "exception"
)

// CustomerStatus represents a userID's manual override of an SDN
type CustomerStatus struct {
	UserID    string              `json:"userID"`
	Note      string              `json:"note"`
	Status    CustomerBlockStatus `json:"block"`
	CreatedAt time.Time           `json:"createdAt"`
}

type customerWatchResponse struct {
	WatchID string `json:"watchID"`
}

func addCustomerRoutes(logger log.Logger, r *mux.Router, searcher *searcher, custRepo *sqliteCustomerRepository, watchRepo *sqliteWatchRepository) {
	r.Methods("GET").Path("/customers/{customerID}").HandlerFunc(getCustomer(logger, searcher, custRepo))
	r.Methods("PUT").Path("/customers/{customerID}").HandlerFunc(updateCustomerStatus(logger, searcher, custRepo))

	r.Methods("POST").Path("/customers/{customerID}/watch").HandlerFunc(addCustomerWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/customers/{customerID}/watch/{watchID}").HandlerFunc(removeCustomerWatch(logger, searcher, watchRepo))

	r.Methods("POST").Path("/customers/watch").HandlerFunc(addCustomerNameWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/customers/watch/{watchID}").HandlerFunc(removeCustomerNameWatch(logger, searcher, watchRepo))
}

func getCustomerID(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["customerID"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoCustomerID)
		return ""
	}
	return v
}

// getCustomerID returns an SDN for an individual and any addresses or alt names by the entity id provided.a
func getCustomerByID(id string, searcher *searcher, custRepo customerRepository) (*Customer, error) {
	sdn := searcher.FindSDN(id)
	if sdn == nil {
		return nil, fmt.Errorf("Customer %s not found", id)
	}
	if !strings.EqualFold(sdn.SDNType, "individual") {
		// SDN wasn't an individual, so don't return it for method that only returns individuals
		return nil, nil
	}
	if custRepo == nil {
		return nil, errors.New("nil customerRepository provided - logic bug")
	}
	status, err := custRepo.getCustomerStatus(sdn.EntityID)
	if err != nil {
		return nil, fmt.Errorf("problem reading Customer=%s block status: %v", id, err)
	}
	return &Customer{
		ID:        id,
		SDN:       sdn,
		Addresses: searcher.FindAddresses(100, id),
		Alts:      searcher.FindAlts(100, id),
		Status:    status,
	}, nil
}

// customerRepository holds the current status (i.e. unsafe or exception) for a given customer
// (individual) and is expected to save metadata about each time the status is changed.
type customerRepository interface {
	getCustomerStatus(customerID string) (*CustomerStatus, error)
	upsertCustomerStatus(customerID string, status *CustomerStatus) error
}

type sqliteCustomerRepository struct {
	db *sql.DB
}

func (r *sqliteCustomerRepository) close() error {
	return r.db.Close()
}

func (r *sqliteCustomerRepository) getCustomerStatus(customerID string) (*CustomerStatus, error) {
	if customerID == "" {
		return nil, errors.New("getCustomerStatus: no Customer.ID")
	}
	query := `select user_id, note, status, created_at from customer_status where customer_id = ? and deleted_at is null order by created_at desc limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(customerID)

	var status CustomerStatus
	err = row.Scan(&status.UserID, &status.Note, &status.Status, &status.CreatedAt)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, fmt.Errorf("getCustomerStatus: %v", err)
	}
	if status.UserID == "" {
		return nil, nil // not found
	}
	return &status, nil
}

func (r *sqliteCustomerRepository) upsertCustomerStatus(customerID string, status *CustomerStatus) error {
	query := `insert or replace into customer_status (customer_id, user_id, note, status, created_at) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customerID, status.UserID, status.Note, status.Status, status.CreatedAt)
	return err
}

func getCustomer(logger log.Logger, searcher *searcher, custRepo customerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		id := getCustomerID(w, r)
		if id == "" {
			return
		}
		customer, err := getCustomerByID(id, searcher, custRepo)
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
		watchID, err := repo.addCustomerNameWatch(name, webhook, req.AuthToken)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customerWatchResponse{watchID}); err != nil {
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

		customerID := getCustomerID(w, r)
		watchID, err := repo.addCustomerWatch(customerID, req)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(customerWatchResponse{watchID}); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

type customerStatusRequest struct {
	Notes string `json:"notes,omitempty"`

	// Status represents a manual exception or unsafe designation
	Status string `json:"status"`
}

func updateCustomerStatus(logger log.Logger, searcher *searcher, custRepo customerRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		custID := getCustomerID(w, r)
		userID := moovhttp.GetUserId(r)
		if userID == "" {
			moovhttp.Problem(w, errNoUserID)
			return
		}

		var req customerStatusRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		status := CustomerBlockStatus(strings.ToLower(strings.TrimSpace(req.Status)))
		switch status {
		case CustomerUnsafe, CustomerException:
			custStatus := &CustomerStatus{
				UserID:    userID,
				Note:      req.Notes,
				Status:    status,
				CreatedAt: time.Now(),
			}
			if err := custRepo.upsertCustomerStatus(custID, custStatus); err != nil {
				moovhttp.Problem(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "unknown status"}`))
		}
	}
}

func removeCustomerWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		customerID, watchID := getCustomerID(w, r), getWatchID(w, r)
		if customerID == "" || watchID == "" {
			return
		}
		if err := repo.removeCustomerWatch(customerID, watchID); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func removeCustomerNameWatch(logger log.Logger, searcher *searcher, repo *sqliteWatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		watchID := getWatchID(w, r)
		if watchID == "" {
			return
		}
		if err := repo.removeCustomerNameWatch(watchID); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
