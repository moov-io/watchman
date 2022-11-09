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
	// No address found for this record

	expectedBirthDate := "1937-04-28"
	expectedBirthCity := "al-Awja, near Tikrit"
	expectedBirthCountryDescription := "IRAQ"

	assert.Greater(t, len(euCSL), 0)
	assert.NotNil(t, euCSL[testLogicalID].Entity)
	assert.NotNil(t, euCSL[testLogicalID].NameAliases)
	assert.NotNil(t, euCSL[testLogicalID].Addresses)
	assert.NotNil(t, euCSL[testLogicalID].BirthDates)
	assert.NotNil(t, euCSL[testLogicalID].Identifications)

	assert.Equal(t, euCSL[testLogicalID].FileGenerationDate, expectedFileGenerationDate)

	// Entity
	assert.Equal(t, euCSL[testLogicalID].Entity.ReferenceNumber, expectedReferenceNumber)
	assert.Equal(t, euCSL[testLogicalID].Entity.Remark, expectedEntityRemark)
	assert.Equal(t, euCSL[testLogicalID].Entity.SubjectType.ClassificationCode, expectedClassificationCode)
	assert.Equal(t, euCSL[testLogicalID].Entity.Regulation.PublicationURL, expectedPublicationURL)

	// Name Alias
	assert.Equal(t, euCSL[testLogicalID].NameAliases[0].WholeName, expectedNameAliasWholeName1)
	assert.Equal(t, euCSL[testLogicalID].NameAliases[1].WholeName, expectedNameAliasWholeName2)
	assert.Equal(t, euCSL[testLogicalID].NameAliases[2].WholeName, expectedNameAliasWholeName3)
	assert.Equal(t, euCSL[testLogicalID].NameAliases[0].Title, expectedNameAliasTitle)

	// Address
	assert.Len(t, euCSL[testLogicalID].Addresses, 0)

	// BirthDate
	assert.Equal(t, euCSL[testLogicalID].BirthDates[0].BirthDate, expectedBirthDate)
	assert.Equal(t, euCSL[testLogicalID].BirthDates[0].City, expectedBirthCity)
	assert.Equal(t, euCSL[testLogicalID].BirthDates[0].CountryDescription, expectedBirthCountryDescription)

	// Identification
	assert.Len(t, euCSL[testLogicalID].Identifications, 0)
}
