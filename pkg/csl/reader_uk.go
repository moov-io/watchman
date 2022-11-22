package csl

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadUKFile(path string) ([]*UKCSLRecord, UKCSL, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()

	rows, rowsMap, err := ParseUK(fd)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsMap, nil
}

func ParseUK(r io.Reader) ([]*UKCSLRecord, UKCSL, error) {
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
			unmarshalUKRecord(record, row)

			report[groupID] = row
		} else {
			// we found an entry in the map and need to append
			unmarshalUKRecord(record, val)
		}

	}
	var totalReport []*UKCSLRecord
	for _, row := range report {
		totalReport = append(totalReport, row)
	}
	return totalReport, report, nil
}

func arrayContains(checkArray []string, nameToCheck string) bool {
	var nameAlreadyExists bool = false
	for _, name := range checkArray {
		if name == nameToCheck {
			nameAlreadyExists = true
			break
		}
	}
	return nameAlreadyExists
}

func unmarshalUKRecord(csvRecord []string, ukCSLRecord *UKCSLRecord) {
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
	if csvRecord[AddressOneIdx] != "" {
		if !arrayContains(ukCSLRecord.Addresses, csvRecord[AddressOneIdx]) {
			ukCSLRecord.Addresses = append(ukCSLRecord.Addresses, csvRecord[AddressOneIdx])
		}
	}
	if csvRecord[AddressTwoIdx] != "" {
		if !arrayContains(ukCSLRecord.AddressesTwo, csvRecord[AddressTwoIdx]) {
			ukCSLRecord.AddressesTwo = append(ukCSLRecord.AddressesTwo, csvRecord[AddressTwoIdx])
		}
	}
	if csvRecord[AddressThreeIdx] != "" {
		if !arrayContains(ukCSLRecord.AddressesThree, csvRecord[AddressThreeIdx]) {
			ukCSLRecord.AddressesThree = append(ukCSLRecord.AddressesThree, csvRecord[AddressThreeIdx])
		}
	}
	if csvRecord[AddressFourIdx] != "" {
		if !arrayContains(ukCSLRecord.AddressesFour, csvRecord[AddressFourIdx]) {
			ukCSLRecord.AddressesFour = append(ukCSLRecord.AddressesFour, csvRecord[AddressFourIdx])
		}
	}
	if csvRecord[AddressFiveIdx] != "" {
		if !arrayContains(ukCSLRecord.AddressesFive, csvRecord[AddressFiveIdx]) {
			ukCSLRecord.AddressesFive = append(ukCSLRecord.AddressesFive, csvRecord[AddressFiveIdx])
		}
	}
	if csvRecord[AddressSixIdx] != "" {
		if !arrayContains(ukCSLRecord.AddressesSix, csvRecord[AddressSixIdx]) {
			ukCSLRecord.AddressesSix = append(ukCSLRecord.AddressesSix, csvRecord[AddressSixIdx])
		}
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
	if csvRecord[GroupTypeIdx] != "" {
		if !arrayContains(ukCSLRecord.GroupTypes, csvRecord[GroupTypeIdx]) {
			ukCSLRecord.GroupTypes = append(ukCSLRecord.GroupTypes, csvRecord[GroupTypeIdx])
		}
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
