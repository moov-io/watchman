// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/dpl"
	"github.com/moov-io/watchman/pkg/ofac"

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

	// BIS
	DPs []*DP

	// TODO: this could be refactored into sub structs that have us/eu (and eventually others)

	// US Consolidated Screening List
	BISEntities      []*Result[csl.EL]
	MilitaryEndUsers []*Result[csl.MEU]
	SSIs             []*Result[csl.SSI]
	UVLs             []*Result[csl.UVL]
	ISNs             []*Result[csl.ISN]
	FSEs             []*Result[csl.FSE]
	PLCs             []*Result[csl.PLC]
	CAPs             []*Result[csl.CAP]
	DTCs             []*Result[csl.DTC]
	CMICs            []*Result[csl.CMIC]
	NS_MBSs          []*Result[csl.NS_MBS]

	// EU Consolidated List of Sactions
	EUCSL []*Result[csl.EUCSLRecord]

	// UK Consolidated List of Sactions - OFSI
	UKCSL []*Result[csl.UKCSLRecord]

	// UK Sanctions List
	UKSanctionsList []*Result[csl.UKSanctionsListRecord]

	// metadata
	lastRefreshedAt time.Time
	sync.RWMutex    // protects all above fields
	*syncutil.Gate  // limits concurrent processing

	pipe *pipeliner

	logger log.Logger
}

func newSearcher(logger log.Logger, pipeline *pipeliner, workers int) *searcher {
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
				weight:  bestPairsJaroWinkler(altTokens, s.Alts[i].name),
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
			alt.matchedName = v.matched
			out = append(out, alt)
		}
	}
	return out
}

// bestPairsJaroWinkler compares a search query to an indexed term (name, address, etc) and returns a decimal fraction
// score.
//
// The algorithm splits each string into tokens, and does a pairwise Jaro-Winkler score of all token combinations
// (outer product). The best match for each search token is chosen, such that each index token can be matched at most
// once.
//
// The pairwise scores are combined into an average in a way that corrects for character length, and the fraction of the
// indexed term that didn't match.
func bestPairsJaroWinkler(searchTokens []string, indexed string) float64 {
	type Score struct {
		score          float64
		searchTokenIdx int
		indexTokenIdx  int
	}

	indexedTokens := strings.Fields(indexed)
	searchTokensLength := sumLength(searchTokens)
	indexTokensLength := sumLength(indexedTokens)

	//Compare each search token to each indexed token. Sort the results in descending order
	scores := make([]Score, 0)
	for searchIdx, searchToken := range searchTokens {
		for indexIdx, indexedToken := range indexedTokens {
			score := customJaroWinkler(indexedToken, searchToken)
			scores = append(scores, Score{score, searchIdx, indexIdx})
		}
	}
	sort.Slice(scores[:], func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	//Pick the highest score for each search term, where the indexed token hasn't yet been matched
	matchedSearchTokens := make([]bool, len(searchTokens))
	matchedIndexTokens := make([]bool, len(indexedTokens))
	matchedIndexTokensLength := 0
	totalWeightedScores := 0.0
	for _, score := range scores {
		//If neither the search token nor index token have been matched so far
		if !matchedSearchTokens[score.searchTokenIdx] && !matchedIndexTokens[score.indexTokenIdx] {
			//Weight the importance of this word score by its character length
			searchToken := searchTokens[score.searchTokenIdx]
			indexToken := indexedTokens[score.indexTokenIdx]
			totalWeightedScores += score.score * float64(len(searchToken)+len(indexToken))

			matchedSearchTokens[score.searchTokenIdx] = true
			matchedIndexTokens[score.indexTokenIdx] = true
			matchedIndexTokensLength += len(indexToken)
		}
	}
	lengthWeightedAverageScore := totalWeightedScores / float64(searchTokensLength+matchedIndexTokensLength)

	//If some index tokens weren't matched by any search token, penalise this search a small amount. If this isn't done,
	//a query of "John Doe" will match "John Doe" and "John Bartholomew Doe" equally well.
	//Calculate the fraction of the index name that wasn't matched, apply a weighting to reduce the importance of
	//unmatched portion, then scale down the final score.
	matchedIndexLength := 0
	for i, str := range indexedTokens {
		if matchedIndexTokens[i] {
			matchedIndexLength += len(str)
		}
	}
	matchedFraction := float64(matchedIndexLength) / float64(indexTokensLength)
	return lengthWeightedAverageScore * scalingFactor(matchedFraction, unmatchedIndexPenaltyWeight)
}

func customJaroWinkler(s1 string, s2 string) float64 {
	score := smetrics.JaroWinkler(s1, s2, boostThreshold, prefixSize)

	if lengthMetric := lengthDifferenceFactor(s1, s2); lengthMetric < lengthDifferenceCutoffFactor {
		//If there's a big difference in matched token lengths, punish the score. Jaro-Winkler is quite permissive about
		//different lengths
		score = score * scalingFactor(lengthMetric, lengthDifferencePenaltyWeight)
	}
	if s1[0] != s2[0] {
		//Penalise words that start with a different characters. Jaro-Winkler is too lenient on this
		//TODO should use a phonetic comparison here, like Soundex
		score = score * differentLetterPenaltyWeight
	}
	return score
}

// scalingFactor returns a float [0,1] that can be used to scale another number down, given some metric and a desired
// weight
// e.g. If a score has a 50% value according to a metric, and we want a 10% weight to the metric:
//
//	scaleFactor := scalingFactor(0.5, 0.1)  // 0.95
//	scaledScore := score * scaleFactor
func scalingFactor(metric float64, weight float64) float64 {
	return 1.0 - (1.0-metric)*weight
}

func sumLength(strs []string) int {
	totalLength := 0
	for _, str := range strs {
		totalLength += len(str)
	}
	return totalLength
}

func lengthDifferenceFactor(s1 string, s2 string) float64 {
	ls1 := float64(len(s1))
	ls2 := float64(len(s2))
	min := math.Min(ls1, ls2)
	max := math.Max(ls1, ls2)
	return min / max
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
				weight:  bestPairsJaroWinkler(nameTokens, s.SDNs[i].name),
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
			sdn.matchedName = v.matched
			out = append(out, &sdn)
		}
	}
	return out
}

