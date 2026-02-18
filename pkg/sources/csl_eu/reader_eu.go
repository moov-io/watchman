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

// minRecordColumns is the minimum number of columns required for a valid CSV record
const minRecordColumns = 2

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

		if len(record) < minRecordColumns {
			continue // skip empty or malformed records
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
	recordLen := len(csvRecord)

	// Minimum required columns check
	if recordLen <= EntityLogicalIdx {
		return
	}
	euCSLRecord.EntityLogicalID, _ = strconv.Atoi(csvRecord[EntityLogicalIdx])

	// === Entity fields (columns 0-15) ===
	if recordLen > FileGenerationDateIdx && csvRecord[FileGenerationDateIdx] != "" {
		euCSLRecord.FileGenerationDate = csvRecord[FileGenerationDateIdx]
	}
	if recordLen > ReferenceNumberIdx && csvRecord[ReferenceNumberIdx] != "" {
		euCSLRecord.EntityReferenceNumber = csvRecord[ReferenceNumberIdx]
	}
	if recordLen > EntityUnitedNationIDIdx && csvRecord[EntityUnitedNationIDIdx] != "" {
		euCSLRecord.EntityUnitedNationID = csvRecord[EntityUnitedNationIDIdx]
	}
	if recordLen > EntityDesignationDateIdx && csvRecord[EntityDesignationDateIdx] != "" {
		euCSLRecord.EntityDesignationDate = csvRecord[EntityDesignationDateIdx]
	}
	if recordLen > EntityDesignationDetailsIdx && csvRecord[EntityDesignationDetailsIdx] != "" {
		euCSLRecord.EntityDesignationDetails = csvRecord[EntityDesignationDetailsIdx]
	}
	if recordLen > EntityRemarkIdx && csvRecord[EntityRemarkIdx] != "" {
		euCSLRecord.EntityRemark = csvRecord[EntityRemarkIdx]
	}
	if recordLen > EntitySubjectTypeIdx && csvRecord[EntitySubjectTypeIdx] != "" {
		euCSLRecord.EntitySubjectType = csvRecord[EntitySubjectTypeIdx]
	}
	if recordLen > EntitySubjectTypeCodeIdx && csvRecord[EntitySubjectTypeCodeIdx] != "" {
		euCSLRecord.EntitySubjectTypeCode = csvRecord[EntitySubjectTypeCodeIdx]
	}
	if recordLen > EntityRegulationTypeIdx && csvRecord[EntityRegulationTypeIdx] != "" {
		euCSLRecord.EntityRegulationType = csvRecord[EntityRegulationTypeIdx]
	}
	if recordLen > EntityRegulationOrgTypeIdx && csvRecord[EntityRegulationOrgTypeIdx] != "" {
		euCSLRecord.EntityRegulationOrgType = csvRecord[EntityRegulationOrgTypeIdx]
	}
	if recordLen > EntityRegulationPubDateIdx && csvRecord[EntityRegulationPubDateIdx] != "" {
		euCSLRecord.EntityRegulationPubDate = csvRecord[EntityRegulationPubDateIdx]
	}
	if recordLen > EntityRegulationEntryDateIdx && csvRecord[EntityRegulationEntryDateIdx] != "" {
		euCSLRecord.EntityRegulationEntryIntoForce = csvRecord[EntityRegulationEntryDateIdx]
	}
	if recordLen > EntityRegulationNumberTitleIdx && csvRecord[EntityRegulationNumberTitleIdx] != "" {
		euCSLRecord.EntityRegulationNumberTitle = csvRecord[EntityRegulationNumberTitleIdx]
	}
	if recordLen > EntityRegulationProgrammeIdx && csvRecord[EntityRegulationProgrammeIdx] != "" {
		euCSLRecord.EntityRegulationProgramme = csvRecord[EntityRegulationProgrammeIdx]
	}
	if recordLen > EntityRegulationPubURLIdx && csvRecord[EntityRegulationPubURLIdx] != "" {
		euCSLRecord.EntityPublicationURL = csvRecord[EntityRegulationPubURLIdx]
	}

	// === NameAlias fields (columns 16-33) ===
	unmarshalNameAlias(csvRecord, recordLen, euCSLRecord)

	// === Address fields (columns 34-53) ===
	unmarshalAddress(csvRecord, recordLen, euCSLRecord)

	// === BirthDate fields (columns 54-77) ===
	unmarshalBirthDate(csvRecord, recordLen, euCSLRecord)

	// === Identification fields (columns 78-104) ===
	unmarshalIdentification(csvRecord, recordLen, euCSLRecord)

	// === Citizenship fields (columns 105-117) ===
	unmarshalCitizenship(csvRecord, recordLen, euCSLRecord)
}

