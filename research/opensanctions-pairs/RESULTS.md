# Watchman on OpenSanctions Pairs

Results of scoring the [OpenSanctions Pairs](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs) benchmark (Smith, Sesodia, Lindenberg, Schroeder de Witt, 2026; arXiv:2603.11051) with [Moov Watchman](https://github.com/moov-io/watchman). Runs are from 2026-09-24 against snapshot `pairs-20251209.json.gz` (755,540 labeled pairs). Code lives in this directory; the scorer is production `pkg/search.Similarity`.

This document covers Watchman's scoring method, how the benchmark was mapped onto it, results by slice of the corpus, comparison with the paper's baselines, and what Watchman exposes for customization.

## 1. What this evaluation is

OpenSanctions Pairs is a **pairwise entity-matching** task: given two FollowTheMoney records, decide whether they are the same real-world entity. Labels are `positive` or `negative`, produced by OpenSanctions analysts in production deduplication (not crowd workers). The corpus is 755,540 pairs over 1,002,093 entity records from 293 sources in 45 jurisdictions; 76.9% of pairs are positive.

Watchman is a **sanctions screening engine**. A client sends one query entity; Watchman scores it against an in-memory watchlist (OFAC, EU, UK, UN, and optional ingest files) and returns ranked hits with a match score in `[0, 1]`.

The two problems share a scorer. They do not share a pipeline. This eval calls `Similarity(left, right)` on every labeled pair. It does **not** measure candidate blocking, list coverage, or live `/v2/search` latency against a loaded OFAC/EU/UK/UN index. A pair that Watchman would never retrieve in production can still receive a high pairwise score here, and a pair Watchman would retrieve can still score low.

Predicted positive means `score >= threshold`. Default screening-style cutoff is **0.80** (`minMatch`). We also sweep 0.50–0.99 and report the F1-maximizing threshold.

## 2. The paper, in brief

**OpenSanctions Pairs: A Large-Scale Dataset for Pairwise Entity Matching.** Analyst-labeled pairs, multi-script, sparse attributes (median 5 populated fields). 64.7% of pairs and 82.0% of positives are cross-source.

Table 2 of the paper splits the corpus:

| Tier | Pairs | Share | Pos / Neg |
|------|------:|------:|-----------|
| Analyst-judged subjects (Person, Company/Org/LegalEntity, Position, Vessel, some cross-schema) | 501,298 | 66.4% | 328k / 173k |
| Auto-merge (Occupancy, Succession, Ownership/Family/Directorship, 100% positive by construction) | 254,242 | 33.6% | 254k / 0 |
| **Total** | **755,540** | 100% | **581,149 / 174,391** |

Our "subjects" slice is 472,477 pairs: analyst-judged people, companies, organizations, legal entities, and vessels. We also drop Position (26,330), which the paper kept in the analyst-judged block. Occupancy/Succession/Family/Ownership inflate absolute F1 if left in, because both captions are often the schema word (`Occupancy`) and score ~0.855.

Paper Table 3 (label-stratified samples, seed 42; RegressionV1 threshold 0.15; auto-merge included):

| Method | Acc | F1 | Prec | Rec |
|--------|----:|---:|-----:|----:|
| nomenklatura RegressionV1 | 85.45 | 91.33 | 84.46 | 99.42 |
| Llama-3.1-8B 0-shot | 90.40 | 94.05 | 89.84 | 98.67 |
| DeepSeek-R1-Distill-Qwen-14B 0-shot | 96.53 | 97.76 | 96.24 | 99.33 |
| GPT-4o 0-shot | 98.38 | 98.95 | 98.78 | 99.11 |
| Constant-positive (paper §3.4) | 76.9 | 87.0 | — | 100 |

RegressionV1 is a logistic model over 18 features (token overlap, Levenshtein, phonetic, dates, identifier overlap, geography). It is high-recall / low-precision, matching OpenSanctions' preference to catch every merge and spend analyst time on false positives. LLMs fail on cross-script transliteration; the rule matcher over-fires on common names without unique IDs (Figure 1: two Khalid Mehmoods on Pakistan's proscribed list, RegressionV1 score 0.98).

