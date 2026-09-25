// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/moov-io/watchman/pkg/search"
)

func printReport(w io.Writer, results []pairResult, cfg evalConfig, mode string, sweep bool) {
	applyThreshold(results, cfg.threshold)
	m := computeMetrics(results, cfg.threshold)

	fmt.Fprintf(w, "OpenSanctions Pairs × Watchman (%s / %s)\n", mode, cfg.label())
	fmt.Fprintf(w, "pairs=%d  positive=%d  negative=%d  threshold=%.2f\n\n",
		m.N, m.Positive, m.Negative, m.Threshold)

	fmt.Fprintln(w, "Classification at threshold")
	printMetrics(w, m)

	fmt.Fprintln(w, "Score distribution (mean / median)")
	fmt.Fprintf(w, "  labeled positive: %.3f / %.3f\n", m.PosMean, m.PosMedian)
	fmt.Fprintf(w, "  labeled negative: %.3f / %.3f\n\n", m.NegMean, m.NegMedian)

	fmt.Fprintln(w, "Histogram (score buckets of 0.05)")
	printHistogram(w, "pos", histogram(results, true))
	printHistogram(w, "neg", histogram(results, false))
	fmt.Fprintln(w)

	if n := len(subset(results, func(r pairResult) bool { return r.CrossScript })); n > 0 {
		cross := subset(results, func(r pairResult) bool { return r.CrossScript })
		cm := computeMetrics(cross, cfg.threshold)
		fmt.Fprintf(w, "Cross-script pairs (%d)\n", n)
		printMetrics(w, cm)
	}

	if subjects := subjectResults(results); len(subjects) > 0 && len(subjects) < len(results) {
		sm := computeMetrics(subjects, cfg.threshold)
		fmt.Fprintf(w, "Analyst-judged subjects only (%d of %d; Occupancy/Position/auto-merge dropped)\n", sm.N, len(results))
		printMetrics(w, sm)
		fmt.Fprintf(w, "  labeled positive mean/median: %.3f / %.3f\n", sm.PosMean, sm.PosMedian)
		fmt.Fprintf(w, "  labeled negative mean/median: %.3f / %.3f\n", sm.NegMean, sm.NegMedian)
		best, _ := sweepThresholds(subjects)
		fmt.Fprintf(w, "  best F1: %.4f at %.2f (acc=%.3f prec=%.3f rec=%.3f)\n\n",
			best.F1, best.Threshold, best.Accuracy, best.Precision, best.Recall)
	}

	rows := schemaBreakdown(results, cfg.threshold)
	if len(rows) > 0 {
		fmt.Fprintln(w, "By schema")
		tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
		fmt.Fprintln(tw, "schema\tn\tacc\tf1\tprec\trec\tpos-mean\tneg-mean")
		for _, row := range rows {
			mm := row.Metrics
			fmt.Fprintf(tw, "%s\t%d\t%.3f\t%.3f\t%.3f\t%.3f\t%.3f\t%.3f\n",
				row.Label, mm.N, mm.Accuracy, mm.F1, mm.Precision, mm.Recall, mm.PosMean, mm.NegMean)
		}
		tw.Flush()
		fmt.Fprintln(w)
	}

	if sweep {
		best, all := sweepThresholds(results)
		fmt.Fprintf(w, "Best F1 on threshold sweep: %.4f at %.2f (acc=%.3f prec=%.3f rec=%.3f)\n",
			best.F1, best.Threshold, best.Accuracy, best.Precision, best.Recall)
		fmt.Fprintln(w, "threshold  acc     f1      prec    rec     fp   fn")
		for _, s := range all {
			if int(s.Threshold*100)%5 != 0 {
				continue
			}
			fmt.Fprintf(w, "  %.2f     %.3f   %.3f   %.3f   %.3f   %4d %4d\n",
				s.Threshold, s.Accuracy, s.F1, s.Precision, s.Recall, s.FP, s.FN)
		}
		fmt.Fprintln(w)
	}

	if mode == "search" {
		found := subset(results, func(r pairResult) bool { return r.SearchFound })
		fmt.Fprintf(w, "Search coverage: %d/%d pairs had the right-hand entity in Watchman results (%.1f%%)\n\n",
			len(found), len(results), 100*float64(len(found))/max(1, float64(len(results))))
	}

	n := cfg.debugErrors
	if n <= 0 {
		n = 5
	}
	fps := errorExamples(results, true, n)
	fns := errorExamples(results, false, n)
	if len(fps) > 0 {
		fmt.Fprintln(w, "Highest-scoring false positives (labeled negative)")
		printErrors(w, fps)
	}
	if len(fns) > 0 {
		fmt.Fprintln(w, "Lowest-scoring false negatives (labeled positive)")
		printErrors(w, fns)
	}
}

