package csl_us

import (
	"fmt"
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
		panic(fmt.Sprintf("invalid %s as baseDetailsURL: %v", baseDetailsURL, err)) //nolint:forbidigo
	}

	return u.String()
}
