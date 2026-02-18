// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ReadCSLFile(fd io.ReadCloser) ([]CSLRecord, CSL, error) {
	if fd == nil {
		return nil, nil, errors.New("uk CSL file is empty or missing")
	}
	defer fd.Close()

	rows, rowsMap, err := ParseCSL(fd)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func ParseCSL(r io.Reader) ([]CSLRecord, CSL, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 36

	report := make(CSL)
	// read and ignore first two rows
	for i := 0; i <= 1; i++ {
		reader.Read()
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
		groupID, err := strconv.Atoi(record[GroupdIdx])
		if err != nil {
			return nil, nil, err
		}

		// check if entry does not exist
		if val, ok := report[groupID]; !ok {
			// creates the initial record
			row := new(CSLRecord)
			unmarshalCSLRecord(record, row)

			report[groupID] = row
		} else {
			// we found an entry in the map and need to append
			unmarshalCSLRecord(record, val)
		}

	}
	totalReport := make([]CSLRecord, 0, len(report))
	for _, row := range report {
		totalReport = append(totalReport, *row)
	}
	return totalReport, report, nil
}

func unmarshalCSLRecord(csvRecord []string, ukCSLRecord *CSLRecord) {
	var names []string
	if csvRecord[UKNameIdx] != "" {
		names = append(names, csvRecord[UKNameIdx])
	}
	if csvRecord[UKNameTwoIdx] != "" {
		names = append(names, csvRecord[UKNameTwoIdx])
	}
	if csvRecord[UKNameThreeIdx] != "" {
		names = append(names, csvRecord[UKNameThreeIdx])
	}
	if csvRecord[UKNameFourIdx] != "" {
		names = append(names, csvRecord[UKNameFourIdx])
	}
	if csvRecord[UKNameFiveIdx] != "" {
		names = append(names, csvRecord[UKNameFiveIdx])
	}
	name := strings.Join(names, " ")
	if !arrayContains(ukCSLRecord.Names, name) {
		ukCSLRecord.Names = append(ukCSLRecord.Names, name)
	}

	if csvRecord[UKTitleIdx] != "" {
		if !arrayContains(ukCSLRecord.Titles, csvRecord[UKTitleIdx]) {
			ukCSLRecord.Titles = append(ukCSLRecord.Titles, csvRecord[UKTitleIdx])
		}
	}
	if csvRecord[DOBhIdx] != "" {
		if !arrayContains(ukCSLRecord.DatesOfBirth, csvRecord[DOBhIdx]) {
			ukCSLRecord.DatesOfBirth = append(ukCSLRecord.DatesOfBirth, csvRecord[DOBhIdx])
		}
	}
	if csvRecord[TownOfBirthIdx] != "" {
		if !arrayContains(ukCSLRecord.TownsOfBirth, csvRecord[TownOfBirthIdx]) {
			ukCSLRecord.TownsOfBirth = append(ukCSLRecord.TownsOfBirth, csvRecord[TownOfBirthIdx])
		}
	}
	if csvRecord[CountryOfBirthIdx] != "" {
		if !arrayContains(ukCSLRecord.CountriesOfBirth, csvRecord[CountryOfBirthIdx]) {
			ukCSLRecord.CountriesOfBirth = append(ukCSLRecord.CountriesOfBirth, csvRecord[CountryOfBirthIdx])
		}
	}
	if csvRecord[UKNationalitiesIdx] != "" {
		if !arrayContains(ukCSLRecord.Nationalities, csvRecord[UKNationalitiesIdx]) {
			ukCSLRecord.Nationalities = append(ukCSLRecord.Nationalities, csvRecord[UKNationalitiesIdx])
		}
	}

	var addresses []string
	if csvRecord[AddressOneIdx] != "" {
		addresses = append(addresses, csvRecord[AddressOneIdx])
	}
	if csvRecord[AddressTwoIdx] != "" {
		addresses = append(addresses, csvRecord[AddressTwoIdx])
	}
	if csvRecord[AddressThreeIdx] != "" {
		addresses = append(addresses, csvRecord[AddressThreeIdx])
	}
	if csvRecord[AddressFourIdx] != "" {
		addresses = append(addresses, csvRecord[AddressFourIdx])
	}
	if csvRecord[AddressFiveIdx] != "" {
		addresses = append(addresses, csvRecord[AddressFiveIdx])
	}
	if csvRecord[AddressSixIdx] != "" {
		addresses = append(addresses, csvRecord[AddressSixIdx])
	}
	address := strings.Join(addresses, ", ")
	if !arrayContains(ukCSLRecord.Addresses, address) {
		ukCSLRecord.Addresses = append(ukCSLRecord.Addresses, address)
	}

	if csvRecord[PostalCodeIdx] != "" {
		if !arrayContains(ukCSLRecord.PostalCodes, csvRecord[PostalCodeIdx]) {
			ukCSLRecord.PostalCodes = append(ukCSLRecord.PostalCodes, csvRecord[PostalCodeIdx])
		}
	}
	if csvRecord[CountryIdx] != "" {
		if !arrayContains(ukCSLRecord.Countries, csvRecord[CountryIdx]) {
			ukCSLRecord.Countries = append(ukCSLRecord.Countries, csvRecord[CountryIdx])
		}
	}
	if csvRecord[OtherInfoIdx] != "" {
		if !arrayContains(ukCSLRecord.OtherInfos, csvRecord[OtherInfoIdx]) {
			ukCSLRecord.OtherInfos = append(ukCSLRecord.OtherInfos, csvRecord[OtherInfoIdx])
		}
	}

	if csvRecord[GroupTypeIdx] != "" && ukCSLRecord.GroupType == "" {
		ukCSLRecord.GroupType = csvRecord[GroupTypeIdx]
	}

	if csvRecord[ListedDateIdx] != "" {
		if !arrayContains(ukCSLRecord.ListedDates, csvRecord[ListedDateIdx]) {
			ukCSLRecord.ListedDates = append(ukCSLRecord.ListedDates, csvRecord[ListedDateIdx])
		}
	}
	if csvRecord[LastUpdatedIdx] != "" {
		if !arrayContains(ukCSLRecord.LastUpdates, csvRecord[LastUpdatedIdx]) {
			ukCSLRecord.LastUpdates = append(ukCSLRecord.LastUpdates, csvRecord[LastUpdatedIdx])
		}
	}
	if csvRecord[UKSancListDateIdx] != "" {
		if !arrayContains(ukCSLRecord.SanctionListDates, csvRecord[UKSancListDateIdx]) {
			ukCSLRecord.SanctionListDates = append(ukCSLRecord.SanctionListDates, csvRecord[UKSancListDateIdx])
		}
	}
	if csvRecord[GroupdIdx] != "" {
		groupID, _ := strconv.Atoi(csvRecord[GroupdIdx])
		ukCSLRecord.GroupID = groupID
	}
}

