// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSL(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", "ConList.csv"))
	if err != nil {
		t.Error(err)
	}
	ukCSL, ukCSLMap, err := ReadCSLFile(fd)
	if err != nil {
		t.Fatal(err)
	}
	if ukCSLMap == nil || ukCSL == nil {
		t.Fatal("failed to parse ConList.csv")
	}

	if len(ukCSL) == 0 {
		t.Fatal("failed to convert map to sheet")
	}

	groupID := 12431

	// expectations
	expectedName := "(GENERAL) ORGANIZATION FOR ENGINEERING INDUSTRIES"
	expectedAddressOne := "PO Box 21120"
	expectedAddressFive := "Baramkeh"
	expectedAddressSix := "Damascus"
	expectedFullAddress := fmt.Sprintf("%s, %s, %s", expectedAddressOne, expectedAddressFive, expectedAddressSix)
	expectedCountry := "Syria"
	expectedOtherInfo := "(UK Sanctions List Ref):SYR0306 (UK Statement of Reasons):Front company for the acquisition of sensitive equipment related to the production of chemical weapons by the Syrian Scientific Studies and Research Center (SSRC) also known as Centre d'études et de recherches syrien (CERS). Owned or controlled by or otherwise associated with the SSRC. (Phone number):(1) +963112121816 (2) +963112121834 (3) +963112212743 (4) +963112214650 (5) +963115110117"
	expectedGroupType := "Entity"
	expectedListedDate := "02/12/2011"
	expectedUKSancListDate := "31/12/2020"
	expectedLastUpdatedDate := "13/05/2022"
	expectedGroupID := 12431

	testRow := ukCSLMap[groupID]

	assert.Greater(t, len(testRow.Names), 0)
	assert.Nil(t, testRow.Titles)
	assert.Nil(t, testRow.DatesOfBirth)
	assert.Nil(t, testRow.TownsOfBirth)
	assert.Nil(t, testRow.CountriesOfBirth)
	assert.Nil(t, testRow.Nationalities)
	assert.Greater(t, len(testRow.Addresses), 0)
	// assert.Greater(t, len(testRow.AddressesTwo), 0)
	// assert.Nil(t, testRow.AddressesThree)
	// assert.Nil(t, testRow.AddressesFour)
	// assert.Greater(t, len(testRow.AddressesFive), 0)
	// assert.Greater(t, len(testRow.AddressesSix), 0)
	assert.Nil(t, testRow.PostalCodes)
	assert.Greater(t, len(testRow.Countries), 0)
	assert.Greater(t, len(testRow.OtherInfos), 0)

	// assertions
	assert.Equal(t, expectedName, testRow.Names[0])
	assert.Equal(t, expectedFullAddress, testRow.Addresses[0])
	// assert.Equal(t, expectedAddressFive, testRow.AddressesFive[0])
	// assert.Equal(t, expectedAddressSix, testRow.AddressesSix[0])
	assert.Equal(t, expectedCountry, testRow.Countries[0])
	assert.Equal(t, expectedOtherInfo, testRow.OtherInfos[0])
	assert.Equal(t, expectedGroupType, testRow.GroupType)
	assert.Equal(t, expectedListedDate, testRow.ListedDates[0])
	assert.Equal(t, expectedLastUpdatedDate, testRow.LastUpdates[0])
	assert.Equal(t, expectedUKSancListDate, testRow.SanctionListDates[0])
	assert.Equal(t, expectedGroupID, testRow.GroupID)
}

