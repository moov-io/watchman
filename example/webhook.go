// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string                    `json:"id"`
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Comments  []*ofac.SDNComments       `json:"comments"`
}

func addWebhookRoute(logger log.Logger, r *mux.Router) {
	r.Methods("POST").Path("/ofac").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cust Customer

		if err := json.NewDecoder(r.Body).Decode(&cust); err != nil {
			logger.Log("webhook", fmt.Sprintf("problem reading request: %v", err))

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}
		if cust.ID == "" {
			logger.Log("webhook", "malformed webhook request")

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "malformed Customer JSON"}`))
			return
		}
		logger.Log("webhook", fmt.Sprintf("got webhook for Customer %s", cust.ID))
		w.WriteHeader(http.StatusOK)
	})
}
