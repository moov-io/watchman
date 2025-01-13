package search

import (
	"io"
	"strings"

	"github.com/moov-io/watchman/internal/stringscore"
)

func compareSupportingInfo[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	fieldsCompared := 0
	var scores []float64

	// Compare sanctions
	if query.SanctionsInfo != nil && index.SanctionsInfo != nil {
		fieldsCompared++
		if score := compareSanctionsPrograms(w, query.SanctionsInfo, index.SanctionsInfo); score > 0 {
			scores = append(scores, score)
		}
	}

	// Compare historical info
	if len(query.HistoricalInfo) > 0 && len(index.HistoricalInfo) > 0 {
		fieldsCompared++
		if score := compareHistoricalValues(query.HistoricalInfo, index.HistoricalInfo); score > 0 {
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

func compareSanctionsPrograms(w io.Writer, query *SanctionsInfo, index *SanctionsInfo) float64 {
	if query == nil || index == nil {
		return 0
	}

	// Compare programs
	programScore := 0.0
	if len(query.Programs) > 0 && len(index.Programs) > 0 {
		matches := 0
		for _, qp := range query.Programs {
			for _, ip := range index.Programs {
				if strings.EqualFold(qp, ip) {
					matches++
					break
				}
			}
		}
		programScore = float64(matches) / float64(len(query.Programs))
	}

	// Adjust score based on secondary sanctions match
	if query.Secondary != index.Secondary {
		programScore *= 0.8
	}

	return programScore
}

func compareHistoricalValues(queryHist, indexHist []HistoricalInfo) float64 {
	bestScore := 0.0
	for _, qh := range queryHist {
		for _, ih := range indexHist {
			// Type must match exactly
			if !strings.EqualFold(qh.Type, ih.Type) {
				continue
			}

			// Compare values
			similarity := stringscore.JaroWinkler(qh.Value, ih.Value)
			if similarity > bestScore {
				bestScore = similarity
			}
		}
	}
	return bestScore
}
