// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_eu

// CLS - Consolidated List Sanctions from European Union

// struct to hold the rows from the csv data before merge
type CSL map[int]*CSLRecord

// CSLRecord contains all parsed fields from EU CSL CSV (118 columns)
// Fields are grouped by their section in the CSV
type CSLRecord struct {
	// === Entity fields (columns 0-15) ===
	FileGenerationDate             string `json:"fileGenerationDate"`
	EntityLogicalID                int    `json:"entityLogicalId"`
	EntityReferenceNumber          string `json:"entityReferenceNumber"`
	EntityUnitedNationID           string `json:"entityUnitedNationId"`
	EntityDesignationDate          string `json:"entityDesignationDate"`
	EntityDesignationDetails       string `json:"entityDesignationDetails"`
	EntityRemark                   string `json:"entityRemark"`
	EntitySubjectType              string `json:"entitySubjectType"`
	EntitySubjectTypeCode          string `json:"entitySubjectTypeCode"`
	EntityRegulationType           string `json:"entityRegulationType"`
	EntityRegulationOrgType        string `json:"entityRegulationOrgType"`
	EntityRegulationPubDate        string `json:"entityRegulationPubDate"`
	EntityRegulationEntryIntoForce string `json:"entityRegulationEntryIntoForce"`
	EntityRegulationNumberTitle    string `json:"entityRegulationNumberTitle"`
	EntityRegulationProgramme      string `json:"entityRegulationProgramme"`
	EntityPublicationURL           string `json:"entityPublicationUrl"`

	// === NameAlias fields (columns 16-33) ===
	// These are collected as arrays since entities can have multiple names
	NameAliasLastNames       []string `json:"nameAliasLastNames"`
	NameAliasFirstNames      []string `json:"nameAliasFirstNames"`
	NameAliasMiddleNames     []string `json:"nameAliasMiddleNames"`
	NameAliasWholeNames      []string `json:"nameAliasWholeNames"`
	NameAliasNameLanguages   []string `json:"nameAliasNameLanguages"`
	NameAliasGenders         []string `json:"nameAliasGenders"` // M/F values
	NameAliasTitles          []string `json:"nameAliasTitles"`
	NameAliasFunctions       []string `json:"nameAliasFunctions"`
	NameAliasLogicalIDs      []int    `json:"nameAliasLogicalIds"`
	NameAliasRegLanguages    []string `json:"nameAliasRegLanguages"`
	NameAliasRemarks         []string `json:"nameAliasRemarks"`
	NameAliasRegTypes        []string `json:"nameAliasRegTypes"`
	NameAliasRegOrgTypes     []string `json:"nameAliasRegOrgTypes"`
	NameAliasRegPubDates     []string `json:"nameAliasRegPubDates"`
	NameAliasRegEntryDates   []string `json:"nameAliasRegEntryDates"`
	NameAliasRegNumberTitles []string `json:"nameAliasRegNumberTitles"`
	NameAliasRegProgrammes   []string `json:"nameAliasRegProgrammes"`
	NameAliasRegPubURLs      []string `json:"nameAliasRegPubUrls"`

	// === Address fields (columns 34-53) ===
	AddressCities              []string `json:"addressCities"`
	AddressStreets             []string `json:"addressStreets"`
	AddressPoBoxes             []string `json:"addressPoBoxes"`
	AddressZipCodes            []string `json:"addressZipCodes"`
	AddressRegions             []string `json:"addressRegions"`
	AddressPlaces              []string `json:"addressPlaces"`
	AddressAsAtListingTimes    []string `json:"addressAsAtListingTimes"`
	AddressContactInfos        []string `json:"addressContactInfos"`
	AddressCountryISOs         []string `json:"addressCountryIsos"`
	AddressCountryDescriptions []string `json:"addressCountryDescriptions"`
	AddressLogicalIDs          []int    `json:"addressLogicalIds"`
	AddressRegLanguages        []string `json:"addressRegLanguages"`
	AddressRemarks             []string `json:"addressRemarks"`
	AddressRegTypes            []string `json:"addressRegTypes"`
	AddressRegOrgTypes         []string `json:"addressRegOrgTypes"`
	AddressRegPubDates         []string `json:"addressRegPubDates"`
	AddressRegEntryDates       []string `json:"addressRegEntryDates"`
	AddressRegNumberTitles     []string `json:"addressRegNumberTitles"`
	AddressRegProgrammes       []string `json:"addressRegProgrammes"`
	AddressRegPubURLs          []string `json:"addressRegPubUrls"`

	// === BirthDate fields (columns 54-77) ===
	BirthDates          []string `json:"birthDates"`
	BirthDays           []string `json:"birthDays"`
	BirthMonths         []string `json:"birthMonths"`
	BirthYears          []string `json:"birthYears"`
	BirthYearRangeFroms []string `json:"birthYearRangeFroms"`
	BirthYearRangeTos   []string `json:"birthYearRangeTos"`
	BirthCircas         []string `json:"birthCircas"`
	BirthCalendarTypes  []string `json:"birthCalendarTypes"`
	BirthZipCodes       []string `json:"birthZipCodes"`
	BirthRegions        []string `json:"birthRegions"`
	BirthPlaces         []string `json:"birthPlaces"`
	BirthCities         []string `json:"birthCities"`
	BirthCountryISOs    []string `json:"birthCountryIsos"`
	BirthCountries      []string `json:"birthCountries"`
	BirthLogicalIDs     []int    `json:"birthLogicalIds"`
	BirthRegLanguages   []string `json:"birthRegLanguages"`
	BirthRemarks        []string `json:"birthRemarks"`
	BirthRegTypes       []string `json:"birthRegTypes"`
	BirthRegOrgTypes    []string `json:"birthRegOrgTypes"`
	BirthRegPubDates    []string `json:"birthRegPubDates"`
	BirthRegEntryDates  []string `json:"birthRegEntryDates"`
	BirthRegNumberTitle []string `json:"birthRegNumberTitles"`
	BirthRegProgrammes  []string `json:"birthRegProgrammes"`
	BirthRegPubURLs     []string `json:"birthRegPubUrls"`

	// === Identification fields (columns 78-104) ===
	// Using IdentificationInfo struct for the main fields, plus additional arrays
	Identifications []IdentificationInfo `json:"identifications"`

	// === Citizenship fields (columns 105-117) ===
	CitizenshipRegions       []string `json:"citizenshipRegions"`
	CitizenshipCountryISOs   []string `json:"citizenshipCountryIsos"`
	Citizenships             []string `json:"citizenships"` // CountryDescription
	CitizenshipLogicalIDs    []int    `json:"citizenshipLogicalIds"`
	CitizenshipRegLanguages  []string `json:"citizenshipRegLanguages"`
	CitizenshipRemarks       []string `json:"citizenshipRemarks"`
	CitizenshipRegTypes      []string `json:"citizenshipRegTypes"`
	CitizenshipRegOrgTypes   []string `json:"citizenshipRegOrgTypes"`
	CitizenshipRegPubDates   []string `json:"citizenshipRegPubDates"`
	CitizenshipRegEntryDates []string `json:"citizenshipRegEntryDates"`
	CitizenshipRegNumTitles  []string `json:"citizenshipRegNumTitles"`
	CitizenshipRegProgrammes []string `json:"citizenshipRegProgrammes"`
	CitizenshipRegPubURLs    []string `json:"citizenshipRegPubUrls"`
}

