// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"cmp"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/watchman/internal/indices"
	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/pkg/address"
	"github.com/moov-io/watchman/pkg/search"
)

// Regular expressions for parsing various fields
var (
	akaRegex = regexp.MustCompile(`(?i)a\.k\.a\.\s+'([^']+)'`)
	// Matches both "citizen Venezuela" and "nationality: Russia"
	citizenshipRegex = regexp.MustCompile(`(?i)(citizen|nationality)[:\s]+([^;,]+)`)
	// Matches both "POB Baghdad, Iraq" and "Alt. POB: Keren Eritrea"
	pobRegex = regexp.MustCompile(`(?i)(?:Alt\.)?\s*POB:?\s+([^;]+)`)
	// Contact information patterns
	emailRegex = regexp.MustCompile(`(?i)(?:Email|EMAIL)[:\s](Address)?[:\s]+([^;]+)`)
	phoneRegex = regexp.MustCompile(`(?i)(?:Telephone|Phone|PHONE)[:\s](No.?)?[:\s]+([^;]+)`)
	faxRegex   = regexp.MustCompile(`(?i)Fax[:\s](No.?)?[:\s]+([^;]+)`)
	// Website patterns
	websiteRegex = regexp.MustCompile(`(?i)(?:Website|http)[:\s]+([^;\s]+)`)

	// Identifier Regex
	identifierRegex = regexp.MustCompile(`(?i)%s[\s.:]+([^\s](?:[^;()]*[^\s;()])?)(?:\s*\(([^)]+)\))?`)

	// Country extraction pattern
	countryParenRegex = regexp.MustCompile(`\(([\w\s]+)\)`)
)

var (
	dobPatterns = []string{"02 Jan 2006", "Jan 2006", "2006"}
)

func findDateStamp(matchingRemarks []remark) *time.Time {
	return withFirstP(matchingRemarks, func(in remark) *time.Time {
		t, err := parseTime(dobPatterns, in.value)
		if err != nil {
			return nil
		}
		return &t
	})
}

func parseTime(acceptedLayouts []string, value string) (time.Time, error) {
	value = strings.TrimSpace(strings.ReplaceAll(value, "circa", ""))
	parts := strings.Split(value, "to")
	if len(parts) > 1 {
		value = parts[0]
	} else {
		parts = strings.Split(value, "-")
		if len(parts) > 1 {
			value = parts[0]
		}
	}
	value = strings.TrimSpace(value)

	for i := range acceptedLayouts {
		tt, err := time.Parse(acceptedLayouts[i], value)
		if !invalidDate(tt) && err == nil {
			return tt, nil
		}
	}
	return time.Time{}, nil
}

