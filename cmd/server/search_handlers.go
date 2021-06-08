// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	matchHist = prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
		Name:    "match_percentages",
		Help:    "Histogram representing the match percent of search routes",
		Buckets: []float64{0.0, 0.5, 0.8, 0.9, 0.99},
	}, []string{"type"})
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
		w = wrapResponseWriter(logger, w, r)
		requestID := moovhttp.GetRequestID(r)

		// Search over all fields
		if q := strings.TrimSpace(r.URL.Query().Get("q")); q != "" {
			logger.Log("search", fmt.Sprintf("searching all names and address for %s", q), "requestID", requestID)
			searchViaQ(logger, searcher, q)(w, r)
			return
		}

		// Search by ID (found in an SDN's Remarks property)
		if id := strings.TrimSpace(r.URL.Query().Get("id")); id != "" {
			logger.Log("search", fmt.Sprintf("searching SDNs by remarks ID for %s", id))
			searchByRemarksID(logger, searcher, id)(w, r)
			return
		}

		// Search by Name
		if name := strings.TrimSpace(r.URL.Query().Get("name")); name != "" {
			if req := readAddressSearchRequest(r.URL); !req.empty() {
				logger.Log("search", fmt.Sprintf("searching SDN names='%s' and addresses", name), "requestID", requestID)
				searchViaAddressAndName(logger, searcher, name, req)(w, r)
				return
			}

			logger.Log("search", fmt.Sprintf("searching SDN names for %s", name), "requestID", requestID)
			searchByName(logger, searcher, name)(w, r)
			return
		}

		// Search by Alt Name
		if alt := strings.TrimSpace(r.URL.Query().Get("altName")); alt != "" {
			logger.Log("search", fmt.Sprintf("searching SDN alt names for %s", alt), "requestID", requestID)
			searchByAltName(logger, searcher, alt)(w, r)
			return
		}

		// Search Addresses
		if req := readAddressSearchRequest(r.URL); !req.empty() {
			logger.Log("search", fmt.Sprintf("searching address for %#v", req), "requestID", requestID)
			searchByAddress(logger, searcher, req)(w, r)
			return
		}

		// Fallback if no search params were found
		moovhttp.Problem(w, errNoSearchParams)
	}
}

type searchResponse struct {
	// OFAC
	SDNs              []*SDN    `json:"SDNs"`
	AltNames          []Alt     `json:"altNames"`
	Addresses         []Address `json:"addresses"`
	SectoralSanctions []SSI     `json:"sectoralSanctions"`
	// BIS
	DeniedPersons []DP        `json:"deniedPersons"`
	BISEntities   []BISEntity `json:"bisEntities"`
	// Metadata
	RefreshedAt time.Time `json:"refreshedAt"`
}

func buildAddressCompares(req addressSearchRequest) []func(*Address) *item {
	var compares []func(*Address) *item
	if req.Address != "" {
		compares = append(compares, topAddressesAddress(req.Address))
	}
	if req.City != "" {
		compares = append(compares, topAddressesCityState(req.City))
	}
	if req.State != "" {
		compares = append(compares, topAddressesCityState(req.State))
	}
	if req.Providence != "" {
		compares = append(compares, topAddressesCityState(req.Providence))
	}
	if req.Zip != "" {
		compares = append(compares, topAddressesCityState(req.Zip))
	}
	if req.Country != "" {
		compares = append(compares, topAddressesCountry(req.Country))
	}
	return compares
}

