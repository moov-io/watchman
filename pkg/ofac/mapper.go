package ofac

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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
	emailRegex = regexp.MustCompile(`(?i)(?:Email|EMAIL)[:\s]+([^;]+)`)
	phoneRegex = regexp.MustCompile(`(?i)(?:Telephone|Phone|PHONE)[:\s]+([^;]+)`)
	faxRegex   = regexp.MustCompile(`(?i)Fax[:\s]+([^;]+)`)
	// Website patterns
	websiteRegex = regexp.MustCompile(`(?i)(?:Website|http)[:\s]+([^;\s]+)`)
	// Country extraction pattern
	countryParenRegex = regexp.MustCompile(`\(([\w\s]+)\)`)
)

var (
	dobPatterns = []string{
		"02 Jan 2006", // 01 Apr 1950
		"Jan 2006",    // Sep 1958
		"2006",        // 1928
	}
)

func makeIdentifiers(remarks []string, needles []string) []search.Identifier {
	seen := make(map[string]bool)
	var out []search.Identifier

	for i := range needles {
		if id := makeIdentifier(remarks, needles[i]); id != nil {
			// Create unique key from name and country
			key := id.Name + "|" + id.Country
			if !seen[key] {
				seen[key] = true
				out = append(out, *id)
			}
		}
	}
	return out
}

func makeIdentifier(remarks []string, suffix string) *search.Identifier {
	found := findMatchingRemarks(remarks, suffix)
	if len(found) == 0 {
		return nil
	}

	// Often the country is in parenthesis at the end, so let's look for that
	// Example: Business Number 51566843 (Hong Kong)
	country := ""
	value := found[0].value

	if matches := countryParenRegex.FindStringSubmatch(value); len(matches) > 1 {
		country = matches[1]
		// Remove the country part from the value
		value = strings.TrimSpace(countryParenRegex.ReplaceAllString(value, ""))
	}

	return &search.Identifier{
		Name:       strings.TrimSpace(found[0].fullName),
		Country:    country,
		Identifier: value,
	}
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
		if !tt.IsZero() && err == nil {
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

func ToEntity(sdn SDN, addresses []Address, comments []SDNComments) search.Entity[SDN] {
	out := search.Entity[SDN]{
		Name:       sdn.SDNName,
		Source:     search.SourceUSOFAC,
		SourceData: sdn,
	}

	remarks := splitRemarks(sdn.Remarks)
	affiliations, sanctionsInfo, historicalInfo, titles := parseRemarks(remarks)

	out.Affiliations = affiliations
	out.SanctionsInfo = sanctionsInfo
	out.HistoricalInfo = historicalInfo
	out.Titles = titles
	out.CryptoAddresses = parseCryptoAddresses(remarks)

	// Extract common fields regardless of entity type
	out.Addresses = parseAddresses(addresses)

	switch strings.ToLower(strings.TrimSpace(sdn.SDNType)) {
	case "-0-", "":
		out.Type = search.EntityBusiness
		out.Business = &search.Business{
			Name: sdn.SDNName,
		}
		out.Business.Identifier = makeIdentifiers(remarks, []string{
			"Branch Unit Number",
			"Business Number",
			"Business Registration Document",
			"Business Registration Number",
			"Certificate of Incorporation Number",
			"Chamber of Commerce Number",
			"Chinese Commercial Code",
			"Registered Charity No.",
			"Commercial Registry Number",
			"Company Number",
			"Enterprise Number",
			"Legal Entity Number",
			"Registration Number",
		})

	case "individual":
		out.Type = search.EntityPerson
		out.Person = &search.Person{
			Name:   sdn.SDNName,
			Gender: search.Gender(strings.ToLower(firstValue(findMatchingRemarks(remarks, "Gender")))),
		}

		// Title from SDN field needs to be prepended if non-empty
		if sdn.Title != "" {
			titles = append([]string{sdn.Title}, titles...)
		}
		out.Titles = titles

		// Extract alternative names
		out.Person.AltNames = parseAltNames(remarks)

		// Handle birth date
		out.Person.BirthDate = withFirstP(findMatchingRemarks(remarks, "DOB"), func(in remark) *time.Time {
			t, err := parseTime(dobPatterns, in.value)
			if t.IsZero() || err != nil {
				return nil
			}
			return &t
		})

		// Parse government IDs
		out.Person.GovernmentIDs = parseGovernmentIDs(remarks)

	case "vessel":
		out.Type = search.EntityVessel
		out.Vessel = &search.Vessel{
			Name:      sdn.SDNName,
			IMONumber: firstValue(findMatchingRemarks(remarks, "IMO")),
			Type: withFirstF(findMatchingRemarks(remarks, "Vessel Type"), func(r remark) search.VesselType {
				return normalizeVesselType(r.value)
			}),
			Flag:                   normalizeCountryCode(firstValue(findMatchingRemarks(remarks, "Flag"))),
			MMSI:                   firstValue(findMatchingRemarks(remarks, "MMSI")),
			Tonnage:                parseTonnage(firstValue(findMatchingRemarks(remarks, "Tonnage"))),
			CallSign:               sdn.CallSign,
			GrossRegisteredTonnage: parseTonnage(sdn.GrossRegisteredTonnage),
			Owner:                  sdn.VesselOwner,
		}

	case "aircraft":
		out.Type = search.EntityAircraft
		out.Aircraft = &search.Aircraft{
			Name: sdn.SDNName,
			Type: normalizeAircraftType(firstValue(findMatchingRemarks(remarks, "Aircraft Type"))),
			Flag: normalizeCountryCode(firstValue(findMatchingRemarks(remarks, "Flag"))),
			Built: withFirstP(findMatchingRemarks(remarks, "Manufacture Date"), func(in remark) *time.Time {
				t, err := parseTime(dobPatterns, in.value)
				if t.IsZero() || err != nil {
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
	out := make([]search.Address, len(inputs))
	for i := range inputs {
		addr := fmt.Sprintf("%s %s %s", inputs[i].Address, inputs[i].CityStateProvincePostalCode, inputs[i].Country)

		out[i] = address.ParseAddress(addr)
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
	governmentIDBusinessRegistrationRegex = regexp.MustCompile(`(?i)Business\s+Registration\s+(?:No\.|Number|Document)?\s*([A-Z0-9-]+)`)
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
