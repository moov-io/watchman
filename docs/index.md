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
      - [Denied Persons List](http://www.bis.doc.gov/dpl/default.shtm)
      - [Unverified List](http://www.bis.doc.gov/enforcement/unverifiedlist/unverified_parties.html)
      - [Entity List](http://www.bis.doc.gov/entities/default.htm)
   - Department of State – Bureau of International Security and Non-proliferation
      - [Nonproliferation Sanctions](http://www.state.gov/t/isn/c15231.htm)
   - Department of State – Directorate of Defense Trade Controls
      - ITAR Debarred (DTC)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Specially Designated Nationals List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/default.aspx)
      - [Foreign Sanctions Evaders List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/fse_list.aspx)
      - [Sectoral Sanctions Identifications List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/ssi_list.aspx)
      - [Palestinian Legislative Council List](https://www.treasury.gov/resource-center/sanctions/Terrorism-Proliferation-Narcotics/Pages/pa.aspx)
   - Department of the Treasury – Office of Foreign Assets Control
      - [Sectoral Sanctions Identifications List](http://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/ssi_list.aspx)

All United States and European Union companies are required to comply with various regulations and sanction lists (such as the US Patriot Act requiring compliance with the BIS Denied Persons List).