func unmarshalNameAlias(csvRecord []string, recordLen int, r *CSLRecord) {
	if recordLen > NameAliasLastNameIdx && csvRecord[NameAliasLastNameIdx] != "" {
		appendUnique(&r.NameAliasLastNames, csvRecord[NameAliasLastNameIdx])
	}
	if recordLen > NameAliasFirstNameIdx && csvRecord[NameAliasFirstNameIdx] != "" {
		appendUnique(&r.NameAliasFirstNames, csvRecord[NameAliasFirstNameIdx])
	}
	if recordLen > NameAliasMiddleNameIdx && csvRecord[NameAliasMiddleNameIdx] != "" {
		appendUnique(&r.NameAliasMiddleNames, csvRecord[NameAliasMiddleNameIdx])
	}
	if recordLen > NameAliasWholeNameIdx && csvRecord[NameAliasWholeNameIdx] != "" {
		appendUnique(&r.NameAliasWholeNames, csvRecord[NameAliasWholeNameIdx])
	}
	if recordLen > NameAliasNameLanguageIdx && csvRecord[NameAliasNameLanguageIdx] != "" {
		appendUnique(&r.NameAliasNameLanguages, csvRecord[NameAliasNameLanguageIdx])
	}
	if recordLen > NameAliasGenderIdx && csvRecord[NameAliasGenderIdx] != "" {
		appendUnique(&r.NameAliasGenders, csvRecord[NameAliasGenderIdx])
	}
	if recordLen > NameAliasTitleIdx && csvRecord[NameAliasTitleIdx] != "" {
		appendUnique(&r.NameAliasTitles, csvRecord[NameAliasTitleIdx])
	}
	if recordLen > NameAliasFunctionIdx && csvRecord[NameAliasFunctionIdx] != "" {
		appendUnique(&r.NameAliasFunctions, csvRecord[NameAliasFunctionIdx])
	}
	if recordLen > NameAliasLogicalIDIdx && csvRecord[NameAliasLogicalIDIdx] != "" {
		if id, err := strconv.Atoi(csvRecord[NameAliasLogicalIDIdx]); err == nil {
			appendUniqueInt(&r.NameAliasLogicalIDs, id)
		}
	}
	if recordLen > NameAliasRegLanguageIdx && csvRecord[NameAliasRegLanguageIdx] != "" {
		appendUnique(&r.NameAliasRegLanguages, csvRecord[NameAliasRegLanguageIdx])
	}
	if recordLen > NameAliasRemarkIdx && csvRecord[NameAliasRemarkIdx] != "" {
		appendUnique(&r.NameAliasRemarks, csvRecord[NameAliasRemarkIdx])
	}
	if recordLen > NameAliasRegTypeIdx && csvRecord[NameAliasRegTypeIdx] != "" {
		appendUnique(&r.NameAliasRegTypes, csvRecord[NameAliasRegTypeIdx])
	}
	if recordLen > NameAliasRegOrgTypeIdx && csvRecord[NameAliasRegOrgTypeIdx] != "" {
		appendUnique(&r.NameAliasRegOrgTypes, csvRecord[NameAliasRegOrgTypeIdx])
	}
	if recordLen > NameAliasRegPubDateIdx && csvRecord[NameAliasRegPubDateIdx] != "" {
		appendUnique(&r.NameAliasRegPubDates, csvRecord[NameAliasRegPubDateIdx])
	}
	if recordLen > NameAliasRegEntryDateIdx && csvRecord[NameAliasRegEntryDateIdx] != "" {
		appendUnique(&r.NameAliasRegEntryDates, csvRecord[NameAliasRegEntryDateIdx])
	}
	if recordLen > NameAliasRegNumberTitleIdx && csvRecord[NameAliasRegNumberTitleIdx] != "" {
		appendUnique(&r.NameAliasRegNumberTitles, csvRecord[NameAliasRegNumberTitleIdx])
	}
	if recordLen > NameAliasRegProgrammeIdx && csvRecord[NameAliasRegProgrammeIdx] != "" {
		appendUnique(&r.NameAliasRegProgrammes, csvRecord[NameAliasRegProgrammeIdx])
	}
	if recordLen > NameAliasRegPubURLIdx && csvRecord[NameAliasRegPubURLIdx] != "" {
		appendUnique(&r.NameAliasRegPubURLs, csvRecord[NameAliasRegPubURLIdx])
	}
}

