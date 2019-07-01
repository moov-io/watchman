// Copyright 2018 The Moov Authors
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
	"strings"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/cardonator/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ofac"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ofac"), "Admin HTTP listen address")

	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")

	ofacDataRefreshInterval = 12 * time.Hour
)

func main() {
	flag.Parse()

	var logger log.Logger
	if v := os.Getenv("LOG_FORMAT"); v != "" {
		*flagLogFormat = v
	}
	if strings.ToLower(*flagLogFormat) == "json" {
		logger = log.NewJSONLogger(os.Stderr)
	} else {
		logger = log.NewLogfmtLogger(os.Stderr)
	}
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("startup", fmt.Sprintf("Starting ofac server version %s", ofac.Version))

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Setup SQLite database
	if sqliteVersion, _, _ := sqlite3.Version(); sqliteVersion != "" {
		logger.Log("main", fmt.Sprintf("sqlite version %s", sqliteVersion))
	}
	db, err := createSqliteConnection(logger, getSqlitePath())
	if err != nil {
		logger.Log("main", err)
		os.Exit(1)
	}
	if err := migrate(logger, db); err != nil {
		logger.Log("main", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Log("main", err)
		}
	}()

	// Setup business HTTP routes
	router := mux.NewRouter()
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
		ReadTimeout:  readTimeout,
		WriteTimeout: writTimeout,
		IdleTimeout:  idleTimeout,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.Log("shutdown", err)
		}
	}

	// Check to see if our -admin.addr flag has been overridden
	if v := os.Getenv("HTTP_ADMIN_BIND_ADDRESS"); v != "" {
		*adminAddr = v
	}

	// Start Admin server (with Prometheus metrics)
	adminServer := admin.NewServer(*adminAddr)
	go func() {
		logger.Log("admin", fmt.Sprintf("listening on %s", adminServer.BindAddr()))
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			logger.Log("admin", err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	// Setup download repository
	downloadRepo := &sqliteDownloadRepository{db, logger}
	defer downloadRepo.close()

	// Start our searcher (and downloader)
	searcher := &searcher{
		logger: logger,
	}
	if stats, err := searcher.refreshData(); err != nil {
		logger.Log("main", fmt.Sprintf("ERROR: failed to download/parse initial OFAC data: %v", err))
		os.Exit(1)
	} else {
		downloadRepo.recordStats(stats)
		logger.Log("main", fmt.Sprintf("OFAC data refreshed - Addresses=%d AltNames=%d SDNs=%d", stats.Addresses, stats.Alts, stats.SDNs))
	}

	// Setup Watch and Webhook database wrapper
	watchRepo := &sqliteWatchRepository{db, logger}
	defer watchRepo.close()
	webhookRepo := &sqliteWebhookRepository{db}
	defer webhookRepo.close()

	// Setup company / customer repositories
	companyRepo := &sqliteCompanyRepository{db}
	defer companyRepo.close()
	custRepo := &sqliteCustomerRepository{db}
	defer custRepo.close()

	// Setup periodic download and re-search
	updates := make(chan *downloadStats)
	ofacDataRefreshInterval = getOFACRefreshInterval(logger, os.Getenv("OFAC_DATA_REFRESH"))
	go searcher.periodicDataRefresh(ofacDataRefreshInterval, downloadRepo, updates)
	go searcher.spawnResearching(logger, companyRepo, custRepo, watchRepo, webhookRepo, updates)

	// Add manual OFAC data refresh endpoint
	adminServer.AddHandler(manualRefreshPath, manualRefreshHandler(logger, searcher, downloadRepo))

	// Add searcher for HTTP routes
	addCompanyRoutes(logger, router, searcher, companyRepo, watchRepo)
	addCustomerRoutes(logger, router, searcher, custRepo, watchRepo)
	addSDNRoutes(logger, router, searcher)
	addSearchRoutes(logger, router, searcher)
	addDownloadRoutes(logger, router, downloadRepo)

	// Start business logic HTTP server
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- serve.ListenAndServe()
		// TODO(adam): support TLS
		// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
		logger.Log("exit", err)
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

// env is the value from an environmental variable
func getOFACRefreshInterval(logger log.Logger, env string) time.Duration {
	if env != "" {
		dur, _ := time.ParseDuration(env)
		if dur > 0 {
			if logger != nil {
				logger.Log("main", fmt.Sprintf("Setting OFAC data refresh interval to %v", dur))
			}
			return dur
		}
	}
	return ofacDataRefreshInterval
}
