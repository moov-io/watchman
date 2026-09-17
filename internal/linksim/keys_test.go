// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package linksim

import (
	"regexp"
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

var keyShape = regexp.MustCompile(`^(TYPE|NAME|GOVID|ADDR|CONTACT|IMO|MMSI|AIR):([CSTPXYLE][0-9a-f]{8}(\|[CSTPXYLE][0-9a-f]{8})*|[0-9a-f]{8})$`)

func TestKeys_ShapeAndNoPII(t *testing.T) {
	keys := Keys(john)
	require.NotEmpty(t, keys)

	joined := strings.Join(keys, "\n")
	for _, leak := range []string{
		"John Smith", "john smith", "Johnathon",
		"1234567890",
		"541 First", "541 first",
		"Anytown", "anytown",
		"john.smith123@example.com",
		"Apt 301",
	} {
		require.NotContains(t, joined, leak, "blocking keys must not contain %q", leak)
	}

	var kinds []string
	for _, k := range keys {
		require.Truef(t, keyShape.MatchString(k), "key %q does not match KIND:hashed-segments", k)
		kind, _, _ := strings.Cut(k, ":")
		kinds = append(kinds, kind)
	}
	require.Contains(t, kinds, KindType)
	require.Contains(t, kinds, KindName)
	require.Contains(t, kinds, KindGovID)
	require.Contains(t, kinds, KindAddr)
	require.Contains(t, kinds, KindContact)
}

func TestKeys_SimilarShareAddressPrefixAndGovID(t *testing.T) {
	johnKeys := Keys(john)
	johnathonKeys := Keys(johnathon)

	require.Equal(t, govIDKey(johnKeys), govIDKey(johnathonKeys), "same passport must produce the same GOVID key")

	jp := addrPrefixes(johnKeys)
	ap := addrPrefixes(johnathonKeys)
	require.NotEmpty(t, jp)
	require.NotEmpty(t, ap)

	// Same country/state/postal/city/line1; johnathon has Line2 so the full
	// keys differ, but they share the prefix through line1.
	require.Equal(t, jp[0], ap[0], "country")
	require.GreaterOrEqual(t, len(jp), 5)
	require.GreaterOrEqual(t, len(ap), 5)
	require.Equal(t, jp[4], ap[4], "prefix through line1 (C|S|P|Y|L)")
	require.NotEqual(t, last(jp), last(ap), "line2 makes the full ADDR keys differ")
}

func TestKeys_PhoneticNameTokens(t *testing.T) {
	smith := search.Entity[search.Value]{
		Name: "John Smith",
		Type: search.EntityPerson,
	}.Normalize()
	smythe := search.Entity[search.Value]{
		Name: "Jon Smythe",
		Type: search.EntityPerson,
	}.Normalize()

	s1 := nameKeys(Keys(smith))
	s2 := nameKeys(Keys(smythe))
	require.NotEmpty(t, s1)
	require.NotEmpty(t, s2)

	shared := intersect(s1, s2)
	require.NotEmpty(t, shared, "Smith/Smythe and John/Jon should share hashed Soundex NAME keys")
}

func TestKeys_DissimilarGovernmentIDs(t *testing.T) {
	a := Keys(john)
	b := Keys(jane)
	require.NotEqual(t, govIDKey(a), govIDKey(b))
	require.NotEqual(t, addrKey(a), addrKey(b), "different country/city should not share a full ADDR key")
}

func TestKeys_OrganizationGovernmentIDs(t *testing.T) {
	org := search.Entity[search.Value]{
		Name: "Acme Corp",
		Type: search.EntityOrganization,
		Organization: &search.Organization{
			Name: "Acme Corp",
			GovernmentIDs: []search.GovernmentID{
				{
					Type:       search.GovernmentIDBusinessRegisration,
					Country:    "US",
					Identifier: "BR-999888",
				},
			},
		},
	}.Normalize()

	keys := Keys(org)
	require.NotEmpty(t, govIDKey(keys))
	require.NotContains(t, strings.Join(keys, " "), "BR-999888")
	require.NotContains(t, strings.Join(keys, " "), "999888")
}

func TestKeys_Deterministic(t *testing.T) {
	require.Equal(t, Keys(john), Keys(john))
}

func TestKeys_CountryAliases(t *testing.T) {
	us := search.Entity[search.Value]{
		Name: "Pat Doe",
		Type: search.EntityPerson,
		Person: &search.Person{
			GovernmentIDs: []search.GovernmentID{
				{Type: search.GovernmentIDPassport, Country: "US", Identifier: "AA111"},
			},
		},
	}.Normalize()
	usa := search.Entity[search.Value]{
		Name: "Pat Doe",
		Type: search.EntityPerson,
		Person: &search.Person{
			GovernmentIDs: []search.GovernmentID{
				{Type: search.GovernmentIDPassport, Country: "United States", Identifier: "AA-111"},
			},
		},
	}.Normalize()

	require.Equal(t, govIDKey(Keys(us)), govIDKey(Keys(usa)), "US / United States and dashed identifiers must hash equal")
}

func TestPrefixes(t *testing.T) {
	t.Run("composite", func(t *testing.T) {
		key := "ADDR:Caaaaaaa1|Sbbbbbbb2|Pcccccccc"
		require.Equal(t, []string{
			"ADDR:Caaaaaaa1",
			"ADDR:Caaaaaaa1|Sbbbbbbb2",
			"ADDR:Caaaaaaa1|Sbbbbbbb2|Pcccccccc",
		}, Prefixes(key))
	})
	t.Run("single", func(t *testing.T) {
		require.Equal(t, []string{"NAME:deadbeef"}, Prefixes("NAME:deadbeef"))
	})
	t.Run("empty", func(t *testing.T) {
		require.Nil(t, Prefixes(""))
		require.Nil(t, Prefixes("NOCOLON"))
	})
}

func TestKeys_EmptyEntity(t *testing.T) {
	require.Empty(t, Keys(search.Entity[search.Value]{}))
}

func govIDKey(keys []string) string {
	for _, k := range keys {
		if strings.HasPrefix(k, KindGovID+":") {
			return k
		}
	}
	return ""
}

func addrKey(keys []string) string {
	for _, k := range keys {
		if strings.HasPrefix(k, KindAddr+":") {
			return k
		}
	}
	return ""
}

func addrPrefixes(keys []string) []string {
	k := addrKey(keys)
	if k == "" {
		return nil
	}
	return Prefixes(k)
}

func nameKeys(keys []string) []string {
	var out []string
	for _, k := range keys {
		if strings.HasPrefix(k, KindName+":") {
			out = append(out, k)
		}
	}
	return out
}

func last(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[len(ss)-1]
}

func intersect(a, b []string) []string {
	seen := make(map[string]struct{}, len(a))
	for _, x := range a {
		seen[x] = struct{}{}
	}
	var out []string
	for _, x := range b {
		if _, ok := seen[x]; ok {
			out = append(out, x)
		}
	}
	return out
}
