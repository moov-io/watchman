package search

type SearchedEntity[T any] struct {
	Entity[T]

	Match float64 `json:"match"`
}
