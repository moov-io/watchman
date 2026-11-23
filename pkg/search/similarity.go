package search

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/moov-io/watchman/internal/tfidf"
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

var (
	lowCoveragePenaltyMultiplier           = readFloat("FINAL_SCORE_LOW_COVERAGE_MULTIPLIER", 0.95)
	minimumRequiredFieldsPenaltyMultiplier = readFloat("FINAL_SCORE_MIN_REQUIRED_FIELDS_MULTIPLIER", 0.90)
	nameOnlyPenaltyMultiplier              = readFloat("FINAL_SCORE_NAME_ONLY_MULTIPLIER", 0.95)
	idConflictPenaltyMultiplier            = readFloat("ID_CONFLICT_PENALTY_MULTIPLIER", 0.70)
)

func readFloat(envVar string, defaultValue float64) float64 {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARN: could not parse %s=%s as float64: %v. Using default %f\n", envVar, value, err, defaultValue)
		return defaultValue
	}
	return n
}

// SimilarityOpts configures similarity scoring for a comparison.
type SimilarityOpts struct {
	TFIDF     *tfidf.Index
	Algorithm StringMatchAlgorithm
}

// Similarity calculates a match score between a query and an index entity.
func Similarity[Q any, I any](query Entity[Q], index Entity[I]) float64 {
	return SimilarityWithTFIDF(query, index, nil)
}

// SimilarityWithTFIDF calculates a match score with optional TF-IDF weighting for name matching.
// When tfidfIndex is nil or disabled, falls back to standard scoring.
// This path avoids allocating the detailed ScorePiece slice returned by DebugSimilarity.
func SimilarityWithTFIDF[Q any, I any](query Entity[Q], index Entity[I], tfidfIndex *tfidf.Index) float64 {
	return SimilarityWithOpts(query, index, SimilarityOpts{TFIDF: tfidfIndex})
}

// SimilarityWithOpts calculates a match score with TF-IDF weighting and a string-match algorithm.
func SimilarityWithOpts[Q any, I any](query Entity[Q], index Entity[I], opts SimilarityOpts) float64 {
	return scoreSimilarityFast(query, index, opts)
}

// DebugSimilarity does the same as Similarity, but writes SimilarityScore.String() to w.
//
// The text is intended for humans to read and will evolve over time. Prefer the
// returned SimilarityScore (name, steps, pieces) for programmatic use.
func DebugSimilarity[Q any, I any](w io.Writer, query Entity[Q], index Entity[I]) float64 {
	return DebugSimilarityWithTFIDF(w, query, index, nil).FinalScore
}

// DebugSimilarityWithTFIDF does the same as DebugSimilarity, with optional TF-IDF weighting.
func DebugSimilarityWithTFIDF[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], tfidfIndex *tfidf.Index) SimilarityScore {
	return DebugSimilarityWithOpts(w, query, index, SimilarityOpts{TFIDF: tfidfIndex})
}

// DebugSimilarityWithOpts does the same as DebugSimilarity, with scoring options.
func DebugSimilarityWithOpts[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], opts SimilarityOpts) SimilarityScore {
	details := DetailedSimilarityWithOpts(nil, query, index, opts)

	switch len(details.Pieces) {
	case 0:
		// No pieces, nothing to debug in terms of pieces
	case 1:
		// Exact SourceID match short-circuit
	case 9:
		// Full set of pieces
	default:
		panic(fmt.Sprintf("BUG: got an unexpected amount of %d ScorePieces", len(details.Pieces))) //nolint:forbidigo
	}

	if w != nil {
		fmt.Fprint(w, details.String())
	}

	return details
}

var (
	emptyPieces     = make([]ScorePiece, 9)
	emptyEntityType = EntityType("")

	exactScore = []ScorePiece{
		{
			Score:          1.0,
			Weight:         criticalIdWeight,
			Matched:        true,
			Required:       true,
			Exact:          true,
			FieldsCompared: 1,
			PieceType:      "exact",
		},
	}
)

// DetailedSimilarity returns the scoring details of each query piece against the index Entity.
//
// This is intended to give detailed results of which fields matched and how they were scored against each other.
// The fields returned in SimilarityScore may change as the general similarity algorithm and scoring methodologies evolve.
// There is no API stability guarantee for SimilarityScore.
func DetailedSimilarity[Q any, I any](w io.Writer, query Entity[Q], index Entity[I]) SimilarityScore {
	return DetailedSimilarityWithTFIDF(w, query, index, nil)
}

// DetailedSimilarityWithTFIDF returns scoring details with optional TF-IDF weighting for name matching.
func DetailedSimilarityWithTFIDF[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], tfidfIndex *tfidf.Index) SimilarityScore {
	return DetailedSimilarityWithOpts(w, query, index, SimilarityOpts{TFIDF: tfidfIndex})
}

