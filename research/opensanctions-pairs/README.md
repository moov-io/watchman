# OpenSanctions Pairs × Watchman scoring

Research tool that scores [OpenSanctions Pairs](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs)
(Smith, Sesodia, Lindenberg, Schroeder de Witt, 2026) with Watchman's similarity function and, optionally, a
running Watchman search.

The paper is a pairwise entity-matching benchmark: 755,540 analyst-labeled pairs (`positive` / `negative`) over
FollowTheMoney records from 293 sources. Watchman is a screening engine, but the same `pkg/search.Similarity`
path used at query time can score left vs right directly. That is the default mode here.

Paper and code:

- Dataset: https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs
- Raw dump: https://data.opensanctions.org/contrib/training/pairs-20251209.json.gz
- Reproduction code: https://github.com/chansmi/OSINT_entity_resolution
- Watchman scores **name, IDs, dates, addresses, contact, and programs** — not the paper's LLM matcher.

Full write-up of method, per-slice scores, paper comparison, and Watchman features: [RESULTS.md](RESULTS.md).

## Quick start

From the Watchman repo root:

```
go test ./research/opensanctions-pairs

go run ./research/opensanctions-pairs \
  -download sample \
  -mode pairwise \
  -threshold 0.80
```

That downloads the 1,000-pair stratified sample (`seed=42`, 769 positive / 231 negative) and scores each pair
in-process. Occupancy / Position / other auto-merge schemas are skipped by default (`-subjects-only`), matching
the paper's 66.4% analyst-judged subset.

Full corpus (~390MB gzip, 755,540 pairs):

```
go run ./research/opensanctions-pairs \
  -download full \
  -mode pairwise \
  -subjects-only=false \
  -report research/opensanctions-pairs/out/full-report.json
```

