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
		{
			name: "business",
			input: Entity[Value]{
				Name: "ACME Corporation",
				Type: EntityBusiness,
				Business: &Business{
					Name:     "ACME Corporation",
					AltNames: []string{"Acme Industries", "ACME Holdings"},
				},
			},
			expected: Entity[Value]{
				Name: "ACME Corporation",
				Type: "business",
				Business: &Business{
					Name:     "ACME Corporation",
					AltNames: []string{"Acme Industries", "ACME Holdings"},
				},
				PreparedFields: PreparedFields{
					Name:       "acme corp",
					NameFields: []string{"acme", "corp"},
					AltNames:   []string{"acme industries", "acme holdings"},
					AltNameFields: [][]string{
						{"acme", "industries"},
						{"acme", "holdings"},
					},
				},
			},
		},
		{
			name: "organization",
			input: Entity[Value]{
				Name: "International Trade Organization",
				Type: EntityOrganization,
				Organization: &Organization{
					Name:     "International Trade Organization",
					AltNames: []string{"Trade Org", "ITO Group"},
				},
			},
			expected: Entity[Value]{
				Name: "International Trade Organization",
				Type: "organization",
				Organization: &Organization{
					Name:     "International Trade Organization",
					AltNames: []string{"Trade Org", "ITO Group"},
				},
				PreparedFields: PreparedFields{
					Name:       "international trade organization",
					NameFields: []string{"international", "trade", "organization"},
					AltNames:   []string{"trade org", "ito group"},
					AltNameFields: [][]string{
						{"trade", "org"},
						{"ito", "group"},
					},
				},
			},
		},
		{
			name: "vessel prefixes stripped from prepared name",
			input: Entity[Value]{
				Name: "MV Solenne Harbour, Bangkok",
				Type: EntityVessel,
				Vessel: &Vessel{
					Name:     "MV Solenne Harbour",
					AltNames: []string{"M/T Solenne Harbour"},
				},
			},
			expected: Entity[Value]{
				Name: "MV Solenne Harbour, Bangkok",
				Type: "vessel",
				Vessel: &Vessel{
					Name:     "MV Solenne Harbour",
					AltNames: []string{"M/T Solenne Harbour"},
				},
				PreparedFields: PreparedFields{
					Name:       "solenne harbour",
					NameFields: []string{"solenne", "harbour"},
					AltNames:   []string{"solenne harbour"},
					AltNameFields: [][]string{
						{"solenne", "harbour"},
					},
				},
			},
		},
		{
			name: "english legal-form expansions canonicalized on businesses",
			input: Entity[Value]{
				Name: "Harrowfield Bearings Limited",
				Type: EntityBusiness,
				Business: &Business{
					Name: "Harrowfield Bearings Limited",
				},
			},
			expected: Entity[Value]{
				Name: "Harrowfield Bearings Limited",
				Type: "business",
				Business: &Business{
					Name: "Harrowfield Bearings Limited",
				},
				PreparedFields: PreparedFields{
					Name:       "harrowfield bearings ltd",
					NameFields: []string{"harrowfield", "bearings", "ltd"},
				},
			},
		},
		{
			name: "ex-former name lifted into historical prepared fields",
			input: Entity[Value]{
				Name: "Ocean Pioneer (ex-Cape Diamond)",
				Type: EntityBusiness,
				Business: &Business{
					Name: "Ocean Pioneer (ex-Cape Diamond)",
				},
			},
			expected: Entity[Value]{
				Name: "Ocean Pioneer (ex-Cape Diamond)",
				Type: "business",
				Business: &Business{
					Name: "Ocean Pioneer (ex-Cape Diamond)",
				},
				PreparedFields: PreparedFields{
					Name:                 "ocean pioneer",
					NameFields:           []string{"ocean", "pioneer"},
					HistoricalNames:      []string{"cape diamond"},
					HistoricalNameFields: [][]string{{"cape", "diamond"}},
				},
			},
		},
		{
			name: "websites prepared from urls",
			input: Entity[Value]{
				Name: "SUEX OTC",
				Type: EntityBusiness,
				Business: &Business{
					Name: "SUEX OTC",
				},
				Contact: ContactInfo{
					Websites: []string{"https://www.suex.io/about", "WWW.SUEX.IO", "suex.io"},
				},
			},
			expected: Entity[Value]{
				Name: "SUEX OTC",
				Type: "business",
				Business: &Business{
					Name: "SUEX OTC",
				},
				Contact: ContactInfo{
					Websites: []string{"https://www.suex.io/about", "WWW.SUEX.IO", "suex.io"},
				},
				PreparedFields: PreparedFields{
					Name:       "suex otc",
					NameFields: []string{"suex", "otc"},
					Contact: ContactInfo{
						Websites: []string{"suex.io"},
					},
				},
			},
		},
		{
			name: "f.k.a. former name lifted into historical prepared fields",
			input: Entity[Value]{
				Name: "Ostrowski Grain Partners LLC (f.k.a. Ostrowski Feed & Grain LLC)",
				Type: EntityBusiness,
				Business: &Business{
					Name: "Ostrowski Grain Partners LLC (f.k.a. Ostrowski Feed & Grain LLC)",
				},
			},
			expected: Entity[Value]{
				Name: "Ostrowski Grain Partners LLC (f.k.a. Ostrowski Feed & Grain LLC)",
				Type: "business",
				Business: &Business{
					Name: "Ostrowski Grain Partners LLC (f.k.a. Ostrowski Feed & Grain LLC)",
				},
				PreparedFields: PreparedFields{
					Name:                 "ostrowski grain partners llc",
					NameFields:           []string{"ostrowski", "grain", "partners", "llc"},
					HistoricalNames:      []string{"ostrowski feed grain llc"},
					HistoricalNameFields: [][]string{{"ostrowski", "feed", "grain", "llc"}},
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
