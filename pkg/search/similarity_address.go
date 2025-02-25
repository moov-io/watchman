package search

import (
	"io"
	"strings"

	"github.com/moov-io/watchman/internal/stringscore"
)

const (
	// Field weights for addresses
	line1Weight   = 5.0 // Primary address line - most important
	line2Weight   = 2.0 // Secondary address info - less important
	cityWeight    = 4.0 // City - highly important for location
	stateWeight   = 2.0 // State - helps confirm location
	postalWeight  = 3.0 // Postal code - strong verification
	countryWeight = 4.0 // Country - critical for international
)

func compareAddresses[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
	fieldsCompared := 0
	var scores []float64

	if w != nil {
		debug(w, "address comparison details: query=%d  index=%d\n", len(query.Addresses), len(index.Addresses))
	}

	// Compare addresses
	if len(query.Addresses) > 0 && len(index.Addresses) > 0 {
		fieldsCompared++
		if score := findBestAddressMatch(w, query.Addresses, index.Addresses); score > 0 {
			scores = append(scores, score)
		}
	}

	if len(scores) == 0 {
		return ScorePiece{Score: 0, Weight: weight, FieldsCompared: 0, PieceType: "address"}
	}

	avgScore := calculateAverage(scores)
	return ScorePiece{
		Score:          avgScore,
		Weight:         weight,
		Matched:        avgScore > 0.5,
		Required:       false,
		Exact:          avgScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "address",
	}
}

func findBestAddressMatch(w io.Writer, queryAddrs, indexAddrs []Address) float64 {
	bestScore := 0.0
	for i, qa := range queryAddrs {
		for j, ia := range indexAddrs {
			if w != nil {
				debug(w, "Comparing Query Address %d with Index Address %d:\n", i+1, j+1)
			}
			if score := compareAddress(w, qa, ia); score > bestScore {
				bestScore = score
				if score > highConfidenceThreshold {
					if w != nil {
						debug(w, "Found high confidence match (%.2f), stopping search\n", score)
					}
					return score
				}
			}
		}
	}
	return bestScore
}

func compareAddress(w io.Writer, query, index Address) float64 {
	var totalScore, totalWeight float64

	// Compare line1 (highest weight)
	if query.Line1 != "" && index.Line1 != "" {
		similarity := stringscore.JaroWinkler(query.Line1, index.Line1)
		totalScore += similarity * line1Weight
		totalWeight += line1Weight
		if w != nil {
			debug(w, "  Line1: %.3f (weight: %.1f) [%s] vs [%s]\n",
				similarity, line1Weight, query.Line1, index.Line1)
		}
	}

	// Compare line2
	if query.Line2 != "" && index.Line2 != "" {
		similarity := stringscore.JaroWinkler(query.Line2, index.Line2)
		totalScore += similarity * line2Weight
		totalWeight += line2Weight
		if w != nil {
			debug(w, "  Line2: %.3f (weight: %.1f) [%s] vs [%s]\n",
				similarity, line2Weight, query.Line2, index.Line2)
		}
	}

	// Compare city
	if query.City != "" && index.City != "" {
		similarity := stringscore.JaroWinkler(query.City, index.City)
		totalScore += similarity * cityWeight
		totalWeight += cityWeight
		if w != nil {
			debug(w, "  City: %.3f (weight: %.1f) [%s] vs [%s]\n",
				similarity, cityWeight, query.City, index.City)
		}
	}

	// Compare state
	if query.State != "" && index.State != "" {
		match := strings.EqualFold(query.State, index.State)
		score := boolToScore(match)
		totalScore += score * stateWeight
		totalWeight += stateWeight
		if w != nil {
			debug(w, "  State: %.3f (weight: %.1f) [%s] vs [%s]\n",
				score, stateWeight, query.State, index.State)
		}
	}

	// Compare postal code
	if query.PostalCode != "" && index.PostalCode != "" {
		match := strings.EqualFold(query.PostalCode, index.PostalCode)
		score := boolToScore(match)
		totalScore += score * postalWeight
		totalWeight += postalWeight
		if w != nil {
			debug(w, "  Postal: %.3f (weight: %.1f) [%s] vs [%s]\n",
				score, postalWeight, query.PostalCode, index.PostalCode)
		}
	}

	// Compare country
	if query.Country != "" && index.Country != "" {
		// We assume the query and index Country fields are normalized before Similarity
		match := strings.EqualFold(query.Country, index.Country)
		score := boolToScore(match)
		totalScore += score * countryWeight
		totalWeight += countryWeight
		if w != nil {
			debug(w, "  Country: %.3f (weight: %.1f) [%s] vs [%s]\n",
				score, countryWeight, query.Country, index.Country)
		}
	}

	if totalWeight == 0 {
		if w != nil {
			debug(w, "  No fields compared\n")
		}
		return 0
	}

	finalScore := totalScore / totalWeight
	if w != nil {
		debug(w, "  Final Score: %.3f (total score: %.2f / total weight: %.2f)\n",
			finalScore, totalScore, totalWeight)
	}
	return finalScore
}
