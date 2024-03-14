package ofac

import (
	"regexp"
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/search"
)

func PtrToEntity(sdn *SDN) search.Entity[SDN] {
	if sdn != nil {
		return ToEntity(*sdn)
	}
	return search.Entity[SDN]{}
}

// TODO(adam): Accept Addresses, Alts, Comments

func ToEntity(sdn SDN) search.Entity[SDN] {
	out := search.Entity[SDN]{
		Name:       sdn.SDNName,
		Source:     search.SourceUSOFAC,
		SourceData: sdn,
	}

	remarks := splitRemarks(sdn.Remarks)

	switch strings.ToLower(strings.TrimSpace(sdn.SDNType)) {
	case "-0-", "":
		out.Type = search.EntityBusiness
		// Set properties
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
		})

	case "individual":
		out.Type = search.EntityPerson
		out.Person = &search.Person{
			Name:   sdn.SDNName,
			Gender: search.Gender(strings.ToLower(firstValue(findMatchingRemarks(remarks, "Gender")))),
		}
		out.Person.BirthDate = withFirstP(findMatchingRemarks(remarks, "DOB"), func(in remark) *time.Time {
			t, err := parseTime(dobPatterns, in.value)
			if t.IsZero() || err != nil {
				return nil
			}
			return &t
		})

		// TODO(adam):
		// citizen Venezuela
		//
		// nationality Russia;
		// nationality: Eritrean
		//
		// POB 'Adlun, Lebanon
		// Alt. POB: Keren Eritrea
		// POB Abadan, Iran

	case "vessel":
		out.Type = search.EntityVessel
		out.Vessel = &search.Vessel{
			Name:      sdn.SDNName,
			IMONumber: firstValue(findMatchingRemarks(remarks, "IMO")),
			Type: withFirstF(findMatchingRemarks(remarks, "Vessel Type"), func(r remark) search.VesselType {
				return search.VesselType(r.value) // TODO(adam): OFAC values are not an enum
			}),
			Flag: firstValue(findMatchingRemarks(remarks, "Flag")), // TODO(adam): ISO-3166
			// Built     *time.Time `json:"built"`
			// Model     string     `json:"model"`
			// Tonnage   int        `json:"tonnage"` // TODO(adam): remove , and ParseInt
			MMSI: firstValue(findMatchingRemarks(remarks, "MMSI")),
		}

	case "aircraft":
		out.Type = search.EntityAircraft
		out.Aircraft = &search.Aircraft{
			Name: sdn.SDNName,
			// Type         AircraftType `json:"type"`
			Flag: firstValue(findMatchingRemarks(remarks, "Flag")), // TODO(adam): ISO-3166
			Built: withFirstP(findMatchingRemarks(remarks, "Manufacture Date"), func(in remark) *time.Time {
				t, err := parseTime(dobPatterns, in.value)
				if t.IsZero() || err != nil {
					return nil
				}
				return &t
			}),
			// ICAOCode     string       `json:"icaoCode"` // ICAO aircraft type designator
			Model: firstValue(findMatchingRemarks(remarks, "Aircraft Model")),
			SerialNumber: withFirstF(findMatchingRemarks(remarks, "Serial Number"), func(r remark) string {
				// Trim parens from these remarks
				// e.g. "Aircraft Manufacturer's Serial Number (MSN) 1023409321;"
				idx := strings.Index(r.value, ")")
				if idx > -1 && len(r.value) > idx+1 {
					r.value = strings.TrimSpace(r.value[idx+1:])
				}
				return r.value
			}),
		}

		// TODO(adam):
		// Aircraft Operator YAS AIR;
		// Previous Aircraft Tail Number 2-WGLP
	}

	return out
}

var parenCountryRegex = regexp.MustCompile(`\(([\w\s]+)\)`)

func makeIdentifier(remarks []string, suffix string) *search.Identifier {
	found := findMatchingRemarks(remarks, suffix)
	if len(found) == 0 {
		return nil
	}

	// Often the country is in parenthesis at the end, so let's look for that
	//
	// Business Number 51566843 (Hong Kong)
	country := parenCountryRegex.FindString(found[0].value)
	country = strings.TrimPrefix(strings.TrimSuffix(country, ")"), "(")

	return &search.Identifier{
		Name:       found[0].fullName,
		Country:    country, // ISO-3166 // TODO(adam):
		Identifier: found[0].value,
	}
}

func makeIdentifiers(remarks []string, needles []string) []search.Identifier {
	var out []search.Identifier
	for i := range needles {
		if id := makeIdentifier(remarks, needles[i]); id != nil {
			out = append(out, *id)
		}
	}
	return out
}

var (
	dobPatterns = []string{
		"02 Jan 2006", // 01 Apr 1950
		"Jan 2006",    // Sep 1958
		"2006",        // 1928
	}
)

func parseTime(acceptedLayouts []string, value string) (time.Time, error) {
	// We don't currently support ranges for birth dates, so take the first date provided
	// Examples include:
	// 01 Feb 1958 to 28 Feb 1958
	// circa 1934
	// circa 1979-1982
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

// TODO(adam):
// Drop "alt. "

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
