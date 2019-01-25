moov-io/ofac
===

[![GoDoc](https://godoc.org/github.com/moov-io/ofac?status.svg)](https://godoc.org/github.com/moov-io/ofac)
[![Build Status](https://travis-ci.com/moov-io/ofac.svg?branch=master)](https://travis-ci.com/moov-io/ofac)
[![Coverage Status](https://codecov.io/gh/moov-io/ofac/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/ofac)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ofac)](https://goreportcard.com/report/github.com/moov-io/ofac)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ofac/master/LICENSE)

*project is under active development and is not production ready*


### Configuration

**Search Similarity Metrics**

OFAC computes string similarity using the Levenshtein algorithm and can match sensitivity can be configured with environment variables.

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `ADDRESS_SIMILARITY` | Ratio of Levenshtein distance for two strings to be considered equal. | 0.85 |
| `ALT_SIMILARITY` | Ratio of Levenshtein distance for two strings to be considered equal. | 0.85 |
| `NAME_SIMILARITY` | Ratio of Levenshtein distance for two strings to be considered equal. | 0.85 |
| `OFAC_DATA_REFRESH` | Interval for OFAC data redownload and reparse. | 12h |


### OFAC Data

#### Refresh

OFAC supports a one-off re-download and refresh of OFAC data from the federal website, simply issue a request like: `curl -v http://a.b.c.d:9094/ofac/refresh`. The response will be `200 OK` only if successful and an errors will be logged to stdout.
