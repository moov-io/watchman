package tfidf_test

import (
	"sort"
	"testing"

	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/moov-io/watchman/internal/tfidf"
)

// TestWeightedScoringDistinguishesRareTerms verifies that TF-IDF weighted scoring
// helps distinguish rare names from common business terms in search results.
func TestWeightedScoringDistinguishesRareTerms(t *testing.T) {
	// Simulate OFAC-like corpus with many "Trading Company" entries
	// but only one "ZEUS" - a rare distinctive name
	documents := [][]string{
		// Many generic trading companies (common pattern in OFAC)
		{"alpha", "trading", "company", "ltd"},
		{"beta", "trading", "company", "ltd"},
		{"gamma", "trading", "company", "ltd"},
		{"delta", "trading", "company", "ltd"},
		{"epsilon", "trading", "company", "ltd"},
		{"omega", "trading", "company", "ltd"},
		{"sigma", "trading", "company", "ltd"},
		{"theta", "trading", "company", "ltd"},
		{"kappa", "trading", "company", "ltd"},
		{"lambda", "trading", "company", "ltd"},
		{"global", "trading", "company", "inc"},
		{"world", "trading", "company", "inc"},
		{"international", "trading", "company", "inc"},
		{"united", "trading", "company", "inc"},
		{"national", "trading", "company", "inc"},
		// One unique entry with distinctive name
		{"zeus", "trading", "company", "ltd"},
		// Some other patterns
		{"acme", "holdings", "group"},
		{"apex", "solutions", "inc"},
		{"prime", "industries", "corp"},
		{"mega", "enterprises", "llc"},
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	// Log IDF values for key terms
	t.Logf("Corpus size: %d documents", len(documents))
	t.Logf("\nIDF values:")
	t.Logf("  'trading' (appears in 16/20): %.4f", idx.GetIDF("trading"))
	t.Logf("  'company' (appears in 16/20): %.4f", idx.GetIDF("company"))
	t.Logf("  'ltd'     (appears in 11/20): %.4f", idx.GetIDF("ltd"))
	t.Logf("  'zeus'    (appears in 1/20):  %.4f", idx.GetIDF("zeus"))
	t.Logf("  'acme'    (appears in 1/20):  %.4f", idx.GetIDF("acme"))

	// THE KEY TEST: Search for "zeus trading"
	// Without TF-IDF: many "X trading company" might score similarly
	// With TF-IDF: "zeus trading company" should clearly win because "zeus" is rare
	query := []string{"zeus", "trading"}

	type result struct {
		name           string
		tokens         []string
		scoreNoTFIDF   float64
		scoreWithTFIDF float64
	}

	var results []result
	for _, doc := range documents {
		scoreNoTFIDF := stringscore.BestPairCombinationJaroWinkler(query, doc)

		queryWeights := idx.GetWeights(query)
		docWeights := idx.GetWeights(doc)
		scoreWithTFIDF := stringscore.BestPairCombinationJaroWinklerWeighted(query, doc, queryWeights, docWeights)

		results = append(results, result{
			name:           joinTokens(doc),
			tokens:         doc,
			scoreNoTFIDF:   scoreNoTFIDF,
			scoreWithTFIDF: scoreWithTFIDF,
		})
	}

	// Sort by score WITHOUT TF-IDF
	sort.Slice(results, func(i, j int) bool {
		return results[i].scoreNoTFIDF > results[j].scoreNoTFIDF
	})

	t.Logf("\n=== Search: 'zeus trading' ===")
	t.Logf("\nTop 5 WITHOUT TF-IDF:")
	for i := 0; i < 5 && i < len(results); i++ {
		t.Logf("  %d. %.4f - %s", i+1, results[i].scoreNoTFIDF, results[i].name)
	}

	// Find rank of "zeus trading company" without TF-IDF
	rankNoTFIDF := -1
	for i, r := range results {
		if r.tokens[0] == "zeus" {
			rankNoTFIDF = i + 1
			break
		}
	}

	// Sort by score WITH TF-IDF
	sort.Slice(results, func(i, j int) bool {
		return results[i].scoreWithTFIDF > results[j].scoreWithTFIDF
	})

	t.Logf("\nTop 5 WITH TF-IDF:")
	for i := 0; i < 5 && i < len(results); i++ {
		t.Logf("  %d. %.4f - %s", i+1, results[i].scoreWithTFIDF, results[i].name)
	}

	// Find rank of "zeus trading company" with TF-IDF
	rankWithTFIDF := -1
	for i, r := range results {
		if r.tokens[0] == "zeus" {
			rankWithTFIDF = i + 1
			break
		}
	}

	t.Logf("\n=== RESULT ===")
	t.Logf("'zeus trading company ltd' rank:")
	t.Logf("  Without TF-IDF: #%d", rankNoTFIDF)
	t.Logf("  With TF-IDF:    #%d", rankWithTFIDF)

	// The key assertion: with TF-IDF, "zeus" should rank #1
	if rankWithTFIDF != 1 {
		t.Errorf("Expected 'zeus trading company' to rank #1 with TF-IDF, got #%d", rankWithTFIDF)
	}

	// TF-IDF should improve the rank (lower is better)
	if rankWithTFIDF > rankNoTFIDF {
		t.Errorf("TF-IDF should improve rank: was #%d, now #%d", rankNoTFIDF, rankWithTFIDF)
	}
}

func joinTokens(tokens []string) string {
	result := ""
	for i, t := range tokens {
		if i > 0 {
			result += " "
		}
		result += t
	}
	return result
}

// TestTFIDFReducesFalsePositivesForCommonTerms shows that searching
// for common terms alone should produce lower confidence scores.
func TestTFIDFReducesFalsePositivesForCommonTerms(t *testing.T) {
	documents := [][]string{
		{"mohammed", "ali", "trading"},
		{"ahmed", "trading", "company"},
		{"global", "trading", "company"},
		{"world", "trading", "company"},
		{"best", "trading", "company"},
	}

	cfg := tfidf.Config{
		Enabled:         true,
		SmoothingFactor: 1.0,
		MinIDF:          0.1,
		MaxIDF:          5.0,
	}
	idx := tfidf.NewIndex(cfg)
	idx.Build(documents)

	// Search for just "trading" - a very common term
	query := []string{"trading"}

	queryWeights := idx.GetWeights(query)

	t.Logf("Search: 'trading' (common term)")
	t.Logf("IDF('trading') = %.4f (low = common)", idx.GetIDF("trading"))

	// All documents contain "trading", but with TF-IDF the scores should be lower
	// because "trading" has low IDF
	for _, doc := range documents {
		scoreNoTFIDF := stringscore.BestPairCombinationJaroWinkler(query, doc)
		docWeights := idx.GetWeights(doc)
		scoreWithTFIDF := stringscore.BestPairCombinationJaroWinklerWeighted(query, doc, queryWeights, docWeights)

		t.Logf("  %s: no-tfidf=%.4f, tfidf=%.4f",
			joinTokens(doc), scoreNoTFIDF, scoreWithTFIDF)
	}

	// Now search for "mohammed" - a rare term
	query2 := []string{"mohammed"}
	query2Weights := idx.GetWeights(query2)

	t.Logf("\nSearch: 'mohammed' (rare term)")
	t.Logf("IDF('mohammed') = %.4f (high = rare)", idx.GetIDF("mohammed"))

	for _, doc := range documents {
		scoreNoTFIDF := stringscore.BestPairCombinationJaroWinkler(query2, doc)
		docWeights := idx.GetWeights(doc)
		scoreWithTFIDF := stringscore.BestPairCombinationJaroWinklerWeighted(query2, doc, query2Weights, docWeights)

		t.Logf("  %s: no-tfidf=%.4f, tfidf=%.4f",
			joinTokens(doc), scoreNoTFIDF, scoreWithTFIDF)
	}
}
