// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package us_tel

import (
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

func addPersonToEntity(name string, aliases []string, entity *search.Entity[search.Value]) {
	if entity.Type != search.EntityPerson {
		return
	}

	entity.Person = &search.Person{
		Name: name,
	}

	for _, alias := range aliases {
		entity.Person.AltNames = append(entity.Person.AltNames, alias)
	}

	// FIX: Removed "return entity" since this function has no return type specified
}

func addBusinessToEntity(name string, aliases []string, entity *search.Entity[search.Value]) {
	if entity.Type != search.EntityBusiness {
		return
	}

	entity.Business = &search.Business{
		Name: name,
	}

	for _, alias := range aliases {
		entity.Business.AltNames = append(entity.Business.AltNames, alias)
	}

	// FIX: Removed "return entity" since this function has no return type specified
}

func (r TELRecord) ToEntity() search.Entity[search.Value] {
	name := r.primaryName()
	aliases := r.altNames(name)

	entity := search.Entity[search.Value]{
		SourceID:   r.ID,
		Name:       name,
		Type:       entityTypeForSchema(r.Schema),
		Source:     search.SourceUSTEL,
		SourceData: r,
	}

	// FIX: Call them directly. The functions modify 'entity' in-place via its pointer memory address.
	addPersonToEntity(name, aliases, &entity)
	addBusinessToEntity(name, aliases, &entity)

	return entity.Normalize()
}

func (r TELRecord) primaryName() string {
	for _, name := range r.Properties.Name {
		name = strings.TrimSpace(name)
		if name != "" {
			return name
		}
	}

	return strings.TrimSpace(r.Caption)
}

func (r TELRecord) altNames(primary string) []string {
	seen := map[string]bool{}
	alts := []string{}

	primary = strings.TrimSpace(primary)
	if primary != "" {
		seen[strings.ToLower(primary)] = true
	}

	for _, name := range r.Properties.Name {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		lower := strings.ToLower(name)
		if seen[lower] {
			continue
		}

		alts = append(alts, name)
		seen[lower] = true
	}

	return alts
}

func entityTypeForSchema(schema string) search.EntityType {
	switch strings.ToLower(strings.TrimSpace(schema)) {
	case "person":
		return search.EntityPerson
	default:
		return search.EntityBusiness
	}
}
