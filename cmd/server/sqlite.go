// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/lopezator/migrator"
	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteMigrator = migrator.New(
		createTable(
			"create_customer_name_watches",
			`create table if not exists customer_name_watches(id primary key, name, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_customer_status",
			`create table if not exists customer_status(customer_id, user_id, note, status, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_customer_watches",
			`create table if not exists customer_watches(id primary key, customer_id, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_company_name_watches",
			`create table if not exists company_name_watches(id primary key, name, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_company_status",
			`create table if not exists company_status(company_id, user_id, note, status, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_company_watches",
			`create table if not exists company_watches(id primary key, company_id, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		createTable(
			"create_ofac_download_stats",
			`create table if not exists ofac_download_stats(downloaded_at datetime, sdns, alt_names, addresses);`,
		),
		createTable(
			"create_webhook_stats",
			`create table if not exists webhook_stats(watch_id string, attempted_at datetime, status);`,
		),
		addColumn("ofac_download_stats", "denied_persons"),
	)
)

func createTable(name, raw string) *migrator.MigrationNoTx {
	return &migrator.MigrationNoTx{
		Name: name,
		Func: func(db *sql.DB) error {
			_, err := db.Exec(raw)
			return err
		},
	}
}

func addColumn(tableName, columnDesc string) *migrator.Migration {
	colName := strings.Fields(columnDesc)[0] // take column name ('deleted_at' or 'deleted_at timestamp')

	return &migrator.Migration{
		Name: fmt.Sprintf("add__%s__to__%s", colName, tableName),
		Func: func(tx *sql.Tx) error {
			stmt, err := tx.Prepare(fmt.Sprintf(`select 1 from pragma_table_info('%s') where name = ? limit 1;`, tableName))
			if err != nil {
				return fmt.Errorf("addColumn: column=%s failed: %v", colName, err)
			}
			var n int
			if err := stmt.QueryRow(colName).Scan(&n); err != nil && err != sql.ErrNoRows {
				return fmt.Errorf("addColumn: query column=%s failed: %#v", colName, err)
			}
			if n == 0 {
				_, err := tx.Exec(fmt.Sprintf(`alter table %s add column %s`, tableName, columnDesc))
				return err
			}
			return nil
		},
	}
}

func getSqlitePath() string {
	path := os.Getenv("SQLITE_DB_PATH")
	if path == "" || strings.Contains(path, "..") {
		// set default if empty or trying to escape
		// don't filepath.ABS to avoid full-fs reads
		path = "ofac.db"
	}
	return path
}

func createSqliteConnection(logger log.Logger, path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		err = fmt.Errorf("problem opening sqlite3 file %s: %v", path, err)
		if logger != nil {
			logger.Log("sqlite", err)
		}
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("problem with Ping against *sql.DB %s: %v", path, err)
	}
	return db, nil
}

// migrate runs our database migrations (defined at the top of this file)
// over a sqlite database it creates first.
// To configure where on disk the sqlite db is set SQLITE_DB_PATH.
//
// You use db like any other database/sql driver.
//
// https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.3.html
func migrate(logger log.Logger, db *sql.DB) error {
	if logger != nil {
		logger.Log("sqlite", "starting database migrations")
	}
	if err := sqliteMigrator.Migrate(db); err != nil {
		return err
	}
	if logger != nil {
		logger.Log("sqlite", "finished migrations")
	}
	return nil
}
