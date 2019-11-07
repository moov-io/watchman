## v0.12.0 (Unreleased)

BREAKING CHANGES

- `Address` in our OpenAPI spec and generated Go client was renamed `EntiyAddress` to provide a more specific naming when combined in Moov's larger OpenAPI specification.

ADDITIONS

- cmd/server: add histogram for match percentages
- cmd/ofaccheck: initial setup of a cli tool for batch searches
- cmd/server: AND name and address queries if params are provided

BUILD

- chore(deps): update mui monorepo to v4.5.1
- build: upgrade openapi-generator to 4.2.0

## v0.11.0 (Released 2019-10-08)

This release adds a web interface for OFAC (developed by [Linden Lab](https://www.lindenlab.com/)) which allows for easier querying from desktop and mobile browsers. Also added are query params to apply additional filtering (exmaple: `sdnType=individual`) and we have improved match percentages to closer mirror the [official OFAC search tool](https://sanctionssearch.ofac.treas.gov/).

![](docs/images/webui.png)

ADDITIONS

- cmd/server: add the web interface developed by Linden Lab
- cmd/server: accept additional query params to filter SDN search results
  - `?sdnType=individual` and `?program=example`
- cmd/server: add endpoint for applications to grab distinct sets of column values
  - `GET /ui/values/sdnType` returns `["aircraft","individual","vessel"]`
- cmd/server: add `/search?id=NNN` endpoint for matching remark IDs
- cmd/server: return the oldest refresh time for our data in search results
- cmd/server: add `-base-path` for serving HTTP routes and web UI from non-root paths

IMPROVEMENTS

- api,client: specify x-request-id and x-user-id as optional HTTP headers
- cmd/ofactest: set x-request-id and x-user-id HTTP headers if CLI flags are set
- cmd/server: use a non-nil logger in search HTTP route tests
- cmd/server: adjust jaro weighting to maximize total weight in strong single word matches

BUG FIXES

- cmd/server: fix spelling of jaroWinkler
- cmd/server: never allow jaroWinkler to return NaN
- cmd/server: match treasury.gov match percentages

BUILD

- cmd/server: update github.com/moov-io/base to v0.10.0
- build: download CI tools rather than install
- build: remove mysql setup from docker-compose.yml

## v0.10.0 (Released 2019-08-16)

This release contains improvements to OFAC's match percentages to tone down false positives. We do this by adjusting the match percentage according to the query and SDN lenth ratio -- if the two strings are different lengths (after normalization) they cannot be equal. We are looking for feedback to further improve the matching code.

BREAKING CHANGES

- api,client: We've renamed all fields like `*Id` to `*ID` which is consistent with Go's style.

BUG FIXES

- attempt retries when downloading files
- download: return after successful download
- cmd/server: search: fix match percents after jaroWinkler change
- cmd/server: when reordering names handle multiple first or last names

ADDITIONS

- internal/database: log which provider we're using
- cmd/server: bind HTTP server with TLS if HTTPS_* variables are defined
- cmd/server: disable perioidic refresh via OFAC_DATA_REFRESH=off
- download: check for initial files on first refresh

IMPROVEMENTS

- build: upgrade openapi-generator to 4.1.0
- all: skip more tests on -short

## v0.9.0 (Released 2019-07-18)

BREAKING CHANGES

- The admin endpoint `:9092/ofac/refresh` was renamed to `:9092/data/refresh`

ADDITIONS

- cmd/server: Include the BIS Denied Person's List data in search endpoints
- cmd/server: Support MySQL as a storage layer via `DATABASE_TYPE=mysql` (See: [#100](https://github.com/moov-io/ofac/pull/100))

IMPROVEMENTS

- build: push moov/ofac:latest on 'make release-push'
- docs: update docs.moov.io links after design refresh
- build: update dependencies
- cmd/server: rename manual refresh endpoint to /data/refresh

## v0.8.0 (Released 2019-06-19)

BREAKING CHANGES

- api: Rename 'searchSDNs' to 'search' to reflect capability of other search capabilities
- api: Add `*OFAC*` on Company and Customer OpenAPI operations (and thus generated clients)

ADDITIONS

- api: Added Company routes and generated client code
- cmd/server: async: send webhooks on company/customer name watches (with match %)
- docs: add "A Framework for OFAC Compliance Commitments" from the US Treasury
- cmd/server: search: handle ?country and average weight with ?address
- cmd/server: handle city, state, providence, and zip in address search

BUG FIXES

- cmd/server: return database/sql Rows.Err
- api: add missing 'notes' field on UpdateCompanyStatus and UpdateCustomerStatus
- docs: describe moov's primary usage of OFAC (for paygate)

## v0.7.0 (Released 2019-04-25)

ADDITIONS

- cmd/server: support GET /search?q=... for searching SDN names, alt names, and address field

IMPROVEMENTS

- cmd/server: set CORS headers in ping route
- cmd/server: UTF-8 normalize names during pre-computation
- cmd/server: re-order SDN names where surname precedes "first name"
- cmd/server: support -log.format=json

BUG FIXES

- cmd/server: add missing database/sql Rows.Close()

BUILD

- Update to Go 1.12

## v0.6.0 (Released 2019-02-26)

IMPROVEMENTS

- Setup automated releases of binaries and Docker image

## v0.5.2 (Released 2019-02-21)

ADDITIONS

- Added `moov/ofactest` docker image

## v0.5.1 (Released 2019-02-21)

BUG FIXES

- `cmd/ofactest`, `client`: fix Watch create JSON and cleanup ofactest watches

## v0.5.0 (Released 2019-02-15)

ADDITIONS

- Add Company routes, persistence, client code, and watches
- `cmd/ofactest`: add -local and -webhook flags
- `cmd/ofactest`: read optional search query parameter

IMPROVEMENTS

- Ignore whitespace in string similarity
- `example`: Read Company or Customer JSON from webhook

BUG FIXES

- Fix capitalization of various JSON properties

## v0.4.0 (Released 2019-02-12)

ADDITIONS

- Added `"match": 0.91` as a percent of SDN name match to parameter.
- Support setting a customer to `unsafe` or `exception` status.

## v0.3.0 (Released 2019-02-07)

CHANGES

- Require `authToken` when creating a Watch (webhook notification)

ADDITIONS

- Added `last_ofac_data_refresh_success` Prometheus metric
- `cmd/ofactest`: for testing OFAC HTTP endpoints

## v0.2.1 (Released 2019-02-04)

IMPROVEMENTS

- Stop expecting `X-User-Id` header to be present (and non-empty)

## v0.2.0 (Released 2019-02-04)

IMPROVEMENTS

- Implement fuzzy search with Levenshtein for word to word comparisons
- Periodically refresh data (according to `OFAC_DATA_REFRESH`, default: `12h`)
- Write OpenAPI v3 specification and generate a Go client
- SQLite persistence for downloads, watches, and webhook results
- Add `GET /downloads?limit=N` for latest N data download metadata

BUG FIXES

- Remove OFAC null characters (`-0-`) from data (and HTTP api)

## v0.1.0 (Released 2019-01-18)

- Initial release
