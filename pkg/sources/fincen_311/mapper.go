// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"strings"

	"github.com/moov-io/watchman/pkg/search"
)

// ConvertSpecialMeasures converts parsed data to search entities
func ConvertSpecialMeasures(data *ListData) []search.Entity[search.Value] {
	var out []search.Entity[search.Value]
	for _, sm := range data.SpecialMeasures {
		out = append(out, ToEntity(sm))
	}
	return out
}

// ToEntity converts a SpecialMeasure to a search Entity
func ToEntity(sm SpecialMeasure) search.Entity[search.Value] {
	entity := search.Entity[search.Value]{
		Name:       sm.EntityName,
		Source:     search.SourceUSFinCEN311,
		SourceID:   generateSourceID(sm),
		SourceData: sm,
	}

	// Map entity type based on classification
	switch sm.EntityType {
	case SMTypeFinancialInstitution:
		entity.Type = search.EntityBusiness
		entity.Business = &search.Business{
			Name: sm.EntityName,
		}

	case SMTypeJurisdiction:
		entity.Type = search.EntityOrganization
		entity.Organization = &search.Organization{
			Name: sm.EntityName,
		}

	case SMTypeTransactionClass:
		entity.Type = search.EntityOrganization
		entity.Organization = &search.Organization{
			Name: sm.EntityName,
		}

	default:
		// Default to business for unknown types
		entity.Type = search.EntityBusiness
		entity.Business = &search.Business{
			Name: sm.EntityName,
		}
	}

	// Store PDF links in Contact.Websites for audit/compliance access
	entity.Contact = buildContactInfo(sm)

	// Build sanctions info with program details
	entity.SanctionsInfo = buildSanctionsInfo(sm)

	return entity.Normalize()
}

func generateSourceID(sm SpecialMeasure) string {
	// Create a stable ID from the entity name
	id := strings.ToLower(sm.EntityName)
	id = strings.ReplaceAll(id, " ", "_")
	id = strings.ReplaceAll(id, "'", "")
	id = strings.ReplaceAll(id, ",", "")
	id = strings.ReplaceAll(id, ".", "")
	id = strings.ReplaceAll(id, "(", "")
	id = strings.ReplaceAll(id, ")", "")
	return "fincen311_" + id
}

func buildContactInfo(sm SpecialMeasure) search.ContactInfo {
	var websites []string

	// Store PDF document links as "websites" for compliance access
	if sm.FindingURL != "" {
		websites = append(websites, sm.FindingURL)
	}
	if sm.NPRMURL != "" {
		websites = append(websites, sm.NPRMURL)
	}
	if sm.FinalRuleURL != "" {
		websites = append(websites, sm.FinalRuleURL)
	}
	if sm.RescindedURL != "" {
		websites = append(websites, sm.RescindedURL)
	}

	return search.ContactInfo{
		Websites: websites,
	}
}

func buildSanctionsInfo(sm SpecialMeasure) *search.SanctionsInfo {
	programs := []string{"FinCEN Special Measures"}

	// Add specific program based on entity type
	switch sm.EntityType {
	case SMTypeFinancialInstitution:
		programs = append(programs, "Section 311")
	case SMTypeJurisdiction:
		programs = append(programs, "Section 311")
	case SMTypeTransactionClass:
		programs = append(programs, "Section 9714")
	}

	var descParts []string
	descParts = append(descParts, "FinCEN Special Measures Target")

	if sm.FindingDate != "" {
		descParts = append(descParts, "Finding: "+sm.FindingDate)
	}
	if sm.FinalRuleDate != "" {
		descParts = append(descParts, "Final Rule: "+sm.FinalRuleDate)
	}
	if sm.IsRescinded {
		status := "Status: RESCINDED"
		if sm.RescindedDate != "" {
			status += " (" + sm.RescindedDate + ")"
		}
		descParts = append(descParts, status)
	} else {
		descParts = append(descParts, "Status: ACTIVE")
	}

	return &search.SanctionsInfo{
		Programs:    programs,
		Description: strings.Join(descParts, ". ") + ".",
	}
}
