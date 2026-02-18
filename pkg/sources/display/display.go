package display

import (
	"net/url"
	"strings"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_eu"
	"github.com/moov-io/watchman/pkg/sources/csl_uk"
	"github.com/moov-io/watchman/pkg/sources/csl_us"
	"github.com/moov-io/watchman/pkg/sources/ofac"
	"github.com/moov-io/watchman/pkg/sources/opensanctions"
)

// DetailsURL returns a URL where you can view the entity on the source's website
func DetailsURL(entity search.Entity[search.Value]) string {
	switch entity.Source {
	case search.SourceEUCSL:
		return csl_eu.DetailsURL(entity)

	case search.SourceUKCSL:
		return csl_uk.DetailsURL(entity.SourceID)

	case search.SourceUSCSL:
		return csl_us.DetailsURL(entity.SourceID)

	case search.SourceUSOFAC, search.SourceUSNonSDN:
		return ofac.DetailsURL(entity.SourceID)

	case search.SourceAPIRequest:
		// do nothing
	}

	// Shortcut for open sanctions
	if strings.HasPrefix(string(entity.Source), "opensanctions_") {
		return opensanctions.DetailsURL(entity.SourceID)
	}

	// Format the entity as a Watchman search URL
	u, _ := url.Parse("/v2/search")

	params := search.BuildQueryParameters(u.Query(), entity)
	if v := params.Get("source"); v == string(search.SourceAPIRequest) {
		params.Del("source")
	}

	u.RawQuery = params.Encode()

	return u.String()
}
