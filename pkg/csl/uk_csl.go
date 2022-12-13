// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl

type UKCSL map[int]*UKCSLRecord

// Indices we care about for UK - CSL row data
const (
	UKNameIdx      = 0
	UKNameTwoIdx   = 1
	UKNameThreeIdx = 2
	UKNameFourIdx  = 3
	UKNameFiveIdx  = 4

	UKTitleIdx         = 6
	DOBhIdx            = 10
	TownOfBirthIdx     = 11
	CountryOfBirthIdx  = 12
	UKNationalitiesIdx = 13

	AddressOneIdx   = 19
	AddressTwoIdx   = 20
	AddressThreeIdx = 21
	AddressFourIdx  = 22
	AddressFiveIdx  = 23
	AddressSixIdx   = 24

	PostalCodeIdx     = 25
	CountryIdx        = 26
	OtherInfoIdx      = 27
	GroupTypeIdx      = 28
	ListedDateIdx     = 32
	UKSancListDateIdx = 33
	LastUpdatedIdx    = 34
	GroupdIdx         = 35
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
	PostalCodes       []string `json:"postalCodes"`
	Countries         []string `json:"countries"`
	OtherInfos        []string `json:"otherInfo"`
	GroupType         string   `json:"groupType"`
	ListedDates       []string `json:"listedDate"`
	SanctionListDates []string `json:"sanctionListDate"`
	LastUpdates       []string `json:"lastUpdated"`
	GroupID           int      `json:"groupId"`
}

type UKSanctionsListMap map[string]*UKSanctionsListRecord

const (
	UKSL_LastUpdatedIdx       = 0
	UKSL_UniqueIDIdx          = 1
	UKSL_OFSI_GroupIDIdx      = 2 // this is the group ID from the consolidated sanctions list
	UKSL_UNReferenceNumberIdx = 3
	// Name info
	UKSL_Name6Idx            = 4
	UKSL_Name1Idx            = 5
	UKSL_Name2Idx            = 6
	UKSL_Name3Idx            = 7
	UKSL_Name4Idx            = 8
	UKSL_Name5Idx            = 9
	UKSL_NameTypeIdx         = 10 // either Primary Name or Alias
	UKSL_AliasStrengthIdx    = 11
	UKSL_TitleIdx            = 12
	UKSL_NonLatinScriptIdx   = 13
	UKSL_NonLatinTypeIdx     = 14
	UKSL_NonLatinLanguageIdx = 15
	UKSL_EntityTypeIdx       = 17 // individual, entity, ship
	UKSL_OtherInfoIdx        = 20
	// Address Info
	UKSL_AddressLine1Idx   = 22
	UKSL_AddressLine2Idx   = 23
	UKSL_AddressLine3Idx   = 24
	UKSL_AddressLine4Idx   = 25
	UKSL_AddressLine5Idx   = 26
	UKSL_AddressLine6Idx   = 27
	UKSL_PostalCodeIdx     = 28
	UKSL_AddressCountryIdx = 29
	UKSL_CountryOfBirthIdx = 43
)

type UKSanctionsListRecord struct {
	LastUpdated         string
	UniqueID            string
	OFSIGroupID         string
	UNReferenceNumber   string
	Names               []string
	NameTitle           string
	NonLatinScriptNames []string
	EntityType          *UKSLEntityType
	Addresses           []string
	StateLocalities     []string
	AddressCountries    []string
	CountryOfBirth      string
}

type UKSLEntityType string

var EntityStringMap map[string]UKSLEntityType = map[string]UKSLEntityType{
	"Individual": UKSLIndividual,
	"Entity":     UKSLEntity,
	"Ship":       UKSLShip,
}

var EntityEnumMap map[UKSLEntityType]string = map[UKSLEntityType]string{
	UKSLIndividual: "Individual",
	UKSLEntity:     "Entity",
	UKSLShip:       "Ship",
}

const (
	Undefined      UKSLEntityType = ""
	UKSLIndividual UKSLEntityType = "Individual"
	UKSLEntity     UKSLEntityType = "Entity"
	UKSLShip       UKSLEntityType = "Ship"
)

func (et UKSLEntityType) String() string {
	return string(et)
}
