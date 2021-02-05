# Moov Watchman

**[GitHub](https://github.com/moov-io/watchman)** | **[Running](https://github.com/moov-io/watchman#usage)** | **[Configuration](https://github.com/moov-io/watchman#configuration-settings)**

## Purpose

Moov Watchman is an HTTP API and Go library that offers download, parse, and search functions over numerous trade sanction lists from the United States, agencies, and nonprofits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.

## Table of Contents

- [Searching](./search.md)
- [Production Runbook](./runbook.md)
   - [Prometheus Metrics](./metrics.md)
- [Pre-compute Pipeline](./pipeline.md)
- [High Availability](./ha.md)
- [File Structure](./file-structure.md)

## Web Interface

Moov Watchman provides a [web interface](./webui.md) for easy browsing of the SDN and related data for mobile and desktop clients. Simply load the address of Watchman in a browser.

![](images/webui.png)

## API Documentation

See our documentation for Watchman's [API](https://moov-io.github.io/watchman/api/) or [admin endpoints](https://moov-io.github.io/watchman/admin/).

## Webhooks

Watchman supports registering a callback URL (also called a [webhook](https://en.wikipedia.org/wiki/Webhook)) for searches or a given entity ID ([Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/-companyID-/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/-customerID-/watch)). This allows services to monitor for changes to the OFAC data. There's an example [app that receives webhooks](https://github.com/moov-io/watchman/blob/master/examples/webhook/webhook.go) written in Go. Watchman sends either a [Company](https://godoc.org/github.com/moov-io/watchman/client#OfacCompany) or [Customer](https://godoc.org/github.com/moov-io/watchman/client#OfacCustomer) model in JSON to the webhook URL.

Webhook URLs MUST be secure (https://...) and an `Authorization` header is sent with an auth token provided when setting up the webhook. Callers should always verify this auth token matches what was originally provided.

## FAQ

<ul>
    <li>How are entities from the list indexed and used in search?</li>
    <ul>
        <li>
            Entities from sanction lists and other data files are folded through various pre-computations prior to inclusion in the search index.
            This means the following steps will occur (in order):
            <p>
                <strong>SDN Reordering</strong><br />
                Each individual's SDN name is re-ordered (Example: from "MADURO MOROS, Nicolas" to "Nicolas MADURO MOROS").
            </p>
            <p>
                <strong>Company Name Cleanup</strong><br />
                Suffixes from company names such as: "CO.", "INC.", "L.L.C.", etc are removed.
            </p>
            <p>
                <strong>Stopword Removal</strong><br />
                <a href="https://en.wikipedia.org/wiki/Stop_words">Stopwords</a> are removed. See <a href="https://github.com/bbalet/stopwords">bbalet/stopwords</a> for a full list of supported languages and words subject to removal.
            </p>
            <p>
                <strong>UTF-8 Normalization</strong><br />
                Punctuation is removed along with extra spaces on both ends of the entity name.
                Using <a href="https://godoc.org/golang.org/x/text/unicode/norm#Form">Go's /x/text normalization</a> methods we consolidate entity names and search queries for better searching across multiple languages.
            </p>
        </li>
    </ul>
    <li>Why are exact matches of words not ranked higher?</li>
    <ul>
        <li>
            Watchman offers an environmental variable called `EXACT_MATCH_FAVORITISM` that can adjust the weight of exact matches within a query.
            This value is a percentage (float64) added to exact matches prior to computing the final match percentage. Try using 0.1, 0.25 or 0.5 with your testing.
        </li>
    </ul>
</ul>

## Links

- [US Sanctions Search FAQ](https://home.treasury.gov/policy-issues/financial-sanctions/faqs#basic)
- [OFAC Sanctions Search Page](https://sanctionssearch.ofac.treas.gov/)

## Getting Help

 channel | info
 ------- | -------
 [Project Documentation](https://moov-io.github.io/watchman/) | Our project documentation available online.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/watchman/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#watchman`) to have an interactive discussion about the development of the project.
