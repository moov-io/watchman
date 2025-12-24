---
layout: page
title: Data preparation Pipeline
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Search Indexing

Watchman holds the searchable entities in memory (or in the database for ingested files) to speed up search requests. The index is re-computed every time lists are refreshed.

## TF-IDF

Watchman builds an in-memory [TF-IDF](https://en.wikipedia.org/wiki/Tf%E2%80%93idf) (Term Frequency–Inverse Document Frequency) index over the corpus of searchable entities, treating each entity's normalized name (including primary names and alternate known-as entries from lists like OFAC SDN) as a "document" composed of tokenized terms. This precomputes IDF values to weight terms inversely by their commonality across the entire watchlist—rare terms (e.g., unique surnames or identifiers) receive higher weights, while common ones (e.g., "bank" or frequent first names) are downweighted. During name searches, TF-IDF weighting is optionally applied to boost the contribution of discriminative rare terms in similarity scoring, improving match precision for fuzzy name matching against global sanctions and watchlists without overemphasizing ubiquitous words.

The feature is opt-in via [configuration, with tunable parameters](/watchman/config/) like smoothing and IDF bounds to adapt to corpus size and growth.