The paper's samples are 10,000 and 1,000 pairs. We scored the **entire 755,540-pair dump**. Numbers are therefore not a line-for-line replica of Table 3; they are the same labels, full population.

## 3. What Watchman is

Watchman is an open-source Go service (Apache 2.0) that downloads sanctions lists, indexes them in memory, and serves ranked search.

**Lists.** OFAC SDN, US Consolidated Screening List, EU, UK, UN, FinCEN 311, plus CSV ingest of custom files. Refresh is periodic or on demand (`POST /v2/data/refresh`).

**API.** `GET /v2/search` with a structured entity (name, type, dates, government IDs, addresses, crypto, contact). Optional Senzing JSON / JSONL. WASM admin UI. MCP tools.

**Index.** Corpus partitioned by source × type. Candidate selection uses name-token inverted index (distinctive-token intersection), hashed government-ID and address blocks, crypto exact keys, and prefix/QWERTY-near indexes on IMO, MMSI, aircraft serial, email, and phone. Empty type under a known source with no entities of that type returns nothing; it does not scan other types. Empty `type=` selects the all-types partition for that source.

**Scoring.** Every candidate is compared to the query with `Similarity`. Score is a weighted blend of field groups, then coverage penalties, then optional exact-ID override or ID-conflict penalty.

**Runtime knobs.** Per-request `algorithm`, `minMatch`, `limit`, `debug`. Process-wide TF-IDF, embeddings provider, Jaro-Winkler penalties, goroutine counts, ingest schemas. See §8.

## 4. How Watchman scores a pair

`Normalize()` prepares names (lowercase, punctuation, stopwords), phones, and addresses. `Similarity` then builds nine score pieces:

| Piece | Weight | What it compares |
|-------|-------:|------------------|
| Exact identifiers | 50 | Type+country+identifier on passports, tax IDs, IMO/MMSI, serials |
| Crypto addresses | 50 | Currency + address |
| Government IDs (loose) | 50 | Identifier match with optional country (0.7–1.0) |
| Contact | 50 | Email, phone, fax |
| Name | 35 | Best pairwise token alignment (default Jaro-Winkler), including alt and former names |
| Titles | 35 | Person titles / positions |
| Dates | 15 | Birth, death, incorporation, dissolution (tolerant of fudged days) |
| Addresses | 25 | Structured address fields |
| Supporting | 15 | Sanctions programs, historical names |

Final score is a coverage-aware weighted average (`FINAL_SCORE_*` multipliers). Empty query fields are not compared.

**Exact override (after this work).** A piece forces 1.0 only when it is **Exact** (identifier *and* country) **and** a unique identity key: passport, national ID, SSN-like IDs, IMO/MMSI, aircraft serial, crypto. Tax ID, business registration, commercial registry, email, and phone never override; they stay high-weight evidence. Call-sign-only vessel matches do not override.

**ID conflict.** If both records populate the same ID type (and country, when both set) with different values, the blended score is multiplied by `ID_CONFLICT_PENALTY_MULTIPLIER` (default 0.70). Missing IDs on one side are not a conflict. This is the Khalid Mehmood case: identical names, distinct national IDs.

**Type recast.** Person, business, and organization can be projected onto each other so a FollowTheMoney `LegalEntity` encoded as `business` still scores against a `person`. Vessel and aircraft mismatches remain 0. Typed `/v2/search?type=person` is unaffected: candidates already come from the person partition.

**Name algorithms** (`?algorithm=`):

| Value | Role |
|-------|------|
| `jaro-winkler` (default) | Token pairwise Jaro-Winkler |
| `soundex`, `double-metaphone`, `beider-morse` | Same alignment, phonetic boost when encodings overlap |
| `soft-bidist`, `soft-bisim`, `editex`, `nsim`, `nsim-3` | Replace the token-pair metric; BestPairs / length / first-letter filters stay |

**TF-IDF.** Optional. IDF is built at list refresh; rare tokens outweigh `Limited` / `ООО` / `GmbH`. Query weights computed once per search.

