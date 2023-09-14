package csl

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/knieriem/odf/ods"
)

func ReadUKCSLFile(path string) ([]*UKCSLRecord, UKCSL, error) {
	if path == "" {
		return nil, nil, errors.New("path was empty for ukcsl file")
	}
	fd, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	rows, rowsMap, err := ParseUKCSL(fd)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func ParseUKCSL(r io.Reader) ([]*UKCSLRecord, UKCSL, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = 36

	report := make(UKCSL)
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
			row := new(UKCSLRecord)
			unmarshalUKCSLRecord(record, row)

			report[groupID] = row
		} else {
			// we found an entry in the map and need to append
			unmarshalUKCSLRecord(record, val)
		}

	}
	var totalReport []*UKCSLRecord
	for _, row := range report {
		totalReport = append(totalReport, row)
	}
	return totalReport, report, nil
}

func unmarshalUKCSLRecord(csvRecord []string, ukCSLRecord *UKCSLRecord) {
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

func ReadUKSanctionsListFile(path string) ([]*UKSanctionsListRecord, UKSanctionsListMap, error) {
	if path == "" {
		return nil, nil, errors.New("path was empty for uk sanctions list file")
	}
	fd, err := ods.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	doc := new(ods.Doc)
	err = fd.ParseContent(doc)
	if err != nil {
		return nil, nil, err
	}

	rows, rowsMap, err := parseUKSanctionsList(doc)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func parseUKSanctionsList(doc *ods.Doc) ([]*UKSanctionsListRecord, UKSanctionsListMap, error) {
	// read from the ods document
	var totalReport []*UKSanctionsListRecord
	report := UKSanctionsListMap{}

	// unmarshal each row into a uk sanctions list record
	if len(doc.Table) > 0 {
		for i, record := range doc.Table[0].Row {

			// manually skip the header and extra rows
			if record.IsEmpty() || i <= 2 {
				continue
			}

			// need a length of row check since we are using the string representation
			uniqueIDCell := record.Cell[UKSL_UniqueIDIdx]
			b := new(bytes.Buffer)
			uniqueID := uniqueIDCell.PlainText(b)

			if val, ok := report[uniqueID]; !ok {
				row := new(UKSanctionsListRecord)
				row.UniqueID = uniqueID
				unmarshalUKSanctionsListRecord(record.Cell, row)

				report[uniqueID] = row
			} else {
				unmarshalUKSanctionsListRecord(record.Cell, val)
			}
		}
	}

	for _, row := range report {
		totalReport = append(totalReport, row)
	}
	return totalReport, report, nil
}

func unmarshalUKSanctionsListRecord(record []ods.Cell, ukSLRecord *UKSanctionsListRecord) {
	if len(record) < UKSL_CountryOfBirthIdx {
		return
	}

	b := new(bytes.Buffer)
	if !record[UKSL_LastUpdatedIdx].IsEmpty() && ukSLRecord.LastUpdated == "" {
		ukSLRecord.LastUpdated = record[UKSL_LastUpdatedIdx].PlainText(b)
	}

	if !record[UKSL_OFSI_GroupIDIdx].IsEmpty() && ukSLRecord.OFSIGroupID == "" {
		ukSLRecord.OFSIGroupID = record[UKSL_OFSI_GroupIDIdx].PlainText(b)
	}

	if !record[UKSL_UNReferenceNumberIdx].IsEmpty() && ukSLRecord.UNReferenceNumber == "" {
		ukSLRecord.UNReferenceNumber = record[UKSL_UNReferenceNumberIdx].PlainText(b)
	}

	// consolidate names
	var names []string
	if !record[UKSL_Name6Idx].IsEmpty() {
		names = append(names, record[UKSL_Name6Idx].PlainText(b))
	}
	if !record[UKSL_Name1Idx].IsEmpty() {
		names = append(names, record[UKSL_Name1Idx].PlainText(b))
	}
	if !record[UKSL_Name2Idx].IsEmpty() {
		names = append(names, record[UKSL_Name2Idx].PlainText(b))
	}
	if !record[UKSL_Name3Idx].IsEmpty() {
		names = append(names, record[UKSL_Name3Idx].PlainText(b))
	}
	if !record[UKSL_Name4Idx].IsEmpty() {
		names = append(names, record[UKSL_Name4Idx].PlainText(b))
	}
	if !record[UKSL_Name5Idx].IsEmpty() {
		names = append(names, record[UKSL_Name5Idx].PlainText(b))
	}
	name := strings.Join(names, " ")
	if !strings.EqualFold(strings.TrimSpace(name), "") && !arrayContains(ukSLRecord.Names, name) {
		ukSLRecord.Names = append(ukSLRecord.Names, name)
	}

	if !record[UKSL_NameTypeIdx].IsEmpty() && ukSLRecord.NameTitle == "" {
		ukSLRecord.NameTitle = record[UKSL_NameTypeIdx].PlainText(b)
	}

	if !record[UKSL_NonLatinScriptIdx].IsEmpty() && !arrayContains(ukSLRecord.NonLatinScriptNames, record[UKSL_NonLatinScriptIdx].PlainText(b)) {
		ukSLRecord.NonLatinScriptNames = append(ukSLRecord.NonLatinScriptNames, record[UKSL_NonLatinScriptIdx].PlainText(b))
	}

	if !record[UKSL_EntityTypeIdx].IsEmpty() && ukSLRecord.EntityType == nil {
		cellValue := record[UKSL_EntityTypeIdx].PlainText(b)
		entityType := EntityStringMap[cellValue]
		ukSLRecord.EntityType = &entityType
	}

	// consolidate addresses
	var addresses []string
	addr1 := record[UKSL_AddressLine1Idx]
	addr2 := record[UKSL_AddressLine2Idx]
	addr3 := record[UKSL_AddressLine3Idx]
	addr4 := record[UKSL_AddressLine4Idx]
	addr5 := record[UKSL_AddressLine5Idx]
	addr6 := record[UKSL_AddressLine6Idx]

	if !addr1.IsEmpty() {
		addresses = append(addresses, addr1.PlainText(b))
	}
	if !addr2.IsEmpty() {
		addresses = append(addresses, addr2.PlainText(b))
	}
	if !addr3.IsEmpty() {
		addresses = append(addresses, addr3.PlainText(b))
	}
	if !addr4.IsEmpty() {
		addresses = append(addresses, addr4.PlainText(b))
	}
	if !addr5.IsEmpty() {
		addresses = append(addresses, addr5.PlainText(b))
	}
	addr6Value := addr6.PlainText(b)
	if !addr6.IsEmpty() {
		addresses = append(addresses, addr6Value)
		if !arrayContains(ukSLRecord.StateLocalities, addr6Value) {
			ukSLRecord.StateLocalities = append(ukSLRecord.StateLocalities, addr6Value)
		}
	}
	postalCode := record[UKSL_PostalCodeIdx]
	postalCodeValue := record[UKSL_PostalCodeIdx].PlainText(b)
	if !postalCode.IsEmpty() {
		addresses = append(addresses, postalCodeValue)
	}
	addrCountries := record[UKSL_AddressCountryIdx]
	addrCountriesValue := record[UKSL_AddressCountryIdx].PlainText(b)
	if !addrCountries.IsEmpty() {
		addresses = append(addresses, addrCountriesValue)
		if !arrayContains(ukSLRecord.AddressCountries, addrCountriesValue) {
			ukSLRecord.AddressCountries = append(ukSLRecord.AddressCountries, addrCountriesValue)
		}
	}

	address := strings.Join(addresses, ", ")
	if !strings.EqualFold(strings.TrimSpace(address), "") && !arrayContains(ukSLRecord.Addresses, address) {
		ukSLRecord.Addresses = append(ukSLRecord.Addresses, address)
	}

	cob := record[UKSL_CountryOfBirthIdx]
	cobValue := record[UKSL_CountryOfBirthIdx].PlainText(b)
	if !cob.IsEmpty() && ukSLRecord.CountryOfBirth != "" {
		ukSLRecord.CountryOfBirth = cobValue
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
