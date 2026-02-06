package embeddings

import (
	"unicode"
)

// Script represents a Unicode script/writing system.
type Script string

const (
	ScriptLatin    Script = "Latin"
	ScriptArabic   Script = "Arabic"
	ScriptCyrillic Script = "Cyrillic"
	ScriptHan      Script = "Han"      // Chinese
	ScriptHangul   Script = "Hangul"   // Korean
	ScriptHiragana Script = "Hiragana" // Japanese
	ScriptKatakana Script = "Katakana" // Japanese
	ScriptThai     Script = "Thai"
	ScriptHebrew   Script = "Hebrew"
	ScriptGreek    Script = "Greek"
	ScriptUnknown  Script = "Unknown"
)

// IsNonLatin returns true if the text contains primarily non-Latin characters.
// This is used to determine whether to use embedding-based search (for cross-script)
// or Jaro-Winkler (for Latin-only queries).
func IsNonLatin(text string) bool {
	var latinCount, nonLatinCount int

	for _, r := range text {
		if unicode.IsLetter(r) {
			if unicode.Is(unicode.Latin, r) {
				latinCount++
			} else {
				nonLatinCount++
			}
		}
	}

	total := latinCount + nonLatinCount
	if total == 0 {
		return false // No letters, treat as Latin
	}

	// Consider non-Latin if >30% of letters are non-Latin
	// This threshold catches mixed names like "محمد Ali"
	return float64(nonLatinCount)/float64(total) > 0.3
}

// DetectScript returns the primary Unicode script of the text.
// If the text contains multiple scripts, returns the most common one.
func DetectScript(text string) Script {
	counts := map[Script]int{
		ScriptLatin:    0,
		ScriptArabic:   0,
		ScriptCyrillic: 0,
		ScriptHan:      0,
		ScriptHangul:   0,
		ScriptHiragana: 0,
		ScriptKatakana: 0,
		ScriptThai:     0,
		ScriptHebrew:   0,
		ScriptGreek:    0,
	}

	for _, r := range text {
		if !unicode.IsLetter(r) {
			continue
		}

		switch {
		case unicode.Is(unicode.Latin, r):
			counts[ScriptLatin]++
		case unicode.Is(unicode.Arabic, r):
			counts[ScriptArabic]++
		case unicode.Is(unicode.Cyrillic, r):
			counts[ScriptCyrillic]++
		case unicode.Is(unicode.Han, r):
			counts[ScriptHan]++
		case unicode.Is(unicode.Hangul, r):
			counts[ScriptHangul]++
		case unicode.Is(unicode.Hiragana, r):
			counts[ScriptHiragana]++
		case unicode.Is(unicode.Katakana, r):
			counts[ScriptKatakana]++
		case unicode.Is(unicode.Thai, r):
			counts[ScriptThai]++
		case unicode.Is(unicode.Hebrew, r):
			counts[ScriptHebrew]++
		case unicode.Is(unicode.Greek, r):
			counts[ScriptGreek]++
		}
	}

	// Find the dominant script
	maxScript := ScriptUnknown
	maxCount := 0

	for script, count := range counts {
		if count > maxCount {
			maxScript = script
			maxCount = count
		}
	}

	return maxScript
}

// ContainsScript returns true if the text contains any characters from the specified script.
func ContainsScript(text string, script Script) bool {
	var rangeTable *unicode.RangeTable

	switch script {
	case ScriptLatin:
		rangeTable = unicode.Latin
	case ScriptArabic:
		rangeTable = unicode.Arabic
	case ScriptCyrillic:
		rangeTable = unicode.Cyrillic
	case ScriptHan:
		rangeTable = unicode.Han
	case ScriptHangul:
		rangeTable = unicode.Hangul
	case ScriptHiragana:
		rangeTable = unicode.Hiragana
	case ScriptKatakana:
		rangeTable = unicode.Katakana
	case ScriptThai:
		rangeTable = unicode.Thai
	case ScriptHebrew:
		rangeTable = unicode.Hebrew
	case ScriptGreek:
		rangeTable = unicode.Greek
	default:
		return false
	}

	for _, r := range text {
		if unicode.Is(rangeTable, r) {
			return true
		}
	}

	return false
}

// IsCrossScript returns true if the text contains characters from multiple scripts.
// This is useful for detecting transliterated or mixed-script names.
func IsCrossScript(text string) bool {
	scripts := make(map[Script]bool)

	for _, r := range text {
		if !unicode.IsLetter(r) {
			continue
		}

		switch {
		case unicode.Is(unicode.Latin, r):
			scripts[ScriptLatin] = true
		case unicode.Is(unicode.Arabic, r):
			scripts[ScriptArabic] = true
		case unicode.Is(unicode.Cyrillic, r):
			scripts[ScriptCyrillic] = true
		case unicode.Is(unicode.Han, r):
			scripts[ScriptHan] = true
		case unicode.Is(unicode.Hangul, r):
			scripts[ScriptHangul] = true
		case unicode.Is(unicode.Hiragana, r):
			scripts[ScriptHiragana] = true
		case unicode.Is(unicode.Katakana, r):
			scripts[ScriptKatakana] = true
		case unicode.Is(unicode.Thai, r):
			scripts[ScriptThai] = true
		case unicode.Is(unicode.Hebrew, r):
			scripts[ScriptHebrew] = true
		case unicode.Is(unicode.Greek, r):
			scripts[ScriptGreek] = true
		}
	}

	return len(scripts) > 1
}
