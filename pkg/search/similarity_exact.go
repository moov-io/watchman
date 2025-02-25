package search

import (
	"io"
	"strings"
)

// compareExactIdentifiers covers exact matches for identifiers across all entity types
func compareExactIdentifiers[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	// If types don't match, return early
	if query.Type != index.Type {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 1,
			PieceType:      "identifiers",
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
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
	}
}

func normalizeIdentifier(id string) string {
	return strings.ReplaceAll(id, "-", "")
}

// comparePersonExactIDs checks exact matches for Person-specific identifiers
func comparePersonExactIDs(w io.Writer, query *Person, index *Person, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       fieldsCompared > 0,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "identifiers",
	}
}

// compareBusinessExactIDs checks exact matches for Business-specific identifiers
func compareBusinessExactIDs(w io.Writer, query *Business, index *Business, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       fieldsCompared > 0,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "identifiers",
	}
}

// compareOrgExactIDs checks exact matches for Organization-specific identifiers
func compareOrgExactIDs(w io.Writer, query *Organization, index *Organization, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       fieldsCompared > 0,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "identifiers",
	}
}

// compareVesselExactIDs checks exact matches for Vessel-specific identifiers
func compareVesselExactIDs(w io.Writer, query *Vessel, index *Vessel, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       fieldsCompared > 0,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "identifiers",
	}
}

// compareAircraftExactIDs checks exact matches for Aircraft-specific identifiers
func compareAircraftExactIDs(w io.Writer, query *Aircraft, index *Aircraft, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: 0, FieldsCompared: 0, PieceType: "identifiers"}
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       fieldsCompared > 0,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "identifiers",
	}
}

func compareExactCryptoAddresses[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	fieldsCompared := 0
	hasMatch := false
	score := 0.0

	// Early return if either list is empty
	if len(query.CryptoAddresses) == 0 || len(index.CryptoAddresses) == 0 {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "crypto-exact",
		}
	}

	fieldsCompared++

	// First try exact matches (both currency and address)
	for q := range query.CryptoAddresses {
		if query.CryptoAddresses[q].Address == "" {
			continue // Skip empty addresses
		}

		for i := range index.CryptoAddresses {
			// both have currency specified - need both to match
			if query.CryptoAddresses[q].Currency != "" && index.CryptoAddresses[i].Currency != "" {
				// Check currency
				if strings.EqualFold(query.CryptoAddresses[q].Currency, index.CryptoAddresses[i].Currency) {
					goto checkAddresses
				} else {
					continue
				}
			}
			// Check addresses
		checkAddresses:
			if strings.EqualFold(query.CryptoAddresses[q].Address, index.CryptoAddresses[i].Address) {
				score = 1.0
				hasMatch = true
				goto Done
			}
		}
	}

Done:
	return ScorePiece{
		Score:          score,
		Weight:         weight,
		Matched:        hasMatch,
		Required:       false,
		Exact:          score > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "crypto-exact",
	}
}

