// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToEntity_Person(t *testing.T) {
	record := CSLRecord{
		EntityLogicalID:   123,
		EntitySubjectTypeCode: "person",
		NameAliasWholeNames: []string{
			"John Doe",
			"Johnny D",
			"J. Doe",
		},
		NameAliasGenders: []string{"M"},
		NameAliasTitles:  []string{"Dr.", "Prof."},
		BirthDates:       []string{"1985-06-15"},
		BirthCities:      []string{"Berlin"},
		BirthCountries:   []string{"Germany"},
		EntityRemark:     "Test remark",
		AddressStreets:   []string{"123 Main St"},
		AddressCities:    []string{"Munich"},
		AddressZipCodes:  []string{"80331"},
		AddressCountryDescriptions: []string{"Germany"},
		Identifications: []IdentificationInfo{
			{
				Number:          "AB123456",
				TypeCode:        "passport",
				TypeDescription: "National passport",
				CountryDesc:     "Germany",
			},
		},
	}

	entity := ToEntity(record)

	assert.Equal(t, search.SourceEUCSL, entity.Source)
	assert.Equal(t, "123", entity.SourceID)
	assert.Equal(t, search.EntityPerson, entity.Type)
	assert.Equal(t, "John Doe", entity.Name)

	require.NotNil(t, entity.Person)
	assert.Equal(t, "John Doe", entity.Person.Name)
	assert.Equal(t, []string{"Johnny D", "J. Doe"}, entity.Person.AltNames)
	assert.Equal(t, search.GenderMale, entity.Person.Gender)
	assert.Equal(t, []string{"Dr.", "Prof."}, entity.Person.Titles)
	assert.Equal(t, "Berlin, Germany", entity.Person.PlaceOfBirth)

	require.NotNil(t, entity.Person.BirthDate)
	expectedDate := time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, expectedDate, *entity.Person.BirthDate)

	// Check government IDs
	require.Len(t, entity.Person.GovernmentIDs, 1)
	assert.Equal(t, "AB123456", entity.Person.GovernmentIDs[0].Identifier)
	assert.Equal(t, search.GovernmentIDPassport, entity.Person.GovernmentIDs[0].Type)
	assert.Equal(t, "National passport", entity.Person.GovernmentIDs[0].Name)

	// Check addresses
	require.Len(t, entity.Addresses, 1)
	assert.Equal(t, "123 Main St", entity.Addresses[0].Line1)
	assert.Equal(t, "Munich", entity.Addresses[0].City)
	assert.Equal(t, "80331", entity.Addresses[0].PostalCode)

	// Check sanctions info
	require.NotNil(t, entity.SanctionsInfo)
	assert.Contains(t, entity.SanctionsInfo.Description, "Test remark")
}

func TestToEntity_Enterprise(t *testing.T) {
	record := CSLRecord{
		EntityLogicalID:   456,
		EntitySubjectTypeCode: "enterprise",
		NameAliasWholeNames: []string{
			"Acme Corporation",
			"Acme Corp",
		},
		AddressStreets:             []string{"456 Business Ave"},
		AddressCities:              []string{"Paris"},
		AddressCountryDescriptions: []string{"France"},
		EntityReferenceNumber:      "EU.123.456",
	}

	entity := ToEntity(record)

	assert.Equal(t, search.SourceEUCSL, entity.Source)
	assert.Equal(t, "456", entity.SourceID)
	assert.Equal(t, search.EntityBusiness, entity.Type)
	assert.Equal(t, "Acme Corporation", entity.Name)

	require.NotNil(t, entity.Business)
	assert.Equal(t, "Acme Corporation", entity.Business.Name)
	assert.Equal(t, []string{"Acme Corp"}, entity.Business.AltNames)

	// Check addresses
	require.Len(t, entity.Addresses, 1)
	assert.Equal(t, "456 Business Ave", entity.Addresses[0].Line1)
	assert.Equal(t, "Paris", entity.Addresses[0].City)

	// Check sanctions info with reference
	require.NotNil(t, entity.SanctionsInfo)
	assert.Contains(t, entity.SanctionsInfo.Description, "Ref: EU.123.456")
}

func TestToEntity_Vessel(t *testing.T) {
	record := CSLRecord{
		EntityLogicalID:   789,
		EntitySubjectTypeCode: "unknown",
		NameAliasWholeNames: []string{
			"Sea Explorer",
			"Ocean Wanderer",
		},
		EntityRemark:               "IMO: 1234567",
		AddressCountryDescriptions: []string{"Panama"},
	}

	entity := ToEntity(record)

	assert.Equal(t, search.SourceEUCSL, entity.Source)
	assert.Equal(t, "789", entity.SourceID)
	assert.Equal(t, search.EntityVessel, entity.Type)
	assert.Equal(t, "Sea Explorer", entity.Name)

	require.NotNil(t, entity.Vessel)
	assert.Equal(t, "Sea Explorer", entity.Vessel.Name)
	assert.Equal(t, []string{"Ocean Wanderer"}, entity.Vessel.AltNames)
	assert.Equal(t, "1234567", entity.Vessel.IMONumber)
}

