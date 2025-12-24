package senzing

import (
	"github.com/moov-io/watchman/pkg/search"
)

type List struct {
	SourceList search.SourceList
	Entities   []search.Entity[search.Value]

	Hash string
}