func TestReadUKSanctionsList(t *testing.T) {
	// Test CSV parsing with inline test data
	csvData := `Report Date: 13-Feb-2026
Last Updated,Unique ID,OFSI Group ID,UN Reference Number,Name 6,Name 1,Name 2,Name 3,Name 4,Name 5,Name type,Alias strength,Title,Name non-latin script,Non-latin script type,Non-latin script language,Regime Name,Designation Type,Designation source,Sanctions Imposed,Other Information,UK Statement of Reasons,Address Line 1,Address Line 2,Address Line 3,Address Line 4,Address Line 5,Address Line 6,Address Postal Code,Address Country,Phone number,Website,Email address,Date Designated,D.O.B,Nationality(/ies),National Identifier number,National Identifier additional information,Passport number,Passport additional information,Position,Gender,Town of birth,Country of birth,Type of entity,Subsidiaries,Parent company,Business registration number (s),IMO number,Current owner/operator (s),Previous owner/operator (s),Current believed flag of ship,Previous flags,Type of ship,Tonnage of ship,Length of ship,Year Built,Hull identification number (HIN)
12/01/2022,AFG0001,12703,TAe.010,HAJI KHAIRULLAH HAJI SATTAR MONEY EXCHANGE,,,,,,Primary Name,,,"حاجی خيرالله و حاجی ستار صرافی",,,The Afghanistan (Sanctions) (EU Exit) Regulations 2020,Entity,UN,Asset freeze,Test info,,Chaman Central Bazaar,Chaman,Shah Zada Market,,,Baluchistan Province,12345,Pakistan,+123456789,www.example.com,test@example.com,29/06/2012,,,,,,,,,,,,,,,9000000,Test Owner,,Panama,,Cargo,5000,100m,2010,HIN123
26/01/2022,AFG0006,7172,TAi.002,AKHUND,MOHAMMAD,HASSAN,,,,Primary Name,,Mullah,"محمد حسن آخوند",,,The Afghanistan (Sanctions) (EU Exit) Regulations 2020,Individual,UN,Asset freeze|Travel Ban,Test info 2,,,,,,,,,,,,,25/01/2001,01/01/1958,Afghanistan,12345,Mali number,A1234567,Mali passport,"Deputy Minister",Male,Kandahar,Afghanistan,,,,,,,,,,,,,
`

	fd := nopCloser{strings.NewReader(csvData)}
	totalReport, report, err := ReadSanctionsListFile(fd)
	assert.NoError(t, err)

	// test that we get something more than an empty sanctions list record
	assert.NotEmpty(t, totalReport)
	assert.NotEmpty(t, report)
	assert.Len(t, totalReport, 2)

	// Test Entity record
	if record, ok := report["AFG0001"]; ok {
		assert.Equal(t, "12/01/2022", record.LastUpdated)
		assert.Equal(t, "12703", record.OFSIGroupID)
		assert.Equal(t, "TAe.010", record.UNReferenceNumber)
		assert.Equal(t, "HAJI KHAIRULLAH HAJI SATTAR MONEY EXCHANGE", record.Names[0])
		assert.NotEmpty(t, record.NonLatinScriptNames)
		assert.Equal(t, UKSLEntity, *record.EntityType)
		assert.NotEmpty(t, record.Addresses)
		assert.NotEmpty(t, record.AddressCountries)
		assert.Equal(t, "Pakistan", record.AddressCountries[0])
		assert.Equal(t, "12345", record.AddressPostalCodes[0])
		assert.Equal(t, "The Afghanistan (Sanctions) (EU Exit) Regulations 2020", record.Regime)
		// Vessel fields for Entity type
		assert.Equal(t, "9000000", record.IMONumber)
		assert.Equal(t, "Cargo", record.VesselType)
		assert.Equal(t, "5000", record.Tonnage)
		assert.Equal(t, "Panama", record.VesselFlag)
	} else {
		t.Fatal("AFG0001 record not found")
	}

	// Test Individual record with personal info
	if record, ok := report["AFG0006"]; ok {
		assert.Equal(t, UKSLIndividual, *record.EntityType)
		assert.Equal(t, "AKHUND MOHAMMAD HASSAN", record.Names[0])
		assert.Equal(t, "Mullah", record.NameTitle)
		assert.Equal(t, "01/01/1958", record.DOB)
		assert.Equal(t, "Afghanistan", record.Nationality)
		assert.Equal(t, "12345", record.NationalIDNumber)
		assert.Equal(t, "Mali number", record.NationalIDAdditionalInfo)
		assert.Equal(t, "A1234567", record.PassportNumber)
		assert.Equal(t, "Mali passport", record.PassportAdditionalInfo)
		assert.Equal(t, "Deputy Minister", record.Position)
		assert.Equal(t, "Male", record.Gender)
		assert.Equal(t, "Kandahar", record.TownOfBirth)
		assert.Equal(t, "Afghanistan", record.CountryOfBirth)
	} else {
		t.Fatal("AFG0006 record not found")
	}
}

type nopCloser struct {
	*strings.Reader
}

func (nopCloser) Close() error { return nil }

func TestReadUKSanctionsListFromString(t *testing.T) {
	csvData := `Report Date: 13-Feb-2026
Last Updated,Unique ID,OFSI Group ID,UN Reference Number,Name 6,Name 1,Name 2,Name 3,Name 4,Name 5,Name type,Alias strength,Title,Name non-latin script,Non-latin script type,Non-latin script language,Regime Name,Designation Type,Designation source,Sanctions Imposed,Other Information,UK Statement of Reasons,Address Line 1,Address Line 2,Address Line 3,Address Line 4,Address Line 5,Address Line 6,Address Postal Code,Address Country,Phone number,Website,Email address,Date Designated,D.O.B,Nationality(/ies),National Identifier number,National Identifier additional information,Passport number,Passport additional information,Position,Gender,Town of birth,Country of birth,Type of entity,Subsidiaries,Parent company,Business registration number (s),IMO number,Current owner/operator (s),Previous owner/operator (s),Current believed flag of ship,Previous flags,Type of ship,Tonnage of ship,Length of ship,Year Built,Hull identification number (HIN)
12/01/2022,TEST001,12703,,TEST,PERSON,ONE,,,,Primary Name,,Dr,,,,Test Regime,Individual,UN,Asset freeze,,,123 Test St,,,,,,SW1A 1AA,United Kingdom,,,,,15/03/1980,British,ID123,UK ID,PP999,UK passport,CEO,Female,London,United Kingdom,,,,,,,,,,,,,
`

	fd := nopCloser{strings.NewReader(csvData)}
	totalReport, report, err := ReadSanctionsListFile(fd)
	assert.NoError(t, err)
	assert.Len(t, totalReport, 1)

	record := report["TEST001"]
	assert.NotNil(t, record)
	assert.Equal(t, "TEST PERSON ONE", record.Names[0])
	assert.Equal(t, "Dr", record.NameTitle)
	assert.Equal(t, "15/03/1980", record.DOB)
	assert.Equal(t, "British", record.Nationality)
	assert.Equal(t, "ID123", record.NationalIDNumber)
	assert.Equal(t, "PP999", record.PassportNumber)
	assert.Equal(t, "CEO", record.Position)
	assert.Equal(t, "Female", record.Gender)
	assert.Equal(t, "SW1A 1AA", record.AddressPostalCodes[0])
	assert.Equal(t, "Test Regime", record.Regime)
}
