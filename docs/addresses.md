---
layout: page
title: Addresses
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Address Parsing and Normalization

Watchman parses free-text addresses on ingest and on `/v2/search`.

- **Docker images, OpenShift images, and Linux/macOS GitHub release binaries** are built with [libpostal](https://github.com/openvenues/libpostal) (Senzing data). Address parse is in-process libpostal (~3GB RAM for models).
- **Any build** can use [deepparse](https://github.com/GRAAL-Research/deepparse) instead: set `Watchman.Deepparse.Enabled: true` and run the HTTP sidecar.
- **Otherwise** (Windows GitHub `.exe`, `go build` / `go run` without `-tags libpostal`) Watchman uses the built-in US-oriented parser (`usaddress`).

libpostal settings are under `Watchman.PostalPool` in the [config file](/watchman/config/#libpostal). Leave `Enabled: false` for in-process parse in Docker. Set `Enabled: true` if you need extra libpostal worker processes.

### Deepparse (optional)

[GRAAL-Research/deepparse](https://github.com/GRAAL-Research/deepparse) is a multinational neural address parser that runs as an HTTP sidecar. Watchman calls it with [deepparse-go](https://github.com/adamdecaf/deepparse-go).

Start a local instance from this repo:

```
make setup-deepparse   # docker compose --profile deepparse up -d
```

Then enable it in config (`Watchman.Deepparse.Enabled: true`, `BaseURL: http://localhost:8000`). See [Configuration](/watchman/config/#deepparse). When Deepparse is enabled, Watchman uses it instead of libpostal or usaddress.
