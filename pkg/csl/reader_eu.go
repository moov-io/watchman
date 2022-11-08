package csl

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadEUFile(path string) (EUCSL, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	return ParseEU(fd)
}

func ParseEU(r io.Reader) (EUCSL, error) {
	reader := csv.NewReader(r)
	reader.Comma = ';'

	report := make(EUCSL)
	_, err := reader.Read()
	if err != nil {
		fmt.Println("we made an oopsie: ", err)
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
			return nil, err
		}

		if len(record) <= 1 {
			continue // skip empty records
		}

		// merge rows at this point
		// for each record we need to add that to the map
		logicalID, _ := strconv.Atoi(record[EntityLogicalIdx])
		row, err := unmarshalEUCSLRow(record)
		if err != nil {
			return nil, err
		}

		// check if entry does not exist
		report[logicalID] = append(report[logicalID], row)

	}
	return report, nil
}

func unmarshalEUCSLRow(csvRecord []string) (*EUCSLRow, error) {
	row := NewEUCSLRow()

	row.FileGenerationDate = csvRecord[FileGenerationDateIdx]
	row.Entity.ReferenceNumber = csvRecord[ReferenceNumberIdx]
	row.Entity.Remark = csvRecord[EntityRemarkIdx]
	row.Entity.SubjectType.ClassificationCode = csvRecord[EntitySubjectTypeIdx]
	row.Entity.Regulation.PublicationURL = csvRecord[EntityRegulationPublicationURLIdx]

	row.NameAlias.WholeName = csvRecord[NameAliasWholeNameIdx]
	row.NameAlias.Title = csvRecord[NameAliasTitleIdx]

	row.Address.City = csvRecord[AddressCityIdx]
	row.Address.Street = csvRecord[AddressStreetIdx]
	row.Address.PoBox = csvRecord[AddressPoBoxIdx]
	row.Address.ZipCode = csvRecord[AddressZipCodeIdx]
	row.Address.CountryDescription = csvRecord[AddressCountryDescriptionIdx]

	row.BirthDate.BirthDate = csvRecord[BirthDateIdx]
	row.BirthDate.City = csvRecord[BirthDateCityIdx]
	row.BirthDate.CountryDescription = csvRecord[BirthDateCountryIdx]

	row.Identification.ValidFrom = csvRecord[IdentificationValidFromIdx]
	row.Identification.ValidTo = csvRecord[IdentificationValidToIdx]

	return row, nil
}
