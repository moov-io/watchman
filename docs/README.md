## What is Watchman

Moov Watchman is an HTTP API and Go library to download, parse and offer search functions over numerous trade sanction lists from the United States, European Union governments, agencies, and non profits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.

[Live Demo](https://demo.moov.io/watchman/) | [API Endpoints](https://moov-io.github.io/watchman/api/) | [Admin API Endpoints](https://moov-io.github.io/watchman/admin/)

![](images/webui.png)

## Running Moov Watchman

You can download a [binary from GitHub](https://github.com/moov-io/watchman/releases) or a [Docker image](https://hub.docker.com/r/moov/watchman) for Watchman. Once downloaded you can start making requests against Watchman. The service will download the latest data on startup.

```
$ docker run -p 8084:8084 -p 9094:9094 -it moov/watchman:latest
ts=2019-10-01T20:35:31.301254Z caller=main.go:54 startup="Starting watchman server version v0.13.0"
ts=2019-10-01T20:35:31.301338Z caller=database.go:18 database="looking for  database provider"
ts=2019-10-01T20:35:31.301376Z caller=sqlite.go:119 main="sqlite version 3.25.2"
ts=2019-10-01T20:35:31.302651Z caller=download.go:80 download="Starting refresh of data"
ts=2019-10-01T20:35:31.302651Z caller=main.go:118 admin="listening on :9094"
ts=2019-10-01T20:35:31.530729Z caller=download.go:132 download="Finished refresh of data"
ts=2019-10-01T20:35:31.532927Z caller=main.go:142 main="data refreshed - Addresses=11696 AltNames=9682 SDNs=7379 DeniedPersons=547"
ts=2019-10-01T20:35:31.532962Z caller=main.go:218 main="Setting data refresh interval to 12h0m0s (default)"
ts=2019-10-01T20:35:31.533312Z caller=main.go:182 startup="binding to :8084 for HTTP server"

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

An SDN (or entity) is an individual, group, or company which has or could do business with United States companies or individuals. US law requires checking OFAC data before transactions.

## Web Interface

Moov Watchman provides a web interface for easy browsing of the SDN and related data for mobile and desktop clients. Simply load the address of Watchman in a browser.

## API Documentation

See our documentation for Watchman's [API](https://moov-io.github.io/watchman/api/) or [admin endpoints](https://api.moov.io/admin/watchman/).

## Webhooks

Watchman supports registering a callback url (also called [webhook](https://en.wikipedia.org/wiki/Webhook)) for searches or a given entity ID. (API docs: [company](https://api.moov.io/#operation/addCompanyWatch) or [customers](https://api.moov.io/#operation/addCustomerWatch)) This allows services to monitor for changes to the OFAC data. There's an example [app that receives webhooks](https://github.com/moov-io/watchman/blob/master/examples/webhook/webhook.go) written in Go. Watchman sends either a [Company](https://godoc.org/github.com/moov-io/watchman/client#OFacCompany) or [Customer](https://godoc.org/github.com/moov-io/watchman/client#OfacCustomer) model in JSON to the webhook URL.

Webhook URLs MUST be secure (https://...) and an `Authorization` header is sent with an auth token provided when setting up the webhook. Callers should always verify this auth token matches what was originally provided.

## FAQ

<dl>
<dt>How are entities from the list indexed and used in search</dt>
<dd>
  Entities from sanction lists and other data files are folded through various pre-computations prior to inclusion in the search index.
  What this means is the following steps (in order):

   - SDN Reordering
      - Each SDN name of an individual is re-ordered from "MADURO MOROS, Nicolas" to "Nicolas MADURO MOROS"
   - Company Name Cleanup
      - Suffixes from company names such as: "CO.", "INC.", "L.L.C.", etc are removed
   - Stopword Removal
      - [Stopwords](https://en.wikipedia.org/wiki/Stop_words) are removed. See [bbalet/stopwords](https://github.com/bbalet/stopwords) for a full list of supported languages and words subject to removal.
   - UTF-8 Normalization
      - Punctuation is removed along with extra spaces on both ends of the entity name.
      - Using [Go's /x/text normalization](https://godoc.org/golang.org/x/text/unicode/norm#Form) methods we consolidate entity names and search queries for better searching across multiple languages.

</dd>
</dl>

### Links

- [US Sanctions Search FAQ](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/ques_index.aspx)
- [US Sanctions Search General Questions](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/faq_general.aspx)
- [US Sanctions Compliance](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/faq_compliance.aspx)
- [US Sanction List and Files](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/faq_lists.aspx)
- [US DEPARTMENT OF THE TREASURY](https://www.treasury.gov/resource-center/faqs/Sanctions/Pages/faq_general.aspx#basic)

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://moov-io.github.io/watchman/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#watchman`) to have an interactive discussion about the development of the project.