func unmarshalAddress(csvRecord []string, recordLen int, r *CSLRecord) {
	if recordLen > AddressCityIdx && csvRecord[AddressCityIdx] != "" {
		appendUnique(&r.AddressCities, csvRecord[AddressCityIdx])
	}
	if recordLen > AddressStreetIdx && csvRecord[AddressStreetIdx] != "" {
		appendUnique(&r.AddressStreets, csvRecord[AddressStreetIdx])
	}
	if recordLen > AddressPoBoxIdx && csvRecord[AddressPoBoxIdx] != "" {
		appendUnique(&r.AddressPoBoxes, csvRecord[AddressPoBoxIdx])
	}
	if recordLen > AddressZipCodeIdx && csvRecord[AddressZipCodeIdx] != "" {
		appendUnique(&r.AddressZipCodes, csvRecord[AddressZipCodeIdx])
	}
	if recordLen > AddressRegionIdx && csvRecord[AddressRegionIdx] != "" {
		appendUnique(&r.AddressRegions, csvRecord[AddressRegionIdx])
	}
	if recordLen > AddressPlaceIdx && csvRecord[AddressPlaceIdx] != "" {
		appendUnique(&r.AddressPlaces, csvRecord[AddressPlaceIdx])
	}
	if recordLen > AddressAsAtListingTimeIdx && csvRecord[AddressAsAtListingTimeIdx] != "" {
		appendUnique(&r.AddressAsAtListingTimes, csvRecord[AddressAsAtListingTimeIdx])
	}
	if recordLen > AddressContactInfoIdx && csvRecord[AddressContactInfoIdx] != "" {
		appendUnique(&r.AddressContactInfos, csvRecord[AddressContactInfoIdx])
	}
	if recordLen > AddressCountryISOIdx && csvRecord[AddressCountryISOIdx] != "" {
		appendUnique(&r.AddressCountryISOs, csvRecord[AddressCountryISOIdx])
	}
	if recordLen > AddressCountryDescIdx && csvRecord[AddressCountryDescIdx] != "" {
		appendUnique(&r.AddressCountryDescriptions, csvRecord[AddressCountryDescIdx])
	}
	if recordLen > AddressLogicalIDIdx && csvRecord[AddressLogicalIDIdx] != "" {
		if id, err := strconv.Atoi(csvRecord[AddressLogicalIDIdx]); err == nil {
			appendUniqueInt(&r.AddressLogicalIDs, id)
		}
	}
	if recordLen > AddressRegLanguageIdx && csvRecord[AddressRegLanguageIdx] != "" {
		appendUnique(&r.AddressRegLanguages, csvRecord[AddressRegLanguageIdx])
	}
	if recordLen > AddressRemarkIdx && csvRecord[AddressRemarkIdx] != "" {
		appendUnique(&r.AddressRemarks, csvRecord[AddressRemarkIdx])
	}
	if recordLen > AddressRegTypeIdx && csvRecord[AddressRegTypeIdx] != "" {
		appendUnique(&r.AddressRegTypes, csvRecord[AddressRegTypeIdx])
	}
	if recordLen > AddressRegOrgTypeIdx && csvRecord[AddressRegOrgTypeIdx] != "" {
		appendUnique(&r.AddressRegOrgTypes, csvRecord[AddressRegOrgTypeIdx])
	}
	if recordLen > AddressRegPubDateIdx && csvRecord[AddressRegPubDateIdx] != "" {
		appendUnique(&r.AddressRegPubDates, csvRecord[AddressRegPubDateIdx])
	}
	if recordLen > AddressRegEntryDateIdx && csvRecord[AddressRegEntryDateIdx] != "" {
		appendUnique(&r.AddressRegEntryDates, csvRecord[AddressRegEntryDateIdx])
	}
	if recordLen > AddressRegNumberTitleIdx && csvRecord[AddressRegNumberTitleIdx] != "" {
		appendUnique(&r.AddressRegNumberTitles, csvRecord[AddressRegNumberTitleIdx])
	}
	if recordLen > AddressRegProgrammeIdx && csvRecord[AddressRegProgrammeIdx] != "" {
		appendUnique(&r.AddressRegProgrammes, csvRecord[AddressRegProgrammeIdx])
	}
	if recordLen > AddressRegPubURLIdx && csvRecord[AddressRegPubURLIdx] != "" {
		appendUnique(&r.AddressRegPubURLs, csvRecord[AddressRegPubURLIdx])
	}
}

