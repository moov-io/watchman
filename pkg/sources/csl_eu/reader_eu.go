// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func ParseEU(r io.ReadCloser) ([]CSLRecord, CSL, error) {
	if r == nil {
		return nil, nil, errors.New("EU CSL file is empty or missing")
	}
	defer r.Close()
	reader := csv.NewReader(r)
	// sets comma delim to ; and ignores " in non quoted field and size of columns
	// https://stackoverflow.com/questions/31326659/golang-csv-error-bare-in-non-quoted-field
	// https://stackoverflow.com/questions/61336787/how-do-i-fix-the-wrong-number-of-fields-with-the-missing-commas-in-csv-file-in
	reader.Comma = ';'
	reader.LazyQuotes = true

	report := make(CSL)
	_, err := reader.Read()
	if err != nil {
		return nil, report, fmt.Errorf("failed to read csv: %w", err)
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
			continue // skip empty records
		}

		// merge rows at this point
		// for each record we need to add that to the map
		logicalID, _ := strconv.Atoi(record[EntityLogicalIdx])
		// check if entry does not exist
		if val, ok := report[logicalID]; !ok {
			// creates the initial record
			row := new(CSLRecord)
			unmarshalRecord(record, row)

			report[logicalID] = row
		} else {
			// we found an entry in the map and need to append
			unmarshalRecord(record, val)
		}

	}
	totalReport := make([]CSLRecord, 0, len(report))
	for _, row := range report {
		totalReport = append(totalReport, *row)
	}
	return totalReport, report, nil
}

func unmarshalRecord(csvRecord []string, euCSLRecord *CSLRecord) {
	euCSLRecord.EntityLogicalID, _ = strconv.Atoi(csvRecord[EntityLogicalIdx])

	// entity
	if csvRecord[FileGenerationDateIdx] != "" {
		euCSLRecord.FileGenerationDate = csvRecord[FileGenerationDateIdx]
	}
	if csvRecord[ReferenceNumberIdx] != "" {
		euCSLRecord.EntityReferenceNumber = csvRecord[ReferenceNumberIdx]
	}
	if csvRecord[EntityRemarkIdx] != "" {
		euCSLRecord.EntityRemark = csvRecord[EntityRemarkIdx]
	}
	if csvRecord[EntitySubjectTypeIdx] != "" {
		euCSLRecord.EntitySubjectType = csvRecord[EntitySubjectTypeIdx]
	}
	if csvRecord[EntityRegulationPublicationURLIdx] != "" {
		euCSLRecord.EntityPublicationURL = csvRecord[EntityRegulationPublicationURLIdx]
	}

	// name alias
	if csvRecord[NameAliasWholeNameIdx] != "" {
		if !arrayContains(euCSLRecord.NameAliasWholeNames, csvRecord[NameAliasWholeNameIdx]) {
			euCSLRecord.NameAliasWholeNames = append(euCSLRecord.NameAliasWholeNames, csvRecord[NameAliasWholeNameIdx])
		}
	}
	if csvRecord[NameAliasTitleIdx] != "" {
		if !arrayContains(euCSLRecord.NameAliasTitles, csvRecord[NameAliasTitleIdx]) {
			euCSLRecord.NameAliasTitles = append(euCSLRecord.NameAliasTitles, csvRecord[NameAliasTitleIdx])
		}
	}
	// address
	if csvRecord[AddressCityIdx] != "" {
		if !arrayContains(euCSLRecord.AddressCities, csvRecord[AddressCityIdx]) {
			euCSLRecord.AddressCities = append(euCSLRecord.AddressCities, csvRecord[AddressCityIdx])
		}
	}
	if csvRecord[AddressStreetIdx] != "" {
		if !arrayContains(euCSLRecord.AddressStreets, csvRecord[AddressStreetIdx]) {
			euCSLRecord.AddressStreets = append(euCSLRecord.AddressStreets, csvRecord[AddressStreetIdx])
		}
	}
	if csvRecord[AddressPoBoxIdx] != "" {
		if !arrayContains(euCSLRecord.AddressPoBoxes, csvRecord[AddressPoBoxIdx]) {
			euCSLRecord.AddressPoBoxes = append(euCSLRecord.AddressPoBoxes, csvRecord[AddressPoBoxIdx])
		}
	}
	if csvRecord[AddressZipCodeIdx] != "" {
		if !arrayContains(euCSLRecord.AddressZipCodes, csvRecord[AddressZipCodeIdx]) {
			euCSLRecord.AddressZipCodes = append(euCSLRecord.AddressZipCodes, csvRecord[AddressZipCodeIdx])
		}
	}
	if csvRecord[AddressCountryDescriptionIdx] != "" {
		if !arrayContains(euCSLRecord.AddressCountryDescriptions, csvRecord[AddressCountryDescriptionIdx]) {
			euCSLRecord.AddressCountryDescriptions = append(euCSLRecord.AddressCountryDescriptions, csvRecord[AddressCountryDescriptionIdx])
		}
	}

	// birthdate
	if csvRecord[BirthDateIdx] != "" {
		if !arrayContains(euCSLRecord.BirthDates, csvRecord[BirthDateIdx]) {
			euCSLRecord.BirthDates = append(euCSLRecord.BirthDates, csvRecord[BirthDateIdx])
		}
	}
	if csvRecord[BirthDateCityIdx] != "" {
		if !arrayContains(euCSLRecord.BirthCities, csvRecord[BirthDateCityIdx]) {
			euCSLRecord.BirthCities = append(euCSLRecord.BirthCities, csvRecord[BirthDateCityIdx])
		}
	}
	if csvRecord[BirthDateCountryIdx] != "" {
		if !arrayContains(euCSLRecord.BirthCountries, csvRecord[BirthDateCountryIdx]) {
			euCSLRecord.BirthCountries = append(euCSLRecord.BirthCountries, csvRecord[BirthDateCountryIdx])
		}
	}

	// identifications
	if csvRecord[IdentificationValidFromIdx] != "" {
		euCSLRecord.ValidFromTo = make(map[string]string)
		euCSLRecord.ValidFromTo[csvRecord[IdentificationValidFromIdx]] = csvRecord[IdentificationValidToIdx]
	}
}

func arrayContains(checkArray []string, nameToCheck string) bool {
	var nameAlreadyExists bool = false
	if nameToCheck == "" {
		return true
	}
	for _, name := range checkArray {
		if name == nameToCheck {
			nameAlreadyExists = true
			break
		}
	}
	return nameAlreadyExists
}
