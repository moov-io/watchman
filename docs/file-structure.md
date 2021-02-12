---
layout: page
title: File Structure
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# CSV Data Specifications - Specially Designated National (SDN)

Format *.csv consists of records separated by carriage returns (ASCII character 
13), with fields (values) within records delimited by the "," (comma) symbol 
(ASCII character 44).

Null values consist of "-0-" (ASCII characters 45, 48, 45).

The Comma Separated Values (.csv), release consists of three  ASCII text files -- a main
file listing the name of the SDN and other information unique to that entity
(sdn.csv), a file of addresses (add.csv),  and a file of alternate names (alt.csv).
Addresses and alternate names are linked to particular SDNs using unique integer
values in a linking or primary key column. The integers used are assigned for
linking purposes only and do not represent an official reference to that entity.

Releases of the database-format files are intended as a service to the user
community.  OFAC's SDN list is published in the Federal Register.  All of OFAC's
lists are drawn from the same underlying data and every effort has been made to
ensure consistency.  The Federal Register will govern should differences arise. 
Due to the nature, urgency, and sensitivity of the programs  which OFAC
administers and enforces, it may not always be possible to  provide advanced
notice to users of format changes to the database  structure.

## CSV Files

Comma delimited: sdn.csv, add.csv, alt.csv, sdn_comments.csv
 
### sdn.csv

| Column Sequence | Column Name | Type | Size| Description | 
| :---: | :---: | :---: | :--- | :--- |
| 1 | ent_num | number | | unique record, identifier/unique, listing identifier |
| 2 | SDN_Name | text | 350 | name of SDN |
| 3 | SDN_Type | text | 12 | type of SDN |
| 4 | Program | text | 50 | sanctions program name |
| 5 | Title | text | 200 | title of an individual |
| 6 | Call_Sign | text | 8 | vessel call sign |
| 7 | Vess_type | text | 25 | vessel type |
| 8 | Tonnage | text | 14 | vessel tonnage |
| 9 | GRT | text | 8 | gross registered tonnage |
| 10 | Vess_flag | text | 40 | vessel flag |
| 11 | Vess_owner | text | 150 | vessel owner |
| 12 | Remarks | text | 1000 | remarks on SDN* |

### add.csv

| Column Sequence | Column Name | Type | Size| Description | 
| :---: | :---: | :---: | :--- | :--- |
| 1 | Ent_num | number | | link to unique listing |
| 2 | Add_num | number | | unique record identifier |
| 3 | Address | text | 750 | street address of SDN |
| 4 | City/State/Province/Postal Code | text | 116 | city, state/province, zip/postal code |
| 5 | Country | text | 250 | country of address |
| 6 | Add_remarks | text | 200 | remarks on address |

### alt.csv

| Column Sequence | Column Name | Type | Size| Description | 
| :---: | :---: | :---: | :--- | :--- |
| 1 | Ent_num | number | | link to unique listing |
| 2 | alt_num | number | | unique record identifier |
| 3 | alt_type | text | 8 | type of alternate identity (aka, fka, nka) |             
| 4 | alt_name | text | 350 | alternate identity name |
| 5 | alt_remarks | text | 200 | remarks on alternate identity |

### sdn_comments.csv

| Column Sequence | Column Name | Type | Size| Description | 
| :---: | :---: | :---: | :--- | :--- |
| 1 | Ent_num | number | | link to unique listing |
| 2 | RemarksExtended | text | | remarks extended on a SDN |

## Definitions

| Item | CHAR | ASCII DEC |
| :---: | :---: | :---: |
| Record separator | CR (carriage return) | 13 |
| Field (value) delimiter | , | 44 | |
| Text value quotes | " | 34 | |
| Null | -0- | 45, 48, 45 | |

## ASCII Table and Description

[ASCII Table and Description](http://www.asciitable.com/)