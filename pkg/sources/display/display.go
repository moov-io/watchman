package display

import (
	"net/url"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/csl_us"
	"github.com/moov-io/watchman/pkg/sources/ofac"
)

// DetailsURL returns a URL where you can view the entity on the source's website
func DetailsURL(entity search.Entity[search.Value]) string {
	switch entity.Source {
	case search.SourceEUCSL:
		return "TODO"

	case search.SourceUKCSL:
		return "TODO"

	case search.SourceUSCSL:
		return csl_us.DetailsURL(entity.SourceID)

	case search.SourceUSOFAC:
		return ofac.DetailsURL(entity.SourceID)

	case search.SourceAPIRequest:
		// do nothing
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
