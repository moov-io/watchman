# Cascade company and vessel name pairs

Labelled name pairs from [Cascade Screening](https://github.com/ArslaneSempai-ui/cascade-screening/tree/main/contrib/opensanctions), offered for Watchman in [discussion #919](https://github.com/moov-io/watchman/discussions/919) and [opensanctions/nomenklatura#380](https://github.com/opensanctions/nomenklatura/pull/380).

Source file: [`cascade-name-pairs.csv`](https://github.com/ArslaneSempai-ui/cascade-screening/blob/main/contrib/opensanctions/cascade-name-pairs.csv)

A second lot of 185 Hebrew, Burmese, and former-name pairs is [`cascade-name-pairs-2.csv`](cascade-name-pairs-2.csv). Notes for that file are in [`cascade-name-pairs-2.md`](cascade-name-pairs-2.md).

The original 300 rows are MIT-licensed. Watchman testdata adds `baseline` and `score` columns:

- `baseline` — `pkg/search.Similarity` on name-only entities the first time this fixture was recorded. Frozen.
- `score` — the same measurement after later scoring changes. Tests assert this column.

Both sides of a row use the CSV `schema` (`Company` → `business`, `Vessel` → `vessel`). Queries are name-only (no IDs, dates, or addresses), so an exact name lands around 0.855. Cross-script positives score 0 without embeddings, which pulls the match mean below the non-match mean.

## Score history

Each scoring change rewrites `score` and the summary table. `baseline` stays at the first recorded run.

| Step | Match mean | Non-match mean |
|---|---:|---:|
| baseline | 0.5574 | 0.6279 |
| strip vessel prefixes (`MV`, `M/T`, `SS`, …) | 0.5574 | 0.6296 |
| strip trailing vessel port/place | 0.5710 | 0.6296 |
| strip leading SHIPPER:/BENEFICIARY:/MESSRS./FIELD 59: | 0.5784 | 0.6296 |
| treat (ex-Name) as a former name on the query and the list | 0.5919 | 0.6296 |
| penalize disagreeing name numerals (digit, roman, spelled) | 0.5919 | 0.6124 |
| canonicalize English legal-form phrases (`Limited`→`ltd`, …) | 0.6149 | 0.6123 |

Rewrite `score` (and fill `baseline` when empty) with:

```
UPDATE_CASCADE_SCORES=yes go test ./pkg/search -run TestCascadeNamePairs -count=1
```

<!-- cascade-score-summary -->

Recorded `Similarity` on 300 name-only pairs (schema → Watchman type).

| Slice | N | Matches | Mean score | Mean baseline | Mean delta | Min | Max |
|---|---:|---:|---:|---:|---:|---:|---:|
| all | 300 | 140 | 0.6135 | 0.5950 | +0.0185 | 0.0000 | 0.8550 |
| is_match=true | 140 | 140 | 0.6149 | 0.5574 | +0.0575 | 0.0000 | 0.8550 |
| is_match=false | 160 | 0 | 0.6123 | 0.6279 | -0.0155 | 0.0000 | 0.8451 |
| abbreviation | 4 | 4 | 0.6052 | 0.6052 | +0.0000 | 0.4538 | 0.8408 |
| abbreviation-expanded | 4 | 4 | 0.6942 | 0.6942 | +0.0000 | 0.5577 | 0.8311 |
| accent-diacritic-dropped | 4 | 4 | 0.8394 | 0.8394 | +0.0000 | 0.7927 | 0.8550 |
| accent-variant | 4 | 4 | 0.8550 | 0.8550 | +0.0000 | 0.8550 | 0.8550 |
| annotation | 4 | 4 | 0.5077 | 0.3994 | +0.1083 | 0.3597 | 0.7837 |
| branch-same-entity | 4 | 4 | 0.6165 | 0.6003 | +0.0162 | 0.5332 | 0.8060 |
| burmese-script-vs-latin | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| case-punctuation-spacing | 4 | 4 | 0.8427 | 0.8427 | +0.0000 | 0.8058 | 0.8550 |
| different-business-sector-word | 4 | 0 | 0.5002 | 0.5002 | +0.0000 | 0.4032 | 0.5664 |
| different-country-registration | 4 | 0 | 0.5843 | 0.5843 | +0.0000 | 0.4994 | 0.8160 |
| different-name-word | 4 | 0 | 0.6922 | 0.6922 | +0.0000 | 0.5778 | 0.7936 |
| different-place | 4 | 0 | 0.7344 | 0.7344 | +0.0000 | 0.5911 | 0.7830 |
| different-place-same-business-word | 4 | 0 | 0.5747 | 0.5747 | +0.0000 | 0.5549 | 0.5908 |
| different-surname | 4 | 0 | 0.6267 | 0.6267 | +0.0000 | 0.5171 | 0.7855 |
| different-surname-same-form | 4 | 0 | 0.6163 | 0.6163 | +0.0000 | 0.5169 | 0.7700 |
| different-surname-same-form-chat | 4 | 0 | 0.4049 | 0.4049 | +0.0000 | 0.3226 | 0.5340 |
| different-trade-word | 4 | 0 | 0.6262 | 0.6262 | +0.0000 | 0.5264 | 0.8037 |
| doubled-letter | 4 | 4 | 0.8350 | 0.8350 | +0.0000 | 0.8240 | 0.8404 |
| ex-name-alias | 4 | 4 | 0.8122 | 0.4476 | +0.3646 | 0.8122 | 0.8122 |
| family-name-different-person | 4 | 0 | 0.6329 | 0.6329 | +0.0000 | 0.5546 | 0.8173 |
| foreign-subsidiary | 4 | 0 | 0.7463 | 0.7463 | +0.0000 | 0.5931 | 0.8200 |
| generic-words-only | 4 | 0 | 0.5664 | 0.5664 | +0.0000 | 0.5450 | 0.5942 |
| holding-vs-operating | 4 | 0 | 0.4746 | 0.4746 | +0.0000 | 0.0000 | 0.8219 |
| homophone | 4 | 0 | 0.7140 | 0.7140 | +0.0000 | 0.3676 | 0.8333 |
| legal-form-abbreviated | 4 | 4 | 0.8550 | 0.5712 | +0.2838 | 0.8550 | 0.8550 |
| legal-form-spelled-out | 4 | 4 | 0.8550 | 0.6777 | +0.1772 | 0.8550 | 0.8550 |
| legal-form-spelt-out | 4 | 4 | 0.4991 | 0.4991 | +0.0000 | 0.3762 | 0.5616 |
| legal-form-translated | 4 | 4 | 0.7980 | 0.7500 | +0.0479 | 0.7798 | 0.8310 |
| missing-space | 4 | 4 | 0.5524 | 0.5524 | +0.0000 | 0.3869 | 0.8550 |
| native-vs-english | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| near-string | 4 | 0 | 0.7090 | 0.7090 | +0.0000 | 0.5742 | 0.8384 |
| near-string-real-words | 4 | 0 | 0.7091 | 0.7108 | -0.0017 | 0.5688 | 0.8415 |
| numbered-spv | 4 | 0 | 0.5929 | 0.8470 | -0.2541 | 0.5871 | 0.5985 |
| numeral-changed | 4 | 0 | 0.5733 | 0.8189 | -0.2457 | 0.5652 | 0.5786 |
| one-word-sibling | 4 | 0 | 0.5922 | 0.5922 | +0.0000 | 0.4954 | 0.8112 |
| one-word-sibling-activity | 4 | 0 | 0.5580 | 0.5580 | +0.0000 | 0.5180 | 0.6118 |
| other-business-word | 4 | 0 | 0.6530 | 0.6530 | +0.0000 | 0.5198 | 0.7866 |
| other-name | 4 | 0 | 0.6488 | 0.6488 | +0.0000 | 0.5583 | 0.8218 |
| other-place | 4 | 0 | 0.6190 | 0.6190 | +0.0000 | 0.5233 | 0.7980 |
| other-proper-name | 4 | 0 | 0.6906 | 0.6906 | +0.0000 | 0.5745 | 0.8046 |
| other-surname | 4 | 0 | 0.8040 | 0.8040 | +0.0000 | 0.7800 | 0.8229 |
| other-word | 4 | 0 | 0.6161 | 0.6161 | +0.0000 | 0.5177 | 0.8110 |
| plural-vs-singular | 4 | 0 | 0.8379 | 0.8379 | +0.0000 | 0.8314 | 0.8451 |
| port-of-registry-appended | 4 | 4 | 0.8550 | 0.7304 | +0.1246 | 0.8550 | 0.8550 |
| punctuation-spacing | 4 | 4 | 0.7672 | 0.7672 | +0.0000 | 0.5845 | 0.8550 |
| punctuation-spacing-legal-form | 4 | 4 | 0.8536 | 0.8536 | +0.0000 | 0.8494 | 0.8550 |
| registry-written-out | 4 | 4 | 0.6520 | 0.6348 | +0.0171 | 0.5599 | 0.8550 |
| residue-party-label | 4 | 4 | 0.8444 | 0.5951 | +0.2492 | 0.8173 | 0.8550 |
| romanisation | 4 | 4 | 0.8183 | 0.8137 | +0.0046 | 0.7980 | 0.8415 |
| romanisation-variant | 4 | 4 | 0.7473 | 0.6853 | +0.0620 | 0.5224 | 0.8380 |
| russian-based-spelling | 4 | 4 | 0.1257 | 0.1257 | +0.0000 | 0.0000 | 0.5028 |
| same-name-other-jurisdiction | 4 | 0 | 0.6864 | 0.6864 | +0.0000 | 0.5972 | 0.7785 |
| same-word-unrelated-business | 4 | 0 | 0.3951 | 0.3951 | +0.0000 | 0.2766 | 0.5071 |
| script | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| script-translation | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| sibling-one-word | 4 | 0 | 0.6140 | 0.6140 | +0.0000 | 0.5226 | 0.8112 |
| sister-ship | 4 | 0 | 0.5223 | 0.5593 | -0.0370 | 0.3648 | 0.6035 |
| subsidiary-by-activity | 4 | 0 | 0.4873 | 0.4873 | +0.0000 | 0.3933 | 0.5456 |
| subsidiary-by-country | 4 | 0 | 0.5706 | 0.5706 | +0.0000 | 0.5128 | 0.5984 |
| suffix-abbreviation | 4 | 4 | 0.7966 | 0.6433 | +0.1533 | 0.7762 | 0.8550 |
| translation | 4 | 4 | 0.6113 | 0.5718 | +0.0395 | 0.5093 | 0.8059 |
| transliteration | 4 | 4 | 0.6133 | 0.6133 | +0.0000 | 0.3716 | 0.8143 |
| two-words-differ | 4 | 0 | 0.5112 | 0.5112 | +0.0000 | 0.3872 | 0.5810 |
| uppercase-export | 4 | 4 | 0.8550 | 0.8475 | +0.0075 | 0.8550 | 0.8550 |
| vessel-numeral | 4 | 0 | 0.6030 | 0.7899 | -0.1869 | 0.5786 | 0.6674 |
| vessel-other-word | 4 | 0 | 0.5045 | 0.4959 | +0.0086 | 0.0000 | 0.7981 |
| vessel-plural | 4 | 0 | 0.7921 | 0.7921 | +0.0000 | 0.7772 | 0.8192 |
| vessel-prefix | 4 | 4 | 0.8550 | 0.8265 | +0.0285 | 0.8550 | 0.8550 |
| vessel-prefix-and-port | 4 | 4 | 0.8024 | 0.4791 | +0.3234 | 0.6662 | 0.8550 |
| vessel-script | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-sister-ship | 4 | 0 | 0.5229 | 0.5229 | +0.0000 | 0.4412 | 0.5858 |
| vessel-transliteration | 4 | 4 | 0.5753 | 0.5753 | +0.0000 | 0.3680 | 0.8284 |
| vessel-vs-manager | 4 | 0 | 0.5151 | 0.4872 | +0.0278 | 0.4841 | 0.5388 |
| vessel-vs-owner | 4 | 0 | 0.6717 | 0.6032 | +0.0684 | 0.5055 | 0.8007 |
| word-abbreviation | 4 | 4 | 0.5810 | 0.5779 | +0.0031 | 0.5678 | 0.6064 |

<!-- /cascade-score-summary -->
