package decisions

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func clearDecisionEnv(t *testing.T) {
	t.Helper()
	t.Setenv("DECISIONS_ENABLED", "")
	t.Setenv("DECISIONS_API_KEY", "")
	t.Setenv("DECISIONS_BASE_URL", "")
	t.Setenv("DECISIONS_MODEL", "")
	t.Setenv("DECISIONS_TIMEOUT", "")
	t.Setenv("DECISIONS_MAX_ENTITIES", "")
	t.Setenv("DECISIONS_MIN_MATCH", "")
}

func TestNewService_Disabled(t *testing.T) {
	clearDecisionEnv(t)

	svc, err := NewService(log.NewTestLogger(), Config{Enabled: false})
	require.NoError(t, err)
	require.Nil(t, svc)
}

func TestNewService_Invalid(t *testing.T) {
	clearDecisionEnv(t)

	_, err := NewService(log.NewTestLogger(), Config{Enabled: true, Model: "nimble"})
	require.ErrorIs(t, err, ErrBaseURLRequired)
}

func TestAnnotate_AttachesDecisions(t *testing.T) {
	clearDecisionEnv(t)

	var got Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/v1/systemone", r.URL.Path)
		require.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(body, &got))

		resp := Response{
			Model: "nimble",
			Answers: map[string]Answer{
				"same_entity_0": {Type: "noul", Noul: 0.97},
				"action_0": {
					Type:          "choice",
					Choice:        search.DecisionActionClearHit,
					Confidence:    0.88,
					Probabilities: map[string]float64{search.DecisionActionClearHit: 0.91, search.DecisionActionNeedsReview: 0.07, search.DecisionActionLikelyFalsePositive: 0.02},
				},
				"quality_0":     {Type: "score", Score: 2.8},
				"same_entity_1": {Type: "noul", Noul: 0.12},
				"action_1": {
					Type:   "choice",
					Choice: search.DecisionActionLikelyFalsePositive,
				},
				"quality_1": {Type: "score", Score: 0.4},
			},
			Usage: Usage{InputTokens: 200, OutputTokens: 6},
		}
		require.NoError(t, json.NewEncoder(w).Encode(resp))
	}))
	t.Cleanup(server.Close)

	svc, err := NewService(log.NewTestLogger(), Config{
		Enabled: true,
		BaseURL: server.URL,
		APIKey:  "test-key",
		Model:   "nimble",
		Timeout: time.Second,
	})
	require.NoError(t, err)
	require.NotNil(t, svc)

	born := time.Date(1993, 4, 17, 0, 0, 0, 0, time.UTC)
	query := search.Entity[search.Value]{
		Name: "Dmitry Khoroshev",
		Type: search.EntityPerson,
		Person: &search.Person{
			Name:      "Dmitry Khoroshev",
			BirthDate: &born,
			GovernmentIDs: []search.GovernmentID{{
				Type:       search.GovernmentIDPassport,
				Country:    "RU",
				Identifier: "2018278055",
			}},
		},
	}
	results := []search.SearchedEntity[search.Value]{
		{
			Entity: search.Entity[search.Value]{
				Name:     "Dmitry Yuryevich KHOROSHEV",
				Type:     search.EntityPerson,
				Source:   search.SourceUSOFAC,
				SourceID: "48603",
				Person:   &search.Person{Name: "Dmitry Yuryevich KHOROSHEV"},
			},
			Match: 1.0,
		},
		{
			Entity: search.Entity[search.Value]{
				Name:     "Someone Else",
				Type:     search.EntityPerson,
				Source:   search.SourceUSOFAC,
				SourceID: "1",
			},
			Match: 0.61,
		},
	}

	svc.Annotate(t.Context(), query, results)

	require.Equal(t, "nimble", got.Model)
	require.Len(t, got.Questions, 6)

	state, err := json.Marshal(got.State)
	require.NoError(t, err)
	require.Contains(t, string(state), `"name":"Dmitry Khoroshev"`)
	require.Contains(t, string(state), `"sourceID":"48603"`)
	require.NotContains(t, string(state), "sourceData")

	require.NotNil(t, results[0].Decision)
	require.Equal(t, 0.97, results[0].Decision.SameEntity)
	require.Equal(t, search.DecisionActionClearHit, results[0].Decision.Action)
	require.Equal(t, 2.8, results[0].Decision.MatchQuality)
	require.Equal(t, "nimble", results[0].Decision.Model)

	require.NotNil(t, results[1].Decision)
	require.Equal(t, 0.12, results[1].Decision.SameEntity)
	require.Equal(t, search.DecisionActionLikelyFalsePositive, results[1].Decision.Action)
}

