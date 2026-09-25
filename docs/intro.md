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

**Matcher.** Names are compared token by token (Jaro–Winkler by default), then combined with IDs, dates, and addresses. A matching passport, IMO number, or crypto address can score 1.0. A matching tax number or email raises the score without forcing a match. Two IDs of the same type that disagree lower the score. Optional extras: TF-IDF (down-weight common words) and embeddings (better matches across writing systems).

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

Watchman is not a web search engine. It splits names into tokens, compares those tokens, then combines that result with IDs, dates, addresses, and contact. A name-only query is scored lower than the same name plus a date of birth or passport. Add `debug=true` to see which fields produced the score. A public labeled dataset we used to measure this is described in [OpenSanctions Pairs](/watchman/opensanctions-pairs/).

## Next

[Using Watchman](/watchman/using-watchman/) · [Search](/watchman/search/) · [Docker](/watchman/usage-docker/) · [Configuration](/watchman/config/)
