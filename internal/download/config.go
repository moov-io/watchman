package download

import (
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

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

	// IgnoredDownloadErrors lists source-list names for which download or parse
	// errors are suppressed during RefreshAll. A failed list that is named here
	// will produce a warning log but will not cause RefreshAll to return an error.
	//
	// Valid values are the standard known downloadable lists plus any source-list
	// names configured in OpenSanctions.Lists or Senzing.
	//
	// Use the IGNORED_DOWNLOAD_ERRORS environment variable to set this at runtime
	// using a comma-separated list of source-list names.
	IgnoredDownloadErrors []search.SourceList

	OpenSanctions OpenSanctionsConfig
	Senzing       []SenzingList
}

type OpenSanctionsConfig struct {
	ApiKey string
	Lists  []SenzingList
}

type SenzingList struct {
	SourceList search.SourceList
	Location   string
}
