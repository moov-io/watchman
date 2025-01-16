package search

import (
	"io"
	"strings"
	"unicode"
)

// compareExactIdentifiers covers exact matches for identifiers across all entity types
func compareExactIdentifiers[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	// If types don't match, return early
	if query.Type != index.Type {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 1,
			pieceType:      "identifiers",
		}
	}

	// Call appropriate helper based on entity type
	switch query.Type {
	case EntityPerson:
		return comparePersonExactIDs(w, query.Person, index.Person, weight)
	case EntityBusiness:
		return compareBusinessExactIDs(w, query.Business, index.Business, weight)
	case EntityOrganization:
		return compareOrgExactIDs(w, query.Organization, index.Organization, weight)
	case EntityVessel:
		return compareVesselExactIDs(w, query.Vessel, index.Vessel, weight)
	case EntityAircraft:
		return compareAircraftExactIDs(w, query.Aircraft, index.Aircraft, weight)
	default:
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}
}

func normalizeIdentifier(id string) string {
	return strings.ReplaceAll(id, "-", "")
}

// comparePersonExactIDs checks exact matches for Person-specific identifiers
func comparePersonExactIDs(w io.Writer, query *Person, index *Person, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}

	fieldsCompared := 0
	totalWeight := 0.0
	score := 0.0
	hasMatch := false

	// Government IDs (extremely high weight for exact matches)
	if len(query.GovernmentIDs) > 0 && len(index.GovernmentIDs) > 0 {
		fieldsCompared++
		totalWeight += 15.0
		for _, qID := range query.GovernmentIDs {
			for _, iID := range index.GovernmentIDs {
				if strings.EqualFold(string(qID.Type), string(iID.Type)) &&
					strings.EqualFold(qID.Country, iID.Country) &&
					strings.EqualFold(normalizeIdentifier(qID.Identifier), normalizeIdentifier(iID.Identifier)) {
					score += 15.0
					hasMatch = true
					goto GovIDDone // Break both loops on first match
				}
			}
		}
	}
GovIDDone:

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

// compareBusinessExactIDs checks exact matches for Business-specific identifiers
func compareBusinessExactIDs(w io.Writer, query *Business, index *Business, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}

	fieldsCompared := 0
	totalWeight := 0.0
	score := 0.0
	hasMatch := false

	// Business Registration/Tax IDs
	if len(query.GovernmentIDs) > 0 && len(index.GovernmentIDs) > 0 {
		fieldsCompared++
		totalWeight += 15.0
		for _, qID := range query.GovernmentIDs {
			for _, iID := range index.GovernmentIDs {
				// Exact match on all identifier fields
				if strings.EqualFold(string(qID.Type), string(iID.Type)) &&
					strings.EqualFold(qID.Country, iID.Country) &&
					strings.EqualFold(normalizeIdentifier(qID.Identifier), normalizeIdentifier(iID.Identifier)) {
					score += 15.0
					hasMatch = true
					goto IdentifierDone
				}
			}
		}
	}
IdentifierDone:

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

// compareOrgExactIDs checks exact matches for Organization-specific identifiers
func compareOrgExactIDs(w io.Writer, query *Organization, index *Organization, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}

	fieldsCompared := 0
	totalWeight := 0.0
	score := 0.0
	hasMatch := false

	// Organization Registration/Tax IDs
	if len(query.GovernmentIDs) > 0 && len(index.GovernmentIDs) > 0 {
		fieldsCompared++
		totalWeight += 15.0
		for _, qID := range query.GovernmentIDs {
			for _, iID := range index.GovernmentIDs {
				// Exact match on all identifier fields
				if strings.EqualFold(string(qID.Type), string(iID.Type)) &&
					strings.EqualFold(qID.Country, iID.Country) &&
					strings.EqualFold(normalizeIdentifier(qID.Identifier), normalizeIdentifier(iID.Identifier)) {
					score += 15.0
					hasMatch = true
					goto IdentifierDone
				}
			}
		}
	}
IdentifierDone:

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

