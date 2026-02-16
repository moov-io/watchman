// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestToEntity_Individual(t *testing.T) {
	entityType := UKSLIndividual
	record := SanctionsListRecord{
		UniqueID:            "UK-123",
		Names:               []string{"John SMITH", "Johnny SMITH"},
		NameTitle:           "Dr",
		EntityType:          &entityType,
		CountryOfBirth:      "United Kingdom",
		NonLatinScriptNames: []string{"Иван СМИРНОВ"},
	}

	entity := ToEntity(record)

	require.Equal(t, search.SourceUKCSL, entity.Source)
	require.Equal(t, "UK-123", entity.SourceID)
	require.Equal(t, search.EntityPerson, entity.Type)
	require.NotNil(t, entity.Person)
	require.Equal(t, "John SMITH", entity.Person.Name)
	require.Equal(t, "John SMITH", entity.Name)
	require.Contains(t, entity.Person.AltNames, "Johnny SMITH")
	require.Contains(t, entity.Person.AltNames, "Иван СМИРНОВ")
	require.Equal(t, []string{"Dr"}, entity.Person.Titles)
	require.Equal(t, "United Kingdom", entity.Person.PlaceOfBirth)
}

func TestToEntity_Entity(t *testing.T) {
	entityType := UKSLEntity
	record := SanctionsListRecord{
		UniqueID:            "UK-456",
		Names:               []string{"ACME Corporation", "ACME Corp"},
		EntityType:          &entityType,
		NonLatinScriptNames: []string{"شركة أكمي"},
	}

	entity := ToEntity(record)

	require.Equal(t, search.SourceUKCSL, entity.Source)
	require.Equal(t, "UK-456", entity.SourceID)
	require.Equal(t, search.EntityBusiness, entity.Type)
	require.NotNil(t, entity.Business)
	require.Equal(t, "ACME Corporation", entity.Business.Name)
	require.Equal(t, "ACME Corporation", entity.Name)
	require.Contains(t, entity.Business.AltNames, "ACME Corp")
	require.Contains(t, entity.Business.AltNames, "شركة أكمي")
}

func TestToEntity_Ship(t *testing.T) {
	entityType := UKSLShip
	record := SanctionsListRecord{
		UniqueID:         "UK-789",
		Names:            []string{"MV EXAMPLE", "EXAMPLE VESSEL"},
		EntityType:       &entityType,
		AddressCountries: []string{"Panama"},
	}

	entity := ToEntity(record)

	require.Equal(t, search.SourceUKCSL, entity.Source)
	require.Equal(t, "UK-789", entity.SourceID)
	require.Equal(t, search.EntityVessel, entity.Type)
	require.NotNil(t, entity.Vessel)
	require.Equal(t, "MV EXAMPLE", entity.Vessel.Name)
	require.Equal(t, "MV EXAMPLE", entity.Name)
	require.Contains(t, entity.Vessel.AltNames, "EXAMPLE VESSEL")
	require.Equal(t, "Panama", entity.Vessel.Flag)
}

func TestToEntity_NilEntityType(t *testing.T) {
	record := SanctionsListRecord{
		UniqueID: "UK-000",
		Names:    []string{"Unknown Entity"},
	}

	entity := ToEntity(record)

	require.Equal(t, search.SourceUKCSL, entity.Source)
	require.Equal(t, "UK-000", entity.SourceID)
	require.Empty(t, entity.Type)
	require.Nil(t, entity.Person)
	require.Nil(t, entity.Business)
	require.Nil(t, entity.Vessel)
}

func TestMapPerson(t *testing.T) {
	record := SanctionsListRecord{
		Names:               []string{"Jane DOE", "Janet DOE"},
		NameTitle:           "Mrs",
		CountryOfBirth:      "France",
		NonLatinScriptNames: []string{"ジェーン・ドウ"},
	}

	person := mapPerson(record)

	require.Equal(t, "Jane DOE", person.Name)
	require.Len(t, person.AltNames, 2)
	require.Contains(t, person.AltNames, "Janet DOE")
	require.Contains(t, person.AltNames, "ジェーン・ドウ")
	require.Equal(t, []string{"Mrs"}, person.Titles)
	require.Equal(t, "France", person.PlaceOfBirth)
}

