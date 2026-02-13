// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

import (
	"strings"

	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/pkg/search"
)

// ConvertSanctionsListData converts all UK Sanctions List records to search entities.
// This is the main entry point called by the download pipeline.
func ConvertSanctionsListData(records []SanctionsListRecord) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, record := range records {
		entity := ToEntity(record)
		out = append(out, entity)
	}
	return out
}

// ToEntity converts a single UK SanctionsListRecord to a search.Entity.
// Handles Individual (Person), Entity (Business), and Ship (Vessel) types.
func ToEntity(record SanctionsListRecord) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Source:     search.SourceUKCSL,
		SourceID:   record.UniqueID,
		SourceData: record,
	}

	// Determine entity type and map type-specific data
	if record.EntityType != nil {
		switch *record.EntityType {
		case UKSLIndividual:
			entity.Type = search.EntityPerson
			entity.Person = mapPerson(record)
			entity.Name = entity.Person.Name

		case UKSLEntity:
			entity.Type = search.EntityBusiness
			entity.Business = mapBusiness(record)
			entity.Name = entity.Business.Name

		case UKSLShip:
			entity.Type = search.EntityVessel
			entity.Vessel = mapVessel(record)
			entity.Name = entity.Vessel.Name

		case Undefined:
			entity.Type = search.EntityUnknown
		}
	}

	// Map common fields
	entity.Addresses = mapAddresses(record)
	entity.SanctionsInfo = mapSanctionsInfo(record)

	return entity.Normalize()
}

// collectAltNames extracts alternate names from a record.
// It takes all names except the first (primary) one, plus any non-Latin script names.
func collectAltNames(record SanctionsListRecord) []string {
	var altNames []string
	if len(record.Names) > 1 {
		for _, name := range record.Names[1:] {
			if trimmed := strings.TrimSpace(name); trimmed != "" {
				altNames = append(altNames, trimmed)
			}
		}
	}
	for _, name := range record.NonLatinScriptNames {
		if trimmed := strings.TrimSpace(name); trimmed != "" {
			altNames = append(altNames, trimmed)
		}
	}
	return altNames
}

func mapPerson(record SanctionsListRecord) *search.Person {
	person := &search.Person{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		person.Name = strings.TrimSpace(record.Names[0])
	}

	person.AltNames = collectAltNames(record)

	// Title
	if record.NameTitle != "" {
		person.Titles = []string{record.NameTitle}
	}

	// Place of birth
	if record.CountryOfBirth != "" {
		person.PlaceOfBirth = record.CountryOfBirth
	}

	return person
}

func mapBusiness(record SanctionsListRecord) *search.Business {
	business := &search.Business{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		business.Name = strings.TrimSpace(record.Names[0])
	}

	business.AltNames = collectAltNames(record)

	return business
}

func mapVessel(record SanctionsListRecord) *search.Vessel {
	vessel := &search.Vessel{}

	// Primary name (first in list)
	if len(record.Names) > 0 {
		vessel.Name = strings.TrimSpace(record.Names[0])
	}

	vessel.AltNames = collectAltNames(record)

	// Flag from address countries (ships often have flag country)
	if len(record.AddressCountries) > 0 && record.AddressCountries[0] != "" {
		vessel.Flag = norm.Country(record.AddressCountries[0])
	}

	return vessel
}

func mapAddresses(record SanctionsListRecord) []search.Address {
	if len(record.Addresses) == 0 {
		return nil
	}

	var addresses []search.Address

	// Each address in record.Addresses is already a combined string
	for i, addrStr := range record.Addresses {
		if strings.TrimSpace(addrStr) == "" {
			continue
		}

		addr := search.Address{
			Line1: addrStr,
		}

		// Add country if available (parallel array)
		if i < len(record.AddressCountries) && record.AddressCountries[i] != "" {
			addr.Country = norm.Country(record.AddressCountries[i])
		}

		// Add state/locality if available
		if i < len(record.StateLocalities) && record.StateLocalities[i] != "" {
			addr.State = record.StateLocalities[i]
		}

		addresses = append(addresses, addr)
	}

	return addresses
}

func mapSanctionsInfo(record SanctionsListRecord) *search.SanctionsInfo {
	if record.UNReferenceNumber == "" {
		return nil
	}
	return &search.SanctionsInfo{
		Description: "UN Reference: " + record.UNReferenceNumber,
	}
}
