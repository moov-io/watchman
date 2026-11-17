package decisions

import "errors"

var (
	ErrBaseURLRequired    = errors.New("decisions: base URL is required")
	ErrModelRequired      = errors.New("decisions: model is required")
	ErrInvalidTimeout     = errors.New("decisions: timeout must be non-negative")
	ErrInvalidMaxEntities = errors.New("decisions: max entities must be non-negative")
	ErrInvalidMinMatch    = errors.New("decisions: min match must be between 0 and 1")
	ErrUnexpectedStatus   = errors.New("decisions: unexpected HTTP status")
	ErrInvalidResponse    = errors.New("decisions: invalid response")
)
