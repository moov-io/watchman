package csl

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadEUFile(path string) ([]*EUCSLRow, EUCSL, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	rows, rowsMap, err := ParseEU(fd)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func ParseEU(r io.Reader) ([]*EUCSLRow, EUCSL, error) {
	reader := csv.NewReader(r)
	// sets comma delim to ; and ignores " in non quoted field and size of columns
	// https://stackoverflow.com/questions/31326659/golang-csv-error-bare-in-non-quoted-field
	// https://stackoverflow.com/questions/61336787/how-do-i-fix-the-wrong-number-of-fields-with-the-missing-commas-in-csv-file-in
	reader.Comma = ';'
	reader.LazyQuotes = true
	// reader.FieldsPerRecord = -1

	report := make(EUCSL)
	_, err := reader.Read()
	if err != nil {
		fmt.Println("failed to read csv: ", err)
	}
	for {
		record, err := reader.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, nil, err
		}

		if len(record) <= 1 {
			fmt.Println("record is <= 1", record)
			continue // skip empty records
		}

		// merge rows at this point
		// for each record we need to add that to the map
		logicalID, _ := strconv.Atoi(record[EntityLogicalIdx])
		// check if entry does not exist
		if val, ok := report[logicalID]; !ok {
			row := unmarshalFirstEUCSLRow(record)

			report[logicalID] = row
		} else {
			// we found an entry in the map and need to append
			unmarshalNextEUCSLRow(record, val)
		}

	}
	var totalReport []*EUCSLRow
	for _, row := range report {
		totalReport = append(totalReport, row)
	}
	return totalReport, report, nil
}

func unmarshalFirstEUCSLRow(csvRecord []string) *EUCSLRow {
	row := NewEUCSLRow()

	row.FileGenerationDate = csvRecord[FileGenerationDateIdx]
	row.Entity.ReferenceNumber = csvRecord[ReferenceNumberIdx]
	row.Entity.LogicalID, _ = strconv.Atoi(csvRecord[EntityLogicalIdx])
	row.Entity.Remark = csvRecord[EntityRemarkIdx]
	row.Entity.SubjectType.ClassificationCode = csvRecord[EntitySubjectTypeIdx]
	row.Entity.Regulation.PublicationURL = csvRecord[EntityRegulationPublicationURLIdx]

	var newNameAlias *NameAlias
	if csvRecord[NameAliasWholeNameIdx] != "" {
		newNameAlias = new(NameAlias)
		newNameAlias.WholeName = csvRecord[NameAliasWholeNameIdx]
	}
	if csvRecord[NameAliasTitleIdx] != "" {
		if newNameAlias == nil {
			newNameAlias = new(NameAlias)
		}
		newNameAlias.Title = csvRecord[NameAliasTitleIdx]
	}
	if newNameAlias != nil {
		row.NameAliases = append(row.NameAliases, newNameAlias)
	}

	var newAddress *Address
	if csvRecord[AddressCityIdx] != "" {
		newAddress = new(Address)
		newAddress.City = csvRecord[AddressCityIdx]
	}
	if csvRecord[AddressStreetIdx] != "" {
		if newAddress == nil {
			newAddress = new(Address)
		}
		newAddress.Street = csvRecord[AddressStreetIdx]
	}
	if csvRecord[AddressPoBoxIdx] != "" {
		if newAddress == nil {
			newAddress = new(Address)
		}
		newAddress.PoBox = csvRecord[AddressPoBoxIdx]
	}
	if csvRecord[AddressZipCodeIdx] != "" {
		if newAddress == nil {
			newAddress = new(Address)
		}
		newAddress.ZipCode = csvRecord[AddressZipCodeIdx]
	}
	if csvRecord[AddressCountryDescriptionIdx] != "" {
		if newAddress == nil {
			newAddress = new(Address)
		}
		newAddress.CountryDescription = csvRecord[AddressCountryDescriptionIdx]
	}
	if newAddress != nil {
		row.Addresses = append(row.Addresses, newAddress)
	}

	var newBirthDate *BirthDate
	if csvRecord[BirthDateIdx] != "" {
		newBirthDate = new(BirthDate)
		newBirthDate.BirthDate = csvRecord[BirthDateIdx]
	}
	if csvRecord[BirthDateCityIdx] != "" {
		if newBirthDate == nil {
			newBirthDate = new(BirthDate)
		}
		newBirthDate.City = csvRecord[BirthDateCityIdx]
	}
	if csvRecord[BirthDateCountryIdx] != "" {
		if newBirthDate == nil {
			newBirthDate = new(BirthDate)
		}
		newBirthDate.CountryDescription = csvRecord[BirthDateCountryIdx]
	}
	if newBirthDate != nil {
		row.BirthDates = append(row.BirthDates, newBirthDate)
	}

	var newIdentification *Identification
	if csvRecord[IdentificationValidFromIdx] != "" {
		newIdentification = new(Identification)
		newIdentification.ValidFrom = csvRecord[IdentificationValidFromIdx]
	}
	if csvRecord[IdentificationValidToIdx] != "" {
		if newIdentification == nil {
			newIdentification = new(Identification)
		}
		newIdentification.ValidTo = csvRecord[IdentificationValidToIdx]
	}
	if newIdentification != nil {
		row.Identifications = append(row.Identifications, newIdentification)
	}

	return row
}

