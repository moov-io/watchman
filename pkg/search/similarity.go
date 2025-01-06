package search

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/stringscore"
)

// Similarity calculates a match score between a query and an index entity.
func Similarity[Q any, I any](query Entity[Q], index Entity[I]) float64 {
	pieces := make([]scorePiece, 0)

	// Primary identifiers (IMO number, Call Sign, etc.) - highest weight
	exactMatchWeight := 50.0
	pieces = append(pieces, compareExactIdentifiers(query, index, exactMatchWeight))

	// Name match is critical
	nameWeight := 30.0
	pieces = append(pieces, compareName(query, index, nameWeight))

	// Entity-specific comparisons (type, flag, etc)
	entityWeight := 15.0
	pieces = append(pieces, compareEntitySpecific(query, index, entityWeight))

	// Supporting information (addresses, sanctions, etc)
	supportingWeight := 5.0
	pieces = append(pieces, compareSupportingInfo(query, index, supportingWeight))

	// Compute final score with coverage logic
	return calculateFinalScore(pieces, index)
}

// scorePiece is a partial scoring result from one comparison function
type scorePiece struct {
	score          float64 // 0-1 score for this piece
	weight         float64 // Weight for this piece
	matched        bool    // Whether there was a "match"
	required       bool    // Whether this piece is "required" for a high overall score
	exact          bool    // Whether this was an exact match
	fieldsCompared int     // Number of fields actually compared in this piece
	pieceType      string  // e.g. "name", "entity", "identifiers", etc.
}

func compareExactIdentifiers[Q any, I any](query Entity[Q], index Entity[I], weight float64) scorePiece {
	matches := 0
	totalWeight := 0.0
	score := 0.0
	fieldsCompared := 0
	hasMatch := false

	// For Vessel example:
	if query.Vessel != nil && index.Vessel != nil {
		// IMO Number
		if query.Vessel.IMONumber != "" {
			fieldsCompared++
			if index.Vessel.IMONumber != "" && strings.EqualFold(query.Vessel.IMONumber, index.Vessel.IMONumber) {
				matches++
				hasMatch = true
				score += 10.0
				totalWeight += 10.0
			}
		}
		// Call Sign
		if query.Vessel.CallSign != "" {
			fieldsCompared++
			if index.Vessel.CallSign != "" && strings.EqualFold(query.Vessel.CallSign, index.Vessel.CallSign) {
				matches++
				hasMatch = true
				score += 8.0
				totalWeight += 8.0
			}
		}
		// MMSI
		if query.Vessel.MMSI != "" {
			fieldsCompared++
			if index.Vessel.MMSI != "" && strings.EqualFold(query.Vessel.MMSI, index.Vessel.MMSI) {
				matches++
				hasMatch = true
				score += 6.0
				totalWeight += 6.0
			}
		}
	}

	// Similar logic for Person, Aircraft, etc. as needed...

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}
	// Penalty if we compared but found no matches
	if fieldsCompared > 0 && matches == 0 {
		finalScore = 0.1
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.95,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

func compareName[Q any, I any](query Entity[Q], index Entity[I], weight float64) scorePiece {
	qName := strings.TrimSpace(strings.ToLower(query.Name))
	iName := strings.TrimSpace(strings.ToLower(index.Name))

	// If the query name is empty, skip
	if qName == "" {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "name"}
	}

	// Exact match
	if qName == iName {
		return scorePiece{
			score:          1.0,
			weight:         weight,
			matched:        true,
			required:       true,
			exact:          true,
			fieldsCompared: 1,
			pieceType:      "name",
		}
	}

	// Fuzzy match
	qTerms := strings.Fields(qName)
	bestScore := stringscore.BestPairsJaroWinkler(qTerms, iName)

	// Check alternate names if both are Person entities
	if query.Person != nil && index.Person != nil {
		for _, altName := range index.Person.AltNames {
			altScore := stringscore.BestPairsJaroWinkler(qTerms, strings.ToLower(altName))
			if altScore > bestScore {
				bestScore = altScore
			}
		}
	}
	// Check historical info for "Former Name"
	for _, hist := range index.HistoricalInfo {
		if strings.EqualFold(hist.Type, "Former Name") {
			histScore := stringscore.BestPairsJaroWinkler(qTerms, strings.ToLower(hist.Value))
			if histScore > bestScore {
				bestScore = histScore
			}
		}
	}

	return scorePiece{
		score:          bestScore,
		weight:         weight,
		matched:        bestScore > 0.8,
		required:       true,
		exact:          bestScore > 0.99,
		fieldsCompared: 1,
		pieceType:      "name",
	}
}

