package search

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"
	"github.com/moov-io/watchman/internal/tfidf"
)

const (
	// Minimum number of significant terms that must match for a high confidence match
	minMatchingTerms = 2

	// Score thresholds for term matching
	termMatchThreshold = 0.90 // Individual term match threshold
	nameMatchThreshold = 0.85 // Overall name match threshold
)

// nameMatch tracks detailed matching information
type nameMatch struct {
	score              float64
	tokenScore         float64
	afterHist          float64
	afterQuality       float64
	matchingTerms      int
	totalTerms         int
	isExact            bool
	isHistorical       bool
	closeEnoughApplied bool

	querySide     string
	indexSide     string
	queryOriginal string
	queryPrepared string
	queryTerms    []string
	indexOriginal string
	indexPrepared string
	indexTerms    []string
}

func compareName[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	piece, _ := compareNameWithTFIDF(w, query, index, weight, SimilarityOpts{}, false)
	return piece
}

func compareNameWithTFIDF[Q any, I any](_ io.Writer, query Entity[Q], index Entity[I], weight float64, opts SimilarityOpts, explain bool) (ScorePiece, *NameDebug) {
	// Early return for empty query
	if query.PreparedFields.Name == "" {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "name"}, nil
	}

	queryTerms := query.PreparedFields.NameFields
	queryWeights := query.PreparedFields.NameWeights
	cfg := opts.Algorithm.scoringConfig()

	bestMatch := nameMatch{
		querySide:     nameAgainstPrimary,
		indexSide:     nameAgainstPrimary,
		queryOriginal: query.Name,
		queryPrepared: query.PreparedFields.Name,
		queryTerms:    queryTerms,
		indexOriginal: index.Name,
		indexPrepared: index.PreparedFields.Name,
		indexTerms:    index.PreparedFields.NameFields,
	}

	// Exact match fast path
	if query.PreparedFields.Name == index.PreparedFields.Name {
		bestMatch.score = 1.0
		bestMatch.tokenScore = 1.0
		bestMatch.isExact = true
		bestMatch.matchingTerms = len(queryTerms)
		bestMatch.totalTerms = len(queryTerms)
		var debug *NameDebug
		if explain {
			debug = explainName(query, index, bestMatch, opts, true)
		}
		return namePiece(bestMatch, weight), debug
	}

	// Check primary name
	primary := compareNameTermsWeighted(queryTerms, index.PreparedFields.NameFields, queryWeights, index.PreparedFields.NameWeights, opts.TFIDF, cfg)
	takeNameMatch(&bestMatch, primary, false, nameAgainstPrimary, nameAgainstPrimary, query.Name, query.PreparedFields.Name, queryTerms, index.Name, index.PreparedFields.Name, index.PreparedFields.NameFields)

	var altOriginals, histOriginals, queryHistOriginals []string
	if explain {
		altOriginals = originalAltNames(index)
		histOriginals = originalHistoricalNames(index)
		queryHistOriginals = originalHistoricalNames(query)
	}

	// Check alternate names unless the primary is already an exact-quality match.
	if bestMatch.score < exactMatchThreshold {
		for idx := range index.PreparedFields.AltNameFields {
			var indexWeights []float64
			if idx < len(index.PreparedFields.AltNameWeights) {
				indexWeights = index.PreparedFields.AltNameWeights[idx]
			}
			orig := pickName(altOriginals, idx)
			if orig == "" {
				orig = pickName(index.PreparedFields.AltNames, idx)
			}
			prep := pickName(index.PreparedFields.AltNames, idx)
			takeNameMatch(&bestMatch, compareNameTermsWeighted(queryTerms, index.PreparedFields.AltNameFields[idx], queryWeights, indexWeights, opts.TFIDF, cfg), false, nameAgainstPrimary, nameAgainstAlt, query.Name, query.PreparedFields.Name, queryTerms, orig, prep, index.PreparedFields.AltNameFields[idx])
			if bestMatch.score >= exactMatchThreshold {
				break
			}
		}
	}

	// Index former names (query primary vs list historical)
	if bestMatch.score < exactMatchThreshold {
		for idx := range index.PreparedFields.HistoricalNameFields {
			var indexWeights []float64
			if idx < len(index.PreparedFields.HistoricalNameWeights) {
				indexWeights = index.PreparedFields.HistoricalNameWeights[idx]
			}
			orig := pickName(histOriginals, idx)
			if orig == "" {
				orig = pickName(index.PreparedFields.HistoricalNames, idx)
			}
			prep := pickName(index.PreparedFields.HistoricalNames, idx)
			takeNameMatch(&bestMatch, compareNameTermsWeighted(queryTerms, index.PreparedFields.HistoricalNameFields[idx], queryWeights, indexWeights, opts.TFIDF, cfg), true, nameAgainstPrimary, nameAgainstHistorical, query.Name, query.PreparedFields.Name, queryTerms, orig, prep, index.PreparedFields.HistoricalNameFields[idx])
			if bestMatch.score >= exactMatchThreshold {
				break
			}
		}
	}

	// Query former names (query historical vs list primary / alt / historical).
	// Needed when the query is "Ocean Pioneer (ex-Cape Diamond)" and the list row is "Cape Diamond".
	if bestMatch.score < exactMatchThreshold {
		for qIdx := range query.PreparedFields.HistoricalNameFields {
			var qHistWeights []float64
			if qIdx < len(query.PreparedFields.HistoricalNameWeights) {
				qHistWeights = query.PreparedFields.HistoricalNameWeights[qIdx]
			}
			qHistTerms := query.PreparedFields.HistoricalNameFields[qIdx]
			qOrig := pickName(queryHistOriginals, qIdx)
			if qOrig == "" {
				qOrig = pickName(query.PreparedFields.HistoricalNames, qIdx)
			}
			qPrep := pickName(query.PreparedFields.HistoricalNames, qIdx)

			takeNameMatch(&bestMatch, compareNameTermsWeighted(qHistTerms, index.PreparedFields.NameFields, qHistWeights, index.PreparedFields.NameWeights, opts.TFIDF, cfg), true, nameAgainstHistorical, nameAgainstPrimary, qOrig, qPrep, qHistTerms, index.Name, index.PreparedFields.Name, index.PreparedFields.NameFields)
			if bestMatch.score >= exactMatchThreshold {
				break
			}
			for idx := range index.PreparedFields.AltNameFields {
				var indexWeights []float64
				if idx < len(index.PreparedFields.AltNameWeights) {
					indexWeights = index.PreparedFields.AltNameWeights[idx]
				}
				orig := pickName(altOriginals, idx)
				if orig == "" {
					orig = pickName(index.PreparedFields.AltNames, idx)
				}
				prep := pickName(index.PreparedFields.AltNames, idx)
				takeNameMatch(&bestMatch, compareNameTermsWeighted(qHistTerms, index.PreparedFields.AltNameFields[idx], qHistWeights, indexWeights, opts.TFIDF, cfg), true, nameAgainstHistorical, nameAgainstAlt, qOrig, qPrep, qHistTerms, orig, prep, index.PreparedFields.AltNameFields[idx])
				if bestMatch.score >= exactMatchThreshold {
					break
				}
			}
			if bestMatch.score >= exactMatchThreshold {
				break
			}
			for idx := range index.PreparedFields.HistoricalNameFields {
				var indexWeights []float64
				if idx < len(index.PreparedFields.HistoricalNameWeights) {
					indexWeights = index.PreparedFields.HistoricalNameWeights[idx]
				}
				orig := pickName(histOriginals, idx)
				if orig == "" {
					orig = pickName(index.PreparedFields.HistoricalNames, idx)
				}
				prep := pickName(index.PreparedFields.HistoricalNames, idx)
				takeNameMatch(&bestMatch, compareNameTermsWeighted(qHistTerms, index.PreparedFields.HistoricalNameFields[idx], qHistWeights, indexWeights, opts.TFIDF, cfg), true, nameAgainstHistorical, nameAgainstHistorical, qOrig, qPrep, qHistTerms, orig, prep, index.PreparedFields.HistoricalNameFields[idx])
				if bestMatch.score >= exactMatchThreshold {
					break
				}
			}
		}
	}

	termCount := len(queryTerms)
	if bestMatch.isHistorical && termCount < minMatchingTerms {
		termCount = bestMatch.totalTerms
	}

	afterHist := bestMatch.score
	afterQuality := adjustScoreBasedOnQuality(bestMatch, termCount)
	bestMatch.score = afterQuality
	closeEnough := bestMatch.isHistorical || isNameCloseEnough(query.PreparedFields, index.PreparedFields)
	if !closeEnough {
		bestMatch.score *= 0.85
	}
	bestMatch.afterHist = afterHist
	bestMatch.afterQuality = afterQuality
	bestMatch.closeEnoughApplied = !closeEnough

	var debug *NameDebug
	if explain {
		debug = explainName(query, index, bestMatch, opts, false)
	}
	return namePiece(bestMatch, weight), debug
}

