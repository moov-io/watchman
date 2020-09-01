package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/moov-io/watchman/internal/database"
	"strings"
)

// customerRepository holds the current status (i.e. unsafe or exception) for a given customer
// (individual) and is expected to save metadata about each time the status is changed.
type customerRepository interface {
	getCustomerStatus(customerID string) (*CustomerStatus, error)
	upsertCustomerStatus(customerID string, status *CustomerStatus) error
	close() error
}

const (
	genericSelectCustomerStatus = `select user_id, note, status, created_at from customer_status where customer_id = ? and deleted_at is null order by created_at desc limit 1;`
	genericInsertCustomerStatus = `insert into customer_status (customer_id, user_id, note, status, created_at) values (?, ?, ?, ?, ?);`

	postgresSelectCustomerStatus = `select user_id, note, status, created_at from customer_status where customer_id = $1 and deleted_at is null order by created_at desc limit 1;`
	postgresInsertCustomerStatus = `insert into customer_status (customer_id, user_id, note, status, created_at) values ($1, $2, $3, $4, $5);`
)

type genericSQLCustomerRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *genericSQLCustomerRepository) close() error {
	return r.db.Close()
}

func (r *genericSQLCustomerRepository) getCustomerStatus(customerID string) (*CustomerStatus, error) {
	if customerID == "" {
		return nil, errors.New("getCustomerStatus: no Customer.ID")
	}
	query := genericSelectCustomerStatus
	switch dbType {
	case `postgres`:
		query = postgresSelectCustomerStatus
	}
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

func (r *genericSQLCustomerRepository) upsertCustomerStatus(customerID string, status *CustomerStatus) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("upsertCustomerStatus: begin: %v", err)
	}

	query := genericInsertCustomerStatus
	switch dbType {
	case `postgres`:
		query = postgresInsertCustomerStatus
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("upsertCustomerStatus: prepare error=%v rollback=%v", err, tx.Rollback())
	}
	_, err = stmt.Exec(customerID, status.UserID, status.Note, status.Status, status.CreatedAt)
	stmt.Close()
	if err == nil {
		return tx.Commit()
	}
	if database.UniqueViolation(err) {
		query = `update customer_status set note = ?, status = ? where customer_id = ? and user_id = ?`
		stmt, err = tx.Prepare(query)
		if err != nil {
			return fmt.Errorf("upsertCustomerStatus: inner prepare error=%v rollback=%v", err, tx.Rollback())
		}
		_, err := stmt.Exec(status.Note, status.Status, customerID, status.UserID)
		stmt.Close()
		if err != nil {
			return fmt.Errorf("upsertCustomerStatus: unique error=%v rollback=%v", err, tx.Rollback())
		}
	}
	return tx.Commit()
}
