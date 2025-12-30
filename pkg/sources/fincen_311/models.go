// Copyright The Moov Authors
// SPDX-License-Identifier: Apache-2.0

package fincen_311

// SpecialMeasure represents a single FinCEN 311/9714 Special Measures entry
type SpecialMeasure struct {
	EntityName    string `json:"entityName"`
	EntityType    SMType `json:"entityType"`
	FindingURL    string `json:"findingUrl,omitempty"`
	FindingDate   string `json:"findingDate,omitempty"`
	NPRMURL       string `json:"nprmUrl,omitempty"`
	NPRMDate      string `json:"nprmDate,omitempty"`
	FinalRuleURL  string `json:"finalRuleUrl,omitempty"`
	FinalRuleDate string `json:"finalRuleDate,omitempty"`
	RescindedURL  string `json:"rescindedUrl,omitempty"`
	RescindedDate string `json:"rescindedDate,omitempty"`
	IsRescinded   bool   `json:"isRescinded"`
}

// SMType categorizes the type of special measure target
type SMType string

const (
	SMTypeFinancialInstitution SMType = "financial_institution"
	SMTypeJurisdiction         SMType = "jurisdiction"
	SMTypeTransactionClass     SMType = "transaction_class"
)

// ListData holds parsed data and hash for change detection
type ListData struct {
	SpecialMeasures []SpecialMeasure
	ListHash        string
}
