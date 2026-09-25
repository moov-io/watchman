---
layout: page
title: Using Watchman
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Using Watchman

A practical path from first search to a production screening setup. Technical scoring is in [Similarity methodology](/watchman/methodology/). Program language is in [For compliance and risk](/watchman/methodology/for-compliance/).

## What you are running

Watchman is a **list-driven screener**. It downloads OFAC, EU, UK, UN, US CSL, FinCEN 311, and optional OpenSanctions/Senzing files, indexes them in memory, and scores each query with a multi-field matcher. You get a ranked hit list and a score in `[0, 1]`. You decide what to do with hits.

That is the control Wolfsberg calls sanctions screening: compare customer/counterparty text (and IDs) to designated-party lists, then investigate.

## Five-minute start

```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
```

Open [http://localhost:8084](http://localhost:8084) for the WASM UI. Then:

```
curl -s "http://localhost:8084/v2/search?type=person&name=Nicolas+Maduro&birthDate=1962-11-23&limit=5&minMatch=0.80" | jq .
```

Always send **`type=`**. Add **IDs and dates when CDD has them**. `minMatch=0.80` is the production-shaped cutoff from the OpenSanctions Pairs evaluation.

Confirm lists with `GET /v2/listinfo` (counts, hashes, refresh window, version).

Admin metrics stay on **:9094**. Do not put Watchman on the public internet. See [Network access](/watchman/network/).

## Recommended production search

| Goal | How |
|------|-----|
| Catch designated parties | `type` + name + government IDs + DOB/address |
| Keep the queue reviewable | `minMatch=0.80` (or 0.59 if miss-rate is the binding constraint) |
| Transliteration (Arabic, Cyrillic, CJK) | Embeddings on, `EMBEDDINGS_CROSS_SCRIPT_ONLY=true` |
| One list only | `source=us_ofac` (faster) |
| Exact SDN row | `sourceID=22790` |
| Explain a hit | `debug=true` |

Government IDs on the query string:

```
gov_passport=IR:Y53914915
gov_national=PK:35201114139885
gov_tax=RU:9709063550
```

Format is `gov_<type>=COUNTRY:IDENTIFIER`. Passport / national ID / IMO still short-circuit to **1.0** when type, country, and identifier all match. Tax IDs and emails do **not** — they stay in the weighted blend so related companies sharing an INN are not treated as the same person.

Unknown customer type: call `type=person` and `type=business` (and `vessel` / `aircraft` when IMO or serial is present). Empty `type=` is slower and is not the production path.

## Thresholds (policy, not magic)

On 472,477 analyst-judged OpenSanctions subject pairs (people, companies, vessels):

| `minMatch` | Precision | Recall | When to use |
|-----------:|----------:|-------:|-------------|
| **0.80** | 0.986 (JW) / 0.946 (hybrid) | 0.689 / 0.815 | Default production queue |
| **0.59** | 0.945 / 0.876 | 0.920 / 0.942 | When missing a designation is costlier than extra review |

Precision = share of hits that are real. Recall = share of true matches you catch. Write the cutoff into the risk assessment. Full tables: [OpenSanctions Pairs](/watchman/opensanctions-pairs/).

## Knobs that actually move results

Most Jaro–Winkler env flags (prefix size, length penalty, Soundex) moved **F1 by ~0.001** on the 755k-pair dump. Spend time here instead:

| Knob | Default | Tip |
|------|---------|-----|
| Query fields | name only | Send IDs and dates. Name-only is down-ranked (`FINAL_SCORE_NAME_ONLY_MULTIPLIER=0.95`). |
| `minMatch` | 0 | Set 0.80 in production. |
| `type` / `source` | empty | Always set. Partitions the in-memory corpus. |
| `EMBEDDINGS_*` | off | Largest recall win on non-Latin names. Use `qwen3-embedding:0.6b` (or similar) via Ollama; keep `CROSS_SCRIPT_ONLY=true`. |
| `TFIDF_ENABLED` | false | Small recall bump, more false positives. Useful for ranking, not a substitute for embeddings. |
| `algorithm` | jaro-winkler | Per-request. Phonetic boosts and n-gram scorers did not close transliteration. |
| `ID_CONFLICT_PENALTY_MULTIPLIER` | 0.70 | Same ID type + country, different values (two CNICs). Set `1` to disable. |
| `INCLUDED_LISTS` | all downloadable | Start with `us_ofac`. Add `us_csl`, `eu_csl`, `uk_csl`, `un_csl` as the risk assessment names them. |
| `SEARCH_MAX_IN_FLIGHT` | GOMAXPROCS | Caps concurrent *large* searches. Tight queries (≤100 candidates) skip the queue. |

Full tables: [Configuration](/watchman/config/).

## Example queries

Person with CDD fields:

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

Onboarding, periodic refresh, and (where required) pre-transaction screening against named lists. Hits go to investigation with score pieces. Watchman does not replace CDD, transaction monitoring, or beneficial-ownership analysis.

## Next

- [Docker](/watchman/usage-docker/) · [Go client](/watchman/usage-go/) · [MCP](/watchman/mcp/)
- [Search](/watchman/search/) · [Indexing](/watchman/indexing/) · [Performance](/watchman/performance/)
- [For compliance and risk](/watchman/methodology/for-compliance/)
