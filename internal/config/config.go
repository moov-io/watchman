// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package config

import (
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/ingest"
	"github.com/moov-io/watchman/internal/postalpool"
	"github.com/moov-io/watchman/internal/search"
	"github.com/moov-io/watchman/internal/webui"

	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
)

type GlobalConfig struct {
	Watchman Config
}

type Config struct {
	Servers   ServerConfig
	Telemetry telemetry.Config

	Webui webui.Config

	Download   download.Config
	Search     search.Config
	PostalPool postalpool.Config

	Ingest ingest.Config
}

type ServerConfig struct {
	BindAddress  string
	AdminAddress string

	TLSCertFile string
	TLSKeyFile  string
}

func LoadConfig(logger log.Logger) (*Config, error) {
	configService := config.NewService(logger)

	global := &GlobalConfig{}
	if err := configService.LoadFromFS(global, watchman.ConfigDefaults); err != nil {
		return nil, err
	}

	return &global.Watchman, nil
}
