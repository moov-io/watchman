// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

// SDN is a specially Designated National
type SDN struct {
	// EntityID (ent_num) is the unique record identifier/unique listing identifier
	EntityID string `json:"entityID"`
	// SDNName (SDN_name)  is the name of the specially designated national
	SDNName string `json:"sdnName"`
	// SDNType (SDN_Type) is the type of SDN
	SDNType string `json:"sdnType"`
	// Programs is the sanction programs this SDN was added from
	Programs []string `json:"program"`
	// Title is the title of an individual
	Title string `json:"title"`
	// CallSign (Call_Sign) is vessel call sign
	CallSign string `json:"callSign"`
	// VesselType (Vess_type) is the vessel type
	VesselType string `json:"vesselType"`
	// Tonnage is the vessel tonnage
	Tonnage string `json:"tonnage"`
	// GrossRegisteredTonnage (GRT) is gross registered tonnage
	GrossRegisteredTonnage string `json:"grossRegisteredTonnage"`
	// VesselFlag (Vess_flag) is vessel flag
	VesselFlag string `json:"vesselFlag"`
	// VesselOwner  (Vess_owner) is vessel owner
	VesselOwner string `json:"vesselOwner"`
	//  Remarks is remarks on specially designated national
	Remarks string `json:"remarks"`
}

// Address is OFAC SDN Addresses
type Address struct {
	// EntityID (ent_num) is the unique record identifier/unique listing identifier
	EntityID string `json:"entityID"`
	// AddressID (add_num) is the unique record identifier for the address
	AddressID string `json:"addressID"`
	// Address is the street address of the specially designated national
	Address string `json:"address"`
	// CityStateProvincePostalCode is the city, state/province, zip/postal code for the address of the
	// specially designated national
	CityStateProvincePostalCode string `json:"cityStateProvincePostalCode"`
	// Country is the country for the address of the specially designated national
	Country string `json:"country"`
	//AddressRemarks (Add_remarks) is remarks on the address
	AddressRemarks string `json:"addressRemarks"`
}

// AlternateIdentity is OFAC SDN Alternate Identity object
type AlternateIdentity struct {
	// EntityID (ent_num) is the unique record identifier/unique listing identifier
	EntityID string `json:"entityID"`
	// AlternateID (alt_num) is the unique record identifier for the alternate identity
	AlternateID string `json:"alternateID"`
	// AlternateIdentityType (alt_type) is the type of alternate identity (aka, fka, nka)
	AlternateType string `json:"alternateType"`
	// AlternateIdentityName (alt_name) is the alternate identity name of the specially designated national
	AlternateName string `json:"alternateName"`
	// AlternateIdentityRemarks (alt_remarks) is remarks on alternate identity of the specially designated national
	AlternateRemarks string `json:"alternateRemarks"`
}

// SDNComments is OFAC SDN Additional Comments
type SDNComments struct {
	// EntityID (ent_num) is the unique record identifier/unique listing identifier
	EntityID string `json:"entityID"`
	// RemarksExtended is remarks extended on a Specially Designated National
	RemarksExtended string `json:"remarksExtended"`
	// DigitalCurrencyAddresses are wallet addresses for digital currencies
	DigitalCurrencyAddresses []DigitalCurrencyAddress `json:"digitalCurrencyAddresses"`
}

type DigitalCurrencyAddress struct {
	// Currency is the name of the digital currency.
	// Examples: XBT (Bitcoin), ETH (Ethereum)
	Currency string

	// Address is a digital wallet address
	Address string
}
