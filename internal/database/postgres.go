// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	kitprom "github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lopezator/migrator"
	"github.com/moov-io/base/docker"
	"github.com/ory/dockertest/v3"
	stdprom "github.com/prometheus/client_golang/prometheus"
)

var (
	postgresConnections = kitprom.NewGaugeFrom(stdprom.GaugeOpts{
		Name: "postgres_connections",
		Help: "How many postgres connections and what status they're in.",
	}, []string{"state"})

	// postgresErrDuplicateKey is the error code for duplicate entries
	// https://www.postgresql.org/docs/11/errcodes-appendix.html
	postgresErrDuplicateKey string = "23505"

	postgresMigrations = migrator.Migrations(
		execsql(
			"create_customer_name_watches",
			`create table if not exists customer_name_watches(id VARCHAR(40) primary key, name varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_customer_status",
			`create table if not exists customer_status(customer_id varchar(40) primary key, user_id varchar(40), note varchar(1024), status varchar(10), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_customer_watches",
			`create table if not exists customer_watches(id varchar(40) primary key, customer_id varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_company_name_watches",
			`create table if not exists company_name_watches(id varchar(40) primary key, name varchar(256), webhook varchar(512), auth_token varchar(128), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_company_status",
			`create table if not exists company_status(company_id varchar(40) primary key, user_id varchar(40), note varchar(1024), status varchar(10), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_company_watches",
			`create table if not exists company_watches(id varchar(40) primary key, company_id varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(6), deleted_at timestamp(6));`,
		),
		execsql(
			"create_ofac_download_stats",
			`create table if not exists ofac_download_stats(downloaded_at timestamp(6), sdns int4, alt_names int4, addresses int4);`,
		),
		execsql(
			"create_webhook_stats",
			`create table if not exists webhook_stats(watch_id varchar(40), attempted_at timestamp(6), status int4);`,
		),
		execsql(
			"add__denied_persons__to__ofac_download_stats",
			`alter table ofac_download_stats add column denied_persons int4 not null default 0;`,
		),
		execsql(
			"rename_ofac_download_stats",
			`alter table ofac_download_stats rename to download_stats`,
		),
		execsql(
			"add_sectoral_sanctions_to_download_stats",
			`alter table download_stats add column sectoral_sanctions int4 not null default 0;`,
		),
		execsql(
			"add__bis_entities__to_download_stats",
			`alter table download_stats add column bis_entities int4 not null default 0;`,
		),
	)
)

type postgres struct {
	dsn    string
	logger log.Logger

	connections *kitprom.Gauge
}

func (my *postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", my.dsn)
	if err != nil {
		return nil, err
	}

	// Check out DB is up and working
	/*if err := db.Ping(); err != nil {
		return nil, err
	}*/

	// Migrate our database
	if m, err := migrator.New(postgresMigrations); err != nil {
		return nil, err
	} else {
		if err := m.Migrate(db); err != nil {
			return nil, err
		}
	}

	// Setup metrics after the database is setup
	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			stats := db.Stats()
			my.connections.With("state", "idle").Set(float64(stats.Idle))
			my.connections.With("state", "inuse").Set(float64(stats.InUse))
			my.connections.With("state", "open").Set(float64(stats.OpenConnections))
		}
	}()

	return db, nil
}

func postgresConnection(logger log.Logger, user, pass string, address string, database string) *postgres {
	timeout := "10"
	if v := os.Getenv("POSTGRES_TIMEOUT"); v != "" {
		timeout = v
	}
	sslenabled := "disable"
	if sse := os.Getenv("POSTGRES_SSL"); sse != "" {
		sslenabled = sse
	}
	params := fmt.Sprintf("sslmode=%s&connect_timeout=%s", sslenabled, timeout)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", user, pass, address, database, params)
	return &postgres{
		dsn:         dsn,
		logger:      logger,
		connections: postgresConnections,
	}
}

// TestPostgresDB is a wrapper around sql.DB for Postgres connections designed for tests to provide
// a clean database for each testcase.  Callers should cleanup with Close() when finished.
type TestPostgresDB struct {
	DB        *sql.DB
	container *dockertest.Resource
}

func (r *TestPostgresDB) Close() error {
	r.container.Close()
	return r.DB.Close()
}

// CreateTestPostgresDB returns a TestPostgresDB which can be used in tests
// as a clean Postgres database. All migrations are ran on the db before.
//
// Callers should call close on the returned *TestPostgresDB.
func CreateTestPostgresDB(t *testing.T) *TestPostgresDB {
	if testing.Short() {
		t.Skip("-short flag enabled")
	}
	if !docker.Enabled() {
		t.Skip("Docker not enabled")
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal(err)
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.4",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_DB=watchman",
			"PGSSLMODE=disable",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = pool.Retry(func() error {
		db, err := sql.Open("pgx", fmt.Sprintf("postgres://postgres:secret@localhost:%s/watchman?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		defer db.Close()
		return db.Ping()
	})
	if err != nil {
		resource.Close()
		t.Fatal(err)
	}

	logger := log.NewNopLogger()
	address := fmt.Sprintf("localhost:%s", resource.GetPort("5432/tcp"))

	db, err := postgresConnection(logger, "postgres", "secret", address, "watchman").Connect()
	if err != nil {
		t.Fatal(err)
	}
	return &TestPostgresDB{db, resource}
}

// PostgresUniqueViolation returns true when the provided error matches the Postgres code
// for duplicate entries (violating a unique table constraint).
func PostgresUniqueViolation(err error) bool {
	match := strings.Contains(err.Error(), "[Err] ERROR:  duplicate key value violates unique constraint")
	if e, ok := err.(*pgx.SerializationError); ok {
		return match || e.Error() == postgresErrDuplicateKey
	}
	return match
}
