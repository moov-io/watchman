---
layout: page
title: Configuration
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Configuration settings

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `DATA_REFRESH_INTERVAL` | Interval for data redownload and reparse. `off` disables this refreshing. | 12h |
| `INITIAL_DATA_DIRECTORY` | Directory filepath with initial files to use instead of downloading. Periodic downloads will replace the initial files. | Empty |
| `SEARCH_MAX_WORKERS` | Maximum number of goroutines used for search. | 1024 |
| `ADJACENT_SIMILARITY_POSITIONS` | How many nearby words to search for highest max similarly score. | 3 |
| `EXACT_MATCH_FAVORITISM` | Extra weighting assigned to exact matches. | 0.0 |
| `DISABLE_PHONETIC_FILTERING` | Force scoring search terms against every indexed record. | `false` |
| `LENGTH_DIFFERENCE_CUTOFF_FACTOR` | Minimum ratio for the length of two matching tokens, before they score is penalised. | 0.9       |
| `LENGTH_DIFFERENCE_PENALTY_WEIGHT` | Weight of penalty applied to scores when two matching tokens have different lengths. | 0.3    |
| `DIFFERENT_LETTER_PENALTY_WEIGHT` | Weight of penalty applied to scores when two matching tokens begin with different letters. | 0.9   |
| `UNMATCHED_INDEX_TOKEN_WEIGHT` | Weight of penalty applied to scores when part of the indexed name isn't matched. | 0.15    |
| `JARO_WINKLER_BOOST_THRESHOLD` | Jaro-Winkler boost threshold. | 0.7 |
| `JARO_WINKLER_PREFIX_SIZE` | Jaro-Winkler prefix size. | 4 |
| `LOG_FORMAT` | Format for logging lines to be written as. | Options: `json`, `plain` - Default: `plain` |
| `BASE_PATH` | HTTP path to serve API and web UI from. | `/` |
| `HTTP_BIND_ADDRESS` | Address to bind HTTP server on. This overrides the command-line flag `-http.addr`. | Default: `:8084` |
| `HTTP_ADMIN_BIND_ADDRESS` | Address to bind admin HTTP server on. This overrides the command-line flag `-admin.addr`. | Default: `:9094` |
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `DISABLE_WEB_UI` | Skip serving and setup of the web UI. | Default: `false` |
| `WEB_ROOT` | Directory to serve web UI from. | Default: `webui/` |
| `WEBHOOK_MAX_WORKERS` | Maximum number of workers processing webhooks. | Default: 10 |
| `DOWNLOAD_WEBHOOK_URL` | Optional webhook URL called when data downloads / refreshes occur. | Empty |
| `DOWNLOAD_WEBHOOK_AUTH_TOKEN` | Optional `Authorization` header included on download webhooks. | Empty |

## List configurations

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files. | `https://www.treasury.gov/ofac/downloads/%s` |
| `DPL_DOWNLOAD_TEMPLATE` | HTTP address for downloading the DPL. | `https://www.bis.doc.gov/dpl/%s` |
| `EU_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading EU Consolidated Screening List | Subresource of `webgate.ec.europa.eu` |
| `WITH_OFAC_LIST` | Download and parse the US OFAC List | Default: `true` |
| `WITH_US_DPL_LIST` | Download and parse the US Denied Persons List (DPL) | Default: `true` |
| `WITH_US_CSL_SANCTIONS_LIST` | Download and parse the US Consolidated Screening List | Default: `true` |
| `WITH_EU_SCREENING_LIST` | Download and parse the EU Consolidated Screening List | Default: `true` |
| `WITH_UK_CSL_SANCTIONS_LIST` | Download and parse the UK CSL Sanctions List on startup. | Default: `true` |
| `UK_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading UK Consolidated Screening List | Subresource of `www.gov.uk` |
| `UK_SANCTIONS_LIST_URL` | Use an alternate URL for downloading UK Sanctions List | Subresource of `www.gov.uk` |
| `WITH_UK_SANCTIONS_LIST` | Download and parse the UK Sanctions List on startup. | Default: `false` |
| `US_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading US Consolidated Screening List | Subresource of `api.trade.gov` |
| `CSL_DOWNLOAD_TEMPLATE` | Same as `US_CSL_DOWNLOAD_URL` | |
| `KEEP_STOPWORDS` | Boolean to keep stopwords in names. | `false` |
| `DEBUG_NAME_PIPELINE` | Boolean to print debug messages for each name (SDN, SSI) processing step. | `false` |

## Data persistence

By design, Watchman  **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.
