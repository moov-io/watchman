---
layout: page
title: Data preparation Pipeline
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Data preparation pipeline

Watchman prepares list records and search queries the same way so a query like `nicolas maduro` can match a list name stored as `MADURO MOROS, Nicolas`. This page is that preparation. Scoring is [Similarity methodology](/watchman/methodology/). Candidate indexes are [Indexing](/watchman/indexing/).

There are two layers:

1. **List ingest** — extra steps some sources run when a file is loaded. OFAC is the main one.
2. **`Normalize()`** — runs on every list entity after ingest and on every search query.

## List ingest (OFAC)

When Watchman loads the OFAC SDN file it applies two name steps that other lists (EU, UK, UN) do not.

**Name reordering** (people only, `SDNType=individual`):

- `MADURO MOROS, Nicolas` → `Nicolas MADURO MOROS`

**Company suffix stripping** (companies):

- `SAI ADVISORS INC.` → `SAI ADVISORS`
- Suffixes include `INC.`, `LLC`, `LTD.`, `GMBH`, `CO.`, and similar.

A query is **not** reordered or suffix-stripped. You can search `Nicolas Maduro` or `MADURO, Nicolas`; after `Normalize()` both become comparable to the OFAC person name that was already reordered at ingest.

## Normalize (every list record and every query)

`Entity.Normalize()` fills `PreparedFields` used by scoring and indexes.

### Names

1. Trim, lowercase, turn `.` `,` `-` and other punctuation/symbols into spaces.
2. Unicode NFD → strip combining marks → NFC, so `Raúl` and `raul` match.
3. Split on whitespace into tokens.
4. Drop **stopwords** (`of`, `the`, `and`, and the same idea in other languages). Language is guessed from the name. Numeric tokens (`11420`) are kept. Set `KEEP_STOPWORDS=true` to skip this step.

The same steps run on the primary name, aliases, and former names (`HistoricalInfo` of type `Former Name`).

Example: `BANK OF AMERICA` → tokens `bank`, `america`.

### Phones and fax

Digits are kept and, when possible, normalized with a country guess (`+1 (555) 010-0100` and `555-010-0100` can compare).

### Addresses

1. Free-text `address=` on the query is parsed into line, city, state, postal code, country. **Docker images** use libpostal. **GitHub binaries / `go run`** use usaddress. **Any build** can enable deepparse. See [Addresses](/watchman/addresses/).
2. Parsed fields are lowercased. Commas are stripped from street lines. Country codes (`US`, `USA`) fold to a common name (`United States`).
3. Street, city, and similar fields are split into tokens for scoring.

### Gender (query parameter)

`gender=` is mapped to `male`, `female`, or `unknown` (`m` / `male` / `f` / `female`, and a few synonyms).

### Government IDs

Query IDs are `gov_<type>=COUNTRY:IDENTIFIER` (for example `gov_passport=IR:Y53914915`). Country on the ID is normalized the same way as address country when scoring.

## What this means for searches

- Case, accents, and punctuation do not need to match the list file: `José` and `jose` compare as the same letters.
- Common words can be omitted: `Bank of America` and `Bank America`.
- For **OFAC people**, list names in `SURNAME, Given` order are stored as `Given SURNAME`, so a natural-order query works.
- For **OFAC companies**, `Inc.` / `LLC` on the list name is stripped at ingest. A query of `Acme Inc` still has `inc` as a token unless you omit it; it is a weak token, not a blocker.
- Send IDs and dates when you have them. Preparation does not invent them.

## Debugging

`debug=true` on `/v2/search` returns **score pieces** (name, identifiers, dates, addresses), not a dump of these pipeline stages. If a name does not match, check type (`person` vs `business`), whether you sent the same fields as the list record, and the prepared tokens implied above.

## Related

[Using Watchman](/watchman/using-watchman/) · [Search](/watchman/search/) · [Similarity methodology](/watchman/methodology/) · [Indexing](/watchman/indexing/)
