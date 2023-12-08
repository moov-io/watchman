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
	query := "vladimir levenshtein"
	indexedName := "vladimirov vladimir vladimirovich"
	score1 := jaroWinkler(indexedName, query)
	score2 := bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.961)
	eql(t, "New score is better: "+query, score2, 0.603)

	// 2. SDN Entity 7788 "SHAQIRI, Shaqir"
	query = "zaid shakir"
	indexedName = "shaqiri shaqir"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.908)
	eql(t, "New score is better: "+query, score2, 0.704)

	// Single-word sanctioned names shouldn't match any query with that name part
	// 1. SDN Entity 15050 "HADI"
	query = "hadi alwai"
	indexedName = "hadi"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.900)
	eql(t, "New score is better: "+query, score2, 0.615)

	// Name-part scores should be weighted by the character length. If not, small words can have unfair weight
	// 1. SDN Entity "LI, Shangfu"
	query = "li shanlan"
	indexedName = "li shangfu"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.914)
	eql(t, "New score is better: "+query, score2, 0.867)

	// Words with different lengths shouldn't match very highly
	query = "brown"
	indexedName = "browningweight"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.871)
	eql(t, "New score is better: "+query, score2, 0.703)

	// Words that start with different letters shouldn't match very highly
	query = "jimenez"
	indexedName = "dominguez"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too high: "+query, score1, 0.690)
	eql(t, "New score is better: "+query, score2, 0.580)
}

func TestBestPairsJaroWinkler__TruePositives(t *testing.T) {
	// Unmatched indexed words had a large weight, causing false negatives for missing "middle names"
	// 1. Saddam Hussein
	query := "saddam hussien"
	indexedName := "saddam hussein al tikriti"
	score1 := jaroWinkler(indexedName, query)
	score2 := bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too low: "+query, score1, 0.656)
	eql(t, "New score is better: "+query, score2, 0.924)

	// 2. SDN Entity 7574 "VALENCIA TRUJILLO, Joaquin Mario"
	query = "valencia trujillo joaquin"
	indexedName = "valencia trujillo joaquin mario"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too low: "+query, score1, 0.868)
	eql(t, "New score is better: "+query, score2, 0.973)

	// 3. SDN Entity 9760 "LUKASHENKO, Alexander Grigoryevich"
	query = "alexander lukashenko"
	indexedName = "lukashenko alexander grigoryevich"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too low: "+query, score1, 0.765)
	eql(t, "New score is better: "+query, score2, 0.942)

	// Small words had too much weight, causing false negatives
	// 1. SDN Entity 4691 "A.I.C. SOGO KENKYUSHO"
	query = "sogo kenkyusho"
	indexedName = "a i c sogo kenkyusho"
	score1 = jaroWinkler(indexedName, query)
	score2 = bestPairsJaroWinkler(strings.Fields(query), indexedName)
	eql(t, "Score is too low: "+query, score1, 0.400)
	eql(t, "New score is better: "+query, score2, 0.972)
}
