package csl

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

func ReadFile(path string) (*CSL, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	return Parse(fd)
}

func Parse(r io.Reader) (*CSL, error) {
	reader := csv.NewReader(r)

	var report CSL
	for {
		record, err := reader.Read()
		if err != nil {
			// reached the last line
			if errors.Is(err, io.EOF) {
				break
			}
			// malformed row
			if errors.Is(err, csv.ErrFieldCount) ||
				errors.Is(err, csv.ErrBareQuote) ||
				errors.Is(err, csv.ErrQuote) {
				continue
			}
			return nil, err
		}

		if len(record) <= 1 {
			continue // skip empty records
		}

		// CSL datafiles have added a unique identifier as the first column. Thus
		// we need to check either column 0 or 1 contains the identifier.
		for i := 0; i <= 1; i++ {
			switch record[i] {
			case "Entity List (EL) - Bureau of Industry and Security":
				report.ELs = append(report.ELs, unmarshalEL(record, i))

			case "Military End User (MEU) List - Bureau of Industry and Security":
				report.MEUs = append(report.MEUs, unmarshalMEU(record, i))

			case "Sectoral Sanctions Identifications List (SSI) - Treasury Department":
				report.SSIs = append(report.SSIs, unmarshalSSI(record, i))

			case "Unverified List (UVL) - Bureau of Industry and Security":
				report.UVLs = append(report.UVLs, unmarshalUVL(record, i))

			case "Nonproliferation Sanctions (ISN) - State Department":
				report.ISNs = append(report.ISNs, unmarshalISN(record, i))

			case "AECA Debarred List": // TODO: Not found
				// TODO(adam): https://github.com/moov-io/watchman/issues/414

			case "Foreign Sanctions Evaders (FSE) - Treasury Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/415
				report.FSEs = append(report.FSEs, unmarshalFSE(record, i))

			case "Palestinian Legislative Council List (PLC) - Treasury Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/416
				report.PLCs = append(report.PLCs, unmarshalPLC(record, i))

			case "Capta List (CAP) - Treasury Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/417
				report.CAPs = append(report.CAPs, unmarshalCAP(record, i))

			case "Non-SDN Menu-Based Sanctions List (NS-MBS List) - Treasury Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/418
				report.NS_MBSs = append(report.NS_MBSs, unmarshalNS_MBS(record, i))

			case "Non-SDN Chinese Military-Industrial Complex Companies List (CMIC) - Treasury Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/419
				report.CMICs = append(report.CMICs, unmarshalCMIC(record, i))

			case "ITAR Debarred (DTC) - State Department":
				// TODO(adam): https://github.com/moov-io/watchman/issues/422
				report.DTCs = append(report.DTCs, unmarshalDTC(record, i))

			default:
				// Other lists are:
				// "Specially Designated Nationals (SDN) - Treasury Department"
				// "Denied Persons List (DPL) - Bureau of Industry and Security"
			}
		}
	}
	return &report, nil
}

func unmarshalEL(row []string, offset int) *EL {
	id := ""
	if offset == 1 {
		id = row[0] // set the ID from the newer CSV format
	}
	return &EL{
		ID:                 id,
		Name:               row[NameIdx+offset],
		Addresses:          expandField(row[AddressesIdx+offset]),
		AlternateNames:     expandField(row[AltNamesIdx+offset]),
		StartDate:          row[StartDateIdx+offset],
		LicenseRequirement: row[LicenseRequirementIdx+offset],
		LicensePolicy:      row[LicensePolicyIdx+offset],
		FRNotice:           row[FRNoticeIdx+offset],
		SourceListURL:      row[SourceListURLIdx+offset],
		SourceInfoURL:      row[SourceInformationURLIdx+offset],
	}
}

func unmarshalMEU(record []string, offset int) *MEU {
	return &MEU{
		EntityID:  record[0],
		Name:      record[NameIdx+offset],
		Addresses: record[AddressesIdx+offset],
		FRNotice:  record[FRNoticeIdx+offset],
		StartDate: record[StartDateIdx+offset],
		EndDate:   record[EndDateIdx+offset],
	}
}

func unmarshalSSI(record []string, offset int) *SSI {
	return &SSI{
		EntityID:       record[EntityNumberIdx+offset],
		Type:           record[TypeIdx+offset],
		Programs:       expandProgramsList(record[ProgramsIdx+offset]),
		Name:           record[NameIdx+offset],
		Addresses:      expandField(record[AddressesIdx+offset]),
		Remarks:        expandField(record[RemarksIdx+offset]),
		AlternateNames: expandField(record[AltNamesIdx+offset]),
		IDsOnRecord:    expandField(record[IDsIdx+offset]),
		SourceListURL:  record[SourceListURLIdx+offset],
		SourceInfoURL:  record[SourceInformationURLIdx+offset],
	}
}

func unmarshalUVL(record []string, offset int) *UVL {
	return &UVL{
		EntityID:      record[0],
		Name:          record[NameIdx+offset],
		Addresses:     expandField(record[AddressesIdx+offset]),
		SourceListURL: record[SourceListURLIdx+offset],
		SourceInfoURL: record[SourceInformationURLIdx+offset],
	}
}

