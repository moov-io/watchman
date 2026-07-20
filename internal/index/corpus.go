package index

import (
	"slices"
	"strings"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

// corpus holds precomputed search structures built when lists are updated.
// All fields are immutable after Build and safe for concurrent readers.
type corpus struct {
	entities []search.Entity[search.Value]
	tfidf    *tfidf.Index

	// bySourceType maps source -> entityType -> indices into entities.
	// Empty string keys mean "all sources" / "all types".
	bySourceType map[string]map[string][]int

	// nameTokens maps a prepared name token to entity indices that contain it
	// in primary, alt, or historical names.
	nameTokens map[string][]int

	// exactNames maps prepared full name -> entity indices.
	exactNames map[string][]int

	// cryptoKeys maps "CURRENCY:address" (upper currency) -> entity indices.
	cryptoKeys map[string][]int
}

// buildCorpus constructs partitions and inverted indexes from the entity list.
// It also attaches precomputed TF-IDF term weights onto each entity's PreparedFields.
func buildCorpus(entities []search.Entity[search.Value], tfidfIndex *tfidf.Index) *corpus {
	c := &corpus{
		entities:     entities,
		tfidf:        tfidfIndex,
		bySourceType: make(map[string]map[string][]int),
		nameTokens:   make(map[string][]int),
		exactNames:   make(map[string][]int),
		cryptoKeys:   make(map[string][]int),
	}

	tfidfEnabled := tfidfIndex != nil && tfidfIndex.Enabled()

	for i := range entities {
		e := &entities[i]

		// Precompute TF-IDF weights once at index time
		if tfidfEnabled {
			e.PreparedFields.NameWeights = tfidfIndex.GetWeights(e.PreparedFields.NameFields)
			if len(e.PreparedFields.AltNameFields) > 0 {
				e.PreparedFields.AltNameWeights = make([][]float64, len(e.PreparedFields.AltNameFields))
				for j := range e.PreparedFields.AltNameFields {
					e.PreparedFields.AltNameWeights[j] = tfidfIndex.GetWeights(e.PreparedFields.AltNameFields[j])
				}
			}
			if len(e.PreparedFields.HistoricalNameFields) > 0 {
				e.PreparedFields.HistoricalNameWeights = make([][]float64, len(e.PreparedFields.HistoricalNameFields))
				for j := range e.PreparedFields.HistoricalNameFields {
					e.PreparedFields.HistoricalNameWeights[j] = tfidfIndex.GetWeights(e.PreparedFields.HistoricalNameFields[j])
				}
			}
		}

		src := string(e.Source)
		typ := string(e.Type)
		c.addToPartition("", "", i)
		c.addToPartition(src, "", i)
		c.addToPartition("", typ, i)
		c.addToPartition(src, typ, i)

		// Exact prepared name
		if name := e.PreparedFields.Name; name != "" {
			c.exactNames[name] = append(c.exactNames[name], i)
		}

		// Name tokens (primary, alt, historical) — one posting per token per entity
		seenTokens := make(map[string]struct{})
		addTokens := func(tokens []string) {
			for _, tok := range tokens {
				if tok == "" {
					continue
				}
				if _, dup := seenTokens[tok]; dup {
					continue
				}
				seenTokens[tok] = struct{}{}
				c.nameTokens[tok] = append(c.nameTokens[tok], i)
			}
		}
		addTokens(e.PreparedFields.NameFields)
		for _, alt := range e.PreparedFields.AltNameFields {
			addTokens(alt)
		}
		for _, hist := range e.PreparedFields.HistoricalNameFields {
			addTokens(hist)
		}

		// Crypto addresses for exact lookup
		for _, addr := range e.CryptoAddresses {
			key := cryptoKey(addr.Currency, addr.Address)
			if key != "" {
				c.cryptoKeys[key] = append(c.cryptoKeys[key], i)
			}
		}
	}

	return c
}

func (c *corpus) addToPartition(source, entityType string, idx int) {
	byType, ok := c.bySourceType[source]
	if !ok {
		byType = make(map[string][]int)
		c.bySourceType[source] = byType
	}
	byType[entityType] = append(byType[entityType], idx)
}

func cryptoKey(currency, address string) string {
	currency = strings.ToUpper(strings.TrimSpace(currency))
	address = strings.TrimSpace(address)
	if currency == "" || address == "" {
		return ""
	}
	return currency + ":" + address
}

// partitionIndices returns entity indices for the given source and type filters.
// Empty source or type means "all".
func (c *corpus) partitionIndices(source search.SourceList, entityType search.EntityType) []int {
	if c == nil {
		return nil
	}

	src := string(source)
	if source.IsRequestType() {
		src = ""
	}
	typ := string(entityType)

	byType := c.bySourceType[src]
	if byType == nil {
		// Unknown source with no in-memory partition
		return nil
	}
	return byType[typ]
}

// CandidateOpts controls candidate selection.
type CandidateOpts struct {
	// MaxFraction is the maximum fraction of the partition that candidates may
	// cover before falling back to the full partition. Default 0.5.
	MaxFraction float64
}

// selectCandidates returns entities to score for the query.
//
// Strategy (never reduces recall below a full partition scan):
//  1. Restrict to source/type partition.
//  2. Exact crypto address hits short-circuit to those entities.
//  3. Name-token inverted index: union of postings for query name tokens,
//     intersected with the partition. If empty or too large, use full partition.
//  4. Identifier-only / empty-name queries use the full partition.
func (c *corpus) selectCandidates(query search.Entity[search.Value], opts CandidateOpts) []search.Entity[search.Value] {
	if c == nil || len(c.entities) == 0 {
		return nil
	}

	if opts.MaxFraction <= 0 || opts.MaxFraction > 1 {
		opts.MaxFraction = 0.5
	}

	partition := c.partitionIndices(query.Source, query.Type)
	if partition == nil {
		// Fall back to scanning everything when partitions missing
		return c.entities
	}
	if len(partition) == 0 {
		return nil
	}

	// Crypto exact-address fast path
	if len(query.CryptoAddresses) > 0 {
		var hits []int
		seen := make(map[int]struct{})
		for _, addr := range query.CryptoAddresses {
			key := cryptoKey(addr.Currency, addr.Address)
			for _, idx := range c.cryptoKeys[key] {
				// partition is sorted ascending (built by appending increasing indices)
				if _, found := slices.BinarySearch(partition, idx); !found {
					continue
				}
				if _, ok := seen[idx]; ok {
					continue
				}
				seen[idx] = struct{}{}
				hits = append(hits, idx)
			}
		}
		// Crypto-only query: return exact hits (may be empty → fall through)
		if len(hits) > 0 && len(query.PreparedFields.NameFields) == 0 && query.PreparedFields.Name == "" {
			return c.materialize(hits)
		}
		if len(hits) > 0 {
			// Merge with name candidates so multi-field queries stay complete
			for _, idx := range c.nameCandidateIndices(query, partition, opts) {
				if _, ok := seen[idx]; !ok {
					seen[idx] = struct{}{}
					hits = append(hits, idx)
				}
			}
			return c.materialize(hits)
		}
	}

	// Name-based candidates
	if len(query.PreparedFields.NameFields) > 0 {
		idxs := c.nameCandidateIndices(query, partition, opts)
		return c.materialize(idxs)
	}

	// Exact prepared name shortcut (name set but fields empty after stopwords)
	if name := query.PreparedFields.Name; name != "" {
		if exact := c.exactNames[name]; len(exact) > 0 {
			filtered := intersectSorted(exact, partition)
			if len(filtered) > 0 {
				return c.materialize(filtered)
			}
		}
	}

	// Identifier / type-only / empty query: full partition
	return c.materialize(partition)
}

func (c *corpus) nameCandidateIndices(query search.Entity[search.Value], partition []int, opts CandidateOpts) []int {
	tokens := query.PreparedFields.NameFields
	if len(tokens) == 0 {
		return partition
	}

	// Union postings for query tokens, restricted to partition.
	// Membership uses binary search over the sorted partition (no map alloc per query).
	seen := make(map[int]struct{})
	var candidates []int
	for _, tok := range tokens {
		for _, idx := range c.nameTokens[tok] {
			if _, found := slices.BinarySearch(partition, idx); !found {
				continue
			}
			if _, ok := seen[idx]; ok {
				continue
			}
			seen[idx] = struct{}{}
			candidates = append(candidates, idx)
		}
	}

	// No token hits (e.g. pure typos) → full partition to preserve recall
	if len(candidates) == 0 {
		return partition
	}

	// If candidates cover too much of the partition, scoring them is no cheaper
	maxCount := int(float64(len(partition)) * opts.MaxFraction)
	if maxCount < 1 {
		maxCount = 1
	}
	if len(candidates) > maxCount {
		return partition
	}

	return candidates
}

func (c *corpus) materialize(idxs []int) []search.Entity[search.Value] {
	out := make([]search.Entity[search.Value], len(idxs))
	for i, idx := range idxs {
		out[i] = c.entities[idx]
	}
	return out
}

// intersectSorted returns intersection of two ascending slices.
func intersectSorted(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	var out []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			out = append(out, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return out
}
