---
layout: page
title: Addresses
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Address Parsing and Normalization

Watchman optionally uses **[libpostal](https://github.com/openvenues/libpostal)** to deliver advanced address parsing, classification, and normalization. By integrating the latest data from **[Senzing’s classifier, data, and parser](https://github.com/Senzing/libpostal-data)**, it accurately handles global addresses during indexing and screening requests.

To enable libpostal support, activate the **[PostalPool](/watchman/config/#postalpool)** option in Watchman’s configuration.
