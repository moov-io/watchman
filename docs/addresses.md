---
layout: page
title: Addresses
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Address Parsing and Normalization

Watchman optionally uses **[libpostal](https://github.com/openvenues/libpostal)** to deliver advanced address parsing, classification, and normalization. By integrating the latest data from **[Senzing’s classifier, data, and parser](https://github.com/Senzing/libpostal-data)**, it accurately handles global addresses during indexing and screening requests.

To enable libpostal support, activate the **[PostalPool](/watchman/config/#postalpool)** option in Watchman’s configuration.

**Docker / OpenShift images** compile with `-tags libpostal`, so `ParseAddress` uses libpostal in-process. **GitHub release binaries** (`make dist`) and a normal `go build` / `go run` do not set that tag, so they use `usaddress`. PostalPool (`Watchman.PostalPool.Enabled`) is a separate, optional pool of `postal-server` processes; it is off by default even in the Docker image.

### Deepparse (optional)

[GRAAL-Research/deepparse](https://github.com/GRAAL-Research/deepparse) is an alternative multinational address parser that runs as an HTTP sidecar. Watchman calls it with [deepparse-go](https://github.com/adamdecaf/deepparse-go). It is **opt-in**; the default remains libpostal / usaddress.

Start a local instance from this repo:

```
make setup-deepparse   # docker compose --profile deepparse up -d
```

Then enable it in config (`Watchman.Deepparse.Enabled: true`, `BaseURL: http://localhost:8000`). See [Configuration](/watchman/config/#deepparse). Do not enable Deepparse and PostalPool together; Deepparse wins and PostalPool is skipped.
