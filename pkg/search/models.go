package search

import (
	"fmt"
	"time"
)

type Value interface{}

type Entity[T Value] struct {
	Name   string     `json:"name"`
	Type   EntityType `json:"entityType"`
	Source SourceList `json:"sourceList"`

	// SourceID is the source data's identifier.
	SourceID string `json:"sourceID"`

	// TODO(adam): What has opensanctions done to normalize and join this data
	// Review https://www.opensanctions.org/reference/

	Person       *Person       `json:"person"`
	Business     *Business     `json:"business"`
	Organization *Organization `json:"organization"`
	Aircraft     *Aircraft     `json:"aircraft"`
	Vessel       *Vessel       `json:"vessel"`

	Contact   ContactInfo `json:"contact"`
	Addresses []Address   `json:"addresses"`

	CryptoAddresses []CryptoAddress `json:"cryptoAddresses"`

	Affiliations   []Affiliation    `json:"affiliations"`
	SanctionsInfo  *SanctionsInfo   `json:"sanctionsInfo"`
	HistoricalInfo []HistoricalInfo `json:"historicalInfo"`

	SourceData T `json:"sourceData"` // Contains all original list data with source list naming
}

type EntityType string

var (
	EntityPerson       EntityType = "person"
	EntityBusiness     EntityType = "business"
	EntityOrganization EntityType = "organization"
	EntityAircraft     EntityType = "aircraft"
	EntityVessel       EntityType = "vessel"
)

type SourceList string

var (
	SourceAPIRequest SourceList = "api-request"

	SourceEUCSL  SourceList = "eu_csl"
	SourceUKCSL  SourceList = "uk_csl"
	SourceUSCSL  SourceList = "us_csl"
	SourceUSOFAC SourceList = "us_ofac"
)

type Person struct {
	Name      string     `json:"name"`
	AltNames  []string   `json:"altNames"`
	Gender    Gender     `json:"gender"`
	BirthDate *time.Time `json:"birthDate"`
	DeathDate *time.Time `json:"deathDate"`
	Titles    []string   `json:"titles"`

	GovernmentIDs []GovernmentID `json:"governmentIDs"`
}

type Gender string

var (
	GenderUnknown Gender = "unknown"
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
)

type ContactInfo struct {
	EmailAddresses []string `json:"emailAddresses"`
	PhoneNumbers   []string `json:"phoneNumbers"`
	FaxNumbers     []string `json:"faxNumbers"`
	Websites       []string `json:"websites"`
}

// TODO(adam):
//
// Website www.tidewaterco.com;
// Email Address info@tidewaterco.com;    alt. Email Address info@tidewaterco.ir;
// Telephone: 982188553321; Alt. Telephone: 982188554432;
// Fax: 982188717367; Alt. Fax: 982188708761;
//
// 12803,"TIDEWATER MIDDLE EAST CO.",-0- ,"SDGT] [NPWMD] [IRGC] [IFSR] [IFCA",-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,-0- ,"  alt. Email Address info@tidewaterco.ir; IFCA Determination - Port Operator; Additional Sanctions Information - Subject to Secondary Sanctions; Business Registration Document # 18745 (Iran);   Alt. Fax: 982188708911."

type GovernmentID struct {
	Type       GovernmentIDType `json:"type"`
	Country    string           `json:"country"` // ISO-3166 // TODO(adam):
	Identifier string           `json:"identifier"`
}

type GovernmentIDType string

var (
	GovernmentIDPassport            GovernmentIDType = "passport"
	GovernmentIDDriversLicense      GovernmentIDType = "drivers-license"
	GovernmentIDNational            GovernmentIDType = "national-id"
	GovernmentIDTax                 GovernmentIDType = "tax-id"
	GovernmentIDSSN                 GovernmentIDType = "ssn"
	GovernmentIDCedula              GovernmentIDType = "cedula"
	GovernmentIDCURP                GovernmentIDType = "curp"
	GovernmentIDCUIT                GovernmentIDType = "cuit"
	GovernmentIDElectoral           GovernmentIDType = "electoral"
	GovernmentIDBusinessRegisration GovernmentIDType = "business-registration"
	GovernmentIDCommercialRegistry  GovernmentIDType = "commercial-registry"
	GovernmentIDBirthCert           GovernmentIDType = "birth-certificate"
	GovernmentIDRefugee             GovernmentIDType = "refugee-id"
	GovernmentIDDiplomaticPass      GovernmentIDType = "diplomatic-passport"
	GovernmentIDPersonalID          GovernmentIDType = "personal-id"
)