// compareVesselExactIDs checks exact matches for Vessel-specific identifiers
func compareVesselExactIDs(w io.Writer, query *Vessel, index *Vessel, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}

	fieldsCompared := 0
	totalWeight := 0.0
	score := 0.0
	hasMatch := false

	// IMO Number (highest weight)
	if query.IMONumber != "" {
		fieldsCompared++
		totalWeight += 15.0
		if strings.EqualFold(query.IMONumber, index.IMONumber) {
			score += 15.0
			hasMatch = true
		}
	}

	// Call Sign
	if query.CallSign != "" {
		fieldsCompared++
		totalWeight += 12.0
		if strings.EqualFold(query.CallSign, index.CallSign) {
			score += 12.0
			hasMatch = true
		}
	}

	// MMSI (Maritime Mobile Service Identity)
	if query.MMSI != "" {
		fieldsCompared++
		totalWeight += 12.0
		if strings.EqualFold(query.MMSI, index.MMSI) {
			score += 12.0
			hasMatch = true
		}
	}

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

// compareAircraftExactIDs checks exact matches for Aircraft-specific identifiers
func compareAircraftExactIDs(w io.Writer, query *Aircraft, index *Aircraft, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: 0, fieldsCompared: 0, pieceType: "identifiers"}
	}

	fieldsCompared := 0
	totalWeight := 0.0
	score := 0.0
	hasMatch := false

	// Serial Number (highest weight)
	if query.SerialNumber != "" {
		fieldsCompared++
		totalWeight += 15.0
		if strings.EqualFold(query.SerialNumber, index.SerialNumber) {
			score += 15.0
			hasMatch = true
		}
	}

	// ICAO Code
	if query.ICAOCode != "" {
		fieldsCompared++
		totalWeight += 12.0
		if strings.EqualFold(query.ICAOCode, index.ICAOCode) {
			score += 12.0
			hasMatch = true
		}
	}

	finalScore := 0.0
	if totalWeight > 0 {
		finalScore = score / totalWeight
	}

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        hasMatch,
		required:       fieldsCompared > 0,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "identifiers",
	}
}

func compareExactCryptoAddresses[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	fieldsCompared := 0
	hasMatch := false
	score := 0.0

	qCAs := query.CryptoAddresses
	iCAs := index.CryptoAddresses

	// Early return if either list is empty
	if len(qCAs) == 0 || len(iCAs) == 0 {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "crypto-exact",
		}
	}

	fieldsCompared++

	// First try exact matches (both currency and address)
	for _, qCA := range qCAs {
		qAddr := strings.ToLower(strings.TrimSpace(qCA.Address))
		qCurr := strings.ToLower(strings.TrimSpace(qCA.Currency))

		if qAddr == "" {
			continue // Skip empty addresses
		}

		for _, iCA := range iCAs {
			iAddr := strings.ToLower(strings.TrimSpace(iCA.Address))
			iCurr := strings.ToLower(strings.TrimSpace(iCA.Currency))

			// Case 1: Both have currency specified - need both to match
			if qCurr != "" && iCurr != "" {
				if qCurr == iCurr && qAddr == iAddr {
					score = 1.0
					hasMatch = true
					goto Done
				}
			} else {
				// Case 2: At least one currency empty - match on address only
				if qAddr == iAddr {
					score = 1.0
					hasMatch = true
					goto Done
				}
			}
		}
	}

Done:
	return scorePiece{
		score:          score,
		weight:         weight,
		matched:        hasMatch,
		required:       false,
		exact:          score > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "crypto-exact",
	}
}

// compareExactGovernmentIDs compares government IDs across entity types
func compareExactGovernmentIDs[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	if query.Type != index.Type {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "gov-ids-exact",
		}
	}

	switch query.Type {
	case EntityPerson:
		return comparePersonGovernmentIDs(query.Person, index.Person, weight)
	case EntityBusiness:
		return compareBusinessGovernmentIDs(query.Business, index.Business, weight)
	case EntityOrganization:
		return compareOrgGovernmentIDs(query.Organization, index.Organization, weight)
	default:
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "gov-ids-exact",
		}
	}
}

// idMatch represents the result of comparing two identifiers
type idMatch struct {
	score      float64
	found      bool
	exact      bool
	hasCountry bool
}

