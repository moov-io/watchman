package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecastToType_BusinessQueryScoresPerson(t *testing.T) {
	person := Entity[Value]{
		Name: "EMAD MIKHAIL TEWFIK BISHAI",
		Type: EntityPerson,
		Person: &Person{
			Name: "EMAD MIKHAIL TEWFIK BISHAI",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDNational, Country: "US", Identifier: "12345"},
			},
		},
	}.Normalize()

	// FollowTheMoney LegalEntity encoding: same record as a business.
	asBusiness := Entity[Value]{
		Name: person.Name,
		Type: EntityBusiness,
		Business: &Business{
			Name:          person.Name,
			GovernmentIDs: person.Person.GovernmentIDs,
		},
	}.Normalize()

	require.Equal(t, 0.0, func() float64 {
		// Vessel vs person stays a hard zero.
		vessel := Entity[Value]{Name: person.Name, Type: EntityVessel, Vessel: &Vessel{Name: person.Name}}.Normalize()
		return Similarity(vessel, person)
	}())

	got := Similarity(asBusiness, person)
	require.Greater(t, got, 0.80, "LegalEntity-as-business must score against a person instead of returning 0")
	require.InDelta(t, Similarity(person, person), got, 0.15)
}

func TestRecastToType_SameTypeUnchanged(t *testing.T) {
	e := Entity[Value]{Name: "Acme", Type: EntityBusiness, Business: &Business{Name: "Acme"}}.Normalize()
	got, ok := recastToType(e, EntityBusiness)
	require.True(t, ok)
	require.Equal(t, EntityBusiness, got.Type)
	require.Equal(t, e.Business, got.Business)
}

func TestRecastToType_VesselRejected(t *testing.T) {
	e := Entity[Value]{Name: "NS LEADER", Type: EntityPerson, Person: &Person{Name: "NS LEADER"}}.Normalize()
	_, ok := recastToType(e, EntityVessel)
	require.False(t, ok)
}
