// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/xrash/smetrics"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")

	softResultsLimit, hardResultsLimit = 10, 100
)

// searcher holds precomputed data for each object available to search against.
// This data comes from various US Federal agencies, such as: OFAC and BIS
type searcher struct {
	SDNs            []*SDN
	Addresses       []*Address
	Alts            []*Alt
	DPs             []*DP
	lastRefreshedAt time.Time
	sync.RWMutex    // protects all above fields

	logger log.Logger
}

func (s *searcher) FindAddresses(limit int, id string) []*ofac.Address {
	s.RLock()
	defer s.RUnlock()

	var out []*ofac.Address
	for i := range s.Addresses {
		if len(out) > limit {
			break
		}
		if s.Addresses[i].Address.EntityID == id {
			out = append(out, s.Addresses[i].Address)
		}
	}
	return out
}

func (s *searcher) TopAddresses(limit int, reqAddress string) []Address {
	return s.TopAddressesFn(limit, topAddressesAddress(reqAddress))
}

var (
	// topAddressesAddress is a compare method for TopAddressesFn to extract and rank .Address
	topAddressesAddress = func(needleAddr string) func(*Address) *item {
		return func(add *Address) *item {
			return &item{
				value:  add,
				weight: jaroWinkler(add.address, precompute(needleAddr)),
			}
		}
	}

	// topAddressesCityState is a compare method for TopAddressesFn to extract and rank
	// .City, .State, .Providence, and .Zip to return the average match between non-empty
	// search criteria.
	topAddressesCityState = func(needleCityState string) func(*Address) *item {
		return func(add *Address) *item {
			return &item{
				value:  add,
				weight: jaroWinkler(add.citystate, precompute(needleCityState)),
			}
		}
	}

	// topAddressesCountry is a compare method for TopAddressesFn to extract and rank .Country
	topAddressesCountry = func(needleCountry string) func(*Address) *item {
		return func(add *Address) *item {
			return &item{
				value:  add,
				weight: jaroWinkler(add.country, precompute(needleCountry)),
			}
		}
	}

	// multiAddressCompare is a compare method for taking N higher-order compare methods
	// and returning an average weight after computing them all.
	multiAddressCompare = func(cmps ...func(*Address) *item) func(*Address) *item {
		return func(add *Address) *item {
			weight := 0.00
			for i := range cmps {
				weight += cmps[i](add).weight
			}
			return &item{
				value:  add,
				weight: weight / float64(len(cmps)),
			}
		}
	}
)

// TopAddressesFn performs an Address search over an arbitrary member of Address. It's mainly used to rank
// and search over .Country, .CityStateProvincePostalCode.
//
// compare takes an Address (from s.Addresses) and is expected to extract some property to be compared
// against a captured parameter (in a closure calling compare) to return an *item for final sorting.
// See searchByAddress in search_handlers.go for an example
func (s *searcher) TopAddressesFn(limit int, compare func(*Address) *item) []Address {
	s.RLock()
	defer s.RUnlock()

	if len(s.Addresses) == 0 {
		return nil
	}
	xs := newLargest(limit)

	for i := range s.Addresses {
		xs.add(compare(s.Addresses[i]))
	}
	return largestToAddresses(xs)
}

func largestToAddresses(xs *largest) []Address {
	out := make([]Address, 0)
	for i := range xs.items {
		if v := xs.items[i]; v != nil {
			aa, ok := v.value.(*Address)
			if !ok {
				continue
			}
			address := *aa
			address.match = v.weight
			out = append(out, address)
		}
	}
	return out
}

func (s *searcher) FindAlts(limit int, id string) []*ofac.AlternateIdentity {
	s.RLock()
	defer s.RUnlock()

	var out []*ofac.AlternateIdentity
	for i := range s.Alts {
		if len(out) > limit {
			break
		}
		if s.Alts[i].AlternateIdentity.EntityID == id {
			out = append(out, s.Alts[i].AlternateIdentity)
		}
	}
	return out
}

func (s *searcher) TopAltNames(limit int, alt string) []Alt {
	alt = precompute(alt)

	s.RLock()
	defer s.RUnlock()

	if len(s.Alts) == 0 {
		return nil
	}
	xs := newLargest(limit)

	for i := range s.Alts {
		xs.add(&item{
			value:  s.Alts[i],
			weight: jaroWinkler(s.Alts[i].name, alt),
		})
	}

	out := make([]Alt, 0)
	for i := range xs.items {
		if v := xs.items[i]; v != nil {
			aa, ok := v.value.(*Alt)
			if !ok {
				continue
			}
			alt := *aa
			alt.match = v.weight
			out = append(out, alt)
		}
	}
	return out
}

