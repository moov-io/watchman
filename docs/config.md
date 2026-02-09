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

  # Database:
  #   DatabaseName: "watchman"
  #   MySQL:
  #     Address: "tcp(mysql:3306)"
  #     User: "watchman"
  #     Password: "watchman"
  #   Postgres:
  #     Address: "postgres:5432"
  #     User: "watchman"
  #     Password: "watchman"
  #     Connections:
  #       MaxOpen: 50
  #       MaxIdle: 50
  #       MaxLifetime: "60s"
  #       MaxIdleTime: "60s"

  Download:
    RefreshInterval: "12h"
    InitialDataDirectory: ""

    # Specify which lists to download and include in Watchman results
    # Examples: us_csl, us_fincen_311, us_non_sdn, us_ofac, etc...
    IncludedLists:
      - "us_csl"
      - "us_ofac"

    # Include any senzing formatted OpenSanctions lists
    # OpenSanctions:
    #   ApiKey: "secret"
    #   Lists:
    #     - SourceList: opensanctions_peps
    #       Location: https://data.opensanctions.org/datasets/latest/us_congress/senzing.json

    # Include any senzing formatted lists in Watchman's corpus
    # Senzing:
    #   - SourceList: "senzing-persons"
    #     Location: "file://pkg/sources/senzing/testdata/persons.jsonl"

  Search:
    # Tune these settings based on your available resources (CPUs, etc).
    # Watchman will dynamically find an optimal goroutine count for faster responses.
    # Usually a multiple (i.e. 2x, 4x) of GOMAXPROCS is optimal.
    Goroutines:
      Default: 10
      Min: 1
      Max: 25

  Embeddings:
    Enabled: false # See below for Cross-Script Embeddings

  Geocoding:
    Enabled: false
    Provider:
      # Specify the provider name, which can be one of "opencage", "google", or "nominatim"
      Name: ""
      ApiKey: ""  # Can also set GEOCODING_API_KEY
      BaseURL: "" # Useful for self-hosted Nominatim instances.
      Timeout: "10s"
    RateLimit:
      RequestsPerSecond: 1.5 # req/s
      Burst: 5
    Cache:
      L1MaxSize: 10000
      L1TTL: "24h"     # time-to-live for L1 cache entries.
      L2Enabled: false # Uses the database connection for a persistent cache.
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

### Included Lists

