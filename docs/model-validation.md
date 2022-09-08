---
layout: page
title: Model Validation
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Model Validation

A model is defined as "a quantitative method, system, or approach that applies statistical, economic, financial, or mathematical theories, techniques, and assumptions to process input data into quantitative estimates." Watchman offers a statistical scoring of names, addresses, government identifiers, and other data against watchlists from numerous public sources. Watchman is a modern software application used by many companies in production deployments at cloud scale.

## Data Sources

Watchman's default data sources are several government agency and public data sources. These are typically lists of entity data (names, addresses, government IDs, etc) published regularly. Watchman will periodically download these data files and re-index the data. By default this refresh occurs on a 12-hour interval and can be configured or initiated manually. This allows for a high degree of uptime and continual improvement.

### Sources List

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

After the data files are refreshed users can [configure webhook notifications](https://moov-io.github.io/watchman/webhook-notifications/) to be notified and initiate custom processes. Custom data files can be used with Watchman.

Watchman will index the data sources in a normalized form for improved search rankings. These [steps are documented](https://moov-io.github.io/watchman/pipeline/#pipeline-steps) for data cleanup and typical search patterns.

## Scoring

Watchman uses the [Jaro–Winkler](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) string comparison scoring for each query. Each word part is ordered and compared to the indexed data. The model, scoring, and search rankings are verified on every source code commit and release of Watchman. Monthly checks are performed to verify no unexpected changes have occurred. Changes to the scoring are thoroughly analyzed prior to inclusion as the results returned can have large impacts. Users of Watchman should experiment with different tolerances with positive / negative matches.

> Jaro-Winkler distance is a public algorithm for comparing two strings of text to determine their similarity. Results range from 0.0 (completely unequal) to 1.0 (completely equal). Jaro-Winkler has been optimized for human and street names and is a modification of the Jaro algorithm with an additional boost on exact matches.
> There are two parameters with their defaults specified as: `boostThreshold=0.7` and `prefixSize=4`. See "Other Links" below for references.

Periodic searches of names, addresses, IDs, etc can be performed by two different methods. Watchman supports "watches" which are performed after source data is refreshed and delivers results via webhooks. Otherwise the HTTP endpoints can be called to get the current scoring. Watchman is highly performant to support large amounts of queries.

Search queries return better results when multiple criteria are included with the query. Simple name queries will return false positive matches, so including addresses, alternate names, and other fields are suggested.

### Filtering

OFAC searches can add filters to include results of a certain type. These types can be individuals, businesses, aircraft, and vessels. SDN results can also be filtered by their OFAC program. Address searches can filter by Country

### Deep Inspection

OFAC searches can include exact matches on ID values (e.g. Government ID). These are in the "Remarks" section of each entity.

## Checks Not Performed

BSA/AML programs have requirements that are outside of Watchman, such as ownership calculations (thresholds, shell corporations, indirect majority shares), family relationships, and other risk analysis.

## Reporting

Watchman does not store search results or rankings. It's expected that users of Watchman store this information according to their risk and compliance needs.

A web UI is included with Watchman for inspecting OFAC results.

## Other Links

- [FDIC Bank Secrecy Act / Anti-Money Laundering](https://www.fdic.gov/resources/bankers/bank-secrecy-act/)
- [FFIEC BSA/AML Risk Assessment](https://bsaaml.ffiec.gov/manual/BSAAMLRiskAssessment/01)
- [Frequently Asked Questions Regarding Customer Due Diligence Requirements for Financial Institutions](https://www.fincen.gov/sites/default/files/2018-04/FinCEN_Guidance_CDD_FAQ_FINAL_508_2.pdf)
- [OFAC FAQ #249 - How is the Score calculated?](https://home.treasury.gov/policy-issues/financial-sanctions/faqs/topic/1636)
- [Sound Practices for Model Risk Management: Supervisory Guidance on Model Risk Management](https://www.occ.gov/news-issuances/bulletins/2011/bulletin-2011-12.html)
- [Application of Jaro-Winkler String Comparator in Enhancing Veterans Administrative Records](https://nces.ed.gov/FCSM/pdf/H_4HyoParkFCSM2018final.pdf)
- [Efficient Approximate Entity Matching Using Jaro-Winkler Distance](https://jqin.gitee.io/files/wise2017-wang.pdf)
- [On the Efficient Execution of Bounded Jaro-Winkler Distances](http://www.semantic-web-journal.net/system/files/swj1128.pdf)
