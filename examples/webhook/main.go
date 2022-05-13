// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman"

	"github.com/gorilla/mux"
)

var (
	httpAddr = flag.String("http.addr", ":10101", "HTTP listen address")
)

func main() {
	flag.Parse()

	logger := log.NewDefaultLogger()
	logger.Logf("Starting moov/watchman-webhook-example:%s", watchman.Version)

	// Listen for application termination.
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Setup HTTP handler
	handler := mux.NewRouter()
	addPingRoute(handler)
	addWebhookRoute(logger, handler)

	// Create main HTTP server
	serve := &http.Server{
		Addr:         *httpAddr,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.LogError(err)
		}
	}
	defer shutdownServer()

	// Start main HTTP server
	go func() {
		logger.Logf("listening on %s", *httpAddr)
		if err := serve.ListenAndServe(); err != nil {
			logger.LogError(err)
		}
	}()

	// Wait for app shutdown
	if err := <-errs; err != nil {
		logger.LogError(err)
	}
	os.Exit(0)
}

func addPingRoute(r *mux.Router) {
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("PONG"))
	})
}
