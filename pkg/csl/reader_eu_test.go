package csl

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEU(t *testing.T) {
	euCSL, euCSLMap, err := ReadEUFile(filepath.Join("..", "..", "test", "testdata", "eu_csl.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if euCSL == nil || euCSLMap == nil {
		t.Fatal("failed to parse eu_csl.csv")
	}

	if len(euCSL) == 0 {
		t.Fatal("failed to convert map to sheet")
	}

	testLogicalID := 13

	if euCSLMap[testLogicalID] == nil {
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
	expectedWholeNames := []string{expectedNameAliasWholeName1, expectedNameAliasWholeName2, expectedNameAliasWholeName3}

	// Address
	// No address found for this record

	expectedBirthDate := "1937-04-28"
	expectedBirthCity := "al-Awja, near Tikrit"
	expectedBirthCountryDescription := "IRAQ"

	assert.Equal(t, euCSLMap[testLogicalID].FileGenerationDate, expectedFileGenerationDate)

	// Entity
	assert.Equal(t, expectedReferenceNumber, euCSLMap[testLogicalID].EntityReferenceNumber)
	assert.Equal(t, expectedEntityRemark, euCSLMap[testLogicalID].EntityRemark)
	assert.Equal(t, expectedClassificationCode, euCSLMap[testLogicalID].EntitySubjectType)
	assert.Equal(t, expectedPublicationURL, euCSLMap[testLogicalID].EntityPublicationURL)

	// Name Alias
	assert.Equal(t, len(expectedWholeNames), len(euCSLMap[testLogicalID].NameAliasWholeNames))
	assert.Equal(t, expectedNameAliasWholeName1, euCSLMap[testLogicalID].NameAliasWholeNames[0])
	assert.Equal(t, expectedNameAliasWholeName2, euCSLMap[testLogicalID].NameAliasWholeNames[1])
	assert.Equal(t, expectedNameAliasWholeName3, euCSLMap[testLogicalID].NameAliasWholeNames[2])

	// BirthDate
	assert.Equal(t, expectedBirthDate, euCSLMap[testLogicalID].BirthDates[0])
	assert.Equal(t, expectedBirthCity, euCSLMap[testLogicalID].BirthCities[0])
	assert.Equal(t, expectedBirthCountryDescription, euCSLMap[testLogicalID].BirthCountries[0])
}
