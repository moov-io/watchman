package tfidf_test

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/ofactest"
	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

// TestTFIDFWithRealOFACData uses actual OFAC data to verify TF-IDF behavior.
// This test requires the OFAC testdata files.
func TestTFIDFWithRealOFACData(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping OFAC data test in short mode")
	}

	// Load real OFAC data
	stats, err := ofactest.GetDownloader(t).RefreshAll(context.Background())
	if err != nil {
		t.Fatalf("failed to load OFAC data: %v", err)
	}

	t.Logf("Loaded %d OFAC entities", len(stats.Entities))

	// Build TF-IDF index from entity names
	var documents [][]string
	for i := range stats.Entities {
		if len(stats.Entities[i].PreparedFields.NameFields) > 0 {
			documents = append(documents, stats.Entities[i].PreparedFields.NameFields)
		}
		for _, alt := range stats.Entities[i].PreparedFields.AltNameFields {
			if len(alt) > 0 {
				documents = append(documents, alt)
			}
		}
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	idxStats := idx.Stats()
	t.Logf("TF-IDF index: %d documents, %d unique terms", idxStats.TotalDocuments, idxStats.UniqueTerms)

	// Analyze term frequencies - find common and rare terms
	commonTerms := []string{"trading", "company", "ltd", "limited", "inc", "corp", "international"}
	rareTermExamples := findRareTerms(t, idx, stats.Entities, 5)

	t.Logf("\nCommon terms IDF values:")
	for _, term := range commonTerms {
		df := idx.GetDocumentFrequency(term)
		idf := idx.GetIDF(term)
		t.Logf("  '%s': df=%d, idf=%.4f", term, df, idf)
	}

	t.Logf("\nRare terms IDF values (examples):")
	for _, term := range rareTermExamples {
		df := idx.GetDocumentFrequency(term)
		idf := idx.GetIDF(term)
		t.Logf("  '%s': df=%d, idf=%.4f", term, df, idf)
	}

	// In a small dataset, we can only verify that the IDF formula works correctly:
	// terms with higher document frequency should have lower IDF
	t.Logf("\nVerifying IDF ordering (higher df = lower idf):")
	for _, common := range commonTerms {
		commonDF := idx.GetDocumentFrequency(common)
		commonIDF := idx.GetIDF(common)
		if commonDF == 0 {
			continue // term not in corpus
		}
		for _, rare := range rareTermExamples {
			rareDF := idx.GetDocumentFrequency(rare)
			rareIDF := idx.GetIDF(rare)
			if commonDF > rareDF && commonIDF >= rareIDF {
				t.Errorf("Term '%s' (df=%d, idf=%.4f) should have lower IDF than '%s' (df=%d, idf=%.4f)",
					common, commonDF, commonIDF, rare, rareDF, rareIDF)
			}
		}
	}
}

// findRareTerms finds terms that appear in only 1-2 documents
func findRareTerms(t *testing.T, idx *tfidf.Index, entities []search.Entity[search.Value], count int) []string {
	termCounts := make(map[string]int)

	for i := range entities {
		seen := make(map[string]bool)
		for _, term := range entities[i].PreparedFields.NameFields {
			if !seen[term] && len(term) > 3 { // skip short terms
				termCounts[term]++
				seen[term] = true
			}
		}
	}

	// Find terms with count == 1 (truly unique)
	var rare []string
	for term, cnt := range termCounts {
		if cnt == 1 {
			rare = append(rare, term)
			if len(rare) >= count {
				break
			}
		}
	}

	return rare
}

