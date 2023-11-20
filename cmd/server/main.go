// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman"
	"github.com/moov-io/watchman/internal/database"

	"github.com/gorilla/mux"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ofac"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ofac"), "Admin HTTP listen address")

	flagBasePath  = flag.String("base-path", "/", "Base path to serve HTTP routes and webui from")
	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")
	flagMaxProcs  = flag.Int("max-procs", runtime.NumCPU(), "Maximum number of CPUs used for search and endpoints")
	flagWorkers   = flag.Int("workers", 1024, "Maximum number of goroutines used for search")

	dataRefreshInterval = 12 * time.Hour
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(*flagMaxProcs)

	var logger log.Logger
	if v := os.Getenv("LOG_FORMAT"); v != "" {
		*flagLogFormat = v
	}
	if strings.ToLower(*flagLogFormat) == "json" {
		logger = log.NewJSONLogger()
	} else {
		logger = log.NewDefaultLogger()
	}

	logger.Logf("Starting watchman server version %s", watchman.Version)

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("signal: %v", <-c)
	}()

	// Setup database connection
	db, err := database.New(logger, os.Getenv("DATABASE_TYPE"))
	if err != nil {
		logger.Logf("database problem: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.LogError(err)
		}
	}()

	// Setup business HTTP routes
	if v := os.Getenv("BASE_PATH"); v != "" {
		*flagBasePath = v
	}
	router := mux.NewRouter().PathPrefix(*flagBasePath).Subrouter()
	moovhttp.AddCORSHandler(router)
	addPingRoute(router)

	// Start business HTTP server
	readTimeout, _ := time.ParseDuration("30s")
	writTimeout, _ := time.ParseDuration("30s")
	idleTimeout, _ := time.ParseDuration("60s")

	// Check to see if our -http.addr flag has been overridden
	if v := os.Getenv("HTTP_BIND_ADDRESS"); v != "" {
		*httpAddr = v
	}

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readTimeout,
		WriteTimeout:      writTimeout,
		IdleTimeout:       idleTimeout,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.LogError(err)
		}
	}

	// Check to see if our -admin.addr flag has been overridden
	if v := os.Getenv("HTTP_ADMIN_BIND_ADDRESS"); v != "" {
		*adminAddr = v
	}

	// Start Admin server (with Prometheus metrics)
	adminServer, err := admin.New(admin.Opts{
		Addr: *adminAddr,
	})
	if err != nil {
		errs <- fmt.Errorf("problem starting admin server: %v", err)
	}
	adminServer.AddVersionHandler(watchman.Version) // Setup 'GET /version'
	go func() {
		logger.Logf("listening on %s", adminServer.BindAddr())
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			logger.LogError(err)
			errs <- fmt.Errorf("admin shutdown: %v", err)
		}
	}()
	defer adminServer.Shutdown()

	// Setup download repository
	downloadRepo := &sqliteDownloadRepository{db, logger}
	defer downloadRepo.close()

	var pipeline *pipeliner
	if debug, err := strconv.ParseBool(os.Getenv("DEBUG_NAME_PIPELINE")); debug && err == nil {
		pipeline = newPipeliner(logger)
	} else {
		pipeline = newPipeliner(log.NewNopLogger())
	}
	searcher := newSearcher(logger, pipeline, *flagWorkers)

	// Add debug routes
	adminServer.AddHandler(debugSDNPath, debugSDNHandler(logger, searcher))

	// Initial download of data
	if stats, err := searcher.refreshData(os.Getenv("INITIAL_DATA_DIRECTORY")); err != nil {
		logger.LogErrorf("ERROR: failed to download/parse initial data: %v", err)
		os.Exit(1)
	} else {
		if err := downloadRepo.recordStats(stats); err != nil {
			logger.LogErrorf("ERROR: failed to record download stats: %v", err)
			os.Exit(1)
		}
		logger.Info().With(log.Fields{
			"SDNs":             log.Int(stats.SDNs),
			"AltNames":         log.Int(stats.Alts),
			"Addresses":        log.Int(stats.Addresses),
			"SSI":              log.Int(stats.SectoralSanctions),
			"DPL":              log.Int(stats.DeniedPersons),
			"BISEntities":      log.Int(stats.BISEntities),
			"UVL":              log.Int(stats.Unverified),
			"ISN":              log.Int(stats.NonProliferationSanctions),
			"FSE":              log.Int(stats.ForeignSanctionsEvaders),
			"PLC":              log.Int(stats.PalestinianLegislativeCouncil),
			"CAP":              log.Int(stats.CAPTA),
			"DTC":              log.Int(stats.ITARDebarred),
			"CMIC":             log.Int(stats.ChineseMilitaryIndustrialComplex),
			"NS_MBS":           log.Int(stats.NonSDNMenuBasedSanctions),
			"EU_CSL":           log.Int(stats.EUCSL),
			"UK_CSL":           log.Int(stats.UKCSL),
			"UK_SanctionsList": log.Int(stats.UKSanctionsList),
		}).Logf("data refreshed %v ago", time.Since(stats.RefreshedAt))
	}

	// Setup periodic download and re-search
	updates := make(chan *DownloadStats)
	dataRefreshInterval = getDataRefreshInterval(logger, os.Getenv("DATA_REFRESH_INTERVAL"))
	go searcher.periodicDataRefresh(dataRefreshInterval, downloadRepo, updates)
	go handleDownloadStats(updates, func(stats *DownloadStats) {
		callDownloadWebook(logger, stats)
	})

	// Add manual data refresh endpoint
	adminServer.AddHandler(manualRefreshPath, manualRefreshHandler(logger, searcher, updates, downloadRepo))

	// Add searcher for HTTP routes
	addSDNRoutes(logger, router, searcher)
	addSearchRoutes(logger, router, searcher)
	addDownloadRoutes(logger, router, downloadRepo)
	addValuesRoutes(logger, router, searcher)

	// Setup our web UI to be served as well
	setupWebui(logger, router, *flagBasePath)

	// Start business logic HTTP server
	go func() {
		if certFile, keyFile := os.Getenv("HTTPS_CERT_FILE"), os.Getenv("HTTPS_KEY_FILE"); certFile != "" && keyFile != "" {
			logger.Logf("binding to %s for secure HTTP server", *httpAddr)
			if err := serve.ListenAndServeTLS(certFile, keyFile); err != nil {
				logger.LogErrorf("https shutdown: %v", err)
			}
		} else {
			logger.Logf("binding to %s for HTTP server", *httpAddr)
			if err := serve.ListenAndServe(); err != nil {
				logger.LogErrorf("http shutdown: %v", err)
			}
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
		moovhttp.SetAccessControlAllowHeaders(w, r.Header.Get("Origin"))
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})
}

// getDataRefreshInterval returns a time.Duration for how often OFAC should refresh data
//
// env is the value from an environmental variable
func getDataRefreshInterval(logger log.Logger, env string) time.Duration {
	if env != "" {
		if strings.EqualFold(env, "off") {
			return 0 * time.Second
		}
		if dur, _ := time.ParseDuration(env); dur > 0 {
			logger.Logf("Setting data refresh interval to %v", dur)
			return dur
		}
	}
	logger.Logf("Setting data refresh interval to %v (default)", dataRefreshInterval)
	return dataRefreshInterval
}

func setupWebui(logger log.Logger, r *mux.Router, basePath string) {
	dir := os.Getenv("WEB_ROOT")
	if dir == "" {
		dir = filepath.Join("webui", "build")
	}
	if _, err := os.Stat(dir); err != nil {
		logger.Logf("problem with webui=%s: %v", dir, err)
		os.Exit(1)
	}
	r.PathPrefix("/").Handler(http.StripPrefix(basePath, http.FileServer(http.Dir(dir))))
}

func handleDownloadStats(updates chan *DownloadStats, handle func(stats *DownloadStats)) {
	for {
		stats := <-updates
		if stats != nil {
			handle(stats)
		}
	}
}
