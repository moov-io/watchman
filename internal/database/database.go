// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/moov-io/base/log"

	"github.com/lopezator/migrator"
)

func New(logger log.Logger, _type string) (*sql.DB, error) {
	logger.Logf("looking for %s database provider", _type)

	switch strings.ToLower(_type) {
	case "sqlite", "":
		return sqliteConnection(logger.With(log.Fields{
			"database": log.String("sqlite"),
		}), getSqlitePath()).Connect()

	case "mysql":
		return NewMySQLConnection(
			logger.With(log.Fields{
				"database": log.String("mysql"),
			}),
			MySQLConfig{
				Address:  os.Getenv("MYSQL_ADDRESS"),
				Username: os.Getenv("MYSQL_USER"),
				Password: os.Getenv("MYSQL_PASSWORD"),
				Database: os.Getenv("MYSQL_DATABASE"),
			},
		)
	}
	return nil, fmt.Errorf("unknown database type %q", _type)
}

func execsql(name, raw string) *migrator.MigrationNoTx {
	return &migrator.MigrationNoTx{
		Name: name,
		Func: func(db *sql.DB) error {
			_, err := db.Exec(raw)
			return err
		},
	}
}

// UniqueViolation returns true when the provided error matches a database error
// for duplicate entries (violating a unique table constraint).
func UniqueViolation(err error) bool {
	return MySQLUniqueViolation(err) || SqliteUniqueViolation(err)
}
