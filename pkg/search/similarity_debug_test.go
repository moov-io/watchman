package search_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestNameDebug_FromJSON(t *testing.T) {
	query := readEntity(t, "1-query.json").Normalize()
	index := readEntity(t, "1-index.json").Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.InDelta(t, 0.690, got.FinalScore, 0.001)

	require.NotNil(t, got.Name)
	require.False(t, got.Name.Exact)
	require.Equal(t, "primary", got.Name.QueryName)
	require.Equal(t, "alt", got.Name.IndexName)
	require.Equal(t, "Mohamed Salem", got.Name.Query.Original)
	require.Equal(t, []string{"mohamed", "salem"}, got.Name.Query.Tokens)
	require.Equal(t, index.Name, got.Name.Index.Original)
	require.Equal(t, "Mohamed SALEM", got.Name.Index.Compared)
	require.Equal(t, "mohamed salem", nameStageValue(got.Name.Index, "normalize"))
	require.NotContains(t, nameStageValue(got.Name.Index, "normalize"), "rahman")
	require.NotEmpty(t, got.Name.Pairs)
	require.NotEmpty(t, got.Name.Steps)
	require.Equal(t, "token-pairs", got.Name.Steps[0].Stage)
	require.Greater(t, got.Name.Score, 0.0)
	require.InDelta(t, got.Name.Score, got.Name.Steps[len(got.Name.Steps)-1].Score, 0.0001)

	require.NotEmpty(t, got.Steps)
	require.Equal(t, "base", got.Steps[0].Stage)
	require.InDelta(t, got.FinalScore, got.Steps[len(got.Steps)-1].Score, 0.0001)

	text := got.String()
	require.Contains(t, text, "match")
	require.Contains(t, text, "mohamed")
	require.Contains(t, text, "token-pairs")
	require.Contains(t, text, "name-only")
	require.Contains(t, text, "compared")
	require.NotContains(t, text, "search.ScorePiece")

	raw, err := json.Marshal(got)
	require.NoError(t, err)
	require.Contains(t, string(raw), `"pairs"`)
	require.Contains(t, string(raw), `"token-pairs"`)
	require.Contains(t, string(raw), `"name-only"`)
	t.Log("\n" + text)
}

func TestNameDebug_ExactPrepared(t *testing.T) {
	query := search.Entity[search.Value]{Name: "AEROCARIBBEAN AIRLINES", Type: search.EntityBusiness}.Normalize()
	index := search.Entity[search.Value]{Name: "Aerocaribbean Airlines", Type: search.EntityBusiness}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)
	require.True(t, got.Name.Exact)
	require.InDelta(t, 1.0, got.Name.Score, 0.0001)
	require.Equal(t, "exact-prepared", got.Name.Steps[0].Stage)
	require.NotEmpty(t, got.Name.Pairs)
	for _, pair := range got.Name.Pairs {
		require.Equal(t, pair.Query, pair.Index)
		require.InDelta(t, 1.0, pair.Score, 0.0001)
	}
}

func TestNameDebug_HistoricalAndPrepareTrace(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "Ocean Pioneer (ex-Cape Diamond)",
		Type: search.EntityVessel,
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "Cape Diamond",
		Type: search.EntityVessel,
	}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)
	require.True(t, got.Name.Historical)
	require.False(t, got.Name.Exact)
	require.Equal(t, "historical", got.Name.QueryName)
	require.Equal(t, "primary", got.Name.IndexName)
	require.Equal(t, "Ocean Pioneer (ex-Cape Diamond)", got.Name.Query.Original)
	require.Equal(t, "Cape Diamond", got.Name.Query.Compared)
	require.Equal(t, []string{"cape", "diamond"}, got.Name.Query.Tokens)

	require.Contains(t, nameStages(got.Name.Query), "former-names")
	require.Contains(t, nameStages(got.Name.Query), "extracted")
	require.Equal(t, "cape diamond", nameStageValue(got.Name.Query, "normalize"))
	text := got.String()
	require.Contains(t, text, "historical")
	require.Contains(t, text, "Cape Diamond")
	t.Log("\n" + text)
}