**Embeddings.** Optional neural name vectors (Ollama, OpenAI, OpenRouter, Azure). Default `EMBEDDINGS_CROSS_SCRIPT_ONLY=true`: Latin queries stay on Jaro-Winkler; non-Latin queries also search the vector index. This eval's `embed-hybrid` is the pairwise analog: `max(Similarity, cosine)` when the pair is cross-script.

## 5. How we mapped the benchmark

FollowTheMoney JSON → Watchman `Entity`:

| FtM | Watchman |
|-----|----------|
| Person | `person` |
| Company, Organization, PublicBody | `business` |
| LegalEntity / unknown schema | Fan-out `person` + `business` (plus vessel/aircraft when IMO/serial present); best score. Same idea as multiple `/v2/search?type=` calls when the query type is unknown |
| Vessel, Airplane | `vessel`, `aircraft` |
| Occupancy, Position, Succession, Family, … | Dropped in the subjects slice |

Caption is the primary name; `name` / `alias` / `weakAlias` / `previousName` become alt names (capped at 20). `passportNumber`, `idNumber`, `innCode` / `taxNumber`, `registrationNumber` / `ogrnCode` / `leiCode` become government IDs. Dates, addresses, email/phone/website, `programId` are copied when present. `SourceID` is left empty so two list records of the same person are not collapsed by the exact-SourceID short-circuit.

A pair is predicted positive if `Similarity(left, right) >= threshold`. Cross-script means one side's caption/names are Latin-only and the other's are not.

The paper's 1,000-pair sample (`seed=42`, 769 pos / 231 neg) is also reported after dropping auto-merge (626 pairs).

## 6. Results by slice

Unless noted, scores are Jaro-Winkler, no TF-IDF, no embeddings, threshold 0.80, after type fan-out for `LegalEntity`.

### 6.1 Population counts

| Slice | n | Positive | Negative |
|-------|--:|---------:|---------:|
| Full dump | 755,540 | 581,149 | 174,391 |
| Subjects (this eval) | 472,477 | 303,189 | 169,288 |
| Cross-script | 127,829 | 94,224 | 33,605 |
| Paper 1k sample, subjects | 626 | — | — |

### 6.2 Jaro-Winkler headlines (type fan-out on, before ID-conflict tightening)

| Slice | Thr | Acc | Prec | Rec | F1 | TP | FP | FN |
|-------|----:|----:|-----:|----:|---:|---:|---:|---:|
| All pairs | 0.80 | 0.875 | 0.983 | 0.853 | **0.913** | 495,783 | 8,813 | 85,366 |
| All pairs, best F1 | 0.59 | 0.942 | 0.962 | 0.962 | **0.962** | 558,930 | 21,977 | 22,219 |
| Subjects | 0.80 | 0.811 | 0.972 | 0.725 | **0.831** | 219,860 | 6,229 | 83,329 |
| Subjects, best F1 | 0.59 | 0.918 | 0.940 | 0.933 | **0.936** | 282,863 | 18,213 | 20,326 |
| Cross-script | 0.80 | 0.650 | 0.987 | 0.532 | **0.691** | 50,101 | 656 | 44,123 |
| 1k sample, subjects | 0.80 | 0.808 | 0.973 | 0.719 | 0.827 | — | — | — |
| 1k sample, subjects, best F1 | 0.58 | 0.923 | 0.944 | 0.935 | 0.939 | — | — | — |

Score separation on subjects at 0.80: positive mean/median **0.857 / 0.929**, negative **0.352 / 0.301**. When Watchman is looking at the same Latin-script person with overlapping IDs, it is usually sure. Misses concentrate in transliteration, stubs, and remaining type issues.

After ID-override tightening and the conflict penalty, Jaro-Winkler on the full dump at 0.80 moved to F1 **0.905**, precision **0.989**, FP **5,539** (from 8,813). True positives also dropped (tax IDs no longer force 1.0). That is the scorer used in §7.

### 6.3 By schema (threshold 0.80)

