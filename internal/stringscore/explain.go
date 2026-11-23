package stringscore

import "math"

// TokenPair is a greedy token alignment used to explain a BestPairs score.
// An empty Search or Index means that token was unmatched.
type TokenPair struct {
	Search string
	Index  string
	Score  float64
}

// ExplainBestPairsJaroWinkler is BestPairsJaroWinklerWithConfig plus the token pairing.
func ExplainBestPairsJaroWinkler(searchTokens, indexedTokens []string, cfg ScoringConfig) (float64, []TokenPair) {
	return bestPairsJaroWinkler(searchTokens, indexedTokens, cfg, true)
}

// ExplainBestPairCombinationJaroWinkler is BestPairCombinationJaroWinklerWithConfig plus the winning token pairing.
func ExplainBestPairCombinationJaroWinkler(searchTokens, indexedTokens []string, cfg ScoringConfig) (float64, []TokenPair) {
	return bestPairCombinationJaroWinkler(searchTokens, indexedTokens, cfg, true)
}

// ExplainBestPairCombinationJaroWinklerWeighted is BestPairCombinationJaroWinklerWeightedWithConfig plus the winning token pairing.
func ExplainBestPairCombinationJaroWinklerWeighted(searchTokens, indexedTokens []string, searchWeights, indexWeights []float64, cfg ScoringConfig) (float64, []TokenPair) {
	return bestPairCombinationJaroWinklerWeighted(searchTokens, indexedTokens, searchWeights, indexWeights, cfg, true)
}

func bestPairCombinationJaroWinkler(searchTokens, indexedTokens []string, cfg ScoringConfig, collect bool) (float64, []TokenPair) {
	searchCombinations := GenerateWordCombinations(searchTokens)
	indexedCombinations := GenerateWordCombinations(indexedTokens)

	var maxScore float64
	var bestPairs []TokenPair
	have := false
	for _, searchVariation := range searchCombinations {
		for _, indexedVariation := range indexedCombinations {
			score, pairs := bestPairsJaroWinkler(searchVariation, indexedVariation, cfg, collect)
			if !have || score > maxScore {
				maxScore = score
				bestPairs = pairs
				have = true
			}
		}
	}
	return maxScore, bestPairs
}

func bestPairCombinationJaroWinklerWeighted(searchTokens, indexedTokens []string, searchWeights, indexWeights []float64, cfg ScoringConfig, collect bool) (float64, []TokenPair) {
	if len(searchWeights) != len(searchTokens) || len(indexWeights) != len(indexedTokens) {
		return bestPairCombinationJaroWinkler(searchTokens, indexedTokens, cfg, collect)
	}

	searchCombinations := GenerateWordCombinations(searchTokens)
	indexedCombinations := GenerateWordCombinations(indexedTokens)
	searchWeightCombinations := generateWeightCombinations(searchTokens, searchWeights)
	indexWeightCombinations := generateWeightCombinations(indexedTokens, indexWeights)

	var maxScore float64
	var bestPairs []TokenPair
	have := false
	for si, searchVariation := range searchCombinations {
		for ii, indexedVariation := range indexedCombinations {
			score, pairs := bestPairsJaroWinklerWeighted(searchVariation, indexedVariation,
				searchWeightCombinations[si], indexWeightCombinations[ii], cfg, collect)
			if !have || score > maxScore {
				maxScore = score
				bestPairs = pairs
				have = true
			}
		}
	}
	return maxScore, bestPairs
}

