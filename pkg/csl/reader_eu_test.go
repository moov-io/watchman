package csl

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEU(t *testing.T) {
	euCSL, err := ReadEUFile(filepath.Join("..", "..", "test", "testdata", "eu_csl.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if euCSL == nil {
		t.Fatal("failed to parse eu_csl.csv")
	}

	testLogicalID := 13

	// TODO: add tests for failure cases
	if euCSL[testLogicalID] == nil {
		t.Fatalf("expected a record at %d and got nil", testLogicalID)
	}
	expectedFileGenerationDate := "28/10/2022"
	expectedReferenceNumber := "EU.27.28"
	expectedEntityRemark := "(UNSC RESOLUTION 1483)"
	expectedClassificationCode := "person"
	expectedPublicationURL := "http://eur-lex.europa.eu/LexUriServ/LexUriServ.do?uri=OJ:L:2003:169:0006:0023:EN:PDF"

	// Name alias
	expectedNameAliasWholeName1 := "Saddam Hussein Al-Tikriti"
	expectedNameAliasWholeName2 := "Abu Ali"
	expectedNameAliasWholeName3 := "Abou Ali"
	expectedNameAliasTitle := ""

	// Address
	expectedAddressCity := ""
	expectedAddressStreet := ""
	expectedAddressPoBox := ""
	expectedAddressZipCode := ""
	expectedAddressCountryDescription := ""

	expectedBirthDate := "1937-04-28"
	expectedBirthCity := "al-Awja, near Tikrit"
	expectedBirthCountryDescription := "IRAQ"

	assert.Greater(t, len(euCSL), 0)
	assert.Greater(t, len(euCSL[testLogicalID]), 0)
	assert.NotNil(t, euCSL[testLogicalID][0].Entity)
	assert.NotNil(t, euCSL[testLogicalID][0].NameAlias)
	assert.NotNil(t, euCSL[testLogicalID][0].Address)
	assert.NotNil(t, euCSL[testLogicalID][0].BirthDate)
	assert.NotNil(t, euCSL[testLogicalID][0].Identification)

	assert.Equal(t, euCSL[testLogicalID][0].FileGenerationDate, expectedFileGenerationDate)

	// Entity
	assert.Equal(t, euCSL[testLogicalID][0].Entity.ReferenceNumber, expectedReferenceNumber)
	assert.Equal(t, euCSL[testLogicalID][0].Entity.Remark, expectedEntityRemark)
	assert.Equal(t, euCSL[testLogicalID][0].Entity.SubjectType.ClassificationCode, expectedClassificationCode)
	assert.Equal(t, euCSL[testLogicalID][0].Entity.Regulation.PublicationURL, expectedPublicationURL)

	// Name Alias
	assert.Equal(t, euCSL[testLogicalID][0].NameAlias.WholeName, expectedNameAliasWholeName1)
	assert.Equal(t, euCSL[testLogicalID][1].NameAlias.WholeName, expectedNameAliasWholeName2)
	assert.Equal(t, euCSL[testLogicalID][2].NameAlias.WholeName, expectedNameAliasWholeName3)
	assert.Equal(t, euCSL[testLogicalID][0].NameAlias.Title, expectedNameAliasTitle)

	// Address
	assert.Equal(t, euCSL[testLogicalID][0].Address.City, expectedAddressCity)
	assert.Equal(t, euCSL[testLogicalID][0].Address.Street, expectedAddressStreet)
	assert.Equal(t, euCSL[testLogicalID][0].Address.PoBox, expectedAddressPoBox)
	assert.Equal(t, euCSL[testLogicalID][0].Address.ZipCode, expectedAddressZipCode)
	assert.Equal(t, euCSL[testLogicalID][0].Address.CountryDescription, expectedAddressCountryDescription)

	// BirthDate
	assert.Equal(t, euCSL[testLogicalID][3].BirthDate.BirthDate, expectedBirthDate)
	assert.Equal(t, euCSL[testLogicalID][3].BirthDate.City, expectedBirthCity)
	assert.Equal(t, euCSL[testLogicalID][3].BirthDate.CountryDescription, expectedBirthCountryDescription)

	// Identification
	assert.Equal(t, euCSL[testLogicalID][0].Identification.ValidFrom, "")
	assert.Equal(t, euCSL[testLogicalID][0].Identification.ValidTo, "")
}
