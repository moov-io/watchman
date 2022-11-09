package csl

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadEUFile(path string) ([]*EUCSLRow, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	return ParseEU(fd)
}

func ParseEU(r io.Reader) ([]*EUCSLRow, error) {
	reader := csv.NewReader(r)
	reader.Comma = ';'

	report := make(EUCSL)
	_, err := reader.Read()
	if err != nil {
		fmt.Println("we made an oopsie: ", err)
	}
	// TODO: change this back to while
	for i := 0; i <= 4; i++ {
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
			return nil, err
		}

		if len(record) <= 1 {
			continue // skip empty records
		}

		// merge rows at this point
		// for each record we need to add that to the map
		logicalID, _ := strconv.Atoi(record[EntityLogicalIdx])
		// check if entry does not exist
		if val, ok := report[logicalID]; !ok {
			fmt.Println("unmarshalling first row")
			row := unmarshalFirstEUCSLRow(record)

			report[logicalID] = row
		} else {
			// we found an entry in the map and need to append
			fmt.Println("unmarshalling next row")
			unmarshalNextEUCSLRow(record, val)
		}

	}
	var totalReport []*EUCSLRow
	for _, row := range report {
		totalReport = append(totalReport, row)
	}
	return totalReport, nil
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
	row.NameAliases = append(row.NameAliases, newNameAlias)

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
