// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/watchman/pkg/csl_us/gen/ENHANCED_XML"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestGetPrimaryNameFromXML(t *testing.T) {
	namesXML := `
      <names>
        <name id="17178">
          <isPrimary>true</isPrimary>
          <isLowQuality>false</isLowQuality>
          <translations>
            <translation id="1979">
              <isPrimary>true</isPrimary>
              <script refId="20122">Latin</script>
              <formattedFirstName>Mohammed</formattedFirstName>
              <formattedLastName>ABU TEIR</formattedLastName>
              <formattedFullName>ABU TEIR, Mohammed</formattedFullName>
              <nameParts>
                <namePart id="38178">
                  <type refId="1520">Last Name</type>
                  <value>ABU TEIR</value>
                </namePart>
                <namePart id="38179">
                  <type refId="1521">First Name</type>
                  <value>Mohammed</value>
                </namePart>
              </nameParts>
            </translation>
          </translations>
        </name>
      </names>`

	var names ENHANCED_XML.EntityNames
	if err := xml.Unmarshal([]byte(namesXML), &names); err != nil {
		t.Fatalf("failed to unmarshal names XML: %v", err)
	}

	got := getPrimaryName(names, "individual")
	require.Equal(t, "Mohammed ABU TEIR", got)
}

func TestGetAllNamesFromXML(t *testing.T) {
	namesXML := `
      <names>
        <name id="17178">
          <isPrimary>true</isPrimary>
          <isLowQuality>false</isLowQuality>
          <translations>
            <translation id="1979">
              <isPrimary>true</isPrimary>
              <script refId="20122">Latin</script>
              <formattedFirstName>Mohammed</formattedFirstName>
              <formattedLastName>ABU TEIR</formattedLastName>
              <formattedFullName>ABU TEIR, Mohammed</formattedFullName>
            </translation>
          </translations>
        </name>
        <name id="9152">
          <isPrimary>false</isPrimary>
          <aliasType refId="1400">A.K.A.</aliasType>
          <isLowQuality>false</isLowQuality>
          <translations>
            <translation id="10144">
              <isPrimary>true</isPrimary>
              <script refId="20122">Latin</script>
              <formattedFirstName>Mohammed Mahmud</formattedFirstName>
              <formattedLastName>ABU TAIR</formattedLastName>
              <formattedFullName>ABU TAIR, Mohammed Mahmud</formattedFullName>
            </translation>
          </translations>
        </name>
        <name id="9153">
          <isPrimary>false</isPrimary>
          <aliasType refId="1400">A.K.A.</aliasType>
          <isLowQuality>false</isLowQuality>
          <translations>
            <translation id="10145">
              <isPrimary>true</isPrimary>
              <script refId="20122">Latin</script>
              <formattedFirstName>Mohammad Mahmoud</formattedFirstName>
              <formattedLastName>ABOU TAYR</formattedLastName>
              <formattedFullName>ABOU TAYR, Mohammad Mahmoud</formattedFullName>
            </translation>
          </translations>
        </name>
      </names>`

	var names ENHANCED_XML.EntityNames
	if err := xml.Unmarshal([]byte(namesXML), &names); err != nil {
		t.Fatalf("failed to unmarshal names XML: %v", err)
	}

	got := getAllNames(names, "individual")
	want := []string{
		"Mohammed Mahmud ABU TAIR",
		"Mohammad Mahmoud ABOU TAYR",
	}

	if len(got) != len(want) {
		t.Errorf("getAllNames() returned %d names, want %d", len(got), len(want))
		return
	}

	for i, name := range got {
		if name != want[i] {
			t.Errorf("getAllNames()[%d] = %v, want %v", i, name, want[i])
		}
	}
}

func TestMapFeaturesFromXML(t *testing.T) {
	featuresXML := `
      <features>
        <feature id="4698">
          <type featureTypeId="8">Birthdate</type>
          <versionId>1612</versionId>
          <value>1951</value>
          <valueDate id="757">
            <fromDateBegin>1951-01-01</fromDateBegin>
            <fromDateEnd>1951-01-01</fromDateEnd>
            <toDateBegin>1951-12-31</toDateBegin>
            <toDateEnd>1951-12-31</toDateEnd>
            <isApproximate>false</isApproximate>
            <isDateRange>false</isDateRange>
          </valueDate>
          <isPrimary>true</isPrimary>
        </feature>
        <feature id="4699">
          <type featureTypeId="9">Place of Birth</type>
          <versionId>1613</versionId>
          <value>Umm Tuba</value>
          <isPrimary>true</isPrimary>
        </feature>
      </features>`

	var features ENHANCED_XML.EntityFeatures
	if err := xml.Unmarshal([]byte(featuresXML), &features); err != nil {
		t.Fatalf("failed to unmarshal features XML: %v", err)
	}

	person := &search.Person{}
	mapPersonFeatures(&features, person)

	// Check birth date year
	if person.BirthDate == nil {
		t.Fatal("BirthDate is nil, want non-nil")
	}
	if person.BirthDate.Year() != 1951 {
		t.Errorf("BirthDate.Year() = %v, want 1951", person.BirthDate.Year())
	}

	// // Check birth place
	// wantPlace := "Umm Tuba"
	// if person.PlaceOfBirth != wantPlace {
	// 	t.Errorf("PlaceOfBirth = %v, want %v", person.PlaceOfBirth, wantPlace)
	// }
}

func TestMapSanctionsProgramsFromXML(t *testing.T) {
	programsXML := `
      <sanctionsPrograms>
        <sanctionsProgram refId="91055" id="6215">NS-PLC</sanctionsProgram>
      </sanctionsPrograms>`

	var programs ENHANCED_XML.EntitySanctionsPrograms
	if err := xml.Unmarshal([]byte(programsXML), &programs); err != nil {
		t.Fatalf("failed to unmarshal programs XML: %v", err)
	}

	typesXML := `
      <sanctionsTypes>
        <sanctionsType refId="1706" id="40">Reject</sanctionsType>
      </sanctionsTypes>`

	var types ENHANCED_XML.EntitySanctionsTypes
	if err := xml.Unmarshal([]byte(typesXML), &types); err != nil {
		t.Fatalf("failed to unmarshal types XML: %v", err)
	}

	entity := ENHANCED_XML.EntitiesEntity{
		SanctionsPrograms: programs,
		SanctionsTypes:    types,
	}

	got := mapSanctionsInfo(entity)
	require.Len(t, got.Programs, 1)
	require.Equal(t, "NS-PLC", got.Programs[0])
}

func TestMapIDTypeFromXML(t *testing.T) {
	idDocXML := `
      <identityDocuments>
        <identityDocument id="123">
          <type refId="1571">Passport</type>
          <documentNumber>ABC123</documentNumber>
          <issuingCountry refId="11065">China</issuingCountry>
        </identityDocument>
      </identityDocuments>`

	var docs ENHANCED_XML.EntityIdentityDocuments
	if err := xml.Unmarshal([]byte(idDocXML), &docs); err != nil {
		t.Fatalf("failed to unmarshal identity documents XML: %v", err)
	}

	ids := mapGovernmentIDs(&docs)
	require.Len(t, ids, 1)

	want := search.GovernmentID{
		Type:       search.GovernmentIDPassport,
		Country:    "China",
		Identifier: "ABC123",
	}
	require.Equal(t, want, ids[0])
}
