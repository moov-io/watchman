// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	watchman "github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/download"

	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
)

type GlobalConfig struct {
	Watchman Config
}

type Config struct {
	Download download.Config

	Servers ServerConfig
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
