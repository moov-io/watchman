// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/pkg/search"
)

// ConvertSanctionsData converts CSV sanctions data to search entities
func ConvertSanctionsData(data *ListData) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, entity := range data.SanctionsData {
		out = append(out, ToEntity(entity))
	}
	return out
}

// ToEntity converts a CSV SanctionsEntry to a search Entity
func ToEntity(src SanctionsEntry) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Source:     search.SourceUSCSL,
		SourceID:   src.ID,
		SourceData: src,
	}

	// Map entity type based on CSV 'Type' field
	switch strings.ToLower(src.Type) {
	case "individual":
		entity.Type = search.EntityPerson
		entity.Person = mapPerson(src)

	case "entity":
		// Determine if it's a business or organization based on source or name
		// For simplicity, assume entities are businesses unless specific indicators suggest otherwise
		if strings.Contains(strings.ToLower(src.Source), "military-industrial") || strings.Contains(strings.ToLower(src.Name), "bank") {
			entity.Type = search.EntityOrganization
			entity.Organization = mapOrganization(src)
		} else {
			entity.Type = search.EntityBusiness
			entity.Business = mapBusiness(src)
		}

	case "vessel":
		entity.Type = search.EntityVessel
		entity.Vessel = mapVessel(src)

	case "aircraft":
		entity.Type = search.EntityAircraft
		entity.Aircraft = mapAircraft(src)

	default:
		// Inspect some lists in detail
		switch {
		case
			strings.Contains(src.Source, "Entity List"),
			strings.Contains(src.Source, "ITAR Debarred"),
			strings.Contains(src.Source, "Military End User"),
			strings.Contains(src.Source, "Unverified List"):

			name := strings.Fields(strings.ToLower(src.Name))

			// Does the entity look like a business?
			for n := range name {
				for c := range companyNeedles {
					if strings.EqualFold(name[n], companyNeedles[c]) {
						goto business
					}
				}
			}
			// Otherwise they're an individual
			goto individual

		case strings.Contains(src.Source, "Nonproliferation Sanctions"):
			goto business

		case strings.Contains(src.Source, "Denied Persons List"):
			goto individual

		default:
			goto common
		}

	individual:
		entity.Type = search.EntityPerson
		entity.Person = mapPerson(src)
		goto common

	business:
		entity.Type = search.EntityBusiness
		entity.Business = mapBusiness(src)
		goto common
	}

common:
	// Map common fields
	entity.Name = cleanSubsidiarySuffix(src)
	entity.Contact = mapContactInfo(src)
	entity.Addresses = mapAddresses(src)
	entity.Affiliations = mapAffiliations(src)
	entity.SanctionsInfo = mapSanctionsInfo(src)

	return entity.Normalize()
}

var (
	companyNeedles = []string{
		"academy",
		"aviation",
		"bank",
		"business",
		"co.",
		"committee",
		"company",
		"corporation",
		"defense",
		"electronics",
		"equipment",
		"export",
		"group",
		"guard",
		"holding",
		"import",
		"import",
		"industrial",
		"industries",
		"industry",
		"institute",
		"intelligence",
		"international",
		"investment",
		"lab",
		"limited",
		"limited",
		"llc",
		"logistics",
		"ltd",
		"ltd.",
		"partnership",
		"revolutionary",
		"solutions",
		"subsidiary",
		"supply",
		"technology",
		"trading",
		"university",
	}
)

var (
	subsidiaryCleaner = strings.NewReplacer(
		"; and any successor,", "", "sub-unit,", "", "or subsidiary thereof", "",
	)
)

func cleanSubsidiarySuffix(src SanctionsEntry) string {
	return strings.TrimSpace(subsidiaryCleaner.Replace(src.Name))
}

// getAllNames returns all alternate names from the CSV
func getAllNames(src SanctionsEntry, entityType string) []string {
	if src.AltNames == "" {
		return nil
	}
	// Split alt_names by semicolon and normalize
	names := strings.Split(src.AltNames, ";")
	var out []string
	for _, name := range names {
		name = strings.TrimSpace(name)
		if name != "" {
			out = append(out, prepare.ReorderSDNName(name, strings.ToLower(entityType)))
		}
	}
	return out
}

