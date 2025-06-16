package db

import (
	"os"
	"runtime"
	"testing"

	"github.com/moov-io/base/database"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
)

func mysqlConfig() database.DatabaseConfig {
	return database.DatabaseConfig{
		DatabaseName: "watchman",
		MySQL: &database.MySQLConfig{
			Address:  "tcp(127.0.0.1:3306)",
			User:     "watchman",
			Password: "watchman",
		},
	}
}

func postgresConfig() database.DatabaseConfig {
	return database.DatabaseConfig{
		DatabaseName: "watchman",
		Postgres: &database.PostgresConfig{
			Address:  "127.0.0.1:5432",
			User:     "watchman",
			Password: "watchman",
		},
	}
}

func ForEachDatabase(t *testing.T, fn func(db DB)) {
	t.Helper()

	run := strx.Yes(os.Getenv("RUN_DATABASE_TESTS"))
	if !run {
		t.Skipf("not running ForEachDatabase on %v", runtime.GOOS)
	}

	configs := make(map[string]database.DatabaseConfig)
	configs["mysql"] = mysqlConfig()
	configs["postgres"] = postgresConfig()

	for name, config := range configs {
		t.Run(name, func(t *testing.T) {
			logger := log.NewTestLogger()

			db, shutdown, err := New(config, logger)
			if err != nil {
				t.Fatal(err)
				return
			}
			t.Cleanup(func() { shutdown() })

			fn(db)
		})
	}
}
