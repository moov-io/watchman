// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string                    `json:"id"`
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Match     float64                   `json:"match,omitempty"`
}

type Company struct {
	ID        string                    `json:"id,omitempty"`
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Match     float64                   `json:"match,omitempty"`
}

func addWebhookRoute(logger log.Logger, r *mux.Router) {
	r.Methods("POST").Path("/ofac").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(io.LimitReader(r.Body, 5*1024*1024))
		if err != nil {
			logger.Logf("problem reading request: %v", err)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		}

		if cust := readCustomer(bytes.NewReader(bs)); cust != nil {
			logger.Logf("got webhook for Customer %s (%s) match=%.2f", cust.ID, cust.SDN.SDNName, cust.Match)
			w.WriteHeader(http.StatusOK)
			return
		}
		if company := readCompany(bytes.NewReader(bs)); company != nil {
			logger.Logf("got webhook for Company %s (%s) match=%.2f", company.ID, company.SDN.SDNName, company.Match)
			w.WriteHeader(http.StatusOK)
		}

		logger.Log("malformed webhook request")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "malformed JSON"}`))
	})
}

func readCustomer(r io.Reader) *Customer {
	var cust Customer
	if err := json.NewDecoder(r).Decode(&cust); err != nil || cust.ID == "" {
		return nil
	}
	return &cust
}

func readCompany(r io.Reader) *Company {
	var company Company
	if err := json.NewDecoder(r).Decode(&company); err != nil || company.ID == "" {
		return nil
	}
	return &company
}
