package db

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/base/database"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/stretchr/testify/require"
)

func mysqlConfig() database.DatabaseConfig {
	return database.DatabaseConfig{
		DatabaseName: "watchman",
		MySQL: &database.MySQLConfig{
			Address:  "tcp(127.0.0.1:3306)",
			User:     "root",
			Password: "root",
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
	inCI := strx.Yes(os.Getenv("GITHUB_ACTIONS"))

	if inCI && !run {
		t.Skipf("not running ForEachDatabase on %v", runtime.GOOS)
	}

	configs := make(map[string]database.DatabaseConfig)
	configs["mysql"] = mysqlConfig()
	configs["postgres"] = postgresConfig()

	for name, config := range configs {
		t.Run(name, func(t *testing.T) {
			logger := log.NewTestLogger()

			// Connect and creat a test database
			db, _, err := New(config, logger)
			require.NoError(t, err)

			ctx := context.Background()
			dbName := "test" + base.ID()[0:26]

			_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
			require.NoError(t, err)

			t.Cleanup(func() {
				db.ExecContext(ctx, fmt.Sprintf("DROP DATABASE %s", dbName))
				db.Close()
			})

			// Close and reconnect
			db.Close()
			config.DatabaseName = dbName

			// Reconnect
			db, shutdown, err := New(config, logger)
			t.Cleanup(func() { shutdown() })

			// Run our test
			fn(db)

			err = db.Close()
			require.NoError(t, err)
		})
	}
}
