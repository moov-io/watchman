package search

import (
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
				Name: "john doe",
				Type: EntityPerson,
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

func ptr[T any](in T) *T {
	return &in
}