Results from that full run (2026-09-24, type fan-out for `LegalEntity`): **F1 0.913 at threshold 0.80**, **best F1 0.962 at 0.59** on all 755,540 pairs; **F1 0.831 / 0.936** on the 472,477 analyst-judged subject pairs. Configuration sweep (TF-IDF, name algorithms, Ollama embeddings) is in [Configuration comparison](#configuration-comparison) (that matrix was scored before type fan-out; Jaro-Winkler all-pairs F1 there is 0.910).

Sweep every scorer Watchman can turn on:

```
# Ollama must be running. Default model is qwen3-embedding:0.6b (pull if needed).
ollama pull qwen3-embedding:0.6b

go run ./research/opensanctions-pairs \
  -input research/opensanctions-pairs/data/pairs-20251209.json.gz \
  -subjects-only=false \
  -compare default \
  -embed-model qwen3-embedding:0.6b \
  -report research/opensanctions-pairs/out/compare-report.json
```

`-compare default` runs Jaro-Winkler ± TF-IDF, phonetic boosts, n-gram token scorers, Beider-Morse, and three embedding mixes. `-tfidf` / `-embed hybrid|max|only` apply one option on a single pairwise run.

## Modes

### pairwise (default)

Maps each FollowTheMoney side to a Watchman `Entity`, calls `search.SimilarityWithOpts` after `Normalize()`,
and treats `score >= threshold` as a predicted positive. Optional `-tfidf` builds an IDF table from this
corpus's prepared name tokens (same path production uses after `index.Lists.Update`). Optional `-embed`
mixes Ollama cosine similarity on primary names.

### search

Sends the left entity to a running Watchman (`GET /v2/search`) and looks for the right-hand record in the
result page (OFAC/UK/EU/UN native IDs extracted from `id` / `referents`, then caption). The pair score is
that hit's `match`, or 0 if the right entity is not on the lists Watchman has loaded.

```
go run ./research/opensanctions-pairs \
  -input research/opensanctions-pairs/data/sample_1000.json \
  -mode search \
  -watchman http://localhost:8084 \
  -search-limit 10
```

Coverage is reported separately: many pairs are Wikidata PEPs, corporate registries, or other sources Watchman
does not ingest.

## Flags

| Flag | Default | Meaning |
|------|---------|---------|
| `-input` | | JSON array, `{metadata,pairs}` sample file, JSONL, or `.json.gz` |
| `-download` | | `sample` or `full` into `-data-dir` |
| `-mode` | `pairwise` | `pairwise` or `search` |
| `-algorithm` | `jaro-winkler` | Same values as `?algorithm=` on `/v2/search` |
| `-tfidf` | `false` | Build a TF-IDF index from this corpus and weight name tokens |
| `-embed` | | `hybrid` (cross-script max), `max` (always max), or `only` (cosine, no Similarity) |
| `-embed-model` | `qwen3-embedding:0.6b` | Ollama embedding model |
| `-embed-url` | `http://127.0.0.1:11434` | Ollama base URL |
| `-compare` | | `default` or a comma-separated list of configs (`jaro-winkler+tfidf,embed-hybrid`, …) |
| `-threshold` | `0.80` | Positive cutoff |
| `-sweep` | `true` | F1 across 0.50–0.99 |
| `-subjects-only` | `true` | Drop Occupancy/Position/auto-merge pairs |
| `-schemas` | | e.g. `person` or `company` (org-like schemas group together) |
| `-name-only` | `false` | Names only; isolates the string matcher |
| `-limit` / `-offset` | | Slice after filtering |
| `-out` | | Per-pair JSONL |
| `-report` | | JSON metrics summary |
| `-debug-errors` | `5` | Highest-score false positives and lowest-score false negatives, with score pieces |

Single-algorithm or TF-IDF-only runs:

```
go run ./research/opensanctions-pairs -input ... -algorithm nsim
go run ./research/opensanctions-pairs -input ... -algorithm jaro-winkler -tfidf
go run ./research/opensanctions-pairs -input ... -embed hybrid
```

## Full corpus results

Run 2026-09-24 against `pairs-20251209.json.gz` (755,540 pairs, 581,149 positive / 174,391 negative — the paper's published counts). Pairwise `Similarity` with default Jaro-Winkler, cap of 20 alt names, no TF-IDF corpus weights. Occupancy / Position / Succession and other relational schemas are included in "all pairs" and dropped in "subjects".

The paper’s Table 3 uses 1k–10k stratified samples (auto-merge included). RegressionV1: F1 score 0.913, precision 0.845, recall 0.994. GPT-4o: F1 score 0.990. Watchman numbers in this file are the full 755,540-pair dump.

### Headline metrics

| Slice | n | pos / neg | threshold | Acc | Prec | Rec | F1 | TP | FP | FN |
|-------|--:|----------:|----------:|----:|-----:|----:|---:|---:|---:|---:|
| All pairs | 755,540 | 581,149 / 174,391 | 0.80 | 0.875 | 0.983 | 0.853 | **0.913** | 495,783 | 8,813 | 85,366 |
| All pairs, best F1 | 755,540 | 581,149 / 174,391 | **0.59** | 0.942 | 0.962 | 0.962 | **0.962** | 558,930 | 21,977 | 22,219 |
| Analyst-judged subjects | 472,477 | 303,189 / 169,288 | 0.80 | 0.811 | 0.972 | 0.725 | **0.831** | 219,860 | 6,229 | 83,329 |
| Subjects, best F1 | 472,477 | 303,189 / 169,288 | **0.59** | 0.918 | 0.940 | 0.933 | **0.936** | 282,863 | 18,213 | 20,326 |
| Cross-script | 127,829 | 94,224 / 33,605 | 0.80 | 0.648 | 0.988 | 0.530 | **0.689** | 49,887 | 627 | 44,337 |

Score distribution at the 0.80 cutoff:

| Label | mean | median |
|-------|-----:|-------:|
| All pairs, positive | 0.849 | 0.855 |
| All pairs, negative | 0.360 | 0.305 |
| Subjects, positive | 0.845 | 0.928 |
| Subjects, negative | 0.351 | 0.301 |

The all-pairs F1 is pulled up by Occupancy (213,448), Succession (33,607), Family, Ownership, and similar auto-merge rows, which the paper marks 100% positive. Both sides often share the schema name as caption (`Occupancy`), so Watchman scores them ~0.855 and predicts positive at 0.80. **Subjects-only is the number that describes Watchman screening of people, companies, and vessels.**

### Threshold sweep (all 755,540 pairs)

| threshold | acc | f1 | prec | rec | fp | fn |
|----------:|----:|---:|-----:|----:|--:|--:|
| 0.50 | 0.927 | 0.953 | 0.943 | 0.964 | 33,965 | 21,010 |
| 0.55 | 0.934 | 0.957 | 0.954 | 0.959 | 26,644 | 23,548 |
| 0.60 | 0.926 | 0.952 | 0.964 | 0.939 | 20,418 | 35,174 |
| 0.65 | 0.925 | 0.950 | 0.972 | 0.929 | 15,406 | 41,395 |
| 0.70 | 0.915 | 0.943 | 0.978 | 0.910 | 11,631 | 52,438 |
| 0.75 | 0.876 | 0.914 | 0.981 | 0.856 | 9,667 | 83,750 |
| 0.80 | 0.871 | 0.910 | 0.982 | 0.847 | 8,783 | 88,750 |
| 0.85 | 0.851 | 0.894 | 0.984 | 0.819 | 7,682 | 105,104 |
| 0.90 | 0.488 | 0.506 | 0.978 | 0.341 | 4,366 | 382,796 |
| 0.95 | 0.422 | 0.404 | 0.976 | 0.254 | 3,595 | 433,249 |

A cliff at 0.90 matches `exactMatchThreshold` (0.99) / `highConfidenceThreshold` (0.95) in `pkg/search/similarity.go`: occupancy-style ~0.855 scores and many name-only positives fall out of the predicted-positive set.

### By schema (threshold 0.80)

Largest subject slices:

| schema | n | acc | f1 | prec | rec | pos mean | neg mean |
|--------|--:|----:|---:|-----:|----:|---------:|---------:|
| Person/Person | 284,808 | 0.765 | 0.806 | 0.983 | 0.683 | 0.860 | 0.362 |
| Company/Company | 70,885 | 0.849 | 0.816 | 0.882 | 0.759 | 0.800 | 0.341 |
| Organization/Organization | 35,429 | 0.921 | 0.909 | 0.980 | 0.848 | 0.868 | 0.378 |
| Organization/Company | 25,849 | 0.930 | 0.919 | 0.991 | 0.857 | 0.891 | 0.308 |
| Company/Organization | 21,953 | 0.907 | 0.906 | 0.995 | 0.831 | 0.885 | 0.291 |
| Vessel/Vessel | 7,550 | 0.990 | 0.993 | 0.998 | 0.988 | 0.991 | 0.286 |
| Occupancy/Occupancy | 213,448 | 1.000 | 1.000 | 1.000 | 1.000 | 0.855 | — |
| Succession/Succession | 33,607 | 1.000 | 1.000 | 1.000 | 1.000 | 0.855 | — |
| Position/Position | 26,330 | 0.898 | 0.946 | 0.913 | 0.980 | 0.888 | 0.811 |
| LegalEntity/Person | 3,969 | 0.814 | 0.891 | 0.986 | 0.813 | 0.835 | 0.545 |
| Person/LegalEntity | 466 | 0.639 | 0.749 | 0.936 | 0.623 | 0.769 | 0.641 |

FollowTheMoney `LegalEntity` is treated as an **unknown Watchman type**. Production `Similarity` now recasts person/business/organization onto the index type instead of returning 0 (vessel/aircraft stays a hard zero). Typed `/v2/search?type=person` is unchanged: candidates still come from that partition only. A client with an unknown query type should still issue `type=person`, `type=business`, …; empty `type=` already selects the all-types partition. The pairwise eval fans out shared types and takes the max. `LegalEntity` vs `Person` therefore scores as `person`. `Person` vs `Company` stays 0 in the eval (no shared type, Similarity is not called).

### Error modes

**False positives at 1.0** (measured before the scorer change below) were almost all shared identifiers forcing `exactOverride`:

- Related companies with the same INN / registration / address (Aforra Development vs Aforra Property; Mayak vs Vulkan; Massandra winery branches).
- Two `Aung Moe Myint` person records labeled negative despite identical name, dates, address, and IDs — possible label noise.
- One pair where `left.id == right.id` (`REVIVAL OF ISLAMIC HERITAGE SOCIETY`) labeled negative.

**False negatives at 0.0** are empty-stub right-hand records (caption is the schema name `Company` / `Person` and `properties.name` is missing), or two **known** Watchman types that do not share a partition (`Person` vs `Company`). `LegalEntity` vs `Person` is no longer a structural zero; see type fan-out above.

### 1,000-pair sample

The paper's seed-42 stratified sample (769 pos / 231 neg), subjects-only (626 pairs after dropping auto-merge), same scorer:

| | n | threshold | Acc | Prec | Rec | F1 |
|--|--:|----------:|----:|-----:|----:|---:|
| Subjects | 626 | 0.80 | 0.808 | 0.973 | 0.719 | 0.827 |
| Subjects, best F1 | 626 | 0.58 | 0.923 | 0.944 | 0.935 | 0.939 |

Tracks the full subjects-only table (0.823 / 0.929) within a couple of points.

## Screening recommendation

Re-run 2026-09-24 on the full 755,540-pair dump **after** exact-ID tightening, tax-as-evidence, ID-conflict penalty, and person/org recast. Slice that matters for screening: **472,477 analyst-judged subjects** (people, companies, vessels). Occupancy/Succession auto-merge rows are dropped.

**Pick: Jaro-Winkler + `embed-hybrid` (qwen3-embedding:0.6b, `CrossScriptOnly` analog), `minMatch=0.80`.**

That is the only config that closes the transliteration hole without flooding review:

| Configuration (threshold 0.80, subjects) | Precision | Recall | F1 score | False positives | False negatives | Cross-script recall |
|------------------------------------------|----------:|-------:|---------:|----------------:|----------------:|--------------------:|
| Jaro-Winkler | **0.986** | 0.689 | 0.811 | 2,955 | 94,286 | 0.499 |
| Jaro-Winkler + TF-IDF | 0.968 | 0.707 | 0.817 | 7,129 | 88,785 | 0.516 |
| Soundex / Double Metaphone / Beider-Morse | 0.986 | 0.691 | 0.813 | ~3,000 | ~93,500 | 0.503 |
| **Jaro-Winkler + embeddings for different writing systems** | 0.946 | **0.815** | **0.876** | 14,014 | **56,072** | **0.905** |
| Same embeddings + TF-IDF | 0.935 | 0.829 | 0.879 | 17,437 | 51,697 | 0.910 |
| Embeddings on every pair | 0.784 | 0.933 | 0.852 | 78,141 | 20,179 | 0.905 |

Without embeddings, half of cross-script true matches miss at 0.80. Hybrid trades ~11k extra subject FPs for ~38k fewer FNs and takes cross-script recall from 0.50 to 0.91. Precision stays 0.95.

Do **not** use `embed-max` or `embed-only` for screening: ~78k subject FPs. Name-algorithm swaps (nsim, editex, soft-bidist) do not move the needle. TF-IDF alone is a small recall bump paid for in precision.

If missing a hit is worse than extra review, drop `minMatch` to **0.59** on the same hybrid scorer (subjects prec 0.876, rec 0.942, F1 0.908). If embeddings are unavailable, use **jaro-winkler at 0.59** (subjects prec 0.945, rec 0.920, F1 0.932) rather than 0.80.

Production mapping: `algorithm` default (jaro-winkler), embeddings enabled with `crossScriptOnly`, `minMatch=0.80`. Typed searches stay `type=person` / `type=business`; unknown query type still needs the client fan-out.

## Configuration comparison

Same dump, first matrix 2026-09-24 before ID-override/conflict/recast (see screening table above for the later run). TF-IDF is built from this corpus (name + alt-name token documents). Embeddings are **qwen3-embedding:0.6b** via local Ollama (`/api/embed`, 1024-d, 368,359 unique primary names cached). The 8B `qwen3-embedding:latest` already on the machine is the same series at higher quality; it was too slow for a full unique-name pass (~1 embedding/s vs ~70–130/s for 0.6b).

Embedding mixes:

- `embed-hybrid` — Watchman `Similarity`, then `max(sim, cosine)` when the pair is cross-script (production `CrossScriptOnly` analog)
- `embed-max` — `max(sim, cosine)` on every pair
- `embed-only` — cosine of primary names, IDs/dates/addresses ignored

### All 755,540 pairs @ 0.80

| config | acc | f1 | prec | rec | best F1 | best thr | FP | FN |
|--------|----:|---:|-----:|----:|--------:|---------:|--:|--:|
| jaro-winkler | 0.871 | 0.910 | **0.982** | 0.847 | 0.959 | 0.59 | 8,783 | 88,749 |
| jaro-winkler+tfidf | 0.870 | 0.910 | 0.973 | 0.854 | 0.949 | 0.59 | 13,529 | 84,764 |
| soundex | 0.872 | 0.911 | 0.982 | 0.849 | 0.959 | 0.59 | 8,893 | 88,021 |
| soundex+tfidf | 0.870 | 0.910 | 0.973 | 0.855 | 0.949 | 0.59 | 13,863 | 84,245 |
| double-metaphone | 0.872 | 0.911 | 0.982 | 0.849 | 0.959 | 0.59 | 8,885 | 87,997 |
| double-metaphone+tfidf | 0.870 | 0.910 | 0.973 | 0.855 | 0.949 | 0.59 | 13,866 | 84,243 |
| soft-bidist | 0.870 | 0.910 | 0.983 | 0.847 | 0.958 | 0.59 | 8,715 | 89,142 |
| nsim | 0.868 | 0.908 | 0.983 | 0.844 | **0.960** | 0.59 | 8,560 | 90,843 |
| editex | 0.869 | 0.909 | 0.983 | 0.845 | 0.960 | 0.59 | 8,608 | 90,232 |
| beider-morse | 0.871 | 0.910 | 0.982 | 0.848 | 0.959 | 0.59 | 8,845 | 88,281 |
| **embed-hybrid** | **0.906** | **0.937** | 0.964 | 0.912 | 0.945 | 0.59 | 19,860 | **51,383** |
| embed-max | 0.866 | 0.917 | 0.870 | **0.970** | 0.933 | 0.90 | 83,862 | 17,521 |
| embed-only | 0.846 | 0.904 | 0.871 | 0.940 | 0.907 | 0.87 | 81,024 | 35,084 |

### Subjects (472,477) and cross-script (127,829) @ 0.80

| config | subjects F1 | subj prec | subj rec | subj best F1 | cross F1 | cross rec | cross prec |
|--------|------------:|----------:|---------:|-------------:|---------:|----------:|-----------:|
| jaro-winkler | 0.823 | 0.972 | 0.714 | **0.929** | 0.689 | 0.529 | **0.988** |
| jaro-winkler+tfidf | 0.826 | 0.955 | 0.727 | 0.913 | 0.698 | 0.544 | 0.971 |
| soundex | 0.825 | 0.972 | 0.716 | 0.930 | 0.692 | 0.533 | 0.987 |
| nsim | 0.819 | 0.973 | 0.707 | 0.932 | 0.683 | 0.522 | 0.988 |
| **embed-hybrid** | **0.884** | 0.936 | **0.837** | 0.904 | **0.903** | **0.926** | 0.882 |
| embed-max | 0.857 | 0.781 | 0.948 | 0.879 | 0.903 | 0.926 | 0.882 |
| embed-only | 0.829 | 0.776 | 0.890 | 0.831 | 0.878 | 0.879 | 0.877 |

### What moved

**TF-IDF** is a small recall bump and a precision cost. On all pairs, FN 88,749 → 84,764 and FP 8,783 → 13,529; F1 stays 0.910. Negative mean rises 0.360 → 0.422: common tokens (`limited`, company legal forms) lose weight, so some related-but-distinct companies score higher and some true matches with rare tokens recover. Best-F1 at 0.59 drops 0.959 → 0.949. Useful as a production feature for ranking, not a headline F1 win on this labeled set.

**Phonetic boosts** (soundex, double-metaphone, beider-morse) are within 0.001 F1 of Jaro-Winkler. They recover ~700–750 extra true positives and add ~100 extra FPs. Token-scorer swaps (soft-bidist, nsim, editex) trade a little recall for a little precision; nsim has the highest tuned F1 (0.960 at 0.59) and the worst 0.80 recall. None of these close the cross-script gap (recall stays ~0.52–0.55).

**Embeddings are the only lever that moves cross-script recall.** Hybrid lifts cross-script recall 0.529 → **0.926** (FN 44,338 → 6,972) and subjects F1 0.823 → **0.884**. All-pairs F1 goes 0.910 → **0.937**. Precision falls 0.982 → 0.964 (FP 8,783 → 19,860), mostly extra cross-script false positives (627 → 11,704). That is the paper's LLM-shaped trade: catch transliterations, accept more lookalikes.

`embed-max` / `embed-only` over-fire on Latin pairs (FP ~81–84k). Cosine on primary names scores Occupancy stubs and similar company names very high, and it cannot use the ID short-circuit. Keep embeddings behind the cross-script gate (`hybrid` / production `CrossScriptOnly`).

## What the scores mean

Watchman is high-precision / moderate-recall at the 0.80 screening-style cutoff, and closest to RegressionV1's F1 when the threshold is swept to ~0.59. GPT-4o's 99 F1 on the paper's pairwise task is a different problem (binary same-entity with the full FtM record in context).

With configurations flexed on the same dump: name-algorithm choice barely moves F1; TF-IDF is a modest recall/precision slide; **cross-script embeddings are the only change that recovers tens of thousands of true matches**. Hybrid at 0.80 (F1 0.937 all-pairs, 0.884 subjects) is the best default Watchman can show on this benchmark without giving up ID-aware scoring.

Recurring mechanics:

- Passport, IMO, and crypto still short-circuit to 1.0. Tax ID / business registration are weighted evidence only and never force 1.0. When both records have the same ID type and country with different values, the score is multiplied by `ID_CONFLICT_PENALTY_MULTIPLIER` (default 0.70).
- Conflicting IDs withhold that bonus; they do not veto a strong name match.
- Jaro-Winkler on tokens is weak on the 127,829 cross-script pairs (recall 0.53 at 0.80); `embed-hybrid` takes that recall to 0.93.
- Caption-only stubs are structural zeros. Known-type mismatches (`Person` vs `Company`) stay 0. Unknown/`LegalEntity` queries fan out across `person` and `business` the way a Watchman client would.

## Mapping notes

FollowTheMoney → Watchman (best-effort, not a full FtM importer):

- `Person` → `person`; `Company` / `Organization` / `PublicBody` → `business`
- `LegalEntity` (and other unknown schemas) → fan-out `person` + `business` (plus `vessel`/`aircraft` when IMO/serial is present), best score wins — same pattern as multiple `/v2/search?type=` calls when the query type is unknown
- `Vessel` / `Airplane` → `vessel` / `aircraft`
- Caption is the primary name; `name`, `alias`, `weakAlias`, `previousName` become alt names (capped)
- `passportNumber`, `idNumber`, `innCode` / `taxNumber`, `registrationNumber` / `ogrnCode` / `leiCode` → government IDs
- `birthDate` / `incorporationDate`, `address`, `email` / `phone` / `website`, `programId`

`SourceID` is left empty in pairwise mode so two records of the same person from different lists are not
collapsed by the exact SourceID short-circuit.

## License

The pairs corpus is CC-BY-NC 4.0 (OpenSanctions). This tool is Apache 2.0 with the rest of Watchman. Do not
commit the downloaded dataset.
