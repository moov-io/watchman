// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
)

func TestUNIndividual_ToEntity(t *testing.T) {
	p := UNIndividual{
		DataID:     "123",
		FirstName:  "John",
		SecondName: "Q",
		ThirdName:  "Public",
		FourthName: "",
		Aliases: []Alias{
			{Name: "Johnny"},
			{Name: ""}, // empty alias should be ignored
		},
		Addresses: []Address{
			{City: "New York", Country: "US"},
		},
	}

	ent := p.ToEntity()

	expName := "John Q Public"
	if ent.Name != expName {
		t.Errorf("expected Name %q, got %q", expName, ent.Name)
	}

	expID := "UN-123"
	if ent.SourceID != expID {
		t.Errorf("expected SourceID %q, got %q", expID, ent.SourceID)
	}

	if ent.Type != search.EntityPerson {
		t.Errorf("expected EntityPerson, got %v", ent.Type)
	}

	if ent.Person == nil {
		t.Fatal("Person pointer should not be nil")
	}

	if ent.Person.Name != expName {
		t.Errorf("person name mismatch: %q", ent.Person.Name)
	}

	expAlts := []string{"Johnny"}
	if len(ent.Person.AltNames) != len(expAlts) {
		t.Fatalf("expected %d alt names, got %d", len(expAlts), len(ent.Person.AltNames))
	}
	for i, v := range expAlts {
		if ent.Person.AltNames[i] != v {
			t.Errorf("alt name[%d] expected %q got %q", i, v, ent.Person.AltNames[i])
		}
	}

	if len(ent.Addresses) != 1 {
		t.Fatalf("expected 1 address, got %d", len(ent.Addresses))
	}
	if ent.Addresses[0].City != "New York" || ent.Addresses[0].Country != "US" {
		t.Errorf("address mismatch: %+v", ent.Addresses[0])
	}
}

func TestUNEntity_ToEntity(t *testing.T) {
	e := UNEntity{
		DataID:    "456",
		FirstName: "BigCorp",
		Aliases:   []Alias{{Name: "BC"}},
	}

	ent := e.ToEntity()

	expName := "BigCorp"
	if ent.Name != expName {
		t.Errorf("expected Name %q, got %q", expName, ent.Name)
	}

	expID := "UN-456"
	if ent.SourceID != expID {
		t.Errorf("expected SourceID %q, got %q", expID, ent.SourceID)
	}

	if ent.Type != search.EntityBusiness {
		t.Errorf("expected EntityBusiness, got %v", ent.Type)
	}

	if ent.Business == nil {
		t.Fatal("Business pointer should not be nil")
	}

	if ent.Business.Name != expName {
		t.Errorf("business name mismatch: %q", ent.Business.Name)
	}

	expAlts := []string{"BC"}
	if len(ent.Business.AltNames) != len(expAlts) {
		t.Fatalf("expected %d alt names, got %d", len(expAlts), len(ent.Business.AltNames))
	}
	for i, v := range expAlts {
		if ent.Business.AltNames[i] != v {
			t.Errorf("alt name[%d] expected %q got %q", i, v, ent.Business.AltNames[i])
		}
	}
}
