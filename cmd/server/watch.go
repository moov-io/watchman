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

// watchRepository holds information about each company and/or customer that another service wants notifications
// for every time we re-download OFAC data.
type watchRepository interface {
	// getWatchesCursor returns a watchCursor which traverses both customer and company watches
	getWatchesCursor(logger log.Logger, batchSize int) *watchCursor

	// Company watches
	addCompanyWatch(companyId string, params watchRequest) (string, error)
	addCompanyNameWatch(name string, webhook string, authToken string) (string, error)
	removeCompanyWatch(companyId string, watchId string) error
	removeCompanyNameWatch(watchId string) error

	// Customer watches
	addCustomerWatch(customerId string, params watchRequest) (string, error)
	addCustomerNameWatch(name string, webhook string, authToken string) (string, error)
	removeCustomerWatch(customerId string, watchId string) error
	removeCustomerNameWatch(watchId string) error
}

type sqliteWatchRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqliteWatchRepository) close() error {
	return r.db.Close()
}

func (r *sqliteWatchRepository) getWatchesCursor(logger log.Logger, batchSize int) *watchCursor {
	return &watchCursor{
		batchSize: batchSize,
		db:        r.db,
		logger:    logger,
	}
}

// Company methods

func (r *sqliteWatchRepository) addCompanyWatch(companyId string, params watchRequest) (string, error) {
	if companyId == "" {
		return "", errNoCompanyId
	}
	id := base.ID()

	query := `insert or ignore into company_watches (id, company_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, companyId, params.Webhook, params.AuthToken, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *sqliteWatchRepository) removeCompanyWatch(companyId string, watchId string) error {
	if watchId == "" {
		return errNoWatchId
	}

	query := `update company_watches set deleted_at = ? where company_id = ? and id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), companyId, watchId)
	return err
}

func (r *sqliteWatchRepository) addCompanyNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert or ignore into company_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
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

func (r *sqliteWatchRepository) removeCompanyNameWatch(watchId string) error {
	if watchId == "" {
		return errNoWatchId
	}

	query := `update company_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchId)
	return err
}

// Customer methods

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
	id                       string
	customerId, customerName string
	companyId, companyName   string
	webhook                  string
	authToken                string
}

type watchCursor struct {
	batchSize int
	db        *sql.DB

	logger log.Logger

	// '*NewerThan' values represent the minimum (oldest) created_at value to return in the batch.
	//
	// These values start at "zero time" (an empty time.Time) and progresses towards time.Now()
	// with each batch by being set to the batch's newest time.
	customerNewerThan, customerNameNewerThan time.Time
	companyNewerThan, companyNameNewerThan   time.Time
}

// Next returns a batch of watches that will be sent off to their respective webhook URL.
func (cur *watchCursor) Next() ([]watch, error) {
	var watches []watch
	limit := cur.batchSize / 4 // 4 SQL queries
	if cur.batchSize < 4 {
		limit = 1 // return one if batchSize is invalid
	}

	// Companies
	companyWatches, err := cur.getCompanyBatch(limit)
	if err != nil {
		cur.logger.Log("watchCursor", "problem reading company watches", "error", err)
	}
	watches = append(watches, companyWatches...)

	companyNameWatches, err := cur.getCompanyNameBatch(limit)
	if err != nil {
		cur.logger.Log("watchCursor", "problem reading company name watches", "error", err)
	}
	watches = append(watches, companyNameWatches...)

	// Customers
	customerWatches, err := cur.getCustomerBatch(limit)
	if err != nil {
		cur.logger.Log("watchCursor", "problem reading customer watches", "error", err)
	}
	watches = append(watches, customerWatches...)

	customerNameWatches, err := cur.getCustomerNameBatch(limit)
	if err != nil {
		cur.logger.Log("watchCursor", "problem reading customer name watches", "error", err)
	}
	watches = append(watches, customerNameWatches...)

	return watches, nil
}

func (cur *watchCursor) getCompanyBatch(limit int) ([]watch, error) {
	query := `select id, company_id, webhook, auth_token, created_at from company_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	stmt, err := cur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cur.companyNewerThan, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	max := cur.companyNewerThan

	var watches []watch
	for rows.Next() {
		var createdAt time.Time
		var watch watch
		if err := rows.Scan(&watch.id, &watch.companyId, &watch.webhook, &watch.authToken, &createdAt); err == nil {
			watches = append(watches, watch)
		}
		if createdAt.After(max) {
			// advance max to newest time
			max = createdAt
		}
	}
	cur.companyNewerThan = max

	return watches, rows.Err()
}

func (cur *watchCursor) getCompanyNameBatch(limit int) ([]watch, error) {
	query := `select id, name, webhook, auth_token, created_at from company_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	stmt, err := cur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cur.companyNameNewerThan, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	max := cur.companyNameNewerThan

	var watches []watch
	for rows.Next() {
		var createdAt time.Time
		var watch watch
		if err := rows.Scan(&watch.id, &watch.companyName, &watch.webhook, &watch.authToken, &createdAt); err == nil {
			watches = append(watches, watch)
		}
		if createdAt.After(max) {
			// advance max to newest time
			max = createdAt
		}
	}
	cur.companyNameNewerThan = max

	return watches, rows.Err()
}

func (cur *watchCursor) getCustomerBatch(limit int) ([]watch, error) {
	query := `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	stmt, err := cur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cur.customerNewerThan, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	max := cur.customerNewerThan

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
	cur.customerNewerThan = max

	return watches, rows.Err()
}

func (cur *watchCursor) getCustomerNameBatch(limit int) ([]watch, error) {
	query := `select id, name, webhook, auth_token, created_at from customer_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	stmt, err := cur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cur.customerNameNewerThan, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	max := cur.customerNameNewerThan

	var watches []watch
	for rows.Next() {
		var createdAt time.Time
		var watch watch
		if err := rows.Scan(&watch.id, &watch.customerName, &watch.webhook, &watch.authToken, &createdAt); err == nil {
			watches = append(watches, watch)
		}
		if createdAt.After(max) {
			// advance max to newest time
			max = createdAt
		}
	}
	cur.customerNameNewerThan = max

	return watches, rows.Err()
}
