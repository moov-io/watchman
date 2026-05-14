package stringscore

import (
	"strings"
	"unicode"
)

// soundexTable maps letters to their Soundex digit.
// 0 = ignored (vowels, H, W, Y)
var soundexTable = map[rune]rune{
	'B': '1', 'F': '1', 'P': '1', 'V': '1',
	'C': '2', 'G': '2', 'J': '2', 'K': '2', 'Q': '2', 'S': '2', 'X': '2', 'Z': '2',
	'D': '3', 'T': '3',
	'L': '4',
	'M': '5', 'N': '5',
	'R': '6',
}

// EncodeSoundex returns the standard Soundex code for a string.
//
// Standard rules:
//  1. Retain the first letter of the name.
//  2. Map letters to digits per phonetic group.
//  3. H and W are transparent (ignored, do NOT break adjacency/duplicate detection).
//  4. Vowels (A,E,I,O,U) and Y separate consonants (reset duplicate tracking).
//  5. Remove consecutive duplicate digits.
//  6. Pad or truncate to letter + 3 digits.
//
// Examples:
//   - "Smith"    → "S530"
//   - "Lloyd"    → "L300"  (second L is duplicate of first; O resets; Y ignored like vowel; D=3)
//   - "Miller"   → "M400"  (L=4, second L is dup; E resets; R=6 but we stop at 3 digits)
//   - "Ashcraft" → "A261"  (S=2; H transparent so C=2 is dup of S, skip; R=6; A resets; F=1; T=3 but stop)
func EncodeSoundex(s string) string {
	if s == "" {
		return ""
	}

	// Uppercase and strip non-letters
	s = strings.ToUpper(s)
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)
	if s == "" {
		return ""
	}

	firstLetter := rune(s[0])

	// Initialize lastDigit to the Soundex digit of the first letter.
	// This ensures the second letter (if same phonetic group) is treated as a duplicate.
	lastDigit, _ := soundexTable[firstLetter] // 0 if first letter is a vowel

	var code []rune

	for i := 1; i < len(s); i++ {
		char := rune(s[i])

		// H and W are transparent: skip them WITHOUT resetting lastDigit.
		// This means consonants on either side of H/W are treated as adjacent.
		// e.g. Ashcraft: S-H-C → S=2, H ignored, C=2 (duplicate of S, skip)
		if char == 'H' || char == 'W' {
			continue
		}

		digit, exists := soundexTable[char]
		if !exists {
			// Vowels and Y: they separate consonants, so reset duplicate tracking.
			lastDigit = 0
			continue
		}

		// Only append if different from the last consonant digit
		if digit != lastDigit {
			code = append(code, digit)
			lastDigit = digit
		}

		if len(code) >= 3 {
			break
		}
	}

	// Pad with zeros
	for len(code) < 3 {
		code = append(code, '0')
	}

	return string(firstLetter) + string(code[:3])
}

// SoundexMatch returns true if two strings encode to the same Soundex code.
// Two empty strings are considered a match.
func SoundexMatch(s1, s2 string) bool {
	code1 := EncodeSoundex(s1)
	code2 := EncodeSoundex(s2)
	// Two empty strings both return "" — treat as equal (both inputs were empty)
	if s1 == "" && s2 == "" {
		return true
	}
	return code1 == code2 && code1 != ""
}

// SoundexScore returns 1.0 if Soundex codes match, 0.0 otherwise.
func SoundexScore(s1, s2 string) float64 {
	if SoundexMatch(s1, s2) {
		return 1.0
	}
	return 0.0
}