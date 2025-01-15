package prepare

import (
	"strings"
)

func NormalizeGender(input string) string {
	v := strings.ToLower(strings.TrimSpace(input))

	// returned values need to match pkg/search.Gender values
	switch v {
	case "m", "male", "man", "guy":
		return "male"
	case "f", "female", "woman", "gal", "girl":
		return "female"
	}
	return "unknown"
}
