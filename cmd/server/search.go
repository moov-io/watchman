// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"
	// "github.com/moov-io/ofac/pkg/strcmp"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")
)

func addSearchRoutes(logger log.Logger, r *mux.Router, searcher *searcher) {
	r.Methods("GET").Path("/search").HandlerFunc(search(logger, searcher))
}

type addressSearchRequest struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
}

func (req addressSearchRequest) empty() bool {
	return req.Address == "" && req.City == "" && req.State == "" &&
		req.Providence == "" && req.Zip == "" && req.Country == ""
}

func readAddressSearchRequest(u *url.URL) addressSearchRequest {
	return addressSearchRequest{
		Address:    strings.ToLower(strings.TrimSpace(u.Query().Get("address"))),
		City:       strings.ToLower(strings.TrimSpace(u.Query().Get("city"))),
		State:      strings.ToLower(strings.TrimSpace(u.Query().Get("state"))),
		Providence: strings.ToLower(strings.TrimSpace(u.Query().Get("providence"))),
		Zip:        strings.ToLower(strings.TrimSpace(u.Query().Get("zip"))),
		Country:    strings.ToLower(strings.TrimSpace(u.Query().Get("country"))),
	}
}

func search(logger log.Logger, searcher *searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w, err := wrapResponseWriter(logger, w, r)
		if err != nil {
			return
		}

		// Search by Name // TODO(adam): handle multiple?
		if name := strings.TrimSpace(r.URL.Query().Get("name")); name != "" {
			if logger != nil {
				logger.Log("search", fmt.Sprintf("searching SDN names for %s", name))
			}
			searchByName(logger, searcher, name)(w, r)
			return
		}

		// Search by Alt Name
		if alt := strings.TrimSpace(r.URL.Query().Get("altName")); alt != "" {
			if logger != nil {
				logger.Log("search", fmt.Sprintf("searching SDN alt names for %s", alt))
			}
			searchByAltName(logger, searcher, alt)(w, r)
			return
		}

		// Search Addresses
		if req := readAddressSearchRequest(r.URL); !req.empty() {
			if logger != nil {
				logger.Log("search", fmt.Sprintf("searching address for %#v", req))
			}
			searchByAddress(logger, searcher, req)(w, r)
			return
		}

		// Fallback if no search params were found
		moovhttp.Problem(w, errNoSearchParams)
	}
}

type searcher struct {
	SDNs []*SDN

	Addresses []*Address

	Alts []*Alt
}

// SDN is ofac.SDN wrapped with precomputed search metadata
type SDN struct {
	SDN *ofac.SDN

	// name is precomputed as lowercase'd and split on words
	name []string
}

func precomputeSDNs(sdns []*ofac.SDN) []*SDN {
	out := make([]*SDN, len(sdns))
	for i := range sdns {
		out[i] = &SDN{
			SDN:  sdns[i],
			name: precompute(sdns[i].SDNName),
		}
	}
	return out
}

// Address is ofac.Address wrapped with precomputed search metadata
type Address struct {
	Address *ofac.Address

	// precomputed (lowercase and split) fields for speed
	address, citystate, country []string
}

func precomputeAddresses(adds []*ofac.Address) []*Address {
	out := make([]*Address, len(adds))
	for i := range adds {
		out[i] = &Address{
			Address:   adds[i],
			address:   precompute(adds[i].Address),
			citystate: precompute(adds[i].CityStateProvincePostalCode),
			country:   precompute(adds[i].Country),
		}
	}
	return out
}

// Alt is an ofac.AlternateIdentity wrapped with precomputed search metadata
type Alt struct {
	AlternateIdentity *ofac.AlternateIdentity

	// name is precomputed (lowercase and split) for speed
	name []string
}

func precomputeAlts(alts []*ofac.AlternateIdentity) []*Alt {
	out := make([]*Alt, len(alts))
	for i := range alts {
		out[i] = &Alt{
			AlternateIdentity: alts[i],
			name:              precompute(alts[i].AlternateName),
		}
	}
	return out
}

// precompute will split s on white space and lowercase each substring
func precompute(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

func searchByAddress(logger log.Logger, searcher *searcher, req addressSearchRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hasAddress := req.Address != ""
		reqAdds := strings.Fields(strings.ToLower(req.Address))

		var answer []*ofac.Address
		for i := range searcher.Addresses {
			add := searcher.Addresses[i]
			if hasAddress {
				// Count matches for collection if over threshold
				matches := 0
				for k := range add.address {
					for j := range reqAdds {
						if strings.Contains(add.address[k], reqAdds[j]) {
							matches++
						}
					}
				}
				// If over 25% of words from query match (via strings.Contains not full string equality) save as an address.
				// This is arbitrary, but given the following examples only one partial word match is required:
				//  123 Scott Ave
				//  1600 N Penn St
				if (float64(matches) / float64(len(add.address))) >= 0.25 {
					answer = append(answer, add.Address)
				}
				continue
			}
		}

		// score := strcmp.Levenshtein(add.Address, req.Address)
		// if score > .75 {
		// 	acc.add(add)
		// }

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answer); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func searchByName(logger log.Logger, searcher *searcher, nameSlug string) http.HandlerFunc { // TODO(Adam): split nameSlug
	return func(w http.ResponseWriter, r *http.Request) {
		nameSlugs := strings.Fields(strings.ToLower(nameSlug))
		if len(nameSlugs) == 0 {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		var answers []*ofac.SDN
		for i := range searcher.SDNs {
			sdn := searcher.SDNs[i]

			// Count matches for nameSlugs fields
			matches := 0
			for k := range sdn.name {
				for j := range nameSlugs {
					if strings.Contains(sdn.name[k], nameSlugs[j]) {
						matches++
					}
				}
			}
			if matches > 0 {
				answers = append(answers, sdn.SDN)
			}
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answers); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}

func searchByAltName(logger log.Logger, searcher *searcher, altSlug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		altSlugs := strings.Fields(strings.ToLower(altSlug))
		if len(altSlugs) == 0 {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		var answers []*ofac.AlternateIdentity
		for i := range searcher.Alts {
			alt := searcher.Alts[i]

			matches := 0
			for k := range alt.name {
				for j := range altSlugs {
					if strings.Contains(alt.name[k], altSlugs[j]) {
						matches++
					}
				}
			}
			if matches > 0 {
				answers = append(answers, alt.AlternateIdentity)
			}
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(answers); err != nil {
			moovhttp.Problem(w, err)
			return
		}
	}
}
