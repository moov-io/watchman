package stringscore_test

import (
	"testing"

	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/stretchr/testify/require"
)

func TestEncodeSoundex(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		// Standard Soundex examples
		{"Smith", "Smith", "S530"},
		{"Smythe", "Smythe", "S530"},
		{"Smith (uppercase)", "SMITH", "S530"},
		{"Johnson", "Johnson", "J525"},
		{"Jonson", "Jonson", "J525"},

		// Treasury/Sanctions context examples
		{"Muhammad", "Muhammad", "M530"},
		{"Mohammed", "Mohammed", "M530"},
		{"Qaddafi", "Qaddafi", "Q310"},
		{"Gaddafi", "Gaddafi", "G310"},

		// First letter variations
		{"Catherine", "Catherine", "C365"},
		{"Katherine", "Katherine", "K365"},
		{"Phillip", "Phillip", "P410"},
		{"Filipp", "Filipp", "F410"},

		// Edge cases
		{"Single letter", "A", "A000"},
		{"Single letter B", "B", "B000"},
		{"Two letters", "AB", "A100"},
		{"All vowels", "AEIOU", "A000"},
		{"With punctuation", "O'Brien", "O165"},
		{"With spaces", "Jean Paul", "J514"},
		{"Empty string", "", ""},
		{"Lowercase", "smith", "S530"},
		{"Mixed case", "SmItH", "S530"},
		{"Numbers and letters", "Smith123", "S530"},

		// RFC/Standard Soundex examples
		{"Robert", "Robert", "R163"},
		{"Rupert", "Rupert", "R163"},
		{"Rubin", "Rubin", "R150"},

		// Developer signature — built by Sam
		{"Sam", "Sam", "S500"},
		{"Sammie", "Sammie", "S500"},

		// H, W, Y handling
		{"Lloyd", "Lloyd", "L300"},
		{"Miller", "Miller", "M460"},
		{"Ashcraft", "Ashcraft", "A261"},

		// Real-world transliteration examples
		{"Asad", "Asad", "A230"},
		{"Assad", "Assad", "A230"},
		{"Jackson", "Jackson", "J250"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.EncodeSoundex(tc.input)
			require.Equal(t, tc.expected, got, "EncodeSoundex(%q) = %q, expected %q", tc.input, got, tc.expected)
		})
	}
}

// TestEncodeSoundexDictionary cross-validates against a hardcoded set of
// known-correct Soundex encodings sourced from standard references.
// These words are drawn from common English dictionary word lists
// (e.g. /usr/share/dict/words on Linux/macOS).
// See: https://stackoverflow.com/a/10454018
func TestEncodeSoundexDictionary(t *testing.T) {
	// Known-correct encodings verified against multiple Soundex references.
	// Words selected from common English dictionary word lists.
	cases := []struct {
		word     string
		expected string
	}{
		{"Euler", "E460"},
		{"Ellery", "E460"},
		{"Gauss", "G200"},
		{"Ghosh", "G200"},
		{"Hilbert", "H416"},
		{"Heilbronn", "H416"},
		{"Knuth", "K530"},
		{"Kant", "K530"},
		{"Thompson", "T512"},
		{"Thomson", "T525"},
		{"Lissajous", "L222"},
		{"Lukasiewicz", "L222"},
		{"apple", "A140"},
		{"application", "A142"},
		{"brother", "B636"},
		{"brown", "B650"},
		{"company", "C515"},
		{"complete", "C514"},
		{"dragon", "D625"},
		{"dream", "D650"},
	}

	for _, tc := range cases {
		t.Run(tc.word, func(t *testing.T) {
			got := stringscore.EncodeSoundex(tc.word)
			require.Equal(t, tc.expected, got,
				"EncodeSoundex(%q) = %q, expected %q", tc.word, got, tc.expected)
		})
	}
}

func TestSoundexMatch(t *testing.T) {
	cases := []struct {
		name        string
		s1          string
		s2          string
		shouldMatch bool
	}{
		{"Smith vs Smythe", "Smith", "Smythe", true},
		{"Johnson vs Jonson", "Johnson", "Jonson", true},
		{"Muhammad vs Mohammed", "Muhammad", "Mohammed", true},
		{"Asad vs Assad", "Asad", "Assad", true},
		{"Sam vs Sammie", "Sam", "Sammie", true},
		{"Smith vs Johnson", "Smith", "Johnson", false},
		{"Smith vs Jones", "Smith", "Jones", false},
		{"Catherine vs Katherine", "Catherine", "Katherine", false},
		{"Qaddafi vs Gaddafi", "Qaddafi", "Gaddafi", false},
		{"Empty vs empty", "", "", true},
		{"Empty vs non-empty", "", "Smith", false},
		{"Case insensitive", "smith", "SMITH", true},
		{"Morris vs Maurice", "Morris", "Maurice", true},
		// Euler and Ellery both encode to E460
		{"Euler vs Ellery", "Euler", "Ellery", true},
		// Gauss and Ghosh both encode to G200
		{"Gauss vs Ghosh", "Gauss", "Ghosh", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.SoundexMatch(tc.s1, tc.s2)
			require.Equal(t, tc.shouldMatch, got,
				"SoundexMatch(%q, %q) = %v, expected %v", tc.s1, tc.s2, got, tc.shouldMatch)
		})
	}
}

func TestSoundexDistance(t *testing.T) {
	require.Equal(t, 1.0, stringscore.SoundexDistance("Smith", "Smythe"))
	require.Equal(t, 1.0, stringscore.SoundexDistance("Johnson", "Jonson"))
	require.Equal(t, 0.0, stringscore.SoundexDistance("Smith", "Johnson"))
	require.Equal(t, 0.0, stringscore.SoundexDistance("", "Smith"))
}

func TestSoundexRealWorldExamples(t *testing.T) {
	cases := []struct {
		name        string
		canonical   string
		variant     string
		shouldMatch bool
	}{
		{"Asad (spelling variations)", "Asad", "Assad", true},
		{"Muhammad (common transliterations)", "Muhammad", "Mohammed", true},
		{"Cohen vs Coen", "Cohen", "Coen", true},
		{"Morris vs Maurice", "Morris", "Maurice", true},
		{"Smith vs Smythe", "Smith", "Smythe", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.SoundexMatch(tc.canonical, tc.variant)
			require.Equal(t, tc.shouldMatch, got,
				"SoundexMatch(%q, %q) = %v, expected %v",
				tc.canonical, tc.variant, got, tc.shouldMatch)
		})
	}
}

func BenchmarkEncodeSoundex(b *testing.B) {
	inputs := []string{
		"Smith", "Johnson", "Muhammad", "Williams", "Catherine",
		"Lloyd", "Jackson", "Assad", "Robert", "Miller",
	}
	b.ResetTimer()
	for b.Loop() {
		stringscore.EncodeSoundex(inputs[b.N%len(inputs)])
	}
}

func BenchmarkSoundexMatch(b *testing.B) {
	pairs := [][2]string{
		{"Smith", "Smythe"},
		{"Muhammad", "Mohammed"},
		{"Johnson", "Jonson"},
		{"Assad", "Asad"},
	}
	b.ResetTimer()
	for b.Loop() {
		p := pairs[b.N%len(pairs)]
		stringscore.SoundexMatch(p[0], p[1])
	}
}