---
layout: page
title: Caching Data Files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Caching Data Files for Optimal Performance

Watchman downloads sanctions lists on startup and on a refresh interval (default 12 hours). Government sites can be slow, redirect to short-lived S3 URLs, or go down for days. There are two ways to keep Watchman from depending on those origins at every start.

1. **[watchman-cache](https://github.com/moov-io/watchman-cache)** — an nginx reverse proxy that Watchman talks to with the same `*_DOWNLOAD_TEMPLATE` / `*_DOWNLOAD_URL` variables. Images: `moov/watchman-cache`, `ghcr.io/moov-io/watchman-cache`, `quay.io/moov/watchman-cache`.
2. **`INITIAL_DATA_DIRECTORY`** — a local folder of list files Watchman reads at startup. Periodic refresh still hits the network unless those locations are local `file://` paths.

Use watchman-cache in production. Use `INITIAL_DATA_DIRECTORY` for air-gapped starts or when you already vendor the files.

## watchman-cache

[moov-io/watchman-cache](https://github.com/moov-io/watchman-cache) sits in front of Watchman’s downloads. It follows 302 redirects internally, caches complete list files on a Docker volume (48 hours fresh, stale serving for days if the origin is down), and only allows the filenames Watchman actually requests.

Point Watchman at the cache (on the Compose network the service is `cache` or `watchman-cache` on port 8080; from the host, `localhost:3000`):

```
OFAC_DOWNLOAD_TEMPLATE=http://watchman-cache:8080/%s
US_CSL_DOWNLOAD_TEMPLATE=http://watchman-cache:8080/%s
US_NON_SDN_DOWNLOAD_TEMPLATE=http://watchman-cache:8080/%s
EU_CSL_DOWNLOAD_URL=http://watchman-cache:8080/eu_csl.csv
UK_SANCTIONS_LIST_URL=http://watchman-cache:8080/UK_Sanctions_List.csv
UN_CONSOLIDATED_LIST_URL=http://watchman-cache:8080/un_consolidated.xml
FINCEN_311_DOWNLOAD_URL=http://watchman-cache:8080/fincen_311.html
DOWNLOAD_TIMEOUT=180s
```

Origin-shaped paths (for example `http://watchman-cache:8080/api/PublicationPreview/exports/%s`) also work; the cache rewrites them to the same files. Set `DOWNLOAD_TIMEOUT=180s` so Watchman waits for a cold fetch. Keep the cache volume across deploys (`docker compose down` without `-v`).

Quick start and list table: [watchman-cache README](https://github.com/moov-io/watchman-cache). Env vars: [Download reliability](/watchman/config/#download-reliability).

## INITIAL_DATA_DIRECTORY

By setting `INITIAL_DATA_DIRECTORY` to a local directory Watchman will look for the following files.

**OFAC**

- `add.csv` - Address
- `alt.csv` - Alternate ID
- `sdn.csv` - Specially Designated National
- `sdn_comments.csv` - Specially Designated National Comments

Download the files from [Data Center - SDN List](https://sanctionslist.ofac.treas.gov/Home/SdnList)

**US Consolidated Screening List**

- `consolidated.csv` - US Consolidated Screening List

Download the [US Consolidated Screening List](https://www.trade.gov/consolidated-screening-list)

**EU Consolidated Screening List**

- `eu_csl.csv` - EU Consolidated Screening List

Download the [EU Consolidated Screening List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)

**UK Consolidated Screening List**

- `UK_Sanctions_List.csv` - UK Sanctions List

Download the [UK Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list)

**UN Consolidated Sanctions List**

- `un_consolidated.xml` - UN Consolidated Sanctions List

Download from [UN Sanctions](https://www.un.org/sc/resources/sc-sanctions)

**FinCEN 311 Special Measures**

- `fincen_311.html` - FinCEN 311/9714 Special Measures page (parsed for actions)

Point `FINCEN_311_DOWNLOAD_URL` or place the HTML snapshot in your `INITIAL_DATA_DIRECTORY`.

---

**Senzing / OpenSanctions formatted lists**

For any `opensanctions_*` or custom `Download.Senzing` entries, the `Location` (URL or `file://...`) is used. When using `INITIAL_DATA_DIRECTORY`, relative file references or exact filenames referenced in the location can be resolved from that directory for startup without network access. Periodic refreshes will still attempt live fetches unless the location is a stable local file path.

`Location` and the `*_DOWNLOAD_TEMPLATE` / `*_DOWNLOAD_URL` variables are operator configuration. `file://` and internal HTTP caches (including watchman-cache) are supported on purpose; Watchman does not allowlist hosts or jail local paths. See [Network access](/watchman/network/).
