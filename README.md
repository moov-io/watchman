moov-io/ofac
===

[![GoDoc](https://godoc.org/github.com/moov-io/ofac?status.svg)](https://godoc.org/github.com/moov-io/ofac)
[![Build Status](https://travis-ci.com/moov-io/ofac.svg?branch=master)](https://travis-ci.com/moov-io/ofac)
[![Coverage Status](https://codecov.io/gh/moov-io/ofac/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/ofac)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ofac)](https://goreportcard.com/report/github.com/moov-io/ofac)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ofac/master/LICENSE)

Office of Foreign Asset Control (OFAC) is an HTTP API and Go library to download, [parse and serve United States OFAC sanction data](https://docs.moov.io/en/latest/ofac/file-structure/) for applications and humans. Also supported is an async webhook notification service to initiate processes on remote systems connected with OFAC.

To get started using OFAC download [the latest release](https://github.com/moov-io/ofac/releases) or our [Docker image](https://hub.docker.com/r/moov/ofac/tags). We also have a [demo OFAC instance](https://moov.io/ofac/) as part of Moov's demo environment.

```
# Run as a binary
$ wget https://github.com/moov-io/ofac/releases/download/v0.7.0/ofac-darwin-amd64
$ chmod +x ofac-darwin-amd64
$ ./ofac-darwin-amd64
ts=2019-02-05T00:03:31.9583844Z caller=main.go:42 startup="Starting ofac server version v0.7.0"
...

# Run as a Docker image
$ docker run -p 8084:8084 -p 9094:9094 -it moov/ofac:v0.7.0
ts=2019-02-05T00:03:31.9583844Z caller=main.go:42 startup="Starting ofac server version v0.7.0"
...

# Perform a basic search
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

We offer [hosted api docs as part of Moov's tools](https://api.moov.io/#tag/OFAC) and an [OpenAPI specification](https://github.com/moov-io/ofac/blob/master/openapi.yaml) for use with generated clients.

### Configuration

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `OFAC_DATA_REFRESH` | Interval for OFAC data redownload and reparse. | 12h |
| `OFAC_DOWNLOAD_TEMPLATE` | HTTP address for downloading raw OFAC files. | (OFAC website) |
| `SQLITE_DB_PATH`| Local filepath location for the paygate SQLite database. | `ofac.db` |
| `WEBHOOK_BATCH_SIZE` | How many watches to read from database per batch of async searches. | 100 |

### Features

- Download OFAC data on startup
  - Admin endpoint to [manually refresh OFAC data](docs/runbook.md#force-ofac-data-refresh)
- Index data for searches
- Async searches and notifications (webhooks)
- Manual overrides to mark a `Company` or `Customer` as `unsafe` (blocked) or `exception` (never blocked).
- Library to download and parse OFAC files

#### Webhook Notifications

When OFAC sends a [webhook](https://en.wikipedia.org/wiki/Webhook) to your application the body will contain a JSON representation of the [Company](https://godoc.org/github.com/moov-io/ofac/client#Company) or [Customer](https://godoc.org/github.com/moov-io/ofac/client#Customer) model as the body to a POST request. You can see an [example in Go](examples/webhook/webhook.go).

An `Authorization` header will also be sent with the `authToken` provided when setting up the watch. Clients should verify this token to ensure authenticated communicated.

##### Watching a specific Customer or Company by ID

OFAC supports sending a webhook periodically when a specific [Company](https://api.moov.io/#operation/addCompanyWatch) or [Customer](https://api.moov.io/#operation/addCustomerWatch) is to be watched. This is designed to update another system about an OFAC entry's sanction status.

##### Watching a customer or company name

OFAC supports sending a webhook periodically with a free-form name of a [Company](https://api.moov.io/#operation/addCompanyNameWatch) or [Customer](https://api.moov.io/#operation/addCustomerNameWatch). This allows external applications to be notified when an entity matching that name is added to the OFAC list.

## Getting Help

We maintain a [runbook for common issues](docs/runbook.md) and configuration options. Also, if you've encountered a security issue please contact us at [`security@moov.io`](mailto:security@moov.io).

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/en/latest/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher, but we ship the vendor directory in our repository.

## Links

- [Sanctions Search Page](https://sanctionssearch.ofac.treas.gov/)
- [Subscribe for OFAC updates](https://service.govdelivery.com/accounts/USTREAS/subscriber/new)
- [When should I call the OFAC Hotline?](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/directions.aspx)

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