func unmarshalBirthDate(csvRecord []string, recordLen int, r *CSLRecord) {
	if recordLen > BirthDateIdx && csvRecord[BirthDateIdx] != "" {
		appendUnique(&r.BirthDates, csvRecord[BirthDateIdx])
	}
	if recordLen > BirthDayIdx && csvRecord[BirthDayIdx] != "" {
		appendUnique(&r.BirthDays, csvRecord[BirthDayIdx])
	}
	if recordLen > BirthMonthIdx && csvRecord[BirthMonthIdx] != "" {
		appendUnique(&r.BirthMonths, csvRecord[BirthMonthIdx])
	}
	if recordLen > BirthYearIdx && csvRecord[BirthYearIdx] != "" {
		appendUnique(&r.BirthYears, csvRecord[BirthYearIdx])
	}
	if recordLen > BirthYearRangeFromIdx && csvRecord[BirthYearRangeFromIdx] != "" {
		appendUnique(&r.BirthYearRangeFroms, csvRecord[BirthYearRangeFromIdx])
	}
	if recordLen > BirthYearRangeToIdx && csvRecord[BirthYearRangeToIdx] != "" {
		appendUnique(&r.BirthYearRangeTos, csvRecord[BirthYearRangeToIdx])
	}
	if recordLen > BirthCircaIdx && csvRecord[BirthCircaIdx] != "" {
		appendUnique(&r.BirthCircas, csvRecord[BirthCircaIdx])
	}
	if recordLen > BirthCalendarTypeIdx && csvRecord[BirthCalendarTypeIdx] != "" {
		appendUnique(&r.BirthCalendarTypes, csvRecord[BirthCalendarTypeIdx])
	}
	if recordLen > BirthZipCodeIdx && csvRecord[BirthZipCodeIdx] != "" {
		appendUnique(&r.BirthZipCodes, csvRecord[BirthZipCodeIdx])
	}
	if recordLen > BirthRegionIdx && csvRecord[BirthRegionIdx] != "" {
		appendUnique(&r.BirthRegions, csvRecord[BirthRegionIdx])
	}
	if recordLen > BirthPlaceIdx && csvRecord[BirthPlaceIdx] != "" {
		appendUnique(&r.BirthPlaces, csvRecord[BirthPlaceIdx])
	}
	if recordLen > BirthCityIdx && csvRecord[BirthCityIdx] != "" {
		appendUnique(&r.BirthCities, csvRecord[BirthCityIdx])
	}
	if recordLen > BirthCountryISOIdx && csvRecord[BirthCountryISOIdx] != "" {
		appendUnique(&r.BirthCountryISOs, csvRecord[BirthCountryISOIdx])
	}
	if recordLen > BirthCountryDescIdx && csvRecord[BirthCountryDescIdx] != "" {
		appendUnique(&r.BirthCountries, csvRecord[BirthCountryDescIdx])
	}
	if recordLen > BirthLogicalIDIdx && csvRecord[BirthLogicalIDIdx] != "" {
		if id, err := strconv.Atoi(csvRecord[BirthLogicalIDIdx]); err == nil {
			appendUniqueInt(&r.BirthLogicalIDs, id)
		}
	}
	if recordLen > BirthRegLanguageIdx && csvRecord[BirthRegLanguageIdx] != "" {
		appendUnique(&r.BirthRegLanguages, csvRecord[BirthRegLanguageIdx])
	}
	if recordLen > BirthRemarkIdx && csvRecord[BirthRemarkIdx] != "" {
		appendUnique(&r.BirthRemarks, csvRecord[BirthRemarkIdx])
	}
	if recordLen > BirthRegTypeIdx && csvRecord[BirthRegTypeIdx] != "" {
		appendUnique(&r.BirthRegTypes, csvRecord[BirthRegTypeIdx])
	}
	if recordLen > BirthRegOrgTypeIdx && csvRecord[BirthRegOrgTypeIdx] != "" {
		appendUnique(&r.BirthRegOrgTypes, csvRecord[BirthRegOrgTypeIdx])
	}
	if recordLen > BirthRegPubDateIdx && csvRecord[BirthRegPubDateIdx] != "" {
		appendUnique(&r.BirthRegPubDates, csvRecord[BirthRegPubDateIdx])
	}
	if recordLen > BirthRegEntryDateIdx && csvRecord[BirthRegEntryDateIdx] != "" {
		appendUnique(&r.BirthRegEntryDates, csvRecord[BirthRegEntryDateIdx])
	}
	if recordLen > BirthRegNumberTitleIdx && csvRecord[BirthRegNumberTitleIdx] != "" {
		appendUnique(&r.BirthRegNumberTitle, csvRecord[BirthRegNumberTitleIdx])
	}
	if recordLen > BirthRegProgrammeIdx && csvRecord[BirthRegProgrammeIdx] != "" {
		appendUnique(&r.BirthRegProgrammes, csvRecord[BirthRegProgrammeIdx])
	}
	if recordLen > BirthRegPubURLIdx && csvRecord[BirthRegPubURLIdx] != "" {
		appendUnique(&r.BirthRegPubURLs, csvRecord[BirthRegPubURLIdx])
	}
}