func TestMapPerson_WithNewFields(t *testing.T) {
	record := SanctionsListRecord{
		Names:                    []string{"John SMITH"},
		DOB:                      "15/03/1980",
		Gender:                   "Male",
		TownOfBirth:              "London",
		CountryOfBirth:           "United Kingdom",
		PassportNumber:           "AB123456",
		PassportAdditionalInfo:   "UK passport",
		NationalIDNumber:         "ID999888",
		NationalIDAdditionalInfo: "British ID",
	}

	person := mapPerson(record)

	// Check DOB parsing
	require.NotNil(t, person.BirthDate)
	require.Equal(t, 1980, person.BirthDate.Year())
	require.Equal(t, 3, int(person.BirthDate.Month()))
	require.Equal(t, 15, person.BirthDate.Day())

	// Check Gender
	require.Equal(t, search.GenderMale, person.Gender)

	// Check PlaceOfBirth combines town and country
	require.Equal(t, "London, United Kingdom", person.PlaceOfBirth)

	// Check GovernmentIDs
	require.Len(t, person.GovernmentIDs, 2)

	// Passport
	var passportFound, nationalIDFound bool
	for _, id := range person.GovernmentIDs {
		if id.Type == search.GovernmentIDPassport {
			passportFound = true
			require.Equal(t, "AB123456", id.Identifier)
			require.Equal(t, "UK", id.Country)
		}
		if id.Type == search.GovernmentIDNational {
			nationalIDFound = true
			require.Equal(t, "ID999888", id.Identifier)
			require.Equal(t, "British", id.Country)
		}
	}
	require.True(t, passportFound, "Passport not found in GovernmentIDs")
	require.True(t, nationalIDFound, "National ID not found in GovernmentIDs")
}

func TestMapPerson_PartialDOB(t *testing.T) {
	// Test partial date with dd/mm placeholders
	record := SanctionsListRecord{
		Names: []string{"Test Person"},
		DOB:   "dd/mm/1958",
	}

	person := mapPerson(record)

	require.NotNil(t, person.BirthDate)
	require.Equal(t, 1958, person.BirthDate.Year())
}

func TestMapPerson_FemaleGender(t *testing.T) {
	record := SanctionsListRecord{
		Names:  []string{"Jane DOE"},
		Gender: "Female",
	}

	person := mapPerson(record)
	require.Equal(t, search.GenderFemale, person.Gender)
}

func TestParseUKDate(t *testing.T) {
	tests := []struct {
		input    string
		wantNil  bool
		wantYear int
	}{
		{"15/03/1980", false, 1980},
		{"01/12/2000", false, 2000},
		{"dd/mm/1958", false, 1958},
		{"", true, 0},
		{"invalid", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseUKDate(tt.input)
			if tt.wantNil {
				require.Nil(t, result)
			} else {
				require.NotNil(t, result)
				require.Equal(t, tt.wantYear, result.Year())
			}
		})
	}
}

func TestExtractCountryFromDetails(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"UK passport", "UK"},
		{"Mali number", "Mali"},
		{"British ID", "British"},
		{"France issued", "France"},
		{"", ""},
		{"Some very long description that does not have a country pattern", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := extractCountryFromDetails(tt.input)
			require.Equal(t, tt.want, result)
		})
	}
}

func TestMapPerson_EmptyFields(t *testing.T) {
	record := SanctionsListRecord{}

	person := mapPerson(record)

	require.Empty(t, person.Name)
	require.Empty(t, person.AltNames)
	require.Empty(t, person.Titles)
	require.Empty(t, person.PlaceOfBirth)
}

func TestMapBusiness(t *testing.T) {
	record := SanctionsListRecord{
		Names:               []string{"Test Company Ltd", "Test Co"},
		NonLatinScriptNames: []string{"测试公司"},
	}

	business := mapBusiness(record)

	require.Equal(t, "Test Company Ltd", business.Name)
	require.Len(t, business.AltNames, 2)
	require.Contains(t, business.AltNames, "Test Co")
	require.Contains(t, business.AltNames, "测试公司")
}

func TestMapVessel(t *testing.T) {
	record := SanctionsListRecord{
		Names:            []string{"SS FREEDOM", "FREEDOM"},
		AddressCountries: []string{"Liberia"},
	}

	vessel := mapVessel(record)

	require.Equal(t, "SS FREEDOM", vessel.Name)
	require.Len(t, vessel.AltNames, 1)
	require.Contains(t, vessel.AltNames, "FREEDOM")
	require.Equal(t, "Liberia", vessel.Flag)
}

func TestMapVessel_NoFlag(t *testing.T) {
	record := SanctionsListRecord{
		Names: []string{"VESSEL ONE"},
	}

	vessel := mapVessel(record)

	require.Equal(t, "VESSEL ONE", vessel.Name)
	require.Empty(t, vessel.Flag)
}

func TestMapVessel_WithNewFields(t *testing.T) {
	record := SanctionsListRecord{
		Names:       []string{"MV CARGO SHIP"},
		IMONumber:   "9876543",
		VesselType:  "Bulk Carrier",
		Tonnage:     "50000",
		VesselFlag:  "Panama",
		VesselOwner: "Test Shipping Co",
	}

	vessel := mapVessel(record)

	require.Equal(t, "MV CARGO SHIP", vessel.Name)
	require.Equal(t, "9876543", vessel.IMONumber)
	require.Equal(t, search.VesselType("Bulk Carrier"), vessel.Type)
	require.Equal(t, 50000, vessel.Tonnage)
	require.Equal(t, "Panama", vessel.Flag)
	require.Equal(t, "Test Shipping Co", vessel.Owner)
}

func TestMapVessel_InvalidTonnage(t *testing.T) {
	record := SanctionsListRecord{
		Names:   []string{"TEST VESSEL"},
		Tonnage: "not-a-number",
	}

	vessel := mapVessel(record)

	require.Equal(t, "TEST VESSEL", vessel.Name)
	require.Equal(t, 0, vessel.Tonnage) // Should be 0 when parsing fails
}

