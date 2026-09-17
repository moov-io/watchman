---
layout: page
title: Network access
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Network access

Watchman is not designed to be served directly on the internet. Run it on a private network, or put a reverse proxy / API gateway in front of it. Authentication, ACLs, and rate limiting belong at the edge of the deployment, not inside Watchman.

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
