[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://moov-io.github.io/watchman/">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/watchman/api/#overview">API Endpoints</a>
  ·
  <a href="https://moov.io/blog/education/watchman-api-guide/">API Guide</a>
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
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/watchman)](https://hub.docker.com/r/moov/watchman)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/watchman

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

## What is Watchman?

Moov Watchman is a high-performance sanctions screening and compliance tool that helps businesses meet their regulatory obligations. It provides an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/watchman/pkg/search#Client) for searching against multiple global sanctions and screening lists.

## Key Features

- **Comprehensive Coverage**: Integrates multiple global watchlists in one unified system
- **High-Performance Search**: Optimized for speed and accuracy using advanced matching algorithms
- **Flexible Integration**: HTTP API and Go library for easy integration into your systems
- **Automated Updates**: Regular refreshes of watchlist data to ensure compliance

## Included Lists

Watchman integrates the following lists to help you maintain global compliance:

| Source | List |
|--------|------|
| US Treasury | [Office of Foreign Assets Control (OFAC)](https://ofac.treasury.gov/sanctions-list-service) |
| US Government | [Consolidated Screening List (CSL)](https://www.trade.gov/consolidated-screening-list) |

### Future Lists

The v0.5X series of Watchman has revamped its search engine. The following lists are being re-added into Watchman.

| Source | List |
|--------|------|
| European Union | [Consolidated Sanctions List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en) |
| United Kingdom | [OFSI Sanctions List](https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents) |
| United Kingdom | [Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list) (Disabled by default) |

## v2 Endpoints (v0.5X series and beyond)

The v0.5X series of Watchman has introduced a new v2 search endpoint and removed the older endpoint. This was done to offer a unified response model, improve overall performance, and work towards a stable v1.0 release.

We encourage you to try out the new Watchman and [report any issues or requests in slack](https://slack.moov.io) (`#watchman` channel).

## Table of contents

- [Project status](#project-status)
- [Usage](#usage)
  - As an API
    - [Docker](#docker) ([Config](#configuration-settings))
    - [Data persistence](#data-persistence)
- [Configuration](#configuration)
   - [Similarity Configuration](#similarity-configuration)
   - [Source List Configuration](#source-list-configuration)
- [Useful resources](#useful-resources)
- [FAQ](#faq)
- [Getting help](#getting-help)
- [Supported and tested platforms](#supported-and-tested-platforms)
- [Contributing](#contributing)
- [Related projects](#related-projects)
- [License](#license)

## Project status

Moov Watchman is actively used in multiple production environments. Please star the project if you are interested in its progress. If you have layers above Watchman to simplify tasks, perform business operations, or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

The Watchman project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/watchman/pkg/search#Client) for searching against Watchman.

Government lists are downloaded (and refreshed), parsed, prepared, normalized, and indexed in-memory. Searches operate concurrently and do not require an external database or connection.

### Docker

We publish a [public Docker image `moov/watchman`](https://hub.docker.com/r/moov/watchman/) from Docker Hub or use this repository. No configuration is required to serve on `:8084`. We also have Docker images for [OpenShift](https://quay.io/repository/moov/watchman?tab=tags) published as `quay.io/moov/watchman`. Lastly, we offer a `moov/watchman:static` Docker image with files from 2019. This image can be useful for faster local testing or consistent results.

Pull & start the Docker image:
```
docker pull moov/watchman:latest
docker run -p 8084:8084 moov/watchman:latest
```

Run a search for an individual or business:
```
curl -s "http://localhost:8084/v2/search?name=Nicolas+Maduro&type=person&limit=1&minMatch=0.75" | jq .
```

<details>

```json
{
  "entities": [
    {
      "name": "Nicolas MADURO MOROS",
      "entityType": "person",
      "sourceList": "us_ofac",
      "sourceID": "22790",
      "person": {
        "name": "Nicolas MADURO MOROS",
        "altNames": null,
        "gender": "male",
        "birthDate": "1962-11-23T00:00:00Z",
        "deathDate": null,
        "titles": [
          "President of the Bolivarian Republic of Venezuela"
        ],
        "governmentIDs": [
          {
            "type": "cedula",
            "country": "Venezuela",
            "identifier": "5892464"
          }
        ]
      },
      "business": null,
      "organization": null,
      "aircraft": null,
      "vessel": null,
      "contact": {
        "emailAddresses": null,
        "phoneNumbers": null,
        "faxNumbers": null,
        "websites": null
      },
      "addresses": null,
      "cryptoAddresses": null,
      "affiliations": null,
      "sanctionsInfo": null,
      "historicalInfo": null,
      "sourceData": {
        "entityID": "22790",
        "sdnName": "MADURO MOROS, Nicolas",
        "sdnType": "individual",
        "program": [
          "VENEZUELA",
          "IRAN-CON-ARMS-EO"
        ],
        "title": "President of the Bolivarian Republic of Venezuela",
        "callSign": "",
        "vesselType": "",
        "tonnage": "",
        "grossRegisteredTonnage": "",
        "vesselFlag": "",
        "vesselOwner": "",
        "remarks": "DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela."
      },
      "match": 0.7784062500000001
    }
  ]
}
```

</details>

### Data persistence

By design, Watchman **does not persist** (save) any data about the search queries or actions created. The only storage occurs in memory of the process and upon restart Watchman will have no files or data saved. Also, no in-memory encryption of the data is performed.

## Configuration

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

### FAQ

#### Reporting hits to OFAC

OFAC requires [reporting of positive hits](https://ofac.treasury.gov/ofac-reporting-system). Work with your Financial Institution for complete details.

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

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/watchman/contribute) for something to help out with.

Building Watchman's source code follows standard Go commands. You can use `make build` to compile the code and `make check` to run linters and tests.

Run `make install` to setup [gopostal](https://github.com/openvenues/gopostal) / [libpostal](https://github.com/openvenues/libpostal) for Watchman.

## Related projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
