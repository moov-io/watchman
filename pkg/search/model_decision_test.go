package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchedEntity_DecisionJSON(t *testing.T) {
	t.Run("omitted when nil", func(t *testing.T) {
		raw, err := json.Marshal(SearchedEntity[Value]{
			Entity: Entity[Value]{Name: "Ada"},
			Match:  0.8,
		})
		require.NoError(t, err)
		require.NotContains(t, string(raw), "decision")
	})

	t.Run("round trip", func(t *testing.T) {
		in := SearchedEntity[Value]{
			Entity: Entity[Value]{Name: "Ada"},
			Match:  0.8,
			Decision: &EntityDecision{
				SameEntity:   0.91,
				Action:       DecisionActionNeedsReview,
				MatchQuality: 2.1,
				Model:        "nimble",
			},
		}
		raw, err := json.Marshal(in)
		require.NoError(t, err)
		require.Contains(t, string(raw), `"decision"`)
		require.Contains(t, string(raw), `"sameEntity":0.91`)

		var out SearchedEntity[Value]
		require.NoError(t, json.Unmarshal(raw, &out))
		require.NotNil(t, out.Decision)
		require.Equal(t, 0.91, out.Decision.SameEntity)
		require.Equal(t, DecisionActionNeedsReview, out.Decision.Action)
		require.Equal(t, "nimble", out.Decision.Model)
	})

	t.Run("zero quality and confidence stay in JSON", func(t *testing.T) {
		raw, err := json.Marshal(EntityDecision{
			MatchQuality: 0,
			Confidence:   0,
		})
		require.NoError(t, err)
		require.Contains(t, string(raw), `"matchQuality":0`)
		require.Contains(t, string(raw), `"confidence":0`)
	})
}