func searchByAddress(logger log.Logger, searcher *searcher, req addressSearchRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if req.empty() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := searchResponse{
			RefreshedAt: searcher.lastRefreshedAt,
		}
		limit := extractSearchLimit(r)
		minMatch := extractSearchMinMatch(r)

		// Perform our ranking across all accumulated compare functions
		//
		// TODO(adam): Is there something in the (SDN?) files which signal to block an entire country? (i.e. Needing to block Iran all together)
		// https://www.treasury.gov/resource-center/sanctions/CivPen/Documents/20190327_decker_settlement.pdf
		compares := buildAddressCompares(req)

		filtered := searcher.FilterCountries(req.Country)
		resp.Addresses = TopAddressesFn(limit, minMatch, filtered, multiAddressCompare(compares...))

		// record Prometheus metrics
		if len(resp.Addresses) > 0 {
			matchHist.With("type", "address").Observe(resp.Addresses[0].match)
		} else {
			matchHist.With("type", "address").Observe(0.0)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func searchViaQ(logger log.Logger, searcher *searcher, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name = strings.TrimSpace(name)
		if name == "" {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}
		limit := extractSearchLimit(r)
		minMatch := extractSearchMinMatch(r)

		// Perform multiple searches over the set of SDNs
		resp := buildFullSearchResponse(searcher, buildFilterRequest(r.URL), limit, minMatch, name)

		// record Prometheus metrics
		if len(resp.SDNs) > 0 {
			matchHist.With("type", "q").Observe(resp.SDNs[0].match)
		} else {
			matchHist.With("type", "q").Observe(0.0)
		}

		// Build our big response object
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

// searchGather performs an inmem search with *searcher and mutates *searchResponse by setting a specific field
type searchGather func(searcher *searcher, filters filterRequest, limit int, minMatch float64, name string, resp *searchResponse)

var (
	gatherings = []searchGather{
		// OFAC SDN Search
		func(s *searcher, filters filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			sdns := s.FindSDNsByRemarksID(limit, name)
			if len(sdns) == 0 {
				sdns = s.TopSDNs(limit, minMatch, name, keepSDN(filters))
			}
			resp.SDNs = filterSDNs(sdns, filters)
		},
		// OFAC SDN Alt Names
		func(s *searcher, _ filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			resp.AltNames = s.TopAltNames(limit, minMatch, name)
		},
		// OFAC Addresses
		func(s *searcher, _ filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			resp.Addresses = s.TopAddresses(limit, minMatch, name)
		},
		// OFAC Sectoral Sanctions Identifications
		func(s *searcher, _ filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			resp.SectoralSanctions = s.TopSSIs(limit, minMatch, name)
		},
		// BIS Denied Persons
		func(s *searcher, _ filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			resp.DeniedPersons = s.TopDPs(limit, minMatch, name)
		},
		// BIS Entity List
		func(s *searcher, _ filterRequest, limit int, minMatch float64, name string, resp *searchResponse) {
			resp.BISEntities = s.TopBISEntities(limit, minMatch, name)
		},
	}
)

func buildFullSearchResponse(searcher *searcher, filters filterRequest, limit int, minMatch float64, name string) *searchResponse {
	resp := searchResponse{
		RefreshedAt: searcher.lastRefreshedAt,
	}
	var wg sync.WaitGroup
	wg.Add(len(gatherings))
	for i := range gatherings {
		go func(i int) {
			gatherings[i](searcher, filters, limit, minMatch, name, &resp)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return &resp
}

func searchViaAddressAndName(logger log.Logger, searcher *searcher, name string, req addressSearchRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name = strings.TrimSpace(name)
		if name == "" || req.empty() {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		limit, minMatch := extractSearchLimit(r), extractSearchMinMatch(r)

		resp := &searchResponse{
			RefreshedAt: searcher.lastRefreshedAt,
		}

		resp.SDNs = searcher.TopSDNs(limit, minMatch, name, keepSDN(buildFilterRequest(r.URL)))

		compares := buildAddressCompares(req)
		filtered := searcher.FilterCountries(req.Country)
		resp.Addresses = TopAddressesFn(limit, minMatch, filtered, multiAddressCompare(compares...))

		// record Prometheus metrics
		if len(resp.SDNs) > 0 && len(resp.Addresses) > 0 {
			matchHist.With("type", "addressname").Observe(math.Max(resp.SDNs[0].match, resp.Addresses[0].match))
		} else {
			matchHist.With("type", "addressname").Observe(0.0)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func searchByRemarksID(logger log.Logger, searcher *searcher, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if id == "" {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		limit := extractSearchLimit(r)

		sdns := searcher.FindSDNsByRemarksID(limit, id)
		sdns = filterSDNs(sdns, buildFilterRequest(r.URL))

		// record Prometheus metrics
		if len(sdns) > 0 {
			matchHist.With("type", "remarksID").Observe(sdns[0].match)
		} else {
			matchHist.With("type", "remarksID").Observe(0.0)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&searchResponse{
			SDNs:        sdns,
			RefreshedAt: searcher.lastRefreshedAt,
		})
	}
}

func searchByName(logger log.Logger, searcher *searcher, nameSlug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameSlug = strings.TrimSpace(nameSlug)
		if nameSlug == "" {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		limit := extractSearchLimit(r)
		minMatch := extractSearchMinMatch(r)

		// Grab the SDN's and then filter any out based on query params
		sdns := searcher.TopSDNs(limit, minMatch, nameSlug, keepSDN(buildFilterRequest(r.URL)))

		// record Prometheus metrics
		if len(sdns) > 0 {
			matchHist.With("type", "name").Observe(sdns[0].match)
		} else {
			matchHist.With("type", "name").Observe(0.0)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&searchResponse{
			// OFAC
			SDNs:              sdns,
			AltNames:          searcher.TopAltNames(limit, minMatch, nameSlug),
			SectoralSanctions: searcher.TopSSIs(limit, minMatch, nameSlug),
			// BIS
			DeniedPersons: searcher.TopDPs(limit, minMatch, nameSlug),
			BISEntities:   searcher.TopBISEntities(limit, minMatch, nameSlug),
			// Metadata
			RefreshedAt: searcher.lastRefreshedAt,
		})
	}
}

func searchByAltName(logger log.Logger, searcher *searcher, altSlug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		altSlug = strings.TrimSpace(altSlug)
		if altSlug == "" {
			moovhttp.Problem(w, errNoSearchParams)
			return
		}

		limit := extractSearchLimit(r)
		minMatch := extractSearchMinMatch(r)
		alts := searcher.TopAltNames(limit, minMatch, altSlug)

		// record Prometheus metrics
		if len(alts) > 0 {
			matchHist.With("type", "altName").Observe(alts[0].match)
		} else {
			matchHist.With("type", "altName").Observe(0.0)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&searchResponse{
			AltNames:    alts,
			RefreshedAt: searcher.lastRefreshedAt,
		})
	}
}
