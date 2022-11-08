package csl

// CLS - Consolidated List Sanctions from European Union
// TODO: does this need to be in csl (ask Adam from moov)

// TODO: get this from env
// download uri
// https://webgate.ec.europa.eu/fsd/fsf/public/files/csvFullSanctionsList_1_1/content?token=dG9rZW4tMjAxNw
// protocol: https://
// hostname: webgate.ec.europa.eu
// path: /fsd/fsf/public/files/csvFullSanctionsList_1_1/content
// query: ?token=dG9rZW4tMjAxNw

// struct to hold the rows from the csv data
type EUCSL map[int][]*EUCSLRow

type EUCSLRow struct {
	FileGenerationDate string
	Entity             *Entity
	NameAlias          *NameAlias
	Address            *Address
	BirthDate          *BirthDate
	Identification     *Identification
}

type Entity struct {
	LogicalID          int
	ReferenceNumber    string // EntityNumber
	UnitiedNationsID   string
	DesignationDate    string
	DesignationDetails string
	Remark             string       // Remark - other remarks exist but this one is most pertinent
	SubjectType        *SubjectType // Type
	Regulation         *Regulation
}
type SubjectType struct {
	SingleLetter       string
	ClassificationCode string
}
type Regulation struct {
	Type               string
	OrganizationType   string
	PublicationDate    string
	EntryInfoForceDate string
	NumberTitle        string
	Programme          string
	PublicationURL     string // SourceListURL
}

type NameAlias struct { // AltNames
	LastName           string
	FirstName          string
	MiddleName         string
	WholeName          string // Name
	NameLanguage       string
	Gender             string
	Title              string // Title
	Function           string
	LogicalID          int64
	RegulationLanguage string
	Remark             string
	Regulation         *Regulation
}
type Address struct { // addresses
	City               string // keep
	Street             string // keep
	PoBox              string // keep
	ZipCode            string // keep
	Region             string
	Place              string
	AsAtListingTime    string
	ContactInfo        string
	CountryIso2code    string
	CountryDescription string // keep
	LogicalID          int64
	RegulationLanguage string
	Remark             string
	Regulation         *Regulation
}
type BirthDate struct {
	BirthDate          string // keep
	Day                int64
	Month              int64
	Year               int64
	YearRangeFrom      string
	YearRangeTo        string
	Circa              string
	CaldendarType      string // TODO: this could be an enum
	ZipCode            string
	Region             string
	Place              string
	City               string // keep?
	CountryIso2code    string
	CountryDescription string // Nationality?
	LogicalID          int64
	Regulation         *Regulation
}

type Identification struct {
	Regulation         *Regulation
	Number             int64
	KnownExpired       bool
	KnownFalse         bool
	ReportedLost       bool
	RevokedByIssuer    bool
	LogicalID          int64
	Diplomatic         string // TODO: not sure about this field
	IssuedBy           string
	IssuedDate         string
	ValidFrom          string // StartDate
	ValidTo            string // EndDate
	NameOnDocument     string
	TypeCode           string
	TypeDescription    string
	Region             string
	CountryIso2code    string
	CountryDescription string
	RegulationLanguage string
	Remark             string
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

func NewEUCSLRow() *EUCSLRow {
	row := new(EUCSLRow)
	row.Entity = new(Entity)
	row.Entity.SubjectType = new(SubjectType)
	row.Entity.Regulation = new(Regulation)
	row.NameAlias = new(NameAlias)
	row.Address = new(Address)
	row.BirthDate = new(BirthDate)
	row.Identification = new(Identification)

	return row
}
