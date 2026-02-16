// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/pkg/search"
)

// ConvertSanctionsListData converts all UK Sanctions List records to search entities.
// This is the main entry point called by the download pipeline.
func ConvertSanctionsListData(records []SanctionsListRecord) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, record := range records {
		entity := ToEntity(record)
		out = append(out, entity)
	}
	return out
}

// ToEntity converts a single UK SanctionsListRecord to a search.Entity.
// Handles Individual (Person), Entity (Business), and Ship (Vessel) types.
func ToEntity(record SanctionsListRecord) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Source:     search.SourceUKCSL,
		SourceID:   record.UniqueID,
		SourceData: record,
	}

	// Determine entity type and map type-specific data
	if record.EntityType != nil {
		switch *record.EntityType {
		case UKSLIndividual:
			entity.Type = search.EntityPerson
			entity.Person = mapPerson(record)
			entity.Name = entity.Person.Name

		case UKSLEntity:
			entity.Type = search.EntityBusiness
			entity.Business = mapBusiness(record)
			entity.Name = entity.Business.Name

		case UKSLShip:
			entity.Type = search.EntityVessel
			entity.Vessel = mapVessel(record)
			entity.Name = entity.Vessel.Name

		case Undefined:
			entity.Type = search.EntityUnknown
		}
	}

	// Map common fields
	entity.Addresses = mapAddresses(record)
	entity.SanctionsInfo = mapSanctionsInfo(record)

	return entity.Normalize()
}

// collectAltNames extracts alternate names from a record.
// It takes all names except the first (primary) one, plus any non-Latin script names.
func collectAltNames(record SanctionsListRecord) []string {
	var altNames []string
	if len(record.Names) > 1 {
		for _, name := range record.Names[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				altNames = append(altNames, trimmed)
			}
		}
	}
	for _, name := range record.NonLatinScriptNames {
		if trimmed := strings.TrimSpace(name); trimmed != "" {
			altNames = append(altNames, trimmed)
		}
	}
	return altNames
}

func mapPerson(record SanctionsListRecord) *search.Person {
	person := &search.Person{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		person.Name = strings.TrimSpace(record.Names[0])
	}

	person.AltNames = collectAltNames(record)

	// Title
	if record.NameTitle != "" {
		person.Titles = []string{record.NameTitle}
	}

	// Place of birth (combine town and country if available)
	if record.TownOfBirth != "" && record.CountryOfBirth != "" {
		person.PlaceOfBirth = record.TownOfBirth + ", " + record.CountryOfBirth
	} else if record.CountryOfBirth != "" {
		person.PlaceOfBirth = record.CountryOfBirth
	} else if record.TownOfBirth != "" {
		person.PlaceOfBirth = record.TownOfBirth
	}

	// Date of birth
	if record.DOB != "" {
		if t := parseUKDate(record.DOB); t != nil {
			person.BirthDate = t
		}
	}

	// Gender
	if record.Gender != "" {
		switch strings.ToLower(record.Gender) {
		case "male":
			person.Gender = search.GenderMale
		case "female":
			person.Gender = search.GenderFemale
		}
	}

	// Government IDs
	var govIDs []search.GovernmentID

	// Passport
	if record.PassportNumber != "" {
		govID := search.GovernmentID{
			Type:       search.GovernmentIDPassport,
			Identifier: record.PassportNumber,
		}
		// Try to extract country from additional info
		if record.PassportAdditionalInfo != "" {
			govID.Country = extractCountryFromDetails(record.PassportAdditionalInfo)
		}
		govIDs = append(govIDs, govID)
	}

	// National ID
	if record.NationalIDNumber != "" {
		govID := search.GovernmentID{
			Type:       search.GovernmentIDNational,
			Identifier: record.NationalIDNumber,
		}
		// Try to extract country from additional info
		if record.NationalIDAdditionalInfo != "" {
			govID.Country = extractCountryFromDetails(record.NationalIDAdditionalInfo)
		}
		govIDs = append(govIDs, govID)
	}

	if len(govIDs) > 0 {
		person.GovernmentIDs = govIDs
	}

	return person
}

