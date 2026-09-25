// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

type pairResult struct {
	Index          int                 `json:"index"`
	LeftID         string              `json:"leftID"`
	RightID        string              `json:"rightID"`
	LeftName       string              `json:"leftName"`
	RightName      string              `json:"rightName"`
	Schema         string              `json:"schema"`
	Judgement      string              `json:"judgement"`
	Score          float64             `json:"score"`
	Predicted      string              `json:"predicted"`
	CrossScript    bool                `json:"crossScript"`
	LeftDatasets   []string            `json:"leftDatasets,omitempty"`
	RightDatasets  []string            `json:"rightDatasets,omitempty"`
	SearchFound    bool                `json:"searchFound,omitempty"`
	SearchRank     int                 `json:"searchRank,omitempty"`
	SearchSource   string              `json:"searchSource,omitempty"`
	SearchSourceID string              `json:"searchSourceID,omitempty"`
	Pieces         []search.ScorePiece `json:"pieces,omitempty"`
}

type metrics struct {
	N         int     `json:"n"`
	Positive  int     `json:"positive"`
	Negative  int     `json:"negative"`
	Threshold float64 `json:"threshold"`
	Accuracy  float64 `json:"accuracy"`
	Precision float64 `json:"precision"`
	Recall    float64 `json:"recall"`
	F1        float64 `json:"f1"`
	TP        int     `json:"tp"`
	TN        int     `json:"tn"`
	FP        int     `json:"fp"`
	FN        int     `json:"fn"`
	PosMean   float64 `json:"positiveMean"`
	NegMean   float64 `json:"negativeMean"`
	PosMedian float64 `json:"positiveMedian"`
	NegMedian float64 `json:"negativeMedian"`
}

type evalConfig struct {
	name        string
	algorithm   search.StringMatchAlgorithm
	threshold   float64
	nameOnly    bool
	maxAltNames int
	workers     int
	debugErrors int
	tfidf       *tfidf.Index
	embed       *embedCache
	embedMode   string // "", "max", "hybrid", "only"
}

func (c evalConfig) label() string {
	if c.name != "" {
		return c.name
	}
	s := c.algorithm.Name()
	if c.tfidf != nil && c.tfidf.Enabled() {
		s += "+tfidf"
	}
	if c.embedMode != "" {
		s += "+embed-" + c.embedMode
	}
	if c.nameOnly {
		s += "+name-only"
	}
	return s
}

func scorePairwise(pairs []Pair, cfg evalConfig) []pairResult {
	out := make([]pairResult, len(pairs))
	workers := cfg.workers
	if workers < 1 {
		workers = 1
	}

	jobs := make(chan int, workers*4)
	var done atomic.Int64
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			opts := search.SimilarityOpts{Algorithm: cfg.algorithm, TFIDF: cfg.tfidf}
			conv := convertOpts{maxAltNames: cfg.maxAltNames, nameOnly: cfg.nameOnly}
			for i := range jobs {
				p := pairs[i]
				left, right := convertPair(p, conv)
				var score float64
				if cfg.embedMode == "only" && cfg.embed != nil {
					score = cfg.embed.cosine(left.Name, right.Name)
				} else {
					score = similarityFanout(p, conv, opts)
					if cfg.embed != nil && cfg.embedMode != "" {
						cos := cfg.embed.cosine(left.Name, right.Name)
						score = mixEmbed(score, cos, cfg.embedMode, pairCrossScript(p))
					}
				}
				out[i] = makeResult(i, p, score, cfg.threshold)
				n := done.Add(1)
				if n%50000 == 0 {
					fmt.Fprintf(os.Stderr, "%s: %d/%d\n", cfg.label(), n, len(pairs))
				}
			}
		}()
	}
	for i := range pairs {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return out
}

// similarityFanout mirrors a client with unknown query type: one Similarity
// call per Watchman type the two records share (person, business, …), best score wins.
func similarityFanout(p Pair, conv convertOpts, opts search.SimilarityOpts) float64 {
	types := commonTypes(p.Left, p.Right)
	if len(types) == 0 {
		// Known-type mismatch (Person vs Company, Person vs Vessel, …):
		// a client would not fan out. Similarity may recast person/org internally,
		// so do not call it here.
		return 0
	}
	best := 0.0
	for _, typ := range types {
		left := convertEntityAs(p.Left, conv, typ).Normalize()
		right := convertEntityAs(p.Right, conv, typ).Normalize()
		s := search.SimilarityWithOpts(left, right, opts)
		if s > best {
			best = s
		}
	}
	return best
}

func makeResult(i int, p Pair, score float64, threshold float64) pairResult {
	pred := "negative"
	if score >= threshold {
		pred = "positive"
	}
	judgement := "negative"
	if p.isPositive() {
		judgement = "positive"
	}
	return pairResult{
		Index:         i,
		LeftID:        p.Left.ID,
		RightID:       p.Right.ID,
		LeftName:      p.Left.Caption,
		RightName:     p.Right.Caption,
		Schema:        p.schemaKey(),
		Judgement:     judgement,
		Score:         score,
		Predicted:     pred,
		CrossScript:   pairCrossScript(p),
		LeftDatasets:  p.Left.Datasets,
		RightDatasets: p.Right.Datasets,
	}
}

