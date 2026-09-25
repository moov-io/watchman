// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScorePairwise_Fixture(t *testing.T) {
	pairs, err := loadPairs(filepath.Join("testdata", "pairs.json"))
	require.NoError(t, err)
	pairs = filterPairs(pairs, true, nil, 0, 0)
	require.Len(t, pairs, 4)

	results := scorePairwise(pairs, evalConfig{
		threshold:   0.80,
		maxAltNames: 20,
		workers:     2,
	})
	require.Len(t, results, 4)

	byLeft := map[string]pairResult{}
	for _, r := range results {
		byLeft[r.LeftID] = r
	}

	pos := byLeft["ofac-40604"]
	require.Equal(t, "positive", pos.Judgement)
	require.Greater(t, pos.Score, 0.90, "shared passport + similar name should score high")

	sameName := byLeft["pk-cnic-1"]
	require.Equal(t, "negative", sameName.Judgement)
	require.Less(t, sameName.Score, 0.50, "conflicting national IDs pull the score down")

	nameOnly := scorePairwise(pairs, evalConfig{
		threshold:   0.80,
		nameOnly:    true,
		maxAltNames: 20,
		workers:     1,
	})
	var nameOnlyKhalid float64
	for _, r := range nameOnly {
		if r.LeftID == "pk-cnic-1" {
			nameOnlyKhalid = r.Score
		}
	}
	require.Greater(t, nameOnlyKhalid, sameName.Score, "name-only scoring ignores the ID conflict")

	company := byLeft["ru-inn-7811550255"]
	require.Equal(t, "positive", company.Judgement)
	require.Greater(t, company.Score, 0.90, "shared INN should exact-match")

	diffPeople := byLeft["koval-1"]
	require.Equal(t, "negative", diffPeople.Judgement)
	require.Less(t, diffPeople.Score, pos.Score)
}

func TestParseCompareSpec(t *testing.T) {
	base := evalConfig{threshold: 0.8, maxAltNames: 5, workers: 1}
	cfg, err := parseCompareSpec("jaro-winkler+tfidf", base, buildTFIDF([]Pair{{
		Left:      FTMEntity{Caption: "Alpha Ltd", Schema: "Company", Properties: map[string][]string{"name": {"Alpha Ltd"}}},
		Right:     FTMEntity{Caption: "Alpha Limited", Schema: "Company", Properties: map[string][]string{"name": {"Alpha Limited"}}},
		Judgement: "positive",
	}}, 5), nil)
	require.NoError(t, err)
	require.True(t, cfg.tfidf.Enabled())
	require.Equal(t, "jaro-winkler+tfidf", cfg.label())

	emb := newEmbedCache("", "test", "", 2)
	emb.vecs = map[string][]float32{"A": {1, 0}, "B": {0, 1}}
	cfg, err = parseCompareSpec("embed-hybrid", base, nil, emb)
	require.NoError(t, err)
	require.Equal(t, "hybrid", cfg.embedMode)

	idx := buildTFIDF([]Pair{{
		Left:      FTMEntity{Caption: "Alpha Ltd", Schema: "Company", Properties: map[string][]string{"name": {"Alpha Ltd"}}},
		Right:     FTMEntity{Caption: "Alpha Limited", Schema: "Company", Properties: map[string][]string{"name": {"Alpha Limited"}}},
		Judgement: "positive",
	}}, 5)
	cfg, err = parseCompareSpec("embed-hybrid+tfidf", base, idx, emb)
	require.NoError(t, err)
	require.Equal(t, "hybrid", cfg.embedMode)
	require.True(t, cfg.tfidf.Enabled())
	require.Equal(t, "embed-hybrid+tfidf", cfg.label())
}

func TestMixEmbed(t *testing.T) {
	require.InDelta(t, 0.9, mixEmbed(0.4, 0.9, "only", true), 1e-9)
	require.InDelta(t, 0.9, mixEmbed(0.4, 0.9, "max", false), 1e-9)
	require.InDelta(t, 0.4, mixEmbed(0.4, 0.9, "hybrid", false), 1e-9)
	require.InDelta(t, 0.9, mixEmbed(0.4, 0.9, "hybrid", true), 1e-9)
}

func TestComputeMetrics(t *testing.T) {
	results := []pairResult{
		{Judgement: "positive", Predicted: "positive", Score: 0.9},
		{Judgement: "positive", Predicted: "negative", Score: 0.2},
		{Judgement: "negative", Predicted: "negative", Score: 0.1},
		{Judgement: "negative", Predicted: "positive", Score: 0.85},
	}
	m := computeMetrics(results, 0.8)
	require.Equal(t, 1, m.TP)
	require.Equal(t, 1, m.TN)
	require.Equal(t, 1, m.FP)
	require.Equal(t, 1, m.FN)
	require.InDelta(t, 0.5, m.Accuracy, 1e-9)
	require.InDelta(t, 0.5, m.Precision, 1e-9)
	require.InDelta(t, 0.5, m.Recall, 1e-9)
}
