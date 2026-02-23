package csl_us

import (
	"net/url"
)

const (
	baseDetailsURL = "https://www.trade.gov/data-visualization/csl-search"
)

// DetailsURL returns the US CSL Search URL
// There is not a page for viewing individual records we can link to.
func DetailsURL(entityID string) string {
	u, err := url.Parse(baseDetailsURL)
	if err != nil {
		return ""
	}

	return u.String()
}
