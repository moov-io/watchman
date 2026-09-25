---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Moov Watchman

![Moov Watchman Logo](https://repository-images.githubusercontent.com/163885848/41101f80-c6d9-11ea-9ab5-dc9f51b849df)

Watchman is an open-source **sanctions screening engine**. It downloads government watchlists (OFAC, EU, UK, UN, and others), keeps them in memory, and compares each customer or counterparty you send to those lists. You get a ranked list of possible matches and a score from 0 to 1. HTTP API, Go library, browser UI, optional MCP.

Start here: [Using Watchman](/watchman/using-watchman/) · [Docker](/watchman/usage-docker/) · [For compliance and risk](/watchman/methodology/for-compliance/)

## Why teams pick it

- **Named lists** — OFAC SDN and Non-SDN, US Consolidated Screening List, FinCEN 311, EU, UK, UN, plus your own CSV files.
- **Structured search** — person, business, organization, vessel, or aircraft, with name, aliases, IDs, dates, addresses, and contact.
- **How IDs work** — a matching passport, IMO number, or crypto address (with country, when it applies) scores 1.0. A matching tax number or email raises the score; it does not declare a match by itself. Two national IDs that disagree lower the score.
- **A cutoff you choose** — `minMatch` is the minimum score to return. 0.80 is a typical screening line; about 0.59 returns more possible hits.
- **Explainable hits** — `debug=true` shows which fields drove the score.
- **Built for onboarding and refresh** — lists are partitioned by source and type; name and ID indexes pick candidates before scoring.
- **Apache 2.0** — read the scorer, pin a release tag, run it in your network.

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
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&gov_passport=RU:2018278055&minMatch=0.80&limit=1" \
  | jq '{name: .entities[0].name, match: .entities[0].match}'
```

UI at `http://localhost:8084`. Production checklist: [Using Watchman](/watchman/using-watchman/). API: [Search](/watchman/search/). Knobs: [Configuration](/watchman/config/).

## About Moov

Moov builds open-source libraries for a single financial-services job each — ACH, wire, Watchman, and related formats — around correctness, performance, and code you can read.
