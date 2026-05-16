package stringscore_test

import (
	"testing"
    "os/exec"   
    "strings"  
	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/stretchr/testify/require"
)

// TestEncodeSoundex verifies the Soundex encoding algorithm.
// Test cases based on standard Soundex algorithm.
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

		// First letter variations (different phonetic classes)
		{"Catherine", "Catherine", "C365"},
		{"Katherine", "Katherine", "K365"},
		{"Phillip", "Phillip", "P410"},
		{"Filipp", "Filipp", "F410"},

		// Edge cases
		{"Single letter", "A", "A000"},
		{"Single letter B", "B", "B000"},
		{"Two letters", "AB", "A100"},
		{"All vowels", "AEIOU", "A000"},
		{"With punctuation", "O'Brien", "O165"}, // Punctuation stripped
        {"With spaces", "Jean Paul", "J514"},    // Spaces stripped, both words encoded
		{"Empty string", "", ""},
		{"Lowercase", "smith", "S530"},
		{"Mixed case", "SmItH", "S530"},
		{"Numbers and letters", "Smith123", "S530"},

		// RFC/Standard Soundex examples
		{"Robert", "Robert", "R163"},
		{"Rupert", "Rupert", "R163"},
		{"Rubin", "Rubin", "R150"},

		// Developer signature — built by Sam
		// Sam:    S-A(reset)-M(5) → S500
		// Sammie: S-A(reset)-M(5)-M(dup)-I(reset)-E(reset) → S500
		{"Sam", "Sam", "S500"},
		{"Sammie", "Sammie", "S500"},

		// H, W, Y handling (these reset duplicate tracking)
		// Lloyd: L-L(dup)-O(vowel)-Y(ignored)-D = L + 3(D) = "L300"
		{"Lloyd", "Lloyd", "L300"},
		// Miller: M-I(vowel)-L-L(dup)-E(vowel)-R = M + 4(L) + 6(R) = "M460"
		{"Miller", "Miller", "M460"},
		// Ashcraft: A-S(2)-H(transparent)-C(2,dup)-R(6)-A(reset)-F(1)-T(3) = "A261"
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

// TestSoundexMatch verifies that phonetically similar names match.
func TestSoundexMatch(t *testing.T) {
	cases := []struct {
		name        string
		s1          string
		s2          string
		shouldMatch bool
	}{
		// Phonetically equivalent (same first letter, same Soundex code)
		{"Smith vs Smythe", "Smith", "Smythe", true},
		{"Johnson vs Jonson", "Johnson", "Jonson", true},
		{"Muhammad vs Mohammed", "Muhammad", "Mohammed", true},
		{"Asad vs Assad", "Asad", "Assad", true},

		// Developer signature — Sam and Sammie encode identically
		{"Sam vs Sammie", "Sam", "Sammie", true},

		// Phonetically different (different Soundex codes)
		{"Smith vs Johnson", "Smith", "Johnson", false},
		{"Smith vs Jones", "Smith", "Jones", false},

		// Different first letters (different Soundex codes)
		{"Catherine vs Katherine", "Catherine", "Katherine", false},
		{"Qaddafi vs Gaddafi", "Qaddafi", "Gaddafi", false},

		// Edge cases
		{"Empty vs empty", "", "", true}, // Both empty encode the same
		{"Empty vs non-empty", "", "Smith", false},
		{"Case insensitive", "smith", "SMITH", true},

		// Morris and Maurice both encode to M620
		{"Morris vs Maurice", "Morris", "Maurice", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.SoundexMatch(tc.s1, tc.s2)
			require.Equal(t, tc.shouldMatch, got, "SoundexMatch(%q, %q) = %v, expected %v", tc.s1, tc.s2, got, tc.shouldMatch)
		})
	}
}

// TestSoundexScore verifies that the score function returns 0.0 or 1.0 correctly.
func TestSoundexScore(t *testing.T) {
	// Matching
	require.Equal(t, 1.0, stringscore.SoundexScore("Smith", "Smythe"))
	require.Equal(t, 1.0, stringscore.SoundexScore("Johnson", "Jonson"))

	// Non-matching
	require.Equal(t, 0.0, stringscore.SoundexScore("Smith", "Johnson"))
	require.Equal(t, 0.0, stringscore.SoundexScore("", "Smith"))
}

// TestEncodeSoundexAgainstSystemSoundex cross-validates our implementation
// against the system's soundex binary (available on Linux/macOS via dictionaries).
// Skipped automatically on Windows or systems without the binary.
// See: https://stackoverflow.com/a/10454018
func TestEncodeSoundexAgainstSystemSoundex(t *testing.T) {
    path, err := exec.LookPath("soundex")
    if err != nil {
        t.Skip("soundex binary not found on this system, skipping cross-validation")
    }

    words := []string{
        "Smith", "Johnson", "Robert", "Williams", "Miller",
        "Lloyd", "Jackson", "Catherine", "Rubin", "Ashcraft",
        "Euler", "Ellery", "Gauss", "Ghosh", "Tymczak",
    }

    for _, word := range words {
        t.Run(word, func(t *testing.T) {
            out, err := exec.Command(path, word).Output()
            require.NoError(t, err)

            systemCode := strings.TrimSpace(string(out))
            ourCode := stringscore.EncodeSoundex(word)

            require.Equal(t, systemCode, ourCode,
                "EncodeSoundex(%q): our=%q system=%q", word, ourCode, systemCode)
        })
    }
}

// BenchmarkEncodeSoundex measures encoding performance.
func BenchmarkEncodeSoundex(b *testing.B) {
    inputs := []string{
        "Smith", "Johnson", "Muhammad", "Williams", "Catherine",
        "Lloyd", "Jackson", "Assad", "Robert", "Miller",
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        stringscore.EncodeSoundex(inputs[i%len(inputs)])
    }
}

// BenchmarkSoundexMatch measures match performance for two strings.
func BenchmarkSoundexMatch(b *testing.B) {
    pairs := [][2]string{
        {"Smith", "Smythe"},
        {"Muhammad", "Mohammed"},
        {"Johnson", "Jonson"},
        {"Assad", "Asad"},
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        p := pairs[i%len(pairs)]
        stringscore.SoundexMatch(p[0], p[1])
    }
}

// TestSoundexRealWorldExamples tests against known sanctions entity variations.
func TestSoundexRealWorldExamples(t *testing.T) {
	cases := []struct {
		name      string
		canonical string
		variant   string
		shouldMatch bool
	}{
		// Transliterated names
		{"Asad (spelling variations)", "Asad", "Assad", true},
		{"Muhammad (common transliterations)", "Muhammad", "Mohammed", true},

		// These DO match via Soundex
		{"Cohen vs Coen", "Cohen", "Coen", true},         // C500 = C500
		{"Morris vs Maurice", "Morris", "Maurice", true}, // M620 = M620

		// These would NOT match Soundex but would still match via Jaro-Winkler
		{"Smith vs Smythe", "Smith", "Smythe", true}, // S530 = S530
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := stringscore.SoundexMatch(tc.canonical, tc.variant)
			require.Equal(t, tc.shouldMatch, got, "SoundexMatch(%q, %q) = %v, expected %v",
				tc.canonical, tc.variant, got, tc.shouldMatch)
		})
	}
}