package search

import (
	"fmt"
	"strings"

	"github.com/moov-io/watchman/internal/stringscore"
)

// StringMatchAlgorithm selects the name-matching algorithm used during search.
// The empty value is the default Watchman Jaro-Winkler setup.
type StringMatchAlgorithm string

const (
	// AlgorithmJaroWinkler is the default token pairwise Jaro-Winkler scorer.
	AlgorithmJaroWinkler StringMatchAlgorithm = "jaro-winkler"

	// AlgorithmSoundex uses the Jaro-Winkler setup with a Soundex phonetic boost
	// for token pairs that encode to the same Soundex code.
	AlgorithmSoundex StringMatchAlgorithm = "soundex"

	// AlgorithmSoftBidist is Soft-Bidist (Hadwan 2021), a character-bigram
	// edit distance. Aliases: soft-bigram, bidist.
	AlgorithmSoftBidist StringMatchAlgorithm = "soft-bidist"

	// AlgorithmSoftBisim is Soft-Bisim (Millán-Hernández 2019), a
	// character-bigram similarity. Alias: bisim.
	AlgorithmSoftBisim StringMatchAlgorithm = "soft-bisim"

	// defaultSoundexBoostWeight is applied when algorithm=soundex is requested
	// and SOUNDEX_BOOST_WEIGHT is unset (0).
	defaultSoundexBoostWeight = 0.12
)

const supportedAlgorithms = "jaro-winkler, soundex, soft-bidist, soft-bisim"

// ParseStringMatchAlgorithm parses a caller-provided algorithm name.
// Empty input is the default Jaro-Winkler setup (process env flags still apply).
func ParseStringMatchAlgorithm(raw string) (StringMatchAlgorithm, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	v = strings.ReplaceAll(v, "_", "-")
	switch v {
	case "", string(AlgorithmJaroWinkler), "jarowinkler", "jw":
		if v == "" {
			return "", nil
		}
		return AlgorithmJaroWinkler, nil
	case string(AlgorithmSoundex):
		return AlgorithmSoundex, nil
	case string(AlgorithmSoftBidist), "soft-bigram", "bidist":
		return AlgorithmSoftBidist, nil
	case string(AlgorithmSoftBisim), "bisim":
		return AlgorithmSoftBisim, nil
	default:
		return "", fmt.Errorf("unknown algorithm %q (supported: %s)", raw, supportedAlgorithms)
	}
}

// Name returns a stable identifier for telemetry and logs.
func (a StringMatchAlgorithm) Name() string {
	if a == "" {
		return string(AlgorithmJaroWinkler)
	}
	return string(a)
}

func (a StringMatchAlgorithm) scoringConfig() stringscore.ScoringConfig {
	cfg := stringscore.DefaultScoringConfig()
	switch a {
	case AlgorithmSoundex:
		cfg.UseSoundexBoost = true
		if cfg.SoundexBoostWeight <= 0 {
			cfg.SoundexBoostWeight = defaultSoundexBoostWeight
		}
	case AlgorithmJaroWinkler:
		cfg.UseSoundexBoost = false
	case AlgorithmSoftBidist:
		cfg.UseSoundexBoost = false
		cfg.TokenScorer = stringscore.TokenScorerSoftBidist
	case AlgorithmSoftBisim:
		cfg.UseSoundexBoost = false
		cfg.TokenScorer = stringscore.TokenScorerSoftBisim
	}
	return cfg
}
