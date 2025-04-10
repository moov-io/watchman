---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Moov Watchman: Compliance Screening Made Simple

![Moov Watchman Logo](https://repository-images.githubusercontent.com/163885848/41101f80-c6d9-11ea-9ab5-dc9f51b849df)

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

The v0.5x series of Watchman has revamped its search engine. The following lists are being re-added into Watchman.

| Source | List |
|--------|------|
| European Union | [Consolidated Sanctions List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en) |
| United Kingdom | [OFSI Sanctions List](https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents) |
| United Kingdom | [Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list) (Disabled by default) |

## Regulatory Context

Organizations in the United States and European Union are required to comply with various sanctions programs and screening requirements, including:

- US Patriot Act compliance with the BIS Denied Persons List
- OFAC sanctions compliance requirements
- EU and UK financial sanctions regimes

Watchman helps you manage these compliance obligations efficiently through a single, powerful platform.

## Getting Started

To begin using Watchman, see our:

- [Docker Usage](/watchman/usage-docker/)
- [API Documentation](/watchman/api/)
- [Configuration Options](/watchman/config/)

## About Moov

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.
