## ofaccheck

`ofaccheck` is a cli tool used for testing batches of names against Moov's OFAC service.

With no arguments the contaier runs tests against the production API, but we strongly ask you run ofaccheck against local instances of OFAC.

```
$ go install ./cmd/ofaccheck
$ ofaccheck -allowed-file users.txt -blocked-file criminals.txt -threshold 0.99 -sdn-type individual -v
2019/10/09 17:36:16.160025 main.go:61: Starting moov/ofaccheck v0.12.0-dev
2019/10/09 17:36:16.160055 main.go:64: [INFO] using http://localhost:8084 for address
2019/10/09 17:36:16.161818 main.go:73: [SUCCESS] ping
2019/10/09 17:36:16.174108 main.go:156: [INFO] didn't block 'Husein HAZEM'
2019/10/09 17:36:16.212986 main.go:148: [INFO] blocked 'Nicolas Ernesto MADURO GUERRA'
2019/10/09 17:36:16.213423 main.go:146: [ERROR] 'Maria Alexandra PERDOMO' wasn't blocked (match=0.96)
exit status 1
```

__ofaccheck is not a stable tool. Please contact Moov developers if you intend to use this tool, otherwise we might change the tool (or remove it) without notice.__
