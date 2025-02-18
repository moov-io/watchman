package norm

import (
	"strings"
	"unicode"
)

// Name performs thorough name normalization
func Name(name string) string {
	// Convert to lowercase and trim spaces
	name = strings.ToLower(strings.TrimSpace(name))

	// Remove all punctuation and normalize whitespace
	var normalized strings.Builder
	lastWasSpace := true // Start with true to trim leading spaces

	for _, r := range name {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			if !lastWasSpace {
				normalized.WriteRune(' ')
				lastWasSpace = true
			}
			continue
		}

		if unicode.IsSpace(r) {
			if !lastWasSpace {
				normalized.WriteRune(' ')
				lastWasSpace = true
			}
			continue
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			normalized.WriteRune(r)
			lastWasSpace = false
		}
	}

	return strings.TrimSpace(normalized.String())
}

var (
	// Minimum length for a name term to be considered significant
	minTermLength = 3

	noiseTerms = []string{"the", "and", "or", "of", "in", "at", "by"}
)

func RemoveInsignificantTerms(terms []string) []string {
	out := make([]string, 0, len(terms))
	for t := range terms {
		var found bool
		for n := range noiseTerms {
			if strings.EqualFold(terms[t], noiseTerms[n]) {
				found = true
				break
			}
		}
		if !found {
			if len(terms[t]) >= minTermLength {
				out = append(out, terms[t])
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
