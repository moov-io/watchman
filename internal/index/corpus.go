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

	// blockKeys maps PII-safe composite blocking keys and their segment
	// prefixes (see internal/linksim) to entity indices.
	blockKeys map[string][]int

	// Plaintext identifier indexes for prefix and QWERTY-near queries.
	// Hashed linksim keys cannot match a typed prefix.
	imo   idIndex
	mmsi  idIndex
	air   idIndex
	email idIndex
	phone idIndex
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
		blockKeys:    make(map[string][]int),
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
		// Always index under the "all sources / all types" key once.
		// Only add more specific keys when they are non-empty so empty Source/Type
		// does not append the same entity index multiple times to one partition.
		c.addToPartition("", "", i)
		if src != "" {
			c.addToPartition(src, "", i)
		}
		if typ != "" {
			c.addToPartition("", typ, i)
		}
		if src != "" && typ != "" {
			c.addToPartition(src, typ, i)
		}

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

		c.indexBlockingKeys(*e, i)
		c.indexPlainIdentifiers(*e, i)
	}

	c.imo.sort()
	c.mmsi.sort()
	c.air.sort()
	c.email.sort()
	c.phone.sort()

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

const (
	// distinctiveMaxDF is the maximum fraction of a partition a token may cover
	// and still be required in the name-token intersection. Common words above
	// this are optional (Limited, ООО, GmbH, 有限公司 — whatever is frequent here).
	distinctiveMaxDF = 0.20

	// optionalDFRatio: when a query token is this many times more frequent than
	// the next-rarest query token, it is optional. Extra legal-form words in any
	// language are usually the most common token in the query.
	optionalDFRatio = 2
)