| Schema | n | Acc | F1 | Prec | Rec | Pos mean | Neg mean |
|--------|--:|----:|---:|-----:|----:|---------:|---------:|
| Person/Person | 284,808 | 0.765 | 0.806 | 0.983 | 0.683 | 0.860 | 0.362 |
| Company/Company | 70,885 | 0.849 | 0.816 | 0.882 | 0.759 | 0.800 | 0.341 |
| Organization/Organization | 35,429 | 0.921 | 0.909 | 0.980 | 0.848 | 0.868 | 0.378 |
| Organization/Company | 25,849 | 0.930 | 0.919 | 0.991 | 0.857 | 0.891 | 0.308 |
| Company/Organization | 21,953 | 0.907 | 0.906 | 0.995 | 0.831 | 0.885 | 0.291 |
| Vessel/Vessel | 7,550 | 0.990 | **0.993** | 0.998 | 0.988 | 0.991 | 0.286 |
| LegalEntity/Person | 3,969 | 0.814 | 0.891 | 0.986 | 0.813 | 0.835 | 0.545 |
| Person/LegalEntity | 466 | 0.639 | 0.749 | 0.936 | 0.623 | 0.769 | 0.641 |
| Occupancy/Occupancy | 213,448 | 1.000 | 1.000 | 1.000 | 1.000 | 0.855 | — |
| Succession/Succession | 33,607 | 1.000 | 1.000 | 1.000 | 1.000 | 0.855 | — |
| Position/Position | 26,330 | 0.899 | 0.946 | 0.913 | 0.980 | 0.888 | 0.811 |
| Person/Company | 44 | 1.000 | 0.000 | — | 0.000 | 0.000 | — |
| Organization/Person | 123 | 1.000 | 0.000 | — | 0.000 | 0.000 | — |

Vessels are essentially solved (IMO/MMSI). Organization matching is strong once Company/Organization are coerced to `business`. Person/Person is the bulk of subject false negatives (64,536 of ~83k at 0.80), driven by cross-script names. `LegalEntity` vs `Person` recovered from F1 0.000 to **0.891** after type fan-out. Known-type mismatches (`Person` vs `Company`) stay 0: a client would not search `type=person` against companies.

Occupancy/Succession score ~0.855 because both captions are the schema name. They are 100% positive in the paper and inflate all-pairs F1. **Subjects is the screening number.**

### 6.4 Threshold sweep (all pairs, Jaro-Winkler, type fan-out)

| Thr | Acc | F1 | Prec | Rec | FP | FN |
|----:|----:|---:|-----:|----:|--:|--:|
| 0.50 | 0.932 | 0.956 | 0.942 | 0.971 | 34,467 | 17,082 |
| 0.55 | 0.938 | 0.960 | 0.954 | 0.966 | 27,093 | 19,631 |
| 0.60 | 0.931 | 0.955 | 0.964 | 0.946 | 20,681 | 31,269 |
| 0.70 | 0.920 | 0.946 | 0.979 | 0.916 | 11,695 | 48,736 |
| 0.80 | 0.875 | 0.913 | 0.983 | 0.853 | 8,813 | 85,366 |
| 0.85 | 0.857 | 0.899 | 0.988 | 0.824 | 5,640 | 102,031 |
| 0.90 | 0.489 | 0.508 | 0.979 | 0.343 | 4,354 | 381,641 |

The cliff at 0.90 is occupancy-style ~0.855 scores and name-only positives falling out of the predicted-positive set (`highConfidenceThreshold` 0.95 / `exactMatchThreshold` 0.99).

## 7. Configuration matrix (current scorer)

Full dump, 13 configs, after exact-ID tightening, tax-as-evidence, ID-conflict penalty, and person/org recast. Embeddings: local Ollama `qwen3-embedding:0.6b` (1024-d, 368,359 unique primary names). Each non-embed config scored the 755,540 pairs in ~44s on an M4 Max.

`embed-hybrid` = Similarity, then `max(sim, cosine)` on cross-script pairs (production `CrossScriptOnly`). `embed-max` = always max. `embed-only` = cosine of primary names, IDs ignored.

### 7.1 All 755,540 pairs @ 0.80

