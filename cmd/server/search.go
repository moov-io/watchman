// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/moov-io/watchman/pkg/csl_eu"
	"github.com/moov-io/watchman/pkg/csl_uk"
	"github.com/moov-io/watchman/pkg/csl_us"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

	"go4.org/syncutil"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")

	softResultsLimit, hardResultsLimit = 10, 100
)

// searcher holds prepare.LowerAndRemovePunctuationd data for each object available to search against.
// This data comes from various US and EU Federal agencies
type searcher struct {
	// OFAC
	SDNs        []*SDN
	Addresses   []*Address
	Alts        []*Alt
	SDNComments []*ofac.SDNComments

	// BIS
	DPs []*DP

	// TODO: this could be refactored into sub structs that have us/eu (and eventually others)

	// US Consolidated Screening List
	BISEntities      []*Result[csl_us.EL]
	MilitaryEndUsers []*Result[csl_us.MEU]
	SSIs             []*Result[csl_us.SSI]
	UVLs             []*Result[csl_us.UVL]
	ISNs             []*Result[csl_us.ISN]
	FSEs             []*Result[csl_us.FSE]
	PLCs             []*Result[csl_us.PLC]
	CAPs             []*Result[csl_us.CAP]
	DTCs             []*Result[csl_us.DTC]
	CMICs            []*Result[csl_us.CMIC]
	NS_MBSs          []*Result[csl_us.NS_MBS]

	// EU Consolidated List of Sactions
	EUCSL []*Result[csl_eu.CSLRecord]

	// UK Consolidated List of Sactions - OFSI
	UKCSL []*Result[csl_uk.CSLRecord]

	// UK Sanctions List
	UKSanctionsList []*Result[csl_uk.SanctionsListRecord]

	// metadata
	lastRefreshedAt time.Time
	sync.RWMutex    // protects all above fields
	*syncutil.Gate  // limits concurrent processing

	pipe *prepare.Pipeliner

	logger log.Logger
}