func attachDebugPieces(pairs []Pair, results []pairResult, cfg evalConfig) {
	if cfg.debugErrors <= 0 {
		return
	}
	want := make(map[int]struct{})
	for _, r := range errorExamples(results, true, cfg.debugErrors) {
		want[r.Index] = struct{}{}
	}
	for _, r := range errorExamples(results, false, cfg.debugErrors) {
		want[r.Index] = struct{}{}
	}
	if len(want) == 0 {
		return
	}
	opts := search.SimilarityOpts{Algorithm: cfg.algorithm, TFIDF: cfg.tfidf}
	conv := convertOpts{maxAltNames: cfg.maxAltNames, nameOnly: cfg.nameOnly}
	for i, r := range results {
		if _, ok := want[r.Index]; !ok {
			continue
		}
		if r.Index < 0 || r.Index >= len(pairs) {
			continue
		}
		pair := pairs[r.Index]
		left, right := convertPair(pair, conv)
		best := search.DetailedSimilarityWithOpts(nil, left, right, opts)
		for _, typ := range commonTypes(pair.Left, pair.Right) {
			l := convertEntityAs(pair.Left, conv, typ).Normalize()
			rgt := convertEntityAs(pair.Right, conv, typ).Normalize()
			d := search.DetailedSimilarityWithOpts(nil, l, rgt, opts)
			if d.FinalScore > best.FinalScore {
				best = d
			}
		}
		results[i].Pieces = best.Pieces
		results[i].Score = best.FinalScore
	}
}

func printMetrics(w io.Writer, m metrics) {
	fmt.Fprintf(w, "  accuracy=%.4f  precision=%.4f  recall=%.4f  f1=%.4f\n",
		m.Accuracy, m.Precision, m.Recall, m.F1)
	fmt.Fprintf(w, "  TP=%d  TN=%d  FP=%d  FN=%d\n\n", m.TP, m.TN, m.FP, m.FN)
}

func printHistogram(w io.Writer, label string, buckets []int) {
	fmt.Fprintf(w, "  %s ", label)
	for _, n := range buckets {
		if n == 0 {
			fmt.Fprint(w, ".")
			continue
		}
		fmt.Fprint(w, bar(n))
	}
	fmt.Fprintf(w, "  (0                    1)\n")
}

func bar(n int) string {
	switch {
	case n >= 50:
		return "#"
	case n >= 10:
		return "="
	default:
		return "+"
	}
}

func printErrors(w io.Writer, rows []pairResult) {
	for _, r := range rows {
		fmt.Fprintf(w, "  %.3f  %s  %s  (%s)\n    left  %s  %s\n    right %s  %s\n",
			r.Score, r.Judgement, r.Schema, strings.Join(r.LeftDatasets, ","),
			r.LeftID, r.LeftName, r.RightID, r.RightName)
		if len(r.Pieces) > 0 {
			for _, p := range r.Pieces {
				if p.FieldsCompared == 0 {
					continue
				}
				fmt.Fprintf(w, "      %-16s score=%.3f weight=%.0f matched=%v exact=%v fields=%d\n",
					p.PieceType, p.Score, p.Weight, p.Matched, p.Exact, p.FieldsCompared)
			}
		}
	}
	fmt.Fprintln(w)
}

func writeJSONL(w io.Writer, results []pairResult) error {
	enc := json.NewEncoder(w)
	for _, r := range results {
		if err := enc.Encode(r); err != nil {
			return err
		}
	}
	return nil
}

type runReport struct {
	Mode           string       `json:"mode"`
	Algorithm      string       `json:"algorithm"`
	All            metrics      `json:"all"`
	BestF1         metrics      `json:"bestF1"`
	Sweep          []metrics    `json:"sweep"`
	Subjects       *metrics     `json:"subjects,omitempty"`
	SubjectsBestF1 *metrics     `json:"subjectsBestF1,omitempty"`
	CrossScript    *metrics     `json:"crossScript,omitempty"`
	Schema         []metrics    `json:"schema"`
	SchemaLabels   []string     `json:"schemaLabels"`
	FalsePositives []pairResult `json:"falsePositives"`
	FalseNegatives []pairResult `json:"falseNegatives"`
}

func buildReport(results []pairResult, cfg evalConfig, mode string) runReport {
	applyThreshold(results, cfg.threshold)
	best, sweep := sweepThresholds(results)
	rep := runReport{
		Mode:           mode,
		Algorithm:      cfg.label(),
		All:            computeMetrics(results, cfg.threshold),
		BestF1:         best,
		Sweep:          everyFifth(sweep),
		FalsePositives: errorExamples(results, true, cfg.debugErrors),
		FalseNegatives: errorExamples(results, false, cfg.debugErrors),
	}
	if n := len(subset(results, func(r pairResult) bool { return r.CrossScript })); n > 0 {
		m := computeMetrics(subset(results, func(r pairResult) bool { return r.CrossScript }), cfg.threshold)
		rep.CrossScript = &m
	}
	if subjects := subjectResults(results); len(subjects) > 0 && len(subjects) < len(results) {
		m := computeMetrics(subjects, cfg.threshold)
		b, _ := sweepThresholds(subjects)
		rep.Subjects = &m
		rep.SubjectsBestF1 = &b
	}
	for _, row := range schemaBreakdown(results, cfg.threshold) {
		rep.SchemaLabels = append(rep.SchemaLabels, row.Label)
		rep.Schema = append(rep.Schema, row.Metrics)
	}
	return rep
}

func everyFifth(all []metrics) []metrics {
	var out []metrics
	for _, s := range all {
		if int(s.Threshold*100)%5 == 0 {
			out = append(out, s)
		}
	}
	return out
}

func writeJSON(path string, v any) error {
	if path == "" {
		return nil
	}
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, data)
}
