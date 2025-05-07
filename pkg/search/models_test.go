package search

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEntityJSON(t *testing.T) {
	type SDN struct {
		EntityID string `json:"entityID"`
	}
	bs, err := json.MarshalIndent(Entity[SDN]{
		SourceData: SDN{
			EntityID: "12345",
		},
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
  "contact": {
    "emailAddresses": null,
    "phoneNumbers": null,
    "faxNumbers": null,
    "websites": null
  },
  "addresses": null,
  "cryptoAddresses": null,
  "affiliations": null,
  "sanctionsInfo": null,
  "historicalInfo": null,
  "sourceData": {
    "entityID": "12345"
  }
}`)
	require.Equal(t, expected, string(bs))
}

func TestEntity_Normalize(t *testing.T) {
	birthDate := time.Date(1993, time.April, 17, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		name            string
		input, expected Entity[Value]
	}{
		{
			name:     "empty",
			input:    Entity[Value]{},
			expected: Entity[Value]{},
		},
		{
			name: "person",
			input: Entity[Value]{
				Name: "Dmitry Yuryevich KHOROSHEV",
				Type: EntityPerson,
				Person: &Person{
					Name:      "Dmitry Yuryevich KHOROSHEV",
					BirthDate: &birthDate,
					Gender:    GenderMale,
				},
				Contact: ContactInfo{
					EmailAddresses: []string{"khoroshev1@icloud.com"},
					PhoneNumbers:   []string{"1-555-123-4567"},
				},
				Addresses: []Address{
					{
						Line1:      "1234 Broadway Suite 500",
						City:       "New York",
						State:      "NY",
						PostalCode: "10013",
						Country:    "US",
					},
				},
			},
			expected: Entity[Value]{
				Name: "Dmitry Yuryevich KHOROSHEV",
				Type: "person",
				Person: &Person{
					Name:      "Dmitry Yuryevich KHOROSHEV",
					Gender:    "male",
					BirthDate: &birthDate,
				},
				Contact: ContactInfo{
					EmailAddresses: []string{"khoroshev1@icloud.com"},
					PhoneNumbers:   []string{"1-555-123-4567"},
				},
				Addresses: []Address{
					{Line1: "1234 Broadway Suite 500", City: "New York", PostalCode: "10013", State: "NY", Country: "US"},
				},
				PreparedFields: PreparedFields{
					Name:       "dmitry yuryevich khoroshev",
					NameFields: []string{"dmitry", "yuryevich", "khoroshev"},
					Contact: ContactInfo{
						PhoneNumbers: []string{"15551234567"},
					},
					Addresses: []PreparedAddress{
						{
							Line1:       "1234 broadway suite 500",
							Line1Fields: []string{"1234", "broadway", "suite", "500"},
							City:        "new york",
							CityFields:  []string{"new", "york"},
							PostalCode:  "10013",
							State:       "ny",
							Country:     "united states",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.input.Normalize())
		})
	}
}
