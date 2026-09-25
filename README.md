[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/watchman/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/watchman/using-watchman/">Using Watchman</a>
  ·
  <a href="https://moov-io.github.io/watchman/api/#overview">API Endpoints</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/watchman?utm_source=godoc)](https://pkg.go.dev/github.com/moov-io/watchman/pkg/search#Client)
[![Build Status](https://github.com/moov-io/watchman/workflows/Go/badge.svg)](https://github.com/moov-io/watchman/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/watchman)](https://goreportcard.com/report/github.com/moov-io/watchman)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/watchman/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/watchman)](https://hub.docker.com/r/moov/watchman)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/moov-io/watchman)

# moov-io/watchman

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

## What is Watchman?

Moov Watchman is an open-source **sanctions screening engine**. It downloads OFAC, EU, UK, UN, and related lists, indexes them in memory, and scores each customer or counterparty with an inspectable multi-field matcher. Use the HTTP API, [Go client](https://pkg.go.dev/github.com/moov-io/watchman/pkg/search#Client), WASM UI, or experimental [MCP](https://moov-io.github.io/watchman/mcp/) server.

How to run it: [Using Watchman](https://moov-io.github.io/watchman/using-watchman/). For BSA/AML and sanctions officers: [For compliance and risk](https://moov-io.github.io/watchman/methodology/for-compliance/).

We measured the matcher on 472,477 labeled people, companies, and vessels from [OpenSanctions Pairs](https://moov-io.github.io/watchman/opensanctions-pairs/). At `minMatch=0.80`, almost every returned hit was a real match (precision 0.99). Enabling embeddings for names in different writing systems raised the share of true matches found from 0.69 to 0.82, with precision 0.95.

## Key Features

- **Lists you can name** — OFAC SDN and Non-SDN, US CSL, FinCEN 311, EU, UK, UN, OpenSanctions Senzing files, plus CSV ingest
- **Structured search** — `type` (person, business, organization, vessel, aircraft), name, aliases, government IDs, dates, addresses, crypto, contact
- **How IDs work** — a matching passport, IMO number, or crypto address (with country) scores 1.0; a matching tax number or email raises the score; two national IDs that disagree lower the score
- **You set the cutoff** — `minMatch` is policy (0.80 screening, ~0.59 high recall)
- **Explainable hits** — `debug=true` returns field-level score pieces
- **Fast candidate search** — source/type partitions, name-token and ID indexes, parallel scoring ([Performance](https://moov-io.github.io/watchman/performance/), [Indexing](https://moov-io.github.io/watchman/indexing/))
- **Optional** — TF-IDF, cross-script embeddings, geocoding, deepparse, Senzing request/response format, MCP. Docker images and Linux/macOS GitHub releases parse addresses with libpostal; Windows `.exe` and `go run` use usaddress; any build can enable deepparse.

## Included Lists

Watchman integrates the following lists to help you maintain global compliance. Use the env variable `INCLUDED_LISTS` or [config file](https://moov-io.github.io/watchman/config/#download) to customize which lists are loaded.

| Source            | List                                                                                                                                                                                    |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **OpenSanctions** | [Any Senzing formatted list from OpenSanctions](https://www.opensanctions.org/datasets/)                                                                                                |
| European Union    | [Consolidated Sanctions List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)                        |
| US Government     | [Consolidated Screening List (CSL)](https://www.trade.gov/consolidated-screening-list), [FinCEN 311](https://home.treasury.gov/policy-issues/terrorism-and-illicit-finance/311-actions) |
| US Treasury       | [Office of Foreign Assets Control (OFAC)](https://ofac.treasury.gov/sanctions-list-service) and Non-SDN list                                                                            |
| United Kingdom    | [UK sanctions list](https://www.gov.uk/government/publications/the-uk-sanctions-list)                                                                                                    |
| United Nations    | [Consolidated Sanctions List](https://www.un.org/sc/resources/sc-sanctions)                                                                                                             |

When loading multiple OpenSanctions or custom Senzing-formatted lists, set the `SENZING_CONCURRENT_DOWNLOADS` environment variable to control parallelism during downloads (defaults to 5 concurrent).

## Project status

Moov Watchman is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above Watchman to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

Send `type` (`person`, `business`, `vessel`, …) on every search. It is not a required query parameter: omitting it searches every type, which is slower. A wrong type searches only that partition, so the designated party can be missing from the page. On GET, fields such as `birthDate` and `gov_*` are only read when `type` is set; sending them without `type` is HTTP 400.

See [Using Watchman](https://moov-io.github.io/watchman/using-watchman/), [Search](https://moov-io.github.io/watchman/search/), [Performance](https://moov-io.github.io/watchman/performance/), and [Indexing](https://moov-io.github.io/watchman/indexing/).

### Docker

We publish [`moov/watchman`](https://hub.docker.com/r/moov/watchman/) on Docker Hub and [`quay.io/moov/watchman`](https://quay.io/repository/moov/watchman?tab=tags) for OpenShift. `moov/watchman:v2-static` ships frozen 2019 files for fast local tests. The WASM UI is at `/` on `:8084`.

Start the Docker image [using a tag](https://hub.docker.com/r/moov/watchman/tags):
```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
```

That example publishes only the business API (`:8084`). Do not expose Watchman on the public internet. See [Network access](#network-access).

Search is `GET /v2/search`. Adding fields raises the score. These queries use OFAC SDN 48603 (Dmitry Yuryevich KHOROSHEV). `jq` prints the top hit:

```
# Name only — often under the 0.80 screening line
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":0.767}

# Name + date of birth
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&birthDate=1993-04-17&limit=1" \
  | jq '{name: .entities[0].name, match: (.entities[0].match*1000|round/1000)}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":0.867}

# Name + passport (unique identity key → 1.0)
curl -s "http://localhost:8084/v2/search?type=person&name=Dmitry+Khoroshev&gov_passport=RU:2018278055&limit=1" \
  | jq '{name: .entities[0].name, match: .entities[0].match}'
# {"name":"Dmitry Yuryevich KHOROSHEV","match":1}
```

`minMatch=0.80` is a typical production cutoff. The name-only query above would return no rows at that cutoff; the passport query would. Each entity in `entities` is the list record with a `match` field (0 to 1) on the same object.

### Network access

Watchman is not designed to be served directly on the internet. Run it on a private network or behind a reverse proxy / API gateway. Authentication, ACLs, and rate limiting belong at the edge of the deployment, not inside Watchman.

The HTTP API (`BindAddress`, `:8084`) is unauthenticated by design (search, ingest, export, refresh, web UI, MCP). The admin server (`AdminAddress`, `:9094`) is a **separate port** so you can firewall it, bind it to an internal interface, or block it entirely. Prometheus `/metrics` and `/version` live on the admin port on purpose and are unauthenticated.

Download URLs (`file://` locations, `*_DOWNLOAD_TEMPLATE`, `*_DOWNLOAD_URL`) are operator configuration, not API input. Watchman does not allowlist hosts or jail `file://` paths.

See [Network access](https://moov-io.github.io/watchman/network/), [issue #875](https://github.com/moov-io/watchman/issues/875), and [issue #876](https://github.com/moov-io/watchman/issues/876).

### Data persistence

By design, Watchman **does not persist** search queries. Your application must retain screening logs (who was screened, when, against which list hashes, at which cutoff). Watchman can store ingested files (the individual records) in
a MySQL or PostgreSQL database for concurrent access. External lists that are downloaded on startup (and refreshed periodically) are only kept in memory. No encryption of data in-memory
is performed.

### Download reliability

External sanctions lists are fetched over the network on startup and during periodic refreshes (default every 12h). Government endpoints can be flaky or unavailable, causing startup failures
or refresh errors. For production use, see [Caching Data Files](https://moov-io.github.io/watchman/cache-data-files/) for local pre-loaded files and the companion project
[moov-io/watchman-cache](https://github.com/moov-io/watchman-cache), which provides a nginx reverse proxy + long-lived cache designed to sit in front of Watchman.
It works with the existing `*_DOWNLOAD_TEMPLATE` / `*_DOWNLOAD_URL` environment variables with no changes required to Watchman.

## Configuration

Watchman recommends [file-based configuration](https://moov-io.github.io/watchman/config/) but supports environmental variable options.

### FAQ

#### Reporting hits to OFAC

OFAC requires [reporting of positive hits](https://ofac.treasury.gov/ofac-reporting-system) and has a [voluntary self-disclosure](https://disclosure.ofac.treas.gov/) portal. Work with your Financial Institution for complete details.

#### Useful resources

- [OFAC Sanctions Search Page](https://sanctionssearch.ofac.treas.gov/)
- [Subscribe for OFAC email updates](https://service.govdelivery.com/accounts/USTREAS/subscriber/new)
- [When should I call the OFAC Hotline?](https://home.treasury.gov/policy-issues/financial-sanctions/contact-ofac/when-should-i-call-the-ofac-hotline#:~:text=If%20it's%20hitting%20against%20OFAC's,the%20match%20is%20hitting%20against.)
- [BIS Denied Persons List with Denied US Export Privileges (DPL)](https://bis.data.commerce.gov/dataset/Denied-Persons-List-with-Denied-US-Export-Privileg/xwtd-wd7a/data)
- [BIS Entity List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/entity-list)
- [Sectoral Sanctions Identifications (SSI)](https://home.treasury.gov/policy-issues/financial-sanctions/consolidated-sanctions-list/sectoral-sanctions-identifications-ssi-list)
- [US Sanctions Search FAQ](https://home.treasury.gov/policy-issues/financial-sanctions/faqs#basic)

## Getting help

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/watchman/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/watchman/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#watchman`) to have an interactive discussion about the development of the project.

If you find a security issue please contact us at [`security@moov.io`](mailto:security@moov.io).

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Checkout our [issues for first time contributors](https://github.com/moov-io/watchman/contribute) for something to help out with.

Building Watchman's source code follows standard Go commands. You can use `make build` to compile the code and `make check` to run linters and tests.

Run `make install` to setup [gopostal](https://github.com/moov-io/gopostal) / [libpostal](https://github.com/openvenues/libpostal) for Watchman.

Run `make setup-deepparse` to start an optional [deepparse](https://github.com/GRAAL-Research/deepparse) HTTP sidecar (`ghcr.io/graal-research/deepparse:0.11.0`) as an alternative address parser. It is disabled by default.

## Related projects

As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