func unmarshalIdentification(csvRecord []string, recordLen int, r *CSLRecord) {
	// Check if there's identification data
	if recordLen <= IdentificationNumberIdx || csvRecord[IdentificationNumberIdx] == "" {
		return
	}

	id := IdentificationInfo{
		Number: csvRecord[IdentificationNumberIdx],
	}

	if recordLen > IdentificationDiplomaticIdx {
		id.Diplomatic = csvRecord[IdentificationDiplomaticIdx]
	}
	if recordLen > IdentificationKnownExpiredIdx {
		id.KnownExpired = csvRecord[IdentificationKnownExpiredIdx]
	}
	if recordLen > IdentificationKnownFalseIdx {
		id.KnownFalse = csvRecord[IdentificationKnownFalseIdx]
	}
	if recordLen > IdentificationReportedLostIdx {
		id.ReportedLost = csvRecord[IdentificationReportedLostIdx]
	}
	if recordLen > IdentificationRevokedIdx {
		id.RevokedByIssuer = csvRecord[IdentificationRevokedIdx]
	}
	if recordLen > IdentificationIssuedByIdx {
		id.IssuedBy = csvRecord[IdentificationIssuedByIdx]
	}
	if recordLen > IdentificationIssuedDateIdx {
		id.IssuedDate = csvRecord[IdentificationIssuedDateIdx]
	}
	if recordLen > IdentificationValidFromIdx {
		id.ValidFrom = csvRecord[IdentificationValidFromIdx]
	}
	if recordLen > IdentificationValidToIdx {
		id.ValidTo = csvRecord[IdentificationValidToIdx]
	}
	if recordLen > IdentificationLatinNumberIdx {
		id.LatinNumber = csvRecord[IdentificationLatinNumberIdx]
	}
	if recordLen > IdentificationNameOnDocIdx {
		id.NameOnDocument = csvRecord[IdentificationNameOnDocIdx]
	}
	if recordLen > IdentificationTypeCodeIdx {
		id.TypeCode = csvRecord[IdentificationTypeCodeIdx]
	}
	if recordLen > IdentificationTypeDescIdx {
		id.TypeDescription = csvRecord[IdentificationTypeDescIdx]
	}
	if recordLen > IdentificationRegionIdx {
		id.Region = csvRecord[IdentificationRegionIdx]
	}
	if recordLen > IdentificationCountryISOIdx {
		id.CountryISO = csvRecord[IdentificationCountryISOIdx]
	}
	if recordLen > IdentificationCountryDescIdx {
		id.CountryDesc = csvRecord[IdentificationCountryDescIdx]
	}
	if recordLen > IdentificationLogicalIDIdx && csvRecord[IdentificationLogicalIDIdx] != "" {
		id.LogicalID, _ = strconv.Atoi(csvRecord[IdentificationLogicalIDIdx])
	}
	if recordLen > IdentificationRegLanguageIdx {
		id.RegLanguage = csvRecord[IdentificationRegLanguageIdx]
	}
	if recordLen > IdentificationRemarkIdx {
		id.Remark = csvRecord[IdentificationRemarkIdx]
	}
	if recordLen > IdentificationRegTypeIdx {
		id.RegType = csvRecord[IdentificationRegTypeIdx]
	}
	if recordLen > IdentificationRegOrgTypeIdx {
		id.RegOrgType = csvRecord[IdentificationRegOrgTypeIdx]
	}
	if recordLen > IdentificationRegPubDateIdx {
		id.RegPubDate = csvRecord[IdentificationRegPubDateIdx]
	}
	if recordLen > IdentificationRegEntryDateIdx {
		id.RegEntryDate = csvRecord[IdentificationRegEntryDateIdx]
	}
	if recordLen > IdentificationRegNumTitleIdx {
		id.RegNumberTitle = csvRecord[IdentificationRegNumTitleIdx]
	}
	if recordLen > IdentificationRegProgrammeIdx {
		id.RegProgramme = csvRecord[IdentificationRegProgrammeIdx]
	}
	if recordLen > IdentificationRegPubURLIdx {
		id.RegPubURL = csvRecord[IdentificationRegPubURLIdx]
	}

	// Check if this identification already exists (by number + type)
	for _, existing := range r.Identifications {
		if existing.Number == id.Number && existing.TypeCode == id.TypeCode {
			return
		}
	}
	r.Identifications = append(r.Identifications, id)
}

