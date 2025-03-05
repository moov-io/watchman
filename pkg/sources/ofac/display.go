package ofac

import (
	"fmt"
	"net/url"
)

const (
	baseDetailsURL = "https://sanctionssearch.ofac.treas.gov/Details.aspx"
)

func DetailsURL(entityID string) string {
	u, err := url.Parse(baseDetailsURL)
	if err != nil {
		panic(fmt.Sprintf("invalid %s as baseDetailsURL: %v", baseDetailsURL, err)) //nolint:forbidigo
	}

	q := u.Query()
	q.Set("id", entityID)

	u.RawQuery = q.Encode()

	return u.String()
}
