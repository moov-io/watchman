package senzing

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

// WriteEntities exports Watchman entities to Senzing format.
// Supports both JSON Lines ("jsonl") and JSON Array ("json") output formats.
func WriteEntities(w io.Writer, entities []search.Entity[search.Value], opts ExportOptions) error {
	records := make([]SenzingRecord, 0, len(entities))

	for _, entity := range entities {
		rec := FromWatchmanEntity(entity, opts.DataSource)
		records = append(records, rec)
	}

	return writeSenzingRecords(w, records, opts)
}

// WriteEntities exports a Watchman SearchedEntity to Senzing format.
// Supports both JSON Lines ("jsonl") and JSON Array ("json") output formats.
func WriteSearchedEntities(w io.Writer, entities []search.SearchedEntity[search.Value], opts ExportOptions) error {
	records := make([]SenzingRecord, 0, len(entities))

	for _, entity := range entities {
		rec := FromWatchmanEntity(entity.Entity, opts.DataSource)
		records = append(records, rec)
	}

	return writeSenzingRecords(w, records, opts)
}

func writeSenzingRecords(w io.Writer, records []SenzingRecord, opts ExportOptions) error {
	format := strings.ToLower(opts.Format)
	switch format {
	case "jsonl", "json-lines", "ndjson":
		return writeJSONLines(w, records, opts.Pretty)
	case "json", "array", "":
		return writeJSONArray(w, records, opts.Pretty)
	default:
		return fmt.Errorf("unknown senzing export format: %s", opts.Format)
	}
}

func writeJSONArray(w io.Writer, records []SenzingRecord, pretty bool) error {
	encoder := json.NewEncoder(w)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	return encoder.Encode(records)
}

func writeJSONLines(w io.Writer, records []SenzingRecord, pretty bool) error {
	for _, rec := range records {
		var data []byte
		var err error

		if pretty {
			data, err = json.MarshalIndent(rec, "", "  ")
		} else {
			data, err = json.Marshal(rec)
		}

		if err != nil {
			return fmt.Errorf("encoding record %s: %w", rec.RecordID, err)
		}

		if _, err := w.Write(data); err != nil {
			return fmt.Errorf("writing record %s: %w", rec.RecordID, err)
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}

// FromWatchmanEntity converts a Watchman Entity to a SenzingRecord.
func FromWatchmanEntity(entity search.Entity[search.Value], defaultDataSource string) SenzingRecord {
	rec := SenzingRecord{
		DataSource: cmp.Or(string(entity.Source), defaultDataSource),
		RecordID:   entity.SourceID,
	}

	switch entity.Type {
	case search.EntityPerson:
		rec.RecordType = RecordTypePerson
		if entity.Person != nil {
			populatePersonFields(&rec, entity.Person)
		} else {
			rec.NameFull = entity.Name
		}

	case search.EntityBusiness:
		rec.RecordType = RecordTypeOrganization
		rec.NameOrg = entity.Name
		if entity.Business != nil {
			populateGovernmentIDs(&rec, entity.Business.GovernmentIDs)
		}

	case search.EntityOrganization:
		rec.RecordType = RecordTypeOrganization
		rec.NameOrg = entity.Name
		if entity.Organization != nil {
			populateGovernmentIDs(&rec, entity.Organization.GovernmentIDs)
		}

	default:
		// For vessels, aircraft, etc. - treat as organization
		rec.RecordType = RecordTypeOrganization
		rec.NameOrg = entity.Name
	}

	// Populate addresses (use first address for flat format)
	if len(entity.Addresses) > 0 {
		addr := entity.Addresses[0]
		rec.AddrLine1 = addr.Line1
		rec.AddrLine2 = addr.Line2
		rec.AddrCity = addr.City
		rec.AddrState = addr.State
		rec.AddrPostalCode = addr.PostalCode
		rec.AddrCountry = addr.Country
	}

	// Populate contact info (use first values for flat format)
	if len(entity.Contact.PhoneNumbers) > 0 {
		rec.PhoneNumber = entity.Contact.PhoneNumbers[0]
	}
	if len(entity.Contact.EmailAddresses) > 0 {
		rec.Email = entity.Contact.EmailAddresses[0]
	}
	if len(entity.Contact.Websites) > 0 {
		rec.Website = entity.Contact.Websites[0]
	}

	return rec
}

func populatePersonFields(rec *SenzingRecord, person *search.Person) {
	// Try to parse name into components using simple heuristics
	rec.NameFull = person.Name

	parts := strings.Fields(person.Name)
	switch len(parts) {
	case 1:
		rec.NameLast = parts[0]
	case 2:
		rec.NameFirst = parts[0]
		rec.NameLast = parts[1]
	case 3:
		rec.NameFirst = parts[0]
		rec.NameMiddle = parts[1]
		rec.NameLast = parts[2]
	default:
		if len(parts) > 3 {
			rec.NameFirst = parts[0]
			rec.NameMiddle = strings.Join(parts[1:len(parts)-1], " ")
			rec.NameLast = parts[len(parts)-1]
		}
	}

	// Gender
	rec.Gender = formatGender(person.Gender)

	// Dates
	if person.BirthDate != nil {
		rec.DateOfBirth = person.BirthDate.Format("2006-01-02")
	}
	if person.DeathDate != nil {
		rec.DateOfDeath = person.DeathDate.Format("2006-01-02")
	}

	// Government IDs
	populateGovernmentIDs(rec, person.GovernmentIDs)
}

func populateGovernmentIDs(rec *SenzingRecord, ids []search.GovernmentID) {
	for _, id := range ids {
		switch id.Type {
		case search.GovernmentIDSSN:
			rec.SSN = id.Identifier
		case search.GovernmentIDPassport:
			rec.PassportNumber = id.Identifier
			rec.PassportCountry = id.Country
		case search.GovernmentIDTax:
			rec.TaxID = id.Identifier
			rec.TaxIDCountry = id.Country
		case search.GovernmentIDNational:
			rec.NationalID = id.Identifier
			rec.NationalIDCountry = id.Country
		case search.GovernmentIDDriversLicense:
			rec.DriversLicenseNumber = id.Identifier
			rec.DriversLicenseState = id.Country
		}
	}
}

func formatGender(g search.Gender) string {
	switch g {
	case search.GenderMale:
		return "M"
	case search.GenderFemale:
		return "F"
	default:
		return ""
	}
}
