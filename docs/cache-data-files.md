---
layout: page
title: Caching Data Files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Caching Data Files

Watchman supports sourcing data files from the local filesystem and falling back to the remote servers when local files do not exist. This is helpful for flakey servers or when operating a large number of Watchman instances.

By setting `INITIAL_DATA_DIRECTORY` to a local directory Watchman will look for the following files.

**OFAC**

- `add.csv` - Address
- `alt.csv` - Alternate ID
- `sdn.csv` - Specially Designated National
- `sdn_comments.csv` - Specially Designated National Comments

Download the files from [Data Center - SDN List](https://sanctionslist.ofac.treas.gov/Home/SdnList)

**US Consolidated Screening List**

- `csl.csv` - US Consolidated Screening List
- `dpl.txt` - US Denied Persons List

Download the [US Consolidated Screening List](https://www.trade.gov/consolidated-screening-list)

**EU Consolidated Screening List**

- `eu_csl.csv` - EU Consolidated Screening List

Download the [EU Consolidated Screening List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)

**UK Consolidated Screening List**

- `ConList.csv` - UK Consolidated Screening List
- `UK_Sanctions_List.ods` - UK Sanctions List

Download the [UK Consolidated Screening List](https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents) and the [UK Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list)