// IdentificationInfo holds parsed identification document data
type IdentificationInfo struct {
	Number          string `json:"number"`
	Diplomatic      string `json:"diplomatic"`
	KnownExpired    string `json:"knownExpired"`
	KnownFalse      string `json:"knownFalse"`
	ReportedLost    string `json:"reportedLost"`
	RevokedByIssuer string `json:"revokedByIssuer"`
	IssuedBy        string `json:"issuedBy"`
	IssuedDate      string `json:"issuedDate"`
	ValidFrom       string `json:"validFrom"`
	ValidTo         string `json:"validTo"`
	LatinNumber     string `json:"latinNumber"`
	NameOnDocument  string `json:"nameOnDocument"`
	TypeCode        string `json:"typeCode"`        // passport, id, other
	TypeDescription string `json:"typeDescription"` // National passport, etc.
	Region          string `json:"region"`
	CountryISO      string `json:"countryIso"`
	CountryDesc     string `json:"countryDesc"`
	LogicalID       int    `json:"logicalId"`
	RegLanguage     string `json:"regLanguage"`
	Remark          string `json:"remark"`
	RegType         string `json:"regType"`
	RegOrgType      string `json:"regOrgType"`
	RegPubDate      string `json:"regPubDate"`
	RegEntryDate    string `json:"regEntryDate"`
	RegNumberTitle  string `json:"regNumberTitle"`
	RegProgramme    string `json:"regProgramme"`
	RegPubURL       string `json:"regPubUrl"`
}