const (
	debugEntitySpecific = true
)

func compareEntitySpecific[Q any, I any](query Entity[Q], index Entity[I], weight float64) scorePiece {
	// If types don't match, it's an immediate 0
	if query.Type != index.Type {
		return scorePiece{
			score:          0,
			weight:         weight,
			pieceType:      "entity",
			fieldsCompared: 1,
		}
	}

	var typeScore float64
	var matched bool
	var fieldsCompared int

	switch query.Type {
	case EntityVessel:
		typeScore, matched, fieldsCompared = compareVesselFields(query.Vessel, index.Vessel)
	case EntityPerson:
		typeScore, matched, fieldsCompared = comparePersonFields(query.Person, index.Person)
	case EntityAircraft:
		typeScore, matched, fieldsCompared = compareAircraftFields(query.Aircraft, index.Aircraft)
	case EntityBusiness:
		typeScore, matched, fieldsCompared = compareBusinessFields(query.Business, index.Business)
	case EntityOrganization:
		typeScore, matched, fieldsCompared = compareOrganizationFields(query.Organization, index.Organization)
	}
	if debugEntitySpecific {
		debug("%v  typeScore=%.4f  matched=%v  fieldsCompared=%v\n", query.Type, typeScore, matched, fieldsCompared)
	}

	return scorePiece{
		score:          typeScore,
		weight:         weight,
		matched:        matched,
		required:       fieldsCompared > 0,
		exact:          typeScore > 0.95,
		fieldsCompared: fieldsCompared + 1, // +1 for the type comparison
		pieceType:      "entity",
	}
}

const (
	day = 24 * time.Hour
)

// -------------------------------
// Person-Specific Fields
// -------------------------------
func comparePersonFields(query *Person, index *Person) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	scores := make([]float64, 0)
	fieldsCompared := 0

	// Birthdate
	if query.BirthDate != nil && index.BirthDate != nil {
		fieldsCompared++

		qb := query.BirthDate.Truncate(day)
		ib := index.BirthDate.Truncate(day)

		if qb.Equal(ib) {
			scores = append(scores, 1.0)
		} else {
			scores = append(scores, 0.0)
		}
	}

	// Gender
	if query.Gender != "" {
		fieldsCompared++
		if strings.EqualFold(string(query.Gender), string(index.Gender)) {
			scores = append(scores, 1.0)
		} else {
			scores = append(scores, 0.0)
		}
	}

	// Titles
	if len(query.Titles) > 0 && len(index.Titles) > 0 {
		fieldsCompared++
		matches := 0
		for _, qTitle := range query.Titles {
			for _, iTitle := range index.Titles {
				if strings.EqualFold(qTitle, iTitle) {
					matches++
					break
				}
			}
		}
		scores = append(scores, float64(matches)/float64(len(query.Titles)))
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	sum := 0.0
	for _, s := range scores {
		sum += s
	}
	avg := sum / float64(len(scores))

	return avg, avg > 0.5, fieldsCompared
}

