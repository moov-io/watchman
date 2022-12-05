package csl

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/knieriem/odf/ods"
)

func ReadUKCSLFile(path string) ([]*UKCSLRecord, UKCSL, error) {
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
			fmt.Println("record is <= 1", record)
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
	fd, err := ods.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	rows, rowsMap, err := ParseUKSanctionsList(fd)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

// func ParseUKSanctionsList(r io.Reader, fileSize int64) ([]*UKSanctionsListRecord, UKSanctionsListMap, error) {
func ParseUKSanctionsList(file *ods.File) ([]*UKSanctionsListRecord, UKSanctionsListMap, error) {
	// read from the ods document
	var totalReport []*UKSanctionsListRecord
	var report UKSanctionsListMap

	var doc ods.Doc
	if err := file.ParseContent(&doc); err != nil {
		if err != nil {
			return totalReport, report, err
		}
	}

	// unmarshal each row into a uk sanctions list record
	if len(doc.Table) > 0 {
		for _, record := range doc.Table[0].Row {
			if record.IsEmpty() {
				continue
			}

			// need a length of row check since we are using the string representation
			uniqueIDCell := record.Cell[UKSL_UniqueIDIdx]
			uniqueID := uniqueIDCell.Value

			if val, ok := report[uniqueID]; !ok {
				row := new(UKSanctionsListRecord)
				row.UniqueID = uniqueID
				unmarshalUKSanctionsListRecord(record.Cell, row)
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
	if record[UKSL_LastUpdatedIdx].Value != "" && ukSLRecord.LastUpdated == "" {
		ukSLRecord.LastUpdated = record[UKSL_LastUpdatedIdx].Value
	}

	if record[UKSL_OFSI_GroupIDIdx].Value != "" && ukSLRecord.OFSIGroupID == "" {
		ukSLRecord.OFSIGroupID = record[UKSL_OFSI_GroupIDIdx].Value
	}

	if record[UKSL_UNReferenceNumberIdx].Value != "" && ukSLRecord.UNReferenceNumber == "" {
		ukSLRecord.UNReferenceNumber = record[UKSL_UNReferenceNumberIdx].Value
	}

	// consolidate names
	var names []string
	if record[UKSL_Name6Idx].Value != "" {
		names = append(names, record[UKSL_Name6Idx].Value)
	}
	if record[UKSL_Name1Idx].Value != "" {
		names = append(names, record[UKSL_Name1Idx].Value)
	}
	if record[UKSL_Name2Idx].Value != "" {
		names = append(names, record[UKSL_Name2Idx].Value)
	}
	if record[UKSL_Name3Idx].Value != "" {
		names = append(names, record[UKSL_Name3Idx].Value)
	}
	if record[UKSL_Name4Idx].Value != "" {
		names = append(names, record[UKSL_Name4Idx].Value)
	}
	if record[UKSL_Name5Idx].Value != "" {
		names = append(names, record[UKSL_Name5Idx].Value)
	}
	name := strings.Join(names, " ")
	if !arrayContains(ukSLRecord.Names, name) {
		ukSLRecord.Names = append(ukSLRecord.Names, name)
	}
	// consolidate addresses
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