func (s *searcher) FindSDN(entityID string) *ofac.SDN {
	s.RLock()
	defer s.RUnlock()

	for i := range s.SDNs {
		if s.SDNs[i].EntityID == entityID {
			return s.SDNs[i].SDN
		}
	}
	return nil
}

// FindSDNsByRemarksID looks for SDN's whose remarks property contains an ID matching
// what is provided to this function. It's typically used with values assigned by a local
// government. (National ID, Drivers License, etc)
func (s *searcher) FindSDNsByRemarksID(limit int, id string) []SDN {
	var out []SDN
	for i := range s.SDNs {
		// If our remarks ID contains a space then just see if the query ID matches any part.
		// Otherwise ID searches need to be exact matches.
		if strings.Contains(s.SDNs[i].id, " ") && strings.Contains(s.SDNs[i].id, id) {
			sdn := *s.SDNs[i]
			sdn.match = 1.0
			out = append(out, sdn)
		} else {
			if s.SDNs[i].id == id {
				sdn := *s.SDNs[i]
				sdn.match = 1.0
				out = append(out, sdn)
			}
		}
		// quit if we're at our max result size
		if len(out) >= limit {
			return out
		}
	}
	return out
}

func (s *searcher) TopSDNs(limit int, name string) []SDN {
	name = precompute(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.SDNs) == 0 {
		return nil
	}
	xs := newLargest(limit)

	for i := range s.SDNs {
		xs.add(&item{
			value:  s.SDNs[i],
			weight: jaroWinkler(s.SDNs[i].name, name),
		})
	}

	out := make([]SDN, 0)
	for i := range xs.items {
		if v := xs.items[i]; v != nil {
			ss, ok := v.value.(*SDN)
			if !ok {
				continue
			}
			sdn := *ss // deref for a copy
			sdn.match = v.weight
			out = append(out, sdn)
		}
	}
	return out
}

func (s *searcher) TopDPs(limit int, name string) []DP {
	name = precompute(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.DPs) == 0 {
		return nil
	}
	xs := newLargest(limit)

	for _, dp := range s.DPs {
		xs.add(&item{
			value:  dp,
			weight: jaroWinkler(dp.name, name),
		})
	}

	out := make([]DP, 0)
	for _, thisItem := range xs.items {
		if v := thisItem; v != nil {
			ss, ok := v.value.(*DP)
			if !ok {
				continue
			}
			dp := *ss
			dp.match = v.weight
			out = append(out, dp)
		}
	}
	return out
}

// SDN is ofac.SDN wrapped with precomputed search metadata
type SDN struct {
	*ofac.SDN

	// match holds the match ratio for an SDN in search results
	match float64

	// name is precomputed for speed
	name string

	// id is the parseed ID value from an SDN's remarks field. Often this
	// is a National ID, Drivers License, or similar government value
	// ueed to uniquely identify an entiy.
	//
	// Typically the form of this is 'No. NNNNN' where NNNNN is alphanumeric.
	id string
}

// MarshalJSON is a custom method for marshaling a SDN search result
func (s SDN) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*ofac.SDN
		Match float64 `json:"match"`
	}{
		s.SDN,
		s.match,
	})
}

func precomputeSDNs(sdns []*ofac.SDN) []*SDN {
	out := make([]*SDN, len(sdns))
	for i := range sdns {
		out[i] = &SDN{
			SDN:  sdns[i],
			name: precompute(reorderSDNName(sdns[i].SDNName, sdns[i].SDNType)),
			id:   extractIDFromRemark(strings.TrimSpace(sdns[i].Remarks)),
		}
	}
	return out
}

var (
	surnamePrecedes = regexp.MustCompile(`(,?[\s?a-zA-Z\.]{1,})$`)
)