// parseUKDate parses dates in UK format (DD/MM/YYYY) or partial dates
func parseUKDate(s string) *time.Time {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	// Try full date DD/MM/YYYY
	if t, err := time.Parse("02/01/2006", s); err == nil {
		return &t
	}

	// Try partial date with "dd" placeholder (dd/mm/yyyy)
	// e.g., "dd/mm/1945" means only year is known
	s = strings.ReplaceAll(s, "dd", "01")
	s = strings.ReplaceAll(s, "mm", "01")
	if t, err := time.Parse("02/01/2006", s); err == nil {
		return &t
	}

	return nil
}

// extractCountryFromDetails tries to extract a country name from passport/ID details
func extractCountryFromDetails(details string) string {
	details = strings.TrimSpace(details)
	if details == "" {
		return ""
	}

	// Common patterns: "Mali number", "issued in UK", country name at start
	lowerDetails := strings.ToLower(details)

	// Check if it starts with a country name followed by common words
	commonSuffixes := []string{" number", " passport", " id", " issued", ","}
	for _, suffix := range commonSuffixes {
		if idx := strings.Index(lowerDetails, suffix); idx > 0 {
			potential := strings.TrimSpace(details[:idx])
			if len(potential) > 1 && len(potential) < 50 {
				return potential
			}
		}
	}

	// If details is short enough, it might just be a country name
	if len(details) < 30 && !strings.Contains(details, " ") {
		return details
	}

	return ""
}

func mapBusiness(record SanctionsListRecord) *search.Business {
	business := &search.Business{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		business.Name = strings.TrimSpace(record.Names[0])
	}

	business.AltNames = collectAltNames(record)

	return business
}

func mapVessel(record SanctionsListRecord) *search.Vessel {
	vessel := &search.Vessel{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		vessel.Name = strings.TrimSpace(record.Names[0])
	}

	vessel.AltNames = collectAltNames(record)

	// IMO number
	if record.IMONumber != "" {
		vessel.IMONumber = record.IMONumber
	}

	// Vessel type
	if record.VesselType != "" {
		vessel.Type = search.VesselType(record.VesselType)
	}

	// Tonnage
	if record.Tonnage != "" {
		if tonnage, err := strconv.Atoi(record.Tonnage); err == nil {
			vessel.Tonnage = tonnage
		}
	}

	// Flag - prefer vessel-specific flag, fall back to address country
	if record.VesselFlag != "" {
		vessel.Flag = norm.Country(record.VesselFlag)
	} else if len(record.AddressCountries) > 0 && record.AddressCountries[0] != "" {
		vessel.Flag = norm.Country(record.AddressCountries[0])
	}

	// Owner
	if record.VesselOwner != "" {
		vessel.Owner = record.VesselOwner
	}

	return vessel
}

func mapAddresses(record SanctionsListRecord) []search.Address {
	if len(record.Addresses) == 0 {
		return nil
	}

	var addresses []search.Address

	// Each address in record.Addresses is already a combined string
	for i, addrStr := range record.Addresses {
		if strings.TrimSpace(addrStr) == "" {
			continue
		}

		addr := search.Address{
			Line1: addrStr,
		}

		// Add country if available (parallel array)
		if i < len(record.AddressCountries) && record.AddressCountries[i] != "" {
			addr.Country = norm.Country(record.AddressCountries[i])
		}

		// Add state/locality if available
		if i < len(record.StateLocalities) && record.StateLocalities[i] != "" {
			addr.State = record.StateLocalities[i]
		}

		// Add postal code if available (parallel array)
		if i < len(record.AddressPostalCodes) && record.AddressPostalCodes[i] != "" {
			addr.PostalCode = record.AddressPostalCodes[i]
		}

		addresses = append(addresses, addr)
	}

	return addresses
}

func mapSanctionsInfo(record SanctionsListRecord) *search.SanctionsInfo {
	info := &search.SanctionsInfo{}
	hasContent := false

	// Add Regime as a program
	if record.Regime != "" {
		info.Programs = []string{record.Regime}
		hasContent = true
	}

	// Add UN Reference to description
	if record.UNReferenceNumber != "" {
		info.Description = "UN Reference: " + record.UNReferenceNumber
		hasContent = true
	}

	if !hasContent {
		return nil
	}
	return info
}
