---
layout: page
title: List history
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# List history

Watchman screens from an in-memory corpus. With MySQL or PostgreSQL configured, it also stores a snapshot of each downloaded list after a successful refresh. Search does not read this store.

Use it when you need to know whether a list actually changed, and which `sourceID`s were added, removed, or updated, so you re-screen only then. That is the request in [issue #486](https://github.com/moov-io/watchman/issues/486).

Ingested files are separate. See [File Dataset Ingestion](/watchman/ingest/).

## What is stored

After each successful download refresh, Watchman copies the prepared entities, groups them by list, and writes a delta in the background:

- A **ref** is the list file hash from `GET /v2/listinfo` (`listHashes`).
- If that hash is already stored for the list, the write is skipped. `latest` is the last *new* hash, so a rollback to an older file does not move `latest`. Pass `listinfo.listHashes` when you want the snapshot that matches search.
- The first snapshot for a list treats every row as added.
- Each entity is hashed from its public JSON (`PreparedFields` are omitted, then rebuilt with `Normalize()` on read).
- Empty `sourceID`s are skipped. Duplicate `sourceID`s in one refresh keep the last row.

Search still uses the in-memory corpus. History endpoints do not run Jaro-Winkler scoring.

Configure the same `Database` block used for ingest. There is no extra YAML key. Without a database the history endpoints return HTTP 501.

## Endpoints

| Method | Path | Notes |
|--------|------|--------|
| GET | `/v2/lists` | Latest ref per downloaded list. Empty `lists` array when nothing is stored yet. |
| GET | `/v2/lists/{listName}/refs` | Snapshots for one list, newest first. 404 if that list was never stored. |
| GET | `/v2/lists/{listName}/entities` | Rows at a ref (default `latest`). |
| GET | `/v2/lists/{listName}/entities/{sourceID}` | One row at a ref. |
| GET | `/v2/lists/{listName}/diff` | Added, removed, and changed rows between two refs. |

`listName` values match `sourceList` on entities (`us_ofac`, `eu_csl`, …). Extra query parameters return HTTP 400.

### Entities

`GET /v2/lists/us_ofac/entities?ref=latest&q=khoroshev&type=person&limit=50`

| Query | Default | Description |
|-------|---------|-------------|
| `ref` | `latest` | List hash, or `latest`. |
| `q` | | Case-insensitive substring on name or `sourceID`. |
| `type` | | `person`, `business`, `organization`, `aircraft`, `vessel`. |
| `limit` | 50 | Max 100. |
| `after` | | `sourceID` cursor from the previous page. |

This is a filter over the stored snapshot, not `/v2/search`.

### Diff

`GET /v2/lists/us_ofac/diff?to=latest&detail=summary`

| Query | Default | Description |
|-------|---------|-------------|
| `to` | `latest` | Newer ref. |
| `from` | previous of `to` | Older ref. If `to` is the first snapshot, `from` equals `to` and the diff is empty. |
| `detail` | `summary` | `summary` (ids and names) or `full` (entity bodies on `before` / `after`). |
| `limit` | 100 | Max 500, across added then changed then removed. |
| `after` | | Cursor `added:sourceID`, `changed:sourceID`, or `removed:sourceID`. |

`counts` is the full diff size. The `added`, `changed`, and `removed` arrays are the current page. If `from` is newer than `to`, Watchman swaps them.

Record hashes are on each entity summary as `entityHash`. Callers that store screens can skip `sourceID`s whose hash is unchanged.

## Go client

```go
lists, err := client.Lists(ctx)
refs, err := client.ListRefs(ctx, "us_ofac")
page, err := client.ListEntities(ctx, "us_ofac", search.ListEntitiesOpts{Query: "khoroshev"})
diff, err := client.ListDiff(ctx, "us_ofac", search.ListDiffOpts{Detail: "full"})
```

## Shutdown

In-flight stores finish for up to 30 seconds after the HTTP server stops. The search index swaps first; history copies from that refresh and writes in the background. A full history queue can delay the next refresh (single-flight) after search is already live.
