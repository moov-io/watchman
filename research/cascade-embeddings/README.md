# Cascade embedding model comparison

Scores the Cascade name-pair fixtures with Watchman Jaro–Winkler, then mixes in cosine from local Ollama embedding models. Hybrid is max(JW, cosine) only when the two names differ in script.

From the Watchman repo root:

```
go test ./research/cascade-embeddings

go run ./research/cascade-embeddings \
  -csv pkg/search/testdata/cascade-name-pairs-2.csv \
  -models qwen3-embedding:0.6b,qwen3-embedding:4b,qwen3-embedding:latest,bge-m3,nomic-embed-text \
  -out research/cascade-embeddings/RESULTS.md
```

Ollama must already have those models (`ollama pull …`). A missing model is skipped. This is pairwise scoring, not `/v2/search`.

Measured numbers: [RESULTS.md](RESULTS.md).