func extractCountry(remark string) string {
	matches := countryParenRegex.FindStringSubmatch(remark)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func GroupIntoEntities(sdns []SDN, addresses map[string][]Address, comments map[string][]SDNComments, altIds map[string][]AlternateIdentity) []search.Entity[search.Value] {
	fn := func(sdn SDN) search.Entity[search.Value] {
		return ToEntity(sdn, addresses[sdn.EntityID], comments[sdn.EntityID], altIds[sdn.EntityID])
	}

	groups := runtime.NumCPU()
	return indices.ProcessSlice(sdns, groups, fn)
}

func ToEntity(sdn SDN, addresses []Address, comments []SDNComments, altIds []AlternateIdentity) search.Entity[search.Value] {
	out := search.Entity[search.Value]{
		Name:       prepare.ReorderSDNName(sdn.SDNName, sdn.SDNType),
		Source:     search.SourceUSOFAC,
		SourceData: sdn,
		SourceID:   sdn.EntityID,
	}

	remarks := splitRemarks(sdn.Remarks)
	affiliations, sanctionsInfo, historicalInfo, titles := parseRemarks(remarks)

	out.Affiliations = affiliations
	out.SanctionsInfo = sanctionsInfo
	out.HistoricalInfo = historicalInfo
	out.CryptoAddresses = parseCryptoAddresses(remarks)

	// Extract common fields regardless of entity type
	out.Contact = parseContactInfo(remarks)
	out.Addresses = parseAddresses(addresses)

	// Get all alternate names from both remarks and AlternateIdentity entries
	var altNames []string
	altNames = append(altNames, parseAltNames(remarks)...)
	altNames = append(altNames, parseAltIdentities(altIds)...)
	altNames = deduplicateStrings(altNames)
	altNames = prepare.ReorderSDNNames(altNames, sdn.SDNType)

	switch strings.ToLower(strings.TrimSpace(sdn.SDNType)) {
	case "-0-", "":
		out.Type = search.EntityBusiness
		out.Business = &search.Business{
			Name:     prepare.RemoveCompanyTitles(sdn.SDNName),
			AltNames: altNames,
		}
		out.Business.Created = findDateStamp(findMatchingRemarks(remarks, "Organization Established Date"))
		// out.Business.Dissolved = findDateStamp(findMatchingRemarks(remarks, "TODO(adam)"))

		out.Business.GovernmentIDs = parseGovernmentIDs(remarks)

		// out.Business.Identifiers = makeIdentifiers(remarks, []string{
		// 	"Branch Unit Number",
		// 	"Business Number",
		// 	"Business Registration Document",
		// 	"Business Registration Number",
		// 	"Certificate of Incorporation Number",
		// 	"Chamber of Commerce Number",
		// 	"Chinese Commercial Code",
		// 	"Commercial Registry Number",
		// 	"Company Number",
		// 	"Company ID",                              // new: e.g., "Company ID: No. 59 531..."
		// 	"D-U-N-S Number",                          // new: e.g., "D-U-N-S Number 33-843-5672"
		// 	"Dubai Chamber of Commerce Membership No", // new
		// 	"Enterprise Number",
		// 	"Fiscal Code",        // new: business tax identifiers
		// 	"Folio Mercantil No", // new: Mexican business registration
		// 	"Legal Entity Number",
		// 	"Matricula Mercantil No",     // new: Colombian business registration
		// 	"Public Registration Number", // new
		// 	"Registration Number",
		// 	"RIF",                                   // new: Venezuelan tax ID
		// 	"RUC",                                   // new: Panama business registration
		// 	"Romanian C.R",                          // new: Romanian Commercial Registry
		// 	"Tax ID No.",                            // new: Important business identifier
		// 	"Trade License No",                      // new
		// 	"UK Company Number",                     // new: Specific UK format
		// 	"US FEIN",                               // new: US Federal Employer ID Number
		// 	"United Social Credit Code Certificate", // new: Chinese business ID
		// 	"V.A.T. Number",                         // new: VAT registration numbers
		// })

	case "individual":
		out.Type = search.EntityPerson
		out.Person = &search.Person{
			Name:     out.Name,
			AltNames: altNames,
			Gender:   search.Gender(strings.ToLower(firstValue(findMatchingRemarks(remarks, "Gender")))),
		}

		// Title from SDN field needs to be prepended if non-empty
		if sdn.Title != "" {
			titles = append([]string{sdn.Title}, titles...)
		}
		out.Person.Titles = titles

		// Handle birth date
		out.Person.BirthDate = withFirstP(findMatchingRemarks(remarks, "DOB"), func(in remark) *time.Time {
			t, err := parseTime(dobPatterns, in.value)
			if err != nil {
				return nil
			}
			return &t
		})

		// Parse government IDs
		out.Person.GovernmentIDs = parseGovernmentIDs(remarks)

	case "vessel":
		out.Type = search.EntityVessel
		out.Vessel = &search.Vessel{
			Name:                   sdn.SDNName,
			AltNames:               altNames,
			IMONumber:              firstValue(findMatchingRemarks(remarks, "IMO")),
			Type:                   normalizeVesselType(cmp.Or(sdn.VesselType, firstValue(findMatchingRemarks(remarks, "Vessel Type")))),
			Flag:                   normalizeCountryCode(cmp.Or(sdn.VesselFlag, firstValue(findMatchingRemarks(remarks, "Flag")))),
			MMSI:                   firstValue(findMatchingRemarks(remarks, "MMSI")),
			Tonnage:                parseTonnage(cmp.Or(sdn.Tonnage, firstValue(findMatchingRemarks(remarks, "Tonnage")))),
			CallSign:               sdn.CallSign,
			GrossRegisteredTonnage: parseTonnage(sdn.GrossRegisteredTonnage),
			Owner:                  sdn.VesselOwner,
		}

	case "aircraft":
		out.Type = search.EntityAircraft
		out.Aircraft = &search.Aircraft{
			Name:     sdn.SDNName,
			AltNames: altNames,
			Type:     normalizeAircraftType(firstValue(findMatchingRemarks(remarks, "Aircraft Type"))),
			Flag:     normalizeCountryCode(firstValue(findMatchingRemarks(remarks, "Flag"))),
			Built: withFirstP(findMatchingRemarks(remarks, "Manufacture Date"), func(in remark) *time.Time {
				t, err := parseTime(dobPatterns, in.value)
				if err != nil {
					return nil
				}
				return &t
			}),
			Model:        firstValue(findMatchingRemarks(remarks, "Aircraft Model")),
			SerialNumber: parseSerialNumber(remarks),
			ICAOCode:     firstValue(findMatchingRemarks(remarks, "ICAO Code")),
		}
	}

	return out
}

func parseAddresses(inputs []Address) []search.Address {
	var out []search.Address
	for i := range inputs {
		input := fmt.Sprintf("%s %s %s", inputs[i].Address, inputs[i].CityStateProvincePostalCode, inputs[i].Country)
		addr := address.ParseAddress(input)
		if addr.Line1 != "" {
			out = append(out, addr)
		}
	}
	return out
}

func parseAltNames(remarks []string) []string {
	var names []string
	for _, r := range remarks {
		matches := akaRegex.FindAllStringSubmatch(r, -1)
		for _, m := range matches {
			if len(m) > 1 {
				name := strings.Trim(m[1], "'")
				names = append(names, name)
			}
		}
	}
	return names
}

func parseAltIdentities(altIds []AlternateIdentity) []string {
	var names []string
	for _, alt := range altIds {
		if alt.AlternateName != "" {
			names = append(names, strings.TrimSpace(alt.AlternateName))
		}
	}
	return names
}

func deduplicateStrings(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, str := range input {
		// Normalize the string for comparison
		normalized := strings.ToLower(strings.TrimSpace(str))
		if !seen[normalized] {
			seen[normalized] = true
			result = append(result, str) // Keep original string formatting
		}
	}
	return result
}

var (
	// Passports
	governmentIDPassportRegex       = regexp.MustCompile(`(?i)Passport\s+(?:#|No\.|Number)?\s*([A-Z0-9]+)`)
	governmentIDDiplomaticPassRegex = regexp.MustCompile(`(?i)Diplomatic\s+Passport\s+([A-Z0-9]+)`)

	// Drivers Licenses
	governmentIDDriversLicenseRegex = regexp.MustCompile(`(?i)Driver'?s?\s+License\s+(?:No\.|Number)?\s*(?:[A-Z]-)?([A-Z0-9]+)`)

	// National IDs
	governmentIDNationalRegex   = regexp.MustCompile(`(?i)National\s+ID\s+(?:No\.|Number)?\s*([A-Z0-9-]+)`)
	governmentIDPersonalIDRegex = regexp.MustCompile(`(?i)Personal\s+ID\s+(?:Card)?\s*(?:No\.|Number)?\s*([A-Z0-9-]+)`)

	// Tax IDs
	governmentIDTaxRegex  = regexp.MustCompile(`(?i)Tax\s+ID\s+(?:No\.|Number)?\s*([A-Z0-9-]+)`)
	governmentIDCUITRegex = regexp.MustCompile(`(?i)C\.?U\.?I\.?T\.?\s+(?:No\.|Number)?\s*([A-Z0-9-]+)`)

	// Social Security Numbers
	governmentIDSSNRegex = regexp.MustCompile(`(?i)SSN\s+([0-9-]+)`)

	// Latin American IDs
	governmentIDCedulaRegex = regexp.MustCompile(`(?i)Cedula\s+(?:No\.|Number)?\s*([A-Z0-9-]+)`)
	governmentIDCURPRegex   = regexp.MustCompile(`(?i)C\.?U\.?R\.?P\.?\s+(?:#|No\.|Number)?\s*([A-Z0-9-]+)`)

	// Electoral IDs
	governmentIDElectoralRegex = regexp.MustCompile(`(?i)Electoral\s+Registry\s+(?:No\.|Number)?\s*([A-Z0-9-]+)`)

	// Business Registration
	governmentIDBusinessRegistrationRegex = regexp.MustCompile(`(?i)Business\s+Registration\s+(?:No\.|Number|Document)?\s*([A-Z0-9-\.]+)`)
	governmentIDCompanyNumberRegex        = regexp.MustCompile(`(?i)Company\s+Number\s+([0-9]+)`)
	governmentIDLegalEntityNumberRegex    = regexp.MustCompile(`(?i)Legal\s+Entity\s+Number\s+([A-Za-z0-9\-\.]+)`)
	governmentIDCommercialRegistryRegex   = regexp.MustCompile(`(?i)Commercial\s+Registry\s+(?:No\.|Number)?\s*([A-Z0-9-./]+)`)

	// Birth Certificates
	governmentIDBirthCertRegex = regexp.MustCompile(`(?i)Birth\s+Certificate\s+(?:No\.|Number)?\s*([A-Z0-9]+)`)

	// Refugee Documents
	governmentIDRefugeeRegex = regexp.MustCompile(`(?i)Refugee\s+ID\s+(?:Card)?\s*([A-Z0-9]+)`)
)

func parseGovernmentIDs(remarks []string) []search.GovernmentID {
	var ids []search.GovernmentID

	// Map of regex patterns to GovernmentIDType
	idPatterns := map[*regexp.Regexp]search.GovernmentIDType{
		governmentIDPassportRegex:             search.GovernmentIDPassport,
		governmentIDDriversLicenseRegex:       search.GovernmentIDDriversLicense,
		governmentIDPassportRegex:             search.GovernmentIDPassport,
		governmentIDDiplomaticPassRegex:       search.GovernmentIDDiplomaticPass,
		governmentIDDriversLicenseRegex:       search.GovernmentIDDriversLicense,
		governmentIDNationalRegex:             search.GovernmentIDNational,
		governmentIDPersonalIDRegex:           search.GovernmentIDPersonalID,
		governmentIDTaxRegex:                  search.GovernmentIDTax,
		governmentIDCUITRegex:                 search.GovernmentIDCUIT,
		governmentIDSSNRegex:                  search.GovernmentIDSSN,
		governmentIDCedulaRegex:               search.GovernmentIDCedula,
		governmentIDCURPRegex:                 search.GovernmentIDCURP,
		governmentIDElectoralRegex:            search.GovernmentIDElectoral,
		governmentIDBusinessRegistrationRegex: search.GovernmentIDBusinessRegisration,
		governmentIDCompanyNumberRegex:        search.GovernmentIDBusinessRegisration,
		governmentIDLegalEntityNumberRegex:    search.GovernmentIDBusinessRegisration,
		governmentIDCommercialRegistryRegex:   search.GovernmentIDCommercialRegistry,
		governmentIDBirthCertRegex:            search.GovernmentIDBirthCert,
		governmentIDRefugeeRegex:              search.GovernmentIDRefugee,
	}

	for _, r := range remarks {
		// Extract country first
		countryRaw := extractCountry(r)
		country := normalizeCountry(countryRaw)

		// Remove the country from the remark for cleaner ID extraction
		remarkWithoutCountry := r
		if countryRaw != "" {
			remarkWithoutCountry = strings.TrimSpace(strings.ReplaceAll(r, "("+countryRaw+")", ""))
		}

		for re, idType := range idPatterns {
			if matches := re.FindStringSubmatch(remarkWithoutCountry); len(matches) > 1 {
				identifier := strings.TrimRight(matches[1], ".;,")

				ids = append(ids, search.GovernmentID{
					Type:       idType,
					Country:    country, // Use the extracted and normalized country
					Identifier: identifier,
				})
			}
		}
	}

	return ids
}

func normalizeCountry(country string) string {
	// Mapping of common country name variations to standard names
	countryMap := map[string]string{
		"USA":    "United States",
		"U.S.A.": "United States",
		"US":     "United States",
		"U.S.":   "United States",
		"UK":     "United Kingdom",
		"U.K.":   "United Kingdom",
		"UAE":    "United Arab Emirates",
		"ROK":    "South Korea",
		"DPRK":   "North Korea",
		"PRC":    "China",
		"ROC":    "Taiwan",
		"россия": "Russia",
		"РОССИЯ": "Russia",
		"中国":     "China",
		"日本":     "Japan",
		"한국":     "South Korea",
		"España": "Spain",
		"ESPAÑA": "Spain",
	}

	// First try direct mapping
	if normalized, exists := countryMap[strings.ToUpper(strings.TrimSpace(country))]; exists {
		return normalized
	}

	// If no direct mapping, return original (could be extended with more sophisticated matching)
	return country
}

func normalizeCountryCode(country string) string {
	// TODO: Implement conversion to ISO-3166
	return strings.TrimSpace(country)
}

func normalizeVesselType(vesselType string) search.VesselType {
	switch strings.ToLower(strings.TrimSpace(vesselType)) {
	case "cargo":
		return search.VesselTypeCargo
	default:
		return search.VesselTypeUnknown
	}
}

func normalizeAircraftType(aircraftType string) search.AircraftType {
	switch strings.ToLower(strings.TrimSpace(aircraftType)) {
	case "cargo":
		return search.AircraftCargo
	default:
		return search.AircraftTypeUnknown
	}
}

func parseTonnage(value string) int {
	// Remove commas and convert to int
	value = strings.ReplaceAll(value, ",", "")
	tonnage, _ := strconv.Atoi(value)
	return tonnage
}

func parseSerialNumber(remarks []string) string {
	for _, r := range findMatchingRemarks(remarks, "Serial Number") {
		// Remove parenthetical content and clean
		idx := strings.Index(r.value, ")")
		if idx > -1 && len(r.value) > idx+1 {
			return strings.TrimSpace(r.value[idx+1:])
		}
		return strings.TrimSpace(r.value)
	}
	return ""
}

var (
	// Regular expressions for parsing relationships and sanctions
	linkedToRegex   = regexp.MustCompile(`(?i)Linked\s+To:\s+([^;]+)`)
	subsidiaryRegex = regexp.MustCompile(`(?i)Subsidiary\s+Of:\s+([^;]+)`)
	ownedByRegex    = regexp.MustCompile(`(?i)(?:Owned|Controlled)\s+By:\s+([^;]+)`)
	sanctionsRegex  = regexp.MustCompile(`(?i)Additional\s+Sanctions\s+Information\s+-\s+([^;]+)`)
	formerNameRegex = regexp.MustCompile(`(?i)(?:Former|Previous|f\.k\.a\.|p\.k\.a\.)\s+(?:Name|Vessel):\s+([^;]+)`)
	titleRegex      = regexp.MustCompile(`(?i)Title:\s+([^;]+)`)

	cryptoAddressRegex = regexp.MustCompile(`(?i)Digital\s+Currency\s+Address\s+-\s+([A-Z0-9]+)\s+([A-Z0-9]+)`) // Matches "Digital Currency Address - XBT 1234abc..."
)

func parseRemarks(remarks []string) ([]search.Affiliation, *search.SanctionsInfo, []search.HistoricalInfo, []string) {
	var affiliations []search.Affiliation
	var historicalInfo []search.HistoricalInfo
	var titles []string
	sanctionsInfo := &search.SanctionsInfo{}

	for _, remark := range remarks {
		// Parse affiliations
		if matches := linkedToRegex.FindAllStringSubmatch(remark, -1); matches != nil {
			for _, m := range matches {
				affiliations = append(affiliations, search.Affiliation{
					EntityName: strings.TrimSpace(m[1]),
					Type:       "Linked To",
				})
			}
		}

		// Parse subsidiary relationships
		if matches := subsidiaryRegex.FindAllStringSubmatch(remark, -1); matches != nil {
			for _, m := range matches {
				affiliations = append(affiliations, search.Affiliation{
					EntityName: strings.TrimSpace(m[1]),
					Type:       "Subsidiary Of",
				})
			}
		}

		// Parse owned/controlled by relationships
		if matches := ownedByRegex.FindAllStringSubmatch(remark, -1); matches != nil {
			for _, m := range matches {
				affiliations = append(affiliations, search.Affiliation{
					EntityName: strings.TrimSpace(m[1]),
					Type:       "Subsidiary Of",
				})
			}
		}

		// Parse sanctions information
		if matches := sanctionsRegex.FindStringSubmatch(remark); matches != nil {
			info := strings.TrimSpace(matches[1])
			sanctionsInfo.Description = info
			if strings.Contains(strings.ToLower(info), "secondary sanctions") {
				sanctionsInfo.Secondary = true
			}
		}

		// Parse historical information
		if matches := formerNameRegex.FindAllStringSubmatch(remark, -1); matches != nil {
			for _, m := range matches {
				historicalInfo = append(historicalInfo, search.HistoricalInfo{
					Type:  "Former Name",
					Value: strings.TrimSpace(m[1]),
				})
			}
		}

		// Parse titles
		if matches := titleRegex.FindAllStringSubmatch(remark, -1); matches != nil {
			for _, m := range matches {
				titles = append(titles, strings.TrimSpace(m[1]))
			}
		}
	}

	return deduplicateAffiliations(affiliations),
		sanitizeSanctionsInfo(sanctionsInfo),
		deduplicateHistoricalInfo(historicalInfo),
		deduplicateTitles(titles)
}

func deduplicateAffiliations(affiliations []search.Affiliation) []search.Affiliation {
	seen := make(map[string]bool)
	var result []search.Affiliation

	for _, aff := range affiliations {
		key := aff.Type + "|" + aff.EntityName
		if !seen[key] {
			seen[key] = true
			result = append(result, aff)
		}
	}
	return result
}

func sanitizeSanctionsInfo(info *search.SanctionsInfo) *search.SanctionsInfo {
	if info == nil || (info.Description == "" && !info.Secondary && len(info.Programs) == 0) {
		return nil
	}
	return info
}

func deduplicateHistoricalInfo(info []search.HistoricalInfo) []search.HistoricalInfo {
	seen := make(map[string]bool)
	var result []search.HistoricalInfo

	for _, hi := range info {
		key := hi.Type + "|" + hi.Value
		if !seen[key] {
			seen[key] = true
			result = append(result, hi)
		}
	}
	return result
}

func deduplicateTitles(titles []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, title := range titles {
		if !seen[title] {
			seen[title] = true
			result = append(result, title)
		}
	}
	return result
}

func parseContactInfo(remarks []string) search.ContactInfo {
	var out search.ContactInfo

	// Look for emails, phone numbers, fax numbers, websites, etc..
	for _, remark := range remarks {
		remark = strings.TrimSuffix(remark, ".")

		// Emails
		matches := emailRegex.FindAllStringSubmatch(remark, -1)
		for _, m := range matches {
			if len(m) > 1 {
				out.EmailAddresses = append(out.EmailAddresses, m[len(m)-1])
			}
		}

		// Phone Numbers
		matches = phoneRegex.FindAllStringSubmatch(remark, -1)
		for _, m := range matches {
			if len(m) > 1 {
				out.PhoneNumbers = append(out.PhoneNumbers, m[len(m)-1])
			}
		}

		// Fax Numbers
		matches = faxRegex.FindAllStringSubmatch(remark, -1)
		for _, m := range matches {
			if len(m) > 1 {
				out.FaxNumbers = append(out.FaxNumbers, m[len(m)-1])
			}
		}

		// Websites
		matches = websiteRegex.FindAllStringSubmatch(remark, -1)
		for _, m := range matches {
			if len(m) > 1 {
				website := m[len(m)-1]
				if strings.Contains(website, "//www.treasury.gov") {
					continue
				}
				out.Websites = append(out.Websites, website)
			}
		}
	}

	return out
}

func parseCryptoAddresses(remarks []string) []search.CryptoAddress {
	var addresses []search.CryptoAddress

	for _, remark := range remarks {
		matches := cryptoAddressRegex.FindAllStringSubmatch(remark, -1)
		for _, m := range matches {
			if len(m) > 2 {
				addresses = append(addresses, search.CryptoAddress{
					Currency: strings.TrimSpace(m[1]),
					Address:  strings.TrimSpace(m[2]),
				})
			}
		}
	}

	// Deduplicate addresses
	seen := make(map[string]bool)
	var unique []search.CryptoAddress

	for _, addr := range addresses {
		key := addr.Currency + "|" + addr.Address
		if !seen[key] {
			seen[key] = true
			unique = append(unique, addr)
		}
	}

	return unique
}

var (
	oldestAllowedDate = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
)

func invalidDate(when time.Time) bool {
	return when.Before(oldestAllowedDate)
}

// ContactInfo
// Fax: 0097282858208.
// Fax No. (022) 7363196.
// Fax (356)(25990640)
// FAX 850 2 381 4431/4432
// Alt. Fax: 9221227700019
//
// Alt. Telephone: 982188554432
// Telephone (356)(21241232)
// Telephone + 97165749996
// Telephone +31 010-4951863
// Telephone No. (022) 7363030
// Telephone Number: (971) (4) (3248000).
// PHONE 850 2 18111 8204/8208
// PHONE 850 2 18111 ext. 8221
// Phone No. 263-4-486946
// Phone Number 982188526300
//
// EMAIL daesong@co.chesin.com.
// Email Address EnExchanger@gmail.com
// Email:adelb@shabakah.net.sa.
// info@sanabel.org.uk (email).
//
// Website Oboronlogistika.ru
// Website http://comitet.su/about/
// http://www.saraproperties.co.uk (website).
// a.k.a. 'ABU AHMAD ISHAB'.
// a.k.a. 'ZAMANI, Aziz Shah'
// GovernmentIDs
//
// Cedula No. 94428531 (Colombia)
// Cedula No. 94487319 (Colombia) issued 31 Oct 1994
// Birth Certificate Number 32270 (Iran)
// Bosnian Personal ID No. 1005967953038
// British National Overseas Passport 750200421 (United Kingdom)
// C.I.F. B84758374 (Spain).
// C.R. No. 03-B-1620
// C.R. No. J10/623/1997 (Romania)
// C.U.I.P. AOIR671020H1374898 (Mexico).
// C.U.I.T. 20-60357110-0 (Argentina)
// C.U.R.P. # HESU430525HBCRMR13 (Mexico)
// CNP (Personal Numerical Code) 7460301380011 (Romania)
// Cartilla de Servicio Militar Nacional 607092 (Mexico).
// Citizen's Card Number 210222198011096648 (China).
// Commercial Registry Number 0411518776478 (Iran)
// Commercial Registry Number CH-020.1.066.499-9 (Switzerland)
// Company ID: No. 59 531 at Commercial Registry of the Civil Court of First Instance at Baabda, Lebanon.
// Company Number 05527424 (United Kingdom)
// Company Number IMO 1991835.
// Credencial electoral 073855815496 (Mexico).
// D-U-N-S Number 33-843-5672
// D.N.I. 00263695-T (Spain)
// Diplomatic Passport 00000017 (Yemen) issued 27 Oct 2008 expires 26 Oct 2014
// Driver's License No. 04900377 (Moldova) issued 02 Jul 2004
// Driver's License No. 07442833 (United States) expires 15 Mar 2016
// Driver's License No. 1-1-22-07-00030905-3 (Guatemala) expires 2010.
// Driver's License No. M600161650080 (United States) issued 07 Apr 2006 expires 08 Jan 2011.
// Driver's License is issued by the State of Texas.
// Dubai Chamber of Commerce Membership No. 123076 (United Arab Emirates).
// Electoral Registry No. 07385114 (Afghanistan).
// Enterprise Number 0430.033.662 (Belgium).
// Fiscal Code 9896460 (Romania).
// Folio Mercantil No. 10328 (Jalisco) (Mexico).
// Government Gazette Number 00132598 (Russia).
// I.F.E. 05116040222575 (Mexico).
// Identification Number 0-16 Reg 53089 (Guatemala)
// Immigration No. A38839964 (United States).
// Interpol: Red Notice. File No. 2009/3599. March 24, 2009. Orange Notice. File No. 2009/52/OS/CCC. February 10, 2009.
// Italian Fiscal Code BCHMHT69R13Z352T.
// Kenyan ID No. 12773667
// LE Number 07541863 (Peru).
// Legal Entity Number 851683897 (Netherlands)
// License 1249 (Russia).
// Matricula Mercantil No 0000104026 (Colombia).
// N.I.E. X-1552120-B (Spain).
// NIT # 16215230-1 (Colombia).
// National Foreign ID Number 210602197107153012 (China)
// National ID No. (HWI)040182 (Burma)
// Passport #H0044232 (Iraq).
// Passport 00016161 (Yemen) issued 19 Jun 2012 expires 18 Jun 2018
// Passport No.: 0310857, Eritrea, Issue Date 21 August 2006, Expire Date 20 August 2008)
// Personal ID Card 00246412491303975500493 (Slovenia) expires 17 Dec 2018
// Pilot License Number 2326384
// Public Registration Number 1021801434380.
// Public Registration Number 1041202 (Virgin Islands, British)
// R.F.C. # IES-870805 (Mexico).
// RFC AAIJ810808SX4 (Mexico)
// RIF # J-00317392-4 (Venezuela).
// RTN 01019995013319 (Honduras)
// RUC # 1008619-1-537654 (Panama).
// Refugee ID Card A88000043 (Moldova) issued 16 Dec 2005.
// Registration ID 0000421465 (Poland)
// Registration Number 1027700499903 (Russia)
// Residency Number 003-5506420-0100028 (Costa Rica).
// Romanian C.R. J23/242/2004 (Romania).
// Romanian Permanent Resident CAN 0125477 (Romania) issued 13 Jul 2007.
// Romanian Tax Registration 14637977 (Romania).
// SSN 156-92-9858 (United States)
// SSN 33-3208848-3 (Philippines)
// SWIFT/BIC AFABAFKA
// Stateless Person ID Card CC00200261 (Moldova) issued 09 Sep 2000
// Stateless Person Passport C000375 (Moldova) issued 09 Sep 2000
// Tax ID No. 002235933 (Canada).
// Trade License No. 04110179 (United Kingdom).
// Travel Document Number A0003900 (Germany)
// Turkish Identificiation Number 10298480866 (Turkey).
// U.S.A. Passport issued 21 Jun 1992 in Amman, Jordan.
// UAE Identification 784-1968-9720837-5
// UK Company Number 01019769 (United Kingdom)
// US FEIN 000920912 (United States).
// United Social Credit Code Certificate (USCCC) 91420112711981060J (China)
// V.A.T. Number 0430.033.662 (Belgium)
// VisaNumberID 2024702 (Mexico).
