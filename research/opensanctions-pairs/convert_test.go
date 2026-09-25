// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestConvertEntity_Person(t *testing.T) {
	e := FTMEntity{
		ID:      "ofac-40604",
		Caption: "Aliasghar Norouzi",
		Schema:  "Person",
		Properties: map[string][]string{
			"name":           {"Aliasghar Norouzi", "علی اصغر نوروزی"},
			"alias":          {"Ali Asghar Norowzi"},
			"birthDate":      {"1962-11-11"},
			"nationality":    {"ir"},
			"passportNumber": {"Y53914915"},
			"gender":         {"male"},
		},
	}
	got := convertEntity(e, convertOpts{maxAltNames: 20}).Normalize()
	require.Equal(t, "Aliasghar Norouzi", got.Name)
	require.Equal(t, search.EntityPerson, got.Type)
	require.NotNil(t, got.Person)
	require.Equal(t, search.GenderMale, got.Person.Gender)
	require.NotNil(t, got.Person.BirthDate)
	require.Equal(t, 1962, got.Person.BirthDate.Year())
	require.Equal(t, "Iran", got.Person.GovernmentIDs[0].Country)
	require.Equal(t, "Y53914915", got.Person.GovernmentIDs[0].Identifier)
	require.Contains(t, got.Person.AltNames, "علی اصغر نوروزی")
}

func TestConvertPair_OrgLikeTypesAlign(t *testing.T) {
	p := Pair{
		Left: FTMEntity{
			Caption:    "Fuel Co",
			Schema:     "Company",
			Properties: map[string][]string{"name": {"Fuel Co"}, "innCode": {"7811550255"}, "jurisdiction": {"ru"}},
		},
		Right: FTMEntity{
			Caption:    "Fuel Company LLC",
			Schema:     "Organization",
			Properties: map[string][]string{"name": {"Fuel Company LLC"}, "innCode": {"7811550255"}, "jurisdiction": {"ru"}},
		},
		Judgement: "positive",
	}
	left, right := convertPair(p, convertOpts{maxAltNames: 5})
	require.Equal(t, search.EntityBusiness, left.Type)
	require.Equal(t, search.EntityBusiness, right.Type)
	require.NotNil(t, left.Business)
	require.Equal(t, "7811550255", left.Business.GovernmentIDs[0].Identifier)
}

func TestFilterPairs_SubjectsOnly(t *testing.T) {
	pairs := []Pair{
		{Left: FTMEntity{Schema: "Person"}, Right: FTMEntity{Schema: "Person"}, Judgement: "positive"},
		{Left: FTMEntity{Schema: "Occupancy"}, Right: FTMEntity{Schema: "Occupancy"}, Judgement: "positive"},
	}
	got := filterPairs(pairs, true, nil, 0, 0)
	require.Len(t, got, 1)
	require.Equal(t, "Person", got[0].Left.Schema)
}

func TestPickPrimaryName_IgnoresSchemaCaption(t *testing.T) {
	e := FTMEntity{
		Caption: "Company",
		Schema:  "Company",
		Properties: map[string][]string{
			"name": {"Vasha Toplivnaya Kompaniya LLC"},
		},
	}
	require.Equal(t, "Vasha Toplivnaya Kompaniya LLC", pickPrimaryName(e))
}

func TestQueryTypes_LegalEntityFansOut(t *testing.T) {
	le := FTMEntity{Schema: "LegalEntity", Caption: "AMANDA MARIE ORLOSKI"}
	require.Equal(t, []search.EntityType{search.EntityPerson, search.EntityBusiness}, queryTypes(le))
	person := FTMEntity{Schema: "Person", Caption: "Amanda Orloski"}
	require.Equal(t, []search.EntityType{search.EntityPerson}, commonTypes(le, person))
	company := FTMEntity{Schema: "Company", Caption: "Acme"}
	require.Equal(t, []search.EntityType{search.EntityBusiness}, commonTypes(le, company))
	require.Empty(t, commonTypes(person, company))
}

func TestSimilarityFanout_LegalEntityVsPerson(t *testing.T) {
	p := Pair{
		Left: FTMEntity{
			Caption:    "BISHAI EMAD MIKHAIL TEWFIK",
			Schema:     "LegalEntity",
			Properties: map[string][]string{"name": {"BISHAI EMAD MIKHAIL TEWFIK"}},
		},
		Right: FTMEntity{
			Caption:    "EMAD MIKHAIL TEWFIK BISHAI",
			Schema:     "Person",
			Properties: map[string][]string{"name": {"EMAD MIKHAIL TEWFIK BISHAI"}},
		},
		Judgement: "positive",
	}
	left, right := convertPair(p, convertOpts{maxAltNames: 5})
	require.Equal(t, search.EntityBusiness, left.Type)
	require.Equal(t, search.EntityPerson, right.Type)

	got := similarityFanout(p, convertOpts{maxAltNames: 5}, search.SimilarityOpts{})
	require.Greater(t, got, 0.5, "type fan-out scores LegalEntity as person")
}

func TestRightIDs_ExtractsOFAC(t *testing.T) {
	ids := rightIDs(FTMEntity{
		ID:        "NK-abc",
		Referents: []string{"ofac-40604", "gb-hmt-15851"},
	})
	_, hasOFAC := ids["40604"]
	_, hasUK := ids["15851"]
	require.True(t, hasOFAC)
	require.True(t, hasUK)
}