// -------------------------------
// Vessel-Specific Fields
// -------------------------------
func compareVesselFields(query *Vessel, index *Vessel) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	type fieldScore struct {
		score  float64
		weight float64
	}
	var (
		scores         []fieldScore
		fieldsCompared int
	)

	// Compare only fields present in query
	if query.CallSign != "" {
		fieldsCompared++
		if strings.EqualFold(query.CallSign, index.CallSign) {
			scores = append(scores, fieldScore{1.0, 4.0})
		} else {
			scores = append(scores, fieldScore{0.0, 4.0})
		}
	}
	if query.IMONumber != "" && index.IMONumber != "" {
		fieldsCompared++
		if strings.EqualFold(query.IMONumber, index.IMONumber) {
			scores = append(scores, fieldScore{1.0, 4.0})
		} else {
			scores = append(scores, fieldScore{0.0, 4.0})
		}
	}
	if query.Owner != "" {
		fieldsCompared++
		ownerTerms := strings.Fields(strings.ToLower(query.Owner))
		ownerScore := stringscore.BestPairsJaroWinkler(ownerTerms, strings.ToLower(index.Owner))
		scores = append(scores, fieldScore{ownerScore, 2.0})
	}
	if query.Flag != "" {
		fieldsCompared++
		if strings.EqualFold(query.Flag, index.Flag) {
			scores = append(scores, fieldScore{1.0, 1.5})
		} else {
			scores = append(scores, fieldScore{0.0, 1.5})
		}
	}
	if query.Type != "" {
		fieldsCompared++
		if strings.EqualFold(string(query.Type), string(index.Type)) {
			scores = append(scores, fieldScore{1.0, 1.0})
		} else {
			scores = append(scores, fieldScore{0.0, 1.0})
		}
	}
	if query.Tonnage > 0 && index.Tonnage > 0 {
		fieldsCompared++
		diff := math.Abs(float64(query.Tonnage - index.Tonnage))
		s := vesselTonnageScore(diff)
		scores = append(scores, fieldScore{s, 1.0})
	}
	if query.GrossRegisteredTonnage > 0 && index.GrossRegisteredTonnage > 0 {
		fieldsCompared++
		diff := math.Abs(float64(query.GrossRegisteredTonnage - index.GrossRegisteredTonnage))
		s := vesselTonnageScore(diff)
		scores = append(scores, fieldScore{s, 1.0})
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	var totalScore, totalWeight float64
	for _, fs := range scores {
		totalScore += fs.score * fs.weight
		totalWeight += fs.weight
	}
	avgScore := totalScore / totalWeight

	return avgScore, avgScore > 0.5, fieldsCompared
}

// Helper for vessel tonnage diffs
func vesselTonnageScore(diff float64) float64 {
	switch {
	case diff == 0:
		return 1.0
	case diff < 100:
		return 0.8
	case diff < 500:
		return 0.5
	default:
		return 0.0
	}
}

// -------------------------------
// Aircraft-Specific Fields
// -------------------------------
func compareAircraftFields(query *Aircraft, index *Aircraft) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	var scores []float64
	fieldsCompared := 0

	// ICAO
	if query.ICAOCode != "" {
		fieldsCompared++
		if strings.EqualFold(query.ICAOCode, index.ICAOCode) {
			scores = append(scores, 1.0)
		} else {
			scores = append(scores, 0.0)
		}
	}
	// Model
	if query.Model != "" {
		fieldsCompared++
		if strings.EqualFold(query.Model, index.Model) {
			scores = append(scores, 1.0)
		} else {
			// fuzzy
			qTerms := strings.Fields(strings.ToLower(query.Model))
			modelScore := stringscore.BestPairsJaroWinkler(qTerms, strings.ToLower(index.Model))
			scores = append(scores, modelScore)
		}
	}
	// Flag
	if query.Flag != "" {
		fieldsCompared++
		if strings.EqualFold(query.Flag, index.Flag) {
			scores = append(scores, 1.0)
		} else {
			scores = append(scores, 0.0)
		}
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	sum := 0.0
	for _, s := range scores {
		sum += s
	}
	avg := sum / float64(len(scores))

	return avg, avg > 0.5, fieldsCompared
}