// compareExactGovernmentIDs compares government IDs across entity types
func compareExactGovernmentIDs[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	if query.Type != index.Type {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "gov-ids-exact",
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
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "gov-ids-exact",
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
	if !strings.EqualFold(queryID, indexID) {
		return idMatch{score: 0, found: false, exact: false}
	}

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

func comparePersonGovernmentIDs(query *Person, index *Person, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
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
	return ScorePiece{
		Score:          bestMatch.score,
		Weight:         weight,
		Matched:        bestMatch.found,
		Required:       false,
		Exact:          bestMatch.exact,
		FieldsCompared: fieldsCompared,
		PieceType:      "gov-ids-exact",
	}
}

func compareBusinessGovernmentIDs(query *Business, index *Business, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
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
	return ScorePiece{
		Score:          bestMatch.score,
		Weight:         weight,
		Matched:        bestMatch.found,
		Required:       false,
		Exact:          bestMatch.exact,
		FieldsCompared: fieldsCompared,
		PieceType:      "gov-ids-exact",
	}
}

func compareOrgGovernmentIDs(query *Organization, index *Organization, weight float64) ScorePiece {
	if query == nil || index == nil {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
	}

	qIDs := query.GovernmentIDs
	iIDs := index.GovernmentIDs

	if len(qIDs) == 0 || len(iIDs) == 0 {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "gov-ids-exact"}
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
	return ScorePiece{
		Score:          bestMatch.score,
		Weight:         weight,
		Matched:        bestMatch.found,
		Required:       false,
		Exact:          bestMatch.exact,
		FieldsCompared: fieldsCompared,
		PieceType:      "gov-ids-exact",
	}
}

func compareExactSourceID[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	// Early return if query has no source ID
	if query.SourceID == "" {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "source-id-exact",
		}
	}

	// Always count as field compared if query has a source ID
	fieldsCompared := 1

	// Handle case where index has no source ID
	if index.SourceID == "" {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: fieldsCompared,
			PieceType:      "source-id-exact",
		}
	}

	// Compare normalized source IDs
	hasMatch := strings.EqualFold(query.SourceID, index.SourceID)

	return ScorePiece{
		Score:          boolToScore(hasMatch),
		Weight:         weight,
		Matched:        hasMatch,
		Required:       false,
		Exact:          hasMatch, // If matched, it's exact by definition
		FieldsCompared: fieldsCompared,
		PieceType:      "source-id-exact",
	}
}

func compareExactSourceList[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	// Early return if query has no source
	if query.Source == "" {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "source-list",
		}
	}

	// Always count as field compared if query has a source
	fieldsCompared := 1

	// Handle case where index has no source
	if index.Source == "" {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: fieldsCompared,
			PieceType:      "source-list",
		}
	}

	// Compare normalized sources
	hasMatch := strings.EqualFold(string(query.Source), string(index.Source))

	return ScorePiece{
		Score:          boolToScore(hasMatch),
		Weight:         weight,
		Matched:        hasMatch,
		Required:       false,
		Exact:          hasMatch, // If matched, it's exact by definition
		FieldsCompared: fieldsCompared,
		PieceType:      "source-list",
	}
}

// contactFieldMatch handles matching logic for a single contact field type (email, phone, fax)
type contactFieldMatch struct {
	matches    int
	totalQuery int
	score      float64
}

func compareExactContactInfo[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	fieldsCompared := 0
	var matches []contactFieldMatch

	// Compare emails (exact match)
	if len(query.Contact.EmailAddresses) > 0 && len(index.Contact.EmailAddresses) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(
			query.Contact.EmailAddresses,
			index.Contact.EmailAddresses,
		))
	}

	// Compare phone numbers
	if len(query.PreparedFields.Contact.PhoneNumbers) > 0 && len(index.PreparedFields.Contact.PhoneNumbers) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(query.PreparedFields.Contact.PhoneNumbers, index.PreparedFields.Contact.PhoneNumbers))
	}

	// Compare fax numbers
	if len(query.PreparedFields.Contact.FaxNumbers) > 0 && len(index.PreparedFields.Contact.FaxNumbers) > 0 {
		fieldsCompared++
		matches = append(matches, compareContactField(query.PreparedFields.Contact.FaxNumbers, index.PreparedFields.Contact.FaxNumbers))
	}

	if fieldsCompared == 0 {
		return ScorePiece{
			Score:          0,
			Weight:         weight,
			Matched:        false,
			Required:       false,
			Exact:          false,
			FieldsCompared: 0,
			PieceType:      "contact-exact",
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

	return ScorePiece{
		Score:          finalScore,
		Weight:         weight,
		Matched:        totalMatches > 0,
		Required:       false,
		Exact:          finalScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "contact-exact",
	}
}

// compareContactField handles the comparison logic for a single type of contact field
func compareContactField(queryValues, indexValues []string) contactFieldMatch {
	matches := 0

	for q := range queryValues {
		for i := range indexValues {
			if strings.EqualFold(queryValues[q], indexValues[i]) {
				matches++
			}
		}
	}

	score := float64(matches) / float64(len(queryValues))

	return contactFieldMatch{
		matches:    matches,
		totalQuery: len(queryValues),
		score:      score,
	}
}
