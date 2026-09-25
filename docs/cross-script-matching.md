---
layout: page
title: Cross-Script Name Matching
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Cross-script name matching

Jaro–Winkler compares characters. A name in Arabic, Cyrillic, or Chinese shares almost no characters with its Latin spelling, so `محمد علي` versus `Mohamed Ali` scores near 0 even when they are the same person.

Watchman can add **neural embeddings**: each name becomes a numeric vector, and similar names land near each other regardless of writing system. Cosine similarity on those vectors is combined with the usual matcher. This feature is **off** until you enable it.

On [OpenSanctions Pairs](/watchman/opensanctions-pairs/), Jaro–Winkler at `minMatch=0.80` had cross-script recall about **0.50**. The hybrid below raised that to **0.91**, with subject precision 0.95 and recall 0.82. Soundex and other name-algorithm flags do not close that gap. See [Using Watchman](/watchman/using-watchman/) for the screening cutoff.

## Hybrid scoring

Embeddings are slower than Jaro–Winkler, so Watchman does not use them on every query. With `EMBEDDINGS_CROSS_SCRIPT_ONLY=true` (the default when embeddings are on):

- A **non-Latin** query (Arabic, Cyrillic, Chinese, and similar) also searches the vector index.
- A **Latin** query stays on Jaro–Winkler.

Using embeddings on every pair (embed-max / embed-only) over-fires on Latin names. Keep the cross-script gate on.

The screening pick measured in this repo is Jaro–Winkler plus hybrid embeddings (`qwen3-embedding:0.6b` via Ollama) at `minMatch=0.80`.

## Setup (Ollama)

This is the configuration used for the OpenSanctions numbers above.

1. Install [Ollama](https://ollama.com/download).
2. Pull the model (1024 dimensions, about 639MB):

```
ollama pull qwen3-embedding:0.6b
```

3. Enable embeddings in Watchman:

```yaml
  Search:
    Embeddings:
      Enabled: true
      Provider:
        Name: "ollama"
        BaseURL: "http://localhost:11434/v1"
        Model: "qwen3-embedding:0.6b"
        Dimension: 1024
        NormalizeVectors: true
        Timeout: "10s"
      Cache:
        Type: "memory"
      CrossScriptOnly: true
      SimilarityThreshold: 0.70
```

`SimilarityThreshold` is the minimum **cosine** similarity for a vector hit. `minMatch` on `/v2/search` is still the Watchman score that decides what is returned. Use `minMatch=0.80` for a production-shaped queue.

Environment variables: [Configuration](/watchman/config/#cross-script-embeddings-configuration). `EMBEDDINGS_API_KEY` is optional for local Ollama.

## Other providers

Watchman calls any OpenAI-compatible embeddings API.

| Provider | Base URL | Notes |
|----------|----------|-------|
| [Ollama](https://ollama.com/search?c=embedding) | `http://localhost:11434/v1` | Local. Use `qwen3-embedding:0.6b` for the measured pick. |
| [OpenAI](https://developers.openai.com/api/docs/guides/embeddings#embedding-models) | `https://api.openai.com/v1` | Hosted. Example: `text-embedding-3-small` (1536-d). |
| [OpenRouter](https://openrouter.ai/models?fmt=cards&output_modalities=embeddings) | `https://openrouter.ai/api/v1` | Hosted router. Set `EMBEDDINGS_API_KEY`. |
| [Chutes](https://chutes.ai/app?type=embedding) | `https://{model}.chutes.ai/v1` | Hosted. |

Larger models (for example Qwen3 Embedding 8B, 4096-d) can be higher quality and are slower and more expensive. Dimension in config must match the model.

## Search

Search is still `GET /v2/search`. With embeddings enabled, a non-Latin `name` uses the vector index. Wait until lists are loaded (`GET /v2/listinfo`).

```
curl -s --get "http://localhost:8084/v2/search" \
  --data-urlencode "type=person" \
  --data-urlencode "name=Владимир Путин" \
  --data-urlencode "minMatch=0.80" \
  --data-urlencode "limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Vladimir Vladimirovich PUTIN","match":0.949}
```

The `match` field is the Watchman score in `[0, 1]`, the same field as a Latin query. It is not a separate cosine percentage.

UTF-8 names belong on the query string (`curl --get --data-urlencode`). There is no JSON POST body on `/v2/search`.

## Limitations

- Embeddings are off in the default `docker run`. Enable them before judging transliteration recall.
- The first query after enable pays model warm-up and an API round-trip.
- Very short names (one or two characters) score poorly.
- Quality depends on the model. The published tables used `qwen3-embedding:0.6b`.
- Rare scripts may recall less than Arabic, Cyrillic, or Chinese.

Full tables: [OpenSanctions Pairs](/watchman/opensanctions-pairs/). Scoring: [Similarity methodology](/watchman/methodology/).
