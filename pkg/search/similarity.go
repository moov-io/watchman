package search

import (
	"fmt"
	"io"
	"math"
)

const (
	// Score thresholds
	exactMatchThreshold     = 0.99
	highConfidenceThreshold = 0.95

	// Weights for different categories
	criticalIdWeight     = 50.0
	nameWeight           = 35.0
	addressWeight        = 25.0
	supportingInfoWeight = 15.0
)

// Similarity calculates a match score between a query and an index entity.
func Similarity[Q any, I any](query Entity[Q], index Entity[I]) float64 {
	return DebugSimilarity(nil, query, index)
}

// DebugSimilarity does the same as Similarity, but logs debug info to w.
//
// The format written to w is not machine readable and is intended for humans to read.
// The format will evolve over time. No stability guarantee is given over what is written.
func DebugSimilarity[Q any, I any](w io.Writer, query Entity[Q], index Entity[I]) float64 {
	details := DetailedSimilarity(w, query, index)
	if len(details.Pieces) != 9 {
		panic(fmt.Sprintf("BUG: got an unexpected amount of %d ScorePieces", len(details.Pieces))) //nolint:forbidigo
	}

	if w != nil {
		// Critical comparisons (exact matches)
		debug(w, "Critical pieces\n")
		debug(w, "  exact identifiers: %#v\n", details.Pieces[0])
		debug(w, "  crypto addresses: %#v\n", details.Pieces[1])
		debug(w, "  gov IDs: %#v\n", details.Pieces[2])
		debug(w, "  contact info: %#v\n", details.Pieces[3])

		// Name comparison (second highest weight)
		debug(w, "name comparison\n")
		debug(w, "  name: %#v\n", details.Pieces[4])
		debug(w, "  titles: %#v\n", details.Pieces[5])

		// Supporting information (lower weight)
		debug(w, "supporting info\n")
		debug(w, "  dates: %#v\n", details.Pieces[6])
		debug(w, "  addresses: %#v\n", details.Pieces[7])
		debug(w, "  supporting into: %#v\n", details.Pieces[8])

		// Final Score
		debug(w, "finalScore=%.2f", details.FinalScore)
	}

	return details.FinalScore
}

// DetailedSimilarity returns the scoring details of each query piece against the index Entity.
//
// This is intended to give detailed results of which fields matched and how they were scored against each other.
// The fields returned in SimilarityScore may change as the general similarity algorithm and scoring methodologies evolve.
// There is no API stability guarantee for SimilarityScore.
func DetailedSimilarity[Q any, I any](w io.Writer, query Entity[Q], index Entity[I]) SimilarityScore {
	out := SimilarityScore{
		Pieces: make([]ScorePiece, 0, 9),
	}

	// Critical identifiers (highest weight)
	var exactOverride bool
	exactIdentifiers := compareExactIdentifiers(w, query, index, criticalIdWeight)
	if exactIdentifiers.Matched && exactIdentifiers.FieldsCompared > 0 {
		exactOverride = true
		if math.IsNaN(exactIdentifiers.Score) {
			exactIdentifiers.Score = 1.0
		}
	}

	exactCryptoAddresses := compareExactCryptoAddresses(w, query, index, criticalIdWeight)
	if exactCryptoAddresses.Matched && exactCryptoAddresses.FieldsCompared > 0 {
		exactOverride = true
		if math.IsNaN(exactCryptoAddresses.Score) {
			exactCryptoAddresses.Score = 1.0
		}
	}

	exactGovernmentIDs := compareExactGovernmentIDs(w, query, index, criticalIdWeight)
	if exactGovernmentIDs.Matched && exactGovernmentIDs.FieldsCompared > 0 {
		exactOverride = true
		if math.IsNaN(exactGovernmentIDs.Score) {
			exactGovernmentIDs.Score = 1.0
		}
	}

	exactContactInfo := compareExactContactInfo(w, query, index, criticalIdWeight)
	if exactContactInfo.Matched && exactContactInfo.FieldsCompared > 0 {
		exactOverride = true
		if math.IsNaN(exactContactInfo.Score) {
			exactContactInfo.Score = 1.0
		}
	}
	out.Pieces = append(out.Pieces, exactIdentifiers, exactCryptoAddresses, exactGovernmentIDs, exactContactInfo)

	// Name comparison (second highest weight)
	out.Pieces = append(out.Pieces,
		compareName(w, query, index, nameWeight),
		compareEntityTitlesFuzzy(w, query, index, nameWeight),
	)

	// Supporting information (lower weight)
	out.Pieces = append(out.Pieces,
		compareEntityDates(w, query, index, supportingInfoWeight),
		compareAddresses(w, query, index, addressWeight),
		compareSupportingInfo(w, query, index, supportingInfoWeight),
	)

	out.FinalScore = calculateFinalScore(w, out.Pieces, exactOverride, query, index)
	if math.IsNaN(out.FinalScore) {
		out.FinalScore = 1.0
	}
	return out
}

