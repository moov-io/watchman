// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/pkg/address"
	"github.com/moov-io/watchman/pkg/csl_us/gen/ENHANCED_XML"
	"github.com/moov-io/watchman/pkg/search"
)

func ConvertSanctionsData(data *ENHANCED_XML.SanctionsData) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, entity := range data.Entities.Entity {
		out = append(out, ToEntity(entity))
	}
	return out
}

// ToEntity converts an OFAC XML entity to a search Entity
func ToEntity(src ENHANCED_XML.EntitiesEntity) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Source:     search.SourceUSCSL,
		SourceID:   strconv.Itoa(src.Id),
		SourceData: src,
	}

	// Map entity type and create corresponding struct
	// Map based on SDN TYPE reference values
	switch src.GeneralInfo.EntityType.RefId {
	case 600: // Individual
		entity.Type = search.EntityPerson
		entity.Person = mapPerson(src)
	case 601: // Entity
		// For entities, we need to determine if it's a business or organization
		// Could use Organization Type field or other attributes to distinguish
		entity.Type = search.EntityBusiness
		entity.Business = mapBusiness(src)
	case 602: // Vessel
		entity.Type = search.EntityVessel
		entity.Vessel = mapVessel(src)
	case 91120: // Aircraft
		entity.Type = search.EntityAircraft
		entity.Aircraft = mapAircraft(src)
	}

	// Map common fields
	entity.Name = getPrimaryName(src.Names, src.GeneralInfo.EntityType.Text)
	entity.Contact = mapContactInfo(src)
	entity.Addresses = mapAddresses(src.Addresses)
	entity.Affiliations = mapAffiliations(src.Relationships)
	entity.SanctionsInfo = mapSanctionsInfo(src)

	return entity
}

// getPrimaryName returns the primary formatted name from translations
func getPrimaryName(names ENHANCED_XML.EntityNames, entityType string) string {
	entityType = strings.ToLower(entityType)

	for _, name := range names.Name {
		if name.IsPrimary {
			// Find primary translation
			for _, translation := range name.Translations.Translation {
				if translation.IsPrimary {
					// Return formatted full name for primary translation
					if translation.FormattedFullName != "" {
						return prepare.ReorderSDNName(translation.FormattedFullName, entityType)
					}
					// If no formatted name, build from parts
					var firstName, lastName string
					for _, part := range translation.NameParts.NamePart {
						switch part.Type.RefId {
						case 1520: // Last Name
							lastName = part.Value
						case 1521: // First Name
							firstName = part.Value
						}
					}
					if lastName != "" && firstName != "" {
						return fmt.Sprintf("%s %s", firstName, lastName)
					}
					return firstName + lastName // Return whatever we have
				}
			}
		}
	}
	return ""
}

// getAllNames returns all names including aliases
func getAllNames(names ENHANCED_XML.EntityNames, entityType string) []string {
	entityType = strings.ToLower(entityType)

	var allNames []string
	for _, name := range names.Name {
		// Skip primary name as it's handled separately
		if name.IsPrimary {
			continue
		}

		// Check alias type by refId
		var isAlias bool
		switch name.AliasType.RefId {
		case 1400: // A.K.A.
			isAlias = true
		case 1401: // F.K.A
			isAlias = true
		case 1402: // N.K.A
			isAlias = true
		}

		if isAlias {
			for _, translation := range name.Translations.Translation {
				if translation.IsPrimary && translation.FormattedFullName != "" {
					name := prepare.ReorderSDNName(translation.FormattedFullName, entityType)

					allNames = append(allNames, name)
				}
			}
		}
	}
	return allNames
}

func mapPerson(src ENHANCED_XML.EntitiesEntity) *search.Person {
	person := &search.Person{
		Name:     getPrimaryName(src.Names, "individual"),
		AltNames: getAllNames(src.Names, "individual"),
	}

	// Map features for birth date, death date, gender
	mapPersonFeatures(src.Features, person)

	// Map government IDs
	if src.IdentityDocuments != nil {
		person.GovernmentIDs = mapGovernmentIDs(src.IdentityDocuments)
	}

	person.Titles = mapTitles(src)

	return person
}

func mapPersonFeatures(features *ENHANCED_XML.EntityFeatures, person *search.Person) {
	if features == nil {
		return
	}

	for _, feature := range features.Feature {
		if !feature.IsPrimary {
			continue
		}

		switch feature.Type.FeatureTypeId {
		case 8: // Birthdate
			if feature.ValueDate != nil {
				if date, err := time.Parse("2006-01-02", feature.ValueDate.FromDateBegin); err == nil {
					person.BirthDate = &date
				}
			}
		// case 9: // Place of Birth
		// 	person.PlaceOfBirth = feature.Value
		case 91526: // Gender - Male
			person.Gender = search.GenderMale
		case 91527: // Gender - Female
			person.Gender = search.GenderFemale
		}
	}
}

