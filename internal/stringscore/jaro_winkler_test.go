// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package stringscore_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"

	"github.com/stretchr/testify/require"
)

func TestJaroWinkler(t *testing.T) {
	cases := []struct {
		indexed, search string
		match           float64
	}{
		// examples
		{"wei, zhao", "wei, Zhao", 0.875},
		{"WEI, Zhao", "WEI, Zhao", 1.0},
		{"WEI Zhao", "WEI Zhao", 1.0},
		{strings.ToLower("WEI Zhao"), prepare.LowerAndRemovePunctuation("WEI, Zhao"), 1.0},

		// apply jaroWinkler in both directions
		{"jane doe", "jan lahore", 0.439},
		{"jan lahore", "jane doe", 0.549},

		// real world case
		{"john doe", "paul john", 0.624},
		{"john doe", "john othername", 0.440},

		// close match
		{"jane doe", "jane doe2", 0.940},

		// real-ish world examples
		{"kalamity linden", "kala limited", 0.687},
		{"kala limited", "kalamity linden", 0.687},

		// examples used in demos / commonly
		{"nicolas", "nicolas", 1.0},
		{"nicolas moros maduro", "nicolas maduro", 0.958},
		{"nicolas maduro", "nicolas moros maduro", 0.839},

		// customer examples
		{"ian", "ian mckinley", 0.429},
		{"iap", "ian mckinley", 0.352},
		{"ian mckinley", "ian", 0.891},
		{"ian mckinley", "iap", 0.733},
		{"ian mckinley", "tian xiang 7", 0.000},
		{"bindaree food group pty", prepare.LowerAndRemovePunctuation("independent insurance group ltd"), 0.269}, // removes ltd
		{"bindaree food group pty ltd", "independent insurance group ltd", 0.401},                                // only matches higher from 'ltd'
		{"p.c.c. (singapore) private limited", "culver max entertainment private limited", 0.514},
		{"zincum llc", "easy verification inc.", 0.000},
		{"transpetrochart co ltd", "jx metals trading co.", 0.431},
		{"technolab", "moomoo technologies inc", 0.565},
		{"sewa security services", "sesa - safety & environmental services australia pty ltd", 0.480},
		{"bueno", "20/f rykadan capital twr135 hoi bun rd, kwun tong 135 hoi bun rd., kwun tong", 0.094},

		// example cases
		{"nicolas maduro", "nicolás maduro", 0.937},
		{"nicolas maduro", prepare.LowerAndRemovePunctuation("nicolás maduro"), 1.0},
		{"nic maduro", "nicolas maduro", 0.872},
		{"nick maduro", "nicolas maduro", 0.859},
		{"nicolas maduroo", "nicolas maduro", 0.966},
		{"nicolas maduro", "nicolas maduro", 1.0},
		{"maduro, nicolas", "maduro, nicolas", 1.0},
		{"maduro moros, nicolas", "maduro moros, nicolas", 1.0},
		{"maduro moros, nicolas", "nicolas maduro", 0.953},
		{"nicolas maduro moros", "maduro", 0.900},
		{"nicolas maduro moros", "nicolás maduro", 0.898},
		{"nicolas, maduro moros", "maduro", 0.897},
		{"nicolas, maduro moros", "nicolas maduro", 0.928},
		{"nicolas, maduro moros", "nicolás", 0.822},
		{"nicolas, maduro moros", "maduro", 0.897},
		{"nicolas, maduro moros", "nicolás maduro", 0.906},
		{"africada financial services bureau change", "skylight", 0.441},
		{"africada financial services bureau change", "skylight financial inc", 0.658},
		{"africada financial services bureau change", "skylight services inc", 0.599},
		{"africada financial services bureau change", "skylight financial services", 0.761},
		{"africada financial services bureau change", "skylight financial services inc", 0.730},

		// stopwords tests
		{"the group for the preservation of the holy sites", "the bridgespan group", 0.682},
		{prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"), prepare.LowerAndRemovePunctuation("the bridgespan group"), 0.682},
		{"group preservation holy sites", "bridgespan group", 0.652},

		{"the group for the preservation of the holy sites", "the logan group", 0.670},
		{prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"), prepare.LowerAndRemovePunctuation("the logan group"), 0.670},
		{"group preservation holy sites", "logan group", 0.586},

		{"the group for the preservation of the holy sites", "the anything group", 0.546},
		{prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"), prepare.LowerAndRemovePunctuation("the anything group"), 0.546},
		{"group preservation holy sites", "anything group", 0.488},

		{"the group for the preservation of the holy sites", "the hello world group", 0.637},
		{prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"), prepare.LowerAndRemovePunctuation("the hello world group"), 0.637},
		{"group preservation holy sites", "hello world group", 0.577},

		{"the group for the preservation of the holy sites", "the group", 0.880},
		{prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"), prepare.LowerAndRemovePunctuation("the group"), 0.880},
		{"group preservation holy sites", "group", 0.879},

		{"the group for the preservation of the holy sites", "The flibbity jibbity flobbity jobbity grobbity zobbity group", 0.345},
		{
			prepare.LowerAndRemovePunctuation("the group for the preservation of the holy sites"),
			prepare.LowerAndRemovePunctuation("the flibbity jibbity flobbity jobbity grobbity zobbity group"),
			0.366,
		},
		{"group preservation holy sites", "flibbity jibbity flobbity jobbity grobbity zobbity group", 0.263},

		// prepare.LowerAndRemovePunctuation
		{"i c sogo kenkyusho", prepare.LowerAndRemovePunctuation("A.I.C. SOGO KENKYUSHO"), 0.968},
		{prepare.LowerAndRemovePunctuation("A.I.C. SOGO KENKYUSHO"), "sogo kenkyusho", 0.972},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11 420 2 1 corp", 1.000},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11 420 21 corp", 0.947},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11420 2 1 corp", 0.849},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11420 21 corp", 0.787},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "114202 1 corp", 0.802},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "1142021 corp", 0.752},

		// From https://github.com/moov-io/watchman/issues/594
		{"JSCARGUMENT", "JSC ARGUMENT", 0.413},
		{"ARGUMENTJSC", "JSC ARGUMENT", 0.750},
	}
	for i := range cases {
		v := cases[i]
		// Only need to call chomp on s1, see jaroWinkler doc
		eql(t, fmt.Sprintf("#%d %s vs %s", i, v.indexed, v.search),
			stringscore.BestPairsJaroWinkler(strings.Fields(v.search), strings.Fields(v.indexed)), v.match)
	}
}

func TestBestPairCombinationJaroWinkler(t *testing.T) {
	cases := []struct {
		indexed, search string
		match           float64
	}{
		// prepare.LowerAndRemovePunctuation
		{"i c sogo kenkyusho", prepare.LowerAndRemovePunctuation("A.I.C. SOGO KENKYUSHO"), 0.968},
		{prepare.LowerAndRemovePunctuation("A.I.C. SOGO KENKYUSHO"), "sogo kenkyusho", 0.972},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11 420 2 1 corp", 1.0},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11 420 21 corp", 1.0},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11420 2 1 corp", 1.0},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "11420 21 corp", 1.0},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "114202 1 corp", 1.0},
		{prepare.LowerAndRemovePunctuation("11,420.2-1 CORP."), "1142021 corp", 1.0},

		// From https://github.com/moov-io/watchman/issues/594
		{"JSCARGUMENT", "JSC ARGUMENT", 1.000},
		{"ARGUMENTJSC", "JSC ARGUMENT", 0.750},
	}
	for i := range cases {
		v := cases[i]

		// Only need to call chomp on s1, see jaroWinkler doc
		eql(t, fmt.Sprintf("#%d %s vs %s", i, v.indexed, v.search),
			stringscore.BestPairCombinationJaroWinkler(strings.Fields(v.search), strings.Fields(v.indexed)), v.match)
	}
}

