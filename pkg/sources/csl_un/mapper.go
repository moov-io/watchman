// Copyright Bloomfielddev Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_un

import (
	"fmt"
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

// ToEntity converts a UNIndividual to the Moov search.Entity format.
func (p UNIndividual) ToEntity() search.Entity[search.Value] {
	// Construct full name from all parts
	nameParts := []string{p.FirstName, p.SecondName, p.ThirdName, p.FourthName}
	fullName := strings.TrimSpace(strings.Join(nameParts, " "))

	entity := search.Entity[search.Value]{
		SourceID:   fmt.Sprintf("UN-%s", p.DataID), // Prefix to avoid collisions
		Name:       fullName,
		Type:       search.EntityPerson,
		Source:     search.SourceUNCSL,
		SourceData: p,
	}

	// Ensure Person pointer is initialized before appending to its fields
	if entity.Person == nil {
		entity.Person = &search.Person{}
	}
	//set the full name in the Person struct as well for easier access
	entity.Person.Name = fullName
	// Map Aliases
	for _, alias := range p.Aliases {
		if alias.Name != "" {
			entity.Person.AltNames = append(entity.Person.AltNames, alias.Name)
		}
	}

	// Map Addresses
	for _, addr := range p.Addresses {
		entity.Addresses = append(entity.Addresses, search.Address{
			City:    addr.City,
			Country: addr.Country,
		})
	}

	return entity
}

// ToEntity converts a UNEntity to the Moov search.Entity format.
func (e UNEntity) ToEntity() search.Entity[search.Value] {

	entity := search.Entity[search.Value]{
		SourceID:   fmt.Sprintf("UN-%s", e.DataID),
		Name:       e.FirstName, // Entities store their name in FIRST_NAME
		Type:       search.EntityBusiness,
		Source:     search.SourceUNCSL,
		SourceData: e,
	}

	// Ensure Business pointer is initialized before appending to its fields
	if entity.Business == nil {
		entity.Business = &search.Business{}
	}

	//set business name in the Business struct as well for easier access
	entity.Business.Name = e.FirstName
	for _, alias := range e.Aliases {
		if alias.Name != "" {
			entity.Business.AltNames = append(entity.Business.AltNames, alias.Name)
		}
	}

	return entity
}
