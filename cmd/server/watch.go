// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoWatchId = errors.New("no watchId found")
)

func getWatchId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["watchId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoWatchId)
		return ""
	}
	return v
}

type watchRequest struct {
	Webhook string `json:"webhook"`
}

type watchRepository interface {
	// addCustomerWatch takes a customerId (EntityID), creates a watch and
	// returns the watchId.
	addCustomerWatch(customerId string) (string, error)
	removeCustomerWatch(watchId string) error

	// addCustomerNameWatch takes a customerId (EntityID), creates a watch and
	// returns the watchId.
	addCustomerNameWatch(name string) (string, error)
	removeCustomerNameWatch(watchId string) error
}

type sqliteWatchRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r sqliteWatchRepository) close() error {
	return r.db.Close()
}

func (r sqliteWatchRepository) addCustomerWatch(customerId string, params watchRequest) (string, error) {
	if customerId == "" {
		return "", errNoCustomerId
	}
	id := base.ID()

	query := `insert or ignore into customer_watches (id, customer_id, webhook, created_at) values (?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, customerId, params.Webhook, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r sqliteWatchRepository) removeCustomerWatch(watchId string) error {
	if watchId == "" {
		return errNoWatchId
	}

	query := `update customer_watches set deleted_at = ? where id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchId)
	if err != nil {
		return err
	}
	return nil
}

func (r sqliteWatchRepository) addCustomerNameWatch(name string) (string, error) {
	return "", nil
}

func (r sqliteWatchRepository) removeCustomerNameWatch(watchId string) error {
	return nil
}