func TestAnnotate_FailOpen(t *testing.T) {
	clearDecisionEnv(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	t.Cleanup(server.Close)

	svc, err := NewService(log.NewTestLogger(), Config{
		Enabled: true,
		BaseURL: server.URL,
		Model:   "nimble",
		Timeout: time.Second,
	})
	require.NoError(t, err)

	results := []search.SearchedEntity[search.Value]{
		{Entity: search.Entity[search.Value]{Name: "Ada"}, Match: 0.9},
	}
	svc.Annotate(t.Context(), search.Entity[search.Value]{Name: "Ada"}, results)
	require.Nil(t, results[0].Decision)
}

func TestAnnotate_RespectsMaxEntitiesAndMinMatch(t *testing.T) {
	clearDecisionEnv(t)

	var n int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		n = len(req.Questions) / questionsPerHit
		require.NoError(t, json.NewEncoder(w).Encode(Response{
			Model: "nimble",
			Answers: map[string]Answer{
				"same_entity_0": {Type: "noul", Noul: 0.8},
			},
		}))
	}))
	t.Cleanup(server.Close)

	svc, err := NewService(log.NewTestLogger(), Config{
		Enabled:     true,
		BaseURL:     server.URL,
		Model:       "nimble",
		Timeout:     time.Second,
		MaxEntities: 1,
		MinMatch:    0.7,
	})
	require.NoError(t, err)

	results := []search.SearchedEntity[search.Value]{
		{Entity: search.Entity[search.Value]{Name: "weak"}, Match: 0.2},
		{Entity: search.Entity[search.Value]{Name: "strong"}, Match: 0.9},
		{Entity: search.Entity[search.Value]{Name: "also-strong"}, Match: 0.85},
	}
	svc.Annotate(t.Context(), search.Entity[search.Value]{Name: "q"}, results)

	require.Equal(t, 1, n)
	require.Nil(t, results[0].Decision)
	require.NotNil(t, results[1].Decision)
	require.Nil(t, results[2].Decision)
}

func TestAnnotate_EmptyResults(t *testing.T) {
	clearDecisionEnv(t)

	called := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	t.Cleanup(server.Close)

	svc, err := NewService(log.NewTestLogger(), Config{
		Enabled: true,
		BaseURL: server.URL,
		Model:   "nimble",
		Timeout: time.Second,
	})
	require.NoError(t, err)

	svc.Annotate(t.Context(), search.Entity[search.Value]{Name: "q"}, nil)
	require.False(t, called)
}

func TestCompactEntity_OmitsSourceData(t *testing.T) {
	entity := search.Entity[search.Value]{
		Name:       "Example Co",
		Type:       search.EntityBusiness,
		Source:     search.SourceUSOFAC,
		SourceID:   "1",
		SourceData: map[string]string{"remarks": "secret-row"},
		Business: &search.Business{
			Name:     "Example Co",
			AltNames: []string{"Example"},
		},
		Addresses: []search.Address{{Line1: "1 Main", City: "Town", Country: "US"}},
	}

	got := compactEntity(entity, 0, 0.81)
	require.Equal(t, "Example Co", got.Name)
	require.Equal(t, 0, got.Index)
	require.Equal(t, 0.81, got.Match)
	require.Equal(t, []string{"Example"}, got.AltNames)
	require.Contains(t, got.Addresses[0], "1 Main")

	raw, err := json.Marshal(got)
	require.NoError(t, err)
	require.NotContains(t, string(raw), "secret-row")
	require.NotContains(t, string(raw), "sourceData")
}

func TestApplyAnswers_Missing(t *testing.T) {
	require.Nil(t, applyAnswers(nil, 0, "nimble"))
	require.Nil(t, applyAnswers(map[string]Answer{"same_entity_1": {Noul: 0.9}}, 0, "nimble"))
}

func TestApplyAnswers_KeepsZeroActionConfidence(t *testing.T) {
	got := applyAnswers(map[string]Answer{
		"same_entity_0": {Type: "noul", Noul: 0.5},
		"action_0": {
			Type:       "choice",
			Choice:     search.DecisionActionNeedsReview,
			Confidence: 0,
			Probabilities: map[string]float64{
				search.DecisionActionClearHit:            1.0 / 3,
				search.DecisionActionNeedsReview:         1.0 / 3,
				search.DecisionActionLikelyFalsePositive: 1.0 / 3,
			},
		},
		"quality_0": {Type: "score", Score: 1.2, Confidence: 0.9},
	}, 0, "nimble")
	require.NotNil(t, got)
	require.Equal(t, 0.0, got.Confidence)
	require.Equal(t, 1.2, got.MatchQuality)
	require.Equal(t, search.DecisionActionNeedsReview, got.Action)
}

func TestApplyAnswers_QualityOnly(t *testing.T) {
	got := applyAnswers(map[string]Answer{
		"quality_0": {Type: "score", Score: 0, Confidence: 0.4},
	}, 0, "nimble")
	require.NotNil(t, got)
	require.Equal(t, 0.0, got.MatchQuality)
	require.Equal(t, 0.4, got.Confidence)
	require.Empty(t, got.Action)
}