func mapPerson(src SanctionsEntry) *search.Person {
	person := &search.Person{
		Name:     cleanSubsidiarySuffix(src),
		AltNames: getAllNames(src, "individual"),
	}

	// Map gender from IDs field if available
	if gender := extractGender(src.IDs); gender != "" {
		person.Gender = mapGender(gender)
	}

	// Map birth date
	if src.DatesOfBirth != "" {
		if date, err := time.Parse("2006-01-02", src.DatesOfBirth); err == nil {
			person.BirthDate = &date
		} else if year, err := strconv.Atoi(src.DatesOfBirth); err == nil {
			// Handle year-only birth dates (e.g., "1962")
			date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
			person.BirthDate = &date
		}
	}

	person.PlaceOfBirth = src.PlacesOfBirth

	// Map government IDs
	person.GovernmentIDs = mapGovernmentIDs(src)

	// Map titles
	if src.Title != "" {
		person.Titles = []string{src.Title}
	}

	// Map Citizenships and Nationalities to GovernmentIDs
	if src.Citizenships != "" {
		for _, citizenship := range strings.Split(src.Citizenships, ";") {
			citizenship = strings.TrimSpace(citizenship)
			if citizenship != "" {
				person.GovernmentIDs = append(person.GovernmentIDs, search.GovernmentID{
					Type:       search.GovernmentIDCitizenship,
					Country:    norm.Country(citizenship),
					Identifier: citizenship,
				})
			}
		}
	}

	if src.Nationalities != "" {
		for _, nationality := range strings.Split(src.Nationalities, ";") {
			nationality = strings.TrimSpace(nationality)
			if nationality != "" {
				person.GovernmentIDs = append(person.GovernmentIDs, search.GovernmentID{
					Type:       search.GovernmentIDNationality,
					Country:    norm.Country(nationality),
					Identifier: nationality,
				})
			}
		}
	}

	return person
}

func mapBusiness(src SanctionsEntry) *search.Business {
	business := &search.Business{
		Name:     cleanSubsidiarySuffix(src),
		AltNames: getAllNames(src, "business"),
	}

	// Map creation/dissolution dates
	if src.StartDate != "" {
		if date, err := time.Parse("2006-01-02", src.StartDate); err == nil {
			business.Created = &date
		}
	}
	if src.EndDate != "" {
		if date, err := time.Parse("2006-01-02", src.EndDate); err == nil {
			business.Dissolved = &date
		}
	}

	// Map government IDs
	business.GovernmentIDs = mapGovernmentIDs(src)

	return business
}

func mapOrganization(src SanctionsEntry) *search.Organization {
	org := &search.Organization{
		Name:     cleanSubsidiarySuffix(src),
		AltNames: getAllNames(src, "organization"),
	}

	// Map creation/dissolution dates
	if src.StartDate != "" {
		if date, err := time.Parse("2006-01-02", src.StartDate); err == nil {
			org.Created = &date
		}
	}
	if src.EndDate != "" {
		if date, err := time.Parse("2006-01-02", src.EndDate); err == nil {
			org.Dissolved = &date
		}
	}

	// Map government IDs
	org.GovernmentIDs = mapGovernmentIDs(src)

	return org
}

func mapAircraft(src SanctionsEntry) *search.Aircraft {
	aircraft := &search.Aircraft{
		Name:     cleanSubsidiarySuffix(src),
		AltNames: getAllNames(src, "aircraft"),
		Type:     mapAircraftType(src.VesselType), // Using VesselType for simplicity, assuming aircraft type is similar
		Flag:     src.VesselFlag,                  // Map to country code
	}

	// Map additional fields if available
	if src.StartDate != "" {
		if date, err := time.Parse("2006-01-02", src.StartDate); err == nil {
			aircraft.Built = &date
		}
	}

	return aircraft
}

