---
layout: page
title: Go library
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Go library

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/watchman?utm_source=godoc)](https://pkg.go.dev/github.com/moov-io/watchman)

```go
import (
    "github.com/moov-io/watchman/pkg/search"
)

func main() {
    client := search.NewClient(nil, "http://localhost:8084")

    ctx := context.Background()
    query := public.Entity[public.Value]{
    	Name: "Flight",
    	Type: public.EntityAircraft,
    }
    var opts public.SearchOpts

    response, err := scope.client.SearchByEntity(ctx, query, opts)
    if err != nil {
         // deal with errors
    }

    for _, entity := response.Entities {
        // do something with each entity
    }
}
```
