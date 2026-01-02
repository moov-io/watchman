package download

import (
	"time"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

type Stats struct {
	Entities []search.Entity[search.Value] `json:"-"`

	// TFIDFIndex contains precomputed IDF values for name term weighting.
	// This is built from all entity names after loading and used during search.
	TFIDFIndex *tfidf.Index `json:"-"`

	Lists      map[string]int    `json:"lists"`
	ListHashes map[string]string `json:"listHashes"`

	StartedAt time.Time `json:"startedAt"`
	EndedAt   time.Time `json:"endedAt"`

	Version string `json:"version"`
}
