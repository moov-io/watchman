---
layout: page
title: Configuration
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Configuration

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

### File

Set the `APP_CONFIG` environment variable with a filepath to a yaml file containing the following:

```yaml
Watchman:
  Servers:
    BindAddress: ":8084"
    AdminAddress: ":9094"

  Telemetry:
    ServiceName: "watchman"

  Download:
    RefreshInterval: "12h"
    InitialDataDirectory: ""
    IncludedLists:
      - "us_csl"
      - "us_ofac"

  Search:
    # Tune these settings based on your available resources (CPUs, etc).
    # Watchman will dynamically find an optimal goroutine count for faster responses.
    # Usually a multiple (i.e. 2x, 4x) of GOMAXPROCS is optimal.
    Goroutines:
      Default: 10
      Min: 1
      Max: 25
```

#### PostalPool

PostalPool is an experiment for improving address parsing. It's optional configuration and may be removed in the future.

```yaml
  PostalPool:
    Enabled: false
    Instances: 2
    StartingPort: 10000
    StartupTimeout: "60s"
    RequestTimeout: "10s"
    BinaryPath: "" # POSTAL_SERVER_BIN_PATH is set in Dockerfile
    CGOSelfInstances: 1
```

### Environment Variables

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `DATA_REFRESH_INTERVAL` | Interval for data redownload and reparse. `off` disables this refreshing. | 12h |
| `INITIAL_DATA_DIRECTORY` | Directory filepath with initial files to use instead of downloading. Periodic downloads will replace the initial files. | Empty |
| `HTTPS_CERT_FILE` | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty |
| `HTTPS_KEY_FILE`  | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`. | Empty |
| `LOG_FORMAT` | Format for logging lines to be written as. | Options: `json`, `plain` - Default: `plain` |
| `LOG_LEVEL` | Level of logging to emit. | Options: `trace`, `info` - Default: `info` |

### Similarity Configuration

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `SEARCH_GOROUTINE_COUNT` | Set a fixed number of goroutines used for each search. Default is to dynamically optimize for faster results. | Empty |
| `KEEP_STOPWORDS` | Boolean to keep stopwords in names. | `false` |
| `JARO_WINKLER_BOOST_THRESHOLD` | Jaro-Winkler boost threshold. | 0.7 |
| `JARO_WINKLER_PREFIX_SIZE` | Jaro-Winkler prefix size. | 4 |
| `LENGTH_DIFFERENCE_CUTOFF_FACTOR` | Minimum ratio for the length of two matching tokens, before they score is penalised. | 0.9       |
| `LENGTH_DIFFERENCE_PENALTY_WEIGHT` | Weight of penalty applied to scores when two matching tokens have different lengths. | 0.3    |
| `DIFFERENT_LETTER_PENALTY_WEIGHT` | Weight of penalty applied to scores when two matching tokens begin with different letters. | 0.9   |
| `UNMATCHED_INDEX_TOKEN_WEIGHT` | Weight of penalty applied to scores when part of the indexed name isn't matched. | 0.15    |
| `ADJACENT_SIMILARITY_POSITIONS` | How many nearby words to search for highest max similarly score. | 3 |
| `EXACT_MATCH_FAVORITISM` | Extra weighting assigned to exact matches. | 0.0 |
| `DISABLE_PHONETIC_FILTERING` | Force scoring search terms against every indexed record. | `false` |

#### Source List Configuration

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `DOWNLOAD_TIMEOUT` | Duration of time allowed for a list to fully download. | `45s` |
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files. | `https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/%s` |
| `EU_CSL_TOKEN` | Token used to download the EU Consolidated Screening List | `<valid-token>` |
| `EU_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading EU Consolidated Screening List | Subresource of `webgate.ec.europa.eu` |
| `UK_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading UK Consolidated Screening List | Subresource of `www.gov.uk` |
| `UK_SANCTIONS_LIST_URL` | Use an alternate URL for downloading UK Sanctions List | Subresource of `www.gov.uk` |
| `WITH_UK_SANCTIONS_LIST` | Download and parse the UK Sanctions List on startup. | Default: `false` |
| `US_CSL_DOWNLOAD_URL` | Use an alternate URL for downloading US Consolidated Screening List | Subresource of `api.trade.gov` |
| `CSL_DOWNLOAD_TEMPLATE` | Same as `US_CSL_DOWNLOAD_URL` | |

## Data persistence

By design, Watchman **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.
