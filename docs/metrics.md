---
layout: page
title: Prometheus metrics
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Metrics

This list represents the series of Prometheus metrics that Watchman reports. All metrics are served on the admin server (`:9094/metrics` by default). Outside of this list, the standard Go metrics are reported.

## Last data refresh count

`last_data_refresh_count` holds the count of records parsed from the latest download and indexing of the specified data source.

```
# HELP last_data_refresh_count Count of records for a given sanction or entity list
# TYPE last_data_refresh_count gauge
last_data_refresh_count{source="BISEntities"} 1503
last_data_refresh_count{source="DPs"} 584
last_data_refresh_count{source="SDNs"} 8497
last_data_refresh_count{source="SSIs"} 290
```

## Last data refresh success

`last_data_refresh_success` holds a unix timestamp of the last successful data refresh.

```
# HELP last_data_refresh_success Unix timestamp of when data was last refreshed successfully
# TYPE last_data_refresh_success gauge
last_data_refresh_success 1.604635215e+09
```

## Last data refresh failure

`last_data_refresh_failure` holds a unix timestamp of the last data refresh that failed.

```
# HELP last_data_refresh_failure Unix timestamp of the most recent failure to refresh data
# TYPE last_data_refresh_failure gauge
last_data_refresh_failure{source="SDNs"} 1.60520363e+09
```

## Match percentages

This histogram holds the results of searches performed. The exposed types represent the differnet endpoints and query parameters used such as: `address`, `addressname`, `altName`, `name`, `q`, `remarksID`.

```
# HELP match_percentages Histogram representing the match percent of search routes
# TYPE match_percentages histogram
match_percentages_bucket{type="address",le="0"} 0
match_percentages_bucket{type="address",le="0.5"} 0
match_percentages_bucket{type="address",le="0.8"} 2924
match_percentages_bucket{type="address",le="0.9"} 2924
match_percentages_bucket{type="address",le="0.99"} 2924
match_percentages_bucket{type="address",le="+Inf"} 2927
match_percentages_sum{type="address"} 2277.222222222333
match_percentages_count{type="address"} 2927
```

## MySQL connections

`mysql_connections` represents the current count of idle, active, and open connections to the configured MySQL database.

```
# HELP mysql_connections How many MySQL connections and what status they're in.
# TYPE mysql_connections gauge
mysql_connections{state="idle"} 1
mysql_connections{state="inuse"} 0
mysql_connections{state="open"} 1
```
