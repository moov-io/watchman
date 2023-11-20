## watchmantest

`watchmantest` is a cli tool used for testing the Moov Sanction Search service.

With no arguments the contaier runs tests against the production API. This tool requires an OAuth token provided by github.com/moov-io/api written to the local disk, but running apitest first will write this token.

This tool can be used to query with custom searches:

```
$ go install ./cmd/watchmantest
$ watchmantest -local moh
2019/02/14 23:37:44.432334 main.go:44: Starting moov/watchmantest v0.4.1-dev
2019/02/14 23:37:44.432366 main.go:60: [INFO] using http://localhost:8084 for address
2019/02/14 23:37:44.434534 main.go:76: [SUCCESS] ping
2019/02/14 23:37:44.435204 main.go:83: [SUCCESS] last download was: 3h45m58s ago
2019/02/14 23:37:44.440230 main.go:96: [SUCCESS] name search passed, query="moh"
2019/02/14 23:37:44.445473 main.go:118: [SUCCESS] alt name search passed
2019/02/14 23:37:44.449367 main.go:123: [SUCCESS] address search passed
```

__watchmantest is not a stable tool. Please contact Moov developers if you intend to use this tool, otherwise we might change the tool (or remove it) without notice.__
