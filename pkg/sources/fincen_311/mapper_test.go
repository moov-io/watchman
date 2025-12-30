// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

import (
	"testing"

	"github.com/moov-io/watchman/pkg/search"
	"github.com/stretchr/testify/require"
)

func TestConvertSpecialMeasures(t *testing.T) {
	data := &ListData{
		SpecialMeasures: []SpecialMeasure{
			{
				EntityName:    "ABLV Bank",
				EntityType:    SMTypeFinancialInstitution,
				FindingURL:    "https://www.fincen.gov/finding.pdf",
				FindingDate:   "02/13/2018",
				FinalRuleURL:  "https://www.fincen.gov/final.pdf",
				FinalRuleDate: "06/25/2018",
				IsRescinded:   false,
			},
			{
				EntityName:    "Islamic Republic of Iran",
				EntityType:    SMTypeJurisdiction,
				FindingURL:    "https://www.fincen.gov/iran.pdf",
				FindingDate:   "11/21/2011",
				FinalRuleURL:  "https://www.fincen.gov/iran_final.pdf",
				FinalRuleDate: "01/01/2012",
				IsRescinded:   false,
			},
		},
		ListHash: "abc123",
	}

	entities := ConvertSpecialMeasures(data)
	require.Len(t, entities, 2)

	// Check first entity
	ablv := entities[0]
	require.Equal(t, "ABLV Bank", ablv.Name)
	require.Equal(t, search.SourceUSFinCEN311, ablv.Source)
	require.Equal(t, search.EntityBusiness, ablv.Type)
	require.NotNil(t, ablv.Business)
	require.Equal(t, "ABLV Bank", ablv.Business.Name)
	require.Contains(t, ablv.SourceID, "fincen311_ablv_bank")

	// Check websites contain PDF links
	require.Contains(t, ablv.Contact.Websites, "https://www.fincen.gov/finding.pdf")
	require.Contains(t, ablv.Contact.Websites, "https://www.fincen.gov/final.pdf")

	// Check sanctions info
	require.NotNil(t, ablv.SanctionsInfo)
	require.Contains(t, ablv.SanctionsInfo.Programs, "FinCEN Special Measures")
	require.Contains(t, ablv.SanctionsInfo.Programs, "Section 311")
	require.Contains(t, ablv.SanctionsInfo.Description, "ACTIVE")

	// Check second entity (jurisdiction)
	iran := entities[1]
	require.Equal(t, "Islamic Republic of Iran", iran.Name)
	require.Equal(t, search.EntityOrganization, iran.Type)
	require.NotNil(t, iran.Organization)
}

func TestToEntity_FinancialInstitution(t *testing.T) {
	sm := SpecialMeasure{
		EntityName:    "Test Bank Ltd.",
		EntityType:    SMTypeFinancialInstitution,
		FindingURL:    "https://example.com/finding.pdf",
		FindingDate:   "01/01/2020",
		FinalRuleURL:  "https://example.com/final.pdf",
		FinalRuleDate: "06/01/2020",
		IsRescinded:   false,
	}

	entity := ToEntity(sm)

	require.Equal(t, search.EntityBusiness, entity.Type)
	require.NotNil(t, entity.Business)
	require.Nil(t, entity.Organization)
	require.Equal(t, "Test Bank Ltd.", entity.Business.Name)
}

func TestToEntity_Jurisdiction(t *testing.T) {
	sm := SpecialMeasure{
		EntityName:  "Test Country",
		EntityType:  SMTypeJurisdiction,
		IsRescinded: false,
	}

	entity := ToEntity(sm)

	require.Equal(t, search.EntityOrganization, entity.Type)
	require.NotNil(t, entity.Organization)
	require.Nil(t, entity.Business)
	require.Equal(t, "Test Country", entity.Organization.Name)
}

func TestToEntity_TransactionClass(t *testing.T) {
	sm := SpecialMeasure{
		EntityName:  "Virtual Currency Mixing",
		EntityType:  SMTypeTransactionClass,
		IsRescinded: false,
	}

	entity := ToEntity(sm)

	require.Equal(t, search.EntityOrganization, entity.Type)
	require.NotNil(t, entity.Organization)

	// Transaction class should have Section 9714 in programs
	require.Contains(t, entity.SanctionsInfo.Programs, "Section 9714")
}

func TestToEntity_Rescinded(t *testing.T) {
	sm := SpecialMeasure{
		EntityName:    "Rescinded Bank",
		EntityType:    SMTypeFinancialInstitution,
		FinalRuleURL:  "https://example.com/final.pdf",
		FinalRuleDate: "01/01/2020",
		RescindedURL:  "https://example.com/rescind.pdf",
		RescindedDate: "12/01/2024",
		IsRescinded:   true,
	}

	entity := ToEntity(sm)

	require.Contains(t, entity.SanctionsInfo.Description, "RESCINDED")
	require.Contains(t, entity.SanctionsInfo.Description, "12/01/2024")

	// Rescinded URL should be in websites
	require.Contains(t, entity.Contact.Websites, "https://example.com/rescind.pdf")
}

func TestGenerateSourceID(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"ABLV Bank", "fincen311_ablv_bank"},
		{"Islamic Republic of Iran", "fincen311_islamic_republic_of_iran"},
		{"Bank's Name, Inc.", "fincen311_banks_name_inc"},
		{"Test (Entity)", "fincen311_test_entity"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sm := SpecialMeasure{EntityName: tc.name}
			result := generateSourceID(sm)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestBuildSanctionsInfo(t *testing.T) {
	// Test financial institution
	sm := SpecialMeasure{
		EntityName:    "Test Bank",
		EntityType:    SMTypeFinancialInstitution,
		FindingDate:   "01/01/2020",
		FinalRuleDate: "06/01/2020",
		IsRescinded:   false,
	}

	info := buildSanctionsInfo(sm)
	require.Contains(t, info.Programs, "FinCEN Special Measures")
	require.Contains(t, info.Programs, "Section 311")
	require.Contains(t, info.Description, "Finding: 01/01/2020")
	require.Contains(t, info.Description, "Final Rule: 06/01/2020")
	require.Contains(t, info.Description, "Status: ACTIVE")
}