func ReadSanctionsListFile(f io.ReadCloser) ([]SanctionsListRecord, SanctionsListMap, error) {
	if f == nil {
		return nil, nil, errors.New("uk sanctions list file is empty or missing")
	}
	defer f.Close()

	rows, rowsMap, err := parseSanctionsListCSV(f)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func parseSanctionsListCSV(r io.Reader) ([]SanctionsListRecord, SanctionsListMap, error) {
	reader := csv.NewReader(r)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // Allow variable field counts

	report := SanctionsListMap{}

	rowNum := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			// Skip malformed rows
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				rowNum++
				continue
			}
			return nil, nil, err
		}
		rowNum++

		// Skip first two rows (Report Date and Header)
		if rowNum <= 2 {
			continue
		}

		if len(record) <= UKSL_UniqueIDIdx {
			continue // skip empty or malformed records
		}

		uniqueID := strings.TrimSpace(record[UKSL_UniqueIDIdx])
		if uniqueID == "" {
			continue
		}

		// Group by UniqueID - multiple rows may exist for same entity
		if val, ok := report[uniqueID]; !ok {
			row := new(SanctionsListRecord)
			row.UniqueID = uniqueID
			unmarshalSanctionsListRecord(record, row)
			report[uniqueID] = row
		} else {
			unmarshalSanctionsListRecord(record, val)
		}
	}

	var totalReport []SanctionsListRecord
	for _, row := range report {
		totalReport = append(totalReport, *row)
	}
	return totalReport, report, nil
}

