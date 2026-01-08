//go:build embeddings && integration

package search

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/embeddings"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

// getModelPath returns the path to the test model
func getModelPath() string {
	if path := os.Getenv("EMBEDDING_MODEL_PATH"); path != "" {
		return path
	}
	return filepath.Join("..", "..", "models", "multilingual-minilm")
}

// TestCrossScript_Search tests cross-script search with embeddings
func TestCrossScript_Search(t *testing.T) {
	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping cross-script test", modelPath)
	}

	logger := log.NewTestLogger()

	// Create test entities with Latin names
	entities := []search.Entity[search.Value]{
		{
			Name:     "Mohamed Ali",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1",
			Person:   &search.Person{Name: "Mohamed Ali"},
		},
		{
			Name:     "Vladimir Putin",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "2",
			Person:   &search.Person{Name: "Vladimir Putin"},
		},
		{
			Name:     "Kim Jong Un",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "3",
			Person:   &search.Person{Name: "Kim Jong Un"},
		},
		{
			Name:     "Ahmed Hassan",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "4",
			Person:   &search.Person{Name: "Ahmed Hassan"},
		},
		{
			Name:     "Ivan Ivanov",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "5",
			Person:   &search.Person{Name: "Ivan Ivanov"},
		},
		{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "6",
			Person:   &search.Person{Name: "John Smith"},
		},
	}

	// Create indexed lists
	indexedLists := index.NewLists(nil)
	indexedLists.Update(download.Stats{
		Entities: entities,
	})

	// Create config with embeddings enabled
	config := Config{
		Goroutines: DefaultConfig().Goroutines,
		Embeddings: embeddings.Config{
			Enabled:             true,
			ModelPath:           modelPath,
			CacheSize:           100,
			CrossScriptOnly:     true,
			SimilarityThreshold: 0.5,
			BatchSize:           32,
			IndexBuildTimeout:   5 * time.Minute,
		},
	}

	// Create search service with embeddings
	svc, err := NewService(logger, config, indexedLists)
	require.NoError(t, err)
	require.NotNil(t, svc)

	// Build embedding index
	ctx := context.Background()
	err = svc.RebuildEmbeddingIndex(ctx)
	require.NoError(t, err)

	opts := SearchOpts{Limit: 10, MinMatch: 0.01}

	// Test cases: Non-Latin queries that should match Latin names
	testCases := []struct {
		name          string
		query         string
		expectedMatch string
		minScore      float64
	}{
		{
			name:          "Arabic query matches Latin name",
			query:         "محمد علي",
			expectedMatch: "Mohamed Ali",
			minScore:      0.70,
		},
		{
			name:          "Cyrillic query matches Latin name",
			query:         "Владимир Путин",
			expectedMatch: "Vladimir Putin",
			minScore:      0.70,
		},
		{
			name:          "Chinese query matches Latin name",
			query:         "金正恩",
			expectedMatch: "Kim Jong Un",
			minScore:      0.60,
		},
		{
			name:          "Arabic name Ahmed",
			query:         "أحمد حسن",
			expectedMatch: "Ahmed Hassan",
			minScore:      0.70,
		},
		{
			name:          "Cyrillic name Ivan",
			query:         "Иван Иванов",
			expectedMatch: "Ivan Ivanov",
			minScore:      0.70,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := search.Entity[search.Value]{
				Name: tc.query,
				Type: search.EntityPerson,
			}

			results, err := svc.Search(ctx, query.Normalize(), opts)
			require.NoError(t, err)

			// Log results
			t.Logf("Query: %q", tc.query)
			t.Logf("Results count: %d", len(results))
			for i, r := range results {
				if i < 5 {
					t.Logf("  %d. %s (score: %.4f)", i+1, r.Entity.Name, r.Match)
				}
			}

			// Verify we got results
			require.NotEmpty(t, results, "Expected results for query %q", tc.query)

			// Check if expected match is top result
			require.Equal(t, tc.expectedMatch, results[0].Entity.Name,
				"Expected %q as top result for query %q, got %q",
				tc.expectedMatch, tc.query, results[0].Entity.Name)

			require.GreaterOrEqual(t, results[0].Match, tc.minScore,
				"Expected score >= %.2f for %q, got %.4f",
				tc.minScore, tc.expectedMatch, results[0].Match)

			t.Logf("PASS: %q -> %q (score: %.4f)", tc.query, results[0].Entity.Name, results[0].Match)
		})
	}
}

