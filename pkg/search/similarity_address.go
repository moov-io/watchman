package search

import (
	"io"
	"strings"

	"github.com/moov-io/watchman/internal/stringscore"
)

const (
	// Field weights for addresses
	line1Weight   = 3.0 // Primary address line - most important
	line2Weight   = 1.0 // Secondary address info - less important
	cityWeight    = 2.0 // City - moderately important
	stateWeight   = 1.0 // State - helps confirm location
	postalWeight  = 1.5 // Postal code - good verification
	countryWeight = 2.0 // Country - important for international
)

func compareAddresses[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	fieldsCompared := 0
	var scores []float64

	// Compare addresses
	if len(query.Addresses) > 0 && len(index.Addresses) > 0 {
		fieldsCompared++
		if score := findBestAddressMatch(query.Addresses, index.Addresses); score > 0 {
			scores = append(scores, score)
		}
	}

	if len(scores) == 0 {
		return scorePiece{score: 0, weight: weight, fieldsCompared: 0, pieceType: "supporting"}
	}

	avgScore := calculateAverage(scores)
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

func findBestAddressMatch(queryAddrs, indexAddrs []Address) float64 {
	bestScore := 0.0
	for _, qa := range queryAddrs {
		for _, ia := range indexAddrs {
			if score := compareAddress(qa, ia); score > bestScore {
				bestScore = score
				if score > highConfidenceThreshold {
					return score // Early exit on high confidence match
				}
			}
		}
	}
	return bestScore
}

func compareAddress(query, index Address) float64 {
	var totalScore, totalWeight float64

	// Compare line1 (highest weight)
	if query.Line1 != "" && index.Line1 != "" {
		similarity := stringscore.JaroWinkler(query.Line1, index.Line1)
		totalScore += similarity * line1Weight
		totalWeight += line1Weight
	}

	// Compare line2
	if query.Line2 != "" && index.Line2 != "" {
		similarity := stringscore.JaroWinkler(query.Line2, index.Line2)
		totalScore += similarity * line2Weight
		totalWeight += line2Weight
	}

	// Compare city
	if query.City != "" && index.City != "" {
		similarity := stringscore.JaroWinkler(query.City, index.City)
		totalScore += similarity * cityWeight
		totalWeight += cityWeight
	}

	// Compare state (exact match)
	if query.State != "" && index.State != "" {
		if strings.EqualFold(query.State, index.State) {
			totalScore += stateWeight
		}
		totalWeight += stateWeight
	}

	// Compare postal code (exact match)
	if query.PostalCode != "" && index.PostalCode != "" {
		if strings.EqualFold(query.PostalCode, index.PostalCode) {
			totalScore += postalWeight
		}
		totalWeight += postalWeight
	}

	// Compare country (exact match)
	if query.Country != "" && index.Country != "" {
		if strings.EqualFold(query.Country, index.Country) {
			totalScore += countryWeight
		}
		totalWeight += countryWeight
	}

	if totalWeight == 0 {
		return 0
	}

	return totalScore / totalWeight
}
