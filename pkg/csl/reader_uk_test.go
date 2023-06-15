package csl

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUKCSL(t *testing.T) {
	ukCSL, ukCSLMap, err := ReadUKCSLFile(filepath.Join("..", "..", "test", "testdata", "ConList.csv"))
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
	t.Setenv("WITH_UK_SANCTIONS_LIST", "false")
	// test we don't err on parsing the content
	totalReport, report, err := ReadUKSanctionsListFile("../../test/testdata/UK_Sanctions_List.ods")
	assert.NoError(t, err)

	// test that we get something more than an empty sanctions list record
	assert.NotEmpty(t, totalReport)
	assert.NotEmpty(t, report)
	assert.GreaterOrEqual(t, len(totalReport), 3728)

	if record, ok := report["AFG0001"]; ok {
		assert.Equal(t, "12/01/2022", record.LastUpdated)
		assert.Equal(t, "12703", record.OFSIGroupID)
		assert.Equal(t, "TAe.010", record.UNReferenceNumber)
		assert.Equal(t, "HAJI KHAIRULLAH HAJI SATTAR MONEY EXCHANGE", record.Names[0])
		assert.Len(t, record.Names, 9)
		assert.Equal(t, "Primary Name", record.NameTitle)
		assert.NotEmpty(t, record.NonLatinScriptNames)
		assert.Equal(t, UKSLEntity, *record.EntityType)
		assert.NotEmpty(t, record.Addresses)
		assert.NotEmpty(t, record.StateLocalities)
		assert.NotEmpty(t, record.AddressCountries)
		assert.Empty(t, record.CountryOfBirth)
	} else {
		t.Fatal("record not found")
	}
}
