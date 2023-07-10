---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

![Moov Watchman Logo](https://repository-images.githubusercontent.com/163885848/41101f80-c6d9-11ea-9ab5-dc9f51b849df)

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

Moov Watchman is an HTTP API and Go library that offers download, parse, and search functions over numerous trade sanction lists from the United States, agencies, and nonprofits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.

Lists included in search are:

- US Treasury - Office of Foreign Assets Control
  - [Specially Designated Nationals](https://home.treasury.gov/policy-issues/financial-sanctions/specially-designated-nationals-and-blocked-persons-list-sdn-human-readable-lists)
    - Includes SDN, SDN Alternative Names, SDN Addresses
- [United States Consolidated Screening List](https://www.export.gov/article2?id=Consolidated-Screening-List)
   - Department of Commerce – Bureau of Industry and Security
      - [Denied Persons List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/denied-persons-list)
      - [Unverified List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/unverified-list)
      - [Entity List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/entity-list)
   - Department of State – Bureau of International Security and Non-proliferation
      - [Nonproliferation Sanctions](http://www.state.gov/t/isn/c15231.htm)
   - Department of State – Directorate of Defense Trade Controls
      - ITAR Debarred (DTC)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Specially Designated Nationals List](https://ofac.treasury.gov/specially-designated-nationals-list-data-formats-data-schemas)
      - [Foreign Sanctions Evaders List](https://ofac.treasury.gov/consolidated-sanctions-list-non-sdn-lists/foreign-sanctions-evaders-fse-list)
      - [Sectoral Sanctions Identifications List](https://ofac.treasury.gov/consolidated-sanctions-list-non-sdn-lists/sectoral-sanctions-identifications-ssi-list)
      - [Palestinian Legislative Council List](https://ofac.treasury.gov/consolidated-sanctions-list/non-sdn-palestinian-legislative-council-ns-plc-list)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Sectoral Sanctions Identifications List](https://ofac.treasury.gov/consolidated-sanctions-list-non-sdn-lists/sectoral-sanctions-identifications-ssi-list)
- [EU - Consolidated Sanctions List](https://data.europa.eu/data/datasets/consolidated-list-of-persons-groups-and-entities-subject-to-eu-financial-sanctions?locale=en)
   - NOTE: it is recommended to [create your own europa.eu account](https://webgate.ec.europa.eu/cas/login) and then access the [EU Financial Sanctions Files](https://webgate.ec.europa.eu/fsd/fsf)
      - Use the token described under the "Show settings for crawler/robot" section
- [UK - OFSI Sactions List](https://www.gov.uk/government/publications/financial-sanctions-consolidated-list-of-targets/consolidated-list-of-targets#contents)
- [UK - Sanctions List](https://www.gov.uk/government/publications/the-uk-sanctions-list) (Disabled by default)

All United States and European Union companies are required to comply with various regulations and sanction lists (such as the US Patriot Act requiring compliance with the BIS Denied Persons List).
