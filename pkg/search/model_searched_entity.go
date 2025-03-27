package search

type SearchedEntity[T any] struct {
	Entity[T]

	Match float64 `json:"match"`

	// Debug is an optional base64 encoded humand-readable field that contains
	// detailed field-level match scores and weight adjustments.
	//
	// Adding ?debug to /v2/search will populated this field, but more memory
	// will be used for each request.
	//
	// The format will change over time and should not be parsed by machines.
	Debug string `json:"debug,omitempty"`

	// Details includes field level matching and scoring results
	//
	// The fields returned may change as the general similarity algorithm and scoring methodologies evolve.
	// There is no API stability guarantee for Details or SimilarityScore.
	Details SimilarityScore `json:"details,omitempty,omitzero"`
}
