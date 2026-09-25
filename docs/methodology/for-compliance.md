---
layout: page
title: For compliance and risk
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Sanctions screening for compliance and risk

Watchman is an open-source screening engine you can inspect, tune, and defend in an exam. It downloads OFAC, EU, UK, UN, and related lists, scores each customer or counterparty against those lists, and returns a ranked hit with a **field-level score** — not a black-box yes/no.

This page is for BSA/AML officers, sanctions officers, model-risk reviewers, and internal audit. The technical scorer is described in [Similarity methodology](/watchman/methodology/). Numbers below come from scoring the full **OpenSanctions Pairs** dump (755,540 analyst-labeled pairs, 293 sources, 45 jurisdictions) with production `Similarity` in September 2026.

## The job screening has to do

Targeted financial sanctions require countries to **freeze without delay** the funds of designated persons and to stop those funds being made available to them ([FATF Recommendations 6 and 7](https://www.fatf-gafi.org/en/topics/fatf-recommendations.html); [FATF best practices on Recommendation 6](https://www.fatf-gafi.org/en/publications/Fatfrecommendations/Bpp-finsanctions-tf-r6.html)). In the United States, OFAC administers those lists and evaluates apparent violations under the [Economic Sanctions Enforcement Guidelines](https://www.ecfr.gov/current/title-31/subtitle-B/chapter-V/part-501/appendix-Appendix%20A%20to%20Part%20501) (31 CFR Part 501, Appendix A). OFAC’s [Framework for OFAC Compliance Commitments](https://ofac.treasury.gov/media/16331/download) (2019) expects a written program with management commitment, risk assessment, internal controls, testing, and training.

The Wolfsberg Group is explicit about what screening **is**: “the comparison of one string of text against another to detect similarities which would suggest a possible match,” used as one control inside a wider sanctions program, with a **risk-based** approach and known limits ([Wolfsberg Guidance on Sanctions Screening](https://wolfsberg-group.org/news/publication-of-guidance-on-sanctions-screening), 2019). Screening is not CDD, not transaction monitoring, and not a substitute for investigating a true hit.

Two failure modes dominate operations:

1. **A missed designation** (false negative) can become a prohibited transaction. That is the “without delay” problem.
2. **A flood of false hits** (false positives) burns analyst hours, delays good customers, and hides real alerts in noise. Name-only tools are famous for this. OFAC’s own public search returned **21 matches at 100%** for the query `Khamis Al`; Watchman’s name scorer on those same SDN names ranged from **0.26 to 0.87** ([portal comparison](/watchman/methodology/pages/ofac-name-comparison/)).

A usable engine has to move **both** numbers, and it has to show **why** a hit scored the way it did. That last point is model risk: OCC/Fed [SR 11-7 / OCC Bulletin 2011-12](https://www.occ.gov/news-issuances/bulletins/2011/bulletin-2011-12.html) expects conceptual soundness, ongoing monitoring, and effective challenge for models used in risk management.

## What Watchman gives a program

**Lists you can name.** OFAC SDN, the US Consolidated Screening List, EU, UK, UN, FinCEN 311, plus CSV ingest of your own files. Refresh is periodic or on demand. You can point examiners at which file was loaded and when.

**A score built from the fields CDD already collects.** FinCEN’s CDD rule expects institutions to know who they are dealing with. Watchman takes that same record — name, type, government IDs, date of birth, address, contact, vessel IMO — and scores it as a structured entity, not a single string. Name-only queries are down-ranked on purpose.

**Identity keys vs evidence.** A matching **passport, national ID, IMO, MMSI, aircraft serial, or crypto address** (identifier *and* country) still scores **1.0**. A matching **tax ID, business registration, or email** does not. Shared INNs across related companies were a primary false-positive driver on OpenSanctions Pairs when every identifier forced 1.0. Watchman now keeps those as high-weight evidence. If both sides have the **same ID type and country with different values** (two Khalid Mehmoods, two CNICs), the score is cut by `ID_CONFLICT_PENALTY_MULTIPLIER` (default 0.70). That is the paper’s Figure 1 case: OpenSanctions’ production rule matcher scored 0.98; Watchman does not treat the name as identity when the national IDs disagree ([Smith et al., 2026](https://arxiv.org/abs/2603.11051)).

**A threshold you set as policy.** `minMatch` is not a hidden neural cutoff. 0.80 is a high-precision screening line; ~0.59 is a high-recall line. You can document that choice in the risk assessment.

**An audit trail.** `debug=true` returns the pieces: which identifiers matched, what the name score was, whether an override or conflict penalty fired. That is the “effective challenge” artifact SR 11-7 asks for.

**Open source.** Apache 2.0. The scorer is Go you can read, test, and pin. You are not locked to a vendor’s unpublished matcher.

## Evidence: 755,540 labeled pairs

OpenSanctions Pairs is the first large public pairwise benchmark on real sanctions/OSINT data: **755,540** expert-labeled pairs over **1,002,093** records from **293 sources** in **45 jurisdictions**, 76.9% positive, 64.7% cross-source, seven Unicode scripts ([Smith, Sesodia, Lindenberg, Schroeder de Witt, 2026](https://arxiv.org/abs/2603.11051); [dataset](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs)). Labels come from OpenSanctions analysts in production deduplication, not crowd workers.

We scored the **entire dump** with Watchman’s production `Similarity` (pairwise: if both records are already in front of the scorer, does the score agree with the analyst?). That is not a live OFAC-coverage test and not a blocking test. It is the cleanest public answer to “how does this matcher treat hard, labeled sanctions pairs?”

**Subjects** (472,477 people, companies, vessels — Occupancy/Succession auto-merge dropped) at threshold **0.80**:

| Configuration (threshold 0.80) | Precision | Recall | F1 score | Cross-script recall |
|--------------------------------|----------:|-------:|---------:|--------------------:|
| Jaro–Winkler | **0.986** | 0.689 | 0.811 | 0.50 |
| Jaro–Winkler + TF-IDF | 0.968 | 0.707 | 0.817 | 0.52 |
| **Jaro–Winkler + cross-script embeddings** | 0.946 | **0.815** | **0.876** | **0.91** |

Read that as a **review-capacity vs miss-rate** trade:

- Default Jaro–Winkler at 0.80: about **3,000** subject false positives and **94,000** missed labeled positives. Precision is exam-friendly; recall is not, especially on transliteration.
- Adding **cross-script embeddings** (`qwen3-embedding:0.6b`, used only when scripts differ — the production `EMBEDDINGS_CROSS_SCRIPT_ONLY` analog): about **14,000** false positives and **56,000** misses. Precision stays **0.95**. Cross-script recall goes from **half the true matches** to **91%**.

Name-algorithm swaps (Soundex, Double Metaphone, Beider-Morse, nsim, Editex) moved F1 by about **0.001**. They do not fix Arabic/Cyrillic/Latin pairs. That matches the paper: rule matchers over-fire on common Latin names; learned methods fail on transliteration unless you add a representation that is not character-based.

**Vessels** (7,550 pairs): F1 **0.993**, recall 0.988 — IMO/MMSI as unique keys.

**Person/Person** (284,808 pairs): precision 0.983, recall 0.683 at 0.80 without embeddings. That is where most remaining misses live, and where embeddings earn their keep.

**Threshold as a documented control.** Same Jaro–Winkler scorer, subjects only:

| `minMatch` | Precision | Recall | F1 score |
|-----------:|----------:|-------:|---------:|
| 0.80 | 0.986 | 0.689 | 0.811 |
| 0.59 | 0.945 | 0.920 | 0.932 |

If the risk assessment says “we would rather review more alerts than miss a designation,” write down 0.59. If analyst capacity is the binding constraint, write down 0.80 and turn on cross-script embeddings.

### How this sits next to the paper’s baselines

OpenSanctions’ production rule matcher (nomenklatura RegressionV1) on their sampled sets: F1 91.3, **precision 84.5, recall 99.4** — catch everything, pay in review. GPT-4o pairwise: F1 99.0. That LLM number is a **binary judgment on the full FollowTheMoney JSON**, not a ranked screener with a 20 ms budget. Watchman at 0.80 is the other side of the curve (precision 0.99, recall 0.83 on the full dump). At 0.59, full-dump F1 is **0.960** with precision still above 0.96.

Use the LLM paper as a ceiling on **pairwise identity**, not as a replacement for an inspectable, list-driven control. Wolfsberg’s point still holds: screening compares strings (and, here, IDs and dates); it does not decide beneficial ownership.

Full tables: [OpenSanctions Pairs evaluation](/watchman/opensanctions-pairs/) and [`RESULTS.md`](https://github.com/moov-io/watchman/blob/master/research/opensanctions-pairs/RESULTS.md).

## What to tell an examiner

- **Lists and refresh.** Which files, which timestamp, how ingest of internal lists works.
- **Fields used.** Customer type, name and aliases, government IDs, DOB, address. Name-only is a degraded mode, not the design.
- **Cutoff.** `minMatch` and why (capacity vs miss-rate), with the OpenSanctions subject table as independent evidence.
- **Transliteration.** Whether embeddings are on, `CROSS_SCRIPT_ONLY`, and that hybrid recall on labeled cross-script pairs was 0.91 at 0.80.
- **False-positive design.** Tax IDs do not auto-confirm identity; conflicting national IDs penalize the score.
- **Explainability.** Stored `debug` pieces or equivalent logging of which fields drove the hit.
- **Change control.** Scoring-policy changes are in git, tested, and described in [methodology](/watchman/methodology/).

## How to run it in the program

1. Screen **onboarding, periodic refresh, and (where required) transactions** against OFAC plus the other lists your risk assessment names.
2. Always send **`type=`** (`person`, `business`, `vessel`, …). Unknown type: call person and business (and vessel/aircraft when those IDs exist).
3. Send **IDs and dates when CDD has them**. That is the difference between a 1.0 passport hit and a name-only 0.81.
4. Set **`minMatch=0.80`** with cross-script embeddings for a production-shaped queue; lower it if the board has accepted more review.
5. Route hits to investigation with the score pieces attached. True matches become blocks, rejects, and — where applicable — [FinCEN](https://www.fincen.gov/) reporting. False hits become tuning evidence.

Docker and config: [Usage](/watchman/usage-docker/), [Configuration](/watchman/config/). Search API: [Search](/watchman/search/).

## Sources

- Smith, Sesodia, Lindenberg, Schroeder de Witt. *OpenSanctions Pairs: A Large-Scale Dataset for Pairwise Entity Matching*. arXiv:2603.11051, 2026. [Paper](https://arxiv.org/abs/2603.11051) · [Dataset](https://huggingface.co/datasets/sanctions-er-anon/opensanctions_pairs)
- Fellegi, I. P., and Sunter, A. B. “A Theory for Record Linkage.” *JASA* 64 (328), 1969. [DOI](https://doi.org/10.1080/01621459.1969.10501049)
- Jaro, M. A. “Advances in Record-Linkage Methodology as Applied to Matching the 1985 Census of Tampa, Florida.” *JASA* 84 (406), 1989. [DOI](https://doi.org/10.1080/01621459.1989.10478785)
- Winkler, W. E. *String Comparator Metrics and Enhanced Decision Rules in the Fellegi-Sunter Model of Record Linkage.* ASA Proceedings, 1990. [ERIC ED325505](https://eric.ed.gov/?id=ED325505)
- FATF. *International Standards on Combating Money Laundering and the Financing of Terrorism & Proliferation* (Recommendations 6 and 7). [fatf-gafi.org](https://www.fatf-gafi.org/en/topics/fatf-recommendations.html)
- FATF. *International Best Practices: Targeted Financial Sanctions Related to Terrorism and Terrorist Financing (Recommendation 6).* [Publication](https://www.fatf-gafi.org/en/publications/Fatfrecommendations/Bpp-finsanctions-tf-r6.html)
- OFAC. *A Framework for OFAC Compliance Commitments* (May 2019). [PDF](https://ofac.treasury.gov/media/16331/download)
- OFAC. Economic Sanctions Enforcement Guidelines, 31 CFR Part 501 Appendix A. [eCFR](https://www.ecfr.gov/current/title-31/subtitle-B/chapter-V/part-501/appendix-Appendix%20A%20to%20Part%20501)
- OFAC. FAQ 249, SDN search scoring. [ofac.treasury.gov/faqs/249](https://ofac.treasury.gov/faqs/249)
- OCC / Federal Reserve. *Supervisory Guidance on Model Risk Management* (OCC Bulletin 2011-12 / SR 11-7). [OCC](https://www.occ.gov/news-issuances/bulletins/2011/bulletin-2011-12.html)
- The Wolfsberg Group. *Guidance on Sanctions Screening* (2019). [wolfsberg-group.org](https://wolfsberg-group.org/news/publication-of-guidance-on-sanctions-screening)
- FinCEN. *Frequently Asked Questions Regarding Customer Due Diligence Requirements for Financial Institutions* (2018). [PDF](https://www.fincen.gov/sites/default/files/2018-04/FinCEN_Guidance_CDD_FAQ_FINAL_508_2.pdf)
- FFIEC. *BSA/AML Examination Manual — Risk Assessment.* [bsaaml.ffiec.gov](https://bsaaml.ffiec.gov/manual/BSAAMLRiskAssessment/01)
- FDIC. Bank Secrecy Act / Anti-Money Laundering resources. [fdic.gov](https://www.fdic.gov/resources/bankers/bank-secrecy-act/)
