package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/moov-io/base/database"
	"github.com/moov-io/base/log"
	root "github.com/moov-io/watchman"
)

type DB interface {
	Ping() error
	Close() error

	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

const (
	maxRetries = 22
)

func New(config database.DatabaseConfig, logger log.Logger, options ...Option) (DB, func(), error) {
	ctx, cancelFunc := context.WithCancel(context.Background())

	// connect to the database and keep retrying
	data, err := database.New(ctx, logger, config)
	for i := 0; err != nil && i < maxRetries; i++ {
		logger.Info().LogErrorf("attempt %d/%d to connect to database again: %v", i+1, maxRetries, err)

		time.Sleep(time.Second * 5)
		data, err = database.New(ctx, logger, config)
	}
	if err != nil {
		return nil, cancelFunc, logger.Fatal().LogErrorf("Error creating database: %w", err).Err()
	}

	shutdown := func() {
		logger.Info().Log("Shutting down the db")
		cancelFunc()
		if err := data.Close(); err != nil {
			logger.Fatal().LogErrorf("Error closing DB: %w", err)
		}
	}

	var dbType string
	var dbOpts []Option
	var migOpts []database.MigrateOption

	if config.MySQL != nil {
		dbType = "MySQL"
		migOpts = append(migOpts, database.WithEmbeddedMigrations(root.MySQLMigrations))
		dbOpts = append(dbOpts, MySQLRebind())
	} else if config.Postgres != nil {
		dbType = "Postgres"
		migOpts = append(migOpts, database.WithEmbeddedMigrations(root.PostgresMigrations))
		dbOpts = append(dbOpts, PostgresRebind())
	}

	// Run the migrations
	if err := database.RunMigrations(logger, config, migOpts...); err != nil {
		return nil, shutdown, logger.Fatal().LogErrorf("Error running %s migrations: %w", dbType, err).Err()
	}

	logger.Info().Logf("finished initializing %s db", dbType)

	out := &db{
		db:      data,
		onQuery: func(query string) string { return query },
	}

	return out, shutdown, nil
}

type db struct {
	db      DB
	onQuery func(query string) string
}

func (db *db) Ping() error {
	return db.db.Ping()
}

func (db *db) Close() error {
	return db.db.Close()
}

func (db *db) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	query = db.onQuery(query)

	return db.db.ExecContext(ctx, query, args...)
}

func (db *db) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	query = db.onQuery(query)

	return db.db.QueryContext(ctx, query, args...)
}

func (db *db) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	query = db.onQuery(query)

	return db.db.QueryRowContext(ctx, query, args...)
}
