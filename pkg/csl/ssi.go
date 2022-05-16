package csl

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
