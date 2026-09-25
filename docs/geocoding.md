---
layout: page
title: Geocoding
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Geocoding

Watchman can attach latitude and longitude to **list** addresses at refresh time. It is optional and **off by default** (since v0.57.0). Providers: OpenCage, Nominatim, or Google Maps.

Geocoding does **not** change match scores today. Address similarity still compares parsed text fields (line, city, state, postal code, country). Coordinates are stored on each `Address` (`latitude` / `longitude`), returned in search JSON, and shown in the UI.

Use it when a downstream system wants a map pin, a distance filter you apply yourself, or a stable lat/long on exported records. Skip it if you only need Watchman to score names and IDs.

## What happens when it is on

On each list refresh, Watchman walks every loaded entity that has addresses and calls the provider (one goroutine per entity). Hits go through an in-memory LRU (L1) and, if you enable it, a database cache (L2). Failed lookups leave the address unchanged; they do not fail the refresh.

Query-time `address=` on `/v2/search` is parsed (libpostal / usaddress / deepparse). Those query addresses are **not** geocoded. Scoring still compares the parsed fields.

## Tradeoffs

| Factor | What you pay | Why it matters |
|--------|--------------|----------------|
| Startup / refresh time | One provider call per distinct address, rate-limited (default 1.5 req/s, burst 5) | A list with tens of thousands of addresses can take hours on a cold cache. L1/L2 caches make later refreshes cheap. |
| Money / quota | OpenCage and Google are paid APIs. Nominatim is free with a [usage policy](https://operations.osmfoundation.org/policies/nominatim/) | Default 1.5 req/s is meant to stay polite. Self-host Nominatim if you cannot use a commercial key. |
| Memory / DB | L1 holds up to `L1MaxSize` entries (default 10,000) for `L1TTL` (default 24h). L2 writes to MySQL/Postgres | L2 needs the `Database` block. Without it, coordinates are recomputed after restart. |
| Score | No change | `Similarity` does not use lat/long. Turning geocoding on will not move `minMatch` hits. |
| Privacy | Full street addresses leave your network | Prefer a self-hosted Nominatim if addresses must stay inside your network. |

Rate-limit and cache YAML: [Configuration](/watchman/config/#geocoding).

## When it would change a score

It would only matter if scoring compared coordinates (for example a distance threshold between two pins). That path is not implemented. Two records at “123 Main St, Springfield” vs “123 Main Street, Springfield, IL” still match on tokens and country folding, with or without geocoding.

If you need distance logic, read `latitude` / `longitude` from the search response (or from ingest/export) and apply it in your application.

## Providers

Set `Watchman.Geocoding.Enabled: true` and a provider name. `GEOCODING_API_KEY` overrides `Provider.ApiKey`.

### OpenCage

[OpenCage Geocoding API](https://opencagedata.com/). Config name: `opencage`. Paid key required for production volume.

### Google

[Google Geocoding API](https://developers.google.com/maps/documentation/geocoding). Config name: `google`. Paid key required.

### Nominatim

[Nominatim](https://nominatim.org/) (OpenStreetMap). Config name: `nominatim`. No API key. Set `Provider.BaseURL` to a [self-hosted](https://nominatim.org/release-docs/latest/admin/Installation/) instance for production. The public OSM endpoint is for light, policy-compliant use only.

To add a provider, [open an issue](https://github.com/moov-io/watchman/issues/new).
