// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// GET /customers/:id
// - get customer information and matches

// POST /customers/name/:name/watch
// - monitor customer by name, re-parse on each search

// POST /customers/:id/watch
// - monitor customer

// PUT /customers/:id
// - mark customer as blocked or unblocked

// DELETE /customers/:id/watch
// - stop watching customer

// DELETE /customers/name/:name/watch
// - stop watching customer name

func addCustomerRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("GET").Path("/customers/{id}").HandlerFunc(getCustomer(logger))
	r.Methods("POST").Path("/customers/name/{name}/watch").HandlerFunc(addCustomerNameWatch(logger))
	r.Methods("POST").Path("/customers/{id}/watch").HandlerFunc(addCustomerWatch(logger))
	r.Methods("PUT").Path("/customers/{id}").HandlerFunc(updateCustomer(logger))
	r.Methods("DELETE").Path("/customers/{id}/watch").HandlerFunc(removeCustomerWatch(logger))
	r.Methods("DELETE").Path("/customers/name/{name}/watch").HandlerFunc(removeCustomerNameWatch(logger))
}

func getCustomer(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func addCustomerNameWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func addCustomerWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func updateCustomer(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func removeCustomerWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func removeCustomerNameWatch(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
