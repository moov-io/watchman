//go:build !js

package arch

import (
	"net/url"
)

func PrefillValues() url.Values {
	return make(url.Values)
}
