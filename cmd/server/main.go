// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"cmp"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/postalpool"
	"github.com/moov-io/watchman/internal/search"

	"github.com/gorilla/mux"
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	logger := log.NewDefaultLogger().With(log.Fields{
		"app":     log.String("watchman"),
		"version": log.String(watchman.Version),
	})
	logger.Log("Starting watchman server")

	// Set runtime.GOMAXPROCS
	maxprocs.Set(maxprocs.Logger(logger.Info().Logf))

	config, err := LoadConfig(logger)
	if err != nil {
		logger.Fatal().LogErrorf("problem loading config: %v", err)
		os.Exit(1)
	}

	// Setup telemetry
	telemetryShutdownFunc, err := telemetry.SetupTelemetry(context.Background(), config.Telemetry, watchman.Version)
	if err != nil {
		logger.Fatal().LogErrorf("setting up telemetry failed: %w", err)
		os.Exit(1)
	}
	defer telemetryShutdownFunc()

	downloader, err := download.NewDownloader(logger, config.Download)
	if err != nil {
		logger.Fatal().LogErrorf("problem setting up downloader: %v", err)
		os.Exit(1)
	}

	// Setup signal listener
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// Set up a channel to listen for system interrupt signals
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Listen for errors
	errs := make(chan error, 1)

	// Setup search service and endpoints
	searchService, err := search.NewService(logger, config.Search)
	if err != nil {
		logger.Fatal().LogErrorf("problem setting up search service: %v", err)
		os.Exit(1)
	}
	err = setupPeriodicRefreshing(ctx, logger, errs, config.Download, downloader, searchService)
	if err != nil {
		logger.Fatal().LogErrorf("problem during initial download: %v", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	addPingRoute(router)

	addressParsingPool, err := postalpool.NewService(logger, config.PostalPool)
	if err != nil {
		logger.Fatal().LogErrorf("problem setting up address parsing pool: %v", err)
		os.Exit(1)
	}
	searchController := search.NewController(logger, searchService, addressParsingPool)
	searchController.AppendRoutes(router)

	// Start Admin server (with Prometheus metrics)
	adminServer, err := admin.New(admin.Opts{
		Addr: config.Servers.AdminAddress,
	})
	if err != nil {
		errs <- fmt.Errorf("problem starting admin server: %v", err)
	} else {
		adminServer.AddVersionHandler(watchman.Version) // Setup 'GET /version'
	}
	go func() {
		if adminServer == nil {
			return
		}

		logger.Logf("listening on %s", adminServer.BindAddr())

		if err := adminServer.Listen(); err != nil {
			errs <- logger.Error().LogErrorf("admin server shutdown: %v", err).Err()
		}
	}()
	defer func() {
		if adminServer != nil {
			adminServer.Shutdown()
		}
	}()

	// Setup HTTP server
	defaultTimeout := 20 * time.Second
	serve := &http.Server{
		Addr:    config.Servers.BindAddress,
		Handler: router,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       defaultTimeout,
		ReadHeaderTimeout: defaultTimeout,
		WriteTimeout:      defaultTimeout,
		IdleTimeout:       defaultTimeout,
	}
	shutdownServer := func() {
		if serve != nil {
			serve.Shutdown(context.TODO())
		}
	}

	// Start business logic HTTP server
	go func() {
		certFile := cmp.Or(os.Getenv("HTTPS_CERT_FILE"), config.Servers.TLSCertFile)
		keyFile := cmp.Or(os.Getenv("HTTPS_KEY_FILE"), config.Servers.TLSKeyFile)

		if certFile != "" && keyFile != "" {
			logger.Logf("binding to %s for secure HTTP server", config.Servers.BindAddress)
			errs <- serve.ListenAndServeTLS(certFile, keyFile)
		} else {
			logger.Logf("binding to %s for HTTP server", config.Servers.BindAddress)
			errs <- serve.ListenAndServe()
		}
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
		logger.LogErrorf("final exit: %v", err)
	}
}

func addPingRoute(r *mux.Router) {
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})
}