func namePiece(match nameMatch, weight float64) ScorePiece {
	return ScorePiece{
		Score:          match.score,
		Weight:         weight,
		Matched:        match.score > 0.6,
		Required:       true,
		Exact:          match.isExact,
		FieldsCompared: 1,
		PieceType:      "name",
	}
}

func pickName(names []string, idx int) string {
	if idx < len(names) {
		return names[idx]
	}
	return ""
}

func takeNameMatch(best *nameMatch, candidate nameMatch, historical bool, querySide, indexSide, queryOrig, queryPrep string, queryTerms []string, indexOrig, indexPrep string, indexTerms []string) {
	candidate.tokenScore = candidate.score
	if historical {
		candidate.score *= 0.95
		candidate.isHistorical = true
	}
	if candidate.score <= best.score {
		return
	}
	candidate.querySide = querySide
	candidate.indexSide = indexSide
	candidate.queryOriginal = queryOrig
	candidate.queryPrepared = queryPrep
	candidate.queryTerms = queryTerms
	candidate.indexOriginal = indexOrig
	candidate.indexPrepared = indexPrep
	candidate.indexTerms = indexTerms
	*best = candidate
}

func explainName[Q any, I any](query Entity[Q], index Entity[I], match nameMatch, opts SimilarityOpts, exactPrepared bool) *NameDebug {
	debug := &NameDebug{
		Score:      match.score,
		Exact:      match.score >= exactMatchThreshold,
		Historical: match.isHistorical,
		QueryName:  match.querySide,
		IndexName:  match.indexSide,
		Query:      nameSideFrom(query.Name, match.queryOriginal, match.queryPrepared, match.queryTerms, query.Type),
		Index:      nameSideFrom(index.Name, match.indexOriginal, match.indexPrepared, match.indexTerms, index.Type),
	}

	if exactPrepared {
		debug.Pairs = identityTokenPairs(match.queryTerms)
		debug.Steps = []ScoreStep{{Stage: "exact-prepared", Score: 1.0, Reason: "prepared names are identical"}}
		return debug
	}

	cfg := opts.Algorithm.scoringConfig()
	queryWeights := matchQueryWeights(query, match)
	indexWeights := matchIndexWeights(index, match)
	useWeighted := len(queryWeights) == len(match.queryTerms) && len(indexWeights) == len(match.indexTerms)
	if !useWeighted && opts.TFIDF != nil && opts.TFIDF.Enabled() {
		queryWeights = opts.TFIDF.GetWeights(match.queryTerms)
		indexWeights = opts.TFIDF.GetWeights(match.indexTerms)
		useWeighted = len(queryWeights) == len(match.queryTerms) && len(indexWeights) == len(match.indexTerms)
	}
	if useWeighted {
		_, pairs := stringscore.ExplainBestPairCombinationJaroWinklerWeighted(match.queryTerms, match.indexTerms, queryWeights, indexWeights, cfg)
		debug.Pairs = tokenPairsFromExplain(pairs)
	} else {
		_, pairs := stringscore.ExplainBestPairCombinationJaroWinkler(match.queryTerms, match.indexTerms, cfg)
		debug.Pairs = tokenPairsFromExplain(pairs)
	}

	prev := match.tokenScore
	debug.Steps = appendNameStep(debug.Steps, "token-pairs", "", 0, match.tokenScore)
	if match.isHistorical {
		debug.Steps = appendNameStep(debug.Steps, "historical", "former-name penalty 0.95", prev, match.afterHist)
		prev = match.afterHist
	}
	if match.afterQuality != prev {
		debug.Steps = appendNameStep(debug.Steps, "few-matching-terms", "fewer than 2 terms above the term-match threshold", prev, match.afterQuality)
		prev = match.afterQuality
	}
	if match.closeEnoughApplied {
		qCount := strings.Count(query.PreparedFields.Name, " ") + 1
		iCount := strings.Count(index.PreparedFields.Name, " ") + 1
		debug.Steps = appendNameStep(debug.Steps, "close-enough", fmt.Sprintf("primary name term counts %d vs %d", qCount, iCount), prev, match.score)
	}
	return debug
}

