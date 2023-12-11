// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

func TestBestPairsJaroWinkler__FalsePositives(t *testing.T) {
	// Words in the query should be matched against at most one indexed word. Doubled names on the sanctioned list can
	// skew results
	// 1. SDN Entity 40273, VLADIMIROV, Vladimir Vladimirovich
	oldScore, newScore := compareAlgorithms("vladimirov vladimir vladimirovich", "vladimir levenshtein")
	eql(t, "Score is too high", oldScore, 0.961)
	eql(t, "New score is better", newScore, 0.603)

	// 2. SDN Entity 7788 "SHAQIRI, Shaqir"
	oldScore, newScore = compareAlgorithms("shaqiri shaqir", "zaid shakir")
	eql(t, "Score is too high", oldScore, 0.908)
	eql(t, "New score is better", newScore, 0.704)

	// Single-word sanctioned names shouldn't match any query with that name part
	// 1. SDN Entity 15050 "HADI"
	oldScore, newScore = compareAlgorithms("hadi", "hadi alwai")
	eql(t, "Score is too high", oldScore, 0.900)
	eql(t, "New score is better", newScore, 0.615)

	// Name-part scores should be weighted by the character length. If not, small words can have unfair weight
	// 1. SDN Entity "LI, Shangfu"
	oldScore, newScore = compareAlgorithms("li shangfu", "li shanlan")
	eql(t, "Score is too high", oldScore, 0.914)
	eql(t, "New score is better", newScore, 0.867)

	// Words with different lengths shouldn't match very highly
	oldScore, newScore = compareAlgorithms("browningweight", "brown")
	eql(t, "Score is too high", oldScore, 0.871)
	eql(t, "New score is better", newScore, 0.703)

	// Words that start with different letters shouldn't match very highly
	oldScore, newScore = compareAlgorithms("dominguez", "jimenez")
	eql(t, "Score is too high", oldScore, 0.690)
	eql(t, "New score is better", newScore, 0.580)
}

func TestBestPairsJaroWinkler__TruePositives(t *testing.T) {
	// Unmatched indexed words had a large weight, causing false negatives for missing "middle names"
	// 1. Saddam Hussein
	oldScore, newScore := compareAlgorithms("saddam hussein al tikriti", "saddam hussien")
	eql(t, "Score is too low", oldScore, 0.656)
	eql(t, "New score is better", newScore, 0.924)

	// 2. SDN Entity 7574 "VALENCIA TRUJILLO, Joaquin Mario"
	oldScore, newScore = compareAlgorithms("valencia trujillo joaquin mario", "valencia trujillo joaquin")
	eql(t, "Score is too low", oldScore, 0.868)
	eql(t, "New score is better", newScore, 0.973)

	// 3. SDN Entity 9760 "LUKASHENKO, Alexander Grigoryevich"
	oldScore, newScore = compareAlgorithms("lukashenko alexander grigoryevich", "alexander lukashenko")
	eql(t, "Score is too low", oldScore, 0.765)
	eql(t, "New score is better", newScore, 0.942)

	// Small words had too much weight, causing false negatives
	// 1. SDN Entity 4691 "A.I.C. SOGO KENKYUSHO"
	oldScore, newScore = compareAlgorithms("a i c sogo kenkyusho", "sogo kenkyusho")
	eql(t, "Score is too low", oldScore, 0.400)
	eql(t, "New score is better", newScore, 0.972)
}

func compareAlgorithms(indexedName string, query string) (float64, float64) {
	oldScore := jaroWinkler(indexedName, query)
	newScore := bestPairsJaroWinkler(strings.Fields(query), indexedName)
	return oldScore, newScore
}
