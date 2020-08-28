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

// SQLite implementation of company repository
type sqliteCompanyRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *sqliteCompanyRepository) close() error {
	return r.db.Close()
}

func (r *sqliteCompanyRepository) getCompanyStatus(companyID string) (*CompanyStatus, error) {
	if companyID == "" {
		return nil, errors.New("getCompanyStatus: no Company.ID")
	}
	query := `select user_id, note, status, created_at from company_status where company_id = ? and deleted_at is null order by created_at desc limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return queryCompanyStatus(companyID, stmt, err)
}

func (r *sqliteCompanyRepository) upsertCompanyStatus(companyID string, status *CompanyStatus) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("upsertCompanyStatus: begin: %v", err)
	}

	query := `insert into company_status (company_id, user_id, note, status, created_at) values (?, ?, ?, ?, ?);`
	return insertCompanyStatus(companyID, status, err, tx, query)
}

// Postgres implementation of company repository
type postgresCompanyRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *postgresCompanyRepository) close() error {
	return r.db.Close()
}

func (r *postgresCompanyRepository) getCompanyStatus(companyID string) (*CompanyStatus, error) {
	if companyID == "" {
		return nil, errors.New("getCompanyStatus: no Company.ID")
	}
	query := `select user_id, note, status, created_at from company_status where company_id = $1 and deleted_at is null order by created_at desc limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	return queryCompanyStatus(companyID, stmt, err)
}

func (r *postgresCompanyRepository) upsertCompanyStatus(companyID string, status *CompanyStatus) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("upsertCompanyStatus: begin: %v", err)
	}

	query := `insert into company_status (company_id, user_id, note, status, created_at) values ($1, $2, $3, $4, $5);`
	return insertCompanyStatus(companyID, status, err, tx, query)
}

func getCompanyRepo(dbType string, db *sql.DB, logger log.Logger) companyRepository {
	if dbType == "postgres" {
		return &postgresCompanyRepository{db, logger}
	} else if dbType == "mysql" {
		return nil
	}
	return &sqliteCompanyRepository{db, logger}
}

// Common access code across DB
func queryCompanyStatus(companyID string, stmt *sql.Stmt, err error) (*CompanyStatus, error) {
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

func insertCompanyStatus(companyID string, status *CompanyStatus, err error, tx *sql.Tx, query string) error {
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