func unmarshalSanctionsListRecord(record []string, ukSLRecord *SanctionsListRecord) {
	getField := func(idx int) string {
		if idx < len(record) {
			return strings.TrimSpace(record[idx])
		}
		return ""
	}

	// Basic fields (only set if not already set)
	if ukSLRecord.LastUpdated == "" {
		ukSLRecord.LastUpdated = getField(UKSL_LastUpdatedIdx)
	}
	if ukSLRecord.OFSIGroupID == "" {
		ukSLRecord.OFSIGroupID = getField(UKSL_OFSI_GroupIDIdx)
	}
	if ukSLRecord.UNReferenceNumber == "" {
		ukSLRecord.UNReferenceNumber = getField(UKSL_UNReferenceNumberIdx)
	}

	// Consolidate names (Name6 is surname, Name1-5 are given names)
	var names []string
	if v := getField(UKSL_Name6Idx); v != "" {
		names = append(names, v)
	}
	if v := getField(UKSL_Name1Idx); v != "" {
		names = append(names, v)
	}
	if v := getField(UKSL_Name2Idx); v != "" {
		names = append(names, v)
	}
	if v := getField(UKSL_Name3Idx); v != "" {
		names = append(names, v)
	}
	if v := getField(UKSL_Name4Idx); v != "" {
		names = append(names, v)
	}
	if v := getField(UKSL_Name5Idx); v != "" {
		names = append(names, v)
	}
	name := strings.Join(names, " ")
	if name != "" && !arrayContains(ukSLRecord.Names, name) {
		ukSLRecord.Names = append(ukSLRecord.Names, name)
	}

	// Title
	if ukSLRecord.NameTitle == "" {
		ukSLRecord.NameTitle = getField(UKSL_TitleIdx)
	}

	// Non-Latin script names
	if v := getField(UKSL_NonLatinScriptIdx); v != "" && !arrayContains(ukSLRecord.NonLatinScriptNames, v) {
		ukSLRecord.NonLatinScriptNames = append(ukSLRecord.NonLatinScriptNames, v)
	}

	// Entity type (from Designation Type column - Individual, Entity, or Ship)
	if ukSLRecord.EntityType == nil {
		cellValue := getField(UKSL_EntityTypeIdx)
		if cellValue != "" {
			entityType := EntityStringMap[cellValue]
			ukSLRecord.EntityType = &entityType
		}
	}

	// Consolidate addresses
	var addressParts []string
	if v := getField(UKSL_AddressLine1Idx); v != "" {
		addressParts = append(addressParts, v)
	}
	if v := getField(UKSL_AddressLine2Idx); v != "" {
		addressParts = append(addressParts, v)
	}
	if v := getField(UKSL_AddressLine3Idx); v != "" {
		addressParts = append(addressParts, v)
	}
	if v := getField(UKSL_AddressLine4Idx); v != "" {
		addressParts = append(addressParts, v)
	}
	if v := getField(UKSL_AddressLine5Idx); v != "" {
		addressParts = append(addressParts, v)
	}
	if v := getField(UKSL_AddressLine6Idx); v != "" {
		addressParts = append(addressParts, v)
		if !arrayContains(ukSLRecord.StateLocalities, v) {
			ukSLRecord.StateLocalities = append(ukSLRecord.StateLocalities, v)
		}
	}

	// Postal code
	postalCode := getField(UKSL_PostalCodeIdx)
	if postalCode != "" && !arrayContains(ukSLRecord.AddressPostalCodes, postalCode) {
		ukSLRecord.AddressPostalCodes = append(ukSLRecord.AddressPostalCodes, postalCode)
	}

	// Address country
	addrCountry := getField(UKSL_AddressCountryIdx)
	if addrCountry != "" && !arrayContains(ukSLRecord.AddressCountries, addrCountry) {
		ukSLRecord.AddressCountries = append(ukSLRecord.AddressCountries, addrCountry)
	}

	// Full address string
	address := strings.Join(addressParts, ", ")
	if address != "" && !arrayContains(ukSLRecord.Addresses, address) {
		ukSLRecord.Addresses = append(ukSLRecord.Addresses, address)
	}

	// Birth information
	if ukSLRecord.CountryOfBirth == "" {
		ukSLRecord.CountryOfBirth = getField(UKSL_CountryOfBirthIdx)
	}
	if ukSLRecord.TownOfBirth == "" {
		ukSLRecord.TownOfBirth = getField(UKSL_TownOfBirthIdx)
	}

	// New fields from CSV
	if ukSLRecord.DOB == "" {
		ukSLRecord.DOB = getField(UKSL_DOBIdx)
	}
	if ukSLRecord.Nationality == "" {
		ukSLRecord.Nationality = getField(UKSL_NationalityIdx)
	}
	if ukSLRecord.PassportNumber == "" {
		ukSLRecord.PassportNumber = getField(UKSL_PassportNumberIdx)
	}
	if ukSLRecord.PassportAdditionalInfo == "" {
		ukSLRecord.PassportAdditionalInfo = getField(UKSL_PassportAdditionalIdx)
	}
	if ukSLRecord.NationalIDNumber == "" {
		ukSLRecord.NationalIDNumber = getField(UKSL_NationalIDNumberIdx)
	}
	if ukSLRecord.NationalIDAdditionalInfo == "" {
		ukSLRecord.NationalIDAdditionalInfo = getField(UKSL_NationalIDAdditionalIdx)
	}
	if ukSLRecord.Position == "" {
		ukSLRecord.Position = getField(UKSL_PositionIdx)
	}
	if ukSLRecord.Gender == "" {
		ukSLRecord.Gender = getField(UKSL_GenderIdx)
	}
	if ukSLRecord.Regime == "" {
		ukSLRecord.Regime = getField(UKSL_RegimeNameIdx)
	}
	if ukSLRecord.DateDesignated == "" {
		ukSLRecord.DateDesignated = getField(UKSL_DateDesignatedIdx)
	}
	if ukSLRecord.OtherInfo == "" {
		ukSLRecord.OtherInfo = getField(UKSL_OtherInfoIdx)
	}

	// Vessel specific fields
	if ukSLRecord.IMONumber == "" {
		ukSLRecord.IMONumber = getField(UKSL_IMONumberIdx)
	}
	if ukSLRecord.VesselType == "" {
		ukSLRecord.VesselType = getField(UKSL_TypeOfShipIdx)
	}
	if ukSLRecord.Tonnage == "" {
		ukSLRecord.Tonnage = getField(UKSL_TonnageIdx)
	}
	if ukSLRecord.VesselFlag == "" {
		ukSLRecord.VesselFlag = getField(UKSL_CurrentFlagIdx)
	}
	if ukSLRecord.VesselOwner == "" {
		ukSLRecord.VesselOwner = getField(UKSL_CurrentOwnerIdx)
	}

	// Business specific fields
	if ukSLRecord.BusinessRegNumber == "" {
		ukSLRecord.BusinessRegNumber = getField(UKSL_BusinessRegNumberIdx)
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