func unmarshalISN(record []string, offset int) *ISN {
	return &ISN{
		EntityID:              record[0],
		Programs:              expandProgramsList(record[ProgramsIdx+offset]),
		Name:                  record[NameIdx+offset],
		FederalRegisterNotice: record[FRNoticeIdx+offset],
		StartDate:             record[StartDateIdx+offset],
		Remarks:               expandField(record[RemarksIdx+offset]),
		SourceListURL:         record[SourceListURLIdx+offset],
		AlternateNames:        expandField(record[AltNamesIdx+offset]),
		SourceInfoURL:         record[SourceInformationURLIdx+offset],
	}
}

func unmarshalFSE(record []string, offset int) *FSE {
	return &FSE{
		EntityID:      record[0],
		EntityNumber:  record[EntityNumberIdx+offset],
		Type:          record[TypeIdx+offset],
		Programs:      expandProgramsList(record[ProgramsIdx+offset]),
		Name:          record[NameIdx+offset],
		Addresses:     expandField(record[AddressesIdx+offset]),
		SourceListURL: record[SourceListURLIdx+offset],
		Citizenships:  record[CitizenshipsIdx+offset],
		DatesOfBirth:  record[DatesOfBirthIdx+offset],
		SourceInfoURL: record[SourceInformationURLIdx+offset],
		IDs:           expandField(record[IDsIdx+offset]),
	}
}

func unmarshalPLC(record []string, offset int) *PLC {
	return &PLC{
		EntityID:       record[0],
		EntityNumber:   record[EntityNumberIdx+offset],
		Type:           record[TypeIdx+offset],
		Programs:       expandProgramsList(record[ProgramsIdx+offset]),
		Name:           record[NameIdx+offset],
		Addresses:      expandField(record[AddressesIdx+offset]),
		Remarks:        record[RemarksIdx+offset],
		SourceListURL:  record[SourceListURLIdx+offset],
		AlternateNames: expandField(record[AltNamesIdx+offset]),
		DatesOfBirth:   record[DatesOfBirthIdx+offset],
		PlacesOfBirth:  record[PlacesOfBirthIdx+offset],
		SourceInfoURL:  record[SourceInformationURLIdx+offset],
	}
}

func unmarshalCAP(record []string, offset int) *CAP {
	return &CAP{
		EntityID:       record[0],
		EntityNumber:   record[EntityNumberIdx+offset],
		Type:           record[TypeIdx+offset],
		Programs:       expandProgramsList(record[ProgramsIdx+offset]),
		Name:           record[NameIdx+offset],
		Addresses:      expandField(record[AddressesIdx+offset]),
		Remarks:        expandField(record[RemarksIdx+offset]),
		SourceListURL:  record[SourceListURLIdx+offset],
		AlternateNames: expandField(record[AltNamesIdx+offset]),
		SourceInfoURL:  record[SourceInformationURLIdx+offset],
		IDs:            expandField(record[IDsIdx+offset]),
	}
}

func unmarshalNS_MBS(record []string, offset int) *NS_MBS {
	return &NS_MBS{
		EntityID:       record[0],
		EntityNumber:   record[EntityNumberIdx+offset],
		Type:           record[TypeIdx+offset],
		Programs:       expandProgramsList(record[ProgramsIdx+offset]),
		Name:           record[NameIdx+offset],
		Addresses:      expandField(record[AddressesIdx+offset]),
		Remarks:        expandField(record[RemarksIdx+offset]),
		AlternateNames: expandField(record[AltNamesIdx+offset]),
		SourceInfoURL:  record[SourceInformationURLIdx+offset],
		IDs:            expandField(record[IDsIdx+offset]),
	}
}

func unmarshalCMIC(record []string, offset int) *CMIC {
	return &CMIC{
		EntityID:       record[0],
		EntityNumber:   record[EntityNumberIdx+offset],
		Type:           record[TypeIdx+offset],
		Programs:       expandProgramsList(record[ProgramsIdx+offset]),
		Name:           record[NameIdx+offset],
		Addresses:      expandField(record[AddressesIdx+offset]),
		Remarks:        expandField(record[RemarksIdx+offset]),
		SourceListURL:  record[SourceListURLIdx+offset],
		AlternateNames: expandField(record[AltNamesIdx+offset]),
		SourceInfoURL:  record[SourceInformationURLIdx+offset],
		IDs:            expandField(record[IDsIdx+offset]),
	}
}

func unmarshalDTC(record []string, offset int) *DTC {
	return &DTC{
		EntityID:              record[0],
		Name:                  record[NameIdx+offset],
		FederalRegisterNotice: record[FRNoticeIdx+offset],
		SourceListURL:         record[SourceListURLIdx+offset],
		AlternateNames:        expandField(record[AltNamesIdx+offset]),
		SourceInfoURL:         record[SourceInformationURLIdx+offset],
	}
}

// Some columns in a CSL row are actually lists delimited by ';'.
// These helper methods split these fields out and clean up the results.

func expandField(addrs string) []string {
	var result []string
	for _, a := range strings.Split(addrs, ";") {
		if res := strings.TrimSpace(a); res != "" {
			result = append(result, res)
		}
	}
	return result
}

var prgmReplacer = strings.NewReplacer("]", "", "[", "")

func expandProgramsList(prgms string) []string {
	prgms = strings.ReplaceAll(prgms, "] [", ";")
	prgms = prgmReplacer.Replace(prgms)
	return expandField(prgms)
}
