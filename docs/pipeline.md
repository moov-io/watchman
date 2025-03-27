---
layout: page
title: Data preparation Pipeline
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Data Preparation Pipeline

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

## Overview

Before entities are added to the search index, Watchman processes their names and associated data to improve search accuracy. This processing pipeline standardizes different text variations to ensure that similar entities will match during searches, even when their original formats differ.

## Pipeline Stages

The pipeline consists of four main stages that transform entity data:

### 1. Name Reordering

Transforms inverted names (typically found in sanctions lists) into standard order.

**Example**:
- `MADURO MOROS, Nicolas` → `Nicolas MADURO MOROS`

This makes it easier to match names entered in natural order during searches.

### 2. Company Name Cleanup

Removes legal suffixes and business entity types from company names.

**Example**:
- `AMD CO. LTD AGENCY` → `AMD AGENCY`

This focuses matching on the core business name rather than common legal designations.

### 3. Stopword Removal

Removes common words (like "the", "and", "of") that don't help distinguish entities.

**Example**:
- `BANK OF AMERICA` → `BANK AMERICA`

This prevents common words from reducing match quality.

### 4. Text Normalization

Standardizes text by removing accents, converting to lowercase, and handling special characters.

**Example**:
- `Raúl Castro` → `raul castro`

This ensures that searches work regardless of capitalization or special characters.

## Entity-Specific Processing

Different entity types (person, business, vessel, etc.) receive specialized processing:

- **Person entities**: Name reordering, gender standardization, proper handling of titles
- **Business entities**: Company suffix removal, registration number standardization
- **Vessels/Aircraft**: Specialized identifier formatting (IMO numbers, call signs)

## Impact on Searching

Understanding this pipeline helps you build more effective search queries:

1. **Don't worry about name format** - Both `John Smith` and `SMITH, John` will match
2. **Company suffixes don't matter** - `Acme Inc` and `ACME Corporation` will match
3. **Accents and case are ignored** - `José` and `jose` are treated the same
4. **Common words can be omitted** - `Bank of America` and `Bank America` are equivalent

## Debugging

If you need to see how your search terms are being processed:

1. Enable debug mode in your searches with `debug=true`

This will help you understand why particular searches are or aren't matching expected results.
