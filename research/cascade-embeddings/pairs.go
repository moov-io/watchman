// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/moov-io/watchman/pkg/search"
)

type pair struct {
	schema   string
	name1    string
	name2    string
	isMatch  bool
	category string
	jw       float64
}

func loadPairs(path string) ([]pair, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, fmt.Errorf("%s: no data rows", path)
	}

	header := rows[0]
	col := func(name string) int {
		for i, h := range header {
			if h == name {
				return i
			}
		}
		return -1
	}
	idxSchema := col("schema")
	idxName1 := col("name1")
	idxName2 := col("name2")
	idxMatch := col("is_match")
	idxCat := col("category")
	if idxSchema < 0 || idxName1 < 0 || idxName2 < 0 || idxMatch < 0 {
		return nil, fmt.Errorf("%s: missing required columns", path)
	}

	out := make([]pair, 0, len(rows)-1)
	for i, row := range rows[1:] {
		p := pair{
			schema:   field(row, idxSchema),
			name1:    field(row, idxName1),
			name2:    field(row, idxName2),
			category: field(row, idxCat),
		}
		p.isMatch, err = strconv.ParseBool(field(row, idxMatch))
		if err != nil {
			return nil, fmt.Errorf("%s row %d is_match: %w", path, i+2, err)
		}
		p.jw = cascadeJW(p)
		out = append(out, p)
	}
	return out, nil
}

func field(row []string, idx int) string {
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return row[idx]
}

func cascadeJW(p pair) float64 {
	query := cascadeEntity(p.schema, p.name1)
	index := cascadeEntity(p.schema, p.name2)
	return search.Similarity(query, index)
}

func cascadeEntity(schema, name string) search.Entity[search.Value] {
	e := search.Entity[search.Value]{Name: name}
	switch strings.ToLower(strings.TrimSpace(schema)) {
	case "vessel":
		e.Type = search.EntityVessel
		e.Vessel = &search.Vessel{Name: name}
	default:
		e.Type = search.EntityBusiness
		e.Business = &search.Business{Name: name}
	}
	return e.Normalize()
}

func hasNonLatin(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return true
		}
	}
	return false
}

func pairCrossScript(p pair) bool {
	return hasNonLatin(p.name1) != hasNonLatin(p.name2)
}

func hasScript(s string, table *unicode.RangeTable) bool {
	for _, r := range s {
		if unicode.Is(table, r) {
			return true
		}
	}
	return false
}

func pairHebrewLatin(p pair) bool {
	a, b := hasScript(p.name1, unicode.Hebrew), hasScript(p.name2, unicode.Hebrew)
	return a != b && pairCrossScript(p)
}

func pairBurmeseLatin(p pair) bool {
	a, b := hasScript(p.name1, unicode.Myanmar), hasScript(p.name2, unicode.Myanmar)
	return a != b && pairCrossScript(p)
}

func uniqueNames(pairs []pair) []string {
	seen := make(map[string]struct{})
	out := make([]string, 0)
	for _, p := range pairs {
		for _, n := range []string{p.name1, p.name2} {
			n = strings.TrimSpace(n)
			if n == "" {
				continue
			}
			if _, ok := seen[n]; ok {
				continue
			}
			seen[n] = struct{}{}
			out = append(out, n)
		}
	}
	return out
}
