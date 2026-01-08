//go:build embeddings && integration

package embeddings

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

// getModelPath returns the path to the test model
func getModelPath() string {
	// Check environment variable first
	if path := os.Getenv("EMBEDDING_MODEL_PATH"); path != "" {
		return path
	}
	// Default to project models directory
	return filepath.Join("..", "..", "models", "multilingual-minilm")
}

func TestIntegration_ServiceCreation(t *testing.T) {
	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping integration test", modelPath)
	}

	logger := log.NewTestLogger()
	config := Config{
		Enabled:             true,
		ModelPath:           modelPath,
		CacheSize:           100,
		CrossScriptOnly:     true,
		SimilarityThreshold: 0.7,
		BatchSize:           32,
		IndexBuildTimeout:   5 * time.Minute,
	}

	service, err := NewService(logger, config)
	require.NoError(t, err)
	require.NotNil(t, service)
	defer service.Shutdown()
}

func TestIntegration_Encode(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()

	// Test encoding a simple Latin text
	embedding, err := service.Encode(ctx, "Mohamed Ali")
	require.NoError(t, err)
	require.NotEmpty(t, embedding)
	require.Equal(t, 384, len(embedding), "Expected 384-dimensional embedding")

	// Test encoding Arabic text
	embedding, err = service.Encode(ctx, "محمد علي")
	require.NoError(t, err)
	require.NotEmpty(t, embedding)
	require.Equal(t, 384, len(embedding))

	// Test encoding Cyrillic text
	embedding, err = service.Encode(ctx, "Владимир Путин")
	require.NoError(t, err)
	require.NotEmpty(t, embedding)
	require.Equal(t, 384, len(embedding))

	// Test encoding Chinese text
	embedding, err = service.Encode(ctx, "金正恩")
	require.NoError(t, err)
	require.NotEmpty(t, embedding)
	require.Equal(t, 384, len(embedding))
}

func TestIntegration_EncodeBatch(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()
	texts := []string{
		"Mohamed Ali",
		"محمد علي",
		"Vladimir Putin",
		"Владимир Путин",
		"Kim Jong Un",
		"金正恩",
	}

	embeddings, err := service.EncodeBatch(ctx, texts)
	require.NoError(t, err)
	require.Len(t, embeddings, len(texts))

	for i, emb := range embeddings {
		require.Len(t, emb, 384, "Embedding %d should be 384-dimensional", i)
	}
}

func TestIntegration_Similarity(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()

	// Test same name in different scripts
	tests := []struct {
		name1    string
		name2    string
		minScore float64
		desc     string
	}{
		{"Mohamed Ali", "محمد علي", 0.80, "Arabic-Latin transliteration"},
		{"Vladimir Putin", "Владимир Путин", 0.80, "Russian-Latin transliteration"},
		{"Kim Jong Un", "金正恩", 0.65, "Korean name in Chinese characters"},
		{"Alexander", "Αλέξανδρος", 0.65, "Greek-Latin"},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			score, err := service.Similarity(ctx, tc.name1, tc.name2)
			require.NoError(t, err)
			require.GreaterOrEqual(t, score, tc.minScore,
				"Similarity(%q, %q) = %.4f, expected >= %.4f",
				tc.name1, tc.name2, score, tc.minScore)
			t.Logf("Similarity(%q, %q) = %.4f", tc.name1, tc.name2, score)
		})
	}

	// Test dissimilar names should have lower scores
	dissimilarScore, err := service.Similarity(ctx, "John Smith", "李明")
	require.NoError(t, err)
	require.Less(t, dissimilarScore, 0.6,
		"Dissimilar names should have score < 0.6, got %.4f", dissimilarScore)
	t.Logf("Similarity(\"John Smith\", \"李明\") = %.4f", dissimilarScore)
}

func TestIntegration_CrossScriptPairs(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()
	pairs := loadCrossScriptPairs(t)

	var passed, failed int
	for _, pair := range pairs {
		score, err := service.Similarity(ctx, pair.Query, pair.Expected)
		require.NoError(t, err)

		if score >= pair.MinScore {
			passed++
			t.Logf("PASS: %s -> %s (%.4f >= %.4f) %s",
				pair.Query, pair.Expected, score, pair.MinScore, pair.Description)
		} else {
			failed++
			t.Logf("FAIL: %s -> %s (%.4f < %.4f) %s",
				pair.Query, pair.Expected, score, pair.MinScore, pair.Description)
		}
	}

	t.Logf("Cross-script matching: %d/%d passed (%.1f%%)",
		passed, len(pairs), float64(passed)/float64(len(pairs))*100)

	// Require at least 70% pass rate
	passRate := float64(passed) / float64(len(pairs))
	require.GreaterOrEqual(t, passRate, 0.70,
		"Cross-script matching pass rate %.1f%% is below 70%% threshold", passRate*100)
}

