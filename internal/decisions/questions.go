package decisions

import (
	"fmt"

	"github.com/moov-io/watchman/pkg/search"
)

func screeningQuestions(n int) map[string]Question {
	questions := make(map[string]Question, n*3)
	for i := 0; i < n; i++ {
		questions[sameEntityKey(i)] = Question{
			Type: "noul",
			Instructions: fmt.Sprintf(
				"Is candidate %d the same real-world person, business, vessel, or aircraft as the query? Matching unique identity documents (passport, national ID, IMO, MMSI, aircraft serial, crypto address, SWIFT/BIC) are strong evidence. Name-only similarity is not enough for a high probability. Conflicting identifiers should lower the probability.",
				i,
			),
			Criteria: map[string]string{
				"true":  fmt.Sprintf("The query and candidate %d are the same legal or natural person, vessel, or aircraft.", i),
				"false": fmt.Sprintf("They are different parties, including sister ships, related companies, or a weak name resemblance. Candidate %d is not the query.", i),
			},
		}
		questions[actionKey(i)] = Question{
			Type: "choice",
			Instructions: fmt.Sprintf(
				"Which review action should a sanctions screener take for candidate %d against the query?",
				i,
			),
			Criteria: map[string]string{
				search.DecisionActionClearHit:            "Same entity with identity evidence: matching unique IDs, or name plus an agreeing date of birth or registration.",
				search.DecisionActionNeedsReview:         "Plausible match a human should review: similar name, partial identifiers, aliases, or missing fields.",
				search.DecisionActionLikelyFalsePositive: "Different entity: conflicting identifiers, different type, sister ship or related company, or only a weak name resemblance.",
			},
		}
		questions[qualityKey(i)] = Question{
			Type: "score",
			Instructions: fmt.Sprintf(
				"How strong is the identity match between the query and candidate %d?",
				i,
			),
			Criteria: []string{
				"Unrelated",
				"Weak name resemblance only",
				"Plausible, needs corroboration",
				"Strong identity match",
			},
		}
	}
	return questions
}
