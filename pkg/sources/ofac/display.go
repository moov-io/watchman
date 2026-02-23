package ofac

import (
	"net/url"
)

const (
	baseDetailsURL = "https://sanctionssearch.ofac.treas.gov/Details.aspx"
)

func DetailsURL(entityID string) string {
	u, err := url.Parse(baseDetailsURL)
	if err != nil {
		return ""
	}

	q := u.Query()
	q.Set("id", entityID)

	u.RawQuery = q.Encode()

	return u.String()
}
