// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/pkg/search"
)

// imoPattern matches IMO numbers (7 digits) in text
// Compiled once at package init for performance
var imoPattern = regexp.MustCompile(`(?i)IMO[:\s]*(\d{7})`)

// ConvertEUCSLData converts all EU CSL records to search entities.
// This is the main entry point called by the download pipeline.
func ConvertEUCSLData(records []CSLRecord) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, record := range records {
		entity := ToEntity(record)
		out = append(out, entity)
	}
	return out
}

// ToEntity converts a single EU CSLRecord to a search.Entity.
// Handles Person, Business/Enterprise, and Vessel types.
func ToEntity(record CSLRecord) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Source:     search.SourceEUCSL,
		SourceID:   strconv.Itoa(record.EntityLogicalID),
		SourceData: record,
	}

	// Determine entity type based on EntitySubjectTypeCode (classification code like "person", "enterprise")
	// EntitySubjectType only contains single letters like "P" or "E"
	entityType := strings.ToLower(strings.TrimSpace(record.EntitySubjectTypeCode))

	switch entityType {
	case "person":
		entity.Type = search.EntityPerson
		entity.Person = mapPerson(record)
		entity.Name = entity.Person.Name

	case "enterprise":
		entity.Type = search.EntityBusiness
		entity.Business = mapBusiness(record)
		entity.Name = entity.Business.Name

	default:
		// Check if it might be a vessel based on IMO number in remarks
		if isVessel(record) {
			entity.Type = search.EntityVessel
			entity.Vessel = mapVessel(record)
			entity.Name = entity.Vessel.Name
		} else {
			// Unknown type, treat as organization/business
			entity.Type = search.EntityOrganization
			entity.Organization = mapOrganization(record)
			entity.Name = entity.Organization.Name
		}
	}

	// Map common fields
	entity.Addresses = mapAddresses(record)
	entity.SanctionsInfo = mapSanctionsInfo(record)
	entity.Contact = mapContactInfo(record)

	return entity.Normalize()
}

// isVessel checks if a record represents a vessel based on IMO number pattern in remarks
func isVessel(record CSLRecord) bool {
	return imoPattern.MatchString(record.EntityRemark)
}

// extractIMONumber extracts IMO number from text
func extractIMONumber(text string) string {
	matches := imoPattern.FindStringSubmatch(text)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}

func mapPerson(record CSLRecord) *search.Person {
	person := &search.Person{}

	// Primary name (first in list)
	if len(record.NameAliasWholeNames) > 0 {
		person.Name = strings.TrimSpace(record.NameAliasWholeNames[0])
	}

	// Alt names (all except first)
	if len(record.NameAliasWholeNames) > 1 {
		person.AltNames = make([]string, 0, len(record.NameAliasWholeNames)-1)
		for _, name := range record.NameAliasWholeNames[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				person.AltNames = append(person.AltNames, trimmed)
			}
		}
	}

	// Gender (M/F -> male/female)
	if len(record.NameAliasGenders) > 0 {
		person.Gender = mapGender(record.NameAliasGenders[0])
	}

	// Titles - combine NameAliasTitles and NameAliasFunctions
	var titles []string
	titleSet := make(map[string]struct{})

	for _, title := range record.NameAliasTitles {
		if trimmed := strings.TrimSpace(title); trimmed != "" {
			if _, exists := titleSet[trimmed]; !exists {
				titles = append(titles, trimmed)
				titleSet[trimmed] = struct{}{}
			}
		}
	}
	for _, function := range record.NameAliasFunctions {
		if trimmed := strings.TrimSpace(function); trimmed != "" {
			if _, exists := titleSet[trimmed]; !exists {
				titles = append(titles, trimmed)
				titleSet[trimmed] = struct{}{}
			}
		}
	}
	if len(titles) > 0 {
		person.Titles = titles
	}

	// Birth date (format: YYYY-MM-DD)
	if len(record.BirthDates) > 0 && record.BirthDates[0] != "" {
		if t := parseEUDate(record.BirthDates[0]); t != nil {
			person.BirthDate = t
		}
	}

	// Place of birth - combine multiple sources for the most complete info
	person.PlaceOfBirth = buildPlaceOfBirth(record)

	// Government IDs (passports, national IDs, etc.)
	person.GovernmentIDs = mapGovernmentIDs(record.Identifications)

	return person
}

