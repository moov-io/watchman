---
layout: page
title: Webhook notifications
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Webhooks

Watchman supports registering a callback URL (also called a [webhook](https://en.wikipedia.org/wiki/Webhook)) for notifications on list download and reindexing.

Webhook URLs MUST be secure (https://...) and an `Authorization` header is sent with an auth token provided when setting up the webhook. Callers should always verify this auth token matches what was originally provided.

When Watchman sends a [webhook](https://en.wikipedia.org/wiki/Webhook) to your application, the body will contain a JSON representation of the `DownloadStats` model (described below) as the body to a POST request.

An `Authorization` header will also be sent with the `authToken` provided when setting up the watch. Clients should verify this token to ensure authenticated communicated.

Webhook notifications are ran after the OFAC data is successfully refreshed, which is determined by the `DATA_REFRESH_INTERVAL` environmental variable. `WEBHOOK_MAX_WORKERS` can be set to control how many goroutines can process webhooks concurrently

## Download / Refresh

Watchman can notify when the OFAC, CSL, etc lists are downloaded and re-indexed. The address specified at `DOWNLOAD_WEBHOOK_URL` will be sent a POST HTTP request with the following body. An Authorization header can be specified with `DOWNLOAD_WEBHOOK_AUTH_TOKEN`.

```json
{
    "SDNs": 0,
    "altNames": 0,
    "addresses": 0,

	"deniedPersons": 0,
    "bisEntities": 0,
    "militaryEndUsers": 0,
    "sectoralSanctions": 0,
    "unverifiedCSL": 0,
    "nonProliferationSanctions": 0,
    "foreignSanctionsEvaders": 0,
    "palestinianLegislativeCouncil": 0,
    "CAPTA": 0,
    "ITARDebarred": 0,
    "chineseMilitaryIndustrialComplex": 0,
    "nonSDNMenuBasedSanctions": 0,

    "europeanSanctionsList": 0,
    "ukConsolidatedSanctionsList": 0,
	"ukSanctionsList": 0,

    "errors": [
        "CSL: unexpected error 429"
    ],

    "timestamp": "2009-11-10T23:00:00Z"
}
```
