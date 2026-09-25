---
layout: page
title: Cross-Script Name Matching
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Cross-Script Name Matching

Watchman can match names across writing systems (Arabic, Cyrillic, Chinese, and others) using neural embeddings. A query of `محمد علي` can return `Mohamed Ali` on the OFAC list.

Jaro-Winkler compares characters. Different scripts share almost no characters, so that pair scores near 0 even when they are the same name.

Embeddings map each name to a numeric vector. The same name in two scripts lands near itself in that space. Watchman ranks those vectors with cosine similarity.

```
"Mohamed Ali"  → [0.12, -0.45, 0.78, ...]
"محمد علي"     → [0.11, -0.44, 0.79, ...]
```

On [OpenSanctions Pairs](/watchman/opensanctions-pairs/), this is what moves cross-script recall from about 0.50 to 0.91 at `minMatch=0.80`. Name-algorithm swaps (Soundex, and others) do not.

### Hybrid approach

Embeddings are slower than Jaro-Winkler, so Watchman does not use them on every query. With `crossScriptOnly: true` (the default):

- **Non-Latin query** (Arabic, Cyrillic, Chinese, …) also searches the vector index
- **Latin query** stays on Jaro-Winkler

Using embeddings on every pair (embed-max / embed-only) over-fires on Latin names. Keep the cross-script gate on.

The screening pick measured in this repo is Jaro-Winkler plus hybrid embeddings (`qwen3-embedding:0.6b` via Ollama) at `minMatch=0.80`. Embeddings are **off** until you enable them.

## Supported Providers

Watchman supports any OpenAI-compatible embeddings API:

| Provider                                                                                | Base URL                                  | Notes              |
|-----------------------------------------------------------------------------------------|-------------------------------------------|--------------------|
| [**Chutes**](https://chutes.ai/app?type=embedding)                                      | `https://{model}.chutes.ai/v1`            | Many models, paid  |
| [**Ollama**](https://ollama.com/search?c=embedding) (local)                             | `http://localhost:11434/v1`               | Free, runs locally |
| [**OpenAI**](https://developers.openai.com/api/docs/guides/embeddings#embedding-models) | `https://api.openai.com/v1`               | High quality, paid |
| [**OpenRouter**](https://openrouter.ai/models?fmt=cards&output_modalities=embeddings)   | `https://openrouter.ai/api/v1`            | Many models, paid  |

## Setup

### Choose a provider

**Option A: Ollama (local, open-source models)**

Install or [Download Ollama](https://ollama.com/download)
```bash
curl -fsSL https://ollama.com/install.sh | sh
```

Pull the model used in the OpenSanctions evaluation:
```
ollama pull qwen3-embedding:0.6b
```

**Option B: OpenAI (paid, best quality)**
```bash
  Search:
    # Tune these settings based on your available resources (CPUs, etc).
    # Usually a multiple (i.e. 2x, 4x) of GOMAXPROCS is optimal.
    Goroutines:
      Default: 10
      Min: 1
      Max: 25
    Embeddings:
      Enabled: true # Opt-in feature
      Provider:
        Name: "openrouter"                      # ollama, openai, openrouter, azure
        BaseURL: "https://openrouter.ai/api/v1" # API endpoint (required when enabled)
        APIKey: "<api-key>"                     # Can be set via EMBEDDINGS_API_KEY env var
        Model: "qwen/qwen3-embedding-8b"        # Required: e.g., "text-embedding-3-small" (OpenAI)
        Dimension: 4096                         # Required: must match model (e.g., 1536 for OpenAI, 1024 for e5-large)
        NormalizeVectors: true                  # L2 normalize if API doesn't
        Timeout: "10s"
        RateLimit:
          RequestsPerSecond: 100
          Burst: 75
        Retry:
          MaxRetries: 3
          InitialBackoff: "1s"
          MaxBackoff: "30s"
      Cache:
        # Cache type can be one of Blank (disabled), memory, sql
        Type: "sql"
      CrossScriptOnly: true # Hybrid approach: embeddings for cross-script only
      SimilarityThreshold: 0.70
      BatchSize: 32
      IndexBuildTimeout: "10m"
```

## Configuration

| Env Variable | Default | What it does |
|--------------|---------|--------------|
| `EMBEDDINGS_ENABLED` | `false` | Turn on/off |
| `EMBEDDINGS_BASE_URL` | — | API endpoint (required) |
| `EMBEDDINGS_API_KEY` | — | API key (optional for Ollama) |
| `EMBEDDINGS_MODEL` | — | Model name (required) |
| `EMBEDDINGS_DIMENSION` | — | Vector dimension (required, must match model) |
| `EMBEDDINGS_CROSS_SCRIPT_ONLY` | `true` | Only use for non-Latin queries |
| `EMBEDDINGS_SIMILARITY_THRESHOLD` | `0.7` | Min score to return a match |
| `EMBEDDINGS_CACHE_SIZE` | `10000` | How many vectors to cache |

### Recommended models

Cross-script name matching quality varies significantly between models. Models with [embedding support on Ollama](https://ollama.com/search?c=embedding&o=newest).

The OpenSanctions subject numbers in this repo used `qwen3-embedding:0.6b` (1024-d). Larger models can be higher quality and are slower and more expensive.

| Model                                                             | Provider            | Dimension | Notes |
|-------------------------------------------------------------------|---------------------|-----------|-------|
| `qwen3-embedding:0.6b`                                            | Ollama              | 1024      | Screening pick in [OpenSanctions Pairs](/watchman/opensanctions-pairs/) |
| [Qwen3 Embedding 8B](https://huggingface.co/Qwen/Qwen3-Embedding-8B) | Ollama & OpenRouter | 4096      | Larger local/router model |
| `text-embedding-3-small`                                          | OpenAI              | 1536      | Hosted |
| `text-embedding-3-large`                                          | OpenAI              | 3072      | Hosted, slower |
| `multilingual-e5-large`                                           | Ollama              | 1024      | Open multilingual model |
| `nomic-embed-text`                                                | Ollama              | 768       | General-purpose, weaker on names |

## API

Search is still `GET /v2/search`. With embeddings enabled, a non-Latin `name` uses the vector index automatically:

```bash
curl -s --get "http://localhost:8084/v2/search" \
  --data-urlencode "type=person" \
  --data-urlencode "name=Владимир Путин" \
  --data-urlencode "limit=1" \
  | jq '{name: .entities[0].name, match: .entities[0].match}'
```
```
Vladimir Vladimirovich PUTIN
0.949172991083859
```

## Known limitations

- First query is slower (API round-trip + model warm-up)
- Very short names (1-2 chars) don't work well
- Quality depends heavily on the model used
- Some rare scripts may have lower accuracy
