---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Moov Watchman

![Moov Watchman Logo](https://repository-images.githubusercontent.com/163885848/41101f80-c6d9-11ea-9ab5-dc9f51b849df)

Watchman is an open-source **sanctions screening engine**: download OFAC, EU, UK, UN, and related lists, index them in memory, and score each customer or counterparty with an inspectable multi-field matcher. HTTP API, Go library, WASM UI, optional MCP.

On **755,540** analyst-labeled OpenSanctions pairs, Jaro–Winkler at `minMatch=0.80` had **subject precision 0.99**. With cross-script embeddings, subject recall rose from **0.69 to 0.82** while precision stayed **0.95**. That is a production-shaped queue: few junk alerts, transliteration covered when embeddings are on.

[Using Watchman](/watchman/using-watchman/) · [For compliance and risk](/watchman/methodology/for-compliance/) · [Docker](/watchman/usage-docker/)

## Why teams pick it

- **Lists you can name** — OFAC SDN and Non-SDN, US CSL, FinCEN 311, EU, UK, UN, OpenSanctions Senzing files, plus CSV ingest of your own data.
- **Structured search** — person, business, organization, vessel, aircraft. Names, aliases, government IDs, dates, addresses, crypto, contact.
- **Identity vs evidence** — matching passport / IMO / crypto (type + country + identifier) scores 1.0. Shared tax IDs and emails do not force a match. Conflicting national IDs penalize the score.
- **You set the cutoff** — `minMatch` is policy (0.80 default screening, ~0.59 high recall), not a hidden model parameter.
- **Explainable hits** — `debug=true` returns field-level pieces for investigation and model-risk review.
- **Fast enough for onboarding and refresh** — source/type partitions, name-token and ID candidate indexes, parallel scoring. Tight queries skip the admission queue.
- **Apache 2.0** — read the scorer, pin a tag, run it in your VPC.

## Included lists

Use `INCLUDED_LISTS` or the [config file](/watchman/config/#included-lists) to choose what loads. Empty means the built-in downloadable lists.

| Source | List |
|--------|------|
| **OpenSanctions** | [Senzing-formatted datasets](https://www.opensanctions.org/datasets/) |
| European Union | [Consolidated financial sanctions](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en) |
| US Government | [Consolidated Screening List](https://www.trade.gov/consolidated-screening-list), [FinCEN 311](https://home.treasury.gov/policy-issues/terrorism-and-illicit-finance/311-actions) |
| US Treasury | [OFAC](https://ofac.treasury.gov/sanctions-list-service) SDN and Non-SDN |
| United Kingdom | [UK sanctions list](https://www.gov.uk/government/publications/the-uk-sanctions-list) |
| United Nations | [UN consolidated list](https://www.un.org/sc/resources/sc-sanctions) |

## Start here

```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
curl -s "http://localhost:8084/v2/search?type=person&name=Nicolas+Maduro&minMatch=0.80&limit=1" | jq .
```

UI at `http://localhost:8084`. Production checklist: [Using Watchman](/watchman/using-watchman/). API: [Search](/watchman/search/). Knobs: [Configuration](/watchman/config/).

## About Moov

Moov builds open-source libraries for a single financial-services job each — ACH, wire, Watchman, and related formats — around correctness, performance, and code you can read.
