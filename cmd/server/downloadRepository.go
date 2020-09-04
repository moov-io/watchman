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

const (
	genericInsertDownloadStats = `insert into download_stats (downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities) values (?, ?, ?, ?, ?, ?, ?);`
	genericSelectDownloadStats = `select downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities from download_stats order by downloaded_at desc limit ?;`

	postgresInsertDownloadStats = `insert into download_stats (downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities) values ($1, $2, $3, $4, $5, $6, $7);`
	postgresSelectDownloadStats = `select downloaded_at, sdns, alt_names, addresses, sectoral_sanctions, denied_persons, bis_entities from download_stats order by downloaded_at desc limit $1;`
)

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

	query := genericInsertDownloadStats
	switch dbType {
	case `postgres`:
		query = postgresInsertDownloadStats
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stats.RefreshedAt, stats.SDNs, stats.Alts, stats.Addresses, stats.SectoralSanctions, stats.DeniedPersons, stats.BISEntities)
	return err
}

func (r *genericSQLDownloadRepository) latestDownloads(limit int) ([]Download, error) {
	query := genericSelectDownloadStats
	switch dbType {
	case `postgres`:
		query = postgresSelectDownloadStats
	}
	stmt, err := r.db.Prepare(query)
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
