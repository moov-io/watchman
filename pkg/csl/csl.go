// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

// CSL contains each record from the Consolidate Screening List, broken down by the record's original source
type CSL struct {
	SSIs []*SSI // Sectoral Sanctions Identifications List (SSI) - Treasury Department
	ELs  []*EL  // Entity List – Bureau of Industry and Security
	// []*UL (Unverified List – Bureau of Industry and Security)
	// []*PSE (Foreign Sanctions Evaders (FSE) - Treasury Department)
	// []*ISN (Nonproliferation Sanctions (ISN) - State Department)
	// []*PLC (Palestinian Legislative Council List (PLC) - Treasury Department)
	// []*CAPTA (CAPTA (formerly Foreign Financial Institutions Subject to Part 561 - Treasury Department))
	// []*ADL (AECA Debarred List - State Department)
}

// This is the order of the columns in the CSL
const (
	SourceIdx = iota
	EntityNumberIdx
	TypeIdx
	ProgramsIdx
	NameIdx
	TitleIdx
	AddressesIdx
	FRNoticeIdx
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
