package csl

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

func Read(path string) (*CSL, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var ssis []*SSI
	var els []*EL
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		switch record[0] {
		case "Sectoral Sanctions Identifications List (SSI) - Treasury Department":
			ssis = append(ssis, unmarshalSSI(record))
		case "Entity List (EL) - Bureau of Industry and Security":
			els = append(els, unmarshalEL(record))
		default:
			continue
		}
	}

	return &CSL{
		SSIs: ssis,
		ELs:  els,
	}, nil
}

func unmarshalSSI(record []string) *SSI {
	return &SSI{
		EntityID:       record[EntityNumberIdx],
		Type:           record[TypeIdx],
		Programs:       expandProgramsList(record[ProgramsIdx]),
		Name:           record[NameIdx],
		Addresses:      expandField(record[AddressesIdx]),
		Remarks:        expandField(record[RemarksIdx]),
		AlternateNames: expandField(record[AltNamesIdx]),
		IDsOnRecord:    expandField(record[IDsIdx]),
		SourceListURL:  record[SourceListURLIdx],
		SourceInfoURL:  record[SourceInformationURLIdx],
	}
}

func unmarshalEL(row []string) *EL {
	return &EL{
		Name:               row[NameIdx],
		Addresses:          expandField(row[AddressesIdx]),
		AlternateNames:     expandField(row[AltNamesIdx]),
		StartDate:          row[StartDateIdx],
		LicenceRequirement: row[LicenseRequirementIdx],
		LicensePolicy:      row[LicensePolicyIdx],
		FRNotice:           row[FRNoticeIdx],
		SourceListURL:      row[SourceListURLIdx],
		SourceInfoURL:      row[SourceInformationURLIdx],
	}
}

// Some columns in a CSL row are actually lists delimited by ';'.
// These helper methods split these fields out and clean up the results.

func expandField(addrs string) []string {
	var result []string
	for _, a := range strings.Split(addrs, ";") {
		result = append(result, strings.TrimSpace(a))
	}
	return result
}

var prgmReplacer = strings.NewReplacer("]", "", "[", "")

func expandProgramsList(prgms string) []string {
	prgms = strings.ReplaceAll(prgms, "] [", ";")
	prgms = prgmReplacer.Replace(prgms)
	return expandField(prgms)
}
