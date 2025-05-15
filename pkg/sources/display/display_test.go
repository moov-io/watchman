package display_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/moov-io/watchman/pkg/sources/display"

	"github.com/stretchr/testify/require"
)

func TestDetailsURL(t *testing.T) {
	apiRequest := ofactest.FindEntity(t, "10278")
	apiRequest.Source = search.SourceAPIRequest

	cases := []struct {
		Entity   search.Entity[search.Value]
		Expected string
	}{
		{
			Entity: search.Entity[search.Value]{
				Source:   search.SourceUSOFAC,
				SourceID: "12345",
			},
			Expected: "https://sanctionssearch.ofac.treas.gov/Details.aspx?id=12345",
		},
		{
			Entity: search.Entity[search.Value]{
				Source:   search.SourceUSCSL,
				SourceID: "wont-be-added",
			},
			Expected: "https://www.trade.gov/data-visualization/csl-search",
		},
		{
			Entity:   apiRequest,
			Expected: "/v2/search?altNames=BURTON+BURGESS&birthDate=1963-07-28&name=Elvis+Angus+LOGAN+MOREY&type=person",
		},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%s/%s", tc.Entity.Source, tc.Entity.SourceID)

		t.Run(name, func(t *testing.T) {
			got := display.DetailsURL(tc.Entity)
			require.Equal(t, tc.Expected, got)
		})
	}
}