// buildPlaceOfBirth combines birth place data from multiple fields
func buildPlaceOfBirth(record CSLRecord) string {
	var parts []string

	// Add place if available (more specific than city)
	if len(record.BirthPlaces) > 0 && record.BirthPlaces[0] != "" {
		parts = append(parts, strings.TrimSpace(record.BirthPlaces[0]))
	}

	// Add city
	if len(record.BirthCities) > 0 && record.BirthCities[0] != "" {
		city := strings.TrimSpace(record.BirthCities[0])
		// Don't duplicate if place already contains this
		if len(parts) == 0 || !strings.Contains(parts[0], city) {
			parts = append(parts, city)
		}
	}

	// Add region
	if len(record.BirthRegions) > 0 && record.BirthRegions[0] != "" {
		parts = append(parts, strings.TrimSpace(record.BirthRegions[0]))
	}

	// Add country
	if len(record.BirthCountries) > 0 && record.BirthCountries[0] != "" {
		parts = append(parts, strings.TrimSpace(record.BirthCountries[0]))
	}

	return strings.Join(parts, ", ")
}

// mapGender converts EU CSL gender codes to search.Gender
func mapGender(code string) search.Gender {
	switch strings.ToUpper(strings.TrimSpace(code)) {
	case "M", "MALE":
		return search.GenderMale
	case "F", "FEMALE":
		return search.GenderFemale
	default:
		return search.GenderUnknown
	}
}

// mapGovernmentIDs converts EU CSL identification info to search.GovernmentID
func mapGovernmentIDs(ids []IdentificationInfo) []search.GovernmentID {
	if len(ids) == 0 {
		return nil
	}

	result := make([]search.GovernmentID, 0, len(ids))
	for _, id := range ids {
		if id.Number == "" {
			continue
		}

		govID := search.GovernmentID{
			Identifier: extractIDNumber(id.Number),
			Type:       mapIDType(id.TypeCode, id.TypeDescription),
			Country:    norm.Country(coalesce(id.CountryDesc, id.CountryISO)),
		}

		// Set name based on type description or type code
		if id.TypeDescription != "" {
			govID.Name = id.TypeDescription
		} else if id.TypeCode != "" {
			govID.Name = id.TypeCode
		} else {
			govID.Name = "Unknown"
		}

		result = append(result, govID)
	}

	return result
}

// extractIDNumber cleans the identification number
// EU CSL often includes type info in the number field like "488555 (passport-National passport)"
func extractIDNumber(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}

	// Check if there's a parenthesis - the actual number is before it
	if idx := strings.Index(raw, "("); idx > 0 {
		return strings.TrimSpace(raw[:idx])
	}

	return raw
}

// mapIDType maps EU CSL type codes to search.GovernmentIDType
func mapIDType(typeCode, typeDesc string) search.GovernmentIDType {
	code := strings.ToLower(strings.TrimSpace(typeCode))
	desc := strings.ToLower(strings.TrimSpace(typeDesc))

	// Check type code first
	switch code {
	case "passport":
		return search.GovernmentIDPassport
	case "id":
		return search.GovernmentIDNational
	case "ssn":
		return search.GovernmentIDSSN
	case "tax":
		return search.GovernmentIDTax
	}

	// Fall back to description
	if strings.Contains(desc, "passport") {
		if strings.Contains(desc, "diplomatic") {
			return search.GovernmentIDDiplomaticPass
		}
		return search.GovernmentIDPassport
	}
	if strings.Contains(desc, "national") && strings.Contains(desc, "identification") {
		return search.GovernmentIDNational
	}
	if strings.Contains(desc, "driver") || strings.Contains(desc, "license") {
		return search.GovernmentIDDriversLicense
	}
	if strings.Contains(desc, "tax") {
		return search.GovernmentIDTax
	}
	if strings.Contains(desc, "birth") {
		return search.GovernmentIDBirthCert
	}
	if strings.Contains(desc, "business") || strings.Contains(desc, "registration") {
		return search.GovernmentIDBusinessRegisration
	}

	// Default to personal ID for unknown types
	return search.GovernmentIDPersonalID
}

