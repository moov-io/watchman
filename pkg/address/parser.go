package address

import (
	"context"

	"github.com/moov-io/watchman/pkg/search"
)

// Parser turns a free-form address string into structured fields.
// postalpool.Service and the optional deepparse client implement this.
type Parser interface {
	ParseAddress(ctx context.Context, input string) (search.Address, error)
}

// Parse uses parser when it is non-nil and succeeds, otherwise the default
// ParseAddress (usaddress, or libpostal when built with that tag).
func Parse(ctx context.Context, parser Parser, input string) search.Address {
	if parser != nil {
		addr, err := parser.ParseAddress(ctx, input)
		if err == nil {
			return addr
		}
	}
	return ParseAddress(ctx, input)
}