func newSearcher(logger log.Logger, pipeline *prepare.Pipeliner, workers int) *searcher {
	logger.Logf("allowing only %d workers for search", workers)
	return &searcher{
		logger: logger.With(log.Fields{
			"component": log.String("pipeline"),
		}),
		pipe: pipeline,
		Gate: syncutil.NewGate(workers),
	}
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

func (s *searcher) TopAddresses(limit int, minMatch float64, reqAddress string) []Address {
	s.RLock()
	defer s.RUnlock()

	return TopAddressesFn(limit, minMatch, s.Addresses, topAddressesAddress(reqAddress))
}

var (
	// topAddressesAddress is a compare method for TopAddressesFn to extract and rank .Address
	topAddressesAddress = func(needleAddr string) func(*Address) *item {
		return func(add *Address) *item {
			return &item{
				value:  add,
				weight: stringscore.JaroWinkler(add.address, prepare.LowerAndRemovePunctuation(needleAddr)),
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
				weight: stringscore.JaroWinkler(add.citystate, prepare.LowerAndRemovePunctuation(needleCityState)),
			}
		}
	}

	// topAddressesCountry is a compare method for TopAddressesFn to extract and rank .Country
	topAddressesCountry = func(needleCountry string) func(*Address) *item {
		return func(add *Address) *item {
			return &item{
				value:  add,
				weight: stringscore.JaroWinkler(add.country, prepare.LowerAndRemovePunctuation(needleCountry)),
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

// FilterCountries returns Addresses that match a given country name.
//
// If name is blank all Addresses are returned.
//
// This filtering ignore case differences, but does require the name matches
// to the underlying data.
func (s *searcher) FilterCountries(name string) []*Address {
	s.RLock()
	defer s.RUnlock()

	if len(s.Addresses) == 0 {
		return nil
	}

	if name == "" {
		out := make([]*Address, len(s.Addresses))
		copy(out, s.Addresses)
		return out
	}
	var out []*Address
	for i := range s.Addresses {
		if strings.EqualFold(s.Addresses[i].country, name) {
			out = append(out, s.Addresses[i])
		}
	}
	return out
}

// TopAddressesFn performs a ranked search over an arbitrary set of Address fields.
//
// compare takes an Address (from s.Addresses) and is expected to extract some property to be compared
// against a captured parameter (in a closure calling compare) to return an *item for final sorting.
// See searchByAddress in search_handlers.go for an example
func TopAddressesFn(limit int, minMatch float64, addresses []*Address, compare func(*Address) *item) []Address {
	if len(addresses) == 0 {
		return nil
	}
	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(addresses))

	for i := range addresses {
		go func(i int) {
			defer wg.Done()
			xs.add(compare(addresses[i]))
		}(i)
	}

	wg.Wait()

	return largestToAddresses(xs)
}

func largestToAddresses(xs *largest) []Address {
	items := xs.getItems()
	out := make([]Address, 0, xs.capacity)
	for _, item := range items {
		if item == nil {
			continue
		}

		aa, ok := item.value.(*Address)
		if !ok {
			continue
		}
		address := *aa
		address.match = item.weight
		out = append(out, address)
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

func (s *searcher) TopAltNames(limit int, minMatch float64, alt string) []Alt {
	alt = prepare.LowerAndRemovePunctuation(alt)
	altTokens := strings.Fields(alt)

	s.RLock()
	defer s.RUnlock()

	if len(s.Alts) == 0 {
		return nil
	}
	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(s.Alts))

	for i := range s.Alts {
		s.Gate.Start()
		go func(i int) {
			defer wg.Done()
			defer s.Gate.Done()
			xs.add(&item{
				matched: s.Alts[i].name,
				value:   s.Alts[i],
				weight:  stringscore.BestPairsJaroWinkler(altTokens, s.Alts[i].name),
			})
		}(i)
	}
	wg.Wait()

	items := xs.getItems()
	out := make([]Alt, 0, limit)
	for _, item := range items {
		if item == nil {
			continue
		}

		aa, ok := item.value.(*Alt)
		if !ok {
			continue
		}
		alt := *aa
		alt.match = item.weight
		alt.matchedName = item.matched
		out = append(out, alt)
	}
	return out
}

func (s *searcher) FindSDN(entityID string) *ofac.SDN {
	if sdn := s.debugSDN(entityID); sdn != nil {
		return sdn.SDN
	}
	return nil
}

func (s *searcher) debugSDN(entityID string) *SDN {
	s.RLock()
	defer s.RUnlock()

	return s.findSDNWithoutLock(entityID)
}

func (s *searcher) findSDNWithoutLock(entityID string) *SDN {
	for i := range s.SDNs {
		if s.SDNs[i].EntityID == entityID {
			return s.SDNs[i]
		}
	}
	return nil
}

// FindSDNsByRemarksID looks for SDN's whose remarks property contains an ID matching
// what is provided to this function. It's typically used with values assigned by a local
// government. (National ID, Drivers License, etc)
func (s *searcher) FindSDNsByRemarksID(limit int, id string) []*SDN {
	if id == "" {
		return nil
	}

	var out []*SDN
	for i := range s.SDNs {
		// If the SDN's remarks ID contains a space then we need to ensure "all the numeric
		// parts have to exactly match" between our query and the parsed ID.
		if strings.Contains(s.SDNs[i].id, " ") {
			qParts := strings.Fields(id)
			sdnParts := strings.Fields(s.SDNs[i].id)

			matched, expected := 0, 0
			for j := range sdnParts {
				if n, _ := strconv.ParseInt(sdnParts[j], 10, 64); n > 0 {
					// This part of the SDN's remarks is a number so it must exactly
					// match to a query's part
					expected += 1

					for k := range qParts {
						if sdnParts[j] == qParts[k] {
							matched += 1
						}
					}
				}
			}

			// If all the numeric parts match between query and SDN return the match
			if matched == expected {
				sdn := *s.SDNs[i]
				sdn.match = 1.0
				out = append(out, &sdn)
			}
		} else {
			// The query and remarks ID must exactly match
			if s.SDNs[i].id == id {
				sdn := *s.SDNs[i]
				sdn.match = 1.0
				out = append(out, &sdn)
			}
		}

		// quit if we're at our max result size
		if len(out) >= limit {
			return out
		}
	}
	return out
}

func (s *searcher) TopSDNs(limit int, minMatch float64, name string, keepSDN func(*SDN) bool) []*SDN {
	name = prepare.LowerAndRemovePunctuation(name)
	nameTokens := strings.Fields(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.SDNs) == 0 {
		return nil
	}
	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(s.SDNs))

	for i := range s.SDNs {
		if !keepSDN(s.SDNs[i]) {
			wg.Done()
			continue
		}
		s.Gate.Start()
		go func(i int) {
			defer wg.Done()
			defer s.Gate.Done()
			xs.add(&item{
				matched: s.SDNs[i].name,
				value:   s.SDNs[i],
				weight:  stringscore.BestPairsJaroWinkler(nameTokens, s.SDNs[i].name),
			})
		}(i)
	}
	wg.Wait()

	items := xs.getItems()
	out := make([]*SDN, 0, limit)
	for _, item := range items {
		if item == nil {
			continue
		}

		ss, ok := item.value.(*SDN)
		if !ok {
			continue
		}
		sdn := *ss
		sdn.match = item.weight
		sdn.matchedName = item.matched
		out = append(out, &sdn)
	}
	return out
}

func (s *searcher) TopDPs(limit int, minMatch float64, name string) []DP {
	name = prepare.LowerAndRemovePunctuation(name)
	nameTokens := strings.Fields(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.DPs) == 0 {
		return nil
	}
	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(s.DPs))

	for i := range s.DPs {
		s.Gate.Start()
		go func(i int) {
			defer wg.Done()
			defer s.Gate.Done()
			xs.add(&item{
				matched: s.DPs[i].name,
				value:   s.DPs[i],
				weight:  stringscore.BestPairsJaroWinkler(nameTokens, s.DPs[i].name),
			})
		}(i)
	}
	wg.Wait()

	items := xs.getItems()
	out := make([]DP, 0, limit)
	for _, item := range items {
		if item == nil {
			continue
		}

		ss, ok := item.value.(*DP)
		if !ok {
			continue
		}
		dp := *ss
		dp.match = item.weight
		dp.matchedName = item.matched
		out = append(out, dp)
	}
	return out
}

// SDN is ofac.SDN wrapped with prepare.LowerAndRemovePunctuationd search metadata
type SDN struct {
	*ofac.SDN

	// match holds the match ratio for an SDN in search results
	match float64

	// matchedName holds the highest scoring term from the search query
	matchedName string

	// name is prepare.LowerAndRemovePunctuationd for speed
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
		Match       float64 `json:"match"`
		MatchedName string  `json:"matchedName"`
	}{
		s.SDN,
		s.match,
		s.matchedName,
	})
}

func findAddresses(entityID string, addrs []*ofac.Address) []*ofac.Address {
	var out []*ofac.Address
	for i := range addrs {
		if entityID == addrs[i].EntityID {
			out = append(out, addrs[i])
		}
	}
	return out
}

func precomputeSDNs(sdns []*ofac.SDN, addrs []*ofac.Address, pipe *prepare.Pipeliner) []*SDN {
	out := make([]*SDN, len(sdns))
	for i := range sdns {
		nn := prepare.SdnName(sdns[i], findAddresses(sdns[i].EntityID, addrs))

		if err := pipe.Do(nn); err != nil {
			continue
		}

		out[i] = &SDN{
			SDN:  sdns[i],
			name: nn.Processed,
			id:   extractIDFromRemark(strings.TrimSpace(sdns[i].Remarks)),
		}
	}
	return out
}

// Address is ofac.Address wrapped with prepare.LowerAndRemovePunctuationd search metadata
type Address struct {
	Address *ofac.Address

	match float64 // match %

	// prepare.LowerAndRemovePunctuationd fields for speed
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
			address:   prepare.LowerAndRemovePunctuation(adds[i].Address),
			citystate: prepare.LowerAndRemovePunctuation(adds[i].CityStateProvincePostalCode),
			country:   prepare.LowerAndRemovePunctuation(adds[i].Country),
		}
	}
	return out
}