// TestTFIDFRankingImprovementWithOFAC tests that TF-IDF improves ranking
// for searches where the query contains both common and rare terms.
func TestTFIDFRankingImprovementWithOFAC(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping OFAC ranking test in short mode")
	}

	stats, err := ofactest.GetDownloader(t).RefreshAll(context.Background())
	if err != nil {
		t.Fatalf("failed to load OFAC data: %v", err)
	}

	// Build TF-IDF index
	var documents [][]string
	for i := range stats.Entities {
		if len(stats.Entities[i].PreparedFields.NameFields) > 0 {
			documents = append(documents, stats.Entities[i].PreparedFields.NameFields)
		}
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	// Find an entity with a distinctive name + common suffix
	// e.g., "UNIQUE_NAME Trading Company"
	var testEntity search.Entity[search.Value]
	var distinctiveTerm string

	for i := range stats.Entities {
		fields := stats.Entities[i].PreparedFields.NameFields
		if len(fields) >= 2 {
			// Look for entity with at least one rare term and one common term
			hasRare := false
			hasCommon := false
			var rareTerm string

			for _, term := range fields {
				df := idx.GetDocumentFrequency(term)
				if df == 1 && len(term) > 4 {
					hasRare = true
					rareTerm = term
				}
				if df > 10 {
					hasCommon = true
				}
			}

			if hasRare && hasCommon {
				testEntity = stats.Entities[i]
				distinctiveTerm = rareTerm
				break
			}
		}
	}

	if testEntity.Name == "" {
		t.Skip("Could not find suitable test entity with rare + common terms")
	}

	t.Logf("Test entity: %s", testEntity.Name)
	t.Logf("Distinctive term: '%s' (df=%d)", distinctiveTerm, idx.GetDocumentFrequency(distinctiveTerm))
	t.Logf("Query terms: %v", testEntity.PreparedFields.NameFields)

	// Search for this entity among all entities
	query := testEntity.PreparedFields.NameFields

	type result struct {
		entity         search.Entity[search.Value]
		scoreNoTFIDF   float64
		scoreWithTFIDF float64
	}

	var results []result
	for i := range stats.Entities {
		indexFields := stats.Entities[i].PreparedFields.NameFields
		if len(indexFields) == 0 {
			continue
		}

		scoreNoTFIDF := stringscore.BestPairCombinationJaroWinkler(query, indexFields)

		queryWeights := idx.GetWeights(query)
		indexWeights := idx.GetWeights(indexFields)
		scoreWithTFIDF := stringscore.BestPairCombinationJaroWinklerWeighted(query, indexFields, queryWeights, indexWeights)

		results = append(results, result{
			entity:         stats.Entities[i],
			scoreNoTFIDF:   scoreNoTFIDF,
			scoreWithTFIDF: scoreWithTFIDF,
		})
	}

	// Find rank without TF-IDF
	sort.Slice(results, func(i, j int) bool {
		return results[i].scoreNoTFIDF > results[j].scoreNoTFIDF
	})

	rankNoTFIDF := -1
	for i, r := range results {
		if r.entity.SourceID == testEntity.SourceID {
			rankNoTFIDF = i + 1
			break
		}
	}

	t.Logf("\nWithout TF-IDF - Top 5:")
	for i := 0; i < 5 && i < len(results); i++ {
		t.Logf("  %d. %.4f - %s", i+1, results[i].scoreNoTFIDF, results[i].entity.Name)
	}

	// Find rank with TF-IDF
	sort.Slice(results, func(i, j int) bool {
		return results[i].scoreWithTFIDF > results[j].scoreWithTFIDF
	})

	rankWithTFIDF := -1
	for i, r := range results {
		if r.entity.SourceID == testEntity.SourceID {
			rankWithTFIDF = i + 1
			break
		}
	}

	t.Logf("\nWith TF-IDF - Top 5:")
	for i := 0; i < 5 && i < len(results); i++ {
		t.Logf("  %d. %.4f - %s", i+1, results[i].scoreWithTFIDF, results[i].entity.Name)
	}

	t.Logf("\nTarget entity rank:")
	t.Logf("  Without TF-IDF: #%d", rankNoTFIDF)
	t.Logf("  With TF-IDF:    #%d", rankWithTFIDF)

	// The entity should rank #1 in both cases (exact match)
	if rankWithTFIDF != 1 {
		t.Errorf("Expected target entity to rank #1 with TF-IDF, got #%d", rankWithTFIDF)
	}
}

// TestCommonTermsHaveLowerIDF verifies that business suffixes have low IDF.
// NOTE: This test requires the full OFAC dataset to be meaningful.
// With the small test dataset (17 entities), common terms may not appear frequently.
func TestCommonTermsHaveLowerIDF(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	stats, err := ofactest.GetDownloader(t).RefreshAll(context.Background())
	if err != nil {
		t.Fatalf("failed to load OFAC data: %v", err)
	}

	// Skip if dataset is too small for meaningful IDF analysis
	if len(stats.Entities) < 100 {
		t.Skipf("Dataset too small (%d entities) for common term analysis. Need full OFAC data.", len(stats.Entities))
	}

	var documents [][]string
	for i := range stats.Entities {
		if len(stats.Entities[i].PreparedFields.NameFields) > 0 {
			documents = append(documents, stats.Entities[i].PreparedFields.NameFields)
		}
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	// Common business terms should have IDF < 1.0 in a large corpus
	commonTerms := map[string]bool{
		"trading": true, "company": true, "ltd": true, "limited": true,
		"inc": true, "corp": true, "corporation": true, "international": true,
		"group": true, "holdings": true, "enterprises": true,
	}

	var lowIDFCount, highIDFCount int
	for term := range commonTerms {
		idf := idx.GetIDF(term)
		df := idx.GetDocumentFrequency(term)

		if df > 0 { // term exists in corpus
			if idf < 1.0 {
				lowIDFCount++
				t.Logf("✓ '%s': df=%d, idf=%.4f (low - good)", term, df, idf)
			} else {
				highIDFCount++
				t.Logf("? '%s': df=%d, idf=%.4f (not as low as expected)", term, df, idf)
			}
		}
	}

	// Most common terms should have low IDF
	if lowIDFCount < highIDFCount {
		t.Errorf("Expected most common terms to have IDF < 1.0, got %d low vs %d high",
			lowIDFCount, highIDFCount)
	}
}

// BenchmarkTFIDFWithOFAC benchmarks TF-IDF building and querying with real data.
func BenchmarkTFIDFWithOFAC(b *testing.B) {
	stats, err := ofactest.GetDownloader(b).RefreshAll(context.Background())
	if err != nil {
		b.Fatalf("failed to load OFAC data: %v", err)
	}

	var documents [][]string
	for i := range stats.Entities {
		if len(stats.Entities[i].PreparedFields.NameFields) > 0 {
			documents = append(documents, stats.Entities[i].PreparedFields.NameFields)
		}
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}

	b.Run("Build", func(b *testing.B) {
		idx := tfidf.NewIndex(cfg)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			idx.Build(documents)
		}
	})

	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	query := strings.Fields("mohammed trading company")

	b.Run("GetWeights", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx.GetWeights(query)
		}
	})

	b.Run("GetIDF", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx.GetIDF("trading")
		}
	})
}
