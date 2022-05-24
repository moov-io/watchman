// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

// CSL contains each record from the Consolidate Screening List, broken down by the record's original source
type CSL struct {
	ELs  []*EL  // Entity List – Bureau of Industry and Security
	MEUs []*MEU // Military End User List
	SSIs []*SSI // Sectoral Sanctions Identifications List (SSI) - Treasury Department

	// []*UL (Unverified List – Bureau of Industry and Security)
	// []*PSE (Foreign Sanctions Evaders (FSE) - Treasury Department)
	// []*ISN (Nonproliferation Sanctions (ISN) - State Department)
	// []*PLC (Palestinian Legislative Council List (PLC) - Treasury Department)
	// []*CAPTA (CAPTA (formerly Foreign Financial Institutions Subject to Part 561 - Treasury Department))
	// []*ADL (AECA Debarred List - State Department)
}

// This is the order of the columns in the CSL
// Source: https://legacy.trade.gov/CSL_Download_Instructions.pdf
const (
	SourceIdx = iota
	EntityNumberIdx
	TypeIdx
	ProgramsIdx
	NameIdx
	TitleIdx
	AddressesIdx
	FRNoticeIdx // Federal Register Notice
	StartDateIdx
	EndDateIdx
	StandardOrderIdx
	LicenseRequirementIdx
	LicensePolicyIdx
	CallSignIdx
	VesselTypeIdx
	GrossTonnageIdx
	GrossRegisteredTonnageIdx
	VesselFlagIdx
	VesselOwnerIdx
	RemarksIdx
	SourceListURLIdx
	AltNamesIdx
	CitizenshipsIdx
	DatesOfBirthIdx
	NationalitiesIdx
	PlacesOfBirthIdx
	SourceInformationURLIdx
	IDsIdx
)

// EL is the Entity List (EL) - Bureau of Industry and Security
type EL struct {
	// ID is the unique identifier for the entity
	ID string `json:"id"`
	// Name is the primary name of the entity
	Name string `json:"name"`
	// AlternateNames is a list of aliases associated with the entity
	AlternateNames []string `json:"alternateNames"`
	// Addresses is a list of known addresses associated with the entity
	Addresses []string `json:"addresses"`
	// StartDate is the effective date
	StartDate string `json:"startDate"`
	// LicenseRequirement specifies the license requirements that it imposes on each listed person
	LicenseRequirement string `json:"licenseRequirement"`
	// LicensePolicy is the policy with which BIS reviews the requirements set forth in License Requirements
	LicensePolicy string `json:"licensePolicy"`
	// FRNotice identifies the notice in the Federal Register
	FRNotice string `json:"FRNotice"`
	// SourceListURL is a link to the official SSI list
	SourceListURL string `json:"sourceListURL"`
	// SourceInfoURL is a link to information about the list
	SourceInfoURL string `json:"sourceInfoURL"`
}

type MEU struct {
	EntityID  string `json:"entityID"`
	Name      string `json:"name"`
	Addresses string `json:"addresses"`
	FRNotice  string `json:"FRNotice"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// SSI is the Sectoral Sanctions Identifications List - Treasury Department
type SSI struct {
	// EntityID (ent_num) is the unique record identifier/unique listing identifier
	EntityID string `json:"entityID"`
	// Type is the entity type (e.g. individual, vessel, aircraft, etc)
	Type string `json:"type"`
	// Programs is the list of sanctions program for which the entity is flagged
	Programs []string `json:"programs"`
	// Name is the entity's name (e.g. given name for individual, company name, etc.)
	Name string `json:"name"`
	// Addresses is a list of known addresses associated with the entity
	Addresses []string `json:"addresses"`
	// Remarks is used to provide additional details for the entity
	Remarks []string `json:"remarks"`
	// AlternateNames is a list of aliases associated with the entity
	AlternateNames []string `json:"alternateNames"`
	// IDsOnRecord is a list of the forms of identification on file for the entity
	IDsOnRecord []string `json:"ids"`
	// SourceListURL is a link to the official SSI list
	SourceListURL string `json:"sourceListURL"`
	// SourceInfoURL is a link to information about the list
	SourceInfoURL string `json:"sourceInfoURL"`
}
