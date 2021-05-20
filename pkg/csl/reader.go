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
		if len(record) <= 1 {
			continue // skip empty records
		}

		// CSL datafiles have added a unique identifier as the first column. Thus
		// we need to check either column 0 or 1 contains the identifier.
		for i := 0; i <= 1; i++ {
			switch record[i] {
			case "Sectoral Sanctions Identifications List (SSI) - Treasury Department":
				ssis = append(ssis, unmarshalSSI(record, i))

			case "Entity List (EL) - Bureau of Industry and Security":
				els = append(els, unmarshalEL(record, i))
			}
		}
	}

	return &CSL{
		SSIs: ssis,
		ELs:  els,
	}, nil
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
