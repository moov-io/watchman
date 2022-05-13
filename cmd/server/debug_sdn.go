// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
)

const (
	debugSDNPath = "/debug/sdn/{sdnId}"
)

func debugSDNHandler(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sdnID := getSDNId(w, r)
		if sdnID == "" {
			return
		}

		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger.Info().With(log.Fields{
				"requestID": log.String(requestID),
			}).Logf("debug route for SDN=%s", sdnID)
		}

		var response struct {
			SDN   *SDN `json:"SDN"`
			Debug struct {
				IndexedName     string `json:"indexedName"`
				ParsedRemarksID string `json:"parsedRemarksId"`
			} `json:"debug"`
		}
		response.SDN = searcher.debugSDN(sdnID)
		if response.SDN != nil {
			response.Debug.IndexedName = response.SDN.name
			response.Debug.ParsedRemarksID = response.SDN.id
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
