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

WASM UI: [http://localhost:8084](http://localhost:8084). Full recipe: [Using Watchman](/watchman/using-watchman/). Env vars: [Configuration](/watchman/config/).

For an optional [deepparse](/watchman/config/#deepparse) sidecar used in tests and examples:

```
make setup-deepparse
```

Search is `GET /v2/search`. Send `type`. Adding fields raises the score (OFAC SDN 48603):

```
# Name only
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":0.767}

# Name + date of birth
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&birthDate=1993-04-17&limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":0.867}

# Name + passport → 1.0
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&gov_passport=RU:2018278055&limit=1" \
  | jq '{name: .entities[0].name, match: .entities[0].match}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":1}
```

`minMatch=0.80` is a typical production cutoff. The name-only query would return no rows at that cutoff; the passport query would. Each result is the list record with `match` on the same object. Full recipe: [Using Watchman](/watchman/using-watchman/).
