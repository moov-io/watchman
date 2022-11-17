// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

type UKCSL map[int]*UKCSLRecord

// Indices we care about for UK - CSL row data
const (
	UKNameIdx          = 0
	UKTitleIdx         = 6
	DOBhIdx            = 10
	TownOfBirthIdx     = 11
	CountryOfBirthIdx  = 12
	UKNationalitiesIdx = 13
	AddressOneIdx      = 19
	AddressTwoIdx      = 20
	AddressThreeIdx    = 21
	AddressFourIdx     = 22
	AddressFiveIdx     = 23
	AddressSixIdx      = 24
	PostalCodeIdx      = 25
	CountryIdx         = 26
	OtherInfoIdx       = 27
	GroupTypeidx       = 28
	ListedDateIdx      = 32
	UKSancListDateIdx  = 33
	LastUpdatedIdx     = 34
	GroupdIdx          = 35
)

// UK is the UK Consolidated List of Financial Sanctions Targets
type UKCSLRecord struct {
	Names             []string `json:"names"`
	Titles            []string `json:"titles"`
	DatesOfBirth      []string `json:"datesOfBirth"`
	TownsOfBirth      []string `json:"townsOfBirth"`
	CountriesOfBirth  []string `json:"countriesOfBirth"`
	Nationalities     []string `json:"nationalities"`
	Addresses         []string `json:"addresses"`
	AddressesTwo      []string `json:"addressesTwo"`
	AddressesThree    []string `json:"addressesThree"`
	AddressesFour     []string `json:"addressesFour"`
	AddressesFive     []string `json:"addressesFive"`
	AddressesSix      []string `json:"addressesSix"`
	PostalCodes       []string `json:"postalCodes"`
	Countries         []string `json:"countries"`
	OtherInfos        []string `json:"otherInfo"`
	GroupTypes        []string `json:"groupType"`
	ListedDates       []string `json:"listedDate"`
	SanctionListDates []string `json:"sanctionListDate"`
	LastUpdates       []string `json:"lastUpdated"`
	GroupID           int      `json:"groupId"`
}
