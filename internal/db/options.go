package db

import (
	"github.com/jmoiron/sqlx"
)

type Option func(db *db) error

func MySQLRebind() Option {
	return func(db *db) error {
		original := db.onQuery
		db.onQuery = func(query string) string {
			query = original(query)
			return sqlx.Rebind(sqlx.QUESTION, query)
		}
		return nil
	}
}

func PostgresRebind() Option {
	return func(db *db) error {
		original := db.onQuery
		db.onQuery = func(query string) string {
			query = original(query)
			return sqlx.Rebind(sqlx.DOLLAR, query)
		}
		return nil
	}
}
