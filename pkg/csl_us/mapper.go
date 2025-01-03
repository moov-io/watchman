// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"github.com/moov-io/watchman/pkg/search"
)

// Entity List – Bureau of Industry and Security
func EL_ToEntity(record EL) search.Entity[EL] {
	out := search.Entity[EL]{
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
func MEU_ToEntity(record MEU) search.Entity[MEU] {
	out := search.Entity[MEU]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Name      string `json:"name"`
	// Addresses string `json:"addresses"`

	return out
}

// Sectoral Sanctions Identifications List (SSI) - Treasury Department
func SSI_ToEntity(record SSI) search.Entity[SSI] {
	out := search.Entity[SSI]{
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
func UVL_ToEntity(record UVL) search.Entity[UVL] {
	out := search.Entity[UVL]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name
	// Addresses     []string `json:"addresses"`

	return out
}

// Foreign Sanctions Evaders (FSE) - Treasury Department
func FSE_ToEntity(record FSE) search.Entity[FSE] {
	out := search.Entity[FSE]{
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
func ISN_ToEntity(record ISN) search.Entity[ISN] {
	out := search.Entity[ISN]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// Remarks               []string `json:"remarks,omitempty"`
	// AlternateNames        []string `json:"alternateNames,omitempty"`

	return out
}

// Palestinian Legislative Council List (PLC) - Treasury Department
func PLC_ToEntity(record PLC) search.Entity[PLC] {
	out := search.Entity[PLC]{
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
func CAP_ToEntity(record CAP) search.Entity[CAP] {
	out := search.Entity[CAP]{
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
func DTC_ToEntity(record DTC) search.Entity[DTC] {
	out := search.Entity[DTC]{
		Source:     search.SourceUSCSL,
		SourceData: record,
	}

	out.Name = record.Name

	// AlternateNames        []string `json:"alternateNames,omitempty"`

	return out
}

// Non-SDN Chinese Military-Industrial Complex Companies List (CMIC) - Treasury Department
func CMIC_ToEntity(record CMIC) search.Entity[CMIC] {
	out := search.Entity[CMIC]{
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
func NS_MBS_ToEntity(record NS_MBS) search.Entity[NS_MBS] {
	out := search.Entity[NS_MBS]{
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
