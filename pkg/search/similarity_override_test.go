package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExactOverride_TaxIDDoesNotForceOneWhenNamesDisagree(t *testing.T) {
	query := Entity[Value]{
		Name: "Aforra Development",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Aforra Development",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Russia", Identifier: "1234567890"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Aforra Property",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Aforra Property",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Russia", Identifier: "1234567890"},
			},
		},
	}.Normalize()

	got := Similarity(query, index)
	require.Less(t, got, 1.0, "shared tax ID with disagreeing names must not short-circuit to 1.0")
	require.Greater(t, got, 0.3)
	require.InDelta(t, got, scoreSimilarityFast(query, index, SimilarityOpts{}), 0.001)
}

func TestExactOverride_TaxIDIsEvidenceNotIdentity(t *testing.T) {
	query := Entity[Value]{
		Name: "Dialog Regions",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Dialog Regions",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Russia", Identifier: "9709063550"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Dialog Regions",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Dialog Regions",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Russia", Identifier: "9709063550"},
			},
		},
	}.Normalize()

	got := Similarity(query, index)
	require.Less(t, got, 1.0, "matching tax ID is weighted evidence, not a 1.0 identity key")
	require.Greater(t, got, 0.85)
}

func TestExactOverride_PassportForcesOneDespiteName(t *testing.T) {
	query := Entity[Value]{
		Name: "Aliasghar Norouzi",
		Type: EntityPerson,
		Person: &Person{
			Name: "Aliasghar Norouzi",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "Iran", Identifier: "Y53914915"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Completely Different Person",
		Type: EntityPerson,
		Person: &Person{
			Name: "Completely Different Person",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "Iran", Identifier: "Y53914915"},
			},
		},
	}.Normalize()

	got := Similarity(query, index)
	require.InDelta(t, 1.0, got, 0.001)
	require.InDelta(t, 1.0, scoreSimilarityFast(query, index, SimilarityOpts{}), 0.001)
}

func TestExactOverride_IMOForcesOneDespiteName(t *testing.T) {
	query := Entity[Value]{
		Name: "NS LEADER",
		Type: EntityVessel,
		Vessel: &Vessel{
			Name:      "NS LEADER",
			IMONumber: "9339301",
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "RENAMED TANKER",
		Type: EntityVessel,
		Vessel: &Vessel{
			Name:      "RENAMED TANKER",
			IMONumber: "9339301",
		},
	}.Normalize()

	got := Similarity(query, index)
	require.InDelta(t, 1.0, got, 0.001)
}

func TestExactOverride_CountryMismatchDoesNotForceOne(t *testing.T) {
	query := Entity[Value]{
		Name: "Same Name Co",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Same Name Co",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Russia", Identifier: "111"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Same Name Co",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Same Name Co",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDTax, Country: "Ukraine", Identifier: "111"},
			},
		},
	}.Normalize()

	got := Similarity(query, index)
	require.Less(t, got, 1.0, "same identifier with disagreeing countries must not override")
}

func TestExactOverride_ContactDoesNotForceOne(t *testing.T) {
	query := Entity[Value]{
		Name: "Alpha LLC",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Alpha LLC",
		},
		Contact: ContactInfo{EmailAddresses: []string{"office@example.com"}},
	}.Normalize()
	index := Entity[Value]{
		Name: "Beta GmbH",
		Type: EntityBusiness,
		Business: &Business{
			Name: "Beta GmbH",
		},
		Contact: ContactInfo{EmailAddresses: []string{"office@example.com"}},
	}.Normalize()

	got := Similarity(query, index)
	require.Less(t, got, 1.0, "shared email must not force 1.0")
}

func TestShouldExactOverride(t *testing.T) {
	require.False(t, shouldExactOverride(nil))
	require.True(t, shouldExactOverride([]ScorePiece{
		{PieceType: "crypto-exact", Exact: true, FieldsCompared: 1, UniqueIdentity: true, Score: 1},
	}))
	require.False(t, shouldExactOverride([]ScorePiece{
		{PieceType: "identifiers", Exact: true, FieldsCompared: 1, UniqueIdentity: false, Score: 1},
		{PieceType: "name", FieldsCompared: 1, Score: 0.29},
	}))
	require.False(t, shouldExactOverride([]ScorePiece{
		{PieceType: "identifiers", Exact: true, FieldsCompared: 1, UniqueIdentity: false, Score: 1},
		{PieceType: "name", FieldsCompared: 1, Score: 0.90},
	}), "tax/registration must not override even when names agree")
	require.True(t, shouldExactOverride([]ScorePiece{
		{PieceType: "identifiers", Exact: true, FieldsCompared: 1, UniqueIdentity: true, Score: 1},
		{PieceType: "name", FieldsCompared: 1, Score: 0.10},
	}))
}

func TestIdentifierConflict_NationalIDsDiffer(t *testing.T) {
	sameName := Entity[Value]{
		Name: "Khalid Mehmood",
		Type: EntityPerson,
		Person: &Person{
			Name: "Khalid Mehmood",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDNational, Country: "Pakistan", Identifier: "35201114139885"},
			},
		},
	}.Normalize()
	other := Entity[Value]{
		Name: "Khalid Mehmood",
		Type: EntityPerson,
		Person: &Person{
			Name: "Khalid Mehmood",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDNational, Country: "Pakistan", Identifier: "3710502620181"},
			},
		},
	}.Normalize()

	require.True(t, hasIdentifierConflict(sameName, other))
	withConflict := Similarity(sameName, other)

	matching := other
	matching.Person.GovernmentIDs[0].Identifier = "35201114139885"
	matching = matching.Normalize()
	withoutConflict := Similarity(sameName, matching)
	require.Less(t, withConflict, withoutConflict)
	require.Less(t, withConflict, 0.80, "same name + conflicting national IDs should not look like a strong match")
}

func TestIdentifierConflict_MissingIDIsNotConflict(t *testing.T) {
	withID := Entity[Value]{
		Name: "Jane Doe",
		Type: EntityPerson,
		Person: &Person{
			Name: "Jane Doe",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "US", Identifier: "A123"},
			},
		},
	}.Normalize()
	noID := Entity[Value]{
		Name:   "Jane Doe",
		Type:   EntityPerson,
		Person: &Person{Name: "Jane Doe"},
	}.Normalize()
	require.False(t, hasIdentifierConflict(withID, noID))
}
