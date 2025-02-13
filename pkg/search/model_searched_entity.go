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
}
