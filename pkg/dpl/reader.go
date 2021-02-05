// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package dpl

import (
	"encoding/csv"
	"os"
)

// Reader parses DPL records from a TXT file and populates the associated arrays.
//
// For more details on the raw DPL files see https://moov-io.github.io/watchman/file-structure.html
func Read(path string) ([]*DPL, error) {
	// open txt file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create a new csv.Reader and set the delim char to txtDelim char
	reader := csv.NewReader(f)
	reader.Comma = '\t'

	// Read File into a Variable
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var out []*DPL
	for _, txtLine := range lines {
		if txtLine[1] == "Street_Address" {
			continue // skip the headers
		}

		deniedPerson := &DPL{
			Name:           txtLine[0],
			StreetAddress:  txtLine[1],
			City:           txtLine[2],
			State:          txtLine[3],
			Country:        txtLine[4],
			PostalCode:     txtLine[5],
			EffectiveDate:  txtLine[6],
			ExpirationDate: txtLine[7],
			StandardOrder:  txtLine[8],
			LastUpdate:     txtLine[9],
			Action:         txtLine[10],
			FRCitation:     txtLine[11],
		}
		out = append(out, deniedPerson)
	}
	return out, nil
}