// -------------------------------
// Business-Specific Fields
// -------------------------------
// compareBusinessFields compares fields for the Business entity
func compareBusinessFields(query *Business, index *Business) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	// We'll collect sub-scores with weights
	type fieldScore struct {
		score  float64
		weight float64
	}
	var scores []fieldScore
	fieldsCompared := 0

	fmt.Println("compareBusinessFields")

	// 1) Primary Name check (fuzzy or exact)
	if query.Name != "" {
		fieldsCompared++
		if strings.EqualFold(query.Name, index.Name) {
			// exact match
			scores = append(scores, fieldScore{score: 1.0, weight: 4.0})
		} else {
			// fuzzy match
			qTerms := strings.Fields(strings.ToLower(query.Name))
			iName := strings.ToLower(index.Name)
			nameScore := stringscore.BestPairsJaroWinkler(qTerms, iName)
			scores = append(scores, fieldScore{score: nameScore, weight: 4.0})
		}
	}
	fmt.Printf(".Name  scores=%v  fieldsCompared=%v\n", scores, fieldsCompared)

	// 2) AltNames check
	// If the query has alt names, let's see if any overlap. Or, if the index has alt names,
	// we can see if the query.Name matches them.
	// Typically you'd do something like your `Person.AltNames` logic. For simplicity:
	if len(query.AltNames) > 0 && len(index.AltNames) > 0 {
		fieldsCompared++
		bestAltScore := 0.0
		for _, qAlt := range query.AltNames {
			for _, iAlt := range index.AltNames {
				altScore := stringscore.BestPairsJaroWinkler(
					strings.Fields(strings.ToLower(qAlt)),
					strings.ToLower(iAlt),
				)
				if altScore > bestAltScore {
					bestAltScore = altScore
				}
			}
		}
		// Weight alt names a bit lower than primary name
		scores = append(scores, fieldScore{score: bestAltScore, weight: 2.0})
	}
	fmt.Printf(".AltName  scores=%v  fieldsCompared=%v\n", scores, fieldsCompared)

	// 3) Created date
	if query.Created != nil && index.Created != nil {
		fieldsCompared++
		if query.Created.Equal(*index.Created) {
			scores = append(scores, fieldScore{score: 1.0, weight: 1.0})
		} else {
			// partial credit if close
			diffDays := math.Abs(query.Created.Sub(*index.Created).Hours() / 24)
			switch {
			case diffDays <= 1:
				scores = append(scores, fieldScore{score: 0.9, weight: 1.0})
			case diffDays <= 7:
				scores = append(scores, fieldScore{score: 0.7, weight: 1.0})
			default:
				scores = append(scores, fieldScore{score: 0.0, weight: 1.0})
			}
		}
	}
	fmt.Printf(".Created  scores=%v  fieldsCompared=%v\n", scores, fieldsCompared)

	// 4) Dissolved date
	if query.Dissolved != nil && index.Dissolved != nil {
		fieldsCompared++
		if query.Dissolved.Equal(*index.Dissolved) {
			scores = append(scores, fieldScore{score: 1.0, weight: 1.0})
		} else {
			// partial logic if you want to consider near-dates as partial matches
			diffDays := math.Abs(query.Dissolved.Sub(*index.Dissolved).Hours() / 24)
			switch {
			case diffDays <= 1:
				scores = append(scores, fieldScore{score: 0.9, weight: 1.0})
			case diffDays <= 7:
				scores = append(scores, fieldScore{score: 0.7, weight: 1.0})
			default:
				scores = append(scores, fieldScore{score: 0.0, weight: 1.0})
			}
		}
	}
	fmt.Printf(".Dissolved  scores=%v  fieldsCompared=%v\n", scores, fieldsCompared)

	// 5) Identifiers
	// If you have multiple IDs in each, you might do a best match approach.
	// For each query.Identifier, find best match in index.Identifier.
	// Possibly weigh "Tax ID" or other critical IDs more heavily.
	if len(query.Identifier) > 0 && len(index.Identifier) > 0 {
		fieldsCompared++
		bestIDScore := 0.0
		for _, qID := range query.Identifier {
			for _, iID := range index.Identifier {
				// Example logic: exact match of "Identifier" + Country -> 1.0
				// partial or mismatch -> 0.
				// Could also do fuzzy or partial logic for qID.Identifier vs. iID.Identifier
				if strings.EqualFold(qID.Identifier, iID.Identifier) &&
					strings.EqualFold(qID.Country, iID.Country) &&
					strings.EqualFold(qID.Name, iID.Name) {
					// perfect
					bestIDScore = 1.0
					break
				} else if strings.EqualFold(qID.Identifier, iID.Identifier) {
					// partial
					if bestIDScore < 0.8 {
						bestIDScore = 0.8
					}
				} else {
					// could do fuzzy on qID.Identifier vs. iID.Identifier if you want
				}
			}
			if bestIDScore == 1.0 {
				break
			}
		}
		// Weight ID matches strongly
		scores = append(scores, fieldScore{score: bestIDScore, weight: 5.0})
	}
	fmt.Printf(".Identifier  scores=%v  fieldsCompared=%v\n", scores, fieldsCompared)

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	// Weighted average
	var totalScore, totalWeight float64
	for _, fs := range scores {
		totalScore += fs.score * fs.weight
		totalWeight += fs.weight
	}
	avgScore := totalScore / totalWeight

	// We'll say it's "matched" if > 0.5 on average // TODO(adam): why so low?
	return avgScore, avgScore > 0.5, fieldsCompared
}

