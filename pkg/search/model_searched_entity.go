package search

type SearchedEntity[T any] struct {
	Entity[T]

	Match float64 `json:"match"`

	// Debug is a human-readable rendering of Details: the name comparison,
	// token pairs, and the score trail with deltas. Prefer Details for
	// programmatic use.
	//
	// Adding ?debug=true (or ?debug=0.80) to /v2/search populates this field
	// for returned hits that meet the threshold. Debug is computed after
	// ranking, only for those hits.
	//
	// The format will change over time and should not be parsed by machines.
	Debug string `json:"debug,omitempty"`

	// Details is the typed score breakdown: field pieces, name comparison
	// (query vs index, token pairs, name-score steps), and the final-score trail.
	// RankedBy names the ranker (similarity or embeddings). When embeddings
	// ranked the hit, FinalScore is the field-similarity diagnostic and
	// RankingScore is the cosine used for match.
	//
	// The fields returned may change as the general similarity algorithm and scoring methodologies evolve.
	// There is no API stability guarantee for Details or SimilarityScore.
	Details SimilarityScore `json:"details,omitempty,omitzero"`
}
