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
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ofac"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ofac"), "Admin HTTP listen address")

	ofacDataRefreshInterval = 12 * time.Hour
)

func main() {
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
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

	// Start our searcher (and downloader)
	searcher := &searcher{
		logger: logger,
	}
	if err := searcher.refreshData(); err != nil {
		logger.Log("main", fmt.Sprintf("ERROR: failed to download/parse initial OFAC data: %v", err))
		os.Exit(1)
	}
	go func() {
		// Override refresh interval if set
		ofacDataRefreshInterval = getOFACRefreshInterval(logger, os.Getenv("OFAC_DATA_REFRESH"))
		for {
			time.Sleep(ofacDataRefreshInterval)
			if err := searcher.refreshData(); err != nil {
				logger.Log("main", fmt.Sprintf("ERROR: refreshing OFAC data: %v", err))
			} else {
				searcher.RLock()
				logger.Log("main", fmt.Sprintf("OFAC data refreshed - Addresses=%d AltNames=%d SDNs=%d", len(searcher.Addresses), len(searcher.Alts), len(searcher.SDNs)))
				searcher.RUnlock()
			}
		}
	}()

	// Setup Database wrappers
	watchRepo := &sqliteWatchRepository{db, logger}

	// Add searcher for HTTP routes
	addCustomerRoutes(logger, router, searcher, watchRepo)
	addSDNRoutes(logger, router, searcher)
	addSearchRoutes(logger, router, searcher)

	// Add admin server OFAC data refresh endpoint
	adminServer.AddHandler("/ofac/refresh", func(w http.ResponseWriter, r *http.Request) {
		logger.Log("main", "admin: refreshing OFAC data")
		if err := searcher.refreshData(); err != nil {
			logger.Log("main", fmt.Sprintf("ERROR: admin: problem refreshing OFAC data: %v", err))
			w.WriteHeader(http.StatusOK)
		} else {
			logger.Log("main", "admin: finished OFAC data refresh")
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

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
