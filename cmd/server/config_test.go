// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	logger := log.NewTestLogger()

	conf, err := LoadConfig(logger)
	require.NoError(t, err)
	require.NotNil(t, conf)

	require.Equal(t, ":8084", conf.Servers.BindAddress)
	require.Equal(t, 12*time.Hour, conf.Download.RefreshInterval)
}