// -------------------------------
// Organization-Specific Fields
// -------------------------------
func compareOrganizationFields(query *Organization, index *Organization) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	fieldsCompared := 0
	scores := make([]float64, 0)

	// Created date
	if query.Created != nil && index.Created != nil {
		fieldsCompared++
		if query.Created.Equal(*index.Created) {
			scores = append(scores, 1.0)
		} else {
			diff := math.Abs(query.Created.Sub(*index.Created).Hours() / 24)
			switch {
			case diff <= 1:
				scores = append(scores, 0.9)
			case diff <= 7:
				scores = append(scores, 0.7)
			default:
				scores = append(scores, 0.0)
			}
		}
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	sum := 0.0
	for _, s := range scores {
		sum += s
	}
	avg := sum / float64(len(scores))

	return avg, avg > 0.5, fieldsCompared
}

// -------------------------------
// Supporting Info (addresses, etc.)
// -------------------------------
func compareSupportingInfo[Q any, I any](query Entity[Q], index Entity[I], weight float64) scorePiece {
	var pieces []float64
	fieldsCompared := 0

	// Compare addresses
	if len(query.Addresses) > 0 && len(index.Addresses) > 0 {
		bestAddress := 0.0
		fieldsCompared++
		for _, qAddr := range query.Addresses {
			for _, iAddr := range index.Addresses {
				addrScore := compareAddress(qAddr, iAddr)
				if addrScore > bestAddress {
					bestAddress = addrScore
				}
			}
		}
		pieces = append(pieces, bestAddress)
	}

	// Compare sanctions programs
	if query.SanctionsInfo != nil && index.SanctionsInfo != nil {
		fieldsCompared++
		programScore := compareSanctionsPrograms(query.SanctionsInfo, index.SanctionsInfo)
		pieces = append(pieces, programScore)
	}

	// Compare crypto addresses (exact matches only)
	if len(query.CryptoAddresses) > 0 && len(index.CryptoAddresses) > 0 {
		fieldsCompared++
		matches := 0
		for _, qCA := range query.CryptoAddresses {
			for _, iCA := range index.CryptoAddresses {
				if strings.EqualFold(qCA.Currency, iCA.Currency) &&
					strings.EqualFold(qCA.Address, iCA.Address) {
					matches++
				}
			}
		}
		score := float64(matches) / float64(len(query.CryptoAddresses))
		pieces = append(pieces, score)
	}

	if len(pieces) == 0 {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "supporting"}
	}

	// Average of these pieces
	sum := 0.0
	for _, s := range pieces {
		sum += s
	}
	avgScore := sum / float64(len(pieces))

	return scorePiece{
		score:          avgScore,
		weight:         weight,
		matched:        avgScore > 0.5,
		required:       false,
		exact:          avgScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "supporting",
	}
}