// Column indices for all 118 EU CSL CSV columns (0-based)
const (
	// Entity fields (0-15)
	FileGenerationDateIdx          = 0
	EntityLogicalIdx               = 1
	ReferenceNumberIdx             = 2
	EntityUnitedNationIDIdx        = 3
	EntityDesignationDateIdx       = 4
	EntityDesignationDetailsIdx    = 5
	EntityRemarkIdx                = 6
	EntitySubjectTypeIdx           = 7
	EntitySubjectTypeCodeIdx       = 8
	EntityRegulationTypeIdx        = 9
	EntityRegulationOrgTypeIdx     = 10
	EntityRegulationPubDateIdx     = 11
	EntityRegulationEntryDateIdx   = 12
	EntityRegulationNumberTitleIdx = 13
	EntityRegulationProgrammeIdx   = 14
	EntityRegulationPubURLIdx      = 15

	// NameAlias fields (16-33)
	NameAliasLastNameIdx       = 16
	NameAliasFirstNameIdx      = 17
	NameAliasMiddleNameIdx     = 18
	NameAliasWholeNameIdx      = 19
	NameAliasNameLanguageIdx   = 20
	NameAliasGenderIdx         = 21
	NameAliasTitleIdx          = 22
	NameAliasFunctionIdx       = 23
	NameAliasLogicalIDIdx      = 24
	NameAliasRegLanguageIdx    = 25
	NameAliasRemarkIdx         = 26
	NameAliasRegTypeIdx        = 27
	NameAliasRegOrgTypeIdx     = 28
	NameAliasRegPubDateIdx     = 29
	NameAliasRegEntryDateIdx   = 30
	NameAliasRegNumberTitleIdx = 31
	NameAliasRegProgrammeIdx   = 32
	NameAliasRegPubURLIdx      = 33

	// Address fields (34-53)
	AddressCityIdx           = 34
	AddressStreetIdx         = 35
	AddressPoBoxIdx          = 36
	AddressZipCodeIdx        = 37
	AddressRegionIdx         = 38
	AddressPlaceIdx          = 39
	AddressAsAtListingTimeIdx= 40
	AddressContactInfoIdx    = 41
	AddressCountryISOIdx     = 42
	AddressCountryDescIdx    = 43
	AddressLogicalIDIdx      = 44
	AddressRegLanguageIdx    = 45
	AddressRemarkIdx         = 46
	AddressRegTypeIdx        = 47
	AddressRegOrgTypeIdx     = 48
	AddressRegPubDateIdx     = 49
	AddressRegEntryDateIdx   = 50
	AddressRegNumberTitleIdx = 51
	AddressRegProgrammeIdx   = 52
	AddressRegPubURLIdx      = 53

	// BirthDate fields (54-77)
	BirthDateIdx           = 54
	BirthDayIdx            = 55
	BirthMonthIdx          = 56
	BirthYearIdx           = 57
	BirthYearRangeFromIdx  = 58
	BirthYearRangeToIdx    = 59
	BirthCircaIdx          = 60
	BirthCalendarTypeIdx   = 61
	BirthZipCodeIdx        = 62
	BirthRegionIdx         = 63
	BirthPlaceIdx          = 64
	BirthCityIdx           = 65
	BirthCountryISOIdx     = 66
	BirthCountryDescIdx    = 67
	BirthLogicalIDIdx      = 68
	BirthRegLanguageIdx    = 69
	BirthRemarkIdx         = 70
	BirthRegTypeIdx        = 71
	BirthRegOrgTypeIdx     = 72
	BirthRegPubDateIdx     = 73
	BirthRegEntryDateIdx   = 74
	BirthRegNumberTitleIdx = 75
	BirthRegProgrammeIdx   = 76
	BirthRegPubURLIdx      = 77

	// Identification fields (78-104)
	IdentificationNumberIdx      = 78
	IdentificationDiplomaticIdx  = 79
	IdentificationKnownExpiredIdx= 80
	IdentificationKnownFalseIdx  = 81
	IdentificationReportedLostIdx= 82
	IdentificationRevokedIdx     = 83
	IdentificationIssuedByIdx    = 84
	IdentificationIssuedDateIdx  = 85
	IdentificationValidFromIdx   = 86
	IdentificationValidToIdx     = 87
	IdentificationLatinNumberIdx = 88
	IdentificationNameOnDocIdx   = 89
	IdentificationTypeCodeIdx    = 90
	IdentificationTypeDescIdx    = 91
	IdentificationRegionIdx      = 92
	IdentificationCountryISOIdx  = 93
	IdentificationCountryDescIdx = 94
	IdentificationLogicalIDIdx   = 95
	IdentificationRegLanguageIdx = 96
	IdentificationRemarkIdx      = 97
	IdentificationRegTypeIdx     = 98
	IdentificationRegOrgTypeIdx  = 99
	IdentificationRegPubDateIdx  = 100
	IdentificationRegEntryDateIdx= 101
	IdentificationRegNumTitleIdx = 102
	IdentificationRegProgrammeIdx= 103
	IdentificationRegPubURLIdx   = 104

	// Citizenship fields (105-117)
	CitizenshipRegionIdx       = 105
	CitizenshipCountryISOIdx   = 106
	CitizenshipCountryDescIdx  = 107
	CitizenshipLogicalIDIdx    = 108
	CitizenshipRegLanguageIdx  = 109
	CitizenshipRemarkIdx       = 110
	CitizenshipRegTypeIdx      = 111
	CitizenshipRegOrgTypeIdx   = 112
	CitizenshipRegPubDateIdx   = 113
	CitizenshipRegEntryDateIdx = 114
	CitizenshipRegNumTitleIdx  = 115
	CitizenshipRegProgrammeIdx = 116
	CitizenshipRegPubURLIdx    = 117

	// Total expected columns
	TotalCSVColumns = 118
)