func matchQueryWeights[Q any](query Entity[Q], match nameMatch) []float64 {
	if match.querySide == nameAgainstHistorical {
		for i, prepared := range query.PreparedFields.HistoricalNames {
			if prepared == match.queryPrepared && i < len(query.PreparedFields.HistoricalNameWeights) {
				return query.PreparedFields.HistoricalNameWeights[i]
			}
		}
		return nil
	}
	return query.PreparedFields.NameWeights
}

func matchIndexWeights[I any](index Entity[I], match nameMatch) []float64 {
	switch match.indexSide {
	case nameAgainstAlt:
		for i, prepared := range index.PreparedFields.AltNames {
			if prepared == match.indexPrepared && i < len(index.PreparedFields.AltNameWeights) {
				return index.PreparedFields.AltNameWeights[i]
			}
		}
		return nil
	case nameAgainstHistorical:
		for i, prepared := range index.PreparedFields.HistoricalNames {
			if prepared == match.indexPrepared && i < len(index.PreparedFields.HistoricalNameWeights) {
				return index.PreparedFields.HistoricalNameWeights[i]
			}
		}
		return nil
	}
	return index.PreparedFields.NameWeights
}

// compareNameFields performs detailed term-by-term comparison
func compareNameTerms(queryTerms, indexTerms []string) nameMatch {
	return compareNameTermsWeighted(queryTerms, indexTerms, nil, nil, nil, stringscore.DefaultScoringConfig())
}

