window.BENCHMARK_DATA = {
  "lastUpdate": 1792689024449,
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
        "date": 1791997790877,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7480.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "152596 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7480.5,
            "unit": "ns/op",
            "extra": "152596 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "152596 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "152596 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 24158.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "48945 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 24158.5,
            "unit": "ns/op",
            "extra": "48945 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "48945 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "48945 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11612.5,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "101439 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11612.5,
            "unit": "ns/op",
            "extra": "101439 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "101439 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "101439 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 25752.5,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "46425 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 25752.5,
            "unit": "ns/op",
            "extra": "46425 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "46425 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "46425 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1651,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "700213 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1651,
            "unit": "ns/op",
            "extra": "700213 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "700213 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "700213 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12607.5,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "94338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12607.5,
            "unit": "ns/op",
            "extra": "94338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "94338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "94338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 52822,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "22719 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 52822,
            "unit": "ns/op",
            "extra": "22719 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "22719 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "22719 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 83052,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "14316 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 83052,
            "unit": "ns/op",
            "extra": "14316 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "14316 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "14316 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 19764238,
            "unit": "ns/op\t6742328 B/op\t152713 allocs/op",
            "extra": "61 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 19764238,
            "unit": "ns/op",
            "extra": "61 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6742328,
            "unit": "B/op",
            "extra": "61 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152713,
            "unit": "allocs/op",
            "extra": "61 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 667655.5,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1902 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 667655.5,
            "unit": "ns/op",
            "extra": "1902 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1902 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1902 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 848314.5,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1578 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 848314.5,
            "unit": "ns/op",
            "extra": "1578 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1578 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1578 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 665864,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1947 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 665864,
            "unit": "ns/op",
            "extra": "1947 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1947 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1947 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 821037,
            "unit": "ns/op\t274898 B/op\t5451 allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 821037,
            "unit": "ns/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 274898,
            "unit": "B/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5451,
            "unit": "allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 946368.5,
            "unit": "ns/op\t366219 B/op\t5656 allocs/op",
            "extra": "1322 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 946368.5,
            "unit": "ns/op",
            "extra": "1322 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 366219,
            "unit": "B/op",
            "extra": "1322 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5656,
            "unit": "allocs/op",
            "extra": "1322 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 809046.5,
            "unit": "ns/op\t277094 B/op\t5458 allocs/op",
            "extra": "1540 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 809046.5,
            "unit": "ns/op",
            "extra": "1540 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 277094,
            "unit": "B/op",
            "extra": "1540 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5458,
            "unit": "allocs/op",
            "extra": "1540 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 783747.5,
            "unit": "ns/op\t276262 B/op\t5465 allocs/op",
            "extra": "1598 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 783747.5,
            "unit": "ns/op",
            "extra": "1598 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 276262,
            "unit": "B/op",
            "extra": "1598 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5465,
            "unit": "allocs/op",
            "extra": "1598 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 779700,
            "unit": "ns/op\t275624 B/op\t5458 allocs/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 779700,
            "unit": "ns/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 275624,
            "unit": "B/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5458,
            "unit": "allocs/op",
            "extra": "1534 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 785687.5,
            "unit": "ns/op\t275837 B/op\t5460 allocs/op",
            "extra": "1560 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 785687.5,
            "unit": "ns/op",
            "extra": "1560 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 275837,
            "unit": "B/op",
            "extra": "1560 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5460,
            "unit": "allocs/op",
            "extra": "1560 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 2525.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "452673 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 2525.5,
            "unit": "ns/op",
            "extra": "452673 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "452673 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "452673 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 3815,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "334282 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 3815,
            "unit": "ns/op",
            "extra": "334282 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "334282 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "334282 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 56.225,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "20586148 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 56.225,
            "unit": "ns/op",
            "extra": "20586148 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "20586148 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "20586148 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 119.7,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "9932757 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 119.7,
            "unit": "ns/op",
            "extra": "9932757 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "9932757 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "9932757 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 31570.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "38316 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 31570.5,
            "unit": "ns/op",
            "extra": "38316 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "38316 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "38316 times\n2 procs"
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
          "id": "6df95570dc82728afb7d0327e142b21c86f6c1ed",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.55 (#970)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T05:09:24Z",
          "url": "https://github.com/moov-io/watchman/commit/6df95570dc82728afb7d0327e142b21c86f6c1ed"
        },
        "date": 1792083494570,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 5154.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "227506 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 5154.5,
            "unit": "ns/op",
            "extra": "227506 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "227506 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "227506 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 15845.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "76533 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 15845.5,
            "unit": "ns/op",
            "extra": "76533 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "76533 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "76533 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 8369,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "141850 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 8369,
            "unit": "ns/op",
            "extra": "141850 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "141850 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "141850 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 16602.5,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "72290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 16602.5,
            "unit": "ns/op",
            "extra": "72290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "72290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "72290 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1104,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1104,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 7602,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "157407 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 7602,
            "unit": "ns/op",
            "extra": "157407 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "157407 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "157407 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 39448,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "30060 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 39448,
            "unit": "ns/op",
            "extra": "30060 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "30060 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "30060 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 57491,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "20728 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 57491,
            "unit": "ns/op",
            "extra": "20728 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "20728 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "20728 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 13006245,
            "unit": "ns/op\t6740257 B/op\t152705 allocs/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 13006245,
            "unit": "ns/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740257,
            "unit": "B/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152705,
            "unit": "allocs/op",
            "extra": "97 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 424770.5,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "3136 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 424770.5,
            "unit": "ns/op",
            "extra": "3136 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "3136 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "3136 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 528301,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "2545 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 528301,
            "unit": "ns/op",
            "extra": "2545 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "2545 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "2545 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 421053,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "3151 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 421053,
            "unit": "ns/op",
            "extra": "3151 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "3151 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "3151 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 492444,
            "unit": "ns/op\t275276 B/op\t5454 allocs/op",
            "extra": "2530 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 492444,
            "unit": "ns/op",
            "extra": "2530 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275276,
            "unit": "B/op",
            "extra": "2530 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5454,
            "unit": "allocs/op",
            "extra": "2530 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 549845.5,
            "unit": "ns/op\t367617 B/op\t5669 allocs/op",
            "extra": "2344 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 549845.5,
            "unit": "ns/op",
            "extra": "2344 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 367617,
            "unit": "B/op",
            "extra": "2344 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5669,
            "unit": "allocs/op",
            "extra": "2344 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 491796,
            "unit": "ns/op\t279518 B/op\t5482 allocs/op",
            "extra": "2611 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 491796,
            "unit": "ns/op",
            "extra": "2611 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 279518,
            "unit": "B/op",
            "extra": "2611 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5482,
            "unit": "allocs/op",
            "extra": "2611 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 489566,
            "unit": "ns/op\t273313 B/op\t5435 allocs/op",
            "extra": "2784 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 489566,
            "unit": "ns/op",
            "extra": "2784 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273313,
            "unit": "B/op",
            "extra": "2784 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5435,
            "unit": "allocs/op",
            "extra": "2784 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 489162,
            "unit": "ns/op\t274278 B/op\t5444 allocs/op",
            "extra": "2701 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 489162,
            "unit": "ns/op",
            "extra": "2701 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 274278,
            "unit": "B/op",
            "extra": "2701 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5444,
            "unit": "allocs/op",
            "extra": "2701 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 488835,
            "unit": "ns/op\t276556 B/op\t5468 allocs/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 488835,
            "unit": "ns/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 276556,
            "unit": "B/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5468,
            "unit": "allocs/op",
            "extra": "2601 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 1681,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "699142 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 1681,
            "unit": "ns/op",
            "extra": "699142 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "699142 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "699142 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 2427,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "564127 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 2427,
            "unit": "ns/op",
            "extra": "564127 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "564127 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "564127 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 35.465,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "32810798 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 35.465,
            "unit": "ns/op",
            "extra": "32810798 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "32810798 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32810798 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 75.445,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "15790701 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 75.445,
            "unit": "ns/op",
            "extra": "15790701 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "15790701 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "15790701 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 22420.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "54242 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 22420.5,
            "unit": "ns/op",
            "extra": "54242 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "54242 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "54242 times\n2 procs"
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
          "id": "f263df92f7ce93cc8b44176fb63b4463c60ba241",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.59 (#972)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-16T01:06:59Z",
          "url": "https://github.com/moov-io/watchman/commit/f263df92f7ce93cc8b44176fb63b4463c60ba241"
        },
        "date": 1792168771189,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8239.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "142522 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8239.5,
            "unit": "ns/op",
            "extra": "142522 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "142522 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "142522 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 26885,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "44433 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 26885,
            "unit": "ns/op",
            "extra": "44433 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "44433 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "44433 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 12986.5,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "91921 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 12986.5,
            "unit": "ns/op",
            "extra": "91921 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "91921 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "91921 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 28476,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "41872 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 28476,
            "unit": "ns/op",
            "extra": "41872 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "41872 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "41872 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1654.5,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "718335 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1654.5,
            "unit": "ns/op",
            "extra": "718335 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "718335 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "718335 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13537.5,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "88264 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13537.5,
            "unit": "ns/op",
            "extra": "88264 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "88264 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "88264 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 61405,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "19480 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 61405,
            "unit": "ns/op",
            "extra": "19480 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "19480 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "19480 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 95377,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "12607 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 95377,
            "unit": "ns/op",
            "extra": "12607 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "12607 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "12607 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20137315,
            "unit": "ns/op\t6739226 B/op\t152700 allocs/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20137315,
            "unit": "ns/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6739226,
            "unit": "B/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152700,
            "unit": "allocs/op",
            "extra": "60 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 725654,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1621 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 725654,
            "unit": "ns/op",
            "extra": "1621 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1621 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1621 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 917945,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 917945,
            "unit": "ns/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1395 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 717898.5,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1849 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 717898.5,
            "unit": "ns/op",
            "extra": "1849 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1849 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1849 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 893132.5,
            "unit": "ns/op\t275750 B/op\t5460 allocs/op",
            "extra": "1488 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 893132.5,
            "unit": "ns/op",
            "extra": "1488 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275750,
            "unit": "B/op",
            "extra": "1488 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5460,
            "unit": "allocs/op",
            "extra": "1488 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1050621,
            "unit": "ns/op\t365151 B/op\t5642 allocs/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1050621,
            "unit": "ns/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 365151,
            "unit": "B/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5642,
            "unit": "allocs/op",
            "extra": "1231 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 870625.5,
            "unit": "ns/op\t279390 B/op\t5481 allocs/op",
            "extra": "1519 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 870625.5,
            "unit": "ns/op",
            "extra": "1519 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 279390,
            "unit": "B/op",
            "extra": "1519 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5481,
            "unit": "allocs/op",
            "extra": "1519 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 860681.5,
            "unit": "ns/op\t277298 B/op\t5476 allocs/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 860681.5,
            "unit": "ns/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 277298,
            "unit": "B/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5476,
            "unit": "allocs/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 862258.5,
            "unit": "ns/op\t277108 B/op\t5474 allocs/op",
            "extra": "1453 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 862258.5,
            "unit": "ns/op",
            "extra": "1453 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 277108,
            "unit": "B/op",
            "extra": "1453 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5474,
            "unit": "allocs/op",
            "extra": "1453 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 872722.5,
            "unit": "ns/op\t276578 B/op\t5468 allocs/op",
            "extra": "1430 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 872722.5,
            "unit": "ns/op",
            "extra": "1430 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 276578,
            "unit": "B/op",
            "extra": "1430 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5468,
            "unit": "allocs/op",
            "extra": "1430 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3396.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "352524 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3396.5,
            "unit": "ns/op",
            "extra": "352524 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "352524 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "352524 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4780.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "247156 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4780.5,
            "unit": "ns/op",
            "extra": "247156 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "247156 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "247156 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 67.265,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "17648908 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 67.265,
            "unit": "ns/op",
            "extra": "17648908 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "17648908 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17648908 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 145.6,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8281557 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 145.6,
            "unit": "ns/op",
            "extra": "8281557 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8281557 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8281557 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35289.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "34074 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35289.5,
            "unit": "ns/op",
            "extra": "34074 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "34074 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "34074 times\n2 procs"
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
          "id": "b90c7739146cbdf882964ca3cc9707a189efdaac",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.61 (#973)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-17T01:17:10Z",
          "url": "https://github.com/moov-io/watchman/commit/b90c7739146cbdf882964ca3cc9707a189efdaac"
        },
        "date": 1792251261851,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 4829,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "242160 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 4829,
            "unit": "ns/op",
            "extra": "242160 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "242160 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "242160 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 14827,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "79826 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 14827,
            "unit": "ns/op",
            "extra": "79826 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "79826 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "79826 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 7946,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "148662 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 7946,
            "unit": "ns/op",
            "extra": "148662 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "148662 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "148662 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 16421.5,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "72376 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 16421.5,
            "unit": "ns/op",
            "extra": "72376 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "72376 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "72376 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1047,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1047,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 7308.5,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "161410 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 7308.5,
            "unit": "ns/op",
            "extra": "161410 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "161410 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "161410 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 39081.5,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "30710 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 39081.5,
            "unit": "ns/op",
            "extra": "30710 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "30710 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "30710 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 59268.5,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "20306 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 59268.5,
            "unit": "ns/op",
            "extra": "20306 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28933,
            "unit": "B/op",
            "extra": "20306 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 338,
            "unit": "allocs/op",
            "extra": "20306 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 13756464,
            "unit": "ns/op\t6741714 B/op\t152712 allocs/op",
            "extra": "87 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 13756464,
            "unit": "ns/op",
            "extra": "87 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6741714,
            "unit": "B/op",
            "extra": "87 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152712,
            "unit": "allocs/op",
            "extra": "87 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 427416,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "2934 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 427416,
            "unit": "ns/op",
            "extra": "2934 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "2934 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "2934 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 523815,
            "unit": "ns/op\t341039 B/op\t5330 allocs/op",
            "extra": "2463 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 523815,
            "unit": "ns/op",
            "extra": "2463 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341039,
            "unit": "B/op",
            "extra": "2463 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "2463 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 423365,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "3202 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 423365,
            "unit": "ns/op",
            "extra": "3202 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "3202 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "3202 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 500471,
            "unit": "ns/op\t275122 B/op\t5454 allocs/op",
            "extra": "2666 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 500471,
            "unit": "ns/op",
            "extra": "2666 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275122,
            "unit": "B/op",
            "extra": "2666 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5454,
            "unit": "allocs/op",
            "extra": "2666 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 559499.5,
            "unit": "ns/op\t367208 B/op\t5666 allocs/op",
            "extra": "2300 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 559499.5,
            "unit": "ns/op",
            "extra": "2300 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 367208,
            "unit": "B/op",
            "extra": "2300 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5666,
            "unit": "allocs/op",
            "extra": "2300 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 488021.5,
            "unit": "ns/op\t277504 B/op\t5462 allocs/op",
            "extra": "2656 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 488021.5,
            "unit": "ns/op",
            "extra": "2656 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 277504,
            "unit": "B/op",
            "extra": "2656 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5462,
            "unit": "allocs/op",
            "extra": "2656 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 481131.5,
            "unit": "ns/op\t275878 B/op\t5462 allocs/op",
            "extra": "2816 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 481131.5,
            "unit": "ns/op",
            "extra": "2816 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 275878,
            "unit": "B/op",
            "extra": "2816 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5462,
            "unit": "allocs/op",
            "extra": "2816 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 486807,
            "unit": "ns/op\t275228 B/op\t5454 allocs/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 486807,
            "unit": "ns/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 275228,
            "unit": "B/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5454,
            "unit": "allocs/op",
            "extra": "2815 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 473974,
            "unit": "ns/op\t276119 B/op\t5464 allocs/op",
            "extra": "2797 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 473974,
            "unit": "ns/op",
            "extra": "2797 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 276119,
            "unit": "B/op",
            "extra": "2797 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5464,
            "unit": "allocs/op",
            "extra": "2797 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 1781.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "700622 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 1781.5,
            "unit": "ns/op",
            "extra": "700622 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "700622 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "700622 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 2607,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "433658 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 2607,
            "unit": "ns/op",
            "extra": "433658 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "433658 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "433658 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 36.615,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "32411034 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 36.615,
            "unit": "ns/op",
            "extra": "32411034 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "32411034 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32411034 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 76.9,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "15493678 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 76.9,
            "unit": "ns/op",
            "extra": "15493678 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "15493678 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "15493678 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 22417,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "54285 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 22417,
            "unit": "ns/op",
            "extra": "54285 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "54285 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "54285 times\n2 procs"
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
          "id": "1d31b0054f6f879bb2ae06e195d93aba80bed354",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.63 (#975)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-18T05:31:29Z",
          "url": "https://github.com/moov-io/watchman/commit/1d31b0054f6f879bb2ae06e195d93aba80bed354"
        },
        "date": 1792338844897,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8691.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "134830 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8691.5,
            "unit": "ns/op",
            "extra": "134830 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "134830 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "134830 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 29365,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "40918 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 29365,
            "unit": "ns/op",
            "extra": "40918 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "40918 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "40918 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13730,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "86703 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13730,
            "unit": "ns/op",
            "extra": "86703 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "86703 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "86703 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30993,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "38559 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30993,
            "unit": "ns/op",
            "extra": "38559 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "38559 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "38559 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1790,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "655963 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1790,
            "unit": "ns/op",
            "extra": "655963 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "655963 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "655963 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14858,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "79836 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14858,
            "unit": "ns/op",
            "extra": "79836 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "79836 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "79836 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 64525,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "18429 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 64525,
            "unit": "ns/op",
            "extra": "18429 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "18429 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18429 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 103622.5,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 103622.5,
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
            "value": 21396256,
            "unit": "ns/op\t6740370 B/op\t152706 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21396256,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740370,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152706,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 724447.5,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1678 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 724447.5,
            "unit": "ns/op",
            "extra": "1678 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1678 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1678 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 916987.5,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1386 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 916987.5,
            "unit": "ns/op",
            "extra": "1386 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1386 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1386 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 733588,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1723 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 733588,
            "unit": "ns/op",
            "extra": "1723 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1723 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1723 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 874158,
            "unit": "ns/op\t276324 B/op\t5466 allocs/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 874158,
            "unit": "ns/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 276324,
            "unit": "B/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5466,
            "unit": "allocs/op",
            "extra": "1459 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1033478.5,
            "unit": "ns/op\t363796 B/op\t5631 allocs/op",
            "extra": "1213 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1033478.5,
            "unit": "ns/op",
            "extra": "1213 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 363796,
            "unit": "B/op",
            "extra": "1213 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5631,
            "unit": "allocs/op",
            "extra": "1213 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 838245.5,
            "unit": "ns/op\t275341 B/op\t5440 allocs/op",
            "extra": "1544 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 838245.5,
            "unit": "ns/op",
            "extra": "1544 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 275341,
            "unit": "B/op",
            "extra": "1544 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5440,
            "unit": "allocs/op",
            "extra": "1544 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 841897.5,
            "unit": "ns/op\t273306 B/op\t5435 allocs/op",
            "extra": "1506 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 841897.5,
            "unit": "ns/op",
            "extra": "1506 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273306,
            "unit": "B/op",
            "extra": "1506 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5435,
            "unit": "allocs/op",
            "extra": "1506 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 834544.5,
            "unit": "ns/op\t273296 B/op\t5435 allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 834544.5,
            "unit": "ns/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 273296,
            "unit": "B/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5435,
            "unit": "allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 843594.5,
            "unit": "ns/op\t273329 B/op\t5435 allocs/op",
            "extra": "1480 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 843594.5,
            "unit": "ns/op",
            "extra": "1480 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 273329,
            "unit": "B/op",
            "extra": "1480 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5435,
            "unit": "allocs/op",
            "extra": "1480 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4333.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "280449 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4333.5,
            "unit": "ns/op",
            "extra": "280449 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "280449 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "280449 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5923.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "211402 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5923.5,
            "unit": "ns/op",
            "extra": "211402 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "211402 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "211402 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 70.34,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "16350357 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 70.34,
            "unit": "ns/op",
            "extra": "16350357 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16350357 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16350357 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.45,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8150318 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.45,
            "unit": "ns/op",
            "extra": "8150318 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8150318 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8150318 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37370.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "32238 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37370.5,
            "unit": "ns/op",
            "extra": "32238 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32238 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32238 times\n2 procs"
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
          "id": "44ad7b13d148afbac8d92019ec6e42ecfbff4dfe",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.64 (#976)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-19T01:37:25Z",
          "url": "https://github.com/moov-io/watchman/commit/44ad7b13d148afbac8d92019ec6e42ecfbff4dfe"
        },
        "date": 1792435676813,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8864,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "132927 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8864,
            "unit": "ns/op",
            "extra": "132927 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "132927 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "132927 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30118.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "39489 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30118.5,
            "unit": "ns/op",
            "extra": "39489 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "39489 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "39489 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13866,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "84424 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13866,
            "unit": "ns/op",
            "extra": "84424 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "84424 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "84424 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30670,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "39207 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30670,
            "unit": "ns/op",
            "extra": "39207 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "39207 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "39207 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1791,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "645244 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1791,
            "unit": "ns/op",
            "extra": "645244 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "645244 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "645244 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14452,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "82136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14452,
            "unit": "ns/op",
            "extra": "82136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "82136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "82136 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 64177,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "18614 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 64177,
            "unit": "ns/op",
            "extra": "18614 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "18614 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18614 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 102261.5,
            "unit": "ns/op\t28934 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 102261.5,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28934,
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
            "value": 20505691.5,
            "unit": "ns/op\t6740549 B/op\t152706 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20505691.5,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740549,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152706,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 734836,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 734836,
            "unit": "ns/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1639 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 922590,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1368 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 922590,
            "unit": "ns/op",
            "extra": "1368 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1368 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1368 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 741553.5,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1754 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 741553.5,
            "unit": "ns/op",
            "extra": "1754 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1754 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1754 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 866719,
            "unit": "ns/op\t273417 B/op\t5436 allocs/op",
            "extra": "1456 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 866719,
            "unit": "ns/op",
            "extra": "1456 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 273417,
            "unit": "B/op",
            "extra": "1456 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5436,
            "unit": "allocs/op",
            "extra": "1456 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1060368,
            "unit": "ns/op\t363464 B/op\t5626 allocs/op",
            "extra": "1192 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1060368,
            "unit": "ns/op",
            "extra": "1192 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 363464,
            "unit": "B/op",
            "extra": "1192 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5626,
            "unit": "allocs/op",
            "extra": "1192 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 871996.5,
            "unit": "ns/op\t278694 B/op\t5474 allocs/op",
            "extra": "1414 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 871996.5,
            "unit": "ns/op",
            "extra": "1414 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 278694,
            "unit": "B/op",
            "extra": "1414 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5474,
            "unit": "allocs/op",
            "extra": "1414 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 859641.5,
            "unit": "ns/op\t273547 B/op\t5438 allocs/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 859641.5,
            "unit": "ns/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273547,
            "unit": "B/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5438,
            "unit": "allocs/op",
            "extra": "1473 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 861095.5,
            "unit": "ns/op\t273504 B/op\t5437 allocs/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 861095.5,
            "unit": "ns/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 273504,
            "unit": "B/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5437,
            "unit": "allocs/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 869972,
            "unit": "ns/op\t275898 B/op\t5461 allocs/op",
            "extra": "1384 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 869972,
            "unit": "ns/op",
            "extra": "1384 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 275898,
            "unit": "B/op",
            "extra": "1384 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5461,
            "unit": "allocs/op",
            "extra": "1384 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4373.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "303423 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4373.5,
            "unit": "ns/op",
            "extra": "303423 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "303423 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "303423 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5742.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "204002 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5742.5,
            "unit": "ns/op",
            "extra": "204002 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "204002 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "204002 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 69.255,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "16778127 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 69.255,
            "unit": "ns/op",
            "extra": "16778127 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16778127 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16778127 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.55,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8226420 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.55,
            "unit": "ns/op",
            "extra": "8226420 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8226420 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8226420 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36660.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "32844 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36660.5,
            "unit": "ns/op",
            "extra": "32844 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32844 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32844 times\n2 procs"
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
          "id": "bea69b6c311e7c74898daa668a571cd96fb2e38a",
          "message": "chore(deps): update go toolchain directive to v1.27.3 (#977)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-20T14:03:08Z",
          "url": "https://github.com/moov-io/watchman/commit/bea69b6c311e7c74898daa668a571cd96fb2e38a"
        },
        "date": 1792515336801,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8716.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "134305 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8716.5,
            "unit": "ns/op",
            "extra": "134305 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "134305 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "134305 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 28888,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "41304 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 28888,
            "unit": "ns/op",
            "extra": "41304 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "41304 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "41304 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13762,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "86082 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13762,
            "unit": "ns/op",
            "extra": "86082 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "86082 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "86082 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30658.5,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "38542 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30658.5,
            "unit": "ns/op",
            "extra": "38542 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "38542 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "38542 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1875,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "635227 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1875,
            "unit": "ns/op",
            "extra": "635227 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "635227 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "635227 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 15818.5,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "76214 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 15818.5,
            "unit": "ns/op",
            "extra": "76214 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "76214 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "76214 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 63882,
            "unit": "ns/op\t10009 B/op\t205 allocs/op",
            "extra": "18844 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 63882,
            "unit": "ns/op",
            "extra": "18844 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10009,
            "unit": "B/op",
            "extra": "18844 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18844 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 102339.5,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 102339.5,
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
            "value": 20783629,
            "unit": "ns/op\t6740203 B/op\t152704 allocs/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20783629,
            "unit": "ns/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740203,
            "unit": "B/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152704,
            "unit": "allocs/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 760311.5,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1688 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 760311.5,
            "unit": "ns/op",
            "extra": "1688 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1688 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1688 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 937406,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1323 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 937406,
            "unit": "ns/op",
            "extra": "1323 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1323 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1323 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 743527.5,
            "unit": "ns/op\t258753 B/op\t5002 allocs/op",
            "extra": "1677 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 743527.5,
            "unit": "ns/op",
            "extra": "1677 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258753,
            "unit": "B/op",
            "extra": "1677 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1677 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 863377.5,
            "unit": "ns/op\t275040 B/op\t5452 allocs/op",
            "extra": "1435 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 863377.5,
            "unit": "ns/op",
            "extra": "1435 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275040,
            "unit": "B/op",
            "extra": "1435 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5452,
            "unit": "allocs/op",
            "extra": "1435 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1064383,
            "unit": "ns/op\t364448 B/op\t5636 allocs/op",
            "extra": "1113 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1064383,
            "unit": "ns/op",
            "extra": "1113 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 364448,
            "unit": "B/op",
            "extra": "1113 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5636,
            "unit": "allocs/op",
            "extra": "1113 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 857565,
            "unit": "ns/op\t275433 B/op\t5440 allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 857565,
            "unit": "ns/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 275433,
            "unit": "B/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5440,
            "unit": "allocs/op",
            "extra": "1551 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 858723,
            "unit": "ns/op\t273444 B/op\t5436 allocs/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 858723,
            "unit": "ns/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 273444,
            "unit": "B/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5436,
            "unit": "allocs/op",
            "extra": "1455 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 867870,
            "unit": "ns/op\t273512 B/op\t5437 allocs/op",
            "extra": "1431 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 867870,
            "unit": "ns/op",
            "extra": "1431 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 273512,
            "unit": "B/op",
            "extra": "1431 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5437,
            "unit": "allocs/op",
            "extra": "1431 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 870299.5,
            "unit": "ns/op\t274952 B/op\t5452 allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 870299.5,
            "unit": "ns/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 274952,
            "unit": "B/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5452,
            "unit": "allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4170.5,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "302864 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4170.5,
            "unit": "ns/op",
            "extra": "302864 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "302864 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "302864 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5840,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "224494 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5840,
            "unit": "ns/op",
            "extra": "224494 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "224494 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "224494 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 69.285,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "15564973 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 69.285,
            "unit": "ns/op",
            "extra": "15564973 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "15564973 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "15564973 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.6,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8115614 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.6,
            "unit": "ns/op",
            "extra": "8115614 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8115614 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8115614 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36540.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "32900 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36540.5,
            "unit": "ns/op",
            "extra": "32900 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32900 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32900 times\n2 procs"
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
          "id": "1c10ad1c075bf745f4f2ed97dcc3f3bf89de5d19",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.66 (#979)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T05:43:22Z",
          "url": "https://github.com/moov-io/watchman/commit/1c10ad1c075bf745f4f2ed97dcc3f3bf89de5d19"
        },
        "date": 1792602699077,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8627.5,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "136059 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8627.5,
            "unit": "ns/op",
            "extra": "136059 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "136059 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "136059 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 28631.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "41529 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 28631.5,
            "unit": "ns/op",
            "extra": "41529 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "41529 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "41529 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13828,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "86161 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13828,
            "unit": "ns/op",
            "extra": "86161 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "86161 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "86161 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30738,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "38884 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30738,
            "unit": "ns/op",
            "extra": "38884 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "38884 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "38884 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1800.5,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "644270 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1800.5,
            "unit": "ns/op",
            "extra": "644270 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "644270 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "644270 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14595,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "81338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14595,
            "unit": "ns/op",
            "extra": "81338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "81338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "81338 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 63451,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "18913 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 63451,
            "unit": "ns/op",
            "extra": "18913 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "18913 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18913 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 100599,
            "unit": "ns/op\t28933 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 100599,
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
            "value": 20937839,
            "unit": "ns/op\t6740538 B/op\t152706 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20937839,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740538,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152706,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 734192,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1772 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 734192,
            "unit": "ns/op",
            "extra": "1772 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1772 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1772 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 922511,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 922511,
            "unit": "ns/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1424 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 739224,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1776 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 739224,
            "unit": "ns/op",
            "extra": "1776 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1776 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1776 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 869785.5,
            "unit": "ns/op\t275976 B/op\t5462 allocs/op",
            "extra": "1441 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 869785.5,
            "unit": "ns/op",
            "extra": "1441 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 275976,
            "unit": "B/op",
            "extra": "1441 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5462,
            "unit": "allocs/op",
            "extra": "1441 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1035000.5,
            "unit": "ns/op\t366730 B/op\t5660 allocs/op",
            "extra": "1218 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1035000.5,
            "unit": "ns/op",
            "extra": "1218 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 366730,
            "unit": "B/op",
            "extra": "1218 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5660,
            "unit": "allocs/op",
            "extra": "1218 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 875408.5,
            "unit": "ns/op\t278128 B/op\t5468 allocs/op",
            "extra": "1474 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 875408.5,
            "unit": "ns/op",
            "extra": "1474 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 278128,
            "unit": "B/op",
            "extra": "1474 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5468,
            "unit": "allocs/op",
            "extra": "1474 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 861914.5,
            "unit": "ns/op\t275594 B/op\t5458 allocs/op",
            "extra": "1437 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 861914.5,
            "unit": "ns/op",
            "extra": "1437 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 275594,
            "unit": "B/op",
            "extra": "1437 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5458,
            "unit": "allocs/op",
            "extra": "1437 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 857218,
            "unit": "ns/op\t277780 B/op\t5480 allocs/op",
            "extra": "1538 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 857218,
            "unit": "ns/op",
            "extra": "1538 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 277780,
            "unit": "B/op",
            "extra": "1538 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5480,
            "unit": "allocs/op",
            "extra": "1538 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 846349,
            "unit": "ns/op\t277980 B/op\t5482 allocs/op",
            "extra": "1558 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 846349,
            "unit": "ns/op",
            "extra": "1558 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 277980,
            "unit": "B/op",
            "extra": "1558 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5482,
            "unit": "allocs/op",
            "extra": "1558 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4335,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "307620 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4335,
            "unit": "ns/op",
            "extra": "307620 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "307620 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "307620 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5776.5,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "207547 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5776.5,
            "unit": "ns/op",
            "extra": "207547 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "207547 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "207547 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 69.39,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "16653424 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 69.39,
            "unit": "ns/op",
            "extra": "16653424 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16653424 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16653424 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.5,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8232630 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.5,
            "unit": "ns/op",
            "extra": "8232630 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8232630 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8232630 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36807,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "32655 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36807,
            "unit": "ns/op",
            "extra": "32655 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32655 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32655 times\n2 procs"
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
          "id": "3c9bd53ea466c587b261b307137d1d8ed897457c",
          "message": "fix(deps): update module github.com/openrouterteam/go-sdk to v0.9.69 (#981)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T23:29:51Z",
          "url": "https://github.com/moov-io/watchman/commit/3c9bd53ea466c587b261b307137d1d8ed897457c"
        },
        "date": 1792689023276,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8529,
            "unit": "ns/op\t2112 B/op\t48 allocs/op",
            "extra": "138189 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8529,
            "unit": "ns/op",
            "extra": "138189 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2112,
            "unit": "B/op",
            "extra": "138189 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 48,
            "unit": "allocs/op",
            "extra": "138189 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 28656.5,
            "unit": "ns/op\t11817 B/op\t100 allocs/op",
            "extra": "41587 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 28656.5,
            "unit": "ns/op",
            "extra": "41587 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11817,
            "unit": "B/op",
            "extra": "41587 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 100,
            "unit": "allocs/op",
            "extra": "41587 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 13682.5,
            "unit": "ns/op\t1614 B/op\t35 allocs/op",
            "extra": "86835 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 13682.5,
            "unit": "ns/op",
            "extra": "86835 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1614,
            "unit": "B/op",
            "extra": "86835 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 35,
            "unit": "allocs/op",
            "extra": "86835 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30379.5,
            "unit": "ns/op\t10095 B/op\t71 allocs/op",
            "extra": "39400 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30379.5,
            "unit": "ns/op",
            "extra": "39400 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10095,
            "unit": "B/op",
            "extra": "39400 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 71,
            "unit": "allocs/op",
            "extra": "39400 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1781,
            "unit": "ns/op\t320 B/op\t9 allocs/op",
            "extra": "651793 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1781,
            "unit": "ns/op",
            "extra": "651793 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 320,
            "unit": "B/op",
            "extra": "651793 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "651793 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14555,
            "unit": "ns/op\t5528 B/op\t28 allocs/op",
            "extra": "81789 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14555,
            "unit": "ns/op",
            "extra": "81789 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5528,
            "unit": "B/op",
            "extra": "81789 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "81789 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 64335,
            "unit": "ns/op\t10008 B/op\t205 allocs/op",
            "extra": "18640 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 64335,
            "unit": "ns/op",
            "extra": "18640 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 10008,
            "unit": "B/op",
            "extra": "18640 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 205,
            "unit": "allocs/op",
            "extra": "18640 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 103480.5,
            "unit": "ns/op\t28934 B/op\t338 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 103480.5,
            "unit": "ns/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 28934,
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
            "value": 20170636,
            "unit": "ns/op\t6740278 B/op\t152705 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20170636,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 6740278,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 152705,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off",
            "value": 728173,
            "unit": "ns/op\t256914 B/op\t5000 allocs/op",
            "extra": "1652 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - ns/op",
            "value": 728173,
            "unit": "ns/op",
            "extra": "1652 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - B/op",
            "value": 256914,
            "unit": "B/op",
            "extra": "1652 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/off - allocs/op",
            "value": 5000,
            "unit": "allocs/op",
            "extra": "1652 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned",
            "value": 913373.5,
            "unit": "ns/op\t341038 B/op\t5330 allocs/op",
            "extra": "1382 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - ns/op",
            "value": 913373.5,
            "unit": "ns/op",
            "extra": "1382 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - B/op",
            "value": 341038,
            "unit": "B/op",
            "extra": "1382 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_all_returned - allocs/op",
            "value": 5330,
            "unit": "allocs/op",
            "extra": "1382 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80",
            "value": 730050.5,
            "unit": "ns/op\t258754 B/op\t5002 allocs/op",
            "extra": "1822 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - ns/op",
            "value": 730050.5,
            "unit": "ns/op",
            "extra": "1822 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - B/op",
            "value": 258754,
            "unit": "B/op",
            "extra": "1822 times\n2 procs"
          },
          {
            "name": "Benchmark_SearchDebug/debug_threshold_0.80 - allocs/op",
            "value": 5002,
            "unit": "allocs/op",
            "extra": "1822 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 876577,
            "unit": "ns/op\t276581 B/op\t5468 allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 876577,
            "unit": "ns/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 276581,
            "unit": "B/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 5468,
            "unit": "allocs/op",
            "extra": "1419 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 1012039.5,
            "unit": "ns/op\t366580 B/op\t5658 allocs/op",
            "extra": "1287 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 1012039.5,
            "unit": "ns/op",
            "extra": "1287 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 366580,
            "unit": "B/op",
            "extra": "1287 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 5658,
            "unit": "allocs/op",
            "extra": "1287 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold",
            "value": 866947.5,
            "unit": "ns/op\t277956 B/op\t5466 allocs/op",
            "extra": "1406 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - ns/op",
            "value": 866947.5,
            "unit": "ns/op",
            "extra": "1406 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - B/op",
            "value": 277956,
            "unit": "B/op",
            "extra": "1406 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug-threshold - allocs/op",
            "value": 5466,
            "unit": "allocs/op",
            "extra": "1406 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 849006.5,
            "unit": "ns/op\t276368 B/op\t5466 allocs/op",
            "extra": "1417 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 849006.5,
            "unit": "ns/op",
            "extra": "1417 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 276368,
            "unit": "B/op",
            "extra": "1417 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 5466,
            "unit": "allocs/op",
            "extra": "1417 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 841619.5,
            "unit": "ns/op\t275122 B/op\t5453 allocs/op",
            "extra": "1508 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 841619.5,
            "unit": "ns/op",
            "extra": "1508 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 275122,
            "unit": "B/op",
            "extra": "1508 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 5453,
            "unit": "allocs/op",
            "extra": "1508 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 840306.5,
            "unit": "ns/op\t275106 B/op\t5453 allocs/op",
            "extra": "1567 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 840306.5,
            "unit": "ns/op",
            "extra": "1567 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 275106,
            "unit": "B/op",
            "extra": "1567 times\n2 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 5453,
            "unit": "allocs/op",
            "extra": "1567 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4426,
            "unit": "ns/op\t50 B/op\t1 allocs/op",
            "extra": "269522 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4426,
            "unit": "ns/op",
            "extra": "269522 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 50,
            "unit": "B/op",
            "extra": "269522 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "269522 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5834,
            "unit": "ns/op\t351 B/op\t6 allocs/op",
            "extra": "201717 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5834,
            "unit": "ns/op",
            "extra": "201717 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 351,
            "unit": "B/op",
            "extra": "201717 times\n2 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "201717 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex",
            "value": 69.75,
            "unit": "ns/op\t4 B/op\t1 allocs/op",
            "extra": "16252119 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - ns/op",
            "value": 69.75,
            "unit": "ns/op",
            "extra": "16252119 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - B/op",
            "value": 4,
            "unit": "B/op",
            "extra": "16252119 times\n2 procs"
          },
          {
            "name": "BenchmarkEncodeSoundex - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "16252119 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch",
            "value": 147.5,
            "unit": "ns/op\t8 B/op\t2 allocs/op",
            "extra": "8118660 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - ns/op",
            "value": 147.5,
            "unit": "ns/op",
            "extra": "8118660 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - B/op",
            "value": 8,
            "unit": "B/op",
            "extra": "8118660 times\n2 procs"
          },
          {
            "name": "BenchmarkSoundexMatch - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "8118660 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36728.5,
            "unit": "ns/op\t16 B/op\t1 allocs/op",
            "extra": "32706 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36728.5,
            "unit": "ns/op",
            "extra": "32706 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32706 times\n2 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32706 times\n2 procs"
          }
        ]
      }
    ]
  }
}