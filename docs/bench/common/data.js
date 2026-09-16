window.BENCHMARK_DATA = {
  "lastUpdate": 1789575985272,
  "repoUrl": "https://github.com/samofolabi/watchman_moov",
  "entries": {
    "moov-io/watchman Common Benchmarks": [
      {
        "commit": {
          "author": {
            "name": "user.email",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "user.email",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "c099c4268c526807e69330750db3ad4a5bc5312a",
          "message": "bench: keep skip-fetch-gh-pages on master\n\ngithub-action-benchmark fetches master:master, which git refuses\nwhen master is already checked out.",
          "timestamp": "2026-09-14T19:17:33Z",
          "url": "https://github.com/moov-io/watchman/commit/c099c4268c526807e69330750db3ad4a5bc5312a"
        },
        "date": 1789414360738,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search)",
            "value": 8601,
            "unit": "ns/op\t    2904 B/op\t      78 allocs/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 8601,
            "unit": "ns/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 2904,
            "unit": "B/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 27701,
            "unit": "ns/op\t   12394 B/op\t     129 allocs/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 27701,
            "unit": "ns/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 12394,
            "unit": "B/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 129,
            "unit": "allocs/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search)",
            "value": 15056,
            "unit": "ns/op\t    3472 B/op\t      86 allocs/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 15056,
            "unit": "ns/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 3472,
            "unit": "B/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 86,
            "unit": "allocs/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 32162,
            "unit": "ns/op\t   10714 B/op\t     122 allocs/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 32162,
            "unit": "ns/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 122,
            "unit": "allocs/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search)",
            "value": 1099,
            "unit": "ns/op\t     224 B/op\t       7 allocs/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 1099,
            "unit": "ns/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 12555,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 12555,
            "unit": "ns/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search)",
            "value": 73279,
            "unit": "ns/op\t   19835 B/op\t     465 allocs/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 73279,
            "unit": "ns/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 19835,
            "unit": "B/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 465,
            "unit": "allocs/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 112359,
            "unit": "ns/op\t   38555 B/op\t     599 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 112359,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 38555,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search)",
            "value": 17932265,
            "unit": "ns/op\t13404484 B/op\t  231470 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 17932265,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 13404484,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 231470,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search)",
            "value": 836605,
            "unit": "ns/op\t  552935 B/op\t    8141 allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 836605,
            "unit": "ns/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 552935,
            "unit": "B/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search)",
            "value": 2639811,
            "unit": "ns/op\t 2262142 B/op\t   13376 allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 2639811,
            "unit": "ns/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 2262142,
            "unit": "B/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 13376,
            "unit": "allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search)",
            "value": 839189,
            "unit": "ns/op\t  552497 B/op\t    8140 allocs/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 839189,
            "unit": "ns/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 552497,
            "unit": "B/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8140,
            "unit": "allocs/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search)",
            "value": 833293,
            "unit": "ns/op\t  552384 B/op\t    8141 allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 833293,
            "unit": "ns/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 552384,
            "unit": "B/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search)",
            "value": 835341,
            "unit": "ns/op\t  552205 B/op\t    8141 allocs/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 835341,
            "unit": "ns/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 552205,
            "unit": "B/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore)",
            "value": 4504,
            "unit": "ns/op\t     191 B/op\t       5 allocs/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 4504,
            "unit": "ns/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 191,
            "unit": "B/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore)",
            "value": 5786,
            "unit": "ns/op\t     556 B/op\t      12 allocs/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 5786,
            "unit": "ns/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 556,
            "unit": "B/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore)",
            "value": 67.33,
            "unit": "ns/op\t       4 B/op\t       1 allocs/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 67.33,
            "unit": "ns/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore)",
            "value": 143.5,
            "unit": "ns/op\t       8 B/op\t       2 allocs/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 143.5,
            "unit": "ns/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm)",
            "value": 36393,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - ns/op",
            "value": 36393,
            "unit": "ns/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32834 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "a9e04384fa11196351464a96b12ee335a1ab626e",
          "message": "Merge pull request #868 from toruiwasa/embeddings-env\n\nembeddings: read EMBEDDINGS_* from environment",
          "timestamp": "2026-09-15T19:54:54Z",
          "url": "https://github.com/samofolabi/watchman_moov/commit/a9e04384fa11196351464a96b12ee335a1ab626e"
        },
        "date": 1789575984833,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search)",
            "value": 8431,
            "unit": "ns/op\t    2904 B/op\t      78 allocs/op",
            "extra": "136132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 8431,
            "unit": "ns/op",
            "extra": "136132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 2904,
            "unit": "B/op",
            "extra": "136132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "136132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 27773,
            "unit": "ns/op\t   12394 B/op\t     129 allocs/op",
            "extra": "42752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 27773,
            "unit": "ns/op",
            "extra": "42752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 12394,
            "unit": "B/op",
            "extra": "42752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 129,
            "unit": "allocs/op",
            "extra": "42752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search)",
            "value": 14471,
            "unit": "ns/op\t    3472 B/op\t      86 allocs/op",
            "extra": "80772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 14471,
            "unit": "ns/op",
            "extra": "80772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 3472,
            "unit": "B/op",
            "extra": "80772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 86,
            "unit": "allocs/op",
            "extra": "80772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 29789,
            "unit": "ns/op\t   10714 B/op\t     122 allocs/op",
            "extra": "40020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 29789,
            "unit": "ns/op",
            "extra": "40020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "40020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 122,
            "unit": "allocs/op",
            "extra": "40020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search)",
            "value": 1117,
            "unit": "ns/op\t     224 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 1117,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 12545,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "94050 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 12545,
            "unit": "ns/op",
            "extra": "94050 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "94050 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "94050 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search)",
            "value": 72114,
            "unit": "ns/op\t   19836 B/op\t     466 allocs/op",
            "extra": "16186 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 72114,
            "unit": "ns/op",
            "extra": "16186 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 19836,
            "unit": "B/op",
            "extra": "16186 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "16186 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search)",
            "value": 109574,
            "unit": "ns/op\t   38556 B/op\t     599 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - ns/op",
            "value": 109574,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - B/op",
            "value": 38556,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug (github.com/moov-io/watchman/pkg/search) - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search)",
            "value": 17883834,
            "unit": "ns/op\t13418299 B/op\t  231463 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 17883834,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 13418299,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 231463,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search)",
            "value": 852192,
            "unit": "ns/op\t  584985 B/op\t    8163 allocs/op",
            "extra": "1327 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 852192,
            "unit": "ns/op",
            "extra": "1327 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 584985,
            "unit": "B/op",
            "extra": "1327 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8163,
            "unit": "allocs/op",
            "extra": "1327 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search)",
            "value": 2640591,
            "unit": "ns/op\t 2289701 B/op\t   13405 allocs/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 2640591,
            "unit": "ns/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 2289701,
            "unit": "B/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 13405,
            "unit": "allocs/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search)",
            "value": 853452,
            "unit": "ns/op\t  588626 B/op\t    8166 allocs/op",
            "extra": "1413 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 853452,
            "unit": "ns/op",
            "extra": "1413 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 588626,
            "unit": "B/op",
            "extra": "1413 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8166,
            "unit": "allocs/op",
            "extra": "1413 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search)",
            "value": 858079,
            "unit": "ns/op\t  588672 B/op\t    8166 allocs/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 858079,
            "unit": "ns/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 588672,
            "unit": "B/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8166,
            "unit": "allocs/op",
            "extra": "1417 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search)",
            "value": 956046,
            "unit": "ns/op\t  590346 B/op\t    8167 allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - ns/op",
            "value": 956046,
            "unit": "ns/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - B/op",
            "value": 590346,
            "unit": "B/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email (github.com/moov-io/watchman/internal/search) - allocs/op",
            "value": 8167,
            "unit": "allocs/op",
            "extra": "1404 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore)",
            "value": 4878,
            "unit": "ns/op\t     191 B/op\t       5 allocs/op",
            "extra": "213361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 4878,
            "unit": "ns/op",
            "extra": "213361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 191,
            "unit": "B/op",
            "extra": "213361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "213361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore)",
            "value": 6194,
            "unit": "ns/op\t     556 B/op\t      12 allocs/op",
            "extra": "172872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 6194,
            "unit": "ns/op",
            "extra": "172872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 556,
            "unit": "B/op",
            "extra": "172872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "172872 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore)",
            "value": 67.15,
            "unit": "ns/op\t       4 B/op\t       1 allocs/op",
            "extra": "17744864 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 67.15,
            "unit": "ns/op",
            "extra": "17744864 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "17744864 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17744864 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore)",
            "value": 142.5,
            "unit": "ns/op\t       8 B/op\t       2 allocs/op",
            "extra": "8407917 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - ns/op",
            "value": 142.5,
            "unit": "ns/op",
            "extra": "8407917 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8407917 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch (github.com/moov-io/watchman/internal/stringscore) - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8407917 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm)",
            "value": 36940,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32464 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - ns/op",
            "value": 36940,
            "unit": "ns/op",
            "extra": "32464 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32464 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber (github.com/moov-io/watchman/internal/norm) - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32464 times\n4 procs"
          }
        ]
      }
    ]
  }
}