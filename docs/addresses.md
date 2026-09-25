---
layout: page
title: Addresses
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Address Parsing and Normalization

Watchman parses free-text addresses on ingest and on `/v2/search`.

- **Docker and OpenShift images** (`moov/watchman`, `quay.io/moov/watchman`) are built with [libpostal](https://github.com/openvenues/libpostal) (Senzing data). Address parse is in-process libpostal.
- **Any build** can use [deepparse](https://github.com/GRAAL-Research/deepparse) instead: set `Watchman.Deepparse.Enabled: true` and run the HTTP sidecar. Deepparse wins if PostalPool is also enabled.
- **Otherwise** (GitHub release binaries, `go build` / `go run` without `-tags libpostal`) Watchman uses the built-in US-oriented parser (`usaddress`).

PostalPool is an optional extra on libpostal builds: a pool of `postal-server` processes (`Watchman.PostalPool.Enabled`). It is off by default, including in Docker. See [Configuration](/watchman/config/#postalpool).

### Deepparse (optional)

[GRAAL-Research/deepparse](https://github.com/GRAAL-Research/deepparse) is a multinational neural address parser that runs as an HTTP sidecar. Watchman calls it with [deepparse-go](https://github.com/adamdecaf/deepparse-go).

Start a local instance from this repo:

```
make setup-deepparse   # docker compose --profile deepparse up -d
```

Then enable it in config (`Watchman.Deepparse.Enabled: true`, `BaseURL: http://localhost:8000`). See [Configuration](/watchman/config/#deepparse). Do not enable Deepparse and PostalPool together; Deepparse wins and PostalPool is skipped.
