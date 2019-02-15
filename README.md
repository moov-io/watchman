moov-io/ofac
===

[![GoDoc](https://godoc.org/github.com/moov-io/ofac?status.svg)](https://godoc.org/github.com/moov-io/ofac)
[![Build Status](https://travis-ci.com/moov-io/ofac.svg?branch=master)](https://travis-ci.com/moov-io/ofac)
[![Coverage Status](https://codecov.io/gh/moov-io/ofac/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/ofac)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ofac)](https://goreportcard.com/report/github.com/moov-io/ofac)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ofac/master/LICENSE)

*project is under active development and is not production ready*

OFAC is an HTTP API and Go library to download, parse and serve OFAC sanction data for applications and humans. Also supported is an async webhook notification service to initiate processes on remote systems connected with OFAC.

To get started using OFAC download our [Docker image](https://hub.docker.com/r/moov/ofac/tags) or [the latest release](https://github.com/moov-io/ofac/releases).

```
$ docker run -p 8084:8084 -p 9094:9094 -it moov/ofac:v0.2.1
ts=2019-02-05T00:03:31.9583844Z caller=main.go:42 startup="Starting ofac server version v0.2.1"
...

$ curl -s localhost:8084/search?name=...
{
  "SDNs": [
    {
      "entityID": "...",
      "sdnName": "...",
      "sdnType": "...",
      "program": "...",
      "title": "...",
      "callSign": "...",
      "vesselType": "...",
      "tonnage": "...",
      "grossRegisteredTonnage": "...",
      "vesselFlag": "...",
      "vesselOwner": "...",
      "remarks": "..."
    }
  ],
  "altNames": null,
  "addresses": null
}
```


### Configuration

**Search Similarity Metrics**

OFAC computes string similarity using the Levenshtein algorithm and can match sensitivity can be configured with environment variables.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `OFAC_DATA_REFRESH` | Interval for OFAC data redownload and reparse. | 12h |
| `SQLITE_DB_PATH`| Local filepath location for the paygate SQLite database. | `ofac.db` |
| `WEBHOOK_BATCH_SIZE` | How many watches to read from database per batch of async searches. | 100 |

### Features

- Download data on startup
  - admin endpoint to [manually refresh OFAC data](docs/runbook.md#force-ofac-data-refresh)
- Index data for searches
- async searches and notifications (webhooks)
- Library to download and parse OFAC files

#### Webhook Notifications

When OFAC sends a [webhook](https://en.wikipedia.org/wiki/Webhook) to your application the body will contain a [JSON representation of the Customer model](https://godoc.org/github.com/moov-io/ofac/client#Customer) as the body to a POST request. You can see an [example in Go](example/webhook.go).

An `Authorization` header will also be sent with the `authToken` provided when setting up the watch. Clients should verify this token to ensure authenticated communicated.

## Links

- [Sanctions Search Page](https://sanctionssearch.ofac.treas.gov/)
- [Subscribe for OFAC updates](https://service.govdelivery.com/accounts/USTREAS/subscriber/new)
