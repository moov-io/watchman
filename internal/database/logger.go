// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package database

import (
	"testing"

	"github.com/moov-io/base/log"

	"github.com/lopezator/migrator"
)

type migrationLogger struct {
	logger log.Logger
}

func newMigrationLogger(logger log.Logger) migrator.Logger {
	return &migrationLogger{
		logger: logger,
	}
}

func (ml *migrationLogger) Printf(pattern string, args ...interface{}) {
	if ml != nil && ml.logger != nil {
		if testing.Verbose() {
			ml.logger.Info().Logf(pattern, args...)
		}
	}
}