// compareIdentifiers handles the core logic of comparing two identifier values
func compareIdentifiers(queryID, indexID string, queryCountry, indexCountry string) idMatch {
	// Early return if identifiers don't match
	if !strings.EqualFold(strings.TrimSpace(queryID), strings.TrimSpace(indexID)) {
		return idMatch{score: 0, found: false, exact: false}
	}

	// If we get here, identifiers match exactly
	queryCountry = strings.TrimSpace(queryCountry)
	indexCountry = strings.TrimSpace(indexCountry)

	// If neither has country, it's an exact match but flag no country
	if queryCountry == "" && indexCountry == "" {
		return idMatch{score: 1.0, found: true, exact: true, hasCountry: false}
	}

	// If only one has country, slight penalty
	if (queryCountry == "" && indexCountry != "") || (queryCountry != "" && indexCountry == "") {
		return idMatch{score: 0.9, found: true, exact: false, hasCountry: true}
	}

	// Both have country - check if they match
	if strings.EqualFold(queryCountry, indexCountry) {
		return idMatch{score: 1.0, found: true, exact: true, hasCountry: true}
	}

	// Countries don't match - significant penalty but still count as a match
	return idMatch{score: 0.7, found: true, exact: false, hasCountry: true}
}

func comparePersonGovernmentIDs(query *Person, index *Person, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	fieldsCompared := 1
	bestMatch := idMatch{score: 0}

	for _, qID := range qIDs {
		for _, iID := range iIDs {
			match := compareIdentifiers(qID.Identifier, iID.Identifier, qID.Country, iID.Country)
			if match.found && match.score > bestMatch.score {
				bestMatch = match
			}
			if bestMatch.exact {
				goto Done
			}
		}
	}

Done:
	return scorePiece{
		score:          bestMatch.score,
		weight:         weight,
		matched:        bestMatch.found,
		required:       false,
		exact:          bestMatch.exact,
		fieldsCompared: fieldsCompared,
		pieceType:      "gov-ids-exact",
	}
}

func compareBusinessGovernmentIDs(query *Business, index *Business, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	fieldsCompared := 1
	bestMatch := idMatch{score: 0}

	for _, qID := range qIDs {
		for _, iID := range iIDs {
			// For business, we'll check the identifier and country, ignoring name for now
			match := compareIdentifiers(qID.Identifier, iID.Identifier, qID.Country, iID.Country)
			if match.found && match.score > bestMatch.score {
				bestMatch = match
			}
			if bestMatch.exact {
				goto Done
			}
		}
	}

Done:
	return scorePiece{
		score:          bestMatch.score,
		weight:         weight,
		matched:        bestMatch.found,
		required:       false,
		exact:          bestMatch.exact,
		fieldsCompared: fieldsCompared,
		pieceType:      "gov-ids-exact",
	}
}

func compareOrgGovernmentIDs(query *Organization, index *Organization, weight float64) scorePiece {
	if query == nil || index == nil {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "gov-ids-exact"}
	}

	fieldsCompared := 1
	bestMatch := idMatch{score: 0}

	for _, qID := range qIDs {
		for _, iID := range iIDs {
			// For orgs, we'll check the identifier and country, ignoring name for now
			match := compareIdentifiers(qID.Identifier, iID.Identifier, qID.Country, iID.Country)
			if match.found && match.score > bestMatch.score {
				bestMatch = match
			}
			if bestMatch.exact {
				goto Done
			}
		}
	}

Done:
	return scorePiece{
		score:          bestMatch.score,
		weight:         weight,
		matched:        bestMatch.found,
		required:       false,
		exact:          bestMatch.exact,
		fieldsCompared: fieldsCompared,
		pieceType:      "gov-ids-exact",
	}
}

func compareExactSourceID[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	// Early return if query has no source ID
	if strings.TrimSpace(query.SourceID) == "" {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "source-id-exact",
		}
	}

	// Always count as field compared if query has a source ID
	fieldsCompared := 1

	// Handle case where index has no source ID
	if strings.TrimSpace(index.SourceID) == "" {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: fieldsCompared,
			pieceType:      "source-id-exact",
		}
	}

	// Compare normalized source IDs
	hasMatch := strings.EqualFold(
		strings.TrimSpace(query.SourceID),
		strings.TrimSpace(index.SourceID),
	)

	return scorePiece{
		score:          boolToScore(hasMatch),
		weight:         weight,
		matched:        hasMatch,
		required:       false,
		exact:          hasMatch, // If matched, it's exact by definition
		fieldsCompared: fieldsCompared,
		pieceType:      "source-id-exact",
	}
}