func TestJaroWinklerWithFavoritism(t *testing.T) {
	favoritism := 1.0
	delta := 0.01

	score := stringscore.JaroWinklerWithFavoritism("Vladimir Putin", "PUTIN, Vladimir Vladimirovich", favoritism)
	require.InDelta(t, score, 1.00, delta)

	score = stringscore.JaroWinklerWithFavoritism("nicolas, maduro moros", "nicolás maduro", 0.25)
	require.InDelta(t, score, 0.96, delta)

	score = stringscore.JaroWinklerWithFavoritism("Vladimir Putin", "A.I.C. SOGO KENKYUSHO", favoritism)
	require.InDelta(t, score, 0.00, delta)
}

func TestJaroWinklerErr(t *testing.T) {
	v := stringscore.JaroWinkler("", "hello")
	eql(t, "NaN #1", v, 0.0)

	v = stringscore.JaroWinkler("hello", "")
	eql(t, "NaN #1", v, 0.0)
}

func eql(t *testing.T, desc string, x, y float64) {
	t.Helper()
	if math.IsNaN(x) || math.IsNaN(y) {
		t.Fatalf("%s: x=%.2f y=%.2f", desc, x, y)
	}
	if math.Abs(x-y) > 0.01 {
		t.Errorf("%s: %.3f != %.3f", desc, x, y)
	}
}