func unmarshalCitizenship(csvRecord []string, recordLen int, r *CSLRecord) {
	if recordLen > CitizenshipRegionIdx && csvRecord[CitizenshipRegionIdx] != "" {
		appendUnique(&r.CitizenshipRegions, csvRecord[CitizenshipRegionIdx])
	}
	if recordLen > CitizenshipCountryISOIdx && csvRecord[CitizenshipCountryISOIdx] != "" {
		appendUnique(&r.CitizenshipCountryISOs, csvRecord[CitizenshipCountryISOIdx])
	}
	if recordLen > CitizenshipCountryDescIdx && csvRecord[CitizenshipCountryDescIdx] != "" {
		appendUnique(&r.Citizenships, csvRecord[CitizenshipCountryDescIdx])
	}
	if recordLen > CitizenshipLogicalIDIdx && csvRecord[CitizenshipLogicalIDIdx] != "" {
		if id, err := strconv.Atoi(csvRecord[CitizenshipLogicalIDIdx]); err == nil {
			appendUniqueInt(&r.CitizenshipLogicalIDs, id)
		}
	}
	if recordLen > CitizenshipRegLanguageIdx && csvRecord[CitizenshipRegLanguageIdx] != "" {
		appendUnique(&r.CitizenshipRegLanguages, csvRecord[CitizenshipRegLanguageIdx])
	}
	if recordLen > CitizenshipRemarkIdx && csvRecord[CitizenshipRemarkIdx] != "" {
		appendUnique(&r.CitizenshipRemarks, csvRecord[CitizenshipRemarkIdx])
	}
	if recordLen > CitizenshipRegTypeIdx && csvRecord[CitizenshipRegTypeIdx] != "" {
		appendUnique(&r.CitizenshipRegTypes, csvRecord[CitizenshipRegTypeIdx])
	}
	if recordLen > CitizenshipRegOrgTypeIdx && csvRecord[CitizenshipRegOrgTypeIdx] != "" {
		appendUnique(&r.CitizenshipRegOrgTypes, csvRecord[CitizenshipRegOrgTypeIdx])
	}
	if recordLen > CitizenshipRegPubDateIdx && csvRecord[CitizenshipRegPubDateIdx] != "" {
		appendUnique(&r.CitizenshipRegPubDates, csvRecord[CitizenshipRegPubDateIdx])
	}
	if recordLen > CitizenshipRegEntryDateIdx && csvRecord[CitizenshipRegEntryDateIdx] != "" {
		appendUnique(&r.CitizenshipRegEntryDates, csvRecord[CitizenshipRegEntryDateIdx])
	}
	if recordLen > CitizenshipRegNumTitleIdx && csvRecord[CitizenshipRegNumTitleIdx] != "" {
		appendUnique(&r.CitizenshipRegNumTitles, csvRecord[CitizenshipRegNumTitleIdx])
	}
	if recordLen > CitizenshipRegProgrammeIdx && csvRecord[CitizenshipRegProgrammeIdx] != "" {
		appendUnique(&r.CitizenshipRegProgrammes, csvRecord[CitizenshipRegProgrammeIdx])
	}
	if recordLen > CitizenshipRegPubURLIdx && csvRecord[CitizenshipRegPubURLIdx] != "" {
		appendUnique(&r.CitizenshipRegPubURLs, csvRecord[CitizenshipRegPubURLIdx])
	}
}

// appendUnique appends value to slice if not already present
func appendUnique(slice *[]string, value string) {
	if value == "" {
		return
	}
	for _, existing := range *slice {
		if existing == value {
			return
		}
	}
	*slice = append(*slice, value)
}

// appendUniqueInt appends value to int slice if not already present
func appendUniqueInt(slice *[]int, value int) {
	for _, existing := range *slice {
		if existing == value {
			return
		}
	}
	*slice = append(*slice, value)
}

// arrayContains checks if a string exists in a slice (kept for backwards compatibility)
func arrayContains(checkArray []string, nameToCheck string) bool {
	if nameToCheck == "" {
		return true
	}
	for _, name := range checkArray {
		if name == nameToCheck {
			return true
		}
	}
	return false
}
