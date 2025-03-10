// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/ENHANCED_XML
package ENHANCED_XML

import (
	"encoding/xml"
)

// Element
type SanctionsData struct {
	XMLName xml.Name `xml:"sanctionsData"`

	PublicationInfo SanctionsDataPublicationInfo `xml:"publicationInfo"`

	ReferenceValues SanctionsDataReferenceValues `xml:"referenceValues"`

	FeatureTypes SanctionsDataFeatureTypes `xml:"featureTypes"`

	Entities SanctionsDataEntities `xml:"entities"`
}

// Element
type FiltersSanctionsLists struct {
	XMLName xml.Name `xml:"sanctionsLists"`

	SanctionsList []ReferenceValueReferenceType `xml:",any"`
}

// Element
type FiltersSanctionsPrograms struct {
	XMLName xml.Name `xml:"sanctionsPrograms"`

	SanctionsProgram []ReferenceValueReferenceType `xml:",any"`
}

// Element
type PublicationInfoFilters struct {
	XMLName xml.Name `xml:"filters"`

	SanctionsLists *FiltersSanctionsLists `xml:"sanctionsLists"`

	SanctionsPrograms *FiltersSanctionsPrograms `xml:"sanctionsPrograms"`
}

// Element
type SanctionsDataPublicationInfo struct {
	XMLName xml.Name `xml:"publicationInfo"`

	DataAsOf string `xml:"dataAsOf"`

	Filters PublicationInfoFilters `xml:"filters"`
}

// Element
type ReferenceValuesReferenceValue struct {
	XMLName xml.Name `xml:"referenceValue"`

	RefId int `xml:"refId,attr"`

	Type string `xml:"type"`

	Value string `xml:"value"`

	Code string `xml:"code"`

	IsoCode string `xml:"isoCode"`

	AdvancedXmlOffset *int `xml:"advancedXmlOffset"`
}

// Element
type SanctionsDataReferenceValues struct {
	XMLName xml.Name `xml:"referenceValues"`

	ReferenceValue []ReferenceValuesReferenceValue `xml:",any"`
}

// Element
type FeatureTypesFeatureType struct {
	XMLName xml.Name `xml:"featureType"`

	FeatureTypeId int `xml:"featureTypeId,attr"`

	Type string `xml:"type"`

	DetailType *ReferenceValueReferenceType `xml:"detailType"`

	PublishExclude *bool `xml:"publishExclude"`

	PublishDescription string `xml:"publishDescription"`

	PublishOrder *int `xml:"publishOrder"`
}

// Element
type SanctionsDataFeatureTypes struct {
	XMLName xml.Name `xml:"featureTypes"`

	FeatureType []FeatureTypesFeatureType `xml:",any"`
}

// Element
type EntityGeneralInfo struct {
	XMLName xml.Name `xml:"generalInfo"`

	IdentityId int `xml:"identityId"`

	EntityType ReferenceValueReferenceType `xml:"entityType"`

	LivingStatus *ReferenceValueReferenceType `xml:"livingStatus"`

	Remarks string `xml:"remarks"`

	Title string `xml:"title"`

	IsUsCitizen *bool `xml:"isUsCitizen"`

	IsUsPerson *bool `xml:"isUsPerson"`
}

// Element
type SanctionsListsSanctionsList struct {
	XMLName xml.Name `xml:"sanctionsList"`

	Id string `xml:"id,attr"`

	DatePublished string `xml:"datePublished,attr"`

	RefId int `xml:"refId,attr"`

	Text string `xml:",chardata"`
}

// Element
type EntitySanctionsLists struct {
	XMLName xml.Name `xml:"sanctionsLists"`

	SanctionsList []SanctionsListsSanctionsList `xml:",any"`
}

// Element
type SanctionsProgramsSanctionsProgram struct {
	XMLName xml.Name `xml:"sanctionsProgram"`

	Id string `xml:"id,attr"`

	RefId int `xml:"refId,attr"`

	Text string `xml:",chardata"`
}

// Element
type EntitySanctionsPrograms struct {
	XMLName xml.Name `xml:"sanctionsPrograms"`

	SanctionsProgram []SanctionsProgramsSanctionsProgram `xml:",any"`
}

// Element
type SanctionsTypesSanctionsType struct {
	XMLName xml.Name `xml:"sanctionsType"`

	Id string `xml:"id,attr"`

	RefId int `xml:"refId,attr"`

	Text string `xml:",chardata"`
}

