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
