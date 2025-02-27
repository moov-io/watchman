---
layout: page
title: Model Methodology
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Moov Watchman's Advanced Similarity Analysis

### Introduction

Moov Watchman implements a sophisticated multi-layered similarity matching system that significantly outperforms basic search tools like OFAC's SDN search portal. The system features a well-architected matching framework designed for high accuracy in sanctions screening and compliance workflows. This document provides a detailed overview of how Watchman's scoring algorithms work.

### Core Matching Architecture

Watchman's matching engine uses a hierarchical approach to entity matching that considers multiple dimensions of similarity with appropriate weighting for each component:

- **Critical Identifiers** (highest priority): Exact matches of government IDs, crypto addresses, and contact information
- **Name Comparison** (high priority): Sophisticated fuzzy matching of entity names and titles
- **Supporting Information** (medium priority): Contextual details like addresses, dates, and other entity metadata

This layered design allows Watchman to make nuanced decisions about match quality, prioritizing the most definitive identifiers while still considering supporting evidence.

### Advanced Name Matching

The core of Watchman's matching capability lies in its sophisticated name matching algorithms:

#### Jaro-Winkler with Customizations

While OFAC or other typical search uses basic keyword matching, Watchman employs enhanced Jaro-Winkler string similarity with custom adjustments:

**Token-based Comparisons**: Names are broken into tokens/words, and each token from the search query is matched against the indexed terms using a pairwise approach

**Length Difference Penalties**: The algorithm penalizes matches between terms with significant length differences
**Different First Letter Penalties**: Additional penalties are applied when words begin with different letters
**Adjacent Term Similarity**: The algorithm checks for matches in adjacent positions, accommodating name inversions and word order differences

#### Phonetic Matching

Watchman implements phonetic matching to catch similar-sounding names regardless of spelling variations:

**Custom Soundex Implementation**: The system uses a modified Soundex algorithm that groups phonetically similar letters
**First Character Analysis**: Names are compared based on their first character's phonetic class, reducing unnecessary comparisons for clearly dissimilar names
**Efficiency Optimization**: Phonetic filtering is applied early in the comparison process to eliminate obviously non-matching pairs

### Entity-Type Specific Matching

Watchman differentiates between various entity types (Person, Business, Organization, Vessel, Aircraft) and applies specialized matching rules for each:

#### Person Matching

**Government ID Comparison**: High priority for exact matches of ID numbers
**Title Comparison**: Fuzzy matching of professional titles with abbreviation expansion
**Birth/Death Date Analysis**: Date comparison with consistency checking between birth and death dates
**Logical Validation**: The system checks that dates make temporal sense (birth precedes death)

#### Business/Organization Matching

**Registration ID Comparison**: Primary focus on government/tax identification numbers
**Name Variations**: Consideration of alternate names and historical names
**Creation/Dissolution Dates**: Temporal validation for business lifecycle

#### Vessel/Aircraft Matching

**Industry-specific Identifiers**: Primary matching using IMO numbers, call signs, MMSI (for vessels) or serial numbers and ICAO codes (for aircraft)
**Registration Details**: Flag country, model, and ownership comparison

#### Address and Contact Information Matching

Watchman applies weighted matching to different address components:

**Primary Address Line**: Highest priority (most important component)
**City and Country**: High priority (critical for location)
**Postal Code**: Strong verification priority
**Secondary Address Info and State**: Supporting information priority

This weighted approach prioritizes the most important address components while still considering all available information.

### Comprehensive Scoring System

The final match score is calculated through a sophisticated process:

**Weighted Component Aggregation**: Each component score is multiplied by its importance weight
**Critical Field Multiplier**: Required fields are given additional weight
**Coverage Analysis: Penalties** are applied if the query doesn't cover enough of the index entity's available fields
**Perfect Match Boosting**: High-quality matches that meet specific criteria receive a boost
**Type Mismatch Handling**: Entity type mismatches result in a zero score

### Score Adjustment Based on Match Quality

Watchman applies various score adjustments based on match quality:

**Minimum Field Requirements**: Penalties for queries with insufficient field coverage
**Name-Only Match Penalties**: Reduced confidence in matches based solely on name
**Historical Name Handling**: Names identified as "Former Name" receive a penalty
**Exact Match Override**: Certain exact matches of critical identifiers automatically receive a perfect score

### Implementation Optimizations

Watchman includes several optimizations that improve both accuracy and performance:

**Early Termination**: The algorithm stops comparing once a high-confidence match is found
**Phonetic Filtering**: Preliminary filtering reduces unnecessary detailed comparisons
**Selective Deep Analysis**: Detailed analysis is applied only to promising candidates
**Configuration Flexibility**: The system allows tuning of algorithm parameters

### Benefits for Compliance Teams

For risk and compliance professionals, Watchman's sophisticated matching offers several benefits:

**Reduced False Positives**: The multi-dimensional scoring and contextual analysis help differentiate between true and false matches
**Improved Match Confidence**: Detailed scoring provides better justification for match decisions
**Audit Trail**: The system's detailed scoring components provide clear rationale for matches
**Configurable Thresholds**: Teams can adjust sensitivity based on risk tolerance
**Enhanced Detection**: The sophisticated algorithms can identify potential sanctions matches that would be missed by conventional systems

### Conclusion

Moov Watchman's advanced match scoring system represents a significant improvement over basic search tools like OFAC's search portal. By combining probabilistic string matching, phonetic analysis, entity-specific logic, and weighted multi-dimensional scoring, Watchman provides compliance teams with more accurate and explainable match results.

The system's attention to detail—from handling name variations to applying logical consistency checks—demonstrates a deep understanding of the challenges in entity resolution for sanctions screening. For auditors and compliance professionals, this means greater confidence in screening results, better justification for match decisions, and a more robust compliance program overall.
