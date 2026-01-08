//go:build embeddings && integration

package embeddings

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

// AccuracyBenchmark runs comprehensive accuracy measurements
// Run with: go test -tags "embeddings integration" -run TestAccuracy -v
func TestAccuracyBenchmark(t *testing.T) {
	service := createAccuracyTestService(t)
	defer service.Shutdown()

	ctx := context.Background()

	// Build a test index with known names
	testIndex := buildTestIndex(t, service)

	// Run accuracy tests
	results := runAccuracyTests(t, ctx, service, testIndex)

	// Print summary
	printAccuracySummary(t, results)
}

type testIndexEntry struct {
	ID          string
	Name        string
	Variants    []string // Alternative names/transliterations
	Script      string
	Description string
}

type accuracyResults struct {
	TotalQueries    int
	PrecisionAt1    float64
	PrecisionAt5    float64
	PrecisionAt10   float64
	MRR             float64 // Mean Reciprocal Rank
	CrossScriptHits int
	CrossScriptMRR  float64
	LatencyP50      time.Duration
	LatencyP95      time.Duration
}

func buildTestIndex(t *testing.T, service Service) []testIndexEntry {
	t.Helper()

	// Create a test dataset with known correct mappings
	entries := []testIndexEntry{
		// Arabic names
		{
			ID:       "AR001",
			Name:     "Mohamed Ali",
			Variants: []string{"محمد علي", "Muhammad Ali", "Mohammed Ali"},
			Script:   "Arabic",
		},
		{
			ID:       "AR002",
			Name:     "Ahmed Hassan",
			Variants: []string{"أحمد حسن", "Ahmad Hassan"},
			Script:   "Arabic",
		},
		{
			ID:       "AR003",
			Name:     "Abdullah Al-Rashid",
			Variants: []string{"عبد الله الراشد", "Abdallah Al Rashid"},
			Script:   "Arabic",
		},
		{
			ID:       "AR004",
			Name:     "Khalid Mahmoud",
			Variants: []string{"خالد محمود", "Khaled Mahmoud"},
			Script:   "Arabic",
		},

		// Russian names
		{
			ID:       "RU001",
			Name:     "Vladimir Putin",
			Variants: []string{"Владимир Путин", "Vladimir Vladimirovich Putin"},
			Script:   "Cyrillic",
		},
		{
			ID:       "RU002",
			Name:     "Ivan Ivanov",
			Variants: []string{"Иван Иванов", "Ivan Ivanovich Ivanov"},
			Script:   "Cyrillic",
		},
		{
			ID:       "RU003",
			Name:     "Alexander Petrov",
			Variants: []string{"Александр Петров", "Aleksandr Petrov"},
			Script:   "Cyrillic",
		},
		{
			ID:       "RU004",
			Name:     "Dmitry Medvedev",
			Variants: []string{"Дмитрий Медведев"},
			Script:   "Cyrillic",
		},

		// Chinese/Korean names
		{
			ID:       "CN001",
			Name:     "Kim Jong Un",
			Variants: []string{"金正恩", "김정은"},
			Script:   "Han/Hangul",
		},
		{
			ID:       "CN002",
			Name:     "Li Ming",
			Variants: []string{"李明"},
			Script:   "Han",
		},
		{
			ID:       "CN003",
			Name:     "Xi Jinping",
			Variants: []string{"习近平"},
			Script:   "Han",
		},
		{
			ID:       "CN004",
			Name:     "Wang Wei",
			Variants: []string{"王伟"},
			Script:   "Han",
		},

		// Latin-only entries (for negative tests)
		{
			ID:       "EN001",
			Name:     "John Smith",
			Variants: []string{},
			Script:   "Latin",
		},
		{
			ID:       "EN002",
			Name:     "Jane Doe",
			Variants: []string{},
			Script:   "Latin",
		},
		{
			ID:       "EN003",
			Name:     "Michael Johnson",
			Variants: []string{},
			Script:   "Latin",
		},
	}

	// Build the index with primary names
	ctx := context.Background()
	names := make([]string, len(entries))
	ids := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name
		ids[i] = e.ID
	}

	err := service.BuildIndex(ctx, names, ids)
	require.NoError(t, err)

	return entries
}

