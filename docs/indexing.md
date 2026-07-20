---
layout: page
title: Indexing and TF-IDF
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Advanced Search Indexing for Superior Performance

Watchman maintains searchable entities in memory (with database options for ingested files) to deliver ultra-fast search responses.
The index automatically rebuilds with each data refresh, ensuring your compliance checks are always based on the latest information.

## What is built on refresh

When lists finish downloading and preparing, Watchman constructs an in-memory **search corpus** from every entity:

1. **Prepared fields** — normalized names, tokenized name fields (stopwords removed), alt names, former names, addresses, and contact data (see [pipeline](/watchman/pipeline/) and entity `Normalize()`).
2. **Source × type partitions** — index slices keyed by list source and entity type so queries with `source` / `type` only score the relevant subset.
3. **Name-token inverted index** — maps each significant prepared name token to entity positions (primary name, alternate names, and historical “Former Name” values).
4. **Crypto address index** — exact lookup by `CURRENCY:address` for fast crypto screening.
5. **Optional TF-IDF weights** — when enabled, term weights for each entity’s name fields are stored on the entity so search does not recompute them per comparison.

These structures are immutable for readers until the next successful refresh replaces the corpus atomically.

## Candidate selection at search time

Before Jaro-Winkler scoring, Watchman selects a **candidate set**:

| Query shape | Candidate strategy |
|-------------|-------------------|
| `type` and/or `source` set | Start from that partition only |
| Name tokens present | Union of inverted-index hits for those tokens, restricted to the partition |
| No token hits (e.g. heavy typos) | **Fall back to the full partition** (preserves recall) |
| Crypto address on query | Exact crypto hits (merged with name candidates when both are present) |
| Name-less / identifier-oriented | Full partition for the filtered source/type |

If candidates would cover most of the partition (default threshold: half the partition size), Watchman scores the full partition instead—token pruning would not save work.

Always pass **`type`** (and **`source`** when appropriate) on `/v2/search` for the best latency. See [Performance](/watchman/performance/) for concurrency and tuning.

## TF-IDF

Starting with `v0.57.0` Watchman builds an in-memory [TF-IDF](https://en.wikipedia.org/wiki/Tf%E2%80%93idf)
(Term Frequency–Inverse Document Frequency) index over the corpus of searchable entities, treating each entity's
normalized name (including primary names and alternate known-as entries from lists like OFAC SDN) as a "document"
composed of tokenized terms. This precomputes IDF values to weight terms inversely by their commonality across
the entire watchlist—rare terms (e.g., unique surnames or identifiers) receive higher weights, while common
ones (e.g., "bank" or frequent first names) are downweighted.

During name searches, TF-IDF weighting is optionally applied to boost the contribution of discriminative rare terms
in similarity scoring, improving match precision for fuzzy name matching against global sanctions and watchlists
without overemphasizing ubiquitous words.

When TF-IDF is enabled, weights for indexed entities are **materialized at corpus build time** and attached to each
entity’s prepared fields. Query-side weights are computed once per search request.

The feature is opt-in via [configuration, with tunable parameters](/watchman/config/#tf-idf-configuration) like
smoothing and IDF bounds to adapt to corpus size and growth.
