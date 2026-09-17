package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimilarity_AlgorithmSoundexBoostsPhoneticNames(t *testing.T) {
	query := Entity[Value]{
		Name:   "Smythe",
		Type:   EntityPerson,
		Person: &Person{Name: "Smythe"},
	}.Normalize()
	index := Entity[Value]{
		Name:   "Smith",
		Type:   EntityPerson,
		Person: &Person{Name: "Smith"},
	}.Normalize()

	jw := SimilarityWithOpts(query, index, SimilarityOpts{Algorithm: AlgorithmJaroWinkler})
	sx := SimilarityWithOpts(query, index, SimilarityOpts{Algorithm: AlgorithmSoundex})

	require.Greater(t, jw, 0.0)
	require.Greater(t, sx, jw, "soundex should boost phonetically similar names")
}
