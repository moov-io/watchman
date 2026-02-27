---
layout: page
title: Caching Data Files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Caching Data Files

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

Watchman supports sourcing data files from the local filesystem and falling back to the remote servers when local files do not exist. This is helpful for flakey servers or when operating a large number of Watchman instances.

By setting `INITIAL_DATA_DIRECTORY` to a local directory Watchman will look for the following files.

**OFAC**

- `add.csv` - Address
- `alt.csv` - Alternate ID
- `sdn.csv` - Specially Designated National
- `sdn_comments.csv` - Specially Designated National Comments

Download the files from [Data Center - SDN List](https://sanctionslist.ofac.treas.gov/Home/SdnList)

**US Consolidated Screening List**

- `consolidated.csv` - US Consolidated Screening List

Download the [US Consolidated Screening List](https://www.trade.gov/consolidated-screening-list)

**EU Consolidated Screening List**

- `eu_csl.csv` - EU Consolidated Screening List

Download the [EU Consolidated Screening List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)

**UK Consolidated Screening List**

- `UK_Sanctions_List.csv` - UK Sanctions List

Download the [UK Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list)

**UN Consolidated Sanctions List**

- `un_consolidated.xml` - UN Consolidated Sanctions List

Download from [UN Sanctions](https://www.un.org/sc/resources/sc-sanctions)
