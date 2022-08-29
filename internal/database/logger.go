// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"os"
	"strings"

	"github.com/moov-io/base/log"

	"github.com/lopezator/migrator"
)

type migrationLogger struct {
	logger    log.Logger
	shouldLog bool
}

func newMigrationLogger(logger log.Logger) migrator.Logger {
	level := strings.TrimSpace(os.Getenv("LOG_LEVEL"))
	shouldLog := strings.EqualFold("trace", level)

	return &migrationLogger{
		logger:    logger,
		shouldLog: shouldLog,
	}
}

func (ml *migrationLogger) Printf(pattern string, args ...interface{}) {
	if ml != nil && ml.logger != nil {
		if ml.shouldLog {
			ml.logger.Info().Logf(pattern, args...)
		}
	}
}
