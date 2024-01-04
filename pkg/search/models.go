package search

import "time"

type Entity[T any] struct {
	Name   string     `json:"name"`
	Type   EntityType `json:"entityType"`
	Source SourceList `json:"sourceList"`

	// TODO(adam): What has opensanctions done to normalize and join this data
	// Review https://www.opensanctions.org/reference/

	Person       *Person       `json:"person"`
	Business     *Business     `json:"business"`
	Organization *Organization `json:"organization"`
	Aircraft     *Aircraft     `json:"aircraft"`
	Vessel       *Vessel       `json:"vessel"`

	CryptoAddresses []CryptoAddress `json:"cryptoAddresses"`

	Addresses []Address `json:"addresses"`

	SourceData T `json:"sourceData"` // Contains all original list data with source list naming
}

type EntityType string

var (
	EntityPerson        EntityType = "person"
	EntityBusiness      EntityType = "business"
	EntityAircraft      EntityType = "aircraft"
	EntityVessel        EntityType = "vessel"
	EntityCryptoAddress EntityType = "crypto-address"
)

type SourceList string

var (
	SourceEUCSL  SourceList = "eu_csl"
	SourceUKCSL  SourceList = "uk_csl"
	SourceUSCSL  SourceList = "us_csl"
	SourceUSOFAC SourceList = "us_ofac"
)

type Person struct {
	Name      string     `json:"name"`
	Gender    Gender     `json:"gender"`
	BirthDate *time.Time `json:"birthDate"`
	DeathDate *time.Time `json:"deathDate"`

	GovernmentIDs []GovernmentID `json:"governmentIDs"`
}

type Gender string

var (
	GenderUnknown Gender = "unknown"
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
)

type GovernmentID struct {
	Type       GovernmentIDType `json:"type"`
	Identifier string           `json:"identifier"`
}

type GovernmentIDType string

var (
	GovernmentIDPassport GovernmentIDType = "passport"
)

type Business struct {
	Name       string       `json:"name"`
	Created    *time.Time   `json:"created"`
	Dissolved  *time.Time   `json:"dissolved"`
	Identifier []Identifier `json:"identifier"`
}

// Identifier
//
// TODO(adam): Look at OpenSanctions for tax ID codes
// https://www.opensanctions.org/reference/#schema.Company
type Identifier struct {
	Type       IdentifierType `json:"type"`
	Identifier string         `json:"value"`
}

type IdentifierType string

var (
	Identifier_US_EIN IdentifierType = "us_ein"
	Identifier_US_SSN IdentifierType = "us_ssn"
)

// Organization
//
// TODO(adam): https://www.opensanctions.org/reference/#schema.Organization
type Organization struct {
	Name       string       `json:"name"`
	Created    *time.Time   `json:"created"`
	Dissolved  *time.Time   `json:"dissolved"`
	Identifier []Identifier `json:"identifier"`
}

type Aircraft struct {
	Name         string       `json:"name"`
	Type         AircraftType `json:"type"`
	Flag         string       `json:"flag"` // ISO-3166
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
	Name      string     `json:"name"`
	IMONumber string     `json:"imoNumber"`
	Type      VesselType `json:"type"`
	Flag      string     `json:"flag"` // ISO-3166
	Built     *time.Time `json:"built"`
	Model     string     `json:"model"`
	Tonnage   int        `json:"tonnage"`
	MMSI      string     `json:"mmsi"` // Maritime Mobile Service Identity
}

type VesselType string

var (
	VesselTypeUnknown VesselType = "unknown"
	VesselTypeCargo   VesselType = "cargo"
)

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
