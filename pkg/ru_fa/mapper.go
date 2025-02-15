package ru_fa

import (
	"fmt"
	"github.com/moov-io/watchman/pkg/search"
)

func WrapEntities(tbls Tables) []search.Entity[search.Value] {
	var entities []search.Entity[search.Value]

	// Wrap Non-commercial Organizations (assumed as organizations)
	for i, org := range tbls.NonCommercialOrgs {
		entity := search.Entity[search.Value]{
			Name:       org.NameTranslit,
			Type:       search.EntityOrganization,
			Source:     search.SourceRUFA,
			SourceID:   fmt.Sprintf("ru_nc_%d", i+1),
			SourceData: org, // org is automatically upcast to search.Value
		}
		entities = append(entities, entity)
	}

	// Wrap Mass Media (assumed as organizations)
	for i, m := range tbls.MassMedias {
		entity := search.Entity[search.Value]{
			Name:       m.NameTranslit,
			Type:       search.EntityOrganization,
			Source:     search.SourceRUFA,
			SourceID:   fmt.Sprintf("ru_mm_%d", i+1),
			SourceData: m,
		}
		entities = append(entities, entity)
	}

	// Wrap Media Individuals (assumed as persons)
	for i, mi := range tbls.MediaIndividuals {
		entity := search.Entity[search.Value]{
			Name:       mi.NameTranslit,
			Type:       search.EntityPerson,
			Source:     search.SourceRUFA,
			SourceID:   fmt.Sprintf("ru_mi_%d", i+1),
			SourceData: mi,
		}
		entities = append(entities, entity)
	}

	// Wrap Foreign Agent Individuals (assumed as persons)
	for i, fai := range tbls.ForeignAgentIndividuals {
		entity := search.Entity[search.Value]{
			Name:       fai.NameTranslit,
			Type:       search.EntityPerson,
			Source:     search.SourceRUFA,
			SourceID:   fmt.Sprintf("ru_fai_%d", i+1),
			SourceData: fai,
		}
		entities = append(entities, entity)
	}

	// Wrap Unregistered Associations (assumed as organizations)
	for i, ua := range tbls.UnregisteredAssociations {
		entity := search.Entity[search.Value]{
			Name:       ua.NameTranslit,
			Type:       search.EntityOrganization,
			Source:     search.SourceRUFA,
			SourceID:   fmt.Sprintf("ru_ua_%d", i+1),
			SourceData: ua,
		}
		entities = append(entities, entity)
	}

	return entities
}
