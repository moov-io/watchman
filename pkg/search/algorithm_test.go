package search

import (
	"testing"

	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/stretchr/testify/require"
)

func TestParseStringMatchAlgorithm(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in      string
		want    StringMatchAlgorithm
		wantErr bool
	}{
		{in: "", want: ""},
		{in: "jaro-winkler", want: AlgorithmJaroWinkler},
		{in: "Jaro_Winkler", want: AlgorithmJaroWinkler},
		{in: "jw", want: AlgorithmJaroWinkler},
		{in: "soundex", want: AlgorithmSoundex},
		{in: " Soundex ", want: AlgorithmSoundex},
		{in: "soft-bidist", want: AlgorithmSoftBidist},
		{in: "soft_bigram", want: AlgorithmSoftBidist},
		{in: "bidist", want: AlgorithmSoftBidist},
		{in: "soft-bisim", want: AlgorithmSoftBisim},
		{in: "bisim", want: AlgorithmSoftBisim},
		{in: "editex", want: AlgorithmEditex},
		{in: "nsim", want: AlgorithmNSim},
		{in: "kondrak", want: AlgorithmNSim},
		{in: "nsim-3", want: AlgorithmNSim3},
		{in: "trigram", want: AlgorithmNSim3},
		{in: "double-metaphone", want: AlgorithmDoubleMetaphone},
		{in: "dmetaphone", want: AlgorithmDoubleMetaphone},
		{in: "beider-morse", want: AlgorithmBeiderMorse},
		{in: "bmpm", want: AlgorithmBeiderMorse},
		{in: "levenshtein", wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := ParseStringMatchAlgorithm(tc.in)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestStringMatchAlgorithm_ScoringConfig(t *testing.T) {
	t.Cleanup(stringscore.ResetEnvConfigForTest)
	stringscore.ResetEnvConfigForTest()

	jw := AlgorithmJaroWinkler.scoringConfig()
	require.False(t, jw.UseSoundexBoost)

	sx := AlgorithmSoundex.scoringConfig()
	require.True(t, sx.UseSoundexBoost)
	require.InDelta(t, defaultSoundexBoostWeight, sx.SoundexBoostWeight, 0.0001)

	bidist := AlgorithmSoftBidist.scoringConfig()
	require.False(t, bidist.UseSoundexBoost)
	require.Equal(t, stringscore.TokenScorerSoftBidist, bidist.TokenScorer)

	bisim := AlgorithmSoftBisim.scoringConfig()
	require.False(t, bisim.UseSoundexBoost)
	require.Equal(t, stringscore.TokenScorerSoftBisim, bisim.TokenScorer)

	ed := AlgorithmEditex.scoringConfig()
	require.Equal(t, stringscore.TokenScorerEditex, ed.TokenScorer)

	ns := AlgorithmNSim.scoringConfig()
	require.Equal(t, stringscore.TokenScorerNSim, ns.TokenScorer)

	ns3 := AlgorithmNSim3.scoringConfig()
	require.Equal(t, stringscore.TokenScorerNSim3, ns3.TokenScorer)

	dm := AlgorithmDoubleMetaphone.scoringConfig()
	require.True(t, dm.UseDoubleMetaphoneBoost)
	require.InDelta(t, defaultSoundexBoostWeight, dm.SoundexBoostWeight, 0.0001)

	bm := AlgorithmBeiderMorse.scoringConfig()
	require.True(t, bm.UseBeiderMorseBoost)
}