func TestToEntity_UnknownType(t *testing.T) {
	record := CSLRecord{
		EntityLogicalID:     999,
		EntitySubjectTypeCode: "other",
		NameAliasWholeNames: []string{"Unknown Entity"},
	}

	entity := ToEntity(record)

	assert.Equal(t, search.SourceEUCSL, entity.Source)
	assert.Equal(t, "999", entity.SourceID)
	assert.Equal(t, search.EntityOrganization, entity.Type)
	assert.Equal(t, "Unknown Entity", entity.Name)

	require.NotNil(t, entity.Organization)
	assert.Equal(t, "Unknown Entity", entity.Organization.Name)
}

func TestMapAddresses_MultipleAddresses(t *testing.T) {
	record := CSLRecord{
		AddressStreets:             []string{"Street 1", "Street 2"},
		AddressCities:              []string{"City 1", "City 2"},
		AddressZipCodes:            []string{"11111", "22222"},
		AddressPoBoxes:             []string{"", "PO123"},
		AddressCountryDescriptions: []string{"Country 1", "Country 2"},
	}

	addresses := mapAddresses(record)

	require.Len(t, addresses, 2)

	assert.Equal(t, "Street 1", addresses[0].Line1)
	assert.Equal(t, "City 1", addresses[0].City)
	assert.Equal(t, "11111", addresses[0].PostalCode)

	assert.Equal(t, "Street 2", addresses[1].Line1)
	assert.Equal(t, "PO Box PO123", addresses[1].Line2)
	assert.Equal(t, "City 2", addresses[1].City)
	assert.Equal(t, "22222", addresses[1].PostalCode)
}

func TestMapAddresses_Empty(t *testing.T) {
	record := CSLRecord{}
	addresses := mapAddresses(record)
	assert.Nil(t, addresses)
}

func TestParseEUDate(t *testing.T) {
	tests := []struct {
		input    string
		expected *time.Time
	}{
		{
			input:    "1985-06-15",
			expected: ptr(time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)),
		},
		{
			input:    "28/10/2022",
			expected: ptr(time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)),
		},
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "invalid",
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := parseEUDate(tc.input)
			if tc.expected == nil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				assert.Equal(t, *tc.expected, *result)
			}
		})
	}
}

func TestExtractIMONumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"IMO: 1234567", "1234567"},
		{"IMO1234567", "1234567"},
		{"imo: 9876543", "9876543"},
		{"The vessel has IMO 1111111 registered", "1111111"},
		{"No IMO number here", ""},
		{"IMO 123", ""}, // Too short
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := extractIMONumber(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMapGender(t *testing.T) {
	tests := []struct {
		input    string
		expected search.Gender
	}{
		{"M", search.GenderMale},
		{"m", search.GenderMale},
		{"MALE", search.GenderMale},
		{"Male", search.GenderMale},
		{"F", search.GenderFemale},
		{"f", search.GenderFemale},
		{"FEMALE", search.GenderFemale},
		{"Female", search.GenderFemale},
		{"", search.GenderUnknown},
		{"X", search.GenderUnknown},
		{"other", search.GenderUnknown},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := mapGender(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMapGovernmentIDs(t *testing.T) {
	tests := []struct {
		name     string
		input    []IdentificationInfo
		expected []search.GovernmentID
	}{
		{
			name:     "empty",
			input:    nil,
			expected: nil,
		},
		{
			name: "passport",
			input: []IdentificationInfo{
				{
					Number:          "AB123456",
					TypeCode:        "passport",
					TypeDescription: "National passport",
					CountryDesc:     "Germany",
				},
			},
			expected: []search.GovernmentID{
				{
					Identifier: "AB123456",
					Type:       search.GovernmentIDPassport,
					Name:       "National passport",
					Country:    "Germany",
				},
			},
		},
		{
			name: "national_id",
			input: []IdentificationInfo{
				{
					Number:          "123456789",
					TypeCode:        "id",
					TypeDescription: "National identification card",
					CountryDesc:     "France",
				},
			},
			expected: []search.GovernmentID{
				{
					Identifier: "123456789",
					Type:       search.GovernmentIDNational,
					Name:       "National identification card",
					Country:    "France",
				},
			},
		},
		{
			name: "number_with_type_in_parens",
			input: []IdentificationInfo{
				{
					Number:      "488555 (passport-National passport)",
					CountryISO:  "RU",
					CountryDesc: "",
				},
			},
			expected: []search.GovernmentID{
				{
					Identifier: "488555",
					Type:       search.GovernmentIDPersonalID,
					Name:       "Unknown",
					Country:    "Russia",
				},
			},
		},
		{
			name: "multiple_ids",
			input: []IdentificationInfo{
				{
					Number:          "PASS123",
					TypeCode:        "passport",
					TypeDescription: "Passport",
					CountryDesc:     "Iraq",
				},
				{
					Number:          "DL456789",
					TypeDescription: "Driver's license",
					CountryDesc:     "Iraq",
				},
			},
			expected: []search.GovernmentID{
				{
					Identifier: "PASS123",
					Type:       search.GovernmentIDPassport,
					Name:       "Passport",
					Country:    "Iraq",
				},
				{
					Identifier: "DL456789",
					Type:       search.GovernmentIDDriversLicense,
					Name:       "Driver's license",
					Country:    "Iraq",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mapGovernmentIDs(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestExtractIDNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"AB123456", "AB123456"},
		{"488555 (passport-National passport)", "488555"},
		{"  123456  ", "123456"},
		{"", ""},
		{"M0003264580 (other-Other identification number)", "M0003264580"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := extractIDNumber(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMapIDType(t *testing.T) {
	tests := []struct {
		typeCode string
		typeDesc string
		expected search.GovernmentIDType
	}{
		{"passport", "", search.GovernmentIDPassport},
		{"id", "", search.GovernmentIDNational},
		{"", "National passport", search.GovernmentIDPassport},
		{"", "Diplomatic passport", search.GovernmentIDDiplomaticPass},
		{"", "National identification card", search.GovernmentIDNational},
		{"", "Driver's license", search.GovernmentIDDriversLicense},
		{"", "Tax ID", search.GovernmentIDTax},
		{"", "Birth certificate", search.GovernmentIDBirthCert},
		{"", "Business registration", search.GovernmentIDBusinessRegisration},
		{"", "Other identification", search.GovernmentIDPersonalID},
		{"", "", search.GovernmentIDPersonalID},
	}

	for _, tc := range tests {
		name := tc.typeCode
		if name == "" {
			name = tc.typeDesc
		}
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			result := mapIDType(tc.typeCode, tc.typeDesc)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestConvertEUCSLData(t *testing.T) {
	records := []CSLRecord{
		{
			EntityLogicalID:       1,
			EntitySubjectTypeCode: "person",
			NameAliasWholeNames:   []string{"Person One"},
		},
		{
			EntityLogicalID:       2,
			EntitySubjectTypeCode: "enterprise",
			NameAliasWholeNames:   []string{"Company Two"},
		},
	}

	entities := ConvertEUCSLData(records)

	require.Len(t, entities, 2)
	assert.Equal(t, "Person One", entities[0].Name)
	assert.Equal(t, search.EntityPerson, entities[0].Type)
	assert.Equal(t, "Company Two", entities[1].Name)
	assert.Equal(t, search.EntityBusiness, entities[1].Type)
}

func TestConvertEUCSLData_WithRealData(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", "eu_csl.csv"))
	if err != nil {
		t.Skip("test data not available")
	}

	records, _, err := ParseEU(fd)
	require.NoError(t, err)
	require.NotEmpty(t, records)

	entities := ConvertEUCSLData(records)
	require.NotEmpty(t, entities)

	// Count entity types and features
	var persons, businesses, vessels, orgs int
	var withGender, withGovIDs int
	for _, e := range entities {
		switch e.Type {
		case search.EntityPerson:
			persons++
			if e.Person != nil {
				if e.Person.Gender != "" && e.Person.Gender != search.GenderUnknown {
					withGender++
				}
				if len(e.Person.GovernmentIDs) > 0 {
					withGovIDs++
				}
			}
		case search.EntityBusiness:
			businesses++
		case search.EntityVessel:
			vessels++
		case search.EntityOrganization:
			orgs++
		}
	}

	t.Logf("Converted %d entities: %d persons, %d businesses, %d vessels, %d organizations",
		len(entities), persons, businesses, vessels, orgs)
	t.Logf("Persons with gender: %d, with government IDs: %d", withGender, withGovIDs)

	// Should have at least some persons
	assert.Greater(t, persons, 0, "expected at least some person entities")

	// Should have some persons with gender (we know there's data)
	assert.Greater(t, withGender, 0, "expected at least some persons with gender")

	// Should have some persons with government IDs
	assert.Greater(t, withGovIDs, 0, "expected at least some persons with government IDs")
}

func ptr[T any](v T) *T {
	return &v
}
