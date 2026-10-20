window.BENCHMARK_DATA = {
  "lastUpdate": 1792515415928,
  "repoUrl": "https://github.com/moov-io/watchman",
  "entries": {
    "moov-io/watchman CSL US Benchmarks": [
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
        "date": 1791988251418,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 52901462,
            "unit": "ns/op\t58558493 B/op\t   46933 allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 52901462,
            "unit": "ns/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558493,
            "unit": "B/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "21 times\n4 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 6134163,
            "unit": "ns/op\t 2134181 B/op\t   11373 allocs/op",
            "extra": "206 times\n4 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 6134163,
            "unit": "ns/op",
            "extra": "206 times\n4 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2134181,
            "unit": "B/op",
            "extra": "206 times\n4 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11373,
            "unit": "allocs/op",
            "extra": "206 times\n4 procs"
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
        "date": 1791990304105,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 48056701.5,
            "unit": "ns/op\t58558216 B/op\t46933 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 48056701.5,
            "unit": "ns/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558216,
            "unit": "B/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5559444.5,
            "unit": "ns/op\t2089144 B/op\t11359 allocs/op",
            "extra": "216 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5559444.5,
            "unit": "ns/op",
            "extra": "216 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2089144,
            "unit": "B/op",
            "extra": "216 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11359,
            "unit": "allocs/op",
            "extra": "216 times\n2 procs"
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
        "date": 1791991472864,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 37135300,
            "unit": "ns/op\t58558218 B/op\t46933 allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 37135300,
            "unit": "ns/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558218,
            "unit": "B/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "30 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 4271667,
            "unit": "ns/op\t2104188 B/op\t11482 allocs/op",
            "extra": "278 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 4271667,
            "unit": "ns/op",
            "extra": "278 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2104188,
            "unit": "B/op",
            "extra": "278 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "278 times\n2 procs"
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
        "date": 1791992405774,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 49644275,
            "unit": "ns/op\t58558216 B/op\t46933 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 49644275,
            "unit": "ns/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558216,
            "unit": "B/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5575007,
            "unit": "ns/op\t2103184 B/op\t11482 allocs/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5575007,
            "unit": "ns/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2103184,
            "unit": "B/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "213 times\n2 procs"
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
        "date": 1791997870614,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 47802268.5,
            "unit": "ns/op\t58558218 B/op\t46933 allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 47802268.5,
            "unit": "ns/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558218,
            "unit": "B/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5535935,
            "unit": "ns/op\t2104989 B/op\t11482 allocs/op",
            "extra": "217 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5535935,
            "unit": "ns/op",
            "extra": "217 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2104989,
            "unit": "B/op",
            "extra": "217 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "217 times\n2 procs"
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
        "date": 1792083568116,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 27468688,
            "unit": "ns/op\t58558219 B/op\t46933 allocs/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 27468688,
            "unit": "ns/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558219,
            "unit": "B/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "39 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 3163746.5,
            "unit": "ns/op\t2103678 B/op\t11482 allocs/op",
            "extra": "378 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 3163746.5,
            "unit": "ns/op",
            "extra": "378 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2103678,
            "unit": "B/op",
            "extra": "378 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "378 times\n2 procs"
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
        "date": 1792168850667,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 49944355,
            "unit": "ns/op\t58558221 B/op\t46933 allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 49944355,
            "unit": "ns/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558221,
            "unit": "B/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "21 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5344731.5,
            "unit": "ns/op\t2101573 B/op\t11480 allocs/op",
            "extra": "222 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5344731.5,
            "unit": "ns/op",
            "extra": "222 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2101573,
            "unit": "B/op",
            "extra": "222 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11480,
            "unit": "allocs/op",
            "extra": "222 times\n2 procs"
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
        "date": 1792251337267,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 29522656.5,
            "unit": "ns/op\t58558218 B/op\t46933 allocs/op",
            "extra": "38 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 29522656.5,
            "unit": "ns/op",
            "extra": "38 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558218,
            "unit": "B/op",
            "extra": "38 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "38 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 3381477.5,
            "unit": "ns/op\t2102104 B/op\t11482 allocs/op",
            "extra": "357 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 3381477.5,
            "unit": "ns/op",
            "extra": "357 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2102104,
            "unit": "B/op",
            "extra": "357 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "357 times\n2 procs"
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
        "date": 1792338923119,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 49195082,
            "unit": "ns/op\t58558221 B/op\t46933 allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 49195082,
            "unit": "ns/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558221,
            "unit": "B/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5964004.5,
            "unit": "ns/op\t2104511 B/op\t11483 allocs/op",
            "extra": "201 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5964004.5,
            "unit": "ns/op",
            "extra": "201 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2104511,
            "unit": "B/op",
            "extra": "201 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11483,
            "unit": "allocs/op",
            "extra": "201 times\n2 procs"
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
        "date": 1792435756296,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 49051490,
            "unit": "ns/op\t58558216 B/op\t46933 allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 49051490,
            "unit": "ns/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558216,
            "unit": "B/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "24 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5910375.5,
            "unit": "ns/op\t2103534 B/op\t11482 allocs/op",
            "extra": "202 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5910375.5,
            "unit": "ns/op",
            "extra": "202 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2103534,
            "unit": "B/op",
            "extra": "202 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11482,
            "unit": "allocs/op",
            "extra": "202 times\n2 procs"
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
        "date": 1792515415403,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRead",
            "value": 50666116.5,
            "unit": "ns/op\t58558216 B/op\t46933 allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - ns/op",
            "value": 50666116.5,
            "unit": "ns/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - B/op",
            "value": 58558216,
            "unit": "B/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkRead - allocs/op",
            "value": 46933,
            "unit": "allocs/op",
            "extra": "22 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity",
            "value": 5637714.5,
            "unit": "ns/op\t2098924 B/op\t11480 allocs/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - ns/op",
            "value": 5637714.5,
            "unit": "ns/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - B/op",
            "value": 2098924,
            "unit": "B/op",
            "extra": "213 times\n2 procs"
          },
          {
            "name": "BenchmarkFindEntity - allocs/op",
            "value": 11480,
            "unit": "allocs/op",
            "extra": "213 times\n2 procs"
          }
        ]
      }
    ]
  }
}