// DetailedSimilarityWithOpts returns scoring details with optional TF-IDF weighting and algorithm.
func DetailedSimilarityWithOpts[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], opts SimilarityOpts) SimilarityScore {
	var out SimilarityScore

	// Quick filters — these are free when the search service already partitioned,
	// but remain here for direct Similarity() callers.
	if query.Source != sourceEmpty && !query.Source.IsRequestType() {
		if query.Source != index.Source {
			out.Pieces = emptyPieces
			return out
		}
	}
	if query.SourceID != "" {
		if query.SourceID == index.SourceID {
			out.FinalScore = 1.0
			out.Pieces = exactScore
			out.Steps = []ScoreStep{{Stage: "exact-source-id", Score: 1.0, Reason: "query sourceID matches the list row"}}
		} else {
			out.Pieces = emptyPieces
		}
		return out
	}
	if query.Type != emptyEntityType && query.Type != index.Type {
		recast, ok := recastToType(query, index.Type)
		if !ok {
			out.Pieces = emptyPieces
			return out
		}
		query = recast
	}

	// Stack-allocated piece buffer avoids per-comparison heap alloc of the slice header backing array
	var pieces [9]ScorePiece

	// Critical identifiers (highest weight)
	pieces[0] = compareExactIdentifiers(w, query, index, criticalIdWeight)
	if math.IsNaN(pieces[0].Score) {
		pieces[0].Score = 1.0
	}

	pieces[1] = compareExactCryptoAddresses(w, query, index, criticalIdWeight)
	if math.IsNaN(pieces[1].Score) {
		pieces[1].Score = 1.0
	}

	pieces[2] = compareExactGovernmentIDs(w, query, index, criticalIdWeight)
	if math.IsNaN(pieces[2].Score) {
		pieces[2].Score = 1.0
	}

	pieces[3] = compareExactContactInfo(w, query, index, criticalIdWeight)
	if math.IsNaN(pieces[3].Score) {
		pieces[3].Score = 1.0
	}

	// Name comparison (second highest weight) - use TF-IDF if provided
	pieces[4], out.Name = compareNameWithTFIDF(w, query, index, nameWeight, opts, true)
	pieces[5] = compareEntityTitlesFuzzy(w, query, index, nameWeight)

	// Supporting information (lower weight)
	pieces[6] = compareEntityDates(w, query, index, supportingInfoWeight)
	pieces[7] = compareAddresses(w, query, index, addressWeight)
	pieces[8] = compareSupportingInfo(w, query, index, supportingInfoWeight)

	pieceSlice := pieces[:]
	out.FinalScore, out.Steps = calculateFinalScore(w, pieceSlice, shouldExactOverride(pieceSlice), query, index, hasIdentifierConflict(query, index), true)

	// Only allocate a heap-backed Pieces slice when the caller needs details (debug writer)
	// or when Exact override short pieces aren't used. Always copy for API stability of Details.
	out.Pieces = make([]ScorePiece, 9)
	copy(out.Pieces, pieceSlice)

	return out
}

// scoreSimilarityFast is the allocation-light path used by bulk search: it returns only the final score.
func scoreSimilarityFast[Q any, I any](query Entity[Q], index Entity[I], opts SimilarityOpts) float64 {
	// Quick filters
	if query.Source != sourceEmpty && !query.Source.IsRequestType() {
		if query.Source != index.Source {
			return 0
		}
	}
	if query.SourceID != "" {
		if query.SourceID == index.SourceID {
			return 1.0
		}
		return 0
	}
	if query.Type != emptyEntityType && query.Type != index.Type {
		recast, ok := recastToType(query, index.Type)
		if !ok {
			return 0
		}
		query = recast
	}

	// Unique identity keys (passport, IMO, SWIFT/BIC, crypto) can still skip name work.
	// Tax/registration and contact (email/phone/website) matches fall through so a disagreeing name
	// can block a 1.0 override.
	p0 := compareExactIdentifiers(nil, query, index, criticalIdWeight)
	p1 := compareExactCryptoAddresses(nil, query, index, criticalIdWeight)
	p2 := compareExactGovernmentIDs(nil, query, index, criticalIdWeight)
	p3 := compareExactContactInfo(nil, query, index, criticalIdWeight)
	if uniqueIdentityExact(p0) || uniqueIdentityExact(p1) || uniqueIdentityExact(p2) {
		return 1.0
	}

	namePiece, _ := compareNameWithTFIDF(nil, query, index, nameWeight, opts, false)
	pieces := [9]ScorePiece{
		p0,
		p1,
		p2,
		p3,
		namePiece,
		compareEntityTitlesFuzzy(nil, query, index, nameWeight),
		compareEntityDates(nil, query, index, supportingInfoWeight),
		compareAddresses(nil, query, index, addressWeight),
		compareSupportingInfo(nil, query, index, supportingInfoWeight),
	}

	score, _ := calculateFinalScore(nil, pieces[:], shouldExactOverride(pieces[:]), query, index, hasIdentifierConflict(query, index), false)
	return score
}

