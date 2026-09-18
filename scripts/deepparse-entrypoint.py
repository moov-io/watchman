"""Start the deepparse FastAPI app without downloading unused FastText models.

The published image's lifespan calls download_models() for every model, including
the 6.8GB FastText embeddings, then skips initializing those parsers. Watchman
only uses BPEmb over HTTP, so skip the unused downloads.
"""

import deepparse.download_tools as download_tools

_BP_EMB_MODELS = ("bpemb", "bpemb-attention")


def download_models(saving_cache_path=None):
    for model_type in _BP_EMB_MODELS:
        download_tools.download_model(model_type, saving_cache_path=saving_cache_path)


download_tools.download_models = download_models

import uvicorn

if __name__ == "__main__":
    uvicorn.run("deepparse.app.app:app", host="0.0.0.0", port=8000)
