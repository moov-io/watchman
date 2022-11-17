// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

const (
	UKNameIdx = 0
	UKTitleIdx = 6
	DOBhIdx = 10
	TownOfBirthIdx = 11
	CountryOfBirthIdx = 12
	UKNationalitiesIdx = 13
	AddressOneIdx = 19 // Address Line 1
	AddressTwoIdx = 20 // Address Line 2
	AddressThreeIdx = 21 // 
	AddressFourIdx = 22 
	AddressFiveIdx = 23
	AddressSixIdx = 24
	PostalCodeIdx = 25
	CountryIdx = 26
	OtherInfoIdx = 27
	GroupTypeidx = 28
	ListedDateIdx = 32
	UKSancListDateIdx = 33
	LastUpdatedIdx = 34
	GroupdIdx = 35
)

// UK is the UK Consolidated List of Financial Sanctions Targets
type UKCSLRecord struct {
	Names []string `json:"names"`
	Titles []string `json:"titles"`
	DatesOfBirth []string `json:"datesOfBirth"`
	TownsOfBirth []string `json:"townsOfBirth"`
	CountriesOfBirth []string `json:"countriesOfBirth"`
	Nationalities []string `json:"nationalities"`
	Addresses []string `json:"addresses"` // Address Line 1
	AddressesTwo []string `json:"addressesTwo"` // Address Line 2
	AddressesThree []string `json:"addressesThree"` // Address Line 3
	AddressesFour []string `json:"addressesFour"` // Address Line 4
	AddressesFive []string `json:"addressesFive"` // Address Line 5
	AddressesSix []string `json:"addressesSix"` // Address Line 6
	PostalCodes []string `json:"postalCodes"`
	Countries []string `json:"countries"`
	OtherInfos []string `json:"otherInfo"`
	GroupTypes []string `json:"groupType"`
	ListedDates []string `json:"listedDate"`
	SanctionListDates []string `json:"sanctionListDate"`
	LastUpdates []string `json:"lastUpdated"`
	GroupIDs []int `json:"groups"`
}


