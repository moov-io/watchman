package stringscore

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/moov-io/base/strx"

	"github.com/xrash/smetrics"
)

var (
	// Jaro-Winkler parameters
	boostThreshold = readFloat(os.Getenv("JARO_WINKLER_BOOST_THRESHOLD"), 0.7)
	prefixSize     = readInt(os.Getenv("JARO_WINKLER_PREFIX_SIZE"), 4)
	// Customised Jaro-Winkler parameters
	lengthDifferenceCutoffFactor  = readFloat(os.Getenv("LENGTH_DIFFERENCE_CUTOFF_FACTOR"), 0.9)
	lengthDifferencePenaltyWeight = readFloat(os.Getenv("LENGTH_DIFFERENCE_PENALTY_WEIGHT"), 0.3)
	differentLetterPenaltyWeight  = readFloat(os.Getenv("DIFFERENT_LETTER_PENALTY_WEIGHT"), 0.9)

	// Watchman parameters
	exactMatchFavoritism        = readFloat(os.Getenv("EXACT_MATCH_FAVORITISM"), 0.0)
	unmatchedIndexPenaltyWeight = readFloat(os.Getenv("UNMATCHED_INDEX_TOKEN_WEIGHT"), 0.15)
)

func readFloat(override string, value float64) float64 {
	if override != "" {
		n, err := strconv.ParseFloat(override, 32)
		if err != nil {
			panic(fmt.Errorf("unable to parse %q as float64", override)) //nolint:forbidigo
		}
		return n
	}
	return value
}

func readInt(override string, value int) int {
	if override != "" {
		n, err := strconv.ParseInt(override, 10, 32)
		if err != nil {
			panic(fmt.Errorf("unable to parse %q as int", override)) //nolint:forbidigo
		}
		return int(n)
	}
	return value
}

// BestPairsJaroWinkler compares a search query to an indexed term (name, address, etc) and returns a decimal fraction
// score.
//
// The algorithm splits each string into tokens, and does a pairwise Jaro-Winkler score of all token combinations
// (outer product). The best match for each search token is chosen, such that each index token can be matched at most
// once.
//
// The pairwise scores are combined into an average in a way that corrects for character length, and the fraction of the
// indexed term that didn't match.
func BestPairsJaroWinkler(searchTokens []string, indexed string) float64 {
	type Score struct {
		score          float64
		searchTokenIdx int
		indexTokenIdx  int
	}

	indexedTokens := strings.Fields(indexed)
	searchTokensLength := sumLength(searchTokens)
	indexTokensLength := sumLength(indexedTokens)

	disablePhoneticFiltering := strx.Yes(os.Getenv("DISABLE_PHONETIC_FILTERING"))

	//Compare each search token to each indexed token. Sort the results in descending order
	scoresCapacity := (len(searchTokens) + len(indexedTokens))
	if !disablePhoneticFiltering {
		scoresCapacity /= 5 // reduce the capacity as many terms don't phonetically match
	}
	scores := make([]Score, 0, scoresCapacity)
	for searchIdx, searchToken := range searchTokens {
		for indexIdx, indexedToken := range indexedTokens {
			// Compare the first letters phonetically and only run jaro-winkler on those which are similar
			if disablePhoneticFiltering || firstCharacterSoundexMatch(indexedToken, searchToken) {
				score := customJaroWinkler(indexedToken, searchToken)
				scores = append(scores, Score{score, searchIdx, indexIdx})
			}
		}
	}
	sort.Slice(scores[:], func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	//Pick the highest score for each search term, where the indexed token hasn't yet been matched
	matchedSearchTokens := make([]bool, len(searchTokens))
	matchedIndexTokens := make([]bool, len(indexedTokens))
	matchedIndexTokensLength := 0
	totalWeightedScores := 0.0
	for _, score := range scores {
		//If neither the search token nor index token have been matched so far
		if !matchedSearchTokens[score.searchTokenIdx] && !matchedIndexTokens[score.indexTokenIdx] {
			//Weight the importance of this word score by its character length
			searchToken := searchTokens[score.searchTokenIdx]
			indexToken := indexedTokens[score.indexTokenIdx]
			totalWeightedScores += score.score * float64(len(searchToken)+len(indexToken))

			matchedSearchTokens[score.searchTokenIdx] = true
			matchedIndexTokens[score.indexTokenIdx] = true
			matchedIndexTokensLength += len(indexToken)
		}
	}
	lengthWeightedAverageScore := totalWeightedScores / float64(searchTokensLength+matchedIndexTokensLength)

	//If some index tokens weren't matched by any search token, penalise this search a small amount. If this isn't done,
	//a query of "John Doe" will match "John Doe" and "John Bartholomew Doe" equally well.
	//Calculate the fraction of the index name that wasn't matched, apply a weighting to reduce the importance of
	//unmatched portion, then scale down the final score.
	matchedIndexLength := 0
	for i, str := range indexedTokens {
		if matchedIndexTokens[i] {
			matchedIndexLength += len(str)
		}
	}
	matchedFraction := float64(matchedIndexLength) / float64(indexTokensLength)
	return lengthWeightedAverageScore * scalingFactor(matchedFraction, unmatchedIndexPenaltyWeight)
}

func customJaroWinkler(s1 string, s2 string) float64 {
	score := smetrics.JaroWinkler(s1, s2, boostThreshold, prefixSize)

	if lengthMetric := lengthDifferenceFactor(s1, s2); lengthMetric < lengthDifferenceCutoffFactor {
		//If there's a big difference in matched token lengths, punish the score. Jaro-Winkler is quite permissive about
		//different lengths
		score = score * scalingFactor(lengthMetric, lengthDifferencePenaltyWeight)
	}
	if s1[0] != s2[0] {
		//Penalise words that start with a different characters. Jaro-Winkler is too lenient on this
		//TODO should use a phonetic comparison here, like Soundex
		score = score * differentLetterPenaltyWeight
	}
	return score
}

// scalingFactor returns a float [0,1] that can be used to scale another number down, given some metric and a desired
// weight
// e.g. If a score has a 50% value according to a metric, and we want a 10% weight to the metric:
//
//	scaleFactor := scalingFactor(0.5, 0.1)  // 0.95
//	scaledScore := score * scaleFactor
func scalingFactor(metric float64, weight float64) float64 {
	return 1.0 - (1.0-metric)*weight
}

func sumLength(strs []string) int {
	totalLength := 0
	for _, str := range strs {
		totalLength += len(str)
	}
	return totalLength
}

func lengthDifferenceFactor(s1 string, s2 string) float64 {
	ls1 := float64(len(s1))
	ls2 := float64(len(s2))
	min := math.Min(ls1, ls2)
	max := math.Max(ls1, ls2)
	return min / max
}

// jaroWinkler runs the similarly named algorithm over the two input strings and averages their match percentages
// according to the second string (assumed to be the user's query)
//
// Terms are compared between a few adjacent terms and accumulate the highest near-neighbor match.
//
// For more details see https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func JaroWinkler(s1, s2 string) float64 {
	return JaroWinklerWithFavoritism(s1, s2, exactMatchFavoritism)
}