func runAccuracyTests(t *testing.T, ctx context.Context, service Service, entries []testIndexEntry) accuracyResults {
	t.Helper()

	var results accuracyResults
	var latencies []time.Duration
	var reciprocalRanks []float64
	var crossScriptRRs []float64

	// Test each variant query
	for _, entry := range entries {
		for _, variant := range entry.Variants {
			results.TotalQueries++

			start := time.Now()
			searchResults, err := service.Search(ctx, variant, 10)
			latency := time.Since(start)
			latencies = append(latencies, latency)

			require.NoError(t, err)

			// Find the rank of the correct answer
			rank := findRank(searchResults, entry.ID)

			if rank == 1 {
				results.PrecisionAt1++
			}
			if rank <= 5 {
				results.PrecisionAt5++
			}
			if rank <= 10 {
				results.PrecisionAt10++
			}

			// Calculate reciprocal rank
			rr := 0.0
			if rank > 0 {
				rr = 1.0 / float64(rank)
			}
			reciprocalRanks = append(reciprocalRanks, rr)

			// Track cross-script specific stats
			if entry.Script != "Latin" {
				crossScriptRRs = append(crossScriptRRs, rr)
				if rank == 1 {
					results.CrossScriptHits++
				}
			}

			// Log each result
			topResult := "none"
			topScore := 0.0
			if len(searchResults) > 0 {
				topResult = searchResults[0].Name
				topScore = searchResults[0].Score
			}
			t.Logf("Query: %q -> Expected: %s, Top: %s (%.4f), Rank: %d",
				variant, entry.Name, topResult, topScore, rank)
		}
	}

	// Calculate final metrics
	if results.TotalQueries > 0 {
		results.PrecisionAt1 /= float64(results.TotalQueries)
		results.PrecisionAt5 /= float64(results.TotalQueries)
		results.PrecisionAt10 /= float64(results.TotalQueries)
		results.MRR = mean(reciprocalRanks)
	}

	if len(crossScriptRRs) > 0 {
		results.CrossScriptMRR = mean(crossScriptRRs)
	}

	// Calculate latency percentiles
	results.LatencyP50 = percentile(latencies, 50)
	results.LatencyP95 = percentile(latencies, 95)

	return results
}

func printAccuracySummary(t *testing.T, results accuracyResults) {
	separator := strings.Repeat("=", 60)
	t.Logf("\n%s", separator)
	t.Logf("ACCURACY BENCHMARK RESULTS")
	t.Logf("%s", separator)
	t.Logf("Total queries: %d", results.TotalQueries)
	t.Logf("")
	t.Logf("Precision Metrics:")
	t.Logf("  Precision@1:  %.2f%% (correct answer is top result)", results.PrecisionAt1*100)
	t.Logf("  Precision@5:  %.2f%% (correct answer in top 5)", results.PrecisionAt5*100)
	t.Logf("  Precision@10: %.2f%% (correct answer in top 10)", results.PrecisionAt10*100)
	t.Logf("")
	t.Logf("Ranking Metrics:")
	t.Logf("  MRR (Mean Reciprocal Rank): %.4f", results.MRR)
	t.Logf("  Cross-Script MRR: %.4f", results.CrossScriptMRR)
	t.Logf("  Cross-Script Top-1 Hits: %d", results.CrossScriptHits)
	t.Logf("")
	t.Logf("Latency:")
	t.Logf("  P50: %v", results.LatencyP50)
	t.Logf("  P95: %v", results.LatencyP95)
	t.Logf("%s", separator)

	// Assert minimum quality thresholds
	require.GreaterOrEqual(t, results.PrecisionAt1, 0.60,
		"Precision@1 should be at least 60%%")
	require.GreaterOrEqual(t, results.MRR, 0.70,
		"MRR should be at least 0.70")
	require.GreaterOrEqual(t, results.CrossScriptMRR, 0.65,
		"Cross-script MRR should be at least 0.65")
}

func findRank(results []SearchResult, targetID string) int {
	for i, r := range results {
		if r.ID == targetID {
			return i + 1
		}
	}
	return 0 // Not found
}

func mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func percentile(durations []time.Duration, p int) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	// Sort
	sorted := make([]time.Duration, len(durations))
	copy(sorted, durations)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	idx := (p * len(sorted)) / 100
	if idx >= len(sorted) {
		idx = len(sorted) - 1
	}
	return sorted[idx]
}