func mapVessel(src SanctionsEntry) *search.Vessel {
	vessel := &search.Vessel{
		Name:                   cleanSubsidiarySuffix(src),
		AltNames:               getAllNames(src, "vessel"),
		Type:                   mapVesselType(src.VesselType),
		Flag:                   src.VesselFlag,
		CallSign:               src.CallSign,
		Owner:                  src.VesselOwner,
		Tonnage:                parseIntOrZero(src.GrossTonnage),
		GrossRegisteredTonnage: parseIntOrZero(src.GrossRegisteredTonnage),
	}

	// Map build date
	if src.StartDate != "" {
		if date, err := time.Parse("2006-01-02", src.StartDate); err == nil {
			vessel.Built = &date
		}
	}

	return vessel
}

func mapContactInfo(src SanctionsEntry) search.ContactInfo {
	info := search.ContactInfo{}

	// Extract contact info from IDs or Remarks if available
	if src.IDs != "" {
		ids := strings.Split(src.IDs, ";")
		for _, id := range ids {
			id = strings.TrimSpace(id)
			if strings.HasPrefix(strings.ToLower(id), "website") {
				info.Websites = append(info.Websites, strings.TrimPrefix(id, "Website "))
			} else if strings.HasPrefix(strings.ToLower(id), "email address") {
				info.EmailAddresses = append(info.EmailAddresses, strings.TrimPrefix(id, "Email Address "))
			} else if strings.HasPrefix(strings.ToLower(id), "telephone") {
				info.PhoneNumbers = append(info.PhoneNumbers, strings.TrimPrefix(id, "Telephone "))
			} else if strings.HasPrefix(strings.ToLower(id), "fax") {
				info.FaxNumbers = append(info.FaxNumbers, strings.TrimPrefix(id, "Fax "))
			}
		}
	}

	return info
}

func mapAddresses(src SanctionsEntry) []search.Address {
	if src.Addresses == "" {
		return nil
	}

	var result []search.Address
	// Split addresses by semicolon for multiple addresses
	addresses := strings.Split(src.Addresses, ";")
	for _, addr := range addresses {
		addr = strings.TrimSpace(addr)
		if addr == "" {
			continue
		}

		// Parse address components (basic parsing, assumes comma-separated parts)
		parts := strings.Split(addr, ",")
		var mappedAddr search.Address
		for i, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			// Basic heuristic: last part is often country, second-to-last is city or postal code
			if i == len(parts)-1 {
				mappedAddr.Country = norm.Country(part)
			} else if i == len(parts)-2 {
				// Check if it's a postal code (basic heuristic)
				if len(part) <= 10 && strings.ContainsAny(part, "0123456789") {
					mappedAddr.PostalCode = part
				} else {
					mappedAddr.City = part
				}
			} else if i == len(parts)-3 && mappedAddr.PostalCode == "" {
				mappedAddr.City = part
			} else {
				// Assume earlier parts are Line1 or Line2
				if mappedAddr.Line1 == "" {
					mappedAddr.Line1 = part
				} else {
					mappedAddr.Line2 += part + " "
				}
			}
		}

		if mappedAddr.Line1 != "" && mappedAddr.Country != "" {
			result = append(result, mappedAddr)
		}
	}

	return result
}

func mapAffiliations(src SanctionsEntry) []search.Affiliation {
	// Affiliations are not explicitly in the CSV; extract from Remarks if applicable
	if src.Remarks == "" {
		return nil
	}

	var affiliations []search.Affiliation
	// Basic heuristic: look for patterns like "Linked To" or "Owned By" in remarks
	remarks := strings.Split(src.Remarks, ";")
	for _, remark := range remarks {
		remark = strings.TrimSpace(remark)
		if strings.Contains(strings.ToLower(remark), "linked to") ||
			strings.Contains(strings.ToLower(remark), "owned by") ||
			strings.Contains(strings.ToLower(remark), "subsidiary of") {
			parts := strings.SplitN(remark, ":", 2)
			if len(parts) == 2 {
				aff := search.Affiliation{
					Type:    strings.TrimSpace(parts[0]),
					Details: strings.TrimSpace(parts[1]),
				}
				affiliations = append(affiliations, aff)
			}
		}
	}

	return affiliations
}

