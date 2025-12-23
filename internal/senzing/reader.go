package senzing

import (
	"bufio"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

// ReadEntities reads Senzing-formatted records and converts them to Watchman entities.
// Supports both JSON Lines (.jsonl) and JSON Array (.json) formats.
// The format is auto-detected based on the first non-whitespace character.
// Uses streaming to handle large files without loading everything into memory.
func ReadEntities(r io.Reader, sourceList search.SourceList) ([]search.Entity[search.Value], error) {
	br := bufio.NewReader(r)

	// Find the first non-whitespace character to detect format
	var firstChar byte
	var err error
	for {
		firstChar, err = br.ReadByte()
		if err != nil {
			if err == io.EOF {
				return nil, nil // Empty input
			}
			return nil, fmt.Errorf("reading senzing input: %w", err)
		}
		if firstChar != ' ' && firstChar != '\n' && firstChar != '\r' && firstChar != '\t' {
			break
		}
	}

	// Put the character back into the stream
	if err := br.UnreadByte(); err != nil {
		return nil, fmt.Errorf("unreading byte: %w", err)
	}

	// Detect format: JSON Array starts with '[', JSON Lines starts with '{'
	if firstChar == '[' {
		return readJSONArray(br, sourceList)
	}
	return readJSONLines(br, sourceList)
}

func readJSONArray(r io.Reader, sourceList search.SourceList) ([]search.Entity[search.Value], error) {
	var entities []search.Entity[search.Value]
	decoder := json.NewDecoder(r)

	// Read opening bracket
	token, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("parsing senzing json array: %w", err)
	}
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		return nil, fmt.Errorf("expected JSON array, got %v", token)
	}

	// Read array elements
	for decoder.More() {
		var rec SenzingRecord
		if err := decoder.Decode(&rec); err != nil {
			return nil, fmt.Errorf("parsing senzing record: %w", err)
		}

		entity, err := ToWatchmanEntity(rec, sourceList)
		if err != nil {
			return nil, fmt.Errorf("converting record %s: %w", rec.RecordID, err)
		}
		entities = append(entities, entity)
	}

	// Read closing bracket
	if _, err := decoder.Token(); err != nil {
		return nil, fmt.Errorf("parsing senzing json array end: %w", err)
	}

	return entities, nil
}

func readJSONLines(r io.Reader, sourceList search.SourceList) ([]search.Entity[search.Value], error) {
	var entities []search.Entity[search.Value]
	scanner := bufio.NewScanner(r)

	// Increase buffer size for large records (1MB max)
	const maxScanTokenSize = 1024 * 1024
	scanner.Buffer(make([]byte, 0, maxScanTokenSize), maxScanTokenSize)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var rec SenzingRecord
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			return nil, fmt.Errorf("parsing line %d: %w", lineNum, err)
		}

		entity, err := ToWatchmanEntity(rec, sourceList)
		if err != nil {
			return nil, fmt.Errorf("converting record at line %d: %w", lineNum, err)
		}
		entities = append(entities, entity)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning senzing jsonl: %w", err)
	}

	return entities, nil
}

// ToWatchmanEntity converts a SenzingRecord to a Watchman Entity.
// Handles both FEATURES array format and flat field format.
func ToWatchmanEntity(rec SenzingRecord, sourceList search.SourceList) (search.Entity[search.Value], error) {
	// Normalize the record (flatten FEATURES if present)
	normalized := normalizeRecord(rec)

	var entity search.Entity[search.Value]

	// Set basic fields
	entity.Source = sourceList
	entity.SourceID = normalized.RecordID
	entity.SourceData = normalized

	// Determine entity type based on RECORD_TYPE
	recordType := strings.ToUpper(normalized.RecordType)

	switch recordType {
	case RecordTypePerson, "INDIVIDUAL":
		populatePersonEntity(&entity, normalized)

	case RecordTypeOrganization, "BUSINESS", "COMPANY":
		populateBusinessEntity(&entity, normalized)

	default:
		// Default: if NAME_ORG is set, treat as business; otherwise as person
		if normalized.NameOrg != "" {
			populateBusinessEntity(&entity, normalized)
		} else {
			populatePersonEntity(&entity, normalized)
		}
	}

	// Extract addresses
	entity.Addresses = extractAddresses(normalized)

	// Extract contact info
	entity.Contact = extractContactInfo(normalized)

	return entity.Normalize(), nil
}