func mapBusiness(src ENHANCED_XML.EntitiesEntity) *search.Business {
	business := &search.Business{
		Name:     getPrimaryName(src.Names, "business"),
		AltNames: getAllNames(src.Names, "business"),
	}

	// Map features for registration dates and identifiers
	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			switch feature.Type.Text {
			case "Registration Date":
				if date := mapDate(feature.ValueDate); date != nil {
					business.Created = date
				}
			case "Dissolution Date":
				if date := mapDate(feature.ValueDate); date != nil {
					business.Dissolved = date
				}
			}
		}
	}

	// Map business identifiers
	if src.IdentityDocuments != nil {
		business.GovernmentIDs = mapGovernmentIDs(src.IdentityDocuments)
	}

	return business
}

func mapOrganization(src ENHANCED_XML.EntitiesEntity) *search.Organization {
	org := &search.Organization{
		Name:     getPrimaryName(src.Names, "organization"),
		AltNames: getAllNames(src.Names, "organization"),
	}

	// Map features similar to business
	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			switch feature.Type.Text {
			case "Formation Date":
				if date := mapDate(feature.ValueDate); date != nil {
					org.Created = date
				}
			case "Dissolution Date":
				if date := mapDate(feature.ValueDate); date != nil {
					org.Dissolved = date
				}
			}
		}
	}

	// Map organization identifiers
	if src.IdentityDocuments != nil {
		org.GovernmentIDs = mapGovernmentIDs(src.IdentityDocuments)
	}

	return org
}

func mapAircraft(src ENHANCED_XML.EntitiesEntity) *search.Aircraft {
	aircraft := &search.Aircraft{
		Name:     getPrimaryName(src.Names, "aircraft"),
		AltNames: getAllNames(src.Names, "aircraft"),
	}

	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			switch feature.Type.Text {
			case "Aircraft Type":
				aircraft.Type = mapAircraftType(feature.Value)
			case "Construction Date":
				if date := mapDate(feature.ValueDate); date != nil {
					aircraft.Built = date
				}
			case "Aircraft Model":
				aircraft.Model = feature.Value
			case "Serial Number":
				aircraft.SerialNumber = feature.Value
			case "Registration Country":
				aircraft.Flag = feature.Value // Assuming country code is stored in value
			case "ICAO Number":
				aircraft.ICAOCode = feature.Value
			}
		}
	}

	return aircraft
}

func mapVessel(src ENHANCED_XML.EntitiesEntity) *search.Vessel {
	vessel := &search.Vessel{
		Name:     getPrimaryName(src.Names, "vessel"),
		AltNames: getAllNames(src.Names, "vessel"),
	}

	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			switch feature.Type.Text {
			case "Vessel Type":
				vessel.Type = mapVesselType(feature.Value)
			case "Build Date":
				if date := mapDate(feature.ValueDate); date != nil {
					vessel.Built = date
				}
			case "IMO Number":
				vessel.IMONumber = feature.Value
			case "MMSI Number":
				vessel.MMSI = feature.Value
			case "Call Sign":
				vessel.CallSign = feature.Value
			case "Tonnage":
				// Convert string to int, handle error in production code
				vessel.Tonnage = parseIntOrZero(feature.Value)
			case "GRT":
				vessel.GrossRegisteredTonnage = parseIntOrZero(feature.Value)
			case "Vessel Flag":
				vessel.Flag = feature.Value // Assuming country code is stored in value
			case "Vessel Owner":
				vessel.Owner = feature.Value
			}
		}
	}

	return vessel
}

func mapContactInfo(src ENHANCED_XML.EntitiesEntity) search.ContactInfo {
	info := search.ContactInfo{}

	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			switch feature.Type.Text {
			case "Email Address":
				info.EmailAddresses = append(info.EmailAddresses, feature.Value)
			case "Phone Number":
				info.PhoneNumbers = append(info.PhoneNumbers, feature.Value)
			case "Fax":
				info.FaxNumbers = append(info.FaxNumbers, feature.Value)
			case "Website":
				info.Websites = append(info.Websites, feature.Value)
			}
		}
	}

	return info
}

func mapAddresses(addresses *ENHANCED_XML.EntityAddresses) []search.Address {
	if addresses == nil {
		return nil
	}

	var result []search.Address
	for _, addr := range addresses.Address {
		var mappedAddr search.Address

		// Get primary translation
		var translation *ENHANCED_XML.AddressTranslationsTranslation
		for _, t := range addr.Translations.Translation {
			if t.IsPrimary {
				translation = &t
				break
			}
		}
		if translation != nil && translation.AddressParts != nil {
			for _, part := range translation.AddressParts.AddressPart {
				switch part.Type.RefId {
				case 1451: // ADDRESS1
					mappedAddr.Line1 = part.Value
				case 1452: // ADDRESS2
					mappedAddr.Line2 = part.Value
				case 1454: // CITY
					mappedAddr.City = part.Value
				case 1456: // POSTAL CODE
					mappedAddr.PostalCode = part.Value
				case 1455: // STATE/PROVINCE
					mappedAddr.State = part.Value
				}
			}
		}
		if addr.Country != nil {
			mappedAddr.Country = addr.Country.Text
		}

		res := address.ParseAddress(mappedAddr.Format())
		if res.Line1 != "" && res.Country != "" {
			result = append(result, res)
		}
	}

	return result
}

