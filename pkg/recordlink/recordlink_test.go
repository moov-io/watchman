// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package recordlink

import (
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestKeys_NoPII(t *testing.T) {
	e := search.Entity[search.Value]{
		Name: "John Smith",
		Type: search.EntityPerson,
		Person: &search.Person{
			Name: "John Smith",
			GovernmentIDs: []search.GovernmentID{{
				Type:       search.GovernmentIDPassport,
				Country:    "US",
				Identifier: "AA-111",
			}},
		},
		Contact: search.ContactInfo{
			EmailAddresses: []string{"john.smith123@example.com"},
			PhoneNumbers:   []string{"+1-555-0100"},
		},
		Addresses: []search.Address{{
			Line1:      "541 First St",
			Line2:      "Apt 301",
			City:       "Anytown",
			State:      "CA",
			PostalCode: "90210",
			Country:    "US",
		}},
	}

	keys := Keys(e)
	require.NotEmpty(t, keys)

	joined := strings.Join(keys, "\n")
	for _, leak := range []string{
		"John Smith", "john smith",
		"AA-111", "AA111",
		"541 First", "Anytown", "90210",
		"john.smith123@example.com",
		"555-0100", "5550100",
	} {
		require.NotContains(t, joined, leak, "keys must not contain %q", leak)
	}

	var kinds []string
	for _, k := range keys {
		kind, _, ok := strings.Cut(k, ":")
		require.True(t, ok)
		kinds = append(kinds, kind)
	}
	require.Contains(t, kinds, KindType)
	require.Contains(t, kinds, KindName)
	require.Contains(t, kinds, KindGovID)
	require.Contains(t, kinds, KindAddr)
	require.Contains(t, kinds, KindContact)
}

func TestKeys_NormalizesCountryAndID(t *testing.T) {
	a := search.Entity[search.Value]{
		Name: "Acme",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "Acme",
			GovernmentIDs: []search.GovernmentID{{
				Type:       search.GovernmentIDTax,
				Country:    "United States",
				Identifier: "12-345",
			}},
		},
	}
	b := search.Entity[search.Value]{
		Name: "Acme",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "Acme",
			GovernmentIDs: []search.GovernmentID{{
				Type:       search.GovernmentIDTax,
				Country:    "US",
				Identifier: "12345",
			}},
		},
	}

	ka, kb := Keys(a), Keys(b)
	require.Equal(t, govID(ka), govID(kb))
}

func TestPrefixes(t *testing.T) {
	require.Nil(t, Prefixes(""))
	require.Nil(t, Prefixes("nocolon"))

	got := Prefixes("ADDR:Caaaaaaa|Sbbbbbbb|Pcccccccc")
	require.Equal(t, []string{
		"ADDR:Caaaaaaa",
		"ADDR:Caaaaaaa|Sbbbbbbb",
		"ADDR:Caaaaaaa|Sbbbbbbb|Pcccccccc",
	}, got)
}

func govID(keys []string) string {
	for _, k := range keys {
		if strings.HasPrefix(k, KindGovID+":") {
			return k
		}
	}
	return ""
}
