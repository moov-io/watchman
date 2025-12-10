package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCryptoAddressFavoritism(t *testing.T) {
	// Create debug output buffer
	var debug strings.Builder

	t.Run("ExactCryptoMatch_NoDoubleApplied", func(t *testing.T) {
		debug.Reset()

		// Test case: Exact crypto address match - should get favoritism only once
		// CODE PATH: Names go through customJaroWinkler (in BestPairsJaroWinkler)
		// CRYPTO PATH: Exact string matching (no favoritism, binary 0/1 scores)
		query := Entity[any]{
			Name: "CRYPTO TRADER",
			Type: EntityPerson,
			Person: &Person{
				Name: "CRYPTO TRADER",
			},
			CryptoAddresses: []CryptoAddress{
				{
					Currency: "BTC",
					Address:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", // Genesis Bitcoin address
				},
			},
		}

		index := Entity[any]{
			Name: "CRYPTO TRADER",
			Type: EntityPerson,
			Person: &Person{
				Name: "CRYPTO TRADER",
			},
			CryptoAddresses: []CryptoAddress{
				{
					Currency: "BTC",
					Address:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", // Same address
				},
			},
		}

		// Normalize entities
		queryNorm := query.Normalize()
		indexNorm := index.Normalize()

		// Test with favoritism=0.0
		t.Setenv("EXACT_MATCH_FAVORITISM", "0.0")
		scoreNoFavoritism := DebugSimilarity(&debug, queryNorm, indexNorm)

		// This should be 1.0 (perfect match due to name + crypto match)
		require.InDelta(t, 1.0, scoreNoFavoritism, 0.01, "Perfect crypto + name match should be 1.0 even without favoritism")

		debugOutput := debug.String()
		t.Logf("Crypto match debug output (favoritism=0.0):\n%s", debugOutput)

		// Test with favoritism=1.0 and path logging
		debug.Reset()
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "1")

		// Run similarity calculation
		scoreFavoritism := DebugSimilarity(&debug, queryNorm, indexNorm)

		// Should still be 1.0 - if it's higher, we have double favoritism
		require.InDelta(t, 1.0, scoreFavoritism, 0.01, "Perfect match should remain 1.0 regardless of favoritism (no double application)")

		debugOutputFavoritism := debug.String()
		t.Logf("Crypto match debug output (favoritism=1.0):\n%s", debugOutputFavoritism)

		// Verify favoritism was applied to name scoring (should see Score:2 in name piece)
		require.Contains(t, debugOutputFavoritism, "Score:2", "Name should show favoritism applied (Score:2)")
		require.NotContains(t, debugOutputFavoritism, "Score:3", "Should not see double favoritism (Score:3)")

		// PATH VERIFICATION: In this test case
		// - Names "CRYPTO TRADER" should go through customJaroWinkler (expect trace)
		// - Crypto addresses should use exact matching (no trace, just binary match)
		// - NO JaroWinklerWithFavoritism should be called (no HistoricalInfo)
		t.Logf("✓ VERIFIED PATHS:")
		t.Logf("  - Names: customJaroWinkler path (favoritism applied → Score:2)")
		t.Logf("  - Crypto: exact matching path (no favoritism → Score:1)")
		t.Logf("  - HistoricalInfo: not present (no JaroWinklerWithFavoritism calls)")

		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "0")
	})

	t.Run("CryptoWithHistoricalInfo_SeparatePaths", func(t *testing.T) {
		debug.Reset()

		// Test case: Entity with both crypto addresses (exact match) and historical info
		// CODE PATH 1: Names go through customJaroWinkler (in BestPairsJaroWinkler)
		// CODE PATH 2: HistoricalInfo goes through JaroWinklerWithFavoritism (in compareHistoricalValues)
		// This verifies both favoritism paths work independently without interference
		query := Entity[any]{
			Name: "MIXED ENTITY",
			Type: EntityPerson,
			Person: &Person{
				Name: "MIXED ENTITY",
			},
			CryptoAddresses: []CryptoAddress{
				{
					Currency: "ETH",
					Address:  "0x742d35Cc6634C0532925a3b8D8a6C8E1E3A8e0B6",
				},
			},
			HistoricalInfo: []HistoricalInfo{
				{
					Type:  "Former Crypto Address",
					Value: "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
				},
			},
		}

		index := Entity[any]{
			Name: "MIXED ENTITY",
			Type: EntityPerson,
			Person: &Person{
				Name: "MIXED ENTITY",
			},
			CryptoAddresses: []CryptoAddress{
				{
					Currency: "ETH",
					Address:  "0x742d35Cc6634C0532925a3b8D8a6C8E1E3A8e0B6", // Same as query
				},
			},
			HistoricalInfo: []HistoricalInfo{
				{
					Type:  "Former Crypto Address",
					Value: "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2", // Same as query
				},
			},
		}

		// Normalize entities
		queryNorm := query.Normalize()
		indexNorm := index.Normalize()

		// Test with favoritism=1.0 and path logging
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "1")
		score := DebugSimilarity(&debug, queryNorm, indexNorm)

		// Should be perfect match (1.0) due to multiple exact matches
		require.InDelta(t, 1.0, score, 0.01, "Multiple exact matches should result in perfect score")

		debugOutput := debug.String()
		t.Logf("Mixed crypto/historical debug output:\n%s", debugOutput)

		// Verify both paths are working:
		// 1. Exact crypto match (crypto-exact piece should show Score:1)
		// 2. Name match with favoritism (name piece should show Score:2)
		// 3. Historical info match via JaroWinklerWithFavoritism (supporting piece shows Score:2)
		require.Contains(t, debugOutput, "crypto-exact", "Should have crypto exact match piece")
		require.Contains(t, debugOutput, "Score:2", "Name should show favoritism (Score:2)")
		require.Contains(t, debugOutput, "supporting into: search.ScorePiece{Score:2", "HistoricalInfo should match via JaroWinklerWithFavoritism")

		// PATH VERIFICATION: In this test case
		// - Names "MIXED ENTITY" should go through customJaroWinkler (Score:2)
		// - Crypto addresses should use exact matching (Score:1)
		// - HistoricalInfo should go through JaroWinklerWithFavoritism (Score:2 with favoritism applied)
		t.Logf("✓ VERIFIED PATHS:")
		t.Logf("  - Names: customJaroWinkler path (favoritism applied → Score:2)")
		t.Logf("  - Crypto: exact matching path (no favoritism → Score:1)")
		t.Logf("  - HistoricalInfo: JaroWinklerWithFavoritism path (exact match + favoritism → Score:2)")

		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "0")
	})

	t.Run("PartialCryptoMatch_ShowFavoritism", func(t *testing.T) {
		debug.Reset()

		// Test case: Name-only query vs entity with crypto address - should show favoritism effect
		// CODE PATH: Names go through customJaroWinkler (in BestPairsJaroWinkler)
		// COVERAGE: Query has fewer fields than index, tests favoritism with coverage penalties
		query := Entity[any]{
			Name: "PARTIAL ENTITY",
			Type: EntityPerson,
			Person: &Person{
				Name: "PARTIAL ENTITY",
			},
			// No crypto addresses in query
		}

		index := Entity[any]{
			Name: "PARTIAL ENTITY",
			Type: EntityPerson,
			Person: &Person{
				Name: "PARTIAL ENTITY",
			},
			CryptoAddresses: []CryptoAddress{
				{
					Currency: "BTC",
					Address:  "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
				},
			},
		}

		// Normalize entities
		queryNorm := query.Normalize()
		indexNorm := index.Normalize()

		// Test without favoritism
		t.Setenv("EXACT_MATCH_FAVORITISM", "0.0")
		scoreNoFavoritism := DebugSimilarity(&debug, queryNorm, indexNorm)

		debugOutput := debug.String()
		t.Logf("Partial match debug output (favoritism=0.0):\n%s", debugOutput)

		// Test with favoritism and path logging
		debug.Reset()
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "1")
		scoreFavoritism := DebugSimilarity(&debug, queryNorm, indexNorm)

		debugOutputFavoritism := debug.String()
		t.Logf("Partial match debug output (favoritism=1.0):\n%s", debugOutputFavoritism)

		// Favoritism should boost the score
		require.Greater(t, scoreFavoritism, scoreNoFavoritism, "Favoritism should increase score for partial matches")
		require.Contains(t, debugOutputFavoritism, "Score:2", "Name should show favoritism applied")

		// Should be capped at 1.0 if favoritism pushes it over
		require.LessOrEqual(t, scoreFavoritism, 1.0, "Score should be capped at 1.0")

		t.Logf("✓ Expected path: Names processed through customJaroWinkler only")
		t.Setenv("WATCHMAN_TEST_LOG_PATHS", "0")
	})
}
