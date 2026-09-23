---
layout: page
title: Record linkage keys
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Record linkage keys

Watchman can emit **composite blocking keys** for an entity: hierarchical, hashed tokens that bucket similar records without writing names, ID numbers, street addresses, emails, or phone numbers into the index.

Use them to shrink a candidate set before Jaro-Winkler scoring — in SQL with prefix filters, or in Watchman's in-memory corpus (government IDs and address-only queries). They are not a substitute for similarity scoring.

## Why hashed segments

Record-linkage blocking classically stores values such as `US|CA|90210|anytown`. That is useful for prefix scans and disastrous if the index is a database column: it is PII.

Watchman hashes each field independently (truncated SHA-256 with a field-specific domain string) and concatenates the digests from coarse to fine:

```
TYPE:<8 hex>
NAME:<8 hex>                          // hashed Soundex of each significant name token
GOVID:C<hex>|T<hex>|X<hex>            // country | ID type | identifier
ADDR:C<hex>|S<hex>|P<hex>|Y<hex>|L<hex>[|E<hex>]
CONTACT:E<hex>                        // email
CONTACT:P<hex>                        // phone
IMO:<hex>   MMSI:<hex>   AIR:<hex>
```

`C`, `S`, `P`, `Y`, `L`, `E`, `T`, `X` are segment tags, not data. The hex is 8 lowercase characters (4 bytes of SHA-256). Collisions only add extra candidates; scoring still decides the match.

The same input always produces the same digest. `US` and `United States` normalize to the same country before hashing. Passport `AA-111` and `AA111` hash equal.

This is **not encryption**. Anyone who can guess the input can recompute the digest. The goal is that a dumped index does not contain the original PII.

## Prefix filters

`internal/linksim.Prefixes` walks a key at `|` boundaries, coarse to fine:

```
ADDR:Caaaaaaa1|Sbbbbbbb2|Pcccccccc|Ydddddddd|Leeeeeeee
```

yields:

1. `ADDR:Caaaaaaa1` — same country
2. `ADDR:Caaaaaaa1|Sbbbbbbb2` — same country + state
3. `…|Pcccccccc` — + postal
4. `…|Ydddddddd` — + city
5. `…|Leeeeeeee` — + line1

Two people at `541 First St` and `541 First St Apt 301` share the prefix through line1; the optional line2 (`E`) segment differs. Strip right-hand segments to widen the block.

SQL (illustrative — Watchman does not require this schema):

```sql
-- Same hashed country + state + postal, any city or street
SELECT entity_id FROM blocking_keys
 WHERE key LIKE 'ADDR:Caaaaaaa1|Sbbbbbbb2|Pcccccccc%';
```

Do not range-scan the hex (`key > x AND key < y`) expecting similar records. Similarity lives in **shared prefixes**, not in numeric order of the digest.

## Name keys

Name keys are **hashed Soundex codes** of prepared name tokens (stopwords already removed), not hashes of the raw name. `Smith` and `Smythe` share a `NAME:` key; `John` and `Jon` share another. The Soundex code itself is not stored.

Name search in Watchman still uses the token inverted index (see [Indexing](/watchman/indexing/)). `NAME:` keys are for external blockers and for phonetic grouping in SQL.

## In-memory search

On each list refresh the corpus indexes every key and every prefix. Candidate selection then:

| Query shape | Blocking behavior |
|-------------|-------------------|
| Government ID present | Exact `GOVID:` lookup (same idea as crypto addresses). No hits → fall back to the source/type partition so recall is preserved. |
| IMO / MMSI / aircraft serial / email / phone | Exact `IMO:` / `MMSI:` / `AIR:` / `CONTACT:` lookup. No hits → fall back to the partition. |
| Address, no name tokens | Finest `ADDR:` prefix that still prunes the partition. Too broad or empty → fall back. |
| Name tokens | Name-token inverted index (intersect from the rarest token; disjoint tokens fall back to union). |

Call `linksim.Keys(entity)` after `entity.Normalize()`.

## Go

```go
import "github.com/moov-io/watchman/internal/linksim"

keys := linksim.Keys(entity.Normalize())
for _, key := range keys {
    for _, prefix := range linksim.Prefixes(key) {
        _ = prefix // store / probe
    }
}
```
