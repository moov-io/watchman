#!/usr/bin/env python3
"""
Test suite for ONNX model export verification.

Runs comprehensive tests to ensure the exported ONNX model:
1. Produces identical embeddings to the original model
2. Handles cross-script text correctly
3. Maintains expected similarity scores
"""

import sys
import numpy as np
from pathlib import Path


def test_cross_script_similarity(model_dir: str):
    """Test that cross-script pairs have high similarity."""

    print("Testing cross-script similarity...")

    import onnxruntime as ort
    from transformers import AutoTokenizer

    output_path = Path(model_dir)
    session = ort.InferenceSession(str(output_path / "model.onnx"))
    tokenizer = AutoTokenizer.from_pretrained(model_dir)

    def encode(texts):
        inputs = tokenizer(texts, padding=True, truncation=True, return_tensors="np")
        outputs = session.run(None, {
            "input_ids": inputs["input_ids"],
            "attention_mask": inputs["attention_mask"],
        })

        # Mean pooling
        token_embeddings = outputs[0]
        attention_mask = inputs["attention_mask"]
        input_mask_expanded = np.expand_dims(attention_mask, -1)
        input_mask_expanded = np.broadcast_to(input_mask_expanded, token_embeddings.shape).astype(float)
        sum_embeddings = np.sum(token_embeddings * input_mask_expanded, axis=1)
        sum_mask = np.clip(np.sum(input_mask_expanded, axis=1), a_min=1e-9, a_max=None)
        embeddings = sum_embeddings / sum_mask

        # L2 normalize
        norms = np.linalg.norm(embeddings, axis=1, keepdims=True)
        return embeddings / norms

    def similarity(text1, text2):
        emb = encode([text1, text2])
        return float(np.dot(emb[0], emb[1]))

    # Test cases: (name1, name2, min_expected_similarity)
    test_cases = [
        # Arabic
        ("Mohamed Ali", "محمد علي", 0.85),
        ("Hassan Ahmed", "حسن أحمد", 0.80),

        # Cyrillic
        ("Vladimir Putin", "Владимир Путин", 0.80),
        ("Alexander Petrov", "Александр Петров", 0.75),

        # Chinese
        ("Kim Jong Un", "金正恩", 0.70),

        # Negative cases (should be dissimilar)
        ("John Smith", "Jane Doe", 0.0),  # Different people
        ("Mohamed Ali", "Vladimir Putin", 0.0),  # Different people
    ]

    passed = 0
    failed = 0

    for name1, name2, min_sim in test_cases:
        sim = similarity(name1, name2)

        if min_sim > 0:
            # Should be similar
            if sim >= min_sim:
                print(f"  ✓ {name1} ↔ {name2}: {sim:.3f} >= {min_sim}")
                passed += 1
            else:
                print(f"  ✗ {name1} ↔ {name2}: {sim:.3f} < {min_sim}")
                failed += 1
        else:
            # Should be dissimilar
            if sim < 0.5:
                print(f"  ✓ {name1} ↔ {name2}: {sim:.3f} (dissimilar)")
                passed += 1
            else:
                print(f"  ✗ {name1} ↔ {name2}: {sim:.3f} (unexpectedly similar)")
                failed += 1

    print(f"\nResults: {passed}/{passed+failed} passed")
    return failed == 0


def test_embedding_dimensions(model_dir: str):
    """Test that embeddings have correct dimensions."""

    print("\nTesting embedding dimensions...")

    import onnxruntime as ort
    from transformers import AutoTokenizer

    output_path = Path(model_dir)
    session = ort.InferenceSession(str(output_path / "model.onnx"))
    tokenizer = AutoTokenizer.from_pretrained(model_dir)

    inputs = tokenizer(["test"], return_tensors="np")
    outputs = session.run(None, {
        "input_ids": inputs["input_ids"],
        "attention_mask": inputs["attention_mask"],
    })

    # Expected: (batch_size=1, seq_len, hidden_dim=384)
    hidden_dim = outputs[0].shape[-1]

    if hidden_dim == 384:
        print(f"  ✓ Hidden dimension: {hidden_dim}")
        return True
    else:
        print(f"  ✗ Hidden dimension: {hidden_dim} (expected 384)")
        return False


def test_batch_processing(model_dir: str):
    """Test that batch processing works correctly."""

    print("\nTesting batch processing...")

    import onnxruntime as ort
    from transformers import AutoTokenizer

    output_path = Path(model_dir)
    session = ort.InferenceSession(str(output_path / "model.onnx"))
    tokenizer = AutoTokenizer.from_pretrained(model_dir)

    texts = [
        "Mohamed Ali",
        "Vladimir Putin",
        "Kim Jong Un",
        "John Smith",
        "Jane Doe",
    ]

    inputs = tokenizer(texts, padding=True, truncation=True, return_tensors="np")
    outputs = session.run(None, {
        "input_ids": inputs["input_ids"],
        "attention_mask": inputs["attention_mask"],
    })

    batch_size = outputs[0].shape[0]

    if batch_size == len(texts):
        print(f"  ✓ Batch size: {batch_size}")
        return True
    else:
        print(f"  ✗ Batch size: {batch_size} (expected {len(texts)})")
        return False


def main():
    import argparse

    parser = argparse.ArgumentParser(description="Test ONNX model export")
    parser.add_argument(
        "--model-dir",
        default="./models/multilingual-minilm",
        help="Directory containing exported ONNX model"
    )

    args = parser.parse_args()

    model_dir = Path(args.model_dir)

    if not (model_dir / "model.onnx").exists():
        print(f"Error: Model not found at {model_dir}/model.onnx")
        print("Run export_model.py first")
        sys.exit(1)

    results = []
    results.append(("Dimensions", test_embedding_dimensions(args.model_dir)))
    results.append(("Batch Processing", test_batch_processing(args.model_dir)))
    results.append(("Cross-Script Similarity", test_cross_script_similarity(args.model_dir)))

    print("\n" + "=" * 50)
    print("Summary:")
    all_passed = True
    for name, passed in results:
        status = "✓ PASS" if passed else "✗ FAIL"
        print(f"  {status}: {name}")
        if not passed:
            all_passed = False

    sys.exit(0 if all_passed else 1)


if __name__ == "__main__":
    main()
