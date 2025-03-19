window.BENCHMARK_DATA = {
  "lastUpdate": 1742388782771,
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
          "id": "5f80e9ca1bf4090c4766d69d354af9e4c82626f1",
          "message": "Merge pull request #605 from adamdecaf/detailed-similarity\n\nsearch: add DetailedSimilarity which exposes direct field comparisons",
          "timestamp": "2025-02-25T21:49:30Z",
          "url": "https://github.com/moov-io/watchman/commit/5f80e9ca1bf4090c4766d69d354af9e4c82626f1"
        },
        "date": 1740574177822,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7664,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "150464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7664,
            "unit": "ns/op",
            "extra": "150464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "150464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "150464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31319,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31319,
            "unit": "ns/op",
            "extra": "38382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4822,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "237009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4822,
            "unit": "ns/op",
            "extra": "237009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "237009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "237009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 23084,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "52450 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 23084,
            "unit": "ns/op",
            "extra": "52450 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "52450 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "52450 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 887.9,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1349814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 887.9,
            "unit": "ns/op",
            "extra": "1349814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1349814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1349814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13279,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "91226 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13279,
            "unit": "ns/op",
            "extra": "91226 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "91226 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "91226 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35396,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "32101 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35396,
            "unit": "ns/op",
            "extra": "32101 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "32101 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "32101 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 82988,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 82988,
            "unit": "ns/op",
            "extra": "14487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14487 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14487 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21198853,
            "unit": "ns/op\t12022515 B/op\t  296854 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21198853,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022515,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19267693,
            "unit": "ns/op\t10561182 B/op\t  205956 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19267693,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10561182,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205956,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 45214091,
            "unit": "ns/op\t42282633 B/op\t  270272 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 45214091,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42282633,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270272,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3389,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "336184 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3389,
            "unit": "ns/op",
            "extra": "336184 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "336184 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "336184 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36883,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32500 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36883,
            "unit": "ns/op",
            "extra": "32500 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32500 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32500 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "e60f73af96aabdf8882a084ccfc1fe07acf2cf8a",
          "message": "add moov-io/watchman Common Benchmarks (go) benchmark result for 5f80e9ca1bf4090c4766d69d354af9e4c82626f1",
          "timestamp": "2025-02-26T12:49:38Z",
          "url": "https://github.com/moov-io/watchman/commit/e60f73af96aabdf8882a084ccfc1fe07acf2cf8a"
        },
        "date": 1740660574315,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7566,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "153802 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7566,
            "unit": "ns/op",
            "extra": "153802 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "153802 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "153802 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31148,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38044 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31148,
            "unit": "ns/op",
            "extra": "38044 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38044 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38044 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4736,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "240333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4736,
            "unit": "ns/op",
            "extra": "240333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "240333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "240333 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22658,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53461 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22658,
            "unit": "ns/op",
            "extra": "53461 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53461 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53461 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 907.3,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1321362 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 907.3,
            "unit": "ns/op",
            "extra": "1321362 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1321362 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1321362 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13643,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "88411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13643,
            "unit": "ns/op",
            "extra": "88411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "88411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "88411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35231,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33717 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35231,
            "unit": "ns/op",
            "extra": "33717 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33717 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33717 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 81767,
            "unit": "ns/op\t   26789 B/op\t     679 allocs/op",
            "extra": "14655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 81767,
            "unit": "ns/op",
            "extra": "14655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26789,
            "unit": "B/op",
            "extra": "14655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14655 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21416647,
            "unit": "ns/op\t12022349 B/op\t  296854 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21416647,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022349,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19705343,
            "unit": "ns/op\t10560643 B/op\t  205957 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19705343,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10560643,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205957,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40680829,
            "unit": "ns/op\t42290892 B/op\t  270275 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40680829,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42290892,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270275,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3450,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "310689 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3450,
            "unit": "ns/op",
            "extra": "310689 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "310689 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "310689 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36931,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36931,
            "unit": "ns/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33097 times\n4 procs"
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
          "id": "ba81c16f6b03e18943ac3c9fe1fdc42f08914f04",
          "message": "build: reading the readme should help",
          "timestamp": "2025-02-27T19:10:30Z",
          "url": "https://github.com/moov-io/watchman/commit/ba81c16f6b03e18943ac3c9fe1fdc42f08914f04"
        },
        "date": 1740746946688,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7517,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "150465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7517,
            "unit": "ns/op",
            "extra": "150465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "150465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "150465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30443,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38929 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30443,
            "unit": "ns/op",
            "extra": "38929 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38929 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38929 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4824,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "232178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4824,
            "unit": "ns/op",
            "extra": "232178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "232178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "232178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22774,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "52472 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22774,
            "unit": "ns/op",
            "extra": "52472 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "52472 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "52472 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 921,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1306284 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 921,
            "unit": "ns/op",
            "extra": "1306284 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1306284 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1306284 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13596,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "87745 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13596,
            "unit": "ns/op",
            "extra": "87745 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "87745 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "87745 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35733,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "32538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35733,
            "unit": "ns/op",
            "extra": "32538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "32538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "32538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 82194,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14518 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 82194,
            "unit": "ns/op",
            "extra": "14518 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14518 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14518 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20620337,
            "unit": "ns/op\t12022545 B/op\t  296854 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20620337,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022545,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18937853,
            "unit": "ns/op\t10561412 B/op\t  205956 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18937853,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10561412,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205956,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 42404369,
            "unit": "ns/op\t42293296 B/op\t  270275 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 42404369,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42293296,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270275,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3407,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "357154 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3407,
            "unit": "ns/op",
            "extra": "357154 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "357154 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "357154 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37020,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32469 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37020,
            "unit": "ns/op",
            "extra": "32469 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32469 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32469 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "4d7bd2fb803879fcedcd48724832cf4ef635b8c6",
          "message": "add moov-io/watchman Common Benchmarks (go) benchmark result for ba81c16f6b03e18943ac3c9fe1fdc42f08914f04",
          "timestamp": "2025-02-28T12:49:07Z",
          "url": "https://github.com/moov-io/watchman/commit/4d7bd2fb803879fcedcd48724832cf4ef635b8c6"
        },
        "date": 1740833048081,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7648,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "153464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7648,
            "unit": "ns/op",
            "extra": "153464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "153464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "153464 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33827,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33827,
            "unit": "ns/op",
            "extra": "38032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38032 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4716,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "241483 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4716,
            "unit": "ns/op",
            "extra": "241483 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "241483 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "241483 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22446,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22446,
            "unit": "ns/op",
            "extra": "53512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 914.9,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1312029 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 914.9,
            "unit": "ns/op",
            "extra": "1312029 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1312029 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1312029 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13572,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "88393 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13572,
            "unit": "ns/op",
            "extra": "88393 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "88393 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "88393 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35355,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "31950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35355,
            "unit": "ns/op",
            "extra": "31950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "31950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "31950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 83222,
            "unit": "ns/op\t   26789 B/op\t     679 allocs/op",
            "extra": "14395 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 83222,
            "unit": "ns/op",
            "extra": "14395 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26789,
            "unit": "B/op",
            "extra": "14395 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14395 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20833395,
            "unit": "ns/op\t12022415 B/op\t  296854 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20833395,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022415,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18811357,
            "unit": "ns/op\t10556785 B/op\t  205954 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18811357,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10556785,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205954,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40041024,
            "unit": "ns/op\t42294348 B/op\t  270277 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40041024,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42294348,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270277,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3403,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "342618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3403,
            "unit": "ns/op",
            "extra": "342618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "342618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "342618 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36583,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32786 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36583,
            "unit": "ns/op",
            "extra": "32786 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32786 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32786 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "10441a6462ab65f25247425dad65630f97b19ced",
          "message": "add moov-io/watchman Common Benchmarks (go) benchmark result for 4d7bd2fb803879fcedcd48724832cf4ef635b8c6",
          "timestamp": "2025-03-01T12:44:08Z",
          "url": "https://github.com/moov-io/watchman/commit/10441a6462ab65f25247425dad65630f97b19ced"
        },
        "date": 1740919439801,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7641,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "150374 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7641,
            "unit": "ns/op",
            "extra": "150374 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "150374 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "150374 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31087,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31087,
            "unit": "ns/op",
            "extra": "38378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38378 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4702,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "235861 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4702,
            "unit": "ns/op",
            "extra": "235861 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "235861 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "235861 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22285,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53421 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22285,
            "unit": "ns/op",
            "extra": "53421 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53421 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53421 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 885,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1357814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 885,
            "unit": "ns/op",
            "extra": "1357814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1357814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1357814 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13072,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "91375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13072,
            "unit": "ns/op",
            "extra": "91375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "91375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "91375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35452,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33696 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35452,
            "unit": "ns/op",
            "extra": "33696 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33696 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33696 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 82888,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14408 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 82888,
            "unit": "ns/op",
            "extra": "14408 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14408 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14408 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20722191,
            "unit": "ns/op\t12022570 B/op\t  296854 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20722191,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022570,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19962521,
            "unit": "ns/op\t10562257 B/op\t  205956 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19962521,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10562257,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205956,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 41447994,
            "unit": "ns/op\t42300706 B/op\t  270277 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 41447994,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42300706,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270277,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3350,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "322023 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3350,
            "unit": "ns/op",
            "extra": "322023 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "322023 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "322023 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36590,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32900 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36590,
            "unit": "ns/op",
            "extra": "32900 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32900 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32900 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "54a9a086dabb8739764ef1c6f6e992d5802348c7",
          "message": "add moov-io/watchman Common Benchmarks (go) benchmark result for 10441a6462ab65f25247425dad65630f97b19ced",
          "timestamp": "2025-03-02T12:44:00Z",
          "url": "https://github.com/moov-io/watchman/commit/54a9a086dabb8739764ef1c6f6e992d5802348c7"
        },
        "date": 1741006197237,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7397,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "152574 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7397,
            "unit": "ns/op",
            "extra": "152574 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "152574 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "152574 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30300,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "39303 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30300,
            "unit": "ns/op",
            "extra": "39303 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "39303 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "39303 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4781,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "240499 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4781,
            "unit": "ns/op",
            "extra": "240499 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "240499 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "240499 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 23155,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "52575 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 23155,
            "unit": "ns/op",
            "extra": "52575 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "52575 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "52575 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 914.6,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1308603 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 914.6,
            "unit": "ns/op",
            "extra": "1308603 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1308603 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1308603 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13528,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "88198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13528,
            "unit": "ns/op",
            "extra": "88198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "88198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "88198 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 34633,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "31411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 34633,
            "unit": "ns/op",
            "extra": "31411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "31411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "31411 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 80585,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14854 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 80585,
            "unit": "ns/op",
            "extra": "14854 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14854 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14854 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23000374,
            "unit": "ns/op\t12022589 B/op\t  296854 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23000374,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022589,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18850013,
            "unit": "ns/op\t10556900 B/op\t  205954 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18850013,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10556900,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205954,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39237293,
            "unit": "ns/op\t42294281 B/op\t  270276 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39237293,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42294281,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270276,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3387,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "365906 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3387,
            "unit": "ns/op",
            "extra": "365906 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "365906 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "365906 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36536,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33000 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36536,
            "unit": "ns/op",
            "extra": "33000 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33000 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33000 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "0da60c685a0e828226a96c866ec9ef56b48013b6",
          "message": "add moov-io/watchman Common Benchmarks (go) benchmark result for 54a9a086dabb8739764ef1c6f6e992d5802348c7",
          "timestamp": "2025-03-03T12:49:57Z",
          "url": "https://github.com/moov-io/watchman/commit/0da60c685a0e828226a96c866ec9ef56b48013b6"
        },
        "date": 1741092590791,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7538,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "151599 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7538,
            "unit": "ns/op",
            "extra": "151599 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "151599 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "151599 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33432,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38692 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33432,
            "unit": "ns/op",
            "extra": "38692 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38692 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38692 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4843,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "240764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4843,
            "unit": "ns/op",
            "extra": "240764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "240764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "240764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22946,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "51906 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22946,
            "unit": "ns/op",
            "extra": "51906 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "51906 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "51906 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 897.2,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1339400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 897.2,
            "unit": "ns/op",
            "extra": "1339400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1339400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1339400 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13613,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "89019 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13613,
            "unit": "ns/op",
            "extra": "89019 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "89019 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "89019 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 36514,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 36514,
            "unit": "ns/op",
            "extra": "33080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 84535,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 84535,
            "unit": "ns/op",
            "extra": "14132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14132 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14132 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21251716,
            "unit": "ns/op\t12022393 B/op\t  296854 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21251716,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022393,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19240126,
            "unit": "ns/op\t10562101 B/op\t  205957 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19240126,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10562101,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205957,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39863299,
            "unit": "ns/op\t42284923 B/op\t  270274 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39863299,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42284923,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270274,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3445,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "375562 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3445,
            "unit": "ns/op",
            "extra": "375562 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "375562 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "375562 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36750,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32739 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36750,
            "unit": "ns/op",
            "extra": "32739 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32739 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32739 times\n4 procs"
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
          "id": "b66c7df45934f21000c1321c8fa742c24c2fb2ac",
          "message": "build: update ./pkg/sources/ path",
          "timestamp": "2025-03-04T15:53:58Z",
          "url": "https://github.com/moov-io/watchman/commit/b66c7df45934f21000c1321c8fa742c24c2fb2ac"
        },
        "date": 1741103824612,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7491,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "154178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7491,
            "unit": "ns/op",
            "extra": "154178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "154178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "154178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30443,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38869 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30443,
            "unit": "ns/op",
            "extra": "38869 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38869 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38869 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4813,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "238150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4813,
            "unit": "ns/op",
            "extra": "238150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "238150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "238150 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 23006,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "52418 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 23006,
            "unit": "ns/op",
            "extra": "52418 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "52418 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "52418 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 917.3,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1294202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 917.3,
            "unit": "ns/op",
            "extra": "1294202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1294202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1294202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13681,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "87775 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13681,
            "unit": "ns/op",
            "extra": "87775 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "87775 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "87775 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 36371,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "34230 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 36371,
            "unit": "ns/op",
            "extra": "34230 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "34230 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "34230 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 80823,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 80823,
            "unit": "ns/op",
            "extra": "14544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14544 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14544 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20957897,
            "unit": "ns/op\t12022513 B/op\t  296855 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20957897,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022513,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 21822557,
            "unit": "ns/op\t10563917 B/op\t  205963 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 21822557,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10563917,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205963,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39782651,
            "unit": "ns/op\t42303690 B/op\t  270280 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39782651,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42303690,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270280,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3286,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "328392 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3286,
            "unit": "ns/op",
            "extra": "328392 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "328392 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "328392 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36488,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32960 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36488,
            "unit": "ns/op",
            "extra": "32960 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32960 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32960 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "038fdf7b6d65a216ca6bee22435ccb8a18284bae",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for afb4676095348c91c03477f0126d80352f5bc6b8",
          "timestamp": "2025-03-04T15:57:12Z",
          "url": "https://github.com/moov-io/watchman/commit/038fdf7b6d65a216ca6bee22435ccb8a18284bae"
        },
        "date": 1741179004932,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7713,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "154243 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7713,
            "unit": "ns/op",
            "extra": "154243 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "154243 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "154243 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31148,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38215 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31148,
            "unit": "ns/op",
            "extra": "38215 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38215 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38215 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4835,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "248386 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4835,
            "unit": "ns/op",
            "extra": "248386 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "248386 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "248386 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22908,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "52366 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22908,
            "unit": "ns/op",
            "extra": "52366 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "52366 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "52366 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 885.6,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1354728 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 885.6,
            "unit": "ns/op",
            "extra": "1354728 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1354728 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1354728 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13007,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "90465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13007,
            "unit": "ns/op",
            "extra": "90465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "90465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "90465 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35386,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35386,
            "unit": "ns/op",
            "extra": "33231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 82232,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 82232,
            "unit": "ns/op",
            "extra": "14529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14529 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20542946,
            "unit": "ns/op\t12022725 B/op\t  296855 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20542946,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022725,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 20341648,
            "unit": "ns/op\t10562838 B/op\t  205962 allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 20341648,
            "unit": "ns/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10562838,
            "unit": "B/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205962,
            "unit": "allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 38932659,
            "unit": "ns/op\t42303301 B/op\t  270277 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 38932659,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42303301,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270277,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3581,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "369369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3581,
            "unit": "ns/op",
            "extra": "369369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "369369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "369369 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36323,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33086 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36323,
            "unit": "ns/op",
            "extra": "33086 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33086 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33086 times\n4 procs"
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
          "id": "182231836f8bc95bdab2fd70d5e31d05d6068ec0",
          "message": "chore: fix compile error",
          "timestamp": "2025-03-05T23:28:21Z",
          "url": "https://github.com/moov-io/watchman/commit/182231836f8bc95bdab2fd70d5e31d05d6068ec0"
        },
        "date": 1741265440158,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7606,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "145618 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7606,
            "unit": "ns/op",
            "extra": "145618 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "145618 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "145618 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31657,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38402 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31657,
            "unit": "ns/op",
            "extra": "38402 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38402 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38402 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4730,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "238833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4730,
            "unit": "ns/op",
            "extra": "238833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "238833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "238833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22047,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "54172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22047,
            "unit": "ns/op",
            "extra": "54172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "54172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "54172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 893.2,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1340822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 893.2,
            "unit": "ns/op",
            "extra": "1340822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1340822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1340822 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13069,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "93381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13069,
            "unit": "ns/op",
            "extra": "93381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "93381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "93381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35187,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33026 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35187,
            "unit": "ns/op",
            "extra": "33026 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33026 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33026 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 83152,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 83152,
            "unit": "ns/op",
            "extra": "14694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14694 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21883753,
            "unit": "ns/op\t12022645 B/op\t  296854 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21883753,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022645,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19236684,
            "unit": "ns/op\t10557248 B/op\t  205958 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19236684,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10557248,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205958,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40119455,
            "unit": "ns/op\t42287122 B/op\t  270273 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40119455,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42287122,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270273,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3513,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "380734 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3513,
            "unit": "ns/op",
            "extra": "380734 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "380734 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "380734 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36200,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33309 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36200,
            "unit": "ns/op",
            "extra": "33309 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33309 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33309 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "fafcf27a6289f25f633445a112036fe2b3340007",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 90ea3326e5e644079d516865d1411dcd78cc830c",
          "timestamp": "2025-03-06T12:50:48Z",
          "url": "https://github.com/moov-io/watchman/commit/fafcf27a6289f25f633445a112036fe2b3340007"
        },
        "date": 1741351744666,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7530,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "154549 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7530,
            "unit": "ns/op",
            "extra": "154549 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "154549 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "154549 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30733,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38559 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30733,
            "unit": "ns/op",
            "extra": "38559 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38559 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38559 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4810,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "239254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4810,
            "unit": "ns/op",
            "extra": "239254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "239254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "239254 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22014,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "54454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22014,
            "unit": "ns/op",
            "extra": "54454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "54454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "54454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 908.4,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1316178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 908.4,
            "unit": "ns/op",
            "extra": "1316178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1316178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1316178 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13399,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "90286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13399,
            "unit": "ns/op",
            "extra": "90286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "90286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "90286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35152,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33882 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35152,
            "unit": "ns/op",
            "extra": "33882 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33882 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33882 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 81677,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "13556 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 81677,
            "unit": "ns/op",
            "extra": "13556 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "13556 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "13556 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21399109,
            "unit": "ns/op\t12022512 B/op\t  296855 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21399109,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022512,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19604520,
            "unit": "ns/op\t10562287 B/op\t  205957 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19604520,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10562287,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205957,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 41502486,
            "unit": "ns/op\t42290449 B/op\t  270275 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 41502486,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42290449,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270275,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3445,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "431910 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3445,
            "unit": "ns/op",
            "extra": "431910 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "431910 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "431910 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35698,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "31624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35698,
            "unit": "ns/op",
            "extra": "31624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "31624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "31624 times\n4 procs"
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
          "id": "024b7a7d0de9634596beddccfe246bc90a7c5fec",
          "message": "fix(deps): update opentelemetry-go monorepo to v1.35.0 (#607)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2025-03-08T05:54:36Z",
          "url": "https://github.com/moov-io/watchman/commit/024b7a7d0de9634596beddccfe246bc90a7c5fec"
        },
        "date": 1741437447083,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7447,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "156751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7447,
            "unit": "ns/op",
            "extra": "156751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "156751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "156751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34450,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "39753 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34450,
            "unit": "ns/op",
            "extra": "39753 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "39753 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "39753 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4703,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "233481 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4703,
            "unit": "ns/op",
            "extra": "233481 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "233481 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "233481 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 21958,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "54238 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 21958,
            "unit": "ns/op",
            "extra": "54238 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "54238 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "54238 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 881.1,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1360610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 881.1,
            "unit": "ns/op",
            "extra": "1360610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1360610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1360610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13103,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "92242 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13103,
            "unit": "ns/op",
            "extra": "92242 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "92242 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "92242 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35100,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "34048 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35100,
            "unit": "ns/op",
            "extra": "34048 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "34048 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "34048 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 80860,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14871 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 80860,
            "unit": "ns/op",
            "extra": "14871 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14871 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14871 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20944442,
            "unit": "ns/op\t12022560 B/op\t  296855 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20944442,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022560,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19586312,
            "unit": "ns/op\t10555189 B/op\t  205955 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19586312,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10555189,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205955,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40525153,
            "unit": "ns/op\t42299944 B/op\t  270276 allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40525153,
            "unit": "ns/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42299944,
            "unit": "B/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270276,
            "unit": "allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3318,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "371972 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3318,
            "unit": "ns/op",
            "extra": "371972 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "371972 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "371972 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36012,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36012,
            "unit": "ns/op",
            "extra": "33762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33762 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33762 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "1c85d9b30f55001773455908f5df112b31141e03",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 510169972fa2f26a2183fd7c1a6f26c6cf87dddf",
          "timestamp": "2025-03-08T12:37:35Z",
          "url": "https://github.com/moov-io/watchman/commit/1c85d9b30f55001773455908f5df112b31141e03"
        },
        "date": 1741523775380,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7691,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "150462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7691,
            "unit": "ns/op",
            "extra": "150462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "150462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "150462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30718,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38974 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30718,
            "unit": "ns/op",
            "extra": "38974 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38974 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38974 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 5284,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "236432 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 5284,
            "unit": "ns/op",
            "extra": "236432 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "236432 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "236432 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22139,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "54105 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22139,
            "unit": "ns/op",
            "extra": "54105 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "54105 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "54105 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 874.1,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1362865 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 874.1,
            "unit": "ns/op",
            "extra": "1362865 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1362865 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1362865 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12770,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "93225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12770,
            "unit": "ns/op",
            "extra": "93225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "93225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "93225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35326,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35326,
            "unit": "ns/op",
            "extra": "33793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 81643,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "13527 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 81643,
            "unit": "ns/op",
            "extra": "13527 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "13527 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "13527 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21088135,
            "unit": "ns/op\t12022586 B/op\t  296854 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21088135,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022586,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19424891,
            "unit": "ns/op\t10559073 B/op\t  205957 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19424891,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10559073,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205957,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40425221,
            "unit": "ns/op\t42297253 B/op\t  270278 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40425221,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42297253,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270278,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3599,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "310226 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3599,
            "unit": "ns/op",
            "extra": "310226 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "310226 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "310226 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35667,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33775 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35667,
            "unit": "ns/op",
            "extra": "33775 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33775 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33775 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "91ef081f13bd846a84e5e469071b2c56d13c40f3",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for dc50d2ba7bd6ecbf3d06f70797c39fcc70d612e2",
          "timestamp": "2025-03-09T12:36:22Z",
          "url": "https://github.com/moov-io/watchman/commit/91ef081f13bd846a84e5e469071b2c56d13c40f3"
        },
        "date": 1741611057218,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7387,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "155492 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7387,
            "unit": "ns/op",
            "extra": "155492 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "155492 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "155492 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30613,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30613,
            "unit": "ns/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "39476 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4745,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "234789 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4745,
            "unit": "ns/op",
            "extra": "234789 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "234789 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "234789 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 21968,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 21968,
            "unit": "ns/op",
            "extra": "53722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 923,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1360530 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 923,
            "unit": "ns/op",
            "extra": "1360530 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1360530 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1360530 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12790,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "91950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12790,
            "unit": "ns/op",
            "extra": "91950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "91950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "91950 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35057,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "31918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35057,
            "unit": "ns/op",
            "extra": "31918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "31918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "31918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 81224,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 81224,
            "unit": "ns/op",
            "extra": "14764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14764 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14764 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21075307,
            "unit": "ns/op\t12022850 B/op\t  296855 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21075307,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022850,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18951423,
            "unit": "ns/op\t10557113 B/op\t  205955 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18951423,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10557113,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205955,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39813952,
            "unit": "ns/op\t42292801 B/op\t  270279 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39813952,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42292801,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270279,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3507,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "313531 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3507,
            "unit": "ns/op",
            "extra": "313531 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "313531 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "313531 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35962,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33427 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35962,
            "unit": "ns/op",
            "extra": "33427 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33427 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33427 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "9b8fdcdff2746b7bdcc9253b458971d103546141",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 4b1a9a05db06fb7c56aeeab81abe89a2bae0d565",
          "timestamp": "2025-03-10T12:51:03Z",
          "url": "https://github.com/moov-io/watchman/commit/9b8fdcdff2746b7bdcc9253b458971d103546141"
        },
        "date": 1741697479336,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7568,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "155529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7568,
            "unit": "ns/op",
            "extra": "155529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "155529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "155529 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30623,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "38862 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30623,
            "unit": "ns/op",
            "extra": "38862 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "38862 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "38862 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4713,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "237126 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4713,
            "unit": "ns/op",
            "extra": "237126 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "237126 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "237126 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 21940,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 21940,
            "unit": "ns/op",
            "extra": "53820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 905.7,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1302585 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 905.7,
            "unit": "ns/op",
            "extra": "1302585 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1302585 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1302585 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13159,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "91180 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13159,
            "unit": "ns/op",
            "extra": "91180 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "91180 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "91180 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 35220,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "33825 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 35220,
            "unit": "ns/op",
            "extra": "33825 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "33825 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "33825 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 84499,
            "unit": "ns/op\t   26790 B/op\t     679 allocs/op",
            "extra": "14804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 84499,
            "unit": "ns/op",
            "extra": "14804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26790,
            "unit": "B/op",
            "extra": "14804 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14804 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 20867530,
            "unit": "ns/op\t12022484 B/op\t  296855 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 20867530,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022484,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296855,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19100499,
            "unit": "ns/op\t10565612 B/op\t  205958 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19100499,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10565612,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205958,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 48073326,
            "unit": "ns/op\t42279383 B/op\t  270274 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 48073326,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42279383,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270274,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3545,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "348759 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3545,
            "unit": "ns/op",
            "extra": "348759 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "348759 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "348759 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35939,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33421 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35939,
            "unit": "ns/op",
            "extra": "33421 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33421 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33421 times\n4 procs"
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
          "id": "b39313c00d412e3de5249019e6d140b498aa790e",
          "message": "build: update codeql action",
          "timestamp": "2025-03-11T19:34:42Z",
          "url": "https://github.com/moov-io/watchman/commit/b39313c00d412e3de5249019e6d140b498aa790e"
        },
        "date": 1741783906076,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7475,
            "unit": "ns/op\t    1808 B/op\t     113 allocs/op",
            "extra": "155877 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7475,
            "unit": "ns/op",
            "extra": "155877 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 1808,
            "unit": "B/op",
            "extra": "155877 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 113,
            "unit": "allocs/op",
            "extra": "155877 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34647,
            "unit": "ns/op\t   10907 B/op\t     169 allocs/op",
            "extra": "37611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34647,
            "unit": "ns/op",
            "extra": "37611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 10907,
            "unit": "B/op",
            "extra": "37611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 169,
            "unit": "allocs/op",
            "extra": "37611 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 4808,
            "unit": "ns/op\t    1362 B/op\t      64 allocs/op",
            "extra": "237058 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 4808,
            "unit": "ns/op",
            "extra": "237058 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 1362,
            "unit": "B/op",
            "extra": "237058 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 64,
            "unit": "allocs/op",
            "extra": "237058 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 22341,
            "unit": "ns/op\t    8164 B/op\t      99 allocs/op",
            "extra": "53676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 22341,
            "unit": "ns/op",
            "extra": "53676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 8164,
            "unit": "B/op",
            "extra": "53676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "53676 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 920.6,
            "unit": "ns/op\t     474 B/op\t       4 allocs/op",
            "extra": "1370672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 920.6,
            "unit": "ns/op",
            "extra": "1370672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 474,
            "unit": "B/op",
            "extra": "1370672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "1370672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 12836,
            "unit": "ns/op\t    5027 B/op\t      22 allocs/op",
            "extra": "93018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 12836,
            "unit": "ns/op",
            "extra": "93018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5027,
            "unit": "B/op",
            "extra": "93018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "93018 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 34515,
            "unit": "ns/op\t    8492 B/op\t     543 allocs/op",
            "extra": "34501 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 34515,
            "unit": "ns/op",
            "extra": "34501 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 8492,
            "unit": "B/op",
            "extra": "34501 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 543,
            "unit": "allocs/op",
            "extra": "34501 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 80957,
            "unit": "ns/op\t   26789 B/op\t     679 allocs/op",
            "extra": "14768 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 80957,
            "unit": "ns/op",
            "extra": "14768 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 26789,
            "unit": "B/op",
            "extra": "14768 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 679,
            "unit": "allocs/op",
            "extra": "14768 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21886082,
            "unit": "ns/op\t12022498 B/op\t  296854 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21886082,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 12022498,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 296854,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 21145358,
            "unit": "ns/op\t10558316 B/op\t  205962 allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 21145358,
            "unit": "ns/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 10558316,
            "unit": "B/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 205962,
            "unit": "allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 40075239,
            "unit": "ns/op\t42294900 B/op\t  270277 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 40075239,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 42294900,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 270277,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3615,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "305516 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3615,
            "unit": "ns/op",
            "extra": "305516 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "305516 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "305516 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35868,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33529 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35868,
            "unit": "ns/op",
            "extra": "33529 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33529 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33529 times\n4 procs"
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
          "id": "c689db731a3da40f896e4dee3fc1f41d6a533770",
          "message": "meta: gofmt",
          "timestamp": "2025-03-12T22:13:23Z",
          "url": "https://github.com/moov-io/watchman/commit/c689db731a3da40f896e4dee3fc1f41d6a533770"
        },
        "date": 1741870385082,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8769,
            "unit": "ns/op\t    2216 B/op\t     126 allocs/op",
            "extra": "132373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8769,
            "unit": "ns/op",
            "extra": "132373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2216,
            "unit": "B/op",
            "extra": "132373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 126,
            "unit": "allocs/op",
            "extra": "132373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34080,
            "unit": "ns/op\t   11317 B/op\t     182 allocs/op",
            "extra": "35082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34080,
            "unit": "ns/op",
            "extra": "35082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11317,
            "unit": "B/op",
            "extra": "35082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 182,
            "unit": "allocs/op",
            "extra": "35082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 10913,
            "unit": "ns/op\t    3228 B/op\t     136 allocs/op",
            "extra": "106056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 10913,
            "unit": "ns/op",
            "extra": "106056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3228,
            "unit": "B/op",
            "extra": "106056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 136,
            "unit": "allocs/op",
            "extra": "106056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30372,
            "unit": "ns/op\t   10029 B/op\t     172 allocs/op",
            "extra": "39796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30372,
            "unit": "ns/op",
            "extra": "39796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10029,
            "unit": "B/op",
            "extra": "39796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 172,
            "unit": "allocs/op",
            "extra": "39796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1446,
            "unit": "ns/op\t     672 B/op\t      13 allocs/op",
            "extra": "753390 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1446,
            "unit": "ns/op",
            "extra": "753390 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "753390 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 13,
            "unit": "allocs/op",
            "extra": "753390 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14515,
            "unit": "ns/op\t    5226 B/op\t      31 allocs/op",
            "extra": "81811 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14515,
            "unit": "ns/op",
            "extra": "81811 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5226,
            "unit": "B/op",
            "extra": "81811 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "81811 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 57410,
            "unit": "ns/op\t   15372 B/op\t     761 allocs/op",
            "extra": "20736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 57410,
            "unit": "ns/op",
            "extra": "20736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 15372,
            "unit": "B/op",
            "extra": "20736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 761,
            "unit": "allocs/op",
            "extra": "20736 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 109549,
            "unit": "ns/op\t   33681 B/op\t     897 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 109549,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 33681,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 897,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 65344011,
            "unit": "ns/op\t34045837 B/op\t 1253241 allocs/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 65344011,
            "unit": "ns/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 34045837,
            "unit": "B/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1253241,
            "unit": "allocs/op",
            "extra": "19 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 33777355,
            "unit": "ns/op\t19997882 B/op\t  520305 allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 33777355,
            "unit": "ns/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19997882,
            "unit": "B/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 520305,
            "unit": "allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 55604367,
            "unit": "ns/op\t51733784 B/op\t  584789 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 55604367,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 51733784,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 584789,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3451,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "348555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3451,
            "unit": "ns/op",
            "extra": "348555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "348555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "348555 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5534,
            "unit": "ns/op\t     618 B/op\t      24 allocs/op",
            "extra": "216066 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5534,
            "unit": "ns/op",
            "extra": "216066 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 618,
            "unit": "B/op",
            "extra": "216066 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "216066 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36556,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36556,
            "unit": "ns/op",
            "extra": "32946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32946 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32946 times\n4 procs"
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
          "id": "ad64a546c938b4c65f5c5e10ec9b404cc88ce462",
          "message": "stringscore: reduce memory allocations in GenerateWordCombinations",
          "timestamp": "2025-03-13T14:20:57Z",
          "url": "https://github.com/moov-io/watchman/commit/ad64a546c938b4c65f5c5e10ec9b404cc88ce462"
        },
        "date": 1741876023679,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8054,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "139532 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8054,
            "unit": "ns/op",
            "extra": "139532 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "139532 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "139532 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34037,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "36712 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34037,
            "unit": "ns/op",
            "extra": "36712 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "36712 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "36712 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11531,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "101094 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11531,
            "unit": "ns/op",
            "extra": "101094 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "101094 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "101094 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30621,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "38380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30621,
            "unit": "ns/op",
            "extra": "38380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "38380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "38380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1401,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "811141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1401,
            "unit": "ns/op",
            "extra": "811141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "811141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "811141 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14462,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "82820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14462,
            "unit": "ns/op",
            "extra": "82820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "82820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "82820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 45091,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26497 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 45091,
            "unit": "ns/op",
            "extra": "26497 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26497 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26497 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 95914,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 95914,
            "unit": "ns/op",
            "extra": "12334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12334 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 64634271,
            "unit": "ns/op\t32930491 B/op\t 1121402 allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 64634271,
            "unit": "ns/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32930491,
            "unit": "B/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121402,
            "unit": "allocs/op",
            "extra": "16 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31106423,
            "unit": "ns/op\t19168934 B/op\t  462244 allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31106423,
            "unit": "ns/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19168934,
            "unit": "B/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462244,
            "unit": "allocs/op",
            "extra": "39 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 53341362,
            "unit": "ns/op\t50903457 B/op\t  526727 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 53341362,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50903457,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526727,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3441,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "302575 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3441,
            "unit": "ns/op",
            "extra": "302575 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "302575 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "302575 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4998,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "246702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4998,
            "unit": "ns/op",
            "extra": "246702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "246702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "246702 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36522,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32985 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36522,
            "unit": "ns/op",
            "extra": "32985 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32985 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32985 times\n4 procs"
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
          "id": "c55e67d84d0d5c479beec02a1623a0b060bb03b6",
          "message": "build: downgrade golang.org/x/net",
          "timestamp": "2025-03-13T21:24:56Z",
          "url": "https://github.com/moov-io/watchman/commit/c55e67d84d0d5c479beec02a1623a0b060bb03b6"
        },
        "date": 1741901370122,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8094,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "142090 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8094,
            "unit": "ns/op",
            "extra": "142090 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "142090 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "142090 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32591,
            "unit": "ns/op\t   11285 B/op\t     175 allocs/op",
            "extra": "36448 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32591,
            "unit": "ns/op",
            "extra": "36448 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11285,
            "unit": "B/op",
            "extra": "36448 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "36448 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11248,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "101827 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11248,
            "unit": "ns/op",
            "extra": "101827 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "101827 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "101827 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30089,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30089,
            "unit": "ns/op",
            "extra": "39538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39538 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1428,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "809200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1428,
            "unit": "ns/op",
            "extra": "809200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "809200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "809200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14621,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "79627 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14621,
            "unit": "ns/op",
            "extra": "79627 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "79627 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "79627 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44833,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26167 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44833,
            "unit": "ns/op",
            "extra": "26167 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26167 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26167 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 95651,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12484 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 95651,
            "unit": "ns/op",
            "extra": "12484 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12484 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12484 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53553800,
            "unit": "ns/op\t32930948 B/op\t 1121403 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53553800,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32930948,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121403,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29568077,
            "unit": "ns/op\t19168360 B/op\t  462218 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29568077,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19168360,
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
            "value": 51739028,
            "unit": "ns/op\t50909503 B/op\t  526708 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 51739028,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50909503,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526708,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30778387,
            "unit": "ns/op\t19162633 B/op\t  462192 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30778387,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19162633,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462192,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 32461327,
            "unit": "ns/op\t19171524 B/op\t  462198 allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 32461327,
            "unit": "ns/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19171524,
            "unit": "B/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462198,
            "unit": "allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30216398,
            "unit": "ns/op\t19169239 B/op\t  462194 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30216398,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19169239,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462194,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3513,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "329221 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3513,
            "unit": "ns/op",
            "extra": "329221 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "329221 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "329221 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4782,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "224618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4782,
            "unit": "ns/op",
            "extra": "224618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "224618 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "224618 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36450,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32907 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36450,
            "unit": "ns/op",
            "extra": "32907 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32907 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32907 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "42e79b572136278e544501e867214589fc3e37f0",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 5146d0e008a4c43faf0ddae5f0a2fff789d4b576",
          "timestamp": "2025-03-13T21:29:37Z",
          "url": "https://github.com/moov-io/watchman/commit/42e79b572136278e544501e867214589fc3e37f0"
        },
        "date": 1741956611296,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8619,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "138616 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8619,
            "unit": "ns/op",
            "extra": "138616 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "138616 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "138616 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33445,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "35922 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33445,
            "unit": "ns/op",
            "extra": "35922 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "35922 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "35922 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11573,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "100381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11573,
            "unit": "ns/op",
            "extra": "100381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "100381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "100381 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30968,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "37960 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30968,
            "unit": "ns/op",
            "extra": "37960 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "37960 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "37960 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1384,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "868024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1384,
            "unit": "ns/op",
            "extra": "868024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "868024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "868024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14558,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "84300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14558,
            "unit": "ns/op",
            "extra": "84300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "84300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "84300 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44706,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26467 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44706,
            "unit": "ns/op",
            "extra": "26467 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26467 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26467 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 95418,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 95418,
            "unit": "ns/op",
            "extra": "12572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12572 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12572 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53172083,
            "unit": "ns/op\t32932422 B/op\t 1121460 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53172083,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32932422,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121460,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29396229,
            "unit": "ns/op\t19170464 B/op\t  462226 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29396229,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19170464,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462226,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 50554874,
            "unit": "ns/op\t50899284 B/op\t  526711 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 50554874,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50899284,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526711,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 28687403,
            "unit": "ns/op\t19172692 B/op\t  462202 allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 28687403,
            "unit": "ns/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19172692,
            "unit": "B/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 28660652,
            "unit": "ns/op\t19166777 B/op\t  462208 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 28660652,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19166777,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462208,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 28646076,
            "unit": "ns/op\t19176218 B/op\t  462211 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 28646076,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19176218,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462211,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3506,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "348272 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3506,
            "unit": "ns/op",
            "extra": "348272 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "348272 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "348272 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4913,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "236230 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4913,
            "unit": "ns/op",
            "extra": "236230 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "236230 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "236230 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36157,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33084 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36157,
            "unit": "ns/op",
            "extra": "33084 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33084 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33084 times\n4 procs"
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
          "id": "93a8a7eb61d71bd74bbf9af86f5a86fc670291fb",
          "message": "ui: PoC to prefill values from query params",
          "timestamp": "2025-03-14T22:37:15Z",
          "url": "https://github.com/moov-io/watchman/commit/93a8a7eb61d71bd74bbf9af86f5a86fc670291fb"
        },
        "date": 1742042813529,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7994,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "138655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7994,
            "unit": "ns/op",
            "extra": "138655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "138655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "138655 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32797,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "35454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32797,
            "unit": "ns/op",
            "extra": "35454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "35454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "35454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11269,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "101826 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11269,
            "unit": "ns/op",
            "extra": "101826 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "101826 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "101826 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30438,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30438,
            "unit": "ns/op",
            "extra": "39543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1391,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "816824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1391,
            "unit": "ns/op",
            "extra": "816824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "816824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "816824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14350,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "84024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14350,
            "unit": "ns/op",
            "extra": "84024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "84024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "84024 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44851,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26491 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44851,
            "unit": "ns/op",
            "extra": "26491 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26491 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26491 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 97862,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12500 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 97862,
            "unit": "ns/op",
            "extra": "12500 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12500 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12500 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 52530209,
            "unit": "ns/op\t32931258 B/op\t 1121439 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 52530209,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931258,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121439,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29097188,
            "unit": "ns/op\t19164073 B/op\t  462215 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29097188,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19164073,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462215,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 51159093,
            "unit": "ns/op\t50887758 B/op\t  526699 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 51159093,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50887758,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526699,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 32552646,
            "unit": "ns/op\t19164629 B/op\t  462193 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 32552646,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19164629,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462193,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 28782606,
            "unit": "ns/op\t19170157 B/op\t  462199 allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 28782606,
            "unit": "ns/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19170157,
            "unit": "B/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462199,
            "unit": "allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 28681472,
            "unit": "ns/op\t19167505 B/op\t  462202 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 28681472,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19167505,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462202,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3433,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "349437 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3433,
            "unit": "ns/op",
            "extra": "349437 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "349437 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "349437 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4867,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "250546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4867,
            "unit": "ns/op",
            "extra": "250546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "250546 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "250546 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36625,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33142 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36625,
            "unit": "ns/op",
            "extra": "33142 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33142 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33142 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "4a76a43d443be2c22abbe2120d0d374cecb28220",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 7da7753c986b2ac96bc8936a4e297d54d0edb0e0",
          "timestamp": "2025-03-15T12:47:01Z",
          "url": "https://github.com/moov-io/watchman/commit/4a76a43d443be2c22abbe2120d0d374cecb28220"
        },
        "date": 1742129222537,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8108,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "138454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8108,
            "unit": "ns/op",
            "extra": "138454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "138454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "138454 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33293,
            "unit": "ns/op\t   11285 B/op\t     175 allocs/op",
            "extra": "35619 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33293,
            "unit": "ns/op",
            "extra": "35619 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11285,
            "unit": "B/op",
            "extra": "35619 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "35619 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11382,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "98953 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11382,
            "unit": "ns/op",
            "extra": "98953 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "98953 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "98953 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30579,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30579,
            "unit": "ns/op",
            "extra": "39064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39064 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1413,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "805442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1413,
            "unit": "ns/op",
            "extra": "805442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "805442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "805442 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14569,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "82065 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14569,
            "unit": "ns/op",
            "extra": "82065 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "82065 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "82065 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 45345,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26312 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 45345,
            "unit": "ns/op",
            "extra": "26312 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26312 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26312 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 97253,
            "unit": "ns/op\t   30448 B/op\t     758 allocs/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 97253,
            "unit": "ns/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30448,
            "unit": "B/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53179435,
            "unit": "ns/op\t32931757 B/op\t 1121471 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53179435,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931757,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121471,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 31797744,
            "unit": "ns/op\t19173779 B/op\t  462235 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 31797744,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19173779,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462235,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 57854555,
            "unit": "ns/op\t50890487 B/op\t  526715 allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 57854555,
            "unit": "ns/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50890487,
            "unit": "B/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526715,
            "unit": "allocs/op",
            "extra": "18 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30140245,
            "unit": "ns/op\t19169407 B/op\t  462211 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30140245,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19169407,
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
            "value": 30466640,
            "unit": "ns/op\t19158170 B/op\t  462209 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30466640,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19158170,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462209,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30615236,
            "unit": "ns/op\t19160387 B/op\t  462207 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30615236,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19160387,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462207,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3511,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "375687 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3511,
            "unit": "ns/op",
            "extra": "375687 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "375687 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "375687 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5023,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "238041 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5023,
            "unit": "ns/op",
            "extra": "238041 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "238041 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "238041 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36474,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33194 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36474,
            "unit": "ns/op",
            "extra": "33194 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33194 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33194 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "9d522218ee5257187e94bb982130c8d65719ad42",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 89235f3ec68cc53c619394adc7ae47ad74160b16",
          "timestamp": "2025-03-16T12:47:09Z",
          "url": "https://github.com/moov-io/watchman/commit/9d522218ee5257187e94bb982130c8d65719ad42"
        },
        "date": 1742216032638,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7983,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "139051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7983,
            "unit": "ns/op",
            "extra": "139051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "139051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "139051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32346,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "37020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32346,
            "unit": "ns/op",
            "extra": "37020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "37020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "37020 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11313,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "101505 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11313,
            "unit": "ns/op",
            "extra": "101505 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "101505 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "101505 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30285,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "39540 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30285,
            "unit": "ns/op",
            "extra": "39540 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "39540 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "39540 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1389,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "770047 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1389,
            "unit": "ns/op",
            "extra": "770047 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "770047 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "770047 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14260,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "82724 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14260,
            "unit": "ns/op",
            "extra": "82724 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "82724 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "82724 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44389,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44389,
            "unit": "ns/op",
            "extra": "26672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 94477,
            "unit": "ns/op\t   30446 B/op\t     758 allocs/op",
            "extra": "12666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 94477,
            "unit": "ns/op",
            "extra": "12666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30446,
            "unit": "B/op",
            "extra": "12666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12666 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 51466968,
            "unit": "ns/op\t32931383 B/op\t 1121424 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 51466968,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931383,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121424,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 28914886,
            "unit": "ns/op\t19167558 B/op\t  462241 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 28914886,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19167558,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462241,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 49424585,
            "unit": "ns/op\t50904976 B/op\t  526727 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 49424585,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50904976,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526727,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 28290400,
            "unit": "ns/op\t19168692 B/op\t  462217 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 28290400,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19168692,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462217,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 28315708,
            "unit": "ns/op\t19172530 B/op\t  462224 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 28315708,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19172530,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462224,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 28157964,
            "unit": "ns/op\t19172184 B/op\t  462223 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 28157964,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19172184,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462223,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3444,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "324607 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3444,
            "unit": "ns/op",
            "extra": "324607 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "324607 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "324607 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4867,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "230014 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4867,
            "unit": "ns/op",
            "extra": "230014 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "230014 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "230014 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36313,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "33118 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36313,
            "unit": "ns/op",
            "extra": "33118 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "33118 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "33118 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "committer": {
            "name": "github-action-benchmark",
            "username": "github",
            "email": "github@users.noreply.github.com"
          },
          "id": "7edfb454f08d6ff21e9d6bce627ce3e251b5980a",
          "message": "add moov-io/watchman OFAC Benchmarks (go) benchmark result for 8a8c89cd74d8c5bb11b441fecceaa3597e8ee137",
          "timestamp": "2025-03-17T12:53:59Z",
          "url": "https://github.com/moov-io/watchman/commit/7edfb454f08d6ff21e9d6bce627ce3e251b5980a"
        },
        "date": 1742302416334,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 8122,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "139868 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 8122,
            "unit": "ns/op",
            "extra": "139868 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "139868 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "139868 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33502,
            "unit": "ns/op\t   11286 B/op\t     175 allocs/op",
            "extra": "35796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33502,
            "unit": "ns/op",
            "extra": "35796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11286,
            "unit": "B/op",
            "extra": "35796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "35796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11575,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "97772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11575,
            "unit": "ns/op",
            "extra": "97772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "97772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "97772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 31548,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "38496 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 31548,
            "unit": "ns/op",
            "extra": "38496 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "38496 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "38496 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1445,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "797864 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1445,
            "unit": "ns/op",
            "extra": "797864 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "797864 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "797864 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14758,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81319 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14758,
            "unit": "ns/op",
            "extra": "81319 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81319 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81319 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 45418,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26214 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 45418,
            "unit": "ns/op",
            "extra": "26214 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26214 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26214 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 96965,
            "unit": "ns/op\t   30447 B/op\t     758 allocs/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 96965,
            "unit": "ns/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30447,
            "unit": "B/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12373 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 53346771,
            "unit": "ns/op\t32931984 B/op\t 1121472 allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 53346771,
            "unit": "ns/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931984,
            "unit": "B/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121472,
            "unit": "allocs/op",
            "extra": "20 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 30059975,
            "unit": "ns/op\t19171425 B/op\t  462253 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 30059975,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19171425,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462253,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 52657178,
            "unit": "ns/op\t50892639 B/op\t  526739 allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 52657178,
            "unit": "ns/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50892639,
            "unit": "B/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526739,
            "unit": "allocs/op",
            "extra": "25 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30958261,
            "unit": "ns/op\t19161550 B/op\t  462227 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30958261,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19161550,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462227,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30911802,
            "unit": "ns/op\t19161324 B/op\t  462227 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30911802,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19161324,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462227,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 30944436,
            "unit": "ns/op\t19164797 B/op\t  462231 allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 30944436,
            "unit": "ns/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19164797,
            "unit": "B/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462231,
            "unit": "allocs/op",
            "extra": "40 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3469,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "316978 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3469,
            "unit": "ns/op",
            "extra": "316978 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "316978 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "316978 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4963,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "246626 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4963,
            "unit": "ns/op",
            "extra": "246626 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "246626 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "246626 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36684,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32598 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36684,
            "unit": "ns/op",
            "extra": "32598 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32598 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32598 times\n4 procs"
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
          "id": "1da0488957ab287f18e61f3eae05831f577e4689",
          "message": "docs: mention a couple makefile commands\n\nFixes: https://github.com/moov-io/watchman/issues/504",
          "timestamp": "2025-03-18T19:15:41Z",
          "url": "https://github.com/moov-io/watchman/commit/1da0488957ab287f18e61f3eae05831f577e4689"
        },
        "date": 1742388782474,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 7982,
            "unit": "ns/op\t    2184 B/op\t     119 allocs/op",
            "extra": "138700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 7982,
            "unit": "ns/op",
            "extra": "138700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 2184,
            "unit": "B/op",
            "extra": "138700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 119,
            "unit": "allocs/op",
            "extra": "138700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32714,
            "unit": "ns/op\t   11285 B/op\t     175 allocs/op",
            "extra": "36819 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32714,
            "unit": "ns/op",
            "extra": "36819 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 11285,
            "unit": "B/op",
            "extra": "36819 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 175,
            "unit": "allocs/op",
            "extra": "36819 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 11445,
            "unit": "ns/op\t    3258 B/op\t     141 allocs/op",
            "extra": "99769 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 11445,
            "unit": "ns/op",
            "extra": "99769 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3258,
            "unit": "B/op",
            "extra": "99769 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "99769 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 30794,
            "unit": "ns/op\t   10059 B/op\t     177 allocs/op",
            "extra": "38968 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 30794,
            "unit": "ns/op",
            "extra": "38968 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10059,
            "unit": "B/op",
            "extra": "38968 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 177,
            "unit": "allocs/op",
            "extra": "38968 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1464,
            "unit": "ns/op\t     680 B/op\t      12 allocs/op",
            "extra": "721460 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1464,
            "unit": "ns/op",
            "extra": "721460 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 680,
            "unit": "B/op",
            "extra": "721460 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "721460 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14656,
            "unit": "ns/op\t    5234 B/op\t      30 allocs/op",
            "extra": "81740 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14656,
            "unit": "ns/op",
            "extra": "81740 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5234,
            "unit": "B/op",
            "extra": "81740 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "81740 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 44469,
            "unit": "ns/op\t   12132 B/op\t     622 allocs/op",
            "extra": "26734 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 44469,
            "unit": "ns/op",
            "extra": "26734 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 12132,
            "unit": "B/op",
            "extra": "26734 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 622,
            "unit": "allocs/op",
            "extra": "26734 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 94846,
            "unit": "ns/op\t   30446 B/op\t     758 allocs/op",
            "extra": "12714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 94846,
            "unit": "ns/op",
            "extra": "12714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 30446,
            "unit": "B/op",
            "extra": "12714 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 758,
            "unit": "allocs/op",
            "extra": "12714 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 52551497,
            "unit": "ns/op\t32931067 B/op\t 1121433 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 52551497,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 32931067,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 1121433,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 29087535,
            "unit": "ns/op\t19167761 B/op\t  462235 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 29087535,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 19167761,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 462235,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 55278948,
            "unit": "ns/op\t50900164 B/op\t  526727 allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 55278948,
            "unit": "ns/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 50900164,
            "unit": "B/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 526727,
            "unit": "allocs/op",
            "extra": "24 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 30615052,
            "unit": "ns/op\t19171227 B/op\t  462213 allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 30615052,
            "unit": "ns/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 19171227,
            "unit": "B/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 462213,
            "unit": "allocs/op",
            "extra": "38 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 30380432,
            "unit": "ns/op\t19165406 B/op\t  462215 allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 30380432,
            "unit": "ns/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 19165406,
            "unit": "B/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 462215,
            "unit": "allocs/op",
            "extra": "44 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 29717804,
            "unit": "ns/op\t19163156 B/op\t  462212 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 29717804,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 19163156,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 462212,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3538,
            "unit": "ns/op\t     185 B/op\t       9 allocs/op",
            "extra": "348452 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3538,
            "unit": "ns/op",
            "extra": "348452 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 185,
            "unit": "B/op",
            "extra": "348452 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 9,
            "unit": "allocs/op",
            "extra": "348452 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4891,
            "unit": "ns/op\t     565 B/op\t      20 allocs/op",
            "extra": "229803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4891,
            "unit": "ns/op",
            "extra": "229803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 565,
            "unit": "B/op",
            "extra": "229803 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "229803 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36614,
            "unit": "ns/op\t      32 B/op\t       2 allocs/op",
            "extra": "32850 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36614,
            "unit": "ns/op",
            "extra": "32850 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 32,
            "unit": "B/op",
            "extra": "32850 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "32850 times\n4 procs"
          }
        ]
      }
    ]
  }
}