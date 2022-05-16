package csl

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
