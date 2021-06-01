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

- US Treasury - Office of Foreign Assets Control (OFAC)
  - [Specially Designated Nationals](https://home.treasury.gov/policy-issues/financial-sanctions/specially-designated-nationals-and-blocked-persons-list-sdn-human-readable-lists) (SDN)
    - Includes SDN, SDN Alternative Names, SDN Addresses
  - [Sectoral Sanctions Identifications](https://home.treasury.gov/policy-issues/financial-sanctions/consolidated-sanctions-list/sectoral-sanctions-identifications-ssi-list) (SSI)
- US Department of Commerce - Bureau of Industry and Security (BIS)
  - [Denied Persons List](https://bis.data.commerce.gov/dataset/Denied-Persons-List-with-Denied-US-Export-Privileg/xwtd-wd7a/data) (DPL)
  - [Entity List](https://www.bis.doc.gov/index.php/policy-guidance/lists-of-parties-of-concern/entity-list) (EL)

All United States and European Union companies are required to comply with various regulations and sanction lists (such as the US Patriot Act requiring compliance with the BIS Denied Persons List).
