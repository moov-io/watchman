---
layout: page
title: Cross-Script Name Matching
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

**Option A: Ollama (free, local)**
```bash
# Install Ollama
curl -fsSL https://ollama.com/install.sh | sh

# Pull the model
ollama pull nomic-embed-text
```

#### Example Models

Models with [embedding support on Ollama](https://ollama.com/search?c=embedding&o=newest).

| Source                                                                              | Ollama                                                                                       | OpenRouter                                                                  |
|-------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| [HuggingFace](https://huggingface.co/sentence-transformers/paraphrase-MiniLM-L6-v2) | [Link](https://ollama.com/koill/sentence-transformers:paraphrase-multilingual-minilm-l12-v2) | [Link](https://openrouter.ai/sentence-transformers/paraphrase-minilm-l6-v2) |

**Option B: OpenAI (paid, best quality)**
```bash
export EMBEDDINGS_API_KEY=ollama
```

### Run

```bash
export EMBEDDINGS_ENABLED=true
# For Ollama (default):
export EMBEDDINGS_BASE_URL=http://localhost:11434/v1
export EMBEDDINGS_MODEL=nomic-embed-text

# For OpenAI:
# export EMBEDDINGS_BASE_URL=https://api.openai.com/v1
# export EMBEDDINGS_MODEL=text-embedding-3-small
# export EMBEDDINGS_DIMENSION=1536

./watchman
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

Cross-script name matching quality varies significantly between models. Choose based on your accuracy requirements:

| Model | Provider | Dimension | Cross-script Quality | Notes |
|-------|----------|-----------|---------------------|-------|
| `text-embedding-3-small` | OpenAI | 1536 | Best | Recommended for production |
| `text-embedding-3-large` | OpenAI | 3072 | Best | Higher accuracy, slower |
| `multilingual-e5-large` | Ollama | 1024 | Good | Best open-source option |
| `nomic-embed-text` | Ollama | 768 | Limited | General-purpose, not optimized for names |

## API

Nothing special — just search as usual. Embeddings kick in automatically for non-Latin queries:

```bash
curl -X POST http://localhost:8084/v2/search \
  -H "Content-Type: application/json" \
  -d '{"name": "محمد علي", "type": "person"}'
```

## Testing

```bash
# Unit tests
go test ./internal/embeddings/...

# Integration tests (needs running provider like Ollama)
go test -tags integration ./internal/embeddings/... -v
```

## Known limitations

- First query is slower (API round-trip + model warm-up)
- Very short names (1-2 chars) don't work well
- Quality depends heavily on the model used
- Some rare scripts may have lower accuracy
