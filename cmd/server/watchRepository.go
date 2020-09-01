package main

import (
	"database/sql"
	"github.com/go-kit/kit/log"
	"github.com/moov-io/base"
	"reflect"
	"time"
)

type watchRequest struct {
	AuthToken string `json:"authToken"`
	Webhook   string `json:"webhook"`
}

// watchRepository holds information about each company and/or customer that another service wants notifications
// for every time we re-download data.
type watchRepository interface {
	// getWatchesCursor returns a watchCursor which traverses both customer and company watches
	getWatchesCursor(logger log.Logger, batchSize int) *watchCursor

	// Company watches
	addCompanyWatch(companyID string, params watchRequest) (string, error)
	addCompanyNameWatch(name string, webhook string, authToken string) (string, error)
	removeCompanyWatch(companyID string, watchID string) error
	removeCompanyNameWatch(watchID string) error

	// Customer watches
	addCustomerWatch(customerID string, params watchRequest) (string, error)
	addCustomerNameWatch(name string, webhook string, authToken string) (string, error)
	removeCustomerWatch(customerID string, watchID string) error
	removeCustomerNameWatch(watchID string) error
	close() error
}

const (
	genericInsertCompanyWatches      = `insert into company_watches (id, company_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	genericUpdateCompanyWatches      = `update company_watches set deleted_at = ? where company_id = ? and id = ? and deleted_at is null`
	genericInsertCompanyNameWatches  = `insert into company_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
	genericUpdateCompanyNameWatches  = `update company_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	genericInsertCustomerWatches     = `insert into customer_watches (id, customer_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	genericUpdateCustomerWatches     = `update customer_watches set deleted_at = ? where customer_id = ? and id = ? and deleted_at is null`
	genericInsertCustomerNameWatches = `insert into customer_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
	genericUpdateCustomerNameWatches = `update customer_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	genericSelectCompanyWatches      = `select id, company_id, webhook, auth_token, created_at from company_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	genericSelectCompanyNameWatches  = `select id, name, webhook, auth_token, created_at from company_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	genericSelectCustomerWatches     = `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
	genericSelectCustomerNameWatches = `select id, name, webhook, auth_token, created_at from customer_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`

	postgresInsertCompanyWatches      = `insert into company_watches (id, company_id, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5)`
	postgresUpdateCompanyWatches      = `update company_watches set deleted_at = $1 where company_id = $2 and id = $3 and deleted_at is null`
	postgresInsertCompanyNameWatches  = `insert into company_name_watches (id, name, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5);`
	postgresUpdateCompanyNameWatches  = `update company_name_watches set deleted_at = $1 where id = $2 and deleted_at is null`
	postgresInsertCustomerWatches     = `insert into customer_watches (id, customer_id, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5)`
	postgresUpdateCustomerWatches     = `update customer_watches set deleted_at = $1 where customer_id = $2 and id = $3 and deleted_at is null`
	postgresInsertCustomerNameWatches = `insert into customer_name_watches (id, name, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5);`
	postgresUpdateCustomerNameWatches = `update customer_name_watches set deleted_at = $1 where id = $2 and deleted_at is null`
	postgresSelectCompanyWatches      = `select id, company_id, webhook, auth_token, created_at from company_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	postgresSelectCompanyNameWatches  = `select id, name, webhook, auth_token, created_at from company_name_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	postgresSelectCustomerWatches     = `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	postgresSelectCustomerNameWatches = `select id, name, webhook, auth_token, created_at from customer_name_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
)

type genericSQLWatchRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *genericSQLWatchRepository) close() error {
	return r.db.Close()
}

func (r *genericSQLWatchRepository) getWatchesCursor(logger log.Logger, batchSize int) *watchCursor {
	return &watchCursor{
		batchSize: batchSize,
		db:        r.db,
		logger:    logger,
	}
}

// Company methods - for generic sql (SQLite and MySQL)

func (r *genericSQLWatchRepository) addCompanyWatch(companyID string, params watchRequest) (string, error) {
	if companyID == "" {
		return "", errNoCompanyID
	}
	id := base.ID()
	query := genericInsertCompanyWatches
	switch dbType {
	case `postgres`:
		query = postgresInsertCompanyWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, companyID, params.Webhook, params.AuthToken, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *genericSQLWatchRepository) removeCompanyWatch(companyID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := genericUpdateCompanyWatches
	switch dbType {
	case `postgres`:
		query = postgresUpdateCompanyWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), companyID, watchID)
	return err
}

func (r *genericSQLWatchRepository) addCompanyNameWatch(name string, webhook string, authToken string) (string, error) {
	query := genericInsertCompanyNameWatches
	switch dbType {
	case `postgres`:
		query = postgresInsertCompanyNameWatches
	}
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

func (r *genericSQLWatchRepository) removeCompanyNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := genericUpdateCompanyNameWatches
	switch dbType {
	case `postgres`:
		query = postgresUpdateCompanyNameWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchID)
	return err
}

// Customer methods - SQLite

func (r *genericSQLWatchRepository) addCustomerWatch(customerID string, params watchRequest) (string, error) {
	r.logger.Log("Repository is of type %s", reflect.TypeOf(r).String())
	if customerID == "" {
		return "", errNoCustomerID
	}
	id := base.ID()

	query := genericInsertCustomerWatches
	switch dbType {
	case `postgres`:
		query = postgresInsertCustomerWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, customerID, params.Webhook, params.AuthToken, time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *genericSQLWatchRepository) removeCustomerWatch(customerID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := genericUpdateCustomerWatches
	switch dbType {
	case `postgres`:
		query = postgresUpdateCustomerWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), customerID, watchID)
	return err
}

func (r *genericSQLWatchRepository) addCustomerNameWatch(name string, webhook string, authToken string) (string, error) {
	query := genericInsertCustomerNameWatches
	switch dbType {
	case `postgres`:
		query = postgresInsertCustomerNameWatches
	}
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

func (r *genericSQLWatchRepository) removeCustomerNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := genericUpdateCustomerNameWatches
	switch dbType {
	case `postgres`:
		query = postgresUpdateCustomerNameWatches
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchID)
	return err
}

// Cursor Methods

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
	query := genericSelectCompanyWatches
	switch dbType {
	case `postgres`:
		query = postgresSelectCompanyWatches
	}
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
		if err := rows.Scan(&watch.id, &watch.companyID, &watch.webhook, &watch.authToken, &createdAt); err == nil {
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
	query := genericSelectCompanyNameWatches
	switch dbType {
	case `postgres`:
		query = postgresSelectCompanyNameWatches
	}

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
	query := genericSelectCustomerWatches
	switch dbType {
	case `postgres`:
		query = postgresSelectCustomerWatches
	}
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
		if err := rows.Scan(&watch.id, &watch.customerID, &watch.webhook, &watch.authToken, &createdAt); err == nil {
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
	query := genericSelectCustomerNameWatches
	switch dbType {
	case `postgres`:
		query = postgresSelectCustomerNameWatches
	}

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
