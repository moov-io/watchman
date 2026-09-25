// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/moov-io/watchman/internal/tfidf"
	"github.com/moov-io/watchman/pkg/search"
)

type configResult struct {
	Name           string  `json:"name"`
	Elapsed        string  `json:"elapsed"`
	AtThreshold    metrics `json:"atThreshold"`
	BestF1         metrics `json:"bestF1"`
	Subjects       metrics `json:"subjects"`
	SubjectsBestF1 metrics `json:"subjectsBestF1"`
	CrossScript    metrics `json:"crossScript"`
}

func defaultCompareSpecs() []string {
	return []string{
		"jaro-winkler",
		"jaro-winkler+tfidf",
		"soundex",
		"soundex+tfidf",
		"double-metaphone",
		"double-metaphone+tfidf",
		"soft-bidist",
		"nsim",
		"editex",
		"beider-morse",
		"embed-hybrid",
		"embed-max",
		"embed-only",
	}
}

func parseCompareSpec(spec string, base evalConfig, idx *tfidf.Index, emb *embedCache) (evalConfig, error) {
	spec = strings.ToLower(strings.TrimSpace(spec))
	cfg := base
	cfg.name = spec
	cfg.debugErrors = 0

	switch spec {
	case "embed-hybrid":
		cfg.embed = emb
		cfg.embedMode = "hybrid"
		cfg.algorithm = search.AlgorithmJaroWinkler
		return cfg, nil
	case "embed-max":
		cfg.embed = emb
		cfg.embedMode = "max"
		cfg.algorithm = search.AlgorithmJaroWinkler
		return cfg, nil
	case "embed-only":
		cfg.embed = emb
		cfg.embedMode = "only"
		return cfg, nil
	}

	wantTFIDF := false
	algoPart := spec
	if strings.HasSuffix(spec, "+tfidf") {
		wantTFIDF = true
		algoPart = strings.TrimSuffix(spec, "+tfidf")
	}
	if strings.HasSuffix(algoPart, "+name-only") {
		cfg.nameOnly = true
		algoPart = strings.TrimSuffix(algoPart, "+name-only")
	}
	algo, err := search.ParseStringMatchAlgorithm(algoPart)
	if err != nil {
		return cfg, fmt.Errorf("compare spec %q: %w", spec, err)
	}
	cfg.algorithm = algo
	if wantTFIDF {
		if idx == nil {
			return cfg, fmt.Errorf("compare spec %q needs a TF-IDF index", spec)
		}
		cfg.tfidf = idx
	}
	return cfg, nil
}

func runCompare(pairs []Pair, specs []string, base evalConfig, idx *tfidf.Index, emb *embedCache) ([]configResult, error) {
	var out []configResult
	for _, spec := range specs {
		cfg, err := parseCompareSpec(spec, base, idx, emb)
		if err != nil {
			return out, err
		}
		if strings.HasPrefix(spec, "embed-") && emb == nil {
			fmt.Fprintf(os.Stderr, "skip %s: embeddings not configured\n", spec)
			continue
		}
		fmt.Fprintf(os.Stderr, "=== %s ===\n", cfg.label())
		start := time.Now()
		results := scorePairwise(pairs, cfg)
		elapsed := time.Since(start)
		applyThreshold(results, cfg.threshold)
		best, _ := sweepThresholds(results)
		subjects := subjectResults(results)
		cross := subset(results, func(r pairResult) bool { return r.CrossScript })
		sb, _ := sweepThresholds(subjects)
		row := configResult{
			Name:           cfg.label(),
			Elapsed:        elapsed.Round(time.Millisecond).String(),
			AtThreshold:    computeMetrics(results, cfg.threshold),
			BestF1:         best,
			Subjects:       computeMetrics(subjects, cfg.threshold),
			SubjectsBestF1: sb,
			CrossScript:    computeMetrics(cross, cfg.threshold),
		}
		out = append(out, row)
		fmt.Fprintf(os.Stderr, "%s done in %s  f1@%.2f=%.4f  best=%.4f@%.2f  subjects=%.4f\n",
			cfg.label(), row.Elapsed, cfg.threshold, row.AtThreshold.F1, row.BestF1.F1, row.BestF1.Threshold, row.Subjects.F1)
	}
	return out, nil
}

func printCompareTable(w io.Writer, rows []configResult, threshold float64) {
	fmt.Fprintf(w, "Watchman configuration comparison (threshold=%.2f)\n\n", threshold)
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "config\telapsed\tn\tacc\tf1\tprec\trec\tbestF1\tbestThr\tsubjF1\tsubjBest\tcrossF1\tcrossRec")
	for _, r := range rows {
		a, b, s, sb, c := r.AtThreshold, r.BestF1, r.Subjects, r.SubjectsBestF1, r.CrossScript
		fmt.Fprintf(tw, "%s\t%s\t%d\t%.3f\t%.3f\t%.3f\t%.3f\t%.3f\t%.2f\t%.3f\t%.3f\t%.3f\t%.3f\n",
			r.Name, r.Elapsed, a.N, a.Accuracy, a.F1, a.Precision, a.Recall,
			b.F1, b.Threshold, s.F1, sb.F1, c.F1, c.Recall)
	}
	tw.Flush()
	fmt.Fprintln(w)
	fmt.Fprintln(w, "All-pairs at threshold (confusion)")
	tw = tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "config\tTP\tTN\tFP\tFN\tposMean\tnegMean")
	for _, r := range rows {
		a := r.AtThreshold
		fmt.Fprintf(tw, "%s\t%d\t%d\t%d\t%d\t%.3f\t%.3f\n",
			r.Name, a.TP, a.TN, a.FP, a.FN, a.PosMean, a.NegMean)
	}
	tw.Flush()
}
