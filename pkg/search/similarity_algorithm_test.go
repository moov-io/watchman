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

func TestSimilarity_AlgorithmSoftNGrams(t *testing.T) {
	// Transposition / doubled-letter pairs where bigram DP is the point of the paper.
	query := Entity[Value]{
		Name:   "Aleksandr",
		Type:   EntityPerson,
		Person: &Person{Name: "Aleksandr"},
	}.Normalize()
	index := Entity[Value]{
		Name:   "Alexander",
		Type:   EntityPerson,
		Person: &Person{Name: "Alexander"},
	}.Normalize()

	jw := SimilarityWithOpts(query, index, SimilarityOpts{Algorithm: AlgorithmJaroWinkler})
	bidist := SimilarityWithOpts(query, index, SimilarityOpts{Algorithm: AlgorithmSoftBidist})
	bisim := SimilarityWithOpts(query, index, SimilarityOpts{Algorithm: AlgorithmSoftBisim})

	require.Greater(t, jw, 0.0)
	require.Greater(t, bidist, 0.5, "soft-bidist should treat Aleksandr/Alexander as a close pair")
	require.Greater(t, bisim, 0.5, "soft-bisim should treat Aleksandr/Alexander as a close pair")
}
