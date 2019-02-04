// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/moov-io/ofac"
	"github.com/moov-io/ofac/pkg/strcmp"

	"github.com/go-kit/kit/log"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")

	nameSimilarity    float64 = 0.90
	altSimilarity     float64 = 0.90
	addressSimilarity float64 = 0.90

	softResultsLimit, hardResultsLimit = 10, 100
)

func init() {
	if v := os.Getenv("NAME_SIMILARITY"); v != "" {
		f, _ := strconv.ParseFloat(v, 64)
		if f > 0 {
			nameSimilarity = f
		}
	}
	if v := os.Getenv("ALT_SIMILARITY"); v != "" {
		f, _ := strconv.ParseFloat(v, 64)
		if f > 0 {
			altSimilarity = f
		}
	}
	if v := os.Getenv("ADDRESS_SIMILARITY"); v != "" {
		f, _ := strconv.ParseFloat(v, 64)
		if f > 0 {
			addressSimilarity = f
		}
	}
}

type searcher struct {
	SDNs         []*SDN
	Addresses    []*Address
	Alts         []*Alt
	sync.RWMutex // protects all above fields

	logger log.Logger
}

func (s *searcher) FindAddresses(limit int, f func(*Address) bool) []*ofac.Address {
	s.RLock()
	defer s.RUnlock()

	var out []*ofac.Address
	for i := range s.Addresses {
		// Break if at results limit
		if len(out) > limit {
			break
		}
		// Check filter func
		if f(s.Addresses[i]) {
			out = append(out, s.Addresses[i].Address)
		}
	}
	return out
}

func (s *searcher) FindAlts(limit int, f func(alt *Alt) bool) []*ofac.AlternateIdentity {
	s.RLock()
	defer s.RUnlock()

	var out []*ofac.AlternateIdentity
	for i := range s.Alts {
		if len(out) > limit {
			break
		}
		if f(s.Alts[i]) {
			out = append(out, s.Alts[i].AlternateIdentity)
		}
	}
	return out
}

func (s *searcher) FindSDNs(limit int, f func(*SDN) bool) []*ofac.SDN {
	s.RLock()
	defer s.RUnlock()

	var out []*ofac.SDN
	for i := range s.SDNs {
		if len(out) > limit {
			break
		}
		if f(s.SDNs[i]) {
			out = append(out, s.SDNs[i].SDN)
		}
	}
	return out
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

func extractSearchLimit(r *http.Request) int {
	limit := softResultsLimit
	if v := r.URL.Query().Get("limit"); v != "" {
		n, _ := strconv.Atoi(v)
		if n > 0 {
			limit = n
		}
	}
	if limit > hardResultsLimit {
		limit = hardResultsLimit
	}
	return limit
}

type searchResponse struct {
	SDNs      []*ofac.SDN               `json:"SDNs"`
	AltNames  []*ofac.AlternateIdentity `json:"altNames"`
	Addresses []*ofac.Address           `json:"addresses"`
}

// addressMatches returns a bool which represents if addressParts matches a given Address.
//
// An Address contains precomputed data to speed up searches and addressParts is split along
// word boundries.
func addressMatches(addressParts []string, inc *Address) bool {
	// Count matches for collection if over threshold
	matches := 0
	for k := range inc.address {
		for j := range addressParts {
			if strcmp.Levenshtein(inc.address[k], addressParts[j]) > addressSimilarity {
				matches++
			}
		}
	}
	// If over 25% of words from query match (via strings.Contains not full string equality) save as an address.
	// This is arbitrary, but given the following examples only one partial word match is required:
	//  123 Scott Ave
	//  1600 N Penn St
	if (float64(matches) / float64(len(inc.address))) >= 0.25 {
		return true
	}
	return false
}

// nameMatches returns a bool representing if nameParts (name split on word boundries) matches a given SDN.
func nameMatches(nameParts []string, inc *SDN) bool {
	for k := range inc.name {
		for j := range nameParts {
			if strcmp.Levenshtein(inc.name[k], nameParts[j]) > nameSimilarity {
				return true
			}
		}
	}
	return false
}

// altMatches returns a bool representing if altParts (alt name split on word boundries) matches a given ofac.AlternateIdentity.
func altMatches(altParts []string, inc *Alt) bool {
	for k := range inc.name {
		for j := range altParts {
			if strcmp.Levenshtein(inc.name[k], altParts[j]) > altSimilarity {
				return true
			}
		}
	}
	return false
}
