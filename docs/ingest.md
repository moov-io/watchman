---
layout: page
title: File Dataset Ingestion
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## File Dataset Ingestion

Watchman has a `POST /v2/ingest/{fileType}` endpoint allows you to upload CSV files containing entity data (e.g., businesses or persons) for ingestion.
The endpoint processes the file based on a predefined schema defined in the Watchman YAML configuration and returns a JSON response with the parsed entities.
The parsed entities are included in Watchman's memory for search, but kept as a separate list. Callers must specify a entity type matching the `fileType`.

### Path Parameters

- `fileType` (required): The type of file being ingested, corresponding to a specific configuration in the Watchman YAML (e.g., `fincen-business` or `fincen-person`).

### Request Body

- A CSV file containing entity data, with headers matching the schema defined in the Watchman configuration.

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

### YAML Configuration

Below is an example Watchman YAML configuration for two file types: fincen-business and fincen-person.

```yaml
Watchman:
  Ingest:
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

- **`format`**: Specifies the file format. Currently, only `csv` is supported.
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

You can then perform searches against the ingested file.

```
GET /v2/search?type=fincen-person&name=John+Doe&type=person
```