| Config | Acc | F1 | Prec | Rec | Best F1 | Best thr | FP | FN |
|--------|----:|---:|-----:|----:|--------:|---------:|--:|--:|
| jaro-winkler | 0.865 | 0.905 | **0.989** | 0.834 | 0.960 | 0.59 | 5,539 | 96,323 |
| jaro-winkler+tfidf | 0.866 | 0.907 | 0.979 | 0.844 | 0.950 | 0.59 | 10,346 | 90,736 |
| soundex | 0.866 | 0.906 | 0.989 | 0.836 | 0.960 | 0.59 | 5,646 | 95,541 |
| double-metaphone | 0.866 | 0.906 | 0.989 | 0.836 | 0.960 | 0.59 | 5,638 | 95,517 |
| beider-morse | 0.866 | 0.905 | 0.989 | 0.835 | 0.960 | 0.59 | 5,599 | 95,824 |
| soft-bidist | 0.864 | 0.904 | 0.989 | 0.833 | 0.960 | 0.59 | 5,471 | 97,132 |
| nsim | 0.861 | 0.902 | 0.989 | 0.829 | **0.961** | 0.59 | 5,313 | 99,612 |
| editex | 0.862 | 0.902 | 0.989 | 0.830 | 0.961 | 0.59 | 5,358 | 98,917 |
| **embed-hybrid** | **0.901** | **0.933** | 0.969 | 0.900 | 0.946 | 0.59 | 16,598 | **58,108** |
| embed-max | 0.863 | 0.915 | 0.873 | **0.962** | 0.924 | 0.90 | 81,331 | 22,049 |
| embed-only | 0.846 | 0.904 | 0.871 | 0.940 | 0.907 | 0.87 | 81,024 | 35,084 |

### 7.2 Subjects (472,477) and cross-script (127,829) @ 0.80

| Config | Subj F1 | Subj prec | Subj rec | Subj best F1 | Cross F1 | Cross rec | Subj FP | Subj FN |
|--------|--------:|----------:|---------:|-------------:|---------:|----------:|--------:|--------:|
| jaro-winkler | 0.811 | **0.986** | 0.689 | **0.932** | 0.663 | 0.499 | 2,955 | 94,286 |
| jaro-winkler+tfidf | 0.817 | 0.968 | 0.707 | 0.915 | 0.674 | 0.516 | 7,129 | 88,785 |
| soundex / dmetaphone | 0.813 | 0.986 | 0.691 | 0.933 | 0.666 | 0.503 | ~3,020 | ~93,560 |
| nsim | 0.804 | 0.987 | 0.678 | 0.934 | 0.656 | 0.491 | 2,788 | 97,517 |
| **embed-hybrid** | **0.876** | 0.946 | **0.815** | 0.908 | **0.892** | **0.905** | 14,014 | **56,072** |
| embed-max | 0.852 | 0.784 | 0.933 | 0.861 | 0.892 | 0.905 | 78,141 | 20,179 |
| embed-only | 0.829 | 0.776 | 0.890 | 0.831 | 0.878 | 0.879 | 77,853 | 33,213 |

At subjects best-F1 (~0.59): Jaro-Winkler prec 0.945 rec 0.920; hybrid prec 0.876 rec 0.942.

### 7.3 What moved

**Name algorithms** are within 0.001–0.003 F1 of Jaro-Winkler. Phonetic boosts recover ~700–800 extra true positives. nsim/editex trade a little recall for a little precision. None of them move cross-script recall off ~0.50.

**TF-IDF** is a small recall bump and a precision cost (subjects FP 2,955 → 7,129). Common legal-form tokens lose weight, so related companies look more alike and some rare-token true matches recover. Useful for ranking, not a headline F1 win here.

**Embeddings are the only lever that moves transliteration.** Hybrid takes cross-script recall from 0.50 to **0.91** (subject FN 94k → 56k) at precision 0.95. That is the paper's complementary failure mode: rules miss Latin/Cyrillic/Arabic pairs; hybrid catches them and accepts more lookalikes.

`embed-max` / `embed-only` over-fire on Latin pairs (~78k subject FPs). Cosine likes Occupancy stubs and similar company names, and it cannot use ID short-circuits. Keep embeddings behind the cross-script gate.

## 8. Comparison with the paper

The paper evaluates a **binary matcher** on pairwise records, including auto-merge. Watchman is a **ranked screener**. Put them on the same axes anyway:

