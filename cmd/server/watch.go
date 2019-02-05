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
	AuthToken string `json:"authToken"`
	Webhook   string `json:"webhook"`
}

type watchRepository interface {
	// addCustomerWatch takes a customerId (EntityID), creates a watch and
	// returns the watchId.
	addCustomerWatch(customerId string, params watchRequest) (string, error)
	removeCustomerWatch(customerId string, watchId string) error

	getWatchesCursor(batchSize int) *watchCursor

	// addCustomerNameWatch takes a customerId (EntityID), creates a watch and returns the watchId.
	addCustomerNameWatch(name string, webhook string, authToken string) (string, error)
	removeCustomerNameWatch(watchId string) error
}

type sqliteWatchRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqliteWatchRepository) close() error {
	return r.db.Close()
}

func (r *sqliteWatchRepository) getWatchesCursor(batchSize int) *watchCursor {
	return &watchCursor{
		batchSize: batchSize,
		db:        r.db,
	}
}

func (r *sqliteWatchRepository) addCustomerWatch(customerId string, params watchRequest) (string, error) {
	if customerId == "" {
		return "", errNoCustomerId
	}
	id := base.ID()

	query := `insert or ignore into customer_watches (id, customer_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, customerId, params.Webhook, params.AuthToken, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *sqliteWatchRepository) removeCustomerWatch(customerId string, watchId string) error {
	if watchId == "" {
		return errNoWatchId
	}

	query := `update customer_watches set deleted_at = ? where customer_id = ? and id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), customerId, watchId)
	return err
}

func (r *sqliteWatchRepository) addCustomerNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert or ignore into customer_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	id := base.ID()
	_, err = stmt.Exec(id, name, webhook, authToken, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *sqliteWatchRepository) removeCustomerNameWatch(watchId string) error {
	if watchId == "" {
		return errNoWatchId
	}

	query := `update customer_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchId)
	return err
}

type watch struct {
	id, customerId string
	webhook        string
	authToken      string
}

type watchCursor struct {
	batchSize int
	db        *sql.DB

	// newerThan represents the minimum (oldest) created_at value
	// to return in the batch.
	//
	// This value starts at "zero time" (an empty time.Time) and progresses
	// towards time.Now() with each batch by being set to the batch's newest time.
	newerThan time.Time
}

func (cur *watchCursor) Next() ([]watch, error) {
	query := `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	stmt, err := cur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cur.newerThan, cur.batchSize)
	if err != nil {
		return nil, err
	}

	max := cur.newerThan

	var watches []watch
	for rows.Next() {
		var createdAt time.Time
		var watch watch
		if err := rows.Scan(&watch.id, &watch.customerId, &watch.webhook, &watch.authToken, &createdAt); err == nil {
			watches = append(watches, watch)
		}
		if createdAt.After(max) {
			// advance max to newest time
			max = createdAt
		}
	}
	cur.newerThan = max

	return watches, nil
}
