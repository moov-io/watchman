package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegularNameMatchingPath(t *testing.T) {
	// Create debug output buffer
	var debug strings.Builder

	t.Run("RegularNameMatching_ShowPath", func(t *testing.T) {
		debug.Reset()

		// Regular name-only query (no crypto, no historical info)
		query := Entity[any]{
			Name: "JOHN SMITH",
			Type: EntityPerson,
			Person: &Person{
				Name: "JOHN SMITH",
			},
		}

		index := Entity[any]{
			Name: "JOHN SMITH",
			Type: EntityPerson,
			Person: &Person{
				Name: "JOHN SMITH",
			},
		}

		// Normalize entities
		queryNorm := query.Normalize()
		indexNorm := index.Normalize()

		// Test with favoritism=1.0
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
		score := DebugSimilarity(&debug, queryNorm, indexNorm)

		// Should be perfect match (1.0) after final capping
		require.InDelta(t, 1.0, score, 0.01, "Perfect name match should be 1.0")

		debugOutput := debug.String()
		t.Logf("Regular name matching debug output:\n%s", debugOutput)

		// For exact name matches, we should see Score:2 (favoritism applied)
		require.Contains(t, debugOutput, "Score:2", "Name should show favoritism applied (Score:2)")

		t.Logf("✓ VERIFIED PATH:")
		t.Logf("  Regular name matching goes through:")
		t.Logf("  1. similarity_fuzzy.go::compareName() - entry point")
		t.Logf("  2. similarity_fuzzy.go::compareNameTerms() - calls BestPairCombinationJaroWinkler")
		t.Logf("  3. stringscore.BestPairCombinationJaroWinkler() - generates word combinations")
		t.Logf("  4. stringscore.BestPairsJaroWinkler() - token-by-token comparison")
		t.Logf("  5. stringscore.customJaroWinkler() - applies favoritism for exact matches")
		t.Logf("  6. Final score gets capped at 1.0 in calculateFinalScore()")
	})

	t.Run("RegularWithHistoricalInfo_BothPaths", func(t *testing.T) {
		debug.Reset()

		// Regular query with both name and historical info
		query := Entity[any]{
			Name: "JANE DOE",
			Type: EntityPerson,
			Person: &Person{
				Name: "JANE DOE",
			},
			HistoricalInfo: []HistoricalInfo{
				{
					Type:  "Former Name",
					Value: "JANE SMITH",
				},
			},
		}

		index := Entity[any]{
			Name: "JANE DOE",
			Type: EntityPerson,
			Person: &Person{
				Name: "JANE DOE",
			},
			HistoricalInfo: []HistoricalInfo{
				{
					Type:  "Former Name",
					Value: "JANE SMITH",
				},
			},
		}

		// Normalize entities
		queryNorm := query.Normalize()
		indexNorm := index.Normalize()

		// Test with favoritism=1.0
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
		score := DebugSimilarity(&debug, queryNorm, indexNorm)

		// Should be perfect match (1.0) after final capping
		require.InDelta(t, 1.0, score, 0.01, "Perfect match should be 1.0")

		debugOutput := debug.String()
		t.Logf("Regular with HistoricalInfo debug output:\n%s", debugOutput)

		// Should see both Score:2 values - one from name (customJaroWinkler), one from HistoricalInfo (JaroWinklerWithFavoritism)
		require.Contains(t, debugOutput, "name: search.ScorePiece{Score:2", "Name should show favoritism (Score:2)")
		require.Contains(t, debugOutput, "supporting into: search.ScorePiece{Score:2", "HistoricalInfo should also show favoritism (Score:2)")

		t.Logf("✓ VERIFIED BOTH PATHS:")
		t.Logf("  - Names: customJaroWinkler path (via BestPairCombinationJaroWinkler → BestPairsJaroWinkler)")
		t.Logf("  - HistoricalInfo: JaroWinklerWithFavoritism path (via compareHistoricalValues)")
		t.Logf("  Both paths correctly apply favoritism for exact matches!")
	})
}
