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

	found := entities[563]
	require.Equal(t, "ABOUD ROGO MOHAMMED", found.Name)
	require.Equal(t, search.EntityPerson, found.Type)
	require.Equal(t, search.SourceUNCSL, found.Source)
	require.Equal(t, "UN-6908043", found.SourceID)

	// Verify phone numbers are extracted for individual DATAID 6909290
	var individual *search.Entity[search.Value]
	var entity *search.Entity[search.Value]
	for i := range entities {
		if entities[i].SourceID == "UN-6909290" {
			individual = &entities[i]
		}
		if entities[i].SourceID == "UN-6908027" {
			entity = &entities[i]
		}
	}

	require.NotNil(t, individual)
	require.Equal(t, "ABUBAKAR SWALLEH", individual.Name)
	require.Equal(t, search.EntityPerson, individual.Type)
	require.Equal(t, search.SourceUNCSL, individual.Source)
	require.NotNil(t, individual.Person)
	require.Nil(t, individual.Business)
	require.Len(t, individual.Contact.PhoneNumbers, 1)
	require.Equal(t, "+963936016952", individual.Contact.PhoneNumbers[0])
	require.Len(t, individual.Addresses, 1)
	require.Equal(t, "Luzira Prison, Luzira, Kampala", individual.Addresses[0].City)
	require.Equal(t, "Uganda", individual.Addresses[0].Country)

	// Verify emails and addresses are extracted for entity DATAID 6908027
	require.NotNil(t, entity)
	require.Equal(t, "FORCES DEMOCRATIQUES DE LIBERATION DU RWANDA (FDLR)", entity.Name)
	require.Equal(t, search.EntityBusiness, entity.Type)
	require.Equal(t, search.SourceUNCSL, entity.Source)
	require.NotNil(t, entity.Business)
	require.Nil(t, entity.Person)
	expectedEmails := []string{"Fdlr@fmx.de", "fldrrse@yahoo.fr", "fdlr@gmx.net", "fdlrsrt@gmail"}
	require.Len(t, entity.Contact.EmailAddresses, len(expectedEmails))
	for i, email := range expectedEmails {
		require.Equal(t, email, entity.Contact.EmailAddresses[i])
	}
	require.Len(t, entity.Addresses, 2)
	require.Equal(t, "North Kivu", entity.Addresses[0].City)
	require.Equal(t, "Democratic Republic of the Congo", entity.Addresses[0].Country)
	require.Equal(t, "South Kivu", entity.Addresses[1].City)
	require.Equal(t, "Democratic Republic of the Congo", entity.Addresses[1].Country)
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