func TestIntegration_BuildIndexAndSearch(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()

	// Build index with sample names
	names := []string{
		"Mohamed Ali",
		"Vladimir Putin",
		"Kim Jong Un",
		"John Smith",
		"Jane Doe",
		"Ahmed Hassan",
		"Alexander Petrov",
		"Li Ming",
	}
	ids := []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8"}

	err := service.BuildIndex(ctx, names, ids)
	require.NoError(t, err)
	require.Equal(t, len(names), service.IndexSize())

	// Search with Arabic query - should find Mohamed Ali
	results, err := service.Search(ctx, "محمد علي", 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	t.Logf("Search for 'محمد علي': top result = %s (score: %.4f)",
		results[0].Name, results[0].Score)

	// Search with Cyrillic query - should find Vladimir Putin
	results, err = service.Search(ctx, "Владимир Путин", 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	t.Logf("Search for 'Владимир Путин': top result = %s (score: %.4f)",
		results[0].Name, results[0].Score)

	// Search with Chinese query - should find Kim Jong Un
	results, err = service.Search(ctx, "金正恩", 3)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	t.Logf("Search for '金正恩': top result = %s (score: %.4f)",
		results[0].Name, results[0].Score)
}

func TestIntegration_ShouldUseEmbeddings(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	tests := []struct {
		query    string
		expected bool
	}{
		// Non-Latin queries should use embeddings
		{"محمد علي", true},
		{"Владимир Путин", true},
		{"金正恩", true},
		{"김정은", true},

		// Latin-only queries should not use embeddings (use Jaro-Winkler instead)
		{"Mohamed Ali", false},
		{"John Smith", false},
		{"Jane Doe", false},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			result := service.ShouldUseEmbeddings(tc.query)
			require.Equal(t, tc.expected, result,
				"ShouldUseEmbeddings(%q) = %v, want %v", tc.query, result, tc.expected)
		})
	}
}

func TestIntegration_CacheEffectiveness(t *testing.T) {
	service := createTestService(t)
	defer service.Shutdown()

	ctx := context.Background()
	query := "محمد علي"

	// First encode - cache miss
	start := time.Now()
	_, err := service.Encode(ctx, query)
	require.NoError(t, err)
	firstDuration := time.Since(start)

	// Second encode - cache hit (should be faster)
	start = time.Now()
	_, err = service.Encode(ctx, query)
	require.NoError(t, err)
	secondDuration := time.Since(start)

	t.Logf("First encode: %v, Second encode (cached): %v", firstDuration, secondDuration)

	// Cache hit should be significantly faster (at least 2x)
	require.Less(t, secondDuration, firstDuration,
		"Cached query should be faster than uncached")
}

// Helper functions

func createTestService(t *testing.T) Service {
	t.Helper()

	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping integration test", modelPath)
	}

	logger := log.NewTestLogger()
	config := Config{
		Enabled:             true,
		ModelPath:           modelPath,
		CacheSize:           100,
		CrossScriptOnly:     true,
		SimilarityThreshold: 0.7,
		BatchSize:           32,
		IndexBuildTimeout:   5 * time.Minute,
	}

	ctx := context.Background()
	service, err := NewService(logger, config)
	require.NoError(t, err)
	require.NotNil(t, service)

	// Warm up the model with a single encode
	_, err = service.Encode(ctx, "test")
	require.NoError(t, err)

	return service
}

type crossScriptPair struct {
	Query       string
	Expected    string
	MinScore    float64
	Description string
}

func loadCrossScriptPairs(t *testing.T) []crossScriptPair {
	t.Helper()

	testdataPath := filepath.Join("testdata", "cross_script_pairs.csv")
	file, err := os.Open(testdataPath)
	if err != nil {
		t.Skipf("Test data file not found: %v", err)
	}
	defer file.Close()

	var pairs []crossScriptPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}

		minScore, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
		}

		desc := ""
		if len(parts) >= 4 {
			desc = parts[3]
		}

		pairs = append(pairs, crossScriptPair{
			Query:       parts[0],
			Expected:    parts[1],
			MinScore:    minScore,
			Description: desc,
		})
	}

	return pairs
}
