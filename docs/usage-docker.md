---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Quick Start with Docker

Images: [`moov/watchman`](https://hub.docker.com/r/moov/watchman/) on Docker Hub, [`quay.io/moov/watchman`](https://quay.io/repository/moov/watchman?tab=tags) for OpenShift. `moov/watchman:v2-static` ships frozen 2019 files for fast local tests.

Business API on **:8084**. Admin/metrics on **:9094**. Do not expose Watchman on the public internet. See [Network access](/watchman/network/).

```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
```

WASM UI: [http://localhost:8084](http://localhost:8084). In another terminal, wait until OFAC is indexed, then screen at a production cutoff:

```
until curl -sf http://localhost:8084/v2/listinfo | jq -e '.lists.us_ofac > 0' >/dev/null; do sleep 2; done

curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&gov_passport=RU:2018278055&minMatch=0.80&limit=1" \
  | jq '{name: .entities[0].name, match: .entities[0].match, sourceID: .entities[0].sourceID}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":1,"sourceID":"48603"}
```

That is OFAC SDN 48603. The passport is an identity key, so the score is **1.0** at `minMatch=0.80`. Full recipe: [Using Watchman](/watchman/using-watchman/). Env vars: [Configuration](/watchman/config/).

For an optional [deepparse](/watchman/config/#deepparse) sidecar used in tests and examples:

```
make setup-deepparse
```

The same name without an ID is below the screening line. Date of birth clears it; the passport is 1.0:

```
# Name only — empty at minMatch=0.80
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&minMatch=0.80&limit=1" | jq .entities
# []

# Name + date of birth
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&birthDate=1993-04-17&minMatch=0.80&limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":0.867}
```

Each result is the list record with `match` on the same object.
