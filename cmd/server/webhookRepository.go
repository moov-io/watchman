package main

import (
	"database/sql"
	"time"
)

type webhookRepository interface {
	recordWebhook(watchID string, attemptedAt time.Time, status int) error
	close() error
}

const (
	genericInsertWebhookStats  = `insert into webhook_stats (watch_id, attempted_at, status) values (?, ?, ?);`
	postgresInsertWebhookStats = `insert into webhook_stats (watch_id, attempted_at, status) values ($1, $2, $3);`
)

type genericSQLWebhookRepository struct {
	db *sql.DB
}

func (r *genericSQLWebhookRepository) close() error {
	return r.db.Close()
}

func (r *genericSQLWebhookRepository) recordWebhook(watchID string, attemptedAt time.Time, status int) error {
	var query = genericInsertWebhookStats
	switch dbType {
	case `postgres`:
		query = postgresInsertWebhookStats
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(watchID, attemptedAt, status)
	return err
}
