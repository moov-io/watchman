// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/watchman/pkg/search"

	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "..", "test", "testdata", "UN_consolidated.xml"))
	require.NoError(t, err)
	t.Cleanup(func() { fd.Close() })

	var entities []search.Entity[search.Value]
	err = NewReader(fd).Read(
		func(p UNIndividual) {
			entities = append(entities, p.ToEntity())

		},
		func(e UNEntity) {
			entities = append(entities, e.ToEntity())
		},
	)
	require.NoError(t, err)

	require.Len(t, entities, 1003)

	// found := entities[563]
	// require.Equal(t, "ABOUD ROGO MOHAMMED", found.Name)
	// require.Equal(t, search.EntityPerson, found.Type)
	// require.Equal(t, search.SourceUNCSL, found.Source)
	// require.Equal(t, "UN-6908043", found.SourceID)

	// t.Logf("%#v", found.Person)

	// require.NotNil(t, found.Person)
	// require.Nil(t, found.Business)
	// require.Nil(t, found.Organization)
	// require.Nil(t, found.Aircraft)
	// require.Nil(t, found.Vessel)

	for _, found := range entities {
		var consider int

		if len(found.Contact.EmailAddresses) > 0 || len(found.Contact.PhoneNumbers) > 0 {
			consider++
		}
		if len(found.Addresses) > 0 {
			consider++
		}

		if found.SourceID == "UN-6908027" {
			t.Logf("%#v", found)
		}

		if consider > 1 {
			t.Logf("%#v", found)
			break
		}
	}

	// TODO(adam): verify phone numbers are extracted
	// <INDIVIDUAL>
	//   <DATAID>6909290</DATAID>
	//   <VERSIONNUM>1</VERSIONNUM>
	//   <FIRST_NAME>ABUBAKAR</FIRST_NAME>
	//   <SECOND_NAME>SWALLEH</SECOND_NAME>
	//   <UN_LIST_TYPE>Al-Qaida</UN_LIST_TYPE>
	//   <REFERENCE_NUMBER>QDi.436</REFERENCE_NUMBER>
	//   <LISTED_ON>2025-06-16</LISTED_ON>
	//   <COMMENTS1>Abubakar Swalleh provides financial, material, or technological support for, or financial or other services to, or in support of, ISIL (listed as Al-Qaida in Iraq (QDe.115). He acted, since 2018, as an ISIL facilitator who provides financial and logistic support including recruitment for ISIL in East and Southern Africa. Phone number: +963936016952. Gender: Male  INTERPOL-UN Security Council Special Notice web link:https://www.interpol.int/en/How-we-work/Notices/View-UN-Notices-Individuals</COMMENTS1>
	// ...
	// <INDIVIDUAL_ADDRESS>
	//   <CITY>Luzira Prison, Luzira, Kampala</CITY>
	//   <COUNTRY>Uganda</COUNTRY>
	// </INDIVIDUAL_ADDRESS>
	// <INDIVIDUAL_DATE_OF_BIRTH>
	//   <TYPE_OF_DATE>EXACT</TYPE_OF_DATE>
	//   <DATE>1992-01-13</DATE>
	// </INDIVIDUAL_DATE_OF_BIRTH>

	// TODO(adam): verify emails and addresses are extracted
	// <ENTITY>
	//       <DATAID>6908027</DATAID>
	//       <VERSIONNUM>1</VERSIONNUM>
	//       <FIRST_NAME>FORCES DEMOCRATIQUES DE LIBERATION DU RWANDA (FDLR)</FIRST_NAME>
	//       <UN_LIST_TYPE>DRC</UN_LIST_TYPE>
	//       <REFERENCE_NUMBER>CDe.005</REFERENCE_NUMBER>
	//       <LISTED_ON>2012-12-31</LISTED_ON>
	//       <COMMENTS1>Email: Fdlr@fmx.de; fldrrse@yahoo.fr; fdlr@gmx.net; fdlrsrt@gmail.com;
	//          humura2020@gmail.com. INTERPOL-UN Security Council Special Notice web link:https://www.interpol.int/en/How-we-work/Notices/View-UN-Notices-Individuals</COMMENTS1>
	// ..
	//   <ENTITY_ADDRESS>
	//     <CITY>North Kivu</CITY>
	//     <COUNTRY>Democratic Republic of the Congo</COUNTRY>
	//   </ENTITY_ADDRESS>
	//   <ENTITY_ADDRESS>
	//     <CITY>South Kivu</CITY>
	//     <COUNTRY>Democratic Republic of the Congo</COUNTRY>
	//   </ENTITY_ADDRESS>

	// Contact:search.ContactInfo{EmailAddresses:[]string(nil), PhoneNumbers:[]string(nil), FaxNumbers:[]string(nil), Websites:[]string(nil)},
	// 	Addresses:[]search.Address{search.Address{Line1:"", Line2:"", City:"", PostalCode:"", State:"", Country:"", Latitude:0, Longitude:0}},
	// 	CryptoAddresses:[]search.CryptoAddress(nil), Affiliations:[]search.Affiliation(nil), SanctionsInfo:(*search.SanctionsInfo)(nil), HistoricalInfo:[]search.HistoricalInfo(nil),

}

func TestReader_Read_success(t *testing.T) {
	// Build a minimal UN Consolidated List XML containing one individual and one entity
	xml := `<?xml version="1.0"?>
<CONSOLIDATED_LIST>
  <INDIVIDUALS>
    <INDIVIDUAL>
      <DATAID>id1</DATAID>
      <FIRST_NAME>Jane</FIRST_NAME>
      <SECOND_NAME>Doe</SECOND_NAME>
      <THIRD_NAME></THIRD_NAME>
      <FOURTH_NAME></FOURTH_NAME>
    </INDIVIDUAL>
  </INDIVIDUALS>
  <ENTITIES>
    <ENTITY>
      <DATAID>id2</DATAID>
      <FIRST_NAME>Acme Corp</FIRST_NAME>
    </ENTITY>
  </ENTITIES>
</CONSOLIDATED_LIST>`

	reader := NewReader(strings.NewReader(xml))

	var inds []UNIndividual
	var ents []UNEntity

	err := reader.Read(
		func(p UNIndividual) { inds = append(inds, p) },
		func(e UNEntity) { ents = append(ents, e) },
	)
	if err != nil {
		t.Fatalf("unexpected error reading xml: %v", err)
	}

	if len(inds) != 1 {
		t.Fatalf("expected 1 individual, got %d", len(inds))
	}
	if inds[0].DataID != "id1" || inds[0].FirstName != "Jane" {
		t.Errorf("individual fields incorrect: %+v", inds[0])
	}

	if len(ents) != 1 {
		t.Fatalf("expected 1 entity, got %d", len(ents))
	}
	if ents[0].DataID != "id2" || ents[0].FirstName != "Acme Corp" {
		t.Errorf("entity fields incorrect: %+v", ents[0])
	}
}

func TestReader_Read_malformed(t *testing.T) {
	// truncated/invalid XML should return an error on Decode
	xml := `<CONSOLIDATED_LIST><INDIVIDUALS><INDIVIDUAL><DATAID>id1` +
		"</CONSOLIDATED_LIST>"

	reader := NewReader(strings.NewReader(xml))

	err := reader.Read(func(p UNIndividual) {}, func(e UNEntity) {})
	if err == nil {
		t.Fatal("expected error for malformed xml")
	}
}
