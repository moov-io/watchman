---
layout: page
title: Cross-Script Name Matching
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Cross-Script Name Matching

Watchman can match names across different writing systems — Arabic, Cyrillic, Chinese, etc. — using neural embeddings. So if someone searches for "محمد علي", we can find "Mohamed Ali" in the OFAC list.

## Why do we need this?

Jaro-Winkler and other string algorithms compare characters. Different scripts = different characters = no match:

```
"محمد علي" vs "Mohamed Ali" → Jaro-Winkler says 0%
```

But they're the same name.

## How it works

We use a neural network (via API) that converts text into vectors. The key insight: similar names get similar vectors, regardless of script.

```
"Mohamed Ali"  → [0.12, -0.45, 0.78, ...]
"محمد علي"     → [0.11, -0.44, 0.79, ...]  ← almost identical!
```

Then we just compare vectors with cosine similarity. Done.

### Hybrid approach

We don't use embeddings for everything — that would be slow. Instead:

- **Non-Latin query** (Arabic, Cyrillic, etc.) → use embeddings
- **Latin query** → use Jaro-Winkler (faster, works great for Latin)

Set `crossScriptOnly: true` (the default) to get this behavior.

## Supported Providers

Watchman supports any OpenAI-compatible embeddings API:

| Provider | Base URL | Notes |
|----------|----------|-------|
| **Ollama** (local) | `http://localhost:11434/v1` | Free, runs locally |
| **OpenAI** | `https://api.openai.com/v1` | Best quality, paid |
| **OpenRouter** | `https://openrouter.ai/api/v1` | Many models, paid |
| **Azure OpenAI** | `https://{resource}.openai.azure.com/...` | Enterprise |

## Setup

### Choose a provider

**Option A: Ollama (local, open-source models)**

```bash
# Install Ollama
curl -fsSL https://ollama.com/install.sh | sh

# Pull the model
ollama pull qwen3-embedding
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

Choose based on your accuracy requirements:

| Model                                                             | Provider            | Dimension | Cross-script Quality | Notes                                    |
|-------------------------------------------------------------------|---------------------|-----------|----------------------|------------------------------------------|
| [Qwen3 Embedding](https://huggingface.co/Qwen/Qwen3-Embedding-8B) | Ollama & OpenRouter | 4096      | Best                 | Open source router, easy to run.         |
| `text-embedding-3-small`                                          | OpenAI              | 1536      | Best                 | Recommended for production               |
| `text-embedding-3-large`                                          | OpenAI              | 3072      | Best                 | Higher accuracy, slower                  |
| `multilingual-e5-large`                                           | Ollama              | 1024      | Good                 | Best open-source option                  |
| `nomic-embed-text`                                                | Ollama              | 768       | Limited              | General-purpose, not optimized for names |

## API

Nothing special — just search as usual. Embeddings kick in automatically for non-Latin queries:

```bash
$ curl -s "http://localhost:8084/v2/search?type=person&limit=1&name=Владимир+Путин+PUTIN" | jq -r '.entities[] | .name,.match'
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
