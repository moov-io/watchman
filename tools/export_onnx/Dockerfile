# Dockerfile for exporting the multilingual sentence-transformers model to ONNX
#
# This image exports paraphrase-multilingual-MiniLM-L12-v2 to ONNX format
# for use with the Go hugot library in watchman's cross-script name matching.
#
# Usage:
#   docker build -t watchman-onnx-export .
#   docker run -v $(pwd)/models:/output watchman-onnx-export
#
# The exported model will be in ./models/

# Force x86_64 platform for consistent builds (torch packages are x86_64 only)
FROM --platform=linux/amd64 python:3.12-slim AS builder

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Copy requirements first (better layer caching)
COPY requirements.txt .

# Install CPU-only torch AND torchvision together for compatibility
# The ONNX export doesn't need GPU
RUN pip install --no-cache-dir \
    --index-url https://download.pytorch.org/whl/cpu \
    torch==2.4.1+cpu \
    torchvision==0.19.1+cpu \
    && pip install --no-cache-dir -r requirements.txt

# Copy export script
COPY export_model.py .

# Export the model during build
RUN python export_model.py --output /model

# Verify the export succeeded
RUN test -f /model/model.onnx && \
    test -f /model/tokenizer.json && \
    echo "Model export verified successfully"

# Final stage: minimal image with just the model
FROM --platform=linux/amd64 alpine:3.20

WORKDIR /model

# Copy exported model from builder
COPY --from=builder /model /model

# Default command: copy model to output volume
CMD ["sh", "-c", "cp -r /model/* /output/ && echo 'Model exported to /output/'"]
