package tfidf

import (
	"math"
	"sync"
)

// Index stores precomputed IDF values for efficient term weighting.
// It is safe for concurrent read access after Build() completes.
type Index struct {
	mu     sync.RWMutex
	config Config

	termDocFreq    map[string]int     // term -> number of documents containing term
	idfValues      map[string]float64 // precomputed IDF values
	totalDocuments int
}

// Stats contains statistics about the TF-IDF index.
type Stats struct {
	TotalDocuments int `json:"totalDocuments"`
	UniqueTerms    int `json:"uniqueTerms"`
	Enabled        bool `json:"enabled"`
}

// NewIndex creates a new TF-IDF index with the given configuration.
func NewIndex(config Config) *Index {
	return &Index{
		config:      config,
		termDocFreq: make(map[string]int),
		idfValues:   make(map[string]float64),
	}
}

// Build computes IDF values from a slice of tokenized documents.
// Each document is represented as a slice of terms (already normalized).
//
// The IDF formula used is: log((N + k) / (df + k))
// where N is total documents, df is document frequency, and k is smoothing factor.
func (idx *Index) Build(documents [][]string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	// Reset
	idx.termDocFreq = make(map[string]int)
	idx.totalDocuments = len(documents)

	if idx.totalDocuments == 0 {
		idx.idfValues = make(map[string]float64)
		return
	}

	// Count document frequency for each term
	// A term is counted once per document, even if it appears multiple times
	for _, doc := range documents {
		seen := make(map[string]bool)
		for _, term := range doc {
			if term == "" {
				continue
			}
			if !seen[term] {
				idx.termDocFreq[term]++
				seen[term] = true
			}
		}
	}

	// Precompute IDF values
	idx.idfValues = make(map[string]float64, len(idx.termDocFreq))
	for term, df := range idx.termDocFreq {
		idx.idfValues[term] = idx.computeIDF(df)
	}
}

// computeIDF calculates the IDF value for a given document frequency.
func (idx *Index) computeIDF(docFreq int) float64 {
	k := idx.config.SmoothingFactor
	N := float64(idx.totalDocuments)
	df := float64(docFreq)

	// Smoothed IDF formula: log((N + k) / (df + k))
	idf := math.Log((N + k) / (df + k))

	// Apply bounds
	idf = math.Max(idx.config.MinIDF, idf)
	idf = math.Min(idx.config.MaxIDF, idf)

	return idf
}

// GetIDF returns the IDF value for a term.
// Unknown terms (not in the corpus) receive MaxIDF value,
// as they are assumed to be rare/unique.
func (idx *Index) GetIDF(term string) float64 {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if idf, ok := idx.idfValues[term]; ok {
		return idf
	}

	// Unknown term gets maximum IDF (it's very rare/unique)
	return idx.config.MaxIDF
}

// GetWeights returns TF-IDF weights for a slice of terms.
// The weight for each term is TF(t,d) * IDF(t), where:
//   - TF(t,d) = 1 + log(frequency of t in this document)
//   - IDF(t) = precomputed inverse document frequency
func (idx *Index) GetWeights(terms []string) []float64 {
	if len(terms) == 0 {
		return nil
	}

	weights := make([]float64, len(terms))

	// Calculate term frequencies within this document
	tf := make(map[string]int)
	for _, term := range terms {
		tf[term]++
	}

	idx.mu.RLock()
	defer idx.mu.RUnlock()

	for i, term := range terms {
		// TF = 1 + log(frequency) for terms that appear
		// Using log-normalized TF to dampen the effect of term frequency
		termFreq := 1.0 + math.Log(float64(tf[term]))

		// Get IDF value
		idf := idx.config.MaxIDF // default for unknown terms
		if val, ok := idx.idfValues[term]; ok {
			idf = val
		}

		weights[i] = termFreq * idf
	}

	return weights
}

// GetDocumentFrequency returns how many documents contain the given term.
// Returns 0 if the term is not in the index.
func (idx *Index) GetDocumentFrequency(term string) int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	return idx.termDocFreq[term]
}

// Stats returns statistics about the index.
func (idx *Index) Stats() Stats {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	return Stats{
		TotalDocuments: idx.totalDocuments,
		UniqueTerms:    len(idx.termDocFreq),
		Enabled:        idx.config.Enabled,
	}
}

// Enabled returns whether TF-IDF weighting is enabled.
func (idx *Index) Enabled() bool {
	return idx.config.Enabled
}

// Config returns the current configuration.
func (idx *Index) Config() Config {
	return idx.config
}