func TestMapVessel_FlagPrecedence(t *testing.T) {
	// VesselFlag should take precedence over AddressCountries
	record := SanctionsListRecord{
		Names:            []string{"TEST VESSEL"},
		VesselFlag:       "Liberia",
		AddressCountries: []string{"Panama"},
	}

	vessel := mapVessel(record)

	require.Equal(t, "Liberia", vessel.Flag)
}

func TestMapAddresses(t *testing.T) {
	record := SanctionsListRecord{
		Addresses:        []string{"123 Main St, London", "456 Side Ave, Manchester"},
		AddressCountries: []string{"United Kingdom", "United Kingdom"},
		StateLocalities:  []string{"Greater London", "Greater Manchester"},
	}

	addresses := mapAddresses(record)

	require.Len(t, addresses, 2)
	require.Equal(t, "123 Main St, London", addresses[0].Line1)
	require.Equal(t, "United Kingdom", addresses[0].Country)
	require.Equal(t, "Greater London", addresses[0].State)
	require.Equal(t, "456 Side Ave, Manchester", addresses[1].Line1)
	require.Equal(t, "United Kingdom", addresses[1].Country)
	require.Equal(t, "Greater Manchester", addresses[1].State)
}

func TestMapAddresses_Empty(t *testing.T) {
	record := SanctionsListRecord{}

	addresses := mapAddresses(record)

	require.Nil(t, addresses)
}

func TestMapAddresses_SkipsEmptyStrings(t *testing.T) {
	record := SanctionsListRecord{
		Addresses:        []string{"", "Valid Address", "  "},
		AddressCountries: []string{"", "Germany", ""},
	}

	addresses := mapAddresses(record)

	require.Len(t, addresses, 1)
	require.Equal(t, "Valid Address", addresses[0].Line1)
	require.Equal(t, "Germany", addresses[0].Country)
}

func TestMapAddresses_WithPostalCode(t *testing.T) {
	record := SanctionsListRecord{
		Addresses:          []string{"123 Main St, London"},
		AddressCountries:   []string{"United Kingdom"},
		StateLocalities:    []string{"Greater London"},
		AddressPostalCodes: []string{"SW1A 1AA"},
	}

	addresses := mapAddresses(record)

	require.Len(t, addresses, 1)
	require.Equal(t, "123 Main St, London", addresses[0].Line1)
	require.Equal(t, "United Kingdom", addresses[0].Country)
	require.Equal(t, "Greater London", addresses[0].State)
	require.Equal(t, "SW1A 1AA", addresses[0].PostalCode)
}

func TestMapSanctionsInfo(t *testing.T) {
	record := SanctionsListRecord{
		UNReferenceNumber: "UN-2024-001",
	}

	info := mapSanctionsInfo(record)

	require.NotNil(t, info)
	require.Equal(t, "UN Reference: UN-2024-001", info.Description)
}

func TestMapSanctionsInfo_Empty(t *testing.T) {
	record := SanctionsListRecord{}

	info := mapSanctionsInfo(record)

	require.Nil(t, info)
}

func TestMapSanctionsInfo_WithRegime(t *testing.T) {
	record := SanctionsListRecord{
		Regime: "The Russia (Sanctions) (EU Exit) Regulations 2019",
	}

	info := mapSanctionsInfo(record)

	require.NotNil(t, info)
	require.Len(t, info.Programs, 1)
	require.Equal(t, "The Russia (Sanctions) (EU Exit) Regulations 2019", info.Programs[0])
}

func TestMapSanctionsInfo_WithBoth(t *testing.T) {
	record := SanctionsListRecord{
		UNReferenceNumber: "UN-2024-001",
		Regime:            "Test Regime",
	}

	info := mapSanctionsInfo(record)

	require.NotNil(t, info)
	require.Equal(t, "UN Reference: UN-2024-001", info.Description)
	require.Len(t, info.Programs, 1)
	require.Equal(t, "Test Regime", info.Programs[0])
}

func TestConvertSanctionsListData(t *testing.T) {
	individualType := UKSLIndividual
	entityType := UKSLEntity

	records := []SanctionsListRecord{
		{
			UniqueID:   "UK-001",
			Names:      []string{"Person One"},
			EntityType: &individualType,
		},
		{
			UniqueID:   "UK-002",
			Names:      []string{"Company Two"},
			EntityType: &entityType,
		},
	}

	entities := ConvertSanctionsListData(records)

	require.Len(t, entities, 2)
	require.Equal(t, "UK-001", entities[0].SourceID)
	require.Equal(t, search.EntityPerson, entities[0].Type)
	require.Equal(t, "UK-002", entities[1].SourceID)
	require.Equal(t, search.EntityBusiness, entities[1].Type)
}

func TestConvertSanctionsListData_Empty(t *testing.T) {
	entities := ConvertSanctionsListData(nil)
	require.Empty(t, entities)

	entities = ConvertSanctionsListData([]SanctionsListRecord{})
	require.Empty(t, entities)
}
