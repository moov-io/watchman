package senzing

// SenzingRecord represents a single record in Senzing Entity Resolution format.
// Senzing supports both a FEATURES array format and flat field format.
// This implementation supports both for maximum compatibility.
//
// Reference: https://www.senzing.com/docs/entity_specification/index.html
type SenzingRecord struct {
	// Required fields
	DataSource string `json:"DATA_SOURCE"`
	RecordID   string `json:"RECORD_ID"`

	// FEATURES array format (optional - alternative to flat fields)
	Features []map[string]any `json:"FEATURES,omitempty"`

	// Record type
	RecordType string `json:"RECORD_TYPE,omitempty"` // PERSON or ORGANIZATION

	// Person name fields (flat format)
	NameFirst  string `json:"NAME_FIRST,omitempty"`
	NameMiddle string `json:"NAME_MIDDLE,omitempty"`
	NameLast   string `json:"NAME_LAST,omitempty"`
	NameFull   string `json:"NAME_FULL,omitempty"`
	NamePrefix string `json:"NAME_PREFIX,omitempty"`
	NameSuffix string `json:"NAME_SUFFIX,omitempty"`

	// Organization name
	NameOrg string `json:"NAME_ORG,omitempty"`

	// Address fields
	AddrLine1      string `json:"ADDR_LINE1,omitempty"`
	AddrLine2      string `json:"ADDR_LINE2,omitempty"`
	AddrLine3      string `json:"ADDR_LINE3,omitempty"`
	AddrCity       string `json:"ADDR_CITY,omitempty"`
	AddrState      string `json:"ADDR_STATE,omitempty"`
	AddrPostalCode string `json:"ADDR_POSTAL_CODE,omitempty"`
	AddrCountry    string `json:"ADDR_COUNTRY,omitempty"`
	AddrFull       string `json:"ADDR_FULL,omitempty"`

	// Government identifiers
	SSN                  string `json:"SSN,omitempty"`
	PassportNumber       string `json:"PASSPORT_NUMBER,omitempty"`
	PassportCountry      string `json:"PASSPORT_COUNTRY,omitempty"`
	TaxID                string `json:"TAX_ID_NUMBER,omitempty"`
	TaxIDCountry         string `json:"TAX_ID_COUNTRY,omitempty"`
	NationalID           string `json:"NATIONAL_ID_NUMBER,omitempty"`
	NationalIDCountry    string `json:"NATIONAL_ID_COUNTRY,omitempty"`
	DriversLicenseNumber string `json:"DRIVERS_LICENSE_NUMBER,omitempty"`
	DriversLicenseState  string `json:"DRIVERS_LICENSE_STATE,omitempty"`

	// Contact information
	PhoneNumber string `json:"PHONE_NUMBER,omitempty"`
	PhoneType   string `json:"PHONE_TYPE,omitempty"` // MOBILE, HOME, WORK, etc.
	Email       string `json:"EMAIL_ADDRESS,omitempty"`
	Website     string `json:"WEBSITE_ADDRESS,omitempty"`

	// Dates
	DateOfBirth string `json:"DATE_OF_BIRTH,omitempty"`
	DateOfDeath string `json:"DATE_OF_DEATH,omitempty"`

	// Demographics
	Gender      string `json:"GENDER,omitempty"` // M or F
	Nationality string `json:"NATIONALITY,omitempty"`

	// Relationships
	RelAnchorDomain  string `json:"REL_ANCHOR_DOMAIN,omitempty"`
	RelAnchorKey     string `json:"REL_ANCHOR_KEY,omitempty"`
	RelPointerDomain string `json:"REL_POINTER_DOMAIN,omitempty"`
	RelPointerKey    string `json:"REL_POINTER_KEY,omitempty"`
	RelPointerRole   string `json:"REL_POINTER_ROLE,omitempty"`
}

// Senzing record type constants
const (
	RecordTypePerson       = "PERSON"
	RecordTypeOrganization = "ORGANIZATION"
)

// Feature field name constants used in the FEATURES array
const (
	FieldRecordType        = "RECORD_TYPE"
	FieldNameFirst         = "NAME_FIRST"
	FieldNameMiddle        = "NAME_MIDDLE"
	FieldNameLast          = "NAME_LAST"
	FieldNameFull          = "NAME_FULL"
	FieldNamePrefix        = "NAME_PREFIX"
	FieldNameSuffix        = "NAME_SUFFIX"
	FieldNameOrg           = "NAME_ORG"
	FieldAddrLine1         = "ADDR_LINE1"
	FieldAddrLine2         = "ADDR_LINE2"
	FieldAddrLine3         = "ADDR_LINE3"
	FieldAddrCity          = "ADDR_CITY"
	FieldAddrState         = "ADDR_STATE"
	FieldAddrPostalCode    = "ADDR_POSTAL_CODE"
	FieldAddrCountry       = "ADDR_COUNTRY"
	FieldAddrFull          = "ADDR_FULL"
	FieldSSN               = "SSN"
	FieldPassportNumber    = "PASSPORT_NUMBER"
	FieldPassportCountry   = "PASSPORT_COUNTRY"
	FieldTaxIDNumber       = "TAX_ID_NUMBER"
	FieldTaxIDCountry      = "TAX_ID_COUNTRY"
	FieldNationalIDNumber  = "NATIONAL_ID_NUMBER"
	FieldNationalIDCountry = "NATIONAL_ID_COUNTRY"
	FieldDriversLicNumber  = "DRIVERS_LICENSE_NUMBER"
	FieldDriversLicState   = "DRIVERS_LICENSE_STATE"
	FieldPhoneNumber       = "PHONE_NUMBER"
	FieldPhoneType         = "PHONE_TYPE"
	FieldEmail             = "EMAIL_ADDRESS"
	FieldWebsite           = "WEBSITE_ADDRESS"
	FieldDateOfBirth       = "DATE_OF_BIRTH"
	FieldDateOfDeath       = "DATE_OF_DEATH"
	FieldGender            = "GENDER"
	FieldNationality       = "NATIONALITY"
	FieldRelAnchorDomain   = "REL_ANCHOR_DOMAIN"
	FieldRelAnchorKey      = "REL_ANCHOR_KEY"
	FieldRelPointerDomain  = "REL_POINTER_DOMAIN"
	FieldRelPointerKey     = "REL_POINTER_KEY"
	FieldRelPointerRole    = "REL_POINTER_ROLE"
)

// ExportOptions controls how entities are exported to Senzing format
type ExportOptions struct {
	// DataSource is used when the entity's Source is empty
	DataSource string

	// Format specifies the output format: "jsonl" for JSON Lines or "json" for JSON Array
	Format string

	// Pretty enables indented JSON output
	Pretty bool
}
