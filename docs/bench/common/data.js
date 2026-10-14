window.BENCHMARK_DATA = {
  "lastUpdate": 1791992320964,
  "repoUrl": "https://github.com/moov-io/watchman",
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
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8601,
            "unit": "ns/op\t    2904 B/op\t      78 allocs/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8601,
            "unit": "ns/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2904,
            "unit": "B/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "134442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 27701,
            "unit": "ns/op\t   12394 B/op\t     129 allocs/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 27701,
            "unit": "ns/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12394,
            "unit": "B/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 129,
            "unit": "allocs/op",
            "extra": "42102 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15056,
            "unit": "ns/op\t    3472 B/op\t      86 allocs/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15056,
            "unit": "ns/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3472,
            "unit": "B/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 86,
            "unit": "allocs/op",
            "extra": "80820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32162,
            "unit": "ns/op\t   10714 B/op\t     122 allocs/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32162,
            "unit": "ns/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 122,
            "unit": "allocs/op",
            "extra": "37992 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1099,
            "unit": "ns/op\t     224 B/op\t       7 allocs/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1099,
            "unit": "ns/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "997822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12555,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12555,
            "unit": "ns/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "93198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 73279,
            "unit": "ns/op\t   19835 B/op\t     465 allocs/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 73279,
            "unit": "ns/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 19835,
            "unit": "B/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 465,
            "unit": "allocs/op",
            "extra": "16189 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 112359,
            "unit": "ns/op\t   38555 B/op\t     599 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 112359,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38555,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 17932265,
            "unit": "ns/op\t13404484 B/op\t  231470 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 17932265,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 13404484,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 231470,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 836605,
            "unit": "ns/op\t  552935 B/op\t    8141 allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 836605,
            "unit": "ns/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 552935,
            "unit": "B/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1339 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 2639811,
            "unit": "ns/op\t 2262142 B/op\t   13376 allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 2639811,
            "unit": "ns/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 2262142,
            "unit": "B/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 13376,
            "unit": "allocs/op",
            "extra": "464 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 839189,
            "unit": "ns/op\t  552497 B/op\t    8140 allocs/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 839189,
            "unit": "ns/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 552497,
            "unit": "B/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 8140,
            "unit": "allocs/op",
            "extra": "1423 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 833293,
            "unit": "ns/op\t  552384 B/op\t    8141 allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 833293,
            "unit": "ns/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 552384,
            "unit": "B/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1434 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 835341,
            "unit": "ns/op\t  552205 B/op\t    8141 allocs/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 835341,
            "unit": "ns/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 552205,
            "unit": "B/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 8141,
            "unit": "allocs/op",
            "extra": "1381 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4504,
            "unit": "ns/op\t     191 B/op\t       5 allocs/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4504,
            "unit": "ns/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 191,
            "unit": "B/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "259255 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5786,
            "unit": "ns/op\t     556 B/op\t      12 allocs/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5786,
            "unit": "ns/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 556,
            "unit": "B/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "218032 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 67.33,
            "unit": "ns/op\t       4 B/op\t       1 allocs/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 67.33,
            "unit": "ns/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17301310 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 143.5,
            "unit": "ns/op\t       8 B/op\t       2 allocs/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 143.5,
            "unit": "ns/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8247566 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36393,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36393,
            "unit": "ns/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
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
          "url": "https://github.com/moov-io/watchman/commit/a9e04384fa11196351464a96b12ee335a1ab626e"
        },
        "date": 1789577060262,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8443,
            "unit": "ns/op\t    2904 B/op\t      78 allocs/op",
            "extra": "136584 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8443,
            "unit": "ns/op",
            "extra": "136584 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2904,
            "unit": "B/op",
            "extra": "136584 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "136584 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 27962,
            "unit": "ns/op\t   12394 B/op\t     129 allocs/op",
            "extra": "40509 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 27962,
            "unit": "ns/op",
            "extra": "40509 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12394,
            "unit": "B/op",
            "extra": "40509 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 129,
            "unit": "allocs/op",
            "extra": "40509 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14624,
            "unit": "ns/op\t    3472 B/op\t      86 allocs/op",
            "extra": "79980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14624,
            "unit": "ns/op",
            "extra": "79980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3472,
            "unit": "B/op",
            "extra": "79980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 86,
            "unit": "allocs/op",
            "extra": "79980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30098,
            "unit": "ns/op\t   10714 B/op\t     122 allocs/op",
            "extra": "39595 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30098,
            "unit": "ns/op",
            "extra": "39595 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "39595 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 122,
            "unit": "allocs/op",
            "extra": "39595 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1131,
            "unit": "ns/op\t     224 B/op\t       7 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1131,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12757,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "94016 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12757,
            "unit": "ns/op",
            "extra": "94016 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "94016 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "94016 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 72456,
            "unit": "ns/op\t   19836 B/op\t     466 allocs/op",
            "extra": "16400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 72456,
            "unit": "ns/op",
            "extra": "16400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 19836,
            "unit": "B/op",
            "extra": "16400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "16400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 110093,
            "unit": "ns/op\t   38555 B/op\t     599 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 110093,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38555,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23781071,
            "unit": "ns/op\t13402848 B/op\t  231470 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23781071,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 13402848,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 231470,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 841194,
            "unit": "ns/op\t  584671 B/op\t    8163 allocs/op",
            "extra": "1341 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 841194,
            "unit": "ns/op",
            "extra": "1341 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 584671,
            "unit": "B/op",
            "extra": "1341 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 8163,
            "unit": "allocs/op",
            "extra": "1341 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 2641175,
            "unit": "ns/op\t 2269530 B/op\t   13385 allocs/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 2641175,
            "unit": "ns/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 2269530,
            "unit": "B/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 13385,
            "unit": "allocs/op",
            "extra": "469 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 840660,
            "unit": "ns/op\t  588583 B/op\t    8166 allocs/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 840660,
            "unit": "ns/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 588583,
            "unit": "B/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 8166,
            "unit": "allocs/op",
            "extra": "1441 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 1189031,
            "unit": "ns/op\t  562557 B/op\t    8147 allocs/op",
            "extra": "1320 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 1189031,
            "unit": "ns/op",
            "extra": "1320 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 562557,
            "unit": "B/op",
            "extra": "1320 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 8147,
            "unit": "allocs/op",
            "extra": "1320 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 924295,
            "unit": "ns/op\t  559837 B/op\t    8146 allocs/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 924295,
            "unit": "ns/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 559837,
            "unit": "B/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 8146,
            "unit": "allocs/op",
            "extra": "1273 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4471,
            "unit": "ns/op\t     191 B/op\t       5 allocs/op",
            "extra": "271171 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4471,
            "unit": "ns/op",
            "extra": "271171 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 191,
            "unit": "B/op",
            "extra": "271171 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "271171 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5815,
            "unit": "ns/op\t     556 B/op\t      12 allocs/op",
            "extra": "206174 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5815,
            "unit": "ns/op",
            "extra": "206174 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 556,
            "unit": "B/op",
            "extra": "206174 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "206174 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 67.39,
            "unit": "ns/op\t       4 B/op\t       1 allocs/op",
            "extra": "17374753 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 67.39,
            "unit": "ns/op",
            "extra": "17374753 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "17374753 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17374753 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 142.8,
            "unit": "ns/op\t       8 B/op\t       2 allocs/op",
            "extra": "8256486 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 142.8,
            "unit": "ns/op",
            "extra": "8256486 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8256486 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8256486 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36559,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32719 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36559,
            "unit": "ns/op",
            "extra": "32719 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32719 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32719 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "4813dbaea2b4dd9657522c90aef2c1ad22c57a9d",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.48 (#965)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-14T04:56:15Z",
          "url": "https://github.com/moov-io/watchman/commit/4813dbaea2b4dd9657522c90aef2c1ad22c57a9d"
        },
        "date": 1791988232334,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9514,
            "unit": "ns/op\t    2112 B/op\t      48 allocs/op",
            "extra": "116751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9514,
            "unit": "ns/op",
            "extra": "116751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "116751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "116751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30781,
            "unit": "ns/op\t   11819 B/op\t     100 allocs/op",
            "extra": "38305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30781,
            "unit": "ns/op",
            "extra": "38305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11819,
            "unit": "B/op",
            "extra": "38305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "38305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14021,
            "unit": "ns/op\t    1614 B/op\t      35 allocs/op",
            "extra": "79784 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14021,
            "unit": "ns/op",
            "extra": "79784 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "79784 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "79784 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32995,
            "unit": "ns/op\t   10097 B/op\t      71 allocs/op",
            "extra": "35995 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32995,
            "unit": "ns/op",
            "extra": "35995 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10097,
            "unit": "B/op",
            "extra": "35995 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "35995 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1735,
            "unit": "ns/op\t     320 B/op\t       9 allocs/op",
            "extra": "712392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1735,
            "unit": "ns/op",
            "extra": "712392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "712392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "712392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 15351,
            "unit": "ns/op\t    5529 B/op\t      28 allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 15351,
            "unit": "ns/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5529,
            "unit": "B/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 65803,
            "unit": "ns/op\t   10009 B/op\t     205 allocs/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 65803,
            "unit": "ns/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18166 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 107824,
            "unit": "ns/op\t   28940 B/op\t     338 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 107824,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28940,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 17702829,
            "unit": "ns/op\t 6740072 B/op\t  152704 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 17702829,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740072,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152704,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 668650,
            "unit": "ns/op\t  256267 B/op\t    5000 allocs/op",
            "extra": "1737 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 668650,
            "unit": "ns/op",
            "extra": "1737 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256267,
            "unit": "B/op",
            "extra": "1737 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1737 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 855692,
            "unit": "ns/op\t  340399 B/op\t    5330 allocs/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 855692,
            "unit": "ns/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 340399,
            "unit": "B/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1429 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 670317,
            "unit": "ns/op\t  258107 B/op\t    5002 allocs/op",
            "extra": "1875 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 670317,
            "unit": "ns/op",
            "extra": "1875 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258107,
            "unit": "B/op",
            "extra": "1875 times\n4 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1875 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 1150513,
            "unit": "ns/op\t  274682 B/op\t    5455 allocs/op",
            "extra": "1544 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 1150513,
            "unit": "ns/op",
            "extra": "1544 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 274682,
            "unit": "B/op",
            "extra": "1544 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5455,
            "unit": "allocs/op",
            "extra": "1544 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1033498,
            "unit": "ns/op\t  365109 B/op\t    5644 allocs/op",
            "extra": "1149 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1033498,
            "unit": "ns/op",
            "extra": "1149 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 365109,
            "unit": "B/op",
            "extra": "1149 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5644,
            "unit": "allocs/op",
            "extra": "1149 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 784097,
            "unit": "ns/op\t  276480 B/op\t    5459 allocs/op",
            "extra": "1563 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 784097,
            "unit": "ns/op",
            "extra": "1563 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 276480,
            "unit": "B/op",
            "extra": "1563 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5459,
            "unit": "allocs/op",
            "extra": "1563 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 785250,
            "unit": "ns/op\t  274301 B/op\t    5454 allocs/op",
            "extra": "1594 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 785250,
            "unit": "ns/op",
            "extra": "1594 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 274301,
            "unit": "B/op",
            "extra": "1594 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5454,
            "unit": "allocs/op",
            "extra": "1594 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 774666,
            "unit": "ns/op\t  274266 B/op\t    5453 allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 774666,
            "unit": "ns/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 274266,
            "unit": "B/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5453,
            "unit": "allocs/op",
            "extra": "1538 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 784053,
            "unit": "ns/op\t  274437 B/op\t    5455 allocs/op",
            "extra": "1318 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 784053,
            "unit": "ns/op",
            "extra": "1318 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 274437,
            "unit": "B/op",
            "extra": "1318 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5455,
            "unit": "allocs/op",
            "extra": "1318 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3728,
            "unit": "ns/op\t      50 B/op\t       1 allocs/op",
            "extra": "303333 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3728,
            "unit": "ns/op",
            "extra": "303333 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "303333 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "303333 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5255,
            "unit": "ns/op\t     351 B/op\t       6 allocs/op",
            "extra": "213199 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5255,
            "unit": "ns/op",
            "extra": "213199 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "213199 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "213199 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 70.39,
            "unit": "ns/op\t       4 B/op\t       1 allocs/op",
            "extra": "16815378 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 70.39,
            "unit": "ns/op",
            "extra": "16815378 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16815378 times\n4 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16815378 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 148.9,
            "unit": "ns/op\t       8 B/op\t       2 allocs/op",
            "extra": "7839385 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 148.9,
            "unit": "ns/op",
            "extra": "7839385 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "7839385 times\n4 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7839385 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36368,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36368,
            "unit": "ns/op",
            "extra": "32762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32762 times\n4 procs"
          }
        ]
      },
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
          "id": "041d814c2a8636559e02412dc26e27877b6e0b9b",
          "message": "ci: store median of 6 bench samples at 2 CPUs\n\nRun go test -count=6 -cpu=2 and collapse each benchmark to a median\nline before github-action-benchmark records it. Raise alert-threshold\nto 150% and pin ubuntu-24.04.",
          "timestamp": "2026-10-14T14:53:19Z",
          "url": "https://github.com/moov-io/watchman/commit/041d814c2a8636559e02412dc26e27877b6e0b9b"
        },
        "date": 1791990230888,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8658,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "136089 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8658,
            "unit": "ns/op",
            "extra": "136089 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "136089 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "136089 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 29043.5,
            "unit": "ns/op\t11818 B/op\t100 allocs/op",
            "extra": "41050 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 29043.5,
            "unit": "ns/op",
            "extra": "41050 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11818,
            "unit": "B/op",
            "extra": "41050 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "41050 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13947.5,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "85632 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13947.5,
            "unit": "ns/op",
            "extra": "85632 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "85632 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "85632 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32110,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "37896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32110,
            "unit": "ns/op",
            "extra": "37896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "37896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "37896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1667,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "704385 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1667,
            "unit": "ns/op",
            "extra": "704385 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "704385 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "704385 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14409,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "82290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14409,
            "unit": "ns/op",
            "extra": "82290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "82290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "82290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 66105,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "18136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 66105,
            "unit": "ns/op",
            "extra": "18136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "18136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 105291,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 105291,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21012857.5,
            "unit": "ns/op\t6738928 B/op\t152702 allocs/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21012857.5,
            "unit": "ns/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6738928,
            "unit": "B/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152702,
            "unit": "allocs/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 728747.5,
            "unit": "ns/op\t256274 B/op\t5000 allocs/op",
            "extra": "1705 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 728747.5,
            "unit": "ns/op",
            "extra": "1705 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256274,
            "unit": "B/op",
            "extra": "1705 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1705 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 894756.5,
            "unit": "ns/op\t340398 B/op\t5330 allocs/op",
            "extra": "1468 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 894756.5,
            "unit": "ns/op",
            "extra": "1468 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 340398,
            "unit": "B/op",
            "extra": "1468 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1468 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 737929.5,
            "unit": "ns/op\t258114 B/op\t5002 allocs/op",
            "extra": "1616 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 737929.5,
            "unit": "ns/op",
            "extra": "1616 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258114,
            "unit": "B/op",
            "extra": "1616 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1616 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 865768,
            "unit": "ns/op\t273982 B/op\t5451 allocs/op",
            "extra": "1464 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 865768,
            "unit": "ns/op",
            "extra": "1464 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 273982,
            "unit": "B/op",
            "extra": "1464 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5451,
            "unit": "allocs/op",
            "extra": "1464 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1007522,
            "unit": "ns/op\t364883 B/op\t5650 allocs/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1007522,
            "unit": "ns/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 364883,
            "unit": "B/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5650,
            "unit": "allocs/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 851629,
            "unit": "ns/op\t277081 B/op\t5466 allocs/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 851629,
            "unit": "ns/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 277081,
            "unit": "B/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5466,
            "unit": "allocs/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 840301.5,
            "unit": "ns/op\t272728 B/op\t5438 allocs/op",
            "extra": "1428 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 840301.5,
            "unit": "ns/op",
            "extra": "1428 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 272728,
            "unit": "B/op",
            "extra": "1428 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5438,
            "unit": "allocs/op",
            "extra": "1428 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 841555,
            "unit": "ns/op\t276084 B/op\t5472 allocs/op",
            "extra": "1467 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 841555,
            "unit": "ns/op",
            "extra": "1467 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 276084,
            "unit": "B/op",
            "extra": "1467 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5472,
            "unit": "allocs/op",
            "extra": "1467 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 839223,
            "unit": "ns/op\t273866 B/op\t5450 allocs/op",
            "extra": "1531 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 839223,
            "unit": "ns/op",
            "extra": "1531 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 273866,
            "unit": "B/op",
            "extra": "1531 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5450,
            "unit": "allocs/op",
            "extra": "1531 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4338,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "286633 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4338,
            "unit": "ns/op",
            "extra": "286633 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "286633 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "286633 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5935.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "193741 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5935.5,
            "unit": "ns/op",
            "extra": "193741 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "193741 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "193741 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 69.3,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "16764919 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 69.3,
            "unit": "ns/op",
            "extra": "16764919 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16764919 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16764919 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.8,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8240264 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.8,
            "unit": "ns/op",
            "extra": "8240264 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8240264 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8240264 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36038.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "33319 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36038.5,
            "unit": "ns/op",
            "extra": "33319 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33319 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33319 times\n2 procs"
          }
        ]
      },
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
          "id": "e9f80f8b047b47db4f63012313de0db2c0dc9590",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:54Z",
          "url": "https://github.com/moov-io/watchman/commit/e9f80f8b047b47db4f63012313de0db2c0dc9590"
        },
        "date": 1791991389037,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 6419,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "181995 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 6419,
            "unit": "ns/op",
            "extra": "181995 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "181995 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "181995 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 20911.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "57360 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 20911.5,
            "unit": "ns/op",
            "extra": "57360 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "57360 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "57360 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 9815,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "119697 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 9815,
            "unit": "ns/op",
            "extra": "119697 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "119697 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "119697 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 21540,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "54488 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 21540,
            "unit": "ns/op",
            "extra": "54488 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "54488 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "54488 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1372,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "840744 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1372,
            "unit": "ns/op",
            "extra": "840744 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "840744 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "840744 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 10431,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "114856 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 10431,
            "unit": "ns/op",
            "extra": "114856 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "114856 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "114856 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44523,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "26820 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44523,
            "unit": "ns/op",
            "extra": "26820 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "26820 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "26820 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 71226,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "16936 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 71226,
            "unit": "ns/op",
            "extra": "16936 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "16936 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "16936 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 17062152.5,
            "unit": "ns/op\t6740014 B/op\t152701 allocs/op",
            "extra": "68 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 17062152.5,
            "unit": "ns/op",
            "extra": "68 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740014,
            "unit": "B/op",
            "extra": "68 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152701,
            "unit": "allocs/op",
            "extra": "68 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 564851.5,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "2200 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 564851.5,
            "unit": "ns/op",
            "extra": "2200 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "2200 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "2200 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 701110.5,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1884 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 701110.5,
            "unit": "ns/op",
            "extra": "1884 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1884 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1884 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 555062.5,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "2377 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 555062.5,
            "unit": "ns/op",
            "extra": "2377 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "2377 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "2377 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 648154,
            "unit": "ns/op\t275291 B/op\t5455 allocs/op",
            "extra": "1899 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 648154,
            "unit": "ns/op",
            "extra": "1899 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275291,
            "unit": "B/op",
            "extra": "1899 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5455,
            "unit": "allocs/op",
            "extra": "1899 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 768607,
            "unit": "ns/op\t363756 B/op\t5628 allocs/op",
            "extra": "1700 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 768607,
            "unit": "ns/op",
            "extra": "1700 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 363756,
            "unit": "B/op",
            "extra": "1700 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5628,
            "unit": "allocs/op",
            "extra": "1700 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 624569.5,
            "unit": "ns/op\t275316 B/op\t5440 allocs/op",
            "extra": "1958 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 624569.5,
            "unit": "ns/op",
            "extra": "1958 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 275316,
            "unit": "B/op",
            "extra": "1958 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5440,
            "unit": "allocs/op",
            "extra": "1958 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 625409,
            "unit": "ns/op\t273382 B/op\t5436 allocs/op",
            "extra": "1996 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 625409,
            "unit": "ns/op",
            "extra": "1996 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273382,
            "unit": "B/op",
            "extra": "1996 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5436,
            "unit": "allocs/op",
            "extra": "1996 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 624240,
            "unit": "ns/op\t273405 B/op\t5436 allocs/op",
            "extra": "1998 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 624240,
            "unit": "ns/op",
            "extra": "1998 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 273405,
            "unit": "B/op",
            "extra": "1998 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5436,
            "unit": "allocs/op",
            "extra": "1998 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 629780,
            "unit": "ns/op\t273488 B/op\t5437 allocs/op",
            "extra": "1923 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 629780,
            "unit": "ns/op",
            "extra": "1923 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 273488,
            "unit": "B/op",
            "extra": "1923 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5437,
            "unit": "allocs/op",
            "extra": "1923 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 2211,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "544388 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 2211,
            "unit": "ns/op",
            "extra": "544388 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "544388 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "544388 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 3285.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "378160 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 3285.5,
            "unit": "ns/op",
            "extra": "378160 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "378160 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "378160 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 48.765,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "24392464 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 48.765,
            "unit": "ns/op",
            "extra": "24392464 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "24392464 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "24392464 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 103.1,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "11543174 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 103.1,
            "unit": "ns/op",
            "extra": "11543174 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "11543174 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "11543174 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 26855.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "45055 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 26855.5,
            "unit": "ns/op",
            "extra": "45055 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "45055 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "45055 times\n2 procs"
          }
        ]
      },
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
          "id": "e9f80f8b047b47db4f63012313de0db2c0dc9590",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:54Z",
          "url": "https://github.com/moov-io/watchman/commit/e9f80f8b047b47db4f63012313de0db2c0dc9590"
        },
        "date": 1791992319908,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8134.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "143896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8134.5,
            "unit": "ns/op",
            "extra": "143896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "143896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "143896 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 26304.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "45603 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 26304.5,
            "unit": "ns/op",
            "extra": "45603 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "45603 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "45603 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 12671.5,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "93390 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 12671.5,
            "unit": "ns/op",
            "extra": "93390 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "93390 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "93390 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 28148,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "42878 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 28148,
            "unit": "ns/op",
            "extra": "42878 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "42878 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "42878 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1776.5,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "650515 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1776.5,
            "unit": "ns/op",
            "extra": "650515 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "650515 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "650515 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13192,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "90222 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13192,
            "unit": "ns/op",
            "extra": "90222 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "90222 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "90222 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 57330,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "20982 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 57330,
            "unit": "ns/op",
            "extra": "20982 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "20982 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "20982 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 92408,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "13076 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 92408,
            "unit": "ns/op",
            "extra": "13076 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "13076 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "13076 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20384394.5,
            "unit": "ns/op\t6743479 B/op\t152722 allocs/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20384394.5,
            "unit": "ns/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6743479,
            "unit": "B/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152722,
            "unit": "allocs/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 754099,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1766 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 754099,
            "unit": "ns/op",
            "extra": "1766 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1766 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1766 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 901032.5,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 901032.5,
            "unit": "ns/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 713161,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1812 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 713161,
            "unit": "ns/op",
            "extra": "1812 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1812 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1812 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 870095,
            "unit": "ns/op\t275610 B/op\t5458 allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 870095,
            "unit": "ns/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275610,
            "unit": "B/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5458,
            "unit": "allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 988992.5,
            "unit": "ns/op\t363434 B/op\t5626 allocs/op",
            "extra": "1274 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 988992.5,
            "unit": "ns/op",
            "extra": "1274 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 363434,
            "unit": "B/op",
            "extra": "1274 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5626,
            "unit": "allocs/op",
            "extra": "1274 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 844234,
            "unit": "ns/op\t277037 B/op\t5457 allocs/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 844234,
            "unit": "ns/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 277037,
            "unit": "B/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5457,
            "unit": "allocs/op",
            "extra": "1504 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 831871,
            "unit": "ns/op\t273961 B/op\t5442 allocs/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 831871,
            "unit": "ns/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273961,
            "unit": "B/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5442,
            "unit": "allocs/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 820747.5,
            "unit": "ns/op\t273481 B/op\t5436 allocs/op",
            "extra": "1615 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 820747.5,
            "unit": "ns/op",
            "extra": "1615 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 273481,
            "unit": "B/op",
            "extra": "1615 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5436,
            "unit": "allocs/op",
            "extra": "1615 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 824028,
            "unit": "ns/op\t273659 B/op\t5438 allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 824028,
            "unit": "ns/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 273659,
            "unit": "B/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5438,
            "unit": "allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 2841.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "414979 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 2841.5,
            "unit": "ns/op",
            "extra": "414979 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "414979 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "414979 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4124.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "286892 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4124.5,
            "unit": "ns/op",
            "extra": "286892 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "286892 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "286892 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 62.92,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "18996970 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 62.92,
            "unit": "ns/op",
            "extra": "18996970 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "18996970 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "18996970 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 133.5,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8945858 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 133.5,
            "unit": "ns/op",
            "extra": "8945858 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8945858 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8945858 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 34831.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "34722 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 34831.5,
            "unit": "ns/op",
            "extra": "34722 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "34722 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "34722 times\n2 procs"
          }
        ]
      }
    ]
  }
}