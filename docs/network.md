---
layout: page
title: Network access
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Network access

Watchman is not designed to be served directly on the internet. Run it on a private network, or put a reverse proxy / API gateway with allowed paths in front of it. Authentication, ACLs, and rate limiting belong at the edge of the deployment, not inside Watchman.

## HTTP API (`BindAddress`)

The business API (`BindAddress`, `:8084` by default) serves search, ingest, export, data refresh, the web UI, and (when enabled) MCP. These routes are unauthenticated by design.

Operators who need auth or rate limits should terminate them on the proxy in front of `:8084`. `POST /v2/ingest/{fileType}` only accepts `fileType` values already defined in config; an unknown type does not create a new list.

See [issue #875](https://github.com/moov-io/watchman/issues/875).

## Admin server (`AdminAddress`)

The admin HTTP server is a **separate port** (`AdminAddress`, `:9094` by default) so deployments can isolate it from the business API:

- Firewall or ACL the admin port
- Bind it to an internal interface (for example `127.0.0.1:9094`)
- Or block the port entirely

Prometheus metrics (`/metrics`) and `/version` live on the admin server on purpose and are unauthenticated. Unauthenticated `/metrics` on that port is expected.

See [issue #875](https://github.com/moov-io/watchman/issues/875).

## Download URLs

List download locations are operator configuration, not API input. Nothing in the HTTP API lets a caller pick a URL.

- `Download.Senzing[].Location` and `OpenSanctions.Lists[].Location` come from YAML
- `OFAC_DOWNLOAD_TEMPLATE` and the other `*_DOWNLOAD_TEMPLATE` / `*_DOWNLOAD_URL` variables come from the environment
- `POST /v2/data/refresh` refetches whatever is already configured

`file://` is supported so operators can load a local Senzing file (see [Configuration](/watchman/config/#download) and [Caching Data Files](/watchman/cache-data-files/)). Pointing templates at an internal HTTP cache, including `http://` on a private network, is the intended way to run [watchman-cache](https://github.com/moov-io/watchman-cache).

Watchman does not allowlist download hosts, reject `file:` paths, or block private / loopback / link-local addresses. A cross-OS denylist of "sensitive" paths is not maintained. Whoever can write `APP_CONFIG` or the process environment already controls the process.

See [issue #876](https://github.com/moov-io/watchman/issues/876).
