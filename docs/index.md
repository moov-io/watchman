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

Watchman empowers your compliance team with powerful tools designed for accuracy, speed, and ease of use:

- **Comprehensive Global Coverage**: Seamlessly integrates multiple international sanctions lists, ensuring thorough compliance screening across jurisdictions.
- **Advanced Fuzzy Matching**: Leverages sophisticated algorithms to detect name variations and reduce false negatives, saving time on manual reviews.
- **High-Performance Search**: Delivers lightning-fast queries with optimized in-memory processing for real-time compliance checks.
- **Flexible Integration Options**: Provides an intuitive HTTP API and native Go library for seamless incorporation into your existing systems.
- **Automated Data Management**: Handles automatic downloads and refreshes of watchlists, keeping your data current without manual intervention.
- **Customizable Configuration**: Fine-tune search parameters, thresholds, and behaviors to align with your organization's specific risk profile. See [Configuration Guide](/watchman/config/) for details.
- **Geocoding Support**: Enhances address matching accuracy with integrated providers like Google Maps and OpenCage.
- **Senzing Compatibility**: Supports Senzing format for advanced entity resolution and data import.

## Included Lists

Watchman integrates the following lists to help you maintain global compliance. Use the env variable `INCLUDED_LISTS` or [config file](https://moov-io.github.io/watchman/config/#download) to customize which lists are loaded.

| Source            | List                                                                                                                                                                                    |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **OpenSanctions** | [Any Senzing formatted list from OpenSanctions](https://www.opensanctions.org/datasets/)                                                                                                |
| European Union    | [Consolidated Sanctions List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)                        |
| US Government     | [Consolidated Screening List (CSL)](https://www.trade.gov/consolidated-screening-list), [FinCEN 311](https://home.treasury.gov/policy-issues/terrorism-and-illicit-finance/311-actions) |
| US Treasury       | [Office of Foreign Assets Control (OFAC)](https://ofac.treasury.gov/sanctions-list-service) and Non-SDN list                                                                            |
| United Kingdom    | [OFSI Sanctions List](https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents)                                |
| United Nations    | [Consolidated Sanctions List](https://www.un.org/sc/resources/sc-sanctions)                                                                                                             |

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
