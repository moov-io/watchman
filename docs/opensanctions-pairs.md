---
layout: page
title: OpenSanctions Pairs evaluation
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# OpenSanctions Pairs evaluation

Watchman was scored against [OpenSanctions Pairs](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs) (Smith, Sesodia, Lindenberg, Schroeder de Witt, 2026; arXiv:2603.11051): 755,540 analyst-labeled entity pairs from 293 sanctions and OSINT sources.

The full method, per-schema tables, configuration matrix, and paper comparison live in the repo at [`research/opensanctions-pairs/RESULTS.md`](https://github.com/moov-io/watchman/blob/master/research/opensanctions-pairs/RESULTS.md). The evaluator is `go run ./research/opensanctions-pairs`. This page is the short version.

## What was measured

The paper is **pairwise** matching (same real-world entity or not). Watchman is a **ranked screener**. The eval calls production `pkg/search.Similarity` on every labeled pair. It does not measure candidate blocking or coverage of Watchman's loaded OFAC/EU/UK/UN lists.

Predicted positive means `score >= threshold`. Default cutoff is 0.80 (`minMatch`). Occupancy/Succession auto-merge rows (100% positive in the paper) are dropped in the **subjects** slice (472,477 people, companies, vessels).

## Screening pick

**Jaro-Winkler + cross-script embeddings (`qwen3-embedding:0.6b`, hybrid / `EMBEDDINGS_CROSS_SCRIPT_ONLY`), `minMatch=0.80`.**

On subjects after scoring-policy updates:

| Configuration (threshold 0.80) | Precision | Recall | F1 score | Cross-script recall |
|--------------------------------|----------:|-------:|---------:|--------------------:|
| Jaro-Winkler | 0.986 | 0.689 | 0.811 | 0.50 |
| Jaro-Winkler + TF-IDF | 0.968 | 0.707 | 0.817 | 0.52 |
| **Jaro-Winkler + cross-script embeddings** | **0.946** | **0.815** | **0.876** | **0.91** |
| Embeddings on every pair | 0.784 | 0.933 | 0.852 | 0.91 |

Without embeddings, about half of cross-script true matches miss at 0.80. Hybrid costs more review (14k subject FPs vs 3k) and recovers ~38k false negatives. Name-algorithm swaps (Soundex, nsim, Editex, Beider-Morse) do not close that gap.

If missing a hit is worse than extra review, drop hybrid `minMatch` to 0.59 (subjects prec 0.876, rec 0.942). If embeddings are off, use Jaro-Winkler at 0.59 rather than 0.80.

## How Watchman compares to the paper

The paper’s Table 3 uses **sampled** pair sets (about 1k–10k) and **includes** Occupancy/Succession auto-merge rows. Watchman numbers below are the **full 755,540-pair dump** unless noted.

| System | What was measured | F1 score | Precision | Recall |
|--------|-------------------|---------:|----------:|-------:|
| nomenklatura RegressionV1 | Paper sample, threshold 0.15 | 0.913 | 0.845 | 0.994 |
| GPT-4o (0-shot) | Paper sample, binary same-entity label from the full FollowTheMoney JSON | 0.990 | 0.988 | 0.991 |
| Watchman Jaro-Winkler | Full dump, threshold 0.80 | 0.905 | 0.989 | 0.834 |
| Watchman Jaro-Winkler | Full dump, threshold 0.59 | 0.960 | 0.960 | 0.960 |

RegressionV1 is built to catch almost every merge and accept a large review queue. Watchman at 0.80 does the reverse: almost every hit is real, and more labeled positives stay below the cutoff. At 0.59, Watchman’s F1 score on the full dump is 0.960, with higher precision than RegressionV1.

GPT-4o is a different task. The model reads the entire record pair and emits yes/no. Watchman returns a ranked score from names, IDs, dates, and addresses for screening.

Vessels (IMO/MMSI) reach an F1 score of 0.993. Person/Person is where most remaining misses sit under Jaro-Winkler (recall 0.68 at 0.80), mainly transliteration.

## Related scoring changes

Work from this evaluation also tightened production scoring (separate PRs): unique identity keys (passport, IMO, crypto) still force 1.0; tax/registration IDs are evidence only; conflicting same-type IDs apply `ID_CONFLICT_PENALTY_MULTIPLIER`; person/business/organization type mismatches recast instead of scoring 0.

See [For compliance and risk](/watchman/methodology/for-compliance/), [Search](/watchman/search/), [Performance](/watchman/performance/), [Cross-script matching](/watchman/cross-script-matching/), and [Configuration](/watchman/config/).
