package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/moov-io/watchman/internal/database"
	"strings"
)

// companyRepository holds the current status (i.e. unsafe or exception) for a given company and
// is expected to save metadata about each time the status is changed.
type companyRepository interface {
	getCompanyStatus(companyID string) (*CompanyStatus, error)
	upsertCompanyStatus(companyID string, status *CompanyStatus) error
	close() error
}

const (
	genericSelectCompanyStatus = `select user_id, note, status, created_at from company_status where company_id = ? and deleted_at is null order by created_at desc limit 1;`
	genericInsertCompanyStatus = `insert into company_status (company_id, user_id, note, status, created_at) values (?, ?, ?, ?, ?);`

	postgresSelectCompanyStatus = `select user_id, note, status, created_at from company_status where company_id = $1 and deleted_at is null order by created_at desc limit 1;`
	postgresInsertCompanyStatus = `insert into company_status (company_id, user_id, note, status, created_at) values ($1, $2, $3, $4, $5);`
)

type genericSQLCompanyRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *genericSQLCompanyRepository) close() error {
	return r.db.Close()
}

func (r *genericSQLCompanyRepository) getCompanyStatus(companyID string) (*CompanyStatus, error) {
	if companyID == "" {
		return nil, errors.New("getCompanyStatus: no Company.ID")
	}
	query := genericSelectCompanyStatus
	switch dbType {
	case `postgres`:
		query = postgresSelectCompanyStatus
	}
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(companyID)

	var status CompanyStatus
	err = row.Scan(&status.UserID, &status.Note, &status.Status, &status.CreatedAt)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, fmt.Errorf("getCompanyStatus: %v", err)
	}
	if status.UserID == "" {
		return nil, nil // not found
	}
	return &status, nil
}

func (r *genericSQLCompanyRepository) upsertCompanyStatus(companyID string, status *CompanyStatus) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("upsertCompanyStatus: begin: %v", err)
	}

	query := genericInsertCompanyStatus
	switch dbType {
	case `postgres`:
		query = postgresInsertCompanyStatus
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("upsertCompanyStatus: prepare error=%v rollback=%v", err, tx.Rollback())
	}
	_, err = stmt.Exec(companyID, status.UserID, status.Note, status.Status, status.CreatedAt)
	stmt.Close()
	if err == nil {
		return tx.Commit()
	}
	if database.UniqueViolation(err) {
		query = `update company_status set note = ?, status = ? where company_id = ? and user_id = ?;`
		stmt, err = tx.Prepare(query)
		if err != nil {
			return fmt.Errorf("upsertCompanyStatus: inner prepare error=%v rollback=%v", err, tx.Rollback())
		}
		_, err := stmt.Exec(status.Note, status.Status, companyID, status.UserID)
		stmt.Close()
		if err != nil {
			return fmt.Errorf("upsertCompanyStatus: unique error=%v rollback=%v", err, tx.Rollback())
		}
	}
	return tx.Commit()
}
