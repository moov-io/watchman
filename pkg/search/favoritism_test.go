package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExactMatchFavoritism(t *testing.T) {
	// Create debug output buffer
	var debug strings.Builder

	// Test case 1: Exact name match WITH supporting fields should get favoritism
	query := Entity[any]{
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

	// Calculate similarity with debug output
	score := DebugSimilarity(&debug, queryNorm, indexNorm)

	// Check that score is 1.0 (perfect match)
	require.InDelta(t, 1.0, score, 0.01, "Expected exact match with supporting fields to get perfect score")

	// Check debug output contains favoritism application
	debugOutput := debug.String()
	t.Logf("Debug output:\n%s", debugOutput)

	// Test case 2: Exact name match WITHOUT sufficient supporting fields should NOT get favoritism

	queryNameOnly := Entity[any]{
		Name: "JOHN DOE",
		Type: EntityPerson,
		Person: &Person{
			Name: "JOHN DOE",
		},
		Addresses: []Address{
			{
				Line1:      "Oak Ave",
				City:       "Los Angeles",
				State:      "CA",
				PostalCode: "90210",
				Country:    "US",
			},
		},
	}

	indexNameOnly := Entity[any]{
		Name: "JANE DOE",
		Type: EntityPerson,
		Person: &Person{
			Name: "JANE DOE",
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

		// With EXACT_MATCH_FAVORITISM=0.0 (default), exact name matches will get penalties with no favoritism boost
		require.InDelta(t, 0.85, scoreNameOnly, 0.02, "Name-only exact match should get penalized without favoritism (default favoritism=0.0)")

		debugOutput2 := debug.String()
		t.Logf("Name-only debug output:\n%s", debugOutput2)
	})

	t.Run("EXACT_MATCH_FAVORITISM=1.0", func(t *testing.T) {
		debug.Reset()
		t.Setenv("EXACT_MATCH_FAVORITISM", "1.0")

		scoreNameOnly := DebugSimilarity(&debug, queryNameOnlyNorm, indexNameOnlyNorm)

		// With EXACT_MATCH_FAVORITISM=0.0 (default), exact name matches will get penalties with no favoritism boost
		require.InDelta(t, 0.85, scoreNameOnly, 0.02, "Name-only exact match should get penalized without favoritism (default favoritism=0.0)")

		debugOutput2 := debug.String()
		t.Logf("Name-only debug output:\n%s", debugOutput2)
	})

}