func applyThreshold(results []pairResult, threshold float64) {
	for i := range results {
		if results[i].Score >= threshold {
			results[i].Predicted = "positive"
		} else {
			results[i].Predicted = "negative"
		}
	}
}

func computeMetrics(results []pairResult, threshold float64) metrics {
	var m metrics
	m.N = len(results)
	m.Threshold = threshold
	var pos, neg []float64
	for _, r := range results {
		truth := r.Judgement == "positive"
		pred := r.Predicted == "positive"
		if truth {
			m.Positive++
			pos = append(pos, r.Score)
		} else {
			m.Negative++
			neg = append(neg, r.Score)
		}
		switch {
		case truth && pred:
			m.TP++
		case !truth && !pred:
			m.TN++
		case !truth && pred:
			m.FP++
		default:
			m.FN++
		}
	}
	if m.N > 0 {
		m.Accuracy = float64(m.TP+m.TN) / float64(m.N)
	}
	if m.TP+m.FP > 0 {
		m.Precision = float64(m.TP) / float64(m.TP+m.FP)
	}
	if m.TP+m.FN > 0 {
		m.Recall = float64(m.TP) / float64(m.TP+m.FN)
	}
	if m.Precision+m.Recall > 0 {
		m.F1 = 2 * m.Precision * m.Recall / (m.Precision + m.Recall)
	}
	m.PosMean, m.PosMedian = meanMedian(pos)
	m.NegMean, m.NegMedian = meanMedian(neg)
	return m
}

func meanMedian(values []float64) (mean, median float64) {
	if len(values) == 0 {
		return 0, 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	mean = sum / float64(len(values))
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		median = (sorted[mid-1] + sorted[mid]) / 2
	} else {
		median = sorted[mid]
	}
	return mean, median
}

func sweepThresholds(results []pairResult) (best metrics, all []metrics) {
	best.F1 = -1
	for t := 50; t <= 99; t++ {
		th := float64(t) / 100
		cloned := append([]pairResult(nil), results...)
		applyThreshold(cloned, th)
		m := computeMetrics(cloned, th)
		all = append(all, m)
		if m.F1 > best.F1 || (m.F1 == best.F1 && m.Threshold > best.Threshold) {
			best = m
		}
	}
	return best, all
}

func histogram(results []pairResult, positive bool) []int {
	buckets := make([]int, 20)
	for _, r := range results {
		if (r.Judgement == "positive") != positive {
			continue
		}
		if r.Score >= 1 {
			buckets[len(buckets)-1]++
			continue
		}
		idx := int(math.Floor(r.Score * float64(len(buckets))))
		if idx < 0 || idx >= len(buckets) {
			continue
		}
		buckets[idx]++
	}
	return buckets
}

func schemaBreakdown(results []pairResult, threshold float64) []metricsRow {
	groups := make(map[string][]pairResult)
	for _, r := range results {
		groups[r.Schema] = append(groups[r.Schema], r)
	}
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var rows []metricsRow
	for _, k := range keys {
		m := computeMetrics(groups[k], threshold)
		rows = append(rows, metricsRow{Label: k, Metrics: m})
	}
	return rows
}

type metricsRow struct {
	Label   string
	Metrics metrics
}

func errorExamples(results []pairResult, wantFP bool, n int) []pairResult {
	var picked []pairResult
	for _, r := range results {
		if wantFP && r.Judgement == "negative" && r.Predicted == "positive" {
			picked = append(picked, r)
		}
		if !wantFP && r.Judgement == "positive" && r.Predicted == "negative" {
			picked = append(picked, r)
		}
	}
	sort.Slice(picked, func(i, j int) bool {
		if wantFP {
			return picked[i].Score > picked[j].Score
		}
		return picked[i].Score < picked[j].Score
	})
	if n > 0 && len(picked) > n {
		picked = picked[:n]
	}
	return picked
}

func subset(results []pairResult, keep func(pairResult) bool) []pairResult {
	var out []pairResult
	for _, r := range results {
		if keep(r) {
			out = append(out, r)
		}
	}
	return out
}

func subjectResults(results []pairResult) []pairResult {
	return subset(results, func(r pairResult) bool {
		left, right, _ := strings.Cut(r.Schema, "/")
		return !isAutoMerge(left) && !isAutoMerge(right)
	})
}

func filterPairs(pairs []Pair, subjectsOnly bool, schemas map[string]struct{}, offset, limit int) []Pair {
	var out []Pair
	for _, p := range pairs {
		if subjectsOnly && (isAutoMerge(p.Left.Schema) || isAutoMerge(p.Right.Schema)) {
			continue
		}
		if len(schemas) > 0 {
			if _, ok := schemas[strings.ToLower(p.Left.Schema)]; !ok {
				continue
			}
			if _, ok := schemas[strings.ToLower(p.Right.Schema)]; !ok {
				continue
			}
		}
		out = append(out, p)
	}
	if offset > 0 {
		if offset >= len(out) {
			return nil
		}
		out = out[offset:]
	}
	if limit > 0 && len(out) > limit {
		out = out[:limit]
	}
	return out
}
