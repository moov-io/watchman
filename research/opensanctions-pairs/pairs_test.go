// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadPairs_Array(t *testing.T) {
	pairs, err := loadPairs(filepath.Join("testdata", "pairs.json"))
	require.NoError(t, err)
	require.Len(t, pairs, 5)
	require.Equal(t, "positive", pairs[0].Judgement)
	require.Equal(t, "ofac-40604", pairs[0].Left.ID)
}

func TestLoadPairs_SampleWrapper(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.json")
	raw := `{"metadata":{"n_total":1},"pairs":[{"left":{"id":"a","caption":"A","schema":"Person"},"right":{"id":"b","caption":"B","schema":"Person"},"judgement":"negative"}]}`
	require.NoError(t, os.WriteFile(path, []byte(raw), 0o644))

	pairs, err := loadPairs(path)
	require.NoError(t, err)
	require.Len(t, pairs, 1)
	require.Equal(t, "a", pairs[0].Left.ID)
	require.Equal(t, "negative", pairs[0].Judgement)
}

func TestLoadPairs_JSONL(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "pairs.jsonl")
	raw := `{"left":{"id":"1","caption":"Ann","schema":"Person"},"right":{"id":"2","caption":"Anne","schema":"Person"},"judgement":"positive"}` + "\n"
	require.NoError(t, os.WriteFile(path, []byte(raw), 0o644))

	pairs, err := loadPairs(path)
	require.NoError(t, err)
	require.Len(t, pairs, 1)
	require.Equal(t, "Ann", pairs[0].Left.Caption)
}
