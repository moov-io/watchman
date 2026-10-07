# Cascade embedding model comparison

Local Ollama models scored against the Cascade name-pair fixtures. Pairwise `pkg/search.Similarity` plus cosine on names, mixed as production embed-hybrid: max(Jaro–Winkler, cosine) only when the two names differ in script. Latin/Latin pairs stay on Jaro–Winkler. This is not `/v2/search`.

Machine: Apple M4 Max. Baseline models on local Ollama (`:11434`). EmbeddingGemma 2 needs Ollama ≥ 0.36. First pass used a 0.40 sidecar; later confirmed on Ollama.app 0.40.0. `qwen3-embedding:latest` is the 7.6B Q4_K_M tag (4096-d). EmbeddingGemma 2 Ollama tags are nvfp4; `/api/embed` returns 768-d L2-normalized vectors (Ollama `show` lists embedding length 512, which is the backbone width before the 512→768 projection).

`embeddinggemma-2:270m` (text-only, 378MB) and `embeddinggemma-2:latest` (740m with vision/audio, 1.3GB) produced identical text embeddings on smoke pairs. Watchman only needs the 270m tag.

```
# existing models
go run ./research/cascade-embeddings \
  -csv pkg/search/testdata/cascade-name-pairs-2.csv \
  -models qwen3-embedding:0.6b,qwen3-embedding:4b,qwen3-embedding:latest,bge-m3,nomic-embed-text

# EmbeddingGemma 2 (Ollama >= 0.36; text-only 270m tag)
go run ./research/cascade-embeddings \
  -csv pkg/search/testdata/cascade-name-pairs-2.csv \
  -models embeddinggemma-2:270m

# STS prefix (model-card SentenceSimilarity; over-fires at 0.80)
go run ./research/cascade-embeddings \
  -csv pkg/search/testdata/cascade-name-pairs-2.csv \
  -models embeddinggemma-2:270m \
  -prefix "task: sentence similarity | query: "
```

351 unique lot-2 names: Gemma 2 270m in 3.3s, qwen3 0.6b in 5.3s, qwen3 8B in 44.7s, Gemma 1 in 14.5s.

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
| embeddinggemma (v1 300m) hybrid | 768 | 0.6990 | 0.5885 | 0.6447 | 0.6357 | 0.6537 | 42 | 3 | 0.933 | 0.356 | 5 | 0.066 |
| **embeddinggemma-2:270m hybrid** | **768** | **0.8054** | **0.7705** | **0.8099** | **0.7964** | **0.8234** | **82** | **25** | **0.766** | **0.695** | **45** | **0.592** |
| embeddinggemma-2:270m hybrid + STS prefix | 768 | 0.8650 | 0.8440 | 0.9024 | 0.9044 | 0.9004 | 113 | 57 | 0.665 | 0.958 | 76 | 1.000 |

At 0.59 (high-recall line):

| Model | TP @ 0.59 | FP @ 0.59 | Cross-script TP @ 0.59 | Cross-script recall @ 0.59 |
|---|---:|---:|---:|---:|
| jaro-winkler | 39 | 8 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 54 | 18 | 15 | 0.197 |
| qwen3-embedding:4b hybrid | 88 | 40 | 49 | 0.645 |
| qwen3-embedding:latest (8B) hybrid | 113 | 56 | 74 | 0.974 |
| bge-m3 hybrid | 73 | 15 | 34 | 0.447 |
| nomic-embed-text hybrid | 59 | 17 | 20 | 0.263 |
| embeddinggemma (v1 300m) hybrid | 93 | 33 | 54 | 0.711 |
| embeddinggemma-2:270m hybrid | 115 | 62 | 76 | 1.000 |
| embeddinggemma-2:270m hybrid + STS prefix | 115 | 62 | 76 | 1.000 |

EmbeddingGemma 1 stays in the qwen3 0.6b band: +5 true pairs at 0.80, zero extra false positives.

EmbeddingGemma 2 (raw names, no task prefix) is a different band. At minMatch=0.80 it recovers 45 of 76 cross-script true pairs (qwen3 8B recovered 15) and adds 22 false positives on the designed near-miss set (25 vs 3 for 0.6b / 7 for 8B). Burmese/Latin mean 0.823 vs 0.689 on 8B and 0.464 on 0.6b. Hebrew/Latin mean 0.796 vs 0.760 / 0.572.

