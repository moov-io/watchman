// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package index

import (
	"slices"
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

// Minimum query length before prefix/QWERTY lookup is used. Shorter values
// match too much of the partition to be useful as a typed filter.
const (
	minIMOPrefix    = 4
	minMMSIPrefix   = 4
	minSerialPrefix = 4
	minEmailPrefix  = 3
	minPhonePrefix  = 4
)

type idEntry struct {
	value string
	idx   int
}

// idIndex is a sorted list of normalized identifier values for prefix scans
// and exact lookup of QWERTY-near variants. Hashing would destroy prefixes.
type idIndex struct {
	entries []idEntry
}

func (ix *idIndex) add(value string, idx int) {
	value = normalizeID(value)
	if value == "" {
		return
	}
	ix.entries = append(ix.entries, idEntry{value: value, idx: idx})
}

func (ix *idIndex) sort() {
	slices.SortFunc(ix.entries, func(a, b idEntry) int {
		if a.value == b.value {
			return a.idx - b.idx
		}
		if a.value < b.value {
			return -1
		}
		return 1
	})
}

func (ix *idIndex) exact(value string) []int {
	if value == "" || len(ix.entries) == 0 {
		return nil
	}
	i := sortSearchLowerBound(ix.entries, value)
	var out []int
	for ; i < len(ix.entries) && ix.entries[i].value == value; i++ {
		out = append(out, ix.entries[i].idx)
	}
	return out
}

func (ix *idIndex) prefix(value string) []int {
	if value == "" || len(ix.entries) == 0 {
		return nil
	}
	i := sortSearchLowerBound(ix.entries, value)
	var out []int
	for ; i < len(ix.entries) && strings.HasPrefix(ix.entries[i].value, value); i++ {
		out = append(out, ix.entries[i].idx)
	}
	return out
}

func sortSearchLowerBound(entries []idEntry, value string) int {
	lo, hi := 0, len(entries)
	for lo < hi {
		mid := (lo + hi) / 2
		if entries[mid].value < value {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// match returns entity indices whose identifier equals q, starts with q, or
// equals/starts with a single QWERTY-adjacent substitution or transposition of q.
func (ix *idIndex) match(raw string, minPrefix int) []int {
	q := normalizeID(raw)
	if q == "" {
		return nil
	}
	var hits []int
	for _, v := range qwertyVariants(q) {
		hits = append(hits, ix.exact(v)...)
		if len(v) >= minPrefix {
			hits = append(hits, ix.prefix(v)...)
		}
	}
	if len(hits) == 0 {
		return nil
	}
	slices.Sort(hits)
	return slices.Compact(hits)
}

func normalizeID(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.NewReplacer(" ", "", "-", "", ".", "").Replace(s)
}

func (c *corpus) indexPlainIdentifiers(e search.Entity[search.Value], idx int) {
	if e.Vessel != nil {
		c.imo.add(e.Vessel.IMONumber, idx)
		c.mmsi.add(e.Vessel.MMSI, idx)
	}
	if e.Aircraft != nil {
		c.air.add(e.Aircraft.SerialNumber, idx)
	}
	for _, email := range e.Contact.EmailAddresses {
		c.email.add(email, idx)
	}
	phones := e.PreparedFields.Contact.PhoneNumbers
	if len(phones) == 0 {
		phones = e.Contact.PhoneNumbers
	}
	for _, phone := range phones {
		c.phone.add(phone, idx)
	}
}

func (c *corpus) plainIdentifierHits(query search.Entity[search.Value]) []int {
	var hits []int
	if query.Vessel != nil {
		hits = append(hits, c.imo.match(query.Vessel.IMONumber, minIMOPrefix)...)
		hits = append(hits, c.mmsi.match(query.Vessel.MMSI, minMMSIPrefix)...)
	}
	if query.Aircraft != nil {
		hits = append(hits, c.air.match(query.Aircraft.SerialNumber, minSerialPrefix)...)
	}
	for _, email := range query.Contact.EmailAddresses {
		hits = append(hits, c.email.match(email, minEmailPrefix)...)
	}
	phones := query.PreparedFields.Contact.PhoneNumbers
	if len(phones) == 0 {
		phones = query.Contact.PhoneNumbers
	}
	for _, phone := range phones {
		hits = append(hits, c.phone.match(phone, minPhonePrefix)...)
	}
	if len(hits) == 0 {
		return nil
	}
	slices.Sort(hits)
	return slices.Compact(hits)
}
