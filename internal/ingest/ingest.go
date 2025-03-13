package ingest

// TODO(adam): basic layout for API driven ingest
//
// POST /ingest/{name}
//  name = fincen-business, fincen-person
//
// In FinCEN example, it's csv, mapper is constructed from YAML config
// Results are loaded into memory, stored in DB? How do we distribute this list to all instances?
//
// Each POST replaces previous list records, useful for weekly/monthly/etc reports over a changing source list.
// e.g. FinCEN, employee searches, etc
//
// So search-screen can consume FileSaved event, POST to a (separate?) Watchman instance and run search against that?
//
// Search
// GET /v2/search?sourceList=fincen-business&name=John&type=person
