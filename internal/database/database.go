// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/lopezator/migrator"
)

func New(logger log.Logger, _type string) (*sql.DB, error) {
	logger.Log("database", fmt.Sprintf("looking for %s database provider", _type))
	switch strings.ToLower(_type) {
	case "postgres":
		// check for ssl
		hasSSL := (os.Getenv("DB_SSL") != "")

		// check for port
		port, _ := getenvInt(os.Getenv("DB_PORT"))
		if port == 0 {
			// default Postgres port
			port = 5432
		}

		return postgreSQLConnection(logger, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_HOST"), hasSSL, port).Connect()
	case "sqlite", "":
		return sqliteConnection(logger, getSqlitePath()).Connect()
	case "mysql":
		return mysqlConnection(logger, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDRESS"), os.Getenv("MYSQL_DATABASE")).Connect()
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
	return MySQLUniqueViolation(err) || SqliteUniqueViolation(err) || PostgreSQLUniqueViolation(err)
}

func getenvInt(key string) (int, error) {
	if key == "" {
		return 0, errors.New("empty environment variable")
	}
	v, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return v, nil
}
