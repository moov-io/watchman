package search

import (
	"fmt"
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
func compareEntityDates[Q any, I any](w io.Writer, query Entity[Q], index Entity[I], weight float64) ScorePiece {
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

	return ScorePiece{
		Score:          dateScore,
		Weight:         weight,
		Matched:        matched,
		Required:       fieldsCompared > 0,
		Exact:          dateScore > 0.99,
		FieldsCompared: fieldsCompared,
		PieceType:      "dates",
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

const (
	yearInHours = 60.0 * 24.0 * 365.25
)

// compareDates calculates a similarity score between two dates, useful for detecting fudged dates
// Returns a score between 0.0 (no similarity) and 1.0 (exact match).
func compareDates(date1, date2 *time.Time) float64 {
	if date1 == nil || date2 == nil {
		return 0.0
	}

	// Normalize to same time of day to focus on date components
	d1 := date1.Truncate(24 * time.Hour)
	d2 := date2.Truncate(24 * time.Hour)

	// Initialize score components
	var yearScore float64
	var monthScore float64
	var dayScore float64

	// Year comparison: ±5 years tolerance
	yearDiff := math.Abs(float64(d1.Year() - d2.Year()))
	if yearDiff <= 5 {
		yearScore = 1.0 - (0.1 * yearDiff) // Linear decay: 1.0 for same year, 0.5 at 5 years
	} else {
		yearScore = 0.2 // Low score for >5 years, but not zero to allow partial matches
	}

	// Month comparison: ±1 month tolerance, special handling for 1 vs. 10/11/12
	month1, month2 := int(d1.Month()), int(d2.Month())
	monthDiff := math.Abs(float64(month1 - month2))
	if monthDiff == 0 {
		monthScore = 1.0 // Exact match
	} else if monthDiff <= 1 {
		monthScore = 0.9 // Adjacent months
	} else if (month1 == 1 && (month2 >= 10 && month2 <= 12)) || (month2 == 1 && (month1 >= 10 && month1 <= 12)) {
		monthScore = 0.7 // Handle 1 vs. 10/11/12 (common typo or obfuscation)
	} else {
		monthScore = 0.3 // Low score for distant months
	}

	// Day comparison: ±3 days tolerance, special handling for similar numbers
	day1, day2 := d1.Day(), d2.Day()
	dayDiff := math.Abs(float64(day1 - day2))
	if dayDiff == 0 {
		dayScore = 1.0 // Exact match
	} else if dayDiff <= 3 {
		dayScore = 0.95 - (0.05 * dayDiff / 3) // Linear decay for close days
	} else {
		// Check for similar numbers (e.g., 1 vs. 11, 2 vs. 22)
		if areDaysSimilar(day1, day2) {
			dayScore = 0.7
		} else {
			dayScore = 0.3 // Low score for distant days
		}
	}

	// Overall score: Weighted average of components
	// Weights: Year (40%), Month (30%), Day (30%)
	overallScore := (0.4 * yearScore) + (0.3 * monthScore) + (0.3 * dayScore)

	// Ensure score is between 0.0 and 1.0
	return math.Max(0.0, math.Min(1.0, overallScore))
}

// areDaysSimilar checks if two days are numerically similar (e.g., 1 vs. 11, 2 vs. 22).
func areDaysSimilar(day1, day2 int) bool {
	// Convert days to strings to check for pattern similarity
	dayStr1, dayStr2 := fmt.Sprintf("%d", day1), fmt.Sprintf("%d", day2)

	// Case 1: Same single digit or repeated digit (e.g., 1 vs. 11, 2 vs. 22)
	if dayStr1 == dayStr2[:1] || dayStr2 == dayStr1[:1] {
		return true
	}

	// Case 2: Transposed digits (e.g., 12 vs. 21)
	if len(dayStr1) == 2 && len(dayStr2) == 2 {
		if dayStr1[0] == dayStr2[1] && dayStr1[1] == dayStr2[0] {
			return true
		}
	}

	return false
}

// areDatesLogical checks if dates make temporal sense
func areDatesLogical(person *Person, index *Person) bool {
	if person.BirthDate != nil && person.DeathDate != nil && index.BirthDate != nil && index.DeathDate != nil {
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
