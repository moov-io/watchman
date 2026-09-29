# Cascade company and vessel name pairs (lot 2)

Labelled name pairs from [Cascade Screening](https://github.com/ArslaneSempai-ui/cascade-screening/tree/main/contrib/opensanctions), offered for Watchman in [discussion #919](https://github.com/moov-io/watchman/discussions/919) after the first 300-row fixture.

Source file: [`cascade-name-pairs-2.csv`](https://github.com/ArslaneSempai-ui/cascade-screening/blob/main/contrib/opensanctions/cascade-name-pairs-2.csv)

The original 185 rows are MIT-licensed. Watchman testdata adds `baseline` and `score` columns:

- `baseline` — `pkg/search.Similarity` on name-only entities the first time this fixture was recorded. Frozen.
- `score` — the same measurement after later scoring changes. Tests assert this column.

Both sides of a row use the CSV `schema` (`Company` → `business`, `Vessel` → `vessel`). Queries are name-only (no IDs, dates, or addresses), so an exact name lands around 0.855. Cross-script positives score 0 without embeddings.

Lot 2 is 118 matches and 67 non-matches, written as a yardstick for leftover Hebrew/Latin, Burmese/Latin, and former-name shapes (`ex-`, `f.k.a.`, `formerly`). Near-miss non-matches share a word, root, or numeral across scripts.

## Score history

Each scoring change rewrites `score` and the summary table. `baseline` stays at the first recorded run.

| Step | Match mean | Non-match mean |
|---|---:|---:|
| baseline | 0.2319 | 0.1295 |
| f.k.a. / formerly / chained and unparenthesized `ex` | 0.2846 | 0.1344 |

Rewrite `score` (and fill `baseline` when empty) with:

```
UPDATE_CASCADE_SCORES=yes go test ./pkg/search -run TestCascadeNamePairs -count=1
```

<!-- cascade-score-summary -->

Recorded `Similarity` on 185 name-only pairs (schema → Watchman type).

| Slice | N | Matches | Mean score | Mean baseline | Mean delta | Min | Max |
|---|---:|---:|---:|---:|---:|---:|---:|
| all | 185 | 118 | 0.2302 | 0.1948 | +0.0354 | 0.0000 | 0.8550 |
| is_match=true | 118 | 118 | 0.2846 | 0.2319 | +0.0527 | 0.0000 | 0.8550 |
| is_match=false | 67 | 0 | 0.1344 | 0.1295 | +0.0050 | 0.0000 | 0.8122 |
| alias-fka | 3 | 3 | 0.8122 | 0.3711 | +0.4411 | 0.8122 | 0.8122 |
| article-joined vs split | 3 | 3 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| barge-former-name-alias | 1 | 1 | 0.5686 | 0.5686 | +0.0000 | 0.5686 | 0.5686 |
| burmese-script-vs-latin | 9 | 9 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| chat-burmese-script | 1 | 1 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| company-script | 1 | 1 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| cross-script-close-root | 4 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| cross-script-close-syllable | 2 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| cross-script-numeral | 9 | 0 | 0.0164 | 0.0164 | +0.0000 | 0.0000 | 0.0742 |
| cross-script-other-name | 2 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| cross-script-other-word | 12 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| cross-script-place | 9 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| current-name-with-ex-residue | 1 | 1 | 0.8550 | 0.8550 | +0.0000 | 0.8550 | 0.8550 |
| ex-name | 3 | 3 | 0.8122 | 0.8122 | +0.0000 | 0.8122 | 0.8122 |
| ex-name-alias | 3 | 3 | 0.7310 | 0.7310 | +0.0000 | 0.5686 | 0.8122 |
| ex-name-alias-with-numeral | 1 | 1 | 0.8122 | 0.8122 | +0.0000 | 0.8122 | 0.8122 |
| ex-name-annotation | 1 | 1 | 0.8550 | 0.5746 | +0.2804 | 0.8550 | 0.8550 |
| ex-name-chain | 1 | 1 | 0.8122 | 0.1604 | +0.6518 | 0.8122 | 0.8122 |
| ex-name-collision-different-vessel | 2 | 0 | 0.8122 | 0.8122 | +0.0000 | 0.8122 | 0.8122 |
| ex-name-current | 1 | 1 | 0.8550 | 0.8550 | +0.0000 | 0.8550 | 0.8550 |
| ex-name-lookalike | 2 | 0 | 0.6550 | 0.6550 | +0.0000 | 0.5121 | 0.7979 |
| ex-name-numeral | 2 | 0 | 0.4291 | 0.4291 | +0.0000 | 0.4032 | 0.4549 |
| ex-name-numeral-trap | 2 | 0 | 0.7918 | 0.7918 | +0.0000 | 0.7901 | 0.7935 |
| ex-name-on-other-side | 2 | 2 | 0.8281 | 0.8281 | +0.0000 | 0.8013 | 0.8550 |
| ex-name-other-word | 1 | 0 | 0.5455 | 0.5455 | +0.0000 | 0.5455 | 0.5455 |
| ex-name-vs-numbered-name | 1 | 0 | 0.7948 | 0.7948 | +0.0000 | 0.7948 | 0.7948 |
| ex-name-vs-other-numeral | 1 | 0 | 0.8122 | 0.8122 | +0.0000 | 0.8122 | 0.8122 |
| ex-name-with-date | 1 | 1 | 0.8122 | 0.5081 | +0.3041 | 0.8122 | 0.8122 |
| ex-name-without-parentheses | 1 | 1 | 0.8550 | 0.4623 | +0.3927 | 0.8550 | 0.8550 |
| fka-former-name-match | 1 | 1 | 0.8122 | 0.4985 | +0.3137 | 0.8122 | 0.8122 |
| former-name | 1 | 1 | 0.8122 | 0.4697 | +0.3425 | 0.8122 | 0.8122 |
| former-name-alias | 3 | 3 | 0.7581 | 0.7581 | +0.0000 | 0.6498 | 0.8122 |
| former-name-appended | 1 | 1 | 0.8550 | 0.4371 | +0.4179 | 0.8550 | 0.8550 |
| former-name-trap | 1 | 0 | 0.7677 | 0.4353 | +0.3324 | 0.7677 | 0.7677 |
| hull-current-name-with-former-name | 3 | 3 | 0.8451 | 0.7341 | +0.1110 | 0.8254 | 0.8550 |
| hull-former-name-alias | 3 | 3 | 0.7581 | 0.7581 | +0.0000 | 0.6498 | 0.8122 |
| legal-form | 7 | 7 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| native-vs-english | 1 | 1 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| registry-former-name | 2 | 2 | 0.8122 | 0.4396 | +0.3726 | 0.8122 | 0.8122 |
| romanization-variant | 5 | 5 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| script-translation | 6 | 6 | 0.0155 | 0.0155 | +0.0000 | 0.0000 | 0.0932 |
| script-variant | 1 | 1 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| second-ex-name-alias | 1 | 1 | 0.8122 | 0.3546 | +0.4576 | 0.8122 | 0.8122 |
| second-former-name-alias | 1 | 1 | 0.8122 | 0.4745 | +0.3377 | 0.8122 | 0.8122 |
| spacing | 4 | 4 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| subsidiary-or-parent | 7 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| translated-generic-word | 10 | 10 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| translit-convention | 5 | 5 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| trap-former-name-alias | 2 | 2 | 0.8122 | 0.8122 | +0.0000 | 0.8122 | 0.8122 |
| vessel-burmese-script | 2 | 2 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-cross-script-numeral | 1 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-cross-script-other-word | 1 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-ex-name | 3 | 3 | 0.7310 | 0.7310 | +0.0000 | 0.5686 | 0.8122 |
| vessel-ex-name-appended | 1 | 1 | 0.8550 | 0.5363 | +0.3187 | 0.8550 | 0.8550 |
| vessel-ex-name-numeral | 1 | 0 | 0.5629 | 0.5629 | +0.0000 | 0.5629 | 0.5629 |
| vessel-ex-name-rename | 1 | 1 | 0.8550 | 0.8550 | +0.0000 | 0.8550 | 0.8550 |
| vessel-name-burmese-vs-latin-caps | 7 | 7 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-name-hebrew-vs-latin-caps | 6 | 6 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-script | 5 | 5 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| vessel-sister-ship | 7 | 0 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |
| word-order or abbreviation found on documents | 3 | 3 | 0.0000 | 0.0000 | +0.0000 | 0.0000 | 0.0000 |

<!-- /cascade-score-summary -->
