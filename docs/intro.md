---
layout: page
title: Introduction
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Introduction to Moov Watchman

## Core Functionality

Watchman is a robust compliance screening tool that provides:

1. **Data Management**:
   - Automatic downloading of sanctions lists (US OFAC, US CSL, UK, EU, etc.)
   - Regular refreshing of data to maintain compliance
   - Custom data file support for specialized screening needs

2. **Search Capabilities**:
   - High-performance indexing and search functionality
   - Advanced fuzzy matching using Jaro-Winkler algorithms
   - Multi-field search with entity type filtering

3. **Integration Options**:
   - HTTP API for web and service integration
   - Native Go library for direct implementation
   - Webhook notifications for automated workflows

## Available Libraries and Components

| Component | Purpose |
|-----------|---------|
| OFAC Library | Parse and process US OFAC sanctions data |
| US CSL  | Process the US Consolidated Screening List |
| UK/EU CSL | Handle UK and European sanctions data formats |

## Search Methodology

### Jaro-Winkler Similarity Algorithm

Watchman uses the [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) algorithm to score the similarity between search queries and list entries. This approach:

- Matches the methodology used by [US Treasury's OFAC Search](https://ofac.treasury.gov/faqs/892)
- Is specifically optimized for person names and other proper nouns
- Produces scores from 0.0 (completely different) to 1.0 (exact match)
- Has been validated by [academic research](https://www.wseas.org/multimedia/journals/computers/2015/a965705-699.pdf) as effective for compliance screening

### Search Customization

Watchman offers environment variables to adjust search behavior:

- `EXACT_MATCH_FAVORITISM`: Controls weight given to exact matches (recommended values: 0.1, 0.25, or 0.5)

## Common Questions

### How are entities prepared for the search index?

Entities undergo a multi-step preparation process before being indexed:

1. **SDN Name Reordering**
   ```
   "MADURO MOROS, Nicolas" → "Nicolas MADURO MOROS"
   ```

2. **Company Name Cleanup**
   ```
   "ACME CORPORATION, INC." → "ACME CORPORATION"
   ```
   *Legal suffixes like "CO.", "INC.", "L.L.C." are removed*

3. **Stopword Removal**
   ```
   "TREES AND EQUIPMENT LTD" → "TREES EQUIPMENT LTD"
   ```
   *Common words like "and", "the", "of" are removed*

4. **UTF-8 Normalization**
   ```
   "Raúl Castro" → "raul castro"
   ```
   *Punctuation is removed, text is lowercased, and diacritical marks are normalized*

The resulting normalized names enable more accurate matching across different formats and variations of the same entity.

### What's the difference between Watchman's search and standard text search?

Standard text search typically relies on exact matches or simple wildcards, which can:
- Miss alternative spellings
- Fail to handle name inversions
- Be overly sensitive to typos
- Require multiple manual searches

Watchman's fuzzy matching approach allows for:
- Identification of similar names despite variations
- Tolerance for typographical errors
- Handling of word order differences
- Normalization of international character sets
- Confidence scoring to prioritize results

This produces more comprehensive screening with fewer false negatives while still providing the tools to manage false positives effectively.

## Next Steps

- See the [Search Documentation](/watchman/search/) for detailed query options
- Explore the [Scoring Methodology](/watchman/methodology/)
- Check the [Configuration Guide](/watchman/config/) for deployment options
