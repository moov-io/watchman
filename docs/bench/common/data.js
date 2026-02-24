window.BENCHMARK_DATA = {
  "lastUpdate": 1771939789540,
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
          "id": "c7a2122b386fa4336cfd03986ba53cc9ac134358",
          "message": "Merge pull request #658 from adamdecaf/fixup-ingest-dedup\n\ningest,search: dedup phone contact info",
          "timestamp": "2025-08-21T21:15:28Z",
          "url": "https://github.com/moov-io/watchman/commit/c7a2122b386fa4336cfd03986ba53cc9ac134358"
        },
        "date": 1755867441796,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9504,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9504,
            "unit": "ns/op",
            "extra": "121305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121305 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33039,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "36672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33039,
            "unit": "ns/op",
            "extra": "36672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "36672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36672 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15093,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "75694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15093,
            "unit": "ns/op",
            "extra": "75694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "75694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "75694 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33337,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33337,
            "unit": "ns/op",
            "extra": "35644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1393,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "891415 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1393,
            "unit": "ns/op",
            "extra": "891415 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "891415 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "891415 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14472,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "83156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14472,
            "unit": "ns/op",
            "extra": "83156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "83156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "83156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78815,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15003 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78815,
            "unit": "ns/op",
            "extra": "15003 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15003 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15003 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 127857,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8683 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 127857,
            "unit": "ns/op",
            "extra": "8683 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8683 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8683 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22952113,
            "unit": "ns/op\t16034437 B/op\t  248427 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22952113,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034437,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248427,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17463059,
            "unit": "ns/op\t12617549 B/op\t  144516 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17463059,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12617549,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144516,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36467614,
            "unit": "ns/op\t44346657 B/op\t  209004 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36467614,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44346657,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209004,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17344278,
            "unit": "ns/op\t12610052 B/op\t  144498 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17344278,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12610052,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144498,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17291156,
            "unit": "ns/op\t12609936 B/op\t  144496 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17291156,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12609936,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144496,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17429241,
            "unit": "ns/op\t12613953 B/op\t  144499 allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17429241,
            "unit": "ns/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12613953,
            "unit": "B/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3612,
            "unit": "ns/op\t     147 B/op\t       4 allocs/op",
            "extra": "299088 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3612,
            "unit": "ns/op",
            "extra": "299088 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 147,
            "unit": "B/op",
            "extra": "299088 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "299088 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4964,
            "unit": "ns/op\t     512 B/op\t      11 allocs/op",
            "extra": "225938 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4964,
            "unit": "ns/op",
            "extra": "225938 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "225938 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "225938 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35987,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33337 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35987,
            "unit": "ns/op",
            "extra": "33337 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33337 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33337 times\n4 procs"
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
          "id": "dbcb862473cd2354b2bb75ba8948b5d1bd50a2e8",
          "message": "Merge pull request #659 from moov-io/dependabot/bundler/docs/nokogiri-1.18.9\n\nbuild(deps-dev): bump nokogiri from 1.18.8 to 1.18.9 in /docs",
          "timestamp": "2025-08-25T14:09:40Z",
          "url": "https://github.com/moov-io/watchman/commit/dbcb862473cd2354b2bb75ba8948b5d1bd50a2e8"
        },
        "date": 1756213254807,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10002,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "114799 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10002,
            "unit": "ns/op",
            "extra": "114799 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "114799 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "114799 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34390,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "34856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34390,
            "unit": "ns/op",
            "extra": "34856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "34856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "34856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16088,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "73812 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16088,
            "unit": "ns/op",
            "extra": "73812 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "73812 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "73812 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 35071,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "33939 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 35071,
            "unit": "ns/op",
            "extra": "33939 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "33939 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "33939 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1435,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "886590 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1435,
            "unit": "ns/op",
            "extra": "886590 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "886590 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "886590 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14600,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "81302 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14600,
            "unit": "ns/op",
            "extra": "81302 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "81302 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "81302 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 82452,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "14493 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 82452,
            "unit": "ns/op",
            "extra": "14493 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "14493 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "14493 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 134238,
            "unit": "ns/op\t   38560 B/op\t     599 allocs/op",
            "extra": "8440 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 134238,
            "unit": "ns/op",
            "extra": "8440 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38560,
            "unit": "B/op",
            "extra": "8440 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8440 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 27797917,
            "unit": "ns/op\t16034214 B/op\t  248426 allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 27797917,
            "unit": "ns/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034214,
            "unit": "B/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248426,
            "unit": "allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18060822,
            "unit": "ns/op\t12617959 B/op\t  144520 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18060822,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12617959,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144520,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37812587,
            "unit": "ns/op\t44347603 B/op\t  209009 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37812587,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44347603,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209009,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17952904,
            "unit": "ns/op\t12617990 B/op\t  144504 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17952904,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12617990,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144504,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17832648,
            "unit": "ns/op\t12613789 B/op\t  144504 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17832648,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12613789,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144504,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18048050,
            "unit": "ns/op\t12611441 B/op\t  144503 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18048050,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12611441,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144503,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 3529,
            "unit": "ns/op\t     147 B/op\t       4 allocs/op",
            "extra": "314686 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 3529,
            "unit": "ns/op",
            "extra": "314686 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 147,
            "unit": "B/op",
            "extra": "314686 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "314686 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 4945,
            "unit": "ns/op\t     512 B/op\t      11 allocs/op",
            "extra": "242767 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 4945,
            "unit": "ns/op",
            "extra": "242767 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 512,
            "unit": "B/op",
            "extra": "242767 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 11,
            "unit": "allocs/op",
            "extra": "242767 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36065,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33110 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36065,
            "unit": "ns/op",
            "extra": "33110 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33110 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33110 times\n4 procs"
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
          "id": "3c349860cb0b0a433970dba5f788b68d1af191d4",
          "message": "Merge pull request #663 from moov-io/build-tini\n\nbuild: wrap process in tini",
          "timestamp": "2025-09-22T20:04:01Z",
          "url": "https://github.com/moov-io/watchman/commit/3c349860cb0b0a433970dba5f788b68d1af191d4"
        },
        "date": 1758632194151,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9493,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "122011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9493,
            "unit": "ns/op",
            "extra": "122011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "122011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "122011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32396,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "37080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32396,
            "unit": "ns/op",
            "extra": "37080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "37080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37080 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15077,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76363 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15077,
            "unit": "ns/op",
            "extra": "76363 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76363 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76363 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32986,
            "unit": "ns/op\t   10715 B/op\t     121 allocs/op",
            "extra": "36495 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32986,
            "unit": "ns/op",
            "extra": "36495 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10715,
            "unit": "B/op",
            "extra": "36495 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36495 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1296,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "914391 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1296,
            "unit": "ns/op",
            "extra": "914391 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "914391 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "914391 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13740,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "87092 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13740,
            "unit": "ns/op",
            "extra": "87092 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "87092 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "87092 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 79027,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 79027,
            "unit": "ns/op",
            "extra": "15235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15235 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 127919,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8570 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 127919,
            "unit": "ns/op",
            "extra": "8570 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8570 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8570 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 24350957,
            "unit": "ns/op\t16034762 B/op\t  248438 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 24350957,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034762,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248438,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17247191,
            "unit": "ns/op\t12613256 B/op\t  144515 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17247191,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12613256,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144515,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37296797,
            "unit": "ns/op\t44350178 B/op\t  209002 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37296797,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44350178,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209002,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17478711,
            "unit": "ns/op\t12615786 B/op\t  144499 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17478711,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12615786,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17513797,
            "unit": "ns/op\t12614092 B/op\t  144499 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17513797,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12614092,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17660393,
            "unit": "ns/op\t12611903 B/op\t  144499 allocs/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17660393,
            "unit": "ns/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12611903,
            "unit": "B/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4467,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "256269 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4467,
            "unit": "ns/op",
            "extra": "256269 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "256269 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "256269 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5928,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "176955 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5928,
            "unit": "ns/op",
            "extra": "176955 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "176955 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "176955 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36258,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33168 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36258,
            "unit": "ns/op",
            "extra": "33168 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33168 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33168 times\n4 procs"
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
          "id": "f0630a4c3045713a296f3b5d77ee17f0250a24f2",
          "message": "fix: USA -> United States",
          "timestamp": "2025-10-01T18:20:57Z",
          "url": "https://github.com/moov-io/watchman/commit/f0630a4c3045713a296f3b5d77ee17f0250a24f2"
        },
        "date": 1759409848545,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9731,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "117246 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9731,
            "unit": "ns/op",
            "extra": "117246 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "117246 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "117246 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32753,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "36582 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32753,
            "unit": "ns/op",
            "extra": "36582 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "36582 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36582 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15355,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "75022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15355,
            "unit": "ns/op",
            "extra": "75022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "75022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "75022 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 34083,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 34083,
            "unit": "ns/op",
            "extra": "35767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1314,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "868202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1314,
            "unit": "ns/op",
            "extra": "868202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "868202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "868202 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13701,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "86073 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13701,
            "unit": "ns/op",
            "extra": "86073 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "86073 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "86073 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78844,
            "unit": "ns/op\t   20284 B/op\t     467 allocs/op",
            "extra": "15068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78844,
            "unit": "ns/op",
            "extra": "15068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20284,
            "unit": "B/op",
            "extra": "15068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 467,
            "unit": "allocs/op",
            "extra": "15068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 129177,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 129177,
            "unit": "ns/op",
            "extra": "8542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8542 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8542 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23399540,
            "unit": "ns/op\t16034703 B/op\t  248438 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23399540,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034703,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248438,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18498670,
            "unit": "ns/op\t12621357 B/op\t  144517 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18498670,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12621357,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144517,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 38206245,
            "unit": "ns/op\t44365992 B/op\t  209006 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 38206245,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44365992,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209006,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18751575,
            "unit": "ns/op\t12615067 B/op\t  144498 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18751575,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12615067,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144498,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18497589,
            "unit": "ns/op\t12612136 B/op\t  144494 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18497589,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12612136,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144494,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18661617,
            "unit": "ns/op\t12612250 B/op\t  144494 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18661617,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12612250,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144494,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 5064,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "294132 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 5064,
            "unit": "ns/op",
            "extra": "294132 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "294132 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "294132 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6580,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "199890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6580,
            "unit": "ns/op",
            "extra": "199890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "199890 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "199890 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36750,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36750,
            "unit": "ns/op",
            "extra": "32624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32624 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32624 times\n4 procs"
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
          "id": "2d62bab7abf7dcab88a79c5194bd214fd28c9ede",
          "message": "fix: install fyne",
          "timestamp": "2025-10-13T17:04:38Z",
          "url": "https://github.com/moov-io/watchman/commit/2d62bab7abf7dcab88a79c5194bd214fd28c9ede"
        },
        "date": 1760446827460,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9163,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "124485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9163,
            "unit": "ns/op",
            "extra": "124485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "124485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "124485 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31637,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "37833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31637,
            "unit": "ns/op",
            "extra": "37833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "37833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37833 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14641,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "78907 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14641,
            "unit": "ns/op",
            "extra": "78907 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "78907 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "78907 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32334,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "37096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32334,
            "unit": "ns/op",
            "extra": "37096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "37096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "37096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1289,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "919156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1289,
            "unit": "ns/op",
            "extra": "919156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "919156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "919156 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13331,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "89269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13331,
            "unit": "ns/op",
            "extra": "89269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "89269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "89269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 75469,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15813 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 75469,
            "unit": "ns/op",
            "extra": "15813 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15813 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15813 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 122577,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "8959 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 122577,
            "unit": "ns/op",
            "extra": "8959 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "8959 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8959 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 27019294,
            "unit": "ns/op\t16034527 B/op\t  248430 allocs/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 27019294,
            "unit": "ns/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034527,
            "unit": "B/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17155068,
            "unit": "ns/op\t12622515 B/op\t  144522 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17155068,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12622515,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144522,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36139257,
            "unit": "ns/op\t44362233 B/op\t  209010 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36139257,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44362233,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209010,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17218342,
            "unit": "ns/op\t12623688 B/op\t  144505 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17218342,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12623688,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144505,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17369131,
            "unit": "ns/op\t12618699 B/op\t  144506 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17369131,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12618699,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144506,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17303199,
            "unit": "ns/op\t12620507 B/op\t  144505 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17303199,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12620507,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144505,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4528,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "269493 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4528,
            "unit": "ns/op",
            "extra": "269493 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "269493 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "269493 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5903,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "183848 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5903,
            "unit": "ns/op",
            "extra": "183848 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "183848 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "183848 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36335,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36335,
            "unit": "ns/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33097 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
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
          "id": "f4f0b531f30223d878dfeeabb5861e58f2d60e47",
          "message": "build: fix docker steps for CI",
          "timestamp": "2025-10-30T21:28:18Z",
          "url": "https://github.com/moov-io/watchman/commit/f4f0b531f30223d878dfeeabb5861e58f2d60e47"
        },
        "date": 1761915549861,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9034,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "125385 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9034,
            "unit": "ns/op",
            "extra": "125385 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "125385 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "125385 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31803,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "37678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31803,
            "unit": "ns/op",
            "extra": "37678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "37678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37678 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14904,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "77791 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14904,
            "unit": "ns/op",
            "extra": "77791 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "77791 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "77791 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32987,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35902 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32987,
            "unit": "ns/op",
            "extra": "35902 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35902 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35902 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1276,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "913934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1276,
            "unit": "ns/op",
            "extra": "913934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "913934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "913934 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13494,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "88824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13494,
            "unit": "ns/op",
            "extra": "88824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "88824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "88824 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 75588,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 75588,
            "unit": "ns/op",
            "extra": "15751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15751 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 122990,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "9379 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 122990,
            "unit": "ns/op",
            "extra": "9379 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "9379 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "9379 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22190390,
            "unit": "ns/op\t16034639 B/op\t  248430 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22190390,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034639,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17439971,
            "unit": "ns/op\t12622775 B/op\t  144517 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17439971,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12622775,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144517,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 42264495,
            "unit": "ns/op\t44374240 B/op\t  209009 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 42264495,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44374240,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209009,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 22709087,
            "unit": "ns/op\t12617150 B/op\t  144500 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 22709087,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12617150,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144500,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18851246,
            "unit": "ns/op\t12621613 B/op\t  144499 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18851246,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12621613,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17959059,
            "unit": "ns/op\t12615696 B/op\t  144498 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17959059,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12615696,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144498,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4762,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "256273 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4762,
            "unit": "ns/op",
            "extra": "256273 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "256273 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "256273 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6042,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "183369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6042,
            "unit": "ns/op",
            "extra": "183369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "183369 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "183369 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36755,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32428 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36755,
            "unit": "ns/op",
            "extra": "32428 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32428 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32428 times\n4 procs"
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
          "id": "7dcaaba14b514bead110605d6d51c0cffaf2898e",
          "message": "build: publish static as \"v2-static\" tag",
          "timestamp": "2025-11-24T20:56:54Z",
          "url": "https://github.com/moov-io/watchman/commit/7dcaaba14b514bead110605d6d51c0cffaf2898e"
        },
        "date": 1764075679860,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9460,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121681 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9460,
            "unit": "ns/op",
            "extra": "121681 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121681 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121681 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32610,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "36231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32610,
            "unit": "ns/op",
            "extra": "36231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "36231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36231 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15485,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "74971 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15485,
            "unit": "ns/op",
            "extra": "74971 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "74971 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "74971 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33936,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33936,
            "unit": "ns/op",
            "extra": "35283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1342,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "876783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1342,
            "unit": "ns/op",
            "extra": "876783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "876783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "876783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14212,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "84550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14212,
            "unit": "ns/op",
            "extra": "84550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "84550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "84550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 77954,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15159 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 77954,
            "unit": "ns/op",
            "extra": "15159 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15159 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15159 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 125646,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8654 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 125646,
            "unit": "ns/op",
            "extra": "8654 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8654 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8654 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 28775252,
            "unit": "ns/op\t16034806 B/op\t  248438 allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 28775252,
            "unit": "ns/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034806,
            "unit": "B/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248438,
            "unit": "allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18404228,
            "unit": "ns/op\t12617582 B/op\t  144521 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18404228,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12617582,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144521,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37369079,
            "unit": "ns/op\t44363194 B/op\t  209007 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37369079,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44363194,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209007,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18175642,
            "unit": "ns/op\t12618010 B/op\t  144504 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18175642,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12618010,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144504,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18348633,
            "unit": "ns/op\t12621540 B/op\t  144510 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18348633,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12621540,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144510,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18287729,
            "unit": "ns/op\t12622058 B/op\t  144512 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18287729,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12622058,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144512,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 5021,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "226796 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 5021,
            "unit": "ns/op",
            "extra": "226796 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "226796 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "226796 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6514,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "159786 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6514,
            "unit": "ns/op",
            "extra": "159786 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "159786 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "159786 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36593,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32902 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36593,
            "unit": "ns/op",
            "extra": "32902 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32902 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32902 times\n4 procs"
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
          "id": "cc9b990f67fdc5ada2e0beba29b0f2175fc72447",
          "message": "Merge pull request #675 from moov-io/renovate/all\n\nchore(deps): update dependency bulma-clean-theme to v0.14.0",
          "timestamp": "2025-12-10T17:04:32Z",
          "url": "https://github.com/moov-io/watchman/commit/cc9b990f67fdc5ada2e0beba29b0f2175fc72447"
        },
        "date": 1765458375963,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9813,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121686 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9813,
            "unit": "ns/op",
            "extra": "121686 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121686 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121686 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 35295,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "34322 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 35295,
            "unit": "ns/op",
            "extra": "34322 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "34322 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "34322 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15096,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15096,
            "unit": "ns/op",
            "extra": "76375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76375 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33946,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33946,
            "unit": "ns/op",
            "extra": "35228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1322,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "922897 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1322,
            "unit": "ns/op",
            "extra": "922897 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "922897 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "922897 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14040,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "84205 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14040,
            "unit": "ns/op",
            "extra": "84205 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "84205 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "84205 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78383,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78383,
            "unit": "ns/op",
            "extra": "15225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15225 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 129449,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8516 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 129449,
            "unit": "ns/op",
            "extra": "8516 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8516 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8516 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23047005,
            "unit": "ns/op\t16034421 B/op\t  248428 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23047005,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034421,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248428,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19029629,
            "unit": "ns/op\t12618814 B/op\t  144524 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19029629,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12618814,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144524,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 38927203,
            "unit": "ns/op\t44364333 B/op\t  209005 allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 38927203,
            "unit": "ns/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44364333,
            "unit": "B/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209005,
            "unit": "allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 21517841,
            "unit": "ns/op\t12619551 B/op\t  144509 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 21517841,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12619551,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144509,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 23681315,
            "unit": "ns/op\t12618141 B/op\t  144511 allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 23681315,
            "unit": "ns/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12618141,
            "unit": "B/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144511,
            "unit": "allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 19073924,
            "unit": "ns/op\t12618483 B/op\t  144515 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 19073924,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12618483,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144515,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4661,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "284896 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4661,
            "unit": "ns/op",
            "extra": "284896 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "284896 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "284896 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6137,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "192872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6137,
            "unit": "ns/op",
            "extra": "192872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "192872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "192872 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36374,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33006 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36374,
            "unit": "ns/op",
            "extra": "33006 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33006 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33006 times\n4 procs"
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
          "id": "baabb1b0ede5abe7cebdff71ba93c0320c3c79d4",
          "message": "ingest: trim unicode replacement characters",
          "timestamp": "2025-12-11T23:05:14Z",
          "url": "https://github.com/moov-io/watchman/commit/baabb1b0ede5abe7cebdff71ba93c0320c3c79d4"
        },
        "date": 1765544592998,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9356,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "120594 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9356,
            "unit": "ns/op",
            "extra": "120594 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "120594 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "120594 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32625,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "36972 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32625,
            "unit": "ns/op",
            "extra": "36972 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "36972 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36972 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15039,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76131 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15039,
            "unit": "ns/op",
            "extra": "76131 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76131 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76131 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33124,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "36228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33124,
            "unit": "ns/op",
            "extra": "36228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "36228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36228 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1278,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "923923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1278,
            "unit": "ns/op",
            "extra": "923923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "923923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "923923 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13670,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "86793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13670,
            "unit": "ns/op",
            "extra": "86793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "86793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "86793 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 76966,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15451 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 76966,
            "unit": "ns/op",
            "extra": "15451 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15451 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15451 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 125321,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "8607 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 125321,
            "unit": "ns/op",
            "extra": "8607 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "8607 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8607 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22143045,
            "unit": "ns/op\t16034603 B/op\t  248429 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22143045,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034603,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248429,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 20524631,
            "unit": "ns/op\t12618995 B/op\t  144518 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 20524631,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12618995,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144518,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37410827,
            "unit": "ns/op\t44362872 B/op\t  209007 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37410827,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44362872,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209007,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17782732,
            "unit": "ns/op\t12618926 B/op\t  144500 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17782732,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12618926,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144500,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18038126,
            "unit": "ns/op\t12616600 B/op\t  144498 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18038126,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12616600,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144498,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18349744,
            "unit": "ns/op\t12617026 B/op\t  144499 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18349744,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12617026,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4609,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "219960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4609,
            "unit": "ns/op",
            "extra": "219960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "219960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "219960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6186,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "203340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6186,
            "unit": "ns/op",
            "extra": "203340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "203340 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "203340 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36165,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33174 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36165,
            "unit": "ns/op",
            "extra": "33174 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33174 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33174 times\n4 procs"
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
          "id": "67115193f3404e9825eda6143bdc59742c3d9db6",
          "message": "Merge pull request #678 from moov-io/ingest-accept-gzip\n\ningest: accept gzip bodies",
          "timestamp": "2025-12-12T19:54:02Z",
          "url": "https://github.com/moov-io/watchman/commit/67115193f3404e9825eda6143bdc59742c3d9db6"
        },
        "date": 1765630646961,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9150,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9150,
            "unit": "ns/op",
            "extra": "121056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121056 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31827,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "37543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31827,
            "unit": "ns/op",
            "extra": "37543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "37543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37543 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14756,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "78739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14756,
            "unit": "ns/op",
            "extra": "78739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "78739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "78739 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32458,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "37051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32458,
            "unit": "ns/op",
            "extra": "37051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "37051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "37051 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1266,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "905821 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1266,
            "unit": "ns/op",
            "extra": "905821 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "905821 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "905821 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13528,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "88423 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13528,
            "unit": "ns/op",
            "extra": "88423 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "88423 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "88423 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 74737,
            "unit": "ns/op\t   20284 B/op\t     467 allocs/op",
            "extra": "15954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 74737,
            "unit": "ns/op",
            "extra": "15954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20284,
            "unit": "B/op",
            "extra": "15954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 467,
            "unit": "allocs/op",
            "extra": "15954 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 122965,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 122965,
            "unit": "ns/op",
            "extra": "8984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8984 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8984 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22875283,
            "unit": "ns/op\t16034385 B/op\t  248429 allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22875283,
            "unit": "ns/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034385,
            "unit": "B/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248429,
            "unit": "allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17788665,
            "unit": "ns/op\t12626554 B/op\t  144522 allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17788665,
            "unit": "ns/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12626554,
            "unit": "B/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144522,
            "unit": "allocs/op",
            "extra": "64 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36679995,
            "unit": "ns/op\t44357345 B/op\t  209006 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36679995,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44357345,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209006,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17607897,
            "unit": "ns/op\t12622804 B/op\t  144501 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17607897,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12622804,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17648946,
            "unit": "ns/op\t12618291 B/op\t  144501 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17648946,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12618291,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17442568,
            "unit": "ns/op\t12618642 B/op\t  144501 allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17442568,
            "unit": "ns/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12618642,
            "unit": "B/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4780,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "239127 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4780,
            "unit": "ns/op",
            "extra": "239127 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "239127 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "239127 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6089,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "188395 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6089,
            "unit": "ns/op",
            "extra": "188395 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "188395 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "188395 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36114,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33246 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36114,
            "unit": "ns/op",
            "extra": "33246 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33246 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33246 times\n4 procs"
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
          "id": "3a394922c773adaa281bd003ec4daa85923dcda3",
          "message": "Merge pull request #679 from moov-io/renovate/major-github-artifact-actions\n\nchore(deps): update github artifact actions (major)",
          "timestamp": "2025-12-15T15:46:24Z",
          "url": "https://github.com/moov-io/watchman/commit/3a394922c773adaa281bd003ec4daa85923dcda3"
        },
        "date": 1765890307683,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9500,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "119208 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9500,
            "unit": "ns/op",
            "extra": "119208 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "119208 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "119208 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32197,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32197,
            "unit": "ns/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15092,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76160 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15092,
            "unit": "ns/op",
            "extra": "76160 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76160 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76160 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32953,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "36338 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32953,
            "unit": "ns/op",
            "extra": "36338 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "36338 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36338 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1311,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "890646 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1311,
            "unit": "ns/op",
            "extra": "890646 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "890646 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "890646 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13586,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "87115 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13586,
            "unit": "ns/op",
            "extra": "87115 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "87115 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "87115 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78910,
            "unit": "ns/op\t   20284 B/op\t     467 allocs/op",
            "extra": "15110 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78910,
            "unit": "ns/op",
            "extra": "15110 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20284,
            "unit": "B/op",
            "extra": "15110 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 467,
            "unit": "allocs/op",
            "extra": "15110 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 127999,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 127999,
            "unit": "ns/op",
            "extra": "8644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8644 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8644 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22520998,
            "unit": "ns/op\t16034643 B/op\t  248430 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22520998,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034643,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 22112136,
            "unit": "ns/op\t12616014 B/op\t  144514 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 22112136,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12616014,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144514,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 48435135,
            "unit": "ns/op\t44349272 B/op\t  209000 allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 48435135,
            "unit": "ns/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44349272,
            "unit": "B/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209000,
            "unit": "allocs/op",
            "extra": "22 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17738335,
            "unit": "ns/op\t12616902 B/op\t  144497 allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17738335,
            "unit": "ns/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12616902,
            "unit": "B/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144497,
            "unit": "allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17818281,
            "unit": "ns/op\t12619587 B/op\t  144496 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17818281,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12619587,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144496,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17789115,
            "unit": "ns/op\t12619736 B/op\t  144497 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17789115,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12619736,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144497,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4655,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "222801 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4655,
            "unit": "ns/op",
            "extra": "222801 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "222801 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "222801 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6134,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "182070 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6134,
            "unit": "ns/op",
            "extra": "182070 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "182070 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "182070 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36113,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32780 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36113,
            "unit": "ns/op",
            "extra": "32780 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32780 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32780 times\n4 procs"
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
          "id": "49a9156aaa1cb278b1039968f219b4addd8f640a",
          "message": "Merge pull request #680 from moov-io/fix-sourceID-match\n\nsearch: score 1.0 for sourceID matches",
          "timestamp": "2025-12-19T18:35:47Z",
          "url": "https://github.com/moov-io/watchman/commit/49a9156aaa1cb278b1039968f219b4addd8f640a"
        },
        "date": 1766235433240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9256,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9256,
            "unit": "ns/op",
            "extra": "121766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121766 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31991,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "37306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31991,
            "unit": "ns/op",
            "extra": "37306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "37306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37306 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 14997,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "77270 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 14997,
            "unit": "ns/op",
            "extra": "77270 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "77270 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "77270 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32916,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "36380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32916,
            "unit": "ns/op",
            "extra": "36380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "36380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36380 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1302,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "877698 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1302,
            "unit": "ns/op",
            "extra": "877698 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "877698 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "877698 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13601,
            "unit": "ns/op\t    5225 B/op\t      26 allocs/op",
            "extra": "87382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13601,
            "unit": "ns/op",
            "extra": "87382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5225,
            "unit": "B/op",
            "extra": "87382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "87382 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 76494,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15553 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 76494,
            "unit": "ns/op",
            "extra": "15553 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15553 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15553 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 124900,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8726 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 124900,
            "unit": "ns/op",
            "extra": "8726 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8726 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8726 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21728241,
            "unit": "ns/op\t16034821 B/op\t  248438 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21728241,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034821,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248438,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17347652,
            "unit": "ns/op\t12620343 B/op\t  144518 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17347652,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12620343,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144518,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36810664,
            "unit": "ns/op\t44359425 B/op\t  209003 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36810664,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44359425,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209003,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17371246,
            "unit": "ns/op\t12621376 B/op\t  144502 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17371246,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12621376,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144502,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17463463,
            "unit": "ns/op\t12616707 B/op\t  144501 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17463463,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12616707,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17190099,
            "unit": "ns/op\t12618587 B/op\t  144502 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17190099,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12618587,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144502,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4700,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "228327 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4700,
            "unit": "ns/op",
            "extra": "228327 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "228327 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "228327 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6082,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "187167 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6082,
            "unit": "ns/op",
            "extra": "187167 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "187167 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "187167 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36824,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32497 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36824,
            "unit": "ns/op",
            "extra": "32497 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32497 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32497 times\n4 procs"
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
          "id": "ec63f856d9947e12f4df91a8c76249ef994e0c30",
          "message": "fix: correctly download and prep us_non_sdn",
          "timestamp": "2025-12-22T15:36:06Z",
          "url": "https://github.com/moov-io/watchman/commit/ec63f856d9947e12f4df91a8c76249ef994e0c30"
        },
        "date": 1766495046591,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9193,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "122407 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9193,
            "unit": "ns/op",
            "extra": "122407 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "122407 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "122407 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31586,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "37398 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31586,
            "unit": "ns/op",
            "extra": "37398 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "37398 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37398 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15532,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15532,
            "unit": "ns/op",
            "extra": "76410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76410 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33966,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33966,
            "unit": "ns/op",
            "extra": "35011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35011 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1276,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "956076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1276,
            "unit": "ns/op",
            "extra": "956076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "956076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "956076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13525,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "88212 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13525,
            "unit": "ns/op",
            "extra": "88212 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "88212 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "88212 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 75996,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15721 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 75996,
            "unit": "ns/op",
            "extra": "15721 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15721 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15721 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 123920,
            "unit": "ns/op\t   38557 B/op\t     599 allocs/op",
            "extra": "8557 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 123920,
            "unit": "ns/op",
            "extra": "8557 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38557,
            "unit": "B/op",
            "extra": "8557 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8557 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 21649706,
            "unit": "ns/op\t16034618 B/op\t  248429 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 21649706,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034618,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248429,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17082819,
            "unit": "ns/op\t12619062 B/op\t  144516 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17082819,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12619062,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144516,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36181381,
            "unit": "ns/op\t44368621 B/op\t  209006 allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36181381,
            "unit": "ns/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44368621,
            "unit": "B/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209006,
            "unit": "allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17843584,
            "unit": "ns/op\t12621543 B/op\t  144501 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17843584,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12621543,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17346896,
            "unit": "ns/op\t12621182 B/op\t  144502 allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17346896,
            "unit": "ns/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12621182,
            "unit": "B/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144502,
            "unit": "allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17251721,
            "unit": "ns/op\t12620328 B/op\t  144501 allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17251721,
            "unit": "ns/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12620328,
            "unit": "B/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144501,
            "unit": "allocs/op",
            "extra": "85 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4512,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "297253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4512,
            "unit": "ns/op",
            "extra": "297253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "297253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "297253 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5935,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "220828 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5935,
            "unit": "ns/op",
            "extra": "220828 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "220828 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "220828 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35958,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33176 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35958,
            "unit": "ns/op",
            "extra": "33176 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33176 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33176 times\n4 procs"
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
          "id": "1269716d3da714c7e1bfbdeb321d044464e57741",
          "message": "download: concurrenly geocode addresses",
          "timestamp": "2025-12-23T21:21:04Z",
          "url": "https://github.com/moov-io/watchman/commit/1269716d3da714c7e1bfbdeb321d044464e57741"
        },
        "date": 1766581332453,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9148,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "123181 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9148,
            "unit": "ns/op",
            "extra": "123181 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "123181 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "123181 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32092,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "36866 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32092,
            "unit": "ns/op",
            "extra": "36866 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "36866 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36866 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15025,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15025,
            "unit": "ns/op",
            "extra": "76524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76524 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33215,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "36206 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33215,
            "unit": "ns/op",
            "extra": "36206 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "36206 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36206 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1297,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "893068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1297,
            "unit": "ns/op",
            "extra": "893068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "893068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "893068 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13927,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "85269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13927,
            "unit": "ns/op",
            "extra": "85269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "85269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "85269 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 79571,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "14889 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 79571,
            "unit": "ns/op",
            "extra": "14889 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "14889 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "14889 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 131411,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "8589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 131411,
            "unit": "ns/op",
            "extra": "8589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "8589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8589 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23015236,
            "unit": "ns/op\t16034640 B/op\t  248430 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23015236,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034640,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17859497,
            "unit": "ns/op\t12619883 B/op\t  144519 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17859497,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12619883,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144519,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 38107000,
            "unit": "ns/op\t44364391 B/op\t  209004 allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 38107000,
            "unit": "ns/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44364391,
            "unit": "B/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209004,
            "unit": "allocs/op",
            "extra": "30 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17892675,
            "unit": "ns/op\t12612953 B/op\t  144499 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17892675,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12612953,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144499,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 22410077,
            "unit": "ns/op\t12621289 B/op\t  144508 allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 22410077,
            "unit": "ns/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12621289,
            "unit": "B/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144508,
            "unit": "allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17559976,
            "unit": "ns/op\t12617270 B/op\t  144507 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17559976,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12617270,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144507,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4533,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "327361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4533,
            "unit": "ns/op",
            "extra": "327361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "327361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "327361 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6188,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "195250 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6188,
            "unit": "ns/op",
            "extra": "195250 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "195250 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "195250 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36358,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33055 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36358,
            "unit": "ns/op",
            "extra": "33055 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33055 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33055 times\n4 procs"
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
          "id": "62ef20875321e80a26058b070442b64e48796007",
          "message": "search: add ExportFile to client",
          "timestamp": "2025-12-24T21:40:31Z",
          "url": "https://github.com/moov-io/watchman/commit/62ef20875321e80a26058b070442b64e48796007"
        },
        "date": 1766667679937,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9886,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "117034 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9886,
            "unit": "ns/op",
            "extra": "117034 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "117034 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "117034 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 34367,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "35649 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 34367,
            "unit": "ns/op",
            "extra": "35649 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "35649 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "35649 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15543,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "74276 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15543,
            "unit": "ns/op",
            "extra": "74276 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "74276 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "74276 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33498,
            "unit": "ns/op\t   10715 B/op\t     121 allocs/op",
            "extra": "35666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33498,
            "unit": "ns/op",
            "extra": "35666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10715,
            "unit": "B/op",
            "extra": "35666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35666 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1359,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "891076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1359,
            "unit": "ns/op",
            "extra": "891076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "891076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "891076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14009,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "84471 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14009,
            "unit": "ns/op",
            "extra": "84471 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "84471 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "84471 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78772,
            "unit": "ns/op\t   20284 B/op\t     467 allocs/op",
            "extra": "15192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78772,
            "unit": "ns/op",
            "extra": "15192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20284,
            "unit": "B/op",
            "extra": "15192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 467,
            "unit": "allocs/op",
            "extra": "15192 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 126177,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8720 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 126177,
            "unit": "ns/op",
            "extra": "8720 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8720 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8720 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23442968,
            "unit": "ns/op\t16034570 B/op\t  248430 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23442968,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034570,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18440989,
            "unit": "ns/op\t12623864 B/op\t  144526 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18440989,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12623864,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144526,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37922989,
            "unit": "ns/op\t44359075 B/op\t  209010 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37922989,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44359075,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209010,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18608087,
            "unit": "ns/op\t12617974 B/op\t  144508 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18608087,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12617974,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144508,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 23297005,
            "unit": "ns/op\t12621130 B/op\t  144514 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 23297005,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12621130,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144514,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18323918,
            "unit": "ns/op\t12615193 B/op\t  144514 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18323918,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12615193,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144514,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4500,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "260036 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4500,
            "unit": "ns/op",
            "extra": "260036 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "260036 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "260036 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6115,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "226960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6115,
            "unit": "ns/op",
            "extra": "226960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "226960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "226960 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36515,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32918 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36515,
            "unit": "ns/op",
            "extra": "32918 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32918 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32918 times\n4 procs"
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
          "id": "d66e42fd69ad5dd9180c4c72b6832a8a4442550b",
          "message": "Merge pull request #694 from moov-io/dependabot/bundler/docs/uri-0.13.3\n\nbuild(deps-dev): bump uri from 0.13.2 to 0.13.3 in /docs",
          "timestamp": "2026-01-02T23:24:19Z",
          "url": "https://github.com/moov-io/watchman/commit/d66e42fd69ad5dd9180c4c72b6832a8a4442550b"
        },
        "date": 1767445161862,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9583,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "114614 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9583,
            "unit": "ns/op",
            "extra": "114614 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "114614 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "114614 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 33332,
            "unit": "ns/op\t   12387 B/op\t     128 allocs/op",
            "extra": "36285 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 33332,
            "unit": "ns/op",
            "extra": "36285 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12387,
            "unit": "B/op",
            "extra": "36285 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36285 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15119,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76525 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15119,
            "unit": "ns/op",
            "extra": "76525 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76525 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76525 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32854,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32854,
            "unit": "ns/op",
            "extra": "35832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35832 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1375,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "838329 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1375,
            "unit": "ns/op",
            "extra": "838329 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "838329 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "838329 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14217,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "86691 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14217,
            "unit": "ns/op",
            "extra": "86691 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "86691 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "86691 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 81548,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "14703 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 81548,
            "unit": "ns/op",
            "extra": "14703 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "14703 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "14703 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 132064,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "9118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 132064,
            "unit": "ns/op",
            "extra": "9118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "9118 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "9118 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22363220,
            "unit": "ns/op\t16034601 B/op\t  248430 allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22363220,
            "unit": "ns/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034601,
            "unit": "B/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "66 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17837130,
            "unit": "ns/op\t12620134 B/op\t  144526 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17837130,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12620134,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144526,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 37676666,
            "unit": "ns/op\t44358946 B/op\t  209011 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 37676666,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44358946,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209011,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17911027,
            "unit": "ns/op\t12616153 B/op\t  144508 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17911027,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12616153,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144508,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 22534085,
            "unit": "ns/op\t12619620 B/op\t  144510 allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 22534085,
            "unit": "ns/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12619620,
            "unit": "B/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144510,
            "unit": "allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18041663,
            "unit": "ns/op\t12622559 B/op\t  144511 allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18041663,
            "unit": "ns/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12622559,
            "unit": "B/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144511,
            "unit": "allocs/op",
            "extra": "58 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4934,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "216368 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4934,
            "unit": "ns/op",
            "extra": "216368 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "216368 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "216368 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6618,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "188581 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6618,
            "unit": "ns/op",
            "extra": "188581 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "188581 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "188581 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36779,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32600 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36779,
            "unit": "ns/op",
            "extra": "32600 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32600 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32600 times\n4 procs"
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
          "id": "b44e7aab96147584ba9b875b345076316f1713d4",
          "message": "docs: fix index",
          "timestamp": "2026-01-05T16:19:35Z",
          "url": "https://github.com/moov-io/watchman/commit/b44e7aab96147584ba9b875b345076316f1713d4"
        },
        "date": 1767704682098,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9376,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "121612 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9376,
            "unit": "ns/op",
            "extra": "121612 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "121612 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "121612 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32453,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "36596 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32453,
            "unit": "ns/op",
            "extra": "36596 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "36596 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36596 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15445,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15445,
            "unit": "ns/op",
            "extra": "76009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76009 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 34402,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 34402,
            "unit": "ns/op",
            "extra": "35287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35287 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1363,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "868987 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1363,
            "unit": "ns/op",
            "extra": "868987 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "868987 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "868987 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13768,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "85076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13768,
            "unit": "ns/op",
            "extra": "85076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "85076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "85076 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 78097,
            "unit": "ns/op\t   20284 B/op\t     467 allocs/op",
            "extra": "15462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 78097,
            "unit": "ns/op",
            "extra": "15462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20284,
            "unit": "B/op",
            "extra": "15462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 467,
            "unit": "allocs/op",
            "extra": "15462 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 129016,
            "unit": "ns/op\t   38559 B/op\t     599 allocs/op",
            "extra": "8550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 129016,
            "unit": "ns/op",
            "extra": "8550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38559,
            "unit": "B/op",
            "extra": "8550 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8550 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 23370086,
            "unit": "ns/op\t16034668 B/op\t  248430 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 23370086,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034668,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248430,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19397547,
            "unit": "ns/op\t12618852 B/op\t  144526 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19397547,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12618852,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144526,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 38850654,
            "unit": "ns/op\t44364332 B/op\t  209008 allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 38850654,
            "unit": "ns/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44364332,
            "unit": "B/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209008,
            "unit": "allocs/op",
            "extra": "26 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17615731,
            "unit": "ns/op\t12614523 B/op\t  144510 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17615731,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12614523,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144510,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17325024,
            "unit": "ns/op\t12619579 B/op\t  144512 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17325024,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12619579,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144512,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17846556,
            "unit": "ns/op\t12618199 B/op\t  144512 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17846556,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12618199,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144512,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 5026,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "218922 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 5026,
            "unit": "ns/op",
            "extra": "218922 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "218922 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "218922 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6628,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "183258 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6628,
            "unit": "ns/op",
            "extra": "183258 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "183258 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "183258 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36435,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32930 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36435,
            "unit": "ns/op",
            "extra": "32930 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32930 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32930 times\n4 procs"
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
          "id": "e6386f0e13df0c467c7cb042e65d30a95ebcf2f6",
          "message": "fix: shutdown refresh ticker only on shutdown signal\n\nCo-Authored-By: Aniket Singh <aniket.singh@v3.cash>",
          "timestamp": "2026-02-04T19:17:14Z",
          "url": "https://github.com/moov-io/watchman/commit/e6386f0e13df0c467c7cb042e65d30a95ebcf2f6"
        },
        "date": 1770297919125,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9311,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "122564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9311,
            "unit": "ns/op",
            "extra": "122564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "122564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "122564 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32721,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "36200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32721,
            "unit": "ns/op",
            "extra": "36200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "36200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "36200 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15723,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "75915 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15723,
            "unit": "ns/op",
            "extra": "75915 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "75915 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "75915 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 34942,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "34528 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 34942,
            "unit": "ns/op",
            "extra": "34528 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "34528 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "34528 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1354,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "871815 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1354,
            "unit": "ns/op",
            "extra": "871815 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "871815 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "871815 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14204,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "84512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14204,
            "unit": "ns/op",
            "extra": "84512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "84512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "84512 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 76926,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 76926,
            "unit": "ns/op",
            "extra": "15283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15283 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 125848,
            "unit": "ns/op\t   38557 B/op\t     599 allocs/op",
            "extra": "9519 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 125848,
            "unit": "ns/op",
            "extra": "9519 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38557,
            "unit": "B/op",
            "extra": "9519 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "9519 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 24936620,
            "unit": "ns/op\t16034849 B/op\t  248431 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 24936620,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034849,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248431,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18739101,
            "unit": "ns/op\t12617992 B/op\t  144526 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18739101,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12617992,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144526,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39317243,
            "unit": "ns/op\t44372262 B/op\t  209017 allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39317243,
            "unit": "ns/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44372262,
            "unit": "B/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209017,
            "unit": "allocs/op",
            "extra": "27 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18712864,
            "unit": "ns/op\t12622550 B/op\t  144509 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18712864,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12622550,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144509,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18894537,
            "unit": "ns/op\t12618198 B/op\t  144509 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18894537,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12618198,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144509,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18522001,
            "unit": "ns/op\t12617060 B/op\t  144508 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18522001,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12617060,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144508,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4701,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "231492 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4701,
            "unit": "ns/op",
            "extra": "231492 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "231492 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "231492 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6127,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "175466 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6127,
            "unit": "ns/op",
            "extra": "175466 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "175466 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "175466 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36436,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33100 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36436,
            "unit": "ns/op",
            "extra": "33100 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33100 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33100 times\n4 procs"
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
          "id": "cf815330fffe89ba3ac1deae8e251f3ce19b7ce4",
          "message": "build: fix arm64 label",
          "timestamp": "2026-02-06T15:34:12Z",
          "url": "https://github.com/moov-io/watchman/commit/cf815330fffe89ba3ac1deae8e251f3ce19b7ce4"
        },
        "date": 1770469878491,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9449,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "126087 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9449,
            "unit": "ns/op",
            "extra": "126087 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "126087 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "126087 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32827,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "34958 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32827,
            "unit": "ns/op",
            "extra": "34958 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "34958 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "34958 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15355,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "76086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15355,
            "unit": "ns/op",
            "extra": "76086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "76086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "76086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33869,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "35874 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33869,
            "unit": "ns/op",
            "extra": "35874 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "35874 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "35874 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1352,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "918334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1352,
            "unit": "ns/op",
            "extra": "918334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "918334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "918334 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 14353,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 14353,
            "unit": "ns/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "83188 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 76847,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "15523 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 76847,
            "unit": "ns/op",
            "extra": "15523 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "15523 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "15523 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 124848,
            "unit": "ns/op\t   38558 B/op\t     599 allocs/op",
            "extra": "8796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 124848,
            "unit": "ns/op",
            "extra": "8796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38558,
            "unit": "B/op",
            "extra": "8796 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "8796 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22384302,
            "unit": "ns/op\t16034684 B/op\t  248429 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22384302,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034684,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248429,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 19488146,
            "unit": "ns/op\t12623472 B/op\t  144527 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 19488146,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12623472,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144527,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 39190551,
            "unit": "ns/op\t44360161 B/op\t  209010 allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 39190551,
            "unit": "ns/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44360161,
            "unit": "B/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209010,
            "unit": "allocs/op",
            "extra": "28 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17831436,
            "unit": "ns/op\t12614373 B/op\t  144502 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17831436,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12614373,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144502,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17839540,
            "unit": "ns/op\t12616075 B/op\t  144503 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17839540,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12616075,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144503,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17792645,
            "unit": "ns/op\t12614816 B/op\t  144502 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17792645,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12614816,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144502,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 5128,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "251134 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 5128,
            "unit": "ns/op",
            "extra": "251134 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "251134 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "251134 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6969,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "167702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6969,
            "unit": "ns/op",
            "extra": "167702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "167702 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "167702 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36477,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32989 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36477,
            "unit": "ns/op",
            "extra": "32989 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32989 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32989 times\n4 procs"
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
          "id": "5ef8a4e1d3cc51b422bb83d049b6072ee4fdb157",
          "message": "Merge pull request #700 from moov-io/dependabot/bundler/docs/faraday-2.14.1\n\nbuild(deps-dev): bump faraday from 2.10.1 to 2.14.1 in /docs",
          "timestamp": "2026-02-09T21:01:52Z",
          "url": "https://github.com/moov-io/watchman/commit/5ef8a4e1d3cc51b422bb83d049b6072ee4fdb157"
        },
        "date": 1770731265082,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9390,
            "unit": "ns/op\t    3336 B/op\t      78 allocs/op",
            "extra": "120980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9390,
            "unit": "ns/op",
            "extra": "120980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3336,
            "unit": "B/op",
            "extra": "120980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 78,
            "unit": "allocs/op",
            "extra": "120980 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 32100,
            "unit": "ns/op\t   12386 B/op\t     128 allocs/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 32100,
            "unit": "ns/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12386,
            "unit": "B/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 128,
            "unit": "allocs/op",
            "extra": "37082 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15136,
            "unit": "ns/op\t    3916 B/op\t      87 allocs/op",
            "extra": "75937 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15136,
            "unit": "ns/op",
            "extra": "75937 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 3916,
            "unit": "B/op",
            "extra": "75937 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 87,
            "unit": "allocs/op",
            "extra": "75937 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33243,
            "unit": "ns/op\t   10714 B/op\t     121 allocs/op",
            "extra": "36172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33243,
            "unit": "ns/op",
            "extra": "36172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 10714,
            "unit": "B/op",
            "extra": "36172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 121,
            "unit": "allocs/op",
            "extra": "36172 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1307,
            "unit": "ns/op\t     672 B/op\t       8 allocs/op",
            "extra": "862746 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1307,
            "unit": "ns/op",
            "extra": "862746 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 672,
            "unit": "B/op",
            "extra": "862746 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "862746 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13962,
            "unit": "ns/op\t    5224 B/op\t      26 allocs/op",
            "extra": "81735 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13962,
            "unit": "ns/op",
            "extra": "81735 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5224,
            "unit": "B/op",
            "extra": "81735 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "81735 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 81037,
            "unit": "ns/op\t   20283 B/op\t     466 allocs/op",
            "extra": "14817 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 81037,
            "unit": "ns/op",
            "extra": "14817 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 20283,
            "unit": "B/op",
            "extra": "14817 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 466,
            "unit": "allocs/op",
            "extra": "14817 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 131213,
            "unit": "ns/op\t   38557 B/op\t     599 allocs/op",
            "extra": "9144 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 131213,
            "unit": "ns/op",
            "extra": "9144 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 38557,
            "unit": "B/op",
            "extra": "9144 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 599,
            "unit": "allocs/op",
            "extra": "9144 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22405740,
            "unit": "ns/op\t16034821 B/op\t  248431 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22405740,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 16034821,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 248431,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17362695,
            "unit": "ns/op\t12619701 B/op\t  144520 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17362695,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 12619701,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 144520,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 36676190,
            "unit": "ns/op\t44362710 B/op\t  209010 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 36676190,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 44362710,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 209010,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 20427998,
            "unit": "ns/op\t12623494 B/op\t  144504 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 20427998,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 12623494,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 144504,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 20823937,
            "unit": "ns/op\t12618881 B/op\t  144505 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 20823937,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 12618881,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 144505,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17868512,
            "unit": "ns/op\t12619466 B/op\t  144505 allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17868512,
            "unit": "ns/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 12619466,
            "unit": "B/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 144505,
            "unit": "allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4600,
            "unit": "ns/op\t     190 B/op\t       5 allocs/op",
            "extra": "240967 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4600,
            "unit": "ns/op",
            "extra": "240967 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 190,
            "unit": "B/op",
            "extra": "240967 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 5,
            "unit": "allocs/op",
            "extra": "240967 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6040,
            "unit": "ns/op\t     557 B/op\t      12 allocs/op",
            "extra": "194517 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6040,
            "unit": "ns/op",
            "extra": "194517 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 557,
            "unit": "B/op",
            "extra": "194517 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "194517 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 35991,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "33448 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 35991,
            "unit": "ns/op",
            "extra": "33448 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "33448 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "33448 times\n4 procs"
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
          "id": "4e388505f81963f2f3fca5bae40af0356aaa7031",
          "message": "build: update darwin binary for amd64",
          "timestamp": "2026-02-10T16:25:14Z",
          "url": "https://github.com/moov-io/watchman/commit/4e388505f81963f2f3fca5bae40af0356aaa7031"
        },
        "date": 1770817467585,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10006,
            "unit": "ns/op\t    3952 B/op\t      91 allocs/op",
            "extra": "117943 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10006,
            "unit": "ns/op",
            "extra": "117943 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3952,
            "unit": "B/op",
            "extra": "117943 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 91,
            "unit": "allocs/op",
            "extra": "117943 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 29775,
            "unit": "ns/op\t   12994 B/op\t     141 allocs/op",
            "extra": "39348 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 29775,
            "unit": "ns/op",
            "extra": "39348 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12994,
            "unit": "B/op",
            "extra": "39348 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "39348 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15930,
            "unit": "ns/op\t    4652 B/op\t     102 allocs/op",
            "extra": "74086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15930,
            "unit": "ns/op",
            "extra": "74086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4652,
            "unit": "B/op",
            "extra": "74086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 102,
            "unit": "allocs/op",
            "extra": "74086 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32108,
            "unit": "ns/op\t   11446 B/op\t     137 allocs/op",
            "extra": "36218 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32108,
            "unit": "ns/op",
            "extra": "36218 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11446,
            "unit": "B/op",
            "extra": "36218 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "36218 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1449,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "827920 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1449,
            "unit": "ns/op",
            "extra": "827920 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "827920 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "827920 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13213,
            "unit": "ns/op\t    5320 B/op\t      28 allocs/op",
            "extra": "89856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13213,
            "unit": "ns/op",
            "extra": "89856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5320,
            "unit": "B/op",
            "extra": "89856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "89856 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 79941,
            "unit": "ns/op\t   24220 B/op\t     549 allocs/op",
            "extra": "14918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 79941,
            "unit": "ns/op",
            "extra": "14918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 24220,
            "unit": "B/op",
            "extra": "14918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 549,
            "unit": "allocs/op",
            "extra": "14918 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 118555,
            "unit": "ns/op\t   42495 B/op\t     681 allocs/op",
            "extra": "8880 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 118555,
            "unit": "ns/op",
            "extra": "8880 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 42495,
            "unit": "B/op",
            "extra": "8880 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 681,
            "unit": "allocs/op",
            "extra": "8880 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 26508323,
            "unit": "ns/op\t18515305 B/op\t  300105 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 26508323,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 18515305,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 300105,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17690262,
            "unit": "ns/op\t13861995 B/op\t  170361 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17690262,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 13861995,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 170361,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 33818409,
            "unit": "ns/op\t45610184 B/op\t  234848 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 33818409,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 45610184,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 234848,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17662177,
            "unit": "ns/op\t13860658 B/op\t  170345 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17662177,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 13860658,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 170345,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17460113,
            "unit": "ns/op\t13862776 B/op\t  170345 allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17460113,
            "unit": "ns/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 13862776,
            "unit": "B/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 170345,
            "unit": "allocs/op",
            "extra": "84 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17522511,
            "unit": "ns/op\t13858004 B/op\t  170344 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17522511,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 13858004,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 170344,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4563,
            "unit": "ns/op\t     239 B/op\t       6 allocs/op",
            "extra": "285940 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4563,
            "unit": "ns/op",
            "extra": "285940 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 239,
            "unit": "B/op",
            "extra": "285940 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "285940 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 5825,
            "unit": "ns/op\t     641 B/op\t      14 allocs/op",
            "extra": "202039 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 5825,
            "unit": "ns/op",
            "extra": "202039 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 641,
            "unit": "B/op",
            "extra": "202039 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "202039 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 36833,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32541 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 36833,
            "unit": "ns/op",
            "extra": "32541 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32541 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32541 times\n4 procs"
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
          "id": "8dc132ce9ad1adba34ba33c490168d68b370f079",
          "message": "build: remove extra steps",
          "timestamp": "2026-02-18T22:34:49Z",
          "url": "https://github.com/moov-io/watchman/commit/8dc132ce9ad1adba34ba33c490168d68b370f079"
        },
        "date": 1771507801902,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 9866,
            "unit": "ns/op\t    3952 B/op\t      91 allocs/op",
            "extra": "116967 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 9866,
            "unit": "ns/op",
            "extra": "116967 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3952,
            "unit": "B/op",
            "extra": "116967 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 91,
            "unit": "allocs/op",
            "extra": "116967 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 29604,
            "unit": "ns/op\t   12994 B/op\t     141 allocs/op",
            "extra": "40252 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 29604,
            "unit": "ns/op",
            "extra": "40252 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12994,
            "unit": "B/op",
            "extra": "40252 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "40252 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16362,
            "unit": "ns/op\t    4652 B/op\t     102 allocs/op",
            "extra": "73722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16362,
            "unit": "ns/op",
            "extra": "73722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4652,
            "unit": "B/op",
            "extra": "73722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 102,
            "unit": "allocs/op",
            "extra": "73722 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 33645,
            "unit": "ns/op\t   11446 B/op\t     137 allocs/op",
            "extra": "35684 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 33645,
            "unit": "ns/op",
            "extra": "35684 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11446,
            "unit": "B/op",
            "extra": "35684 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "35684 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1447,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "783250 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1447,
            "unit": "ns/op",
            "extra": "783250 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "783250 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "783250 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13312,
            "unit": "ns/op\t    5320 B/op\t      28 allocs/op",
            "extra": "88713 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13312,
            "unit": "ns/op",
            "extra": "88713 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5320,
            "unit": "B/op",
            "extra": "88713 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "88713 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 83275,
            "unit": "ns/op\t   24219 B/op\t     548 allocs/op",
            "extra": "14373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 83275,
            "unit": "ns/op",
            "extra": "14373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 24219,
            "unit": "B/op",
            "extra": "14373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 548,
            "unit": "allocs/op",
            "extra": "14373 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 126378,
            "unit": "ns/op\t   42494 B/op\t     681 allocs/op",
            "extra": "8820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 126378,
            "unit": "ns/op",
            "extra": "8820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 42494,
            "unit": "B/op",
            "extra": "8820 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 681,
            "unit": "allocs/op",
            "extra": "8820 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22933309,
            "unit": "ns/op\t18515860 B/op\t  300118 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22933309,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 18515860,
            "unit": "B/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 300118,
            "unit": "allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17736160,
            "unit": "ns/op\t13861374 B/op\t  170366 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17736160,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 13861374,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 170366,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 34559022,
            "unit": "ns/op\t45601002 B/op\t  234852 allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 34559022,
            "unit": "ns/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 45601002,
            "unit": "B/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 234852,
            "unit": "allocs/op",
            "extra": "32 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18840449,
            "unit": "ns/op\t13861672 B/op\t  170351 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18840449,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 13861672,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 170351,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 23293039,
            "unit": "ns/op\t13865866 B/op\t  170352 allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 23293039,
            "unit": "ns/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 13865866,
            "unit": "B/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 170352,
            "unit": "allocs/op",
            "extra": "61 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18762065,
            "unit": "ns/op\t13855952 B/op\t  170350 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18762065,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 13855952,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 170350,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 5030,
            "unit": "ns/op\t     239 B/op\t       6 allocs/op",
            "extra": "228960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 5030,
            "unit": "ns/op",
            "extra": "228960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 239,
            "unit": "B/op",
            "extra": "228960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "228960 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6585,
            "unit": "ns/op\t     641 B/op\t      14 allocs/op",
            "extra": "180872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6585,
            "unit": "ns/op",
            "extra": "180872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 641,
            "unit": "B/op",
            "extra": "180872 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "180872 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37037,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32336 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37037,
            "unit": "ns/op",
            "extra": "32336 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32336 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32336 times\n4 procs"
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
          "id": "4fdd2bf54b887afac62569a4374b7f7e819be346",
          "message": "build: fix apt packages",
          "timestamp": "2026-02-19T22:33:19Z",
          "url": "https://github.com/moov-io/watchman/commit/4fdd2bf54b887afac62569a4374b7f7e819be346"
        },
        "date": 1771593681110,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10021,
            "unit": "ns/op\t    3952 B/op\t      91 allocs/op",
            "extra": "115152 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10021,
            "unit": "ns/op",
            "extra": "115152 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3952,
            "unit": "B/op",
            "extra": "115152 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 91,
            "unit": "allocs/op",
            "extra": "115152 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30154,
            "unit": "ns/op\t   12994 B/op\t     141 allocs/op",
            "extra": "40096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30154,
            "unit": "ns/op",
            "extra": "40096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12994,
            "unit": "B/op",
            "extra": "40096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "40096 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 15654,
            "unit": "ns/op\t    4652 B/op\t     102 allocs/op",
            "extra": "74286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 15654,
            "unit": "ns/op",
            "extra": "74286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4652,
            "unit": "B/op",
            "extra": "74286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 102,
            "unit": "allocs/op",
            "extra": "74286 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 31585,
            "unit": "ns/op\t   11446 B/op\t     137 allocs/op",
            "extra": "37783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 31585,
            "unit": "ns/op",
            "extra": "37783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11446,
            "unit": "B/op",
            "extra": "37783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "37783 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1438,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "851355 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1438,
            "unit": "ns/op",
            "extra": "851355 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "851355 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "851355 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13255,
            "unit": "ns/op\t    5320 B/op\t      28 allocs/op",
            "extra": "89307 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13255,
            "unit": "ns/op",
            "extra": "89307 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5320,
            "unit": "B/op",
            "extra": "89307 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "89307 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 81543,
            "unit": "ns/op\t   24220 B/op\t     549 allocs/op",
            "extra": "14626 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 81543,
            "unit": "ns/op",
            "extra": "14626 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 24220,
            "unit": "B/op",
            "extra": "14626 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 549,
            "unit": "allocs/op",
            "extra": "14626 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 119751,
            "unit": "ns/op\t   42495 B/op\t     681 allocs/op",
            "extra": "9658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 119751,
            "unit": "ns/op",
            "extra": "9658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 42495,
            "unit": "B/op",
            "extra": "9658 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 681,
            "unit": "allocs/op",
            "extra": "9658 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 22613420,
            "unit": "ns/op\t18515307 B/op\t  300105 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 22613420,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 18515307,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 300105,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 17596786,
            "unit": "ns/op\t13861422 B/op\t  170363 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 17596786,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 13861422,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 170363,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 34483361,
            "unit": "ns/op\t45598576 B/op\t  234840 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 34483361,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 45598576,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 234840,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 17487108,
            "unit": "ns/op\t13861268 B/op\t  170345 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 17487108,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 13861268,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 170345,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 17374790,
            "unit": "ns/op\t13858675 B/op\t  170351 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 17374790,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 13858675,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 170351,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 17389624,
            "unit": "ns/op\t13864802 B/op\t  170356 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 17389624,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 13864802,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 170356,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4947,
            "unit": "ns/op\t     239 B/op\t       6 allocs/op",
            "extra": "212266 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4947,
            "unit": "ns/op",
            "extra": "212266 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 239,
            "unit": "B/op",
            "extra": "212266 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "212266 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6553,
            "unit": "ns/op\t     641 B/op\t      14 allocs/op",
            "extra": "197176 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6553,
            "unit": "ns/op",
            "extra": "197176 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 641,
            "unit": "B/op",
            "extra": "197176 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "197176 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37343,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "32244 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37343,
            "unit": "ns/op",
            "extra": "32244 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "32244 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "32244 times\n4 procs"
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
          "id": "45954911e3e2277156f6d80e890efd277213814c",
          "message": "Merge pull request #706 from moov-io/dependabot/go_modules/filippo.io/edwards25519-1.1.1\n\nbuild(deps): bump filippo.io/edwards25519 from 1.1.0 to 1.1.1",
          "timestamp": "2026-02-20T23:20:07Z",
          "url": "https://github.com/moov-io/watchman/commit/45954911e3e2277156f6d80e890efd277213814c"
        },
        "date": 1771679403312,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10694,
            "unit": "ns/op\t    3952 B/op\t      91 allocs/op",
            "extra": "113358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10694,
            "unit": "ns/op",
            "extra": "113358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3952,
            "unit": "B/op",
            "extra": "113358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 91,
            "unit": "allocs/op",
            "extra": "113358 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 31984,
            "unit": "ns/op\t   12994 B/op\t     141 allocs/op",
            "extra": "37767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 31984,
            "unit": "ns/op",
            "extra": "37767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12994,
            "unit": "B/op",
            "extra": "37767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "37767 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16739,
            "unit": "ns/op\t    4652 B/op\t     102 allocs/op",
            "extra": "72700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16739,
            "unit": "ns/op",
            "extra": "72700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4652,
            "unit": "B/op",
            "extra": "72700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 102,
            "unit": "allocs/op",
            "extra": "72700 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 34105,
            "unit": "ns/op\t   11446 B/op\t     137 allocs/op",
            "extra": "34951 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 34105,
            "unit": "ns/op",
            "extra": "34951 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11446,
            "unit": "B/op",
            "extra": "34951 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "34951 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1473,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "825723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1473,
            "unit": "ns/op",
            "extra": "825723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "825723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "825723 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13554,
            "unit": "ns/op\t    5320 B/op\t      28 allocs/op",
            "extra": "87295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13554,
            "unit": "ns/op",
            "extra": "87295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5320,
            "unit": "B/op",
            "extra": "87295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "87295 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 83665,
            "unit": "ns/op\t   24219 B/op\t     548 allocs/op",
            "extra": "14293 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 83665,
            "unit": "ns/op",
            "extra": "14293 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 24219,
            "unit": "B/op",
            "extra": "14293 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 548,
            "unit": "allocs/op",
            "extra": "14293 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 125773,
            "unit": "ns/op\t   42494 B/op\t     681 allocs/op",
            "extra": "8994 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 125773,
            "unit": "ns/op",
            "extra": "8994 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 42494,
            "unit": "B/op",
            "extra": "8994 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 681,
            "unit": "allocs/op",
            "extra": "8994 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 30540101,
            "unit": "ns/op\t18515359 B/op\t  300105 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 30540101,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 18515359,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 300105,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18223666,
            "unit": "ns/op\t13864891 B/op\t  170370 allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18223666,
            "unit": "ns/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 13864891,
            "unit": "B/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 170370,
            "unit": "allocs/op",
            "extra": "67 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 35193993,
            "unit": "ns/op\t45598531 B/op\t  234849 allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 35193993,
            "unit": "ns/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 45598531,
            "unit": "B/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 234849,
            "unit": "allocs/op",
            "extra": "33 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18280633,
            "unit": "ns/op\t13857649 B/op\t  170349 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18280633,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 13857649,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 170349,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 18440575,
            "unit": "ns/op\t13855006 B/op\t  170345 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 18440575,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 13855006,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 170345,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 18170342,
            "unit": "ns/op\t13856463 B/op\t  170349 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 18170342,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 13856463,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 170349,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4648,
            "unit": "ns/op\t     239 B/op\t       6 allocs/op",
            "extra": "253078 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4648,
            "unit": "ns/op",
            "extra": "253078 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 239,
            "unit": "B/op",
            "extra": "253078 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "253078 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6200,
            "unit": "ns/op\t     641 B/op\t      14 allocs/op",
            "extra": "191128 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6200,
            "unit": "ns/op",
            "extra": "191128 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 641,
            "unit": "B/op",
            "extra": "191128 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "191128 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 38212,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31467 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 38212,
            "unit": "ns/op",
            "extra": "31467 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31467 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31467 times\n4 procs"
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
          "id": "b2cbc08dd9439ae80d3ca189e7e88898ee88de03",
          "message": "stringscore: fix NaN's",
          "timestamp": "2026-02-23T22:00:57Z",
          "url": "https://github.com/moov-io/watchman/commit/b2cbc08dd9439ae80d3ca189e7e88898ee88de03"
        },
        "date": 1771939789102,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkDebugSimilarity/individuals",
            "value": 10194,
            "unit": "ns/op\t    3952 B/op\t      91 allocs/op",
            "extra": "113416 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - ns/op",
            "value": 10194,
            "unit": "ns/op",
            "extra": "113416 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - B/op",
            "value": 3952,
            "unit": "B/op",
            "extra": "113416 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals - allocs/op",
            "value": 91,
            "unit": "allocs/op",
            "extra": "113416 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug",
            "value": 30373,
            "unit": "ns/op\t   12994 B/op\t     141 allocs/op",
            "extra": "38589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - ns/op",
            "value": 30373,
            "unit": "ns/op",
            "extra": "38589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - B/op",
            "value": 12994,
            "unit": "B/op",
            "extra": "38589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/individuals-debug - allocs/op",
            "value": 141,
            "unit": "allocs/op",
            "extra": "38589 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses",
            "value": 16261,
            "unit": "ns/op\t    4652 B/op\t     102 allocs/op",
            "extra": "71438 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - ns/op",
            "value": 16261,
            "unit": "ns/op",
            "extra": "71438 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - B/op",
            "value": 4652,
            "unit": "B/op",
            "extra": "71438 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses - allocs/op",
            "value": 102,
            "unit": "allocs/op",
            "extra": "71438 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug",
            "value": 32713,
            "unit": "ns/op\t   11446 B/op\t     137 allocs/op",
            "extra": "36772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - ns/op",
            "value": 32713,
            "unit": "ns/op",
            "extra": "36772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - B/op",
            "value": 11446,
            "unit": "B/op",
            "extra": "36772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/businesses-debug - allocs/op",
            "value": 137,
            "unit": "allocs/op",
            "extra": "36772 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels",
            "value": 1500,
            "unit": "ns/op\t     768 B/op\t      10 allocs/op",
            "extra": "844765 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - ns/op",
            "value": 1500,
            "unit": "ns/op",
            "extra": "844765 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - B/op",
            "value": 768,
            "unit": "B/op",
            "extra": "844765 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "844765 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug",
            "value": 13980,
            "unit": "ns/op\t    5320 B/op\t      28 allocs/op",
            "extra": "85153 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - ns/op",
            "value": 13980,
            "unit": "ns/op",
            "extra": "85153 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - B/op",
            "value": 5320,
            "unit": "B/op",
            "extra": "85153 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/vessels-debug - allocs/op",
            "value": 28,
            "unit": "allocs/op",
            "extra": "85153 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft",
            "value": 86280,
            "unit": "ns/op\t   24219 B/op\t     548 allocs/op",
            "extra": "14035 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - ns/op",
            "value": 86280,
            "unit": "ns/op",
            "extra": "14035 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - B/op",
            "value": 24219,
            "unit": "B/op",
            "extra": "14035 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft - allocs/op",
            "value": 548,
            "unit": "allocs/op",
            "extra": "14035 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug",
            "value": 125456,
            "unit": "ns/op\t   42495 B/op\t     681 allocs/op",
            "extra": "8610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - ns/op",
            "value": 125456,
            "unit": "ns/op",
            "extra": "8610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - B/op",
            "value": 42495,
            "unit": "B/op",
            "extra": "8610 times\n4 procs"
          },
          {
            "name": "BenchmarkDebugSimilarity/aircraft-debug - allocs/op",
            "value": 681,
            "unit": "allocs/op",
            "extra": "8610 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count",
            "value": 25034244,
            "unit": "ns/op\t18515260 B/op\t  300105 allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - ns/op",
            "value": 25034244,
            "unit": "ns/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - B/op",
            "value": 18515260,
            "unit": "B/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "Benchmark_Search/dynamic_goroutine_count - allocs/op",
            "value": 300105,
            "unit": "allocs/op",
            "extra": "45 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal",
            "value": 18909116,
            "unit": "ns/op\t13862360 B/op\t  170360 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - ns/op",
            "value": 18909116,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - B/op",
            "value": 13862360,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/normal - allocs/op",
            "value": 170360,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug",
            "value": 35367700,
            "unit": "ns/op\t45606710 B/op\t  234846 allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - ns/op",
            "value": 35367700,
            "unit": "ns/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - B/op",
            "value": 45606710,
            "unit": "B/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/debug - allocs/op",
            "value": 234846,
            "unit": "allocs/op",
            "extra": "31 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address",
            "value": 18824527,
            "unit": "ns/op\t13856078 B/op\t  170344 allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - ns/op",
            "value": 18824527,
            "unit": "ns/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - B/op",
            "value": 13856078,
            "unit": "B/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address - allocs/op",
            "value": 170344,
            "unit": "allocs/op",
            "extra": "73 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email",
            "value": 19173471,
            "unit": "ns/op\t13859669 B/op\t  170343 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - ns/op",
            "value": 19173471,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - B/op",
            "value": 13859669,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_email - allocs/op",
            "value": 170343,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email",
            "value": 19571177,
            "unit": "ns/op\t13860685 B/op\t  170345 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - ns/op",
            "value": 19571177,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - B/op",
            "value": 13860685,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkAPI_Search/name_address_email - allocs/op",
            "value": 170345,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler",
            "value": 4887,
            "unit": "ns/op\t     239 B/op\t       6 allocs/op",
            "extra": "265189 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - ns/op",
            "value": 4887,
            "unit": "ns/op",
            "extra": "265189 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - B/op",
            "value": 239,
            "unit": "B/op",
            "extra": "265189 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairsJaroWinkler - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "265189 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler",
            "value": 6369,
            "unit": "ns/op\t     641 B/op\t      14 allocs/op",
            "extra": "177398 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - ns/op",
            "value": 6369,
            "unit": "ns/op",
            "extra": "177398 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - B/op",
            "value": 641,
            "unit": "B/op",
            "extra": "177398 times\n4 procs"
          },
          {
            "name": "BenchmarkJaroWinkler/BestPairCombinationJaroWinkler - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "177398 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber",
            "value": 37957,
            "unit": "ns/op\t      16 B/op\t       1 allocs/op",
            "extra": "31972 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - ns/op",
            "value": 37957,
            "unit": "ns/op",
            "extra": "31972 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - B/op",
            "value": 16,
            "unit": "B/op",
            "extra": "31972 times\n4 procs"
          },
          {
            "name": "BenchmarkPhoneNumber - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "31972 times\n4 procs"
          }
        ]
      }
    ]
  }
}