var (
	adjacentSimilarityPositions = readInt(os.Getenv("ADJACENT_SIMILARITY_POSITIONS"), 3)
)

func JaroWinklerWithFavoritism(indexedTerm, query string, favoritism float64) float64 {
	maxMatch := func(indexedWord string, indexedWordIdx int, queryWords []string) (float64, string) {
		if indexedWord == "" || len(queryWords) == 0 {
			return 0.0, ""
		}

		// We're only looking for the highest match close
		start := indexedWordIdx - adjacentSimilarityPositions
		end := indexedWordIdx + adjacentSimilarityPositions

		var max float64
		var maxTerm string
		for i := start; i < end; i++ {
			if i >= 0 && len(queryWords) > i {
				score := smetrics.JaroWinkler(indexedWord, queryWords[i], boostThreshold, prefixSize)
				if score > max {
					max = score
					maxTerm = queryWords[i]
				}
			}
		}
		return max, maxTerm
	}

	indexedWords, queryWords := strings.Fields(indexedTerm), strings.Fields(query)
	if len(indexedWords) == 0 || len(queryWords) == 0 {
		return 0.0 // avoid returning NaN later on
	}

	var scores []float64
	for i := range indexedWords {
		max, term := maxMatch(indexedWords[i], i, queryWords)
		if max >= 1.0 {
			// If the query is longer than our indexed term (and EITHER are longer than most names)
			// we want to reduce the maximum weight proportionally by the term difference, which
			// forces more terms to match instead of one or two dominating the weight.
			if (len(queryWords) > len(indexedWords)) && (len(indexedWords) > 3 || len(queryWords) > 3) {
				max *= (float64(len(indexedWords)) / float64(len(queryWords)))
				goto add
			}
			// If the indexed term is really short cap the match at 90%.
			// This sill allows names to match highly with a couple different characters.
			if len(indexedWords) == 1 && len(queryWords) > 1 {
				max *= 0.9
				goto add
			}
			// Otherwise, apply Perfect match favoritism
			max += favoritism
		add:
			scores = append(scores, max)
		} else {
			// If there are more terms in the user's query than what's indexed then
			// adjust the max lower by the proportion of different terms.
			//
			// We do this to decrease the importance of a short (often common) term.
			if len(queryWords) > len(indexedWords) {
				scores = append(scores, max*float64(len(indexedWords))/float64(len(queryWords)))
				continue
			}

			// Apply an additional weight based on similarity of term lengths,
			// so terms which are closer in length match higher.
			s1 := float64(len(indexedWords[i]))
			t := float64(len(term)) - 1
			weight := math.Min(math.Abs(s1/t), 1.0)

			scores = append(scores, max*weight)
		}
	}

	// average the highest N scores where N is the words in our query (query).
	// Only truncate scores if there are enough words (aka more than First/Last).
	sort.Float64s(scores)
	if len(indexedWords) > len(queryWords) && len(queryWords) > 5 {
		scores = scores[len(indexedWords)-len(queryWords):]
	}

	var sum float64
	for i := range scores {
		sum += scores[i]
	}
	return math.Min(sum/float64(len(scores)), 1.00)
}
