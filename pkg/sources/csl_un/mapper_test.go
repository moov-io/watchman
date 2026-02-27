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
			{City: "New York", Country: "US", Note: "Phone:+1-222-3333 Email:john@example.com"},
		},
		Comments:      "In comments you may reach him at john.doe@domain.org or +44 20 7946 0991 Gender: Male website:https://john.example.com",
		BirthDates:    []BirthDate{{Type: "EXACT", Date: "1980-05-20"}},
		BirthPlaces:   []BirthPlace{{City: "Springfield", State: "IL", Country: "US"}},
		Nationalities: []Value{{Text: "USA"}},
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

	// contact info should be pulled from address notes and comments
	expEmails := []string{"john@example.com", "john.doe@domain.org"}
	if len(ent.Contact.EmailAddresses) != len(expEmails) {
		t.Fatalf("expected %d emails, got %d", len(expEmails), len(ent.Contact.EmailAddresses))
	}
	for i, v := range expEmails {
		if ent.Contact.EmailAddresses[i] != v {
			t.Errorf("email[%d] expected %q got %q", i, v, ent.Contact.EmailAddresses[i])
		}
	}

	expPhones := []string{"+1-222-3333", "+44 20 7946 0991"}
	if len(ent.Contact.PhoneNumbers) != len(expPhones) {
		t.Fatalf("expected %d phones, got %d", len(expPhones), len(ent.Contact.PhoneNumbers))
	}
	for i, v := range expPhones {
		if ent.Contact.PhoneNumbers[i] != v {
			t.Errorf("phone[%d] expected %q got %q", i, v, ent.Contact.PhoneNumbers[i])
		}
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

	// birthdate
	if ent.Person.BirthDate == nil {
		t.Fatalf("expected birthdate to be set")
	}
	if ent.Person.BirthDate.Year() != 1980 || ent.Person.BirthDate.Month() != 5 || ent.Person.BirthDate.Day() != 20 {
		t.Errorf("birthdate mismatch: %v", ent.Person.BirthDate)
	}

	// birthplace
	if ent.Person.PlaceOfBirth != "Springfield, IL, US" {
		t.Errorf("expected birthplace %q got %q", "Springfield, IL, US", ent.Person.PlaceOfBirth)
	}

	// nationality stored in Titles (mapper appends nationalities there)
	found := false
	for _, t := range ent.Person.Titles {
		if t == "USA" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected nationality USA in Person.Titles, got %v", ent.Person.Titles)
	}

	// gender parsed from comments
	if ent.Person.Gender != search.Gender("Male") {
		t.Errorf("expected gender Male, got %v", ent.Person.Gender)
	}

	// website extracted
	expWeb := []string{"https://john.example.com"}
	if len(ent.Contact.Websites) != len(expWeb) {
		t.Fatalf("expected %d websites, got %d", len(expWeb), len(ent.Contact.Websites))
	}
	if ent.Contact.Websites[0] != expWeb[0] {
		t.Errorf("website expected %q got %q", expWeb[0], ent.Contact.Websites[0])
	}
}

func TestUNEntity_ToEntity(t *testing.T) {
	e := UNEntity{
		DataID:    "456",
		FirstName: "BigCorp",
		Aliases:   []Alias{{Name: "BC"}},
		Addresses: []Address{{City: "London", Country: "GB", Note: "email:info@bigcorp.com phone:+442071838750 website:www.bigcorp.com"}},
		Comments:  "Reach us at biz@contact.com or call +1234567890",
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

	// business comments contact info
	expEmailsBiz := []string{"info@bigcorp.com", "biz@contact.com"}
	if len(ent.Contact.EmailAddresses) != len(expEmailsBiz) {
		t.Fatalf("expected %d business emails, got %d", len(expEmailsBiz), len(ent.Contact.EmailAddresses))
	}
	for i, v := range expEmailsBiz {
		if ent.Contact.EmailAddresses[i] != v {
			t.Errorf("business email[%d] expected %q got %q", i, v, ent.Contact.EmailAddresses[i])
		}
	}

	expPhonesBiz := []string{"+442071838750", "+1234567890"}
	if len(ent.Contact.PhoneNumbers) != len(expPhonesBiz) {
		t.Fatalf("expected %d business phones, got %d", len(expPhonesBiz), len(ent.Contact.PhoneNumbers))
	}
	for i, v := range expPhonesBiz {
		if ent.Contact.PhoneNumbers[i] != v {
			t.Errorf("business phone[%d] expected %q got %q", i, v, ent.Contact.PhoneNumbers[i])
		}
	}

	// website from address note
	expBizWeb := []string{"www.bigcorp.com"}
	if len(ent.Contact.Websites) != len(expBizWeb) {
		t.Fatalf("expected %d business websites, got %d", len(expBizWeb), len(ent.Contact.Websites))
	}
	if ent.Contact.Websites[0] != expBizWeb[0] {
		t.Errorf("business website expected %q got %q", expBizWeb[0], ent.Contact.Websites[0])
	}
}