The extra false positives are mostly Burmese near misses (other-word, place, subsidiary/parent, close syllable, numeral, sister ship). Hebrew near misses stay cleaner (about 5 cosine FPs vs 17 Burmese).

The model-card STS prefix (`task: sentence similarity | query: `) is the wrong setup for this cutoff. It puts true pairs and near misses both around 0.90, so hybrid at 0.80 fires on all 76 cross-script true pairs and 57 false positives. Watchman should send raw names, matching the previous bake-off and the production embed client.

At 0.59, Gemma 2 raw already has XS recall 1.0 and 62 FP, in the same collapse as 8B at 0.59 (74 XS TP / 56 FP). Do not drop the cutoff to chase leftover Hebrew/Burmese.

MTEB (multilingual, v2) mean 61.36 is a weak proxy here. Gemma 2 sits below Qwen3-8B and Harrier 0.6B on that board and still beats every local Ollama model we have on this Hebrew/Burmese name set.

## Lot 1 (300 pairs, no cross-script negatives)

140 matches / 160 non-matches. All 23 cross-script rows are labelled matches, so hybrid cannot add Latin false positives. Non-match mean stays 0.6124.

| Model | Dim | Match mean | Non-match mean | Cross-script match mean | Hebrew/Latin match mean | Burmese/Latin match mean | TP @ 0.80 | FP @ 0.80 | Precision @ 0.80 | Recall @ 0.80 | Cross-script TP @ 0.80 | Cross-script recall @ 0.80 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| jaro-winkler | 0 | 0.6149 | 0.6124 | 0.0000 | 0.0000 | 0.0000 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |
| qwen3-embedding:0.6b hybrid | 1024 | 0.7142 | 0.6124 | 0.6047 | 0.5055 | 0.4218 | 72 | 25 | 0.742 | 0.514 | 3 | 0.130 |
| qwen3-embedding:4b hybrid | 2560 | 0.7364 | 0.6124 | 0.7398 | 0.7790 | 0.5254 | 75 | 25 | 0.750 | 0.536 | 6 | 0.261 |
| qwen3-embedding:latest (8B) hybrid | 4096 | 0.7378 | 0.6124 | 0.7482 | 0.7650 | 0.6733 | 77 | 25 | 0.755 | 0.550 | 8 | 0.348 |
| bge-m3 hybrid | 1024 | 0.7095 | 0.6124 | 0.5760 | 0.5796 | 0.5344 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |
| nomic-embed-text hybrid | 768 | 0.7017 | 0.6124 | 0.5283 | 0.5184 | 0.5424 | 69 | 25 | 0.734 | 0.493 | 0 | 0.000 |
| embeddinggemma-2:270m hybrid | 768 | 0.7508 | 0.6124 | 0.8270 | 0.7938 | 0.8367 | 84 | 25 | 0.771 | 0.600 | 15 | 0.652 |

At 0.59 Gemma 2 recovers all 23 lot-1 cross-script true pairs (TP 109 / FP 73, same FP as JW).

## What this says for Watchman

The documented pick (`qwen3-embedding:0.6b`, CROSS_SCRIPT_ONLY, minMatch=0.80) is still the FP-stable local model.

EmbeddingGemma 2 is the first local Ollama model we ran that actually moves the leftover Hebrew/Burmese true pairs at 0.80 (45 XS TP vs 4 for 0.6b and 15 for 8B). It does it by lifting near misses too (25 FP vs 3). That is the same trade as larger Qwen3, with more recall and more FP, at 270M / 768-d / ~3s for 351 names.

Do not use the STS prefix in Watchman. Do not switch the documented default on this run alone. Gemma 2 is the strongest local candidate so far if Hebrew/Burmese recall at 0.80 is worth ~22 extra lot-2 false positives.

Apache 2.0 (v1 used the Gemma license). Card: [EmbeddingGemma 2](https://ai.google.dev/gemma/docs/embeddinggemma/model_card_2), [google/embeddinggemma-2](https://huggingface.co/google/embeddinggemma-2), Ollama `embeddinggemma-2`.
