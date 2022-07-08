---
layout: page
title: API configuration
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Configuration settings

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `DATA_REFRESH_INTERVAL` | Interval for data redownload and reparse. `off` disables this refreshing. | 12h |
| `INITIAL_DATA_DIRECTORY` | Directory filepath with initial files to use instead of downloading. Periodic downloads will replace the initial files. | Empty |
| `EXACT_MATCH_FAVORITISM` | Extra weighting assigned to exact matches. | 0.0 |
| `JARO_WINKLER_BOOST_THRESHOLD` | Jaro-Winkler boost threshold. | 0.7 |
| `JARO_WINKLER_PREFIX_SIZE` | Jaro-Winkler prefix size. | 4 |
| `WEBHOOK_BATCH_SIZE` | How many watches to read from database per batch of async searches. | 100 |
| `LOG_FORMAT` | Format for logging lines to be written as. | Options: `json`, `plain` - Default: `plain` |
| `BASE_PATH` | HTTP path to serve API and web UI from. | `/` |
| `HTTP_BIND_ADDRESS` | Address to bind HTTP server on. This overrides the command-line flag `-http.addr`. | Default: `:8084` |
| `HTTP_ADMIN_BIND_ADDRESS` | Address to bind admin HTTP server on. This overrides the command-line flag `-admin.addr`. | Default: `:9094` |
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `DATABASE_TYPE` | Which database option to use (Options: `sqlite`, `mysql`). | Default: `sqlite` |
| `WEB_ROOT` | Directory to serve web UI from. | Default: `webui/` |
| `WEBHOOK_MAX_WORKERS` | Maximum number of workers processing webhooks. | Default: 10 |
| `DOWNLOAD_WEBHOOK_URL` | Optional webhook URL called when data downloads / refreshes occur. | Empty |
| `DOWNLOAD_WEBHOOK_AUTH_TOKEN` | Optional `Authorization` header included on download webhooks. | Empty |

## List configurations

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files. | `https://www.treasury.gov/ofac/downloads/%s` |
| `DPL_DOWNLOAD_TEMPLATE` | HTTP address for downloading the DPL. | `https://www.bis.doc.gov/dpl/%s` |
| `CSL_DOWNLOAD_TEMPLATE` | HTTP address for downloading the Consolidated Screening List (CSL), which is a collection of US government sanctions lists. | `https://api.trade.gov/consolidated_screening_list/%s` |
| `KEEP_STOPWORDS` | Boolean to keep stopwords in names. | `false` |
| `DEBUG_NAME_PIPELINE` | Boolean to pring debug messages for each name (SDN, SSI) processing step. | `false` |

## Storage

Based on `DATABASE_TYPE`, the following environmental variables will be read to configure connections for a specific database.

### MySQL

- `MYSQL_ADDRESS`: TCP address for connecting to the MySQL server. (example: `tcp(hostname:3306)`)
- `MYSQL_DATABASE`: Name of database to connect into.
- `MYSQL_PASSWORD`: Password of user account for authentication.
- `MYSQL_USER`: Username used for authentication.

Refer to the MySQL driver documentation for [connection parameters](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

- `MYSQL_TIMEOUT`: Timeout parameter specified on (DSN) data source name. (Default: `30s`)

### SQLite

- `SQLITE_DB_PATH`: Local filepath location for the paygate SQLite database. (Default: `watchman.db`)

Refer to the SQLite driver documentation for [connection parameters](https://github.com/mattn/go-sqlite3#connection-string).

## Data persistence

By design, Watchman  **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.
