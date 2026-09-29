package search

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	cascadePairsFile    = "testdata/cascade-name-pairs.csv"
	cascadePairsDocFile = "testdata/cascade-name-pairs.md"

	cascadeScoreSummaryBegin = "<!-- cascade-score-summary -->"
	cascadeScoreSummaryEnd   = "<!-- /cascade-score-summary -->"
)

type cascadePair struct {
	caseGroup string
	schema    string
	name1     string
	name2     string
	isMatch   bool
	quality   string
	category  string
	notes     string
	baseline  float64
	score     float64
	hasBase   bool
	hasScore  bool
	line      int
}

func TestCascadeNamePairs(t *testing.T) {
	pairs := loadCascadePairs(t)
	require.Len(t, pairs, 300)

	update := strings.EqualFold(os.Getenv("UPDATE_CASCADE_SCORES"), "yes")
	var mismatches int

	for i := range pairs {
		p := &pairs[i]
		got := cascadePairScore(p)
		t.Run(fmt.Sprintf("%03d_%s", i+1, p.category), func(t *testing.T) {
			if update {
				if !p.hasBase {
					p.baseline = roundCascadeScore(got)
					p.hasBase = true
				}
				p.score = roundCascadeScore(got)
				p.hasScore = true
				return
			}

			require.True(t, p.hasScore, "row %d (%s) missing score; run UPDATE_CASCADE_SCORES=yes", p.line, p.category)
			require.InDelta(t, p.score, got, 0.00015, "%s %q vs %q", p.category, p.name1, p.name2)
			require.InDelta(t, got, scoreSimilarityFast(cascadeEntity(p.schema, p.name1), cascadeEntity(p.schema, p.name2), SimilarityOpts{}), 0.00015)
		})
		if !update && p.hasScore && math.Abs(p.score-got) > 0.00015 {
			mismatches++
		}
	}

	if update {
		writeCascadePairs(t, pairs)
		writeCascadeScoreSummary(t, pairs)
	}

	t.Logf("cascade name pairs: %d rows, %d score mismatches vs fixture", len(pairs), mismatches)
	t.Log("\n" + cascadeCategoryTable(pairs))
}

func TestCascadeNamePairs_VesselOwnerIsTypedZero(t *testing.T) {
	pairs := loadCascadePairs(t)

	var checked int
	for i := range pairs {
		p := pairs[i]
		if p.category != "vessel-vs-owner" && p.category != "vessel-vs-manager" {
			continue
		}
		checked++
		query := cascadeEntity("Vessel", p.name1)
		index := cascadeEntity("Company", p.name2)
		got := Similarity(query, index)
		require.Equal(t, 0.0, got, "%s %q vs %q (typed vessel vs business)", p.category, p.name1, p.name2)
	}
	require.Equal(t, 8, checked)
}

func cascadePairScore(p *cascadePair) float64 {
	query := cascadeEntity(p.schema, p.name1)
	index := cascadeEntity(p.schema, p.name2)
	return Similarity(query, index)
}

func cascadeEntity(schema, name string) Entity[Value] {
	e := Entity[Value]{Name: name}
	switch strings.ToLower(strings.TrimSpace(schema)) {
	case "vessel":
		e.Type = EntityVessel
		e.Vessel = &Vessel{Name: name}
	default:
		e.Type = EntityBusiness
		e.Business = &Business{Name: name}
	}
	return e.Normalize()
}

func loadCascadePairs(t *testing.T) []cascadePair {
	t.Helper()

	f, err := os.Open(cascadePairsFile)
	require.NoError(t, err)
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	rows, err := r.ReadAll()
	require.NoError(t, err)
	require.Greater(t, len(rows), 1)

	header := rows[0]
	col := func(name string) int {
		for i, h := range header {
			if h == name {
				return i
			}
		}
		return -1
	}

	idxGroup := col("case_group")
	idxSchema := col("schema")
	idxName1 := col("name1")
	idxName2 := col("name2")
	idxMatch := col("is_match")
	idxQuality := col("quality")
	idxCategory := col("category")
	idxNotes := col("notes")
	idxBase := col("baseline")
	idxScore := col("score")
	require.GreaterOrEqual(t, idxGroup, 0)
	require.GreaterOrEqual(t, idxSchema, 0)
	require.GreaterOrEqual(t, idxName1, 0)
	require.GreaterOrEqual(t, idxName2, 0)
	require.GreaterOrEqual(t, idxMatch, 0)
	require.GreaterOrEqual(t, idxCategory, 0)

	out := make([]cascadePair, 0, len(rows)-1)
	for i, row := range rows[1:] {
		p := cascadePair{
			line:      i + 2,
			caseGroup: csvField(row, idxGroup),
			schema:    csvField(row, idxSchema),
			name1:     csvField(row, idxName1),
			name2:     csvField(row, idxName2),
			quality:   csvField(row, idxQuality),
			category:  csvField(row, idxCategory),
			notes:     csvField(row, idxNotes),
		}
		p.isMatch, err = strconv.ParseBool(csvField(row, idxMatch))
		require.NoError(t, err, "row %d is_match", p.line)
		p.baseline, p.hasBase = csvFloat(t, row, idxBase, p.line, "baseline")
		p.score, p.hasScore = csvFloat(t, row, idxScore, p.line, "score")
		out = append(out, p)
	}
	return out
}

