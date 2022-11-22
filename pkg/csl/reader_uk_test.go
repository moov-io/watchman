package csl

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUK(t *testing.T) {
	ukCSL, ukCSLMap, err := ReadUKFile(filepath.Join("..", "..", "test", "testdata", "ConList.csv"))
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
	assert.Greater(t, len(testRow.AddressesTwo), 0)
	assert.Nil(t, testRow.AddressesThree)
	assert.Nil(t, testRow.AddressesFour)
	assert.Greater(t, len(testRow.AddressesFive), 0)
	assert.Greater(t, len(testRow.AddressesSix), 0)
	assert.Nil(t, testRow.PostalCodes)
	assert.Greater(t, len(testRow.Countries), 0)
	assert.Greater(t, len(testRow.OtherInfos), 0)

	// assertions
	assert.Equal(t, expectedName, testRow.Names[0])
	assert.Equal(t, expectedAddressOne, testRow.Addresses[0])
	assert.Equal(t, expectedAddressFive, testRow.AddressesFive[0])
	assert.Equal(t, expectedAddressSix, testRow.AddressesSix[0])
	assert.Equal(t, expectedCountry, testRow.Countries[0])
	assert.Equal(t, expectedOtherInfo, testRow.OtherInfos[0])
	assert.Equal(t, expectedGroupType, testRow.GroupTypes[0])
	assert.Equal(t, expectedListedDate, testRow.ListedDates[0])
	assert.Equal(t, expectedLastUpdatedDate, testRow.LastUpdates[0])
	assert.Equal(t, expectedUKSancListDate, testRow.SanctionListDates[0])
	assert.Equal(t, expectedGroupID, testRow.GroupID)
}