// below is the original struct used to parse the document
// fields commented out are not parsed
// this was refactored to be a flatter structure but is left in for documentation
// use the CSLRecord struct above
type CSLRow struct {
	FileGenerationDate string            `json:"fileGenerationDate"`
	Entity             *Entity           `json:"entity"`
	NameAliases        []*NameAlias      `json:"nameAliases"`
	Addresses          []*Address        `json:"addresses"`
	BirthDates         []*BirthDate      `json:"birthDates"`
	Identifications    []*Identification `json:"identifications"`
}

type Entity struct {
	LogicalID       int    `json:"logicalId"`
	ReferenceNumber string `json:"referenceNumber"`
	// UnitiedNationsID   string
	// DesignationDate    string
	// DesignationDetails string
	Remark      string       `json:"remark"`
	SubjectType *SubjectType `json:"subjectType"`
	Regulation  *Regulation  `json:"regulation"`
}
type SubjectType struct {
	// SingleLetter       string
	ClassificationCode string `json:"classificationCode"`
}
type Regulation struct {
	// Type               string
	// OrganizationType   string
	// PublicationDate    string
	// EntryInfoForceDate string
	// NumberTitle        string
	// Programme          string
	PublicationURL string `json:"publicationURL"`
}

type NameAlias struct { // AltNames
	// LastName           string
	// FirstName          string
	// MiddleName         string
	WholeName string `json:"wholeName"`
	// NameLanguage       string
	// Gender             string
	Title string `json:"title"`
	// Function           string
	// LogicalID          int64
	// RegulationLanguage string
	// Remark             string
	// Regulation         *Regulation
}
type Address struct { // addresses
	City    string `json:"city"`
	Street  string `json:"street"`
	PoBox   string `json:"poBox"`
	ZipCode string `json:"zipCode"`
	// Region             string
	// Place              string
	// AsAtListingTime    string
	// ContactInfo        string
	// CountryIso2code    string
	CountryDescription string `json:"countryDescription"`
	// LogicalID          int64
	// RegulationLanguage string
	// Remark             string
	// Regulation         *Regulation
}
type BirthDate struct {
	BirthDate string // keep
	// Day                int64
	// Month              int64
	// Year               int64
	// YearRangeFrom      string
	// YearRangeTo        string
	// Circa              string
	// CaldendarType      string
	// ZipCode            string
	// Region             string
	// Place              string
	City string `json:"city"`
	// CountryIso2code    string
	CountryDescription string `json:"countryDescription"`
	// LogicalID          int64
	// Regulation         *Regulation
}

// Identification is the legacy struct (kept for documentation)
type Identification struct {
	ValidFrom string `json:"validFrom"`
	ValidTo   string `json:"validTo"`
}