func mapSanctionsInfo(src SanctionsEntry) *search.SanctionsInfo {
	info := &search.SanctionsInfo{
		Description: src.Remarks,
	}

	// Map programs
	if src.Programs != "" {
		info.Programs = strings.Split(src.Programs, ";")
	}

	// Check for secondary sanctions in IDs or Remarks
	if strings.Contains(strings.ToLower(src.IDs), "secondary sanctions") ||
		strings.Contains(strings.ToLower(src.Remarks), "secondary sanctions") {
		info.Secondary = true
	}

	return info
}

func mapGovernmentIDs(src SanctionsEntry) []search.GovernmentID {
	if src.IDs == "" {
		return nil
	}

	var ids []search.GovernmentID
	idEntries := strings.Split(src.IDs, ";")
	for _, idEntry := range idEntries {
		idEntry = strings.TrimSpace(idEntry)
		if idEntry == "" {
			continue
		}

		// Parse ID format (e.g., "Passport, U00242309, TR")
		parts := strings.Split(idEntry, ",")
		if len(parts) < 2 {
			continue
		}

		idType := strings.TrimSpace(parts[0])
		identifier := strings.TrimSpace(parts[1])
		country := ""
		if len(parts) > 2 {
			country = norm.Country(strings.TrimSpace(parts[2]))
		}

		govID := search.GovernmentID{
			Type:       mapIDType(idType),
			Country:    country,
			Identifier: identifier,
		}
		if string(govID.Type) != "" {
			ids = append(ids, govID)
		}
	}

	return ids
}

// Helper functions

func mapDate(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	if t, err := time.Parse("2006-01-02", dateStr); err == nil {
		return &t
	}
	return nil
}

func mapGender(value string) search.Gender {
	switch strings.ToLower(value) {
	case "male", "m":
		return search.GenderMale
	case "female", "f":
		return search.GenderFemale
	default:
		return search.GenderUnknown
	}
}

func mapIDType(idType string) search.GovernmentIDType {
	switch strings.ToLower(idType) {
	case "passport":
		return search.GovernmentIDPassport
	case "driver's license", "drivers license":
		return search.GovernmentIDDriversLicense
	case "national id", "national identification number":
		return search.GovernmentIDNational
	case "tax id", "tax identification number":
		return search.GovernmentIDTax
	case "ssn", "social security number":
		return search.GovernmentIDSSN
	case "cedula":
		return search.GovernmentIDCedula
	case "curp":
		return search.GovernmentIDCURP
	case "cuit":
		return search.GovernmentIDCUIT
	case "electoral registry":
		return search.GovernmentIDElectoral
	case "business registration", "business registration document", "business registration number":
		return search.GovernmentIDBusinessRegisration
	case "commercial registry":
		return search.GovernmentIDCommercialRegistry
	case "birth certificate":
		return search.GovernmentIDBirthCert
	case "refugee id":
		return search.GovernmentIDRefugee
	case "diplomatic passport":
		return search.GovernmentIDDiplomaticPass
	case "personal id":
		return search.GovernmentIDPersonalID
	default:
		return ""
	}
}

func mapAircraftType(value string) search.AircraftType {
	switch strings.ToLower(value) {
	case "cargo":
		return search.AircraftCargo
	default:
		return search.AircraftTypeUnknown
	}
}

func mapVesselType(value string) search.VesselType {
	switch strings.ToLower(value) {
	case "cargo":
		return search.VesselTypeCargo
	default:
		return search.VesselTypeUnknown
	}
}

func parseIntOrZero(value string) int {
	if value == "" {
		return 0
	}
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return 0
}

func extractGender(ids string) string {
	// Extract gender from IDs field (e.g., "Gender, Male")
	idEntries := strings.Split(ids, ";")
	for _, idEntry := range idEntries {
		idEntry = strings.TrimSpace(idEntry)
		if strings.HasPrefix(strings.ToLower(idEntry), "gender") {
			parts := strings.SplitN(idEntry, ",", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}
