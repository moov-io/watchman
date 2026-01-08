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

We use a neural network that converts text into 384-dimensional vectors. The key insight: similar names get similar vectors, regardless of script.

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

## How well does it work?

| Script | Example | Score |
|--------|---------|-------|
| Arabic → Latin | محمد علي → Mohamed Ali | 97% |
| Cyrillic → Latin | Владимир Путин → Vladimir Putin | 99.8% |
| Chinese → Latin | 金正恩 → Kim Jong Un | 79% |
| Hebrew → Latin | דוד → David | 96% |

Cyrillic and Arabic work great (transliteration is straightforward). Chinese is harder (characters ≠ sounds), but still useful.

**Speed:** First query ~180ms (model warm-up), then ~5µs from cache. Search over 1000 names takes ~360µs.

## Setup

### 1. Build with embeddings support

```bash
go build -tags embeddings ./cmd/server
```

Without `-tags embeddings`, the feature is compiled out entirely.

### 2. Get the model

We use `paraphrase-multilingual-MiniLM-L12-v2` in ONNX format (~450MB).

```bash
cd tools/export_onnx
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
python export_model.py
```

This creates `models/multilingual-minilm/`.

### 3. Run

```bash
export EMBEDDINGS_ENABLED=true
export EMBEDDINGS_MODEL_PATH=./models/multilingual-minilm
./watchman
```

## Configuration

| Env Variable | Default | What it does |
|--------------|---------|--------------|
| `EMBEDDINGS_ENABLED` | `false` | Turn on/off |
| `EMBEDDINGS_MODEL_PATH` | — | Where's the ONNX model |
| `EMBEDDINGS_CROSS_SCRIPT_ONLY` | `true` | Only use for non-Latin queries |
| `EMBEDDINGS_SIMILARITY_THRESHOLD` | `0.7` | Min score to return a match |
| `EMBEDDINGS_CACHE_SIZE` | `10000` | How many vectors to cache |

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
go test -tags embeddings ./internal/embeddings/...

# Integration tests (needs the model)
go test -tags "embeddings integration" ./internal/embeddings/... -v

# Cross-script e2e
go test -tags "embeddings integration" ./internal/search/... -run TestCrossScript -v
```

## Code structure

All the embeddings code lives in `internal/embeddings/`:

| File | What it does |
|------|--------------|
| `service.go` | Main interface: Encode, Search, BuildIndex |
| `model.go` | Loads ONNX model via hugot library |
| `index.go` | Vector similarity search (brute-force, good enough for <100k items) |
| `cache.go` | LRU cache for embeddings |
| `script_detect.go` | Detects if text is Latin/Arabic/Cyrillic/etc. |

## Known limitations

- Model is ~450MB (has to be downloaded separately)
- First query is slow (~180ms) while model warms up
- Very short names (1-2 chars) don't work well
- Some rare scripts may have lower accuracy
