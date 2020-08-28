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

////////////////////////////
//sqlite implementation
////////////////////////////
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

// Company methods -SQLite

func (r *sqliteWatchRepository) addCompanyWatch(companyID string, params watchRequest) (string, error) {
	if companyID == "" {
		return "", errNoCompanyID
	}
	id := base.ID()

	query := `insert into company_watches (id, company_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	return insertCompanyWatch(companyID, params, err, stmt, id)
}

func (r *sqliteWatchRepository) removeCompanyWatch(companyID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update company_watches set deleted_at = ? where company_id = ? and id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCompanyWatch(companyID, watchID, err, stmt)
}

func (r *sqliteWatchRepository) addCompanyNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert into company_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	return insertCompanyNameWatch(name, webhook, authToken, err, stmt)
}

func (r *sqliteWatchRepository) removeCompanyNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update company_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCompanyNameWatch(watchID, err, stmt)
}

// Customer methods - SQLite

func (r *sqliteWatchRepository) addCustomerWatch(customerID string, params watchRequest) (string, error) {
	r.logger.Log("Repository is of type %s", reflect.TypeOf(r).String())
	if customerID == "" {
		return "", errNoCustomerID
	}
	id := base.ID()

	query := `insert into customer_watches (id, customer_id, webhook, auth_token, created_at) values (?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	return insertCustomerWatch(customerID, params, err, stmt, id)
}

func (r *sqliteWatchRepository) removeCustomerWatch(customerID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update customer_watches set deleted_at = ? where customer_id = ? and id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCustomerWatch(customerID, watchID, err, stmt)
}

func (r *sqliteWatchRepository) addCustomerNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert into customer_name_watches (id, name, webhook, auth_token, created_at) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	return insertCustomerNameWatch(name, webhook, authToken, err, stmt)
}

func (r *sqliteWatchRepository) removeCustomerNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update customer_name_watches set deleted_at = ? where id = ? and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCustomerNameWatch(watchID, err, stmt)
}

////////////////////////////
// Postgres implementation
////////////////////////////
type postgresWatchRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *postgresWatchRepository) close() error {
	return r.db.Close()
}

func (r *postgresWatchRepository) getWatchesCursor(logger log.Logger, batchSize int) *watchCursor {
	return &watchCursor{
		batchSize: batchSize,
		db:        r.db,
		logger:    logger,
	}
}

// Company methods - Postgres
func (r *postgresWatchRepository) addCompanyWatch(companyID string, params watchRequest) (string, error) {
	if companyID == "" {
		return "", errNoCompanyID
	}
	id := base.ID()

	query := `insert into company_watches (id, company_id, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5)`
	stmt, err := r.db.Prepare(query)
	return insertCompanyWatch(companyID, params, err, stmt, id)
}

func (r *postgresWatchRepository) removeCompanyWatch(companyID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update company_watches set deleted_at = $1 where company_id = $2 and id = $3 and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCompanyWatch(companyID, watchID, err, stmt)
}

func (r *postgresWatchRepository) addCompanyNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert into company_name_watches (id, name, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5);`
	stmt, err := r.db.Prepare(query)
	return insertCompanyNameWatch(name, webhook, authToken, err, stmt)
}

func (r *postgresWatchRepository) removeCompanyNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update company_name_watches set deleted_at = $1 where id = $2 and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCompanyNameWatch(watchID, err, stmt)
}

// Customer methods - Postgres

func (r *postgresWatchRepository) addCustomerWatch(customerID string, params watchRequest) (string, error) {
	if customerID == "" {
		return "", errNoCustomerID
	}
	id := base.ID()

	query := `insert into customer_watches (id, customer_id, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5)`
	stmt, err := r.db.Prepare(query)
	return insertCustomerWatch(customerID, params, err, stmt, id)
}

func (r *postgresWatchRepository) removeCustomerWatch(customerID string, watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update customer_watches set deleted_at = $1 where customer_id = $2 and id = $3 and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCustomerWatch(customerID, watchID, err, stmt)
}

func (r *postgresWatchRepository) addCustomerNameWatch(name string, webhook string, authToken string) (string, error) {
	query := `insert into customer_name_watches (id, name, webhook, auth_token, created_at) values ($1, $2, $3, $4, $5);`
	stmt, err := r.db.Prepare(query)
	return insertCustomerNameWatch(name, webhook, authToken, err, stmt)
}

func (r *postgresWatchRepository) removeCustomerNameWatch(watchID string) error {
	if watchID == "" {
		return errNoWatchID
	}

	query := `update customer_name_watches set deleted_at = $1 where id = $2 and deleted_at is null`
	stmt, err := r.db.Prepare(query)
	return updateCustomerNameWatch(watchID, err, stmt)
}

func getWatchRepo(dbType string, db *sql.DB, logger log.Logger) watchRepository {
	if dbType == "postgres" {
		return &postgresWatchRepository{db, logger}
	} else if dbType == "mysql" {
		return nil
	}
	return &sqliteWatchRepository{db, logger}
}

// Common access code across DB

// Company Methods
func insertCompanyWatch(companyID string, params watchRequest, err error, stmt *sql.Stmt, id string) (string, error) {
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

func updateCompanyWatch(companyID string, watchID string, err error, stmt *sql.Stmt) error {
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), companyID, watchID)
	return err
}

func insertCompanyNameWatch(name string, webhook string, authToken string, err error, stmt *sql.Stmt) (string, error) {
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

func updateCompanyNameWatch(watchID string, err error, stmt *sql.Stmt) error {
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), watchID)
	return err
}

//Customer Methods
func insertCustomerWatch(customerID string, params watchRequest, err error, stmt *sql.Stmt, id string) (string, error) {
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

func updateCustomerWatch(customerID string, watchID string, err error, stmt *sql.Stmt) error {
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), customerID, watchID)
	return err
}

func insertCustomerNameWatch(name string, webhook string, authToken string, err error, stmt *sql.Stmt) (string, error) {
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

func updateCustomerNameWatch(watchID string, err error, stmt *sql.Stmt) error {
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
	var query = ""
	switch dbType {
	case "postgres":
		query = `select id, company_id, webhook, auth_token, created_at from company_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	default:
		query = `select id, company_id, webhook, auth_token, created_at from company_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
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
	var query = ""
	switch dbType {
	case "postgres":
		query = `select id, name, webhook, auth_token, created_at from company_name_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	default:
		query = `select id, name, webhook, auth_token, created_at from company_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
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
	var query = ""
	switch dbType {
	case "postgres":
		query = `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	default:
		query = `select id, customer_id, webhook, auth_token, created_at from customer_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
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
	var query = ""
	switch dbType {
	case "postgres":
		query = `select id, name, webhook, auth_token, created_at from customer_name_watches where created_at > $1 and deleted_at is null order by created_at asc limit $2`
	default:
		query = `select id, name, webhook, auth_token, created_at from customer_name_watches where created_at > ? and deleted_at is null order by created_at asc limit ?`
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
