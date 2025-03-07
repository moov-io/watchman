---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.18 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/watchman/releases/latest) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/watchman.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/watchman

$ go doc github.com/moov-io/watchman/client Search
```
