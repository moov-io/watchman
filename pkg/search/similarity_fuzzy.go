package search

import (
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/internal/stringscore"
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
	score         float64
	matchingTerms int
	totalTerms    int
	isExact       bool
	isHistorical  bool
}

func compareName[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	// Early return for empty query
	if query.PreparedFields.Name == "" {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "name"}
	}

	// Exact match fast path
	if query.PreparedFields.Name == index.PreparedFields.Name {
		return ScorePiece{
			Score:          1.0,
			Weight:         weight,
			Matched:        true,
			Required:       true,
			Exact:          true,
			FieldsCompared: 1,
			PieceType:      "name",
		}
	}

	// Check primary name
	bestMatch := compareNameTerms(query.PreparedFields.NameFields, index.PreparedFields.NameFields)

	// Check alternate names
	for idx := range index.PreparedFields.AltNameFields {
		altMatch := compareNameTerms(query.PreparedFields.NameFields, index.PreparedFields.AltNameFields[idx])
		if altMatch.score > bestMatch.score {
			bestMatch = altMatch
		}
	}

	// Check historical names with penalty
	for _, hist := range index.HistoricalInfo {
		if strings.EqualFold(hist.Type, "Former Name") {
			indexHistoricalTerms := strings.Fields(prepare.LowerAndRemovePunctuation(hist.Value))

			histMatch := compareNameTerms(query.PreparedFields.NameFields, indexHistoricalTerms)
			histMatch.score *= 0.95 // Apply penalty for historical names
			histMatch.isHistorical = true
			if histMatch.score > bestMatch.score {
				bestMatch = histMatch
			}
		}
	}

	// Apply additional criteria for match quality
	bestMatch.score = adjustScoreBasedOnQuality(bestMatch, len(query.PreparedFields.NameFields))
	if !isNameCloseEnough(query.PreparedFields, index.PreparedFields) {
		bestMatch.score *= 0.85
	}

	return ScorePiece{
		Score:          bestMatch.score,
		Weight:         weight,
		Matched:        bestMatch.score > 0.6,
		Required:       true,
		Exact:          bestMatch.isExact,
		FieldsCompared: 1,
		PieceType:      "name",
	}
}

// compareNameFields performs detailed term-by-term comparison
func compareNameTerms(queryTerms, indexTerms []string) nameMatch {
	var score float64
	if len(indexTerms) > 0 {
		score = stringscore.BestPairCombinationJaroWinkler(queryTerms, indexTerms)
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
