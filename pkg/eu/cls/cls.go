package eucls

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
type EUCLS map[string]EUCLSRow

type EUCLSRow struct {
	FileGenerationDate string
	Entity             *Entity
	Address            *Address
	BirthDate          *BirthDate
	Identification     *Identification
}

type Entity struct {
	LogicalID          int64
	ReferenceNumber    string
	UnitiedNationsID   string
	DesignationDate    string
	DesignationDetails string
	Remark             string
	SubjectType        *SubjectType
	Regulation         *Regulation
}
type SubjectType struct {
	SingleLetter       string
	ClassificationCode string
}
type Regulation struct {
	Type               string
	Organization       string
	PublicationDate    string
	EntryInfoForceDate string
	NumberTitle        string
	Programme          string
	PublicationURL     string
}

type NameAlias struct {
	LastName           string
	FirstName          string
	MiddleName         string
	WholeName          string
	NameLanguage       string
	Gender             string
	Title              string
	Function           string
	LogicalID          int64
	RegulationLanguage string
	Remark             string
	Regulation         *Regulation
}
type Address struct {
	City               string
	Street             string
	PoBox              string
	ZipCode            string
	Region             string
	Place              string
	AsAtListingTime    string
	ContactInfo        string
	CountryIso2code    string
	CountryDescription string
	LogicalID          int64
	RegulationLanguage string
	Remark             string
	Regulation         *Regulation
}
type BirthDate struct {
	BirthDate          string
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
	City               string
	CountryIso2code    string
	CountryDescription string
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
	ValidFrom          string
	ValidTo            string
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
const ()

// TODO: need function to merge the like rows based on the Entity_LogicalId
func mergeRowsByLogicalID(rows []EUCLSRow) EUCLS {
	cls := make(EUCLS)
	// using a map for constant lookups (if possible)
	// append to the map every time a new lookup for entity logical id is found
	// otherwise

	return cls
}