// compareNameTermsWithTFIDF performs term-by-term comparison with optional TF-IDF weighting.
// When tfidfIndex is nil or disabled, falls back to unweighted comparison.
func compareNameTermsWithTFIDF(queryTerms, indexTerms []string, tfidfIndex *tfidf.Index) nameMatch {
	return compareNameTermsWeighted(queryTerms, indexTerms, nil, nil, tfidfIndex, stringscore.DefaultScoringConfig())
}

// compareNameTermsWeighted prefers precomputed weights; falls back to tfidfIndex.GetWeights
// only when weights are missing and the index is enabled.
func compareNameTermsWeighted(queryTerms, indexTerms []string, queryWeights, indexWeights []float64, tfidfIndex *tfidf.Index, cfg stringscore.ScoringConfig) nameMatch {
	var score float64
	if len(indexTerms) > 0 {
		useWeighted := false
		if len(queryWeights) == len(queryTerms) && len(indexWeights) == len(indexTerms) {
			useWeighted = true
		} else if tfidfIndex != nil && tfidfIndex.Enabled() {
			queryWeights = tfidfIndex.GetWeights(queryTerms)
			indexWeights = tfidfIndex.GetWeights(indexTerms)
			useWeighted = len(queryWeights) == len(queryTerms) && len(indexWeights) == len(indexTerms)
		}

		if useWeighted {
			score = stringscore.BestPairCombinationJaroWinklerWeightedWithConfig(queryTerms, indexTerms, queryWeights, indexWeights, cfg)
		} else {
			score = stringscore.BestPairCombinationJaroWinklerWithConfig(queryTerms, indexTerms, cfg)
		}
	}

	matchingTerms := 0
	if score > termMatchThreshold {
		matchingTerms = len(queryTerms)
	}

	return nameMatch{
		score:         score,
		matchingTerms: matchingTerms,
		totalTerms:    len(queryTerms),
		isExact:       score > exactMatchThreshold,
	}
}

