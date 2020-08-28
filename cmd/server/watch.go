// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoWatchID = errors.New("no watchID found")
)

func getWatchID(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["watchID"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoWatchID)
		return ""
	}
	return v
}

type watch struct {
	id                       string
	customerID, customerName string
	companyID, companyName   string
	webhook                  string
	authToken                string
}

type watchCursor struct {
	batchSize int
	db        *sql.DB

	logger log.Logger

	// '*NewerThan' values represent the minimum (oldest) created_at value to return in the batch.
	//
	// These values start at "zero time" (an empty time.Time) and progresses towards time.Now()
	// with each batch by being set to the batch's newest time.
	customerNewerThan, customerNameNewerThan time.Time
	companyNewerThan, companyNameNewerThan   time.Time
}