func populatePersonEntity(entity *search.Entity[search.Value], rec SenzingRecord) {
	entity.Type = search.EntityPerson
	entity.Name = buildPersonName(rec)
	entity.Person = &search.Person{
		Name:          entity.Name,
		Gender:        parseGender(rec.Gender),
		BirthDate:     parseDate(rec.DateOfBirth),
		DeathDate:     parseDate(rec.DateOfDeath),
		GovernmentIDs: extractGovernmentIDs(rec),
	}
}

func populateBusinessEntity(entity *search.Entity[search.Value], rec SenzingRecord) {
	entity.Type = search.EntityBusiness
	entity.Name = cmp.Or(rec.NameOrg, rec.NameFull, buildPersonName(rec))
	entity.Business = &search.Business{
		Name:          entity.Name,
		GovernmentIDs: extractGovernmentIDs(rec),
	}
}

// normalizeRecord flattens FEATURES array into direct fields on the record
func normalizeRecord(rec SenzingRecord) SenzingRecord {
	out := rec

	for _, feature := range rec.Features {
		for key, val := range feature {
			strVal, ok := val.(string)
			if !ok {
				continue
			}

			switch strings.ToUpper(key) {
			case FieldRecordType:
				out.RecordType = strVal
			case FieldNameFirst:
				out.NameFirst = strVal
			case FieldNameMiddle:
				out.NameMiddle = strVal
			case FieldNameLast:
				out.NameLast = strVal
			case FieldNameFull:
				out.NameFull = strVal
			case FieldNamePrefix:
				out.NamePrefix = strVal
			case FieldNameSuffix:
				out.NameSuffix = strVal
			case FieldNameOrg:
				out.NameOrg = strVal
			case FieldAddrLine1:
				out.AddrLine1 = strVal
			case FieldAddrLine2:
				out.AddrLine2 = strVal
			case FieldAddrLine3:
				out.AddrLine3 = strVal
			case FieldAddrCity:
				out.AddrCity = strVal
			case FieldAddrState:
				out.AddrState = strVal
			case FieldAddrPostalCode:
				out.AddrPostalCode = strVal
			case FieldAddrCountry:
				out.AddrCountry = strVal
			case FieldAddrFull:
				out.AddrFull = strVal
			case FieldSSN:
				out.SSN = strVal
			case FieldPassportNumber:
				out.PassportNumber = strVal
			case FieldPassportCountry:
				out.PassportCountry = strVal
			case FieldTaxIDNumber:
				out.TaxID = strVal
			case FieldTaxIDCountry:
				out.TaxIDCountry = strVal
			case FieldNationalIDNumber:
				out.NationalID = strVal
			case FieldNationalIDCountry:
				out.NationalIDCountry = strVal
			case FieldDriversLicNumber:
				out.DriversLicenseNumber = strVal
			case FieldDriversLicState:
				out.DriversLicenseState = strVal
			case FieldPhoneNumber:
				out.PhoneNumber = strVal
			case FieldPhoneType:
				out.PhoneType = strVal
			case FieldEmail:
				out.Email = strVal
			case FieldWebsite:
				out.Website = strVal
			case FieldDateOfBirth:
				out.DateOfBirth = strVal
			case FieldDateOfDeath:
				out.DateOfDeath = strVal
			case FieldGender:
				out.Gender = strVal
			case FieldNationality:
				out.Nationality = strVal
			}
		}
	}

	return out
}

