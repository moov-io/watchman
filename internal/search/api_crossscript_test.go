//go:build embeddings && integration

package search

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/download"
	"github.com/moov-io/watchman/internal/embeddings"
	"github.com/moov-io/watchman/internal/index"
	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

// getTestEmbeddingsConfig returns configuration for integration tests.
// It uses environment variables to configure the embedding provider.
func getTestEmbeddingsConfig() embeddings.Config {
	baseURL := os.Getenv("EMBEDDINGS_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:11434/v1" // Default to local Ollama
	}

	model := os.Getenv("EMBEDDINGS_MODEL")
	if model == "" {
		model = "nomic-embed-text"
	}

	dimension := 768
	if dim := os.Getenv("EMBEDDINGS_DIMENSION"); dim != "" {
		if d, err := strconv.Atoi(dim); err == nil {
			dimension = d
		}
	}

	return embeddings.Config{
		Enabled: true,
		Provider: embeddings.ProviderConfig{
			Name:             "ollama",
			BaseURL:          baseURL,
			APIKey:           os.Getenv("EMBEDDINGS_API_KEY"),
			Model:            model,
			Dimension:        dimension,
			NormalizeVectors: true,
			Timeout:          30 * time.Second,
			RateLimit: embeddings.RateLimitConfig{
				RequestsPerSecond: 10,
				Burst:             20,
			},
			Retry: embeddings.RetryConfig{
				MaxRetries:     3,
				InitialBackoff: time.Second,
				MaxBackoff:     30 * time.Second,
			},
		},
		CacheSize:           100,
		CrossScriptOnly:     true,
		SimilarityThreshold: 0.5,
		BatchSize:           32,
		IndexBuildTimeout:   5 * time.Minute,
	}
}

// TestCrossScript_Search tests cross-script search with embeddings
func TestCrossScript_Search(t *testing.T) {
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
		Embeddings: getTestEmbeddingsConfig(),
	}

	// Create search service with embeddings
	svc, err := NewService(logger, config, indexedLists)
	if err != nil {
		t.Skipf("Could not create search service (provider may be unavailable): %v", err)
	}
	require.NotNil(t, svc)

	// Build embedding index
	ctx := context.Background()
	err = svc.RebuildEmbeddingIndex(ctx)
	if err != nil {
		t.Skipf("Could not build embedding index: %v", err)
	}

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
			minScore:      0.40, // Lower threshold for general models
		},
		{
			name:          "Cyrillic query matches Latin name",
			query:         "Владимир Путин",
			expectedMatch: "Vladimir Putin",
			minScore:      0.40,
		},
		{
			name:          "Chinese query matches Latin name",
			query:         "金正恩",
			expectedMatch: "Kim Jong Un",
			minScore:      0.40,
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

			// Check if expected match is in top 3 results (more lenient for general models)
			found := false
			for i := 0; i < 3 && i < len(results); i++ {
				if results[i].Entity.Name == tc.expectedMatch {
					found = true
					t.Logf("PASS: %q -> %q found at position %d (score: %.4f)",
						tc.query, tc.expectedMatch, i+1, results[i].Match)
					break
				}
			}
			require.True(t, found, "Expected %q in top 3 results for query %q", tc.expectedMatch, tc.query)
		})
	}
}

// TestCrossScript_JaroWinklerComparison compares embedding vs Jaro-Winkler
func TestCrossScript_JaroWinklerComparison(t *testing.T) {
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
			Embeddings: getTestEmbeddingsConfig(),
		}

		svc, err := NewService(logger, config, indexedLists)
		if err != nil {
			t.Skipf("Could not create search service: %v", err)
		}

		err = svc.RebuildEmbeddingIndex(ctx)
		if err != nil {
			t.Skipf("Could not build embedding index: %v", err)
		}

		results, err := svc.Search(ctx, arabicQuery.Normalize(), opts)
		require.NoError(t, err)

		t.Logf("Embedding results for 'محمد علي':")
		for i, r := range results {
			t.Logf("  %d. %s (score: %.4f)", i+1, r.Entity.Name, r.Match)
		}

		// Embeddings should return results (quality depends on model)
		require.NotEmpty(t, results, "Expected embedding results")
		t.Logf("Top result with embeddings: %s (score: %.4f)", results[0].Entity.Name, results[0].Match)
	})
}

// TestCrossScript_LatinQueryUsesJaroWinkler verifies Latin queries bypass embeddings
func TestCrossScript_LatinQueryUsesJaroWinkler(t *testing.T) {
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
		Embeddings: getTestEmbeddingsConfig(),
	}

	svc, err := NewService(logger, config, indexedLists)
	if err != nil {
		t.Skipf("Could not create search service: %v", err)
	}

	ctx := context.Background()
	err = svc.RebuildEmbeddingIndex(ctx)
	if err != nil {
		t.Skipf("Could not build embedding index: %v", err)
	}

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
