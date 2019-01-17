// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

// GET /sdn/:id/addresses
// - get addresses for a given SDN

// GET /sdn/:id/alternateNames
// - get alternate names for a given SDN

// GET /sdn/:id
// - get SDN information

func addSDNRoutes(logger log.Logger, r *mux.Router) {
	r.Methods("GET").Path("/sdn/{id}/addresses").HandlerFunc(getSDNAddresses(logger))
	r.Methods("GET").Path("/sdn/{id}/alts").HandlerFunc(getSDNAltNames(logger))
	r.Methods("GET").Path("/sdn/{id}").HandlerFunc(getSDN(logger))
}

func getSDNAddresses(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		addresses := []*ofac.Address{
			{
				// 587,356,"Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku","Tokyo 103","Japan",-0-
				EntityID:                    "587",
				AddressID:                   "356",
				Address:                     "Dai-Ichi Bldg. 6th Floor, 10-2 Nihombashi, 2-chome, Chuo-ku",
				CityStateProvincePostalCode: "Tokyo 103",
				Country:                     "Japan",
			},
			{
				// 651,376,"Case Postale 236, 10 Bis Rue Du Vieux College 12-11","Geneva","Switzerland",-0-
				EntityID:                    "651",
				AddressID:                   "376",
				Address:                     "Case Postale 236, 10 Bis Rue Du Vieux College 12-11",
				CityStateProvincePostalCode: "Geneva",
				Country:                     "Switzerland",
			},
		}
		if err := json.NewEncoder(w).Encode(addresses); err != nil {
			moovhttp.Problem(w, err) // TODO(adam): JSON errors should moovhttp.InternalError (wrapped, see auth's http.go)
			return
		}
	}
}

func getSDNAltNames(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		alts := []*ofac.AlternateIdentity{
			{
				// 815,641,"aka","GALAX INC.",-0-
				EntityID:      "815",
				AlternateID:   "641",
				AlternateType: "aka",
				AlternateName: "GALAX INC",
			},
			{
				// 4359,3565,"fka","INDUSTRIA AVICOLA PALMASECA S.A.",-0-
				EntityID:      "4359",
				AlternateID:   "3565",
				AlternateType: "fkaa",
				AlternateName: "INDUSTRIA AVICOLA PALMASECA S.A.",
			},
		}
		if err := json.NewEncoder(w).Encode(alts); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func getSDN(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		sdn := &ofac.SDN{
			EntityID: "3754",
			SDNName:  "ABU MARZOOK, Mousa Mohammed",
			SDNType:  "individual",
			Program:  "SDGT] [SDT",
			Title:    "Political Leader in Amman, Jordan and Damascus, Syria for HAMAS",
			Remarks:  "DOB 09 Feb 1951; POB Gaza, Egypt; Passport 92/664 (Egypt); SSN 523-33-8386 (United States); Political Leader in Amman, Jordan and Damascus, Syria for HAMAS; a.k.a. 'ABU-'UMAR'.",
		}
		if err := json.NewEncoder(w).Encode(sdn); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