func (s *searcher) TopDPs(limit int, minMatch float64, name string) []DP {
	name = precompute(name)
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
				weight:  bestPairsJaroWinkler(nameTokens, s.DPs[i].name),
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
			dp.matchedName = v.matched
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

	// matchedName holds the highest scoring term from the search query
	matchedName string

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

func precomputeSDNs(sdns []*ofac.SDN, addrs []*ofac.Address, pipe *pipeliner) []*SDN {
	out := make([]*SDN, len(sdns))
	for i := range sdns {
		nn := sdnName(sdns[i], findAddresses(sdns[i].EntityID, addrs))

		if err := pipe.Do(nn); err != nil {
			pipe.logger.Logf("pipeline", fmt.Sprintf("problem pipelining SDN: %v", err))
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

	// match holds the match ratio for an Alt in search results
	match float64

	// matchedName holds the highest scoring term from the search query
	matchedName string

	// name is precomputed for speed
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

func precomputeAlts(alts []*ofac.AlternateIdentity, pipe *pipeliner) []*Alt {
	out := make([]*Alt, len(alts))
	for i := range alts {
		an := altName(alts[i])

		if err := pipe.Do(an); err != nil {
			pipe.logger.LogErrorf("problem pipelining SDN: %v", err)
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

func precomputeDPs(persons []*dpl.DPL, pipe *pipeliner) []*DP {
	out := make([]*DP, len(persons))
	for i := range persons {
		nn := dpName(persons[i])
		if err := pipe.Do(nn); err != nil {
			pipe.logger.LogErrorf("problem pipelining DP: %v", err)
			continue
		}
		out[i] = &DP{
			DeniedPerson: persons[i],
			name:         nn.Processed,
		}
	}
	return out
}

var (
	// Jaro-Winkler parameters
	boostThreshold = readFloat(os.Getenv("JARO_WINKLER_BOOST_THRESHOLD"), 0.7)
	prefixSize     = readInt(os.Getenv("JARO_WINKLER_PREFIX_SIZE"), 4)
	// Customised Jaro-Winkler parameters
	lengthDifferenceCutoffFactor  = readFloat(os.Getenv("LENGTH_DIFFERENCE_CUTOFF_FACTOR"), 0.9)
	lengthDifferencePenaltyWeight = readFloat(os.Getenv("LENGTH_DIFFERENCE_PENALTY_WEIGHT"), 0.3)
	differentLetterPenaltyWeight  = readFloat(os.Getenv("DIFFERENT_LETTER_PENALTY_WEIGHT"), 0.9)

	// Watchman parameters
	exactMatchFavoritism        = readFloat(os.Getenv("EXACT_MATCH_FAVORITISM"), 0.0)
	unmatchedIndexPenaltyWeight = readFloat(os.Getenv("UNMATCHED_INDEX_TOKEN_WEIGHT"), 0.15)
)

func readFloat(override string, value float64) float64 {
	if override != "" {
		n, err := strconv.ParseFloat(override, 32)
		if err != nil {
			panic(fmt.Errorf("unable to parse %q as float64", override)) //nolint:forbidigo
		}
		return n
	}
	return value
}

func readInt(override string, value int) int {
	if override != "" {
		n, err := strconv.ParseInt(override, 10, 32)
		if err != nil {
			panic(fmt.Errorf("unable to parse %q as int", override)) //nolint:forbidigo
		}
		return int(n)
	}
	return value
}

// jaroWinkler runs the similarly named algorithm over the two input strings and averages their match percentages
// according to the second string (assumed to be the user's query)
//
// Terms are compared between a few adjacent terms and accumulate the highest near-neighbor match.
//
// For more details see https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func jaroWinkler(s1, s2 string) float64 {
	return jaroWinklerWithFavoritism(s1, s2, exactMatchFavoritism)
}

var (
	adjacentSimilarityPositions = readInt(os.Getenv("ADJACENT_SIMILARITY_POSITIONS"), 3)
)

func jaroWinklerWithFavoritism(indexedTerm, query string, favoritism float64) float64 {
	maxMatch := func(indexedWord string, indexedWordIdx int, queryWords []string) (float64, string) {
		if indexedWord == "" || len(queryWords) == 0 {
			return 0.0, ""
		}

		// We're only looking for the highest match close
		start := indexedWordIdx - adjacentSimilarityPositions
		end := indexedWordIdx + adjacentSimilarityPositions

		var max float64
		var maxTerm string
		for i := start; i < end; i++ {
			if i >= 0 && len(queryWords) > i {
				score := smetrics.JaroWinkler(indexedWord, queryWords[i], boostThreshold, prefixSize)
				if score > max {
					max = score
					maxTerm = queryWords[i]
				}
			}
		}
		return max, maxTerm
	}

	indexedWords, queryWords := strings.Fields(indexedTerm), strings.Fields(query)
	if len(indexedWords) == 0 || len(queryWords) == 0 {
		return 0.0 // avoid returning NaN later on
	}

	var scores []float64
	for i := range indexedWords {
		max, term := maxMatch(indexedWords[i], i, queryWords)
		//fmt.Printf("%s maxMatch %s %f\n", indexedWords[i], term, max)
		if max >= 1.0 {
			// If the query is longer than our indexed term (and EITHER are longer than most names)
			// we want to reduce the maximum weight proportionally by the term difference, which
			// forces more terms to match instead of one or two dominating the weight.
			if (len(queryWords) > len(indexedWords)) && (len(indexedWords) > 3 || len(queryWords) > 3) {
				max *= (float64(len(indexedWords)) / float64(len(queryWords)))
				goto add
			}
			// If the indexed term is really short cap the match at 90%.
			// This sill allows names to match highly with a couple different characters.
			if len(indexedWords) == 1 && len(queryWords) > 1 {
				max *= 0.9
				goto add
			}
			// Otherwise, apply Perfect match favoritism
			max += favoritism
		add:
			scores = append(scores, max)
		} else {
			// If there are more terms in the user's query than what's indexed then
			// adjust the max lower by the proportion of different terms.
			//
			// We do this to decrease the importance of a short (often common) term.
			if len(queryWords) > len(indexedWords) {
				scores = append(scores, max*float64(len(indexedWords))/float64(len(queryWords)))
				continue
			}

			// Apply an additional weight based on similarity of term lengths,
			// so terms which are closer in length match higher.
			s1 := float64(len(indexedWords[i]))
			t := float64(len(term)) - 1
			weight := math.Min(math.Abs(s1/t), 1.0)

			scores = append(scores, max*weight)
		}
	}

	// average the highest N scores where N is the words in our query (query).
	// Only truncate scores if there are enough words (aka more than First/Last).
	sort.Float64s(scores)
	if len(indexedWords) > len(queryWords) && len(queryWords) > 5 {
		scores = scores[len(indexedWords)-len(queryWords):]
	}

	var sum float64
	for i := range scores {
		sum += scores[i]
	}
	return math.Min(sum/float64(len(scores)), 1.00)
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
