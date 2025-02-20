package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/postalpool/coder"
	"github.com/moov-io/watchman/pkg/address"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No PORT provided")
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Methods("GET").Path("/parse").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := address.ParseAddress(context.Background(), strings.TrimSpace(r.URL.Query().Get("address")))

		enc := coder.GetEncoder()
		enc.Reset(w)
		defer coder.SaveEncoder(enc)

		enc.Encode(addr)
	})

	// Setup HTTP server
	defaultTimeout := readDuration(os.Getenv("REQUEST_DURATION"), 10*time.Second) // default matches config.default.yml
	serve := &http.Server{
		Addr:    ":" + port,
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

	// Listen for errors
	errs := make(chan error, 1)

	// Start business logic HTTP server
	go func() {
		errs <- serve.ListenAndServe()
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
	}
}

func readDuration(input string, empty time.Duration) time.Duration {
	d, err := time.ParseDuration(input)
	if err == nil {
		return d
	}
	return empty
}
