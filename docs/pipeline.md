---
layout: page
title: Precomputation pipeline
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Pipeline

Watchman performs various operations on records prior to their inclusion in the search index and offers some inspection capabilities into the search index.

Setting `DEBUG_NAME_PIPELINE=true` will enable verbose logging of every text processing step performed on records prior to inclusion in the search index.

```
ts=2019-12-19T17:11:25.325105Z caller=pipeline.go:84 pipeline=*main.reorderSDNStep result="ANGLO-CARIBBEAN CO., LTD." original="ANGLO-CARIBBEAN CO., LTD."
ts=2019-12-19T17:11:25.32513Z caller=pipeline.go:84 pipeline=*main.companyNameCleanupStep result=ANGLO-CARIBBEAN. original="ANGLO-CARIBBEAN CO., LTD."
ts=2019-12-19T17:11:25.325583Z caller=pipeline.go:84 pipeline=*main.stopwordsStep result=anglo-caribbean original="ANGLO-CARIBBEAN CO., LTD."
ts=2019-12-19T17:11:25.325613Z caller=pipeline.go:84 pipeline=*main.normalizeStep result="anglo caribbean" original="ANGLO-CARIBBEAN CO., LTD."
```

Note: Some record types are skipped in pipeline steps.

## Debugging SDNs

For a more precise inspection of a specific SDN record, call the following endpoint. The `debug` object is included along with the SDN in question.

```
$ curl -s localhost:9094/debug/sdn/16016 | jq .
{
  "SDN": {
    "entityID": "16016",
    "sdnName": "CYLINDER SYSTEM L.T.D.",
    // ...
    "remarks": "Tax ID No. 27694384517 (Croatia)."
  },
  "debug": {
    "indexedName": "cylinder system",
    "parsedRemarksId": "27694384517"
  }
}
```

## Pipeline steps

**Reordering of individual names**

This step processes SDN and SSI entries to rearrange their name into a "first middle last" ordering.

Example: `MADURO MOROS, Nicolas` into `Nicolas MADURO MOROS`

**Company name cleanup**

This step strips SDN and SSI company suffixes/titles from their indexed name. The original name from their source file is never changed.

Example: `AMD CO. LTD AGENCY` into `AMD AGENCY`

**Stopwords removal**

This step removes stopwords from SDN and SSI entities. [Stopwords](https://en.wikipedia.org/wiki/Stop_words) are typically the most common words in languages and don't convey necessary information in a sentence. They are more typically used for grammatical correctness and thus can be ignored in search rankings.

Example: `COLOMBIANA DE CERDOS LTDA.` into `colombiana cerdos ltda`
Example: `Trees and Trucks` into `trees trucks`

**Normalization**

This step "normalizes" all text passed to it by converting it to lowercase, removing punctuation, and applying [UTF-8 Normalization](https://en.wikipedia.org/wiki/Unicode_equivalence#Normalization) to support searching non-English names with English letters. Watchman has a primary focus on American business which often performs this same conversion as a result of human or computer systems.

Example: `Raúl Castro` into `raul castro`

More information: [Why You Need to Normalize Unicode Strings](https://withblue.ink/2019/03/11/why-you-need-to-normalize-unicode-strings.html)
