package embeddings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNonLatin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Latin-only texts
		{"pure latin", "Mohamed Ali", false},
		{"latin with numbers", "John Smith 123", false},
		{"latin with punctuation", "O'Connor, Jr.", false},

		// Arabic texts
		{"pure arabic", "محمد علي", true},
		{"arabic name", "أحمد", true},

		// Cyrillic texts
		{"pure cyrillic", "Владимир Путин", true},
		{"cyrillic name", "Иванов", true},

		// Chinese texts
		{"chinese", "金正恩", true},
		{"chinese name", "李明", true},

		// Mixed scripts (should be non-Latin if >30% non-Latin)
		{"mixed arabic-latin mostly latin", "Mohamed محمد", true},     // ~50% Arabic
		{"mixed cyrillic-latin mostly latin", "Vladimir Путин", true}, // ~50% Cyrillic

		// Edge cases
		{"empty string", "", false},
		{"only numbers", "12345", false},
		{"only punctuation", "...", false},
		{"korean", "김정은", true},
		{"japanese hiragana", "やまだ", true},
		{"thai", "สมชาย", true},
		{"hebrew", "משה", true},
		{"greek", "Αλέξανδρος", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsNonLatin(tc.input)
			require.Equal(t, tc.expected, result, "IsNonLatin(%q) = %v, want %v", tc.input, result, tc.expected)
		})
	}
}

func TestDetectScript(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Script
	}{
		{"latin name", "Mohamed Ali", ScriptLatin},
		{"arabic name", "محمد علي", ScriptArabic},
		{"cyrillic name", "Владимир Путин", ScriptCyrillic},
		{"chinese name", "金正恩", ScriptHan},
		{"korean name", "김정은", ScriptHangul},
		{"japanese hiragana", "やまだ", ScriptHiragana},
		{"thai name", "สมชาย", ScriptThai},
		{"hebrew name", "משה", ScriptHebrew},
		{"greek name", "Αλέξανδρος", ScriptGreek},
		{"empty string", "", ScriptUnknown},
		{"numbers only", "12345", ScriptUnknown},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := DetectScript(tc.input)
			require.Equal(t, tc.expected, result, "DetectScript(%q) = %v, want %v", tc.input, result, tc.expected)
		})
	}
}

func TestContainsScript(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		script   Script
		expected bool
	}{
		{"latin in latin text", "Mohamed Ali", ScriptLatin, true},
		{"arabic in latin text", "Mohamed Ali", ScriptArabic, false},
		{"arabic in arabic text", "محمد علي", ScriptArabic, true},
		{"latin in arabic text", "محمد علي", ScriptLatin, false},
		{"cyrillic in cyrillic text", "Владимир Путин", ScriptCyrillic, true},
		{"han in chinese text", "金正恩", ScriptHan, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsScript(tc.input, tc.script)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestIsCrossScript(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"single script latin", "Mohamed Ali", false},
		{"single script arabic", "محمد علي", false},
		{"mixed latin-arabic", "Mohamed محمد", true},
		{"mixed latin-cyrillic", "Vladimir Путин", true},
		{"empty string", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsCrossScript(tc.input)
			require.Equal(t, tc.expected, result, "IsCrossScript(%q) = %v, want %v", tc.input, result, tc.expected)
		})
	}
}