// distinctiveQueryTokens picks which hitting query tokens to AND.
// Language-agnostic: uses document frequency in this partition, not a suffix list.
//
//   - Skip empty postings (already done by the caller).
//   - If two or more tokens hit, drop the most common when it is at least
//     optionalDFRatio times as frequent as the next (e.g. "Limited" vs "Shipping").
//   - Of the rest, keep only tokens at or below distinctiveMaxDF when any such
//     token exists, so a corpus-wide common word is not required.
func distinctiveQueryTokens(hitting []tokenPostings, partN float64) [][]int {
	slices.SortFunc(hitting, func(a, b tokenPostings) int {
		return len(a.hits) - len(b.hits)
	})
	if len(hitting) >= 2 {
		most := hitting[len(hitting)-1]
		next := hitting[len(hitting)-2]
		if len(most.hits) >= optionalDFRatio*len(next.hits) {
			hitting = hitting[:len(hitting)-1]
		}
	}
	var distinctive [][]int
	for _, h := range hitting {
		if float64(len(h.hits))/partN <= distinctiveMaxDF {
			distinctive = append(distinctive, h.hits)
		}
	}
	if len(distinctive) > 0 {
		return distinctive
	}
	out := make([][]int, len(hitting))
	for i, h := range hitting {
		out[i] = h.hits
	}
	return out
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
//
// The bool is true when the source key exists in the corpus (even if the type
// slice is empty). Callers must distinguish "unknown source" from "known source,
// zero entities of this type" so they do not fall back to a full corpus scan.
func (c *corpus) partitionIndices(source search.SourceList, entityType search.EntityType) ([]int, bool) {
	if c == nil {
		return nil, false
	}

	src := string(source)
	if source.IsRequestType() {
		src = ""
	}
	typ := string(entityType)

	byType, ok := c.bySourceType[src]
	if !ok {
		// Unknown source with no in-memory partition
		return nil, false
	}
	// Missing type key means an empty partition for that type, not an unknown source.
	idxs, _ := byType[typ]
	if idxs == nil {
		return []int{}, true
	}
	return idxs, true
}

// CandidateOpts controls candidate selection.
type CandidateOpts struct {
	// MaxFraction is the maximum fraction of the partition that candidates may
	// cover before falling back to the full partition. Default 0.5.
	MaxFraction float64
}

// Candidates is a read-only view of corpus entities to score.
// Entities aliases the in-memory generation (no per-search copy of the
// candidate set). The slice remains valid after SelectCandidates returns
// because the caller holds a reference to that generation.
type Candidates struct {
	Entities []search.Entity[search.Value]
	Indices  []int
	TFIDF    *tfidf.Index
}

// Len returns the number of candidates to score.
func (c Candidates) Len() int {
	return len(c.Indices)
}

// At returns the i-th candidate. It copies the entity header.
func (c Candidates) At(i int) search.Entity[search.Value] {
	return c.Entities[c.Indices[i]]
}

func (c *corpus) result(idxs []int) Candidates {
	if len(idxs) == 0 {
		return Candidates{Entities: c.entities, TFIDF: c.tfidf}
	}
	return Candidates{Entities: c.entities, Indices: idxs, TFIDF: c.tfidf}
}

func candidatesFromEntities(entities []search.Entity[search.Value], tfidfIndex *tfidf.Index) Candidates {
	if len(entities) == 0 {
		return Candidates{}
	}
	idxs := make([]int, len(entities))
	for i := range idxs {
		idxs[i] = i
	}
	return Candidates{Entities: entities, Indices: idxs, TFIDF: tfidfIndex}
}

// selectCandidates returns entities to score for the query.
//
// Strategy (never reduces recall below a full partition scan):
//  1. Restrict to source/type partition.
//  2. Crypto and government-ID hits are exact. IMO, MMSI, aircraft serial,
//     email, and phone also match prefixes and single QWERTY-adjacent typos.
//     Identifier hits are merged with name-token hits when the query has a name.
//  3. Name-token inverted index: intersect distinctive tokens using document
//     frequency in this partition (no language-specific suffix list). Extra
//     common query tokens (Limited, ООО, GmbH, …) do not drop a DBA that omits
//     them. If the intersection is empty, fall back to the union of those
//     hitting tokens. If no token hits, or the set is too large, use the full
//     partition.
//  4. Address-only queries use hashed ADDR prefix blocks when they prune the
//     partition; otherwise fall through.
//  5. Identifier-only / empty-name queries use the full partition.
func (c *corpus) selectCandidates(query search.Entity[search.Value], opts CandidateOpts) Candidates {
	if c == nil || len(c.entities) == 0 {
		return Candidates{}
	}

	if opts.MaxFraction <= 0 || opts.MaxFraction > 1 {
		opts.MaxFraction = 0.5
	}

	partition, sourceOK := c.partitionIndices(query.Source, query.Type)
	if !sourceOK {
		// Unknown source key: do not scan unrelated lists
		return Candidates{}
	}
	if len(partition) == 0 {
		// Known source (or all-sources) but no entities of this type
		return c.result(nil)
	}

	// Exact identifier fast path (crypto, GOVID, IMO, MMSI, aircraft serial, contact)
	idHits := c.identifierHits(query, partition)
	if len(idHits) > 0 {
		// Merge name candidates only when the query has name tokens.
		// Otherwise nameCandidateIndices returns the full partition and would
		// defeat the exact-identifier fast path.
		if len(query.PreparedFields.NameFields) > 0 {
			idHits = append(idHits, c.nameCandidateIndices(query, partition, opts)...)
		}
		slices.Sort(idHits)
		idHits = slices.Compact(idHits)
		return c.result(idHits)
	}

	// Name-based candidates
	if len(query.PreparedFields.NameFields) > 0 {
		return c.result(c.nameCandidateIndices(query, partition, opts))
	}

	// Address prefix blocking for address-only queries
	if addr := c.addressHits(query, partition, opts); len(addr) > 0 {
		return c.result(addr)
	}

	// Exact prepared name shortcut (name set but fields empty after stopwords)
	if name := query.PreparedFields.Name; name != "" {
		if exact := c.exactNames[name]; len(exact) > 0 {
			if filtered := intersectSorted(exact, partition); len(filtered) > 0 {
				return c.result(filtered)
			}
		}
	}

	// Identifier / type-only / empty query: full partition
	return c.result(partition)
}

type tokenPostings struct {
	hits []int
}

func (c *corpus) nameCandidateIndices(query search.Entity[search.Value], partition []int, opts CandidateOpts) []int {
	tokens := query.PreparedFields.NameFields
	if len(tokens) == 0 {
		return partition
	}

	// Per-token postings restricted to the partition. Skip tokens with no hits
	// so a misspelled word does not wipe a good match on the remaining tokens.
	partN := float64(len(partition))
	if partN < 1 {
		partN = 1
	}
	hitting := make([]tokenPostings, 0, len(tokens))
	for _, tok := range tokens {
		hits := intersectSorted(c.nameTokens[tok], partition)
		if len(hits) == 0 {
			continue
		}
		hitting = append(hitting, tokenPostings{hits: hits})
	}

	// No token hits (e.g. pure typos) → full partition to preserve recall
	if len(hitting) == 0 {
		return partition
	}

	lists := distinctiveQueryTokens(hitting, partN)

	// Intersect from the rarest distinctive token.
	slices.SortFunc(lists, func(a, b []int) int {
		return len(a) - len(b)
	})
	candidates := lists[0]
	for i := 1; i < len(lists); i++ {
		candidates = intersectSorted(candidates, lists[i])
		if len(candidates) == 0 {
			break
		}
	}

	// Disjoint distinctive tokens (John∩Smith empty while both hit) → union
	// so recall matches the previous union-of-postings behavior.
	if len(candidates) == 0 {
		candidates = unionSorted(lists)
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