func mapAffiliations(relationships *ENHANCED_XML.EntityRelationships) []search.Affiliation {
	if relationships == nil {
		return nil
	}

	var affiliations []search.Affiliation
	for _, rel := range relationships.Relationship {
		aff := search.Affiliation{
			EntityName: rel.RelatedEntity.Text,
			Type:       rel.Type.Text,
			Details:    rel.Comments,
		}
		affiliations = append(affiliations, aff)
	}

	return affiliations
}

func mapSanctionsInfo(src ENHANCED_XML.EntitiesEntity) *search.SanctionsInfo {
	info := &search.SanctionsInfo{
		Secondary:   false, // Default value
		Description: src.GeneralInfo.Remarks,
	}

	// Map sanctions programs
	for _, program := range src.SanctionsPrograms.SanctionsProgram {
		info.Programs = append(info.Programs, program.Text)
	}

	// Check for secondary sanctions in features
	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			if feature.Type.Text == "Secondary Sanctions Risk" {
				info.Secondary = strings.ToLower(feature.Value) == "yes"
			}
		}
	}

	return info
}

func mapTitles(src ENHANCED_XML.EntitiesEntity) []string {
	var titles []string
	if src.GeneralInfo.Title != "" {
		titles = append(titles, src.GeneralInfo.Title)
	}

	// Additional titles might be in features
	if src.Features != nil {
		for _, feature := range src.Features.Feature {
			if feature.Type.Text == "Title" {
				titles = append(titles, feature.Value)
			}
		}
	}

	return titles
}

// Helper functions

func mapDate(datePeriod *ENHANCED_XML.DatePeriodType) *time.Time {
	if datePeriod == nil {
		return nil
	}
	// Use FromDateBegin as the primary date if available
	if datePeriod.FromDateBegin != "" {
		if t, err := time.Parse("2006-01-02", datePeriod.FromDateBegin); err == nil {
			return &t
		}
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

func mapGovernmentIDs(docs *ENHANCED_XML.EntityIdentityDocuments) []search.GovernmentID {
	if docs == nil {
		return nil
	}

	var ids []search.GovernmentID
	for _, doc := range docs.IdentityDocument {
		id := search.GovernmentID{
			Type:       mapIDType(doc.Type),
			Country:    getCountryCode(doc.IssuingCountry),
			Identifier: doc.DocumentNumber,
		}
		if string(id.Type) != "" {
			ids = append(ids, id)
		}
	}
	return ids
}

// Helper to map ID document type reference IDs to GovernmentIDType values
func mapIDType(ref ENHANCED_XML.ReferenceValueReferenceType) search.GovernmentIDType {
	switch ref.RefId {
	case 1571: // Passport
		return search.GovernmentIDPassport
	case 1577: // Driver's License No.
		return search.GovernmentIDDriversLicense
	case 1584: // National ID No.
		return search.GovernmentIDNational
	case 1596: // Tax ID No.
		return search.GovernmentIDTax
	case 1572: // SSN
		return search.GovernmentIDSSN
	case 1570: // Cedula No.
		return search.GovernmentIDCedula
	case 1600: // C.U.R.P.
		return search.GovernmentIDCURP
	case 1595: // C.U.I.T.
		return search.GovernmentIDCUIT
	case 1607: // Electoral Registry No.
		return search.GovernmentIDElectoral
	case 1581, 1585, 91752, 91760, 91761:
		// Business Registration Document, etc
		return search.GovernmentIDBusinessRegisration
	case 1619: // Commercial Registry Number
		return search.GovernmentIDCommercialRegistry
	case 91759: // Birth Certificate Number
		return search.GovernmentIDBirthCert
	case 1649: // Refugee ID Card
		return search.GovernmentIDRefugee
	case 1613: // Diplomatic Passport
		return search.GovernmentIDDiplomaticPass
	case 1627: // Personal ID Card
		return search.GovernmentIDPersonalID
	default:
		return "" // Unknown ID type
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

func getCountryCode(country *ENHANCED_XML.ReferenceValueReferenceType) string {
	if country == nil {
		return ""
	}
	return country.Text
}

func parseIntOrZero(value string) int {
	// Implementation would convert string to int, returning 0 on error
	// Production code should include proper error handling
	return 0
}