// Alt is an ofac.AlternateIdentity wrapped with prepare.LowerAndRemovePunctuationd search metadata
type Alt struct {
	AlternateIdentity *ofac.AlternateIdentity

	// match holds the match ratio for an Alt in search results
	match float64

	// matchedName holds the highest scoring term from the search query
	matchedName string

	// name is prepare.LowerAndRemovePunctuationd for speed
	name string
}

// MarshalJSON is a custom method for marshaling a SDN Alternate Identity search result
func (a Alt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*ofac.AlternateIdentity
		Match       float64 `json:"match"`
		MatchedName string  `json:"matchedName"`
	}{
		a.AlternateIdentity,
		a.match,
		a.matchedName,
	})
}

func precomputeAlts(alts []*ofac.AlternateIdentity, pipe *prepare.Pipeliner) []*Alt {
	out := make([]*Alt, len(alts))
	for i := range alts {
		an := prepare.AltName(alts[i])

		if err := pipe.Do(an); err != nil {
			continue
		}

		out[i] = &Alt{
			AlternateIdentity: alts[i],
			name:              an.Processed,
		}
	}
	return out
}

// DP is a BIS Denied Person wrapped with prepare.LowerAndRemovePunctuationd search metadata
type DP struct {
	DeniedPerson *dpl.DPL
	match        float64
	matchedName  string
	name         string
}

// MarshalJSON is a custom method for marshaling a BIS Denied Person (DP)
func (d DP) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*dpl.DPL
		Match       float64 `json:"match"`
		MatchedName string  `json:"matchedName"`
	}{
		d.DeniedPerson,
		d.match,
		d.matchedName,
	})
}

func precomputeDPs(persons []*dpl.DPL, pipe *prepare.Pipeliner) []*DP {
	out := make([]*DP, len(persons))
	for i := range persons {
		nn := prepare.DPName(persons[i])
		if err := pipe.Do(nn); err != nil {
			continue
		}
		out[i] = &DP{
			DeniedPerson: persons[i],
			name:         nn.Processed,
		}
	}
	return out
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