// coalesce returns the first non-empty string
func coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// parseEUDate parses dates in EU format (YYYY-MM-DD or DD/MM/YYYY)
func parseEUDate(s string) *time.Time {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	// Try full date YYYY-MM-DD
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return &t
	}

	// Try DD/MM/YYYY format (European style)
	if t, err := time.Parse("02/01/2006", s); err == nil {
		return &t
	}

	return nil
}

func mapBusiness(record CSLRecord) *search.Business {
	business := &search.Business{}

	// Primary name (first in list)
	if len(record.NameAliasWholeNames) > 0 {
		business.Name = strings.TrimSpace(record.NameAliasWholeNames[0])
	}

	// Alt names (all except first)
	if len(record.NameAliasWholeNames) > 1 {
		business.AltNames = make([]string, 0, len(record.NameAliasWholeNames)-1)
		for _, name := range record.NameAliasWholeNames[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				business.AltNames = append(business.AltNames, trimmed)
			}
		}
	}

	// Government IDs for business entities
	business.GovernmentIDs = mapGovernmentIDs(record.Identifications)

	return business
}

func mapOrganization(record CSLRecord) *search.Organization {
	org := &search.Organization{}

	// Primary name (first in list)
	if len(record.NameAliasWholeNames) > 0 {
		org.Name = strings.TrimSpace(record.NameAliasWholeNames[0])
	}

	// Alt names (all except first)
	if len(record.NameAliasWholeNames) > 1 {
		org.AltNames = make([]string, 0, len(record.NameAliasWholeNames)-1)
		for _, name := range record.NameAliasWholeNames[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				org.AltNames = append(org.AltNames, trimmed)
			}
		}
	}

	// Government IDs for organizations
	org.GovernmentIDs = mapGovernmentIDs(record.Identifications)

	return org
}

func mapVessel(record CSLRecord) *search.Vessel {
	vessel := &search.Vessel{}

	// Primary name (first in list)
	if len(record.NameAliasWholeNames) > 0 {
		vessel.Name = strings.TrimSpace(record.NameAliasWholeNames[0])
	}

	// Alt names (all except first)
	if len(record.NameAliasWholeNames) > 1 {
		vessel.AltNames = make([]string, 0, len(record.NameAliasWholeNames)-1)
		for _, name := range record.NameAliasWholeNames[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				vessel.AltNames = append(vessel.AltNames, trimmed)
			}
		}
	}

	// Try to extract IMO number from remarks
	if imo := extractIMONumber(record.EntityRemark); imo != "" {
		vessel.IMONumber = imo
	}

	// Try to get flag from address country
	if len(record.AddressCountryDescriptions) > 0 && record.AddressCountryDescriptions[0] != "" {
		vessel.Flag = norm.Country(record.AddressCountryDescriptions[0])
	}

	return vessel
}

func mapAddresses(record CSLRecord) []search.Address {
	// Determine the maximum number of addresses based on available data
	maxAddrs := max(
		len(record.AddressStreets),
		len(record.AddressCities),
		len(record.AddressZipCodes),
		len(record.AddressCountryDescriptions),
	)

	if maxAddrs == 0 {
		return nil
	}

	var addresses []search.Address

	for i := 0; i < maxAddrs; i++ {
		addr := search.Address{}
		hasContent := false

		// Street address
		if i < len(record.AddressStreets) && record.AddressStreets[i] != "" {
			addr.Line1 = strings.TrimSpace(record.AddressStreets[i])
			hasContent = true
		}

		// PO Box (add to Line2 if present)
		if i < len(record.AddressPoBoxes) && record.AddressPoBoxes[i] != "" {
			addr.Line2 = "PO Box " + strings.TrimSpace(record.AddressPoBoxes[i])
			hasContent = true
		}

		// Place (add to Line2 if no PO Box)
		if addr.Line2 == "" && i < len(record.AddressPlaces) && record.AddressPlaces[i] != "" {
			addr.Line2 = strings.TrimSpace(record.AddressPlaces[i])
			hasContent = true
		}

		// City
		if i < len(record.AddressCities) && record.AddressCities[i] != "" {
			addr.City = strings.TrimSpace(record.AddressCities[i])
			hasContent = true
		}

		// Zip code
		if i < len(record.AddressZipCodes) && record.AddressZipCodes[i] != "" {
			addr.PostalCode = strings.TrimSpace(record.AddressZipCodes[i])
			hasContent = true
		}

		// Region/State
		if i < len(record.AddressRegions) && record.AddressRegions[i] != "" {
			addr.State = strings.TrimSpace(record.AddressRegions[i])
			hasContent = true
		}

		// Country
		if i < len(record.AddressCountryDescriptions) && record.AddressCountryDescriptions[i] != "" {
			addr.Country = norm.Country(record.AddressCountryDescriptions[i])
			hasContent = true
		}

		if hasContent {
			addresses = append(addresses, addr)
		}
	}

	return addresses
}