| List ID           | Name                                       | Source                                                                                   |
|-------------------|--------------------------------------------|------------------------------------------------------------------------------------------|
| `us_csl`          | Consolidated Screening List (CSL)          | [URL](https://www.trade.gov/consolidated-screening-list)                                 |
| `us_fincen_311`   | US FinCEN 311 Actions                      | [URL](https://home.treasury.gov/policy-issues/terrorism-and-illicit-finance/311-actions) |
| `us_non_sdn`      | US Office of Foreign Assets Control (OFAC) | [URL](https://ofac.treasury.gov/sanctions-list-service)                                  |
| `us_ofac`         | US Office of Foreign Assets Control (OFAC) | [URL](https://ofac.treasury.gov/sanctions-list-service)                                  |
| `opensanctions_*` | OpenSanctions Datasets                     | [URL](https://www.opensanctions.org/datasets/)                                           |

### Environment Variables

| Environmental Variable   | Description                                                                                                                          | Default                                     |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| `INCLUDED_LISTS`         | Comma separated list of lists to include.                                                                                            | Empty                                       |
| `DATA_REFRESH_INTERVAL`  | Interval for data redownload and reparse. `off` disables this refreshing.                                                            | 12h                                         |
| `INITIAL_DATA_DIRECTORY` | Directory filepath with initial files to use instead of downloading. Periodic downloads will replace the initial files.              | Empty                                       |
| `HTTPS_CERT_FILE`        | Filepath containing a certificate (or intermediate chain) to be served by the HTTP server. Requires all traffic be over secure HTTP. | Empty                                       |
| `HTTPS_KEY_FILE`         | Filepath of a private key matching the leaf certificate from `HTTPS_CERT_FILE`.                                                      | Empty                                       |
| `LOG_FORMAT`             | Format for logging lines to be written as.                                                                                           | Options: `json`, `plain` - Default: `plain` |
| `LOG_LEVEL`              | Level of logging to emit.                                                                                                            | Options: `trace`, `info` - Default: `info`  |

### TF-IDF Configuration

[TF-IDF](https://en.wikipedia.org/wiki/Tf%E2%80%93idf) (Term Frequency–Inverse Document Frequency) is a technique used to measure the importance of a word in a corpus. Less frequent words are more important.

By default Watchman does not employ TF-IDF, but it can be enabled and configured with the following environmental variables.

| Environmental Variable | Description                                                                                                    | Default |
|------------------------|----------------------------------------------------------------------------------------------------------------|---------|
| `TFIDF_ENABLED`        | Enabled controls whether TF-IDF weighting is applied to name matching.                                         | `false` |
| `TFIDF_SMOOTHING`      | SmoothingFactor (k) in the IDF formula: `log((N+k)/(df+k))`. Prevents division by zero and smooths IDF values. | 1.0     |
| `TFIDF_MIN_IDF`        | MinIDF is the floor for IDF values. Prevents very common terms from having zero or negative weight.            | 0.1     |
| `TFIDF_MAX_IDF`        | MaxIDF is the ceiling for IDF values. Prevents single-occurrence terms from dominating the score.              | 10.0    |

### Cross-Script Embeddings Configuration

Watchman can use neural network embeddings to match names across different writing systems (Arabic, Cyrillic, Chinese, etc.). This feature requires configuring an embeddings API provider (Ollama, OpenAI, OpenRouter, etc.).

See [Cross-Script Name Matching](cross-script-matching.md) for detailed setup instructions.

| Environmental Variable            | Description                                                       | Default |
|-----------------------------------|-------------------------------------------------------------------|---------|
| `EMBEDDINGS_ENABLED`              | Enable embedding-based cross-script search.                       | `false` |
| `EMBEDDINGS_BASE_URL`             | API endpoint URL for embeddings provider (required when enabled). | Empty   |
| `EMBEDDINGS_API_KEY`              | API key for authentication (optional for Ollama).                 | Empty   |
| `EMBEDDINGS_MODEL`                | Model name to use for embeddings (required when enabled).         | Empty   |
| `EMBEDDINGS_DIMENSION`            | Embedding vector dimension (required, must match model).          | `0`     |
| `EMBEDDINGS_CACHE_SIZE`           | Number of embedding vectors to cache in memory.                   | `10000` |
| `EMBEDDINGS_CROSS_SCRIPT_ONLY`    | Only use embeddings for non-Latin queries (recommended).          | `true`  |
| `EMBEDDINGS_SIMILARITY_THRESHOLD` | Minimum similarity score (0.0-1.0) for results.                   | `0.7`   |
| `EMBEDDINGS_BATCH_SIZE`           | Batch size for encoding multiple texts.                           | `32`    |
| `EMBEDDINGS_INDEX_BUILD_TIMEOUT`  | Maximum time allowed for building the embedding index.            | `10m`   |

YAML configuration (example with OpenAI):

```yaml
  Search:
    Embeddings:
      Enabled: true
      Provider:
        Name: "openai"
        BaseURL: "https://api.openai.com/v1"
        APIKey: "${OPENAI_API_KEY}"
        Model: "text-embedding-3-small" # Required - no default
        Dimension: 1536                 # Required - must work with model
      Cache:
        Type: "memory" # or sql, which disregards size
        Size: 10000
      CrossScriptOnly: true
      SimilarityThreshold: 0.7
      BatchSize: 32
      IndexBuildTimeout: "10m"
```

### Similarity Configuration

| Environmental Variable             | Description                                                                                                   | Default |
|------------------------------------|---------------------------------------------------------------------------------------------------------------|---------|
| `SEARCH_GOROUTINE_COUNT`           | Set a fixed number of goroutines used for each search. Default is to dynamically optimize for faster results. | Empty   |
| `KEEP_STOPWORDS`                   | Boolean to keep stopwords in names.                                                                           | `false` |
| `JARO_WINKLER_BOOST_THRESHOLD`     | Jaro-Winkler boost threshold.                                                                                 | 0.7     |
| `JARO_WINKLER_PREFIX_SIZE`         | Jaro-Winkler prefix size.                                                                                     | 4       |
| `LENGTH_DIFFERENCE_CUTOFF_FACTOR`  | Minimum ratio for the length of two matching tokens, before they score is penalised.                          | 0.9     |
| `LENGTH_DIFFERENCE_PENALTY_WEIGHT` | Weight of penalty applied to scores when two matching tokens have different lengths.                          | 0.3     |
| `DIFFERENT_LETTER_PENALTY_WEIGHT`  | Weight of penalty applied to scores when two matching tokens begin with different letters.                    | 0.9     |
| `UNMATCHED_INDEX_TOKEN_WEIGHT`     | Weight of penalty applied to scores when part of the indexed name isn't matched.                              | 0.15    |
| `ADJACENT_SIMILARITY_POSITIONS`    | How many nearby words to search for highest max similarly score.                                              | 3       |
| `EXACT_MATCH_FAVORITISM`           | Extra weighting assigned to exact matches.                                                                    | 0.0     |
| `DISABLE_PHONETIC_FILTERING`       | Force scoring search terms against every indexed record.                                                      | `false` |

#### Source List Configuration

| Environmental Variable | Description                                            | Default |
|------------------------|--------------------------------------------------------|---------|
| `DOWNLOAD_TIMEOUT`     | Duration of time allowed for a list to fully download. | `45s`   |

##### United States

| Environmental Variable   | Description                                                         | Default                                                                         |
|--------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------------------|
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files.                        | `https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/%s` |
| `US_CSL_DOWNLOAD_URL`    | Use an alternate URL for downloading US Consolidated Screening List | Subresource of `api.trade.gov`                                                  |
| `CSL_DOWNLOAD_TEMPLATE`  | Same as `US_CSL_DOWNLOAD_URL`                                       |                                                                                 |

##### European Union

| Environmental Variable | Description                                                         | Default                               |
|------------------------|---------------------------------------------------------------------|---------------------------------------|
| `EU_CSL_TOKEN`         | Token used to download the EU Consolidated Screening List           | `<valid-token>`                       |
| `EU_CSL_DOWNLOAD_URL`  | Use an alternate URL for downloading EU Consolidated Screening List | Subresource of `webgate.ec.europa.eu` |

##### United Kingdom

| Environmental Variable   | Description                                                         | Default                     |
|--------------------------|---------------------------------------------------------------------|-----------------------------|
| `UK_CSL_DOWNLOAD_URL`    | Use an alternate URL for downloading UK Consolidated Screening List | Subresource of `www.gov.uk` |
| `UK_SANCTIONS_LIST_URL`  | Use an alternate URL for downloading UK Sanctions List              | Subresource of `www.gov.uk` |
| `WITH_UK_SANCTIONS_LIST` | Download and parse the UK Sanctions List on startup.                | Default: `false`            |

##### Open Sanctions

| Environmental Variable  | Description                                                                                           | Default |
|-------------------------|-------------------------------------------------------------------------------------------------------|---------|
| `OPENSANCTIONS_API_KEY` | API key for OpenSanctions authenticated access. Suggested when an `opensanctions_*` list is included. | Empty   |

## Data persistence

By design, Watchman **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.
