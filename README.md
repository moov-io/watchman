[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/watchman/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/watchman/api/#overview">API Endpoints</a> <a href="https://moov-io.github.io/watchman/admin/">(Admin Endpoints)</a>
  ·
  <a href="https://moov.io/blog/education/watchman-api-guide/">API Guide</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/watchman?status.svg)](https://godoc.org/github.com/moov-io/watchman)
[![Build Status](https://github.com/moov-io/watchman/workflows/Go/badge.svg)](https://github.com/moov-io/watchman/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/watchman/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/watchman)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/watchman)](https://goreportcard.com/report/github.com/moov-io/watchman)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/watchman?label=project%20size)](https://github.com/moov-io/watchman)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/watchman)](https://hub.docker.com/r/moov/watchman)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/watchman)](https://github.com/moov-io/watchman)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/watchman

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Moov Watchman offers download, parse, and search functions over numerous trade sanction lists from the United States, agencies, and nonprofits for complying with regional laws. Also included is a [web UI](#in-browser-watchman-search) and an async [webhook notification service](#webhook-notifications) to initiate processes on remote systems.

Lists included in search are:

- US Treasury - Office of Foreign Assets Control
  - [Specially Designated Nationals](https://home.treasury.gov/policy-issues/financial-sanctions/specially-designated-nationals-and-blocked-persons-list-sdn-human-readable-lists)
    - Includes SDN, SDN Alternative Names, SDN Addresses
- [United States Consolidated Screening List](https://www.export.gov/article2?id=Consolidated-Screening-List)
   - Department of Commerce – Bureau of Industry and Security
      - [Denied Persons List](http://www.bis.doc.gov/dpl/default.shtm)
      - [Unverified List](http://www.bis.doc.gov/enforcement/unverifiedlist/unverified_parties.html)
      - [Entity List](http://www.bis.doc.gov/entities/default.htm)
   - Department of State – Bureau of International Security and Non-proliferation
      - [Nonproliferation Sanctions](http://www.state.gov/t/isn/c15231.htm)
   - Department of State – Directorate of Defense Trade Controls
      - ITAR Debarred (DTC)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Specially Designated Nationals List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/default.aspx)
      - [Foreign Sanctions Evaders List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/fse_list.aspx)
      - [Sectoral Sanctions Identifications List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/ssi_list.aspx)
      - [Palestinian Legislative Council List](https://www.treasury.gov/resource-center/sanctions/Terrorism-Proliferation-Narcotics/Pages/pa.aspx)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Sectoral Sanctions Identifications List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/ssi_list.aspx)

All United States and European Union companies are required to comply with various regulations and sanction lists (such as the US Patriot Act requiring compliance with the BIS Denied Persons List). Moov's primary usage for this project is with ACH origination in our [paygate](https://github.com/moov-io/paygate) project.

## Table of contents

- [Project status](#project-status)
- [Usage](#usage)
  - As an API
    - [Docker](#docker) ([Config](#configuration-settings))
    - [Google Cloud](#google-cloud-run) ([Config](#configuration-settings))
    - [Webhook notifications](#webhook-notifications)
    - [Data persistence](#data-persistence)
  - [As a Go module](#go-library)
  - [As an in-browser search tool](#in-browser-watchman-search)
- [Useful resources](#useful-resources)
- [Getting help](#getting-help)
- [Supported and tested platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [Related projects](#related-projects)

## Project status

Moov Watchman is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above Watchman to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

The Watchman project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/watchman) for searching, parsing, and downloading lists. We also have an [example](https://pkg.go.dev/github.com/moov-io/watchman/examples) of the webhook service. Below, you can find a detailed list of features offered by Watchman:

- Download OFAC, BIS Denied Persons List (DPL), and various other data sources on startup
  - Admin endpoint to [manually refresh OFAC and DPL data](docs/runbook.md#force-data-refresh)
- Index data for searches
- Async searches and notifications (webhooks)
- Manual overrides to mark a `Company` or `Customer` as `unsafe` (blocked) or `exception` (never blocked).
- Library for OFAC and BIS DPL data to download and parse their custom files

### Docker

We publish a [public Docker image `moov/watchman`](https://hub.docker.com/r/moov/watchman/) from Docker Hub or use this repository. No configuration is required to serve on `:8084` and metrics at `:9094/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/watchman?tab=tags) published as `quay.io/moov/watchman`. Lastly, we offer a `moov/watchman:static` Docker image with files from 2019. This image can be useful for faster local testing or consistent results.

Pull & start the Docker image:
```
docker pull moov/watchman:latest
docker run -p 8084:8084 -p 9094:9094 moov/watchman:latest
```

Get information about a company using their entity ID:
```
curl "localhost:8084/ofac/companies/13374"
```
```
{
   "id":"13374",
   "sdn":{
      "entityID":"13374",
      "sdnName":"SYRONICS",
      "sdnType":"",
      "program":[
         "NPWMD"
      ],
      "title":"",
      "callSign":"",
      "vesselType":"",
      "tonnage":"",
      "grossRegisteredTonnage":"",
      "vesselFlag":"",
      "vesselOwner":"",
      "remarks":""
   },
   "addresses":[
      {
         "entityID":"13374",
         "addressID":"21360",
         "address":"Kaboon Street, PO Box 5966",
         "cityStateProvincePostalCode":"Damascus",
         "country":"Syria",
         "addressRemarks":""
      }
   ],
   "alts":[
      {
         "entityID":"13374",
         "alternateID":"15056",
         "alternateType":"aka",
         "alternateName":"SYRIAN ARAB CO. FOR ELECTRONIC INDUSTRIES",
         "alternateRemarks":""
      }
   ],
   "comments":null,
   "status":null
}
```

### Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-watchman-demo
```

Enable the [Container Registry](https://cloud.google.com/container-registry) API for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following Docker commands, substituting your unique project ID:

```
docker pull moov/watchman
docker tag moov/watchman gcr.io/<PROJECT-ID>/watchman
docker push gcr.io/<PROJECT-ID>/watchman
```

Deploy the container to Cloud Run:
```
gcloud run deploy --image gcr.io/<PROJECT-ID>/watchman --port 8084
```

Select your target platform to `1`, service name to `watchman`, and region to the one closest to you (enable Google API service if a prompt appears). Upon a successful build you will be given a URL where the API has been deployed:

```
https://YOUR-WATCHMAN-APP-URL.a.run.app
```

Now you can ping the server:
```
curl https://YOUR-WATCHMAN-APP-URL.a.run.app/ping
```
You should get this response:
```
PONG
```

### Configuration settings

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `DATA_REFRESH_INTERVAL` | Interval for data redownload and reparse. `off` disables this refreshing. | 12h |
| `INITIAL_DATA_DIRECTORY` | Directory filepath with initial files to use instead of downloading. Periodic downloads will replace the initial files. | Empty |
| `EXACT_MATCH_FAVORITISM` | Extra weighting assigned to exact matches. | 0.0 |
| `JARO_WINKLER_BOOST_THRESHOLD` | Jaro-Winkler boost threshold. | 0.7 |
| `JARO_WINKLER_PREFIX_SIZE` | Jaro-Winkler prefix size. | 4 |
| `WEBHOOK_BATCH_SIZE` | How many watches to read from database per batch of async searches. | 100 |
| `LOG_FORMAT` | Format for logging lines to be written as. | Options: `json`, `plain` - Default: `plain` |
| `LOG_LEVEL` | Level of logging to emit. | Options: `trace`, `info` - Default: `info` |
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

#### List configurations

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files. | `https://www.treasury.gov/ofac/downloads/%s` |
| `DPL_DOWNLOAD_TEMPLATE` | HTTP address for downloading the DPL. | `https://www.bis.doc.gov/dpl/%s` |
| `CSL_DOWNLOAD_TEMPLATE` | HTTP address for downloading the Consolidated Screening List (CSL), which is a collection of US government sanctions lists. | `https://api.trade.gov/consolidated_screening_list/%s` |
| `KEEP_STOPWORDS` | Boolean to keep stopwords in names. | `false` |
| `DEBUG_NAME_PIPELINE` | Boolean to print debug messages for each name (SDN, SSI) processing step. | `false` |

#### Storage

Based on `DATABASE_TYPE`, the following environmental variables will be read to configure connections for a specific database.

##### MySQL

- `MYSQL_ADDRESS`: TCP address for connecting to the MySQL server. (example: `tcp(hostname:3306)`)
- `MYSQL_DATABASE`: Name of database to connect into.
- `MYSQL_PASSWORD`: Password of user account for authentication.
- `MYSQL_USER`: Username used for authentication.

Refer to the MySQL driver documentation for [connection parameters](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

- `MYSQL_TIMEOUT`: Timeout parameter specified on (DSN) data source name. (Default: `30s`)

##### SQLite

- `SQLITE_DB_PATH`: Local filepath location for the paygate SQLite database. (Default: `watchman.db`)

Refer to the SQLite driver documentation for [connection parameters](https://github.com/mattn/go-sqlite3#connection-string).

#### Webhook notifications

When Watchman sends a [webhook](https://en.wikipedia.org/wiki/Webhook) to your application, the body will contain a JSON representation of the [Company](https://godoc.org/github.com/moov-io/watchman/client#OfacCompany) or [Customer](https://godoc.org/github.com/moov-io/watchman/client#OfacCustomer) model as the body to a POST request. You can see an [example in Go](examples/webhook/webhook.go).

An `Authorization` header will also be sent with the `authToken` provided when setting up the watch. Clients should verify this token to ensure authenticated communication.

Webhook notifications are run after the OFAC data is successfully refreshed, which is determined by the `DATA_REFRESH_INTERVAL` environmental variable.

##### Downloads

Moov Watchman supports sending a webhook when the underlying data is refreshed. The body will be the count of entities indexed for each list. The body will be in JSON format and the same schema as the manual data refresh endpoint.

##### Watching a specific customer or company by ID

Moov Watchman supports sending a webhook periodically when a specific [Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/-companyID-/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/-customerID-/watch) is to be watched. This is designed to update another system about an OFAC entry's sanction status.

##### Watching a customer or company name

Moov Watchman supports sending a webhook periodically with a free-form name of a [Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/watch). This allows external applications to be notified when an entity matching that name is added to the OFAC list. The match percentage will be included in the JSON payload.

#### Prometheus metrics

- `http_response_duration_seconds`: A histogram of HTTP response timings.
- `last_data_refresh_success`: Unix timestamp of when data was last refreshed successfully.
- `last_data_refresh_count`: Count of records for a given sanction or entity list.
- `match_percentages` A histogram which holds the match percentages with a label (`type`) of searches.
   - `type`: Can be address, q, remarksID, name, altName
- `mysql_connections`: Number of MySQL connections and their statuses.
- `sqlite_connections`: Number of SQLite connections and their statuses.

### Data persistence

By design, Watchman  **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.

### Go library

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go v1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/watchman/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/watchman.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/watchman

$ go doc github.com/moov-io/watchman/client Search
```

### In-browser Watchman search

Using our [in-browser utility](https://oss.moov.io/watchman/), you can instantly perform advanced Watchman searches. Simply fill search fields and generate a detailed report that includes match percentage, alternative names, effective/expiration dates, IDs, addresses, and other useful information. This tool is particularly useful for completing quick searches with the aid of a intuitive interface.

## Reporting blocks to OFAC

OFAC requires annual reports of blocked entities and [offers guidance for this report](https://www.treasury.gov/resource-center/sanctions/Documents/ofac_blocked_property_guidance.pdf). Section [31 C.F.R. § 501.603(b)(2)](https://www.ecfr.gov/cgi-bin/text-idx?SID=be4f2a1608abec5d93170fb03af99939&mc=true&node=se31.3.501_1603&rgn=div8) requires this annual report.

## Useful resources

- [OFAC Sanctions Search Page](https://sanctionssearch.ofac.treas.gov/)
- [Subscribe for OFAC email updates](https://service.govdelivery.com/accounts/USTREAS/subscriber/new)
- [When should I call the OFAC Hotline?](https://home.treasury.gov/policy-issues/financial-sanctions/contact-ofac/when-should-i-call-the-ofac-hotline#:~:text=If%20it's%20hitting%20against%20OFAC's,the%20match%20is%20hitting%20against.)
- [BIS Denied Persons List with Denied US Export Privileges (DPL)](https://bis.data.commerce.gov/dataset/Denied-Persons-List-with-Denied-US-Export-Privileg/xwtd-wd7a/data)
- [BIS Entity List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/entity-list)
- [Sectoral Sanctions Identifications (SSI)](https://home.treasury.gov/policy-issues/financial-sanctions/consolidated-sanctions-list/sectoral-sanctions-identifications-ssi-list)
- [US Sanctions Search FAQ](https://home.treasury.gov/policy-issues/financial-sanctions/faqs#basic)

## Getting help

We maintain a [runbook for common issues](docs/runbook.md) and configuration options. Also, if you've encountered a security issue please contact us at [`security@moov.io`](mailto:security@moov.io).

 channel | info
 ------- | -------
[Project Documentation](https://moov-io.github.io/watchman/) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/watchman/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/watchman/contribute) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/watchman/releases/latest) as well. We highly recommend you use a tagged release for production.

### Releasing

To make a release of ach simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Testing

We maintain a comprehensive suite of unit tests and recommend table-driven testing when a particular function warrants several very similar test cases. To run all test files in the current directory, use `go test`. Current overall coverage can be found on [Codecov](https://app.codecov.io/gh/moov-io/watchman/).

## Related projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