// Element
type EntitySanctionsTypes struct {
	XMLName xml.Name `xml:"sanctionsTypes"`

	SanctionsType []SanctionsTypesSanctionsType `xml:",any"`
}

// Element
type LegalAuthoritiesLegalAuthority struct {
	XMLName xml.Name `xml:"legalAuthority"`

	Id string `xml:"id,attr"`

	RefId int `xml:"refId,attr"`

	Text string `xml:",chardata"`
}

// Element
type EntityLegalAuthorities struct {
	XMLName xml.Name `xml:"legalAuthorities"`

	LegalAuthority []LegalAuthoritiesLegalAuthority `xml:",any"`
}

// Element
type NamePartsNamePart struct {
	XMLName xml.Name `xml:"namePart"`

	Id int `xml:"id,attr"`

	Type ReferenceValueReferenceType `xml:"type"`

	Value string `xml:"value"`
}

// Element
type TranslationNameParts struct {
	XMLName xml.Name `xml:"nameParts"`

	NamePart []NamePartsNamePart `xml:",any"`
}

// Element
type NameTranslationsTranslation struct {
	XMLName xml.Name `xml:"translation"`

	Id int `xml:"id,attr"`

	IsPrimary bool `xml:"isPrimary"`

	Script ReferenceValueReferenceType `xml:"script"`

	FormattedFirstName string `xml:"formattedFirstName"`

	FormattedLastName string `xml:"formattedLastName"`

	FormattedFullName string `xml:"formattedFullName"`

	NameParts TranslationNameParts `xml:"nameParts"`
}

// Element
type NameTranslations struct {
	XMLName xml.Name `xml:"translations"`

	Translation []NameTranslationsTranslation `xml:",any"`
}

// Element
type NamesName struct {
	XMLName xml.Name `xml:"name"`

	Id int `xml:"id,attr"`

	IsPrimary bool `xml:"isPrimary"`

	AliasType *ReferenceValueReferenceType `xml:"aliasType"`

	IsLowQuality bool `xml:"isLowQuality"`

	Translations NameTranslations `xml:"translations"`
}

// Element
type EntityNames struct {
	XMLName xml.Name `xml:"names"`

	Name []NamesName `xml:",any"`
}

// Element
type AddressPartsAddressPart struct {
	XMLName xml.Name `xml:"addressPart"`

	Id int `xml:"id,attr"`

	Type ReferenceValueReferenceType `xml:"type"`

	Value string `xml:"value"`
}

// Element
type TranslationAddressParts struct {
	XMLName xml.Name `xml:"addressParts"`

	AddressPart []AddressPartsAddressPart `xml:",any"`
}

// Element
type AddressTranslationsTranslation struct {
	XMLName xml.Name `xml:"translation"`

	Id int `xml:"id,attr"`

	IsPrimary bool `xml:"isPrimary"`

	Script ReferenceValueReferenceType `xml:"script"`

	AddressParts *TranslationAddressParts `xml:"addressParts"`
}

// Element
type AddressTranslations struct {
	XMLName xml.Name `xml:"translations"`

	Country *ReferenceValueReferenceType `xml:"country"`

	Translation []AddressTranslationsTranslation `xml:"translation"`
}

// Element
type AddressesAddress struct {
	XMLName xml.Name `xml:"address"`

	Id int `xml:"id,attr"`

	Country *ReferenceValueReferenceType `xml:"country"`

	Translations *AddressTranslations `xml:"translations"`
}

// Element
type EntityAddresses struct {
	XMLName xml.Name `xml:"addresses"`

	Address []AddressesAddress `xml:",any"`
}

// Element
type FeatureType struct {
	XMLName xml.Name `xml:"type"`

	FeatureTypeId int `xml:"featureTypeId,attr"`

	Text string `xml:",chardata"`
}

// Element
type FeaturesFeature struct {
	XMLName xml.Name `xml:"feature"`

	Id int `xml:"id,attr"`

	Type FeatureType `xml:"type"`

	VersionId int `xml:"versionId"`

	Value string `xml:"value"`

	ValueRefId *int `xml:"valueRefId"`

	ValueDate *DatePeriodType `xml:"valueDate"`

	IsPrimary bool `xml:"isPrimary"`

	Reliability *ReferenceValueReferenceType `xml:"reliability"`

	Comments string `xml:"comments"`
}

