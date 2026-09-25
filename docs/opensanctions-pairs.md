---
layout: page
title: OpenSanctions Pairs evaluation
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# OpenSanctions Pairs evaluation

This page reports how Watchman’s matcher behaved on a public labeled dataset. You do not need this page to run Watchman. It is here so you can see measured precision and recall.

[OpenSanctions Pairs](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs) (Smith, Sesodia, Lindenberg, Schroeder de Witt, 2026; arXiv:2603.11051) is 755,540 pairs of records from sanctions and related lists. Analysts labeled each pair **positive** (same real-world person or company) or **negative**. Full tables: [`RESULTS.md`](https://github.com/moov-io/watchman/blob/master/research/opensanctions-pairs/RESULTS.md). Evaluator: `go run ./research/opensanctions-pairs`.

## What was measured

Watchman is a **screener**: one query in, ranked list records out. The paper is **pairwise**: two records in, same-entity or not. We used Watchman’s production scorer (`Similarity`) on each labeled pair. We did not measure whether Watchman’s search index would have retrieved the pair, or whether both records appear on OFAC/EU/UK/UN.

A pair counts as a Watchman “hit” when `score >= threshold`. The usual cutoff is **0.80** (`minMatch` on `/v2/search`). Some rows in the dataset are occupancy/job records that are always labeled positive; those are omitted below. The remaining **472,477** pairs are people, companies, and vessels.

## Recommended setup

**Jaro-Winkler names, embeddings only when the two names use different writing systems (`EMBEDDINGS_CROSS_SCRIPT_ONLY`), `minMatch=0.80`.** Example embedding model: `qwen3-embedding:0.6b`.

On those 472,477 people, companies, and vessels:

**Precision** is the share of returned hits that are real matches. **Recall** is the share of real matches that were returned. **Cross-script recall** is recall when the two names use different writing systems (for example Latin and Cyrillic).

| Configuration (threshold 0.80) | Precision | Recall | Recall on different writing systems |
|--------------------------------|----------:|-------:|------------------------------------:|
| Jaro-Winkler | 0.986 | 0.689 | 0.50 |
| Jaro-Winkler + TF-IDF | 0.968 | 0.707 | 0.52 |
| **Jaro-Winkler + embeddings for different writing systems** | **0.946** | **0.815** | **0.91** |
| Embeddings on every pair | 0.784 | 0.933 | 0.91 |

Without embeddings, about half of true matches whose names use different writing systems (Latin vs Cyrillic, Arabic, and similar) score below 0.80. Adding embeddings for those pairs returns about 38,000 more true matches and about 11,000 more false hits. Changing the name algorithm (Soundex, nsim, and others) does not close that gap.

If missing a designation is worse than extra review, use embeddings and `minMatch=0.59` (precision 0.876, recall 0.942 on this set). If embeddings are off, use Jaro-Winkler at 0.59 rather than 0.80.

## How Watchman compares to the paper

The paper’s Table 3 uses **sampled** pair sets (about 1k–10k) and **includes** Occupancy/Succession auto-merge rows. Watchman numbers below are the **full 755,540-pair dump** unless noted.

| System | What was measured | Precision | Recall |
|--------|-------------------|----------:|-------:|
| nomenklatura RegressionV1 | Paper sample, threshold 0.15 | 0.845 | 0.994 |
| GPT-4o (0-shot) | Paper sample, yes/no from the full JSON record | 0.988 | 0.991 |
| Watchman Jaro-Winkler | Full dump, threshold 0.80 | 0.989 | 0.834 |
| Watchman Jaro-Winkler | Full dump, threshold 0.59 | 0.960 | 0.960 |

RegressionV1 is built to catch almost every merge and accept a large review queue. Watchman at 0.80 does the reverse: almost every hit is real, and more labeled positives stay below the cutoff. At 0.59, Watchman catches 96% of labeled positives on the full dump, with 96% of those hits real — higher precision than RegressionV1.

GPT-4o is a different task. The model reads the entire record pair and emits yes/no. Watchman returns a ranked score from names, IDs, dates, and addresses for screening.

Vessels (IMO/MMSI) catch 98.8% of labeled positives. Person/Person is where most remaining misses sit under Jaro-Winkler (recall 0.68 at 0.80), mainly transliteration.

## How Watchman scores identifiers

A matching passport, IMO number, or crypto address (with country, when it applies) scores **1.0**. A matching tax number or email raises the score without forcing 1.0. If both records have the same ID type and country but different numbers, the score is multiplied by `ID_CONFLICT_PENALTY_MULTIPLIER` (default 0.70). Person, business, and organization records can still be compared to each other; a person query does not score against a vessel.

See [For compliance and risk](/watchman/methodology/for-compliance/), [Search](/watchman/search/), [Performance](/watchman/performance/), [Cross-script matching](/watchman/cross-script-matching/), and [Configuration](/watchman/config/).
