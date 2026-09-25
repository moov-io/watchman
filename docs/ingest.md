---
layout: page
title: File Dataset Ingestion
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Custom File Dataset Ingestion

`POST /v2/ingest/{fileType}` uploads a CSV or Senzing file that Watchman parses with a schema you define in YAML. The file type must already exist in config; an unknown name does not create a new list.

Rows are stored in MySQL or PostgreSQL when you configure `Database`. Without a database, ingest lives only in that process and is gone on restart.

After each upload, Watchman loads **every** ingested row into the same in-memory search index used for OFAC and the other downloaded lists (name-token postings, ID blocks, source/type partitions). A per-source checksum in `ingest_sources` is compared on search so Watchman does not scan the entity table again when the snapshot is unchanged. Another process sharing the database picks up a new file on the next search when the checksum changes.

`GET /v2/listinfo` includes ingested sources (row counts and checksums) next to the downloaded lists.

If you persist ingested rows in your own database as well, store [record-linkage keys](/watchman/record-linkage/) (`recordlink.Keys`) rather than raw names or identifiers.

Ingest is on the unauthenticated business API. Body cap is 32MiB by default (`413` if larger). See [Network access](/watchman/network/).

### Path Parameters

- `fileType` (required): The type of file being ingested, corresponding to a specific configuration in the Watchman YAML (e.g., `fincen-business` or `fincen-person`).

### Request Body

The body is CSV (headers matching the YAML mapping) or Senzing JSON/JSONL when `format: senzing` is set.

Oversized uploads return `413`. Override the 32MiB default with `Ingest.MaxBodyBytes` or `INGEST_MAX_BODY_BYTES`.

```yaml
  Ingest:
    Files:
      "senzing-persons":
        Format: "senzing"
```

### Response

The response is a JSON object containing the parsed entities and the file type. The structure is as follows:

```json
{
  "fileType": string,
  "entities": []searchEntity
}
```

- `fileType`: The type of file ingested (e.g., fincen-business)
- `entities`: An array of parsed entities, each containing fields like name, type, sourceID, and additional fields based on the entity type (e.g., person or business).

#### Senzing Formatting

Set the `Accept` header or `format` query parameter to receive responses in [senzing format](https://www.senzing.com/docs/entity_specification/). To receive responses as JSON Lines specify the subformat as seen below.

```
Accept: senzing       # Array of objects [{...}, {...}]

Accept: senzing/jsonl # One object per line {...}\n{...}
```

The `format` query parameter accepts this as well, `?format=senzing` or `?format=senzing/jsonl`.

### YAML Configuration

Below is an example Watchman YAML configuration for two file types: fincen-business and fincen-person.

```yaml
Watchman:
  Ingest:
    MaxBodyBytes: 33554432 # 32MiB default; oversized uploads return HTTP 413
    files:
      fincen-business:
        format: csv
        mapping:
          name:
            column: business_name
          sourceID:
            column: tracking_number
          type:
            default: "business"
          business:
            name:
              column: business_name
            altNames:
              columns: dba_name
            created:
              column: incorporated
            governmentIDs:
              type:
                column: number_type
              identifier:
                column: number
          contact:
            phoneNumbers:
              columns: phone
          addresses:
            line1:
              columns: street
            city:
              columns: city
            state:
              columns: state
            postalCode:
              columns: zip
            country:
              columns: country
      fincen-person:
        format: csv
        mapping:
          name:
            merge: [first_name, suffix, middle_name, last_name]
          sourceID:
            column: tracking_number
          type:
            default: "person"
          person:
            name:
              merge: [first_name, middle_name, last_name]
            altNames:
              merge: [alias_first_name, alias_suffix, alias_middle_name, alias_last_name]
            birthDate:
              column: dob
            governmentIDs:
              type:
                column: number_type
              identifier:
                column: number
          contact:
            phoneNumbers:
              columns: phone
          addresses:
            line1:
              columns: street
            city:
              columns: city
            state:
              columns: state
            postalCode:
              columns: zip
            country:
              columns: country
```

### Schema Explanation

- **`format`**: `csv` or `senzing` (JSON / JSONL). CSV uses `mapping` below. Senzing files skip the column mapping.
- **`mapping`**: Defines how CSV columns map to entity fields. The mapping supports:
  - **`name`**: The entity’s name, either from a single `column` or `merge` of multiple columns (e.g., combining `first_name` and `last_name`).
  - **`sourceID`**: A unique identifier for the entity, mapped to a single `column`.
  - **`type`**: The entity type (e.g., `business` or `person`), set via a `default` value.
  - **`business`** or **`person`**: Entity-specific fields, such as:
    - `name`: The primary name (can be redundant with top-level `name`).
    - `altNames`: Alternate names, mapped to a single `columns` or `merge` of multiple columns.
    - `created` (business) or `birthDate` (person): A date field, parsed from a `column` using formats like `2006-01-02`, `1/2/2006`, or `01/02/2006`.
    - `governmentIDs`: Government-issued IDs, with `type` (e.g., `tax-id`) and `identifier` mapped to columns.
  - **`contact`**: Contact information, currently supporting `phoneNumbers` mapped to a `columns` field.
  - **`addresses`**: Physical addresses, with fields like `line1`, `city`, `state`, `postalCode`, and `country`, each mapped to a `columns` field.

### CSV File Requirements

- The CSV must include headers that match the column names specified in the Watchman configuration.
- Each row represents a single entity.
- Fields like dates must conform to one of the accepted formats (`2006-01-02`, `1/2/2006`, `01/02/2006`).
- Missing or empty fields are handled gracefully (e.g., skipped or set to empty values).

### Example CSV for `fincen-business`

```csv
business_name,tracking_number,dba_name,incorporated,number_type,number,phone,street,city,state,zip,country
"Acme Corp","12345","Acme Inc","2020-01-15","tax-id","EIN123","555-1234","123 Main St","Springfield","IL","62701","US"
```

### Example CSV for `fincen-person`

```csv
first_name,middle_name,last_name,suffix,tracking_number,alias_first_name,alias_middle_name,alias_last_name,alias_suffix,dob,number_type,number,phone,street,city,state,zip,country
"John","A","Doe","Jr","67890","Johnny","B","Doe","","1985-03-22","ssn","123-45-6789","555-5678","456 Oak St","Springfield","IL","62701","US"
```

### Search the File

`fileType` is the **source**, not `type`. `type` is still the entity kind (`person`, `business`, …):

```
GET /v2/search?source=fincen-person&type=person&name=John+Doe&minMatch=0.80
```

All ingested rows of that source are in the search index. Cross-script embeddings still apply to downloaded lists; ingested files score with the same Jaro–Winkler / identifier matcher.

### Exporting Ingested Data

`GET /v2/export/{fileType}` returns the entities previously ingested for a given fileType.

- Supports the same `Accept` header and `?format=` query param as search for Senzing output:
  - `?format=senzing` or `Accept: senzing` → JSON array
  - `?format=senzing/jsonl` or `Accept: senzing/jsonl` → JSON Lines (NDJSON)

Example:

```
curl -s "http://localhost:8084/v2/export/fincen-person" | jq .
```

Or for Senzing format:

```
curl -s -H "Accept: senzing/jsonl" "http://localhost:8084/v2/export/fincen-person"
```

The Go client exposes this as `client.ExportFile(ctx, "fincen-person")`.

This is the read counterpart to `POST /v2/ingest/{fileType}` for backup, migration, or feeding downstream entity resolution systems.
