package mapper

import (
	"github.com/moov-io/watchman/pkg/csl"
	"github.com/moov-io/watchman/pkg/search"
)

// Entity List – Bureau of Industry and Security
func EL_ToEntity(record csl.EL) search.Entity[csl.EL] {
	out := search.Entity[csl.EL]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name
	// out.Type =  // TODO(adam):

	// record.AlternateNames []string // TODO(adam):
	// record.Addresses []string  // TODO(adam):

	return out
}

// Military End User List
func MEU_ToEntity(record csl.MEU) search.Entity[csl.MEU] {
	out := search.Entity[csl.MEU]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Name      string `json:"name"`
	// Addresses string `json:"addresses"`

	return out
}

// Sectoral Sanctions Identifications List (SSI) - Treasury Department
func SSI_ToEntity(record csl.SSI) search.Entity[csl.SSI] {
	out := search.Entity[csl.SSI]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Addresses []string `json:"addresses"`
	// Remarks []string `json:"remarks"`
	// AlternateNames []string `json:"alternateNames"`

	// IDsOnRecord []string `json:"ids"`

	return out
}

// Unverified List – Bureau of Industry and Security
func UVL_ToEntity(record csl.UVL) search.Entity[csl.UVL] {
	out := search.Entity[csl.UVL]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name
	// Addresses     []string `json:"addresses"`

	return out
}

// Foreign Sanctions Evaders (FSE) - Treasury Department
func FSE_ToEntity(record csl.FSE) search.Entity[csl.FSE] {
	out := search.Entity[csl.FSE]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Type          string   `json:"type"`
	// Addresses     []string `json:"addresses,omitempty"`
	// DatesOfBirth  string   `json:"datesOfBirth"`
	// IDs           []string `json:"IDs"`

	return out
}

// Nonproliferation Sanctions (ISN) - State Department
func ISN_ToEntity(record csl.ISN) search.Entity[csl.ISN] {
	out := search.Entity[csl.ISN]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Remarks               []string `json:"remarks,omitempty"`
	// AlternateNames        []string `json:"alternateNames,omitempty"`

	return out
}

// Palestinian Legislative Council List (PLC) - Treasury Department
func PLC_ToEntity(record csl.PLC) search.Entity[csl.PLC] {
	out := search.Entity[csl.PLC]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Type          string   `json:"type"`
	// Addresses     []string `json:"addresses,omitempty"`
	// DatesOfBirth  string   `json:"datesOfBirth"`
	// IDs           []string `json:"IDs"`
	// Remarks               []string `json:"remarks,omitempty"`

	return out
}

// CAPTA (formerly Foreign Financial Institutions Subject to Part 561 - Treasury Department)
func CAP_ToEntity(record csl.CAP) search.Entity[csl.CAP] {
	out := search.Entity[csl.CAP]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Type          string   `json:"type"`
	// Addresses     []string `json:"addresses,omitempty"`
	// DatesOfBirth  string   `json:"datesOfBirth"`
	// IDs           []string `json:"IDs"`
	// Remarks               []string `json:"remarks,omitempty"`

	return out
}

// ITAR Debarred (DTC) - State Department
func DTC_ToEntity(record csl.DTC) search.Entity[csl.DTC] {
	out := search.Entity[csl.DTC]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// AlternateNames        []string `json:"alternateNames,omitempty"`

	return out
}

// Non-SDN Chinese Military-Industrial Complex Companies List (CMIC) - Treasury Department
func CMIC_ToEntity(record csl.CMIC) search.Entity[csl.CMIC] {
	out := search.Entity[csl.CMIC]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Type          string   `json:"type"`
	// Addresses     []string `json:"addresses,omitempty"`
	// DatesOfBirth  string   `json:"datesOfBirth"`
	// IDs           []string `json:"IDs"`
	// Remarks               []string `json:"remarks,omitempty"`

	return out
}

// Non-SDN Menu-Based Sanctions List (NS-MBS List) - Treasury Department
func NS_MBS_ToEntity(record csl.NS_MBS) search.Entity[csl.NS_MBS] {
	out := search.Entity[csl.NS_MBS]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Type          string   `json:"type"`
	// Addresses     []string `json:"addresses,omitempty"`
	// DatesOfBirth  string   `json:"datesOfBirth"`
	// IDs           []string `json:"IDs"`
	// Remarks               []string `json:"remarks,omitempty"`

	return out
}
