package search

// EntityDecision is an optional advisory from a Jev-style decision model
// (Ollama POST /v1/systemone, or a hosted TypeSafe-compatible API).
// It does not change Match.
type EntityDecision struct {
	// SameEntity is the probability (0-1) that this list record is the same
	// real-world entity as the query.
	SameEntity float64 `json:"sameEntity"`

	// Action is one of clear_hit, needs_review, likely_false_positive.
	Action string `json:"action,omitempty"`

	// ActionProbabilities is the full choice distribution for Action.
	ActionProbabilities map[string]float64 `json:"actionProbabilities,omitempty"`

	// MatchQuality is an ordered identity-match score (0 is unrelated).
	MatchQuality float64 `json:"matchQuality,omitempty"`

	// Confidence is the decision model's confidence in Action, when provided.
	Confidence float64 `json:"confidence,omitempty"`

	// Model is the decision model that produced this advisory.
	Model string `json:"model,omitempty"`
}

const (
	DecisionActionClearHit            = "clear_hit"
	DecisionActionNeedsReview         = "needs_review"
	DecisionActionLikelyFalsePositive = "likely_false_positive"
)
