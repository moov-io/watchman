package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountriesEqual(t *testing.T) {
	t.Run("ISO code vs official name", func(t *testing.T) {
		require.True(t, countriesEqual("RU", "Russia"))
		require.True(t, countriesEqual("Russia", "RU"))
		require.True(t, countriesEqual("ru", "RUSSIAN FEDERATION"))
		require.True(t, countriesEqual("US", "United States"))
		require.True(t, countriesEqual("USA", "US"))
	})

	t.Run("mismatched countries", func(t *testing.T) {
		require.False(t, countriesEqual("RU", "US"))
		require.False(t, countriesEqual("Russia", "Ukraine"))
	})

	t.Run("empty country", func(t *testing.T) {
		require.True(t, countriesEqual("", ""))
		require.False(t, countriesEqual("", "Russia"))
		require.False(t, countriesEqual("RU", ""))
	})
}

func TestCountryCompatible(t *testing.T) {
	t.Run("ISO code vs official name", func(t *testing.T) {
		require.True(t, countryCompatible("RU", "Russia"))
	})

	t.Run("mismatched countries", func(t *testing.T) {
		require.False(t, countryCompatible("RU", "US"))
		require.False(t, countryCompatible("Russia", "Ukraine"))
	})

	t.Run("empty country is compatible", func(t *testing.T) {
		require.True(t, countryCompatible("", "Russia"))
		require.True(t, countryCompatible("RU", ""))
		require.True(t, countryCompatible("", ""))
	})
}

func TestCompareIdentifiers_CountryNormalization(t *testing.T) {
	t.Run("ISO code vs official name is exact", func(t *testing.T) {
		got := compareIdentifiers("2018278055", "2018278055", "RU", "Russia")
		require.True(t, got.found)
		require.True(t, got.exact)
		require.True(t, got.hasCountry)
		require.Equal(t, 1.0, got.score)
	})

	t.Run("mismatched countries keep penalty", func(t *testing.T) {
		got := compareIdentifiers("2018278055", "2018278055", "RU", "US")
		require.True(t, got.found)
		require.False(t, got.exact)
		require.Equal(t, 0.7, got.score)
	})

	t.Run("empty country keeps current behavior", func(t *testing.T) {
		oneSided := compareIdentifiers("2018278055", "2018278055", "", "Russia")
		require.True(t, oneSided.found)
		require.False(t, oneSided.exact)
		require.Equal(t, 0.9, oneSided.score)

		neither := compareIdentifiers("2018278055", "2018278055", "", "")
		require.True(t, neither.found)
		require.True(t, neither.exact)
		require.False(t, neither.hasCountry)
		require.Equal(t, 1.0, neither.score)
	})
}

func TestIdentifierConflict_ISOCodeVsCountryName(t *testing.T) {
	query := Entity[Value]{
		Name: "Dmitry Khoroshev",
		Type: EntityPerson,
		Person: &Person{
			Name: "Dmitry Khoroshev",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "RU", Identifier: "2018278055"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Dmitry Yuryevich KHOROSHEV",
		Type: EntityPerson,
		Person: &Person{
			Name: "Dmitry Yuryevich KHOROSHEV",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "Russia", Identifier: "2018278055"},
			},
		},
	}.Normalize()

	require.False(t, hasIdentifierConflict(query, index))
	require.Equal(t, 1.0, Similarity(query, index))

	// Exact identifier piece also treats the pair as a unique-identity match
	got := comparePersonExactIDs(nil, query.Person, index.Person, 50)
	require.True(t, got.Matched)
	require.True(t, got.Exact)
	require.True(t, got.UniqueIdentity)
	require.Equal(t, 1.0, got.Score)
}

func TestIdentifierConflict_MismatchedCountriesStillConflict(t *testing.T) {
	query := Entity[Value]{
		Name: "Jane Doe",
		Type: EntityPerson,
		Person: &Person{
			Name: "Jane Doe",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "RU", Identifier: "2018278055"},
			},
		},
	}.Normalize()
	index := Entity[Value]{
		Name: "Jane Doe",
		Type: EntityPerson,
		Person: &Person{
			Name: "Jane Doe",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "US", Identifier: "2018278055"},
			},
		},
	}.Normalize()

	require.True(t, hasIdentifierConflict(query, index))
	require.Less(t, Similarity(query, index), 1.0)
}

func TestIdentifierConflict_EmptyCountryKeepsCurrentBehavior(t *testing.T) {
	withCountry := Entity[Value]{
		Name: "Jane Doe",
		Type: EntityPerson,
		Person: &Person{
			Name: "Jane Doe",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "Russia", Identifier: "2018278055"},
			},
		},
	}.Normalize()
	emptyCountry := Entity[Value]{
		Name: "Jane Doe",
		Type: EntityPerson,
		Person: &Person{
			Name: "Jane Doe",
			GovernmentIDs: []GovernmentID{
				{Type: GovernmentIDPassport, Country: "", Identifier: "2018278055"},
			},
		},
	}.Normalize()

	// Empty country is compatible, so a matching identifier is not a conflict.
	require.False(t, hasIdentifierConflict(withCountry, emptyCountry))

	// Scoring still applies the one-sided country penalty (0.9, not exact).
	got := compareIdentifiers("2018278055", "2018278055", "", "Russia")
	require.Equal(t, 0.9, got.score)
	require.False(t, got.exact)
}
