// Copyright 2019 DigitalMint [Carlos Saavedra]
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	kitprom "github.com/go-kit/kit/metrics/prometheus"

	_ "github.com/lib/pq"
	"github.com/lopezator/migrator"
	stdprom "github.com/prometheus/client_golang/prometheus"
)

var (
	postgreSQLConnections = kitprom.NewGaugeFrom(stdprom.GaugeOpts{
		Name: "postgresql_connections",
		Help: "How many postgresql connections and what status they're in.",
	}, []string{"state"})

	postgresqlMigrator = migrator.New(
		execsql(
			"create_customer_name_watches",
			`create table if not exists customer_name_watches(id SERIAL PRIMARY KEY, name varchar, webhook varchar, auth_token varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_customer_status",
			`create table if not exists customer_status(customer_id integer primary key, user_id integer, note varchar, status varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_customer_watches",
			`create table if not exists customer_watches(id integer primary key, customer_id integer, webhook varchar, auth_token varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_company_name_watches",
			`create table if not exists company_name_watches(id integer primary key, name varchar, webhook varchar, auth_token varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_company_status",
			`create table if not exists company_status(company_id integer primary key, user_id integer, note varchar, status varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_company_watches",
			`create table if not exists company_watches(id integer primary key, company_id integer, webhook varchar, auth_token varchar, created_at timestamp, deleted_at timestamp);`,
		),
		execsql(
			"create_ofac_download_stats",
			`create table if not exists ofac_download_stats(downloaded_at timestamp, sdns integer, alt_names integer, addresses integer);`,
		),
		execsql(
			"create_webhook_stats",
			`create table if not exists webhook_stats(watch_id varchar, attempted_at timestamp, status varchar);`,
		),
		execsql("add__denied_persons__to__ofac_download_stats", "alter table ofac_download_stats add column denied_persons integer not null default 0;"),
	)
)

type postgresql struct {
	dns string

	connections *kitprom.Gauge
	logger      log.Logger

	err error
}

// Connect to PosgreSQL database
func (s *postgresql) Connect() (*sql.DB, error) {
	if s.err != nil {
		return nil, fmt.Errorf("postgresql had error %v", s.err)
	}
	// Connect to our DB and perform a quick sanity check
	db, err := sql.Open("postgres", s.dns)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return db, err
	}

	// Migrate our database
	if err := postgresqlMigrator.Migrate(db); err != nil {
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

// postgreSQLConnection prepare connection variables
func postgreSQLConnection(logger log.Logger, user, password, dbname, host string, isSsl bool, port int) *postgresql {
	// check for ssl
	var sslmode string
	if isSsl == false {
		sslmode = "disable"
	} else {
		sslmode = "require"
	}

	// create connection string
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	return &postgresql{
		dns:         dns,
		logger:      logger,
		connections: postgreSQLConnections,
	}
}

// PostgreSQLUniqueViolation returns true when the provided error matches the PostgreSQL error
// for duplicate entries (violating a unique table constraint).
func PostgreSQLUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "ERROR: duplicate key")
}