// -------------------------------
// Address comparison
// -------------------------------
func compareAddress(query Address, index Address) float64 {
	var (
		pieces  []float64
		weights []float64
	)

	// Line1
	if query.Line1 != "" {
		qTerms := strings.Fields(query.Line1)
		score := stringscore.BestPairsJaroWinkler(qTerms, index.Line1)
		pieces = append(pieces, score)
		weights = append(weights, 3.0)
	}
	// Line2
	if query.Line2 != "" {
		qTerms := strings.Fields(query.Line2)
		score := stringscore.BestPairsJaroWinkler(qTerms, index.Line2)
		pieces = append(pieces, score)
		weights = append(weights, 1.0)
	}
	// City
	if query.City != "" {
		qTerms := strings.Fields(query.City)
		score := stringscore.BestPairsJaroWinkler(qTerms, index.City)
		pieces = append(pieces, score)
		weights = append(weights, 2.0)
	}
	// State (exact)
	if query.State != "" {
		if strings.EqualFold(query.State, index.State) {
			pieces = append(pieces, 1.0)
		} else {
			pieces = append(pieces, 0.0)
		}
		weights = append(weights, 1.0)
	}
	// Postal code (exact)
	if query.PostalCode != "" {
		if strings.EqualFold(query.PostalCode, index.PostalCode) {
			pieces = append(pieces, 1.0)
		} else {
			pieces = append(pieces, 0.0)
		}
		weights = append(weights, 1.5)
	}
	// Country (exact)
	if query.Country != "" {
		if strings.EqualFold(query.Country, index.Country) {
			pieces = append(pieces, 1.0)
		} else {
			pieces = append(pieces, 0.0)
		}
		weights = append(weights, 2.0)
	}

	if len(pieces) == 0 {
		return 0
	}

	var totalScore, totalWeight float64
	for i := range pieces {
		totalScore += pieces[i] * weights[i]
		totalWeight += weights[i]
	}
	return totalScore / totalWeight
}

func compareSanctionsPrograms(query *SanctionsInfo, index *SanctionsInfo) float64 {
	if query == nil || index == nil {
		return 0
	}
	if len(query.Programs) == 0 {
		return 0
	}

	matches := 0
	for _, qProgram := range query.Programs {
		for _, iProgram := range index.Programs {
			if strings.EqualFold(qProgram, iProgram) {
				matches++
				break
			}
		}
	}

	score := float64(matches) / float64(len(query.Programs))

	// Adjust for mismatch on "secondary"
	if query.Secondary != index.Secondary {
		score *= 0.8
	}
	return score
}

// -------------------------------
// Coverage Logic
// -------------------------------

// countIndexUniqueFields only counts fields relevant to the entity's type.
// This prevents penalizing a Person for Vessel fields, etc.
func countIndexUniqueFields[I any](index Entity[I]) int {
	count := 0

	switch index.Type {
	case EntityVessel:
		if index.Vessel != nil {
			if index.Vessel.IMONumber != "" {
				count++
			}
			if index.Vessel.CallSign != "" {
				count++
			}
			if index.Vessel.MMSI != "" {
				count++
			}
			if index.Vessel.Owner != "" {
				count++
			}
			// If you want to treat name as part of "unique fields," do that outside or here
		}
	case EntityPerson:
		if index.Person != nil {
			if index.Person.BirthDate != nil {
				count++
			}
			if index.Person.Gender != "" {
				count++
			}
			if len(index.Person.Titles) > 0 {
				count++
			}
		}
	case EntityAircraft:
		if index.Aircraft != nil {
			if index.Aircraft.ICAOCode != "" {
				count++
			}
			if index.Aircraft.Model != "" {
				count++
			}
			if index.Aircraft.Flag != "" {
				count++
			}
		}
	case EntityBusiness:
		if index.Business != nil {
			// If there's a Name
			if strings.TrimSpace(index.Business.Name) != "" {
				count++
			}
			// If there's at least one alt name
			if len(index.Business.AltNames) > 0 {
				count++
			}
			// If there's a created date
			if index.Business.Created != nil {
				count++
			}
			// If there's a dissolved date
			if index.Business.Dissolved != nil {
				count++
			}
			// If there's at least one Identifier
			if len(index.Business.Identifier) > 0 {
				count++
			}
		}
	case EntityOrganization:
		if index.Organization != nil {
			if index.Organization.Created != nil {
				count++
			}
		}
	}

	// Regardless of type, if there's a non-blank Name, count it
	if strings.TrimSpace(index.Name) != "" {
		count++
	}

	// You could also count addresses, sanctions, etc. if relevant:
	// if len(index.Addresses) > 0 { count++ }
	// if index.SanctionsInfo != nil { count++ }
	// etc.

	return count
}

