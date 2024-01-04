package ofac

import (
	"strings"

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

	switch strings.ToLower(strings.TrimSpace(sdn.SDNType)) {
	case "-0-", "":
		out.Type = search.EntityBusiness // TODO(adam): or EntityOrganization
		// TODO(adam): How to tell Business vs Organization ?

	case "individual":
		out.Type = search.EntityPerson
		out.Person = &search.Person{
			Name: sdn.SDNName,
		}

		// TODO(adam):
		// DOB 02 Aug 1991;
		// nationality Russia;
		// Gender Male;
		// Passport 0291622 (Belize);

	case "vessel":
		out.Type = search.EntityVessel
		out.Vessel = &search.Vessel{
			Name: sdn.SDNName,

			// IMONumber string     `json:"imoNumber"`
			// Type      VesselType `json:"type"`
			// Flag      string     `json:"flag"` // ISO-3166
			// Built     *time.Time `json:"built"`
			// Model     string     `json:"model"`
			// Tonnage   int        `json:"tonnage"` // TODO(adam): remove , and ParseInt
			// MMSI      string     `json:"mmsi"` // Maritime Mobile Service Identity
		}

		// TODO(adam):
		// Vessel Registration Identification IMO 9569712;
		// MMSI 572469210;
		//
		// Former Vessel Flag None Identified;        alt. Former Vessel Flag Tanzania;

	case "aircraft":
		out.Type = search.EntityAircraft
		out.Aircraft = &search.Aircraft{
			Name: sdn.SDNName,

			// Type         AircraftType `json:"type"`
			// Flag         string       `json:"flag"` // ISO-3166
			// Built        *time.Time   `json:"built"`
			// ICAOCode     string       `json:"icaoCode"` // ICAO aircraft type designator
			// Model        string       `json:"model"`
			// SerialNumber string       `json:"serialNumber"`
		}

		// TODO(adam):
		// Aircraft Construction Number (also called L/N or S/N or F/N) 8401;
		// Aircraft Manufacture Date 1992;
		// Aircraft Model IL76-TD;
		// Aircraft Operator YAS AIR;
		// Aircraft Manufacturer's Serial Number (MSN) 1023409321;
	}

	return out
}