// adjustScoreBasedOnQuality applies additional quality criteria
func adjustScoreBasedOnQuality(match nameMatch, queryTermCount int) float64 {
	// Require minimum number of matching terms for high scores
	if match.matchingTerms < minMatchingTerms && queryTermCount >= minMatchingTerms {
		return match.score * 0.8 // Significant penalty for too few matching terms
	}

	// Historical names already have a penalty applied
	if match.isHistorical {
		return match.score
	}

	return match.score
}

// isHighConfidenceMatch determines if the match quality is sufficient
func isHighConfidenceMatch(match nameMatch, finalScore float64) bool {
	// Must meet both term matching and score criteria
	return match.matchingTerms >= minMatchingTerms && finalScore > nameMatchThreshold
}

func isNameCloseEnough(query, index PreparedFields) bool {
	qTermCount := strings.Count(query.Name, " ") + 1 // one space == two terms
	iTermCount := strings.Count(index.Name, " ") + 1

	// More lenient length difference check
	if float64(qTermCount)/float64(iTermCount) < 0.3 || float64(qTermCount)/float64(iTermCount) > 3.0 {
		return false
	}

	// More lenient matching threshold
	matchCount := 0
	for q := range query.NameFields {
		for i := range index.NameFields {
			if strings.Contains(index.NameFields[i], query.NameFields[q]) || strings.Contains(query.NameFields[q], index.NameFields[i]) {
				matchCount++
				break
			}
		}
	}

	return float64(matchCount)/float64(qTermCount) >= 0.4
}

const (
	titleMatchThreshold   = 0.85 // Threshold for considering titles matched
	minTitleTermLength    = 2    // Minimum length for title terms
	abbreviationThreshold = 0.92 // Threshold for matching abbreviated titles
)

var (
	// Common title abbreviations and their full forms
	titleAbbreviations = map[string]string{
		"ceo":   "chief executive officer",
		"cfo":   "chief financial officer",
		"coo":   "chief operating officer",
		"pres":  "president",
		"vp":    "vice president",
		"dir":   "director",
		"exec":  "executive",
		"mgr":   "manager",
		"sr":    "senior",
		"jr":    "junior",
		"asst":  "assistant",
		"assoc": "associate",
		"tech":  "technical",
		"admin": "administrator",
		"eng":   "engineer",
		"dev":   "developer",
	}

	// Patterns to clean up titles
	punctRegexp = regexp.MustCompile(`[^\w\s-]`)
	spaceRegexp = regexp.MustCompile(`\s+`)
)

func compareEntityTitlesFuzzy[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	if query.Person == nil || index.Person == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "titles"}
	}

	// Prepare normalized index titles once
	normalizedIndexTitles := make([]string, 0, len(index.Person.Titles))
	for _, title := range index.Person.Titles {
		if normalized := normalizeTitle(title); normalized != "" {
			normalizedIndexTitles = append(normalizedIndexTitles, normalized)
		}
	}

	if len(normalizedIndexTitles) == 0 {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "titles"}
	}

	fieldsCompared := 0
	matches := 0
	total := 0

	for _, qTitle := range query.Person.Titles {
		normalizedQuery := normalizeTitle(qTitle)
		if normalizedQuery == "" {
			continue
		}

		fieldsCompared++
		total++

		// Try exact match first
		if score := findBestTitleMatch(normalizedQuery, normalizedIndexTitles); score > 0 {
			if score > titleMatchThreshold {
				matches++
			}
			continue
		}

		// Try matching with expanded abbreviations
		expandedQuery := expandAbbreviations(normalizedQuery)
		if score := findBestTitleMatch(expandedQuery, normalizedIndexTitles); score > titleMatchThreshold {
			matches++
			continue
		}

		// Try matching each index title with expanded abbreviations
		bestScore := 0.0
		for _, iTitle := range normalizedIndexTitles {
			expandedIndex := expandAbbreviations(iTitle)
			score := calculateTitleSimilarity(normalizedQuery, expandedIndex)
			if score > bestScore {
				bestScore = score
			}
		}

		if bestScore > titleMatchThreshold {
			matches++
		}
	}

	var finalScore float64
	if total > 0 {
		finalScore = float64(matches) / float64(total)
	}

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        finalScore > 0.5,
		Required:       false,
		Exact:          finalScore > exactMatchThreshold,
		FieldsCompared: fieldsCompared,
		PieceType:      "titles",
	}
}

