package api

import (
	"net/http"
	"strings"
)

type EntityFormat string

var (
	EntityWatchman EntityFormat = "watchman"
	EntitySenzing  EntityFormat = "senzing"
)

func ChooseEntityFormat(headers http.Header, queryParam string) (EntityFormat, string) {
	// Look for senzing
	format, sub := findSenzingFormat(headers.Get("Accept"))
	if format != "" {
		return format, sub
	}
	format, sub = findSenzingFormat(queryParam)
	if format != "" {
		return format, sub
	}

	// Default
	return EntityWatchman, "json"
}

func findSenzingFormat(input string) (EntityFormat, string) {
	parts := strings.Fields(input)

	for _, part := range parts {
		remainder, found := strings.CutPrefix(part, "senzing")
		if found {
			return EntitySenzing, strings.TrimPrefix(remainder, "/")
		}
	}

	return EntityFormat(""), ""
}