func uniqueIdentityExact(p ScorePiece) bool {
	return p.FieldsCompared > 0 && p.Exact && p.UniqueIdentity
}

// shouldExactOverride reports whether a 1.0 final score is justified.
// Only unique identity keys (passport, IMO, SWIFT/BIC, crypto, …) short-circuit, and only
// when the piece is Exact (identifier and country). Tax IDs, business
// registrations, and contact (email/phone/website/domain) matches stay in the weighted blend.
func shouldExactOverride(pieces []ScorePiece) bool {
	for i := range pieces {
		p := pieces[i]
		if p.FieldsCompared == 0 || !p.Exact {
			continue
		}
		switch p.PieceType {
		case "crypto-exact":
			return true
		case "identifiers", "gov-ids-exact":
			if p.UniqueIdentity {
				return true
			}
		}
	}
	return false
}

// SimilarityScore gives detailed results of which fields matched and how they were scored against each other.
//
// The fields returned in SimilarityScore may change as the general similarity algorithm and scoring methodologies evolve.
// There is no API stability guarantee for SimilarityScore.
type SimilarityScore struct {
	Pieces     []ScorePiece `json:"pieces,omitzero"`
	Name       *NameDebug   `json:"name,omitempty"`
	Steps      []ScoreStep  `json:"steps,omitempty"`
	FinalScore float64      `json:"finalScore,omitzero"`
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
	// UniqueIdentity is true when the matched identifier is unique to one
	// legal person or asset (passport, IMO, crypto address) rather than a
	// reusable tax or registration number.
	UniqueIdentity bool `json:"uniqueIdentity,omitempty"`
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

func calculateFinalScore[Q any, I any](w io.Writer, pieces []ScorePiece, exactOverride bool, query Entity[Q], index Entity[I], idConflict bool, collect bool) (float64, []ScoreStep) {
	if len(pieces) == 0 {
		return 0, nil
	}

	fields := countFieldsByImportance(pieces)
	coverage := calculateCoverage(w, pieces, index)
	baseScore := calculateBaseScore(pieces, fields)

	trail := scoreTrail{collect: collect}
	applyPenaltiesAndBonuses(&trail, baseScore, coverage, fields)

	if exactOverride {
		trail.applyForced("exact-override", "unique identity key matched exactly", 1.0)
		return 1.0, trail.steps
	}

	if idConflict && idConflictPenaltyMultiplier > 0 && idConflictPenaltyMultiplier < 1 {
		trail.apply("id-conflict", "same-type identifiers disagree", trail.score*idConflictPenaltyMultiplier)
	}

	if hasNameNumeralConflict(query, index) && nameNumeralConflictPenaltyMultiplier > 0 && nameNumeralConflictPenaltyMultiplier < 1 {
		trail.apply("name-numeral-conflict", "prepared names have disagreeing numerals", trail.score*nameNumeralConflictPenaltyMultiplier)
	}

	if math.IsNaN(trail.score) {
		trail.applyForced("nan-to-one", "score was NaN", 1.0)
		return 1.0, trail.steps
	}

	return trail.score, trail.steps
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
			if piece.Exact && piece.UniqueIdentity && (piece.PieceType == "identifiers" || piece.PieceType == "gov-ids-exact" || piece.PieceType == "crypto-exact") {
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

func applyPenaltiesAndBonuses(trail *scoreTrail, baseScore float64, cov coverage, fields entityFields) {
	trail.base(baseScore)

	if cov.ratio < minCoverageThreshold {
		trail.applyf("low-coverage", trail.score*lowCoveragePenaltyMultiplier, "compared %.0f%% of index fields", cov.ratio*100)
	}
	if cov.criticalRatio < criticalCovThreshold {
		trail.applyf("critical-coverage", trail.score*0.90, "critical field coverage %.0f%%", cov.criticalRatio*100)
	}

	if fields.required < 2 {
		trail.apply("min-required-fields", "fewer than 2 required fields compared", trail.score*minimumRequiredFieldsPenaltyMultiplier)
	}

	if !fields.hasID && !fields.hasAddress && fields.hasName {
		trail.apply("name-only", "name-only query", trail.score*nameOnlyPenaltyMultiplier)
	}

	if fields.hasName && fields.hasID && fields.hasCritical && cov.ratio > 0.70 && trail.score > highConfidenceThreshold {
		trail.apply("perfect-match", "name, unique ID, and high coverage", math.Min(1.0, trail.score*perfectMatchBoost))
	}
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
	if len(index.Contact.Websites) > 0 {
		count++
	}
	if len(index.CryptoAddresses) > 0 {
		count++
	}
	if len(index.SwiftBICs) > 0 {
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
