## v0.9.0 (Unreleased)

ADDITIONS

- cmd/server: Include the BIS Denied Person's List data in search endpoints

IMPROVEMENTS

- build: push moov/ofac:latest on 'make release-push'
- docs: update docs.moov.io links after design refresh
- build: update dependencies

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
