package search

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestSearchedEntityJSON(t *testing.T) {
	type SDN struct {
		EntityID string `json:"entityID"`
	}

	bs, err := json.MarshalIndent(SearchedEntity[SDN]{
		Entity: search.Entity[SDN]{
			SourceData: SDN{
				EntityID: "12345",
			},
		},
		Match: 0.6401,
	}, "", "  ")
	require.NoError(t, err)

	expected := strings.TrimSpace(`{
  "name": "",
  "entityType": "",
  "sourceList": "",
  "sourceID": "",
  "person": null,
  "business": null,
  "organization": null,
  "aircraft": null,
  "vessel": null,
  "contact": null,
  "addresses": null,
  "cryptoAddresses": null,
  "affiliations": null,
  "sanctionsInfo": null,
  "historicalInfo": null,
  "titles": null,
  "sourceData": {
    "entityID": "12345"
  },
  "match": 0.6401
}`)
	require.Equal(t, expected, string(bs))
}
