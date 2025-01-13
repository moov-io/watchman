package search

import (
	"io"
	"math"
	"time"
)

const (
	// Date thresholds in days
	exactMatch = 0
	veryClose  = 2   // Within 2 days
	close      = 7   // Within a week
	moderate   = 30  // Within a month
	distant    = 365 // Within a year
)

// compareEntityDates performs date comparisons based on entity type
func compareEntityDates[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) scorePiece {
	if query.Type != index.Type {
		return scorePiece{score: 0, weight: weight, pieceType: "dates", fieldsCompared: 0}
	}

	var dateScore float64
	var fieldsCompared int
	var matched bool

	switch query.Type {
	case EntityPerson:
		dateScore, matched, fieldsCompared = comparePersonDates(query.Person, index.Person)
	case EntityBusiness:
		dateScore, matched, fieldsCompared = compareBusinessDates(query.Business, index.Business)
	case EntityOrganization:
		dateScore, matched, fieldsCompared = compareOrgDates(query.Organization, index.Organization)
	case EntityVessel, EntityAircraft:
		dateScore, matched, fieldsCompared = compareAssetDates(query.Type, query, index)
	}

	return scorePiece{
		score:          dateScore,
		weight:         weight,
		matched:        matched,
		required:       fieldsCompared > 0,
		exact:          dateScore > 0.99,
		fieldsCompared: fieldsCompared,
		pieceType:      "dates",
	}
}

// comparePersonDates handles birth and death dates
func comparePersonDates(query *Person, index *Person) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	fieldsCompared := 0
	var scores []float64

	// Birth date comparison
	if query.BirthDate != nil && index.BirthDate != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.BirthDate, index.BirthDate))
	}

	// Death date comparison
	if query.DeathDate != nil && index.DeathDate != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.DeathDate, index.DeathDate))
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	// Calculate average score
	avgScore := calculateAverage(scores)

	// Check date consistency if both dates are present
	if fieldsCompared == 2 && !areDatesLogical(query, index) {
		avgScore *= 0.5 // Penalty for illogical dates
	}

	return avgScore, avgScore > 0.7, fieldsCompared
}

// compareBusinessDates handles business dates
func compareBusinessDates(query *Business, index *Business) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	fieldsCompared := 0
	var scores []float64

	// Created date comparison
	if query.Created != nil && index.Created != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.Created, index.Created))
	}

	// Dissolved date comparison (if present)
	if query.Dissolved != nil && index.Dissolved != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.Dissolved, index.Dissolved))
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	avgScore := calculateAverage(scores)
	return avgScore, avgScore > 0.7, fieldsCompared
}

// compareOrgDates handles organization dates
func compareOrgDates(query *Organization, index *Organization) (float64, bool, int) {
	if query == nil || index == nil {
		return 0, false, 0
	}

	fieldsCompared := 0
	var scores []float64

	// Created date comparison
	if query.Created != nil && index.Created != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.Created, index.Created))
	}

	// Dissolved date comparison
	if query.Dissolved != nil && index.Dissolved != nil {
		fieldsCompared++
		scores = append(scores, compareDates(query.Dissolved, index.Dissolved))
	}

	if len(scores) == 0 {
		return 0, false, fieldsCompared
	}

	avgScore := calculateAverage(scores)
	return avgScore, avgScore > 0.7, fieldsCompared
}

// compareAssetDates handles vessel and aircraft dates
func compareAssetDates[Q any, I any](entityType EntityType, query Entity[Q], index Entity[I]) (float64, bool, int) {
	fieldsCompared := 0
	var builtDate1, builtDate2 *time.Time

	switch entityType {
	case EntityVessel:
		if query.Vessel != nil && index.Vessel != nil {
			builtDate1 = query.Vessel.Built
			builtDate2 = index.Vessel.Built
		}
	case EntityAircraft:
		if query.Aircraft != nil && index.Aircraft != nil {
			builtDate1 = query.Aircraft.Built
			builtDate2 = index.Aircraft.Built
		}
	}

	if builtDate1 == nil || builtDate2 == nil {
		return 0, false, fieldsCompared
	}

	fieldsCompared = 1
	score := compareDates(builtDate1, builtDate2)
	return score, score > 0.7, fieldsCompared
}

// compareDates calculates similarity score between two dates
func compareDates(date1, date2 *time.Time) float64 {
	// Normalize to same time of day
	d1 := date1.Truncate(24 * time.Hour)
	d2 := date2.Truncate(24 * time.Hour)

	// Calculate difference in days
	diffDays := math.Abs(d1.Sub(d2).Hours() / 24)

	switch {
	case diffDays <= float64(exactMatch):
		return 1.0
	case diffDays <= float64(veryClose):
		return 0.95 - (0.05 * (diffDays / float64(veryClose)))
	case diffDays <= float64(close):
		return 0.9 - (0.1 * (diffDays / float64(close)))
	case diffDays <= float64(moderate):
		return 0.8 - (0.2 * (diffDays / float64(moderate)))
	case diffDays <= float64(distant):
		return 0.6 - (0.3 * (diffDays / float64(distant)))
	default:
		return 0.0
	}
}

// areDatesLogical checks if dates make temporal sense
func areDatesLogical(person *Person, index *Person) bool {
	if person.BirthDate != nil && person.DeathDate != nil &&
		index.BirthDate != nil && index.DeathDate != nil {

		// Check that birth precedes death in both records
		personValid := person.BirthDate.Before(*person.DeathDate)
		indexValid := index.BirthDate.Before(*index.DeathDate)

		if !personValid || !indexValid {
			return false
		}

		// Check relative lifespans (within 20%)
		personSpan := person.DeathDate.Sub(*person.BirthDate)
		indexSpan := index.DeathDate.Sub(*index.BirthDate)

		ratio := math.Max(personSpan.Hours(), indexSpan.Hours()) / math.Max(1, math.Min(personSpan.Hours(), indexSpan.Hours()))
		return ratio <= 1.2
	}
	return true
}
