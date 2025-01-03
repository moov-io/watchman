package search

import (
	"math"
	"strings"
	"time"
)

func Similarity[Q any, I any](query Entity[Q], index Entity[I]) float64 {
	var parts []partial

	// 1) Compare top-level entity fields
	score, weight := compareStringField(query.Name, index.Name, 2.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	score, weight = compareStringField(string(query.Type), string(index.Type), 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	score, weight = compareStringField(string(query.Source), string(index.Source), 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	score, weight = compareStringField(query.SourceID, index.SourceID, 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	// Titles (slice of strings)
	score, weight = compareStringSlice(query.Titles, index.Titles, 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	// Person
	if query.Person != nil && index.Person != nil {
		// Name
		score, weight = compareStringField(query.Person.Name, index.Person.Name, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// Gender
		score, weight = compareStringField(string(query.Person.Gender), string(index.Person.Gender), 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// BirthDate
		score, weight = compareDateField(query.Person.BirthDate, index.Person.BirthDate, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// DeathDate
		score, weight = compareDateField(query.Person.DeathDate, index.Person.DeathDate, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		// AltNames
		score, weight = compareStringSlice(query.Person.AltNames, index.Person.AltNames, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		// Titles
		score, weight = compareStringSlice(query.Person.Titles, index.Person.Titles, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// GovernmentIDs (naive)
		if len(query.Person.GovernmentIDs) > 0 && len(index.Person.GovernmentIDs) > 0 {
			var qIDs, iIDs []string
			for _, gid := range query.Person.GovernmentIDs {
				qIDs = append(qIDs, string(gid.Type), gid.Country, gid.Identifier)
			}
			for _, gid := range index.Person.GovernmentIDs {
				iIDs = append(iIDs, string(gid.Type), gid.Country, gid.Identifier)
			}
			score, weight = compareStringSlice(qIDs, iIDs, 2.0)
			parts = append(parts, partial{Score: score, Weight: weight})
		}
	}

	// Business
	if query.Business != nil && index.Business != nil {
		score, weight = compareStringField(query.Business.Name, index.Business.Name, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Business.Created, index.Business.Created, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Business.Dissolved, index.Business.Dissolved, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// Compare Identifiers
		if len(query.Business.Identifier) > 0 && len(index.Business.Identifier) > 0 {
			var qIDs, iIDs []string
			for _, id := range query.Business.Identifier {
				qIDs = append(qIDs, id.Name, id.Country, id.Identifier)
			}
			for _, id := range index.Business.Identifier {
				iIDs = append(iIDs, id.Name, id.Country, id.Identifier)
			}
			score, weight = compareStringSlice(qIDs, iIDs, 2.0)
			parts = append(parts, partial{Score: score, Weight: weight})
		}
	}

	// Organization
	if query.Organization != nil && index.Organization != nil {
		score, weight = compareStringField(query.Organization.Name, index.Organization.Name, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Organization.Created, index.Organization.Created, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Organization.Dissolved, index.Organization.Dissolved, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		if len(query.Organization.Identifier) > 0 && len(index.Organization.Identifier) > 0 {
			var qIDs, iIDs []string
			for _, id := range query.Organization.Identifier {
				qIDs = append(qIDs, id.Name, id.Country, id.Identifier)
			}
			for _, id := range index.Organization.Identifier {
				iIDs = append(iIDs, id.Name, id.Country, id.Identifier)
			}
			score, weight = compareStringSlice(qIDs, iIDs, 2.0)
			parts = append(parts, partial{Score: score, Weight: weight})
		}
	}

	// Aircraft
	if query.Aircraft != nil && index.Aircraft != nil {
		score, weight = compareStringField(query.Aircraft.Name, index.Aircraft.Name, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(string(query.Aircraft.Type), string(index.Aircraft.Type), 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Aircraft.Flag, index.Aircraft.Flag, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Aircraft.Built, index.Aircraft.Built, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Aircraft.ICAOCode, index.Aircraft.ICAOCode, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Aircraft.Model, index.Aircraft.Model, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Aircraft.SerialNumber, index.Aircraft.SerialNumber, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
	}

	// Vessel
	if query.Vessel != nil && index.Vessel != nil {
		score, weight = compareStringField(query.Vessel.Name, index.Vessel.Name, 2.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Vessel.IMONumber, index.Vessel.IMONumber, 1.5)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(string(query.Vessel.Type), string(index.Vessel.Type), 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Vessel.Flag, index.Vessel.Flag, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareDateField(query.Vessel.Built, index.Vessel.Built, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Vessel.Model, index.Vessel.Model, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// Tonnage
		if query.Vessel.Tonnage > 0 && index.Vessel.Tonnage > 0 {
			diff := math.Abs(float64(query.Vessel.Tonnage) - float64(index.Vessel.Tonnage))
			var matchScore float64
			if diff == 0 {
				matchScore = 1.0
			} else if diff < 500 {
				matchScore = 0.5
			} else {
				matchScore = 0
			}
			parts = append(parts, partial{Score: matchScore, Weight: 1.0})
		}
		score, weight = compareStringField(query.Vessel.MMSI, index.Vessel.MMSI, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		score, weight = compareStringField(query.Vessel.CallSign, index.Vessel.CallSign, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
		// GrossRegisteredTonnage
		if query.Vessel.GrossRegisteredTonnage > 0 && index.Vessel.GrossRegisteredTonnage > 0 {
			diff := math.Abs(float64(query.Vessel.GrossRegisteredTonnage) -
				float64(index.Vessel.GrossRegisteredTonnage))
			var matchScore float64
			if diff == 0 {
				matchScore = 1.0
			} else if diff < 500 {
				matchScore = 0.5
			} else {
				matchScore = 0
			}
			parts = append(parts, partial{Score: matchScore, Weight: 1.0})
		}
		score, weight = compareStringField(query.Vessel.Owner, index.Vessel.Owner, 1.0)
		parts = append(parts, partial{Score: score, Weight: weight})
	}

	// CryptoAddresses
	score, weight = compareCryptoAddresses(query.CryptoAddresses, index.CryptoAddresses, 1.5)
	parts = append(parts, partial{Score: score, Weight: weight})

	// Addresses
	score, weight = compareAddresses(query.Addresses, index.Addresses, 1.5)
	parts = append(parts, partial{Score: score, Weight: weight})

	// Affiliations
	score, weight = compareAffiliations(query.Affiliations, index.Affiliations, 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	// SanctionsInfo
	score, weight = compareSanctionsInfo(query.SanctionsInfo, index.SanctionsInfo, 1.5)
	parts = append(parts, partial{Score: score, Weight: weight})

	// HistoricalInfo
	score, weight = compareHistoricalInfo(query.HistoricalInfo, index.HistoricalInfo, 1.0)
	parts = append(parts, partial{Score: score, Weight: weight})

	// SourceData (T) fields are not included in the scoring

	return combineScores(parts)
}

// jaroWinklerDistance is a placeholder for your advanced string comparison.
func jaroWinklerDistance(a, b string) float64 {
	a = strings.TrimSpace(strings.ToLower(a))
	b = strings.TrimSpace(strings.ToLower(b))
	if a == "" || b == "" {
		return 0.0
	}
	if a == b {
		return 1.0
	}
	// Replace with real logic
	return 0.5
}

// compareStringField returns (score, weight).
// If the query is empty, the field is skipped (no penalty, no weight).
func compareStringField(queryVal, indexVal string, weight float64) (float64, float64) {
	q := strings.TrimSpace(queryVal)
	i := strings.TrimSpace(indexVal)

	if q == "" {
		return 0, 0
	}
	if i == "" {
		return 0, weight
	}
	dist := jaroWinklerDistance(q, i)
	return dist * weight, weight
}

// compareDateField treats nil query as "skip," nil index as mismatch, and
// otherwise uses a simplistic "within 1 year => full match, within 5 => partial" strategy.
func compareDateField(queryVal, indexVal *time.Time, weight float64) (float64, float64) {
	if queryVal == nil {
		return 0, 0
	}
	if indexVal == nil {
		return 0, weight
	}
	diffYears := math.Abs(queryVal.Sub(*indexVal).Hours() / 24 / 365)
	switch {
	case diffYears < 1:
		return weight, weight
	case diffYears < 5:
		return 0.5 * weight, weight
	default:
		return 0, weight
	}
}

// compareStringSlice does a naive match by concatenating slices and comparing as one big string.
// In real systems, you might do a best-match approach or measure how many elements overlap, etc.
func compareStringSlice(queryVals, indexVals []string, weight float64) (float64, float64) {
	if len(queryVals) == 0 {
		return 0, 0
	}
	if len(indexVals) == 0 {
		return 0, weight
	}
	query := strings.Join(queryVals, " ")
	index := strings.Join(indexVals, " ")
	return compareStringField(query, index, weight)
}

// compareAddresses is a placeholder. In reality, you'd want something more sophisticated
// (e.g., best-match across addresses, geospatial closeness, etc.).
func compareAddresses(qAddrs, iAddrs []Address, weight float64) (float64, float64) {
	if len(qAddrs) == 0 {
		return 0, 0
	}
	if len(iAddrs) == 0 {
		return 0, weight
	}
	// Naive approach: just compare the first address line1, line2, city, etc. as a concatenated string
	var qParts, iParts []string
	for _, a := range qAddrs {
		qParts = append(qParts, a.Line1, a.Line2, a.City, a.State, a.PostalCode, a.Country)
	}
	for _, a := range iAddrs {
		iParts = append(iParts, a.Line1, a.Line2, a.City, a.State, a.PostalCode, a.Country)
	}
	query := strings.Join(qParts, " ")
	index := strings.Join(iParts, " ")
	return compareStringField(query, index, weight)
}

// compareCryptoAddresses is a placeholder that just compares them all as one big string.
func compareCryptoAddresses(qAddrs, iAddrs []CryptoAddress, weight float64) (float64, float64) {
	if len(qAddrs) == 0 {
		return 0, 0
	}
	if len(iAddrs) == 0 {
		return 0, weight
	}
	var qParts, iParts []string
	for _, ca := range qAddrs {
		qParts = append(qParts, ca.Currency, ca.Address)
	}
	for _, ca := range iAddrs {
		iParts = append(iParts, ca.Currency, ca.Address)
	}
	query := strings.Join(qParts, " ")
	index := strings.Join(iParts, " ")
	return compareStringField(query, index, weight)
}

// compareAffiliations is another naive approach.
// You could do more advanced logic to match entity names, types, etc.
func compareAffiliations(qAffs, iAffs []Affiliation, weight float64) (float64, float64) {
	if len(qAffs) == 0 {
		return 0, 0
	}
	if len(iAffs) == 0 {
		return 0, weight
	}
	// Just combine them all
	var qParts, iParts []string
	for _, aff := range qAffs {
		qParts = append(qParts, aff.EntityName, aff.Type, aff.Details)
	}
	for _, aff := range iAffs {
		iParts = append(iParts, aff.EntityName, aff.Type, aff.Details)
	}
	query := strings.Join(qParts, " ")
	index := strings.Join(iParts, " ")
	return compareStringField(query, index, weight)
}

// compareSanctionsInfo is naive. Potentially you'd do fuzzy set matching of programs, etc.
func compareSanctionsInfo(qInfo, iInfo *SanctionsInfo, weight float64) (float64, float64) {
	if qInfo == nil {
		return 0, 0
	}
	if iInfo == nil {
		return 0, weight
	}
	// Combine programs and description
	query := strings.Join(qInfo.Programs, " ") + " " + qInfo.Description
	index := strings.Join(iInfo.Programs, " ") + " " + iInfo.Description
	score, w := compareStringField(query, index, weight)
	// If one is "secondary" and the other isn't, reduce score
	if qInfo.Secondary != iInfo.Secondary && score > 0 {
		score *= 0.5
	}
	return score, w
}

// compareHistoricalInfo is naive. You might want date checks, type matching, etc.
func compareHistoricalInfo(qHist, iHist []HistoricalInfo, weight float64) (float64, float64) {
	if len(qHist) == 0 {
		return 0, 0
	}
	if len(iHist) == 0 {
		return 0, weight
	}
	var qParts, iParts []string
	for _, h := range qHist {
		qParts = append(qParts, h.Type, h.Value)
		// If you want to compare the date, you'd do it similarly to compareDateField
	}
	for _, h := range iHist {
		iParts = append(iParts, h.Type, h.Value)
	}
	query := strings.Join(qParts, " ")
	index := strings.Join(iParts, " ")
	return compareStringField(query, index, weight)
}

type partial struct {
	Score  float64
	Weight float64
}

// combineScores sums partials into a final ratio in [0..1].
func combineScores(partials []partial) float64 {
	var totalScore, totalWeight float64
	for _, p := range partials {
		totalScore += p.Score
		totalWeight += p.Weight
	}
	if totalWeight == 0 {
		return 0
	}
	ratio := totalScore / totalWeight
	if ratio < 0 {
		return 0
	} else if ratio > 1 {
		return 1
	}
	return ratio
}
