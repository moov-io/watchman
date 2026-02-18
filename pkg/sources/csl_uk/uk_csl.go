// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_uk

type CSL map[int]*CSLRecord

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
type CSLRecord struct {
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

type SanctionsListMap map[string]*SanctionsListRecord

// CSV column indices for UK Sanctions List
// Source: https://sanctionslist.fcdo.gov.uk/docs/UK-Sanctions-List.csv
const (
	UKSL_LastUpdatedIdx       = 0
	UKSL_UniqueIDIdx          = 1
	UKSL_OFSI_GroupIDIdx      = 2 // this is the group ID from the consolidated sanctions list
	UKSL_UNReferenceNumberIdx = 3
	// Name info
	UKSL_Name6Idx              = 4
	UKSL_Name1Idx              = 5
	UKSL_Name2Idx              = 6
	UKSL_Name3Idx              = 7
	UKSL_Name4Idx              = 8
	UKSL_Name5Idx              = 9
	UKSL_NameTypeIdx           = 10 // either Primary Name or Alias
	UKSL_AliasStrengthIdx      = 11
	UKSL_TitleIdx              = 12
	UKSL_NonLatinScriptIdx     = 13
	UKSL_NonLatinTypeIdx       = 14
	UKSL_NonLatinLanguageIdx   = 15
	UKSL_RegimeNameIdx         = 16
	UKSL_EntityTypeIdx         = 17 // Individual, Entity, Ship (was "Designation Type" in old docs)
	UKSL_DesignationSourceIdx  = 18
	UKSL_SanctionsImposedIdx   = 19
	UKSL_OtherInfoIdx          = 20
	UKSL_StatementOfReasonsIdx = 21
	// Address Info
	UKSL_AddressLine1Idx   = 22
	UKSL_AddressLine2Idx   = 23
	UKSL_AddressLine3Idx   = 24
	UKSL_AddressLine4Idx   = 25
	UKSL_AddressLine5Idx   = 26
	UKSL_AddressLine6Idx   = 27
	UKSL_PostalCodeIdx     = 28
	UKSL_AddressCountryIdx = 29
	// Contact Info
	UKSL_PhoneNumberIdx  = 30
	UKSL_WebsiteIdx      = 31
	UKSL_EmailAddressIdx = 32
	// Dates and Personal Info
	UKSL_DateDesignatedIdx       = 33
	UKSL_DOBIdx                  = 34
	UKSL_NationalityIdx          = 35
	UKSL_NationalIDNumberIdx     = 36
	UKSL_NationalIDAdditionalIdx = 37
	UKSL_PassportNumberIdx       = 38
	UKSL_PassportAdditionalIdx   = 39
	UKSL_PositionIdx             = 40
	UKSL_GenderIdx               = 41
	UKSL_TownOfBirthIdx          = 42
	UKSL_CountryOfBirthIdx       = 43
	// Note: Column 44 "Type of entity" duplicates column 17 "Designation Type" - we use column 17
	// Business/Ship specific
	UKSL_SubsidiariesIdx      = 45
	UKSL_ParentCompanyIdx     = 46
	UKSL_BusinessRegNumberIdx = 47
	UKSL_IMONumberIdx         = 48
	UKSL_CurrentOwnerIdx      = 49
	UKSL_PreviousOwnerIdx     = 50
	UKSL_CurrentFlagIdx       = 51
	UKSL_PreviousFlagsIdx     = 52
	UKSL_TypeOfShipIdx        = 53
	UKSL_TonnageIdx           = 54
	UKSL_LengthIdx            = 55
	UKSL_YearBuiltIdx         = 56
	UKSL_HullIDNumberIdx      = 57
)

type SanctionsListRecord struct {
	LastUpdated         string
	UniqueID            string
	OFSIGroupID         string
	UNReferenceNumber   string
	Names               []string
	NameTitle           string
	NonLatinScriptNames []string
	EntityType          *SLEntityType
	Addresses           []string
	StateLocalities     []string
	AddressCountries    []string
	AddressPostalCodes  []string
	CountryOfBirth      string
	TownOfBirth         string
	// New fields from CSV
	DOB                      string
	Nationality              string
	PassportNumber           string
	PassportAdditionalInfo   string
	NationalIDNumber         string
	NationalIDAdditionalInfo string
	Position                 string
	Gender                   string
	Regime                   string
	DateDesignated           string
	OtherInfo                string
	// Vessel specific
	IMONumber   string
	VesselType  string
	Tonnage     string
	VesselFlag  string
	VesselOwner string
	// Business specific
	BusinessRegNumber string
}

type SLEntityType string

var EntityStringMap map[string]SLEntityType = map[string]SLEntityType{
	"Individual": UKSLIndividual,
	"Entity":     UKSLEntity,
	"Ship":       UKSLShip,
}

var EntityEnumMap map[SLEntityType]string = map[SLEntityType]string{
	UKSLIndividual: "Individual",
	UKSLEntity:     "Entity",
	UKSLShip:       "Ship",
}

const (
	Undefined      SLEntityType = ""
	UKSLIndividual SLEntityType = "Individual"
	UKSLEntity     SLEntityType = "Entity"
	UKSLShip       SLEntityType = "Ship"
)

func (et SLEntityType) String() string {
	return string(et)
}
