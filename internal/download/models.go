package download

import (
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

type Stats struct {
	Entities []search.Entity[search.Value] `json:"-"`

	Lists map[string]int `json:"lists"`

	StartedAt time.Time `json:"startedAt"`
	EndedAt   time.Time `json:"endedAt"`
}

type Config struct {
	RefreshInterval      time.Duration
	InitialDataDirectory string

	IncludedLists []search.SourceList // us_ofac, eu_csl, etc...
}
