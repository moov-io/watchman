// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMixEmbed_HybridCrossScriptOnly(t *testing.T) {
	require.Equal(t, 0.9, mixEmbed(0.2, 0.9, "hybrid", true))
	require.Equal(t, 0.8, mixEmbed(0.8, 0.9, "hybrid", false))
	require.Equal(t, 0.9, mixEmbed(0.2, 0.9, "max", false))
	require.Equal(t, 0.9, mixEmbed(0.2, 0.9, "only", false))
}

func TestPairCrossScript(t *testing.T) {
	require.True(t, pairCrossScript(pair{name1: "צפרירים קירור", name2: "Zafririm Kirur"}))
	require.False(t, pairCrossScript(pair{name1: "Meadow Spirit", name2: "Meadow Spirit"}))
	require.True(t, pairHebrewLatin(pair{name1: "רוח קדים", name2: "RUAH QADIM"}))
	require.False(t, pairHebrewLatin(pair{name1: "Meadow Spirit", name2: "Meadow Spirit"}))
	require.True(t, pairBurmeseLatin(pair{name1: "မြင့်လှိုင် သစ်လုပ်ငန်း ကုမ္ပဏီလီမိတက်", name2: "Myint Hlaing Timber Co., Ltd."}))
}

func TestLoadPairs_Lot2(t *testing.T) {
	pairs, err := loadPairs("../../pkg/search/testdata/cascade-name-pairs-2.csv")
	require.NoError(t, err)
	require.Len(t, pairs, 185)

	var matches, xs int
	for _, p := range pairs {
		if p.isMatch {
			matches++
		}
		if pairCrossScript(p) {
			xs++
		}
	}
	require.Equal(t, 118, matches)
	require.Greater(t, xs, 50)
}
