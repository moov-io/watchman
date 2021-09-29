// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/strx"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

	"github.com/go-kit/kit/log"
	"github.com/xrash/smetrics"
	"go4.org/syncutil"
)

var (
	errNoSearchParams = errors.New("missing search parameter(s)")

	softResultsLimit, hardResultsLimit = 10, 100
)

// searcher holds precomputed data for each object available to search against.
// This data comes from various US and EU Federal agencies
type searcher struct {
	// OFAC
	SDNs      []*SDN
	Addresses []*Address
	Alts      []*Alt
	SSIs      []*SSI

	// BIS
	DPs         []*DP
	BISEntities []*BISEntity

	// metadata
	lastRefreshedAt time.Time
	sync.RWMutex    // protects all above fields
	*syncutil.Gate  // limits concurrent processing

	pipe *pipeliner

	logger log.Logger
}

func newSearcher(logger log.Logger, pipeline *pipeliner, workers int) *searcher {
	logger.Log("search", fmt.Sprintf("allowing only %d workers", workers))
	return &searcher{
		logger: logger,
		pipe:   pipeline,
		Gate:   syncutil.NewGate(workers),
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

func (s *searcher) TopAltNames(limit int, minMatch float64, alt string) []Alt {
	alt = precompute(alt)

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
				value:  s.Alts[i],
				weight: jaroWinkler(s.Alts[i].name, alt),
			})
		}(i)
	}
	wg.Wait()

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
	if sdn := s.debugSDN(entityID); sdn != nil {
		return sdn.SDN
	}
	return nil
}

func (s *searcher) debugSDN(entityID string) *SDN {
	s.RLock()
	defer s.RUnlock()

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
	name = precompute(name)

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
				value:  s.SDNs[i],
				weight: jaroWinkler(s.SDNs[i].name, name),
			})
		}(i)
	}
	wg.Wait()

	out := make([]*SDN, 0)
	for i := range xs.items {
		if v := xs.items[i]; v != nil {
			ss, ok := v.value.(*SDN)
			if !ok {
				continue
			}
			sdn := *ss // deref for a copy
			sdn.match = v.weight
			out = append(out, &sdn)
		}
	}
	return out
}

func (s *searcher) TopDPs(limit int, minMatch float64, name string) []DP {
	name = precompute(name)

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
				value:  s.DPs[i],
				weight: jaroWinkler(s.DPs[i].name, name),
			})
		}(i)
	}
	wg.Wait()

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

// TopSSIs searches Sectoral Sanctions records by Name and Alias
func (s *searcher) TopSSIs(limit int, minMatch float64, name string) []SSI {
	name = precompute(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.SSIs) == 0 {
		return nil
	}
	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(s.SSIs))

	for i := range s.SSIs {
		s.Gate.Start()
		go func(i int) {
			defer wg.Done()
			defer s.Gate.Done()
			it := &item{
				value:  s.SSIs[i],
				weight: jaroWinkler(s.SSIs[i].name, name),
			}
			for _, alt := range s.SSIs[i].SectoralSanction.AlternateNames {
				if alt == "" {
					continue
				}
				currWeight := jaroWinkler(alt, name)
				if currWeight > it.weight {
					it.weight = currWeight
				}
			}
			xs.add(it)
		}(i)
	}
	wg.Wait()

	out := make([]SSI, 0)
	for _, thisItem := range xs.items {
		if v := thisItem; v != nil {
			ss, ok := v.value.(*SSI)
			if !ok {
				continue
			}
			ssi := *ss
			ssi.match = v.weight
			out = append(out, ssi)
		}
	}
	return out
}

