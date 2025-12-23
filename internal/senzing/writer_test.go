package senzing

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestWriteEntities_JSONLines(t *testing.T) {
	birthDate := time.Date(1985, 3, 15, 0, 0, 0, 0, time.UTC)
	entities := []search.Entity[search.Value]{
		{
			Name:     "John Smith",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1001",
			Person: &search.Person{
				Name:      "John Smith",
				Gender:    search.GenderMale,
				BirthDate: &birthDate,
				GovernmentIDs: []search.GovernmentID{
					{Type: search.GovernmentIDSSN, Identifier: "123-45-6789"},
				},
			},
			Addresses: []search.Address{
				{Line1: "123 Main St", City: "Las Vegas", State: "NV", PostalCode: "89101", Country: "US"},
			},
			Contact: search.ContactInfo{
				PhoneNumbers:   []string{"555-123-4567"},
				EmailAddresses: []string{"john@example.com"},
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{
		DataSource: "TEST",
		Format:     "jsonl",
	})
	require.NoError(t, err)

	// Parse output
	var rec SenzingRecord
	err = json.Unmarshal(buf.Bytes(), &rec)
	require.NoError(t, err)

	require.Equal(t, "test", rec.DataSource) // Uses entity source, not default
	require.Equal(t, "1001", rec.RecordID)
	require.Equal(t, RecordTypePerson, rec.RecordType)
	require.Equal(t, "John Smith", rec.NameFull)
	require.Equal(t, "John", rec.NameFirst)
	require.Equal(t, "Smith", rec.NameLast)
	require.Equal(t, "M", rec.Gender)
	require.Equal(t, "1985-03-15", rec.DateOfBirth)
	require.Equal(t, "123-45-6789", rec.SSN)
	require.Equal(t, "123 Main St", rec.AddrLine1)
	require.Equal(t, "Las Vegas", rec.AddrCity)
	require.Equal(t, "555-123-4567", rec.PhoneNumber)
	require.Equal(t, "john@example.com", rec.Email)
}

func TestWriteEntities_JSONArray(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{
			Name:     "Acme Corp",
			Type:     search.EntityBusiness,
			Source:   "test",
			SourceID: "2001",
			Business: &search.Business{
				Name: "Acme Corp",
				GovernmentIDs: []search.GovernmentID{
					{Type: search.GovernmentIDTax, Identifier: "12-3456789", Country: "US"},
				},
			},
		},
		{
			Name:     "Widget Inc",
			Type:     search.EntityBusiness,
			Source:   "test",
			SourceID: "2002",
			Business: &search.Business{
				Name: "Widget Inc",
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{
		DataSource: "COMPANIES",
		Format:     "json",
	})
	require.NoError(t, err)

	// Parse output
	var records []SenzingRecord
	err = json.Unmarshal(buf.Bytes(), &records)
	require.NoError(t, err)
	require.Len(t, records, 2)

	require.Equal(t, "Acme Corp", records[0].NameOrg)
	require.Equal(t, RecordTypeOrganization, records[0].RecordType)
	require.Equal(t, "12-3456789", records[0].TaxID)

	require.Equal(t, "Widget Inc", records[1].NameOrg)
}

func TestWriteEntities_UsesDefaultDataSource(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{
			Name:     "John Doe",
			Type:     search.EntityPerson,
			Source:   "", // Empty source
			SourceID: "1001",
			Person: &search.Person{
				Name: "John Doe",
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{
		DataSource: "DEFAULT_SOURCE",
		Format:     "jsonl",
	})
	require.NoError(t, err)

	var rec SenzingRecord
	err = json.Unmarshal(buf.Bytes(), &rec)
	require.NoError(t, err)

	require.Equal(t, "DEFAULT_SOURCE", rec.DataSource)
}

func TestWriteEntities_AllGovernmentIDTypes(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{
			Name:     "John Doe",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1001",
			Person: &search.Person{
				Name: "John Doe",
				GovernmentIDs: []search.GovernmentID{
					{Type: search.GovernmentIDSSN, Identifier: "123-45-6789"},
					{Type: search.GovernmentIDPassport, Identifier: "AB123456", Country: "US"},
					{Type: search.GovernmentIDTax, Identifier: "98-7654321", Country: "US"},
					{Type: search.GovernmentIDNational, Identifier: "NAT123", Country: "FR"},
					{Type: search.GovernmentIDDriversLicense, Identifier: "DL789", Country: "CA"},
				},
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{Format: "jsonl"})
	require.NoError(t, err)

	var rec SenzingRecord
	err = json.Unmarshal(buf.Bytes(), &rec)
	require.NoError(t, err)

	require.Equal(t, "123-45-6789", rec.SSN)
	require.Equal(t, "AB123456", rec.PassportNumber)
	require.Equal(t, "US", rec.PassportCountry)
	require.Equal(t, "98-7654321", rec.TaxID)
	require.Equal(t, "US", rec.TaxIDCountry)
	require.Equal(t, "NAT123", rec.NationalID)
	require.Equal(t, "FR", rec.NationalIDCountry)
	require.Equal(t, "DL789", rec.DriversLicenseNumber)
	require.Equal(t, "CA", rec.DriversLicenseState)
}

func TestWriteEntities_NameParsing(t *testing.T) {
	tests := []struct {
		name          string
		expectedFirst string
		expectedMid   string
		expectedLast  string
	}{
		{"Smith", "", "", "Smith"},
		{"John Smith", "John", "", "Smith"},
		{"John James Smith", "John", "James", "Smith"},
		{"John James Michael Smith", "John", "James Michael", "Smith"},
	}

	for _, tc := range tests {
		entities := []search.Entity[search.Value]{
			{
				Name:     tc.name,
				Type:     search.EntityPerson,
				Source:   "test",
				SourceID: "1001",
				Person: &search.Person{
					Name: tc.name,
				},
			},
		}

		var buf bytes.Buffer
		err := WriteEntities(&buf, entities, ExportOptions{Format: "jsonl"})
		require.NoError(t, err)

		var rec SenzingRecord
		err = json.Unmarshal(buf.Bytes(), &rec)
		require.NoError(t, err)

		require.Equal(t, tc.name, rec.NameFull, "NameFull mismatch for: %s", tc.name)
		require.Equal(t, tc.expectedFirst, rec.NameFirst, "NameFirst mismatch for: %s", tc.name)
		require.Equal(t, tc.expectedMid, rec.NameMiddle, "NameMiddle mismatch for: %s", tc.name)
		require.Equal(t, tc.expectedLast, rec.NameLast, "NameLast mismatch for: %s", tc.name)
	}
}

func TestWriteEntities_Organization(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{
			Name:     "Test Org",
			Type:     search.EntityOrganization,
			Source:   "test",
			SourceID: "3001",
			Organization: &search.Organization{
				Name: "Test Org",
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{Format: "jsonl"})
	require.NoError(t, err)

	var rec SenzingRecord
	err = json.Unmarshal(buf.Bytes(), &rec)
	require.NoError(t, err)

	require.Equal(t, RecordTypeOrganization, rec.RecordType)
	require.Equal(t, "Test Org", rec.NameOrg)
}

func TestWriteEntities_VesselAndAircraft(t *testing.T) {
	// Vessels and aircraft are treated as organizations
	entities := []search.Entity[search.Value]{
		{
			Name:     "SS Enterprise",
			Type:     search.EntityVessel,
			Source:   "test",
			SourceID: "4001",
			Vessel: &search.Vessel{
				Name: "SS Enterprise",
			},
		},
		{
			Name:     "Boeing 747",
			Type:     search.EntityAircraft,
			Source:   "test",
			SourceID: "4002",
			Aircraft: &search.Aircraft{
				Name: "Boeing 747",
			},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{Format: "jsonl"})
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	require.Len(t, lines, 2)

	var rec1, rec2 SenzingRecord
	require.NoError(t, json.Unmarshal([]byte(lines[0]), &rec1))
	require.NoError(t, json.Unmarshal([]byte(lines[1]), &rec2))

	require.Equal(t, RecordTypeOrganization, rec1.RecordType)
	require.Equal(t, "SS Enterprise", rec1.NameOrg)

	require.Equal(t, RecordTypeOrganization, rec2.RecordType)
	require.Equal(t, "Boeing 747", rec2.NameOrg)
}

func TestWriteEntities_Pretty(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{
			Name:     "John Doe",
			Type:     search.EntityPerson,
			Source:   "test",
			SourceID: "1001",
			Person:   &search.Person{Name: "John Doe"},
		},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{
		Format: "jsonl",
		Pretty: true,
	})
	require.NoError(t, err)

	// Pretty output should contain newlines within the JSON
	require.Contains(t, buf.String(), "\n  ")
}

func TestWriteEntities_Empty(t *testing.T) {
	var buf bytes.Buffer
	err := WriteEntities(&buf, nil, ExportOptions{Format: "json"})
	require.NoError(t, err)
	require.Equal(t, "[]\n", buf.String())

	buf.Reset()
	err = WriteEntities(&buf, []search.Entity[search.Value]{}, ExportOptions{Format: "jsonl"})
	require.NoError(t, err)
	require.Equal(t, "", buf.String())
}

func TestWriteEntities_InvalidFormat(t *testing.T) {
	entities := []search.Entity[search.Value]{
		{Name: "Test", Type: search.EntityPerson, SourceID: "1"},
	}

	var buf bytes.Buffer
	err := WriteEntities(&buf, entities, ExportOptions{Format: "invalid"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "unknown senzing export format")
}

func TestRoundTrip(t *testing.T) {
	// Test that import -> export -> import produces equivalent entities
	original := `{"DATA_SOURCE":"TEST","RECORD_ID":"1001","RECORD_TYPE":"PERSON","NAME_FIRST":"John","NAME_LAST":"Smith","SSN":"123-45-6789","DATE_OF_BIRTH":"1985-03-15","ADDR_LINE1":"123 Main St","ADDR_CITY":"Las Vegas","PHONE_NUMBER":"555-1234"}`

	// Import
	entities, err := ReadEntities(strings.NewReader(original), "test")
	require.NoError(t, err)
	require.Len(t, entities, 1)

	// Export
	var buf bytes.Buffer
	err = WriteEntities(&buf, entities, ExportOptions{Format: "jsonl"})
	require.NoError(t, err)

	// Re-import
	reimported, err := ReadEntities(&buf, "test")
	require.NoError(t, err)
	require.Len(t, reimported, 1)

	// Compare key fields
	require.Equal(t, entities[0].Name, reimported[0].Name)
	require.Equal(t, entities[0].SourceID, reimported[0].SourceID)
	require.Equal(t, entities[0].Type, reimported[0].Type)

	// Compare Person fields
	require.NotNil(t, entities[0].Person)
	require.NotNil(t, reimported[0].Person)
	require.Equal(t, len(entities[0].Person.GovernmentIDs), len(reimported[0].Person.GovernmentIDs))

	// Compare addresses
	require.Equal(t, len(entities[0].Addresses), len(reimported[0].Addresses))
	if len(entities[0].Addresses) > 0 {
		require.Equal(t, entities[0].Addresses[0].Line1, reimported[0].Addresses[0].Line1)
		require.Equal(t, entities[0].Addresses[0].City, reimported[0].Addresses[0].City)
	}
}
