//go:build !libpostal

package address

import (
	"github.com/moov-io/watchman/pkg/search"
)

func ParseAddress(input string) search.Address {
	var out search.Address

	return out
}