func bestPairsJaroWinkler(searchTokens, indexedTokens []string, cfg ScoringConfig, collect bool) (float64, []TokenPair) {
	searchTokensLength := sumLength(searchTokens)
	indexTokensLength := sumLength(indexedTokens)

	skipPhonetic := disablePhoneticFiltering.Load()

	scratch := getBestPairsScratch()
	defer bestPairsPool.Put(scratch)
	scratch.reset(pairScoreCapacity(len(searchTokens), len(indexedTokens), skipPhonetic), len(searchTokens), len(indexedTokens))

	for searchIdx, searchToken := range searchTokens {
		for indexIdx, indexedToken := range indexedTokens {
			if skipPhonetic || firstCharacterSoundexMatch(indexedToken, searchToken) {
				scratch.scores = append(scratch.scores, scorePair{
					score:          tokenSimilarity(indexedToken, searchToken, cfg),
					searchTokenIdx: searchIdx,
					indexTokenIdx:  indexIdx,
				})
			}
		}
	}
	sortScorePairs(scratch.scores)

	var pairs []TokenPair
	if collect {
		pairs = make([]TokenPair, 0, len(searchTokens)+len(indexedTokens))
	}

	matchedIndexTokensLength := 0
	totalWeightedScores := 0.0
	for _, score := range scratch.scores {
		if !scratch.matchedSearch[score.searchTokenIdx] && !scratch.matchedIndex[score.indexTokenIdx] {
			searchToken := searchTokens[score.searchTokenIdx]
			indexToken := indexedTokens[score.indexTokenIdx]
			totalWeightedScores += score.score * float64(len(searchToken)+len(indexToken))

			if collect {
				pairs = append(pairs, TokenPair{Search: searchToken, Index: indexToken, Score: score.score})
			}

			scratch.matchedSearch[score.searchTokenIdx] = true
			scratch.matchedIndex[score.indexTokenIdx] = true
			matchedIndexTokensLength += len(indexToken)
		}
	}
	lengthWeightedAverageScore := totalWeightedScores / float64(searchTokensLength+matchedIndexTokensLength)
	if math.IsNaN(lengthWeightedAverageScore) {
		lengthWeightedAverageScore = 0.0
	}

	matchedIndexLength := 0
	for i, str := range indexedTokens {
		if scratch.matchedIndex[i] {
			matchedIndexLength += len(str)
		}
	}
	matchedFraction := float64(matchedIndexLength) / float64(indexTokensLength)
	if math.IsNaN(matchedFraction) {
		matchedFraction = 0.0
	}

	if collect {
		pairs = appendUnmatchedPairs(pairs, searchTokens, indexedTokens, scratch.matchedSearch, scratch.matchedIndex)
	}

	return lengthWeightedAverageScore * scalingFactor(matchedFraction, unmatchedIndexPenaltyWeight), pairs
}

func bestPairsJaroWinklerWeighted(searchTokens, indexedTokens []string, searchWeights, indexWeights []float64, cfg ScoringConfig, collect bool) (float64, []TokenPair) {
	if len(searchWeights) != len(searchTokens) || len(indexWeights) != len(indexedTokens) {
		return bestPairsJaroWinkler(searchTokens, indexedTokens, cfg, collect)
	}

	skipPhonetic := disablePhoneticFiltering.Load()

	scratch := getBestPairsScratch()
	defer bestPairsPool.Put(scratch)
	scratch.reset(pairScoreCapacity(len(searchTokens), len(indexedTokens), skipPhonetic), len(searchTokens), len(indexedTokens))

	for searchIdx, searchToken := range searchTokens {
		for indexIdx, indexedToken := range indexedTokens {
			if skipPhonetic || firstCharacterSoundexMatch(indexedToken, searchToken) {
				scratch.scores = append(scratch.scores, scorePair{
					score:          tokenSimilarity(indexedToken, searchToken, cfg),
					searchTokenIdx: searchIdx,
					indexTokenIdx:  indexIdx,
				})
			}
		}
	}
	sortScorePairs(scratch.scores)

	var pairs []TokenPair
	if collect {
		pairs = make([]TokenPair, 0, len(searchTokens)+len(indexedTokens))
	}

	var totalWeightedScores float64
	var totalWeight float64

	for _, score := range scratch.scores {
		if !scratch.matchedSearch[score.searchTokenIdx] && !scratch.matchedIndex[score.indexTokenIdx] {
			pairWeight := (searchWeights[score.searchTokenIdx] + indexWeights[score.indexTokenIdx]) / 2.0
			totalWeightedScores += score.score * pairWeight
			totalWeight += pairWeight

			if collect {
				pairs = append(pairs, TokenPair{
					Search: searchTokens[score.searchTokenIdx],
					Index:  indexedTokens[score.indexTokenIdx],
					Score:  score.score,
				})
			}

			scratch.matchedSearch[score.searchTokenIdx] = true
			scratch.matchedIndex[score.indexTokenIdx] = true
		}
	}

	if collect {
		pairs = appendUnmatchedPairs(pairs, searchTokens, indexedTokens, scratch.matchedSearch, scratch.matchedIndex)
	}

	if totalWeight == 0 {
		return 0.0, pairs
	}

	weightedAverageScore := totalWeightedScores / totalWeight

	var matchedWeight, totalIndexWeight float64
	for i := range indexedTokens {
		totalIndexWeight += indexWeights[i]
		if scratch.matchedIndex[i] {
			matchedWeight += indexWeights[i]
		}
	}

	if totalIndexWeight > 0 {
		matchedFraction := matchedWeight / totalIndexWeight
		return weightedAverageScore * scalingFactor(matchedFraction, unmatchedIndexPenaltyWeight), pairs
	}

	return weightedAverageScore, pairs
}

func appendUnmatchedPairs(pairs []TokenPair, searchTokens, indexedTokens []string, matchedSearch, matchedIndex []bool) []TokenPair {
	for i, tok := range searchTokens {
		if !matchedSearch[i] {
			pairs = append(pairs, TokenPair{Search: tok})
		}
	}
	for i, tok := range indexedTokens {
		if !matchedIndex[i] {
			pairs = append(pairs, TokenPair{Index: tok})
		}
	}
	return pairs
}