func TestNameDebug_PartyLabelPrepareTrace(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "SHIPPER: Ocean Pioneer Ltd.",
		Type: search.EntityBusiness,
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "Ocean Pioneer Limited",
		Type: search.EntityBusiness,
	}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)
	require.Equal(t, "SHIPPER: Ocean Pioneer Ltd.", got.Name.Query.Original)
	require.Empty(t, got.Name.Query.Compared)

	require.Contains(t, nameStages(got.Name.Query), "party-labels")
	require.Contains(t, nameStages(got.Name.Index), "legal-forms")
	text := got.String()
	require.Contains(t, text, "party-labels")
	require.Contains(t, text, "legal-forms")
	t.Log("\n" + text)
}

func TestNameDebug_UnmatchedTokens(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "John Smith",
		Type: search.EntityPerson,
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "John Adams",
		Type: search.EntityPerson,
	}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)

	var unmatchedQuery, unmatchedIndex string
	for _, pair := range got.Name.Pairs {
		if pair.Query == "" {
			unmatchedIndex = pair.Index
		}
		if pair.Index == "" {
			unmatchedQuery = pair.Query
		}
	}
	require.Equal(t, "smith", unmatchedQuery)
	require.Equal(t, "adams", unmatchedIndex)

	text := got.String()
	require.Contains(t, text, "unmatched query  smith")
	require.Contains(t, text, "unmatched index  adams")
}

func TestNameDebug_RecastPreservesOriginalPrepareType(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "John Smith Limited",
		Type: search.EntityPerson,
		Person: &search.Person{
			Name: "John Smith Limited",
		},
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "John Smith Limited",
		Type: search.EntityBusiness,
		Business: &search.Business{
			Name: "John Smith Limited",
		},
	}.Normalize()

	require.Contains(t, query.PreparedFields.Name, "limited")
	require.NotContains(t, index.PreparedFields.Name, "limited")

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)
	require.Equal(t, query.PreparedFields.Name, got.Name.Query.Prepared)
	require.NotContains(t, nameStages(got.Name.Query), "legal-forms")
	require.Contains(t, nameStages(got.Name.Index), "legal-forms")
}

func TestNameDebug_FewMatchingTermsReasonIsAggregate(t *testing.T) {
	query := search.Entity[search.Value]{
		Name: "John Smith",
		Type: search.EntityPerson,
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "John Smith Alpha Bravo Charlie Delta Echo Foxtrot Golf Hotel",
		Type: search.EntityPerson,
	}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.NotNil(t, got.Name)

	var few *search.ScoreStep
	for i := range got.Name.Steps {
		if got.Name.Steps[i].Stage == "few-matching-terms" {
			few = &got.Name.Steps[i]
			break
		}
	}
	require.NotNil(t, few, "unmatched index tokens should infer too few matching terms from the aggregate score")
	require.Contains(t, few.Reason, "aggregate token-pair score")
	require.NotContains(t, few.Reason, "term-match threshold")
	require.Contains(t, got.String(), "few-matching-terms")
}

func TestNameDebug_IdentifierOnly(t *testing.T) {
	query := search.Entity[search.Value]{
		Type: search.EntityBusiness,
		CryptoAddresses: []search.CryptoAddress{
			{Currency: "XBT", Address: "12VrYZgS1nmf9KHHped24xBb1aLLRpV2cT"},
		},
	}.Normalize()
	index := search.Entity[search.Value]{
		Name: "SUEX OTC, S.R.O.",
		Type: search.EntityBusiness,
		CryptoAddresses: []search.CryptoAddress{
			{Currency: "XBT", Address: "12VrYZgS1nmf9KHHped24xBb1aLLRpV2cT"},
		},
	}.Normalize()

	got := search.DebugSimilarityWithOpts(nil, query, index, search.SimilarityOpts{})
	require.InDelta(t, search.Similarity(query, index), got.FinalScore, 0.0001)
	require.InDelta(t, 1.0, got.FinalScore, 0.0001)
	require.Nil(t, got.Name)
	require.NotEmpty(t, got.Steps)
	require.Equal(t, "exact-override", got.Steps[len(got.Steps)-1].Stage)
	require.NotContains(t, got.String(), "search.ScorePiece")
}

func nameStages(side search.NameSide) string {
	var stages []string
	for _, step := range side.Steps {
		stages = append(stages, step.Stage)
	}
	return strings.Join(stages, ",")
}

func nameStageValue(side search.NameSide, stage string) string {
	for _, step := range side.Steps {
		if step.Stage == stage {
			return step.Value
		}
	}
	return ""
}
