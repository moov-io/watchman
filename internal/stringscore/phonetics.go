package stringscore

import (
	"unicode"
)

var soundexMap = map[rune]rune{
	'A': 'A', 'E': 'A', 'I': 'A', 'O': 'A', 'U': 'A', 'Y': 'A', // vowels
	'B': 'B', 'F': 'B', 'P': 'B', 'V': 'B', // similar sounds
	'C': 'C', 'G': 'C', 'J': 'C', 'K': 'C', 'Q': 'C', 'S': 'C', 'X': 'C', 'Z': 'C', // sibilants
	'D': 'D', 'T': 'D', // dental sounds
	'L': 'L',           // liquids
	'M': 'M', 'N': 'M', // nasal sounds
	'R': 'R',           // trills
	'H': 'H', 'W': 'H', // breathy sounds
}

// getPhoneticClass returns the phonetic class of the first letter in a string
func getPhoneticClass(s string) rune {
	if s == "" {
		return ' '
	}
	// Return the first rune mapped with partial soundex
	for _, r := range s {
		firstLetter := unicode.ToUpper(r)
		if phonetic, ok := soundexMap[firstLetter]; ok {
			return phonetic
		}
		return firstLetter
	}
	return ' '
}

func firstCharacterSoundexMatch(s1, s2 string) bool {
	if s1 == "" || s2 == "" {
		return false
	}
	return getPhoneticClass(s1) == getPhoneticClass(s2)
}