// normalizeTitle cleans and normalizes a title string
func normalizeTitle(title string) string {
	// Convert to lowercase and trim
	title = strings.TrimSpace(strings.ToLower(title))

	// Remove punctuation except hyphens
	title = punctRegexp.ReplaceAllString(title, " ")

	// Normalize spaces
	title = spaceRegexp.ReplaceAllString(title, " ")

	// Final trim
	return strings.TrimSpace(title)
}

// expandAbbreviations expands known abbreviations in a title
func expandAbbreviations(title string) string {
	words := strings.Fields(title)
	expanded := make([]string, 0, len(words))

	for _, word := range words {
		if full, exists := titleAbbreviations[word]; exists {
			expanded = append(expanded, full)
		} else {
			expanded = append(expanded, word)
		}
	}

	return strings.Join(expanded, " ")
}

// calculateTitleSimilarity computes similarity between two titles
func calculateTitleSimilarity(title1, title2 string) float64 {
	// Handle exact matches
	if title1 == title2 {
		return 1.0
	}

	// Split into terms
	terms1 := strings.Fields(title1)
	terms2 := strings.Fields(title2)

	// Filter out very short terms
	terms1 = filterTerms(terms1)
	terms2 = filterTerms(terms2)

	if len(terms1) == 0 || len(terms2) == 0 {
		return 0.0
	}

	// Use JaroWinkler for term comparison
	score := stringscore.BestPairCombinationJaroWinkler(terms1, terms2)

	// Adjust score based on length difference
	lengthDiff := math.Abs(float64(len(terms1) - len(terms2)))
	if lengthDiff > 0 {
		score *= (1.0 - (lengthDiff * 0.1))
	}

	return score
}

// findBestTitleMatch finds the best matching index title for a query title
func findBestTitleMatch(queryTitle string, indexTitles []string) float64 {
	bestScore := 0.0

	for _, indexTitle := range indexTitles {
		score := calculateTitleSimilarity(queryTitle, indexTitle)
		if score > bestScore {
			bestScore = score
			if score > abbreviationThreshold {
				break // Good enough match found
			}
		}
	}

	return bestScore
}

// filterTerms removes terms that are too short
func filterTerms(terms []string) []string {
	filtered := make([]string, 0, len(terms))
	for _, term := range terms {
		if len(term) >= minTitleTermLength {
			filtered = append(filtered, term)
		}
	}
	return filtered
}

const (
	// Thresholds for name matching
	affiliationNameThreshold = 0.85

	// Type match bonuses/penalties
	exactTypeBonus    = 0.15 // Bonus for exact type match
	relatedTypeBonues = 0.08 // Bonus for related type match
	typeMatchPenalty  = 0.15 // Penalty for type mismatch
)

// affiliationTypeGroup groups similar affiliation types
var affiliationTypeGroups = map[string][]string{
	"ownership": {
		"owned by", "subsidiary of", "parent of", "holding company",
		"owner", "owned", "subsidiary", "parent",
	},
	"control": {
		"controlled by", "controls", "managed by", "manages",
		"operated by", "operates",
	},
	"association": {
		"linked to", "associated with", "affiliated with", "related to",
		"connection to", "connected with",
	},
	"leadership": {
		"led by", "leader of", "directed by", "directs",
		"headed by", "heads",
	},
}

// affiliationMatch tracks match details
type affiliationMatch struct {
	nameScore  float64
	typeScore  float64
	finalScore float64
	exactMatch bool
}

func compareAffiliationsFuzzy[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	// Early return if no affiliations to compare
	if len(query.Affiliations) == 0 {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "affiliations",
		}
	}

	// Validate index affiliations
	if len(index.Affiliations) == 0 {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 1, // We had query affiliations but no index matches
			PieceType:      "affiliations",
		}
	}

	// Process each query affiliation
	var matches []affiliationMatch
	for _, qAff := range query.Affiliations {
		// Skip empty affiliations
		if match := findBestAffiliationMatch(qAff, index.Affiliations); match.nameScore > 0 {
			matches = append(matches, match)
		}
	}

	if len(matches) == 0 {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 1,
			PieceType:      "affiliations",
		}
	}

	// Calculate final score
	finalScore := calculateFinalAffiliateScore(matches)

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        finalScore > affiliationNameThreshold,
		Required:       false,
		Exact:          finalScore > exactMatchThreshold,
		FieldsCompared: 1,
		PieceType:      "affiliations",
	}
}

