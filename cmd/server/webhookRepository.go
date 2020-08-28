package main

import (
	"database/sql"
	"time"
)

type webhookRepository interface {
	recordWebhook(watchID string, attemptedAt time.Time, status int) error
	close() error
}

type sqliteWebhookRepository struct {
	db *sql.DB
}

func (r *sqliteWebhookRepository) close() error {
	return r.db.Close()
}

func (r *sqliteWebhookRepository) recordWebhook(watchID string, attemptedAt time.Time, status int) error {
	query := `insert into webhook_stats (watch_id, attempted_at, status) values (?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(watchID, attemptedAt, status)
	return err
}

// postgres implementation
type postgresWebhookRepository struct {
	db *sql.DB
}

func (r *postgresWebhookRepository) close() error {
	return r.db.Close()
}

func (r *postgresWebhookRepository) recordWebhook(watchID string, attemptedAt time.Time, status int) error {
	query := `insert into webhook_stats (watch_id, attempted_at, status) values ($1, $2, $3);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(watchID, attemptedAt, status)
	return err
}

func getWebhookRepo(dbType string, db *sql.DB) webhookRepository {
	if dbType == "postgres" {
		return &postgresWebhookRepository{db}
	} else if dbType == "mysql" {
		return nil
	}
	return &sqliteWebhookRepository{db}
}
