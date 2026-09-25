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

| Config @ 0.80 | Prec | Rec | F1 | Cross-script rec |
|---------------|-----:|----:|---:|-----------------:|
| Jaro-Winkler | 0.986 | 0.689 | 0.811 | 0.50 |
| Jaro-Winkler + TF-IDF | 0.968 | 0.707 | 0.817 | 0.52 |
| **Embed-hybrid** | **0.946** | **0.815** | **0.876** | **0.91** |
| Embed-max | 0.784 | 0.933 | 0.852 | 0.91 |

Without embeddings, about half of cross-script true matches miss at 0.80. Hybrid costs more review (14k subject FPs vs 3k) and recovers ~38k false negatives. Name-algorithm swaps (Soundex, nsim, Editex, Beider-Morse) do not close that gap.

If missing a hit is worse than extra review, drop hybrid `minMatch` to 0.59 (subjects prec 0.876, rec 0.942). If embeddings are off, use Jaro-Winkler at 0.59 rather than 0.80.

## How Watchman compares to the paper

Paper Table 3 (sampled sets, auto-merge included): nomenklatura RegressionV1 F1 91.3 (prec 84.5, rec 99.4); GPT-4o F1 99.0. Watchman at 0.80 is high-precision / moderate-recall. At 0.59, Jaro-Winkler F1 on the full dump is 0.960 — higher precision than RegressionV1, still below the LLM pairwise ceiling (a different task: full FollowTheMoney JSON in context, binary label).

Vessels are essentially solved (IMO/MMSI, F1 0.993). Person/Person is the bulk of remaining misses under Jaro-Winkler (recall 0.68 at 0.80), mostly transliteration.

## Related scoring changes

Work from this evaluation also tightened production scoring (separate PRs): unique identity keys (passport, IMO, crypto) still force 1.0; tax/registration IDs are evidence only; conflicting same-type IDs apply `ID_CONFLICT_PENALTY_MULTIPLIER`; person/business/organization type mismatches recast instead of scoring 0.

See [Search](/watchman/search/), [Performance](/watchman/performance/), [Cross-script matching](/watchman/cross-script-matching/), and [Configuration](/watchman/config/).
