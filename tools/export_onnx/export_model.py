#!/usr/bin/env python3
"""
Export sentence-transformers model to ONNX format for use with Go hugot library.

Usage:
    python export_model.py --output ./models/multilingual-minilm

This exports:
    - model.onnx: The transformer model
    - tokenizer.json: HuggingFace tokenizer config
    - tokenizer_config.json: Tokenizer settings
    - special_tokens_map.json: Special token mappings
    - vocab.txt: Vocabulary file
"""

import argparse
import os
import sys
from pathlib import Path

def export_model(model_name: str, output_dir: str, opset: int = 14):
    """Export sentence-transformers model to ONNX format."""

    print(f"Exporting {model_name} to {output_dir}")

    # Create output directory
    output_path = Path(output_dir)
    output_path.mkdir(parents=True, exist_ok=True)

    # Use optimum for ONNX export (handles sentence-transformers properly)
    try:
        from optimum.exporters.onnx import main_export
    except ImportError:
        print("Error: optimum not installed. Run: pip install 'optimum[exporters]>=1.17.0,<1.25.0'")
        sys.exit(1)

    # Export model
    print(f"Converting to ONNX (opset {opset})...")
    main_export(
        model_name_or_path=model_name,
        output=output_dir,
        task="feature-extraction",
        opset=opset,
        fp16=False,  # Keep FP32 for accuracy
    )

    print(f"Model exported to {output_dir}")

    # Verify output files exist
    expected_files = [
        "model.onnx",
        "tokenizer.json",
        "tokenizer_config.json",
    ]

    missing = []
    for f in expected_files:
        if not (output_path / f).exists():
            missing.append(f)

    if missing:
        print(f"Warning: Missing expected files: {missing}")
    else:
        print("All expected files created successfully")

    # Print model size
    model_path = output_path / "model.onnx"
    if model_path.exists():
        size_mb = model_path.stat().st_size / (1024 * 1024)
        print(f"Model size: {size_mb:.1f} MB")

    return output_path


def verify_export(output_dir: str, test_texts: list[str] = None):
    """Verify the exported ONNX model produces correct embeddings."""

    print("\nVerifying exported model...")

    if test_texts is None:
        test_texts = [
            "Mohamed Ali",
            "محمد علي",  # Arabic
            "Vladimir Putin",
            "Владимир Путин",  # Cyrillic
            "Kim Jong Un",
            "金正恩",  # Chinese
        ]

    # Load original model
    from sentence_transformers import SentenceTransformer
    original_model = SentenceTransformer("sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2")
    original_embeddings = original_model.encode(test_texts, normalize_embeddings=True)

    # Load ONNX model
    import onnxruntime as ort
    from transformers import AutoTokenizer
    import numpy as np

    output_path = Path(output_dir)
    session = ort.InferenceSession(str(output_path / "model.onnx"))
    tokenizer = AutoTokenizer.from_pretrained(output_dir)

    # Run inference
    inputs = tokenizer(test_texts, padding=True, truncation=True, return_tensors="np")

    onnx_outputs = session.run(
        None,
        {
            "input_ids": inputs["input_ids"],
            "attention_mask": inputs["attention_mask"],
        }
    )

    # Mean pooling + normalization (same as sentence-transformers)
    token_embeddings = onnx_outputs[0]  # (batch, seq_len, hidden_dim)
    attention_mask = inputs["attention_mask"]

    # Expand attention mask
    input_mask_expanded = np.expand_dims(attention_mask, -1)
    input_mask_expanded = np.broadcast_to(input_mask_expanded, token_embeddings.shape).astype(float)

    # Mean pooling
    sum_embeddings = np.sum(token_embeddings * input_mask_expanded, axis=1)
    sum_mask = np.clip(np.sum(input_mask_expanded, axis=1), a_min=1e-9, a_max=None)
    onnx_embeddings = sum_embeddings / sum_mask

    # L2 normalize
    norms = np.linalg.norm(onnx_embeddings, axis=1, keepdims=True)
    onnx_embeddings = onnx_embeddings / norms

    # Compare embeddings
    print("\nComparison (Original vs ONNX):")
    max_diff = 0
    for i, text in enumerate(test_texts):
        diff = np.max(np.abs(original_embeddings[i] - onnx_embeddings[i]))
        max_diff = max(max_diff, diff)
        cosine_sim = np.dot(original_embeddings[i], onnx_embeddings[i])
        status = "✓" if cosine_sim > 0.999 else "✗"
        print(f"  {status} {text[:30]:30s} - cosine: {cosine_sim:.6f}, max_diff: {diff:.6f}")

    print(f"\nMax absolute difference: {max_diff:.6f}")

    if max_diff < 1e-4:
        print("✓ Export verified successfully!")
        return True
    else:
        print("✗ Warning: Embeddings differ significantly")
        return False


def main():
    parser = argparse.ArgumentParser(description="Export sentence-transformers to ONNX")
    parser.add_argument(
        "--model",
        default="sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2",
        help="HuggingFace model name"
    )
    parser.add_argument(
        "--output",
        default="./models/multilingual-minilm",
        help="Output directory for ONNX model"
    )
    parser.add_argument(
        "--opset",
        type=int,
        default=14,
        help="ONNX opset version"
    )
    parser.add_argument(
        "--verify",
        action="store_true",
        help="Verify exported model against original"
    )
    parser.add_argument(
        "--verify-only",
        action="store_true",
        help="Only verify existing export (skip export)"
    )

    args = parser.parse_args()

    if not args.verify_only:
        export_model(args.model, args.output, args.opset)

    if args.verify or args.verify_only:
        success = verify_export(args.output)
        sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()
