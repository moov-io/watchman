// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

type sliceStats struct {
	n, matches             int
	sum, sumPos, sumNeg    float64
	nPos, nNeg             int
	tp80, fp80, tp59, fp59 int
}

func (s sliceStats) mean() float64 {
	if s.n == 0 {
		return 0
	}
	return s.sum / float64(s.n)
}

func (s sliceStats) meanPos() float64 {
	if s.nPos == 0 {
		return 0
	}
	return s.sumPos / float64(s.nPos)
}

func (s sliceStats) meanNeg() float64 {
	if s.nNeg == 0 {
		return 0
	}
	return s.sumNeg / float64(s.nNeg)
}

func (s sliceStats) prec80() float64 {
	if s.tp80+s.fp80 == 0 {
		return 0
	}
	return float64(s.tp80) / float64(s.tp80+s.fp80)
}

func (s sliceStats) rec80() float64 {
	if s.nPos == 0 {
		return 0
	}
	return float64(s.tp80) / float64(s.nPos)
}

func addScore(s *sliceStats, match bool, score float64) {
	s.n++
	s.sum += score
	if match {
		s.matches++
		s.nPos++
		s.sumPos += score
		if score >= 0.80 {
			s.tp80++
		}
		if score >= 0.59 {
			s.tp59++
		}
	} else {
		s.nNeg++
		s.sumNeg += score
		if score >= 0.80 {
			s.fp80++
		}
		if score >= 0.59 {
			s.fp59++
		}
	}
}

type modelReport struct {
	name                     string
	dim                      int
	all, xs, hebrew, burmese sliceStats
}

func scoreModel(name string, dim int, pairs []pair, cos func(a, b string) float64, mode string) modelReport {
	r := modelReport{name: name, dim: dim}
	for _, p := range pairs {
		c := 0.0
		if cos != nil {
			c = cos(p.name1, p.name2)
		}
		xs := pairCrossScript(p)
		score := mixEmbed(p.jw, c, mode, xs)
		addScore(&r.all, p.isMatch, score)
		if xs {
			addScore(&r.xs, p.isMatch, score)
		}
		if pairHebrewLatin(p) {
			addScore(&r.hebrew, p.isMatch, score)
		}
		if pairBurmeseLatin(p) {
			addScore(&r.burmese, p.isMatch, score)
		}
	}
	return r
}

func renderReports(title string, reports []modelReport) string {
	var b strings.Builder
	fmt.Fprintf(&b, "## %s\n\n", title)
	fmt.Fprintf(&b, "Hybrid mix is max(Jaro–Winkler, cosine) only when the two names differ in script (Latin vs not). Latin/Latin pairs stay on Jaro–Winkler.\n\n")
	fmt.Fprintf(&b, "| Model | Dim | Match mean | Non-match mean | Cross-script match mean | Hebrew/Latin match mean | Burmese/Latin match mean | TP @ 0.80 | FP @ 0.80 | Precision @ 0.80 | Recall @ 0.80 | Cross-script TP @ 0.80 | Cross-script recall @ 0.80 |\n")
	fmt.Fprintf(&b, "|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, r := range reports {
		fmt.Fprintf(&b, "| %s | %d | %.4f | %.4f | %.4f | %.4f | %.4f | %d | %d | %.3f | %.3f | %d | %.3f |\n",
			r.name, r.dim, r.all.meanPos(), r.all.meanNeg(), r.xs.meanPos(), r.hebrew.meanPos(), r.burmese.meanPos(),
			r.all.tp80, r.all.fp80, r.all.prec80(), r.all.rec80(), r.xs.tp80, r.xs.rec80())
	}
	fmt.Fprintf(&b, "\nAt 0.59 (high-recall line):\n\n")
	fmt.Fprintf(&b, "| Model | TP @ 0.59 | FP @ 0.59 | Cross-script TP @ 0.59 | Cross-script recall @ 0.59 |\n")
	fmt.Fprintf(&b, "|---|---:|---:|---:|---:|\n")
	for _, r := range reports {
		xsRec := 0.0
		if r.xs.nPos > 0 {
			xsRec = float64(r.xs.tp59) / float64(r.xs.nPos)
		}
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %.3f |\n", r.name, r.all.tp59, r.all.fp59, r.xs.tp59, xsRec)
	}
	b.WriteByte('\n')
	return b.String()
}
