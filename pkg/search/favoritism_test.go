package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExactMatchFavoritism(t *testing.T) {
	// Create debug output buffer
	var debug strings.Builder

	// Test case 1: Exact name match WITH supporting fields but non-matching gov ID
	query := Entity[any]{
		Name: "JOHN SMITH",
		Type: EntityPerson,
		Person: &Person{
			Name: "JOHN SMITH",
		},
		Addresses: []Address{
			{
				Line1:      "123",
				City:       "",
				State:      "NY",
				PostalCode: "",
				Country:    "US",
			},
		},
	}

	index := Entity[any]{
		Name: "JOHN SMITH",
		Type: EntityPerson,
		Person: &Person{
			Name: "JOHN SMITH",
			GovernmentIDs: []GovernmentID{
				{
					Type:       GovernmentIDPassport,
					Country:    "US",
					Identifier: "123456789",
				},
			},
		},
		Addresses: []Address{
			{
				Line1:      "123 Main St",
				City:       "New York",
				State:      "NY",
				PostalCode: "10001",
				Country:    "US",
			},
		},
	}

	// Normalize entities
	queryNorm := query.Normalize()
	indexNorm := index.Normalize()

	// Test with favoritism=0.0 - should get less than 1.0 due to non-matching address
	t.Setenv("EXACT_MATCH_FAVORITISM", "0.0")
	score := DebugSimilarity(&debug, queryNorm, indexNorm)

	// Check that score is less than 1.0 (exact name but slightly mismatched address)
	require.Less(t, score, 0.9, "Expected near-perfect match with mismatched address to get < 0.9 without favoritism")
	t.Logf("Actual score with close gov ID: %.4f", score)
	require.Greater(t, score, 0.8, "Score should still be reasonable with exact name match")

	// Check debug output
	debugOutput := debug.String()
	t.Logf("Rich query debug output (favoritism=0.0):\n%s", debugOutput)

	// Test with favoritism=1.0 - should get higher score due to name favoritism boost
	debug.Reset()
	t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")
	scoreFavoritism := DebugSimilarity(&debug, queryNorm, indexNorm)

	require.Greater(t, scoreFavoritism, 0.95, "Favoritism should boost the score higher than without favoritism")
	debugOutputFavoritism := debug.String()
	t.Logf("Rich query debug output (favoritism=1.0):\n%s", debugOutputFavoritism)

	// Test case 2: Name-only query with exact match to test favoritism effect

	queryNameOnly := Entity[any]{
		Name: "JOHN DOE",
		Type: EntityPerson,
		Person: &Person{
			Name: "JOHN DOE",
		},
	}

	indexNameOnly := Entity[any]{
		Name: "JOHN DOE",
		Type: EntityPerson,
		Person: &Person{
			Name: "JOHN DOE",
		},
		Addresses: []Address{
			{
				Line1:      "456 Oak Ave",
				City:       "Los Angeles",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},
	}

	queryNameOnlyNorm := queryNameOnly.Normalize()
	indexNameOnlyNorm := indexNameOnly.Normalize()

	t.Run("EXACT_MATCH_FAVORITISM=0.0", func(t *testing.T) {
		debug.Reset()
		t.Setenv("EXACT_MATCH_FAVORITISM", "0.0")

		scoreNameOnly := DebugSimilarity(&debug, queryNameOnlyNorm, indexNameOnlyNorm)

		// With EXACT_MATCH_FAVORITISM=0.0 (default), name-only exact matches get penalties (0.85)
		require.InDelta(t, 0.85, scoreNameOnly, 0.02, "Name-only exact match should get penalized without favoritism (default favoritism=0.0)")

		debugOutput2 := debug.String()
		t.Logf("Name-only debug output:\n%s", debugOutput2)
	})

	t.Run("EXACT_MATCH_FAVORITISM=1.0", func(t *testing.T) {
		debug.Reset()
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")

		scoreNameOnly := DebugSimilarity(&debug, queryNameOnlyNorm, indexNameOnlyNorm)

		// With EXACT_MATCH_FAVORITISM=1.0, name-only exact matches get favoritism boost then capped at 1.0
		require.InDelta(t, 1.0, scoreNameOnly, 0.01, "Name-only exact match with favoritism=1.0 should get boosted and capped at 1.0")

		debugOutput2 := debug.String()
		t.Logf("Name-only debug output with favoritism=1.0:\n%s", debugOutput2)
	})

}
