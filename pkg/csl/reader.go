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

	_, err = reader.Read() // read and discard the header row
	if err != nil {
		return nil, err
	}

	var out *CSL
	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

	}
	return out, nil
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
