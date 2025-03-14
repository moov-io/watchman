//go:build wasm

package arch

import (
	"net/url"
	"strings"
	"syscall/js"
)

func PrefillValues() url.Values {
	// In wasm we can access window.location.search
	if js.Global().Get("window").Truthy() {
		// Get the query string from window.location.search
		query := js.Global().Get("window").Get("location").Get("search").String()
		if query != "" {
			// Parse the query string (remove leading "?" if present)
			parsed, _ := url.ParseQuery(strings.TrimPrefix(query, "?"))

			return parsed
		}
	}

	return make(url.Values)
}
