package tfidf

import (
	"math"
	"sync"
	"testing"
)

func TestNewIndex(t *testing.T) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	if idx == nil {
		t.Fatal("expected non-nil index")
	}

	stats := idx.Stats()
	if stats.TotalDocuments != 0 {
		t.Errorf("expected 0 documents, got %d", stats.TotalDocuments)
	}
	if stats.UniqueTerms != 0 {
		t.Errorf("expected 0 unique terms, got %d", stats.UniqueTerms)
	}
}

func TestBuildAndGetIDF(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Enabled = true
	idx := NewIndex(cfg)

	// Create test documents with varying term frequencies
	// "trading" appears in all 4 documents (common term)
	// "acme" appears in 2 documents
	// "unique" appears in 1 document (rare term)
	documents := [][]string{
		{"acme", "trading", "company"},
		{"global", "trading", "ltd"},
		{"acme", "solutions", "trading"},
		{"unique", "enterprises", "trading"},
	}

	idx.Build(documents)

	stats := idx.Stats()
	if stats.TotalDocuments != 4 {
		t.Errorf("expected 4 documents, got %d", stats.TotalDocuments)
	}

	// Verify document frequencies
	if df := idx.GetDocumentFrequency("trading"); df != 4 {
		t.Errorf("expected df(trading)=4, got %d", df)
	}
	if df := idx.GetDocumentFrequency("acme"); df != 2 {
		t.Errorf("expected df(acme)=2, got %d", df)
	}
	if df := idx.GetDocumentFrequency("unique"); df != 1 {
		t.Errorf("expected df(unique)=1, got %d", df)
	}

	// IDF for common term should be lower than for rare term
	idfTrading := idx.GetIDF("trading")
	idfAcme := idx.GetIDF("acme")
	idfUnique := idx.GetIDF("unique")

	if idfTrading >= idfAcme {
		t.Errorf("expected IDF(trading) < IDF(acme), got %f >= %f", idfTrading, idfAcme)
	}
	if idfAcme >= idfUnique {
		t.Errorf("expected IDF(acme) < IDF(unique), got %f >= %f", idfAcme, idfUnique)
	}
}

func TestGetIDFUnknownTerm(t *testing.T) {
	cfg := DefaultConfig()
	cfg.MaxIDF = 5.0

	idx := NewIndex(cfg)

	documents := [][]string{
		{"known", "term"},
	}
	idx.Build(documents)

	// Unknown term should get MaxIDF
	idfUnknown := idx.GetIDF("unknown_term")
	if idfUnknown != cfg.MaxIDF {
		t.Errorf("expected MaxIDF=%f for unknown term, got %f", cfg.MaxIDF, idfUnknown)
	}
}

func TestIDFBounds(t *testing.T) {
	cfg := Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.5,
		MaxIDF:          3.0,
	}
	idx := NewIndex(cfg)

	// Create documents where one term is extremely common
	documents := make([][]string, 100)
	for i := 0; i < 100; i++ {
		documents[i] = []string{"common"}
	}
	// Add one document with a rare term
	documents = append(documents, []string{"rare"})

	idx.Build(documents)

	idfCommon := idx.GetIDF("common")
	idfRare := idx.GetIDF("rare")

	// Check bounds are applied
	if idfCommon < cfg.MinIDF {
		t.Errorf("IDF(common) should be >= MinIDF, got %f < %f", idfCommon, cfg.MinIDF)
	}
	if idfRare > cfg.MaxIDF {
		t.Errorf("IDF(rare) should be <= MaxIDF, got %f > %f", idfRare, cfg.MaxIDF)
	}
}

func TestGetWeights(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Enabled = true
	idx := NewIndex(cfg)

	documents := [][]string{
		{"trading", "company"},
		{"trading", "ltd"},
		{"trading", "corp"},
		{"unique", "name"},
	}
	idx.Build(documents)

	// Test weights for a query
	weights := idx.GetWeights([]string{"trading", "unique"})

	if len(weights) != 2 {
		t.Fatalf("expected 2 weights, got %d", len(weights))
	}

	// Weight for rare term should be higher
	weightTrading := weights[0]
	weightUnique := weights[1]

	if weightTrading >= weightUnique {
		t.Errorf("expected weight(trading) < weight(unique), got %f >= %f", weightTrading, weightUnique)
	}
}

func TestGetWeightsWithRepeatedTerms(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Enabled = true
	idx := NewIndex(cfg)

	documents := [][]string{
		{"word", "other"},
	}
	idx.Build(documents)

	// Query with repeated term
	weights := idx.GetWeights([]string{"word", "word", "word"})

	if len(weights) != 3 {
		t.Fatalf("expected 3 weights, got %d", len(weights))
	}

	// All weights should be the same (same term, same TF in this query)
	for i := 1; i < len(weights); i++ {
		if weights[i] != weights[0] {
			t.Errorf("expected all weights to be equal, got %f != %f", weights[i], weights[0])
		}
	}

	// TF should be 1 + log(3) for a term appearing 3 times
	expectedTF := 1.0 + math.Log(3.0)
	idf := idx.GetIDF("word")
	expectedWeight := expectedTF * idf

	if math.Abs(weights[0]-expectedWeight) > 0.0001 {
		t.Errorf("expected weight %f, got %f", expectedWeight, weights[0])
	}
}