func buildPersonName(rec SenzingRecord) string {
	if rec.NameFull != "" {
		return rec.NameFull
	}

	var parts []string
	if rec.NamePrefix != "" {
		parts = append(parts, rec.NamePrefix)
	}
	if rec.NameFirst != "" {
		parts = append(parts, rec.NameFirst)
	}
	if rec.NameMiddle != "" {
		parts = append(parts, rec.NameMiddle)
	}
	if rec.NameLast != "" {
		parts = append(parts, rec.NameLast)
	}
	if rec.NameSuffix != "" {
		parts = append(parts, rec.NameSuffix)
	}

	return strings.Join(parts, " ")
}

func parseDate(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}

	// Senzing typically uses YYYY-MM-DD, YYYYMMDD, or various date formats
	formats := []string{
		"2006-01-02",
		"20060102",
		"2006/01/02",
		"01/02/2006",
		"1/2/2006",
		"Jan 2, 2006",
		"January 2, 2006",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return &t
		}
	}
	return nil
}

func parseGender(g string) search.Gender {
	switch strings.ToUpper(g) {
	case "M", "MALE":
		return search.GenderMale
	case "F", "FEMALE":
		return search.GenderFemale
	default:
		return search.GenderUnknown
	}
}

func extractGovernmentIDs(rec SenzingRecord) []search.GovernmentID {
	var ids []search.GovernmentID

	if rec.SSN != "" {
		ids = append(ids, search.GovernmentID{
			Type:       search.GovernmentIDSSN,
			Identifier: rec.SSN,
		})
	}
	if rec.PassportNumber != "" {
		ids = append(ids, search.GovernmentID{
			Type:       search.GovernmentIDPassport,
			Country:    rec.PassportCountry,
			Identifier: rec.PassportNumber,
		})
	}
	if rec.TaxID != "" {
		ids = append(ids, search.GovernmentID{
			Type:       search.GovernmentIDTax,
			Country:    rec.TaxIDCountry,
			Identifier: rec.TaxID,
		})
	}
	if rec.NationalID != "" {
		ids = append(ids, search.GovernmentID{
			Type:       search.GovernmentIDNational,
			Country:    rec.NationalIDCountry,
			Identifier: rec.NationalID,
		})
	}
	if rec.DriversLicenseNumber != "" {
		ids = append(ids, search.GovernmentID{
			Type:       search.GovernmentIDDriversLicense,
			Country:    rec.DriversLicenseState,
			Identifier: rec.DriversLicenseNumber,
		})
	}

	return ids
}

func extractAddresses(rec SenzingRecord) []search.Address {
	// Handle ADDR_FULL separately - it's a single unparsed address string
	if rec.AddrFull != "" && rec.AddrLine1 == "" {
		return []search.Address{
			{Line1: rec.AddrFull},
		}
	}

	// Check if any parsed address fields are populated
	if rec.AddrLine1 == "" && rec.AddrCity == "" && rec.AddrCountry == "" {
		return nil
	}

	// Combine Line2 and Line3 if both exist
	line2 := rec.AddrLine2
	if rec.AddrLine3 != "" {
		if line2 != "" {
			line2 = line2 + ", " + rec.AddrLine3
		} else {
			line2 = rec.AddrLine3
		}
	}

	addr := search.Address{
		Line1:      rec.AddrLine1,
		Line2:      line2,
		City:       rec.AddrCity,
		State:      rec.AddrState,
		PostalCode: rec.AddrPostalCode,
		Country:    rec.AddrCountry,
	}

	return []search.Address{addr}
}

func extractContactInfo(rec SenzingRecord) search.ContactInfo {
	var info search.ContactInfo

	if rec.PhoneNumber != "" {
		info.PhoneNumbers = []string{rec.PhoneNumber}
	}
	if rec.Email != "" {
		info.EmailAddresses = []string{rec.Email}
	}
	if rec.Website != "" {
		info.Websites = []string{rec.Website}
	}

	return info
}
