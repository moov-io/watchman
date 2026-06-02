package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_buildQueryParameters(t *testing.T) {
	cases := []struct {
		entity   Entity[Value]
		opts     SearchOpts
		expected map[string][]string
	}{
		{
			entity: Entity[Value]{
				Name:   "john doe",
				Type:   EntityPerson,
				Source: SourceUSOFAC,
				Person: &Person{
					AltNames:  []string{"jon doe", "johnny doe"},
					Gender:    "male",
					BirthDate: ptr(time.Date(1998, time.April, 12, 10, 30, 0, 0, time.UTC)),
				},
			},
			opts: SearchOpts{
				Limit:    3,
				MinMatch: 0.9,
			},
			expected: map[string][]string{
				"name":      []string{"john doe"},
				"source":    []string{"us_ofac"},
				"type":      []string{"person"},
				"altNames":  []string{"jon doe", "johnny doe"},
				"gender":    []string{"male"},
				"birthDate": []string{"1998-04-12"},
				"limit":     []string{"3"},
				"minMatch":  []string{"0.90"},
			},
		},
		{
			entity: Entity[Value]{
				Name: "Acme Crypto Corp",
				Type: EntityBusiness,
				Business: &Business{
					AltNames: []string{"Super Crypto Corp"},
					Created:  ptr(time.Date(2012, time.December, 31, 10, 30, 0, 0, time.UTC)),
				},
				Contact: ContactInfo{
					EmailAddresses: []string{"press@acmecrypto.com"},
					PhoneNumbers:   []string{"123-456-7890"},
					Websites:       []string{"acmecrypto.com"},
				},
				Addresses: []Address{
					{
						Line1:      "123 Acme St",
						City:       "Acmetown",
						PostalCode: "54321",
						State:      "AC",
						Country:    "US",
					},
				},
				CryptoAddresses: []CryptoAddress{
					{
						Currency: "XBT",
						Address:  "abc12345",
					},
				},
			},
			opts: SearchOpts{
				Limit:    5,
				MinMatch: 0.925,
			},
			expected: map[string][]string{
				"name":          []string{"Acme Crypto Corp"},
				"type":          []string{"business"},
				"altNames":      []string{"Super Crypto Corp"},
				"created":       []string{"2012-12-31"},
				"emailAddress":  []string{"press@acmecrypto.com"},
				"phoneNumber":   []string{"123-456-7890"},
				"website":       []string{"acmecrypto.com"},
				"address":       []string{"123 Acme St Acmetown 54321 AC US"},
				"cryptoAddress": []string{"XBT:abc12345"},
				"limit":         []string{"5"},
				"minMatch":      []string{"0.93"}, // rounded
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.entity.Name, func(t *testing.T) {
			q := make(url.Values)
			got := BuildQueryParameters(SetSearchOpts(q, tc.opts), tc.entity)

			require.Equal(t, len(tc.expected), len(got)) // same key size

			for expectedKey, expectedValues := range tc.expected {
				gotValues, found := got[expectedKey]
				require.True(t, found, fmt.Sprintf("looking for %s", expectedKey))
				require.ElementsMatch(t, expectedValues, gotValues)
			}
		})
	}
}

func TestSearchResponse_Unmarshal(t *testing.T) {
	input := []byte(`{"query":{"name":"Nicolas Maduro","entityType":"person","sourceList":"api-request","sourceID":"","person":{"name":"Nicolas Maduro","altNames":null,"gender":"unknown","birthDate":null,"placeOfBirth":"","deathDate":null,"titles":null,"governmentIDs":null},"business":null,"organization":null,"aircraft":null,"vessel":null,"contact":{"emailAddresses":null,"phoneNumbers":null,"faxNumbers":null,"websites":null},"addresses":[],"cryptoAddresses":null,"affiliations":null,"sanctionsInfo":null,"historicalInfo":null,"sourceData":null},"entities":[{"name":"Nicolas MADURO MOROS","entityType":"person","sourceList":"us_ofac","sourceID":"22790","person":{"name":"Nicolas MADURO MOROS","altNames":null,"gender":"male","birthDate":"1962-11-23T00:00:00Z","placeOfBirth":"","deathDate":null,"titles":["President of the Bolivarian Republic of Venezuela"],"governmentIDs":[{"name":"Cedula","type":"cedula","country":"Venezuela","identifier":"5892464"}]},"business":null,"organization":null,"aircraft":null,"vessel":null,"contact":{"emailAddresses":null,"phoneNumbers":null,"faxNumbers":null,"websites":null},"addresses":null,"cryptoAddresses":null,"affiliations":null,"sanctionsInfo":null,"historicalInfo":null,"sourceData":{"entityID":"22790","sdnName":"MADURO MOROS, Nicolas","sdnType":"individual","program":["VENEZUELA","IRAN-CON-ARMS-EO"],"title":"President of the Bolivarian Republic of Venezuela","callSign":"","vesselType":"","tonnage":"","grossRegisteredTonnage":"","vesselFlag":"","vesselOwner":"","remarks":"DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela."},"match":0.7784062500000001},{"name":"Nicolas Ernesto MADURO GUERRA","entityType":"person","sourceList":"us_ofac","sourceID":"26946","person":{"name":"Nicolas Ernesto MADURO GUERRA","altNames":null,"gender":"male","birthDate":"1990-06-21T00:00:00Z","placeOfBirth":"","deathDate":null,"titles":null,"governmentIDs":[{"name":"Cedula","type":"cedula","country":"Venezuela","identifier":"19398759"}]},"business":null,"organization":null,"aircraft":null,"vessel":null,"contact":{"emailAddresses":null,"phoneNumbers":null,"faxNumbers":null,"websites":null},"addresses":null,"cryptoAddresses":null,"affiliations":null,"sanctionsInfo":null,"historicalInfo":null,"sourceData":{"entityID":"26946","sdnName":"MADURO GUERRA, Nicolas Ernesto","sdnType":"individual","program":["VENEZUELA"],"title":"","callSign":"","vesselType":"","tonnage":"","grossRegisteredTonnage":"","vesselFlag":"","vesselOwner":"","remarks":"DOB 21 Jun 1990; Gender Male; Cedula No. 19398759 (Venezuela)."},"match":0.75133125}]}`)

	var resp SearchResponse
	err := json.Unmarshal(input, &resp)
	require.NoError(t, err)

	require.Equal(t, "Nicolas Maduro", resp.Query.Name)
	require.Len(t, resp.Entities, 2)
}

func ptr[T any](in T) *T {
	return &in
}
