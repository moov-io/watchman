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
