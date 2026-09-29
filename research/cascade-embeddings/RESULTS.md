# Cascade embedding model comparison

Local Ollama models scored against the Cascade name-pair fixtures. Pairwise `pkg/search.Similarity` plus cosine on raw names, mixed as production embed-hybrid: max(Jaro–Winkler, cosine) only when the two names differ in script. Latin/Latin pairs stay on Jaro–Winkler. This is not `/v2/search`.

Machine: Apple M4 Max, Ollama. `qwen3-embedding:latest` here is the 7.6B Q4_K_M tag (4096-d).

```
go run ./research/cascade-embeddings \
  -csv pkg/search/testdata/cascade-name-pairs-2.csv \
  -models qwen3-embedding:0.6b,qwen3-embedding:4b,qwen3-embedding:latest,bge-m3,nomic-embed-text
```

## Lot 2 (185 pairs, the Hebrew/Burmese/former-name yardstick)

118 matches / 67 non-matches. 130 pairs are cross-script (76 true, 54 near-miss non-matches). Former-name Latin/Latin rows already sit on Jaro–Winkler after the lot-2 prep PR.

| Model | Dim | Match mean | Non-match mean | Cross-script match mean | Hebrew/Latin match mean | Burmese/Latin match mean | TP @ 0.80 | FP @ 0.80 | Precision @ 0.80 | Recall @ 0.80 | Cross-script TP @ 0.80 | Cross-script recall @ 0.80 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| jaro-winkler | 0 | 0.2846 | 0.1344 | 0.0012 | 0.0025 | 0.0000 | 37 | 3 | 0.925 | 0.314 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 1024 | 0.6173 | 0.5303 | 0.5179 | 0.5720 | 0.4637 | 41 | 3 | 0.932 | 0.347 | 4 | 0.053 |
| qwen3-embedding:4b hybrid | 2560 | 0.7121 | 0.6491 | 0.6650 | 0.7686 | 0.5614 | 51 | 10 | 0.836 | 0.432 | 14 | 0.184 |
| qwen3-embedding:latest (8B) hybrid | 4096 | 0.7504 | 0.6817 | 0.7244 | 0.7599 | 0.6889 | 52 | 7 | 0.881 | 0.441 | 15 | 0.197 |
| bge-m3 hybrid | 1024 | 0.6512 | 0.5450 | 0.5705 | 0.5618 | 0.5792 | 38 | 3 | 0.927 | 0.322 | 1 | 0.013 |
| nomic-embed-text hybrid | 768 | 0.6352 | 0.5595 | 0.5455 | 0.5819 | 0.5092 | 37 | 3 | 0.925 | 0.314 | 0 | 0.000 |

At 0.59 (high-recall line):

| Model | TP @ 0.59 | FP @ 0.59 | Cross-script TP @ 0.59 | Cross-script recall @ 0.59 |
|---|---:|---:|---:|---:|
| jaro-winkler | 39 | 8 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 54 | 18 | 15 | 0.197 |
| qwen3-embedding:4b hybrid | 88 | 40 | 49 | 0.645 |
| qwen3-embedding:latest (8B) hybrid | 113 | 56 | 74 | 0.974 |
| bge-m3 hybrid | 73 | 15 | 34 | 0.447 |
| nomic-embed-text hybrid | 59 | 17 | 20 | 0.263 |

Lot 2 was written so a fix that lifts true Hebrew/Burmese pairs can be checked against lifting the near misses. Larger Qwen3 models do both. Non-match mean rises with model size because 54 of the 67 negatives are themselves cross-script. At minMatch=0.80 the 8B model still only adds 4 false positives (7 vs 3) and still only recovers 15 of 76 cross-script true pairs. Dropping the cutoff to 0.59 recovers almost all of those true pairs on 8B and also 56 false positives.

`bge-m3` and `nomic-embed-text` do not clear 0.80 on this Hebrew/Burmese set. `nomic-embed-text` is an English-centric control.

## Lot 1 (300 pairs, no cross-script negatives)

140 matches / 160 non-matches. All 23 cross-script rows are labelled matches, so hybrid cannot add Latin false positives. Non-match mean stays 0.6124 for every model.

| Model | Dim | Match mean | Non-match mean | Cross-script match mean | Hebrew/Latin match mean | Burmese/Latin match mean | TP @ 0.80 | FP @ 0.80 | Precision @ 0.80 | Recall @ 0.80 | Cross-script TP @ 0.80 | Cross-script recall @ 0.80 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| jaro-winkler | 0 | 0.6149 | 0.6124 | 0.0000 | 0.0000 | 0.0000 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 1024 | 0.7142 | 0.6124 | 0.6047 | 0.5055 | 0.4218 | 72 | 25 | 0.742 | 0.514 | 3 | 0.130 |
| qwen3-embedding:4b hybrid | 2560 | 0.7364 | 0.6124 | 0.7398 | 0.7790 | 0.5254 | 75 | 25 | 0.750 | 0.536 | 6 | 0.261 |
| qwen3-embedding:latest (8B) hybrid | 4096 | 0.7378 | 0.6124 | 0.7482 | 0.7650 | 0.6733 | 77 | 25 | 0.755 | 0.550 | 8 | 0.348 |
| bge-m3 hybrid | 1024 | 0.7095 | 0.6124 | 0.5760 | 0.5796 | 0.5344 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |
| nomic-embed-text hybrid | 768 | 0.7017 | 0.6124 | 0.5283 | 0.5184 | 0.5424 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |

At 0.59:

| Model | TP @ 0.59 | FP @ 0.59 | Cross-script TP @ 0.59 | Cross-script recall @ 0.59 |
|---|---:|---:|---:|---:|
| jaro-winkler | 86 | 73 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 98 | 73 | 12 | 0.522 |
| qwen3-embedding:4b hybrid | 105 | 73 | 19 | 0.826 |
| qwen3-embedding:latest (8B) hybrid | 107 | 73 | 21 | 0.913 |
| bge-m3 hybrid | 96 | 73 | 10 | 0.435 |
| nomic-embed-text hybrid | 90 | 73 | 4 | 0.174 |

## What this says for Watchman

The documented pick (`qwen3-embedding:0.6b`, CROSS_SCRIPT_ONLY, minMatch=0.80) is the FP-stable local model. On lot 2 it adds 4 true pairs at 0.80 and zero extra false positives.

`qwen3-embedding` 8B is the best local model we ran for Hebrew and Burmese true pairs. At 0.80 it is 15 cross-script true positives vs 4 for 0.6b, with 4 extra false positives on the near-miss set. 4B sits between them on recall and is the sloppiest at 0.80 (10 FP). `bge-m3` is not a replacement on this script mix.

No local model we ran separates Hebrew/Burmese true pairs from the designed near misses well enough to use a 0.59 cutoff. That cutoff on 8B is 0.97 cross-script recall and 56 false positives.
