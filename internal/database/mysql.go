// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base"
	"github.com/moov-io/base/log"

	kitprom "github.com/go-kit/kit/metrics/prometheus"
	gomysql "github.com/go-sql-driver/mysql"
	"github.com/lopezator/migrator"
	stdprom "github.com/prometheus/client_golang/prometheus"
)

var (
	mysqlConnections = kitprom.NewGaugeFrom(stdprom.GaugeOpts{
		Name: "mysql_connections",
		Help: "How many MySQL connections and what status they're in.",
	}, []string{"state"})

	// mySQLErrDuplicateKey is the error code for duplicate entries
	// https://dev.mysql.com/doc/refman/8.0/en/server-error-reference.html#error_er_dup_entry
	mySQLErrDuplicateKey uint16 = 1062

	mysqlMigrations = migrator.Migrations(
		execsql(
			"create_customer_name_watches",
			`create table if not exists customer_name_watches(id varchar(40) primary key, name varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_customer_status",
			`create table if not exists customer_status(customer_id varchar(40) primary key, user_id varchar(40), note varchar(1024), status varchar(10), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_customer_watches",
			`create table if not exists customer_watches(id varchar(40) primary key, customer_id varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_company_name_watches",
			`create table if not exists company_name_watches(id varchar(40) primary key, name varchar(256), webhook varchar(512), auth_token varchar(128), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_company_status",
			`create table if not exists company_status(company_id varchar(40) primary key, user_id varchar(40), note varchar(1024), status varchar(10), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_company_watches",
			`create table if not exists company_watches(id varchar(40) primary key, company_id varchar(40), webhook varchar(512), auth_token varchar(128), created_at timestamp(3), deleted_at timestamp(3));`,
		),
		execsql(
			"create_ofac_download_stats",
			`create table if not exists ofac_download_stats(downloaded_at timestamp(3), sdns integer, alt_names integer, addresses integer);`,
		),
		execsql(
			"create_webhook_stats",
			`create table if not exists webhook_stats(watch_id varchar(40), attempted_at timestamp(3), status varchar(10));`,
		),
		execsql(
			"add__denied_persons__to__ofac_download_stats",
			"alter table ofac_download_stats add column denied_persons integer not null default 0;",
		),
		execsql(
			"rename_ofac_download_stats",
			"rename table ofac_download_stats to download_stats",
		),
		execsql(
			"add_sectoral_sanctions_to_download_stats",
			"alter table download_stats add column sectoral_sanctions integer not null default 0;",
		),
		execsql(
			"add__bis_entities__to_download_stats",
			"alter table download_stats add column bis_entities integer not null default 0;",
		),
	)
)

type discardLogger struct{}

func (l discardLogger) Print(v ...interface{}) {}

func init() {
	gomysql.SetLogger(discardLogger{})
}

type mysql struct {
	dsn    string
	logger log.Logger

	connections *kitprom.Gauge
}

func (my *mysql) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", my.dsn)
	if err != nil {
		return nil, err
	}

	// Check out DB is up and working
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Migrate our database
	if m, err := migrator.New(mysqlMigrations); err != nil {
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

func mysqlConnection(logger log.Logger, user, pass string, address string, database string) *mysql {
	timeout := "30s"
	if v := os.Getenv("MYSQL_TIMEOUT"); v != "" {
		timeout = v
	}
	params := fmt.Sprintf("timeout=%s&charset=utf8mb4&parseTime=true&sql_mode=ALLOW_INVALID_DATES", timeout)
	dsn := fmt.Sprintf("%s:%s@%s/%s?%s", user, pass, address, database, params)
	return &mysql{
		dsn:         dsn,
		logger:      logger,
		connections: mysqlConnections,
	}
}

type MySQLConfig struct {
	Address  string
	Username string
	Password string
	Database string
}

func NewMySQLConnection(logger log.Logger, conf MySQLConfig) (*sql.DB, error) {
	my := mysqlConnection(logger, conf.Username, conf.Password, conf.Address, conf.Database)

	return my.Connect()
}

func TestMySQLConnection(t *testing.T) *sql.DB {
	t.Helper()

	if testing.Short() {
		t.Skip("-short flag was specified")
	}

	conf := CreateTestDatabase(t, TestDatabaseConfig())

	db, err := NewMySQLConnection(log.NewNopLogger(), conf)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func TestDatabaseConfig() MySQLConfig {
	return MySQLConfig{
		Address:  "tcp(localhost:3306)",
		Username: "root",
		Password: "root",
		Database: "watchman",
	}
}

func CreateTestDatabase(t *testing.T, config MySQLConfig) MySQLConfig {
	open := func() (*sql.DB, error) {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/", config.Username, config.Password, config.Address))
		if err != nil {
			return nil, err
		}

		if err := db.Ping(); err != nil {
			return nil, err
		}

		return db, nil
	}

	rootDb, err := open()
	for i := 0; err != nil && i < 22; i++ {
		time.Sleep(time.Second * 1)
		rootDb, err = open()
	}
	if err != nil {
		t.Fatal(err)
	}

	dbName := "test" + base.ID()
	_, err = rootDb.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		rootDb.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
		rootDb.Close()
	})

	config.Database = dbName

	return config
}

// MySQLUniqueViolation returns true when the provided error matches the MySQL code
// for duplicate entries (violating a unique table constraint).
func MySQLUniqueViolation(err error) bool {
	match := strings.Contains(err.Error(), fmt.Sprintf("Error %d: Duplicate entry", mySQLErrDuplicateKey))
	if e, ok := err.(*gomysql.MySQLError); ok {
		return match || e.Number == mySQLErrDuplicateKey
	}
	return match
}
