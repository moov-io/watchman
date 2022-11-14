package csl

// CLS - Consolidated List Sanctions from European Union

// struct to hold the rows from the csv data before merge
type EUCSL map[int]*EUCSLRecord

type EUCSLRecord struct {
	FileGenerationDate         string            `json:"fileGenerationDate"`
	EntityLogicalID            int               `json:"entityLogicalId"`
	EntityRemark               string            `json:"entityRemark"`
	EntitySubjectType          string            `json:"entitySubjectType"`
	EntityPublicationURL       string            `json:"entityPublicationURL"`
	EntityReferenceNumber      string            `json:"entityReferenceNumber"`
	NameAliasWholeNames        []string          `json:"nameAliasWholeNames"`
	NameAliasTitles            []string          `json:"nameAliasTitles"`
	AddressCities              []string          `json:"addressCities"`
	AddressStreets             []string          `json:"addressStreets"`
	AddressPoBoxes             []string          `json:"addressPoBoxs"`
	AddressZipCodes            []string          `json:"addressZipCodes"`
	AddressCountryDescriptions []string          `json:"addressCountryDescriptions"`
	BirthDates                 []string          `json:"birthDates"`
	BirthCities                []string          `json:"birthCities"`
	BirthCountries             []string          `json:"birthCountries"`
	ValidFromTo                map[string]string `json:"validFromTo"`
}

// header indicies
const (
	FileGenerationDateIdx             = 0
	EntityLogicalIdx                  = 1
	ReferenceNumberIdx                = 2
	EntityRemarkIdx                   = 6
	EntitySubjectTypeIdx              = 8
	EntityRegulationPublicationURLIdx = 15

	NameAliasWholeNameIdx = 19
	NameAliasTitleIdx     = 22

	AddressCityIdx               = 34
	AddressStreetIdx             = 35
	AddressPoBoxIdx              = 36
	AddressZipCodeIdx            = 37
	AddressCountryDescriptionIdx = 43

	BirthDateIdx        = 54
	BirthDateCityIdx    = 65
	BirthDateCountryIdx = 67

	IdentificationValidFromIdx = 86
	IdentificationValidToIdx   = 87
)

// below is the original struct used to parse the document
// fields commented out are not parsed
// this was refactored to be a flatter structure but is left in for documentation
// use the EUCSLRecord struct above
type EUCSLRow struct {
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

type Identification struct {
	// Regulation         *Regulation
	// Number             int64
	// KnownExpired       bool
	// KnownFalse         bool
	// ReportedLost       bool
	// RevokedByIssuer    bool
	// LogicalID          int64
	// Diplomatic         string
	// IssuedBy           string
	// IssuedDate         string
	ValidFrom string `json:"validFrom"`
	ValidTo   string `json:"validTo"`
	// NameOnDocument     string
	// TypeCode           string
	// TypeDescription    string
	// Region             string
	// CountryIso2code    string
	// CountryDescription string
	// RegulationLanguage string
	// Remark             string
}