// findBestAffiliationMatch finds the best matching index affiliation
func findBestAffiliationMatch(queryAff Affiliation, indexAffs []Affiliation) affiliationMatch {
	qName := normalizeAffiliationName(queryAff.EntityName)
	if qName == "" {
		return affiliationMatch{}
	}

	qFields := strings.Fields(qName)
	if len(qFields) == 0 {
		return affiliationMatch{}
	}

	var bestMatch affiliationMatch

	for _, iAff := range indexAffs {
		iName := normalizeAffiliationName(iAff.EntityName)
		if iName == "" {
			continue
		}

		// Calculate name match score
		iFields := strings.Fields(iName)
		if len(iFields) == 0 {
			continue
		}

		nameScore := calculateNameScore(qFields, iFields)
		if nameScore <= bestMatch.nameScore {
			continue
		}

		// Calculate type match score
		typeScore := calculateTypeScore(queryAff.Type, iAff.Type)

		// Calculate combined score with type influence
		finalScore := calculateCombinedScore(nameScore, typeScore)

		if finalScore > bestMatch.finalScore {
			bestMatch = affiliationMatch{
				nameScore:  nameScore,
				typeScore:  typeScore,
				finalScore: finalScore,
				exactMatch: nameScore > exactMatchThreshold && typeScore > 0.9,
			}
		}
	}

	return bestMatch
}

func calculateNameScore(query, index []string) float64 {
	if len(query) == 0 || len(index) == 0 {
		return 0.0
	}
	return stringscore.BestPairCombinationJaroWinkler(query, index)
}

// normalizeAffiliationName normalizes an entity name for comparison
func normalizeAffiliationName(name string) string {
	// Basic normalization
	name = strings.TrimSpace(strings.ToLower(name))

	// Remove common business suffixes
	suffixes := []string{" inc", " ltd", " llc", " corp", " co", " company"}
	for _, suffix := range suffixes {
		name = strings.TrimSuffix(name, suffix)
	}
	return strings.TrimSpace(name)
}

// calculateTypeScore determines how well affiliation types match
func calculateTypeScore(queryType, indexType string) float64 {
	queryType = prepare.LowerAndRemovePunctuation(queryType)
	indexType = prepare.LowerAndRemovePunctuation(indexType)

	// Exact type match
	if queryType == indexType {
		return 1.0
	}

	// Check if types are in the same group
	queryGroup := getTypeGroup(queryType)
	indexGroup := getTypeGroup(indexType)

	if queryGroup != "" && queryGroup == indexGroup {
		return 0.8
	}

	return 0.0
}

// getTypeGroup finds which group a type belongs to
func getTypeGroup(affType string) string {
	for group, types := range affiliationTypeGroups {
		for _, t := range types {
			if strings.EqualFold(affType, t) {
				return group
			}
		}
	}
	return ""
}

// calculateCombinedScore combines name and type scores
func calculateCombinedScore(nameScore, typeScore float64) float64 {
	// Base score is the name match score
	score := nameScore

	// Apply type match bonus/penalty
	if typeScore > 0.9 {
		score += exactTypeBonus
	} else if typeScore > 0.7 {
		score += relatedTypeBonues
	} else {
		score -= typeMatchPenalty
	}

	// Ensure score stays in valid range
	if score > 1.0 {
		score = 1.0
	}
	if score < 0.0 {
		score = 0.0
	}

	return score
}

// calculateFinalAffiliateScore determines overall affiliation match score
func calculateFinalAffiliateScore(matches []affiliationMatch) float64 {
	if len(matches) == 0 {
		return 0.0
	}

	// Calculate weighted average giving more weight to better matches
	var weightedSum float64
	var totalWeight float64

	for _, match := range matches {
		// Weight is the square of the score to emphasize better matches
		weight := match.finalScore * match.finalScore
		weightedSum += match.finalScore * weight
		totalWeight += weight
	}

	if totalWeight == 0 {
		return 0.0
	}

	return weightedSum / totalWeight
}
