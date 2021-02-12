---
layout: page
title: Webhook Notifications
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Webhooks

Watchman supports registering a callback URL (also called a [webhook](https://en.wikipedia.org/wiki/Webhook)) for searches or a given entity ID ([Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/-companyID-/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/-customerID-/watch)). This allows services to monitor for changes to the OFAC data. There's an example [app that receives webhooks](https://github.com/moov-io/watchman/blob/master/examples/webhook/webhook.go) written in Go. Watchman sends either a [Company](https://godoc.org/github.com/moov-io/watchman/client#OfacCompany) or [Customer](https://godoc.org/github.com/moov-io/watchman/client#OfacCustomer) model in JSON to the webhook URL.

Webhook URLs MUST be secure (https://...) and an `Authorization` header is sent with an auth token provided when setting up the webhook. Callers should always verify this auth token matches what was originally provided.

When Watchman sends a [webhook](https://en.wikipedia.org/wiki/Webhook) to your application, the body will contain a JSON representation of the [Company](https://godoc.org/github.com/moov-io/watchman/client#OfacCompany) or [Customer](https://godoc.org/github.com/moov-io/watchman/client#OfacCustomer) model as the body to a POST request. You can see an [example in Go](https://github.com/moov-io/watchman/blob/master/examples/webhook/webhook.go).

An `Authorization` header will also be sent with the `authToken` provided when setting up the watch. Clients should verify this token to ensure authenticated communicated.

Webhook notifications are ran after the OFAC data is successfully refreshed, which is determined by the `DATA_REFRESH_INTERVAL` environmental variable.

## Watching a specific Customer or Company by ID

Moov Watchman supports sending a webhook periodically when a specific [Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/-companyID-/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/-customerID-/watch) is to be watched. This is designed to update another system about an OFAC entry's sanction status.

## Watching a customer or company name

Moov Watchman supports sending a webhook periodically with a free-form name of a [Company](https://moov-io.github.io/watchman/api/#post-/ofac/companies/watch) or [Customer](https://moov-io.github.io/watchman/api/#post-/ofac/customers/watch). This allows external applications to be notified when an entity matching that name is added to the OFAC list. The match percentage will be included in the JSON payload.