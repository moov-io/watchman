---
layout: page
title: Performance
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Watchman Performance Characteristics

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

Watchman is designed to deliver fast, reliable sanctions and watchlist screening for financial services, balancing compliance needs with performance demands. By leveraging several key optimizations,
Watchman ensures stable query times and efficient resource usage, even under heavy load. Customize these behaviors through the [Configuration Guide](/watchman/config/) to optimize for your specific environment.
These performance traits make it a robust choice for production environments where speed and accuracy are critical.

One of Watchman’s core strengths is its **precomputation and normalization** of data. On startup, Watchman downloads and processes sanctions lists (like US OFAC and US/UK/EU Consolidated Screening)
and applies transformations—such as reordering names (e.g., "MADURO MOROS, Nicolas" to "Nicolas MADURO MOROS") — to provide standardized results.

Watchman can use the [libpostal](https://github.com/openvenues/libpostal) library (with [Senzing's updated classifier, data, and parser](https://github.com/Senzing/libpostal-data)) to parse and normalize postal addresses,
improving match accuracy at the cost of higher memory usage. This upfront work ensures that searches are faster by reducing the need for on-the-fly processing, though it does introduce some memory overhead due to `libpostal`’s requirements.

> The libpostal library can use ~3GB of memory to load into memory. Run Watchman with enough memory to support your load ontop of libpostal's requirements. [Configure the Postal Pool](/watchman/config/#postalpool) to your needs.

Watchman operates entirely with **in-memory lists**, storing all sanction data in memory without disk persistence. This eliminates I/O bottlenecks, enabling rapid search operations.
The trade-off is that data is reloaded on restart, but this ensures freshness and avoids stale data slowing down queries. Combined with a high-performance search implementation using the
Jaro-Winkler algorithm, Watchman delivers quick and accurate fuzzy matching for names and addresses, with scoring from 0.0 (no match) to 1.0 (exact match).
The in-memory approach, paired with precomputed indexes, allows Watchman to handle large query volumes without relying on an external database.

## How a search is executed

Each `/v2/search` request roughly follows this path:

1. **Parse and normalize** the query (names, addresses, IDs).
2. **Select candidates** from the in-memory corpus using prebuilt indexes (see [Indexing](/watchman/indexing/)).
3. **Score candidates in parallel** with Jaro-Winkler similarity (and optional TF-IDF weighting).
4. **Keep a top-N heap** of the best matches above `minMatch`, then return JSON.

Candidate selection is **recall-safe relative to a full source/type partition scan**: if name tokens do not hit the inverted index (for example a pure typo with no shared tokens), Watchman falls back to scoring the entire matching partition rather than returning empty results.

### Candidate selection

On every list refresh Watchman builds:

| Structure | Purpose |
|-----------|---------|
| **Source × type partitions** | Restrict scoring to the requested `source` and/or `type` when provided |
| **Name-token inverted index** | Union of entities whose prepared primary, alt, or former names contain a query token |
| **Crypto address map** | Exact `CURRENCY:address` lookup for crypto-only (or crypto+name) queries |
| **TF-IDF term weights** (optional) | Precomputed per-entity weights so search does not recompute IDF on every comparison |

**Tips for faster queries**

- Always send `type=` (and `source=` when you only need one list). This shrinks the partition before token lookup.
- Prefer multi-token names when possible; shared tokens prune the candidate set aggressively.
- Identifier-heavy queries (crypto addresses, government IDs) use exact paths where available; name-less ID scans still use the type/source partition.

### Concurrency model

Watchman uses two layers of concurrency control:

1. **Admission control (`SEARCH_MAX_IN_FLIGHT`)** — limits how many full searches run at once (default: `GOMAXPROCS`). Extra requests wait rather than oversubscribing every CPU with stacked worker pools. Tune this when you run many concurrent clients against one instance.

2. **Per-search worker pool** — each admitted search fans out scoring across a dynamic number of goroutines (`Search.Goroutines` in config, or a fixed `SEARCH_GOROUTINE_COUNT`). A feedback loop (concurrency champion) adjusts the worker count based on recent search durations for stable latency on shared hardware.

Each worker maintains a **local top-K** of best matches and merges into the final result set after scoring, avoiding a shared mutex on every entity comparison.

As shown in the first graph below, which tracks search requests per second (req/s) over time, Watchman maintains consistent query performance despite fluctuating load.

![Graph 1: Stable query times with search req/s over time](../images/stable-response-times.png)

The second graph illustrates how Watchman dynamically adjusts goroutine group sizes, optimizing for overall time to score and keeping response timings steady.
This concurrency model was a significant improvement over earlier versions, where consistent load could cause slowdowns or crashes. The v0.5x series further refined this by consolidating to a single search model,
encouraging richer query data for better performance.

![Graph 2: Dynamic adjustment of goroutine group sizes for optimal scoring time](../images/dynamic-goroutines.png)

## Scoring hot path

Similarity scoring is allocation-conscious for bulk search:

- Score pieces are computed on the stack for the non-debug path.
- Former names and related prepared fields are normalized at index time (and query normalize), not on every comparison.
- Optional TF-IDF weights are attached to index entities when lists load; query weights are computed once per search.
- Jaro-Winkler feature flags (`DISABLE_PHONETIC_FILTERING`, `USE_SOUNDEX_MATCHING`, `SOUNDEX_BOOST_WEIGHT`) are read at process start (see [Similarity Configuration](/watchman/config/#similarity-configuration)). Changing them requires a restart.

Debug mode (`debug=true`) is substantially more expensive (extra buffers and detailed score pieces) and should stay off in production traffic.

## Related configuration

| Variable / setting | Role |
|--------------------|------|
| `SEARCH_MAX_IN_FLIGHT` | Max concurrent searches (admission control); default `GOMAXPROCS` |
| `SEARCH_GOROUTINE_COUNT` | Fixed workers per search; empty = dynamic champion |
| `Search.Goroutines` | Default / min / max for dynamic workers |
| `TFIDF_ENABLED` | Optional name-term weighting (index-time weights) |

See the full [Configuration Guide](/watchman/config/) and [Indexing](/watchman/indexing/) pages for details.
