---
layout: page
title: Using Watchman
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Using Watchman

This page is a path from first search to a production setup. How scoring works: [Similarity methodology](/watchman/methodology/). For BSA/AML and sanctions officers: [For compliance and risk](/watchman/methodology/for-compliance/).

## What you are running

Watchman downloads government sanctions lists, keeps them in memory, and compares each query you send to those lists. Each possible match has a **score from 0 to 1**. You choose which scores become alerts (`minMatch`) and what your process does with them.

Lists include OFAC, EU, UK, UN, the US Consolidated Screening List, FinCEN 311, and optional extra files you ingest.

## Five-minute start

```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
```

Open [http://localhost:8084](http://localhost:8084) for the WASM UI. Then:

```
curl -s "http://localhost:8084/v2/search?type=person&name=Nicolas+Maduro&birthDate=1962-11-23&limit=5&minMatch=0.80" | jq .
```

Always send **`type=`** (`person`, `business`, `vessel`, …). Include government IDs and dates of birth when you have them. `minMatch=0.80` means “only return hits that score at least 0.80.” That is a typical starting cutoff.

Confirm lists with `GET /v2/listinfo` (counts, hashes, refresh window, version).

Admin metrics stay on **:9094**. Do not put Watchman on the public internet. See [Network access](/watchman/network/).

## Recommended production search

| Goal | How |
|------|-----|
| Catch designated parties | `type` + name + government IDs + DOB/address |
| Keep the alert queue small | `minMatch=0.80` (try 0.59 if you would rather see more possible matches) |
| Names in Arabic, Cyrillic, Chinese, etc. | Enable embeddings; keep `EMBEDDINGS_CROSS_SCRIPT_ONLY=true` |
| One list only | `source=us_ofac` (faster) |
| Exact SDN row | `sourceID=22790` |
| Explain a hit | `debug=true` |

Government IDs on the query string:

```
gov_passport=IR:Y53914915
gov_national=PK:35201114139885
gov_tax=RU:9709063550
```

Format is `gov_<type>=COUNTRY:IDENTIFIER`. When a **passport, national ID, or IMO** matches on type, country, and number, the score is **1.0**. A matching **tax number or email** raises the score; it does not force 1.0, because those values are often shared.

Unknown customer type: call `type=person` and `type=business` (and `vessel` / `aircraft` when IMO or serial is present). Empty `type=` is slower and is not the production path.

## Thresholds

`minMatch` is the lowest score Watchman will return. It is a policy choice.

- **0.80** — typical screening line: most returned hits are real matches; some true matches stay below the line.
- **0.59** — more possible matches returned; more items for analysts to review.

**Precision** is the share of returned hits that are real matches. **Recall** is the share of real matches that were returned.

On a public labeled set of 472,477 people, companies, and vessels ([OpenSanctions Pairs](/watchman/opensanctions-pairs/)):

| `minMatch` | Precision | Recall | When to use |
|-----------:|----------:|-------:|-------------|
| **0.80** | 0.986 without embeddings / 0.946 with cross-script embeddings | 0.689 / 0.815 | Default production queue |
| **0.59** | 0.945 / 0.876 | 0.920 / 0.942 | When missing a designation is costlier than extra review |

## Settings that matter most

Small Jaro–Winkler environment flags (prefix size, Soundex) barely change those numbers. These settings do:

| Setting | Default | Tip |
|------|---------|-----|
| Query fields | name only | Send IDs and dates. Name-only is down-ranked (`FINAL_SCORE_NAME_ONLY_MULTIPLIER=0.95`). |
| `minMatch` | 0 | Set 0.80 in production. |
| `type` / `source` | empty | Always set. Partitions the in-memory corpus. |
| `EMBEDDINGS_*` | off | Improves matching when one name is Latin and the other is Arabic, Cyrillic, Chinese, and similar. Example model: `qwen3-embedding:0.6b` via Ollama. Keep `CROSS_SCRIPT_ONLY=true`. |
| `TFIDF_ENABLED` | false | Down-weights common words (`Limited`, `GmbH`). Slightly more true hits and slightly more false hits. |
| `algorithm` | jaro-winkler | Per-request name metric. Phonetic options (Soundex, and others) help little on transliteration. |
| `ID_CONFLICT_PENALTY_MULTIPLIER` | 0.70 | When both records have the same ID type and country but different numbers. Set `1` to turn this off. |
| `INCLUDED_LISTS` | all downloadable | Start with `us_ofac`. Add `us_csl`, `eu_csl`, `uk_csl`, `un_csl` as needed. |
| `SEARCH_MAX_IN_FLIGHT` | GOMAXPROCS | Caps concurrent *large* searches. Tight queries (≤100 candidates) skip the queue. |

Full tables: [Configuration](/watchman/config/).

## Example queries

Person with IDs and date of birth:

```
GET /v2/search?type=person&name=Aliasghar+Norouzi&birthDate=1962-11-11&gov_passport=IR:Y53914915&minMatch=0.80&limit=10
```

Company:

```
GET /v2/search?type=business&name=Tidewater&gov_tax=RU:1234567890&address=Moscow&minMatch=0.80
```

Vessel:

```
GET /v2/search?type=vessel&name=NS+LEADER&imoNumber=9339301
```

Debug one hit:

```
GET /v2/search?type=person&name=Nicolas+Maduro&minMatch=0.80&debug=true&limit=1
```

JSON body (UTF-8 names):

```
curl -s http://localhost:8084/v2/search \
  -H 'Content-Type: application/json' \
  -d '{"name":"محمد علي","type":"person"}'
```

Senzing output: `Accept: senzing` or `?format=senzing`.

## Lists and ingest

Watchman downloads and refreshes lists on `DATA_REFRESH_INTERVAL` (default 12h). `GET /v2/listinfo` is the freshness check for ops and exams.

Add internal lists with [ingest](/watchman/ingest/) (`POST /v2/ingest/{fileType}`). OpenSanctions Senzing files are first-class via [config](/watchman/config/#open-sanctions).

## Addresses and geocoding

**Docker images** use libpostal in-process (~3GB RAM for models). **Any deployment** can switch to deepparse (`Watchman.Deepparse.Enabled`). **Otherwise** (GitHub binaries, `go run`) Watchman uses `usaddress`. Optional **geocoding** (OpenCage, Nominatim, Google) fills lat/long. See [Addresses](/watchman/addresses/) and [Geocoding](/watchman/geocoding/).

## Where it sits in the program

Use Watchman when you onboard a customer, when you refresh that customer later, and when a payment or other transaction must be checked against sanctions lists. Send hits to investigation with the score details. Customer due diligence, transaction monitoring, and ownership analysis stay in your other systems.

## Next

- [Docker](/watchman/usage-docker/) · [Go client](/watchman/usage-go/) · [MCP](/watchman/mcp/)
- [Search](/watchman/search/) · [Indexing](/watchman/indexing/) · [Performance](/watchman/performance/)
- [For compliance and risk](/watchman/methodology/for-compliance/)
