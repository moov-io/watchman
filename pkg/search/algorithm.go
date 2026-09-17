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

	// AlgorithmEditex is Editex (Zobel & Dart 1996), phonetic-group edit distance.
	AlgorithmEditex StringMatchAlgorithm = "editex"

	// AlgorithmNSim is Kondrak N-SIM with n=2 (BI-SIM). Aliases: n-sim, kondrak.
	AlgorithmNSim StringMatchAlgorithm = "nsim"

	// AlgorithmNSim3 is Kondrak N-SIM with n=3 (Trigram-2B affixing). Alias: nsim-3.
	AlgorithmNSim3 StringMatchAlgorithm = "nsim-3"

	// AlgorithmDoubleMetaphone is Jaro-Winkler with a Double Metaphone boost
	// when primary or alternate codes overlap. Aliases: doublemetaphone, dmetaphone.
	AlgorithmDoubleMetaphone StringMatchAlgorithm = "double-metaphone"

	// AlgorithmBeiderMorse is Jaro-Winkler with a Beider-Morse boost when
	// generic BMPM keys overlap. Aliases: beider-morse, bmpm.
	AlgorithmBeiderMorse StringMatchAlgorithm = "beider-morse"

	// defaultSoundexBoostWeight is applied when algorithm=soundex is requested
	// and SOUNDEX_BOOST_WEIGHT is unset (0). Also used for Double Metaphone
	// and Beider-Morse when those algorithms are selected.
	defaultSoundexBoostWeight = 0.12
)

const supportedAlgorithms = "jaro-winkler, soundex, soft-bidist, soft-bisim, editex, nsim, nsim-3, double-metaphone, beider-morse"

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
	case string(AlgorithmEditex):
		return AlgorithmEditex, nil
	case string(AlgorithmNSim), "n-sim", "kondrak":
		return AlgorithmNSim, nil
	case string(AlgorithmNSim3), "nsim3", "trigram":
		return AlgorithmNSim3, nil
	case string(AlgorithmDoubleMetaphone), "doublemetaphone", "dmetaphone", "metaphone":
		return AlgorithmDoubleMetaphone, nil
	case string(AlgorithmBeiderMorse), "beidermorse", "bmpm":
		return AlgorithmBeiderMorse, nil
	default:
		// Omit the raw value so HTTP error logs are not a log-injection source.
		return "", fmt.Errorf("unknown algorithm (supported: %s)", supportedAlgorithms)
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
	case AlgorithmEditex:
		cfg.UseSoundexBoost = false
		cfg.TokenScorer = stringscore.TokenScorerEditex
	case AlgorithmNSim:
		cfg.UseSoundexBoost = false
		cfg.TokenScorer = stringscore.TokenScorerNSim
	case AlgorithmNSim3:
		cfg.UseSoundexBoost = false
		cfg.TokenScorer = stringscore.TokenScorerNSim3
	case AlgorithmDoubleMetaphone:
		cfg.UseSoundexBoost = false
		cfg.UseDoubleMetaphoneBoost = true
		if cfg.SoundexBoostWeight <= 0 {
			cfg.SoundexBoostWeight = defaultSoundexBoostWeight
		}
	case AlgorithmBeiderMorse:
		cfg.UseSoundexBoost = false
		cfg.UseBeiderMorseBoost = true
		if cfg.SoundexBoostWeight <= 0 {
			cfg.SoundexBoostWeight = defaultSoundexBoostWeight
		}
	}
	return cfg
}
