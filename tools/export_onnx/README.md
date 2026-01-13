# ONNX Model Export for Cross-Script Name Matching

This tool exports the `paraphrase-multilingual-MiniLM-L12-v2` sentence-transformers model to ONNX format for use with the Go `hugot` library.

## Why ONNX?

The embedding model enables **cross-script name matching** - finding that "محمد علي" (Arabic) matches "Mohamed Ali" (Latin). The ONNX export allows native Go inference without Python dependencies.

## Usage

### Option 1: Docker (Recommended)

The easiest way to get the model is using Docker:

```bash
cd tools/export_onnx
docker build -t watchman-onnx-export .
docker run --rm -v $(pwd)/../../models/multilingual-minilm:/output watchman-onnx-export
```

This builds an image with the model pre-exported (~510MB) and copies it to your local directory.

### Option 2: Manual Setup

If you prefer to run the export locally:

```bash
cd tools/export_onnx
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python export_model.py --output ../../models/multilingual-minilm --verify
```

**Note:** The `requirements.txt` has pinned versions to ensure compatibility. If you encounter issues, make sure you're using Python 3.10+.

### Export Options

```bash
python export_model.py --output ../../models/multilingual-minilm --verify
```

This creates:
```
models/multilingual-minilm/
├── model.onnx           # ~140MB ONNX model
├── tokenizer.json       # HuggingFace tokenizer
├── tokenizer_config.json
├── special_tokens_map.json
└── vocab.txt
```

### 3. Verify Export

```bash
python test_export.py --model-dir ../../models/multilingual-minilm
```

Tests:
- Embedding dimensions (384)
- Batch processing
- Cross-script similarity scores

## Model Details

| Property | Value |
|----------|-------|
| Model | `paraphrase-multilingual-MiniLM-L12-v2` |
| Dimensions | 384 |
| Languages | 50+ |
| Size | ~140MB |
| ONNX Opset | 14 |

## Cross-Script Examples

| Query | Match | Similarity |
|-------|-------|------------|
| محمد علي | Mohamed Ali | 0.97 |
| Владимир Путин | Vladimir Putin | 0.92 |
| 金正恩 | Kim Jong Un | 0.85 |

## Go Integration

The exported model is used by `internal/embeddings/` package via the `knights-analytics/hugot` library:

```go
import "github.com/knights-analytics/hugot"

session, _ := hugot.NewORTSession()
pipeline, _ := hugot.NewPipeline(session, hugot.FeatureExtractionConfig{
    ModelPath: "models/multilingual-minilm",
})

embeddings, _ := pipeline.RunPipeline([]string{"Mohamed Ali", "محمد علي"})
```