| System | Setting | F1 | Prec | Rec |
|--------|---------|----:|-----:|----:|
| Constant-positive | Full dump, always yes | 0.870 | 0.769 | 1.00 |
| RegressionV1 | Paper sample, thr 0.15 | 0.913 | 0.845 | 0.994 |
| Watchman JW | Full dump @ 0.80 | 0.905 | 0.989 | 0.834 |
| Watchman JW | Full dump, best F1 @ 0.59 | 0.960 | 0.960 | 0.960 |
| Watchman JW | Subjects @ 0.80 | 0.811 | 0.986 | 0.689 |
| Watchman JW | Subjects, best F1 @ 0.59 | 0.932 | 0.945 | 0.920 |
| Watchman embed-hybrid | Subjects @ 0.80 | 0.876 | 0.946 | 0.815 |
| Watchman embed-hybrid | Subjects @ 0.59 | 0.908 | 0.876 | 0.942 |
| GPT-4o 0-shot | Paper sample | 0.990 | 0.988 | 0.991 |

Watchman at 0.80 is the opposite of RegressionV1: very few false positives, more missed positives. At 0.59, Jaro-Winkler F1 on the full dump (0.960) exceeds RegressionV1's 0.913, with much higher precision, still below the LLM ceiling.

That ceiling is a different task. GPT-4o sees the full FollowTheMoney JSON (father name, sparse fields, provenance) and emits a binary label. Watchman emits a score from names, IDs, dates, and addresses. It does not read `fatherName` as a first-class field, does not reason about "two national IDs that cannot both be true," and does not condition on source provenance.

Figure 1 of the paper (two Khalid Mehmoods, same name, different CNICs and fathers): RegressionV1 0.98. Watchman full scoring ~0.35 even before the conflict penalty, because the required identifier piece scores 0; with `ID_CONFLICT_PENALTY_MULTIPLIER=0.70` it drops further. Name-only scoring was ~0.86. The rule matcher over-weights the name; Watchman already treats conflicting national IDs as evidence against a match.

## 9. Error modes

**False positives that were 1.0 before ID tightening** were almost all shared identifiers:

- Related companies with the same INN/registration (Aforra Development vs Aforra Property; Mayak vs Vulkan; Massandra branches). Tax IDs no longer force 1.0.
- Occasional label noise (`left.id == right.id` labeled negative; two `Aung Moe Myint` records with identical fields labeled negative).

**False negatives at 0.0** are:

- Stub records whose caption is the schema name (`Company`, `Person`) and `properties.name` is empty.
- Known-type mismatches (`Person` vs `Company`): no shared Watchman partition.
- Cross-script names under Jaro-Winkler (the large remaining FN pile without embeddings).

**LegalEntity-encoded people** (US Medicaid exclusions and similar) were a structural zero when mapped only to `business`. Type fan-out / Similarity recast recovered LegalEntity/Person to F1 0.891. Recast cost on an M4 Max: same-type person **224 ns/op**; business-query vs person-index **331 ns/op** (+1 alloc, +144 B). Typed `type=person` search does not recast.

## 10. Screening recommendation

Use **Jaro-Winkler + embed-hybrid (`qwen3-embedding:0.6b`, cross-script only), `minMatch=0.80`.**

It is the only config that closes transliteration without flooding review: subjects precision 0.946, recall 0.815, F1 0.876, cross-script recall 0.905.

If missing a designated party is worse than extra review, drop hybrid to **0.59** (subjects prec 0.876, rec 0.942). If embeddings are unavailable, use **Jaro-Winkler at 0.59** (prec 0.945, rec 0.920), not 0.80.

Do not use embed-max/only. Do not expect nsim/soundex/TF-IDF to substitute for embeddings on this corpus.

Production mapping: default algorithm, `EMBEDDINGS_ENABLED=true`, `EMBEDDINGS_CROSS_SCRIPT_ONLY=true`, `minMatch=0.80`, typed queries. Unknown query type: client still issues `type=person` and `type=business` (and vessel/aircraft when those IDs exist).

## 11. Customizability and features

Watchman is built so a bank can turn pieces on without forking the scorer.

**Per request (`GET /v2/search` and MCP):**