// SimilarityScore gives detailed results of which fields matched and how they were scored against each other.
//
// The fields returned in SimilarityScore may change as the general similarity algorithm and scoring methodologies evolve.
// There is no API stability guarantee for SimilarityScore.
type SimilarityScore struct {
	Pieces     []ScorePiece `json:"pieces"`
	FinalScore float64      `json:"finalScore"`
}

// ScorePiece is a partial scoring result from one comparison function
//
// There is no API stability guarantee for ScorePiece.
type ScorePiece struct {
	Score          float64 `json:"score"`          // 0-1 for this piece
	Weight         float64 `json:"weight"`         // weight for final
	Matched        bool    `json:"matched"`        // whether there's a "match"
	Required       bool    `json:"required"`       // if this piece is "required" for a high overall score
	Exact          bool    `json:"exact"`          // whether it's an exact match
	FieldsCompared int     `json:"fieldsCompared"` // how many fields were actually compared
	PieceType      string  `json:"pieceType"`      // e.g. "identifiers", "name", etc.
}

func boolToScore(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func calculateAverage(scores []float64) float64 {
	if len(scores) == 0 {
		return 0
	}
	var sum float64
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}

// debug prints if w is non-nil
func debug(w io.Writer, pattern string, args ...any) {
	if w != nil {
		fmt.Fprintf(w, pattern, args...)
	}
}

const (
	// Score thresholds
	typeMismatchScore       = 0.667
	criticalCovThreshold    = 0.7
	minCoverageThreshold    = 0.35 // how many fields did the query compare the index against?
	perfectMatchBoost       = 1.15
	criticalFieldMultiplier = 1.2
)

// entityFields tracks required and available fields for an entity
type entityFields struct {
	required    int
	available   int
	hasName     bool
	hasID       bool
	hasCritical bool
	hasAddress  bool
}

func calculateFinalScore[Q any, I any](w io.Writer, pieces []ScorePiece, exactOverride bool, query Entity[Q], index Entity[I]) float64 {
	if len(pieces) == 0 || query.Type != index.Type {
		return 0
	}

	// Get field counts and critical field information
	fields := countFieldsByImportance(pieces)
	coverage := calculateCoverage(w, pieces, index)

	// Calculate base score with weighted importance
	baseScore := calculateBaseScore(pieces, fields)

	// Apply coverage penalties
	finalScore := applyPenaltiesAndBonuses(w, baseScore, coverage, fields, query.Type == index.Type)

	if w != nil {
		debug(w, "calculateFinalScore:\n")
		debug(w, "  exactOverride=%v\n", exactOverride)
		debug(w, "  fields=%#v\n", fields)
		debug(w, "  coverage=%#v\n", coverage)
		debug(w, "  baseScore=%v\n", baseScore)
		debug(w, "  finalScore=%.2f (overridden: %v)\n", finalScore, exactOverride)
	}
	if exactOverride {
		return 1.0
	}
	return finalScore
}

func countFieldsByImportance(pieces []ScorePiece) entityFields {
	var fields entityFields

	for _, piece := range pieces {
		if piece.Weight <= 0 || piece.FieldsCompared == 0 {
			continue
		}

		if piece.Required {
			fields.required += piece.FieldsCompared
		}
		if piece.Matched {
			if piece.PieceType == "name" {
				fields.hasName = true
			}
			if piece.Exact && (piece.PieceType == "identifiers" || piece.PieceType == "gov-ids-exact") {
				fields.hasID = true
			}
			if piece.PieceType == "address" {
				fields.hasAddress = true
			}
			if piece.Exact {
				fields.hasCritical = true
			}
		}
	}

	return fields
}

func calculateBaseScore(pieces []ScorePiece, fields entityFields) float64 {
	var totalScore, totalWeight float64

	for _, piece := range pieces {
		if piece.Weight <= 0 || piece.FieldsCompared == 0 {
			continue
		}

		// Apply importance multiplier for critical fields
		multiplier := 1.0
		if piece.Required {
			multiplier = criticalFieldMultiplier
		}

		totalScore += piece.Score * piece.Weight * multiplier
		totalWeight += piece.Weight * multiplier
	}

	if totalWeight == 0 {
		return 0
	}

	return totalScore / totalWeight
}

func calculateCoverage[I any](w io.Writer, pieces []ScorePiece, index Entity[I]) coverage {
	indexFields := countAvailableFields(index)
	if indexFields == 0 {
		return coverage{ratio: 1.0, criticalRatio: 1.0}
	}

	var fieldsCompared, criticalFieldsCompared int
	var criticalTotal int

	for _, p := range pieces {
		fieldsCompared += p.FieldsCompared
		if p.Required {
			criticalFieldsCompared += p.FieldsCompared
			criticalTotal += p.FieldsCompared
		}
	}

	if w != nil {
		debug(w, "fieldsCompared=%v  indexFields=%v  criticalFieldsCompared=%v  criticalTotal=%v\n",
			fieldsCompared, indexFields, criticalFieldsCompared, criticalTotal)
	}

	return coverage{
		ratio:         float64(fieldsCompared) / float64(indexFields),
		criticalRatio: float64(criticalFieldsCompared) / float64(criticalTotal),
	}
}

type coverage struct {
	ratio         float64 `json:"ratio"`
	criticalRatio float64 `json:"criticalRatio"`
}

func applyPenaltiesAndBonuses(w io.Writer, baseScore float64, cov coverage, fields entityFields, sameType bool) float64 {
	score := baseScore

	if w != nil {
		debug(w, "applyPenaltiesAndBonuses\n")
		debug(w, "  start: %.2f\n", score)
	}

	// Lighter coverage penalties
	if cov.ratio < minCoverageThreshold {
		score *= 0.95

		if w != nil {
			debug(w, "  cov.ratio < minCoverageThreshold = %.2f\n", score)
		}
	}
	if cov.criticalRatio < criticalCovThreshold {
		score *= 0.90

		if w != nil {
			debug(w, "  cov.criticalRatio < criticalCovThreshold = %.2f\n", score)
		}
	}

	// Lighter minimum fields requirement
	if fields.required < 2 {
		score *= 0.90

		if w != nil {
			debug(w, "  fields.required < 2 = %.2f\n", score)
		}
	}

	// Reduced name-only match penalty
	if !fields.hasID && !fields.hasAddress && fields.hasName {
		score *= 0.95

		if w != nil {
			debug(w, "  reduced name-only match penalty = %.2f\n", score)
		}
	}

	// Perfect match requirements
	if fields.hasName && fields.hasID && fields.hasCritical && cov.ratio > 0.70 && score > highConfidenceThreshold {
		score = math.Min(1.0, score*perfectMatchBoost)

		if w != nil {
			debug(w, "  perfect match requirements = %.2f\n", score)
		}
	}

	// Handle type mismatches
	if !sameType {
		score = 0.0

		if w != nil {
			debug(w, "  !sameType = %.2f\n", score)
		}
	}

	return score
}

func countAvailableFields[I any](index Entity[I]) int {
	var count int

	// Count type-specific fields
	switch index.Type {
	case EntityPerson:
		count = countPersonFields(index.Person)
	case EntityBusiness:
		count = countBusinessFields(index.Business)
	case EntityOrganization:
		count = countOrganizationFields(index.Organization)
	case EntityVessel:
		count = countVesselFields(index.Vessel)
	case EntityAircraft:
		count = countAircraftFields(index.Aircraft)
	}

	// Count common fields
	count += countCommonFields(index)

	return count
}

func countCommonFields[I any](index Entity[I]) int {
	count := 0

	if index.Name != "" {
		count++
	}
	if index.Source != "" {
		count++
	}
	if len(index.Contact.EmailAddresses) > 0 {
		count++
	}
	if len(index.Contact.PhoneNumbers) > 0 {
		count++
	}
	if len(index.Contact.FaxNumbers) > 0 {
		count++
	}
	if len(index.CryptoAddresses) > 0 {
		count++
	}
	if len(index.Affiliations) > 0 {
		count++
	}
	if len(index.Addresses) > 0 {
		count++
	}

	return count
}

func countPersonFields(p *Person) int {
	if p == nil {
		return 0
	}

	count := 0
	if p.BirthDate != nil {
		count++
	}
	if p.Gender != "" {
		count++
	}
	if len(p.Titles) > 0 {
		count++
	}
	if len(p.GovernmentIDs) > 0 {
		count++
	}

	return count
}

func countBusinessFields(b *Business) int {
	if b == nil {
		return 0
	}

	count := 0
	if b.Name != "" {
		count++
	}
	if len(b.AltNames) > 0 {
		count++
	}
	if b.Created != nil {
		count++
	}
	if len(b.GovernmentIDs) > 0 {
		count++
	}

	return count
}

func countOrganizationFields(o *Organization) int {
	if o == nil {
		return 0
	}

	count := 0
	if o.Name != "" {
		count++
	}
	if len(o.AltNames) > 0 {
		count++
	}
	if o.Created != nil {
		count++
	}
	if len(o.GovernmentIDs) > 0 {
		count++
	}

	return count
}

func countVesselFields(v *Vessel) int {
	if v == nil {
		return 0
	}

	count := 0
	if v.IMONumber != "" {
		count++
	}
	if v.CallSign != "" {
		count++
	}
	if v.MMSI != "" {
		count++
	}
	if v.Flag != "" {
		count++
	}
	if v.Model != "" {
		count++
	}
	if v.Owner != "" {
		count++
	}

	return count
}

func countAircraftFields(a *Aircraft) int {
	if a == nil {
		return 0
	}

	count := 0
	if a.ICAOCode != "" {
		count++
	}
	if a.Model != "" {
		count++
	}
	if a.Flag != "" {
		count++
	}
	if a.SerialNumber != "" {
		count++
	}

	return count
}
