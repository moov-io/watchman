// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

import (
	"slices"
	"strings"

	"github.com/moov-io/watchman/internal/linksim"
	"github.com/moov-io/watchman/pkg/search"
)

func (c *corpus) indexBlockingKeys(e search.Entity[search.Value], idx int) {
	seen := make(map[string]struct{})
	for _, key := range linksim.Keys(e) {
		for _, p := range linksim.Prefixes(key) {
			if _, dup := seen[p]; dup {
				continue
			}
			seen[p] = struct{}{}
			c.blockKeys[p] = append(c.blockKeys[p], idx)
		}
	}
}

func (c *corpus) lookupBlockKey(key string, partition []int) []int {
	var hits []int
	for _, idx := range c.blockKeys[key] {
		if _, found := slices.BinarySearch(partition, idx); found {
			hits = append(hits, idx)
		}
	}
	return hits
}

func (c *corpus) cryptoHits(query search.Entity[search.Value], partition []int) []int {
	var hits []int
	for _, addr := range query.CryptoAddresses {
		key := cryptoKey(addr.Currency, addr.Address)
		if key == "" {
			continue
		}
		for _, idx := range c.cryptoKeys[key] {
			if _, found := slices.BinarySearch(partition, idx); found {
				hits = append(hits, idx)
			}
		}
	}
	return hits
}

func (c *corpus) governmentIDHits(query search.Entity[search.Value], partition []int) []int {
	var hits []int
	for _, key := range linksim.Keys(query) {
		if !strings.HasPrefix(key, linksim.KindGovID+":") {
			continue
		}
		hits = append(hits, c.lookupBlockKey(key, partition)...)
	}
	if len(hits) == 0 {
		return nil
	}
	slices.Sort(hits)
	return slices.Compact(hits)
}

func (c *corpus) addressHits(query search.Entity[search.Value], partition []int, opts CandidateOpts) []int {
	maxCount := int(float64(len(partition)) * opts.MaxFraction)
	if maxCount < 1 {
		maxCount = 1
	}

	var all []int
	found := false
	for _, key := range linksim.Keys(query) {
		if !strings.HasPrefix(key, linksim.KindAddr+":") {
			continue
		}
		prefixes := linksim.Prefixes(key)
		for i := len(prefixes) - 1; i >= 0; i-- {
			hits := c.lookupBlockKey(prefixes[i], partition)
			if len(hits) == 0 {
				continue
			}
			if len(hits) > maxCount {
				// Coarser prefixes are larger; stop widening this key.
				break
			}
			all = append(all, hits...)
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	slices.Sort(all)
	return slices.Compact(all)
}
