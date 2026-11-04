"""Start the deepparse FastAPI app with only the BPEmb parsers Watchman uses.

The published 1.0.0 image's lifespan calls download_models() for every model,
including the 6.8GB FastText embeddings and Model2Vec, then initializes every
non-FastText parser. Watchman only calls BPEmb over HTTP, so drop the other
MODEL_MAPPING_CHOICES keys before uvicorn loads the app. That skips the unused
downloads and stops AddressParser(offline=True) from failing on Model2Vec
weights that are not in the cache.
"""

import deepparse.download_tools as download_tools

_BP_EMB_MODELS = ("bpemb", "bpemb-attention")


def download_models(saving_cache_path=None):
    for model_type in _BP_EMB_MODELS:
        download_tools.download_model(model_type, saving_cache_path=saving_cache_path)


download_tools.download_models = download_models

for key in list(download_tools.MODEL_MAPPING_CHOICES):
    if key not in _BP_EMB_MODELS:
        del download_tools.MODEL_MAPPING_CHOICES[key]

import uvicorn

if __name__ == "__main__":
    uvicorn.run("deepparse.app.app:app", host="0.0.0.0", port=8000)