// TopBISEntities searches BIS Entity List records by name and alias
func (s *searcher) TopBISEntities(limit int, minMatch float64, name string) []BISEntity {
	name = precompute(name)

	s.RLock()
	defer s.RUnlock()

	if len(s.BISEntities) == 0 {
		return nil
	}

	xs := newLargest(limit, minMatch)

	var wg sync.WaitGroup
	wg.Add(len(s.BISEntities))

	for i := range s.BISEntities {
		s.Gate.Start()
		go func(i int) {
			defer wg.Done()
			defer s.Gate.Done()

			it := &item{
				value:  s.BISEntities[i],
				weight: jaroWinkler(s.BISEntities[i].name, name),
			}
			for _, alt := range s.BISEntities[i].Entity.AlternateNames {
				if alt == "" {
					continue
				}
				currWeight := jaroWinkler(alt, name)
				if currWeight > it.weight {
					it.weight = currWeight
				}
			}

			xs.add(it)
		}(i)
	}
	wg.Wait()

	out := make([]BISEntity, 0)
	for _, thisItem := range xs.items {
		if v := thisItem; v != nil {
			ss, ok := v.value.(*BISEntity)
			if !ok {
				continue
			}
			el := *ss
			el.match = v.weight
			out = append(out, el)
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

func findAddresses(entityID string, addrs []*ofac.Address) []*ofac.Address {
	var out []*ofac.Address
	for i := range addrs {
		if entityID == addrs[i].EntityID {
			out = append(out, addrs[i])
		}
	}
	return out
}

func precomputeSDNs(sdns []*ofac.SDN, addrs []*ofac.Address, pipe *pipeliner) []*SDN {
	out := make([]*SDN, len(sdns))
	for i := range sdns {
		nn := sdnName(sdns[i], findAddresses(sdns[i].EntityID, addrs))

		if err := pipe.Do(nn); err != nil {
			pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining SDN: %v", err))
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

func precomputeAlts(alts []*ofac.AlternateIdentity, pipe *pipeliner) []*Alt {
	out := make([]*Alt, len(alts))
	for i := range alts {
		an := altName(alts[i])

		if err := pipe.Do(an); err != nil {
			pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining SDN: %v", err))
			continue
		}

		out[i] = &Alt{
			AlternateIdentity: alts[i],
			name:              an.Processed,
		}
	}
	return out
}

// DP is a BIS Denied Person wrapped with precomputed search metadata
type DP struct {
	DeniedPerson *dpl.DPL
	match        float64
	name         string
}

// MarshalJSON is a custom method for marshaling a BIS Denied Person (DP)
func (d DP) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*dpl.DPL
		Match float64 `json:"match"`
	}{
		d.DeniedPerson,
		d.match,
	})
}

func precomputeDPs(persons []*dpl.DPL, pipe *pipeliner) []*DP {
	out := make([]*DP, len(persons))
	for i := range persons {
		nn := dpName(persons[i])
		if err := pipe.Do(nn); err != nil {
			pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining DP: %v", err))
			continue
		}
		out[i] = &DP{
			DeniedPerson: persons[i],
			name:         nn.Processed,
		}
	}
	return out
}

type SSI struct {
	SectoralSanction *csl.SSI
	match            float64
	name             string
}

func (s SSI) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*csl.SSI
		Match float64 `json:"match"`
	}{
		s.SectoralSanction,
		s.match,
	})
}

func precomputeSSIs(ssis []*csl.SSI, pipe *pipeliner) []*SSI {
	out := make([]*SSI, len(ssis))
	for i, ssi := range ssis {
		nn := ssiName(ssi)
		if err := pipe.Do(nn); err != nil {
			pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining SSI: %v", err))
			continue
		}

		var altNames []string
		for i := range ssi.AlternateNames {
			altNN := &Name{Processed: ssi.AlternateNames[i]}
			if err := pipe.Do(altNN); err != nil {
				pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining alt: %v", err))
				continue
			}
			altNames = append(altNames, altNN.Processed)
		}
		ssi.AlternateNames = altNames

		out[i] = &SSI{
			SectoralSanction: ssi,
			name:             nn.Processed,
		}
	}
	return out
}

type BISEntity struct {
	Entity *csl.EL
	match  float64
	name   string
}

func (e BISEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*csl.EL
		Match float64 `json:"match"`
	}{
		e.Entity,
		e.match,
	})
}

func precomputeBISEntities(els []*csl.EL, pipe *pipeliner) []*BISEntity {
	out := make([]*BISEntity, len(els))
	for i, el := range els {
		nn := bisEntityName(el)
		if err := pipe.Do(nn); err != nil {
			pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining EL: %v", err))
			continue
		}

		var altNames []string
		for i := range el.AlternateNames {
			altNN := &Name{Processed: el.AlternateNames[i]}
			if err := pipe.Do(altNN); err != nil {
				pipe.logger.Log("pipeline", fmt.Sprintf("problem pipelining alt: %v", err))
				continue
			}
			altNames = append(altNames, altNN.Processed)
		}
		el.AlternateNames = altNames

		out[i] = &BISEntity{
			Entity: el,
			name:   nn.Processed,
		}
	}
	return out
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

func extractSearchMinMatch(r *http.Request) float64 {
	if v := r.URL.Query().Get("minMatch"); v != "" {
		n, _ := strconv.ParseFloat(v, 64)
		return n
	}
	return 0.00
}

var (
	exactMatchFavoritism = readExactMatchFavoritism(os.Getenv("EXACT_MATCH_FAVORITISM"))
)

func readExactMatchFavoritism(input string) float64 {
	weight, _ := strconv.ParseFloat(strx.Or(input, "0.0"), 32)
	return weight
}

// jaroWinkler runs the similarly named algorithm over the two input strings and averages their match percentages
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
		max := maxMatch(s1Parts[i], s2Parts)
		if max >= 1.0 {
			max += exactMatchFavoritism
		}
		scores = append(scores, max)
	}

	// average the highest N scores where N is the words in our query (s2).
	sort.Float64s(scores)
	if len(s1Parts) > len(s2Parts) && len(s2Parts) > 2 {
		scores = scores[len(s1Parts)-len(s2Parts):]
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
