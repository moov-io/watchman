window.BENCHMARK_DATA = {
  "lastUpdate": 1750251631332,
  "repoUrl": "https://github.com/moov-io/watchman",
  "entries": {
    "moov-io/watchman Common Benchmarks": [
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
          "id": "514486764e9f15f6edbc228d71caf9a15beb3a78",
          "message": "Merge pull request #610 from moov-io/feat-ingest-csv\n\ningest: support arbitrary CSV via configuration",
          "timestamp": "2025-05-07T16:27:26Z",
          "url": "https://github.com/moov-io/watchman/commit/514486764e9f15f6edbc228d71caf9a15beb3a78"
        },
        "date": 1746639573319,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7985,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "142452 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7985,
            "unit": "ns/op",
            "extra": "142452 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "142452 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "142452 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32889,
            "unit": "ns/op\t   11285 B/op\t     175 allocs/op",
            "extra": "36333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32889,
            "unit": "ns/op",
            "extra": "36333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11285,
            "unit": "B/op",
            "extra": "36333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "36333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11338,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "102259 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11338,
            "unit": "ns/op",
            "extra": "102259 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "102259 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "102259 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30742,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30742,
            "unit": "ns/op",
            "extra": "39032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1434,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "816625 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1434,
            "unit": "ns/op",
            "extra": "816625 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "816625 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "816625 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 15054,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "79119 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 15054,
            "unit": "ns/op",
            "extra": "79119 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "79119 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "79119 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44933,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44933,
            "unit": "ns/op",
            "extra": "26305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 95937,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12606 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 95937,
            "unit": "ns/op",
            "extra": "12606 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12606 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12606 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 61627922,
            "unit": "ns/op\t32932281 B/op\t 1121480 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 61627922,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932281,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121480,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29787464,
            "unit": "ns/op\t19172542 B/op\t  462241 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29787464,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19172542,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462241,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 50943429,
            "unit": "ns/op\t50900635 B/op\t  526723 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 50943429,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50900635,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526723,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30290710,
            "unit": "ns/op\t19171961 B/op\t  462214 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30290710,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19171961,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462214,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 29673917,
            "unit": "ns/op\t19166513 B/op\t  462214 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 29673917,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166513,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462214,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29387081,
            "unit": "ns/op\t19164571 B/op\t  462214 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29387081,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19164571,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462214,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3495,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "396621 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3495,
            "unit": "ns/op",
            "extra": "396621 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "396621 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "396621 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4870,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "253551 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4870,
            "unit": "ns/op",
            "extra": "253551 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "253551 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "253551 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37186,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32311 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37186,
            "unit": "ns/op",
            "extra": "32311 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32311 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32311 times\n4 procs"
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
          "id": "dc43f05e5f207e6ee1c534b8fdb79f1f4061c0c3",
          "message": "Merge pull request #627 from ZSuiteTech/docs/readme_update_ofac_url\n\ndocs: correct OFAC_DOWNLOAD_TEMPLATE default value",
          "timestamp": "2025-05-14T14:33:29Z",
          "url": "https://github.com/moov-io/watchman/commit/dc43f05e5f207e6ee1c534b8fdb79f1f4061c0c3"
        },
        "date": 1747313809019,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8054,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "143341 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8054,
            "unit": "ns/op",
            "extra": "143341 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "143341 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "143341 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33591,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "34785 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33591,
            "unit": "ns/op",
            "extra": "34785 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "34785 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "34785 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11341,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "100504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11341,
            "unit": "ns/op",
            "extra": "100504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "100504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "100504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30714,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30714,
            "unit": "ns/op",
            "extra": "39141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1413,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "805195 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1413,
            "unit": "ns/op",
            "extra": "805195 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "805195 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "805195 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 15103,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "80210 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 15103,
            "unit": "ns/op",
            "extra": "80210 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "80210 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "80210 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 45065,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26343 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 45065,
            "unit": "ns/op",
            "extra": "26343 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26343 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26343 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 97471,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12326 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 97471,
            "unit": "ns/op",
            "extra": "12326 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12326 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12326 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 54368589,
            "unit": "ns/op\t32931390 B/op\t 1121424 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 54368589,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931390,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121424,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 34465352,
            "unit": "ns/op\t19168826 B/op\t  462253 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 34465352,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19168826,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462253,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 52792918,
            "unit": "ns/op\t50893700 B/op\t  526739 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 52792918,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50893700,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526739,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30909936,
            "unit": "ns/op\t19164142 B/op\t  462230 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30909936,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19164142,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462230,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31356884,
            "unit": "ns/op\t19163196 B/op\t  462230 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31356884,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19163196,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462230,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30938180,
            "unit": "ns/op\t19170215 B/op\t  462231 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30938180,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19170215,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462231,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3428,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "302390 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3428,
            "unit": "ns/op",
            "extra": "302390 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "302390 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "302390 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4779,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "229405 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4779,
            "unit": "ns/op",
            "extra": "229405 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "229405 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "229405 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37713,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31911 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37713,
            "unit": "ns/op",
            "extra": "31911 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31911 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31911 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "dfdf22ef0400ab0c8c607258c3a0e6d084b433d7",
          "message": "search: grab read lock before ProcessSliceFn",
          "timestamp": "2025-05-16T18:03:40Z",
          "url": "https://github.com/moov-io/watchman/commit/dfdf22ef0400ab0c8c607258c3a0e6d084b433d7"
        },
        "date": 1747427163125,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9930,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "114330 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9930,
            "unit": "ns/op",
            "extra": "114330 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "114330 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "114330 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34234,
            "unit": "ns/op\t   12702 B/op\t     174 allocs/op",
            "extra": "35290 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34234,
            "unit": "ns/op",
            "extra": "35290 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12702,
            "unit": "B/op",
            "extra": "35290 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "35290 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 17001,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "69637 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 17001,
            "unit": "ns/op",
            "extra": "69637 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "69637 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69637 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36608,
            "unit": "ns/op\t   11310 B/op\t     210 allocs/op",
            "extra": "33036 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36608,
            "unit": "ns/op",
            "extra": "33036 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33036 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "33036 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1433,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "813716 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1433,
            "unit": "ns/op",
            "extra": "813716 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "813716 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "813716 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14948,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14948,
            "unit": "ns/op",
            "extra": "81378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 86607,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13808 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 86607,
            "unit": "ns/op",
            "extra": "13808 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13808 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13808 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 139503,
            "unit": "ns/op\t   41102 B/op\t     999 allocs/op",
            "extra": "8485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 139503,
            "unit": "ns/op",
            "extra": "8485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41102,
            "unit": "B/op",
            "extra": "8485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8485 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53786728,
            "unit": "ns/op\t32933091 B/op\t 1121493 allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53786728,
            "unit": "ns/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933091,
            "unit": "B/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121493,
            "unit": "allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29396693,
            "unit": "ns/op\t19173408 B/op\t  462245 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29396693,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19173408,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462245,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 53091494,
            "unit": "ns/op\t50889166 B/op\t  526720 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 53091494,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50889166,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526720,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 33645876,
            "unit": "ns/op\t19173558 B/op\t  462220 allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 33645876,
            "unit": "ns/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19173558,
            "unit": "B/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462220,
            "unit": "allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 29412770,
            "unit": "ns/op\t19167330 B/op\t  462224 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 29412770,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19167330,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462224,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29639515,
            "unit": "ns/op\t19170865 B/op\t  462227 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29639515,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19170865,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462227,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3358,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "336049 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3358,
            "unit": "ns/op",
            "extra": "336049 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "336049 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "336049 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4785,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "234644 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4785,
            "unit": "ns/op",
            "extra": "234644 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "234644 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "234644 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37556,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31898 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37556,
            "unit": "ns/op",
            "extra": "31898 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31898 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31898 times\n4 procs"
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
          "id": "0322b24dc5663290a43628791d4af8576334e884",
          "message": "Merge pull request #630 from adamdecaf/fix-ingest-file\n\ningest: files add entities to the in-memory index",
          "timestamp": "2025-05-16T20:35:50Z",
          "url": "https://github.com/moov-io/watchman/commit/0322b24dc5663290a43628791d4af8576334e884"
        },
        "date": 1747427975791,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10168,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "112023 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10168,
            "unit": "ns/op",
            "extra": "112023 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "112023 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "112023 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35571,
            "unit": "ns/op\t   12703 B/op\t     174 allocs/op",
            "extra": "34318 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35571,
            "unit": "ns/op",
            "extra": "34318 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12703,
            "unit": "B/op",
            "extra": "34318 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "34318 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16809,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "68601 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16809,
            "unit": "ns/op",
            "extra": "68601 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "68601 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68601 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36690,
            "unit": "ns/op\t   11310 B/op\t     210 allocs/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36690,
            "unit": "ns/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "32424 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1410,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "790203 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1410,
            "unit": "ns/op",
            "extra": "790203 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "790203 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "790203 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14651,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81511 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14651,
            "unit": "ns/op",
            "extra": "81511 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81511 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81511 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87849,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13689 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87849,
            "unit": "ns/op",
            "extra": "13689 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13689 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13689 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140625,
            "unit": "ns/op\t   41101 B/op\t     999 allocs/op",
            "extra": "8377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140625,
            "unit": "ns/op",
            "extra": "8377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41101,
            "unit": "B/op",
            "extra": "8377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8377 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55674574,
            "unit": "ns/op\t32932563 B/op\t 1121453 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55674574,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932563,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121453,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31140740,
            "unit": "ns/op\t19177325 B/op\t  462248 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31140740,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19177325,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462248,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 53908618,
            "unit": "ns/op\t50875173 B/op\t  526717 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 53908618,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50875173,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526717,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30659799,
            "unit": "ns/op\t19166044 B/op\t  462221 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30659799,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166044,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462221,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 32864183,
            "unit": "ns/op\t19164895 B/op\t  462217 allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 32864183,
            "unit": "ns/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19164895,
            "unit": "B/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462217,
            "unit": "allocs/op",
            "extra": "36 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31015026,
            "unit": "ns/op\t19170837 B/op\t  462221 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31015026,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19170837,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462221,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3511,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "319519 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3511,
            "unit": "ns/op",
            "extra": "319519 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "319519 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "319519 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4907,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "224418 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4907,
            "unit": "ns/op",
            "extra": "224418 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "224418 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "224418 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37164,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32211 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37164,
            "unit": "ns/op",
            "extra": "32211 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32211 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32211 times\n4 procs"
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
          "id": "8dea94239738eff660b84944cc1b7bc24f8c459a",
          "message": "Merge pull request #631 from adamdecaf/fix-initial-data-dir-testing\n\nfix: set INITIAL_DATA_DIRECTORY better in tests",
          "timestamp": "2025-05-16T20:40:05Z",
          "url": "https://github.com/moov-io/watchman/commit/8dea94239738eff660b84944cc1b7bc24f8c459a"
        },
        "date": 1747428232210,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10139,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "113377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10139,
            "unit": "ns/op",
            "extra": "113377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "113377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "113377 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 36297,
            "unit": "ns/op\t   12703 B/op\t     174 allocs/op",
            "extra": "34057 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 36297,
            "unit": "ns/op",
            "extra": "34057 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12703,
            "unit": "B/op",
            "extra": "34057 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "34057 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16823,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "69417 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16823,
            "unit": "ns/op",
            "extra": "69417 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "69417 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69417 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36031,
            "unit": "ns/op\t   11310 B/op\t     210 allocs/op",
            "extra": "33277 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36031,
            "unit": "ns/op",
            "extra": "33277 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33277 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "33277 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1412,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "815524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1412,
            "unit": "ns/op",
            "extra": "815524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "815524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "815524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14766,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14766,
            "unit": "ns/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "80018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87363,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13656 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87363,
            "unit": "ns/op",
            "extra": "13656 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13656 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13656 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140319,
            "unit": "ns/op\t   41101 B/op\t     999 allocs/op",
            "extra": "8469 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140319,
            "unit": "ns/op",
            "extra": "8469 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41101,
            "unit": "B/op",
            "extra": "8469 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8469 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55052790,
            "unit": "ns/op\t32932865 B/op\t 1121454 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55052790,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932865,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121454,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29394152,
            "unit": "ns/op\t19171690 B/op\t  462221 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29394152,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19171690,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462221,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 51143407,
            "unit": "ns/op\t50903053 B/op\t  526698 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 51143407,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50903053,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526698,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 32932726,
            "unit": "ns/op\t19168414 B/op\t  462194 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 32932726,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19168414,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462194,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 29356712,
            "unit": "ns/op\t19164239 B/op\t  462192 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 29356712,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19164239,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462192,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29479638,
            "unit": "ns/op\t19167232 B/op\t  462195 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29479638,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167232,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462195,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3412,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "326546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3412,
            "unit": "ns/op",
            "extra": "326546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "326546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "326546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4881,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "248164 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4881,
            "unit": "ns/op",
            "extra": "248164 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "248164 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "248164 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37114,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32308 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37114,
            "unit": "ns/op",
            "extra": "32308 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32308 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32308 times\n4 procs"
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
          "id": "144742eb1a6766ea02967d9991509fc33a0605d2",
          "message": "Merge pull request #632 from adamdecaf/test-debugging-2025-05-16\n\nbuild: debug slow tests in CI",
          "timestamp": "2025-05-16T22:22:58Z",
          "url": "https://github.com/moov-io/watchman/commit/144742eb1a6766ea02967d9991509fc33a0605d2"
        },
        "date": 1747434417892,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10401,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "111430 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10401,
            "unit": "ns/op",
            "extra": "111430 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "111430 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "111430 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35255,
            "unit": "ns/op\t   12703 B/op\t     174 allocs/op",
            "extra": "33975 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35255,
            "unit": "ns/op",
            "extra": "33975 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12703,
            "unit": "B/op",
            "extra": "33975 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "33975 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 17003,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "68428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 17003,
            "unit": "ns/op",
            "extra": "68428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "68428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 37015,
            "unit": "ns/op\t   11311 B/op\t     210 allocs/op",
            "extra": "32910 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 37015,
            "unit": "ns/op",
            "extra": "32910 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "32910 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "32910 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1427,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "819410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1427,
            "unit": "ns/op",
            "extra": "819410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "819410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "819410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14450,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14450,
            "unit": "ns/op",
            "extra": "83286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 88605,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "12512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 88605,
            "unit": "ns/op",
            "extra": "12512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "12512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "12512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 142242,
            "unit": "ns/op\t   41101 B/op\t     999 allocs/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 142242,
            "unit": "ns/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41101,
            "unit": "B/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8762 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 56654923,
            "unit": "ns/op\t32932414 B/op\t 1121426 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 56654923,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932414,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121426,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31132037,
            "unit": "ns/op\t19174423 B/op\t  462246 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31132037,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19174423,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462246,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 51943302,
            "unit": "ns/op\t50891987 B/op\t  526724 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 51943302,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50891987,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526724,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30965702,
            "unit": "ns/op\t19166018 B/op\t  462220 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30965702,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166018,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462220,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30959017,
            "unit": "ns/op\t19166997 B/op\t  462219 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30959017,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166997,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462219,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30895044,
            "unit": "ns/op\t19167025 B/op\t  462218 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30895044,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167025,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462218,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3447,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "322293 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3447,
            "unit": "ns/op",
            "extra": "322293 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "322293 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "322293 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4810,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "231340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4810,
            "unit": "ns/op",
            "extra": "231340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "231340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "231340 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 38031,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31530 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 38031,
            "unit": "ns/op",
            "extra": "31530 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31530 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31530 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "8af54945bfff50f92a1567b1d23391a5d1ac88ba",
          "message": "release v0.52.0",
          "timestamp": "2025-05-16T22:29:36Z",
          "url": "https://github.com/moov-io/watchman/commit/8af54945bfff50f92a1567b1d23391a5d1ac88ba"
        },
        "date": 1747486296708,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9962,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "115641 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9962,
            "unit": "ns/op",
            "extra": "115641 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "115641 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "115641 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35661,
            "unit": "ns/op\t   12703 B/op\t     174 allocs/op",
            "extra": "34564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35661,
            "unit": "ns/op",
            "extra": "34564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12703,
            "unit": "B/op",
            "extra": "34564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "34564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16773,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "68565 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16773,
            "unit": "ns/op",
            "extra": "68565 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "68565 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68565 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36403,
            "unit": "ns/op\t   11310 B/op\t     210 allocs/op",
            "extra": "33192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36403,
            "unit": "ns/op",
            "extra": "33192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "33192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1419,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "799154 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1419,
            "unit": "ns/op",
            "extra": "799154 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "799154 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "799154 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14477,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83576 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14477,
            "unit": "ns/op",
            "extra": "83576 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83576 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83576 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 86655,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13771 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 86655,
            "unit": "ns/op",
            "extra": "13771 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13771 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13771 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140708,
            "unit": "ns/op\t   41100 B/op\t     999 allocs/op",
            "extra": "8403 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140708,
            "unit": "ns/op",
            "extra": "8403 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41100,
            "unit": "B/op",
            "extra": "8403 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8403 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 54812941,
            "unit": "ns/op\t32932632 B/op\t 1121440 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 54812941,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932632,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121440,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29801234,
            "unit": "ns/op\t19170793 B/op\t  462234 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29801234,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19170793,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462234,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 52110565,
            "unit": "ns/op\t50888983 B/op\t  526714 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 52110565,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50888983,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526714,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 33564709,
            "unit": "ns/op\t19173930 B/op\t  462211 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 33564709,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19173930,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462211,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30031302,
            "unit": "ns/op\t19172155 B/op\t  462211 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30031302,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19172155,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462211,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29576110,
            "unit": "ns/op\t19167417 B/op\t  462211 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29576110,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167417,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462211,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3406,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "343642 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3406,
            "unit": "ns/op",
            "extra": "343642 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "343642 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "343642 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4733,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "219890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4733,
            "unit": "ns/op",
            "extra": "219890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "219890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "219890 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36814,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32635 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36814,
            "unit": "ns/op",
            "extra": "32635 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32635 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32635 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "bafea522ee2c7923836fbd6262e5f37e7cdcdfbd",
          "message": "build: static docker job needs docker hub job [skip ci]",
          "timestamp": "2025-05-19T14:47:07Z",
          "url": "https://github.com/moov-io/watchman/commit/bafea522ee2c7923836fbd6262e5f37e7cdcdfbd"
        },
        "date": 1747666812400,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10219,
            "unit": "ns/op\t    3648 B/op\t     124 allocs/op",
            "extra": "110542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10219,
            "unit": "ns/op",
            "extra": "110542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3648,
            "unit": "B/op",
            "extra": "110542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 124,
            "unit": "allocs/op",
            "extra": "110542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35688,
            "unit": "ns/op\t   12703 B/op\t     174 allocs/op",
            "extra": "34275 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35688,
            "unit": "ns/op",
            "extra": "34275 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12703,
            "unit": "B/op",
            "extra": "34275 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 174,
            "unit": "allocs/op",
            "extra": "34275 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16921,
            "unit": "ns/op\t    4516 B/op\t     176 allocs/op",
            "extra": "68946 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16921,
            "unit": "ns/op",
            "extra": "68946 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4516,
            "unit": "B/op",
            "extra": "68946 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68946 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 37246,
            "unit": "ns/op\t   11311 B/op\t     210 allocs/op",
            "extra": "32031 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 37246,
            "unit": "ns/op",
            "extra": "32031 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "32031 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "32031 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1412,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "780272 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1412,
            "unit": "ns/op",
            "extra": "780272 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "780272 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "780272 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14638,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "80773 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14638,
            "unit": "ns/op",
            "extra": "80773 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "80773 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "80773 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87671,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13609 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87671,
            "unit": "ns/op",
            "extra": "13609 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13609 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13609 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 142108,
            "unit": "ns/op\t   41105 B/op\t     999 allocs/op",
            "extra": "8504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 142108,
            "unit": "ns/op",
            "extra": "8504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41105,
            "unit": "B/op",
            "extra": "8504 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8504 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 56222780,
            "unit": "ns/op\t32933057 B/op\t 1121492 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 56222780,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933057,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121492,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 33602414,
            "unit": "ns/op\t19161921 B/op\t  462242 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 33602414,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19161921,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462242,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 54718283,
            "unit": "ns/op\t50884620 B/op\t  526722 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 54718283,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50884620,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526722,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31547735,
            "unit": "ns/op\t19166040 B/op\t  462217 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31547735,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166040,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462217,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31629175,
            "unit": "ns/op\t19163643 B/op\t  462216 allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31629175,
            "unit": "ns/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19163643,
            "unit": "B/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462216,
            "unit": "allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31777787,
            "unit": "ns/op\t19162456 B/op\t  462217 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31777787,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19162456,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462217,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3453,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "361149 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3453,
            "unit": "ns/op",
            "extra": "361149 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "361149 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "361149 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4827,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "271561 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4827,
            "unit": "ns/op",
            "extra": "271561 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "271561 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "271561 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37467,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32080 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37467,
            "unit": "ns/op",
            "extra": "32080 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32080 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32080 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "6464020ab9482b3965a13772ab02b7450a82a81d",
          "message": "chore: linter fixup",
          "timestamp": "2025-05-19T20:21:19Z",
          "url": "https://github.com/moov-io/watchman/commit/6464020ab9482b3965a13772ab02b7450a82a81d"
        },
        "date": 1747745891913,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10528,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "111512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10528,
            "unit": "ns/op",
            "extra": "111512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "111512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "111512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35287,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35287,
            "unit": "ns/op",
            "extra": "34804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16641,
            "unit": "ns/op\t    4517 B/op\t     176 allocs/op",
            "extra": "69672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16641,
            "unit": "ns/op",
            "extra": "69672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4517,
            "unit": "B/op",
            "extra": "69672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 35995,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33482 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 35995,
            "unit": "ns/op",
            "extra": "33482 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33482 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33482 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1392,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "812150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1392,
            "unit": "ns/op",
            "extra": "812150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "812150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "812150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14379,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83468 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14379,
            "unit": "ns/op",
            "extra": "83468 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83468 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83468 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87975,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13383 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87975,
            "unit": "ns/op",
            "extra": "13383 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13383 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13383 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141643,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8240 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141643,
            "unit": "ns/op",
            "extra": "8240 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8240 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8240 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53824808,
            "unit": "ns/op\t32932798 B/op\t 1121471 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53824808,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932798,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121471,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29932423,
            "unit": "ns/op\t19169586 B/op\t  462246 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29932423,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19169586,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462246,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 51191292,
            "unit": "ns/op\t50888053 B/op\t  526719 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 51191292,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50888053,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526719,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30686178,
            "unit": "ns/op\t19164448 B/op\t  462218 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30686178,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19164448,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462218,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30877794,
            "unit": "ns/op\t19167678 B/op\t  462219 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30877794,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19167678,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462219,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30338587,
            "unit": "ns/op\t19165743 B/op\t  462217 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30338587,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19165743,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462217,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3456,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "341764 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3456,
            "unit": "ns/op",
            "extra": "341764 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "341764 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "341764 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4792,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "219652 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4792,
            "unit": "ns/op",
            "extra": "219652 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "219652 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "219652 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37332,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32085 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37332,
            "unit": "ns/op",
            "extra": "32085 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32085 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32085 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "341484fd5664c0574e33c6d78f320682424051b7",
          "message": "build: upgrade Github actions",
          "timestamp": "2025-05-20T20:43:52Z",
          "url": "https://github.com/moov-io/watchman/commit/341484fd5664c0574e33c6d78f320682424051b7"
        },
        "date": 1747832280845,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10571,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "108358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10571,
            "unit": "ns/op",
            "extra": "108358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "108358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "108358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35637,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "33640 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35637,
            "unit": "ns/op",
            "extra": "33640 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "33640 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33640 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16866,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "68714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16866,
            "unit": "ns/op",
            "extra": "68714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "68714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36542,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "32668 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36542,
            "unit": "ns/op",
            "extra": "32668 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "32668 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "32668 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1389,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "785328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1389,
            "unit": "ns/op",
            "extra": "785328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "785328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "785328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14377,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14377,
            "unit": "ns/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87747,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13629 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87747,
            "unit": "ns/op",
            "extra": "13629 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13629 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13629 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141587,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141587,
            "unit": "ns/op",
            "extra": "8464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8464 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53470678,
            "unit": "ns/op\t32933047 B/op\t 1121491 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53470678,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933047,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121491,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 32713917,
            "unit": "ns/op\t19173752 B/op\t  462226 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 32713917,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19173752,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462226,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 50665096,
            "unit": "ns/op\t50898145 B/op\t  526700 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 50665096,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50898145,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526700,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 29486662,
            "unit": "ns/op\t19168477 B/op\t  462199 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 29486662,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19168477,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462199,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 29568867,
            "unit": "ns/op\t19168583 B/op\t  462202 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 29568867,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19168583,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29596407,
            "unit": "ns/op\t19160611 B/op\t  462198 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29596407,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19160611,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462198,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3490,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "321738 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3490,
            "unit": "ns/op",
            "extra": "321738 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "321738 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "321738 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4716,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "231337 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4716,
            "unit": "ns/op",
            "extra": "231337 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "231337 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "231337 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37266,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32217 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37266,
            "unit": "ns/op",
            "extra": "32217 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32217 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32217 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "4ca30a925e3cedb1cb92573a13a0737d04aa7812",
          "message": "integrity: treat query entity as api-request",
          "timestamp": "2025-05-23T17:20:31Z",
          "url": "https://github.com/moov-io/watchman/commit/4ca30a925e3cedb1cb92573a13a0737d04aa7812"
        },
        "date": 1748091099972,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10759,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "106977 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10759,
            "unit": "ns/op",
            "extra": "106977 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "106977 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "106977 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34703,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "33658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34703,
            "unit": "ns/op",
            "extra": "33658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "33658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16750,
            "unit": "ns/op\t    4517 B/op\t     176 allocs/op",
            "extra": "70060 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16750,
            "unit": "ns/op",
            "extra": "70060 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4517,
            "unit": "B/op",
            "extra": "70060 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "70060 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 35919,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 35919,
            "unit": "ns/op",
            "extra": "33442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1440,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "810097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1440,
            "unit": "ns/op",
            "extra": "810097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "810097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "810097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14889,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "80324 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14889,
            "unit": "ns/op",
            "extra": "80324 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "80324 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "80324 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87236,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87236,
            "unit": "ns/op",
            "extra": "13752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13752 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140268,
            "unit": "ns/op\t   41102 B/op\t     999 allocs/op",
            "extra": "8397 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140268,
            "unit": "ns/op",
            "extra": "8397 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41102,
            "unit": "B/op",
            "extra": "8397 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8397 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 61773760,
            "unit": "ns/op\t32932582 B/op\t 1121435 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 61773760,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932582,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121435,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31190233,
            "unit": "ns/op\t19167932 B/op\t  462231 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31190233,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19167932,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462231,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 68324360,
            "unit": "ns/op\t51410307 B/op\t  559063 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 68324360,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51410307,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559063,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31189945,
            "unit": "ns/op\t19166700 B/op\t  462205 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31189945,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166700,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462205,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30959337,
            "unit": "ns/op\t19168636 B/op\t  462207 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30959337,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19168636,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462207,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31132879,
            "unit": "ns/op\t19166060 B/op\t  462206 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31132879,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19166060,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462206,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3443,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "352555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3443,
            "unit": "ns/op",
            "extra": "352555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "352555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "352555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4814,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "235124 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4814,
            "unit": "ns/op",
            "extra": "235124 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "235124 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "235124 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37519,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31935 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37519,
            "unit": "ns/op",
            "extra": "31935 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31935 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31935 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "68cb487f41edbeda3611884001e00a0c455debe6",
          "message": "download: increase default timeout to 60s",
          "timestamp": "2025-06-02T14:59:45Z",
          "url": "https://github.com/moov-io/watchman/commit/68cb487f41edbeda3611884001e00a0c455debe6"
        },
        "date": 1748876846883,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10530,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "110282 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10530,
            "unit": "ns/op",
            "extra": "110282 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "110282 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "110282 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35073,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "34332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35073,
            "unit": "ns/op",
            "extra": "34332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "34332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16855,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69038 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16855,
            "unit": "ns/op",
            "extra": "69038 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69038 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69038 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36429,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36429,
            "unit": "ns/op",
            "extra": "33283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1437,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "784602 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1437,
            "unit": "ns/op",
            "extra": "784602 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "784602 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "784602 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14618,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "82707 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14618,
            "unit": "ns/op",
            "extra": "82707 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "82707 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "82707 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87558,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13660 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87558,
            "unit": "ns/op",
            "extra": "13660 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13660 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13660 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140184,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8443 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140184,
            "unit": "ns/op",
            "extra": "8443 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8443 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8443 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 54574359,
            "unit": "ns/op\t32933311 B/op\t 1121477 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 54574359,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933311,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121477,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 30095405,
            "unit": "ns/op\t19170479 B/op\t  462214 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 30095405,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19170479,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462214,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 70756698,
            "unit": "ns/op\t51410235 B/op\t  559046 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 70756698,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51410235,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559046,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 34067099,
            "unit": "ns/op\t19166399 B/op\t  462185 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 34067099,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166399,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462185,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 29821200,
            "unit": "ns/op\t19162877 B/op\t  462186 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 29821200,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19162877,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462186,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29705084,
            "unit": "ns/op\t19159090 B/op\t  462185 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29705084,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19159090,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462185,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3492,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "372252 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3492,
            "unit": "ns/op",
            "extra": "372252 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "372252 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "372252 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4871,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "242834 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4871,
            "unit": "ns/op",
            "extra": "242834 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "242834 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "242834 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37401,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32126 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37401,
            "unit": "ns/op",
            "extra": "32126 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32126 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32126 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "d3461d82637cbc508d533c3103a33692a65ddea9",
          "message": "build: setup CSL US benchmarks",
          "timestamp": "2025-06-02T15:14:02Z",
          "url": "https://github.com/moov-io/watchman/commit/d3461d82637cbc508d533c3103a33692a65ddea9"
        },
        "date": 1748877529058,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10315,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109134 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10315,
            "unit": "ns/op",
            "extra": "109134 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109134 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109134 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34614,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34614,
            "unit": "ns/op",
            "extra": "34893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 17393,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "67513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 17393,
            "unit": "ns/op",
            "extra": "67513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "67513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "67513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 37489,
            "unit": "ns/op\t   11311 B/op\t     211 allocs/op",
            "extra": "31851 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 37489,
            "unit": "ns/op",
            "extra": "31851 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "31851 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "31851 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1412,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "795829 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1412,
            "unit": "ns/op",
            "extra": "795829 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "795829 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "795829 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14366,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83262 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14366,
            "unit": "ns/op",
            "extra": "83262 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83262 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83262 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 89222,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 89222,
            "unit": "ns/op",
            "extra": "13332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13332 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 143976,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 143976,
            "unit": "ns/op",
            "extra": "8300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8300 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55942525,
            "unit": "ns/op\t32932953 B/op\t 1121478 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55942525,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932953,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121478,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31648946,
            "unit": "ns/op\t19170006 B/op\t  462242 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31648946,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19170006,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462242,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 68186641,
            "unit": "ns/op\t51402099 B/op\t  559074 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 68186641,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51402099,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559074,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 32233461,
            "unit": "ns/op\t19170218 B/op\t  462216 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 32233461,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19170218,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462216,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31695936,
            "unit": "ns/op\t19161003 B/op\t  462213 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31695936,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19161003,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462213,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 32181973,
            "unit": "ns/op\t19169816 B/op\t  462221 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 32181973,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19169816,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462221,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3596,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "322957 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3596,
            "unit": "ns/op",
            "extra": "322957 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "322957 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "322957 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4998,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "262983 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4998,
            "unit": "ns/op",
            "extra": "262983 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "262983 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "262983 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37423,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32107 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37423,
            "unit": "ns/op",
            "extra": "32107 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32107 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32107 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "fed362349953dc24d1cb88e7a3b876ca7e0a2ade",
          "message": "sources/ofac: add benchmark for FindEntity",
          "timestamp": "2025-06-02T15:44:46Z",
          "url": "https://github.com/moov-io/watchman/commit/fed362349953dc24d1cb88e7a3b876ca7e0a2ade"
        },
        "date": 1748879378448,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10476,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109816 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10476,
            "unit": "ns/op",
            "extra": "109816 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109816 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109816 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35351,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "33014 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35351,
            "unit": "ns/op",
            "extra": "33014 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "33014 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33014 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16581,
            "unit": "ns/op\t    4517 B/op\t     176 allocs/op",
            "extra": "69328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16581,
            "unit": "ns/op",
            "extra": "69328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4517,
            "unit": "B/op",
            "extra": "69328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69328 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 35863,
            "unit": "ns/op\t   11311 B/op\t     211 allocs/op",
            "extra": "33513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 35863,
            "unit": "ns/op",
            "extra": "33513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "33513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33513 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1423,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "778651 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1423,
            "unit": "ns/op",
            "extra": "778651 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "778651 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "778651 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14625,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81757 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14625,
            "unit": "ns/op",
            "extra": "81757 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81757 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81757 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87210,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87210,
            "unit": "ns/op",
            "extra": "13611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140144,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8211 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140144,
            "unit": "ns/op",
            "extra": "8211 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8211 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8211 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55401360,
            "unit": "ns/op\t32933191 B/op\t 1121460 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55401360,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933191,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121460,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31186992,
            "unit": "ns/op\t19166228 B/op\t  462228 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31186992,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19166228,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462228,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66938798,
            "unit": "ns/op\t51384616 B/op\t  559054 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66938798,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51384616,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559054,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30771295,
            "unit": "ns/op\t19165858 B/op\t  462201 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30771295,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19165858,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462201,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31053902,
            "unit": "ns/op\t19167927 B/op\t  462202 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31053902,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19167927,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30976448,
            "unit": "ns/op\t19168615 B/op\t  462201 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30976448,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19168615,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462201,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3496,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "339259 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3496,
            "unit": "ns/op",
            "extra": "339259 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "339259 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "339259 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4812,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "245976 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4812,
            "unit": "ns/op",
            "extra": "245976 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "245976 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "245976 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37123,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32226 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37123,
            "unit": "ns/op",
            "extra": "32226 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32226 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32226 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "cf9230c20b01ada2622a7c5bd694d8b735806827",
          "message": "sources/csl_us: set ReuseRecord on reader",
          "timestamp": "2025-06-02T19:04:15Z",
          "url": "https://github.com/moov-io/watchman/commit/cf9230c20b01ada2622a7c5bd694d8b735806827"
        },
        "date": 1748891367773,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10311,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "110678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10311,
            "unit": "ns/op",
            "extra": "110678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "110678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "110678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34623,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "33117 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34623,
            "unit": "ns/op",
            "extra": "33117 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "33117 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33117 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16999,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16999,
            "unit": "ns/op",
            "extra": "69192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36718,
            "unit": "ns/op\t   11311 B/op\t     211 allocs/op",
            "extra": "33022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36718,
            "unit": "ns/op",
            "extra": "33022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "33022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1401,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "777310 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1401,
            "unit": "ns/op",
            "extra": "777310 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "777310 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "777310 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14314,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "84202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14314,
            "unit": "ns/op",
            "extra": "84202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "84202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "84202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 85437,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 85437,
            "unit": "ns/op",
            "extra": "13954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 138190,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 138190,
            "unit": "ns/op",
            "extra": "8254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8254 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 60356239,
            "unit": "ns/op\t32932517 B/op\t 1121455 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 60356239,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932517,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121455,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 30466357,
            "unit": "ns/op\t19173685 B/op\t  462220 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 30466357,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19173685,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462220,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66757170,
            "unit": "ns/op\t51390262 B/op\t  559046 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66757170,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51390262,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559046,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31559642,
            "unit": "ns/op\t19171819 B/op\t  462191 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31559642,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19171819,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462191,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 33809802,
            "unit": "ns/op\t19171013 B/op\t  462191 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 33809802,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19171013,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462191,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30847100,
            "unit": "ns/op\t19171642 B/op\t  462193 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30847100,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19171642,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462193,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3523,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "384576 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3523,
            "unit": "ns/op",
            "extra": "384576 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "384576 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "384576 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4923,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "233155 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4923,
            "unit": "ns/op",
            "extra": "233155 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "233155 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "233155 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36880,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32562 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36880,
            "unit": "ns/op",
            "extra": "32562 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32562 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32562 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "587452d8d5ef068044c0c61b5bd29f48aaa99dcf",
          "message": "build: fix release yaml",
          "timestamp": "2025-06-02T19:31:36Z",
          "url": "https://github.com/moov-io/watchman/commit/587452d8d5ef068044c0c61b5bd29f48aaa99dcf"
        },
        "date": 1748893848762,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10319,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10319,
            "unit": "ns/op",
            "extra": "109578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34920,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34650 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34920,
            "unit": "ns/op",
            "extra": "34650 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34650 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34650 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16830,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69847 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16830,
            "unit": "ns/op",
            "extra": "69847 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69847 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69847 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36071,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33304 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36071,
            "unit": "ns/op",
            "extra": "33304 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33304 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33304 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1399,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "812498 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1399,
            "unit": "ns/op",
            "extra": "812498 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "812498 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "812498 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14304,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "84297 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14304,
            "unit": "ns/op",
            "extra": "84297 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "84297 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "84297 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87691,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13699 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87691,
            "unit": "ns/op",
            "extra": "13699 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13699 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13699 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140718,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140718,
            "unit": "ns/op",
            "extra": "8287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8287 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55741452,
            "unit": "ns/op\t32933191 B/op\t 1121472 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55741452,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933191,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121472,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 34201375,
            "unit": "ns/op\t19177615 B/op\t  462229 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 34201375,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19177615,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462229,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66005975,
            "unit": "ns/op\t51393892 B/op\t  559056 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66005975,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51393892,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559056,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30964933,
            "unit": "ns/op\t19164557 B/op\t  462202 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30964933,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19164557,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30451948,
            "unit": "ns/op\t19173002 B/op\t  462210 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30451948,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19173002,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462210,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30294446,
            "unit": "ns/op\t19161535 B/op\t  462206 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30294446,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19161535,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462206,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3388,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "358346 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3388,
            "unit": "ns/op",
            "extra": "358346 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "358346 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "358346 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4805,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "249524 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4805,
            "unit": "ns/op",
            "extra": "249524 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "249524 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "249524 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37399,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32186 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37399,
            "unit": "ns/op",
            "extra": "32186 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32186 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32186 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "62f8ed7cf4621a3555c8a88ea7122e24fdcc0f18",
          "message": "docs: add deepwiki badge",
          "timestamp": "2025-06-02T19:34:21Z",
          "url": "https://github.com/moov-io/watchman/commit/62f8ed7cf4621a3555c8a88ea7122e24fdcc0f18"
        },
        "date": 1748955596318,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10691,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109510 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10691,
            "unit": "ns/op",
            "extra": "109510 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109510 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109510 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35942,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "33679 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35942,
            "unit": "ns/op",
            "extra": "33679 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "33679 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33679 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16824,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69315 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16824,
            "unit": "ns/op",
            "extra": "69315 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69315 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69315 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36321,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36321,
            "unit": "ns/op",
            "extra": "33064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1439,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "792486 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1439,
            "unit": "ns/op",
            "extra": "792486 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "792486 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "792486 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14763,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14763,
            "unit": "ns/op",
            "extra": "81766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87071,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87071,
            "unit": "ns/op",
            "extra": "13392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13392 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141208,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8371 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141208,
            "unit": "ns/op",
            "extra": "8371 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8371 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8371 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 56670902,
            "unit": "ns/op\t32932221 B/op\t 1121434 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 56670902,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932221,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121434,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 34995852,
            "unit": "ns/op\t19162567 B/op\t  462207 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 34995852,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19162567,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462207,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 68602590,
            "unit": "ns/op\t51393229 B/op\t  559037 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 68602590,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51393229,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559037,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31477375,
            "unit": "ns/op\t19169373 B/op\t  462181 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31477375,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19169373,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462181,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31599393,
            "unit": "ns/op\t19166167 B/op\t  462178 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31599393,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166167,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462178,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31372989,
            "unit": "ns/op\t19168339 B/op\t  462180 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31372989,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19168339,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462180,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3484,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "326388 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3484,
            "unit": "ns/op",
            "extra": "326388 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "326388 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "326388 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4755,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "215803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4755,
            "unit": "ns/op",
            "extra": "215803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "215803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "215803 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37467,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32235 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37467,
            "unit": "ns/op",
            "extra": "32235 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32235 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32235 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "133a1ca092b172609d329c1eea33946444687a59",
          "message": "build: log go test output during benchmark failures",
          "timestamp": "2025-06-03T13:21:23Z",
          "url": "https://github.com/moov-io/watchman/commit/133a1ca092b172609d329c1eea33946444687a59"
        },
        "date": 1748957176152,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10385,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "108900 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10385,
            "unit": "ns/op",
            "extra": "108900 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "108900 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "108900 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34533,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34657 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34533,
            "unit": "ns/op",
            "extra": "34657 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34657 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34657 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16775,
            "unit": "ns/op\t    4517 B/op\t     176 allocs/op",
            "extra": "69621 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16775,
            "unit": "ns/op",
            "extra": "69621 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4517,
            "unit": "B/op",
            "extra": "69621 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69621 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 35676,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 35676,
            "unit": "ns/op",
            "extra": "33487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1417,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "780411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1417,
            "unit": "ns/op",
            "extra": "780411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "780411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "780411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14408,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "83281 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14408,
            "unit": "ns/op",
            "extra": "83281 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "83281 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "83281 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 86351,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13878 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 86351,
            "unit": "ns/op",
            "extra": "13878 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13878 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13878 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 138270,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 138270,
            "unit": "ns/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8522 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55109722,
            "unit": "ns/op\t32932542 B/op\t 1121436 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55109722,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932542,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121436,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 30265940,
            "unit": "ns/op\t19177607 B/op\t  462244 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 30265940,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19177607,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462244,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 65739485,
            "unit": "ns/op\t51396723 B/op\t  559074 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 65739485,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51396723,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559074,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 32334105,
            "unit": "ns/op\t19169307 B/op\t  462213 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 32334105,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19169307,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462213,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 32303184,
            "unit": "ns/op\t19166608 B/op\t  462216 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 32303184,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166608,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462216,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30593653,
            "unit": "ns/op\t19167323 B/op\t  462214 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30593653,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167323,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462214,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3474,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "337890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3474,
            "unit": "ns/op",
            "extra": "337890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "337890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "337890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4590,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "247988 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4590,
            "unit": "ns/op",
            "extra": "247988 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "247988 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "247988 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37858,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31738 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37858,
            "unit": "ns/op",
            "extra": "31738 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31738 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31738 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "71ee9eb9e76b0286bccca0d626bc2920e2a2a50e",
          "message": "search: allow filtering by sourceID",
          "timestamp": "2025-06-03T17:07:13Z",
          "url": "https://github.com/moov-io/watchman/commit/71ee9eb9e76b0286bccca0d626bc2920e2a2a50e"
        },
        "date": 1748970724661,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 203,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "5624287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 203,
            "unit": "ns/op",
            "extra": "5624287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "5624287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5624287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 8869,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "132534 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 8869,
            "unit": "ns/op",
            "extra": "132534 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "132534 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "132534 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 202.6,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "5695669 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 202.6,
            "unit": "ns/op",
            "extra": "5695669 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "5695669 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5695669 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 8866,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "133562 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 8866,
            "unit": "ns/op",
            "extra": "133562 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "133562 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "133562 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 212.5,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "5674820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 212.5,
            "unit": "ns/op",
            "extra": "5674820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "5674820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5674820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 8932,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "134306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 8932,
            "unit": "ns/op",
            "extra": "134306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "134306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "134306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 212.6,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "5660935 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 212.6,
            "unit": "ns/op",
            "extra": "5660935 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "5660935 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5660935 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 9017,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "133923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 9017,
            "unit": "ns/op",
            "extra": "133923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "133923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "133923 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 57362176,
            "unit": "ns/op\t32932357 B/op\t 1121421 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 57362176,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932357,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121421,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31937972,
            "unit": "ns/op\t19172772 B/op\t  462232 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31937972,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19172772,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462232,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 78090313,
            "unit": "ns/op\t51391737 B/op\t  559060 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 78090313,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51391737,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559060,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 32433898,
            "unit": "ns/op\t19166599 B/op\t  462205 allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 32433898,
            "unit": "ns/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19166599,
            "unit": "B/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462205,
            "unit": "allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 32039336,
            "unit": "ns/op\t19166774 B/op\t  462204 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 32039336,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166774,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462204,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 32215770,
            "unit": "ns/op\t19172933 B/op\t  462206 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 32215770,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19172933,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462206,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3619,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "317920 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3619,
            "unit": "ns/op",
            "extra": "317920 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "317920 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "317920 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4934,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "236511 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4934,
            "unit": "ns/op",
            "extra": "236511 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "236511 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "236511 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37828,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31791 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37828,
            "unit": "ns/op",
            "extra": "31791 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31791 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31791 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "e2abfd18aa9fcef0b01b148ed238942f5005c924",
          "message": "sources/csl_us: fixup address extraction",
          "timestamp": "2025-06-03T18:05:04Z",
          "url": "https://github.com/moov-io/watchman/commit/e2abfd18aa9fcef0b01b148ed238942f5005c924"
        },
        "date": 1748974213007,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 247.4,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "4504770 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 247.4,
            "unit": "ns/op",
            "extra": "4504770 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "4504770 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "4504770 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 9108,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "128832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 9108,
            "unit": "ns/op",
            "extra": "128832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "128832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "128832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 250.7,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "4747347 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 250.7,
            "unit": "ns/op",
            "extra": "4747347 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "4747347 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "4747347 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 9172,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "131054 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 9172,
            "unit": "ns/op",
            "extra": "131054 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "131054 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "131054 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 210.1,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "5476876 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 210.1,
            "unit": "ns/op",
            "extra": "5476876 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "5476876 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "5476876 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 8860,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "134235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 8860,
            "unit": "ns/op",
            "extra": "134235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "134235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "134235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 254.3,
            "unit": "ns/op\t     448 B/op\t       1 allocs/op",
            "extra": "4700982 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 254.3,
            "unit": "ns/op",
            "extra": "4700982 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 448,
            "unit": "B/op",
            "extra": "4700982 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "4700982 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 9139,
            "unit": "ns/op\t    3393 B/op\t      16 allocs/op",
            "extra": "130453 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 9139,
            "unit": "ns/op",
            "extra": "130453 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 3393,
            "unit": "B/op",
            "extra": "130453 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "130453 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 63582808,
            "unit": "ns/op\t32933369 B/op\t 1121462 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 63582808,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933369,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121462,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31132338,
            "unit": "ns/op\t19170634 B/op\t  462232 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31132338,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19170634,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462232,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66888511,
            "unit": "ns/op\t51412544 B/op\t  559063 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66888511,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51412544,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559063,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30518675,
            "unit": "ns/op\t19167658 B/op\t  462204 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30518675,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19167658,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462204,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31059104,
            "unit": "ns/op\t19166900 B/op\t  462203 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31059104,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166900,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462203,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31046002,
            "unit": "ns/op\t19167830 B/op\t  462203 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31046002,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167830,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462203,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3550,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "321852 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3550,
            "unit": "ns/op",
            "extra": "321852 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "321852 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "321852 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4856,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "226755 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4856,
            "unit": "ns/op",
            "extra": "226755 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "226755 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "226755 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37558,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37558,
            "unit": "ns/op",
            "extra": "31946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31946 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "c4025ef23249d0177537bbe3cba011ce3ad4a844",
          "message": "search: force blank .SourceID for benchmarks",
          "timestamp": "2025-06-03T18:15:37Z",
          "url": "https://github.com/moov-io/watchman/commit/c4025ef23249d0177537bbe3cba011ce3ad4a844"
        },
        "date": 1748974859425,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10589,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109376 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10589,
            "unit": "ns/op",
            "extra": "109376 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109376 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109376 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34576,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34455 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34576,
            "unit": "ns/op",
            "extra": "34455 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34455 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34455 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16641,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16641,
            "unit": "ns/op",
            "extra": "69382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36000,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "33420 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36000,
            "unit": "ns/op",
            "extra": "33420 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "33420 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "33420 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1403,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "820934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1403,
            "unit": "ns/op",
            "extra": "820934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "820934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "820934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14257,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "84554 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14257,
            "unit": "ns/op",
            "extra": "84554 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "84554 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "84554 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 88161,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 88161,
            "unit": "ns/op",
            "extra": "13544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141825,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141825,
            "unit": "ns/op",
            "extra": "8295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8295 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 62771550,
            "unit": "ns/op\t32932336 B/op\t 1121420 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 62771550,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932336,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121420,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31303399,
            "unit": "ns/op\t19176724 B/op\t  462229 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31303399,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19176724,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462229,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66094704,
            "unit": "ns/op\t51389544 B/op\t  559056 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66094704,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51389544,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559056,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31423437,
            "unit": "ns/op\t19167724 B/op\t  462199 allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31423437,
            "unit": "ns/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19167724,
            "unit": "B/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462199,
            "unit": "allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30579385,
            "unit": "ns/op\t19169589 B/op\t  462202 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30579385,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19169589,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31219007,
            "unit": "ns/op\t19167388 B/op\t  462199 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31219007,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167388,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462199,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3659,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "297249 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3659,
            "unit": "ns/op",
            "extra": "297249 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "297249 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "297249 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5118,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "248253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5118,
            "unit": "ns/op",
            "extra": "248253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "248253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "248253 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37659,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31944 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37659,
            "unit": "ns/op",
            "extra": "31944 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31944 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31944 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "71ce43ea7a535205fda1c20fb7a17e47a47a392d",
          "message": "ofactest: check EntityForBenchmark",
          "timestamp": "2025-06-03T18:23:01Z",
          "url": "https://github.com/moov-io/watchman/commit/71ce43ea7a535205fda1c20fb7a17e47a47a392d"
        },
        "date": 1748975451805,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10243,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10243,
            "unit": "ns/op",
            "extra": "109790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34416,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "33578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34416,
            "unit": "ns/op",
            "extra": "33578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "33578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33578 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16707,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16707,
            "unit": "ns/op",
            "extra": "69572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36635,
            "unit": "ns/op\t   11311 B/op\t     211 allocs/op",
            "extra": "32790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36635,
            "unit": "ns/op",
            "extra": "32790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11311,
            "unit": "B/op",
            "extra": "32790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "32790 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1401,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "807368 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1401,
            "unit": "ns/op",
            "extra": "807368 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "807368 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "807368 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14374,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "82428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14374,
            "unit": "ns/op",
            "extra": "82428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "82428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "82428 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87805,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13591 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87805,
            "unit": "ns/op",
            "extra": "13591 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13591 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13591 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141679,
            "unit": "ns/op\t   41103 B/op\t     999 allocs/op",
            "extra": "8264 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141679,
            "unit": "ns/op",
            "extra": "8264 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41103,
            "unit": "B/op",
            "extra": "8264 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8264 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 57103864,
            "unit": "ns/op\t32932434 B/op\t 1121454 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 57103864,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932434,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121454,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 34864631,
            "unit": "ns/op\t19173277 B/op\t  462218 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 34864631,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19173277,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462218,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 67411776,
            "unit": "ns/op\t51410068 B/op\t  559051 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 67411776,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51410068,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559051,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30771694,
            "unit": "ns/op\t19161602 B/op\t  462190 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30771694,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19161602,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462190,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30923245,
            "unit": "ns/op\t19169027 B/op\t  462192 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30923245,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19169027,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462192,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30709147,
            "unit": "ns/op\t19173136 B/op\t  462200 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30709147,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19173136,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462200,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3490,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "339205 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3490,
            "unit": "ns/op",
            "extra": "339205 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "339205 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "339205 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4796,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "220810 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4796,
            "unit": "ns/op",
            "extra": "220810 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "220810 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "220810 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37813,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37813,
            "unit": "ns/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31814 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31814 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "1bfff9bb46540cb7ce12e67487407add4143c772",
          "message": "release v0.52.3",
          "timestamp": "2025-06-11T21:04:00Z",
          "url": "https://github.com/moov-io/watchman/commit/1bfff9bb46540cb7ce12e67487407add4143c772"
        },
        "date": 1749733097121,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10736,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "109394 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10736,
            "unit": "ns/op",
            "extra": "109394 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "109394 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "109394 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35568,
            "unit": "ns/op\t   12711 B/op\t     176 allocs/op",
            "extra": "33984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35568,
            "unit": "ns/op",
            "extra": "33984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12711,
            "unit": "B/op",
            "extra": "33984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "33984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16912,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "69361 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16912,
            "unit": "ns/op",
            "extra": "69361 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "69361 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "69361 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36147,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36147,
            "unit": "ns/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "32834 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1430,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "784118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1430,
            "unit": "ns/op",
            "extra": "784118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "784118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "784118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14658,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81409 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14658,
            "unit": "ns/op",
            "extra": "81409 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81409 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81409 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 86443,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 86443,
            "unit": "ns/op",
            "extra": "13833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 140620,
            "unit": "ns/op\t   41101 B/op\t     999 allocs/op",
            "extra": "8739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 140620,
            "unit": "ns/op",
            "extra": "8739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41101,
            "unit": "B/op",
            "extra": "8739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8739 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55954939,
            "unit": "ns/op\t32933708 B/op\t 1121494 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55954939,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32933708,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121494,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 32757375,
            "unit": "ns/op\t19172970 B/op\t  462216 allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 32757375,
            "unit": "ns/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19172970,
            "unit": "B/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462216,
            "unit": "allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 69895439,
            "unit": "ns/op\t51386474 B/op\t  559037 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 69895439,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51386474,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559037,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 34787222,
            "unit": "ns/op\t19165107 B/op\t  462184 allocs/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 34787222,
            "unit": "ns/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19165107,
            "unit": "B/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462184,
            "unit": "allocs/op",
            "extra": "34 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30755807,
            "unit": "ns/op\t19167070 B/op\t  462183 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30755807,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19167070,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462183,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31600384,
            "unit": "ns/op\t19165554 B/op\t  462185 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31600384,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19165554,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462185,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3584,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "298615 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3584,
            "unit": "ns/op",
            "extra": "298615 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "298615 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "298615 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5217,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "239487 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5217,
            "unit": "ns/op",
            "extra": "239487 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "239487 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "239487 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37466,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32138 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37466,
            "unit": "ns/op",
            "extra": "32138 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32138 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32138 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "7ea8bfff08502d2d72606ec4d88953848d73e404",
          "message": "docs: clarify multiple fields are better",
          "timestamp": "2025-06-13T17:00:48Z",
          "url": "https://github.com/moov-io/watchman/commit/7ea8bfff08502d2d72606ec4d88953848d73e404"
        },
        "date": 1749905605925,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10516,
            "unit": "ns/op\t    3657 B/op\t     126 allocs/op",
            "extra": "108279 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10516,
            "unit": "ns/op",
            "extra": "108279 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3657,
            "unit": "B/op",
            "extra": "108279 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "108279 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34637,
            "unit": "ns/op\t   12710 B/op\t     176 allocs/op",
            "extra": "34723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34637,
            "unit": "ns/op",
            "extra": "34723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12710,
            "unit": "B/op",
            "extra": "34723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16992,
            "unit": "ns/op\t    4518 B/op\t     176 allocs/op",
            "extra": "68006 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16992,
            "unit": "ns/op",
            "extra": "68006 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4518,
            "unit": "B/op",
            "extra": "68006 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "68006 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 36480,
            "unit": "ns/op\t   11310 B/op\t     211 allocs/op",
            "extra": "32901 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 36480,
            "unit": "ns/op",
            "extra": "32901 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11310,
            "unit": "B/op",
            "extra": "32901 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "32901 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1445,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "794676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1445,
            "unit": "ns/op",
            "extra": "794676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "794676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "794676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14749,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "80736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14749,
            "unit": "ns/op",
            "extra": "80736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "80736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "80736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 87801,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "13548 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 87801,
            "unit": "ns/op",
            "extra": "13548 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "13548 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "13548 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 141644,
            "unit": "ns/op\t   41102 B/op\t     999 allocs/op",
            "extra": "8592 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 141644,
            "unit": "ns/op",
            "extra": "8592 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41102,
            "unit": "B/op",
            "extra": "8592 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "8592 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55320246,
            "unit": "ns/op\t32931980 B/op\t 1121422 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55320246,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931980,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121422,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 30453370,
            "unit": "ns/op\t19175082 B/op\t  462234 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 30453370,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19175082,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462234,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 66007436,
            "unit": "ns/op\t51390895 B/op\t  559056 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 66007436,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51390895,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559056,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30513045,
            "unit": "ns/op\t19172057 B/op\t  462204 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30513045,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19172057,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462204,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30548457,
            "unit": "ns/op\t19165095 B/op\t  462201 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30548457,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19165095,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462201,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30480226,
            "unit": "ns/op\t19172404 B/op\t  462204 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30480226,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19172404,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462204,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3540,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "346671 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3540,
            "unit": "ns/op",
            "extra": "346671 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "346671 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "346671 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4788,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "223720 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4788,
            "unit": "ns/op",
            "extra": "223720 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "223720 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "223720 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37418,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32120 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37418,
            "unit": "ns/op",
            "extra": "32120 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32120 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32120 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "92da36132718f885246eec99f8af93ef53d976b6",
          "message": "meta: remove outdated TODOs",
          "timestamp": "2025-06-17T21:35:27Z",
          "url": "https://github.com/moov-io/watchman/commit/92da36132718f885246eec99f8af93ef53d976b6"
        },
        "date": 1750251631101,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10400,
            "unit": "ns/op\t    3656 B/op\t     126 allocs/op",
            "extra": "102780 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10400,
            "unit": "ns/op",
            "extra": "102780 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3656,
            "unit": "B/op",
            "extra": "102780 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "102780 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33555,
            "unit": "ns/op\t   12707 B/op\t     176 allocs/op",
            "extra": "34916 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33555,
            "unit": "ns/op",
            "extra": "34916 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12707,
            "unit": "B/op",
            "extra": "34916 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "34916 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16032,
            "unit": "ns/op\t    4517 B/op\t     176 allocs/op",
            "extra": "71002 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16032,
            "unit": "ns/op",
            "extra": "71002 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4517,
            "unit": "B/op",
            "extra": "71002 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 176,
            "unit": "allocs/op",
            "extra": "71002 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 34060,
            "unit": "ns/op\t   11307 B/op\t     211 allocs/op",
            "extra": "34875 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 34060,
            "unit": "ns/op",
            "extra": "34875 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11307,
            "unit": "B/op",
            "extra": "34875 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 211,
            "unit": "allocs/op",
            "extra": "34875 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1326,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "884893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1326,
            "unit": "ns/op",
            "extra": "884893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "884893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "884893 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13689,
            "unit": "ns/op\t    5233 B/op\t      30 allocs/op",
            "extra": "87097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13689,
            "unit": "ns/op",
            "extra": "87097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5233,
            "unit": "B/op",
            "extra": "87097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "87097 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 82563,
            "unit": "ns/op\t   22788 B/op\t     867 allocs/op",
            "extra": "14445 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 82563,
            "unit": "ns/op",
            "extra": "14445 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 22788,
            "unit": "B/op",
            "extra": "14445 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 867,
            "unit": "allocs/op",
            "extra": "14445 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 131063,
            "unit": "ns/op\t   41088 B/op\t     999 allocs/op",
            "extra": "9064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 131063,
            "unit": "ns/op",
            "extra": "9064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 41088,
            "unit": "B/op",
            "extra": "9064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 999,
            "unit": "allocs/op",
            "extra": "9064 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 55351925,
            "unit": "ns/op\t32932654 B/op\t 1121454 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 55351925,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932654,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121454,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31755147,
            "unit": "ns/op\t19166451 B/op\t  462203 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31755147,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19166451,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462203,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 65444796,
            "unit": "ns/op\t51380334 B/op\t  559031 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 65444796,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51380334,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 559031,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 31003647,
            "unit": "ns/op\t19173600 B/op\t  462177 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 31003647,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19173600,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462177,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 31200299,
            "unit": "ns/op\t19168528 B/op\t  462174 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 31200299,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19168528,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462174,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 31455840,
            "unit": "ns/op\t19157889 B/op\t  462170 allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 31455840,
            "unit": "ns/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19157889,
            "unit": "B/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462170,
            "unit": "allocs/op",
            "extra": "43 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3565,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "396946 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3565,
            "unit": "ns/op",
            "extra": "396946 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "396946 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "396946 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4904,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "221434 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4904,
            "unit": "ns/op",
            "extra": "221434 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "221434 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "221434 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37446,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31983 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37446,
            "unit": "ns/op",
            "extra": "31983 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31983 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31983 times\n4 procs"
          }
        ]
      }
    ]
  }
}