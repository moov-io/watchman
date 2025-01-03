package search

import (
	"github.com/moov-io/watchman/pkg/search"
)

type SearchedEntity[T any] struct {
	search.Entity[T]

	Match float64 `json:"match"`
}