func mapContactInfo(record CSLRecord) search.ContactInfo {
	var contact search.ContactInfo

	// Address contact info often contains phone/fax/email
	for _, info := range record.AddressContactInfos {
		info = strings.TrimSpace(info)
		if info == "" {
			continue
		}

		// Parse contact info - it may contain various contact details
		lowerInfo := strings.ToLower(info)
		if strings.Contains(lowerInfo, "@") {
			contact.EmailAddresses = append(contact.EmailAddresses, info)
		} else if strings.Contains(lowerInfo, "fax") {
			contact.FaxNumbers = append(contact.FaxNumbers, info)
		} else if strings.Contains(lowerInfo, "tel") || strings.Contains(lowerInfo, "phone") {
			contact.PhoneNumbers = append(contact.PhoneNumbers, info)
		} else if strings.Contains(lowerInfo, "http") || strings.Contains(lowerInfo, "www") {
			contact.Websites = append(contact.Websites, info)
		}
	}

	return contact
}

func mapSanctionsInfo(record CSLRecord) *search.SanctionsInfo {
	// Check if we have any sanctions-related info
	hasContent := record.EntityRemark != "" ||
		record.EntityReferenceNumber != "" ||
		len(record.Citizenships) > 0 ||
		record.EntityRegulationProgramme != "" ||
		record.EntityDesignationDate != "" ||
		record.EntityUnitedNationID != ""

	if !hasContent {
		return nil
	}

	info := &search.SanctionsInfo{}

	// Parse sanctions programs from EntityRegulationProgramme
	// Format is often like "UKR" or "IRN, TAQA"
	if record.EntityRegulationProgramme != "" {
		programs := strings.Split(record.EntityRegulationProgramme, ",")
		for _, p := range programs {
			if trimmed := strings.TrimSpace(p); trimmed != "" {
				info.Programs = append(info.Programs, trimmed)
			}
		}
	}

	// Build description from various fields
	var parts []string
	if record.EntityRemark != "" {
		parts = append(parts, record.EntityRemark)
	}
	if record.EntityReferenceNumber != "" {
		parts = append(parts, "Ref: "+record.EntityReferenceNumber)
	}
	if record.EntityUnitedNationID != "" {
		parts = append(parts, "UN ID: "+record.EntityUnitedNationID)
	}
	if record.EntityDesignationDate != "" {
		parts = append(parts, "Designated: "+record.EntityDesignationDate)
	}
	if record.EntityDesignationDetails != "" {
		parts = append(parts, record.EntityDesignationDetails)
	}
	if len(record.Citizenships) > 0 {
		citizenships := make([]string, 0, len(record.Citizenships))
		for _, c := range record.Citizenships {
			if trimmed := strings.TrimSpace(c); trimmed != "" {
				citizenships = append(citizenships, trimmed)
			}
		}
		if len(citizenships) > 0 {
			parts = append(parts, "Citizenship: "+strings.Join(citizenships, ", "))
		}
	}

	if len(parts) > 0 {
		info.Description = strings.Join(parts, "; ")
	}

	return info
}