func createAccuracyTestService(t *testing.T) Service {
	t.Helper()

	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		t.Skipf("Model not found at %s, skipping accuracy test", modelPath)
	}

	logger := log.NewTestLogger()
	config := Config{
		Enabled:             true,
		ModelPath:           modelPath,
		CacheSize:           1000,
		CrossScriptOnly:     true,
		SimilarityThreshold: 0.5,
		BatchSize:           32,
		IndexBuildTimeout:   10 * time.Minute,
	}

	service, err := NewService(logger, config)
	require.NoError(t, err)
	require.NotNil(t, service)

	return service
}

// BenchmarkEmbeddings runs performance benchmarks
func BenchmarkEncode(b *testing.B) {
	service := createBenchmarkService(b)
	if service == nil {
		return
	}
	defer service.Shutdown()

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.Encode(ctx, "Mohamed Ali")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeCrossScript(b *testing.B) {
	service := createBenchmarkService(b)
	if service == nil {
		return
	}
	defer service.Shutdown()

	ctx := context.Background()
	queries := []string{"محمد علي", "Владимир Путин", "金正恩", "김정은"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.Encode(ctx, queries[i%len(queries)])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodeBatch(b *testing.B) {
	service := createBenchmarkService(b)
	if service == nil {
		return
	}
	defer service.Shutdown()

	ctx := context.Background()
	batch := []string{
		"Mohamed Ali",
		"محمد علي",
		"Vladimir Putin",
		"Владимир Путин",
		"Kim Jong Un",
		"金正恩",
		"John Smith",
		"Jane Doe",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.EncodeBatch(ctx, batch)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	service := createBenchmarkService(b)
	if service == nil {
		return
	}
	defer service.Shutdown()

	ctx := context.Background()

	// Build index
	names := []string{
		"Mohamed Ali", "Ahmed Hassan", "Abdullah Al-Rashid",
		"Vladimir Putin", "Ivan Ivanov", "Alexander Petrov",
		"Kim Jong Un", "Li Ming", "Xi Jinping",
		"John Smith", "Jane Doe", "Michael Johnson",
	}
	ids := make([]string, len(names))
	for i := range names {
		ids[i] = fmt.Sprintf("id%d", i)
	}
	err := service.BuildIndex(ctx, names, ids)
	if err != nil {
		b.Fatal(err)
	}

	queries := []string{"محمد علي", "Владимир Путин", "金正恩"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.Search(ctx, queries[i%len(queries)], 5)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchLargeIndex(b *testing.B) {
	service := createBenchmarkService(b)
	if service == nil {
		return
	}
	defer service.Shutdown()

	ctx := context.Background()

	// Build larger index (simulating OFAC list size)
	baseNames := []string{
		"Mohamed Ali", "Ahmed Hassan", "Abdullah Al-Rashid", "Khalid Mahmoud",
		"Vladimir Putin", "Ivan Ivanov", "Alexander Petrov", "Dmitry Medvedev",
		"Kim Jong Un", "Li Ming", "Xi Jinping", "Wang Wei",
		"John Smith", "Jane Doe", "Michael Johnson", "Robert Williams",
	}

	// Expand to ~1000 entries
	var names []string
	var ids []string
	for i := 0; i < 60; i++ {
		for j, name := range baseNames {
			names = append(names, fmt.Sprintf("%s %d", name, i))
			ids = append(ids, fmt.Sprintf("id_%d_%d", i, j))
		}
	}

	err := service.BuildIndex(ctx, names, ids)
	if err != nil {
		b.Fatal(err)
	}

	b.Logf("Index size: %d", service.IndexSize())

	queries := []string{"محمد علي", "Владимир Путин", "金正恩"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.Search(ctx, queries[i%len(queries)], 10)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func createBenchmarkService(b *testing.B) Service {
	b.Helper()

	modelPath := getModelPath()
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		b.Skipf("Model not found at %s, skipping benchmark", modelPath)
		return nil
	}

	logger := log.NewNopLogger()
	config := Config{
		Enabled:             true,
		ModelPath:           modelPath,
		CacheSize:           1000,
		CrossScriptOnly:     true,
		SimilarityThreshold: 0.5,
		BatchSize:           32,
		IndexBuildTimeout:   10 * time.Minute,
	}

	ctx := context.Background()
	service, err := NewService(logger, config)
	if err != nil {
		b.Fatalf("Failed to create service: %v", err)
		return nil
	}

	// Warm up
	_, _ = service.Encode(ctx, "warmup")

	return service
}
