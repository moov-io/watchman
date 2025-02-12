package download

import (
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

type Stats struct {
	Entities []search.Entity[search.Value] `json:"-"`

	Lists      map[string]int    `json:"lists"`
	ListHashes map[string]string `json:"listHashes"`

	StartedAt time.Time `json:"startedAt"`
	EndedAt   time.Time `json:"endedAt"`

	Version string `json:"version"`
}

type Config struct {
	RefreshInterval      time.Duration
	InitialDataDirectory string

	// ErrorOnEmptyList determines whether Watchman should raise an error when a list
	// becomes empty during refresh or download operations. An empty list could indicate:
	//   - A parsing error in the list processing logic
	//   - A corrupted or invalid downloaded file
	//   - Network or storage issues during file transfer
	//   - Other issues requiring manual investigation
	//
	// Setting this to true enables early detection of potential data integrity problems.
	// Default: false
	ErrorOnEmptyList bool

	IncludedLists []search.SourceList // us_ofac, eu_csl, etc...
}
