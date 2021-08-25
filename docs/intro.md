---
layout: page
title: Intro
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## What is Watchman?

The Watchman project implements an HTTP server and [Go library](https://pkg.go.dev/github.com/moov-io/watchman) for searching, parsing, and downloading lists. We also have an [example](https://pkg.go.dev/github.com/moov-io/watchman/examples) of the webhook service. Below, you can find a detailed list of features offered by Watchman:

- Download OFAC, BIS Denied Persons List (DPL), and various other data sources on startup
  - Admin endpoint to [manually refresh OFAC and DPL data](https://moov-io.github.io/watchman/runbook/#force-data-refresh)
- Index data for searches
- Async searches and notifications (webhooks)
- Manual overrides to mark a `Company` or `Customer` as `unsafe` (blocked) or `exception` (never blocked).
- Library for OFAC and BIS DPL data to download and parse their custom files

Searching across all sanction lists Watchman uses the [Jaro–Winkler](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) algorithm to score the probability of each search query matching a list entry. This follows after what the [US Treasury OFAC Search](https://home.treasury.gov/policy-issues/financial-sanctions/faqs/topic/1636) uses and what is [recommended in academic literature](https://www.wseas.org/multimedia/journals/computers/2015/a965705-699.pdf).

## FAQ

### How are entities from the list indexed and used in search?

<p>
    Entities from sanction lists and other data files are folded through various pre-computations prior to inclusion in the search index.
    This means the following steps will occur (in order):
    <ul>
        <li>
            <strong>SDN Reordering</strong><br />
            Each individual's SDN name is re-ordered (Example: from "MADURO MOROS, Nicolas" to "Nicolas MADURO MOROS").
        </li>
        <li>
            <strong>Company Name Cleanup</strong><br />
            Suffixes from company names such as: "CO.", "INC.", "L.L.C.", etc are removed.
        </li>
        <li>
            <strong>Stopword Removal</strong><br />
            <a href="https://en.wikipedia.org/wiki/Stop_words">Stopwords</a> are removed. See <a href="https://github.com/bbalet/stopwords">bbalet/stopwords</a> for a full list of supported languages and words subject to removal.
        </li>
        <li>
            <strong>UTF-8 Normalization</strong><br />
            Punctuation is removed along with extra spaces on both ends of the entity name.
            Using <a href="https://godoc.org/golang.org/x/text/unicode/norm#Form">Go's /x/text normalization</a> methods we consolidate entity names and search queries for better searching across multiple languages.
        </li>
    </ul>
</p>

### Why are exact matches of words not ranked higher?

Watchman offers an environmental variable called `EXACT_MATCH_FAVORITISM` that can adjust the weight of exact matches within a query. This value is a percentage (float64) added to exact matches prior to computing the final match percentage. Try using 0.1, 0.25 or 0.5 with your testing.
