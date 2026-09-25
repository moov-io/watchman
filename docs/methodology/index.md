---
layout: page
title: Similarity Methodology
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Similarity Methodology

This page is the technical description of how Watchman scores a query against a watchlist entity. For a program-level briefing written for BSA/AML, sanctions, and model-risk staff, see [For compliance and risk](/watchman/methodology/for-compliance/). Empirical results on analyst-labeled OpenSanctions pairs are in [OpenSanctions Pairs](/watchman/opensanctions-pairs/).

Name scoring defaults to token-pairwise [Jaro–Winkler](https://www.tandfonline.com/doi/abs/10.1080/01621459.1989.10478785). Optional `?algorithm=` scorers are compared in [Algorithm comparison](/watchman/algorithm-comparison/).

## Why multi-field scoring

Watchman scores **identifiers, names, dates, addresses, contact, and supporting fields** together. Name-only queries produce large review queues. OFAC’s public search is a useful illustration: the query `Khamis Al` returns many 100% hits on the portal; Watchman’s name scorer spreads those same SDN names from about 0.26 to 0.87. See [Comparison with the OFAC portal](/watchman/methodology/pages/ofac-name-comparison/).

That design follows classical record linkage: Fellegi and Sunter treat matching as a decision on agreement patterns across fields, not a single string compare ([Fellegi & Sunter, 1969](https://www.tandfonline.com/doi/abs/10.1080/01621459.1969.10501049)). Jaro’s comparator, later extended by Winkler, is the usual way to turn typographical variation into a partial-agreement weight ([Jaro, 1989](https://www.tandfonline.com/doi/abs/10.1080/01621459.1989.10478785); [Winkler, 1990](https://eric.ed.gov/?id=ED325505)).

## Score pieces and weights

After `Normalize()` (case, punctuation, stopwords, phones, addresses), `Similarity` builds these pieces:

| Piece | Weight | Compared fields |
|-------|-------:|-----------------|
| Exact identifiers | 50 | Type + country + identifier (passport, tax ID, IMO/MMSI, serial, …) |
| Crypto addresses | 50 | Currency + address |
| Government IDs (loose) | 50 | Identifier with optional country (0.7–1.0) |
| Contact | 50 | Email, phone, fax |
| Name | 35 | Best pairwise token alignment, including alt and former names |
| Titles | 35 | Person titles / positions |
| Addresses | 25 | Structured address fields |
| Dates | 15 | Birth, death, incorporation, dissolution |
| Supporting | 15 | Sanctions programs, historical names |

Empty query fields are not compared. The blended score is a coverage-aware weighted average (`FINAL_SCORE_*` multipliers). Name-only queries are down-ranked (`FINAL_SCORE_NAME_ONLY_MULTIPLIER`, default 0.95).

### When a match is 1.0

A piece forces **1.0** only when it is **exact** (the identifier *and* country match) **and** the field is unique to one party: passport, national ID, IMO, MMSI, aircraft serial, or crypto address.

A matching tax number, company registration, email, or phone **raises** the blended score. It does not force 1.0. Related companies often share a tax ID; two people can share an office email. A vessel call sign alone also does not force 1.0.

### Identifier conflict

If both records have the same ID type (and country, when both set) with **different** values, the blended score is multiplied by `ID_CONFLICT_PENALTY_MULTIPLIER` (default 0.70). A missing ID on one side is not a conflict. A matching passport or IMO still scores 1.0 before this penalty.

### Entity type

A person query is compared to people, a business query to companies. If the query type is person, business, or organization and the list record is one of those three, Watchman still scores the pair (some source files store a person under a company-like type). A person query is not scored against a vessel or aircraft. `GET /v2/search?type=person` still only searches the person partition of the index.

## Name matching

### Jaro–Winkler

```
sim_jw(s1, s2) = sim_j(s1, s2) + p * l * (1 - sim_j(s1, s2))
```

- `sim_j` is Jaro similarity (common characters and transpositions)
- `p` is the prefix scale (default 0.1)
- `l` is the common prefix length, capped at 4

Watchman applies this **token-wise**, not on the raw full name:

1. Names are tokenized (`John Michael Smith` → `john`, `michael`, `smith`).
2. Tokens are paired with positional preference and a first-letter phonetic filter (first letters are rarely mistranscribed).
3. Length-difference penalties reduce scores when one token is much shorter.
4. Alternate and historical names are tried; scoring can skip remaining aliases once a high-confidence name hit is found.

### Optional algorithms

`?algorithm=` (or MCP `algorithm`) selects the token-pair metric without a restart:

| Value | Role |
|-------|------|
| `jaro-winkler` (default) | Token pairwise Jaro–Winkler |
| `soundex`, `double-metaphone`, `beider-morse` | Same alignment, phonetic boost when encodings overlap |
| `soft-bidist`, `soft-bisim`, `editex`, `nsim`, `nsim-3` | Replace the token-pair metric; BestPairs / length / first-letter filters stay |

Process-wide `USE_SOUNDEX_MATCHING` still applies when `algorithm` is omitted.

### TF-IDF and embeddings

Optional **TF-IDF** down-weights common tokens (`Limited`, `ООО`, `GmbH`) using an index built at list refresh.

Optional **embeddings** (Ollama, OpenAI, OpenRouter, Azure) add a neural name vector. Default `EMBEDDINGS_CROSS_SCRIPT_ONLY=true`: Latin queries stay on Jaro–Winkler; non-Latin queries also search the vector index. On OpenSanctions Pairs, this hybrid is the only config that moves transliteration recall in a material way. See [Cross-script matching](/watchman/cross-script-matching/) and [OpenSanctions Pairs](/watchman/opensanctions-pairs/).

## Entity-type matching

**Person.** Government IDs, given/family names and aliases, birth and death dates, titles, gender.

**Business / organization.** Tax and registration numbers as evidence (not identity), names and aliases, incorporation/dissolution, addresses.

**Vessel / aircraft.** IMO, MMSI, call sign, serial, ICAO, flag. IMO/MMSI/serial remain unique-identity keys.

## Thresholds as policy

The API returns a score in `[0, 1]`. `minMatch` is a **policy** cutoff, not a hidden model parameter.

| Band | Typical `minMatch` | Use |
|------|-------------------:|-----|
| High confidence | 0.95+ | Auto-block / auto-alert after identifier confirmation |
| Screening default | 0.80 | Production hit generation with high precision |
| High recall | ~0.59 | When missing a designation is costlier than extra review |
| Enhanced due diligence | 0.70–0.84 | Broader queues |

On 472,477 analyst-judged OpenSanctions subject pairs, Jaro–Winkler at 0.80 had precision 0.986 and recall 0.689. Dropping the cutoff to 0.59 raised recall to 0.920 at precision 0.945. With cross-script embeddings (hybrid) at 0.80, precision was 0.946 and recall 0.815.

## Candidate selection (before scoring)

Scoring runs only on candidates. See [Indexing](/watchman/indexing/) and [Performance](/watchman/performance/).

1. Partition by source and type.
2. Exact crypto / government-ID hits; prefix and QWERTY-near indexes on IMO, MMSI, serial, email, phone.
3. Name-token inverted index: intersect distinctive tokens; fall back to the union or the full partition.
4. Address prefix blocks for address-only queries.
5. Searches with ≤100 candidates skip the admission queue; larger searches take `SEARCH_MAX_IN_FLIGHT`.

Empty type under a known source with no entities of that type returns nothing. Empty `type=` selects the all-types partition for that source (slower). `type` is not required; send it so GET person/business fields are read and the search stays in one partition. A wrong type can miss the hit.

## Measured results

A public labeled set of 755,540 sanctions/OSINT pairs is described in [OpenSanctions Pairs](/watchman/opensanctions-pairs/). The evaluator in this repo is `go run ./research/opensanctions-pairs`. Full tables: [RESULTS.md](https://github.com/moov-io/watchman/blob/master/research/opensanctions-pairs/RESULTS.md).
