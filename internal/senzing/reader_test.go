package senzing

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestReadEntities_JSONLines(t *testing.T) {
	input := `{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"John","NAME_LAST":"Smith"}
{"DATA_SOURCE":"TEST","RECORD_ID":"1002","RECORD_TYPE":"ORGANIZATION","NAME_ORG":"Acme Corp"}`

	entities, err := ReadEntities(strings.NewReader(input), "test-source")
	require.NoError(t, err)
	require.Len(t, entities, 2)

	// Verify person
	require.Equal(t, search.EntityPerson, entities[0].Type)
	require.Equal(t, "John Smith", entities[0].Name)
	require.Equal(t, "1001", entities[0].SourceID)
	require.Equal(t, search.SourceList("test-source"), entities[0].Source)

	// Verify organization
	require.Equal(t, search.EntityBusiness, entities[1].Type)
	require.Equal(t, "Acme Corp", entities[1].Name)
	require.Equal(t, "1002", entities[1].SourceID)
}

func TestReadEntities_JSONArray(t *testing.T) {
	input := `[
		{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"Jane","NAME_LAST":"Doe"},
		{"DATA_SOURCE":"TEST","RECORD_ID":"1002","RECORD_TYPE":"ORGANIZATION","NAME_ORG":"Widget Inc"}
	]`

	entities, err := ReadEntities(strings.NewReader(input), "test-source")
	require.NoError(t, err)
	require.Len(t, entities, 2)

	require.Equal(t, search.EntityPerson, entities[0].Type)
	require.Equal(t, "Jane Doe", entities[0].Name)

	require.Equal(t, search.EntityBusiness, entities[1].Type)
	require.Equal(t, "Widget Inc", entities[1].Name)
}

func TestReadEntities_WithFeatures(t *testing.T) {
	input := `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","FEATURES":[{"RECORD_TYPE":"PERSON"},{"NAME_FIRST":"Robert","NAME_LAST":"Smith"},{"ADDR_LINE1":"123 Main St","ADDR_CITY":"Las Vegas","ADDR_STATE":"NV"}]}`

	entities, err := ReadEntities(strings.NewReader(input), "customers")
	require.NoError(t, err)
	require.Len(t, entities, 1)

	require.Equal(t, search.EntityPerson, entities[0].Type)
	require.Equal(t, "Robert Smith", entities[0].Name)
	require.Len(t, entities[0].Addresses, 1)
	require.Equal(t, "123 Main St", entities[0].Addresses[0].Line1)
	require.Equal(t, "Las Vegas", entities[0].Addresses[0].City)
	require.Equal(t, "NV", entities[0].Addresses[0].State)
}

func TestReadEntities_GovernmentIDs(t *testing.T) {
	input := `{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"John","NAME_LAST":"Doe","SSN":"123-45-6789","PASSPORT_NUMBER":"AB123456","PASSPORT_COUNTRY":"US","TAX_ID_NUMBER":"98-7654321","DRIVERS_LICENSE_NUMBER":"D1234567","DRIVERS_LICENSE_STATE":"CA"}`

	entities, err := ReadEntities(strings.NewReader(input), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)
	require.NotNil(t, entities[0].Person)
	require.Len(t, entities[0].Person.GovernmentIDs, 4)

	// Check SSN
	var foundSSN, foundPassport, foundTax, foundDL bool
	for _, id := range entities[0].Person.GovernmentIDs {
		switch id.Type {
		case search.GovernmentIDSSN:
			require.Equal(t, "123-45-6789", id.Identifier)
			foundSSN = true
		case search.GovernmentIDPassport:
			require.Equal(t, "AB123456", id.Identifier)
			require.Equal(t, "US", id.Country)
			foundPassport = true
		case search.GovernmentIDTax:
			require.Equal(t, "98-7654321", id.Identifier)
			foundTax = true
		case search.GovernmentIDDriversLicense:
			require.Equal(t, "D1234567", id.Identifier)
			require.Equal(t, "CA", id.Country)
			foundDL = true
		}
	}
	require.True(t, foundSSN, "SSN not found")
	require.True(t, foundPassport, "Passport not found")
	require.True(t, foundTax, "Tax ID not found")
	require.True(t, foundDL, "Drivers License not found")
}

func TestReadEntities_ContactInfo(t *testing.T) {
	input := `{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"John","NAME_LAST":"Doe","PHONE_NUMBER":"555-123-4567","EMAIL_ADDRESS":"john@example.com","WEBSITE_ADDRESS":"https://example.com"}`

	entities, err := ReadEntities(strings.NewReader(input), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)

	require.Equal(t, []string{"555-123-4567"}, entities[0].Contact.PhoneNumbers)
	require.Equal(t, []string{"john@example.com"}, entities[0].Contact.EmailAddresses)
	require.Equal(t, []string{"https://example.com"}, entities[0].Contact.Websites)
}

func TestReadEntities_Dates(t *testing.T) {
	input := `{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"John","NAME_LAST":"Doe","DATE_OF_BIRTH":"1985-03-15","DATE_OF_DEATH":"2020-12-25"}`

	entities, err := ReadEntities(strings.NewReader(input), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)
	require.NotNil(t, entities[0].Person)
	require.NotNil(t, entities[0].Person.BirthDate)
	require.NotNil(t, entities[0].Person.DeathDate)

	expectedBirth := time.Date(1985, 3, 15, 0, 0, 0, 0, time.UTC)
	expectedDeath := time.Date(2020, 12, 25, 0, 0, 0, 0, time.UTC)
	require.Equal(t, expectedBirth, *entities[0].Person.BirthDate)
	require.Equal(t, expectedDeath, *entities[0].Person.DeathDate)
}

