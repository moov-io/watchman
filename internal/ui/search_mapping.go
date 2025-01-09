package ui

import (
	"github.com/moov-io/watchman/pkg/search"
)

var (
	searchName = "Name"
)

func buildQueryEntity(populatedItems []item) search.Entity[search.Value] {
	var out search.Entity[search.Value]

	for _, qry := range populatedItems {
		switch qry.name {
		case searchName:
			out.Name = qry.value
		}
	}

	return out
}