- `type` — person, business, organization, vessel, aircraft (partition)
- `algorithm` — nine name metrics listed in §4
- `minMatch`, `limit`, `debug` (field-level pieces)
- Query fields: name, alt names, gender, birth date, titles, government IDs, addresses, crypto, email/phone/website, IMO/MMSI/call sign, aircraft serial/ICAO
- `source` — restrict to one list
- Response `Accept: senzing` or `senzing/jsonl`

**Process-wide (env / YAML):**

| Area | Knobs |
|------|--------|
| TF-IDF | `TFIDF_ENABLED`, smoothing, min/max IDF |
| Embeddings | Provider (Ollama / OpenAI / OpenRouter / Azure), model, dimension, `CROSS_SCRIPT_ONLY`, similarity threshold, cache |
| Exact-ID / conflict | Unique keys still 1.0; tax/registration never override; `ID_CONFLICT_PENALTY_MULTIPLIER` |
| Coverage penalties | `FINAL_SCORE_LOW_COVERAGE_MULTIPLIER`, `FINAL_SCORE_MIN_REQUIRED_FIELDS_MULTIPLIER`, `FINAL_SCORE_NAME_ONLY_MULTIPLIER` |
| Jaro-Winkler | prefix size, boost threshold, length-difference and first-letter penalties, stopwords |
| Phonetic flags | `USE_SOUNDEX_MATCHING`, `SOUNDEX_BOOST_WEIGHT` (overridden by per-request `algorithm`) |
| Concurrency | `SEARCH_MAX_IN_FLIGHT`, `SEARCH_GOROUTINE_COUNT` |
| Lists | which files to download, `INITIAL_DATA_DIRECTORY`, refresh interval |
| Ingest | CSV templates under `Watchman.Ingest.Files`; `POST /v2/ingest/{fileType}` |
| Addresses | Docker images: libpostal in-process; GitHub binaries / `go run`: usaddress. Optional PostalPool or deepparse |
| Geocoding | optional OpenCage |

**What this benchmark actually exercised:** name algorithm, TF-IDF on/off, three embedding mixes, type fan-out, exact-ID policy, ID-conflict penalty. It did not exercise live list ingest, blocking quality, geocoding, or Senzing output.

## 12. Caveats

1. Pairwise `Similarity` is not end-to-end screening. Blocking can drop pairs this eval scores, and Watchman's loaded lists are a subset of the 293 OpenSanctions sources.
2. Paper Table 3 includes auto-merge and uses sampled sets. Our subjects slice is stricter than their 66.4% analyst-judged block (Position dropped).
3. Hybrid embeddings used primary names only (not all aliases) and a 0.6B Qwen3 model. Production embedding search indexes watchlist names, not pairwise captions.
4. Schema captions like `Occupancy` are an artifact of the dump, not of Watchman production data.
5. Some remaining FNs are empty stubs or possible label noise, not ranking errors.

## 13. Reproduce

```
# Sample (paper seed 42)
go test ./research/opensanctions-pairs
go run ./research/opensanctions-pairs -download sample -mode pairwise -threshold 0.80

# Full dump, one scorer
go run ./research/opensanctions-pairs \
  -input research/opensanctions-pairs/data/pairs-20251209.json.gz \
  -subjects-only=false -threshold 0.80

# Config matrix (needs Ollama + qwen3-embedding:0.6b)
ollama pull qwen3-embedding:0.6b
go run ./research/opensanctions-pairs \
  -input research/opensanctions-pairs/data/pairs-20251209.json.gz \
  -subjects-only=false -compare default \
  -embed-model qwen3-embedding:0.6b
```

Dataset license: CC-BY-NC 4.0 (OpenSanctions). Watchman: Apache 2.0. Do not commit the downloaded dump or embedding cache.

## References

- Smith, Sesodia, Lindenberg, Schroeder de Witt. *OpenSanctions Pairs: A Large-Scale Dataset for Pairwise Entity Matching*. arXiv:2603.11051, 2026.
- Dataset: https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs
- Snapshot: https://data.opensanctions.org/contrib/training/pairs-20251209.json.gz
- Paper code: https://github.com/chansmi/OSINT_entity_resolution
- Watchman: https://github.com/moov-io/watchman