func TestReadEntities_Gender(t *testing.T) {
	tests := []struct {
		input    string
		expected search.Gender
	}{
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","GENDER":"M"}`, search.GenderMale},
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","GENDER":"MALE"}`, search.GenderMale},
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","GENDER":"F"}`, search.GenderFemale},
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","GENDER":"FEMALE"}`, search.GenderFemale},
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","GENDER":""}`, search.GenderUnknown},
		{`{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A"}`, search.GenderUnknown},
	}

	for _, tc := range tests {
		entities, err := ReadEntities(strings.NewReader(tc.input), "test")
		require.NoError(t, err)
		require.Len(t, entities, 1)
		require.NotNil(t, entities[0].Person)
		require.Equal(t, tc.expected, entities[0].Person.Gender)
	}
}

func TestReadEntities_AutoDetectType(t *testing.T) {
	// When RECORD_TYPE is missing, should auto-detect based on fields
	personInput := `{"DATA_SOURCE":"T","RECORD_ID":"1","NAME_FIRST":"John","NAME_LAST":"Doe"}`
	orgInput := `{"DATA_SOURCE":"T","RECORD_ID":"2","NAME_ORG":"Acme Corp"}`

	personEntities, err := ReadEntities(strings.NewReader(personInput), "test")
	require.NoError(t, err)
	require.Len(t, personEntities, 1)
	require.Equal(t, search.EntityPerson, personEntities[0].Type)

	orgEntities, err := ReadEntities(strings.NewReader(orgInput), "test")
	require.NoError(t, err)
	require.Len(t, orgEntities, 1)
	require.Equal(t, search.EntityBusiness, orgEntities[0].Type)
}

func TestReadEntities_EmptyInput(t *testing.T) {
	entities, err := ReadEntities(strings.NewReader(""), "test")
	require.NoError(t, err)
	require.Nil(t, entities)

	entities, err = ReadEntities(strings.NewReader("   \n\n  "), "test")
	require.NoError(t, err)
	require.Nil(t, entities)
}

func TestReadEntities_FromTestFiles(t *testing.T) {
	t.Run("persons.jsonl", func(t *testing.T) {
		f, err := os.Open(filepath.Join("testdata", "persons.jsonl"))
		require.NoError(t, err)
		defer f.Close()

		entities, err := ReadEntities(f, "test-persons")
		require.NoError(t, err)
		require.Len(t, entities, 3)

		// First person
		require.Equal(t, "John Smith", entities[0].Name)
		require.Equal(t, "1001", entities[0].SourceID)

		// Second person with middle name
		require.Equal(t, "Jane Marie Doe", entities[1].Name)

		// Third person with full name
		require.Equal(t, "Robert James Wilson Jr", entities[2].Name)
	})

	t.Run("organizations.json", func(t *testing.T) {
		f, err := os.Open(filepath.Join("testdata", "organizations.json"))
		require.NoError(t, err)
		defer f.Close()

		entities, err := ReadEntities(f, "test-orgs")
		require.NoError(t, err)
		require.Len(t, entities, 2)

		require.Equal(t, "Acme Corporation", entities[0].Name)
		require.Equal(t, search.EntityBusiness, entities[0].Type)

		require.Equal(t, "Global Trading Ltd", entities[1].Name)
	})

	t.Run("features_format.jsonl", func(t *testing.T) {
		f, err := os.Open(filepath.Join("testdata", "features_format.jsonl"))
		require.NoError(t, err)
		defer f.Close()

		entities, err := ReadEntities(f, "features-test")
		require.NoError(t, err)
		require.Len(t, entities, 2)

		// Person with features
		require.Equal(t, search.EntityPerson, entities[0].Type)
		require.Equal(t, "Robert Smith", entities[0].Name)
		require.Len(t, entities[0].Addresses, 1)

		// Organization with features
		require.Equal(t, search.EntityBusiness, entities[1].Type)
		require.Equal(t, "Widget Inc", entities[1].Name)
	})
}

func TestReadEntities_AddressCombine(t *testing.T) {
	input := `{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","ADDR_LINE1":"123 Main","ADDR_LINE2":"Suite 100","ADDR_LINE3":"Floor 5","ADDR_CITY":"NYC","ADDR_STATE":"NY","ADDR_POSTAL_CODE":"10001","ADDR_COUNTRY":"US"}`

	entities, err := ReadEntities(strings.NewReader(input), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)
	require.Len(t, entities[0].Addresses, 1)

	addr := entities[0].Addresses[0]
	require.Equal(t, "123 Main", addr.Line1)
	require.Equal(t, "Suite 100, Floor 5", addr.Line2)
	require.Equal(t, "NYC", addr.City)
	require.Equal(t, "NY", addr.State)
	require.Equal(t, "10001", addr.PostalCode)
	require.Equal(t, "US", addr.Country)
}

func TestReadEntities_FullAddress(t *testing.T) {
	input := `{"DATA_SOURCE":"T","RECORD_ID":"1","RECORD_TYPE":"PERSON","NAME_FIRST":"A","ADDR_FULL":"123 Main St, New York, NY 10001"}`

	entities, err := ReadEntities(strings.NewReader(input), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)
	require.Len(t, entities[0].Addresses, 1)
	require.Equal(t, "123 Main St, New York, NY 10001", entities[0].Addresses[0].Line1)
}

func TestReadEntities_InvalidJSON(t *testing.T) {
	_, err := ReadEntities(strings.NewReader("not json"), "test")
	require.Error(t, err)

	_, err = ReadEntities(strings.NewReader(`{"invalid": json}`), "test")
	require.Error(t, err)
}