func TestEql(t *testing.T) {
	eql(t, "", 0.1, 0.1)
	eql(t, "", 0.0001, 0.00002)
}

func BenchmarkJaroWinkler(b *testing.B) {
	inputs := []string{
		"Seyed Mohammad HASHEMI",
		"KOREA RUNGRADO GENERAL TRADING CORPORATION",
		"KOREA HAEGUMGANG TRADING CORPORATION",
		"Sendy FLORES CASTRO",
		"ERVIN DANESH ARYAN COMPANY",
		"Husam 'Adbd-al-Barr AL-FAKHURI",
		"Oleg Anatolievich KAMSHILOV",
		"SHAHID KARIMI INDUSTRIES",
		"ORAMA PROPERTIES LTD",
		"Pyong Chan KIM",
		"HAO FAN 6",
		"KOTI CORP",
		"IRAN HORMUZ 12",
		"NARI SHIPPING AND CHARTERING GMBH & CO. KG",
		"PROMSYRIOIMPORT",
		"AO ZAVOD FIOLENT",
		"NARIA GENERAL TRADING LLC",
		"Faruq AL-SURI",
		"SANDINO",
		"CAPITAL S.A.L.",
		"Ali DARASSA",
		"Ali Akbar Rezaei TAVANA",
		"CENTRAL PUBLIC PROSECUTORS OFFICE",
		"Seyyed Mohammad ATABAK",
		"PARSIAN TOURISM AND RECREATIONAL CENTERS COMPANY",
		"THE YANGON GALLERY",
	}
	b.ResetTimer()

	b.Run("BestPairsJaroWinkler", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			query := i % (len(inputs) - 1)
			index := (i + 1) % (len(inputs) - 1)

			queryTokens := strings.Fields(inputs[query])
			indexTokens := strings.Fields(inputs[index])
			b.StartTimer()

			score := stringscore.BestPairsJaroWinkler(queryTokens, indexTokens)
			require.Greater(b, score, -0.01)
		}
	})

	b.Run("BestPairCombinationJaroWinkler", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			query := i % (len(inputs) - 1)
			index := (i + 1) % (len(inputs) - 1)

			queryTokens := strings.Fields(inputs[query])
			indexTokens := strings.Fields(inputs[index])
			b.StartTimer()

			score := stringscore.BestPairCombinationJaroWinkler(queryTokens, indexTokens)
			require.Greater(b, score, -0.01)
		}
	})
}

func TestGenerateWordCombinations(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "JSC ARGUMENT",
			input: []string{"JSC", "ARGUMENT"},
			expected: [][]string{
				{"JSC", "ARGUMENT"},
				{"JSCARGUMENT"},
			},
		},
		{
			name:  "11,420.2-1 CORP",
			input: strings.Fields(prepare.LowerAndRemovePunctuation("11,420.2-1 CORP")),
			expected: [][]string{
				{"11", "420", "2", "1", "corp"},
				{"11420", "21", "corp"},
				{"1142021", "corp"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.GenerateWordCombinations(tc.input)
			require.ElementsMatch(t, tc.expected, got)
		})
	}
}