// reorderSDNName will take a given SDN name and if it matches a specific pattern where
// the first name is placed after the last name (surname) to return a string where the first name
// preceedes the last.
//
// Example:
// SDN EntityID: 19147 has 'FELIX B. MADURO S.A.'
// SDN EntityID: 22790 has 'MADURO MOROS, Nicolas'
func reorderSDNName(name string, tpe string) string {
	if !strings.EqualFold(tpe, "individual") {
		return name // only reorder individual names
	}
	v := surnamePrecedes.FindString(name)
	if v == "" {
		return name // no match on 'Doe, John'
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", strings.TrimPrefix(v, ","), strings.TrimSuffix(name, v)))
}

// Address is ofac.Address wrapped with precomputed search metadata
type Address struct {
	Address *ofac.Address

	match float64 // match %

	// precomputed fields for speed
	address, citystate, country string
}

// MarshalJSON is a custom method for marshaling a SDN Address search result
func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*ofac.Address
		Match float64 `json:"match"`
	}{
		a.Address,
		a.match,
	})
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

	match float64 // match %

	// name is precomputed for speed
	name string
}

// MarshalJSON is a custom method for marshaling a SDN Alternate Identity search result
func (a Alt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*ofac.AlternateIdentity
		Match float64 `json:"match"`
	}{
		a.AlternateIdentity,
		a.match,
	})
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

// DP is a BIS Denied Person wrapped with precomputed search metadata
type DP struct {
	DeniedPerson *ofac.DPL
	match        float64
	name         string
}

// MarshalJSON is a custom method for marshaling a BIS Denied Person (DP)
func (d DP) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*ofac.DPL
		Match float64 `json:"match"`
	}{
		d.DeniedPerson,
		d.match,
	})
}

func precomputeDPs(persons []*ofac.DPL) []*DP {
	out := make([]*DP, len(persons))
	for i := range persons {
		out[i] = &DP{
			DeniedPerson: persons[i],
			name:         precompute(persons[i].Name),
		}
	}
	return out
}

var (
	punctuationReplacer = strings.NewReplacer(".", "", ",", "", "-", "", "  ", " ")
)

// precompute will lowercase each substring and remove punctuation
//
// This function is called on every record from the flat files and all
// search requests (i.e. HTTP and searcher.TopNNNs methods).
// See: https://godoc.org/golang.org/x/text/unicode/norm#Form
// See: https://withblue.ink/2019/03/11/why-you-need-to-normalize-unicode-strings.html
func precompute(s string) string {
	trimmed := strings.ToLower(punctuationReplacer.Replace(s))

	// UTF-8 normalization
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC) // Mn: nonspacing marks
	result, _, _ := transform.String(t, trimmed)
	return result
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

// jaroWrinkler runs the similarly named algorithm over the two input strings and averages their match percentages
// according to the second string (assumed to be the user's query)
//
// For more details see https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func jaroWinkler(s1, s2 string) float64 {
	maxMatch := func(word string, parts []string) float64 {
		if len(parts) == 0 {
			return 0.0
		}
		max := smetrics.JaroWinkler(word, parts[0], 0.7, 4)
		for i := 1; i < len(parts); i++ {
			if score := smetrics.JaroWinkler(word, parts[i], 0.7, 4); score > max {
				max = score
			}
		}
		return max
	}

	s1Parts, s2Parts := strings.Fields(s1), strings.Fields(s2)
	if len(s1Parts) == 0 || len(s2Parts) == 0 {
		return 0.0 // avoid returning NaN later on
	}

	var scores []float64
	for i := range s1Parts {
		scores = append(scores, maxMatch(s1Parts[i], s2Parts))
	}

	// average the highest N scores where N is the words in our query (s2).
	sort.Float64s(scores)
	if len(s1Parts) > len(s2Parts) {
		scores = scores[len(s2Parts)-1:]
	}

	var sum float64
	for i := range scores {
		sum += scores[i]
	}
	return sum / float64(len(scores))
}

// extractIDFromRemark attempts to parse out a National ID or similar governmental ID value
// from an SDN's remarks property.
//
// Typically the form of this is 'No. NNNNN' where NNNNN is alphanumeric.
func extractIDFromRemark(remarks string) string {
	if remarks == "" {
		return ""
	}

	var out bytes.Buffer
	parts := strings.Fields(remarks)
	for i := range parts {
		if parts[i] == "No." {
			trimmed := strings.TrimSuffix(strings.TrimSuffix(parts[i+1], "."), ";")

			// Always take the next part
			if strings.HasSuffix(parts[i+1], ".") || strings.HasSuffix(parts[i+1], ";") {
				return trimmed
			} else {
				out.WriteString(trimmed)
			}
			// possibly take additional parts
			for j := i + 2; j < len(parts); j++ {
				if strings.HasPrefix(parts[j], "(") {
					return out.String()
				}
				if _, err := strconv.ParseInt(parts[j], 10, 32); err == nil {
					out.WriteString(" " + parts[j])
				}
			}
		}
	}
	return out.String()
}
