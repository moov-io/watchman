package main

import (
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

type downloadRepository interface {
	latestDownloads(limit int) ([]Download, error)
	recordStats(stats *downloadStats) error
	close() error
}

////////////////////////////////////////////////////////
// generic implementation for most
// databases (SQLite, MySQL)
////////////////////////////////////////////////////////
type genericSQLDownloadRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *genericSQLDownloadRepository) close() error {
	return r.db.Close()
}

func (r *genericSQLDownloadRepository) recordStats(stats *downloadStats) error {
	if stats == nil {
		return errors.New("recordStats: nil downloadStats")
	}

	query := `insert into download_stats (downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities) values (?, ?, ?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	return insertDownloadStat(stats, err, stmt)
}

func (r *genericSQLDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := `select downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities from download_stats order by downloaded_at desc limit ?;`
	stmt, err := r.db.Prepare(query)
	return selectLatestDownload(limit, err, stmt)
}

////////////////////////////////////////////////////////
// postgres implementation
////////////////////////////////////////////////////////
type postgresDownloadRepository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *postgresDownloadRepository) close() error {
	return r.db.Close()
}

func (r *postgresDownloadRepository) recordStats(stats *downloadStats) error {
	if stats == nil {
		return errors.New("recordStats: nil downloadStats")
	}

	query := `insert into download_stats (downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities) values ($1, $2, $3, $4, $5, $6, $7);`
	stmt, err := r.db.Prepare(query)
	return insertDownloadStat(stats, err, stmt)
}

func (r *postgresDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := `select downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities from download_stats order by downloaded_at desc limit ?;`
	stmt, err := r.db.Prepare(query)
	return selectLatestDownload(limit, err, stmt)
}

// Common Code across DB implemation

// This function will return a downloadRepository for a specific database that requires specific handling of
// queries such as Postgres and Oracle. Other databases such as SQLite and MySQL will get a generic repository.
func getDownloadRepo(dbType string, db *sql.DB, logger log.Logger) downloadRepository {
	switch dbType {
	case "postgres":
		return &postgresDownloadRepository{db: db, logger: logger}
	default:
		return &genericSQLDownloadRepository{db: db, logger: logger}
	}
}

func insertDownloadStat(stats *downloadStats, err error, stmt *sql.Stmt) error {
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stats.RefreshedAt, stats.SDNs, stats.Alts, stats.Addresses, stats.SectoralSanctions, stats.DeniedPersons, stats.BISEntities)
	return err
}

func selectLatestDownload(limit int, err error, stmt *sql.Stmt) ([]Download, error) {
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var downloads []Download
	for rows.Next() {
		var dl Download
		if err := rows.Scan(&dl.Timestamp, &dl.SDNs, &dl.Alts, &dl.Addresses, &dl.SectoralSanctions, &dl.DeniedPersons, &dl.BISEntities); err == nil {
			downloads = append(downloads, dl)
		}
	}
	return downloads, rows.Err()
}