func compareExactSourceList[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	// Early return if query has no source
	if query.Source == "" {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "source-list",
		}
	}

	// Always count as field compared if query has a source
	fieldsCompared := 1

	// Handle case where index has no source
	if index.Source == "" {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: fieldsCompared,
			pieceType:      "source-list",
		}
	}

	// Compare normalized sources
	hasMatch := strings.EqualFold(
		string(query.Source),
		string(index.Source),
	)

	return scorePiece{
		score:          boolToScore(hasMatch),
		weight:         weight,
		matched:        hasMatch,
		required:       false,
		exact:          hasMatch, // If matched, it's exact by definition
		fieldsCompared: fieldsCompared,
		pieceType:      "source-list",
	}
}

// contactFieldMatch handles matching logic for a single contact field type (email, phone, fax)
type contactFieldMatch struct {
	matches    int
	totalQuery int
	score      float64
}

func compareExactContactInfo[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	fieldsCompared := 0
	var matches []contactFieldMatch

	// Compare emails (exact match)
	if len(query.Contact.EmailAddresses) > 0 && len(index.Contact.EmailAddresses) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(
			query.Contact.EmailAddresses,
			index.Contact.EmailAddresses,
			normalizeEmail,
		))
	}

	// Compare phone numbers (normalized)
	if len(query.Contact.PhoneNumbers) > 0 && len(index.Contact.PhoneNumbers) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(
			query.Contact.PhoneNumbers,
			index.Contact.PhoneNumbers,
			normalizePhoneNumber,
		))
	}

	// Compare fax numbers (normalized same as phones)
	if len(query.Contact.FaxNumbers) > 0 && len(index.Contact.FaxNumbers) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(
			query.Contact.FaxNumbers,
			index.Contact.FaxNumbers,
			normalizePhoneNumber,
		))
	}

	if fieldsCompared == 0 {
		return scorePiece{
			score:          0,
			weight:         weight,
			matched:        false,
			required:       false,
			exact:          false,
			fieldsCompared: 0,
			pieceType:      "contact-exact",
		}
	}

	// Calculate final scores
	totalScore := 0.0
	totalMatches := 0
	totalQueryItems := 0

	for _, m := range matches {
		totalScore += m.score
		totalMatches += m.matches
		totalQueryItems += m.totalQuery
	}

	finalScore := totalScore / float64(len(matches))

	return scorePiece{
		score:          finalScore,
		weight:         weight,
		matched:        totalMatches > 0,
		required:       false,
		exact:          finalScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "contact-exact",
	}
}

// compareContactField handles the comparison logic for a single type of contact field
func compareContactField(queryValues, indexValues []string, normalize func(string) string) contactFieldMatch {
	matches := 0

	// Create map of normalized index values for faster lookup
	indexMap := make(map[string]bool, len(indexValues))
	for _, iv := range indexValues {
		indexMap[normalize(iv)] = true
	}

	// Check each query value against the map
	for _, qv := range queryValues {
		if indexMap[normalize(qv)] {
			matches++
		}
	}

	score := float64(matches) / float64(len(queryValues))

	return contactFieldMatch{
		matches:    matches,
		totalQuery: len(queryValues),
		score:      score,
	}
}

// normalizeEmail normalizes email addresses for comparison
// TODO(adam):
func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// normalizePhoneNumber strips all non-numeric characters and normalizes phone numbers
// TODO(adam):
func normalizePhoneNumber(phone string) string {
	var normalized strings.Builder

	// Strip everything except digits and plus sign (for international prefix)
	for _, r := range phone {
		if unicode.IsDigit(r) || r == '+' {
			normalized.WriteRune(r)
		}
	}

	// Handle international format
	result := normalized.String()
	if strings.HasPrefix(result, "+") {
		return result // Keep international format as is
	}

	// If it's a 10-digit number without country code, keep as is
	if len(result) == 10 {
		return result
	}

	// If it has more digits but no plus, assume it includes country code
	if len(result) > 10 {
		return "+" + result
	}

	return result
}