func TestEmptyDocuments(t *testing.T) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	// Build with no documents
	idx.Build(nil)

	stats := idx.Stats()
	if stats.TotalDocuments != 0 {
		t.Errorf("expected 0 documents, got %d", stats.TotalDocuments)
	}

	// GetIDF should return MaxIDF for any term
	if idf := idx.GetIDF("any"); idf != cfg.MaxIDF {
		t.Errorf("expected MaxIDF for empty index, got %f", idf)
	}

	// GetWeights should return nil for empty input
	if weights := idx.GetWeights(nil); weights != nil {
		t.Errorf("expected nil weights for nil input, got %v", weights)
	}
}

func TestEmptyTermsInDocument(t *testing.T) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	// Documents with empty strings should be handled
	documents := [][]string{
		{"word", "", "other"},
		{"", "word"},
	}
	idx.Build(documents)

	// Empty string should not be counted
	if df := idx.GetDocumentFrequency(""); df != 0 {
		t.Errorf("expected df('')=0, got %d", df)
	}
	if df := idx.GetDocumentFrequency("word"); df != 2 {
		t.Errorf("expected df(word)=2, got %d", df)
	}
}

func TestConcurrentAccess(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Enabled = true
	idx := NewIndex(cfg)

	documents := [][]string{
		{"term1", "term2"},
		{"term2", "term3"},
		{"term3", "term4"},
	}
	idx.Build(documents)

	// Concurrent reads should be safe
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			idx.GetIDF("term1")
			idx.GetIDF("term2")
			idx.GetWeights([]string{"term1", "term2", "unknown"})
			idx.Stats()
		}()
	}
	wg.Wait()
}

func TestRebuild(t *testing.T) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	// First build
	documents1 := [][]string{
		{"old", "term"},
	}
	idx.Build(documents1)

	if df := idx.GetDocumentFrequency("old"); df != 1 {
		t.Errorf("expected df(old)=1, got %d", df)
	}

	// Rebuild with new documents
	documents2 := [][]string{
		{"new", "term"},
		{"new", "data"},
	}
	idx.Build(documents2)

	// Old terms should be gone
	if df := idx.GetDocumentFrequency("old"); df != 0 {
		t.Errorf("expected df(old)=0 after rebuild, got %d", df)
	}
	// New terms should be present
	if df := idx.GetDocumentFrequency("new"); df != 2 {
		t.Errorf("expected df(new)=2, got %d", df)
	}

	stats := idx.Stats()
	if stats.TotalDocuments != 2 {
		t.Errorf("expected 2 documents after rebuild, got %d", stats.TotalDocuments)
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.SmoothingFactor != 1.0 {
		t.Errorf("expected SmoothingFactor=1.0, got %f", cfg.SmoothingFactor)
	}
	if cfg.MinIDF != 0.1 {
		t.Errorf("expected MinIDF=0.1, got %f", cfg.MinIDF)
	}
	if cfg.MaxIDF != 10.0 {
		t.Errorf("expected MaxIDF=5.0, got %f", cfg.MaxIDF)
	}
}

func TestEnabledFlag(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Enabled = false
	idx := NewIndex(cfg)

	if idx.Enabled() {
		t.Error("expected Enabled()=false")
	}

	cfg.Enabled = true
	idx2 := NewIndex(cfg)
	if !idx2.Enabled() {
		t.Error("expected Enabled()=true")
	}
}

// Benchmark tests
func BenchmarkBuild(b *testing.B) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	// Create realistic corpus
	documents := make([][]string, 10000)
	terms := []string{"trading", "company", "ltd", "corp", "limited", "holdings", "group", "international"}
	for i := 0; i < 10000; i++ {
		doc := make([]string, 3)
		for j := 0; j < 3; j++ {
			doc[j] = terms[(i+j)%len(terms)]
		}
		documents[i] = doc
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx.Build(documents)
	}
}

func BenchmarkGetIDF(b *testing.B) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	documents := make([][]string, 1000)
	for i := 0; i < 1000; i++ {
		documents[i] = []string{"term1", "term2", "term3"}
	}
	idx.Build(documents)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx.GetIDF("term1")
	}
}

func BenchmarkGetWeights(b *testing.B) {
	cfg := DefaultConfig()
	idx := NewIndex(cfg)

	documents := make([][]string, 1000)
	for i := 0; i < 1000; i++ {
		documents[i] = []string{"term1", "term2", "term3"}
	}
	idx.Build(documents)

	query := []string{"term1", "term2", "unknown"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx.GetWeights(query)
	}
}
