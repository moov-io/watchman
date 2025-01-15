package prepare

import (
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

func NormalizeGender(input string) search.Gender {
	v := strings.ToLower(strings.TrimSpace(input))

	switch v {
	case "m", "male", "man", "guy":
		return search.GenderMale

	case "f", "female", "woman", "gal", "girl":
		return search.GenderFemale
	}

	return search.GenderUnknown
}