func unmarshalNextEUCSLRow(csvRecord []string, row *EUCSLRow) {
	// NameAlias
	var newNameAlias *NameAlias
	if csvRecord[NameAliasWholeNameIdx] != "" {
		newNameAlias = new(NameAlias)
		newNameAlias.WholeName = csvRecord[NameAliasWholeNameIdx]
	}
	if csvRecord[NameAliasTitleIdx] != "" {
		if newNameAlias == nil {
			newNameAlias = new(NameAlias)
		}
		newNameAlias.Title = csvRecord[NameAliasTitleIdx]
	}
	if newNameAlias != nil {
		row.NameAliases = append(row.NameAliases, newNameAlias)
	}

	// Address
	var newAddress *Address
	if csvRecord[AddressCityIdx] != "" {
		newAddress = new(Address)
		newAddress.City = csvRecord[AddressCityIdx]
	}
	if csvRecord[AddressStreetIdx] != "" {
		if newAddress == nil {
			newAddress = &Address{}
		}
		newAddress.Street = csvRecord[AddressStreetIdx]
	}
	if csvRecord[AddressPoBoxIdx] != "" {
		if newAddress == nil {
			newAddress = &Address{}
		}
		newAddress.PoBox = csvRecord[AddressPoBoxIdx]
	}
	if csvRecord[AddressZipCodeIdx] != "" {
		if newAddress == nil {
			newAddress.ZipCode = csvRecord[AddressZipCodeIdx]
		}
		newAddress.ZipCode = csvRecord[AddressZipCodeIdx]
	}
	if csvRecord[AddressCountryDescriptionIdx] != "" {
		if newAddress == nil {
			newAddress = &Address{}
		}
		newAddress.CountryDescription = csvRecord[AddressCountryDescriptionIdx]
	}
	if newAddress != nil {
		row.Addresses = append(row.Addresses, newAddress)
	}

	// BirthDate
	var newBirthDate *BirthDate
	if csvRecord[BirthDateIdx] != "" {
		newBirthDate = new(BirthDate)
		newBirthDate.BirthDate = csvRecord[BirthDateIdx]
	}
	if csvRecord[BirthDateCityIdx] != "" {
		if newBirthDate == nil {
			newBirthDate = new(BirthDate)
		}
		newBirthDate.City = csvRecord[BirthDateCityIdx]
	}
	if csvRecord[BirthDateCountryIdx] != "" {
		if newBirthDate == nil {
			newBirthDate = new(BirthDate)
		}
		newBirthDate.CountryDescription = csvRecord[BirthDateCountryIdx]
	}
	if newBirthDate != nil {
		row.BirthDates = append(row.BirthDates, newBirthDate)
	}

	var newIdentification *Identification
	if csvRecord[IdentificationValidFromIdx] != "" {
		newIdentification = new(Identification)
		newIdentification.ValidFrom = csvRecord[IdentificationValidFromIdx]
	}

	if csvRecord[IdentificationValidToIdx] != "" {
		if newIdentification == nil {
			newIdentification = new(Identification)
		}
		newIdentification.ValidTo = csvRecord[IdentificationValidToIdx]
	}

	if newIdentification != nil {
		row.Identifications = append(row.Identifications, newIdentification)
	}
}
