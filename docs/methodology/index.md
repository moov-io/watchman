---
layout: page
title: Matching Methodology
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Watchman Matching Methodology

## Overview

Moov Watchman implements a sophisticated multi-dimensional matching system designed to balance accuracy, performance, and usability for compliance professionals. This document explains the technical foundations of Watchman's matching algorithms.

## Core Matching Architecture

![Matching Architecture](https://via.placeholder.com/800x400?text=Watchman+Matching+Architecture)

Watchman uses a hierarchical matching approach that evaluates entity similarity across multiple dimensions:

| Priority | Component | Description |
|----------|-----------|-------------|
| Highest | Critical Identifiers | Government IDs, passport numbers, registration codes |
| High | Name Matching | Sophisticated fuzzy matching of entity names |
| Medium | Supporting Information | Addresses, dates of birth, and contextual metadata |
| Low | Relationship Data | Connections to other entities, when available |

## Advanced Name Matching

### Jaro-Winkler Algorithm

Watchman uses an enhanced version of the Jaro-Winkler string similarity algorithm:

```
sim_jw(s1, s2) = sim_j(s1, s2) + p * l * (1 - sim_j(s1, s2))
```

Where:
- `sim_j` is the base Jaro similarity
- `p` is the prefix scaling factor (default 0.1)
- `l` is the length of the common prefix (up to 4 characters)

Watchman's implementation includes:

1. **Token-based Comparison**
   - Names are tokenized and compared word-by-word
   - Example: "John Michael Smith" → ["john", "michael", "smith"]

2. **Positional Weighting**
   - Tokens in similar positions receive higher match scores
   - Handles name order variations more effectively

3. **Length Normalization**
   - Shorter token comparisons are weighted differently than longer ones
   - Prevents bias towards long or short names

4. **First-Letter Penalty**
   - Different first letters receive an additional penalty
   - Based on research showing first letters are rarely mistranscribed

### Phonetic Matching

For handling spelling variations, especially in transliterated names, Watchman includes:

1. **Modified Soundex**
   - Groups phonetically similar characters
   - Example: "Smith" and "Smyth" have identical phonetic codes

2. **First Character Analysis**
   - Names with different first-character phonetic classes are less likely to match
   - Improves performance by eliminating obvious non-matches early

## Entity-Type Specific Matching

Watchman applies specialized matching logic based on entity type:

### Person Matching

- **ID Verification**: Exact matches on government IDs
- **Name Components**: First, middle, last name specifics
- **Date Verification**: Birth date comparison when available
- **Title Comparison**: Professional roles and titles

### Business Matching

- **Registration Numbers**: Tax and business identifiers
- **Name Normalization**: Special handling of business entity types
- **Abbreviation Handling**: Common business abbreviations (Inc → Incorporated)

### Vessel/Aircraft Matching

- **Specialized Identifiers**: IMO numbers, call signs, registration codes
- **Flag/Registry** Confirmation: Jurisdictional information
- **Technical Details**: Tonnage, model, etc.

## Scoring System

The final match score is calculated through:

1. **Weighted Component Aggregation**
   - Each component's score is multiplied by its importance weight
   - Formula: `final_score = Σ(component_score * component_weight) / Σ(component_weight)`

2. **Critical Field Multipliers**
   - Required fields receive extra weight
   - Exact matches on certain fields can override fuzzy matching

3. **Coverage Analysis**
   - Penalties applied when query doesn't cover enough entity fields
   - Prevents high scores from partial data

4. **Perfect Match Boosting**
   - High-quality matches that meet specific thresholds receive a boost
   - Configurable via `EXACT_MATCH_FAVORITISM` environment variable

## Threshold Configuration

Watchman allows customizing match thresholds for different risk tolerances:

| Threshold | Default Value | Use Case |
|-----------|--------------|----------|
| High Confidence | 0.95+ | Automatic blocking/alerts |
| Medium Confidence | 0.85-0.94 | Manual review queue |
| Low Confidence | 0.70-0.84 | Enhanced due diligence |

## Performance Optimizations

Watchman includes several performance enhancements:

1. **Early Termination**
   - Once a high-confidence match is found, detailed scoring may be skipped
   - Reduces processing time for obvious matches

2. **Phonetic Pre-filtering**
   - Initial filtering based on phonetic codes
   - Narrows candidate pool before expensive comparison

3. **Caching**
   - Frequently searched entities are cached
   - Improves performance for common searches

4. **Parallel Processing**
   - Multi-threaded search for large entity sets
   - Scales with available CPU cores

## Benefits for Compliance Teams

Watchman's sophisticated matching provides several key advantages:

1. **Reduced False Positives**
   - Multi-dimensional scoring reduces irrelevant matches
   - Context-aware matching prioritizes meaningful similarities

2. **Improved Match Confidence**
   - Detailed scoring provides better justification for match decisions
   - More information for analysts making review decisions

3. **Comprehensive Audit Trail**
   - Score components show exactly why matches occurred
   - Helps demonstrate compliance program effectiveness

4. **Risk-Based Approach**
   - Configurable thresholds align with organizational risk tolerance
   - Different rules can be applied to different entity types or programs

## Validation Methodology

Watchman's matching algorithms are validated through:

1. **Test Suite Verification**
   - Comprehensive test cases covering edge cases
   - Regression testing on algorithm changes

2. **Known Entity Testing**
   - Verification against known sanctions entities and aliases
   - Spelling variation handling

3. **False Positive Analysis**
   - Regular review of common false positives
   - Algorithm tuning to reduce unnecessary matches

For more information on validation, see the [Model Validation](/model-validation) page.
