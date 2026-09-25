---
layout: page
title: Introduction
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# What Watchman is

Watchman downloads sanctions lists and scores customers and counterparties against them with an inspectable, multi-field matcher. It is the watchlist engine in a BSA/AML stack: onboarding, periodic refresh, and (where required) pre-transaction screening.

[Using Watchman](/watchman/using-watchman/) · [For compliance and risk](/watchman/methodology/for-compliance/) · [Similarity methodology](/watchman/methodology/)

## What you get

**Lists.** OFAC SDN and Non-SDN, US CSL, FinCEN 311, EU, UK, UN, OpenSanctions Senzing files, CSV ingest. Refresh on an interval or `POST /v2/data/refresh`. `GET /v2/listinfo` reports counts, hashes, and timestamps.

**Search.** `GET /v2/search` (and JSON POST) with `type`, name, aliases, government IDs (`gov_passport=US:…`), dates, addresses, crypto, contact. Ranked hits with a score in `[0, 1]`. Optional Senzing JSON. WASM UI at `/`. Go client. Experimental [MCP](/watchman/mcp/).

**Matcher.** Default Jaro–Winkler on normalized tokens, plus IDs, dates, and addresses. Unique identity keys (passport, IMO, crypto) can score 1.0. Tax IDs and email do not force a match. Conflicting same-type IDs lower the score. Optional TF-IDF and cross-script embeddings.

**Tuning.** `minMatch` is the policy cutoff. `algorithm` is per request. Embeddings, TF-IDF, ingest, address parsers, and geocoding are process-wide. See [Configuration](/watchman/config/).

## Lists

| Source | List |
|--------|------|
| **OpenSanctions** | [Senzing-formatted datasets](https://www.opensanctions.org/datasets/) |
| European Union | [Consolidated financial sanctions](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en) |
| US Government | [CSL](https://www.trade.gov/consolidated-screening-list), [FinCEN 311](https://home.treasury.gov/policy-issues/terrorism-and-illicit-finance/311-actions) |
| US Treasury | [OFAC](https://ofac.treasury.gov/sanctions-list-service) SDN and Non-SDN |
| United Kingdom | [UK sanctions list](https://www.gov.uk/government/publications/the-uk-sanctions-list) |
| United Nations | [UN consolidated list](https://www.un.org/sc/resources/sc-sanctions) |

`INCLUDED_LISTS=us_ofac,eu_csl,uk_csl,un_csl` to load a subset.

## How names are prepared

Before indexing (and on each query), Watchman normalizes:

1. **SDN name order** — `MADURO MOROS, Nicolas` → `Nicolas MADURO MOROS`
2. **Company suffixes** — strip `INC.`, `LLC`, and similar
3. **Stopwords** — drop `and`, `the`, `of` unless `KEEP_STOPWORDS=true`
4. **UTF-8** — lowercase, strip punctuation, fold diacritics (`Raúl` → `raul`)

That is why `nicolas maduro` hits `MADURO MOROS, Nicolas`. Details: [Pipeline](/watchman/pipeline/).

## Scoring in one paragraph

Watchman does not do “Google-style search.” It tokenizes names, aligns tokens with Jaro–Winkler (optional phonetic or n-gram inner metric), then blends identifier, date, address, and contact pieces. Name-only queries are down-ranked. The score is built so you can log `debug=true` pieces and defend the hit. OpenSanctions Pairs (755,540 labeled pairs) is the public evidence set; see [OpenSanctions Pairs](/watchman/opensanctions-pairs/).

## Next

[Using Watchman](/watchman/using-watchman/) · [Search](/watchman/search/) · [Docker](/watchman/usage-docker/) · [Configuration](/watchman/config/)
