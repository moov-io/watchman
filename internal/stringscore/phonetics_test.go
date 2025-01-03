package stringscore

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstCharacterSoundexMatch(t *testing.T) {
	require.True(t, firstCharacterSoundexMatch("a", "A"))
	require.True(t, firstCharacterSoundexMatch("Catherine", "Katherine"))
	require.True(t, firstCharacterSoundexMatch("Fone", "Phone"))
	require.True(t, firstCharacterSoundexMatch("Vibe", "Bribe"))
	require.True(t, firstCharacterSoundexMatch("mine", "nine"))

	require.False(t, firstCharacterSoundexMatch("a", ""))
	require.False(t, firstCharacterSoundexMatch("", "A"))
	require.False(t, firstCharacterSoundexMatch("Dave", "Eve"))
}

func TestDisablePhoneticFiltering(t *testing.T) {
	search := strings.Fields("ian mckinley")
	indexed := "tian xiang 7"

	t.Setenv("DISABLE_PHONETIC_FILTERING", "no")
	score := BestPairsJaroWinkler(search, indexed)
	require.InDelta(t, 0.00, score, 0.01)

	// Disable filtering (force the compare)
	t.Setenv("DISABLE_PHONETIC_FILTERING", "yes")

	score = BestPairsJaroWinkler(search, indexed)
	require.InDelta(t, 0.544, score, 0.01)
}
