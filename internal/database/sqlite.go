// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	kitprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/lopezator/migrator"
	"github.com/mattn/go-sqlite3"
	stdprom "github.com/prometheus/client_golang/prometheus"
)

var (
	sqliteConnections = kitprom.NewGaugeFrom(stdprom.GaugeOpts{
		Name: "sqlite_connections",
		Help: "How many sqlite connections and what status they're in.",
	}, []string{"state"})

	sqliteVersionLogOnce sync.Once

	sqliteMigrator = migrator.New(
		execsql(
			"create_customer_name_watches",
			`create table if not exists customer_name_watches(id primary key, name, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_customer_status",
			`create table if not exists customer_status(customer_id primary key, user_id, note, status, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_customer_watches",
			`create table if not exists customer_watches(id primary key, customer_id, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_company_name_watches",
			`create table if not exists company_name_watches(id primary key, name, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_company_status",
			`create table if not exists company_status(company_id primary key, user_id, note, status, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_company_watches",
			`create table if not exists company_watches(id primary key, company_id, webhook, auth_token, created_at datetime, deleted_at datetime);`,
		),
		execsql(
			"create_ofac_download_stats",
			`create table if not exists ofac_download_stats(downloaded_at datetime, sdns, alt_names, addresses);`,
		),
		execsql(
			"create_webhook_stats",
			`create table if not exists webhook_stats(watch_id string, attempted_at datetime, status);`,
		),
		addSQLiteColumn("ofac_download_stats", "denied_persons"),
	)
)

func addSQLiteColumn(tableName, columnDesc string) *migrator.Migration {
	colName := strings.Fields(columnDesc)[0] // take column name ('deleted_at' or 'deleted_at timestamp')

	return &migrator.Migration{
		Name: fmt.Sprintf("add__%s__to__%s", colName, tableName),
		Func: func(tx *sql.Tx) error {
			stmt, err := tx.Prepare(fmt.Sprintf(`select 1 from pragma_table_info('%s') where name = ? limit 1;`, tableName))
			if err != nil {
				return fmt.Errorf("addSQLiteColumn: column=%s failed: %v", colName, err)
			}
			var n int
			if err := stmt.QueryRow(colName).Scan(&n); err != nil && err != sql.ErrNoRows {
				return fmt.Errorf("addSQLiteColumn: query column=%s failed: %#v", colName, err)
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

type sqlite struct {
	path string

	connections *kitprom.Gauge
	logger      log.Logger

	err error
}

func (s *sqlite) Connect() (*sql.DB, error) {
	if s.err != nil {
		return nil, fmt.Errorf("sqlite had error %v", s.err)
	}

	sqliteVersionLogOnce.Do(func() {
		if v, _, _ := sqlite3.Version(); v != "" {
			s.logger.Log("main", fmt.Sprintf("sqlite version %s", v))
		}
	})

	// Connect to our DB and perform a quick sanity check
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return db, err
	}

	// Migrate our database
	if err := sqliteMigrator.Migrate(db); err != nil {
		return db, err
	}

	// Spin up metrics only after everything works
	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			stats := db.Stats()
			s.connections.With("state", "idle").Set(float64(stats.Idle))
			s.connections.With("state", "inuse").Set(float64(stats.InUse))
			s.connections.With("state", "open").Set(float64(stats.OpenConnections))
		}
	}()

	return db, err
}

func sqliteConnection(logger log.Logger, path string) *sqlite {
	return &sqlite{
		path:        path,
		logger:      logger,
		connections: sqliteConnections,
	}
}

// TestSQLiteDB is a wrapper around sql.DB for SQLite connections designed for tests to provide
// a clean database for each testcase.  Callers should cleanup with Close() when finished.
type TestSQLiteDB struct {
	DB *sql.DB

	dir string // temp dir created for sqlite files
}

func (r *TestSQLiteDB) Close() error {
	if err := r.DB.Close(); err != nil {
		return err
	}
	return os.RemoveAll(r.dir)
}

// CreateTestSqliteDB returns a TestSQLiteDB which can be used in tests
// as a clean sqlite database. All migrations are ran on the db before.
//
// Callers should call close on the returned *TestSQLiteDB.
func CreateTestSqliteDB(t *testing.T) *TestSQLiteDB {
	dir, err := ioutil.TempDir("", "ofac-sqlite")
	if err != nil {
		t.Fatalf("sqlite test: %v", err)
	}

	db, err := sqliteConnection(log.NewNopLogger(), filepath.Join(dir, "ofac.db")).Connect()
	if err != nil {
		t.Fatalf("sqlite test: %v", err)
	}
	return &TestSQLiteDB{db, dir}
}

// SqliteUniqueViolation returns true when the provided error matches the SQLite error
// for duplicate entries (violating a unique table constraint).
func SqliteUniqueViolation(err error) bool {
	match := strings.Contains(err.Error(), "UNIQUE constraint failed")
	if e, ok := err.(sqlite3.Error); ok {
		return match || e.Code == sqlite3.ErrConstraint
	}
	return match
}