const (
	debugFinalScores = true
)

// calculateFinalScore applies coverage logic and final adjustments.
func calculateFinalScore[I any](pieces []scorePiece, index Entity[I]) float64 {
	if len(pieces) == 0 {
		return 0
	}

	var (
		totalScore    float64
		totalWeight   float64
		hasExactMatch bool
		hasNameMatch  bool
	)

	if debugFinalScores {
		defer debug("\n")
	}

	// Sum up the piece scores
	for _, piece := range pieces {
		if debugFinalScores {
			if debugFinalScores {
				debug("%#v\n", piece)
			}
		}

		// Skip zero-weight pieces entirely
		if piece.weight <= 0 {
			continue
		}

		// If "entity" piece has score=0 but fieldsCompared=1, that indicates a type mismatch => overall 0
		// if piece.pieceType == "entity" && piece.fieldsCompared == 1 && piece.score == 0 {
		// 	debug("entity - mismatch")
		// 	return 0
		// }

		// Only accumulate if we actually compared some fields
		if piece.fieldsCompared > 0 {
			totalScore += piece.score * piece.weight
			totalWeight += piece.weight

			if piece.exact {
				hasExactMatch = true
			}
			// If the piece is "required" and "matched," track if it's the name
			if piece.required && piece.matched && piece.pieceType == "name" {
				hasNameMatch = true
			}
		}
	}

	if totalWeight == 0 {
		return 0
	}

	baseScore := totalScore / totalWeight
	if debugFinalScores {
		debug("baseScore=%.4f  ", baseScore)
	}

	// Coverage check: only count fields relevant to the index type
	coveragePenalty := 1.0
	indexUniqueCount := countIndexUniqueFields(index)
	fieldsCompared := 0
	for _, p := range pieces {
		fieldsCompared += p.fieldsCompared
	}
	if debugFinalScores {
		debug("fieldsCompared=%d  ", fieldsCompared)
	}

	if indexUniqueCount > 0 {
		coverage := float64(fieldsCompared) / float64(indexUniqueCount)

		// If coverage is very low (< 0.5) but the base score is high, reduce a bit
		if coverage < 0.5 && baseScore > 0.6 {
			coveragePenalty = 0.9
		}
	}

	finalScore := baseScore * coveragePenalty
	if debugFinalScores {
		debug("coveragePenalty=%.2f  ", coveragePenalty)
	}

	// Perfect match boost: only if coverage wasn't penalized
	if hasExactMatch && hasNameMatch && finalScore > 0.9 && coveragePenalty == 1.0 {
		if debugFinalScores {
			debug("PERFECT MATCH BOOST  ")
		}

		finalScore = math.Min(1.0, finalScore*1.15)
	}
	if debugFinalScores {
		debug("finalScore=%.2f", finalScore)
	}

	return finalScore
}

func debug(pattern string, args ...any) {
	fmt.Printf(pattern, args...) //nolint:forbidigo
}