func csvField(row []string, idx int) string {
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return row[idx]
}

func csvFloat(t *testing.T, row []string, idx, line int, name string) (float64, bool) {
	t.Helper()
	raw := strings.TrimSpace(csvField(row, idx))
	if raw == "" {
		return 0, false
	}
	n, err := strconv.ParseFloat(raw, 64)
	require.NoError(t, err, "row %d %s", line, name)
	return n, true
}

func writeCascadePairs(t *testing.T, pairs []cascadePair) {
	t.Helper()

	f, err := os.Create(cascadePairsFile)
	require.NoError(t, err)
	defer f.Close()

	w := csv.NewWriter(f)
	require.NoError(t, w.Write([]string{
		"case_group", "schema", "name1", "name2", "is_match", "quality", "category", "notes", "baseline", "score",
	}))
	for i := range pairs {
		p := pairs[i]
		require.NoError(t, w.Write([]string{
			p.caseGroup,
			p.schema,
			p.name1,
			p.name2,
			strconv.FormatBool(p.isMatch),
			p.quality,
			p.category,
			p.notes,
			formatCascadeScore(p.baseline),
			formatCascadeScore(p.score),
		}))
	}
	w.Flush()
	require.NoError(t, w.Error())
}

func writeCascadeScoreSummary(t *testing.T, pairs []cascadePair) {
	t.Helper()

	body, err := os.ReadFile(cascadePairsDocFile)
	require.NoError(t, err)
	text := string(body)
	start := strings.Index(text, cascadeScoreSummaryBegin)
	end := strings.Index(text, cascadeScoreSummaryEnd)
	require.GreaterOrEqual(t, start, 0, "missing %s in %s", cascadeScoreSummaryBegin, cascadePairsDocFile)
	require.Greater(t, end, start)

	var b strings.Builder
	b.WriteString(text[:start])
	b.WriteString(cascadeScoreSummaryBegin)
	b.WriteString("\n\n")
	b.WriteString(cascadeCategoryTable(pairs))
	b.WriteString("\n")
	b.WriteString(text[end:])
	require.NoError(t, os.WriteFile(cascadePairsDocFile, []byte(b.String()), 0o600))
}

func cascadeCategoryTable(pairs []cascadePair) string {
	type agg struct {
		n, matches   int
		sum, sumBase float64
		min, max     float64
		sumDelta     float64
	}
	byCat := make(map[string]*agg)
	var all, pos, neg agg
	all.min, pos.min, neg.min = 1, 1, 1

	add := func(a *agg, p cascadePair) {
		a.n++
		a.sum += p.score
		a.sumBase += p.baseline
		a.sumDelta += p.score - p.baseline
		if p.isMatch {
			a.matches++
		}
		if a.n == 1 || p.score < a.min {
			a.min = p.score
		}
		if p.score > a.max {
			a.max = p.score
		}
	}

	for i := range pairs {
		p := pairs[i]
		g := byCat[p.category]
		if g == nil {
			g = &agg{min: 1}
			byCat[p.category] = g
		}
		add(g, p)
		add(&all, p)
		if p.isMatch {
			add(&pos, p)
		} else {
			add(&neg, p)
		}
	}

	cats := make([]string, 0, len(byCat))
	for c := range byCat {
		cats = append(cats, c)
	}
	sort.Strings(cats)

	var b strings.Builder
	fmt.Fprintf(&b, "Recorded `Similarity` on %d name-only pairs (schema → Watchman type).\n\n", len(pairs))
	fmt.Fprintf(&b, "| Slice | N | Matches | Mean score | Mean baseline | Mean delta | Min | Max |\n")
	fmt.Fprintf(&b, "|---|---:|---:|---:|---:|---:|---:|---:|\n")
	writeAgg := func(name string, a agg) {
		if a.n == 0 {
			return
		}
		fmt.Fprintf(&b, "| %s | %d | %d | %.4f | %.4f | %+.4f | %.4f | %.4f |\n",
			name, a.n, a.matches, a.sum/float64(a.n), a.sumBase/float64(a.n), a.sumDelta/float64(a.n), a.min, a.max)
	}
	writeAgg("all", all)
	writeAgg("is_match=true", pos)
	writeAgg("is_match=false", neg)
	for _, c := range cats {
		writeAgg(c, *byCat[c])
	}
	return b.String()
}

func formatCascadeScore(f float64) string {
	return strconv.FormatFloat(roundCascadeScore(f), 'f', 4, 64)
}

func roundCascadeScore(f float64) float64 {
	return math.Round(f*10000) / 10000
}

func TestRoundCascadeScore(t *testing.T) {
	require.Equal(t, 0.8550, roundCascadeScore(0.85496))
}

func TestCascadeEntityTypes(t *testing.T) {
	v := cascadeEntity("Vessel", "NS LEADER")
	require.Equal(t, EntityVessel, v.Type)
	require.NotNil(t, v.Vessel)

	b := cascadeEntity("Company", "Acme Ltd")
	require.Equal(t, EntityBusiness, b.Type)
	require.NotNil(t, b.Business)
}
