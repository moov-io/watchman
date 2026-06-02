package main

import (
	"context"
	"fmt"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

func main() {
	client := search.NewClient(nil, "http://localhost:8084")
	ctx := context.Background()

	// List the currently loaded sanction lists (from /v2/listinfo)
	info, err := client.ListInfo(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("loaded %d list(s): %v\n\n", len(info.Lists), info.Lists)

	// Search using the unified entity query model (mirrors /v2/search)
	query := search.Entity[search.Value]{
		Name: "Nicolas Maduro",
		Type: search.EntityPerson,
		Person: &search.Person{
			BirthDate: new(time.Date(1962, time.November, 23, 0, 0, 0, 0, time.UTC)),
			Gender:    search.GenderMale,
		},
		// Addresses, Contact, CryptoAddresses, Source, SourceID also supported
	}
	opts := search.SearchOpts{
		Limit:    5,
		MinMatch: 0.75,
		Debug:    false,
	}

	resp, err := client.SearchByEntity(ctx, query, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Results for: %v\n", resp.Query.Name)

	for _, e := range resp.Entities {
		fmt.Printf("  match=%.4f %s (%s) from %s\n", e.Match, e.Name, e.Type, e.Source)
	}
}