// Element
type EntityFeatures struct {
	XMLName xml.Name `xml:"features"`

	Feature []FeaturesFeature `xml:",any"`
}

// Element
type IdentityDocumentName struct {
	XMLName xml.Name `xml:"name"`

	NameId int `xml:"nameId,attr"`

	NameTranslationId int `xml:"nameTranslationId,attr"`

	Text string `xml:",chardata"`
}

// Element
type IdFeaturesIdFeature struct {
	XMLName xml.Name `xml:"idFeature"`

	Id int `xml:"id,attr"`

	FeatureVersionId int `xml:"featureVersionId,attr"`

	Text string `xml:",chardata"`
}

// Element
type IdentityDocumentIdFeatures struct {
	XMLName xml.Name `xml:"idFeatures"`

	IdFeature []IdFeaturesIdFeature `xml:",any"`
}

// Element
type IdentityDocumentsIdentityDocument struct {
	XMLName xml.Name `xml:"identityDocument"`

	Id int `xml:"id,attr"`

	Type ReferenceValueReferenceType `xml:"type"`

	Name IdentityDocumentName `xml:"name"`

	DocumentNumber string `xml:"documentNumber"`

	IsValid bool `xml:"isValid"`

	IssuingAuthority string `xml:"issuingAuthority"`

	IssuingLocation string `xml:"issuingLocation"`

	IssuingCountry *ReferenceValueReferenceType `xml:"issuingCountry"`

	IssueDate *DatePeriodType `xml:"issueDate"`

	ExpirationDate *DatePeriodType `xml:"expirationDate"`

	Comments string `xml:"comments"`

	IdFeatures *IdentityDocumentIdFeatures `xml:"idFeatures"`
}

// Element
type EntityIdentityDocuments struct {
	XMLName xml.Name `xml:"identityDocuments"`

	IdentityDocument []IdentityDocumentsIdentityDocument `xml:",any"`
}

// Element
type RelationshipRelatedEntity struct {
	XMLName xml.Name `xml:"relatedEntity"`

	EntityId int `xml:"entityId,attr"`

	Text string `xml:",chardata"`
}

// Element
type RelationshipsRelationship struct {
	XMLName xml.Name `xml:"relationship"`

	Id int `xml:"id,attr"`

	Type ReferenceValueReferenceType `xml:"type"`

	RelatedEntity RelationshipRelatedEntity `xml:"relatedEntity"`

	Quality *ReferenceValueReferenceType `xml:"quality"`

	DateRange *DatePeriodType `xml:"dateRange"`

	Comments string `xml:"comments"`
}

// Element
type EntityRelationships struct {
	XMLName xml.Name `xml:"relationships"`

	Relationship []RelationshipsRelationship `xml:",any"`
}

// Element
type EntitiesEntity struct {
	XMLName xml.Name `xml:"entity"`

	Id int `xml:"id,attr"`

	GeneralInfo EntityGeneralInfo `xml:"generalInfo"`

	SanctionsLists EntitySanctionsLists `xml:"sanctionsLists"`

	SanctionsPrograms EntitySanctionsPrograms `xml:"sanctionsPrograms"`

	SanctionsTypes EntitySanctionsTypes `xml:"sanctionsTypes"`

	LegalAuthorities EntityLegalAuthorities `xml:"legalAuthorities"`

	Names EntityNames `xml:"names"`

	Addresses *EntityAddresses `xml:"addresses"`

	Features *EntityFeatures `xml:"features"`

	IdentityDocuments *EntityIdentityDocuments `xml:"identityDocuments"`

	Relationships *EntityRelationships `xml:"relationships"`
}

// Element
type SanctionsDataEntities struct {
	XMLName xml.Name `xml:"entities"`

	Entity []EntitiesEntity `xml:",any"`
}

// XSD ComplexType declarations

type DatePeriodType struct {
	XMLName xml.Name

	Id int `xml:"id,attr"`

	FromDateBegin string `xml:"fromDateBegin"`

	FromDateEnd string `xml:"fromDateEnd"`

	ToDateBegin string `xml:"toDateBegin"`

	ToDateEnd string `xml:"toDateEnd"`

	IsApproximate bool `xml:"isApproximate"`

	IsDateRange bool `xml:"isDateRange"`
}

type ReferenceValueReferenceType struct {
	XMLName xml.Name

	RefId int `xml:"refId,attr"`

	Text string `xml:",chardata"`
}

// XSD SimpleType declarations
