// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadEU(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", "eu_csl.csv"))
	if err != nil {
		t.Error(err)
	}
	euCSL, euCSLMap, err := ParseEU(fd)
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
	assert.Equal(t, expectedClassificationCode, euCSLMap[testLogicalID].EntitySubjectTypeCode)
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

func TestParseEU_EmptyFile(t *testing.T) {
	// Test with empty reader - should return error
	_, _, err := ParseEU(nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty or missing")
}

func TestParseEU_ShortRecords(t *testing.T) {
	// CSV with header and records with varying lengths
	// Record needs at least 2 columns (minRecordColumns) to be processed
	csvData := "fileGenerationDate;Entity_LogicalId;Entity_EU_ReferenceNumber\n" +
		"2022-01-01;123;REF123\n" +
		"2022-01-01;456;REF456\n"

	reader := io.NopCloser(strings.NewReader(csvData))
	records, recordMap, err := ParseEU(reader)

	// Should not panic, should handle gracefully
	require.NoError(t, err)

	// Both records should be parsed
	require.Equal(t, 2, len(records), "expected 2 records")
	require.NotNil(t, recordMap[123], "expected record 123")
	require.NotNil(t, recordMap[456], "expected record 456")
	assert.Equal(t, "REF123", recordMap[123].EntityReferenceNumber)
	assert.Equal(t, "REF456", recordMap[456].EntityReferenceNumber)
}

func TestParseEU_OnlyHeader(t *testing.T) {
	// CSV with only header row
	csvData := `fileGenerationDate;Entity_LogicalId;Entity_EU_ReferenceNumber
`
	reader := io.NopCloser(strings.NewReader(csvData))
	records, _, err := ParseEU(reader)

	require.NoError(t, err)
	assert.Empty(t, records)
}

func TestUnmarshalRecord_ShortRecord(t *testing.T) {
	// Test unmarshalRecord with a very short record - should not panic
	record := &CSLRecord{}

	// Empty record
	unmarshalRecord([]string{}, record)
	assert.Equal(t, 0, record.EntityLogicalID)

	// Single element record
	unmarshalRecord([]string{"2022-01-01"}, record)
	assert.Equal(t, 0, record.EntityLogicalID)

	// Record with just ID
	unmarshalRecord([]string{"2022-01-01", "123"}, record)
	assert.Equal(t, 123, record.EntityLogicalID)
	assert.Equal(t, "2022-01-01", record.FileGenerationDate)
}
