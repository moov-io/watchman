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
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/docker"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/base/k8s"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ofac"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ofac"), "Admin HTTP listen address")
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

	// Setup business HTTP routes
	router := mux.NewRouter()
	moovhttp.AddCORSHandler(router)
	addPingRoute(router)
	addCustomerRoutes(logger, router)
	addSDNRoutes(logger, router)
	addSearchRoutes(logger, router)

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

	wg := sync.WaitGroup{}

	// Start Elasticsearch for local dev if Docker exists
	var es *ofac.Elasticsearch
	wg.Add(1)
	go func() {
		defer wg.Done()
		db, err := startElasticsearch(logger)
		if err != nil {
			panic(fmt.Sprintf("ERROR: starting elasticsearch: %v", err))
		}
		es = db
	}()
	defer func() {
		if es != nil {
			es.Stop(logger)
		}
	}()

	// Download OFAC data
	var reader *ofac.Reader
	wg.Add(1)
	go func() {
		defer wg.Done()
		r, err := getAndParseOFACData()
		if err != nil {
			panic(err.Error())
		}
		reader = r
	}()
	logger.Log("main", fmt.Sprintf("OFAC data downloaded and parsed: Addresses=%d AltNames=%d SDNs=%d", len(reader.AddressArray), len(reader.AlternateIdentityArray), len(reader.SDNArray)))

	// Block until startup processes are finished
	wg.Wait()

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

func startElasticsearch(logger log.Logger) (*ofac.Elasticsearch, error) {
	if !k8s.Inside() && docker.Enabled() {
		es, err := ofac.NewElasticsearch(logger)
		if err != nil {
			return nil, err
		}
		logger.Log("main", fmt.Sprintf("ES is up and running! (Docker container ID: %s)", es.ID()))
		return es, nil
	}
	return nil, nil
}

func getAndParseOFACData() (*ofac.Reader, error) {
	dir, err := (&ofac.Downloader{}).GetFiles()
	if err != nil {
		return nil, fmt.Errorf("ERROR: downloading OFAC data: %v", err)
	}

	// Parse each OFAC file
	r := &ofac.Reader{}
	r.FileName = filepath.Join(dir, "add.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading add.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "alt.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading alt.csv: %v", err)
	}
	r.FileName = filepath.Join(dir, "sdn.csv")
	if err := r.Read(); err != nil {
		return nil, fmt.Errorf("ERROR: reading sdn.csv: %v", err)
	}
	return r, nil
}