type Business struct {
	Name          string         `json:"name"`
	AltNames      []string       `json:"altNames"`
	Created       *time.Time     `json:"created"`
	Dissolved     *time.Time     `json:"dissolved"`
	GovernmentIDs []GovernmentID `json:"governmentIDs"`
}

// Organization
//
// TODO(adam): https://www.opensanctions.org/reference/#schema.Organization
type Organization struct {
	Name          string         `json:"name"`
	AltNames      []string       `json:"altNames"`
	Created       *time.Time     `json:"created"`
	Dissolved     *time.Time     `json:"dissolved"`
	GovernmentIDs []GovernmentID `json:"governmentIDs"`
}

type Aircraft struct {
	Name         string       `json:"name"`
	AltNames     []string     `json:"altNames"`
	Type         AircraftType `json:"type"`
	Flag         string       `json:"flag"` // ISO-3166 // TODO(adam):
	Built        *time.Time   `json:"built"`
	ICAOCode     string       `json:"icaoCode"` // ICAO aircraft type designator
	Model        string       `json:"model"`
	SerialNumber string       `json:"serialNumber"`
}

type AircraftType string

var (
	AircraftTypeUnknown AircraftType = "unknown"
	AircraftCargo       AircraftType = "cargo"
)

// Vessel
//
// TODO(adam): https://www.opensanctions.org/reference/#schema.Vessel
type Vessel struct {
	Name                   string     `json:"name"`
	AltNames               []string   `json:"altNames"`
	IMONumber              string     `json:"imoNumber"`
	Type                   VesselType `json:"type"`
	Flag                   string     `json:"flag"` // ISO-3166 // TODO(adam):
	Built                  *time.Time `json:"built"`
	Model                  string     `json:"model"`
	Tonnage                int        `json:"tonnage"`
	MMSI                   string     `json:"mmsi"` // Maritime Mobile Service Identity
	CallSign               string     `json:"callSign"`
	GrossRegisteredTonnage int        `json:"grossRegisteredTonnage"`
	Owner                  string     `json:"owner"`
}

type VesselType string

var (
	VesselTypeUnknown VesselType = "unknown"
	VesselTypeCargo   VesselType = "cargo"
)

// CryptoAddress
//
// &cryptoAddress=XBT:x123456
type CryptoAddress struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
}

// Address is a struct which represents any physical location
//
// TODO(adam): Should probably adopt something like libpostal's naming
// https://github.com/openvenues/libpostal?tab=readme-ov-file#parser-labels
//
// Or OpenSanctions
// https://www.opensanctions.org/reference/#schema.Address
type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
	Country    string `json:"country"` // ISO-3166 code

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (a Address) Format() string {
	return fmt.Sprintf("%s %s %s %s %s %s", a.Line1, a.Line2, a.City, a.PostalCode, a.State, a.Country)
}

type Affiliation struct {
	EntityName string `json:"entityName"`
	Type       string `json:"type"` // e.g., "Linked To", "Subsidiary Of", "Owned By"
	Details    string `json:"details,omitempty"`
}

type SanctionsInfo struct {
	Programs    []string `json:"programs"`    // e.g., "SDGT", "IRGC"
	Secondary   bool     `json:"secondary"`   // Subject to secondary sanctions
	Description string   `json:"description"` // Additional details
}

type HistoricalInfo struct {
	Type  string    `json:"type"`  // e.g., "Former Name", "Previous Flag"
	Value string    `json:"value"` // The historical value
	Date  time.Time `json:"date,omitempty"`
}
