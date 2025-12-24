---
layout: page
title: Geocoding
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Geocoding

Watchman added support in `v0.57.0` for various geocoding providers that populate the `Latitude` and `Longitude` fields of every entity's address.

Refer to the [file configuration](/watchman/config/#file) section for more information.

### Providers

The supported providers are listed below. If you'd like to contribute or recommend a new provider [open an issue](https://github.com/moov-io/watchman/issues/new).

#### Google

The `google` geocoding provider uses Google Maps APIs for geocoding requests.

#### Nominatim

The `nominatim` geocoding provider uses [Nominatim](https://nominatim.org/) from OpenStreetMap to geocode requests.

For information on how to self-host refer to [Nominatim's documentation](https://nominatim.org/release-docs/latest/admin/Installation/).

#### OpenCage

The `opencagedata` geocoding provider uses the [OpenCage](https://opencagedata.com/) Geocoding API.
