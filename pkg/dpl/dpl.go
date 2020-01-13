// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

// DPL is the BIS Denied Persons List
type DPL struct {
	// Name is the name of the Denied Person
	Name string `json:"name"`
	// StreetAddress is the Denied Person's street address
	StreetAddress string `json:"streetAddress"`
	// City is the Denied Person's city
	City string `json:"city"`
	// State is the Denied Person's state
	State string `json:"state"`
	// Country is the Denied Person's country
	Country string `json:"country"`
	// PostalCode is the Denied Person's postal code
	PostalCode string `json:"postalCode"`
	// EffectiveDate is the date the denial came into effect
	EffectiveDate string `json:"effectiveDate"`
	// ExpirationDate is the date the denial expires. If blank, the denial has no expiration
	ExpirationDate string `json:"expirationDate"`
	// StandardOrder denotes whether or not the Person was added to the list by a "standard" order
	StandardOrder string `json:"standardOrder"`
	// LastUpdate is the date of the most recent change to the denial
	LastUpdate string `json:"lastUpdate"`
	// Action is the most recent action taken regarding the denial
	Action string `json:"action"`
	// FRCitation is the reference to the order's citation in the Federal Register
	FRCitation string `json:"frCitation"`
}