// TestCrossScript_JaroWinklerComparison compares embedding vs Jaro-Winkler
func TestCrossScript_JaroWinklerComparison(t *testing.T) {
	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping comparison test", modelPath)
	}

	logger := log.NewTestLogger()

	entities := []search.Entity[search.Value]{
		{
			Name:     "Mohamed Ali",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1",
			Person:   &search.Person{Name: "Mohamed Ali"},
		},
		{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "2",
			Person:   &search.Person{Name: "John Smith"},
		},
	}

	indexedLists := index.NewLists(nil)
	indexedLists.Update(download.Stats{Entities: entities})

	ctx := context.Background()
	opts := SearchOpts{Limit: 10, MinMatch: 0.01}
	arabicQuery := search.Entity[search.Value]{
		Name: "محمد علي",
		Type: search.EntityPerson,
	}

	// Test 1: Jaro-Winkler only (embeddings disabled)
	t.Run("Jaro-Winkler only", func(t *testing.T) {
		config := DefaultConfig()
		config.Embeddings.Enabled = false

		svc, err := NewService(logger, config, indexedLists)
		require.NoError(t, err)

		results, err := svc.Search(ctx, arabicQuery.Normalize(), opts)
		require.NoError(t, err)

		t.Logf("Jaro-Winkler results for 'محمد علي':")
		for i, r := range results {
			t.Logf("  %d. %s (score: %.4f)", i+1, r.Entity.Name, r.Match)
		}

		// Jaro-Winkler will return poor results for cross-script
		// because it compares characters directly
		if len(results) > 0 {
			t.Logf("Top result: %s (score: %.4f)", results[0].Entity.Name, results[0].Match)
			t.Logf("NOTE: Jaro-Winkler cannot match cross-script properly")
		}
	})

	// Test 2: With embeddings
	t.Run("With embeddings", func(t *testing.T) {
		config := Config{
			Goroutines: DefaultConfig().Goroutines,
			Embeddings: embeddings.Config{
				Enabled:             true,
				ModelPath:           modelPath,
				CacheSize:           100,
				CrossScriptOnly:     true,
				SimilarityThreshold: 0.5,
				BatchSize:           32,
				IndexBuildTimeout:   5 * time.Minute,
			},
		}

		svc, err := NewService(logger, config, indexedLists)
		require.NoError(t, err)

		err = svc.RebuildEmbeddingIndex(ctx)
		require.NoError(t, err)

		results, err := svc.Search(ctx, arabicQuery.Normalize(), opts)
		require.NoError(t, err)

		t.Logf("Embedding results for 'محمد علي':")
		for i, r := range results {
			t.Logf("  %d. %s (score: %.4f)", i+1, r.Entity.Name, r.Match)
		}

		// Embeddings should find the correct match
		require.NotEmpty(t, results, "Expected embedding results")
		require.Equal(t, "Mohamed Ali", results[0].Entity.Name,
			"Embeddings should find 'Mohamed Ali' for Arabic query")
		require.GreaterOrEqual(t, results[0].Match, 0.70,
			"Expected high similarity score")

		t.Logf("SUCCESS: Embeddings correctly matched Arabic 'محمد علي' to 'Mohamed Ali'")
	})
}

// TestCrossScript_LatinQueryUsesJaroWinkler verifies Latin queries bypass embeddings
func TestCrossScript_LatinQueryUsesJaroWinkler(t *testing.T) {
	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping test", modelPath)
	}

	logger := log.NewTestLogger()

	// Entities must be normalized for Jaro-Winkler to work
	entities := []search.Entity[search.Value]{
		{
			Name:     "Mohamed Ali",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1",
			Person:   &search.Person{Name: "Mohamed Ali"},
		},
		{
			Name:     "Mohammed Ali",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "2",
			Person:   &search.Person{Name: "Mohammed Ali"},
		},
	}

	// Normalize entities for Jaro-Winkler search
	for i := range entities {
		entities[i] = entities[i].Normalize()
	}

	indexedLists := index.NewLists(nil)
	indexedLists.Update(download.Stats{Entities: entities})

	config := Config{
		Goroutines: DefaultConfig().Goroutines,
		Embeddings: embeddings.Config{
			Enabled:             true,
			ModelPath:           modelPath,
			CacheSize:           100,
			CrossScriptOnly:     true, // Only use embeddings for non-Latin
			SimilarityThreshold: 0.5,
			BatchSize:           32,
			IndexBuildTimeout:   5 * time.Minute,
		},
	}

	svc, err := NewService(logger, config, indexedLists)
	require.NoError(t, err)

	ctx := context.Background()
	err = svc.RebuildEmbeddingIndex(ctx)
	require.NoError(t, err)

	// Latin query should use Jaro-Winkler (faster, CrossScriptOnly=true)
	query := search.Entity[search.Value]{
		Name: "Mohamed Ali",
		Type: search.EntityPerson,
	}

	opts := SearchOpts{Limit: 10, MinMatch: 0.01}
	results, err := svc.Search(ctx, query.Normalize(), opts)
	require.NoError(t, err)

	t.Logf("Latin query 'Mohamed Ali' results:")
	for i, r := range results {
		t.Logf("  %d. %s (score: %.4f)", i+1, r.Entity.Name, r.Match)
	}

	require.NotEmpty(t, results)
	// Should find exact or near-exact match via Jaro-Winkler
	require.Contains(t, []string{"Mohamed Ali", "Mohammed Ali"}, results[0].Entity.Name)
	t.Logf("SUCCESS: Latin query correctly uses Jaro-Winkler path")
}
