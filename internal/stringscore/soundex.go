package stringscore

// getSoundexDigit returns the Soundex digit for a letter.
// Returns 0 for vowels, H, W, Y (ignored/separator characters).
// Using a switch statement instead of a map for better performance and to avoid allocations.
func getSoundexDigit(r rune) rune {
	switch r {
	case 'B', 'F', 'P', 'V':
		return '1'
	case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
		return '2'
	case 'D', 'T':
		return '3'
	case 'L':
		return '4'
	case 'M', 'N':
		return '5'
	case 'R':
		return '6'
	}
	return 0
}

// EncodeSoundex returns the standard Soundex code for a string.
//
// Standard rules:
//  1. Retain the first letter of the name.
//  2. Map letters to digits per phonetic group.
//  3. H and W are transparent (ignored) but do NOT reset duplicate tracking.
//  4. Vowels (A,E,I,O,U) and Y separate consonants (reset duplicate tracking).
//  5. Remove consecutive duplicate digits.
//  6. Pad or truncate to letter + 3 digits.
//
// Examples:
//   - "Smith"    → "S530"
//   - "Lloyd"    → "L300"  (second L is duplicate of first; O resets; Y like vowel; D=3)
//   - "Miller"   → "M460"  (L=4, second L is dup; E resets; R=6)
//   - "Ashcraft" → "A261"  (S=2; H transparent so C=2 is dup of S; R=6; A resets; F=1)
func EncodeSoundex(s string) string {
	if s == "" {
		return ""
	}

	// Process characters directly to avoid multiple allocations
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			out = append(out, r-32) // to uppercase
		} else if r >= 'A' && r <= 'Z' {
			out = append(out, r)
		}
	}

	if len(out) == 0 {
		return ""
	}

	firstLetter := out[0]

	// Initialize lastDigit from the first letter so the second letter
	// (if same phonetic group) is correctly treated as a duplicate.
	lastDigit := getSoundexDigit(firstLetter)

	var code []rune

	for _, char := range out[1:] {
		// H and W are transparent: skip WITHOUT resetting lastDigit.
		// Consonants on either side of H/W are treated as adjacent.
		if char == 'H' || char == 'W' {
			continue
		}

		digit := getSoundexDigit(char)
		if digit == 0 {
			// Vowels and Y: separate consonants, reset duplicate tracking
			lastDigit = 0
			continue
		}

		// Only append if different from the last digit
		if digit != lastDigit {
			code = append(code, digit)
			lastDigit = digit
		}

		if len(code) >= 3 {
			break
		}
	}

	// Pad with zeros to ensure 3 digits
	for len(code) < 3 {
		code = append(code, '0')
	}

	return string(firstLetter) + string(code[:3])
}

// SoundexMatch returns true if two strings encode to the same Soundex code.
// Two empty strings are considered a match.
func SoundexMatch(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	code1 := EncodeSoundex(s1)
	code2 := EncodeSoundex(s2)
	return code1 == code2 && code1 != ""
}

// SoundexDistance returns 1.0 if Soundex codes match, 0.0 otherwise.
// The name "Distance" reflects that this is a phonetic similarity measure
// in the range [0.0, 1.0].
func SoundexDistance(s1, s2 string) float64 {
	if SoundexMatch(s1, s2) {
		return 1.0
	}
	return 0.0
}

// SoundexScore is an alias for SoundexDistance kept for backwards compatibility.
func SoundexScore(s1, s2 string) float64 {
	return SoundexDistance(s1, s2)
}
